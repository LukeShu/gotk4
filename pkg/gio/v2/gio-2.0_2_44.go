// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSimpleIOStream = coreglib.Type(C.g_simple_io_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSimpleIOStream, F: marshalSimpleIOStream},
	})
}

// SimpleIOStream creates a OStream from an arbitrary Stream and Stream. This
// allows any pair of input and output streams to be used with OStream methods.
//
// This is useful when you obtained a Stream and a Stream by other means, for
// instance creating them with platform specific methods as
// g_unix_input_stream_new() or g_win32_input_stream_new(), and you want to take
// advantage of the methods provided by OStream.
type SimpleIOStream struct {
	_ [0]func() // equal guard
	IOStream
}

var (
	_ IOStreamer = (*SimpleIOStream)(nil)
)

func wrapSimpleIOStream(obj *coreglib.Object) *SimpleIOStream {
	return &SimpleIOStream{
		IOStream: IOStream{
			Object: obj,
		},
	}
}

func marshalSimpleIOStream(p uintptr) (interface{}, error) {
	return wrapSimpleIOStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
