// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_model_filter_get_type()), F: marshalTreeModelFilter},
	})
}

// TreeModelFilterVisibleFunc: a function which decides whether the row
// indicated by @iter is visible.
type TreeModelFilterVisibleFunc func(model TreeModel, iter *TreeIter) (ok bool)

//export gotk4_TreeModelFilterVisibleFunc
func gotk4_TreeModelFilterVisibleFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var model TreeModel // out
	var iter *TreeIter  // out

	model = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(TreeModel)
	iter = WrapTreeIter(unsafe.Pointer(arg1))

	fn := v.(TreeModelFilterVisibleFunc)
	ok := fn(model, iter)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// TreeModelFilter: a GtkTreeModel which hides parts of an underlying tree model
//
// A TreeModelFilter is a tree model which wraps another tree model, and can do
// the following things:
//
// - Filter specific rows, based on data from a “visible column”, a column
// storing booleans indicating whether the row should be filtered or not, or
// based on the return value of a “visible function”, which gets a model, iter
// and user_data and returns a boolean indicating whether the row should be
// filtered or not.
//
// - Modify the “appearance” of the model, using a modify function. This is
// extremely powerful and allows for just changing some values and also for
// creating a completely different model based on the given child model.
//
// - Set a different root node, also known as a “virtual root”. You can pass in
// a TreePath indicating the root node for the filter at construction time.
//
// The basic API is similar to TreeModelSort. For an example on its usage, see
// the section on TreeModelSort.
//
// When using TreeModelFilter, it is important to realize that TreeModelFilter
// maintains an internal cache of all nodes which are visible in its clients.
// The cache is likely to be a subtree of the tree exposed by the child model.
// TreeModelFilter will not cache the entire child model when unnecessary to not
// compromise the caching mechanism that is exposed by the reference counting
// scheme. If the child model implements reference counting, unnecessary signals
// may not be emitted because of reference counting rule 3, see the TreeModel
// documentation. (Note that e.g. TreeStore does not implement reference
// counting and will always emit all signals, even when the receiving node is
// not visible).
//
// Because of this, limitations for possible visible functions do apply. In
// general, visible functions should only use data or properties from the node
// for which the visibility state must be determined, its siblings or its
// parents. Usually, having a dependency on the state of any child node is not
// possible, unless references are taken on these explicitly. When no such
// reference exists, no signals may be received for these child nodes (see
// reference counting rule number 3 in the TreeModel section).
//
// Determining the visibility state of a given node based on the state of its
// child nodes is a frequently occurring use case. Therefore, TreeModelFilter
// explicitly supports this. For example, when a node does not have any
// children, you might not want the node to be visible. As soon as the first row
// is added to the node’s child level (or the last row removed), the node’s
// visibility should be updated.
//
// This introduces a dependency from the node on its child nodes. In order to
// accommodate this, TreeModelFilter must make sure the necessary signals are
// received from the child model. This is achieved by building, for all nodes
// which are exposed as visible nodes to TreeModelFilter's clients, the child
// level (if any) and take a reference on the first node in this level.
// Furthermore, for every row-inserted, row-changed or row-deleted signal (also
// these which were not handled because the node was not cached),
// TreeModelFilter will check if the visibility state of any parent node has
// changed.
//
// Beware, however, that this explicit support is limited to these two cases.
// For example, if you want a node to be visible only if two nodes in a child’s
// child level (2 levels deeper) are visible, you are on your own. In this case,
// either rely on TreeStore to emit all signals because it does not implement
// reference counting, or for models that do implement reference counting,
// obtain references on these child levels yourself.
type TreeModelFilter interface {
	gextras.Objector
	TreeDragSource
	TreeModel

	// ClearCache: this function should almost never be called. It clears the
	// @filter of any cached iterators that haven’t been reffed with
	// gtk_tree_model_ref_node(). This might be useful if the child model being
	// filtered is static (and doesn’t change often) and there has been a lot of
	// unreffed access to nodes. As a side effect of this function, all unreffed
	// iters will be invalid.
	ClearCache()
	// ConvertChildIterToIter sets @filter_iter to point to the row in @filter
	// that corresponds to the row pointed at by @child_iter. If @filter_iter
	// was not set, false is returned.
	ConvertChildIterToIter(childIter *TreeIter) (TreeIter, bool)
	// ConvertChildPathToPath converts @child_path to a path relative to
	// @filter. That is, @child_path points to a path in the child model. The
	// rerturned path will point to the same row in the filtered model. If
	// @child_path isn’t a valid path on the child model or points to a row
	// which is not visible in @filter, then nil is returned.
	ConvertChildPathToPath(childPath *TreePath) *TreePath
	// ConvertIterToChildIter sets @child_iter to point to the row pointed to by
	// @filter_iter.
	ConvertIterToChildIter(filterIter *TreeIter) TreeIter
	// ConvertPathToChildPath converts @filter_path to a path on the child model
	// of @filter. That is, @filter_path points to a location in @filter. The
	// returned path will point to the same location in the model not being
	// filtered. If @filter_path does not point to a location in the child
	// model, nil is returned.
	ConvertPathToChildPath(filterPath *TreePath) *TreePath
	// Model returns a pointer to the child model of @filter.
	Model() TreeModel
	// Refilter emits ::row_changed for each row in the child model, which
	// causes the filter to re-evaluate whether a row is visible or not.
	Refilter()
	// SetVisibleColumn sets @column of the child_model to be the column where
	// @filter should look for visibility information. @columns should be a
	// column of type G_TYPE_BOOLEAN, where true means that a row is visible,
	// and false if not.
	//
	// Note that gtk_tree_model_filter_set_visible_func() or
	// gtk_tree_model_filter_set_visible_column() can only be called once for a
	// given filter model.
	SetVisibleColumn(column int)
}

