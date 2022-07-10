// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// RGBToHSV converts a color from RGB space to HSV.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
//
// The function takes the following parameters:
//
//    - r: red.
//    - g: green.
//    - b: blue.
//
// The function returns the following values:
//
//    - h: return value for the hue component.
//    - s: return value for the saturation component.
//    - v: return value for the value component.
//
func RGBToHSV(r, g, b float64) (h, s, v float64) {
	var _args [3]girepository.Argument
	var _outs [3]girepository.Argument

	*(*C.gdouble)(unsafe.Pointer(&_args[0])) = C.gdouble(r)
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(g)
	*(*C.gdouble)(unsafe.Pointer(&_args[2])) = C.gdouble(b)

	_info := girepository.MustFind("Gtk", "rgb_to_hsv")
	_info.InvokeFunction(_args[:], _outs[:])

	runtime.KeepAlive(r)
	runtime.KeepAlive(g)
	runtime.KeepAlive(b)

	var _h float64 // out
	var _s float64 // out
	var _v float64 // out

	_h = float64(*(*C.gdouble)(unsafe.Pointer(&_outs[0])))
	_s = float64(*(*C.gdouble)(unsafe.Pointer(&_outs[1])))
	_v = float64(*(*C.gdouble)(unsafe.Pointer(&_outs[2])))

	return _h, _s, _v
}

// HSVToRGB converts a color from HSV space to RGB.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
//
// The function takes the following parameters:
//
//    - h: hue.
//    - s: saturation.
//    - v: value.
//
// The function returns the following values:
//
//    - r: return value for the red component.
//    - g: return value for the green component.
//    - b: return value for the blue component.
//
func HSVToRGB(h, s, v float64) (r, g, b float64) {
	var _args [3]girepository.Argument
	var _outs [3]girepository.Argument

	*(*C.gdouble)(unsafe.Pointer(&_args[0])) = C.gdouble(h)
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(s)
	*(*C.gdouble)(unsafe.Pointer(&_args[2])) = C.gdouble(v)

	_info := girepository.MustFind("Gtk", "to_rgb")
	_info.InvokeFunction(_args[:], _outs[:])

	runtime.KeepAlive(h)
	runtime.KeepAlive(s)
	runtime.KeepAlive(v)

	var _r float64 // out
	var _g float64 // out
	var _b float64 // out

	_r = float64(*(*C.gdouble)(unsafe.Pointer(&_outs[0])))
	_g = float64(*(*C.gdouble)(unsafe.Pointer(&_outs[1])))
	_b = float64(*(*C.gdouble)(unsafe.Pointer(&_outs[2])))

	return _r, _g, _b
}
