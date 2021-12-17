// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_gesture_get_type()), F: marshalGesturer},
	})
}

// Gesture: GtkGesture is the base class for gesture recognition.
//
// Although GtkGesture is quite generalized to serve as a base for multi-touch
// gestures, it is suitable to implement single-touch and pointer-based gestures
// (using the special NULL GdkEventSequence value for these).
//
// The number of touches that a GtkGesture need to be recognized is controlled
// by the gtk.Gesture:n-points property, if a gesture is keeping track of less
// or more than that number of sequences, it won't check whether the gesture is
// recognized.
//
// As soon as the gesture has the expected number of touches, it will check
// regularly if it is recognized, the criteria to consider a gesture as
// "recognized" is left to GtkGesture subclasses.
//
// A recognized gesture will then emit the following signals:
//
// - gtk.Gesture::begin when the gesture is recognized.
//
// - gtk.Gesture::update, whenever an input event is processed.
//
// - gtk.Gesture::end when the gesture is no longer recognized.
//
//
// Event propagation
//
// In order to receive events, a gesture needs to set a propagation phase
// through gtk.EventController.SetPropagationPhase().
//
// In the capture phase, events are propagated from the toplevel down to the
// target widget, and gestures that are attached to containers above the widget
// get a chance to interact with the event before it reaches the target.
//
// In the bubble phase, events are propagated up from the target widget to the
// toplevel, and gestures that are attached to containers above the widget get a
// chance to interact with events that have not been handled yet.
//
//
// States of a sequence
//
// Whenever input interaction happens, a single event may trigger a cascade of
// GtkGestures, both across the parents of the widget receiving the event and in
// parallel within an individual widget. It is a responsibility of the widgets
// using those gestures to set the state of touch sequences accordingly in order
// to enable cooperation of gestures around the GdkEventSequences triggering
// those.
//
// Within a widget, gestures can be grouped through gtk.Gesture.Group(). Grouped
// gestures synchronize the state of sequences, so calling
// gtk.Gesture.SetSequenceState() on one will effectively propagate the state
// throughout the group.
//
// By default, all sequences start out in the GTK_EVENT_SEQUENCE_NONE state,
// sequences in this state trigger the gesture event handler, but event
// propagation will continue unstopped by gestures.
//
// If a sequence enters into the GTK_EVENT_SEQUENCE_DENIED state, the gesture
// group will effectively ignore the sequence, letting events go unstopped
// through the gesture, but the "slot" will still remain occupied while the
// touch is active.
//
// If a sequence enters in the GTK_EVENT_SEQUENCE_CLAIMED state, the gesture
// group will grab all interaction on the sequence, by:
//
// - Setting the same sequence to GTK_EVENT_SEQUENCE_DENIED on every other
// gesture group within the widget, and every gesture on parent widgets in the
// propagation chain.
//
// - Emitting gtk.Gesture::cancel on every gesture in widgets underneath in the
// propagation chain.
//
// - Stopping event propagation after the gesture group handles the event.
//
// Note: if a sequence is set early to GTK_EVENT_SEQUENCE_CLAIMED on
// GDK_TOUCH_BEGIN/GDK_BUTTON_PRESS (so those events are captured before
// reaching the event widget, this implies GTK_PHASE_CAPTURE), one similar event
// will emulated if the sequence changes to GTK_EVENT_SEQUENCE_DENIED. This way
// event coherence is preserved before event propagation is unstopped again.
//
// Sequence states can't be changed freely. See gtk.Gesture.SetSequenceState()
// to know about the possible lifetimes of a GdkEventSequence.
//
//
// Touchpad gestures
//
// On the platforms that support it, GtkGesture will handle transparently
// touchpad gesture events. The only precautions users of GtkGesture should do
// to enable this support are:
//
// - If the gesture has GTK_PHASE_NONE, ensuring events of type
// GDK_TOUCHPAD_SWIPE and GDK_TOUCHPAD_PINCH are handled by the GtkGesture.
type Gesture struct {
	EventController
}

var (
	_ EventControllerer = (*Gesture)(nil)
)

