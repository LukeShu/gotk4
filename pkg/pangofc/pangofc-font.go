// Code generated by girgen. DO NOT EDIT.

package pangofc

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango pangofc
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangofc-fontmap.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fc_font_get_type()), F: marshalFont},
	})
}

// Font: `PangoFcFont` is a base class for font implementations using the
// Fontconfig and FreeType libraries.
//
// It is used in onjunction with [class@PangoFc.FontMap]. When deriving from
// this class, you need to implement all of its virtual functions other than
// shutdown() along with the get_glyph_extents() virtual function from
// `PangoFont`.
type Font interface {
	Font

	// Glyph gets the glyph index for a given Unicode character for @font.
	//
	// If you only want to determine whether the font has the glyph, use
	// [method@PangoFc.Font.has_char].
	Glyph(wc uint32) uint
	// HasChar determines whether @font has a glyph for the codepoint @wc.
	HasChar(wc uint32) bool
	// UnlockFace releases a font previously obtained with
	// [method@PangoFc.Font.lock_face].
	UnlockFace()
}

// font implements the Font class.
type font struct {
	Font
}

var _ Font = (*font)(nil)

// WrapFont wraps a GObject to the right type. It is
// primarily used internally.
func WrapFont(obj *externglib.Object) Font {
	return font{
		Font: WrapFont(obj),
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFont(obj), nil
}

// Glyph gets the glyph index for a given Unicode character for @font.
//
// If you only want to determine whether the font has the glyph, use
// [method@PangoFc.Font.has_char].
func (f font) Glyph(wc uint32) uint {
	var _arg0 *C.PangoFcFont // out
	var _arg1 C.gunichar     // out

	_arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))
	_arg1 = C.gunichar(wc)

	var _cret C.guint // in

	_cret = C.pango_fc_font_get_glyph(_arg0, _arg1)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// HasChar determines whether @font has a glyph for the codepoint @wc.
func (f font) HasChar(wc uint32) bool {
	var _arg0 *C.PangoFcFont // out
	var _arg1 C.gunichar     // out

	_arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))
	_arg1 = C.gunichar(wc)

	var _cret C.gboolean // in

	_cret = C.pango_fc_font_has_char(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnlockFace releases a font previously obtained with
// [method@PangoFc.Font.lock_face].
func (f font) UnlockFace() {
	var _arg0 *C.PangoFcFont // out

	_arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))

	C.pango_fc_font_unlock_face(_arg0)
}
