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
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_button_accessible_get_type()), F: marshalRadioButtonAccessibler},
	})
}

// RadioButtonAccessibler describes RadioButtonAccessible's methods.
type RadioButtonAccessibler interface {
	privateRadioButtonAccessible()
}

type RadioButtonAccessible struct {
	ToggleButtonAccessible
}

var (
	_ RadioButtonAccessibler = (*RadioButtonAccessible)(nil)
	_ gextras.Nativer        = (*RadioButtonAccessible)(nil)
)

func wrapRadioButtonAccessible(obj *externglib.Object) RadioButtonAccessibler {
	return &RadioButtonAccessible{
		ToggleButtonAccessible: ToggleButtonAccessible{
			ButtonAccessible: ButtonAccessible{
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
				Action: atk.Action{
					Object: obj,
				},
				Image: atk.Image{
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioButtonAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRadioButtonAccessible(obj), nil
}

func (*RadioButtonAccessible) privateRadioButtonAccessible() {}
