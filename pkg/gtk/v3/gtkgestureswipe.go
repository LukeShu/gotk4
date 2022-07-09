// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_GestureSwipe_ConnectSwipe(gpointer, gdouble, gdouble, guintptr);
import "C"

// GTypeGestureSwipe returns the GType for the type GestureSwipe.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGestureSwipe() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "GestureSwipe").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGestureSwipe)
	return gtype
}

// GestureSwipe is a Gesture implementation able to recognize swipes, after a
// press/move/.../move/release sequence happens, the GestureSwipe::swipe signal
// will be emitted, providing the velocity and directionality of the sequence at
// the time it was lifted.
//
// If the velocity is desired in intermediate points,
// gtk_gesture_swipe_get_velocity() can be called on eg. a Gesture::update
// handler.
//
// All velocities are reported in pixels/sec units.
type GestureSwipe struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureSwipe)(nil)
)

func wrapGestureSwipe(obj *coreglib.Object) *GestureSwipe {
	return &GestureSwipe{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureSwipe(p uintptr) (interface{}, error) {
	return wrapGestureSwipe(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_GestureSwipe_ConnectSwipe
func _gotk4_gtk3_GestureSwipe_ConnectSwipe(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(velocityX, velocityY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(velocityX, velocityY float64))
	}

	var _velocityX float64 // out
	var _velocityY float64 // out

	_velocityX = float64(arg1)
	_velocityY = float64(arg2)

	f(_velocityX, _velocityY)
}

// ConnectSwipe: this signal is emitted when the recognized gesture is finished,
// velocity and direction are a product of previously recorded events.
func (gesture *GestureSwipe) ConnectSwipe(f func(velocityX, velocityY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "swipe", false, unsafe.Pointer(C._gotk4_gtk3_GestureSwipe_ConnectSwipe), f)
}

// NewGestureSwipe returns a newly created Gesture that recognizes swipes.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureSwipe: newly created GestureSwipe.
//
func NewGestureSwipe(widget Widgetter) *GestureSwipe {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_gret := girepository.MustFind("Gtk", "GestureSwipe").InvokeMethod("new_GestureSwipe", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _gestureSwipe *GestureSwipe // out

	_gestureSwipe = wrapGestureSwipe(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureSwipe
}

// Velocity: if the gesture is recognized, this function returns TRUE and fill
// in velocity_x and velocity_y with the recorded velocity, as per the last
// event(s) processed.
//
// The function returns the following values:
//
//    - velocityX: return value for the velocity in the X axis, in pixels/sec.
//    - velocityY: return value for the velocity in the Y axis, in pixels/sec.
//    - ok: whether velocity could be calculated.
//
func (gesture *GestureSwipe) Velocity() (velocityX, velocityY float64, ok bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_gret := girepository.MustFind("Gtk", "GestureSwipe").InvokeMethod("get_velocity", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _velocityX float64 // out
	var _velocityY float64 // out
	var _ok bool           // out

	_velocityX = *(*float64)(unsafe.Pointer(_outs[0]))
	_velocityY = *(*float64)(unsafe.Pointer(_outs[1]))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _velocityX, _velocityY, _ok
}
