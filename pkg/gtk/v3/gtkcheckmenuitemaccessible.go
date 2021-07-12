// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
		{T: externglib.Type(C.gtk_check_menu_item_accessible_get_type()), F: marshalCheckMenuItemAccessibler},
	})
}

// CheckMenuItemAccessibler describes CheckMenuItemAccessible's methods.
type CheckMenuItemAccessibler interface {
	privateCheckMenuItemAccessible()
}

type CheckMenuItemAccessible struct {
	MenuItemAccessible
}

var (
	_ CheckMenuItemAccessibler = (*CheckMenuItemAccessible)(nil)
	_ gextras.Nativer          = (*CheckMenuItemAccessible)(nil)
)

func wrapCheckMenuItemAccessible(obj *externglib.Object) CheckMenuItemAccessibler {
	return &CheckMenuItemAccessible{
		MenuItemAccessible: MenuItemAccessible{
			ContainerAccessible: ContainerAccessible{
				WidgetAccessible: WidgetAccessible{
					Accessible: Accessible{
						ObjectClass: atk.ObjectClass{
							Object: obj,
						},
					},
					Component: atk.Component{
						Object: obj,
					},
				},
			},
			Action: atk.Action{
				Object: obj,
			},
			Selection: atk.Selection{
				Object: obj,
			},
		},
	}
}

func marshalCheckMenuItemAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCheckMenuItemAccessible(obj), nil
}

func (*CheckMenuItemAccessible) privateCheckMenuItemAccessible() {}
