// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_shortcut_get_type()), F: marshalShortcut},
	})
}

// Shortcut: `GtkShortcut` describes a keyboard shortcut.
//
// It contains a description of how to trigger the shortcut via a
// [class@Gtk.ShortcutTrigger] and a way to activate the shortcut on a widget
// via a [class@Gtk.ShortcutAction].
//
// The actual work is usually done via [class@Gtk.ShortcutController], which
// decides if and when to activate a shortcut. Using that controller directly
// however is rarely necessary as various higher level convenience APIs exist on
// Widgets that make it easier to use shortcuts in GTK.
//
// `GtkShortcut` does provide functionality to make it easy for users to work
// with shortcuts, either by providing informational strings for display
// purposes or by allowing shortcuts to be configured.
type Shortcut interface {
	gextras.Objector

	// Action gets the action that is activated by this shortcut.
	Action() ShortcutAction
	// Arguments gets the arguments that are passed when activating the
	// shortcut.
	Arguments() *glib.Variant
	// Trigger gets the trigger used to trigger @self.
	Trigger() ShortcutTrigger
	// SetAction sets the new action for @self to be @action.
	SetAction(action ShortcutAction)
	// SetArguments sets the arguments to pass when activating the shortcut.
	SetArguments(args *glib.Variant)
	// SetTrigger sets the new trigger for @self to be @trigger.
	SetTrigger(trigger ShortcutTrigger)
}

// ShortcutClass implements the Shortcut interface.
type ShortcutClass struct {
	*externglib.Object
}

var _ Shortcut = (*ShortcutClass)(nil)

func wrapShortcut(obj *externglib.Object) Shortcut {
	return &ShortcutClass{
		Object: obj,
	}
}

func marshalShortcut(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcut(obj), nil
}

// NewShortcut creates a new `GtkShortcut` that is triggered by @trigger and
// then activates @action.
func NewShortcut(trigger ShortcutTrigger, action ShortcutAction) Shortcut {
	var _arg1 *C.GtkShortcutTrigger // out
	var _arg2 *C.GtkShortcutAction  // out
	var _cret *C.GtkShortcut        // in

	_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(trigger.Native()))
	_arg2 = (*C.GtkShortcutAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_shortcut_new(_arg1, _arg2)

	var _shortcut Shortcut // out

	_shortcut = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Shortcut)

	return _shortcut
}

// Action gets the action that is activated by this shortcut.
func (s *ShortcutClass) Action() ShortcutAction {
	var _arg0 *C.GtkShortcut       // out
	var _cret *C.GtkShortcutAction // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_shortcut_get_action(_arg0)

	var _shortcutAction ShortcutAction // out

	_shortcutAction = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ShortcutAction)

	return _shortcutAction
}

// Arguments gets the arguments that are passed when activating the shortcut.
func (s *ShortcutClass) Arguments() *glib.Variant {
	var _arg0 *C.GtkShortcut // out
	var _cret *C.GVariant    // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_shortcut_get_arguments(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// Trigger gets the trigger used to trigger @self.
func (s *ShortcutClass) Trigger() ShortcutTrigger {
	var _arg0 *C.GtkShortcut        // out
	var _cret *C.GtkShortcutTrigger // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_shortcut_get_trigger(_arg0)

	var _shortcutTrigger ShortcutTrigger // out

	_shortcutTrigger = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ShortcutTrigger)

	return _shortcutTrigger
}

// SetAction sets the new action for @self to be @action.
func (s *ShortcutClass) SetAction(action ShortcutAction) {
	var _arg0 *C.GtkShortcut       // out
	var _arg1 *C.GtkShortcutAction // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkShortcutAction)(unsafe.Pointer(action.Native()))

	C.gtk_shortcut_set_action(_arg0, _arg1)
}

// SetArguments sets the arguments to pass when activating the shortcut.
func (s *ShortcutClass) SetArguments(args *glib.Variant) {
	var _arg0 *C.GtkShortcut // out
	var _arg1 *C.GVariant    // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GVariant)(unsafe.Pointer(args))

	C.gtk_shortcut_set_arguments(_arg0, _arg1)
}

// SetTrigger sets the new trigger for @self to be @trigger.
func (s *ShortcutClass) SetTrigger(trigger ShortcutTrigger) {
	var _arg0 *C.GtkShortcut        // out
	var _arg1 *C.GtkShortcutTrigger // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(trigger.Native()))

	C.gtk_shortcut_set_trigger(_arg0, _arg1)
}
