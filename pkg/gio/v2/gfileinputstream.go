// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
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
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_input_stream_get_type()), F: marshalFileInputStreamer},
	})
}

// FileInputStreamOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileInputStreamOverrider interface {
	CanSeek() bool
	// QueryInfo queries a file input stream the given @attributes. This
	// function blocks while querying the stream. For the asynchronous
	// (non-blocking) version of this function, see
	// g_file_input_stream_query_info_async(). While the stream is blocked, the
	// stream will set the pending flag internally, and any other operations on
	// the stream will fail with G_IO_ERROR_PENDING.
	QueryInfo(attributes string, cancellable Cancellabler) (*FileInfo, error)
	// QueryInfoAsync queries the stream information asynchronously. When the
	// operation is finished @callback will be called. You can then call
	// g_file_input_stream_query_info_finish() to get the result of the
	// operation.
	//
	// For the synchronous version of this function, see
	// g_file_input_stream_query_info().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// QueryInfoFinish finishes an asynchronous info query operation.
	QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
	Seek(offset int64, typ glib.SeekType, cancellable Cancellabler) error
	Tell() int64
}

// FileInputStreamer describes FileInputStream's methods.
type FileInputStreamer interface {
	// QueryInfo queries a file input stream the given @attributes.
	QueryInfo(attributes string, cancellable Cancellabler) (*FileInfo, error)
	// QueryInfoAsync queries the stream information asynchronously.
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// QueryInfoFinish finishes an asynchronous info query operation.
	QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
}

// FileInputStream provides input streams that take their content from a file.
//
// GFileInputStream implements #GSeekable, which allows the input stream to jump
// to arbitrary positions in the file, provided the filesystem of the file
// allows it. To find the position of a file input stream, use
// g_seekable_tell(). To find out if a file input stream supports seeking, use
// g_seekable_can_seek(). To position a file input stream, use
// g_seekable_seek().
type FileInputStream struct {
	InputStream

	Seekable
}

var (
	_ FileInputStreamer = (*FileInputStream)(nil)
	_ gextras.Nativer   = (*FileInputStream)(nil)
)

func wrapFileInputStream(obj *externglib.Object) FileInputStreamer {
	return &FileInputStream{
		InputStream: InputStream{
			Object: obj,
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalFileInputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileInputStream(obj), nil
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *FileInputStream) Native() uintptr {
	return v.InputStream.Object.Native()
}

// QueryInfo queries a file input stream the given @attributes. This function
// blocks while querying the stream. For the asynchronous (non-blocking) version
// of this function, see g_file_input_stream_query_info_async(). While the
// stream is blocked, the stream will set the pending flag internally, and any
// other operations on the stream will fail with G_IO_ERROR_PENDING.
func (stream *FileInputStream) QueryInfo(attributes string, cancellable Cancellabler) (*FileInfo, error) {
	var _arg0 *C.GFileInputStream // out
	var _arg1 *C.char             // out
	var _arg2 *C.GCancellable     // out
	var _cret *C.GFileInfo        // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_file_input_stream_query_info(_arg0, _arg1, _arg2, &_cerr)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}

// QueryInfoAsync queries the stream information asynchronously. When the
// operation is finished @callback will be called. You can then call
// g_file_input_stream_query_info_finish() to get the result of the operation.
//
// For the synchronous version of this function, see
// g_file_input_stream_query_info().
//
// If @cancellable is not nil, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be set
func (stream *FileInputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GFileInputStream   // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.Assign(callback))

	C.g_file_input_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// QueryInfoFinish finishes an asynchronous info query operation.
func (stream *FileInputStream) QueryInfoFinish(result AsyncResulter) (*FileInfo, error) {
	var _arg0 *C.GFileInputStream // out
	var _arg1 *C.GAsyncResult     // out
	var _cret *C.GFileInfo        // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_file_input_stream_query_info_finish(_arg0, _arg1, &_cerr)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}
