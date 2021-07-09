// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_focus_get_type()), F: marshalEventControllerFocus},
	})
}

// EventControllerFocus: `GtkEventControllerFocus` is an event controller to
// keep track of keyboard focus.
//
// The event controller offers [signal@Gtk.EventControllerFocus::enter] and
// [signal@Gtk.EventControllerFocus::leave] signals, as well as
// [property@Gtk.EventControllerFocus:is-focus] and
// [property@Gtk.EventControllerFocus:contains-focus] properties which are
// updated to reflect focus changes inside the widget hierarchy that is rooted
// at the controllers widget.
type EventControllerFocus interface {
	gextras.Objector

	// ContainsFocus returns true if focus is within @self or one of its
	// children.
	ContainsFocus() bool
	// IsFocus returns true if focus is within @self, but not one of its
	// children.
	IsFocus() bool
}

// EventControllerFocusClass implements the EventControllerFocus interface.
type EventControllerFocusClass struct {
	EventControllerClass
}

var _ EventControllerFocus = (*EventControllerFocusClass)(nil)

func wrapEventControllerFocus(obj *externglib.Object) EventControllerFocus {
	return &EventControllerFocusClass{
		EventControllerClass: EventControllerClass{
			Object: obj,
		},
	}
}

func marshalEventControllerFocus(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerFocus(obj), nil
}

// NewEventControllerFocus creates a new event controller that will handle focus
// events.
func NewEventControllerFocus() *EventControllerFocusClass {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_focus_new()

	var _eventControllerFocus *EventControllerFocusClass // out

	_eventControllerFocus = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*EventControllerFocusClass)

	return _eventControllerFocus
}

// ContainsFocus returns true if focus is within @self or one of its children.
func (s *EventControllerFocusClass) ContainsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_event_controller_focus_contains_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsFocus returns true if focus is within @self, but not one of its children.
func (s *EventControllerFocusClass) IsFocus() bool {
	var _arg0 *C.GtkEventControllerFocus // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_event_controller_focus_is_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
