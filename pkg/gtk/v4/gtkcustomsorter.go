// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkcustomsorter.go.
var GTypeCustomSorter = coreglib.Type(C.gtk_custom_sorter_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCustomSorter, F: marshalCustomSorter},
	})
}

// CustomSorterOverrider contains methods that are overridable.
type CustomSorterOverrider interface {
}

// CustomSorter: GtkCustomSorter is a GtkSorter implementation that sorts via a
// callback function.
type CustomSorter struct {
	_ [0]func() // equal guard
	Sorter
}

var (
	_ coreglib.Objector = (*CustomSorter)(nil)
)

func classInitCustomSorterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCustomSorter(obj *coreglib.Object) *CustomSorter {
	return &CustomSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalCustomSorter(p uintptr) (interface{}, error) {
	return wrapCustomSorter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
