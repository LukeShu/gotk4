// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_EventControllerLegacy_ConnectEvent(gpointer, GdkEvent*, guintptr);
import "C"

// GType values.
var (
	GTypeEventControllerLegacy = coreglib.Type(C.gtk_event_controller_legacy_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEventControllerLegacy, F: marshalEventControllerLegacy},
	})
}

// EventControllerLegacy: GtkEventControllerLegacy is an event controller that
// provides raw access to the event stream.
//
// It should only be used as a last resort if none of the other event
// controllers or gestures do the job.
type EventControllerLegacy struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*EventControllerLegacy)(nil)
)

func wrapEventControllerLegacy(obj *coreglib.Object) *EventControllerLegacy {
	return &EventControllerLegacy{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerLegacy(p uintptr) (interface{}, error) {
	return wrapEventControllerLegacy(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEvent is emitted for each GDK event delivered to controller.
func (v *EventControllerLegacy) ConnectEvent(f func(event gdk.Eventer) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "event", false, unsafe.Pointer(C._gotk4_gtk4_EventControllerLegacy_ConnectEvent), f)
}

// NewEventControllerLegacy creates a new legacy event controller.
//
// The function returns the following values:
//
//    - eventControllerLegacy: newly created event controller.
//
func NewEventControllerLegacy() *EventControllerLegacy {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_legacy_new()

	var _eventControllerLegacy *EventControllerLegacy // out

	_eventControllerLegacy = wrapEventControllerLegacy(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerLegacy
}
