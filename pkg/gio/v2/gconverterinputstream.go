// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeConverterInputStream returns the GType for the type ConverterInputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeConverterInputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "ConverterInputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalConverterInputStream)
	return gtype
}

// ConverterInputStreamOverrider contains methods that are overridable.
type ConverterInputStreamOverrider interface {
}

// ConverterInputStream: converter input stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, InputStream implements InputStream.
type ConverterInputStream struct {
	_ [0]func() // equal guard
	FilterInputStream

	*coreglib.Object
	InputStream
	PollableInputStream
}

var (
	_ FilterInputStreamer = (*ConverterInputStream)(nil)
	_ coreglib.Objector   = (*ConverterInputStream)(nil)
	_ InputStreamer       = (*ConverterInputStream)(nil)
)

func classInitConverterInputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapConverterInputStream(obj *coreglib.Object) *ConverterInputStream {
	return &ConverterInputStream{
		FilterInputStream: FilterInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Object: obj,
		InputStream: InputStream{
			Object: obj,
		},
		PollableInputStream: PollableInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
	}
}

func marshalConverterInputStream(p uintptr) (interface{}, error) {
	return wrapConverterInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConverterInputStream creates a new converter input stream for the
// base_stream.
//
// The function takes the following parameters:
//
//    - baseStream: Stream.
//    - converter: #GConverter.
//
// The function returns the following values:
//
//    - converterInputStream: new Stream.
//
func NewConverterInputStream(baseStream InputStreamer, converter Converterer) *ConverterInputStream {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(baseStream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(converter).Native()))

	_gret := girepository.MustFind("Gio", "ConverterInputStream").InvokeMethod("new_ConverterInputStream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(baseStream)
	runtime.KeepAlive(converter)

	var _converterInputStream *ConverterInputStream // out

	_converterInputStream = wrapConverterInputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _converterInputStream
}

// Converter gets the #GConverter that is used by converter_stream.
//
// The function returns the following values:
//
//    - converter of the converter input stream.
//
func (converterStream *ConverterInputStream) Converter() *Converter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(converterStream).Native()))

	_gret := girepository.MustFind("Gio", "ConverterInputStream").InvokeMethod("get_converter", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(converterStream)

	var _converter *Converter // out

	_converter = wrapConverter(coreglib.Take(unsafe.Pointer(_cret)))

	return _converter
}
