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

// GTypeTreeViewAccessible returns the GType for the type TreeViewAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeViewAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_tree_view_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeViewAccessible)
	return gtype
}

// TreeViewAccessibleOverrider contains methods that are overridable.
type TreeViewAccessibleOverrider interface {
}

type TreeViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*coreglib.Object
	atk.Selection
	atk.Table
	CellAccessibleParent
}

var (
	_ coreglib.Objector = (*TreeViewAccessible)(nil)
)

func classInitTreeViewAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeViewAccessible(obj *coreglib.Object) *TreeViewAccessible {
	return &TreeViewAccessible{
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
		Selection: atk.Selection{
			Object: obj,
		},
		Table: atk.Table{
			Object: obj,
		},
		CellAccessibleParent: CellAccessibleParent{
			Object: obj,
		},
	}
}

func marshalTreeViewAccessible(p uintptr) (interface{}, error) {
	return wrapTreeViewAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
