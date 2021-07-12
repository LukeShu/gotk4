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
		{T: externglib.Type(C.gtk_entry_accessible_get_type()), F: marshalEntryAccessibler},
	})
}

// EntryAccessibler describes EntryAccessible's methods.
type EntryAccessibler interface {
	privateEntryAccessible()
}

type EntryAccessible struct {
	WidgetAccessible

	atk.Action
	atk.EditableText
	atk.Text
}

var (
	_ EntryAccessibler = (*EntryAccessible)(nil)
	_ gextras.Nativer  = (*EntryAccessible)(nil)
)

func wrapEntryAccessible(obj *externglib.Object) EntryAccessibler {
	return &EntryAccessible{
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
	}
}

func marshalEntryAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEntryAccessible(obj), nil
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *EntryAccessible) Native() uintptr {
	return v.WidgetAccessible.Accessible.ObjectClass.Object.Native()
}

func (*EntryAccessible) privateEntryAccessible() {}
