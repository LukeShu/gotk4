// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime/cgo"
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// Queue contains the public fields of a [Queue][glib-Double-ended-Queues].
type Queue struct {
	native C.GQueue
}

// Native returns the underlying C source pointer.
func (q *Queue) Native() unsafe.Pointer {
	return unsafe.Pointer(&q.native)
}

// Clear removes all the elements in @queue. If queue elements contain
// dynamically-allocated memory, they should be freed first.
func (queue *Queue) Clear() {
	var _arg0 *C.GQueue // out

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	C.g_queue_clear(_arg0)
}

// Free frees the memory allocated for the #GQueue. Only call this function if
// @queue was created with g_queue_new(). If queue elements contain
// dynamically-allocated memory, they should be freed first.
//
// If queue elements contain dynamically-allocated memory, you should either use
// g_queue_free_full() or free them manually first.
func (queue *Queue) free() {
	var _arg0 *C.GQueue // out

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	C.g_queue_free(_arg0)
}

// Length returns the number of items in @queue.
func (queue *Queue) Length() uint {
	var _arg0 *C.GQueue // out
	var _cret C.guint   // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	_cret = C.g_queue_get_length(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Index returns the position of the first element in @queue which contains
// @data.
func (queue *Queue) Index(data cgo.Handle) int {
	var _arg0 *C.GQueue       // out
	var _arg1 C.gconstpointer // out
	var _cret C.gint          // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = (C.gconstpointer)(unsafe.Pointer(data))

	_cret = C.g_queue_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Init: statically-allocated #GQueue must be initialized with this function
// before it can be used. Alternatively you can initialize it with QUEUE_INIT.
// It is not necessary to initialize queues created with g_queue_new().
func (queue *Queue) Init() {
	var _arg0 *C.GQueue // out

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	C.g_queue_init(_arg0)
}

// IsEmpty returns true if the queue is empty.
func (queue *Queue) IsEmpty() bool {
	var _arg0 *C.GQueue  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	_cret = C.g_queue_is_empty(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PeekHead returns the first element of the queue.
func (queue *Queue) PeekHead() cgo.Handle {
	var _arg0 *C.GQueue  // out
	var _cret C.gpointer // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	_cret = C.g_queue_peek_head(_arg0)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// PeekNth returns the @n'th element of @queue.
func (queue *Queue) PeekNth(n uint) cgo.Handle {
	var _arg0 *C.GQueue  // out
	var _arg1 C.guint    // out
	var _cret C.gpointer // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = C.guint(n)

	_cret = C.g_queue_peek_nth(_arg0, _arg1)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// PeekTail returns the last element of the queue.
func (queue *Queue) PeekTail() cgo.Handle {
	var _arg0 *C.GQueue  // out
	var _cret C.gpointer // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	_cret = C.g_queue_peek_tail(_arg0)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// PopHead removes the first element of the queue and returns its data.
func (queue *Queue) PopHead() cgo.Handle {
	var _arg0 *C.GQueue  // out
	var _cret C.gpointer // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	_cret = C.g_queue_pop_head(_arg0)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// PopNth removes the @n'th element of @queue and returns its data.
func (queue *Queue) PopNth(n uint) cgo.Handle {
	var _arg0 *C.GQueue  // out
	var _arg1 C.guint    // out
	var _cret C.gpointer // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = C.guint(n)

	_cret = C.g_queue_pop_nth(_arg0, _arg1)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// PopTail removes the last element of the queue and returns its data.
func (queue *Queue) PopTail() cgo.Handle {
	var _arg0 *C.GQueue  // out
	var _cret C.gpointer // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	_cret = C.g_queue_pop_tail(_arg0)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// PushHead adds a new element at the head of the queue.
func (queue *Queue) PushHead(data cgo.Handle) {
	var _arg0 *C.GQueue  // out
	var _arg1 C.gpointer // out

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = (C.gpointer)(unsafe.Pointer(data))

	C.g_queue_push_head(_arg0, _arg1)
}

// PushNth inserts a new element into @queue at the given position.
func (queue *Queue) PushNth(data cgo.Handle, n int) {
	var _arg0 *C.GQueue  // out
	var _arg1 C.gpointer // out
	var _arg2 C.gint     // out

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = (C.gpointer)(unsafe.Pointer(data))
	_arg2 = C.gint(n)

	C.g_queue_push_nth(_arg0, _arg1, _arg2)
}

// PushTail adds a new element at the tail of the queue.
func (queue *Queue) PushTail(data cgo.Handle) {
	var _arg0 *C.GQueue  // out
	var _arg1 C.gpointer // out

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = (C.gpointer)(unsafe.Pointer(data))

	C.g_queue_push_tail(_arg0, _arg1)
}

// Remove removes the first element in @queue that contains @data.
func (queue *Queue) Remove(data cgo.Handle) bool {
	var _arg0 *C.GQueue       // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = (C.gconstpointer)(unsafe.Pointer(data))

	_cret = C.g_queue_remove(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveAll: remove all elements whose data equals @data from @queue.
func (queue *Queue) RemoveAll(data cgo.Handle) uint {
	var _arg0 *C.GQueue       // out
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))
	_arg1 = (C.gconstpointer)(unsafe.Pointer(data))

	_cret = C.g_queue_remove_all(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Reverse reverses the order of the items in @queue.
func (queue *Queue) Reverse() {
	var _arg0 *C.GQueue // out

	_arg0 = (*C.GQueue)(unsafe.Pointer(queue))

	C.g_queue_reverse(_arg0)
}
