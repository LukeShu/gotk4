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

// NewGestureDrag returns a newly created Gesture that recognizes drags.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureDrag: newly created GestureDrag.
//
func NewGestureDrag(widget Widgetter) *GestureDrag {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_gesture_drag_new(_arg1)
	runtime.KeepAlive(widget)

	var _gestureDrag *GestureDrag // out

	_gestureDrag = wrapGestureDrag(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureDrag
}

// Offset: if the gesture is active, this function returns TRUE and fills in x
// and y with the coordinates of the current point, as an offset to the starting
// drag point.
//
// The function returns the following values:
//
//    - x (optional): x offset for the current point.
//    - y (optional): y offset for the current point.
//    - ok: TRUE if the gesture is active.
//
func (gesture *GestureDrag) Offset() (x, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.gdouble         // in
	var _arg2 C.gdouble         // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_cret = C.gtk_gesture_drag_get_offset(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg1)
	_y = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// StartPoint: if the gesture is active, this function returns TRUE and fills in
// x and y with the drag start coordinates, in window-relative coordinates.
//
// The function returns the following values:
//
//    - x (optional): x coordinate for the drag start point.
//    - y (optional): y coordinate for the drag start point.
//    - ok: TRUE if the gesture is active.
//
func (gesture *GestureDrag) StartPoint() (x, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.gdouble         // in
	var _arg2 C.gdouble         // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_cret = C.gtk_gesture_drag_get_start_point(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg1)
	_y = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}
