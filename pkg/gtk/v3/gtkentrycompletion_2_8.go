// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// PopupSetWidth returns whether the completion popup window will be resized to
// the width of the entry.
//
// The function returns the following values:
//
//    - ok: TRUE if the popup window will be resized to the width of the entry.
//
func (completion *EntryCompletion) PopupSetWidth() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_cret = C.gtk_entry_completion_get_popup_set_width(_arg0)
	runtime.KeepAlive(completion)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSingleMatch returns whether the completion popup window will appear even
// if there is only a single match.
//
// The function returns the following values:
//
//    - ok: TRUE if the popup window will appear regardless of the number of
//      matches.
//
func (completion *EntryCompletion) PopupSingleMatch() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_cret = C.gtk_entry_completion_get_popup_single_match(_arg0)
	runtime.KeepAlive(completion)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPopupSetWidth sets whether the completion popup window will be resized to
// be the same width as the entry.
//
// The function takes the following parameters:
//
//    - popupSetWidth: TRUE to make the width of the popup the same as the entry.
//
func (completion *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if popupSetWidth {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_set_width(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(popupSetWidth)
}

// SetPopupSingleMatch sets whether the completion popup window will appear even
// if there is only a single match. You may want to set this to FALSE if you are
// using [inline completion][GtkEntryCompletion--inline-completion].
//
// The function takes the following parameters:
//
//    - popupSingleMatch: TRUE if the popup should appear even for a single
//      match.
//
func (completion *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if popupSingleMatch {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_single_match(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(popupSingleMatch)
}
