// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// FormatSizeForDisplay formats a size (for example the size of a file) into a
// human readable string. Sizes are rounded to the nearest size prefix (KB, MB,
// GB) and are displayed rounded to the nearest tenth. E.g. the file size
// 3292528 bytes will be converted into the string "3.1 MB".
//
// The prefix units base is 1024 (i.e. 1 KB is 1024 bytes).
//
// This string should be freed with g_free() when not needed any longer.
//
// Deprecated: This function is broken due to its use of SI suffixes to denote
// IEC units. Use g_format_size() instead.
//
// The function takes the following parameters:
//
//    - size in bytes.
//
// The function returns the following values:
//
//    - utf8: newly-allocated formatted string containing a human readable file
//      size.
//
func FormatSizeForDisplay(size int64) string {
	var _arg1 C.goffset // out
	var _cret *C.gchar  // in

	_arg1 = C.goffset(size)

	_cret = C.g_format_size_for_display(_arg1)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
