// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_SimpleAction_ConnectChangeState(gpointer, GVariant*, guintptr);
// extern void _gotk4_gio2_SimpleAction_ConnectActivate(gpointer, GVariant*, guintptr);
import "C"

// GType values.
var (
	GTypeDBusActionGroup  = coreglib.Type(C.g_dbus_action_group_get_type())
	GTypeDBusMenuModel    = coreglib.Type(C.g_dbus_menu_model_get_type())
	GTypeSimpleAction     = coreglib.Type(C.g_simple_action_get_type())
	GTypeSimplePermission = coreglib.Type(C.g_simple_permission_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDBusActionGroup, F: marshalDBusActionGroup},
		coreglib.TypeMarshaler{T: GTypeDBusMenuModel, F: marshalDBusMenuModel},
		coreglib.TypeMarshaler{T: GTypeSimpleAction, F: marshalSimpleAction},
		coreglib.TypeMarshaler{T: GTypeSimplePermission, F: marshalSimplePermission},
	})
}

// IOErrorQuark gets the GIO Error Quark.
//
// The function returns the following values:
//
//   - quark: #GQuark.
//
func IOErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.g_io_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

// DBusActionGroup is an implementation of the Group interface that can
// be used as a proxy for an action group that is exported over D-Bus with
// g_dbus_connection_export_action_group().
type DBusActionGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object

	RemoteActionGroup
}

var (
	_ coreglib.Objector = (*DBusActionGroup)(nil)
)

func wrapDBusActionGroup(obj *coreglib.Object) *DBusActionGroup {
	return &DBusActionGroup{
		Object: obj,
		RemoteActionGroup: RemoteActionGroup{
			ActionGroup: ActionGroup{
				Object: obj,
			},
		},
	}
}

func marshalDBusActionGroup(p uintptr) (interface{}, error) {
	return wrapDBusActionGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusMenuModel is an implementation of Model that can be used
// as a proxy for a menu model that is exported over D-Bus with
// g_dbus_connection_export_menu_model().
type DBusMenuModel struct {
	_ [0]func() // equal guard
	MenuModel
}

var (
	_ MenuModeller = (*DBusMenuModel)(nil)
)

func wrapDBusMenuModel(obj *coreglib.Object) *DBusMenuModel {
	return &DBusMenuModel{
		MenuModel: MenuModel{
			Object: obj,
		},
	}
}

func marshalDBusMenuModel(p uintptr) (interface{}, error) {
	return wrapDBusMenuModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SimpleAction is the obvious simple implementation of the #GAction interface.
// This is the easiest way to create an action for purposes of adding it to a
// ActionGroup.
//
// See also Action.
type SimpleAction struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Action
}

var (
	_ coreglib.Objector = (*SimpleAction)(nil)
)

func wrapSimpleAction(obj *coreglib.Object) *SimpleAction {
	return &SimpleAction{
		Object: obj,
		Action: Action{
			Object: obj,
		},
	}
}

func marshalSimpleAction(p uintptr) (interface{}, error) {
	return wrapSimpleAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate indicates that the action was just activated.
//
// parameter will always be of the expected type, i.e. the parameter type
// specified when the action was created. If an incorrect type is given when
// activating the action, this signal is not emitted.
//
// Since GLib 2.40, if no handler is connected to this signal then the default
// behaviour for boolean-stated actions with a NULL parameter type is to toggle
// them via the Action::change-state signal. For stateful actions where the
// state type is equal to the parameter type, the default is to forward them
// directly to Action::change-state. This should allow almost all users of
// Action to connect only one handler or the other.
func (simple *SimpleAction) ConnectActivate(f func(parameter *glib.Variant)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(simple, "activate", false, unsafe.Pointer(C._gotk4_gio2_SimpleAction_ConnectActivate), f)
}

// ConnectChangeState indicates that the action just received a request to
// change its state.
//
// value will always be of the correct state type, i.e. the type of the initial
// state passed to g_simple_action_new_stateful(). If an incorrect type is given
// when requesting to change the state, this signal is not emitted.
//
// If no handler is connected to this signal then the default behaviour is to
// call g_simple_action_set_state() to set the state to the requested value.
// If you connect a signal handler then no default action is taken. If the
// state should change then you must call g_simple_action_set_state() from the
// handler.
//
// An example of a 'change-state' handler:
//
//    static void
//    change_volume_state (GSimpleAction *action,
//                         GVariant      *value,
//                         gpointer       user_data)
//    {
//      gint requested;
//
//      requested = g_variant_get_int32 (value);
//
//      // Volume only goes from 0 to 10
//      if (0 <= requested && requested <= 10)
//        g_simple_action_set_state (action, value);
//    }
//
// The handler need not set the state to the requested value. It could set it to
// any value at all, or take some other action.
func (simple *SimpleAction) ConnectChangeState(f func(value *glib.Variant)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(simple, "change-state", false, unsafe.Pointer(C._gotk4_gio2_SimpleAction_ConnectChangeState), f)
}

// NewSimpleAction creates a new action.
//
// The created action is stateless. See g_simple_action_new_stateful() to create
// an action that has state.
//
// The function takes the following parameters:
//
//   - name of the action.
//   - parameterType (optional): type of parameter that will be passed to
//     handlers for the Action::activate signal, or NULL for no parameter.
//
// The function returns the following values:
//
//   - simpleAction: new Action.
//
func NewSimpleAction(name string, parameterType *glib.VariantType) *SimpleAction {
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariantType  // out
	var _cret *C.GSimpleAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameterType != nil {
		_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(parameterType)))
	}

	_cret = C.g_simple_action_new(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(parameterType)

	var _simpleAction *SimpleAction // out

	_simpleAction = wrapSimpleAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleAction
}

// NewSimpleActionStateful creates a new stateful action.
//
// All future state values must have the same Type as the initial state.
//
// If the state #GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//   - name of the action.
//   - parameterType (optional): type of the parameter that will be passed to
//     handlers for the Action::activate signal, or NULL for no parameter.
//   - state: initial state of the action.
//
// The function returns the following values:
//
//   - simpleAction: new Action.
//
func NewSimpleActionStateful(name string, parameterType *glib.VariantType, state *glib.Variant) *SimpleAction {
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariantType  // out
	var _arg3 *C.GVariant      // out
	var _cret *C.GSimpleAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameterType != nil {
		_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(parameterType)))
	}
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(state)))

	_cret = C.g_simple_action_new_stateful(_arg1, _arg2, _arg3)
	runtime.KeepAlive(name)
	runtime.KeepAlive(parameterType)
	runtime.KeepAlive(state)

	var _simpleAction *SimpleAction // out

	_simpleAction = wrapSimpleAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleAction
}

