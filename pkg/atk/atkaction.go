// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_action_get_type()), F: marshalActioner},
	})
}

// ActionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionOverrider interface {
	// DoAction: perform the specified action on the object.
	DoAction(i int) bool
	// Description returns a description of the specified action of the object.
	Description(i int) string
	// Keybinding gets the keybinding which can be used to activate this action,
	// if one exists. The string returned should contain localized,
	// human-readable, key sequences as they would appear when displayed on
	// screen. It must be in the format "mnemonic;sequence;shortcut".
	//
	// - The mnemonic key activates the object if it is presently enabled
	// onscreen. This typically corresponds to the underlined letter within the
	// widget. Example: "n" in a traditional "New..." menu item or the "a" in
	// "Apply" for a button.
	//
	// - The sequence is the full list of keys which invoke the action even if
	// the relevant element is not currently shown on screen. For instance, for
	// a menu item the sequence is the keybindings used to open the parent menus
	// before invoking. The sequence string is colon-delimited. Example:
	// "Alt+F:N" in a traditional "New..." menu item.
	//
	// - The shortcut, if it exists, will invoke the same action without showing
	// the component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
	// traditional "New..." menu item.
	//
	// Example: For a traditional "New..." menu item, the expected return value
	// would be: "N;Alt+F:N;Ctrl+N" for the English locale and
	// "N;Alt+D:N;Strg+N" for the German locale. If, hypothetically, this menu
	// item lacked a mnemonic, it would be represented by ";;Ctrl+N" and
	// ";;Strg+N" respectively.
	Keybinding(i int) string
	// LocalizedName returns the localized name of the specified action of the
	// object.
	LocalizedName(i int) string
	// NActions gets the number of accessible actions available on the object.
	// If there are more than one, the first one is considered the "default"
	// action of the object.
	NActions() int
	// Name returns a non-localized string naming the specified action of the
	// object. This name is generally not descriptive of the end result of the
	// action, but instead names the 'interaction type' which the object
	// supports. By convention, the above strings should be used to represent
	// the actions which correspond to the common point-and-click interaction
	// techniques of the same name: i.e. "click", "press", "release", "drag",
	// "drop", "popup", etc. The "popup" action should be used to pop up a
	// context menu for the object, if one exists.
	//
	// For technical reasons, some toolkits cannot guarantee that the reported
	// action is actually 'bound' to a nontrivial user event; i.e. the result of
	// some actions via atk_action_do_action() may be NIL.
	Name(i int) string
	// SetDescription sets a description of the specified action of the object.
	SetDescription(i int, desc string) bool
}

// Action should be implemented by instances of Object classes with which the
// user can interact directly, i.e. buttons, checkboxes, scrollbars, e.g.
// components which are not "passive" providers of UI information.
//
// Exceptions: when the user interaction is already covered by another
// appropriate interface such as EditableText (insert/delete text, etc.) or
// Value (set value) then these actions should not be exposed by Action as well.
//
// Though most UI interactions on components should be invocable via keyboard as
// well as mouse, there will generally be a close mapping between "mouse
// actions" that are possible on a component and the AtkActions. Where mouse and
// keyboard actions are redundant in effect, Action should expose only one
// action rather than exposing redundant actions if possible. By convention we
// have been using "mouse centric" terminology for Action names.
type Action struct {
	*externglib.Object
}

// Actioner describes Action's abstract methods.
type Actioner interface {
	externglib.Objector

	// DoAction: perform the specified action on the object.
	DoAction(i int) bool
	// Description returns a description of the specified action of the object.
	Description(i int) string
	// Keybinding gets the keybinding which can be used to activate this action,
	// if one exists.
	Keybinding(i int) string
	// LocalizedName returns the localized name of the specified action of the
	// object.
	LocalizedName(i int) string
	// NActions gets the number of accessible actions available on the object.
	NActions() int
	// Name returns a non-localized string naming the specified action of the
	// object.
	Name(i int) string
	// SetDescription sets a description of the specified action of the object.
	SetDescription(i int, desc string) bool
}

var _ Actioner = (*Action)(nil)

func wrapAction(obj *externglib.Object) *Action {
	return &Action{
		Object: obj,
	}
}

func marshalActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAction(obj), nil
}

// DoAction: perform the specified action on the object.
func (action *Action) DoAction(i int) bool {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_do_action(_arg0, _arg1)

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Description returns a description of the specified action of the object.
func (action *Action) Description(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_description(_arg0, _arg1)

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Keybinding gets the keybinding which can be used to activate this action, if
// one exists. The string returned should contain localized, human-readable, key
// sequences as they would appear when displayed on screen. It must be in the
// format "mnemonic;sequence;shortcut".
//
// - The mnemonic key activates the object if it is presently enabled onscreen.
// This typically corresponds to the underlined letter within the widget.
// Example: "n" in a traditional "New..." menu item or the "a" in "Apply" for a
// button.
//
// - The sequence is the full list of keys which invoke the action even if the
// relevant element is not currently shown on screen. For instance, for a menu
// item the sequence is the keybindings used to open the parent menus before
// invoking. The sequence string is colon-delimited. Example: "Alt+F:N" in a
// traditional "New..." menu item.
//
// - The shortcut, if it exists, will invoke the same action without showing the
// component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
// traditional "New..." menu item.
//
// Example: For a traditional "New..." menu item, the expected return value
// would be: "N;Alt+F:N;Ctrl+N" for the English locale and "N;Alt+D:N;Strg+N"
// for the German locale. If, hypothetically, this menu item lacked a mnemonic,
// it would be represented by ";;Ctrl+N" and ";;Strg+N" respectively.
func (action *Action) Keybinding(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_keybinding(_arg0, _arg1)

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LocalizedName returns the localized name of the specified action of the
// object.
func (action *Action) LocalizedName(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_localized_name(_arg0, _arg1)

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NActions gets the number of accessible actions available on the object. If
// there are more than one, the first one is considered the "default" action of
// the object.
func (action *Action) NActions() int {
	var _arg0 *C.AtkAction // out
	var _cret C.gint       // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.atk_action_get_n_actions(_arg0)

	runtime.KeepAlive(action)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name returns a non-localized string naming the specified action of the
// object. This name is generally not descriptive of the end result of the
// action, but instead names the 'interaction type' which the object supports.
// By convention, the above strings should be used to represent the actions
// which correspond to the common point-and-click interaction techniques of the
// same name: i.e. "click", "press", "release", "drag", "drop", "popup", etc.
// The "popup" action should be used to pop up a context menu for the object, if
// one exists.
//
// For technical reasons, some toolkits cannot guarantee that the reported
// action is actually 'bound' to a nontrivial user event; i.e. the result of
// some actions via atk_action_do_action() may be NIL.
func (action *Action) Name(i int) string {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_action_get_name(_arg0, _arg1)

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetDescription sets a description of the specified action of the object.
func (action *Action) SetDescription(i int, desc string) bool {
	var _arg0 *C.AtkAction // out
	var _arg1 C.gint       // out
	var _arg2 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.gint(i)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(desc)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.atk_action_set_description(_arg0, _arg1, _arg2)

	runtime.KeepAlive(action)
	runtime.KeepAlive(i)
	runtime.KeepAlive(desc)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
