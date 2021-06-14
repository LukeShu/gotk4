// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKey},
	})
}

// EventControllerKey is an event controller meant for situations where you need
// access to key events.
//
// This object was added in 3.24.
type EventControllerKey interface {
	EventController

	Forward(widget Widget) bool

	Group() uint
	// ImContext gets the IM context of a key controller.
	ImContext() IMContext

	SetImContext(imContext IMContext)
}

// eventControllerKey implements the EventControllerKey class.
type eventControllerKey struct {
	EventController
}

var _ EventControllerKey = (*eventControllerKey)(nil)

// WrapEventControllerKey wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerKey(obj *externglib.Object) EventControllerKey {
	return eventControllerKey{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerKey(obj), nil
}

// NewEventControllerKey constructs a class EventControllerKey.
func NewEventControllerKey(widget Widget) EventControllerKey {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret C.GtkEventControllerKey // in

	_cret = C.gtk_event_controller_key_new(_arg1)

	var _eventControllerKey EventControllerKey // out

	_eventControllerKey = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(EventControllerKey)

	return _eventControllerKey
}

func (c eventControllerKey) Forward(widget Widget) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c eventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	var _cret C.guint // in

	_cret = C.gtk_event_controller_key_get_group(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// ImContext gets the IM context of a key controller.
func (c eventControllerKey) ImContext() IMContext {
	var _arg0 *C.GtkEventControllerKey // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	var _cret *C.GtkIMContext // in

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)

	var _imContext IMContext // out

	_imContext = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(IMContext)

	return _imContext
}

func (c eventControllerKey) SetImContext(imContext IMContext) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
}
