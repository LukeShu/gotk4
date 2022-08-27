// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_TreeModelFilterModifyFunc
func _gotk4_gtk4_TreeModelFilterModifyFunc(arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter, arg3 *C.GValue, arg4 C.int, arg5 C.gpointer) {
	var fn TreeModelFilterModifyFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeModelFilterModifyFunc)
	}

	var _model TreeModeller // out
	var _iter *TreeIter     // out
	var _column int         // out

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
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_column = int(arg4)

	value := fn(_model, _iter, _column)

	*arg3 = *(*C.GValue)(unsafe.Pointer((&value).Native()))
}

//export _gotk4_gtk4_TreeModelFilterVisibleFunc
func _gotk4_gtk4_TreeModelFilterVisibleFunc(arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter, arg3 C.gpointer) (cret C.gboolean) {
	var fn TreeModelFilterVisibleFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeModelFilterVisibleFunc)
	}

	var _model TreeModeller // out
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
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := fn(_model, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}
