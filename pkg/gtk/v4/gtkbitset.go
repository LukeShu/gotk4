// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// BitsetIter: opaque, stack-allocated struct for iterating over the elements of
// a GtkBitset.
//
// Before a GtkBitsetIter can be used, it needs to be initialized with
// gtk.BitsetIter().InitFirst, gtk.BitsetIter().InitLast or
// gtk.BitsetIter().InitAt.
type BitsetIter struct {
	nocopy gextras.NoCopy
	native *C.GtkBitsetIter
}

// Value gets the current value that iter points to.
//
// If iter is not valid and gtk.BitsetIter.IsValid() returns FALSE, this
// function returns 0.
func (iter *BitsetIter) Value() uint {
	var _arg0 *C.GtkBitsetIter // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkBitsetIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_bitset_iter_get_value(_arg0)

	runtime.KeepAlive(iter)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IsValid checks if iter points to a valid value.
func (iter *BitsetIter) IsValid() bool {
	var _arg0 *C.GtkBitsetIter // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkBitsetIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_bitset_iter_is_valid(_arg0)

	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Next moves iter to the next value in the set.
//
// If it was already pointing to the last value in the set, FALSE is returned
// and iter is invalidated.
func (iter *BitsetIter) Next() (uint, bool) {
	var _arg0 *C.GtkBitsetIter // out
	var _arg1 C.guint          // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkBitsetIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_bitset_iter_next(_arg0, &_arg1)

	runtime.KeepAlive(iter)

	var _value uint // out
	var _ok bool    // out

	_value = uint(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// Previous moves iter to the previous value in the set.
//
// If it was already pointing to the first value in the set, FALSE is returned
// and iter is invalidated.
func (iter *BitsetIter) Previous() (uint, bool) {
	var _arg0 *C.GtkBitsetIter // out
	var _arg1 C.guint          // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkBitsetIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_bitset_iter_previous(_arg0, &_arg1)

	runtime.KeepAlive(iter)

	var _value uint // out
	var _ok bool    // out

	_value = uint(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// BitsetIterInitAt initializes iter to point to target.
//
// If target is not found, finds the next value after it. If no value >= target
// exists in set, this function returns FALSE.
func BitsetIterInitAt(set *Bitset, target uint) (BitsetIter, uint, bool) {
	var _arg1 C.GtkBitsetIter // in
	var _arg2 *C.GtkBitset    // out
	var _arg3 C.guint         // out
	var _arg4 C.guint         // in
	var _cret C.gboolean      // in

	_arg2 = (*C.GtkBitset)(gextras.StructNative(unsafe.Pointer(set)))
	_arg3 = C.guint(target)

	_cret = C.gtk_bitset_iter_init_at(&_arg1, _arg2, _arg3, &_arg4)

	runtime.KeepAlive(set)
	runtime.KeepAlive(target)

	var _iter BitsetIter // out
	var _value uint      // out
	var _ok bool         // out

	_iter = *(*BitsetIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_value = uint(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _iter, _value, _ok
}

// BitsetIterInitFirst initializes an iterator for set and points it to the
// first value in set.
//
// If set is empty, FALSE is returned and value is set to G_MAXUINT.
func BitsetIterInitFirst(set *Bitset) (BitsetIter, uint, bool) {
	var _arg1 C.GtkBitsetIter // in
	var _arg2 *C.GtkBitset    // out
	var _arg3 C.guint         // in
	var _cret C.gboolean      // in

	_arg2 = (*C.GtkBitset)(gextras.StructNative(unsafe.Pointer(set)))

	_cret = C.gtk_bitset_iter_init_first(&_arg1, _arg2, &_arg3)

	runtime.KeepAlive(set)

	var _iter BitsetIter // out
	var _value uint      // out
	var _ok bool         // out

	_iter = *(*BitsetIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_value = uint(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _iter, _value, _ok
}

// BitsetIterInitLast initializes an iterator for set and points it to the last
// value in set.
//
// If set is empty, FALSE is returned.
func BitsetIterInitLast(set *Bitset) (BitsetIter, uint, bool) {
	var _arg1 C.GtkBitsetIter // in
	var _arg2 *C.GtkBitset    // out
	var _arg3 C.guint         // in
	var _cret C.gboolean      // in

	_arg2 = (*C.GtkBitset)(gextras.StructNative(unsafe.Pointer(set)))

	_cret = C.gtk_bitset_iter_init_last(&_arg1, _arg2, &_arg3)

	runtime.KeepAlive(set)

	var _iter BitsetIter // out
	var _value uint      // out
	var _ok bool         // out

	_iter = *(*BitsetIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	_value = uint(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _iter, _value, _ok
}
