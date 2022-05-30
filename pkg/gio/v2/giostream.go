// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GInputStream* _gotk4_gio2_IOStreamClass_get_input_stream(GIOStream*);
// extern GOutputStream* _gotk4_gio2_IOStreamClass_get_output_stream(GIOStream*);
// extern gboolean _gotk4_gio2_IOStreamClass_close_finish(GIOStream*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_IOStreamClass_close_fn(GIOStream*, GCancellable*, GError**);
import "C"

// glib.Type values for giostream.go.
var GTypeIOStream = coreglib.Type(C.g_io_stream_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeIOStream, F: marshalIOStream},
	})
}

// IOStreamOverrider contains methods that are overridable.
type IOStreamOverrider interface {
	// CloseFinish closes a stream.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	CloseFinish(result AsyncResulter) error
	// The function takes the following parameters:
	//
	CloseFn(ctx context.Context) error
	// InputStream gets the input stream for this object. This is used for
	// reading.
	//
	// The function returns the following values:
	//
	//    - inputStream owned by the OStream. Do not free.
	//
	InputStream() InputStreamer
	// OutputStream gets the output stream for this object. This is used for
	// writing.
	//
	// The function returns the following values:
	//
	//    - outputStream owned by the OStream. Do not free.
	//
	OutputStream() OutputStreamer
}

// IOStream represents an object that has both read and write streams. Generally
// the two streams act as separate input and output streams, but they share some
// common resources and state. For instance, for seekable streams, both streams
// may use the same position.
//
// Examples of OStream objects are Connection, which represents a two-way
// network connection; and IOStream, which represents a file handle opened in
// read-write mode.
//
// To do the actual reading and writing you need to get the substreams with
// g_io_stream_get_input_stream() and g_io_stream_get_output_stream().
//
// The OStream object owns the input and the output streams, not the other way
// around, so keeping the substreams alive will not keep the OStream object
// alive. If the OStream object is freed it will be closed, thus closing the
// substreams, so even if the substreams stay alive they will always return
// G_IO_ERROR_CLOSED for all operations.
//
// To close a stream use g_io_stream_close() which will close the common stream
// object and also the individual substreams. You can also close the substreams
// themselves. In most cases this only marks the substream as closed, so further
// I/O on it fails but common state in the OStream may still be open. However,
// some streams may support "half-closed" states where one direction of the
// stream is actually shut down.
//
// Operations on OStreams cannot be started while another operation on the
// OStream or its substreams is in progress. Specifically, an application can
// read from the Stream and write to the Stream simultaneously (either in
// separate threads, or as asynchronous operations in the same thread), but an
// application cannot start any OStream operation while there is a OStream,
// Stream or Stream operation in progress, and an application can’t start any
// Stream or Stream operation while there is a OStream operation in progress.
//
// This is a product of individual stream operations being associated with a
// given Context (the thread-default context at the time the operation was
// started), rather than entire streams being associated with a single Context.
//
// GIO may run operations on OStreams from other (worker) threads, and this may
// be exposed to application code in the behaviour of wrapper streams, such as
// InputStream or Connection. With such wrapper APIs, application code may only
// run operations on the base (wrapped) stream when the wrapper stream is idle.
// Note that the semantics of such operations may not be well-defined due to the
// state the wrapper stream leaves the base stream in (though they are
// guaranteed not to crash).
type IOStream struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*IOStream)(nil)
)

// IOStreamer describes types inherited from class IOStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type IOStreamer interface {
	coreglib.Objector
	baseIOStream() *IOStream
}

var _ IOStreamer = (*IOStream)(nil)

func classInitIOStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GIOStreamClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GIOStreamClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		CloseFinish(result AsyncResulter) error
	}); ok {
		pclass.close_finish = (*[0]byte)(C._gotk4_gio2_IOStreamClass_close_finish)
	}

	if _, ok := goval.(interface {
		CloseFn(ctx context.Context) error
	}); ok {
		pclass.close_fn = (*[0]byte)(C._gotk4_gio2_IOStreamClass_close_fn)
	}

	if _, ok := goval.(interface{ InputStream() InputStreamer }); ok {
		pclass.get_input_stream = (*[0]byte)(C._gotk4_gio2_IOStreamClass_get_input_stream)
	}

	if _, ok := goval.(interface{ OutputStream() OutputStreamer }); ok {
		pclass.get_output_stream = (*[0]byte)(C._gotk4_gio2_IOStreamClass_get_output_stream)
	}
}

