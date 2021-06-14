// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_pollable_input_stream_get_type()), F: marshalPollableInputStream},
	})
}

// PollableInputStreamOverrider contains methods that are overridable. This
// interface is a subset of the interface PollableInputStream.
type PollableInputStreamOverrider interface {
	// CanPoll checks if @stream is actually pollable. Some classes may
	// implement InputStream but have only certain instances of that class be
	// pollable. If this method returns false, then the behavior of other
	// InputStream methods is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	CanPoll() bool
	// IsReadable checks if @stream can be read.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_input_stream_read() after
	// this returns true would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_input_stream_read_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	IsReadable() bool
	// ReadNonblocking attempts to read up to @count bytes from @stream into
	// @buffer, as with g_input_stream_read(). If @stream is not currently
	// readable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you
	// can use g_pollable_input_stream_create_source() to create a #GSource that
	// will be triggered when @stream is readable.
	//
	// Note that since this method never blocks, you cannot actually use
	// @cancellable to cancel it. However, it will return an error if
	// @cancellable has already been cancelled when you call, which may happen
	// if you call this method after a source triggers due to having been
	// cancelled.
	ReadNonblocking(buffer []byte) (int, error)
}

// PollableInputStream is implemented by Streams that can be polled for
// readiness to read. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
type PollableInputStream interface {
	InputStream
	PollableInputStreamOverrider
}

// pollableInputStream implements the PollableInputStream interface.
type pollableInputStream struct {
	InputStream
}

var _ PollableInputStream = (*pollableInputStream)(nil)

// WrapPollableInputStream wraps a GObject to a type that implements interface
// PollableInputStream. It is primarily used internally.
func WrapPollableInputStream(obj *externglib.Object) PollableInputStream {
	return PollableInputStream{
		InputStream: WrapInputStream(obj),
	}
}

func marshalPollableInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPollableInputStream(obj), nil
}

// CanPoll checks if @stream is actually pollable. Some classes may
// implement InputStream but have only certain instances of that class be
// pollable. If this method returns false, then the behavior of other
// InputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant; a
// stream cannot switch from pollable to non-pollable or vice versa.
func (s pollableInputStream) CanPoll() bool {
	var _arg0 *C.GPollableInputStream // out

	_arg0 = (*C.GPollableInputStream)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.g_pollable_input_stream_can_poll(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsReadable checks if @stream can be read.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_input_stream_read() after
// this returns true would still block. To guarantee non-blocking behavior,
// you should always use g_pollable_input_stream_read_nonblocking(), which
// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
func (s pollableInputStream) IsReadable() bool {
	var _arg0 *C.GPollableInputStream // out

	_arg0 = (*C.GPollableInputStream)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.g_pollable_input_stream_is_readable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReadNonblocking attempts to read up to @count bytes from @stream into
// @buffer, as with g_input_stream_read(). If @stream is not currently
// readable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you
// can use g_pollable_input_stream_create_source() to create a #GSource that
// will be triggered when @stream is readable.
//
// Note that since this method never blocks, you cannot actually use
// @cancellable to cancel it. However, it will return an error if
// @cancellable has already been cancelled when you call, which may happen
// if you call this method after a source triggers due to having been
// cancelled.
func (s pollableInputStream) ReadNonblocking(buffer []byte, cancellable Cancellable) (int, error) {
	var _arg0 *C.GPollableInputStream // out
	var _arg1 *C.void
	var _arg2 C.gsize
	var _arg3 *C.GCancellable // out

	_arg0 = (*C.GPollableInputStream)(unsafe.Pointer(s.Native()))
	_arg2 = C.gsize(len(buffer))
	_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cret C.gssize  // in
	var _cerr *C.GError // in

	_cret = C.g_pollable_input_stream_read_nonblocking(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _gssize int  // out
	var _goerr error // out

	_gssize = (int)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gssize, _goerr
}
