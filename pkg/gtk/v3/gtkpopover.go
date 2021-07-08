// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_popover_get_type()), F: marshalPopover},
	})
}

// PopoverOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PopoverOverrider interface {
	Closed()
}

// Popover is a bubble-like context window, primarily meant to provide
// context-dependent information or options. Popovers are attached to a widget,
// passed at construction time on gtk_popover_new(), or updated afterwards
// through gtk_popover_set_relative_to(), by default they will point to the
// whole widget area, although this behavior can be changed through
// gtk_popover_set_pointing_to().
//
// The position of a popover relative to the widget it is attached to can also
// be changed through gtk_popover_set_position().
//
// By default, Popover performs a GTK+ grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Esc key being
// pressed). If no such modal behavior is desired on a popover,
// gtk_popover_set_modal() may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// GtkPopover is often used to replace menus. To facilitate this, it supports
// being populated from a Model, using gtk_popover_new_from_model(). In addition
// to all the regular menu model features, this function supports rendering
// sections in the model in a more compact form, as a row of icon buttons
// instead of menu items.
//
// To use this rendering, set the ”display-hint” attribute of the section to
// ”horizontal-buttons” and set the icons of your items with the ”verb-icon”
// attribute.
//
//    <section>
//      <attribute name="display-hint">horizontal-buttons</attribute>
//      <item>
//        <attribute name="label">Cut</attribute>
//        <attribute name="action">app.cut</attribute>
//        <attribute name="verb-icon">edit-cut-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Copy</attribute>
//        <attribute name="action">app.copy</attribute>
//        <attribute name="verb-icon">edit-copy-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Paste</attribute>
//        <attribute name="action">app.paste</attribute>
//        <attribute name="verb-icon">edit-paste-symbolic</attribute>
//      </item>
//    </section>
//
//
// CSS nodes
//
// GtkPopover has a single css node called popover. It always gets the
// .background style class and it gets the .menu style class if it is menu-like
// (e.g. PopoverMenu or created using gtk_popover_new_from_model().
//
// Particular uses of GtkPopover, such as touch selection popups or magnifiers
// in Entry or TextView get style classes like .touch-selection or .magnifier to
// differentiate from plain popovers.
type Popover interface {
	gextras.Objector

	// BindModel establishes a binding between a Popover and a Model.
	//
	// The contents of @popover are removed and then refilled with menu items
	// according to @model. When @model changes, @popover is updated. Calling
	// this function twice on @popover with different @model will cause the
	// first binding to be replaced with a binding to the new model. If @model
	// is nil then any previous binding is undone and all children are removed.
	//
	// If @action_namespace is non-nil then the effect is as if all actions
	// mentioned in the @model have their names prefixed with the namespace,
	// plus a dot. For example, if the action “quit” is mentioned and
	// @action_namespace is “app” then the effective action name is “app.quit”.
	//
	// This function uses Actionable to define the action name and target values
	// on the created menu items. If you want to use an action group other than
	// “app” and “win”, or if you want to use a MenuShell outside of a
	// ApplicationWindow, then you will need to attach your own action group to
	// the widget hierarchy using gtk_widget_insert_action_group(). As an
	// example, if you created a group with a “quit” action and inserted it with
	// the name “mygroup” then you would use the action name “mygroup.quit” in
	// your Model.
	BindModel(model gio.MenuModel, actionNamespace string)
	// ConstrainTo returns the constraint for placing this popover. See
	// gtk_popover_set_constrain_to().
	ConstrainTo() PopoverConstraint
	// DefaultWidget gets the widget that should be set as the default while the
	// popover is shown.
	DefaultWidget() Widget
	// Modal returns whether the popover is modal, see gtk_popover_set_modal to
	// see the implications of this.
	Modal() bool
	// Position returns the preferred position of @popover.
	Position() PositionType
	// RelativeTo returns the widget @popover is currently attached to
	RelativeTo() Widget
	// TransitionsEnabled returns whether show/hide transitions are enabled on
	// this popover.
	//
	// Deprecated: since version 3.22.
	TransitionsEnabled() bool
	// Popdown pops @popover down.This is different than a gtk_widget_hide()
	// call in that it shows the popover with a transition. If you want to hide
	// the popover without a transition, use gtk_widget_hide().
	Popdown()
	// Popup pops @popover up. This is different than a gtk_widget_show() call
	// in that it shows the popover with a transition. If you want to show the
	// popover without a transition, use gtk_widget_show().
	Popup()
	// SetConstrainTo sets a constraint for positioning this popover.
	//
	// Note that not all platforms support placing popovers freely, and may
	// already impose constraints.
	SetConstrainTo(constraint PopoverConstraint)
	// SetDefaultWidget sets the widget that should be set as default widget
	// while the popover is shown (see gtk_window_set_default()). Popover
	// remembers the previous default widget and reestablishes it when the
	// popover is dismissed.
	SetDefaultWidget(widget Widget)
	// SetModal sets whether @popover is modal, a modal popover will grab all
	// input within the toplevel and grab the keyboard focus on it when being
	// displayed. Clicking outside the popover area or pressing Esc will dismiss
	// the popover and ungrab input.
	SetModal(modal bool)
	// SetPointingTo sets the rectangle that @popover will point to, in the
	// coordinate space of the widget @popover is attached to, see
	// gtk_popover_set_relative_to().
	SetPointingTo(rect *gdk.Rectangle)
	// SetPosition sets the preferred position for @popover to appear. If the
	// @popover is currently visible, it will be immediately updated.
	//
	// This preference will be respected where possible, although on lack of
	// space (eg. if close to the window edges), the Popover may choose to
	// appear on the opposite side
	SetPosition(position PositionType)
	// SetRelativeTo sets a new widget to be attached to @popover. If @popover
	// is visible, the position will be updated.
	//
	// Note: the ownership of popovers is always given to their @relative_to
	// widget, so if @relative_to is set to nil on an attached @popover, it will
	// be detached from its previous widget, and consequently destroyed unless
	// extra references are kept.
	SetRelativeTo(relativeTo Widget)
	// SetTransitionsEnabled sets whether show/hide transitions are enabled on
	// this popover
	//
	// Deprecated: since version 3.22.
	SetTransitionsEnabled(transitionsEnabled bool)
}

