// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeCoordType    = coreglib.Type(C.atk_coord_type_get_type())
	GTypeKeyEventType = coreglib.Type(C.atk_key_event_type_get_type())
	GTypeUtil         = coreglib.Type(C.atk_util_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCoordType, F: marshalCoordType},
		coreglib.TypeMarshaler{T: GTypeKeyEventType, F: marshalKeyEventType},
		coreglib.TypeMarshaler{T: GTypeUtil, F: marshalUtil},
	})
}

// CoordType specifies how xy coordinates are to be interpreted. Used by
// functions such as atk_component_get_position() and
// atk_text_get_character_extents().
type CoordType C.gint

const (
	// XYScreen specifies xy coordinates relative to the screen.
	XYScreen CoordType = iota
	// XYWindow specifies xy coordinates relative to the widget's top-level
	// window.
	XYWindow
	// XYParent specifies xy coordinates relative to the widget's immediate
	// parent. Since: 2.30.
	XYParent
)

func marshalCoordType(p uintptr) (interface{}, error) {
	return CoordType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CoordType.
func (c CoordType) String() string {
	switch c {
	case XYScreen:
		return "Screen"
	case XYWindow:
		return "Window"
	case XYParent:
		return "Parent"
	default:
		return fmt.Sprintf("CoordType(%d)", c)
	}
}

// KeyEventType specifies the type of a keyboard evemt.
type KeyEventType C.gint

const (
	// KeyEventPress specifies a key press event.
	KeyEventPress KeyEventType = iota
	// KeyEventRelease specifies a key release event.
	KeyEventRelease
	// KeyEventLastDefined: not a valid value; specifies end of enumeration.
	KeyEventLastDefined
)

func marshalKeyEventType(p uintptr) (interface{}, error) {
	return KeyEventType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for KeyEventType.
func (k KeyEventType) String() string {
	switch k {
	case KeyEventPress:
		return "Press"
	case KeyEventRelease:
		return "Release"
	case KeyEventLastDefined:
		return "LastDefined"
	default:
		return fmt.Sprintf("KeyEventType(%d)", k)
	}
}

// KeySnoopFunc is a type of callback which is called whenever a key event
// occurs, if registered via atk_add_key_event_listener. It allows for
// pre-emptive interception of key events via the return code as described
// below.
type KeySnoopFunc func(event *KeyEventStruct) (gint int)

// FocusTrackerNotify: cause the focus tracker functions which have been
// specified to be executed for the object.
//
// Deprecated: Focus tracking has been dropped as a feature to be implemented by
// ATK itself. As Object::focus-event was deprecated in favor of a
// Object::state-change signal, in order to notify a focus change on your
// implementation, you can use atk_object_notify_state_change() instead.
//
// The function takes the following parameters:
//
//    - object: Object.
//
func FocusTrackerNotify(object *AtkObject) {
	var _arg1 *C.AtkObject // out

	_arg1 = (*C.AtkObject)(unsafe.Pointer(coreglib.InternObject(object).Native()))

	C.atk_focus_tracker_notify(_arg1)
	runtime.KeepAlive(object)
}

// GetRoot gets the root accessible container for the current application.
//
// The function returns the following values:
//
//    - object: root accessible container for the current application.
//
func GetRoot() *AtkObject {
	var _cret *C.AtkObject // in

	_cret = C.atk_get_root()

	var _object *AtkObject // out

	_object = wrapObject(coreglib.Take(unsafe.Pointer(_cret)))

	return _object
}

// GetToolkitName gets name string for the GUI toolkit implementing ATK for this
// application.
//
// The function returns the following values:
//
//    - utf8: name string for the GUI toolkit implementing ATK for this
//      application.
//
func GetToolkitName() string {
	var _cret *C.gchar // in

	_cret = C.atk_get_toolkit_name()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetToolkitVersion gets version string for the GUI toolkit implementing ATK
// for this application.
//
// The function returns the following values:
//
//    - utf8: version string for the GUI toolkit implementing ATK for this
//      application.
//
func GetToolkitVersion() string {
	var _cret *C.gchar // in

	_cret = C.atk_get_toolkit_version()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// RemoveFocusTracker removes the specified focus tracker from the list of
// functions to be called when any object receives focus.
//
// Deprecated: Focus tracking has been dropped as a feature to be implemented by
// ATK itself. If you need focus tracking on your implementation, subscribe to
// the Object::state-change "focused" signal.
//
// The function takes the following parameters:
//
//    - trackerId: id of the focus tracker to remove.
//
func RemoveFocusTracker(trackerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(trackerId)

	C.atk_remove_focus_tracker(_arg1)
	runtime.KeepAlive(trackerId)
}

// RemoveGlobalEventListener: listener_id is the value returned by
// #atk_add_global_event_listener when you registered that event listener.
//
// Toolkit implementor note: ATK provides a default implementation for this
// virtual method. ATK implementors are discouraged from reimplementing this
// method.
//
// Toolkit implementor note: this method is not intended to be used by ATK
// implementors but by ATK consumers.
//
// Removes the specified event listener.
//
// The function takes the following parameters:
//
//    - listenerId: id of the event listener to remove.
//
func RemoveGlobalEventListener(listenerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(listenerId)

	C.atk_remove_global_event_listener(_arg1)
	runtime.KeepAlive(listenerId)
}

// RemoveKeyEventListener: listener_id is the value returned by
// #atk_add_key_event_listener when you registered that event listener.
//
// Removes the specified event listener.
//
// The function takes the following parameters:
//
//    - listenerId: id of the event listener to remove.
//
func RemoveKeyEventListener(listenerId uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(listenerId)

	C.atk_remove_key_event_listener(_arg1)
	runtime.KeepAlive(listenerId)
}

// UtilOverrides contains methods that are overridable.
type UtilOverrides struct {
}

func defaultUtilOverrides(v *Util) UtilOverrides {
	return UtilOverrides{}
}

// Util: set of ATK utility functions which are used to support event
// registration of various types, and obtaining the 'root' accessible of a
// process and information about the current ATK implementation and toolkit
// version.
type Util struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Util)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Util, *UtilClass, UtilOverrides](
		GTypeUtil,
		initUtilClass,
		wrapUtil,
		defaultUtilOverrides,
	)
}

func initUtilClass(gclass unsafe.Pointer, overrides UtilOverrides, classInitFunc func(*UtilClass)) {
	if classInitFunc != nil {
		class := (*UtilClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapUtil(obj *coreglib.Object) *Util {
	return &Util{
		Object: obj,
	}
}

func marshalUtil(p uintptr) (interface{}, error) {
	return wrapUtil(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// KeyEventStruct encapsulates information about a key event.
//
// An instance of this type is always passed by reference.
type KeyEventStruct struct {
	*keyEventStruct
}

// keyEventStruct is the struct that's finalized.
type keyEventStruct struct {
	native *C.AtkKeyEventStruct
}

// Type: atkKeyEventType, generally one of ATK_KEY_EVENT_PRESS or
// ATK_KEY_EVENT_RELEASE.
func (k *KeyEventStruct) Type() int {
	valptr := &k.native._type
	var _v int // out
	_v = int(*valptr)
	return _v
}

// State: bitmask representing the state of the modifier keys immediately after
// the event takes place. The meaning of the bits is currently defined to match
// the bitmask used by GDK in GdkEventType.state, see
// http://developer.gnome.org/doc/API/2.0/gdk/gdk-Event-Structures.htmlEventKey.
func (k *KeyEventStruct) State() uint {
	valptr := &k.native.state
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Keyval: guint representing a keysym value corresponding to those used by GDK
// and X11: see /usr/X11/include/keysymdef.h.
func (k *KeyEventStruct) Keyval() uint {
	valptr := &k.native.keyval
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Length: length of member #string.
func (k *KeyEventStruct) Length() int {
	valptr := &k.native.length
	var _v int // out
	_v = int(*valptr)
	return _v
}

// String: string containing one of the following: either a string approximating
// the text that would result from this keypress, if the key is a control or
// graphic character, or a symbolic name for this keypress. Alphanumeric and
// printable keys will have the symbolic key name in this string member, for
// instance "A". "0", "semicolon", "aacute". Keypad keys have the prefix "KP".
func (k *KeyEventStruct) String() string {
	valptr := &k.native.string
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}

// Keycode: raw hardware code that generated the key event. This field is raraly
// useful.
func (k *KeyEventStruct) Keycode() uint16 {
	valptr := &k.native.keycode
	var _v uint16 // out
	_v = uint16(*valptr)
	return _v
}

// Timestamp: timestamp in milliseconds indicating when the event occurred.
// These timestamps are relative to a starting point which should be considered
// arbitrary, and only used to compare the dispatch times of events to one
// another.
func (k *KeyEventStruct) Timestamp() uint32 {
	valptr := &k.native.timestamp
	var _v uint32 // out
	_v = uint32(*valptr)
	return _v
}

// Type: atkKeyEventType, generally one of ATK_KEY_EVENT_PRESS or
// ATK_KEY_EVENT_RELEASE.
func (k *KeyEventStruct) SetType(typ int) {
	valptr := &k.native._type
	*valptr = C.gint(typ)
}

// State: bitmask representing the state of the modifier keys immediately after
// the event takes place. The meaning of the bits is currently defined to match
// the bitmask used by GDK in GdkEventType.state, see
// http://developer.gnome.org/doc/API/2.0/gdk/gdk-Event-Structures.htmlEventKey.
func (k *KeyEventStruct) SetState(state uint) {
	valptr := &k.native.state
	*valptr = C.guint(state)
}

// Keyval: guint representing a keysym value corresponding to those used by GDK
// and X11: see /usr/X11/include/keysymdef.h.
func (k *KeyEventStruct) SetKeyval(keyval uint) {
	valptr := &k.native.keyval
	*valptr = C.guint(keyval)
}

// Length: length of member #string.
func (k *KeyEventStruct) SetLength(length int) {
	valptr := &k.native.length
	*valptr = C.gint(length)
}

// Keycode: raw hardware code that generated the key event. This field is raraly
// useful.
func (k *KeyEventStruct) SetKeycode(keycode uint16) {
	valptr := &k.native.keycode
	*valptr = C.guint16(keycode)
}

// Timestamp: timestamp in milliseconds indicating when the event occurred.
// These timestamps are relative to a starting point which should be considered
// arbitrary, and only used to compare the dispatch times of events to one
// another.
func (k *KeyEventStruct) SetTimestamp(timestamp uint32) {
	valptr := &k.native.timestamp
	*valptr = C.guint32(timestamp)
}

// UtilClass: instance of this type is always passed by reference.
type UtilClass struct {
	*utilClass
}

// utilClass is the struct that's finalized.
type utilClass struct {
	native *C.AtkUtilClass
}
