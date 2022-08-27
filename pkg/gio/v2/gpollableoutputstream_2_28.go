// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// GSource* _gotk4_gio2_PollableOutputStream_virtual_create_source(void* fnptr, GPollableOutputStream* arg0, GCancellable* arg1) {
//   return ((GSource* (*)(GPollableOutputStream*, GCancellable*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gio2_PollableOutputStream_virtual_can_poll(void* fnptr, GPollableOutputStream* arg0) {
//   return ((gboolean (*)(GPollableOutputStream*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_PollableOutputStream_virtual_is_writable(void* fnptr, GPollableOutputStream* arg0) {
//   return ((gboolean (*)(GPollableOutputStream*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypePollableOutputStream = coreglib.Type(C.g_pollable_output_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePollableOutputStream, F: marshalPollableOutputStream},
	})
}

// PollableOutputStream is implemented by Streams that can be polled for
// readiness to write. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
//
// PollableOutputStream wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PollableOutputStream struct {
	_ [0]func() // equal guard
	OutputStream
}

var (
	_ OutputStreamer = (*PollableOutputStream)(nil)
)

// PollableOutputStreamer describes PollableOutputStream's interface methods.
type PollableOutputStreamer interface {
	coreglib.Objector

	// CanPoll checks if stream is actually pollable.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when stream can be written,
	// or cancellable is triggered or an error occurs.
	CreateSource(ctx context.Context) *glib.Source
	// IsWritable checks if stream can be written.
	IsWritable() bool
	// WriteNonblocking attempts to write up to count bytes from buffer to
	// stream, as with g_output_stream_write().
	WriteNonblocking(ctx context.Context, buffer []byte) (int, error)
	// WritevNonblocking attempts to write the bytes contained in the n_vectors
	// vectors to stream, as with g_output_stream_writev().
	WritevNonblocking(ctx context.Context, vectors []OutputVector) (uint, PollableReturn, error)
}

var _ PollableOutputStreamer = (*PollableOutputStream)(nil)

func wrapPollableOutputStream(obj *coreglib.Object) *PollableOutputStream {
	return &PollableOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
	}
}

func marshalPollableOutputStream(p uintptr) (interface{}, error) {
	return wrapPollableOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CanPoll checks if stream is actually pollable. Some classes may implement
// OutputStream but have only certain instances of that class be pollable. If
// this method returns FALSE, then the behavior of other OutputStream methods is
// undefined.
//
// For any given stream, the value returned by this method is constant; a stream
// cannot switch from pollable to non-pollable or vice versa.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is pollable, FALSE if not.
//
func (stream *PollableOutputStream) CanPoll() bool {
	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_pollable_output_stream_can_poll(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CreateSource creates a #GSource that triggers when stream can be written, or
// cancellable is triggered or an error occurs. The callback on the source is of
// the SourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that the
// stream may not actually be writable even after the source triggers, so you
// should use g_pollable_output_stream_write_nonblocking() rather than
// g_output_stream_write() from the callback.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func (stream *PollableOutputStream) CreateSource(ctx context.Context) *glib.Source {
	var _arg0 *C.GPollableOutputStream // out
	var _arg1 *C.GCancellable          // out
	var _cret *C.GSource               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.g_pollable_output_stream_create_source(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_source_unref((*C.GSource)(intern.C))
		},
	)

	return _source
}

// IsWritable checks if stream can be written.
//
// Note that some stream types may not be able to implement this 100% reliably,
// and it is possible that a call to g_output_stream_write() after this returns
// TRUE would still block. To guarantee non-blocking behavior, you should always
// use g_pollable_output_stream_write_nonblocking(), which will return a
// G_IO_ERROR_WOULD_BLOCK error rather than blocking.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is writable, FALSE if not. If an error has occurred on
//      stream, this will result in g_pollable_output_stream_is_writable()
//      returning TRUE, and the next attempt to write will return the error.
//
func (stream *PollableOutputStream) IsWritable() bool {
	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_pollable_output_stream_is_writable(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// canPoll checks if stream is actually pollable. Some classes may implement
// OutputStream but have only certain instances of that class be pollable. If
// this method returns FALSE, then the behavior of other OutputStream methods is
// undefined.
//
// For any given stream, the value returned by this method is constant; a stream
// cannot switch from pollable to non-pollable or vice versa.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is pollable, FALSE if not.
//
func (stream *PollableOutputStream) canPoll() bool {
	gclass := (*C.GPollableOutputStreamInterface)(coreglib.PeekParentClass(stream))
	fnarg := gclass.can_poll

	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C._gotk4_gio2_PollableOutputStream_virtual_can_poll(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// createSource creates a #GSource that triggers when stream can be written, or
// cancellable is triggered or an error occurs. The callback on the source is of
// the SourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that the
// stream may not actually be writable even after the source triggers, so you
// should use g_pollable_output_stream_write_nonblocking() rather than
// g_output_stream_write() from the callback.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func (stream *PollableOutputStream) createSource(ctx context.Context) *glib.Source {
	gclass := (*C.GPollableOutputStreamInterface)(coreglib.PeekParentClass(stream))
	fnarg := gclass.create_source

	var _arg0 *C.GPollableOutputStream // out
	var _arg1 *C.GCancellable          // out
	var _cret *C.GSource               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C._gotk4_gio2_PollableOutputStream_virtual_create_source(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_source_unref((*C.GSource)(intern.C))
		},
	)

	return _source
}

// isWritable checks if stream can be written.
//
// Note that some stream types may not be able to implement this 100% reliably,
// and it is possible that a call to g_output_stream_write() after this returns
// TRUE would still block. To guarantee non-blocking behavior, you should always
// use g_pollable_output_stream_write_nonblocking(), which will return a
// G_IO_ERROR_WOULD_BLOCK error rather than blocking.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is writable, FALSE if not. If an error has occurred on
//      stream, this will result in g_pollable_output_stream_is_writable()
//      returning TRUE, and the next attempt to write will return the error.
//
func (stream *PollableOutputStream) isWritable() bool {
	gclass := (*C.GPollableOutputStreamInterface)(coreglib.PeekParentClass(stream))
	fnarg := gclass.is_writable

	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C._gotk4_gio2_PollableOutputStream_virtual_is_writable(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PollableOutputStreamInterface: interface for pollable output streams.
//
// The default implementation of can_poll always returns TRUE.
//
// The default implementation of write_nonblocking calls
// g_pollable_output_stream_is_writable(), and then calls
// g_output_stream_write() if it returns TRUE. This means you only need to
// override it if it is possible that your is_writable implementation may return
// TRUE when the stream is not actually writable.
//
// The default implementation of writev_nonblocking calls
// g_pollable_output_stream_write_nonblocking() for each vector, and converts
// its return value and error (if set) to a Return. You should override this
// where possible to avoid having to allocate a #GError to return
// G_IO_ERROR_WOULD_BLOCK.
//
// An instance of this type is always passed by reference.
type PollableOutputStreamInterface struct {
	*pollableOutputStreamInterface
}

// pollableOutputStreamInterface is the struct that's finalized.
type pollableOutputStreamInterface struct {
	native *C.GPollableOutputStreamInterface
}
