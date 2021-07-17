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
		{T: externglib.Type(C.gtk_boolean_cell_accessible_get_type()), F: marshalBooleanCellAccessibler},
	})
}

type BooleanCellAccessible struct {
	RendererCellAccessible
}

var _ gextras.Nativer = (*BooleanCellAccessible)(nil)

func wrapBooleanCellAccessible(obj *externglib.Object) *BooleanCellAccessible {
	return &BooleanCellAccessible{
		RendererCellAccessible: RendererCellAccessible{
			CellAccessible: CellAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Action: atk.Action{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
				TableCell: atk.TableCell{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalBooleanCellAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBooleanCellAccessible(obj), nil
}

func (*BooleanCellAccessible) privateBooleanCellAccessible() {}
