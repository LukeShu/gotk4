// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_TreeModelFilterClass_visible(GtkTreeModelFilter*, GtkTreeModel*, GtkTreeIter*);
// extern void _gotk4_gtk3_TreeModelFilterClass_modify(GtkTreeModelFilter*, GtkTreeModel*, GtkTreeIter*, GValue*, gint);
import "C"

// glib.Type values for gtktreemodelfilter.go.
var GTypeTreeModelFilter = coreglib.Type(C.gtk_tree_model_filter_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTreeModelFilter, F: marshalTreeModelFilter},
	})
}

// TreeModelFilterVisibleFunc: function which decides whether the row indicated
// by iter is visible.
type TreeModelFilterVisibleFunc func(model TreeModeller, iter *TreeIter) (ok bool)

//export _gotk4_gtk3_TreeModelFilterVisibleFunc
func _gotk4_gtk3_TreeModelFilterVisibleFunc(arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter, arg3 C.gpointer) (cret C.gboolean) {
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

// NewFilter creates a new TreeModel, with child_model as the child_model and
// root as the virtual root.
//
// The function takes the following parameters:
//
//    - root (optional) or NULL.
//
// The function returns the following values:
//
//    - treeModel: new TreeModel.
//
func (childModel *TreeModel) NewFilter(root *TreePath) *TreeModel {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(childModel).Native()))
	if root != nil {
		_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(root)))
	}
	*(**TreeModel)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(childModel)
	runtime.KeepAlive(root)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _treeModel
}

// TreeModelFilterOverrider contains methods that are overridable.
type TreeModelFilterOverrider interface {
	// The function takes the following parameters:
	//
	//    - childModel
	//    - iter
	//    - value
	//    - column
	//
	Modify(childModel TreeModeller, iter *TreeIter, value *coreglib.Value, column int32)
	// The function takes the following parameters:
	//
	//    - childModel
	//    - iter
	//
	// The function returns the following values:
	//
	Visible(childModel TreeModeller, iter *TreeIter) bool
}

// TreeModelFilter is a tree model which wraps another tree model, and can do
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
// reference couting rule number 3 in the TreeModel section).
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
	_ [0]func() // equal guard
	*coreglib.Object

	TreeDragSource
	TreeModel
}

var (
	_ coreglib.Objector = (*TreeModelFilter)(nil)
)

func classInitTreeModelFilterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkTreeModelFilterClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkTreeModelFilterClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Modify(childModel TreeModeller, iter *TreeIter, value *coreglib.Value, column int32)
	}); ok {
		pclass.modify = (*[0]byte)(C._gotk4_gtk3_TreeModelFilterClass_modify)
	}

	if _, ok := goval.(interface {
		Visible(childModel TreeModeller, iter *TreeIter) bool
	}); ok {
		pclass.visible = (*[0]byte)(C._gotk4_gtk3_TreeModelFilterClass_visible)
	}
}

//export _gotk4_gtk3_TreeModelFilterClass_modify
func _gotk4_gtk3_TreeModelFilterClass_modify(arg0 *C.GtkTreeModelFilter, arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter, arg3 *C.GValue, arg4 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Modify(childModel TreeModeller, iter *TreeIter, value *coreglib.Value, column int32)
	})

	var _childModel TreeModeller // out
	var _iter *TreeIter          // out
	var _value *coreglib.Value   // out
	var _column int32            // out

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
		_childModel = rv
	}
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_value = coreglib.ValueFromNative(unsafe.Pointer(arg3))
	_column = int32(arg4)

	iface.Modify(_childModel, _iter, _value, _column)
}

//export _gotk4_gtk3_TreeModelFilterClass_visible
func _gotk4_gtk3_TreeModelFilterClass_visible(arg0 *C.GtkTreeModelFilter, arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Visible(childModel TreeModeller, iter *TreeIter) bool
	})

	var _childModel TreeModeller // out
	var _iter *TreeIter          // out

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
		_childModel = rv
	}
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := iface.Visible(_childModel, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTreeModelFilter(obj *coreglib.Object) *TreeModelFilter {
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

func marshalTreeModelFilter(p uintptr) (interface{}, error) {
	return wrapTreeModelFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClearCache: this function should almost never be called. It clears the filter
// of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// filtered is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (filter *TreeModelFilter) ClearCache() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	*(**TreeModelFilter)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "TreeModelFilter").InvokeMethod("clear_cache", args[:], nil)

	runtime.KeepAlive(filter)
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
// The function returns the following values:
//
//    - treePath (optional): newly allocated TreePath, or NULL.
//
func (filter *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(childPath)))
	*(**TreeModelFilter)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeModelFilter").InvokeMethod("convert_child_path_to_path", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
// The function returns the following values:
//
//    - treePath (optional): newly allocated TreePath, or NULL.
//
func (filter *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(filterPath)))
	*(**TreeModelFilter)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeModelFilter").InvokeMethod("convert_path_to_child_path", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
//
// The function returns the following values:
//
//    - treeModel: pointer to a TreeModel.
//
func (filter *TreeModelFilter) Model() *TreeModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	*(**TreeModelFilter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeModelFilter").InvokeMethod("get_model", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filter)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))

	return _treeModel
}

// Refilter emits ::row_changed for each row in the child model, which causes
// the filter to re-evaluate whether a row is visible or not.
func (filter *TreeModelFilter) Refilter() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	*(**TreeModelFilter)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "TreeModelFilter").InvokeMethod("refilter", args[:], nil)

	runtime.KeepAlive(filter)
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
func (filter *TreeModelFilter) SetVisibleColumn(column int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(filter).Native()))
	_arg1 = C.gint(column)
	*(**TreeModelFilter)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "TreeModelFilter").InvokeMethod("set_visible_column", args[:], nil)

	runtime.KeepAlive(filter)
	runtime.KeepAlive(column)
}
