// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_model_sort_get_type()), F: marshalTreeModelSort},
	})
}

// TreeModelSort: the TreeModelSort is a model which implements the TreeSortable
// interface. It does not hold any data itself, but rather is created with a
// child model and proxies its data. It has identical column types to this child
// model, and the changes in the child are propagated. The primary purpose of
// this model is to provide a way to sort a different model without modifying
// it. Note that the sort function used by TreeModelSort is not guaranteed to be
// stable.
//
// The use of this is best demonstrated through an example. In the following
// sample code we create two TreeView widgets each with a view of the same data.
// As the model is wrapped here by a TreeModelSort, the two TreeViews can each
// sort their view of the data without affecting the other. By contrast, if we
// simply put the same model in each widget, then sorting the first would sort
// the second.
//
// Using a TreeModelSort
//
//    void
//    selection_changed (GtkTreeSelection *selection, gpointer data)
//    {
//      GtkTreeModel *sort_model = NULL;
//      GtkTreeModel *child_model;
//      GtkTreeIter sort_iter;
//      GtkTreeIter child_iter;
//      char *some_data = NULL;
//      char *modified_data;
//
//      // Get the current selected row and the model.
//      if (! gtk_tree_selection_get_selected (selection,
//                                             &sort_model,
//                                             &sort_iter))
//        return;
//
//      // Look up the current value on the selected row and get
//      // a new value to change it to.
//      gtk_tree_model_get (GTK_TREE_MODEL (sort_model), &sort_iter,
//                          COLUMN_1, &some_data,
//                          -1);
//
//      modified_data = change_the_data (some_data);
//      g_free (some_data);
//
//      // Get an iterator on the child model, instead of the sort model.
//      gtk_tree_model_sort_convert_iter_to_child_iter (GTK_TREE_MODEL_SORT (sort_model),
//                                                      &child_iter,
//                                                      &sort_iter);
//
//      // Get the child model and change the value of the row. In this
//      // example, the child model is a GtkListStore. It could be any other
//      // type of model, though.
//      child_model = gtk_tree_model_sort_get_model (GTK_TREE_MODEL_SORT (sort_model));
//      gtk_list_store_set (GTK_LIST_STORE (child_model), &child_iter,
//                          COLUMN_1, &modified_data,
//                          -1);
//      g_free (modified_data);
//    }
type TreeModelSort interface {
	gextras.Objector

	// ClearCache: this function should almost never be called. It clears the
	// @tree_model_sort of any cached iterators that haven’t been reffed with
	// gtk_tree_model_ref_node(). This might be useful if the child model being
	// sorted is static (and doesn’t change often) and there has been a lot of
	// unreffed access to nodes. As a side effect of this function, all unreffed
	// iters will be invalid.
	ClearCache()
	// ConvertChildPathToPath converts @child_path to a path relative to
	// @tree_model_sort. That is, @child_path points to a path in the child
	// model. The returned path will point to the same row in the sorted model.
	// If @child_path isn’t a valid path on the child model, then nil is
	// returned.
	ConvertChildPathToPath(childPath *TreePath) *TreePath
	// ConvertPathToChildPath converts @sorted_path to a path on the child model
	// of @tree_model_sort. That is, @sorted_path points to a location in
	// @tree_model_sort. The returned path will point to the same location in
	// the model not being sorted. If @sorted_path does not point to a location
	// in the child model, nil is returned.
	ConvertPathToChildPath(sortedPath *TreePath) *TreePath
	// Model returns the model the TreeModelSort is sorting.
	Model() *TreeModelInterface
	// IterIsValid: > This function is slow. Only use it for debugging and/or
	// testing > purposes.
	//
	// Checks if the given iter is a valid iter for this TreeModelSort.
	IterIsValid(iter *TreeIter) bool
	// ResetDefaultSortFunc: this resets the default sort function to be in the
	// “unsorted” state. That is, it is in the same order as the child model. It
	// will re-sort the model to be in the same order as the child model only if
	// the TreeModelSort is in “unsorted” state.
	ResetDefaultSortFunc()
}