// treeModelFilter implements the TreeModelFilter class.
type treeModelFilter struct {
	gextras.Objector
	TreeDragSource
	TreeModel
}

var _ TreeModelFilter = (*treeModelFilter)(nil)

// WrapTreeModelFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeModelFilter(obj *externglib.Object) TreeModelFilter {
	return treeModelFilter{
		Objector:       obj,
		TreeDragSource: WrapTreeDragSource(obj),
		TreeModel:      WrapTreeModel(obj),
	}
}

func marshalTreeModelFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeModelFilter(obj), nil
}

// ClearCache: this function should almost never be called. It clears the
// @filter of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// filtered is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (f treeModelFilter) ClearCache() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))

	C.gtk_tree_model_filter_clear_cache(_arg0)
}

// ConvertChildIterToIter sets @filter_iter to point to the row in @filter
// that corresponds to the row pointed at by @child_iter. If @filter_iter
// was not set, false is returned.
func (f treeModelFilter) ConvertChildIterToIter(childIter *TreeIter) (TreeIter, bool) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg2 *C.GtkTreeIter        // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(childIter.Native()))

	var _filterIter TreeIter
	var _cret C.gboolean // in

	_cret = C.gtk_tree_model_filter_convert_child_iter_to_iter(_arg0, _arg2, (*C.GtkTreeIter)(unsafe.Pointer(&_filterIter)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _filterIter, _ok
}

// ConvertChildPathToPath converts @child_path to a path relative to
// @filter. That is, @child_path points to a path in the child model. The
// rerturned path will point to the same row in the filtered model. If
// @child_path isn’t a valid path on the child model or points to a row
// which is not visible in @filter, then nil is returned.
func (f treeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(childPath.Native()))

	var _cret *C.GtkTreePath // in

	_cret = C.gtk_tree_model_filter_convert_child_path_to_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = WrapTreePath(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _treePath
}

// ConvertIterToChildIter sets @child_iter to point to the row pointed to by
// @filter_iter.
func (f treeModelFilter) ConvertIterToChildIter(filterIter *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg2 *C.GtkTreeIter        // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(filterIter.Native()))

	var _childIter TreeIter

	C.gtk_tree_model_filter_convert_iter_to_child_iter(_arg0, _arg2, (*C.GtkTreeIter)(unsafe.Pointer(&_childIter)))

	return _childIter
}

// ConvertPathToChildPath converts @filter_path to a path on the child model
// of @filter. That is, @filter_path points to a location in @filter. The
// returned path will point to the same location in the model not being
// filtered. If @filter_path does not point to a location in the child
// model, nil is returned.
func (f treeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(filterPath.Native()))

	var _cret *C.GtkTreePath // in

	_cret = C.gtk_tree_model_filter_convert_path_to_child_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = WrapTreePath(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _treePath
}

// Model returns a pointer to the child model of @filter.
func (f treeModelFilter) Model() TreeModel {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))

	var _cret *C.GtkTreeModel // in

	_cret = C.gtk_tree_model_filter_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(TreeModel)

	return _treeModel
}

// Refilter emits ::row_changed for each row in the child model, which
// causes the filter to re-evaluate whether a row is visible or not.
func (f treeModelFilter) Refilter() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))

	C.gtk_tree_model_filter_refilter(_arg0)
}

// SetVisibleColumn sets @column of the child_model to be the column where
// @filter should look for visibility information. @columns should be a
// column of type G_TYPE_BOOLEAN, where true means that a row is visible,
// and false if not.
//
// Note that gtk_tree_model_filter_set_visible_func() or
// gtk_tree_model_filter_set_visible_column() can only be called once for a
// given filter model.
func (f treeModelFilter) SetVisibleColumn(column int) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg1 = C.int(column)

	C.gtk_tree_model_filter_set_visible_column(_arg0, _arg1)
}
