// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
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
		{T: externglib.Type(C.gtk_radio_action_get_type()), F: marshalRadioActioner},
	})
}

// RadioActionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RadioActionOverrider interface {
	Changed(current RadioActioner)
}

// RadioActioner describes RadioAction's methods.
type RadioActioner interface {
	// CurrentValue obtains the value property of the currently active member of
	// the group to which @action belongs.
	CurrentValue() int
	// JoinGroup joins a radio action object to the group of another radio
	// action object.
	JoinGroup(groupSource RadioActioner)
	// SetCurrentValue sets the currently active group member to the member with
	// value property @current_value.
	SetCurrentValue(currentValue int)
}

// RadioAction is similar to RadioMenuItem. A number of radio actions can be
// linked together so that only one may be active at any one time.
type RadioAction struct {
	ToggleAction
}

var (
	_ RadioActioner   = (*RadioAction)(nil)
	_ gextras.Nativer = (*RadioAction)(nil)
)

func wrapRadioAction(obj *externglib.Object) RadioActioner {
	return &RadioAction{
		ToggleAction: ToggleAction{
			Action: Action{
				Object: obj,
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRadioAction(obj), nil
}

// NewRadioAction creates a new RadioAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewRadioAction(name string, label string, tooltip string, stockId string, value int) *RadioAction {
	var _arg1 *C.gchar          // out
	var _arg2 *C.gchar          // out
	var _arg3 *C.gchar          // out
	var _arg4 *C.gchar          // out
	var _arg5 C.gint            // out
	var _cret *C.GtkRadioAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.gint(value)

	_cret = C.gtk_radio_action_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _radioAction *RadioAction // out

	_radioAction = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*RadioAction)

	return _radioAction
}

// CurrentValue obtains the value property of the currently active member of the
// group to which @action belongs.
//
// Deprecated: since version 3.10.
func (action *RadioAction) CurrentValue() int {
	var _arg0 *C.GtkRadioAction // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_radio_action_get_current_value(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// JoinGroup joins a radio action object to the group of another radio action
// object.
//
// Use this in language bindings instead of the gtk_radio_action_get_group() and
// gtk_radio_action_set_group() methods
//
// A common way to set up a group of radio actions is the following:
//
//     GtkRadioAction *action;
//     GtkRadioAction *last_action;
//
//     while ( ...more actions to add... /)
//       {
//          action = gtk_radio_action_new (...);
//
//          gtk_radio_action_join_group (action, last_action);
//          last_action = action;
//       }
//
// Deprecated: since version 3.10.
func (action *RadioAction) JoinGroup(groupSource RadioActioner) {
	var _arg0 *C.GtkRadioAction // out
	var _arg1 *C.GtkRadioAction // out

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.GtkRadioAction)(unsafe.Pointer((groupSource).(gextras.Nativer).Native()))

	C.gtk_radio_action_join_group(_arg0, _arg1)
}

// SetCurrentValue sets the currently active group member to the member with
// value property @current_value.
//
// Deprecated: since version 3.10.
func (action *RadioAction) SetCurrentValue(currentValue int) {
	var _arg0 *C.GtkRadioAction // out
	var _arg1 C.gint            // out

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.gint(currentValue)

	C.gtk_radio_action_set_current_value(_arg0, _arg1)
}
