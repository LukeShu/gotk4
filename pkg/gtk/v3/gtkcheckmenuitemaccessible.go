// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_check_menu_item_accessible_get_type()), F: marshalCheckMenuItemAccessible},
	})
}

type CheckMenuItemAccessible interface {
	MenuItemAccessible
	Action
}

// checkMenuItemAccessible implements the CheckMenuItemAccessible class.
type checkMenuItemAccessible struct {
	MenuItemAccessible
	Action
}

var _ CheckMenuItemAccessible = (*checkMenuItemAccessible)(nil)

// WrapCheckMenuItemAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapCheckMenuItemAccessible(obj *externglib.Object) CheckMenuItemAccessible {
	return checkMenuItemAccessible{
		MenuItemAccessible: WrapMenuItemAccessible(obj),
		Action:             WrapAction(obj),
	}
}

func marshalCheckMenuItemAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCheckMenuItemAccessible(obj), nil
}
