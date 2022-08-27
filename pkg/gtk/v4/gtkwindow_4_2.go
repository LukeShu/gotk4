// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

// HandleMenubarAccel returns whether this window reacts to F10 key presses by
// activating a menubar it contains.
//
// The function returns the following values:
//
//    - ok: TRUE if the window handles F10.
//
func (window *Window) HandleMenubarAccel() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gtk_window_get_handle_menubar_accel(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetHandleMenubarAccel sets whether this window should react to F10 key
// presses by activating a menubar it contains.
//
// The function takes the following parameters:
//
//    - handleMenubarAccel: TRUE to make window handle F10.
//
func (window *Window) SetHandleMenubarAccel(handleMenubarAccel bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if handleMenubarAccel {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_handle_menubar_accel(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(handleMenubarAccel)
}
