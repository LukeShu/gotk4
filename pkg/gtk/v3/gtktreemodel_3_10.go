// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// RowsReordered emits the TreeModel::rows-reordered signal on tree_model.
//
// This should be called by models when their rows have been reordered.
//
// The function takes the following parameters:
//
//    - path pointing to the tree node whose children have been reordered.
//    - iter (optional): valid TreeIter-struct pointing to the node whose
//      children have been reordered, or NULL if the depth of path is 0.
//    - newOrder: array of integers mapping the current position of each child to
//      its old position before the re-ordering, i.e. new_order[newpos] = oldpos.
//
func (treeModel *TreeModel) RowsReordered(path *TreePath, iter *TreeIter, newOrder []int) {
	var _arg0 *C.GtkTreeModel // out
	var _arg1 *C.GtkTreePath  // out
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 *C.gint         // out
	var _arg4 C.gint

	_arg0 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(treeModel).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))
	if iter != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}
	_arg4 = (C.gint)(len(newOrder))
	_arg3 = (*C.gint)(C.calloc(C.size_t(len(newOrder)), C.size_t(C.sizeof_gint)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.gint)(_arg3), len(newOrder))
		for i := range newOrder {
			out[i] = C.gint(newOrder[i])
		}
	}

	C.gtk_tree_model_rows_reordered_with_length(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(treeModel)
	runtime.KeepAlive(path)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(newOrder)
}
