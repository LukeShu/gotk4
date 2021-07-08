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
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_expander_accessible_get_type()), F: marshalExpanderAccessible},
	})
}

type ExpanderAccessible interface {
	gextras.Objector

	privateExpanderAccessibleClass()
}

// ExpanderAccessibleClass implements the ExpanderAccessible interface.
type ExpanderAccessibleClass struct {
	*externglib.Object
	ContainerAccessibleClass
	atk.ActionInterface
}

var _ ExpanderAccessible = (*ExpanderAccessibleClass)(nil)

func wrapExpanderAccessible(obj *externglib.Object) ExpanderAccessible {
	return &ExpanderAccessibleClass{
		Object: obj,
		ContainerAccessibleClass: ContainerAccessibleClass{
			WidgetAccessibleClass: WidgetAccessibleClass{
				AccessibleClass: AccessibleClass{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
		ActionInterface: atk.ActionInterface{
			Object: obj,
		},
	}
}

func marshalExpanderAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapExpanderAccessible(obj), nil
}

func (*ExpanderAccessibleClass) privateExpanderAccessibleClass() {}
