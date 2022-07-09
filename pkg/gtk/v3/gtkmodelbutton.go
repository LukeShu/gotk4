// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeButtonRole returns the GType for the type ButtonRole.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeButtonRole() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ButtonRole").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalButtonRole)
	return gtype
}

// GTypeModelButton returns the GType for the type ModelButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeModelButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ModelButton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalModelButton)
	return gtype
}

// ButtonRole: role specifies the desired appearance of a ModelButton.
type ButtonRole C.gint

const (
	// ButtonRoleNormal: plain button.
	ButtonRoleNormal ButtonRole = iota
	// ButtonRoleCheck: check button.
	ButtonRoleCheck
	// ButtonRoleRadio: radio button.
	ButtonRoleRadio
)

func marshalButtonRole(p uintptr) (interface{}, error) {
	return ButtonRole(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ButtonRole.
func (b ButtonRole) String() string {
	switch b {
	case ButtonRoleNormal:
		return "Normal"
	case ButtonRoleCheck:
		return "Check"
	case ButtonRoleRadio:
		return "Radio"
	default:
		return fmt.Sprintf("ButtonRole(%d)", b)
	}
}

// ModelButton is a button class that can use a #GAction as its model. In
// contrast to ToggleButton or RadioButton, which can also be backed by a
// #GAction via the Actionable:action-name property, GtkModelButton will adapt
// its appearance according to the kind of action it is backed by, and appear
// either as a plain, check or radio button.
//
// Model buttons are used when popovers from a menu model with
// gtk_popover_new_from_model(); they can also be used manually in a
// PopoverMenu.
//
// When the action is specified via the Actionable:action-name and
// Actionable:action-target properties, the role of the button (i.e. whether it
// is a plain, check or radio button) is determined by the type of the action
// and doesn't have to be explicitly specified with the ModelButton:role
// property.
//
// The content of the button is specified by the ModelButton:text and
// ModelButton:icon properties.
//
// The appearance of model buttons can be influenced with the
// ModelButton:centered and ModelButton:iconic properties.
//
// Model buttons have built-in support for submenus in PopoverMenu. To make a
// GtkModelButton that opens a submenu when activated, set the
// ModelButton:menu-name property. To make a button that goes back to the parent
// menu, you should set the ModelButton:inverted property to place the submenu
// indicator at the opposite side.
//
// Example
//
//    <object class="GtkPopoverMenu">
//      <child>
//        <object class="GtkBox">
//          <property name="visible">True</property>
//          <property name="margin">10</property>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">view.cut</property>
//              <property name="text" translatable="yes">Cut</property>
//            </object>
//          </child>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">view.copy</property>
//              <property name="text" translatable="yes">Copy</property>
//            </object>
//          </child>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">view.paste</property>
//              <property name="text" translatable="yes">Paste</property>
//            </object>
//          </child>
//        </object>
//      </child>
//    </object>
//
// CSS nodes
//
//    button.model
//    ├── <child>
//    ╰── check
//
// Iconic model buttons (see ModelButton:iconic) change the name of their main
// node to button and add a .model style class to it. The indicator subnode is
// invisible in this case.
type ModelButton struct {
	_ [0]func() // equal guard
	Button
}

var (
	_ Binner            = (*ModelButton)(nil)
	_ coreglib.Objector = (*ModelButton)(nil)
)

func wrapModelButton(obj *coreglib.Object) *ModelButton {
	return &ModelButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Object: obj,
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalModelButton(p uintptr) (interface{}, error) {
	return wrapModelButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewModelButton creates a new GtkModelButton.
//
// The function returns the following values:
//
//    - modelButton: newly created ModelButton widget.
//
func NewModelButton() *ModelButton {
	_gret := girepository.MustFind("Gtk", "ModelButton").InvokeMethod("new_ModelButton", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _modelButton *ModelButton // out

	_modelButton = wrapModelButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _modelButton
}
