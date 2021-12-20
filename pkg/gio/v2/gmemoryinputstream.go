// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_memory_input_stream_get_type()), F: marshalMemoryInputStreamer},
	})
}

// MemoryInputStream is a class for using arbitrary memory chunks as input for
// GIO streaming input operations.
//
// As of GLib 2.34, InputStream implements InputStream.
type MemoryInputStream struct {
	_ [0]func() // equal guard
	InputStream

	*externglib.Object
	PollableInputStream
	Seekable
}

var (
	_ InputStreamer       = (*MemoryInputStream)(nil)
	_ externglib.Objector = (*MemoryInputStream)(nil)
)

func wrapMemoryInputStream(obj *externglib.Object) *MemoryInputStream {
	return &MemoryInputStream{
		InputStream: InputStream{
			Object: obj,
		},
		Object: obj,
		PollableInputStream: PollableInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalMemoryInputStreamer(p uintptr) (interface{}, error) {
	return wrapMemoryInputStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMemoryInputStream creates a new empty InputStream.
//
// The function returns the following values:
//
//    - memoryInputStream: new Stream.
//
func NewMemoryInputStream() *MemoryInputStream {
	var _cret *C.GInputStream // in

	_cret = C.g_memory_input_stream_new()

	var _memoryInputStream *MemoryInputStream // out

	_memoryInputStream = wrapMemoryInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryInputStream
}

// NewMemoryInputStreamFromBytes creates a new InputStream with data from the
// given bytes.
//
// The function takes the following parameters:
//
//    - bytes: #GBytes.
//
// The function returns the following values:
//
//    - memoryInputStream: new Stream read from bytes.
//
func NewMemoryInputStreamFromBytes(bytes *glib.Bytes) *MemoryInputStream {
	var _arg1 *C.GBytes       // out
	var _cret *C.GInputStream // in

	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))

	_cret = C.g_memory_input_stream_new_from_bytes(_arg1)
	runtime.KeepAlive(bytes)

	var _memoryInputStream *MemoryInputStream // out

	_memoryInputStream = wrapMemoryInputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryInputStream
}

// AddBytes appends bytes to data that can be read from the input stream.
//
// The function takes the following parameters:
//
//    - bytes: input data.
//
func (stream *MemoryInputStream) AddBytes(bytes *glib.Bytes) {
	var _arg0 *C.GMemoryInputStream // out
	var _arg1 *C.GBytes             // out

	_arg0 = (*C.GMemoryInputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))

	C.g_memory_input_stream_add_bytes(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(bytes)
}
