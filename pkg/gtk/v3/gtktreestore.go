// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_tree_store_get_type()), F: marshalTreeStore},
	})
}

// TreeStore: the TreeStore object is a list model for use with a TreeView
// widget. It implements the TreeModel interface, and consequentially, can use
// all of the methods available there. It also implements the TreeSortable
// interface so it can be sorted by the view. Finally, it also implements the
// tree [drag and drop][gtk3-GtkTreeView-drag-and-drop] interfaces.
//
//
// GtkTreeStore as GtkBuildable
//
// The GtkTreeStore implementation of the Buildable interface allows to specify
// the model columns with a <columns> element that may contain multiple <column>
// elements, each specifying one model column. The “type” attribute specifies
// the data type for the column.
//
// An example of a UI Definition fragment for a tree store:
//
//    <object class="GtkTreeStore">
//      <columns>
//        <column type="gchararray"/>
//        <column type="gchararray"/>
//        <column type="gint"/>
//      </columns>
//    </object>
type TreeStore interface {
	gextras.Objector

	// Clear removes all rows from @tree_store
	Clear()
	// IsAncestor returns true if @iter is an ancestor of @descendant. That is,
	// @iter is the parent (or grandparent or great-grandparent) of @descendant.
	IsAncestor(iter *TreeIter, descendant *TreeIter) bool
	// IterDepth returns the depth of @iter. This will be 0 for anything on the
	// root level, 1 for anything down a level, etc.
	IterDepth(iter *TreeIter) int
	// IterIsValid: WARNING: This function is slow. Only use it for debugging
	// and/or testing purposes.
	//
	// Checks if the given iter is a valid iter for this TreeStore.
	IterIsValid(iter *TreeIter) bool
	// MoveAfter moves @iter in @tree_store to the position after @position.
	// @iter and @position should be in the same level. Note that this function
	// only works with unsorted stores. If @position is nil, @iter will be moved
	// to the start of the level.
	MoveAfter(iter *TreeIter, position *TreeIter)
	// MoveBefore moves @iter in @tree_store to the position before @position.
	// @iter and @position should be in the same level. Note that this function
	// only works with unsorted stores. If @position is nil, @iter will be moved
	// to the end of the level.
	MoveBefore(iter *TreeIter, position *TreeIter)
	// Remove removes @iter from @tree_store. After being removed, @iter is set
	// to the next valid row at that level, or invalidated if it previously
	// pointed to the last one.
	Remove(iter *TreeIter) bool
	// SetColumnTypes: this function is meant primarily for #GObjects that
	// inherit from TreeStore, and should only be used when constructing a new
	// TreeStore. It will not function after a row has been added, or a method
	// on the TreeModel interface is called.
	SetColumnTypes(types []externglib.Type)
	// SetValue sets the data in the cell specified by @iter and @column. The
	// type of @value must be convertible to the type of the column.
	SetValue(iter *TreeIter, column int, value externglib.Value)
	// Swap swaps @a and @b in the same level of @tree_store. Note that this
	// function only works with unsorted stores.
	Swap(a *TreeIter, b *TreeIter)
}

// TreeStoreClass implements the TreeStore interface.
type TreeStoreClass struct {
	*externglib.Object
	BuildableInterface
	TreeDragDestInterface
	TreeDragSourceInterface
	TreeModelInterface
	TreeSortableInterface
}

var _ TreeStore = (*TreeStoreClass)(nil)

func wrapTreeStore(obj *externglib.Object) TreeStore {
	return &TreeStoreClass{
		Object: obj,
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		TreeDragDestInterface: TreeDragDestInterface{
			Object: obj,
		},
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

func marshalTreeStore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeStore(obj), nil
}

// NewTreeStoreV: non vararg creation function. Used primarily by language
// bindings.
func NewTreeStoreV(types []externglib.Type) TreeStore {
	var _arg2 *C.GType
	var _arg1 C.gint
	var _cret *C.GtkTreeStore // in

	_arg1 = C.gint(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	_cret = C.gtk_tree_store_newv(_arg1, _arg2)

	var _treeStore TreeStore // out

	_treeStore = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(TreeStore)

	return _treeStore
}

// Clear removes all rows from @tree_store
func (t *TreeStoreClass) Clear() {
	var _arg0 *C.GtkTreeStore // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))

	C.gtk_tree_store_clear(_arg0)
}

// IsAncestor returns true if @iter is an ancestor of @descendant. That is,
// @iter is the parent (or grandparent or great-grandparent) of @descendant.
func (t *TreeStoreClass) IsAncestor(iter *TreeIter, descendant *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(descendant))

	_cret = C.gtk_tree_store_is_ancestor(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IterDepth returns the depth of @iter. This will be 0 for anything on the root
// level, 1 for anything down a level, etc.
func (t *TreeStoreClass) IterDepth(iter *TreeIter) int {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_iter_depth(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IterIsValid: WARNING: This function is slow. Only use it for debugging and/or
// testing purposes.
//
// Checks if the given iter is a valid iter for this TreeStore.
func (t *TreeStoreClass) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_iter_is_valid(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MoveAfter moves @iter in @tree_store to the position after @position. @iter
// and @position should be in the same level. Note that this function only works
// with unsorted stores. If @position is nil, @iter will be moved to the start
// of the level.
func (t *TreeStoreClass) MoveAfter(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_after(_arg0, _arg1, _arg2)
}

// MoveBefore moves @iter in @tree_store to the position before @position. @iter
// and @position should be in the same level. Note that this function only works
// with unsorted stores. If @position is nil, @iter will be moved to the end of
// the level.
func (t *TreeStoreClass) MoveBefore(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_before(_arg0, _arg1, _arg2)
}

// Remove removes @iter from @tree_store. After being removed, @iter is set to
// the next valid row at that level, or invalidated if it previously pointed to
// the last one.
func (t *TreeStoreClass) Remove(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_remove(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetColumnTypes: this function is meant primarily for #GObjects that inherit
// from TreeStore, and should only be used when constructing a new TreeStore. It
// will not function after a row has been added, or a method on the TreeModel
// interface is called.
func (t *TreeStoreClass) SetColumnTypes(types []externglib.Type) {
	var _arg0 *C.GtkTreeStore // out
	var _arg2 *C.GType
	var _arg1 C.gint

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	C.gtk_tree_store_set_column_types(_arg0, _arg1, _arg2)
}

// SetValue sets the data in the cell specified by @iter and @column. The type
// of @value must be convertible to the type of the column.
func (t *TreeStoreClass) SetValue(iter *TreeIter, column int, value externglib.Value) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 C.gint          // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = C.gint(column)
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_tree_store_set_value(_arg0, _arg1, _arg2, _arg3)
}

// Swap swaps @a and @b in the same level of @tree_store. Note that this
// function only works with unsorted stores.
func (t *TreeStoreClass) Swap(a *TreeIter, b *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(a))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(b))

	C.gtk_tree_store_swap(_arg0, _arg1, _arg2)
}
