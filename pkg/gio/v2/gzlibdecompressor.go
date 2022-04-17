// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gzlibdecompressor.go.
var GTypeZlibDecompressor = externglib.Type(C.g_zlib_decompressor_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeZlibDecompressor, F: marshalZlibDecompressor},
	})
}

// ZlibDecompressorOverrider contains methods that are overridable.
type ZlibDecompressorOverrider interface {
	externglib.Objector
}

// WrapZlibDecompressorOverrider wraps the ZlibDecompressorOverrider
// interface implementation to access the instance methods.
func WrapZlibDecompressorOverrider(obj ZlibDecompressorOverrider) *ZlibDecompressor {
	return wrapZlibDecompressor(externglib.BaseObject(obj))
}

// ZlibDecompressor: zlib decompression.
type ZlibDecompressor struct {
	_ [0]func() // equal guard
	*externglib.Object

	Converter
}

var (
	_ externglib.Objector = (*ZlibDecompressor)(nil)
)

func classInitZlibDecompressorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapZlibDecompressor(obj *externglib.Object) *ZlibDecompressor {
	return &ZlibDecompressor{
		Object: obj,
		Converter: Converter{
			Object: obj,
		},
	}
}

func marshalZlibDecompressor(p uintptr) (interface{}, error) {
	return wrapZlibDecompressor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewZlibDecompressor creates a new Decompressor.
//
// The function takes the following parameters:
//
//    - format to use for the compressed data.
//
// The function returns the following values:
//
//    - zlibDecompressor: new Decompressor.
//
func NewZlibDecompressor(format ZlibCompressorFormat) *ZlibDecompressor {
	var _arg1 C.GZlibCompressorFormat // out
	var _cret *C.GZlibDecompressor    // in

	_arg1 = C.GZlibCompressorFormat(format)

	_cret = C.g_zlib_decompressor_new(_arg1)
	runtime.KeepAlive(format)

	var _zlibDecompressor *ZlibDecompressor // out

	_zlibDecompressor = wrapZlibDecompressor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _zlibDecompressor
}

// FileInfo retrieves the Info constructed from the GZIP header data of
// compressed data processed by compressor, or NULL if decompressor's
// Decompressor:format property is not G_ZLIB_COMPRESSOR_FORMAT_GZIP, or the
// header data was not fully processed yet, or it not present in the data stream
// at all.
//
// The function returns the following values:
//
//    - fileInfo (optional) or NULL.
//
func (decompressor *ZlibDecompressor) FileInfo() *FileInfo {
	var _arg0 *C.GZlibDecompressor // out
	var _cret *C.GFileInfo         // in

	_arg0 = (*C.GZlibDecompressor)(unsafe.Pointer(externglib.InternObject(decompressor).Native()))

	_cret = C.g_zlib_decompressor_get_file_info(_arg0)
	runtime.KeepAlive(decompressor)

	var _fileInfo *FileInfo // out

	if _cret != nil {
		_fileInfo = wrapFileInfo(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _fileInfo
}
