// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
// extern PangoFont* _gotk4_pango1_FontMapClass_load_font(PangoFontMap*, PangoContext*, PangoFontDescription*);
// extern PangoFontFamily* _gotk4_pango1_FontMapClass_get_family(PangoFontMap*, char*);
// extern PangoFontset* _gotk4_pango1_FontMapClass_load_fontset(PangoFontMap*, PangoContext*, PangoFontDescription*, PangoLanguage*);
// extern guint _gotk4_pango1_FontMapClass_get_serial(PangoFontMap*);
// extern void _gotk4_pango1_FontMapClass_changed(PangoFontMap*);
// extern void _gotk4_pango1_FontMapClass_list_families(PangoFontMap*, PangoFontFamily***, int*);
import "C"

// glib.Type values for pango-fontmap.go.
var GTypeFontMap = coreglib.Type(C.pango_font_map_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFontMap, F: marshalFontMap},
	})
}

// FontMapOverrider contains methods that are overridable.
type FontMapOverrider interface {
	// Changed forces a change in the context, which will cause any PangoContext
	// using this fontmap to change.
	//
	// This function is only useful when implementing a new backend for Pango,
	// something applications won't do. Backends should call this function if
	// they have attached extra data to the context and such data is changed.
	Changed()
	// Family gets a font family by name.
	//
	// The function takes the following parameters:
	//
	//    - name: family name.
	//
	// The function returns the following values:
	//
	//    - fontFamily: PangoFontFamily.
	//
	Family(name string) FontFamilier
	// Serial returns the current serial number of fontmap.
	//
	// The serial number is initialized to an small number larger than zero when
	// a new fontmap is created and is increased whenever the fontmap is
	// changed. It may wrap, but will never have the value 0. Since it can wrap,
	// never compare it with "less than", always use "not equals".
	//
	// The fontmap can only be changed using backend-specific API, like changing
	// fontmap resolution.
	//
	// This can be used to automatically detect changes to a PangoFontMap, like
	// in PangoContext.
	//
	// The function returns the following values:
	//
	//    - guint: current serial number of fontmap.
	//
	Serial() uint32
	// ListFamilies: list all families for a fontmap.
	//
	// The function returns the following values:
	//
	//    - families: location to store a pointer to an array of PangoFontFamily
	//      *. This array should be freed with g_free().
	//
	ListFamilies() []FontFamilier
	// LoadFont: load the font in the fontmap that is the closest match for
	// desc.
	//
	// The function takes the following parameters:
	//
	//    - context: PangoContext the font will be used with.
	//    - desc: PangoFontDescription describing the font to load.
	//
	// The function returns the following values:
	//
	//    - font (optional): newly allocated PangoFont loaded, or NULL if no font
	//      matched.
	//
	LoadFont(context *Context, desc *FontDescription) Fonter
	// LoadFontset: load a set of fonts in the fontmap that can be used to
	// render a font matching desc.
	//
	// The function takes the following parameters:
	//
	//    - context: PangoContext the font will be used with.
	//    - desc: PangoFontDescription describing the font to load.
	//    - language: PangoLanguage the fonts will be used for.
	//
	// The function returns the following values:
	//
	//    - fontset (optional): newly allocated PangoFontset loaded, or NULL if
	//      no font matched.
	//
	LoadFontset(context *Context, desc *FontDescription, language *Language) Fontsetter
}

// FontMap: PangoFontMap represents the set of fonts available for a particular
// rendering system.
//
// This is a virtual object with implementations being specific to particular
// rendering systems.
type FontMap struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FontMap)(nil)
)

// FontMapper describes types inherited from class FontMap.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FontMapper interface {
	coreglib.Objector
	baseFontMap() *FontMap
}

var _ FontMapper = (*FontMap)(nil)

func classInitFontMapper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.PangoFontMapClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.PangoFontMapClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Changed() }); ok {
		pclass.changed = (*[0]byte)(C._gotk4_pango1_FontMapClass_changed)
	}

	if _, ok := goval.(interface {
		Family(name string) FontFamilier
	}); ok {
		pclass.get_family = (*[0]byte)(C._gotk4_pango1_FontMapClass_get_family)
	}

	if _, ok := goval.(interface{ Serial() uint32 }); ok {
		pclass.get_serial = (*[0]byte)(C._gotk4_pango1_FontMapClass_get_serial)
	}

	if _, ok := goval.(interface{ ListFamilies() []FontFamilier }); ok {
		pclass.list_families = (*[0]byte)(C._gotk4_pango1_FontMapClass_list_families)
	}

	if _, ok := goval.(interface {
		LoadFont(context *Context, desc *FontDescription) Fonter
	}); ok {
		pclass.load_font = (*[0]byte)(C._gotk4_pango1_FontMapClass_load_font)
	}

	if _, ok := goval.(interface {
		LoadFontset(context *Context, desc *FontDescription, language *Language) Fontsetter
	}); ok {
		pclass.load_fontset = (*[0]byte)(C._gotk4_pango1_FontMapClass_load_fontset)
	}
}

