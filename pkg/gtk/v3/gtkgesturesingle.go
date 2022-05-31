// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkgesturesingle.go.
var GTypeGestureSingle = coreglib.Type(C.gtk_gesture_single_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeGestureSingle, F: marshalGestureSingle},
	})
}

// GestureSingleOverrider contains methods that are overridable.
type GestureSingleOverrider interface {
}

// GestureSingle is a subclass of Gesture, optimized (although not restricted)
// for dealing with mouse and single-touch gestures. Under interaction, these
// gestures stick to the first interacting sequence, which is accessible through
// gtk_gesture_single_get_current_sequence() while the gesture is being
// interacted with.
//
// By default gestures react to both GDK_BUTTON_PRIMARY and touch events,
// gtk_gesture_single_set_touch_only() can be used to change the touch behavior.
// Callers may also specify a different mouse button number to interact with
// through gtk_gesture_single_set_button(), or react to any mouse button by
// setting 0. While the gesture is active, the button being currently pressed
// can be known through gtk_gesture_single_get_current_button().
type GestureSingle struct {
	_ [0]func() // equal guard
	Gesture
}

var (
	_ Gesturer = (*GestureSingle)(nil)
)

func classInitGestureSingler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGestureSingle(obj *coreglib.Object) *GestureSingle {
	return &GestureSingle{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureSingle(p uintptr) (interface{}, error) {
	return wrapGestureSingle(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Button returns the button number gesture listens for, or 0 if gesture reacts
// to any button press.
//
// The function returns the following values:
//
//    - guint: button number, or 0 for any button.
//
func (gesture *GestureSingle) Button() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	*(**GestureSingle)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("get_button", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// CurrentButton returns the button number currently interacting with gesture,
// or 0 if there is none.
//
// The function returns the following values:
//
//    - guint: current button number.
//
func (gesture *GestureSingle) CurrentButton() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	*(**GestureSingle)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("get_current_button", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// CurrentSequence returns the event sequence currently interacting with
// gesture. This is only meaningful if gtk_gesture_is_active() returns TRUE.
//
// The function returns the following values:
//
//    - eventSequence (optional): current sequence.
//
func (gesture *GestureSingle) CurrentSequence() *gdk.EventSequence {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	*(**GestureSingle)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("get_current_sequence", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _eventSequence *gdk.EventSequence // out

	if _cret != nil {
		_eventSequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_eventSequence)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _eventSequence
}

// Exclusive gets whether a gesture is exclusive. For more information, see
// gtk_gesture_single_set_exclusive().
//
// The function returns the following values:
//
//    - ok: whether the gesture is exclusive.
//
func (gesture *GestureSingle) Exclusive() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	*(**GestureSingle)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("get_exclusive", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TouchOnly returns TRUE if the gesture is only triggered by touch events.
//
// The function returns the following values:
//
//    - ok: TRUE if the gesture only handles touch events.
//
func (gesture *GestureSingle) TouchOnly() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	*(**GestureSingle)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("get_touch_only", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetButton sets the button number gesture listens to. If non-0, every button
// press from a different button number will be ignored. Touch events implicitly
// match with button 1.
//
// The function takes the following parameters:
//
//    - button number to listen to, or 0 for any button.
//
func (gesture *GestureSingle) SetButton(button uint32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	_arg1 = C.guint(button)
	*(**GestureSingle)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("set_button", args[:], nil)

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(button)
}

// SetExclusive sets whether gesture is exclusive. An exclusive gesture will
// only handle pointer and "pointer emulated" touch events, so at any given
// time, there is only one sequence able to interact with those.
//
// The function takes the following parameters:
//
//    - exclusive: TRUE to make gesture exclusive.
//
func (gesture *GestureSingle) SetExclusive(exclusive bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	if exclusive {
		_arg1 = C.TRUE
	}
	*(**GestureSingle)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("set_exclusive", args[:], nil)

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(exclusive)
}

// SetTouchOnly: if touch_only is TRUE, gesture will only handle events of type
// K_TOUCH_BEGIN, K_TOUCH_UPDATE or K_TOUCH_END. If FALSE, mouse events will be
// handled too.
//
// The function takes the following parameters:
//
//    - touchOnly: whether gesture handles only touch events.
//
func (gesture *GestureSingle) SetTouchOnly(touchOnly bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))
	if touchOnly {
		_arg1 = C.TRUE
	}
	*(**GestureSingle)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "GestureSingle").InvokeMethod("set_touch_only", args[:], nil)

	runtime.KeepAlive(gesture)
	runtime.KeepAlive(touchOnly)
}
