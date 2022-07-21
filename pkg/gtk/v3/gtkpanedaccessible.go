// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypePanedAccessible returns the GType for the type PanedAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePanedAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_paned_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalPanedAccessible)
	return gtype
}

// PanedAccessibleOverrider contains methods that are overridable.
type PanedAccessibleOverrider interface {
}

type PanedAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Value
}

var (
	_ coreglib.Objector = (*PanedAccessible)(nil)
)

func classInitPanedAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPanedAccessible(obj *coreglib.Object) *PanedAccessible {
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
	return wrapPanedAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
