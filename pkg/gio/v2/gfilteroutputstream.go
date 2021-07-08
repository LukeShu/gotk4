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
		{T: externglib.Type(C.g_filter_output_stream_get_type()), F: marshalFilterOutputStream},
	})
}

// FilterOutputStream: base class for output stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterOutputStream interface {
	gextras.Objector

	// BaseStream gets the base stream for the filter stream.
	BaseStream() OutputStream
	// CloseBaseStream returns whether the base stream will be closed when
	// @stream is closed.
	CloseBaseStream() bool
	// SetCloseBaseStream sets whether the base stream will be closed when
	// @stream is closed.
	SetCloseBaseStream(closeBase bool)
}

// FilterOutputStreamClass implements the FilterOutputStream interface.
type FilterOutputStreamClass struct {
	OutputStreamClass
}

var _ FilterOutputStream = (*FilterOutputStreamClass)(nil)

func wrapFilterOutputStream(obj *externglib.Object) FilterOutputStream {
	return &FilterOutputStreamClass{
		OutputStreamClass: OutputStreamClass{
			Object: obj,
		},
	}
}

func marshalFilterOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFilterOutputStream(obj), nil
}

// BaseStream gets the base stream for the filter stream.
func (s *FilterOutputStreamClass) BaseStream() OutputStream {
	var _arg0 *C.GFilterOutputStream // out
	var _cret *C.GOutputStream       // in

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_filter_output_stream_get_base_stream(_arg0)

	var _outputStream OutputStream // out

	_outputStream = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(OutputStream)

	return _outputStream
}

// CloseBaseStream returns whether the base stream will be closed when @stream
// is closed.
func (s *FilterOutputStreamClass) CloseBaseStream() bool {
	var _arg0 *C.GFilterOutputStream // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_filter_output_stream_get_close_base_stream(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when @stream
// is closed.
func (s *FilterOutputStreamClass) SetCloseBaseStream(closeBase bool) {
	var _arg0 *C.GFilterOutputStream // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(s.Native()))
	if closeBase {
		_arg1 = C.TRUE
	}

	C.g_filter_output_stream_set_close_base_stream(_arg0, _arg1)
}
