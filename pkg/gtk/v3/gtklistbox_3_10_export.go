// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_ListBoxFilterFunc
func _gotk4_gtk3_ListBoxFilterFunc(arg1 *C.GtkListBoxRow, arg2 C.gpointer) (cret C.gboolean) {
	var fn ListBoxFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxFilterFunc)
	}

	var _row *ListBoxRow // out

	_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))

	ok := fn(_row)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_ListBoxSortFunc
func _gotk4_gtk3_ListBoxSortFunc(arg1 *C.GtkListBoxRow, arg2 *C.GtkListBoxRow, arg3 C.gpointer) (cret C.gint) {
	var fn ListBoxSortFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxSortFunc)
	}

	var _row1 *ListBoxRow // out
	var _row2 *ListBoxRow // out

	_row1 = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))
	_row2 = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg2)))

	gint := fn(_row1, _row2)

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtk3_ListBoxUpdateHeaderFunc
func _gotk4_gtk3_ListBoxUpdateHeaderFunc(arg1 *C.GtkListBoxRow, arg2 *C.GtkListBoxRow, arg3 C.gpointer) {
	var fn ListBoxUpdateHeaderFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ListBoxUpdateHeaderFunc)
	}

	var _row *ListBoxRow    // out
	var _before *ListBoxRow // out

	_row = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_before = wrapListBoxRow(coreglib.Take(unsafe.Pointer(arg2)))
	}

	fn(_row, _before)
}
