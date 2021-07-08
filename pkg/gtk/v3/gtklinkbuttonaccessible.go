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
		{T: externglib.Type(C.gtk_link_button_accessible_get_type()), F: marshalLinkButtonAccessible},
	})
}

type LinkButtonAccessible interface {
	gextras.Objector

	privateLinkButtonAccessibleClass()
}

// LinkButtonAccessibleClass implements the LinkButtonAccessible interface.
type LinkButtonAccessibleClass struct {
	*externglib.Object
	ButtonAccessibleClass
	atk.ActionInterface
	atk.ImageInterface
}

var _ LinkButtonAccessible = (*LinkButtonAccessibleClass)(nil)

func wrapLinkButtonAccessible(obj *externglib.Object) LinkButtonAccessible {
	return &LinkButtonAccessibleClass{
		Object: obj,
		ButtonAccessibleClass: ButtonAccessibleClass{
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
			ImageInterface: atk.ImageInterface{
				Object: obj,
			},
		},
		ActionInterface: atk.ActionInterface{
			Object: obj,
		},
		ImageInterface: atk.ImageInterface{
			Object: obj,
		},
	}
}

func marshalLinkButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLinkButtonAccessible(obj), nil
}

func (*LinkButtonAccessibleClass) privateLinkButtonAccessibleClass() {}
