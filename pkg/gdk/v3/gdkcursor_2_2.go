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

// NewCursorForDisplay creates a new cursor from the set of builtin cursors.
//
// The function takes the following parameters:
//
//    - display for which the cursor will be created.
//    - cursorType: cursor to create.
//
// The function returns the following values:
//
//    - cursor (optional): new Cursor, or NULL on failure.
//
func NewCursorForDisplay(display *Display, cursorType CursorType) *Cursor {
	var _arg1 *C.GdkDisplay   // out
	var _arg2 C.GdkCursorType // out
	var _cret *C.GdkCursor    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg2 = C.GdkCursorType(cursorType)

	_cret = C.gdk_cursor_new_for_display(_arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(cursorType)

	var _cursor *Cursor // out

	if _cret != nil {
		_cursor = wrapCursor(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _cursor
}

// Display returns the display on which the Cursor is defined.
//
// The function returns the following values:
//
//    - display associated to cursor.
//
func (cursor *Cursor) Display() *Display {
	var _arg0 *C.GdkCursor  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))

	_cret = C.gdk_cursor_get_display(_arg0)
	runtime.KeepAlive(cursor)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}
