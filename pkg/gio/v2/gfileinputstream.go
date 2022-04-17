// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GFileInfo* _gotk4_gio2_FileInputStreamClass_query_info(GFileInputStream*, char*, GCancellable*, GError**);
// extern GFileInfo* _gotk4_gio2_FileInputStreamClass_query_info_finish(GFileInputStream*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_FileInputStreamClass_can_seek(GFileInputStream*);
// extern gboolean _gotk4_gio2_FileInputStreamClass_seek(GFileInputStream*, goffset, GSeekType, GCancellable*, GError**);
// extern goffset _gotk4_gio2_FileInputStreamClass_tell(GFileInputStream*);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// glib.Type values for gfileinputstream.go.
var GTypeFileInputStream = externglib.Type(C.g_file_input_stream_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFileInputStream, F: marshalFileInputStream},
	})
}

// FileInputStreamOverrider contains methods that are overridable.
type FileInputStreamOverrider interface {
	externglib.Objector
	// The function returns the following values:
	//
	CanSeek() bool
	// QueryInfo queries a file input stream the given attributes. This function
	// blocks while querying the stream. For the asynchronous (non-blocking)
	// version of this function, see g_file_input_stream_query_info_async().
	// While the stream is blocked, the stream will set the pending flag
	// internally, and any other operations on the stream will fail with
	// G_IO_ERROR_PENDING.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - attributes: file attribute query string.
	//
	// The function returns the following values:
	//
	//    - fileInfo or NULL on error.
	//
	QueryInfo(ctx context.Context, attributes string) (*FileInfo, error)
	// QueryInfoFinish finishes an asynchronous info query operation.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - fileInfo: Info.
	//
	QueryInfoFinish(result AsyncResultOverrider) (*FileInfo, error)
	// The function takes the following parameters:
	//
	//    - ctx (optional)
	//    - offset
	//    - typ
	//
	Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	// The function returns the following values:
	//
	Tell() int64
}

// WrapFileInputStreamOverrider wraps the FileInputStreamOverrider
// interface implementation to access the instance methods.
func WrapFileInputStreamOverrider(obj FileInputStreamOverrider) *FileInputStream {
	return wrapFileInputStream(externglib.BaseObject(obj))
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
	_ [0]func() // equal guard
	InputStream

	*externglib.Object
	Seekable
}

var (
	_ InputStreamer       = (*FileInputStream)(nil)
	_ externglib.Objector = (*FileInputStream)(nil)
)

func classInitFileInputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GFileInputStreamClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GFileInputStreamClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ CanSeek() bool }); ok {
		pclass.can_seek = (*[0]byte)(C._gotk4_gio2_FileInputStreamClass_can_seek)
	}

	if _, ok := goval.(interface {
		QueryInfo(ctx context.Context, attributes string) (*FileInfo, error)
	}); ok {
		pclass.query_info = (*[0]byte)(C._gotk4_gio2_FileInputStreamClass_query_info)
	}

	if _, ok := goval.(interface {
		QueryInfoFinish(result AsyncResultOverrider) (*FileInfo, error)
	}); ok {
		pclass.query_info_finish = (*[0]byte)(C._gotk4_gio2_FileInputStreamClass_query_info_finish)
	}

	if _, ok := goval.(interface {
		Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	}); ok {
		pclass.seek = (*[0]byte)(C._gotk4_gio2_FileInputStreamClass_seek)
	}

	if _, ok := goval.(interface{ Tell() int64 }); ok {
		pclass.tell = (*[0]byte)(C._gotk4_gio2_FileInputStreamClass_tell)
	}
}

//export _gotk4_gio2_FileInputStreamClass_can_seek
func _gotk4_gio2_FileInputStreamClass_can_seek(arg0 *C.GFileInputStream) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CanSeek() bool })

	ok := iface.CanSeek()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_FileInputStreamClass_query_info
