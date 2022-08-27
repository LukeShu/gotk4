// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// Direction gets the text direction at the given character position in layout.
//
// The function takes the following parameters:
//
//    - index: byte index of the char.
//
// The function returns the following values:
//
//    - direction: text direction at index.
//
func (layout *Layout) Direction(index int) Direction {
	var _arg0 *C.PangoLayout   // out
	var _arg1 C.int            // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	_arg1 = C.int(index)

	_cret = C.pango_layout_get_direction(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(index)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}
