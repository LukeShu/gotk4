// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

// IOErrorFromErrno converts errno.h error codes into GIO error codes. The
// fallback value G_IO_ERROR_FAILED is returned for error codes not currently
// handled (but note that future GLib releases may return a more specific value
// instead).
//
// As errno is global and may be modified by intermediate function calls, you
// should save its value as soon as the call which sets it
func IOErrorFromErrno(errNo int32) IOErrorEnum {
	var _arg1 C.gint         // out
	var _cret C.GIOErrorEnum // in

	_arg1 = C.gint(errNo)

	_cret = C.g_io_error_from_errno(_arg1)
	runtime.KeepAlive(errNo)

	var _ioErrorEnum IOErrorEnum // out

	_ioErrorEnum = IOErrorEnum(_cret)

	return _ioErrorEnum
}