//export _gotk4_pango1_FontMapClass_changed
func _gotk4_pango1_FontMapClass_changed(arg0 *C.PangoFontMap) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

//export _gotk4_pango1_FontMapClass_get_family
func _gotk4_pango1_FontMapClass_get_family(arg0 *C.PangoFontMap, arg1 *C.char) (cret *C.PangoFontFamily) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Family(name string) FontFamilier
	})

	var _name string // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	fontFamily := iface.Family(_name)

	cret = (*C.PangoFontFamily)(unsafe.Pointer(coreglib.InternObject(fontFamily).Native()))

	return cret
}

//export _gotk4_pango1_FontMapClass_get_serial
func _gotk4_pango1_FontMapClass_get_serial(arg0 *C.PangoFontMap) (cret C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Serial() uint32 })

	guint := iface.Serial()

	cret = C.guint(guint)

	return cret
}

//export _gotk4_pango1_FontMapClass_list_families
func _gotk4_pango1_FontMapClass_list_families(arg0 *C.PangoFontMap, arg1 ***C.PangoFontFamily, arg2 *C.int) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ListFamilies() []FontFamilier })

	families := iface.ListFamilies()

	*arg2 = (C.int)(len(families))
	*arg1 = (**C.PangoFontFamily)(C.calloc(C.size_t(len(families)), C.size_t(unsafe.Sizeof(uint(0)))))
	{
		out := unsafe.Slice((**C.PangoFontFamily)(*arg1), len(families))
		for i := range families {
			out[i] = (*C.PangoFontFamily)(unsafe.Pointer(coreglib.InternObject(families[i]).Native()))
		}
	}
}

//export _gotk4_pango1_FontMapClass_load_font
func _gotk4_pango1_FontMapClass_load_font(arg0 *C.PangoFontMap, arg1 *C.PangoContext, arg2 *C.PangoFontDescription) (cret *C.PangoFont) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		LoadFont(context *Context, desc *FontDescription) Fonter
	})

	var _context *Context      // out
	var _desc *FontDescription // out

	_context = wrapContext(coreglib.Take(unsafe.Pointer(arg1)))
	_desc = (*FontDescription)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	font := iface.LoadFont(_context, _desc)

	if font != nil {
		cret = (*C.PangoFont)(unsafe.Pointer(coreglib.InternObject(font).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(font).Native()))
	}

	return cret
}

//export _gotk4_pango1_FontMapClass_load_fontset
func _gotk4_pango1_FontMapClass_load_fontset(arg0 *C.PangoFontMap, arg1 *C.PangoContext, arg2 *C.PangoFontDescription, arg3 *C.PangoLanguage) (cret *C.PangoFontset) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		LoadFontset(context *Context, desc *FontDescription, language *Language) Fontsetter
	})

	var _context *Context      // out
	var _desc *FontDescription // out
	var _language *Language    // out

	_context = wrapContext(coreglib.Take(unsafe.Pointer(arg1)))
	_desc = (*FontDescription)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_language = (*Language)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	fontset := iface.LoadFontset(_context, _desc, _language)

	if fontset != nil {
		cret = (*C.PangoFontset)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(fontset).Native()))
	}

	return cret
}

func wrapFontMap(obj *coreglib.Object) *FontMap {
	return &FontMap{
		Object: obj,
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	return wrapFontMap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (fontmap *FontMap) baseFontMap() *FontMap {
	return fontmap
}

// BaseFontMap returns the underlying base object.
func BaseFontMap(obj FontMapper) *FontMap {
	return obj.baseFontMap()
}

// Changed forces a change in the context, which will cause any PangoContext
// using this fontmap to change.
//
// This function is only useful when implementing a new backend for Pango,
// something applications won't do. Backends should call this function if they
// have attached extra data to the context and such data is changed.
func (fontmap *FontMap) Changed() {
	var _arg0 *C.PangoFontMap // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))

	C.pango_font_map_changed(_arg0)
	runtime.KeepAlive(fontmap)
}

// CreateContext creates a PangoContext connected to fontmap.
//
// This is equivalent to pango.Context.New followed by
// pango.Context.SetFontMap().
//
// If you are using Pango as part of a higher-level system, that system may have
// it's own way of create a PangoContext. For instance, the GTK toolkit has,
// among others, gtk_widget_get_pango_context(). Use those instead.
//
// The function returns the following values:
//
//    - context: newly allocated PangoContext, which should be freed with
//      g_object_unref().
//
func (fontmap *FontMap) CreateContext() *Context {
	var _arg0 *C.PangoFontMap // out
	var _cret *C.PangoContext // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))

	_cret = C.pango_font_map_create_context(_arg0)
	runtime.KeepAlive(fontmap)

	var _context *Context // out

	_context = wrapContext(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _context
}

// Family gets a font family by name.
//
// The function takes the following parameters:
//
//    - name: family name.
//
// The function returns the following values:
//
//    - fontFamily: PangoFontFamily.
//
func (fontmap *FontMap) Family(name string) FontFamilier {
	var _arg0 *C.PangoFontMap    // out
	var _arg1 *C.char            // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_font_map_get_family(_arg0, _arg1)
	runtime.KeepAlive(fontmap)
	runtime.KeepAlive(name)

	var _fontFamily FontFamilier // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontFamilier is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(FontFamilier)
			return ok
		})
		rv, ok := casted.(FontFamilier)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFamilier")
		}
		_fontFamily = rv
	}

	return _fontFamily
}

