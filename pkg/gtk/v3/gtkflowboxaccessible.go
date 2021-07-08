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
		{T: externglib.Type(C.gtk_flow_box_accessible_get_type()), F: marshalFlowBoxAccessible},
	})
}

type FlowBoxAccessible interface {
	gextras.Objector

	privateFlowBoxAccessibleClass()
}

// FlowBoxAccessibleClass implements the FlowBoxAccessible interface.
type FlowBoxAccessibleClass struct {
	ContainerAccessibleClass
}

var _ FlowBoxAccessible = (*FlowBoxAccessibleClass)(nil)

func wrapFlowBoxAccessible(obj *externglib.Object) FlowBoxAccessible {
	return &FlowBoxAccessibleClass{
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

func marshalFlowBoxAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFlowBoxAccessible(obj), nil
}

func (*FlowBoxAccessibleClass) privateFlowBoxAccessibleClass() {}
