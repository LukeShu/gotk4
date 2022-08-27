// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_TreeListModelCreateModelFunc
func _gotk4_gtk4_TreeListModelCreateModelFunc(arg1 C.gpointer, arg2 C.gpointer) (cret *C.GListModel) {
	var fn TreeListModelCreateModelFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeListModelCreateModelFunc)
	}

	var _item *coreglib.Object // out

	_item = coreglib.Take(unsafe.Pointer(arg1))

	listModel := fn(_item)

	if listModel != nil {
		cret = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(listModel).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(listModel).Native()))
	}

	return cret
}
