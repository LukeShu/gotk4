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

// GTypeArrowAccessible returns the GType for the type ArrowAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeArrowAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_arrow_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalArrowAccessible)
	return gtype
}

// ArrowAccessibleOverrider contains methods that are overridable.
type ArrowAccessibleOverrider interface {
}

type ArrowAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Image
}

var (
	_ coreglib.Objector = (*ArrowAccessible)(nil)
)

func classInitArrowAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapArrowAccessible(obj *coreglib.Object) *ArrowAccessible {
	return &ArrowAccessible{
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
		Image: atk.Image{
			Object: obj,
		},
	}
}

func marshalArrowAccessible(p uintptr) (interface{}, error) {
	return wrapArrowAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
