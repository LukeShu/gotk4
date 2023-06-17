// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// ExtentsToPixels converts extents from Pango units to device units.
//
// The conversion is done by dividing by the PANGO_SCALE factor and performing
// rounding.
//
// The inclusive rectangle is converted by flooring the x/y coordinates and
// extending width/height, such that the final rectangle completely includes the
// original rectangle.
//
// The nearest rectangle is converted by rounding the coordinates of the
// rectangle to the nearest device unit (pixel).
//
// The rule to which argument to use is: if you want the resulting device-space
// rectangle to completely contain the original rectangle, pass it in as
// inclusive. If you want two touching-but-not-overlapping rectangles stay
// touching-but-not-overlapping after rounding to device units, pass them in as
// nearest.
//
// The function takes the following parameters:
//
//   - inclusive (optional): rectangle to round to pixels inclusively, or NULL.
//   - nearest (optional): rectangle to round to nearest pixels, or NULL.
//
func ExtentsToPixels(inclusive, nearest *Rectangle) {
	var _arg1 *C.PangoRectangle // out
	var _arg2 *C.PangoRectangle // out

	if inclusive != nil {
		_arg1 = (*C.PangoRectangle)(gextras.StructNative(unsafe.Pointer(inclusive)))
	}
	if nearest != nil {
		_arg2 = (*C.PangoRectangle)(gextras.StructNative(unsafe.Pointer(nearest)))
	}

	C.pango_extents_to_pixels(_arg1, _arg2)
	runtime.KeepAlive(inclusive)
	runtime.KeepAlive(nearest)
}

// UnitsFromDouble converts a floating-point number to Pango units.
//
// The conversion is done by multiplying d by PANGO_SCALE and rounding the
// result to nearest integer.
//
// The function takes the following parameters:
//
//   - d: double floating-point value.
//
// The function returns the following values:
//
//   - gint: value in Pango units.
//
func UnitsFromDouble(d float64) int {
	var _arg1 C.double // out
	var _cret C.int    // in

	_arg1 = C.double(d)

	_cret = C.pango_units_from_double(_arg1)
	runtime.KeepAlive(d)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UnitsToDouble converts a number in Pango units to floating-point.
//
// The conversion is done by dividing i by PANGO_SCALE.
//
// The function takes the following parameters:
//
//   - i: value in Pango units.
//
// The function returns the following values:
//
//   - gdouble: double value.
//
func UnitsToDouble(i int) float64 {
	var _arg1 C.int    // out
	var _cret C.double // in

	_arg1 = C.int(i)

	_cret = C.pango_units_to_double(_arg1)
	runtime.KeepAlive(i)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