// PopoverClass implements the Popover interface.
type PopoverClass struct {
	*externglib.Object
	BinClass
	BuildableInterface
}

var _ Popover = (*PopoverClass)(nil)

func wrapPopover(obj *externglib.Object) Popover {
	return &PopoverClass{
		Object: obj,
		BinClass: BinClass{
			Object: obj,
			ContainerClass: ContainerClass{
				Object: obj,
				WidgetClass: WidgetClass{
					Object:           obj,
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPopover(obj), nil
}

// NewPopover creates a new popover to point to @relative_to
func NewPopover(relativeTo Widget) Popover {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))

	_cret = C.gtk_popover_new(_arg1)

	var _popover Popover // out

	_popover = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Popover)

	return _popover
}

// NewPopoverFromModel creates a Popover and populates it according to @model.
// The popover is pointed to the @relative_to widget.
//
// The created buttons are connected to actions found in the ApplicationWindow
// to which the popover belongs - typically by means of being attached to a
// widget that is contained within the ApplicationWindows widget hierarchy.
//
// Actions can also be added using gtk_widget_insert_action_group() on the menus
// attach widget or on any of its parent widgets.
func NewPopoverFromModel(relativeTo Widget, model gio.MenuModel) Popover {
	var _arg1 *C.GtkWidget  // out
	var _arg2 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_popover_new_from_model(_arg1, _arg2)

	var _popover Popover // out

	_popover = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Popover)

	return _popover
}

// BindModel establishes a binding between a Popover and a Model.
//
// The contents of @popover are removed and then refilled with menu items
// according to @model. When @model changes, @popover is updated. Calling this
// function twice on @popover with different @model will cause the first binding
// to be replaced with a binding to the new model. If @model is nil then any
// previous binding is undone and all children are removed.
//
// If @action_namespace is non-nil then the effect is as if all actions
// mentioned in the @model have their names prefixed with the namespace, plus a
// dot. For example, if the action “quit” is mentioned and @action_namespace is
// “app” then the effective action name is “app.quit”.
//
// This function uses Actionable to define the action name and target values on
// the created menu items. If you want to use an action group other than “app”
// and “win”, or if you want to use a MenuShell outside of a ApplicationWindow,
// then you will need to attach your own action group to the widget hierarchy
// using gtk_widget_insert_action_group(). As an example, if you created a group
// with a “quit” action and inserted it with the name “mygroup” then you would
// use the action name “mygroup.quit” in your Model.
func (p *PopoverClass) BindModel(model gio.MenuModel, actionNamespace string) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GMenuModel // out
	var _arg2 *C.gchar      // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg2 = (*C.gchar)(C.CString(actionNamespace))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_popover_bind_model(_arg0, _arg1, _arg2)
}

