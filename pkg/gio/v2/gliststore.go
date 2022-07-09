// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gint _gotk4_gio2_CompareDataFunc(gpointer, gpointer, gpointer);
// extern gint _gotk4_glib2_CompareDataFunc(gpointer, gpointer, gpointer);
import "C"

// GTypeListStore returns the GType for the type ListStore.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeListStore() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "ListStore").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalListStore)
	return gtype
}

// ListStoreOverrider contains methods that are overridable.
type ListStoreOverrider interface {
}

// ListStore is a simple implementation of Model that stores all items in
// memory.
//
// It provides insertions, deletions, and lookups in logarithmic time with a
// fast path for the common case of iterating the list linearly.
type ListStore struct {
	_ [0]func() // equal guard
	*coreglib.Object

	ListModel
}

var (
	_ coreglib.Objector = (*ListStore)(nil)
)

func classInitListStorer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapListStore(obj *coreglib.Object) *ListStore {
	return &ListStore{
		Object: obj,
		ListModel: ListModel{
			Object: obj,
		},
	}
}

func marshalListStore(p uintptr) (interface{}, error) {
	return wrapListStore(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Append appends item to store. item must be of type Store:item-type.
//
// This function takes a ref on item.
//
// Use g_list_store_splice() to append multiple items at the same time
// efficiently.
//
// The function takes the following parameters:
//
//    - item: new item.
//
func (store *ListStore) Append(item *coreglib.Object) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(item.Native()))

	girepository.MustFind("Gio", "ListStore").InvokeMethod("append", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(item)
}

// Find looks up the given item in the list store by looping over the items
// until the first occurrence of item. If item was not found, then position will
// not be set, and this method will return FALSE.
//
// If you need to compare the two items with a custom comparison function, use
// g_list_store_find_with_equal_func() with a custom Func instead.
//
// The function takes the following parameters:
//
//    - item: item.
//
// The function returns the following values:
//
//    - position (optional): first position of item, if it was found.
//    - ok: whether store contains item. If it was found, position will be set to
//      the position where item occurred for the first time.
//
func (store *ListStore) Find(item *coreglib.Object) (uint32, bool) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(item.Native()))

	_gret := girepository.MustFind("Gio", "ListStore").InvokeMethod("find", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(store)
	runtime.KeepAlive(item)

	var _position uint32 // out
	var _ok bool         // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_position = *(*uint32)(unsafe.Pointer(_outs[0]))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _position, _ok
}

// Insert inserts item into store at position. item must be of type
// Store:item-type or derived from it. position must be smaller than the length
// of the list, or equal to it to append.
//
// This function takes a ref on item.
//
// Use g_list_store_splice() to insert multiple items at the same time
// efficiently.
//
// The function takes the following parameters:
//
//    - position at which to insert the new item.
//    - item: new item.
//
func (store *ListStore) Insert(position uint32, item *coreglib.Object) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(*C.gpointer)(unsafe.Pointer(&_args[2])) = C.gpointer(unsafe.Pointer(item.Native()))

	girepository.MustFind("Gio", "ListStore").InvokeMethod("insert", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(position)
	runtime.KeepAlive(item)
}

// InsertSorted inserts item into store at a position to be determined by the
// compare_func.
//
// The list must already be sorted before calling this function or the result is
// undefined. Usually you would approach this by only ever inserting items by
// way of this function.
//
// This function takes a ref on item.
//
// The function takes the following parameters:
//
//    - item: new item.
//    - compareFunc: pairwise comparison function for sorting.
//
// The function returns the following values:
//
//    - guint: position at which item was inserted.
//
func (store *ListStore) InsertSorted(item *coreglib.Object, compareFunc glib.CompareDataFunc) uint32 {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(unsafe.Pointer(item.Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[2])) = (*[0]byte)(C._gotk4_glib2_CompareDataFunc)
	_args[3] = C.gpointer(gbox.Assign(compareFunc))
	defer gbox.Delete(uintptr(_args[3]))

	_gret := girepository.MustFind("Gio", "ListStore").InvokeMethod("insert_sorted", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(store)
	runtime.KeepAlive(item)
	runtime.KeepAlive(compareFunc)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Remove removes the item from store that is at position. position must be
// smaller than the current length of the list.
//
// Use g_list_store_splice() to remove multiple items at the same time
// efficiently.
//
// The function takes the following parameters:
//
//    - position of the item that is to be removed.
//
func (store *ListStore) Remove(position uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	girepository.MustFind("Gio", "ListStore").InvokeMethod("remove", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(position)
}

// RemoveAll removes all items from store.
func (store *ListStore) RemoveAll() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))

	girepository.MustFind("Gio", "ListStore").InvokeMethod("remove_all", _args[:], nil)

	runtime.KeepAlive(store)
}

// Sort the items in store according to compare_func.
//
// The function takes the following parameters:
//
//    - compareFunc: pairwise comparison function for sorting.
//
func (store *ListStore) Sort(compareFunc glib.CompareDataFunc) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_glib2_CompareDataFunc)
	_args[2] = C.gpointer(gbox.Assign(compareFunc))
	defer gbox.Delete(uintptr(_args[2]))

	girepository.MustFind("Gio", "ListStore").InvokeMethod("sort", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(compareFunc)
}

// Splice changes store by removing n_removals items and adding n_additions
// items to it. additions must contain n_additions items of type
// Store:item-type. NULL is not permitted.
//
// This function is more efficient than g_list_store_insert() and
// g_list_store_remove(), because it only emits Model::items-changed once for
// the change.
//
// This function takes a ref on each item in additions.
//
// The parameters position and n_removals must be correct (ie: position +
// n_removals must be less than or equal to the length of the list at the time
// this function is called).
//
// The function takes the following parameters:
//
//    - position at which to make the change.
//    - nRemovals: number of items to remove.
//    - additions items to add.
//
func (store *ListStore) Splice(position, nRemovals uint32, additions []*coreglib.Object) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(store).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(nRemovals)
	*(*C.guint)(unsafe.Pointer(&_args[4])) = (C.guint)(len(additions))
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(C.calloc(C.size_t(len(additions)), C.size_t(C.sizeof_gpointer)))
	defer C.free(unsafe.Pointer(_args[3]))
	{
		out := unsafe.Slice((*C.gpointer)(*(**C.void)(unsafe.Pointer(&_args[3]))), len(additions))
		for i := range additions {
			*(*C.gpointer)(unsafe.Pointer(&out[i])) = C.gpointer(unsafe.Pointer(additions[i].Native()))
		}
	}

	girepository.MustFind("Gio", "ListStore").InvokeMethod("splice", _args[:], nil)

	runtime.KeepAlive(store)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nRemovals)
	runtime.KeepAlive(additions)
}
