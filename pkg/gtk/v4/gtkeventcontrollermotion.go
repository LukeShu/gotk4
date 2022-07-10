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
// #include <glib-object.h>
// extern void _gotk4_gtk4_EventControllerMotion_ConnectEnter(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_EventControllerMotion_ConnectLeave(gpointer, guintptr);
// extern void _gotk4_gtk4_EventControllerMotion_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
import "C"

// GTypeEventControllerMotion returns the GType for the type EventControllerMotion.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEventControllerMotion() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "EventControllerMotion").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalEventControllerMotion)
	return gtype
}

// EventControllerMotion: GtkEventControllerMotion is an event controller
// tracking the pointer position.
//
// The event controller offers gtk.EventControllerMotion::enter and
// gtk.EventControllerMotion::leave signals, as well as
// gtk.EventControllerMotion:is-pointer and
// gtk.EventControllerMotion:contains-pointer properties which are updated to
// reflect changes in the pointer position as it moves over the widget.
type EventControllerMotion struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerMotion)(nil)
)

func wrapEventControllerMotion(obj *coreglib.Object) *EventControllerMotion {
	return &EventControllerMotion{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerMotion(p uintptr) (interface{}, error) {
	return wrapEventControllerMotion(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_EventControllerMotion_ConnectEnter
func _gotk4_gtk4_EventControllerMotion_ConnectEnter(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectEnter signals that the pointer has entered the widget.
func (self *EventControllerMotion) ConnectEnter(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "enter", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerMotion_ConnectEnter), f)
}

//export _gotk4_gtk4_EventControllerMotion_ConnectLeave
func _gotk4_gtk4_EventControllerMotion_ConnectLeave(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectLeave signals that the pointer has left the widget.
func (self *EventControllerMotion) ConnectLeave(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "leave", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerMotion_ConnectLeave), f)
}

//export _gotk4_gtk4_EventControllerMotion_ConnectMotion
func _gotk4_gtk4_EventControllerMotion_ConnectMotion(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectMotion is emitted when the pointer moves inside the widget.
func (self *EventControllerMotion) ConnectMotion(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "motion", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerMotion_ConnectMotion), f)
}

// NewEventControllerMotion creates a new event controller that will handle
// motion events.
//
// The function returns the following values:
//
//    - eventControllerMotion: new GtkEventControllerMotion.
//
func NewEventControllerMotion() *EventControllerMotion {
	_info := girepository.MustFind("Gtk", "EventControllerMotion")
	_gret := _info.InvokeClassMethod("new_EventControllerMotion", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _eventControllerMotion *EventControllerMotion // out

	_eventControllerMotion = wrapEventControllerMotion(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _eventControllerMotion
}

// ContainsPointer returns if a pointer is within self or one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if a pointer is within self or one of its children.
//
func (self *EventControllerMotion) ContainsPointer() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "EventControllerMotion")
	_gret := _info.InvokeClassMethod("contains_pointer", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsPointer returns if a pointer is within self, but not one of its children.
//
// The function returns the following values:
//
//    - ok: TRUE if a pointer is within self but not one of its children.
//
func (self *EventControllerMotion) IsPointer() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "EventControllerMotion")
	_gret := _info.InvokeClassMethod("is_pointer", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
