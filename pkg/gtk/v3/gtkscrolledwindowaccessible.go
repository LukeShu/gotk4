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
		{T: externglib.Type(C.gtk_scrolled_window_accessible_get_type()), F: marshalScrolledWindowAccessible},
	})
}

type ScrolledWindowAccessible interface {
	gextras.Objector

	privateScrolledWindowAccessibleClass()
}

// ScrolledWindowAccessibleClass implements the ScrolledWindowAccessible interface.
type ScrolledWindowAccessibleClass struct {
	ContainerAccessibleClass
}

var _ ScrolledWindowAccessible = (*ScrolledWindowAccessibleClass)(nil)

func wrapScrolledWindowAccessible(obj *externglib.Object) ScrolledWindowAccessible {
	return &ScrolledWindowAccessibleClass{
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

func marshalScrolledWindowAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScrolledWindowAccessible(obj), nil
}

func (*ScrolledWindowAccessibleClass) privateScrolledWindowAccessibleClass() {}
