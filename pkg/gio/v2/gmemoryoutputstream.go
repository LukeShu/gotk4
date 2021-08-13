// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_memory_output_stream_get_type()), F: marshalMemoryOutputStreamer},
	})
}

// MemoryOutputStream is a class for using arbitrary memory chunks as output for
// GIO streaming output operations.
//
// As of GLib 2.34, OutputStream trivially implements OutputStream: it always
// polls as ready.
type MemoryOutputStream struct {
	OutputStream

	PollableOutputStream
	Seekable
	*externglib.Object
}

func wrapMemoryOutputStream(obj *externglib.Object) *MemoryOutputStream {
	return &MemoryOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
		PollableOutputStream: PollableOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalMemoryOutputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMemoryOutputStream(obj), nil
}

// NewMemoryOutputStreamResizable creates a new OutputStream, using g_realloc()
// and g_free() for memory allocation.
func NewMemoryOutputStreamResizable() *MemoryOutputStream {
	var _cret *C.GOutputStream // in

	_cret = C.g_memory_output_stream_new_resizable()

	var _memoryOutputStream *MemoryOutputStream // out

	_memoryOutputStream = wrapMemoryOutputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryOutputStream
}

// Data gets any loaded data from the ostream.
//
// Note that the returned pointer may become invalid on the next write or
// truncate operation on the stream.
func (ostream *MemoryOutputStream) Data() cgo.Handle {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gpointer             // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(ostream.Native()))

	_cret = C.g_memory_output_stream_get_data(_arg0)

	runtime.KeepAlive(ostream)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// DataSize returns the number of bytes from the start up to including the last
// byte written in the stream that has not been truncated away.
func (ostream *MemoryOutputStream) DataSize() uint {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gsize                // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(ostream.Native()))

	_cret = C.g_memory_output_stream_get_data_size(_arg0)

	runtime.KeepAlive(ostream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Size gets the size of the currently allocated data area (available from
// g_memory_output_stream_get_data()).
//
// You probably don't want to use this function on resizable streams. See
// g_memory_output_stream_get_data_size() instead. For resizable streams the
// size returned by this function is an implementation detail and may be change
// at any time in response to operations on the stream.
//
// If the stream is fixed-sized (ie: no realloc was passed to
// g_memory_output_stream_new()) then this is the maximum size of the stream and
// further writes will return G_IO_ERROR_NO_SPACE.
//
// In any case, if you want the number of bytes currently written to the stream,
// use g_memory_output_stream_get_data_size().
func (ostream *MemoryOutputStream) Size() uint {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gsize                // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(ostream.Native()))

	_cret = C.g_memory_output_stream_get_size(_arg0)

	runtime.KeepAlive(ostream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// StealData gets any loaded data from the ostream. Ownership of the data is
// transferred to the caller; when no longer needed it must be freed using the
// free function set in ostream's OutputStream:destroy-function property.
//
// ostream must be closed before calling this function.
func (ostream *MemoryOutputStream) StealData() cgo.Handle {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gpointer             // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(ostream.Native()))

	_cret = C.g_memory_output_stream_steal_data(_arg0)

	runtime.KeepAlive(ostream)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}
