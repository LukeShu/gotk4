// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewScale creates a new Scale.
//
// The function takes the following parameters:
//
//    - orientation scale’s orientation.
//    - adjustment (optional) which sets the range of the scale, or NULL to
//      create a new adjustment.
//
// The function returns the following values:
//
//    - scale: new Scale.
//
func NewScale(orientation Orientation, adjustment *Adjustment) *Scale {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	if adjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_cret = C.gtk_scale_new(_arg1, _arg2)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(adjustment)

	var _scale *Scale // out

	_scale = wrapScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _scale
}

// NewScaleWithRange creates a new scale widget with the given orientation that
// lets the user input a number between min and max (including min and max) with
// the increment step. step must be nonzero; it’s the distance the slider moves
// when using the arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_scale_set_digits() to correct it.
//
// The function takes the following parameters:
//
//    - orientation scale’s orientation.
//    - min: minimum value.
//    - max: maximum value.
//    - step increment (tick size) used with keyboard shortcuts.
//
// The function returns the following values:
//
//    - scale: new Scale.
//
func NewScaleWithRange(orientation Orientation, min, max, step float64) *Scale {
	var _arg1 C.GtkOrientation // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.gdouble(min)
	_arg3 = C.gdouble(max)
	_arg4 = C.gdouble(step)

	_cret = C.gtk_scale_new_with_range(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _scale *Scale // out

	_scale = wrapScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _scale
}
