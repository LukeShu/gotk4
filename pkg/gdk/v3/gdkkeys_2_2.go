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

// KeymapGetForDisplay returns the Keymap attached to display.
//
// The function takes the following parameters:
//
//    - display: Display.
//
// The function returns the following values:
//
//    - keymap attached to display.
//
func KeymapGetForDisplay(display *Display) *Keymap {
	var _arg1 *C.GdkDisplay // out
	var _cret *C.GdkKeymap  // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gdk_keymap_get_for_display(_arg1)
	runtime.KeepAlive(display)

	var _keymap *Keymap // out

	_keymap = wrapKeymap(coreglib.Take(unsafe.Pointer(_cret)))

	return _keymap
}
