// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// DataSize returns the number of bytes from the start up to including the last
// byte written in the stream that has not been truncated away.
//
// The function returns the following values:
//
//    - gsize: number of bytes written to the stream.
//
func (ostream *MemoryOutputStream) DataSize() uint {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gsize                // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(coreglib.InternObject(ostream).Native()))

	_cret = C.g_memory_output_stream_get_data_size(_arg0)
	runtime.KeepAlive(ostream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}
