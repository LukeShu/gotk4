// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKeyer},
	})
}

// EventControllerKey: GtkEventControllerKey is an event controller that
// provides access to key events.
type EventControllerKey struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerKey)(nil)
)

func wrapEventControllerKey(obj *externglib.Object) *EventControllerKey {
	return &EventControllerKey{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerKeyer(p uintptr) (interface{}, error) {
	return wrapEventControllerKey(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectIMUpdate: emitted whenever the input method context filters away a
// keypress and prevents the controller receiving it.
//
// See gtk.EventControllerKey.SetIMContext() and gtk.IMContext.FilterKeypress().
func (controller *EventControllerKey) ConnectIMUpdate(f func()) externglib.SignalHandle {
	return controller.Connect("im-update", f)
}

// ConnectKeyPressed: emitted whenever a key is pressed.
func (controller *EventControllerKey) ConnectKeyPressed(f func(keyval, keycode uint, state gdk.ModifierType) bool) externglib.SignalHandle {
	return controller.Connect("key-pressed", f)
}

// ConnectKeyReleased: emitted whenever a key is released.
func (controller *EventControllerKey) ConnectKeyReleased(f func(keyval, keycode uint, state gdk.ModifierType)) externglib.SignalHandle {
	return controller.Connect("key-released", f)
}

// ConnectModifiers: emitted whenever the state of modifier keys and pointer
// buttons change.
func (controller *EventControllerKey) ConnectModifiers(f func(keyval gdk.ModifierType) bool) externglib.SignalHandle {
	return controller.Connect("modifiers", f)
}

// NewEventControllerKey creates a new event controller that will handle key
// events.
//
// The function returns the following values:
//
//    - eventControllerKey: new GtkEventControllerKey.
//
func NewEventControllerKey() *EventControllerKey {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_key_new()

	var _eventControllerKey *EventControllerKey // out

	_eventControllerKey = wrapEventControllerKey(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerKey
}

// Forward forwards the current event of this controller to a widget.
//
// This function can only be used in handlers for the
// gtk.EventControllerKey::key-pressed, gtk.EventControllerKey::key-released or
// gtk.EventControllerKey::modifiers signals.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget.
//
// The function returns the following values:
//
//    - ok: whether the widget handled the event.
//
func (controller *EventControllerKey) Forward(widget Widgetter) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Group gets the key group of the current event of this controller.
//
// See gdk.KeyEvent.GetLayout().
//
// The function returns the following values:
//
//    - guint: key group.
//
func (controller *EventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret C.guint                  // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_group(_arg0)
	runtime.KeepAlive(controller)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IMContext gets the input method context of the key controller.
//
// The function returns the following values:
//
//    - imContext: GtkIMContext.
//
func (controller *EventControllerKey) IMContext() IMContexter {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret *C.GtkIMContext          // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)
	runtime.KeepAlive(controller)

	var _imContext IMContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.IMContexter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(IMContexter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.IMContexter")
		}
		_imContext = rv
	}

	return _imContext
}

// SetIMContext sets the input method context of the key controller.
//
// The function takes the following parameters:
//
//    - imContext: GtkIMContext.
//
func (controller *EventControllerKey) SetIMContext(imContext IMContexter) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(imContext)
}
