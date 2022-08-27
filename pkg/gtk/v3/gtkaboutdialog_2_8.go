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

// WrapLicense returns whether the license text in about is automatically
// wrapped.
//
// The function returns the following values:
//
//    - ok: TRUE if the license text is wrapped.
//
func (about *AboutDialog) WrapLicense() bool {
	var _arg0 *C.GtkAboutDialog // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(coreglib.InternObject(about).Native()))

	_cret = C.gtk_about_dialog_get_wrap_license(_arg0)
	runtime.KeepAlive(about)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetWrapLicense sets whether the license text in about is automatically
// wrapped.
//
// The function takes the following parameters:
//
//    - wrapLicense: whether to wrap the license.
//
func (about *AboutDialog) SetWrapLicense(wrapLicense bool) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(coreglib.InternObject(about).Native()))
	if wrapLicense {
		_arg1 = C.TRUE
	}

	C.gtk_about_dialog_set_wrap_license(_arg0, _arg1)
	runtime.KeepAlive(about)
	runtime.KeepAlive(wrapLicense)
}
