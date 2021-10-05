// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_toggle_action_get_type()), F: marshalToggleActioner},
	})
}

// ToggleActionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToggleActionOverrider interface {
	// Toggled emits the “toggled” signal on the toggle action.
	//
	// Deprecated: since version 3.10.
	Toggled()
}

// ToggleAction corresponds roughly to a CheckMenuItem. It has an “active” state
// specifying whether the action has been checked or not.
type ToggleAction struct {
	Action
}

func wrapToggleAction(obj *externglib.Object) *ToggleAction {
	return &ToggleAction{
		Action: Action{
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalToggleActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToggleAction(obj), nil
}

// NewToggleAction creates a new ToggleAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewToggleAction(name, label, tooltip, stockId string) *ToggleAction {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _arg3 *C.gchar           // out
	var _arg4 *C.gchar           // out
	var _cret *C.GtkToggleAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if tooltip != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	if stockId != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
		defer C.free(unsafe.Pointer(_arg4))
	}

	_cret = C.gtk_toggle_action_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(stockId)

	var _toggleAction *ToggleAction // out

	_toggleAction = wrapToggleAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _toggleAction
}

// Active returns the checked state of the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) Active() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_toggle_action_get_active(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DrawAsRadio returns whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) DrawAsRadio() bool {
	var _arg0 *C.GtkToggleAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_toggle_action_get_draw_as_radio(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the checked state on the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_active(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(isActive)
}

// SetDrawAsRadio sets whether the action should have proxies like a radio
// action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	var _arg0 *C.GtkToggleAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))
	if drawAsRadio {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_action_set_draw_as_radio(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(drawAsRadio)
}

// Toggled emits the “toggled” signal on the toggle action.
//
// Deprecated: since version 3.10.
func (action *ToggleAction) Toggled() {
	var _arg0 *C.GtkToggleAction // out

	_arg0 = (*C.GtkToggleAction)(unsafe.Pointer(action.Native()))

	C.gtk_toggle_action_toggled(_arg0)
	runtime.KeepAlive(action)
}

// ConnectToggled: should be connected if you wish to perform an action whenever
// the ToggleAction state is changed.
func (t *ToggleAction) ConnectToggled(f func()) glib.SignalHandle {
	return t.Connect("toggled", f)
}
