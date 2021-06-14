// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_font_map_get_type()), F: marshalFontMap},
	})
}

// FontMap: a `PangoFontMap` represents the set of fonts available for a
// particular rendering system.
//
// This is a virtual object with implementations being specific to particular
// rendering systems.
type FontMap interface {
	gextras.Objector

	// Changed forces a change in the context, which will cause any
	// `PangoContext` using this fontmap to change.
	//
	// This function is only useful when implementing a new backend for Pango,
	// something applications won't do. Backends should call this function if
	// they have attached extra data to the context and such data is changed.
	Changed()
	// CreateContext creates a `PangoContext` connected to @fontmap.
	//
	// This is equivalent to [ctor@Pango.Context.new] followed by
	// [method@Pango.Context.set_font_map].
	//
	// If you are using Pango as part of a higher-level system, that system may
	// have it's own way of create a `PangoContext`. For instance, the GTK
	// toolkit has, among others, gtk_widget_get_pango_context(). Use those
	// instead.
	CreateContext() Context
	// Family gets a font family by name.
	Family(name string) FontFamily
	// Serial returns the current serial number of @fontmap.
	//
	// The serial number is initialized to an small number larger than zero when
	// a new fontmap is created and is increased whenever the fontmap is
	// changed. It may wrap, but will never have the value 0. Since it can wrap,
	// never compare it with "less than", always use "not equals".
	//
	// The fontmap can only be changed using backend-specific API, like changing
	// fontmap resolution.
	//
	// This can be used to automatically detect changes to a `PangoFontMap`,
	// like in `PangoContext`.
	Serial() uint
	// ListFamilies: list all families for a fontmap.
	ListFamilies() []FontFamily
	// LoadFont: load the font in the fontmap that is the closest match for
	// @desc.
	LoadFont(context Context, desc *FontDescription) Font
	// LoadFontset: load a set of fonts in the fontmap that can be used to
	// render a font matching @desc.
	LoadFontset(context Context, desc *FontDescription, language *Language) Fontset
}

// fontMap implements the FontMap class.
type fontMap struct {
	gextras.Objector
}

var _ FontMap = (*fontMap)(nil)

// WrapFontMap wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return fontMap{
		Objector: obj,
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}

// Changed forces a change in the context, which will cause any
// `PangoContext` using this fontmap to change.
//
// This function is only useful when implementing a new backend for Pango,
// something applications won't do. Backends should call this function if
// they have attached extra data to the context and such data is changed.
func (f fontMap) Changed() {
	var _arg0 *C.PangoFontMap // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	C.pango_font_map_changed(_arg0)
}

// CreateContext creates a `PangoContext` connected to @fontmap.
//
// This is equivalent to [ctor@Pango.Context.new] followed by
// [method@Pango.Context.set_font_map].
//
// If you are using Pango as part of a higher-level system, that system may
// have it's own way of create a `PangoContext`. For instance, the GTK
// toolkit has, among others, gtk_widget_get_pango_context(). Use those
// instead.
func (f fontMap) CreateContext() Context {
	var _arg0 *C.PangoFontMap // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	var _cret *C.PangoContext // in

	_cret = C.pango_font_map_create_context(_arg0)

	var _context Context // out

	_context = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(Context)

	return _context
}

// Family gets a font family by name.
func (f fontMap) Family(name string) FontFamily {
	var _arg0 *C.PangoFontMap // out
	var _arg1 *C.char         // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.PangoFontFamily // in

	_cret = C.pango_font_map_get_family(_arg0, _arg1)

	var _fontFamily FontFamily // out

	_fontFamily = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FontFamily)

	return _fontFamily
}

// Serial returns the current serial number of @fontmap.
//
// The serial number is initialized to an small number larger than zero when
// a new fontmap is created and is increased whenever the fontmap is
// changed. It may wrap, but will never have the value 0. Since it can wrap,
// never compare it with "less than", always use "not equals".
//
// The fontmap can only be changed using backend-specific API, like changing
// fontmap resolution.
//
// This can be used to automatically detect changes to a `PangoFontMap`,
// like in `PangoContext`.
func (f fontMap) Serial() uint {
	var _arg0 *C.PangoFontMap // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	var _cret C.guint // in

	_cret = C.pango_font_map_get_serial(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// ListFamilies: list all families for a fontmap.
func (f fontMap) ListFamilies() []FontFamily {
	var _arg0 *C.PangoFontMap // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	var _arg1 **C.PangoFontFamily
	var _arg2 C.int // in

	C.pango_font_map_list_families(_arg0, &_arg1, &_arg2)

	var _families []FontFamily

	{
		src := unsafe.Slice(_arg1, _arg2)
		defer C.free(unsafe.Pointer(_arg1))
		_families = make([]FontFamily, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_families[i] = gextras.CastObject(externglib.Take(unsafe.Pointer(src[i].Native()))).(FontFamily)
		}
	}

	return _families
}

// LoadFont: load the font in the fontmap that is the closest match for
// @desc.
func (f fontMap) LoadFont(context Context, desc *FontDescription) Font {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))

	var _cret *C.PangoFont // in

	_cret = C.pango_font_map_load_font(_arg0, _arg1, _arg2)

	var _font Font // out

	_font = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(Font)

	return _font
}

// LoadFontset: load a set of fonts in the fontmap that can be used to
// render a font matching @desc.
func (f fontMap) LoadFontset(context Context, desc *FontDescription, language *Language) Fontset {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out
	var _arg3 *C.PangoLanguage        // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))
	_arg3 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	var _cret *C.PangoFontset // in

	_cret = C.pango_font_map_load_fontset(_arg0, _arg1, _arg2, _arg3)

	var _fontset Fontset // out

	_fontset = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(Fontset)

	return _fontset
}
