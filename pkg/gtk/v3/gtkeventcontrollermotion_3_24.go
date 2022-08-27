// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewEventControllerMotion creates a new event controller that will handle
// motion events for the given widget.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - eventControllerMotion: new EventControllerMotion.
//
func NewEventControllerMotion(widget Widgetter) *EventControllerMotion {
	var _arg1 *C.GtkWidget          // out
	var _cret *C.GtkEventController // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_event_controller_motion_new(_arg1)
	runtime.KeepAlive(widget)

	var _eventControllerMotion *EventControllerMotion // out

	_eventControllerMotion = wrapEventControllerMotion(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerMotion
}
