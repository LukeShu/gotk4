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

// Container: get the #GFile container which is being enumerated.
//
// The function returns the following values:
//
//    - file which is being enumerated.
//
func (enumerator *FileEnumerator) Container() *File {
	var _arg0 *C.GFileEnumerator // out
	var _cret *C.GFile           // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))

	_cret = C.g_file_enumerator_get_container(_arg0)
	runtime.KeepAlive(enumerator)

	var _file *File // out

	_file = wrapFile(coreglib.Take(unsafe.Pointer(_cret)))

	return _file
}
