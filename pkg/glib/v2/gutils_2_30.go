// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// FormatSize formats a size (for example the size of a file) into a human
// readable string. Sizes are rounded to the nearest size prefix (kB, MB,
// GB) and are displayed rounded to the nearest tenth. E.g. the file size
// 3292528 bytes will be converted into the string "3.2 MB". The returned string
// is UTF-8, and may use a non-breaking space to separate the number and units,
// to ensure they aren’t separated when line wrapped.
//
// The prefix units base is 1000 (i.e. 1 kB is 1000 bytes).
//
// This string should be freed with g_free() when not needed any longer.
//
// See g_format_size_full() for more options about how the size might be
// formatted.
//
// The function takes the following parameters:
//
//   - size in bytes.
//
// The function returns the following values:
//
//   - utf8: newly-allocated formatted string containing a human readable file
//     size.
//
func FormatSize(size uint64) string {
	var _arg1 C.guint64 // out
	var _cret *C.gchar  // in

	_arg1 = C.guint64(size)

	_cret = C.g_format_size(_arg1)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FormatSizeFull formats a size.
//
// This function is similar to g_format_size() but allows for flags that modify
// the output. See SizeFlags.
//
// The function takes the following parameters:
//
//   - size in bytes.
//   - flags to modify the output.
//
// The function returns the following values:
//
//   - utf8: newly-allocated formatted string containing a human readable file
//     size.
//
func FormatSizeFull(size uint64, flags FormatSizeFlags) string {
	var _arg1 C.guint64          // out
	var _arg2 C.GFormatSizeFlags // out
	var _cret *C.gchar           // in

	_arg1 = C.guint64(size)
	_arg2 = C.GFormatSizeFlags(flags)

	_cret = C.g_format_size_full(_arg1, _arg2)
	runtime.KeepAlive(size)
	runtime.KeepAlive(flags)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
