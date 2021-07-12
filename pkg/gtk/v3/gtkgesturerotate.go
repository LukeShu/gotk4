// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_rotate_get_type()), F: marshalGestureRotater},
	})
}

// GestureRotater describes GestureRotate's methods.
type GestureRotater interface {
	// AngleDelta: if @gesture is active, this function returns the angle
	// difference in radians since the gesture was first recognized.
	AngleDelta() float64
}

// GestureRotate is a Gesture implementation able to recognize 2-finger
// rotations, whenever the angle between both handled sequences changes, the
// GestureRotate::angle-changed signal is emitted.
type GestureRotate struct {
	Gesture
}

var (
	_ GestureRotater  = (*GestureRotate)(nil)
	_ gextras.Nativer = (*GestureRotate)(nil)
)

func wrapGestureRotate(obj *externglib.Object) GestureRotater {
	return &GestureRotate{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureRotater(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureRotate(obj), nil
}

// NewGestureRotate returns a newly created Gesture that recognizes 2-touch
// rotation gestures.
func NewGestureRotate(widget Widgeter) *GestureRotate {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	_cret = C.gtk_gesture_rotate_new(_arg1)

	var _gestureRotate *GestureRotate // out

	_gestureRotate = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*GestureRotate)

	return _gestureRotate
}

// AngleDelta: if @gesture is active, this function returns the angle difference
// in radians since the gesture was first recognized. If @gesture is not active,
// 0 is returned.
func (gesture *GestureRotate) AngleDelta() float64 {
	var _arg0 *C.GtkGestureRotate // out
	var _cret C.gdouble           // in

	_arg0 = (*C.GtkGestureRotate)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_rotate_get_angle_delta(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
