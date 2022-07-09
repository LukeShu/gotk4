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
// extern gssize _gotk4_gio2_BufferedInputStreamClass_fill(void*, gssize, void*, GError**);
// extern gssize _gotk4_gio2_BufferedInputStreamClass_fill_finish(void*, void*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// GTypeBufferedInputStream returns the GType for the type BufferedInputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBufferedInputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "BufferedInputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBufferedInputStream)
	return gtype
}

// BufferedInputStreamOverrider contains methods that are overridable.
type BufferedInputStreamOverrider interface {
	// Fill tries to read count bytes from the stream into the buffer. Will
	// block during this read.
	//
	// If count is zero, returns zero and does nothing. A value of count larger
	// than G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes read into the buffer is returned. It is
	// not an error if this is not the same as the requested size, as it can
	// happen e.g. near the end of a file. Zero is returned on end of file (or
	// if count is zero), but never otherwise.
	//
	// If count is -1 then the attempted read size is equal to the number of
	// bytes that are required to fill the buffer.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// On error -1 is returned and error is set accordingly.
	//
	// For the asynchronous, non-blocking, version of this function, see
	// g_buffered_input_stream_fill_async().
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - count: number of bytes that will be read from the stream.
	//
	// The function returns the following values:
	//
	//    - gssize: number of bytes read into stream's buffer, up to count, or -1
	//      on error.
	//
	Fill(ctx context.Context, count int) (int, error)
	// FillFinish finishes an asynchronous read.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - gssize of the read stream, or -1 on an error.
	//
	FillFinish(result AsyncResulter) (int, error)
}

// BufferedInputStream: buffered input stream implements InputStream and
// provides for buffered reads.
//
// By default, InputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered input stream, use g_buffered_input_stream_new(), or
// g_buffered_input_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_input_stream_get_buffer_size(). To change the size of a buffered
// input stream's buffer, use g_buffered_input_stream_set_buffer_size(). Note
// that the buffer's size cannot be reduced below the size of the data within
// the buffer.
type BufferedInputStream struct {
	_ [0]func() // equal guard
	FilterInputStream

	Seekable
}

var (
	_ FilterInputStreamer = (*BufferedInputStream)(nil)
)

func classInitBufferedInputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "BufferedInputStreamClass")

	if _, ok := goval.(interface {
		Fill(ctx context.Context, count int) (int, error)
	}); ok {
		o := pclass.StructFieldOffset("fill")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_BufferedInputStreamClass_fill)
	}

	if _, ok := goval.(interface {
		FillFinish(result AsyncResulter) (int, error)
	}); ok {
		o := pclass.StructFieldOffset("fill_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_BufferedInputStreamClass_fill_finish)
	}
}

//export _gotk4_gio2_BufferedInputStreamClass_fill
func _gotk4_gio2_BufferedInputStreamClass_fill(arg0 *C.void, arg1 C.gssize, arg2 *C.void, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Fill(ctx context.Context, count int) (int, error)
	})

	var _cancellable context.Context // out
	var _count int                   // out

	if arg2 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg2))
	}
	_count = int(arg1)

	gssize, _goerr := iface.Fill(_cancellable, _count)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_BufferedInputStreamClass_fill_finish
