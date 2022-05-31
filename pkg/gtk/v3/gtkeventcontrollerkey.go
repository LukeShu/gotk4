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
// extern void _gotk4_gtk3_EventControllerKey_ConnectFocusIn(gpointer, guintptr);
// extern void _gotk4_gtk3_EventControllerKey_ConnectFocusOut(gpointer, guintptr);
// extern void _gotk4_gtk3_EventControllerKey_ConnectIMUpdate(gpointer, guintptr);
import "C"

// glib.Type values for gtkeventcontrollerkey.go.
var GTypeEventControllerKey = coreglib.Type(C.gtk_event_controller_key_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeEventControllerKey, F: marshalEventControllerKey},
	})
}

// EventControllerKeyOverrider contains methods that are overridable.
type EventControllerKeyOverrider interface {
}

// EventControllerKey is an event controller meant for situations where you need
// access to key events.
//
// This object was added in 3.24.
type EventControllerKey struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerKey)(nil)
)

func classInitEventControllerKeyer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEventControllerKey(obj *coreglib.Object) *EventControllerKey {
	return &EventControllerKey{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	return wrapEventControllerKey(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_EventControllerKey_ConnectFocusIn
func _gotk4_gtk3_EventControllerKey_ConnectFocusIn(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

func (controller *EventControllerKey) ConnectFocusIn(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "focus-in", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectFocusIn), f)
}

//export _gotk4_gtk3_EventControllerKey_ConnectFocusOut
func _gotk4_gtk3_EventControllerKey_ConnectFocusOut(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

func (controller *EventControllerKey) ConnectFocusOut(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "focus-out", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectFocusOut), f)
}

//export _gotk4_gtk3_EventControllerKey_ConnectIMUpdate
func _gotk4_gtk3_EventControllerKey_ConnectIMUpdate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

func (controller *EventControllerKey) ConnectIMUpdate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(controller, "im-update", false, unsafe.Pointer(C._gotk4_gtk3_EventControllerKey_ConnectIMUpdate), f)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewEventControllerKey(widget Widgetter) *EventControllerKey {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "EventControllerKey").InvokeMethod("new_EventControllerKey", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _eventControllerKey *EventControllerKey // out

	_eventControllerKey = wrapEventControllerKey(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerKey
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (controller *EventControllerKey) Forward(widget Widgetter) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(controller).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	*(**EventControllerKey)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "EventControllerKey").InvokeMethod("forward", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(controller)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
func (controller *EventControllerKey) Group() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(controller).Native()))
	*(**EventControllerKey)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "EventControllerKey").InvokeMethod("get_group", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(controller)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// IMContext gets the IM context of a key controller.
//
// The function returns the following values:
//
//    - imContext: IM context.
//
func (controller *EventControllerKey) IMContext() IMContexter {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(controller).Native()))
	*(**EventControllerKey)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "EventControllerKey").InvokeMethod("get_im_context", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(controller)

	var _imContext IMContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.IMContexter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(IMContexter)
			return ok
		})
		rv, ok := casted.(IMContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.IMContexter")
		}
		_imContext = rv
	}

	return _imContext
}

// The function takes the following parameters:
//
func (controller *EventControllerKey) SetIMContext(imContext IMContexter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(controller).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(imContext).Native()))
	*(**EventControllerKey)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "EventControllerKey").InvokeMethod("set_im_context", args[:], nil)

	runtime.KeepAlive(controller)
	runtime.KeepAlive(imContext)
}
