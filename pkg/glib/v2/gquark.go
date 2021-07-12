// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// Quark is a non-zero integer which uniquely identifies a particular string. A
// GQuark value of zero is associated to nil.
type Quark = uint32

// InternStaticString returns a canonical representation for @string. Interned
// strings can be compared for equality by comparing the pointers, instead of
// using strcmp(). g_intern_static_string() does not copy the string, therefore
// @string must not be freed or modified.
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func InternStaticString(_string string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_intern_static_string(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// InternString returns a canonical representation for @string. Interned strings
// can be compared for equality by comparing the pointers, instead of using
// strcmp().
//
// This function must not be used before library constructors have finished
// running. In particular, this means it cannot be used to initialize global
// variables in C++.
func InternString(_string string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_intern_string(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
