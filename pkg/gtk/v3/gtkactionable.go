// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_actionable_get_type()), F: marshalActionable},
	})
}

// ActionableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionableOverrider interface {
	// ActionName gets the action name for @actionable.
	//
	// See gtk_actionable_set_action_name() for more information.
	ActionName() string
	// ActionTargetValue gets the current target value of @actionable.
	//
	// See gtk_actionable_set_action_target_value() for more information.
	ActionTargetValue() *glib.Variant
	// SetActionName specifies the name of the action with which this widget
	// should be associated. If @action_name is nil then the widget will be
	// unassociated from any previous action.
	//
	// Usually this function is used when the widget is located (or will be
	// located) within the hierarchy of a ApplicationWindow.
	//
	// Names are of the form “win.save” or “app.quit” for actions on the
	// containing ApplicationWindow or its associated Application, respectively.
	// This is the same form used for actions in the #GMenu associated with the
	// window.
	SetActionName(actionName string)
	// SetActionTargetValue sets the target value of an actionable widget.
	//
	// If @target_value is nil then the target value is unset.
	//
	// The target value has two purposes. First, it is used as the parameter to
	// activation of the action associated with the Actionable widget. Second,
	// it is used to determine if the widget should be rendered as “active” —
	// the widget is active if the state is equal to the given target.
	//
	// Consider the example of associating a set of buttons with a #GAction with
	// string state in a typical “radio button” situation. Each button will be
	// associated with the same action, but with a different target value for
	// that action. Clicking on a particular button will activate the action
	// with the target of that button, which will typically cause the action’s
	// state to change to that value. Since the action’s state is now equal to
	// the target value of the button, the button will now be rendered as active
	// (and the other buttons, with different targets, rendered inactive).
	SetActionTargetValue(targetValue *glib.Variant)
}

// Actionable: this interface provides a convenient way of associating widgets
// with actions on a ApplicationWindow or Application.
//
// It primarily consists of two properties: Actionable:action-name and
// Actionable:action-target. There are also some convenience APIs for setting
// these properties.
//
// The action will be looked up in action groups that are found among the
// widgets ancestors. Most commonly, these will be the actions with the “win.”
// or “app.” prefix that are associated with the ApplicationWindow or
// Application, but other action groups that are added with
// gtk_widget_insert_action_group() will be consulted as well.
type Actionable interface {
	gextras.Objector

	// ActionName gets the action name for @actionable.
	//
	// See gtk_actionable_set_action_name() for more information.
	ActionName() string
	// ActionTargetValue gets the current target value of @actionable.
	//
	// See gtk_actionable_set_action_target_value() for more information.
	ActionTargetValue() *glib.Variant
	// SetActionName specifies the name of the action with which this widget
	// should be associated. If @action_name is nil then the widget will be
	// unassociated from any previous action.
	//
	// Usually this function is used when the widget is located (or will be
	// located) within the hierarchy of a ApplicationWindow.
	//
	// Names are of the form “win.save” or “app.quit” for actions on the
	// containing ApplicationWindow or its associated Application, respectively.
	// This is the same form used for actions in the #GMenu associated with the
	// window.
	SetActionName(actionName string)
	// SetActionTargetValue sets the target value of an actionable widget.
	//
	// If @target_value is nil then the target value is unset.
	//
	// The target value has two purposes. First, it is used as the parameter to
	// activation of the action associated with the Actionable widget. Second,
	// it is used to determine if the widget should be rendered as “active” —
	// the widget is active if the state is equal to the given target.
	//
	// Consider the example of associating a set of buttons with a #GAction with
	// string state in a typical “radio button” situation. Each button will be
	// associated with the same action, but with a different target value for
	// that action. Clicking on a particular button will activate the action
	// with the target of that button, which will typically cause the action’s
	// state to change to that value. Since the action’s state is now equal to
	// the target value of the button, the button will now be rendered as active
	// (and the other buttons, with different targets, rendered inactive).
	SetActionTargetValue(targetValue *glib.Variant)
	// SetDetailedActionName sets the action-name and associated string target
	// value of an actionable widget.
	//
	// @detailed_action_name is a string in the format accepted by
	// g_action_parse_detailed_name().
	//
	// (Note that prior to version 3.22.25, this function is only usable for
	// actions with a simple "s" target, and @detailed_action_name must be of
	// the form `"action::target"` where `action` is the action name and
	// `target` is the string to use as the target.)
	SetDetailedActionName(detailedActionName string)
}

