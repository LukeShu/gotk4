// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeEventControllerScrollFlags = coreglib.Type(C.gtk_event_controller_scroll_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEventControllerScrollFlags, F: marshalEventControllerScrollFlags},
	})
}

// EventControllerScrollFlags describes the behavior of a EventControllerScroll.
type EventControllerScrollFlags C.guint

const (
	// EventControllerScrollNone: don't emit scroll.
	EventControllerScrollNone EventControllerScrollFlags = 0b0
	// EventControllerScrollVertical: emit scroll with vertical deltas.
	EventControllerScrollVertical EventControllerScrollFlags = 0b1
	// EventControllerScrollHorizontal: emit scroll with horizontal deltas.
	EventControllerScrollHorizontal EventControllerScrollFlags = 0b10
	// EventControllerScrollDiscrete: only emit deltas that are multiples of 1.
	EventControllerScrollDiscrete EventControllerScrollFlags = 0b100
	// EventControllerScrollKinetic: emit EventControllerScroll::decelerate
	// after continuous scroll finishes.
	EventControllerScrollKinetic EventControllerScrollFlags = 0b1000
	// EventControllerScrollBothAxes: emit scroll on both axes.
	EventControllerScrollBothAxes EventControllerScrollFlags = 0b11
)

func marshalEventControllerScrollFlags(p uintptr) (interface{}, error) {
	return EventControllerScrollFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for EventControllerScrollFlags.
func (e EventControllerScrollFlags) String() string {
	if e == 0 {
		return "EventControllerScrollFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(176)

	for e != 0 {
		next := e & (e - 1)
		bit := e - next

		switch bit {
		case EventControllerScrollNone:
			builder.WriteString("None|")
		case EventControllerScrollVertical:
			builder.WriteString("Vertical|")
		case EventControllerScrollHorizontal:
			builder.WriteString("Horizontal|")
		case EventControllerScrollDiscrete:
			builder.WriteString("Discrete|")
		case EventControllerScrollKinetic:
			builder.WriteString("Kinetic|")
		case EventControllerScrollBothAxes:
			builder.WriteString("BothAxes|")
		default:
			builder.WriteString(fmt.Sprintf("EventControllerScrollFlags(0b%b)|", bit))
		}

		e = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if e contains other.
func (e EventControllerScrollFlags) Has(other EventControllerScrollFlags) bool {
	return (e & other) == other
}

// NewEventControllerScroll creates a new event controller that will handle
// scroll events for the given widget.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//    - flags: behavior flags.
//
// The function returns the following values:
//
//    - eventControllerScroll: new EventControllerScroll.
//
func NewEventControllerScroll(widget Widgetter, flags EventControllerScrollFlags) *EventControllerScroll {
	var _arg1 *C.GtkWidget                    // out
	var _arg2 C.GtkEventControllerScrollFlags // out
	var _cret *C.GtkEventController           // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = C.GtkEventControllerScrollFlags(flags)

	_cret = C.gtk_event_controller_scroll_new(_arg1, _arg2)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(flags)

	var _eventControllerScroll *EventControllerScroll // out

	_eventControllerScroll = wrapEventControllerScroll(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerScroll
}

// Flags gets the flags conditioning the scroll controller behavior.
//
// The function returns the following values:
//
//    - eventControllerScrollFlags: controller flags.
//
func (controller *EventControllerScroll) Flags() EventControllerScrollFlags {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _cret C.GtkEventControllerScrollFlags // in

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(coreglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_scroll_get_flags(_arg0)
	runtime.KeepAlive(controller)

	var _eventControllerScrollFlags EventControllerScrollFlags // out

	_eventControllerScrollFlags = EventControllerScrollFlags(_cret)

	return _eventControllerScrollFlags
}

// SetFlags sets the flags conditioning scroll controller behavior.
//
// The function takes the following parameters:
//
//    - flags: behavior flags.
//
func (controller *EventControllerScroll) SetFlags(flags EventControllerScrollFlags) {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _arg1 C.GtkEventControllerScrollFlags // out

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(coreglib.InternObject(controller).Native()))
	_arg1 = C.GtkEventControllerScrollFlags(flags)

	C.gtk_event_controller_scroll_set_flags(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(flags)
}
