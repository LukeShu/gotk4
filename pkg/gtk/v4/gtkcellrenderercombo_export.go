// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_CellRendererCombo_ConnectChanged
func _gotk4_gtk4_CellRendererCombo_ConnectChanged(arg0 C.gpointer, arg1 *C.gchar, arg2 *C.GtkTreeIter, arg3 C.guintptr) {
	var f func(pathString string, newIter *TreeIter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pathString string, newIter *TreeIter))
	}

	var _pathString string // out
	var _newIter *TreeIter // out

	_pathString = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_newIter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	f(_pathString, _newIter)
}
