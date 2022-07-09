// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern PangoFont* _gotk4_pango1_FontsetClass_get_font(void*, guint);
// extern PangoFontMetrics* _gotk4_pango1_FontsetClass_get_metrics(void*);
// extern PangoLanguage* _gotk4_pango1_FontsetClass_get_language(void*);
// extern gboolean _gotk4_pango1_FontsetForEachFunc(void*, void*, gpointer);
import "C"

// GTypeFontset returns the GType for the type Fontset.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontset() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "Fontset").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontset)
	return gtype
}

// GTypeFontsetSimple returns the GType for the type FontsetSimple.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontsetSimple() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "FontsetSimple").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontsetSimple)
	return gtype
}

// FontsetForEachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForEachFunc func(fontset Fontsetter, font Fonter) (ok bool)

//export _gotk4_pango1_FontsetForEachFunc
func _gotk4_pango1_FontsetForEachFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) (cret C.gboolean) {
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
	pclass := girepository.MustFind("Pango", "FontsetClass")

	if _, ok := goval.(interface{ Font(wc uint32) Fonter }); ok {
		o := pclass.StructFieldOffset("get_font")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_pango1_FontsetClass_get_font)
	}

	if _, ok := goval.(interface{ Language() *Language }); ok {
		o := pclass.StructFieldOffset("get_language")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_pango1_FontsetClass_get_language)
	}

	if _, ok := goval.(interface{ Metrics() *FontMetrics }); ok {
		o := pclass.StructFieldOffset("get_metrics")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_pango1_FontsetClass_get_metrics)
	}
}

//export _gotk4_pango1_FontsetClass_get_font
func _gotk4_pango1_FontsetClass_get_font(arg0 *C.void, arg1 C.guint) (cret *C.PangoFont) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Font(wc uint32) Fonter })

	var _wc uint32 // out

	_wc = uint32(arg1)

	font := iface.Font(_wc)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(font).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(font).Native()))

	return cret
}

//export _gotk4_pango1_FontsetClass_get_language
func _gotk4_pango1_FontsetClass_get_language(arg0 *C.void) (cret *C.PangoLanguage) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Language() *Language })

	language := iface.Language()

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(language)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(language)), nil)

	return cret
}

//export _gotk4_pango1_FontsetClass_get_metrics
func _gotk4_pango1_FontsetClass_get_metrics(arg0 *C.void) (cret *C.PangoFontMetrics) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Metrics() *FontMetrics })

	fontMetrics := iface.Metrics()

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(fontMetrics)))

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_pango1_FontsetForEachFunc)
	_args[2] = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_args[2]))

	girepository.MustFind("Pango", "Fontset").InvokeMethod("foreach", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(wc)

	_gret := girepository.MustFind("Pango", "Fontset").InvokeMethod("get_font", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))

	_gret := girepository.MustFind("Pango", "Fontset").InvokeMethod("get_metrics", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(language)))

	_gret := girepository.MustFind("Pango", "FontsetSimple").InvokeMethod("new_FontsetSimple", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(font).Native()))

	girepository.MustFind("Pango", "FontsetSimple").InvokeMethod("append", _args[:], nil)

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontset).Native()))

	_gret := girepository.MustFind("Pango", "FontsetSimple").InvokeMethod("size", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontset)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}
