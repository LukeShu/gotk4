// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// GetConsoleCharset obtains the character set used by the console attached to
// the process, which is suitable for printing output to the terminal.
//
// Usually this matches the result returned by g_get_charset(), but in
// environments where the locale's character set does not match the encoding of
// the console this function tries to guess a more suitable value instead.
//
// On Windows the character set returned by this function is the output code
// page used by the console associated with the calling process. If the codepage
// can't be determined (for example because there is no console attached) UTF-8
// is assumed.
//
// The return value is TRUE if the locale's encoding is UTF-8, in that case you
// can perhaps avoid calling g_convert().
//
// The string returned in charset is not allocated, and should not be freed.
//
// The function returns the following values:
//
//   - charset (optional): return location for character set name, or NULL.
//   - ok: TRUE if the returned charset is UTF-8.
//
func GetConsoleCharset() (string, bool) {
	var _arg1 *C.char    // in
	var _cret C.gboolean // in

	_cret = C.g_get_console_charset(&_arg1)

	var _charset string // out
	var _ok bool        // out

	if _arg1 != nil {
		_charset = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	}
	if _cret != 0 {
		_ok = true
	}

	return _charset, _ok
}