// ActionableInterface implements the Actionable interface.
type ActionableInterface struct {
	WidgetClass
}

var _ Actionable = (*ActionableInterface)(nil)

func wrapActionable(obj *externglib.Object) Actionable {
	return &ActionableInterface{
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
	}
}

func marshalActionable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActionable(obj), nil
}

// ActionName gets the action name for @actionable.
//
// See gtk_actionable_set_action_name() for more information.
func (a *ActionableInterface) ActionName() string {
	var _arg0 *C.GtkActionable // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkActionable)(unsafe.Pointer((&a).Native()))

	_cret = C.gtk_actionable_get_action_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ActionTargetValue gets the current target value of @actionable.
//
// See gtk_actionable_set_action_target_value() for more information.
func (a *ActionableInterface) ActionTargetValue() *glib.Variant {
	var _arg0 *C.GtkActionable // out
	var _cret *C.GVariant      // in

	_arg0 = (*C.GtkActionable)(unsafe.Pointer((&a).Native()))

	_cret = C.gtk_actionable_get_action_target_value(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// SetActionName specifies the name of the action with which this widget should
// be associated. If @action_name is nil then the widget will be unassociated
// from any previous action.
//
// Usually this function is used when the widget is located (or will be located)
// within the hierarchy of a ApplicationWindow.
//
// Names are of the form “win.save” or “app.quit” for actions on the containing
// ApplicationWindow or its associated Application, respectively. This is the
// same form used for actions in the #GMenu associated with the window.
func (a *ActionableInterface) SetActionName(actionName string) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer((&a).Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_actionable_set_action_name(_arg0, _arg1)
}

// SetActionTargetValue sets the target value of an actionable widget.
//
// If @target_value is nil then the target value is unset.
//
// The target value has two purposes. First, it is used as the parameter to
// activation of the action associated with the Actionable widget. Second, it is
// used to determine if the widget should be rendered as “active” — the widget
// is active if the state is equal to the given target.
//
// Consider the example of associating a set of buttons with a #GAction with
// string state in a typical “radio button” situation. Each button will be
// associated with the same action, but with a different target value for that
// action. Clicking on a particular button will activate the action with the
// target of that button, which will typically cause the action’s state to
// change to that value. Since the action’s state is now equal to the target
// value of the button, the button will now be rendered as active (and the other
// buttons, with different targets, rendered inactive).
func (a *ActionableInterface) SetActionTargetValue(targetValue *glib.Variant) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.GVariant      // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer((&a).Native()))
	_arg1 = (*C.GVariant)(unsafe.Pointer(targetValue))

	C.gtk_actionable_set_action_target_value(_arg0, _arg1)
}

// SetDetailedActionName sets the action-name and associated string target value
// of an actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// g_action_parse_detailed_name().
//
// (Note that prior to version 3.22.25, this function is only usable for actions
// with a simple "s" target, and @detailed_action_name must be of the form
// `"action::target"` where `action` is the action name and `target` is the
// string to use as the target.)
func (a *ActionableInterface) SetDetailedActionName(detailedActionName string) {
	var _arg0 *C.GtkActionable // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkActionable)(unsafe.Pointer((&a).Native()))
	_arg1 = (*C.gchar)(C.CString(detailedActionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_actionable_set_detailed_action_name(_arg0, _arg1)
}
