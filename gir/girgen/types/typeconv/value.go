package typeconv

import (
	"fmt"
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// ConversionDirection is the conversion direction between Go and C.
type ConversionDirection uint8

const (
	_ ConversionDirection = iota
	ConvertGoToC
	ConvertCToGo
)

// ConversionValueIndex describes an overloaded index type that reserves its
// negative values for special values.
type ConversionValueIndex int8

const (
	_ ConversionValueIndex = -iota // 0
	UnknownValueIndex
	ErrorValueIndex
	ReturnValueIndex
)

// Index returns the actual underlying index if any, or it returns -1.
func (ix ConversionValueIndex) Index() int {
	if ix > UnknownValueIndex {
		return int(ix)
	}
	return -1
}

// Is checks that the index matches. This method should be used as it guarantees
// that the given index isn't special.
func (ix ConversionValueIndex) Is(at int) bool {
	if at < 0 {
		log.Panicln("given index", at, "is invalid")
	}
	return ix.Index() == at
}

// ConversionValue describes the generic properties of a Go or C value for
// conversion.
type ConversionValue struct {
	gir.ParameterAttrs

	InName  string
	OutName string

	// Direction is the direction of conversion.
	Direction ConversionDirection

	// ParameterIndex explicitly gives this value an index used for matching
	// with the given index clues from the GIR files, such as closure, destroy
	// or length.
	ParameterIndex ConversionValueIndex

	// ManualCast, if true, will make GObject conversions from C to Go use the
	// return type given instead of using the value wrapper.
	ManualCast bool
}

// NewValue creates a new ConversionValue from the given parameter attributes.
func NewValue(
	in, out string, i int, dir ConversionDirection, param gir.ParameterAttrs) ConversionValue {

	value := ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      dir,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: param,
	}
	if i > -1 {
		value.ParameterIndex = ConversionValueIndex(i)
	}

	return value
}

// NewReturnValue creates a new ConversionValue from the given return attribute.
func NewReturnValue(in, out string, dir ConversionDirection, ret gir.ReturnValue) ConversionValue {
	return ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      dir,
		ParameterIndex: ReturnValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			Closure:  ret.Closure,
			Destroy:  ret.Destroy,
			Scope:    ret.Scope,
			Skip:     ret.Skip,
			Nullable: ret.Nullable,
			Optional: ret.Skip,

			TransferOwnership: ret.TransferOwnership,
			AnyType:           ret.AnyType,
			Doc:               ret.Doc,
		},
	}
}

// NewFieldValue creates a new ConversionValue from the given C struct field.
// The struct is assumed to have a native field.
func NewFieldValue(recv, out string, field gir.Field) ConversionValue {
	return ConversionValue{
		InName:         fmt.Sprintf("%s.native.%s", recv, strcases.CGoField(field.Name)),
		OutName:        out,
		Direction:      ConvertCToGo,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			Name:    field.Name,
			Skip:    field.Private || !field.IsReadable() || field.Bits > 0,
			AnyType: field.AnyType,
			Doc:     field.Doc,
		},
	}
}

// NewThrowValue creates a new GError value. Thrown values are always assumed
// to be conversions from C to Go. Errors should ALWAYS go AFTER the return!
func NewThrowValue(in, out string) ConversionValue {
	return ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      ConvertCToGo,
		ParameterIndex: ErrorValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			TransferOwnership: gir.TransferOwnership{
				TransferOwnership: "full",
			},
			AnyType: gir.AnyType{
				Type: &gir.Type{
					Name: "GLib.Error",
					// Function parameter type is technically a double-pointer
					// here.
					CType: "GError**",
				},
			},
			Optional:  true,
			Nullable:  true,
			Direction: "out",
		},
	}
}

// IsZero returns true if ConversionValue is empty.
func (value *ConversionValue) IsZero() bool {
	return value.InName == "" || value.OutName == ""
}

// ParameterIsOutput returns true if the direction is out.
func (value *ConversionValue) ParameterIsOutput() bool {
	switch value.ParameterAttrs.Direction {
	case "out":
		return true
	case "in":
		return false
	}

	// GIR is fucking miserable. Why can't they get these properly?
	if value.Doc != nil && strings.HasPrefix(value.Doc.String, "Return location") {
		return true
	}

	return false
}

