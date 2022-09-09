// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_GestureDrag_ConnectDragUpdate(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureDrag_ConnectDragEnd(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureDrag_ConnectDragBegin(gpointer, gdouble, gdouble, guintptr);
import "C"

// GType values.
var (
	GTypeGestureDrag = coreglib.Type(C.gtk_gesture_drag_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGestureDrag, F: marshalGestureDrag},
	})
}

// GestureDrag is a Gesture implementation that recognizes drag operations. The
// drag operation itself can be tracked throught the GestureDrag::drag-begin,
// GestureDrag::drag-update and GestureDrag::drag-end signals, or the relevant
// coordinates be extracted through gtk_gesture_drag_get_offset() and
// gtk_gesture_drag_get_start_point().
type GestureDrag struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureDrag)(nil)
)

func wrapGestureDrag(obj *coreglib.Object) *GestureDrag {
	return &GestureDrag{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureDrag(p uintptr) (interface{}, error) {
	return wrapGestureDrag(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectDragBegin: this signal is emitted whenever dragging starts.
func (gesture *GestureDrag) ConnectDragBegin(f func(startX, startY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "drag-begin", false, unsafe.Pointer(C._gotk4_gtk3_GestureDrag_ConnectDragBegin), f)
}

// ConnectDragEnd: this signal is emitted whenever the dragging is finished.
func (gesture *GestureDrag) ConnectDragEnd(f func(offsetX, offsetY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "drag-end", false, unsafe.Pointer(C._gotk4_gtk3_GestureDrag_ConnectDragEnd), f)
}

// ConnectDragUpdate: this signal is emitted whenever the dragging point moves.
func (gesture *GestureDrag) ConnectDragUpdate(f func(offsetX, offsetY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "drag-update", false, unsafe.Pointer(C._gotk4_gtk3_GestureDrag_ConnectDragUpdate), f)
}

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