//export _gotk4_gio2_IOStreamClass_close_finish
func _gotk4_gio2_IOStreamClass_close_finish(arg0 *C.GIOStream, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.CloseFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_IOStreamClass_close_fn
func _gotk4_gio2_IOStreamClass_close_fn(arg0 *C.GIOStream, arg1 *C.GCancellable, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFn(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.CloseFn(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_IOStreamClass_get_input_stream
func _gotk4_gio2_IOStreamClass_get_input_stream(arg0 *C.GIOStream) (cret *C.GInputStream) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ InputStream() InputStreamer })

	inputStream := iface.InputStream()

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(inputStream).Native()))

	return cret
}

//export _gotk4_gio2_IOStreamClass_get_output_stream
func _gotk4_gio2_IOStreamClass_get_output_stream(arg0 *C.GIOStream) (cret *C.GOutputStream) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ OutputStream() OutputStreamer })

	outputStream := iface.OutputStream()

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(outputStream).Native()))

	return cret
}

func wrapIOStream(obj *coreglib.Object) *IOStream {
	return &IOStream{
		Object: obj,
	}
}

func marshalIOStream(p uintptr) (interface{}, error) {
	return wrapIOStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (stream *IOStream) baseIOStream() *IOStream {
	return stream
}

// BaseIOStream returns the underlying base object.
func BaseIOStream(obj IOStreamer) *IOStream {
	return obj.baseIOStream()
}

// ClearPending clears the pending flag on stream.
func (stream *IOStream) ClearPending() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**IOStream)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gio", "IOStream").InvokeMethod("clear_pending", args[:], nil)

	runtime.KeepAlive(stream)
}

// Close closes the stream, releasing resources related to it. This will also
// close the individual input and output streams, if they are not already
// closed.
//
// Once the stream is closed, all other operations will return
// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an error.
//
// Closing a stream will automatically flush any outstanding buffers in the
// stream.
//
// Streams will be automatically closed when the last reference is dropped, but
// you might want to call this function to make sure resources are released as
// early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file
// descriptor) open after the stream is closed. See the documentation for the
// individual stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to close will
// still return G_IO_ERROR_CLOSED for all operations. Still, it is important to
// check and report the error to the user, otherwise there might be a loss of
// data as all data might not be written.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. Cancelling a close will
// still leave the stream closed, but some streams can use a faster close that
// doesn't block to e.g. check errors.
//
// The default implementation of this method just calls close on the individual
// input/output streams.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
func (stream *IOStream) Close(ctx context.Context) error {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**IOStream)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "IOStream").InvokeMethod("close", args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CloseFinish closes a stream.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (stream *IOStream) CloseFinish(result AsyncResulter) error {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))
	*(**IOStream)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gio", "IOStream").InvokeMethod("close_finish", args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// InputStream gets the input stream for this object. This is used for reading.
//
// The function returns the following values:
//
//    - inputStream owned by the OStream. Do not free.
//
func (stream *IOStream) InputStream() InputStreamer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**IOStream)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "IOStream").InvokeMethod("get_input_stream", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _inputStream InputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.InputStreamer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(InputStreamer)
			return ok
		})
		rv, ok := casted.(InputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
		}
		_inputStream = rv
	}

	return _inputStream
}

// OutputStream gets the output stream for this object. This is used for
// writing.
//
// The function returns the following values:
//
//    - outputStream owned by the OStream. Do not free.
//
func (stream *IOStream) OutputStream() OutputStreamer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**IOStream)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "IOStream").InvokeMethod("get_output_stream", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _outputStream OutputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.OutputStreamer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(OutputStreamer)
			return ok
		})
		rv, ok := casted.(OutputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.OutputStreamer")
		}
		_outputStream = rv
	}

	return _outputStream
}

// HasPending checks if a stream has pending actions.
//
// The function returns the following values:
//
//    - ok: TRUE if stream has pending actions.
//
func (stream *IOStream) HasPending() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**IOStream)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "IOStream").InvokeMethod("has_pending", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed checks if a stream is closed.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream is closed.
//
func (stream *IOStream) IsClosed() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**IOStream)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "IOStream").InvokeMethod("is_closed", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPending sets stream to have actions pending. If the pending flag is
// already set or stream is closed, it will return FALSE and set error.
func (stream *IOStream) SetPending() error {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**IOStream)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gio", "IOStream").InvokeMethod("set_pending", args[:], nil)

	runtime.KeepAlive(stream)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// IOStreamSpliceFinish finishes an asynchronous io stream splice operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func IOStreamSpliceFinish(result AsyncResulter) error {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))
	*(*AsyncResulter)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gio", "splice_finish").Invoke(args[:], nil)

	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