// outputAllocs returns true if the parameter is a value we need to allocate
// ourselves.
func (value *ConversionValue) outputAllocs() bool {
	return value.ParameterIsOutput() && (value.CallerAllocates || value.isTransferring())
}

// isTransferring is true when the ownership is either full or container. If the
// converter code isn't generating for an array, then distinguishing this
// doesn't matter. If the caller hasn't set the ownership yet, then it is
// assumed that we're not getting the ownership, therefore false is returned.
//
// If the generating code is an array, and the conversion is being passed into
// the same generation routine for the inner type, then the ownership should be
// turned into "none" just for that inner type. See TypeConversion.inner().
func (prop *ConversionValue) isTransferring() bool {
	return false ||
		prop.TransferOwnership.TransferOwnership == "full" ||
		prop.TransferOwnership.TransferOwnership == "container"
}

// ValueConverted is the result of conversion for a single value.
//
// Quick convention note:
//
//    - {In,Out}Name is for the original name with no modifications.
//    - {In,Out}Call is used for the C or Go function arguments.
//
// Usually, these are the same, but they're sometimes different depending on the
// edge case.
type ValueConverted struct {
	ConversionValue // original

	InCall    string // use for calls
	InType    string
	InDeclare string

	OutCall    string // use for calls
	OutType    string
	OutDeclare string

	Conversion string

	// internal states
	Resolved       *types.Resolved // only for type conversions
	NeedsNamespace bool

	// output writers
	p       *pen.PaperString
	inDecl  *pen.PaperString
	outDecl *pen.PaperString
	header  file.Header
}

func newValueConverted(value *ConversionValue) ValueConverted {
	return ValueConverted{
		ConversionValue: *value,
		InCall:          value.InName,
		OutCall:         value.OutName,

		p:       pen.NewPaperStringSize(1024), // 1KB
		inDecl:  pen.NewPaperStringSize(128),  // 0.1KB
		outDecl: pen.NewPaperStringSize(128),  // 0.1KB
	}
}

// Header returns the header of the current value.
func (value *ValueConverted) Header() *file.Header { return &value.header }

func (value *ValueConverted) finalize() {
	value.InDeclare = value.inDecl.String()
	value.OutDeclare = value.outDecl.String()
	value.Conversion = value.p.String()

	// // Allow GC to collect the internal buffers.
	// value.inDecl = nil
	// value.outDecl = nil
	// value.p = nil
}

func (value *ValueConverted) logln(conv *Converter, lvl logger.Level, v ...interface{}) {
	conv.Logln(lvl, logger.Prefix(v, value.logPrefix())...)
}

// resolveType resolves the value type to the resolved field. If inputC is true,
// then the input type is set to the CGo type, otherwise the Go type is set.
func (value *ValueConverted) resolveType(conv *Converter) bool {
	if value.AnyType.Type == nil {
		return false
	}

	// ResolveType already checks this, but we can early bail.
	if !value.AnyType.Type.IsIntrospectable() {
		return false
	}

	if value.Resolved != nil {
		// already resolved
		return true
	}

	// Copy Type for mutation.
	typ := *value.AnyType.Type

	// Proritize hard-coded types over ignored types.
	value.Resolved = types.Resolve(types.OverrideNamespace(conv.fgen, conv.sourceNamespace), typ)
	if value.Resolved == nil {
		conv.Logln(logger.Debug, "can't resolve", types.AnyTypeCGo(value.AnyType), typ.Name)
		return false
	}

	// Set the type back for use. We're setting the AnyType struct, which is a
	// copy, so it's fine.
	value.AnyType.Type = &typ

	if value.Resolved.IsCallback() {
		value.header.AddCallback(value.Resolved.Extern.Type.(*gir.Callback))
	}

	// If this is the output parameter, then the pointer count should be less.
	// This only affects the Go 1	ype.
	if value.ParameterIsOutput() && value.Resolved.Ptr > 0 {
		value.Resolved.Ptr--
	}

	value.NeedsNamespace = value.Resolved.NeedsNamespace(conv.currentNamespace)

	cgoType := value.Resolved.CGoType()
	goType := value.Resolved.PublicType(value.NeedsNamespace)

	if value.Direction == ConvertCToGo {
		value.InType = cgoType
		value.OutType = goType
	} else {
		value.OutType = cgoType
		value.InType = goType
	}

	if value.Direction == ConvertCToGo && value.ParameterIsOutput() {
		value.InCall = "&" + value.InCall
		value.InType = strings.TrimPrefix(value.InType, "*")
	}

	value.inDecl.Linef("var %s %s // in", value.InName, value.InType)
	value.outDecl.Linef("var %s %s // out", value.OutName, value.OutType)

	return true
}