// Gesturer describes types inherited from class Gesture.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Gesturer interface {
	externglib.Objector
	baseGesture() *Gesture
}

var _ Gesturer = (*Gesture)(nil)

func wrapGesture(obj *externglib.Object) *Gesture {
	return &Gesture{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalGesturer(p uintptr) (interface{}, error) {
	return wrapGesture(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BoundingBox: if there are touch sequences being currently handled by gesture,
// returns TRUE and fills in rect with the bounding box containing all active
// touches.
//
// Otherwise, FALSE will be returned.
//
// Note: This function will yield unexpected results on touchpad gestures. Since
// there is no correlation between physical and pixel distances, these will look
// as if constrained in an infinitely small area, rect width and height will
// thus be 0 regardless of the number of touchpoints.
func (gesture *Gesture) BoundingBox() (*gdk.Rectangle, bool) {
	var _arg0 *C.GtkGesture  // out
	var _arg1 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_bounding_box(_arg0, &_arg1)
	runtime.KeepAlive(gesture)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// BoundingBoxCenter: if there are touch sequences being currently handled by
// gesture, returns TRUE and fills in x and y with the center of the bounding
// box containing all active touches.
//
// Otherwise, FALSE will be returned.
func (gesture *Gesture) BoundingBoxCenter() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGesture // out
	var _arg1 C.double      // in
	var _arg2 C.double      // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_bounding_box_center(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg1)
	_y = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// Device returns the logical GdkDevice that is currently operating on gesture.
//
// This returns NULL if the gesture is not being interacted.
func (gesture *Gesture) Device() gdk.Devicer {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GdkDevice  // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_device(_arg0)
	runtime.KeepAlive(gesture)

	var _device gdk.Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(gdk.Devicer)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// GetGroup returns all gestures in the group of gesture.
func (gesture *Gesture) GetGroup() []Gesturer {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GList      // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_group(_arg0)
	runtime.KeepAlive(gesture)

	var _list []Gesturer // out

	_list = make([]Gesturer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkGesture)(v)
		var dst Gesturer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.Gesturer is nil")
			}

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Gesturer)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Gesturer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// LastEvent returns the last event that was processed for sequence.
//
// Note that the returned pointer is only valid as long as the sequence is still
// interpreted by the gesture. If in doubt, you should make a copy of the event.
//
// The function takes the following parameters:
//
//    - sequence: GdkEventSequence.
//
func (gesture *Gesture) LastEvent(sequence *gdk.EventSequence) gdk.Eventer {
	var _arg0 *C.GtkGesture       // out
	var _arg1 *C.GdkEventSequence // out
	var _cret *C.GdkEvent         // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	if sequence != nil {
		_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_cret = C.gtk_gesture_get_last_event(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _event gdk.Eventer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(gdk.Eventer)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gdk.Eventer")
			}
			_event = rv
		}
	}

	return _event
}

// LastUpdatedSequence returns the GdkEventSequence that was last updated on
// gesture.
func (gesture *Gesture) LastUpdatedSequence() *gdk.EventSequence {
	var _arg0 *C.GtkGesture       // out
	var _cret *C.GdkEventSequence // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_last_updated_sequence(_arg0)
	runtime.KeepAlive(gesture)

	var _eventSequence *gdk.EventSequence // out

	if _cret != nil {
		_eventSequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _eventSequence
}

// Point: if sequence is currently being interpreted by gesture, returns TRUE
// and fills in x and y with the last coordinates stored for that event
// sequence.
//
// The coordinates are always relative to the widget allocation.
//
// The function takes the following parameters:
//
//    - sequence: GdkEventSequence, or NULL for pointer events.
//
func (gesture *Gesture) Point(sequence *gdk.EventSequence) (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGesture       // out
	var _arg1 *C.GdkEventSequence // out
	var _arg2 C.double            // in
	var _arg3 C.double            // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	if sequence != nil {
		_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_cret = C.gtk_gesture_get_point(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg2)
	_y = float64(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// SequenceState returns the sequence state, as seen by gesture.
//
// The function takes the following parameters:
//
//    - sequence: EventSequence.
//
func (gesture *Gesture) SequenceState(sequence *gdk.EventSequence) EventSequenceState {
	var _arg0 *C.GtkGesture           // out
	var _arg1 *C.GdkEventSequence     // out
	var _cret C.GtkEventSequenceState // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))

	_cret = C.gtk_gesture_get_sequence_state(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _eventSequenceState EventSequenceState // out

	_eventSequenceState = EventSequenceState(_cret)

	return _eventSequenceState
}

// Sequences returns the list of GdkEventSequences currently being interpreted
// by gesture.
func (gesture *Gesture) Sequences() []*gdk.EventSequence {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GList      // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_sequences(_arg0)
	runtime.KeepAlive(gesture)

	var _list []*gdk.EventSequence // out

	_list = make([]*gdk.EventSequence, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkEventSequence)(v)
		var dst *gdk.EventSequence // out
		dst = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// Group adds gesture to the same group than group_gesture.
//
// Gestures are by default isolated in their own groups.
//
// Both gestures must have been added to the same widget before they can be
// grouped.
//
// When gestures are grouped, the state of GdkEventSequences is kept in sync for
// all of those, so calling gtk.Gesture.SetSequenceState(), on one will transfer
// the same value to the others.
//
// Groups also perform an "implicit grabbing" of sequences, if a
// GdkEventSequence state is set to GTK_EVENT_SEQUENCE_CLAIMED on one group,
// every other gesture group attached to the same GtkWidget will switch the
// state for that sequence to GTK_EVENT_SEQUENCE_DENIED.
//
// The function takes the following parameters:
//
//    - gesture: GtkGesture.
//
func (groupGesture *Gesture) Group(gesture Gesturer) {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GtkGesture // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(groupGesture.Native()))
	_arg1 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	C.gtk_gesture_group(_arg0, _arg1)
	runtime.KeepAlive(groupGesture)
	runtime.KeepAlive(gesture)
}

// HandlesSequence returns TRUE if gesture is currently handling events
// corresponding to sequence.
//
// The function takes the following parameters:
//
//    - sequence: GdkEventSequence or NULL.
//
func (gesture *Gesture) HandlesSequence(sequence *gdk.EventSequence) bool {
	var _arg0 *C.GtkGesture       // out
	var _arg1 *C.GdkEventSequence // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	if sequence != nil {
		_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_cret = C.gtk_gesture_handles_sequence(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsActive returns TRUE if the gesture is currently active.
//
// A gesture is active while there are touch sequences interacting with it.
func (gesture *Gesture) IsActive() bool {
	var _arg0 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_is_active(_arg0)
	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsGroupedWith returns TRUE if both gestures pertain to the same group.
//
// The function takes the following parameters:
//
//    - other GtkGesture.
//
func (gesture *Gesture) IsGroupedWith(other Gesturer) bool {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = (*C.GtkGesture)(unsafe.Pointer(other.Native()))

	_cret = C.gtk_gesture_is_grouped_with(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRecognized returns TRUE if the gesture is currently recognized.
//
// A gesture is recognized if there are as many interacting touch sequences as
// required by gesture.
func (gesture *Gesture) IsRecognized() bool {
	var _arg0 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_is_recognized(_arg0)
	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSequenceState sets the state of sequence in gesture.
//
// Sequences start in state GTK_EVENT_SEQUENCE_NONE, and whenever they change
// state, they can never go back to that state. Likewise, sequences in state
// GTK_EVENT_SEQUENCE_DENIED cannot turn back to a not denied state. With these
// rules, the lifetime of an event sequence is constrained to the next four:
//
// * None * None → Denied * None → Claimed * None → Claimed → Denied
//
// Note: Due to event handling ordering, it may be unsafe to set the state on
// another gesture within a gtk.Gesture::begin signal handler, as the callback
// might be executed before the other gesture knows about the sequence. A safe
// way to perform this could be:
//
//    static void
//    first_gesture_begin_cb (GtkGesture       *first_gesture,
//                            GdkEventSequence *sequence,
//                            gpointer          user_data)
//    {
//      gtk_gesture_set_sequence_state (first_gesture, sequence, GTK_EVENT_SEQUENCE_CLAIMED);
//      gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
//    }
//
//    static void
//    second_gesture_begin_cb (GtkGesture       *second_gesture,
//                             GdkEventSequence *sequence,
//                             gpointer          user_data)
//    {
//      if (gtk_gesture_get_sequence_state (first_gesture, sequence) == GTK_EVENT_SEQUENCE_CLAIMED)
//        gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
//    }
//
//
// If both gestures are in the same group, just set the state on the gesture
// emitting the event, the sequence will be already be initialized to the
// group's global state when the second gesture processes the event.
//
// The function takes the following parameters:
//
//    - sequence: GdkEventSequence.
//    - state: sequence state.
//
func (gesture *Gesture) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	var _arg0 *C.GtkGesture           // out
	var _arg1 *C.GdkEventSequence     // out
	var _arg2 C.GtkEventSequenceState // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))
	_arg2 = C.GtkEventSequenceState(state)

	_cret = C.gtk_gesture_set_sequence_state(_arg0, _arg1, _arg2)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)
	runtime.KeepAlive(state)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetState sets the state of all sequences that gesture is currently
// interacting with.
//
// See gtk.Gesture.SetSequenceState() for more details on sequence states.
//
// The function takes the following parameters:
//
//    - state: sequence state.
//
func (gesture *Gesture) SetState(state EventSequenceState) bool {
	var _arg0 *C.GtkGesture           // out
	var _arg1 C.GtkEventSequenceState // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.GtkEventSequenceState(state)

	_cret = C.gtk_gesture_set_state(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(state)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ungroup separates gesture into an isolated group.
func (gesture *Gesture) Ungroup() {
	var _arg0 *C.GtkGesture // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	C.gtk_gesture_ungroup(_arg0)
	runtime.KeepAlive(gesture)
}

func (gesture *Gesture) baseGesture() *Gesture {
	return gesture
}

// BaseGesture returns the underlying base object.
func BaseGesture(obj Gesturer) *Gesture {
	return obj.baseGesture()
}

// ConnectBegin: emitted when the gesture is recognized.
//
// This means the number of touch sequences matches gtk.Gesture:n-points.
//
// Note: These conditions may also happen when an extra touch (eg. a third touch
// on a 2-touches gesture) is lifted, in that situation sequence won't pertain
// to the current set of active touches, so don't rely on this being true.
func (gesture *Gesture) ConnectBegin(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("begin", f)
}

// ConnectCancel: emitted whenever a sequence is cancelled.
//
// This usually happens on active touches when gtk.EventController.Reset() is
// called on gesture (manually, due to grabs...), or the individual sequence was
// claimed by parent widgets' controllers (see gtk.Gesture.SetSequenceState()).
//
// gesture must forget everything about sequence as in response to this signal.
func (gesture *Gesture) ConnectCancel(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("cancel", f)
}

// ConnectEnd: emitted when gesture either stopped recognizing the event
// sequences as something to be handled, or the number of touch sequences became
// higher or lower than gtk.Gesture:n-points.
//
// Note: sequence might not pertain to the group of sequences that were
// previously triggering recognition on gesture (ie. a just pressed touch
// sequence that exceeds gtk.Gesture:n-points). This situation may be detected
// by checking through gtk.Gesture.HandlesSequence().
func (gesture *Gesture) ConnectEnd(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("end", f)
}

// ConnectSequenceStateChanged: emitted whenever a sequence state changes.
//
// See gtk.Gesture.SetSequenceState() to know more about the expectable sequence
// lifetimes.
func (gesture *Gesture) ConnectSequenceStateChanged(f func(sequence *gdk.EventSequence, state EventSequenceState)) externglib.SignalHandle {
	return gesture.Connect("sequence-state-changed", f)
}

// ConnectUpdate: emitted whenever an event is handled while the gesture is
// recognized.
//
// sequence is guaranteed to pertain to the set of active touches.
func (gesture *Gesture) ConnectUpdate(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("update", f)
}
