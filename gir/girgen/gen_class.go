package girgen

import (
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .InterfaceName }}
	type {{ .InterfaceName }} interface {
		{{ $.Ng.PublicType (index .TypeTree 0) }}
		{{ range .Methods }}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail -}}
		{{ end }}
	}

	type {{ .StructName }} struct {
		{{ .StructEmbeds }}
	}

	// Wrap{{ .InterfaceName }} wraps a GObject to the right type. It is
	// primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .Wrap "obj" }}
	}

	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
	}

	{{ range .Constructors }}
	{{ if $.UseConstructor . }}
	// {{ $.Callable.Name }} constructs a class {{ $.InterfaceName }}.
	func {{ $.Callable.Name }}{{ $.Callable.Tail }} {{ $.Callable.Block }}
	{{ end }}
	{{ end }}

	{{ $recv := (FirstLetter $.StructName) }}

	{{ range .Methods }}
	{{ GoDoc .Doc 1 .Name }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

type classGenerator struct {
	gir.Class
	StructName    string
	StructEmbeds  string
	InterfaceName string

	TypeTree []*ResolvedType
	Methods  []callableGenerator

	Callable callableGenerator

	Ng *NamespaceGenerator
}

func (cg *classGenerator) Use(class gir.Class) bool {
	if class.Parent == "" {
		// TODO: check what happens if a class has no parent. It should have a
		// GObject parent, usually.
		return false
	}

	cg.TypeTree = cg.TypeTree[:0]

	cg.Class = class
	cg.InterfaceName = PascalToGo(class.Name)
	cg.StructName = UnexportPascal(cg.InterfaceName)

	// Loop to resolve the parent type, the parent type of that parent type, and
	// so on.
	parent := class.Parent
	for {
		parentType := cg.Ng.ResolveTypeName(parent)
		if parentType == nil {
			cg.Ng.logln(logWarn, "failed to resolve parent", parent, "for", class.Name)
			return false
		}

		cg.TypeTree = append(cg.TypeTree, parentType)

		if parentType.Parent == "" {
			break
		}

		// Use the parent class' parent type.
		parent = parentType.Parent
	}

	cg.Methods = callableGrow(cg.Methods, len(class.Methods))

	for _, method := range class.Methods {
		cbgen := newCallableGenerator(cg.Ng)
		if !cbgen.Use(method.CallableAttrs) {
			continue
		}

		cbgen.Parent = cg.InterfaceName
		cg.Methods = append(cg.Methods, cbgen)
	}

	// Use Go-idiomatic getter names, unless there's a duplicate.
	callableRenameGetters(cg.Methods)

	// Treat StructEmbeds specially: we can only embed our own implementation
	// types, since they're unexported, so we embed interface types if it's not
	// our package.
	if parent := cg.TypeTree[0]; parent.NeedsNamespace(cg.Ng.current) {
		cg.StructEmbeds = parent.PublicType(true)
	} else {
		cg.StructEmbeds = parent.ImplType(false)
	}

	return true
}

// bodgeClassCtor bodges the given constructor to return exactly the class type
// instead of any other. It returns the original ctor if the conditions don't
// match for bodging.
//
// We have to do this to work around some cases where widget constructors would
// return the widget class instead of the actual class.
func bodgeClassCtor(class gir.Class, ctor gir.Constructor) gir.Constructor {
	if ctor.ReturnValue == nil || ctor.ReturnValue.Type == nil {
		return ctor
	}

	retVal := *ctor.ReturnValue
	retTyp := *retVal.AnyType.Type

	retTyp.Name = class.Name
	retTyp.CType = class.CType
	retTyp.Introspectable = class.Introspectable
	retTyp.AnyType = gir.AnyType{}

	retVal.AnyType.Type = &retTyp
	ctor.ReturnValue = &retVal

	ctor.Name = strings.TrimPrefix(ctor.Name, "new")
	ctor.Name = strings.TrimPrefix(ctor.Name, "_")
	if ctor.Name != "" {
		ctor.Name = "_" + ctor.Name
	}

	ctor.Name = "new_" + class.Name + ctor.Name

	return ctor
}

func (cg *classGenerator) UseConstructor(ctor gir.Constructor) bool {
	ctor = bodgeClassCtor(cg.Class, ctor)
	ok := cg.Callable.Use(ctor.CallableAttrs)
	cg.Callable.Parent = cg.InterfaceName
	return ok
}

// Wrap returns the wrap string around the given variable name of type
// *glib.Object.
func (cg *classGenerator) Wrap(objName string) string {
	p := pen.NewPiece()
	p.Write(cg.StructName).Char('{')

	// stack of characters to append afterwards
	stack := make([]byte, 1, 25)
	stack[0] = '}'

	for _, typ := range cg.TypeTree {
		// Extern type is not in the same package, so we can't reference the
		// exported type. Use the Wrap function instead.
		if typ.NeedsNamespace(cg.Ng.current) && typ.Extern != nil {
			p.Writef("%s.Wrap%s(", typ.Package, typ.PublicType(false))
			stack = append(stack, ')')

			break
		}

		p.Writef("%s{", cg.Ng.ImplType(typ))
		stack = append(stack, '}')
	}

	p.Write(objName)

	for i := len(stack) - 1; i >= 0; i-- {
		p.Char(stack[i])
	}

	return p.String()
}

func (ng *NamespaceGenerator) generateClasses() {
	cg := classGenerator{
		TypeTree: make([]*ResolvedType, 15),
		Callable: callableGenerator{Ng: ng},
		Ng:       ng,
	}

	for _, class := range ng.current.Namespace.Classes {
		if !cg.Use(class) {
			ng.logln(logInfo, "skipping class", class.Name)
			continue
		}

		ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
		ng.pen.BlockTmpl(classTmpl, &cg)
	}
}
