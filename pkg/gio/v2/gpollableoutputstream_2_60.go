// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// GPollableReturn _gotk4_gio2_PollableOutputStream_virtual_writev_nonblocking(void* fnptr, GPollableOutputStream* arg0, GOutputVector* arg1, gsize arg2, gsize* arg3, GError** arg4) {
//   return ((GPollableReturn (*)(GPollableOutputStream*, GOutputVector*, gsize, gsize*, GError**))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
import "C"

// WritevNonblocking attempts to write the bytes contained in the n_vectors
// vectors to stream, as with g_output_stream_writev(). If stream is not
// currently writable, this will immediately return
// G_POLLABLE_RETURN_WOULD_BLOCK, and you can use
// g_pollable_output_stream_create_source() to create a #GSource that will be
// triggered when stream is writable. error will *not* be set in that case.
//
// Note that since this method never blocks, you cannot actually use cancellable
// to cancel it. However, it will return an error if cancellable has already
// been cancelled when you call, which may happen if you call this method after
// a source triggers due to having been cancelled.
//
// Also note that if G_POLLABLE_RETURN_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same vectors and n_vectors
// in the next write call.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - vectors: buffer containing the Vectors to write.
//
// The function returns the following values:
//
//    - bytesWritten (optional): location to store the number of bytes that were
//      written to the stream.
//    - pollableReturn: G_POLLABLE_RETURN_OK on success,
//      G_POLLABLE_RETURN_WOULD_BLOCK if the stream is not currently writable
//      (and error is *not* set), or G_POLLABLE_RETURN_FAILED if there was an
//      error in which case error will be set.
//
func (stream *PollableOutputStream) WritevNonblocking(ctx context.Context, vectors []OutputVector) (uint, PollableReturn, error) {
	var _arg0 *C.GPollableOutputStream // out
	var _arg4 *C.GCancellable          // out
	var _arg1 *C.GOutputVector         // out
	var _arg2 C.gsize
	var _arg3 C.gsize           // in
	var _cret C.GPollableReturn // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.gsize)(len(vectors))
	_arg1 = (*C.GOutputVector)(C.calloc(C.size_t(len(vectors)), C.size_t(C.sizeof_GOutputVector)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GOutputVector)(_arg1), len(vectors))
		for i := range vectors {
			out[i] = *(*C.GOutputVector)(gextras.StructNative(unsafe.Pointer((&vectors[i]))))
		}
	}

	_cret = C.g_pollable_output_stream_writev_nonblocking(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(vectors)

	var _bytesWritten uint             // out
	var _pollableReturn PollableReturn // out
	var _goerr error                   // out

	_bytesWritten = uint(_arg3)
	_pollableReturn = PollableReturn(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _pollableReturn, _goerr
}

// writevNonblocking attempts to write the bytes contained in the n_vectors
// vectors to stream, as with g_output_stream_writev(). If stream is not
// currently writable, this will immediately return
// G_POLLABLE_RETURN_WOULD_BLOCK, and you can use
// g_pollable_output_stream_create_source() to create a #GSource that will be
// triggered when stream is writable. error will *not* be set in that case.
//
// Note that since this method never blocks, you cannot actually use cancellable
// to cancel it. However, it will return an error if cancellable has already
// been cancelled when you call, which may happen if you call this method after
// a source triggers due to having been cancelled.
//
// Also note that if G_POLLABLE_RETURN_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same vectors and n_vectors
// in the next write call.
//
// The function takes the following parameters:
//
//    - vectors: buffer containing the Vectors to write.
//
// The function returns the following values:
//
//    - bytesWritten (optional): location to store the number of bytes that were
//      written to the stream.
//    - pollableReturn: G_POLLABLE_RETURN_OK on success,
//      G_POLLABLE_RETURN_WOULD_BLOCK if the stream is not currently writable
//      (and error is *not* set), or G_POLLABLE_RETURN_FAILED if there was an
//      error in which case error will be set.
//
func (stream *PollableOutputStream) writevNonblocking(vectors []OutputVector) (uint, PollableReturn, error) {
	gclass := (*C.GPollableOutputStreamInterface)(coreglib.PeekParentClass(stream))
	fnarg := gclass.writev_nonblocking

	var _arg0 *C.GPollableOutputStream // out
	var _arg1 *C.GOutputVector         // out
	var _arg2 C.gsize
	var _arg3 C.gsize           // in
	var _cret C.GPollableReturn // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg2 = (C.gsize)(len(vectors))
	_arg1 = (*C.GOutputVector)(C.calloc(C.size_t(len(vectors)), C.size_t(C.sizeof_GOutputVector)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GOutputVector)(_arg1), len(vectors))
		for i := range vectors {
			out[i] = *(*C.GOutputVector)(gextras.StructNative(unsafe.Pointer((&vectors[i]))))
		}
	}

	_cret = C._gotk4_gio2_PollableOutputStream_virtual_writev_nonblocking(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, &_arg3, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(vectors)

	var _bytesWritten uint             // out
	var _pollableReturn PollableReturn // out
	var _goerr error                   // out

	_bytesWritten = uint(_arg3)
	_pollableReturn = PollableReturn(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _pollableReturn, _goerr
}
