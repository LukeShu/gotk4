// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_list_row_sorter_get_type()), F: marshalTreeListRowSorterer},
	})
}

// TreeListRowSorter: GtkTreeListRowSorter is a special-purpose sorter that will
// apply a given sorter to the levels in a tree.
//
// Here is an example for setting up a column view with a tree model and a
// GtkTreeListSorter:
//
//    column_sorter = gtk_column_view_get_sorter (view);
//    sorter = gtk_tree_list_row_sorter_new (g_object_ref (column_sorter));
//    sort_model = gtk_sort_list_model_new (tree_model, sorter);
//    selection = gtk_single_selection_new (sort_model);
//    gtk_column_view_set_model (view, G_LIST_MODEL (selection));
type TreeListRowSorter struct {
	Sorter
}

func wrapTreeListRowSorter(obj *externglib.Object) *TreeListRowSorter {
	return &TreeListRowSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalTreeListRowSorterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeListRowSorter(obj), nil
}

// NewTreeListRowSorter: create a special-purpose sorter that applies the
// sorting of sorter to the levels of a GtkTreeListModel.
//
// Note that this sorter relies on gtk.TreeListModel:passthrough being FALSE as
// it can only sort gtk.TreeListRows.
func NewTreeListRowSorter(sorter *Sorter) *TreeListRowSorter {
	var _arg1 *C.GtkSorter            // out
	var _cret *C.GtkTreeListRowSorter // in

	if sorter != nil {
		_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))
		sorter.Ref()
	}

	_cret = C.gtk_tree_list_row_sorter_new(_arg1)

	runtime.KeepAlive(sorter)

	var _treeListRowSorter *TreeListRowSorter // out

	_treeListRowSorter = wrapTreeListRowSorter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _treeListRowSorter
}

// GetSorter returns the sorter used by self.
func (self *TreeListRowSorter) GetSorter() *Sorter {
	var _arg0 *C.GtkTreeListRowSorter // out
	var _cret *C.GtkSorter            // in

	_arg0 = (*C.GtkTreeListRowSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_sorter_get_sorter(_arg0)

	runtime.KeepAlive(self)

	var _sorter *Sorter // out

	if _cret != nil {
		_sorter = wrapSorter(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _sorter
}

// SetSorter sets the sorter to use for items with the same parent.
//
// This sorter will be passed the gtk.TreeListRow:item of the tree list rows
// passed to self.
func (self *TreeListRowSorter) SetSorter(sorter *Sorter) {
	var _arg0 *C.GtkTreeListRowSorter // out
	var _arg1 *C.GtkSorter            // out

	_arg0 = (*C.GtkTreeListRowSorter)(unsafe.Pointer(self.Native()))
	if sorter != nil {
		_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))
	}

	C.gtk_tree_list_row_sorter_set_sorter(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sorter)
}
