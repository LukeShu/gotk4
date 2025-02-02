// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// Listenv gets the names of all variables set in the environment.
//
// Programs that want to be portable to Windows should typically use this
// function and g_getenv() instead of using the environ array from the C library
// directly. On Windows, the strings in the environ array are in system codepage
// encoding, while in most of the typical use cases for environment variables
// in GLib-using programs you want the UTF-8 encoding that this function and
// g_getenv() provide.
//
// The function returns the following values:
//
//   - filenames: a NULL-terminated list of strings which must be freed with
//     g_strfreev().
//
func Listenv() []string {
	var _cret **C.gchar // in

	_cret = C.g_listenv()

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}