// Serial returns the current serial number of fontmap.
//
// The serial number is initialized to an small number larger than zero when a
// new fontmap is created and is increased whenever the fontmap is changed. It
// may wrap, but will never have the value 0. Since it can wrap, never compare
// it with "less than", always use "not equals".
//
// The fontmap can only be changed using backend-specific API, like changing
// fontmap resolution.
//
// This can be used to automatically detect changes to a PangoFontMap, like in
// PangoContext.
//
// The function returns the following values:
//
//    - guint: current serial number of fontmap.
//
func (fontmap *FontMap) Serial() uint32 {
	var _arg0 *C.PangoFontMap // out
	var _cret C.guint         // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))

	_cret = C.pango_font_map_get_serial(_arg0)
	runtime.KeepAlive(fontmap)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// ListFamilies: list all families for a fontmap.
//
// The function returns the following values:
//
//    - families: location to store a pointer to an array of PangoFontFamily *.
//      This array should be freed with g_free().
//
func (fontmap *FontMap) ListFamilies() []FontFamilier {
	var _arg0 *C.PangoFontMap     // out
	var _arg1 **C.PangoFontFamily // in
	var _arg2 C.int               // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))

	C.pango_font_map_list_families(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(fontmap)

	var _families []FontFamilier // out

	defer C.free(unsafe.Pointer(_arg1))
	{
		src := unsafe.Slice((**C.PangoFontFamily)(_arg1), _arg2)
		_families = make([]FontFamilier, _arg2)
		for i := 0; i < int(_arg2); i++ {
			{
				objptr := unsafe.Pointer(src[i])
				if objptr == nil {
					panic("object of type pango.FontFamilier is nil")
				}

				object := coreglib.Take(objptr)
				casted := object.WalkCast(func(obj coreglib.Objector) bool {
					_, ok := obj.(FontFamilier)
					return ok
				})
				rv, ok := casted.(FontFamilier)
				if !ok {
					panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFamilier")
				}
				_families[i] = rv
			}
		}
	}

	return _families
}

// LoadFont: load the font in the fontmap that is the closest match for desc.
//
// The function takes the following parameters:
//
//    - context: PangoContext the font will be used with.
//    - desc: PangoFontDescription describing the font to load.
//
// The function returns the following values:
//
//    - font (optional): newly allocated PangoFont loaded, or NULL if no font
//      matched.
//
func (fontmap *FontMap) LoadFont(context *Context, desc *FontDescription) Fonter {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out
	var _cret *C.PangoFont            // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))

	_cret = C.pango_font_map_load_font(_arg0, _arg1, _arg2)
	runtime.KeepAlive(fontmap)
	runtime.KeepAlive(context)
	runtime.KeepAlive(desc)

	var _font Fonter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Fonter)
				return ok
			})
			rv, ok := casted.(Fonter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fonter")
			}
			_font = rv
		}
	}

	return _font
}

// LoadFontset: load a set of fonts in the fontmap that can be used to render a
// font matching desc.
//
// The function takes the following parameters:
//
//    - context: PangoContext the font will be used with.
//    - desc: PangoFontDescription describing the font to load.
//    - language: PangoLanguage the fonts will be used for.
//
// The function returns the following values:
//
//    - fontset (optional): newly allocated PangoFontset loaded, or NULL if no
//      font matched.
//
func (fontmap *FontMap) LoadFontset(context *Context, desc *FontDescription, language *Language) Fontsetter {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out
	var _arg3 *C.PangoLanguage        // out
	var _cret *C.PangoFontset         // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))
	_arg3 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_font_map_load_fontset(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fontmap)
	runtime.KeepAlive(context)
	runtime.KeepAlive(desc)
	runtime.KeepAlive(language)

	var _fontset Fontsetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Fontsetter)
				return ok
			})
			rv, ok := casted.(Fontsetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fontsetter")
			}
			_fontset = rv
		}
	}

	return _fontset
}
