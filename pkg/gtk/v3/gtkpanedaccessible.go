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

// glib.Type values for gtkpanedaccessible.go.
var GTypePanedAccessible = externglib.Type(C.gtk_paned_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePanedAccessible, F: marshalPanedAccessible},
	})
}

// PanedAccessibleOverrider contains methods that are overridable.
type PanedAccessibleOverrider interface {
	externglib.Objector
}

// WrapPanedAccessibleOverrider wraps the PanedAccessibleOverrider
// interface implementation to access the instance methods.
func WrapPanedAccessibleOverrider(obj PanedAccessibleOverrider) *PanedAccessible {
	return wrapPanedAccessible(externglib.BaseObject(obj))
}

type PanedAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Value
}

var (
	_ externglib.Objector = (*PanedAccessible)(nil)
)

func classInitPanedAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPanedAccessible(obj *externglib.Object) *PanedAccessible {
	return &PanedAccessible{
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
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalPanedAccessible(p uintptr) (interface{}, error) {
	return wrapPanedAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
