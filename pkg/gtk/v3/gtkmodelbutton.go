// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_button_role_get_type()), F: marshalButtonRole},
		{T: externglib.Type(C.gtk_model_button_get_type()), F: marshalModelButton},
	})
}

// ButtonRole: the role specifies the desired appearance of a ModelButton.
type ButtonRole int

const (
	// Normal: plain button
	ButtonRoleNormal ButtonRole = iota
	// Check button
	ButtonRoleCheck
	// Radio button
	ButtonRoleRadio
)

func marshalButtonRole(p uintptr) (interface{}, error) {
	return ButtonRole(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
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
type ModelButton interface {
	gextras.Objector

	privateModelButtonClass()
}

// ModelButtonClass implements the ModelButton interface.
type ModelButtonClass struct {
	*externglib.Object
	ButtonClass
	ActionableInterface
	ActivatableInterface
	BuildableInterface
}

var _ ModelButton = (*ModelButtonClass)(nil)

func wrapModelButton(obj *externglib.Object) ModelButton {
	return &ModelButtonClass{
		Object: obj,
		ButtonClass: ButtonClass{
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
			ActionableInterface: ActionableInterface{
				WidgetClass: WidgetClass{
					Object:           obj,
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
			},
			ActivatableInterface: ActivatableInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		ActionableInterface: ActionableInterface{
			WidgetClass: WidgetClass{
				Object:           obj,
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
		},
		ActivatableInterface: ActivatableInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalModelButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapModelButton(obj), nil
}

// NewModelButton creates a new GtkModelButton.
func NewModelButton() ModelButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_model_button_new()

	var _modelButton ModelButton // out

	_modelButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ModelButton)

	return _modelButton
}

func (*ModelButtonClass) privateModelButtonClass() {}
