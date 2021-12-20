// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_paned_accessible_get_type()), F: marshalPanedAccessibler},
	})
}

type PanedAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Value
}

var (
	_ externglib.Objector = (*PanedAccessible)(nil)
)

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

func marshalPanedAccessibler(p uintptr) (interface{}, error) {
	return wrapPanedAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
