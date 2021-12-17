// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk4_TreeModelFilterVisibleFunc(GtkTreeModel*, GtkTreeIter*, gpointer);
// void _gotk4_gtk4_TreeModelFilterModifyFunc(GtkTreeModel*, GtkTreeIter*, GValue*, int, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_model_filter_get_type()), F: marshalTreeModelFilterer},
	})
}

// TreeModelFilterModifyFunc: function which calculates display values from raw
// values in the model. It must fill value with the display value for the column
// column in the row indicated by iter.
//
// Since this function is called for each data access, it’s not a particularly
// efficient operation.
type TreeModelFilterModifyFunc func(model TreeModeller, iter *TreeIter, column int) (value externglib.Value)

//export _gotk4_gtk4_TreeModelFilterModifyFunc
func _gotk4_gtk4_TreeModelFilterModifyFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 *C.GValue, arg3 C.int, arg4 C.gpointer) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var model TreeModeller // out
	var iter *TreeIter     // out
	var column int         // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.TreeModeller")
		}
		model = rv
	}
	iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	column = int(arg3)

	fn := v.(TreeModelFilterModifyFunc)
	value := fn(model, iter, column)

	*arg2 = *(*C.GValue)(unsafe.Pointer((&value).Native()))
}

// TreeModelFilterVisibleFunc: function which decides whether the row indicated
// by iter is visible.
type TreeModelFilterVisibleFunc func(model TreeModeller, iter *TreeIter) (ok bool)

//export _gotk4_gtk4_TreeModelFilterVisibleFunc
func _gotk4_gtk4_TreeModelFilterVisibleFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var model TreeModeller // out
	var iter *TreeIter     // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.TreeModeller")
		}
		model = rv
	}
	iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	fn := v.(TreeModelFilterVisibleFunc)
	ok := fn(model, iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// TreeModelFilterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TreeModelFilterOverrider interface {
	Modify(childModel TreeModeller, iter *TreeIter, value *externglib.Value, column int)
	Visible(childModel TreeModeller, iter *TreeIter) bool
}

// TreeModelFilter: gtkTreeModel which hides parts of an underlying tree model
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
type TreeModelFilter struct {
	*externglib.Object

	TreeDragSource
	TreeModel
}

var (
	_ externglib.Objector = (*TreeModelFilter)(nil)
)

func wrapTreeModelFilter(obj *externglib.Object) *TreeModelFilter {
	return &TreeModelFilter{
		Object: obj,
		TreeDragSource: TreeDragSource{
			Object: obj,
		},
		TreeModel: TreeModel{
			Object: obj,
		},
	}
}

func marshalTreeModelFilterer(p uintptr) (interface{}, error) {
	return wrapTreeModelFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClearCache: this function should almost never be called. It clears the filter
// of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// filtered is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (filter *TreeModelFilter) ClearCache() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_tree_model_filter_clear_cache(_arg0)
	runtime.KeepAlive(filter)
}

// ConvertChildIterToIter sets filter_iter to point to the row in filter that
// corresponds to the row pointed at by child_iter. If filter_iter was not set,
// FALSE is returned.
//
// The function takes the following parameters:
//
//    - childIter: valid TreeIter pointing to a row on the child model.
//
func (filter *TreeModelFilter) ConvertChildIterToIter(childIter *TreeIter) (*TreeIter, bool) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.GtkTreeIter         // in
	var _arg2 *C.GtkTreeIter        // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(childIter)))

	_cret = C.gtk_tree_model_filter_convert_child_iter_to_iter(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(childIter)

	var _filterIter *TreeIter // out
	var _ok bool              // out

	_filterIter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _filterIter, _ok
}

// ConvertChildPathToPath converts child_path to a path relative to filter. That
// is, child_path points to a path in the child model. The rerturned path will
// point to the same row in the filtered model. If child_path isn’t a valid path
// on the child model or points to a row which is not visible in filter, then
// NULL is returned.
//
// The function takes the following parameters:
//
//    - childPath to convert.
//
func (filter *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GtkTreePath        // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(childPath)))

	_cret = C.gtk_tree_model_filter_convert_child_path_to_path(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(childPath)

	var _treePath *TreePath // out

	if _cret != nil {
		_treePath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_treePath)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_tree_path_free((*C.GtkTreePath)(intern.C))
			},
		)
	}

	return _treePath
}

