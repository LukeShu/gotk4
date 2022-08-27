// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// Ellipsize returns the ellipsizing position of the progress bar. See
// gtk_progress_bar_set_ellipsize().
//
// The function returns the following values:
//
//    - ellipsizeMode: EllipsizeMode.
//
func (pbar *ProgressBar) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkProgressBar    // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))

	_cret = C.gtk_progress_bar_get_ellipsize(_arg0)
	runtime.KeepAlive(pbar)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// SetEllipsize sets the mode used to ellipsize (add an ellipsis: "...") the
// text if there is not enough space to render the entire string.
//
// The function takes the following parameters:
//
//    - mode: EllipsizeMode.
//
func (pbar *ProgressBar) SetEllipsize(mode pango.EllipsizeMode) {
	var _arg0 *C.GtkProgressBar    // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.gtk_progress_bar_set_ellipsize(_arg0, _arg1)
	runtime.KeepAlive(pbar)
	runtime.KeepAlive(mode)
}
