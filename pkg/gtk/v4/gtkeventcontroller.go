// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_event_controller_get_type()), F: marshalEventController},
	})
}

// EventController: `GtkEventController` is the base class for event
// controllers.
//
// These are ancillary objects associated to widgets, which react to
// `GdkEvents`, and possibly trigger actions as a consequence.
//
// Event controllers are added to a widget with
// [method@Gtk.Widget.add_controller]. It is rarely necessary to explicitly
// remove a controller with [method@Gtk.Widget.remove_controller].
//
// See the chapter of input handling (input-handling.html) for an overview of
// the basic concepts, such as the capture and bubble phases of even
// propagation.
type EventController interface {
	gextras.Objector

	// CurrentEvent returns the event that is currently being handled by the
	// controller, and nil at other times.
	CurrentEvent() gdk.Event
	// CurrentEventDevice returns the device of the event that is currently
	// being handled by the controller, and nil otherwise.
	CurrentEventDevice() gdk.Device
	// CurrentEventState returns the modifier state of the event that is
	// currently being handled by the controller, and 0 otherwise.
	CurrentEventState() gdk.ModifierType
	// CurrentEventTime returns the timestamp of the event that is currently
	// being handled by the controller, and 0 otherwise.
	CurrentEventTime() uint32
	// Name gets the name of @controller.
	Name() string
	// PropagationLimit gets the propagation limit of the event controller.
	PropagationLimit() PropagationLimit
	// PropagationPhase gets the propagation phase at which @controller handles
	// events.
	PropagationPhase() PropagationPhase
	// Widget returns the Widget this controller relates to.
	Widget() Widget
	// Reset resets the @controller to a clean state.
	Reset()
	// SetName sets a name on the controller that can be used for debugging.
	SetName(name string)
	// SetPropagationLimit sets the event propagation limit on the event
	// controller.
	//
	// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
	// events that are targeted at widgets on a different surface, such as
	// popovers.
	SetPropagationLimit(limit PropagationLimit)
	// SetPropagationPhase sets the propagation phase at which a controller
	// handles events.
	//
	// If @phase is GTK_PHASE_NONE, no automatic event handling will be
	// performed, but other additional gesture maintenance will.
	SetPropagationPhase(phase PropagationPhase)
}

// EventControllerClass implements the EventController interface.
type EventControllerClass struct {
	*externglib.Object
}

var _ EventController = (*EventControllerClass)(nil)

func wrapEventController(obj *externglib.Object) EventController {
	return &EventControllerClass{
		Object: obj,
	}
}

func marshalEventController(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventController(obj), nil
}

// CurrentEvent returns the event that is currently being handled by the
// controller, and nil at other times.
func (c *EventControllerClass) CurrentEvent() gdk.Event {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GdkEvent           // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_current_event(_arg0)

	var _event gdk.Event // out

	_event = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Event)

	return _event
}

// CurrentEventDevice returns the device of the event that is currently being
// handled by the controller, and nil otherwise.
func (c *EventControllerClass) CurrentEventDevice() gdk.Device {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GdkDevice          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_current_event_device(_arg0)

	var _device gdk.Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Device)

	return _device
}

// CurrentEventState returns the modifier state of the event that is currently
// being handled by the controller, and 0 otherwise.
func (c *EventControllerClass) CurrentEventState() gdk.ModifierType {
	var _arg0 *C.GtkEventController // out
	var _cret C.GdkModifierType     // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_current_event_state(_arg0)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// CurrentEventTime returns the timestamp of the event that is currently being
// handled by the controller, and 0 otherwise.
func (c *EventControllerClass) CurrentEventTime() uint32 {
	var _arg0 *C.GtkEventController // out
	var _cret C.guint32             // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_current_event_time(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Name gets the name of @controller.
func (c *EventControllerClass) Name() string {
	var _arg0 *C.GtkEventController // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PropagationLimit gets the propagation limit of the event controller.
func (c *EventControllerClass) PropagationLimit() PropagationLimit {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationLimit // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_propagation_limit(_arg0)

	var _propagationLimit PropagationLimit // out

	_propagationLimit = PropagationLimit(_cret)

	return _propagationLimit
}

// PropagationPhase gets the propagation phase at which @controller handles
// events.
func (c *EventControllerClass) PropagationPhase() PropagationPhase {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationPhase // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_propagation_phase(_arg0)

	var _propagationPhase PropagationPhase // out

	_propagationPhase = PropagationPhase(_cret)

	return _propagationPhase
}

// Widget returns the Widget this controller relates to.
func (c *EventControllerClass) Widget() Widget {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_event_controller_get_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

// Reset resets the @controller to a clean state.
func (c *EventControllerClass) Reset() {
	var _arg0 *C.GtkEventController // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	C.gtk_event_controller_reset(_arg0)
}

// SetName sets a name on the controller that can be used for debugging.
func (c *EventControllerClass) SetName(name string) {
	var _arg0 *C.GtkEventController // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_event_controller_set_name(_arg0, _arg1)
}

// SetPropagationLimit sets the event propagation limit on the event controller.
//
// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
// events that are targeted at widgets on a different surface, such as popovers.
func (c *EventControllerClass) SetPropagationLimit(limit PropagationLimit) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationLimit // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))
	_arg1 = C.GtkPropagationLimit(limit)

	C.gtk_event_controller_set_propagation_limit(_arg0, _arg1)
}

// SetPropagationPhase sets the propagation phase at which a controller handles
// events.
//
// If @phase is GTK_PHASE_NONE, no automatic event handling will be performed,
// but other additional gesture maintenance will.
func (c *EventControllerClass) SetPropagationPhase(phase PropagationPhase) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationPhase // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))
	_arg1 = C.GtkPropagationPhase(phase)

	C.gtk_event_controller_set_propagation_phase(_arg0, _arg1)
}
