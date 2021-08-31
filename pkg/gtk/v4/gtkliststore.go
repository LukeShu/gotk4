// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_store_get_type()), F: marshalListStorer},
	})
}

// ListStore: list-like data structure that can be used with the GtkTreeView
//
// The ListStore object is a list model for use with a TreeView widget. It
// implements the TreeModel interface, and consequentialy, can use all of the
// methods available there. It also implements the TreeSortable interface so it
// can be sorted by the view. Finally, it also implements the tree [drag and
// drop][gtk4-GtkTreeView-drag-and-drop] interfaces.
//
// The ListStore can accept most GObject types as a column type, though it can’t
// accept all custom types. Internally, it will keep a copy of data passed in
// (such as a string or a boxed pointer). Columns that accept #GObjects are
// handled a little differently. The ListStore will keep a reference to the
// object instead of copying the value. As a result, if the object is modified,
// it is up to the application writer to call gtk_tree_model_row_changed() to
// emit the TreeModel::row_changed signal. This most commonly affects lists with
// Textures stored.
//
// An example for creating a simple list store:
//
//    <object class="GtkListStore">
//      <columns>
//        <column type="gchararray"/>
//        <column type="gchararray"/>
//        <column type="gint"/>
//      </columns>
//      <data>
//        <row>
//          <col id="0">John</col>
//          <col id="1">Doe</col>
//          <col id="2">25</col>
//        </row>
//        <row>
//          <col id="0">Johan</col>
//          <col id="1">Dahlin</col>
//          <col id="2">50</col>
//        </row>
//      </data>
//    </object>
type ListStore struct {
	*externglib.Object

	Buildable
	TreeDragDest
	TreeDragSource
	TreeSortable
}

func wrapListStore(obj *externglib.Object) *ListStore {
	return &ListStore{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
		TreeDragDest: TreeDragDest{
			Object: obj,
		},
		TreeDragSource: TreeDragSource{
			Object: obj,
		},
		TreeSortable: TreeSortable{
			TreeModel: TreeModel{
				Object: obj,
			},
		},
	}
}

func marshalListStorer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListStore(obj), nil
}

// NewListStore: non-vararg creation function. Used primarily by language
// bindings.
func NewListStore(types []externglib.Type) *ListStore {
	var _arg2 *C.GType // out
	var _arg1 C.int
	var _cret *C.GtkListStore // in

	_arg1 = (C.int)(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	_cret = C.gtk_list_store_newv(_arg1, _arg2)
	runtime.KeepAlive(types)

	var _listStore *ListStore // out

	_listStore = wrapListStore(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _listStore
}

// Append appends a new row to list_store. iter will be changed to point to this
// new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
func (listStore *ListStore) Append() TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))

	C.gtk_list_store_append(_arg0, &_arg1)
	runtime.KeepAlive(listStore)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Clear removes all rows from the list store.
func (listStore *ListStore) Clear() {
	var _arg0 *C.GtkListStore // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))

	C.gtk_list_store_clear(_arg0)
	runtime.KeepAlive(listStore)
}

