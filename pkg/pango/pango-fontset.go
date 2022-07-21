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
// extern PangoFont* _gotk4_pango1_FontsetClass_get_font(PangoFontset*, guint);
// extern PangoFontMetrics* _gotk4_pango1_FontsetClass_get_metrics(PangoFontset*);
// extern PangoLanguage* _gotk4_pango1_FontsetClass_get_language(PangoFontset*);
// extern gboolean _gotk4_pango1_FontsetForEachFunc(PangoFontset*, PangoFont*, gpointer);
import "C"

// GTypeFontset returns the GType for the type Fontset.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontset() coreglib.Type {
	gtype := coreglib.Type(C.pango_fontset_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontset)
	return gtype
}

// GTypeFontsetSimple returns the GType for the type FontsetSimple.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontsetSimple() coreglib.Type {
	gtype := coreglib.Type(C.pango_fontset_simple_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontsetSimple)
	return gtype
}

// FontsetForEachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForEachFunc func(fontset Fontsetter, font Fonter) (ok bool)

//export _gotk4_pango1_FontsetForEachFunc
func _gotk4_pango1_FontsetForEachFunc(arg1 *C.PangoFontset, arg2 *C.PangoFont, arg3 C.gpointer) (cret C.gboolean) {
	var fn FontsetForEachFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FontsetForEachFunc)
	}

	var _fontset Fontsetter // out
	var _font Fonter        // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.Fontsetter is nil")
		}

		object := coreglib.Take(objptr)
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
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

		object := coreglib.Take(objptr)
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

	ok := fn(_fontset, _font)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontsetOverrider contains methods that are overridable.
type FontsetOverrider interface {
	// Font returns the font in the fontset that contains the best glyph for a
	// Unicode character.
	//
	// The function takes the following parameters:
	//
	//    - wc: unicode character.
	//
	// The function returns the following values:
	//
	//    - font: PangoFont. The caller must call g_object_unref() when finished
	//      with the font.
	//
	Font(wc uint32) Fonter
	// The function returns the following values:
	//
	Language() *Language
	// Metrics: get overall metric information for the fonts in the fontset.
	//
	// The function returns the following values:
	//
	//    - fontMetrics object. The caller must call pango_font_metrics_unref()
	//      when finished using the object.
	//
	Metrics() *FontMetrics
}

// Fontset: PangoFontset represents a set of PangoFont to use when rendering
// text.
//
// A PAngoFontset is the result of resolving a PangoFontDescription against a
// particular PangoContext. It has operations for finding the component font for
// a particular Unicode character, and for finding a composite set of metrics
// for the entire fontset.
type Fontset struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Fontset)(nil)
)

// Fontsetter describes types inherited from class Fontset.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Fontsetter interface {
	coreglib.Objector
	baseFontset() *Fontset
}

var _ Fontsetter = (*Fontset)(nil)

func classInitFontsetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.PangoFontsetClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ Font(wc uint32) Fonter }); ok {
		pclass.get_font = (*[0]byte)(C._gotk4_pango1_FontsetClass_get_font)
	}

	if _, ok := goval.(interface{ Language() *Language }); ok {
		pclass.get_language = (*[0]byte)(C._gotk4_pango1_FontsetClass_get_language)
	}

	if _, ok := goval.(interface{ Metrics() *FontMetrics }); ok {
		pclass.get_metrics = (*[0]byte)(C._gotk4_pango1_FontsetClass_get_metrics)
	}
}

//export _gotk4_pango1_FontsetClass_get_font
func _gotk4_pango1_FontsetClass_get_font(arg0 *C.PangoFontset, arg1 C.guint) (cret *C.PangoFont) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Font(wc uint32) Fonter })

	var _wc uint32 // out

	_wc = uint32(arg1)

	font := iface.Font(_wc)

	cret = (*C.PangoFont)(unsafe.Pointer(coreglib.InternObject(font).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(font).Native()))

	return cret
}

//export _gotk4_pango1_FontsetClass_get_language
func _gotk4_pango1_FontsetClass_get_language(arg0 *C.PangoFontset) (cret *C.PangoLanguage) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Language() *Language })

	language := iface.Language()

	cret = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(language)), nil)

	return cret
}

//export _gotk4_pango1_FontsetClass_get_metrics
func _gotk4_pango1_FontsetClass_get_metrics(arg0 *C.PangoFontset) (cret *C.PangoFontMetrics) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Metrics() *FontMetrics })

	fontMetrics := iface.Metrics()

	cret = (*C.PangoFontMetrics)(gextras.StructNative(unsafe.Pointer(fontMetrics)))

	return cret
}