// SetEnabled sets the action as enabled or not.
//
// An action must be enabled in order to be activated or in order to have its
// state changed from outside callers.
//
// This should only be called by the implementor of the action. Users of the
// action should not attempt to modify its enabled flag.
//
// The function takes the following parameters:
//
//   - enabled: whether the action is enabled.
//
func (simple *SimpleAction) SetEnabled(enabled bool) {
	var _arg0 *C.GSimpleAction // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GSimpleAction)(unsafe.Pointer(coreglib.InternObject(simple).Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.g_simple_action_set_enabled(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(enabled)
}

// SetState sets the state of the action.
//
// This directly updates the 'state' property to the given value.
//
// This should only be called by the implementor of the action. Users of the
// action should not attempt to directly modify the 'state' property. Instead,
// they should call g_action_change_state() to request the change.
//
// If the value GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//   - value: new #GVariant for the state.
//
func (simple *SimpleAction) SetState(value *glib.Variant) {
	var _arg0 *C.GSimpleAction // out
	var _arg1 *C.GVariant      // out

	_arg0 = (*C.GSimpleAction)(unsafe.Pointer(coreglib.InternObject(simple).Native()))
	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_simple_action_set_state(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(value)
}

// SetStateHint sets the state hint for the action.
//
// See g_action_get_state_hint() for more information about action state hints.
//
// The function takes the following parameters:
//
//   - stateHint (optional) representing the state hint.
//
func (simple *SimpleAction) SetStateHint(stateHint *glib.Variant) {
	var _arg0 *C.GSimpleAction // out
	var _arg1 *C.GVariant      // out

	_arg0 = (*C.GSimpleAction)(unsafe.Pointer(coreglib.InternObject(simple).Native()))
	if stateHint != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(stateHint)))
	}

	C.g_simple_action_set_state_hint(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(stateHint)
}

// SimplePermission is a trivial implementation of #GPermission that represents
// a permission that is either always or never allowed. The value is given at
// construction and doesn't change.
//
// Calling request or release will result in errors.
type SimplePermission struct {
	_ [0]func() // equal guard
	Permission
}

var (
	_ Permissioner = (*SimplePermission)(nil)
)

func wrapSimplePermission(obj *coreglib.Object) *SimplePermission {
	return &SimplePermission{
		Permission: Permission{
			Object: obj,
		},
	}
}

func marshalSimplePermission(p uintptr) (interface{}, error) {
	return wrapSimplePermission(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSimplePermission creates a new #GPermission instance that represents an
// action that is either always or never allowed.
//
// The function takes the following parameters:
//
//   - allowed: TRUE if the action is allowed.
//
// The function returns the following values:
//
//   - simplePermission as a #GPermission.
//
func NewSimplePermission(allowed bool) *SimplePermission {
	var _arg1 C.gboolean     // out
	var _cret *C.GPermission // in

	if allowed {
		_arg1 = C.TRUE
	}

	_cret = C.g_simple_permission_new(_arg1)
	runtime.KeepAlive(allowed)

	var _simplePermission *SimplePermission // out

	_simplePermission = wrapSimplePermission(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simplePermission
}
