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

// GTypeImageCellAccessible returns the GType for the type ImageCellAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeImageCellAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_image_cell_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalImageCellAccessible)
	return gtype
}

// ImageCellAccessibleOverrider contains methods that are overridable.
type ImageCellAccessibleOverrider interface {
}

type ImageCellAccessible struct {
	_ [0]func() // equal guard
	RendererCellAccessible

	*coreglib.Object
	atk.Image
	atk.ObjectClass
}

var (
	_ coreglib.Objector = (*ImageCellAccessible)(nil)
)

func classInitImageCellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapImageCellAccessible(obj *coreglib.Object) *ImageCellAccessible {
	return &ImageCellAccessible{
		RendererCellAccessible: RendererCellAccessible{
			CellAccessible: CellAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
				TableCell: atk.TableCell{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
		Object: obj,
		Image: atk.Image{
			Object: obj,
		},
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
	}
}

func marshalImageCellAccessible(p uintptr) (interface{}, error) {
	return wrapImageCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
