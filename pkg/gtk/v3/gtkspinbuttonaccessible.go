// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
		{T: externglib.Type(C.gtk_spin_button_accessible_get_type()), F: marshalSpinButtonAccessibler},
	})
}

type SpinButtonAccessible struct {
	EntryAccessible

	atk.Value
	*externglib.Object
}

func wrapSpinButtonAccessible(obj *externglib.Object) *SpinButtonAccessible {
	return &SpinButtonAccessible{
		EntryAccessible: EntryAccessible{
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
			Action: atk.Action{
				Object: obj,
			},
			EditableText: atk.EditableText{
				Object: obj,
			},
			Text: atk.Text{
				Object: obj,
			},
			Object: obj,
		},
		Value: atk.Value{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalSpinButtonAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSpinButtonAccessible(obj), nil
}

func (*SpinButtonAccessible) privateSpinButtonAccessible() {}
