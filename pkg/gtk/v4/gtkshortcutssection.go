// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_ShortcutsSection_ConnectChangeCurrentPage(gpointer, gint, guintptr);
import "C"

// glib.Type values for gtkshortcutssection.go.
var GTypeShortcutsSection = externglib.Type(C.gtk_shortcuts_section_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeShortcutsSection, F: marshalShortcutsSection},
	})
}

// ShortcutsSectionOverrider contains methods that are overridable.
type ShortcutsSectionOverrider interface {
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
	_ Widgetter           = (*ShortcutsSection)(nil)
	_ externglib.Objector = (*ShortcutsSection)(nil)
)

func classInitShortcutsSectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapShortcutsSection(obj *externglib.Object) *ShortcutsSection {
	return &ShortcutsSection{
		Box: Box{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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
	return wrapShortcutsSection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ShortcutsSection_ConnectChangeCurrentPage
func _gotk4_gtk4_ShortcutsSection_ConnectChangeCurrentPage(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) (cret C.gboolean) {
	var f func(object int) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object int) (ok bool))
	}

	var _object int // out

	_object = int(arg1)

	ok := f(_object)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func (v *ShortcutsSection) ConnectChangeCurrentPage(f func(object int) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "change-current-page", false, unsafe.Pointer(C._gotk4_gtk4_ShortcutsSection_ConnectChangeCurrentPage), f)
}
