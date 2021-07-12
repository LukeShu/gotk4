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
		{T: externglib.Type(C.gtk_progress_bar_accessible_get_type()), F: marshalProgressBarAccessibler},
	})
}

// ProgressBarAccessibler describes ProgressBarAccessible's methods.
type ProgressBarAccessibler interface {
	privateProgressBarAccessible()
}

type ProgressBarAccessible struct {
	WidgetAccessible

	atk.Value
}

var (
	_ ProgressBarAccessibler = (*ProgressBarAccessible)(nil)
	_ gextras.Nativer        = (*ProgressBarAccessible)(nil)
)

func wrapProgressBarAccessible(obj *externglib.Object) ProgressBarAccessibler {
	return &ProgressBarAccessible{
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
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalProgressBarAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapProgressBarAccessible(obj), nil
}

func (*ProgressBarAccessible) privateProgressBarAccessible() {}
