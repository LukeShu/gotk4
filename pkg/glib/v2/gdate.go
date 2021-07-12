// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// DateValidJulian returns true if the Julian day is valid. Anything greater
// than zero is basically a valid Julian, though there is a 32-bit limit.
func DateValidJulian(julianDate uint32) bool {
	var _arg1 C.guint32  // out
	var _cret C.gboolean // in

	_arg1 = C.guint32(julianDate)

	_cret = C.g_date_valid_julian(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