// Insert creates a new row at position. iter will be changed to point to this
// new row. If position is -1 or is larger than the number of rows on the list,
// then the new row will be appended to the list. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_list_store_set() or gtk_list_store_set_value().
func (listStore *ListStore) Insert(position int32) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 C.int           // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	_arg2 = C.int(position)

	C.gtk_list_store_insert(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(position)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertAfter inserts a new row after sibling. If sibling is NULL, then the row
// will be prepended to the beginning of the list. iter will be changed to point
// to this new row. The row will be empty after this function is called. To fill
// in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
func (listStore *ListStore) InsertAfter(sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	if sibling != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	C.gtk_list_store_insert_after(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(sibling)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertBefore inserts a new row before sibling. If sibling is NULL, then the
// row will be appended to the end of the list. iter will be changed to point to
// this new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
func (listStore *ListStore) InsertBefore(sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	if sibling != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(sibling)))
	}

	C.gtk_list_store_insert_before(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(sibling)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// InsertWithValues: variant of gtk_list_store_insert_with_values() which takes
// the columns and values as two arrays, instead of varargs.
//
// This function is mainly intended for language-bindings.
func (listStore *ListStore) InsertWithValues(position int32, columns []int32, values []externglib.Value) TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 C.int           // out
	var _arg3 *C.int          // out
	var _arg5 C.int
	var _arg4 *C.GValue // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	_arg2 = C.int(position)
	_arg5 = (C.int)(len(columns))
	if len(columns) > 0 {
		_arg3 = (*C.int)(unsafe.Pointer(&columns[0]))
	}
	_arg5 = (C.int)(len(values))
	_arg4 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice((*C.GValue)(_arg4), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_list_store_insert_with_valuesv(_arg0, &_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(position)
	runtime.KeepAlive(columns)
	runtime.KeepAlive(values)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this ListStore.
func (listStore *ListStore) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_list_store_iter_is_valid(_arg0, _arg1)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MoveAfter moves iter in store to the position after position. Note that this
// function only works with unsorted stores. If position is NULL, iter will be
// moved to the start of the list.
func (store *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(position)))
	}

	C.gtk_list_store_move_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// MoveBefore moves iter in store to the position before position. Note that
// this function only works with unsorted stores. If position is NULL, iter will
// be moved to the end of the list.
func (store *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	if position != nil {
		_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(position)))
	}

	C.gtk_list_store_move_before(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(position)
}

// Prepend prepends a new row to list_store. iter will be changed to point to
// this new row. The row will be empty after this function is called. To fill in
// values, you need to call gtk_list_store_set() or gtk_list_store_set_value().
func (listStore *ListStore) Prepend() TreeIter {
	var _arg0 *C.GtkListStore // out
	var _arg1 C.GtkTreeIter   // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))

	C.gtk_list_store_prepend(_arg0, &_arg1)
	runtime.KeepAlive(listStore)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _iter
}

// Remove removes the given row from the list store. After being removed, iter
// is set to be the next valid row, or invalidated if it pointed to the last row
// in list_store.
func (listStore *ListStore) Remove(iter *TreeIter) bool {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_list_store_remove(_arg0, _arg1)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reorder reorders store to follow the order indicated by new_order. Note that
// this function only works with unsorted stores.
func (store *ListStore) Reorder(newOrder []int32) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.int          // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(store.Native()))
	{
		var zero int32
		newOrder = append(newOrder, zero)
		_arg1 = (*C.int)(unsafe.Pointer(&newOrder[0]))
	}

	C.gtk_list_store_reorder(_arg0, _arg1)
	runtime.KeepAlive(store)
	runtime.KeepAlive(newOrder)
}

// SetColumnTypes: this function is meant primarily for #GObjects that inherit
// from ListStore, and should only be used when constructing a new ListStore. It
// will not function after a row has been added, or a method on the TreeModel
// interface is called.
func (listStore *ListStore) SetColumnTypes(types []externglib.Type) {
	var _arg0 *C.GtkListStore // out
	var _arg2 *C.GType        // out
	var _arg1 C.int

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	_arg1 = (C.int)(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.GType)(_arg2), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	C.gtk_list_store_set_column_types(_arg0, _arg1, _arg2)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(types)
}

// SetValue sets the data in the cell specified by iter and column. The type of
// value must be convertible to the type of the column.
func (listStore *ListStore) SetValue(iter *TreeIter, column int32, value *externglib.Value) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 C.int           // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg2 = C.int(column)
	_arg3 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_list_store_set_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(column)
	runtime.KeepAlive(value)
}

// Set: variant of gtk_list_store_set_valist() which takes the columns and
// values as two arrays, instead of varargs. This function is mainly intended
// for language-bindings and in case the number of columns to change is not
// known until run-time.
func (listStore *ListStore) Set(iter *TreeIter, columns []int32, values []externglib.Value) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.int          // out
	var _arg4 C.int
	var _arg3 *C.GValue // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(listStore.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg4 = (C.int)(len(columns))
	if len(columns) > 0 {
		_arg2 = (*C.int)(unsafe.Pointer(&columns[0]))
	}
	_arg4 = (C.int)(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((*C.GValue)(_arg3), len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer((&values[i]).Native()))
		}
	}

	C.gtk_list_store_set_valuesv(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(listStore)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(columns)
	runtime.KeepAlive(values)
}

// Swap swaps a and b in store. Note that this function only works with unsorted
// stores.
func (store *ListStore) Swap(a *TreeIter, b *TreeIter) {
	var _arg0 *C.GtkListStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkListStore)(unsafe.Pointer(store.Native()))
	_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(a)))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(b)))

	C.gtk_list_store_swap(_arg0, _arg1, _arg2)
	runtime.KeepAlive(store)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