func wrapFontset(obj *coreglib.Object) *Fontset {
	return &Fontset{
		Object: obj,
	}
}

func marshalFontset(p uintptr) (interface{}, error) {
	return wrapFontset(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (fontset *Fontset) baseFontset() *Fontset {
	return fontset
}

// BaseFontset returns the underlying base object.
func BaseFontset(obj Fontsetter) *Fontset {
	return obj.baseFontset()
}

// ForEach iterates through all the fonts in a fontset, calling func for each
// one.
//
// If func returns TRUE, that stops the iteration.
//
// The function takes the following parameters:
//
//    - fn: callback function.
//
func (fontset *Fontset) ForEach(fn FontsetForEachFunc) {
	var _arg0 *C.PangoFontset           // out
	var _arg1 C.PangoFontsetForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))
	_arg1 = (*[0]byte)(C._gotk4_pango1_FontsetForEachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.pango_fontset_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(fn)
}

// Font returns the font in the fontset that contains the best glyph for a
// Unicode character.
//
// The function takes the following parameters:
//
//    - wc: unicode character.
//
// The function returns the following values:
//
//    - font: PangoFont. The caller must call g_object_unref() when finished with
//      the font.
//
func (fontset *Fontset) Font(wc uint32) Fonter {
	var _arg0 *C.PangoFontset // out
	var _arg1 C.guint         // out
	var _cret *C.PangoFont    // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))
	_arg1 = C.guint(wc)

	_cret = C.pango_fontset_get_font(_arg0, _arg1)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(wc)

	var _font Fonter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

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

	return _font
}

// Metrics: get overall metric information for the fonts in the fontset.
//
// The function returns the following values:
//
//    - fontMetrics object. The caller must call pango_font_metrics_unref() when
//      finished using the object.
//
func (fontset *Fontset) Metrics() *FontMetrics {
	var _arg0 *C.PangoFontset     // out
	var _cret *C.PangoFontMetrics // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))

	_cret = C.pango_fontset_get_metrics(_arg0)
	runtime.KeepAlive(fontset)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_fontMetrics)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _fontMetrics
}

// FontsetSimpleOverrider contains methods that are overridable.
type FontsetSimpleOverrider interface {
}

// FontsetSimple: PangoFontsetSimple is a implementation of the abstract
// PangoFontset base class as an array of fonts.
//
// When creating a PangoFontsetSimple, you have to provide the array of fonts
// that make up the fontset.
type FontsetSimple struct {
	_ [0]func() // equal guard
	Fontset
}

var (
	_ Fontsetter = (*FontsetSimple)(nil)
)

func classInitFontsetSimpler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFontsetSimple(obj *coreglib.Object) *FontsetSimple {
	return &FontsetSimple{
		Fontset: Fontset{
			Object: obj,
		},
	}
}

func marshalFontsetSimple(p uintptr) (interface{}, error) {
	return wrapFontsetSimple(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontsetSimple creates a new PangoFontsetSimple for the given language.
//
// The function takes the following parameters:
//
//    - language: PangoLanguage tag.
//
// The function returns the following values:
//
//    - fontsetSimple: newly allocated PangoFontsetSimple, which should be freed
//      with g_object_unref().
//
func NewFontsetSimple(language *Language) *FontsetSimple {
	var _arg1 *C.PangoLanguage      // out
	var _cret *C.PangoFontsetSimple // in

	_arg1 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_fontset_simple_new(_arg1)
	runtime.KeepAlive(language)

	var _fontsetSimple *FontsetSimple // out

	_fontsetSimple = wrapFontsetSimple(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fontsetSimple
}

// Append adds a font to the fontset.
//
// The function takes the following parameters:
//
//    - font: PangoFont.
//
func (fontset *FontsetSimple) Append(font Fonter) {
	var _arg0 *C.PangoFontsetSimple // out
	var _arg1 *C.PangoFont          // out

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(coreglib.InternObject(font).Native()))

	C.pango_fontset_simple_append(_arg0, _arg1)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(font)
}

// Size returns the number of fonts in the fontset.
//
// The function returns the following values:
//
//    - gint: size of fontset.
//
func (fontset *FontsetSimple) Size() int32 {
	var _arg0 *C.PangoFontsetSimple // out
	var _cret C.int                 // in

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))

	_cret = C.pango_fontset_simple_size(_arg0)
	runtime.KeepAlive(fontset)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}
