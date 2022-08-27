// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// AddMark adds the mark at position where. The mark must not be added to
// another buffer, and if its name is not NULL then there must not be another
// mark in the buffer with the same name.
//
// Emits the TextBuffer::mark-set signal as notification of the mark's initial
// placement.
//
// The function takes the following parameters:
//
//    - mark to add.
//    - where: location to place mark.
//
func (buffer *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	var _arg0 *C.GtkTextBuffer // out
	var _arg1 *C.GtkTextMark   // out
	var _arg2 *C.GtkTextIter   // out

	_arg0 = (*C.GtkTextBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))
	_arg1 = (*C.GtkTextMark)(unsafe.Pointer(coreglib.InternObject(mark).Native()))
	_arg2 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(where)))

	C.gtk_text_buffer_add_mark(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(mark)
	runtime.KeepAlive(where)
}
