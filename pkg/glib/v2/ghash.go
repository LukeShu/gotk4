// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// glib.Type values for ghash.go.
var GTypeHashTable = coreglib.Type(C.g_hash_table_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeHashTable, F: marshalHashTable},
	})
}

// HashTable struct is an opaque data structure to represent a [Hash
// Table][glib-Hash-Tables]. It should only be accessed via the following
// functions.
//
// An instance of this type is always passed by reference.
type HashTable struct {
	*hashTable
}

// hashTable is the struct that's finalized.
type hashTable struct {
	native *C.GHashTable
}

func marshalHashTable(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &HashTable{&hashTable{(*C.GHashTable)(b)}}, nil
}

// HashTableIter structure represents an iterator that can be used to iterate
// over the elements of a Table. GHashTableIter structures are typically
// allocated on the stack and then initialized with g_hash_table_iter_init().
//
// The iteration order of a TableIter over the keys/values in a hash table is
// not defined.
//
// An instance of this type is always passed by reference.
type HashTableIter struct {
	*hashTableIter
}

// hashTableIter is the struct that's finalized.
type hashTableIter struct {
	native *C.GHashTableIter
}

// Remove removes the key/value pair currently pointed to by the iterator from
// its associated Table. Can only be called after g_hash_table_iter_next()
// returned TRUE, and cannot be called more than once for the same key/value
// pair.
//
// If the Table was created using g_hash_table_new_full(), the key and value are
// freed using the supplied destroy functions, otherwise you have to make sure
// that any dynamically allocated values are freed yourself.
//
// It is safe to continue iterating the Table afterward:
//
//    while (g_hash_table_iter_next (&iter, &key, &value))
//      {
//        if (condition)
//          g_hash_table_iter_remove (&iter);
//      }.
func (iter *HashTableIter) Remove() {
	var _arg0 *C.GHashTableIter // out

	_arg0 = (*C.GHashTableIter)(gextras.StructNative(unsafe.Pointer(iter)))

	C.g_hash_table_iter_remove(_arg0)
	runtime.KeepAlive(iter)
}

// Steal removes the key/value pair currently pointed to by the iterator from
// its associated Table, without calling the key and value destroy functions.
// Can only be called after g_hash_table_iter_next() returned TRUE, and cannot
// be called more than once for the same key/value pair.
func (iter *HashTableIter) Steal() {
	var _arg0 *C.GHashTableIter // out

	_arg0 = (*C.GHashTableIter)(gextras.StructNative(unsafe.Pointer(iter)))

	C.g_hash_table_iter_steal(_arg0)
	runtime.KeepAlive(iter)
}
