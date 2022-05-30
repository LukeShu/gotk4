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
// extern GFileInfo* _gotk4_gio2_FileInputStreamClass_query_info(GFileInputStream*, char*, GCancellable*, GError**);
// extern GFileInfo* _gotk4_gio2_FileInputStreamClass_query_info_finish(GFileInputStream*, GAsyncResult*, GError**);
// extern gboolean _gotk4_gio2_FileInputStreamClass_can_seek(GFileInputStream*);
import "C"

// glib.Type values for gfileinputstream.go.
var GTypeFileInputStream = coreglib.Type(C.g_file_input_stream_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFileInputStream, F: marshalFileInputStream},
	})
}

// FileInputStreamOverrider contains methods that are overridable.
type FileInputStreamOverrider interface {
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
	_ [0]func() // equal guard
	InputStream

	*coreglib.Object
	Seekable
}

var (
	_ InputStreamer     = (*FileInputStream)(nil)
	_ coreglib.Objector = (*FileInputStream)(nil)
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
		QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
	}); ok {
		pclass.query_info_finish = (*[0]byte)(C._gotk4_gio2_FileInputStreamClass_query_info_finish)
	}
}

//export _gotk4_gio2_FileInputStreamClass_can_seek
func _gotk4_gio2_FileInputStreamClass_can_seek(arg0 *C.GFileInputStream) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CanSeek() bool })

	ok := iface.CanSeek()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_FileInputStreamClass_query_info
func _gotk4_gio2_FileInputStreamClass_query_info(arg0 *C.GFileInputStream, arg1 *C.char, arg2 *C.GCancellable, _cerr **C.GError) (cret *C.GFileInfo) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
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

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fileInfo).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(fileInfo).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileInputStreamClass_query_info_finish
func _gotk4_gio2_FileInputStreamClass_query_info_finish(arg0 *C.GFileInputStream, arg1 *C.GAsyncResult, _cerr **C.GError) (cret *C.GFileInfo) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
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

	fileInfo, _goerr := iface.QueryInfoFinish(_result)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fileInfo).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(fileInfo).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapFileInputStream(obj *coreglib.Object) *FileInputStream {
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
	return wrapFileInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg2 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**FileInputStream)(unsafe.Pointer(&args[1])) = _arg1
	*(*context.Context)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gio", "FileInputStream").InvokeMethod("query_info", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(attributes)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _fileInfo, _goerr
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
func (stream *FileInputStream) QueryInfoFinish(result AsyncResulter) (*FileInfo, error) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))
	*(**FileInputStream)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "FileInputStream").InvokeMethod("query_info_finish", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _fileInfo, _goerr
}
