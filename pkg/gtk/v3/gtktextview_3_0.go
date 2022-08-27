// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// CursorLocations: given an iter within a text layout, determine the positions
// of the strong and weak cursors if the insertion point is at that iterator.
// The position of each cursor is stored as a zero-width rectangle. The strong
// cursor location is the location where characters of the directionality equal
// to the base direction of the paragraph are inserted. The weak cursor location
// is the location where characters of the directionality opposite to the base
// direction of the paragraph are inserted.
//
// If iter is NULL, the actual cursor position is used.
//
// Note that if iter happens to be the actual cursor position, and there is
// currently an IM preedit sequence being entered, the returned locations will
// be adjusted to account for the preedit cursor’s offset within the preedit
// sequence.
//
// The rectangle position is in buffer coordinates; use
// gtk_text_view_buffer_to_window_coords() to convert these coordinates to
// coordinates for one of the windows in the text view.
//
// The function takes the following parameters:
//
//    - iter (optional): TextIter.
//
// The function returns the following values:
//
//    - strong (optional): location to store the strong cursor position (may be
//      NULL).
//    - weak (optional): location to store the weak cursor position (may be
//      NULL).
//
func (textView *TextView) CursorLocations(iter *TextIter) (strong, weak *gdk.Rectangle) {
	var _arg0 *C.GtkTextView // out
	var _arg1 *C.GtkTextIter // out
	var _arg2 C.GdkRectangle // in
	var _arg3 C.GdkRectangle // in

	_arg0 = (*C.GtkTextView)(unsafe.Pointer(coreglib.InternObject(textView).Native()))
	if iter != nil {
		_arg1 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}

	C.gtk_text_view_get_cursor_locations(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(textView)
	runtime.KeepAlive(iter)

	var _strong *gdk.Rectangle // out
	var _weak *gdk.Rectangle   // out

	_strong = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_weak = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _strong, _weak
}
