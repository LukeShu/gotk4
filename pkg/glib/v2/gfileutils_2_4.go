// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// FileReadLink reads the contents of the symbolic link filename like the POSIX
// readlink() function. The returned string is in the encoding used for
// filenames. Use g_filename_to_utf8() to convert it to UTF-8.
//
// The function takes the following parameters:
//
//    - filename: symbolic link.
//
// The function returns the following values:
//
//    - ret: newly-allocated string with the contents of the symbolic link, or
//      NULL if an error occurred.
//
func FileReadLink(filename string) (string, error) {
	var _arg1 *C.gchar  // out
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_read_link(_arg1, &_cerr)
	runtime.KeepAlive(filename)

	var _ret string  // out
	var _goerr error // out

	_ret = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ret, _goerr
}
