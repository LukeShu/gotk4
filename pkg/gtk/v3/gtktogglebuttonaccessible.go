// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtktogglebuttonaccessible.go.
var GTypeToggleButtonAccessible = externglib.Type(C.gtk_toggle_button_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeToggleButtonAccessible, F: marshalToggleButtonAccessible},
	})
}

// ToggleButtonAccessibleOverrider contains methods that are overridable.
type ToggleButtonAccessibleOverrider interface {
	externglib.Objector
}

// WrapToggleButtonAccessibleOverrider wraps the ToggleButtonAccessibleOverrider
// interface implementation to access the instance methods.
func WrapToggleButtonAccessibleOverrider(obj ToggleButtonAccessibleOverrider) *ToggleButtonAccessible {
	return wrapToggleButtonAccessible(externglib.BaseObject(obj))
}

type ToggleButtonAccessible struct {
	_ [0]func() // equal guard
	ButtonAccessible
}

var (
	_ externglib.Objector = (*ToggleButtonAccessible)(nil)
)

func classInitToggleButtonAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapToggleButtonAccessible(obj *externglib.Object) *ToggleButtonAccessible {
	return &ToggleButtonAccessible{
		ButtonAccessible: ButtonAccessible{
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
			Object: obj,
			Action: atk.Action{
				Object: obj,
			},
			Image: atk.Image{
				Object: obj,
			},
		},
	}
}

func marshalToggleButtonAccessible(p uintptr) (interface{}, error) {
	return wrapToggleButtonAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
