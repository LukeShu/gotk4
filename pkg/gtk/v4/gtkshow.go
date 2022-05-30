// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// ShowURI: this function launches the default application for showing a given
// uri, or shows an error dialog if that fails.
//
// The function takes the following parameters:
//
//    - parent (optional) window.
//    - uri to show.
//    - timestamp from the event that triggered this call, or GDK_CURRENT_TIME.
//
func ShowURI(parent *Window, uri string, timestamp uint32) {
	var args [3]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 *C.void   // out
	var _arg2 C.guint32 // out

	if parent != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint32(timestamp)
	*(**Window)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1
	*(*uint32)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "show_uri").Invoke(args[:], nil)

	runtime.KeepAlive(parent)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(timestamp)
}

// ShowURIFullFinish finishes the gtk_show_uri() call and returns the result of
// the operation.
//
// The function takes the following parameters:
//
//    - parent passed to gtk_show_uri().
//    - result that was passed to callback.
//
func ShowURIFullFinish(parent *Window, result gio.AsyncResulter) error {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))
	*(**Window)(unsafe.Pointer(&args[0])) = _arg0
	*(*gio.AsyncResulter)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "show_uri_full_finish").Invoke(args[:], nil)

	runtime.KeepAlive(parent)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
