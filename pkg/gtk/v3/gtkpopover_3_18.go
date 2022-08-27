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

// DefaultWidget gets the widget that should be set as the default while the
// popover is shown.
//
// The function returns the following values:
//
//    - widget (optional): default widget, or NULL if there is none.
//
func (popover *Popover) DefaultWidget() Widgetter {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_cret = C.gtk_popover_get_default_widget(_arg0)
	runtime.KeepAlive(popover)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetDefaultWidget sets the widget that should be set as default widget while
// the popover is shown (see gtk_window_set_default()). Popover remembers the
// previous default widget and reestablishes it when the popover is dismissed.
//
// The function takes the following parameters:
//
//    - widget (optional): new default widget, or NULL.
//
func (popover *Popover) SetDefaultWidget(widget Widgetter) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	C.gtk_popover_set_default_widget(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(widget)
}
