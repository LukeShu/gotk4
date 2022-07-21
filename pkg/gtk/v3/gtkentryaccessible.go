// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeEntryAccessible returns the GType for the type EntryAccessible.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEntryAccessible() coreglib.Type {
	gtype := coreglib.Type(C.gtk_entry_accessible_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalEntryAccessible)
	return gtype
}

// EntryAccessibleOverrider contains methods that are overridable.
type EntryAccessibleOverrider interface {
}

type EntryAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	*coreglib.Object
	atk.Action
	atk.EditableText
	atk.Text
}

var (
	_ coreglib.Objector = (*EntryAccessible)(nil)
)

func classInitEntryAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEntryAccessible(obj *coreglib.Object) *EntryAccessible {
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
		Object: obj,
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

func marshalEntryAccessible(p uintptr) (interface{}, error) {
	return wrapEntryAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