// ConvertIterToChildIter sets child_iter to point to the row pointed to by
// filter_iter.
//
// The function takes the following parameters:
//
//    - filterIter: valid TreeIter pointing to a row on filter.
//
func (filter *TreeModelFilter) ConvertIterToChildIter(filterIter *TreeIter) *TreeIter {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.GtkTreeIter         // in
	var _arg2 *C.GtkTreeIter        // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(filterIter)))

	C.gtk_tree_model_filter_convert_iter_to_child_iter(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(filterIter)

	var _childIter *TreeIter // out

	_childIter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _childIter
}

// ConvertPathToChildPath converts filter_path to a path on the child model of
// filter. That is, filter_path points to a location in filter. The returned
// path will point to the same location in the model not being filtered. If
// filter_path does not point to a location in the child model, NULL is
// returned.
//
// The function takes the following parameters:
//
//    - filterPath to convert.
//
func (filter *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GtkTreePath        // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(filterPath)))

	_cret = C.gtk_tree_model_filter_convert_path_to_child_path(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(filterPath)

	var _treePath *TreePath // out

	if _cret != nil {
		_treePath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_treePath)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_tree_path_free((*C.GtkTreePath)(intern.C))
			},
		)
	}

	return _treePath
}

// Model returns a pointer to the child model of filter.
func (filter *TreeModelFilter) Model() TreeModeller {
	var _arg0 *C.GtkTreeModelFilter // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))

	_cret = C.gtk_tree_model_filter_get_model(_arg0)
	runtime.KeepAlive(filter)

	var _treeModel TreeModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.TreeModeller")
		}
		_treeModel = rv
	}

	return _treeModel
}

// Refilter emits ::row_changed for each row in the child model, which causes
// the filter to re-evaluate whether a row is visible or not.
func (filter *TreeModelFilter) Refilter() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_tree_model_filter_refilter(_arg0)
	runtime.KeepAlive(filter)
}

// SetModifyFunc: with the n_columns and types parameters, you give an array of
// column types for this model (which will be exposed to the parent model/view).
// The func, data and destroy parameters are for specifying the modify function.
// The modify function will get called for each data access, the goal of the
// modify function is to return the data which should be displayed at the
// location specified using the parameters of the modify function.
//
// Note that gtk_tree_model_filter_set_modify_func() can only be called once for
// a given filter model.
//
// The function takes the following parameters:
//
//    - types of the columns.
//    - fn: TreeModelFilterModifyFunc.
//
func (filter *TreeModelFilter) SetModifyFunc(types []externglib.Type, fn TreeModelFilterModifyFunc) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg2 *C.GType              // out
	var _arg1 C.int
	var _arg3 C.GtkTreeModelFilterModifyFunc // out
	var _arg4 C.gpointer
	var _arg5 C.GDestroyNotify

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (C.int)(len(types))
	_arg2 = (*C.GType)(C.calloc(C.size_t(len(types)), C.size_t(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}
	_arg3 = (*[0]byte)(C._gotk4_gtk4_TreeModelFilterModifyFunc)
	_arg4 = C.gpointer(gbox.Assign(fn))
	_arg5 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_model_filter_set_modify_func(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(types)
	runtime.KeepAlive(fn)
}

// SetVisibleColumn sets column of the child_model to be the column where filter
// should look for visibility information. columns should be a column of type
// G_TYPE_BOOLEAN, where TRUE means that a row is visible, and FALSE if not.
//
// Note that gtk_tree_model_filter_set_visible_func() or
// gtk_tree_model_filter_set_visible_column() can only be called once for a
// given filter model.
//
// The function takes the following parameters:
//
//    - column which is the column containing the visible information.
//
func (filter *TreeModelFilter) SetVisibleColumn(column int) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = C.int(column)

	C.gtk_tree_model_filter_set_visible_column(_arg0, _arg1)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(column)
}

// SetVisibleFunc sets the visible function used when filtering the filter to be
// func. The function should return TRUE if the given row should be visible and
// FALSE otherwise.
//
// If the condition calculated by the function changes over time (e.g. because
// it depends on some global parameters), you must call
// gtk_tree_model_filter_refilter() to keep the visibility information of the
// model up-to-date.
//
// Note that func is called whenever a row is inserted, when it may still be
// empty. The visible function should therefore take special care of empty rows,
// like in the example below.
//
//    static gboolean
//    visible_func (GtkTreeModel *model,
//                  GtkTreeIter  *iter,
//                  gpointer      data)
//    {
//      // Visible if row is non-empty and first column is “HI”
//      char *str;
//      gboolean visible = FALSE;
//
//      gtk_tree_model_get (model, iter, 0, &str, -1);
//      if (str && strcmp (str, "HI") == 0)
//        visible = TRUE;
//      g_free (str);
//
//      return visible;
//    }
//
// Note that gtk_tree_model_filter_set_visible_func() or
// gtk_tree_model_filter_set_visible_column() can only be called once for a
// given filter model.
//
// The function takes the following parameters:
//
//    - fn the visible function.
//
func (filter *TreeModelFilter) SetVisibleFunc(fn TreeModelFilterVisibleFunc) {
	var _arg0 *C.GtkTreeModelFilter           // out
	var _arg1 C.GtkTreeModelFilterVisibleFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(filter.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_TreeModelFilterVisibleFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_tree_model_filter_set_visible_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(fn)
}