func _gotk4_gio2_BufferedInputStreamClass_fill_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		FillFinish(result AsyncResulter) (int, error)
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

	gssize, _goerr := iface.FillFinish(_result)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapBufferedInputStream(obj *coreglib.Object) *BufferedInputStream {
	return &BufferedInputStream{
		FilterInputStream: FilterInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalBufferedInputStream(p uintptr) (interface{}, error) {
	return wrapBufferedInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBufferedInputStream creates a new Stream from the given base_stream, with
// a buffer set to the default size (4 kilobytes).
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//
// The function returns the following values:
//
//    - bufferedInputStream for the given base_stream.
//
func NewBufferedInputStream(baseStream InputStreamer) *BufferedInputStream {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("new_BufferedInputStream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseStream)

	var _bufferedInputStream *BufferedInputStream // out

	_bufferedInputStream = wrapBufferedInputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedInputStream
}

// NewBufferedInputStreamSized creates a new InputStream from the given
// base_stream, with a buffer set to size.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//    - size: #gsize.
//
// The function returns the following values:
//
//    - bufferedInputStream: Stream.
//
func NewBufferedInputStreamSized(baseStream InputStreamer, size uint) *BufferedInputStream {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))
	*(*C.gsize)(unsafe.Pointer(&_args[1])) = C.gsize(size)

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("new_BufferedInputStream_sized", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(size)

	var _bufferedInputStream *BufferedInputStream // out

	_bufferedInputStream = wrapBufferedInputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bufferedInputStream
}

// Fill tries to read count bytes from the stream into the buffer. Will block
// during this read.
//
// If count is zero, returns zero and does nothing. A value of count larger than
// G_MAXSSIZE will cause a G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned. It is not
// an error if this is not the same as the requested size, as it can happen e.g.
// near the end of a file. Zero is returned on end of file (or if count is
// zero), but never otherwise.
//
// If count is -1 then the attempted read size is equal to the number of bytes
// that are required to fill the buffer.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// On error -1 is returned and error is set accordingly.
//
// For the asynchronous, non-blocking, version of this function, see
// g_buffered_input_stream_fill_async().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - count: number of bytes that will be read from the stream.
//
// The function returns the following values:
//
//    - gssize: number of bytes read into stream's buffer, up to count, or -1 on
//      error.
//
func (stream *BufferedInputStream) Fill(ctx context.Context, count int) (int, error) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.gssize)(unsafe.Pointer(&_args[1])) = C.gssize(count)

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("fill", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(count)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// FillAsync reads data into stream's buffer asynchronously, up to count size.
// io_priority can be used to prioritize reads. For the synchronous version of
// this function, see g_buffered_input_stream_fill().
//
// If count is -1 then the attempted read size is equal to the number of bytes
// that are required to fill the buffer.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object.
//    - count: number of bytes that will be read from the stream.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional): ReadyCallback.
//
func (stream *BufferedInputStream) FillAsync(ctx context.Context, count int, ioPriority int32, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.gssize)(unsafe.Pointer(&_args[1])) = C.gssize(count)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[4])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[5] = C.gpointer(gbox.AssignOnce(callback))
	}

	girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("fill_async", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(count)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// FillFinish finishes an asynchronous read.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - gssize of the read stream, or -1 on an error.
//
func (stream *BufferedInputStream) FillFinish(result AsyncResulter) (int, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("fill_finish", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// Available gets the size of the available data within the stream.
//
// The function returns the following values:
//
//    - gsize: size of the available stream.
//
func (stream *BufferedInputStream) Available() uint {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("get_available", _args[:], nil)
	_cret = *(*C.gsize)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _gsize uint // out

	_gsize = uint(*(*C.gsize)(unsafe.Pointer(&_cret)))

	return _gsize
}

// BufferSize gets the size of the input buffer.
//
// The function returns the following values:
//
//    - gsize: current buffer size.
//
func (stream *BufferedInputStream) BufferSize() uint {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("get_buffer_size", _args[:], nil)
	_cret = *(*C.gsize)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _gsize uint // out

	_gsize = uint(*(*C.gsize)(unsafe.Pointer(&_cret)))

	return _gsize
}

// PeekBuffer returns the buffer with the currently available bytes. The
// returned buffer must not be modified and will become invalid when reading
// from the stream or filling the buffer.
//
// The function returns the following values:
//
//    - guint8s: read-only buffer.
//
func (stream *BufferedInputStream) PeekBuffer() []byte {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("peek_buffer", _args[:], _outs[:])
	_cret = *(*unsafe.Pointer)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _guint8s []byte // out

	_guint8s = make([]byte, _outs[0])
	copy(_guint8s, unsafe.Slice((*byte)(unsafe.Pointer(_cret)), _outs[0]))

	return _guint8s
}

// ReadByte tries to read a single byte from the stream or the buffer. Will
// block during this read.
//
// On success, the byte read from the stream is returned. On end of stream -1 is
// returned but it's not an exceptional error and error is not set.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. If an operation was
// partially finished when the operation was cancelled the partial result will
// be returned, without an error.
//
// On error -1 is returned and error is set accordingly.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
// The function returns the following values:
//
//    - gint: byte read from the stream, or -1 on end of stream or error.
//
func (stream *BufferedInputStream) ReadByte(ctx context.Context) (int32, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_gret := girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("read_byte", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _gint int32  // out
	var _goerr error // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// SetBufferSize sets the size of the internal buffer of stream to size, or to
// the size of the contents of the buffer. The buffer can never be resized
// smaller than its current contents.
//
// The function takes the following parameters:
//
//    - size: #gsize.
//
func (stream *BufferedInputStream) SetBufferSize(size uint) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(*C.gsize)(unsafe.Pointer(&_args[1])) = C.gsize(size)

	girepository.MustFind("Gio", "BufferedInputStream").InvokeMethod("set_buffer_size", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(size)
}
