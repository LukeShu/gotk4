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

// NewOverlay creates a new Overlay.
//
// The function returns the following values:
//
//    - overlay: new Overlay object.
//
func NewOverlay() *Overlay {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_overlay_new()

	var _overlay *Overlay // out

	_overlay = wrapOverlay(coreglib.Take(unsafe.Pointer(_cret)))

	return _overlay
}

// AddOverlay adds widget to overlay.
//
// The widget will be stacked on top of the main widget added with
// gtk_container_add().
//
// The position at which widget is placed is determined from its Widget:halign
// and Widget:valign properties.
//
// The function takes the following parameters:
//
//    - widget to be added to the container.
//
func (overlay *Overlay) AddOverlay(widget Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.gtk_overlay_add_overlay(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
}
