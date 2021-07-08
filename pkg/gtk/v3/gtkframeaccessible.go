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
		{T: externglib.Type(C.gtk_frame_accessible_get_type()), F: marshalFrameAccessible},
	})
}

type FrameAccessible interface {
	gextras.Objector

	privateFrameAccessibleClass()
}

// FrameAccessibleClass implements the FrameAccessible interface.
type FrameAccessibleClass struct {
	ContainerAccessibleClass
}

var _ FrameAccessible = (*FrameAccessibleClass)(nil)

func wrapFrameAccessible(obj *externglib.Object) FrameAccessible {
	return &FrameAccessibleClass{
		ContainerAccessibleClass: ContainerAccessibleClass{
			WidgetAccessibleClass: WidgetAccessibleClass{
				AccessibleClass: AccessibleClass{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalFrameAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFrameAccessible(obj), nil
}

func (*FrameAccessibleClass) privateFrameAccessibleClass() {}
