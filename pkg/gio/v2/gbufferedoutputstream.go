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
		{T: externglib.Type(C.g_buffered_output_stream_get_type()), F: marshalBufferedOutputStream},
	})
}

// BufferedOutputStream: buffered output stream implements OutputStream and
// provides for buffered writes.
//
// By default, OutputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered output stream, use g_buffered_output_stream_new(), or
// g_buffered_output_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_output_stream_get_buffer_size(). To change the size of a buffered
// output stream's buffer, use g_buffered_output_stream_set_buffer_size(). Note
// that the buffer's size cannot be reduced below the size of the data within
// the buffer.
type BufferedOutputStream interface {
	gextras.Objector

	// AutoGrow checks if the buffer automatically grows as data is added.
	AutoGrow() bool
	// BufferSize gets the size of the buffer in the @stream.
	BufferSize() uint
	// SetAutoGrow sets whether or not the @stream's buffer should automatically
	// grow. If @auto_grow is true, then each write will just make the buffer
	// larger, and you must manually flush the buffer to actually write out the
	// data to the underlying stream.
	SetAutoGrow(autoGrow bool)
	// SetBufferSize sets the size of the internal buffer to @size.
	SetBufferSize(size uint)
}

// BufferedOutputStreamClass implements the BufferedOutputStream interface.
type BufferedOutputStreamClass struct {
	*externglib.Object
	FilterOutputStreamClass
	SeekableInterface
}

var _ BufferedOutputStream = (*BufferedOutputStreamClass)(nil)

func wrapBufferedOutputStream(obj *externglib.Object) BufferedOutputStream {
	return &BufferedOutputStreamClass{
		Object: obj,
		FilterOutputStreamClass: FilterOutputStreamClass{
			OutputStreamClass: OutputStreamClass{
				Object: obj,
			},
		},
		SeekableInterface: SeekableInterface{
			Object: obj,
		},
	}
}

func marshalBufferedOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBufferedOutputStream(obj), nil
}

// NewBufferedOutputStream creates a new buffered output stream for a base
// stream.
func NewBufferedOutputStream(baseStream OutputStream) *BufferedOutputStreamClass {
	var _arg1 *C.GOutputStream // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer((&baseStream).Native()))

	_cret = C.g_buffered_output_stream_new(_arg1)

	var _bufferedOutputStream *BufferedOutputStreamClass // out

	_bufferedOutputStream = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*BufferedOutputStreamClass)

	return _bufferedOutputStream
}

// NewBufferedOutputStreamSized creates a new buffered output stream with a
// given buffer size.
func NewBufferedOutputStreamSized(baseStream OutputStream, size uint) *BufferedOutputStreamClass {
	var _arg1 *C.GOutputStream // out
	var _arg2 C.gsize          // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer((&baseStream).Native()))
	_arg2 = C.gsize(size)

	_cret = C.g_buffered_output_stream_new_sized(_arg1, _arg2)

	var _bufferedOutputStream *BufferedOutputStreamClass // out

	_bufferedOutputStream = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*BufferedOutputStreamClass)

	return _bufferedOutputStream
}

// AutoGrow checks if the buffer automatically grows as data is added.
func (s *BufferedOutputStreamClass) AutoGrow() bool {
	var _arg0 *C.GBufferedOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer((&s).Native()))

	_cret = C.g_buffered_output_stream_get_auto_grow(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BufferSize gets the size of the buffer in the @stream.
func (s *BufferedOutputStreamClass) BufferSize() uint {
	var _arg0 *C.GBufferedOutputStream // out
	var _cret C.gsize                  // in

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer((&s).Native()))

	_cret = C.g_buffered_output_stream_get_buffer_size(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// SetAutoGrow sets whether or not the @stream's buffer should automatically
// grow. If @auto_grow is true, then each write will just make the buffer
// larger, and you must manually flush the buffer to actually write out the data
// to the underlying stream.
func (s *BufferedOutputStreamClass) SetAutoGrow(autoGrow bool) {
	var _arg0 *C.GBufferedOutputStream // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer((&s).Native()))
	if autoGrow {
		_arg1 = C.TRUE
	}

	C.g_buffered_output_stream_set_auto_grow(_arg0, _arg1)
}

// SetBufferSize sets the size of the internal buffer to @size.
func (s *BufferedOutputStreamClass) SetBufferSize(size uint) {
	var _arg0 *C.GBufferedOutputStream // out
	var _arg1 C.gsize                  // out

	_arg0 = (*C.GBufferedOutputStream)(unsafe.Pointer((&s).Native()))
	_arg1 = C.gsize(size)

	C.g_buffered_output_stream_set_buffer_size(_arg0, _arg1)
}