// TreeModelSortClass implements the TreeModelSort interface.
type TreeModelSortClass struct {
	*externglib.Object
	TreeDragSourceInterface
	TreeModelInterface
	TreeSortableInterface
}

var _ TreeModelSort = (*TreeModelSortClass)(nil)

func wrapTreeModelSort(obj *externglib.Object) TreeModelSort {
	return &TreeModelSortClass{
		Object: obj,
		TreeDragSourceInterface: TreeDragSourceInterface{
			Object: obj,
		},
		TreeModelInterface: TreeModelInterface{
			Object: obj,
		},
		TreeSortableInterface: TreeSortableInterface{
			TreeModelInterface: TreeModelInterface{
				Object: obj,
			},
		},
	}
}

func marshalTreeModelSort(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeModelSort(obj), nil
}

// NewTreeModelSortWithModel creates a new TreeModelSort, with @child_model as
// the child model.
func NewTreeModelSortWithModel(childModel TreeModel) *TreeModelSortClass {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkTreeModel // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer((&childModel).Native()))

	_cret = C.gtk_tree_model_sort_new_with_model(_arg1)

	var _treeModelSort *TreeModelSortClass // out

	_treeModelSort = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TreeModelSortClass)

	return _treeModelSort
}

// ClearCache: this function should almost never be called. It clears the
// @tree_model_sort of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// sorted is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (t *TreeModelSortClass) ClearCache() {
	var _arg0 *C.GtkTreeModelSort // out

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer((&t).Native()))

	C.gtk_tree_model_sort_clear_cache(_arg0)
}

// ConvertChildPathToPath converts @child_path to a path relative to
// @tree_model_sort. That is, @child_path points to a path in the child model.
// The returned path will point to the same row in the sorted model. If
// @child_path isn’t a valid path on the child model, then nil is returned.
func (t *TreeModelSortClass) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelSort // out
	var _arg1 *C.GtkTreePath      // out
	var _cret *C.GtkTreePath      // in

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer((&t).Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(childPath))

	_cret = C.gtk_tree_model_sort_convert_child_path_to_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
		C.free(unsafe.Pointer(v))
	})

	return _treePath
}

// ConvertPathToChildPath converts @sorted_path to a path on the child model of
// @tree_model_sort. That is, @sorted_path points to a location in
// @tree_model_sort. The returned path will point to the same location in the
// model not being sorted. If @sorted_path does not point to a location in the
// child model, nil is returned.
func (t *TreeModelSortClass) ConvertPathToChildPath(sortedPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelSort // out
	var _arg1 *C.GtkTreePath      // out
	var _cret *C.GtkTreePath      // in

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer((&t).Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(sortedPath))

	_cret = C.gtk_tree_model_sort_convert_path_to_child_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
		C.free(unsafe.Pointer(v))
	})

	return _treePath
}

// Model returns the model the TreeModelSort is sorting.
func (t *TreeModelSortClass) Model() *TreeModelInterface {
	var _arg0 *C.GtkTreeModelSort // out
	var _cret *C.GtkTreeModel     // in

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer((&t).Native()))

	_cret = C.gtk_tree_model_sort_get_model(_arg0)

	var _treeModel *TreeModelInterface // out

	_treeModel = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*TreeModelInterface)

	return _treeModel
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this TreeModelSort.
func (t *TreeModelSortClass) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeModelSort // out
	var _arg1 *C.GtkTreeIter      // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer((&t).Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_model_sort_iter_is_valid(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResetDefaultSortFunc: this resets the default sort function to be in the
// “unsorted” state. That is, it is in the same order as the child model. It
// will re-sort the model to be in the same order as the child model only if the
// TreeModelSort is in “unsorted” state.
func (t *TreeModelSortClass) ResetDefaultSortFunc() {
	var _arg0 *C.GtkTreeModelSort // out

	_arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer((&t).Native()))

	C.gtk_tree_model_sort_reset_default_sort_func(_arg0)
}
