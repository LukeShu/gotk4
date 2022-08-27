// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_ShortcutsSection_ConnectChangeCurrentPage(gpointer, gint, guintptr);
import "C"

// GType values.
var (
	GTypeShortcutsSection = coreglib.Type(C.gtk_shortcuts_section_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutsSection, F: marshalShortcutsSection},
	})
}

// ShortcutsSection: GtkShortcutsSection collects all the keyboard shortcuts and
// gestures for a major application mode.
//
// If your application needs multiple sections, you should give each section a
// unique gtk.ShortcutsSection:section-name and a gtk.ShortcutsSection:title
// that can be shown in the section selector of the gtk.ShortcutsWindow.
//
// The gtk.ShortcutsSection:max-height property can be used to influence how the
// groups in the section are distributed over pages and columns.
//
// This widget is only meant to be used with gtk.ShortcutsWindow.
type ShortcutsSection struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Widgetter         = (*ShortcutsSection)(nil)
	_ coreglib.Objector = (*ShortcutsSection)(nil)
)

func wrapShortcutsSection(obj *coreglib.Object) *ShortcutsSection {
	return &ShortcutsSection{
		Box: Box{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsSection(p uintptr) (interface{}, error) {
	return wrapShortcutsSection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ShortcutsSection) ConnectChangeCurrentPage(f func(object int) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "change-current-page", false, unsafe.Pointer(C._gotk4_gtk4_ShortcutsSection_ConnectChangeCurrentPage), f)
}
