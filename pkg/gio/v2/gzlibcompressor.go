// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_zlib_compressor_get_type()), F: marshalZlibCompressorrer},
	})
}

// ZlibCompressor: zlib decompression
type ZlibCompressor struct {
	*externglib.Object

	Converter
}

func wrapZlibCompressor(obj *externglib.Object) *ZlibCompressor {
	return &ZlibCompressor{
		Object: obj,
		Converter: Converter{
			Object: obj,
		},
	}
}

func marshalZlibCompressorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapZlibCompressor(obj), nil
}

// NewZlibCompressor creates a new Compressor.
func NewZlibCompressor(format ZlibCompressorFormat, level int32) *ZlibCompressor {
	var _arg1 C.GZlibCompressorFormat // out
	var _arg2 C.int                   // out
	var _cret *C.GZlibCompressor      // in

	_arg1 = C.GZlibCompressorFormat(format)
	_arg2 = C.int(level)

	_cret = C.g_zlib_compressor_new(_arg1, _arg2)
	runtime.KeepAlive(format)
	runtime.KeepAlive(level)

	var _zlibCompressor *ZlibCompressor // out

	_zlibCompressor = wrapZlibCompressor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _zlibCompressor
}

// FileInfo returns the Compressor:file-info property.
func (compressor *ZlibCompressor) FileInfo() *FileInfo {
	var _arg0 *C.GZlibCompressor // out
	var _cret *C.GFileInfo       // in

	_arg0 = (*C.GZlibCompressor)(unsafe.Pointer(compressor.Native()))

	_cret = C.g_zlib_compressor_get_file_info(_arg0)
	runtime.KeepAlive(compressor)

	var _fileInfo *FileInfo // out

	if _cret != nil {
		_fileInfo = wrapFileInfo(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _fileInfo
}

// SetFileInfo sets file_info in compressor. If non-NULL, and compressor's
// Compressor:format property is G_ZLIB_COMPRESSOR_FORMAT_GZIP, it will be used
// to set the file name and modification time in the GZIP header of the
// compressed data.
//
// Note: it is an error to call this function while a compression is in
// progress; it may only be called immediately after creation of compressor, or
// after resetting it with g_converter_reset().
func (compressor *ZlibCompressor) SetFileInfo(fileInfo *FileInfo) {
	var _arg0 *C.GZlibCompressor // out
	var _arg1 *C.GFileInfo       // out

	_arg0 = (*C.GZlibCompressor)(unsafe.Pointer(compressor.Native()))
	if fileInfo != nil {
		_arg1 = (*C.GFileInfo)(unsafe.Pointer(fileInfo.Native()))
	}

	C.g_zlib_compressor_set_file_info(_arg0, _arg1)
	runtime.KeepAlive(compressor)
	runtime.KeepAlive(fileInfo)
}
