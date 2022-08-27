// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_TreeSelectionForEachFunc
func _gotk4_gtk4_TreeSelectionForEachFunc(arg1 *C.GtkTreeModel, arg2 *C.GtkTreePath, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	var fn TreeSelectionForEachFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeSelectionForEachFunc)
	}

	var _model TreeModeller // out
	var _path *TreePath     // out
	var _iter *TreeIter     // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	fn(_model, _path, _iter)
}

//export _gotk4_gtk4_TreeSelectionFunc
func _gotk4_gtk4_TreeSelectionFunc(arg1 *C.GtkTreeSelection, arg2 *C.GtkTreeModel, arg3 *C.GtkTreePath, arg4 C.gboolean, arg5 C.gpointer) (cret C.gboolean) {
	var fn TreeSelectionFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeSelectionFunc)
	}

	var _selection *TreeSelection   // out
	var _model TreeModeller         // out
	var _path *TreePath             // out
	var _pathCurrentlySelected bool // out

	_selection = wrapTreeSelection(coreglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	if arg4 != 0 {
		_pathCurrentlySelected = true
	}

	ok := fn(_selection, _model, _path, _pathCurrentlySelected)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_TreeSelection_ConnectChanged
func _gotk4_gtk4_TreeSelection_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