// cgoSetObject generates a glib.Take or glib.AssumeOwnership into a new
// function. This should only be used for C to Go conversion.
func (value *ValueConverted) cgoSetObject(conv *Converter) {
	var gobjectFunction string
	if value.isTransferring() {
		// Full or container means we implicitly own the object, so we must
		// not take another reference.
		gobjectFunction = "AssumeOwnership"
	} else {
		// Else the object is either unowned by us or it's a floating
		// reference. Take our own or sink the object.
		gobjectFunction = "Take"
	}

	value.header.NeedsExternGLib()
	value.header.Import("unsafe")

	if value.ManualCast {
		if !value.Resolved.NeedsNamespace(conv.currentNamespace) {
			value.p.Linef(
				"%s = wrap%s(externglib.%s(unsafe.Pointer(%s)))",
				value.OutName, value.Resolved.PublicType(false), gobjectFunction, value.InName,
			)
			return
		}

		if tree := types.NewTree(conv.fgen); tree.ResolveFromType(value.Resolved) {
			value.p.Descend()
			value.p.Linef("obj := externglib.%s(unsafe.Pointer(%s))", gobjectFunction, value.InName)
			value.p.Linef("%s = %s", value.OutName, tree.Wrap("obj"))
			value.p.Ascend()
			return
		}

		// Fallback.
	}

	value.header.ImportCore("gextras")
	value.p.Linef(
		"%s = gextras.CastObject(externglib.%s(unsafe.Pointer(%s))).(%s)",
		value.OutName, gobjectFunction, value.InName, value.OutType,
	)
}

func (value *ValueConverted) cmalloc(lenOf string, add1 bool) string {
	lenOf = "len(" + lenOf + ")"
	if add1 {
		lenOf += "+1"
	}

	return fmt.Sprintf("C.malloc(C.ulong(%s) * C.ulong(%s))", lenOf, value.csizeof())
}

func (value *ValueConverted) csizeof() string {
	// Arrays are lists of pointers.
	if strings.Contains(types.AnyTypeC(value.AnyType), "*") {
		return value.ptrsz()
	}

	if value.Resolved == nil {
		// Erroneous case.
		return value.ptrsz()
	}

	return "C.sizeof_" + value.Resolved.CType
}

func (value *ValueConverted) ptrsz() string {
	value.header.Import("unsafe")
	// Size of a pointer is the same as uint.
	return "unsafe.Sizeof(uint(0))"
}

func (value *ValueConverted) logPrefix() string {
	var prefix string
	switch value.Direction {
	case ConvertCToGo:
		prefix = fmt.Sprintf("C %s -> Go %s", value.InName, value.OutName)
	case ConvertGoToC:
		prefix = fmt.Sprintf("Go %s -> C %s", value.InName, value.OutName)
	default:
		return ""
	}

	if value.Resolved != nil {
		prefix += fmt.Sprintf(" (%s)", value.Resolved.PublicType(true))
	}

	return prefix
}

// isPtr checks pointer coherency for C types and Go types. It's mostly used to
// guarantee that conversion routines get what they expect.
func (value *ValueConverted) isPtr(wantC int) bool {
	// See this same piece of code in convertRef for more information.
	if types.IsGpointer(value.Resolved.CType) && wantC > 0 {
		wantC--
	}

	switch value.Direction {
	case ConvertCToGo:
		return strings.Count(value.InType, "*") == wantC
	case ConvertGoToC:
		return strings.Count(value.OutType, "*") == wantC
	default:
		return false
	}

	// Rationale for not verifying Go pointer offset is that the pointer offset
	// is already determined in the type resolver routine, so repeating that
	// information is redundant.
	//
	// Edit: this rationale does NOT work because the type resolver only has the
	// wanted Go pointer information up to the point of creating a new
	// ResolvedType, and there's no way we can get it back. This routine may not
	// need to verify the Go pointer, but the conversiron routine will.
}