func _gotk4_gio2_FileInputStreamClass_query_info(arg0 *C.GFileInputStream, arg1 *C.char, arg2 *C.GCancellable, _cerr **C.GError) (cret *C.GFileInfo) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		QueryInfo(ctx context.Context, attributes string) (*FileInfo, error)
	})

	var _cancellable context.Context // out
	var _attributes string           // out

	if arg2 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg2))
	}
	_attributes = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	fileInfo, _goerr := iface.QueryInfo(_cancellable, _attributes)

	cret = (*C.GFileInfo)(unsafe.Pointer(externglib.InternObject(fileInfo).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(fileInfo).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileInputStreamClass_query_info_finish
func _gotk4_gio2_FileInputStreamClass_query_info_finish(arg0 *C.GFileInputStream, arg1 *C.GAsyncResult, _cerr **C.GError) (cret *C.GFileInfo) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		QueryInfoFinish(result AsyncResultOverrider) (*FileInfo, error)
	})

	var _result AsyncResultOverrider // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(AsyncResultOverrider)
			return ok
		})
		rv, ok := casted.(AsyncResultOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	fileInfo, _goerr := iface.QueryInfoFinish(_result)

	cret = (*C.GFileInfo)(unsafe.Pointer(externglib.InternObject(fileInfo).Native()))
	C.g_object_ref(C.gpointer(externglib.InternObject(fileInfo).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileInputStreamClass_seek
func _gotk4_gio2_FileInputStreamClass_seek(arg0 *C.GFileInputStream, arg1 C.goffset, arg2 C.GSeekType, arg3 *C.GCancellable, _cerr **C.GError) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	})

	var _cancellable context.Context // out
	var _offset int64                // out
	var _typ glib.SeekType           // out

	if arg3 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg3))
	}
	_offset = int64(arg1)
	_typ = glib.SeekType(arg2)

	_goerr := iface.Seek(_cancellable, _offset, _typ)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileInputStreamClass_tell
func _gotk4_gio2_FileInputStreamClass_tell(arg0 *C.GFileInputStream) (cret C.goffset) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Tell() int64 })

	gint64 := iface.Tell()

	cret = C.goffset(gint64)

	return cret
}

func wrapFileInputStream(obj *externglib.Object) *FileInputStream {
	return &FileInputStream{
		InputStream: InputStream{
			Object: obj,
		},
		Object: obj,
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalFileInputStream(p uintptr) (interface{}, error) {
	return wrapFileInputStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// QueryInfo queries a file input stream the given attributes. This function
// blocks while querying the stream. For the asynchronous (non-blocking) version
// of this function, see g_file_input_stream_query_info_async(). While the
// stream is blocked, the stream will set the pending flag internally, and any
// other operations on the stream will fail with G_IO_ERROR_PENDING.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - attributes: file attribute query string.
//
// The function returns the following values:
//
//    - fileInfo or NULL on error.
//
func (stream *FileInputStream) QueryInfo(ctx context.Context, attributes string) (*FileInfo, error) {
	var _arg0 *C.GFileInputStream // out
	var _arg2 *C.GCancellable     // out
	var _arg1 *C.char             // out
	var _cret *C.GFileInfo        // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_input_stream_query_info(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(attributes)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _fileInfo, _goerr
}

// QueryInfoAsync queries the stream information asynchronously. When the
// operation is finished callback will be called. You can then call
// g_file_input_stream_query_info_finish() to get the result of the operation.
//
// For the synchronous version of this function, see
// g_file_input_stream_query_info().
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be set.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - attributes: file attribute query string.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *FileInputStream) QueryInfoAsync(ctx context.Context, attributes string, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GFileInputStream   // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_file_input_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// QueryInfoFinish finishes an asynchronous info query operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - fileInfo: Info.
//
func (stream *FileInputStream) QueryInfoFinish(result AsyncResultOverrider) (*FileInfo, error) {
	var _arg0 *C.GFileInputStream // out
	var _arg1 *C.GAsyncResult     // out
	var _cret *C.GFileInfo        // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(externglib.InternObject(stream).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.g_file_input_stream_query_info_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _fileInfo, _goerr
}
