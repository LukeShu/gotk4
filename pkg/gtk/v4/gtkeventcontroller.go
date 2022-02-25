// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
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
}

// EventController: GtkEventController is the base class for event controllers.
//
// These are ancillary objects associated to widgets, which react to GdkEvents,
// and possibly trigger actions as a consequence.
//
// Event controllers are added to a widget with gtk.Widget.AddController(). It
// is rarely necessary to explicitly remove a controller with
// gtk.Widget.RemoveController().
//
// See the chapter of input handling (input-handling.html) for an overview of
// the basic concepts, such as the capture and bubble phases of even
// propagation.
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

// BaseEventController returns the underlying base object.
func BaseEventController(obj EventControllerer) *EventController {
	return obj.baseEventController()
}

// CurrentEvent returns the event that is currently being handled by the
// controller, and NULL at other times.
//
// The function returns the following values:
//
//    - event (optional) that is currently handled by controller.
//
func (controller *EventController) CurrentEvent() gdk.Eventer {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GdkEvent           // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_current_event(_arg0)
	runtime.KeepAlive(controller)

	var _event gdk.Eventer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gdk.Eventer)
				return ok
			})
			rv, ok := casted.(gdk.Eventer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Eventer")
			}
			_event = rv
		}
	}

	return _event
}

// CurrentEventDevice returns the device of the event that is currently being
// handled by the controller, and NULL otherwise.
//
// The function returns the following values:
//
//    - device (optional) of the event is currently handled by controller.
//
func (controller *EventController) CurrentEventDevice() gdk.Devicer {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GdkDevice          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_current_event_device(_arg0)
	runtime.KeepAlive(controller)

	var _device gdk.Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gdk.Devicer)
				return ok
			})
			rv, ok := casted.(gdk.Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// CurrentEventState returns the modifier state of the event that is currently
// being handled by the controller, and 0 otherwise.
//
// The function returns the following values:
//
//    - modifierType: modifier state of the event is currently handled by
//      controller.
//
func (controller *EventController) CurrentEventState() gdk.ModifierType {
	var _arg0 *C.GtkEventController // out
	var _cret C.GdkModifierType     // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_current_event_state(_arg0)
	runtime.KeepAlive(controller)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// CurrentEventTime returns the timestamp of the event that is currently being
// handled by the controller, and 0 otherwise.
//
// The function returns the following values:
//
//    - guint32: timestamp of the event is currently handled by controller.
//
func (controller *EventController) CurrentEventTime() uint32 {
	var _arg0 *C.GtkEventController // out
	var _cret C.guint32             // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_current_event_time(_arg0)
	runtime.KeepAlive(controller)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Name gets the name of controller.
//
// The function returns the following values:
//
func (controller *EventController) Name() string {
	var _arg0 *C.GtkEventController // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_name(_arg0)
	runtime.KeepAlive(controller)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PropagationLimit gets the propagation limit of the event controller.
//
// The function returns the following values:
//
//    - propagationLimit: propagation limit.
//
func (controller *EventController) PropagationLimit() PropagationLimit {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationLimit // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	_cret = C.gtk_event_controller_get_propagation_limit(_arg0)
	runtime.KeepAlive(controller)

	var _propagationLimit PropagationLimit // out

	_propagationLimit = PropagationLimit(_cret)

	return _propagationLimit
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
//    - widget: GtkWidget.
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

// Reset resets the controller to a clean state.
func (controller *EventController) Reset() {
	var _arg0 *C.GtkEventController // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))

	C.gtk_event_controller_reset(_arg0)
	runtime.KeepAlive(controller)
}

// SetName sets a name on the controller that can be used for debugging.
//
// The function takes the following parameters:
//
//    - name for controller.
//
func (controller *EventController) SetName(name string) {
	var _arg0 *C.GtkEventController // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_event_controller_set_name(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(name)
}

// SetPropagationLimit sets the event propagation limit on the event controller.
//
// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
// events that are targeted at widgets on a different surface, such as popovers.
//
// The function takes the following parameters:
//
//    - limit: propagation limit.
//
func (controller *EventController) SetPropagationLimit(limit PropagationLimit) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationLimit // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(externglib.InternObject(controller).Native()))
	_arg1 = C.GtkPropagationLimit(limit)

	C.gtk_event_controller_set_propagation_limit(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(limit)
}

// SetPropagationPhase sets the propagation phase at which a controller handles
// events.
//
// If phase is GTK_PHASE_NONE, no automatic event handling will be performed,
// but other additional gesture maintenance will.
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
