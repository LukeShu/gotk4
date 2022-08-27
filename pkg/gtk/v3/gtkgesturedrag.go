// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
