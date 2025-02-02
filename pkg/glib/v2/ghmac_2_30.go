// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// ComputeHMACForData computes the HMAC for a binary data of length. This is a
// convenience wrapper for g_hmac_new(), g_hmac_get_string() and g_hmac_unref().
//
// The hexadecimal string returned will be in lower case.
//
// The function takes the following parameters:
//
//   - digestType to use for the HMAC.
//   - key to use in the HMAC.
//   - data: binary blob to compute the HMAC of.
//
// The function returns the following values:
//
//   - utf8: HMAC of the binary data as a string in hexadecimal. The returned
//     string should be freed with g_free() when done using it.
//
func ComputeHMACForData(digestType ChecksumType, key, data []byte) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.guchar       // out
	var _arg3 C.gsize
	var _arg4 *C.guchar // out
	var _arg5 C.gsize
	var _cret *C.gchar // in

	_arg1 = C.GChecksumType(digestType)
	_arg3 = (C.gsize)(len(key))
	if len(key) > 0 {
		_arg2 = (*C.guchar)(unsafe.Pointer(&key[0]))
	}
	_arg5 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg4 = (*C.guchar)(unsafe.Pointer(&data[0]))
	}

	_cret = C.g_compute_hmac_for_data(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(digestType)
	runtime.KeepAlive(key)
	runtime.KeepAlive(data)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ComputeHMACForString computes the HMAC for a string.
//
// The hexadecimal string returned will be in lower case.
//
// The function takes the following parameters:
//
//   - digestType to use for the HMAC.
//   - key to use in the HMAC.
//   - str: string to compute the HMAC for.
//   - length of the string, or -1 if the string is nul-terminated.
//
// The function returns the following values:
//
//   - utf8: HMAC as a hexadecimal string. The returned string should be freed
//     with g_free() when done using it.
//
func ComputeHMACForString(digestType ChecksumType, key []byte, str string, length int) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.guchar       // out
	var _arg3 C.gsize
	var _arg4 *C.gchar // out
	var _arg5 C.gssize // out
	var _cret *C.gchar // in

	_arg1 = C.GChecksumType(digestType)
	_arg3 = (C.gsize)(len(key))
	if len(key) > 0 {
		_arg2 = (*C.guchar)(unsafe.Pointer(&key[0]))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.gssize(length)

	_cret = C.g_compute_hmac_for_string(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(digestType)
	runtime.KeepAlive(key)
	runtime.KeepAlive(str)
	runtime.KeepAlive(length)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
