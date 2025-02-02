// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// UTF8CollateKeyForFilename converts a string into a collation key that can
// be compared with other collation keys produced by the same function using
// strcmp().
//
// In order to sort filenames correctly, this function treats the dot '.' as a
// special case. Most dictionary orderings seem to consider it insignificant,
// thus producing the ordering "event.c" "eventgenerator.c" "event.h" instead of
// "event.c" "event.h" "eventgenerator.c". Also, we would like to treat numbers
// intelligently so that "file1" "file10" "file5" is sorted as "file1" "file5"
// "file10".
//
// Note that this function depends on the [current locale][setlocale].
//
// The function takes the following parameters:
//
//   - str: UTF-8 encoded string.
//   - len: length of str, in bytes, or -1 if str is nul-terminated.
//
// The function returns the following values:
//
//   - utf8: newly allocated string. This string should be freed with g_free()
//     when you are done with it.
//
func UTF8CollateKeyForFilename(str string, len int) string {
	var _arg1 *C.gchar // out
	var _arg2 C.gssize // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(len)

	_cret = C.g_utf8_collate_key_for_filename(_arg1, _arg2)
	runtime.KeepAlive(str)
	runtime.KeepAlive(len)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
