// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
		{T: externglib.Type(C.g_unix_input_stream_get_type()), F: marshalUnixInputStream},
	})
}

// UnixInputStream implements Stream for reading from a UNIX file descriptor,
// including asynchronous operations. (If the file descriptor refers to a socket
// or pipe, this will use poll() to do asynchronous I/O. If it refers to a
// regular file, it will fall back to doing asynchronous I/O in another thread.)
//
// Note that `<gio/gunixinputstream.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixInputStream interface {
	gextras.Objector

	// CloseFd returns whether the file descriptor of @stream will be closed
	// when the stream is closed.
	CloseFd() bool
	// Fd: return the UNIX file descriptor that the stream reads from.
	Fd() int
	// SetCloseFd sets whether the file descriptor of @stream shall be closed
	// when the stream is closed.
	SetCloseFd(closeFd bool)
}

// UnixInputStreamClass implements the UnixInputStream interface.
type UnixInputStreamClass struct {
	*externglib.Object
	InputStreamClass
	FileDescriptorBasedInterface
	PollableInputStreamInterface
}

var _ UnixInputStream = (*UnixInputStreamClass)(nil)

func wrapUnixInputStream(obj *externglib.Object) UnixInputStream {
	return &UnixInputStreamClass{
		Object: obj,
		InputStreamClass: InputStreamClass{
			Object: obj,
		},
		FileDescriptorBasedInterface: FileDescriptorBasedInterface{
			Object: obj,
		},
		PollableInputStreamInterface: PollableInputStreamInterface{
			InputStreamClass: InputStreamClass{
				Object: obj,
			},
		},
	}
}

func marshalUnixInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapUnixInputStream(obj), nil
}

// NewUnixInputStream creates a new InputStream for the given @fd.
//
// If @close_fd is true, the file descriptor will be closed when the stream is
// closed.
func NewUnixInputStream(fd int, closeFd bool) *UnixInputStreamClass {
	var _arg1 C.gint          // out
	var _arg2 C.gboolean      // out
	var _cret *C.GInputStream // in

	_arg1 = C.gint(fd)
	if closeFd {
		_arg2 = C.TRUE
	}

	_cret = C.g_unix_input_stream_new(_arg1, _arg2)

	var _unixInputStream *UnixInputStreamClass // out

	_unixInputStream = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*UnixInputStreamClass)

	return _unixInputStream
}

// CloseFd returns whether the file descriptor of @stream will be closed when
// the stream is closed.
func (s *UnixInputStreamClass) CloseFd() bool {
	var _arg0 *C.GUnixInputStream // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GUnixInputStream)(unsafe.Pointer((&s).Native()))

	_cret = C.g_unix_input_stream_get_close_fd(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Fd: return the UNIX file descriptor that the stream reads from.
func (s *UnixInputStreamClass) Fd() int {
	var _arg0 *C.GUnixInputStream // out
	var _cret C.gint              // in

	_arg0 = (*C.GUnixInputStream)(unsafe.Pointer((&s).Native()))

	_cret = C.g_unix_input_stream_get_fd(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetCloseFd sets whether the file descriptor of @stream shall be closed when
// the stream is closed.
func (s *UnixInputStreamClass) SetCloseFd(closeFd bool) {
	var _arg0 *C.GUnixInputStream // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GUnixInputStream)(unsafe.Pointer((&s).Native()))
	if closeFd {
		_arg1 = C.TRUE
	}

	C.g_unix_input_stream_set_close_fd(_arg0, _arg1)
}
