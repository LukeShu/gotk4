// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeZlibCompressor = coreglib.Type(C.g_zlib_compressor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeZlibCompressor, F: marshalZlibCompressor},
	})
}

// ZlibCompressorOverrides contains methods that are overridable.
type ZlibCompressorOverrides struct {
}

func defaultZlibCompressorOverrides(v *ZlibCompressor) ZlibCompressorOverrides {
	return ZlibCompressorOverrides{}
}

// ZlibCompressor: zlib decompression.
type ZlibCompressor struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Converter
}

var (
	_ coreglib.Objector = (*ZlibCompressor)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ZlibCompressor, *ZlibCompressorClass, ZlibCompressorOverrides](
		GTypeZlibCompressor,
		initZlibCompressorClass,
		wrapZlibCompressor,
		defaultZlibCompressorOverrides,
	)
}

func initZlibCompressorClass(gclass unsafe.Pointer, overrides ZlibCompressorOverrides, classInitFunc func(*ZlibCompressorClass)) {
	if classInitFunc != nil {
		class := (*ZlibCompressorClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapZlibCompressor(obj *coreglib.Object) *ZlibCompressor {
	return &ZlibCompressor{
		Object: obj,
		Converter: Converter{
			Object: obj,
		},
	}
}

func marshalZlibCompressor(p uintptr) (interface{}, error) {
	return wrapZlibCompressor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ZlibCompressorClass: instance of this type is always passed by reference.
type ZlibCompressorClass struct {
	*zlibCompressorClass
}

// zlibCompressorClass is the struct that's finalized.
type zlibCompressorClass struct {
	native *C.GZlibCompressorClass
}
