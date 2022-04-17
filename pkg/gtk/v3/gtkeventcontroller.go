// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkeventcontroller.go.
var GTypeEventController = externglib.Type(C.gtk_event_controller_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeEventController, F: marshalEventController},
	})
}

// EventControllerOverrider contains methods that are overridable.
type EventControllerOverrider interface {
	externglib.Objector
}

// WrapEventControllerOverrider wraps the EventControllerOverrider
// interface implementation to access the instance methods.
func WrapEventControllerOverrider(obj EventControllerOverrider) *EventController {
	return wrapEventController(externglib.BaseObject(obj))
}

// EventController is a base, low-level implementation for event controllers.
// Those react to a series of Events, and possibly trigger actions as a
// consequence of those.
type EventController struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*EventController)(nil)
)

// EventControllerer describes types inherited from class EventController.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type EventControllerer interface {
	externglib.Objector
	baseEventController() *EventController
}

var _ EventControllerer = (*EventController)(nil)

func classInitEventControllerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEventController(obj *externglib.Object) *EventController {
	return &EventController{
		Object: obj,
	}
}

func marshalEventController(p uintptr) (interface{}, error) {
	return wrapEventController(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (controller *EventController) baseEventController() *EventController {
	return controller
}

// BaseEventController returns the underlying base object from the
// interface.
func BaseEventController(obj EventControllerer) *EventController {
	return obj.baseEventController()
}

// PropagationPhase gets the propagation phase at which controller handles
// events.
//
// The function returns the following values:
//
//    - propagationPhase: propagation phase.
//
func (controller *EventController) PropagationPhase() PropagationPhase {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationPhase // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_propagation_phase(_arg0)
	runtime.KeepAlive(controller)

	var _propagationPhase PropagationPhase // out

	_propagationPhase = PropagationPhase(_cret)

	return _propagationPhase
}

// Widget returns the Widget this controller relates to.
//
// The function returns the following values:
//
//    - widget: Widget.
//
func (controller *EventController) Widget() Widgetter {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_widget(_arg0)
	runtime.KeepAlive(controller)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// HandleEvent feeds an events into controller, so it can be interpreted and the
// controller actions triggered.
//
// The function takes the following parameters:
//
//    - event: Event.
//
// The function returns the following values:
//
//    - ok: TRUE if the event was potentially useful to trigger the controller
//      action.
//
func (controller *EventController) HandleEvent(event *gdk.Event) bool {
	var _arg0 *C.GtkEventController // out
	var _arg1 *C.GdkEvent           // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))
	_arg1 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))

	_cret = C.gtk_event_controller_handle_event(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(event)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reset resets the controller to a clean state. Every interaction the
// controller did through EventController::handle-event will be dropped at this
// point.
func (controller *EventController) Reset() {
	var _arg0 *C.GtkEventController // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	C.gtk_event_controller_reset(_arg0)
	runtime.KeepAlive(controller)
}

// SetPropagationPhase sets the propagation phase at which a controller handles
// events.
//
// If phase is GTK_PHASE_NONE, no automatic event handling will be performed,
// but other additional gesture maintenance will. In that phase, the events can
// be managed by calling gtk_event_controller_handle_event().
//
// The function takes the following parameters:
//
//    - phase: propagation phase.
//
func (controller *EventController) SetPropagationPhase(phase PropagationPhase) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationPhase // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))
	_arg1 = C.GtkPropagationPhase(phase)

	C.gtk_event_controller_set_propagation_phase(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(phase)
}
