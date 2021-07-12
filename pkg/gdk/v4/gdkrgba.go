// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_rgba_get_type()), F: marshalRGBA},
	})
}

// RGBA: `GdkRGBA` is used to represent a color, in a way that is compatible
// with cairo’s notion of color.
//
// `GdkRGBA` is a convenient way to pass colors around. It’s based on cairo’s
// way to deal with colors and mirrors its behavior. All values are in the range
// from 0.0 to 1.0 inclusive. So the color (0.0, 0.0, 0.0, 0.0) represents
// transparent black and (1.0, 1.0, 1.0, 1.0) is opaque white. Other values will
// be clamped to this range when drawing.
type RGBA struct {
	native C.GdkRGBA
}

func marshalRGBA(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*RGBA)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RGBA) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Copy makes a copy of a `GdkRGBA`.
//
// The result must be freed through [method@Gdk.RGBA.free].
func (rgba *RGBA) Copy() *RGBA {
	var _arg0 *C.GdkRGBA // out
	var _cret *C.GdkRGBA // in

	_arg0 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	_cret = C.gdk_rgba_copy(_arg0)

	var _rgbA *RGBA // out

	_rgbA = (*RGBA)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_rgbA, func(v *RGBA) {
		C.gdk_rgba_free((*C.GdkRGBA)(unsafe.Pointer(v)))
	})

	return _rgbA
}

// Equal compares two `GdkRGBA` colors.
func (p1 *RGBA) Equal(p2 *RGBA) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = C.gconstpointer(unsafe.Pointer(p1))
	_arg1 = C.gconstpointer(unsafe.Pointer(p2))

	_cret = C.gdk_rgba_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Free frees a `GdkRGBA`.
func (rgba *RGBA) free() {
	var _arg0 *C.GdkRGBA // out

	_arg0 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	C.gdk_rgba_free(_arg0)
}

// Hash: hash function suitable for using for a hash table that stores
// `GdkRGBA`s.
func (p *RGBA) Hash() uint {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = C.gconstpointer(unsafe.Pointer(p))

	_cret = C.gdk_rgba_hash(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IsClear checks if an @rgba value is transparent.
//
// That is, drawing with the value would not produce any change.
func (rgba *RGBA) IsClear() bool {
	var _arg0 *C.GdkRGBA // out
	var _cret C.gboolean // in

	_arg0 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	_cret = C.gdk_rgba_is_clear(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsOpaque checks if an @rgba value is opaque.
//
// That is, drawing with the value will not retain any results from previous
// contents.
func (rgba *RGBA) IsOpaque() bool {
	var _arg0 *C.GdkRGBA // out
	var _cret C.gboolean // in

	_arg0 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	_cret = C.gdk_rgba_is_opaque(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Parse parses a textual representation of a color.
//
// The string can be either one of:
//
// - A standard name (Taken from the X11 rgb.txt file). - A hexadecimal value in
// the form “\#rgb”, “\#rrggbb”, “\#rrrgggbbb” or ”\#rrrrggggbbbb” - A
// hexadecimal value in the form “\#rgba”, “\#rrggbbaa”, or ”\#rrrrggggbbbbaaaa”
// - A RGB color in the form “rgb(r,g,b)” (In this case the color will have full
// opacity) - A RGBA color in the form “rgba(r,g,b,a)”
//
// Where “r”, “g”, “b” and “a” are respectively the red, green, blue and alpha
// color values. In the last two cases, “r”, “g”, and “b” are either integers in
// the range 0 to 255 or percentage values in the range 0% to 100%, and a is a
// floating point value in the range 0 to 1.
func (rgba *RGBA) Parse(spec string) bool {
	var _arg0 *C.GdkRGBA // out
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg0 = (*C.GdkRGBA)(unsafe.Pointer(rgba))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(spec)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_rgba_parse(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String returns a textual specification of @rgba in the form `rgb(r,g,b)` or
// `rgba(r,g,b,a)`, where “r”, “g”, “b” and “a” represent the red, green, blue
// and alpha values respectively. “r”, “g”, and “b” are represented as integers
// in the range 0 to 255, and “a” is represented as a floating point value in
// the range 0 to 1.
//
// These string forms are string forms that are supported by the CSS3 colors
// module, and can be parsed by [method@Gdk.RGBA.parse].
//
// Note that this string representation may lose some precision, since “r”, “g”
// and “b” are represented as 8-bit integers. If this is a concern, you should
// use a different representation.
func (rgba *RGBA) String() string {
	var _arg0 *C.GdkRGBA // out
	var _cret *C.char    // in

	_arg0 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	_cret = C.gdk_rgba_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