// ConstrainTo returns the constraint for placing this popover. See
// gtk_popover_set_constrain_to().
func (p *PopoverClass) ConstrainTo() PopoverConstraint {
	var _arg0 *C.GtkPopover          // out
	var _cret C.GtkPopoverConstraint // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_constrain_to(_arg0)

	var _popoverConstraint PopoverConstraint // out

	_popoverConstraint = PopoverConstraint(_cret)

	return _popoverConstraint
}

// DefaultWidget gets the widget that should be set as the default while the
// popover is shown.
func (p *PopoverClass) DefaultWidget() Widget {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_default_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

// Modal returns whether the popover is modal, see gtk_popover_set_modal to see
// the implications of this.
func (p *PopoverClass) Modal() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Position returns the preferred position of @popover.
func (p *PopoverClass) Position() PositionType {
	var _arg0 *C.GtkPopover     // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_position(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// RelativeTo returns the widget @popover is currently attached to
func (p *PopoverClass) RelativeTo() Widget {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_relative_to(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

// TransitionsEnabled returns whether show/hide transitions are enabled on this
// popover.
//
// Deprecated: since version 3.22.
func (p *PopoverClass) TransitionsEnabled() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_transitions_enabled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown pops @popover down.This is different than a gtk_widget_hide() call in
// that it shows the popover with a transition. If you want to hide the popover
// without a transition, use gtk_widget_hide().
func (p *PopoverClass) Popdown() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popdown(_arg0)
}

// Popup pops @popover up. This is different than a gtk_widget_show() call in
// that it shows the popover with a transition. If you want to show the popover
// without a transition, use gtk_widget_show().
func (p *PopoverClass) Popup() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popup(_arg0)
}

// SetConstrainTo sets a constraint for positioning this popover.
//
// Note that not all platforms support placing popovers freely, and may already
// impose constraints.
func (p *PopoverClass) SetConstrainTo(constraint PopoverConstraint) {
	var _arg0 *C.GtkPopover          // out
	var _arg1 C.GtkPopoverConstraint // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = C.GtkPopoverConstraint(constraint)

	C.gtk_popover_set_constrain_to(_arg0, _arg1)
}

// SetDefaultWidget sets the widget that should be set as default widget while
// the popover is shown (see gtk_window_set_default()). Popover remembers the
// previous default widget and reestablishes it when the popover is dismissed.
func (p *PopoverClass) SetDefaultWidget(widget Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_popover_set_default_widget(_arg0, _arg1)
}

// SetModal sets whether @popover is modal, a modal popover will grab all input
// within the toplevel and grab the keyboard focus on it when being displayed.
// Clicking outside the popover area or pressing Esc will dismiss the popover
// and ungrab input.
func (p *PopoverClass) SetModal(modal bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_modal(_arg0, _arg1)
}

// SetPointingTo sets the rectangle that @popover will point to, in the
// coordinate space of the widget @popover is attached to, see
// gtk_popover_set_relative_to().
func (p *PopoverClass) SetPointingTo(rect *gdk.Rectangle) {
	var _arg0 *C.GtkPopover   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect))

	C.gtk_popover_set_pointing_to(_arg0, _arg1)
}

// SetPosition sets the preferred position for @popover to appear. If the
// @popover is currently visible, it will be immediately updated.
//
// This preference will be respected where possible, although on lack of space
// (eg. if close to the window edges), the Popover may choose to appear on the
// opposite side
func (p *PopoverClass) SetPosition(position PositionType) {
	var _arg0 *C.GtkPopover     // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_popover_set_position(_arg0, _arg1)
}

// SetRelativeTo sets a new widget to be attached to @popover. If @popover is
// visible, the position will be updated.
//
// Note: the ownership of popovers is always given to their @relative_to widget,
// so if @relative_to is set to nil on an attached @popover, it will be detached
// from its previous widget, and consequently destroyed unless extra references
// are kept.
func (p *PopoverClass) SetRelativeTo(relativeTo Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))

	C.gtk_popover_set_relative_to(_arg0, _arg1)
}

// SetTransitionsEnabled sets whether show/hide transitions are enabled on this
// popover
//
// Deprecated: since version 3.22.
func (p *PopoverClass) SetTransitionsEnabled(transitionsEnabled bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if transitionsEnabled {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_transitions_enabled(_arg0, _arg1)
}
