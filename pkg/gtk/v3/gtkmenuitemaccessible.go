// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeMenuItemAccessible = coreglib.Type(C.gtk_menu_item_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMenuItemAccessible, F: marshalMenuItemAccessible},
	})
}

// MenuItemAccessibleOverrides contains methods that are overridable.
type MenuItemAccessibleOverrides struct {
}

func defaultMenuItemAccessibleOverrides(v *MenuItemAccessible) MenuItemAccessibleOverrides {
	return MenuItemAccessibleOverrides{}
}

type MenuItemAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*coreglib.Object
	atk.Action
	atk.Selection
}

var (
	_ coreglib.Objector = (*MenuItemAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*MenuItemAccessible, *MenuItemAccessibleClass, MenuItemAccessibleOverrides](
		GTypeMenuItemAccessible,
		initMenuItemAccessibleClass,
		wrapMenuItemAccessible,
		defaultMenuItemAccessibleOverrides,
	)
}

func initMenuItemAccessibleClass(gclass unsafe.Pointer, overrides MenuItemAccessibleOverrides, classInitFunc func(*MenuItemAccessibleClass)) {
	if classInitFunc != nil {
		class := (*MenuItemAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMenuItemAccessible(obj *coreglib.Object) *MenuItemAccessible {
	return &MenuItemAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
		Object: obj,
		Action: atk.Action{
			Object: obj,
		},
		Selection: atk.Selection{
			Object: obj,
		},
	}
}

func marshalMenuItemAccessible(p uintptr) (interface{}, error) {
	return wrapMenuItemAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MenuItemAccessibleClass: instance of this type is always passed by reference.
type MenuItemAccessibleClass struct {
	*menuItemAccessibleClass
}

// menuItemAccessibleClass is the struct that's finalized.
type menuItemAccessibleClass struct {
	native *C.GtkMenuItemAccessibleClass
}

func (m *MenuItemAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &m.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
