// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// WarpPointer warps the pointer of display to the point x,y on the screen
// screen, unless the pointer is confined to a window by a grab, in which case
// it will be moved as far as allowed by the grab. Warping the pointer creates
// events as if the user had moved the mouse instantaneously to the destination.
//
// Note that the pointer should normally be under the control of the user. This
// function was added to cover some rare use cases like keyboard navigation
// support for the color picker in the ColorSelectionDialog.
//
// Deprecated: Use gdk_device_warp() instead.
//
// The function takes the following parameters:
//
//    - screen of display to warp the pointer to.
//    - x coordinate of the destination.
//    - y coordinate of the destination.
//
func (display *Display) WarpPointer(screen *Screen, x, y int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkScreen  // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gdk_display_warp_pointer(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(display)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}
