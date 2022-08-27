// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// GetEnviron gets the list of environment variables for the current process.
//
// The list is NULL terminated and each item in the list is of the form
// 'NAME=VALUE'.
//
// This is equivalent to direct access to the 'environ' global variable, except
// portable.
//
// The return value is freshly allocated and it should be freed with
// g_strfreev() when it is no longer needed.
//
// The function returns the following values:
//
//    - filenames: the list of environment variables.
//
func GetEnviron() []string {
	var _cret **C.gchar // in

	_cret = C.g_get_environ()

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
