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

// DecorationLayout gets the decoration layout set with
// gtk_header_bar_set_decoration_layout().
//
// The function returns the following values:
//
//    - utf8: decoration layout.
//
func (bar *HeaderBar) DecorationLayout() string {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_decoration_layout(_arg0)
	runtime.KeepAlive(bar)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasSubtitle retrieves whether the header bar reserves space for a subtitle,
// regardless if one is currently set or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the header bar reserves space for a subtitle.
//
func (bar *HeaderBar) HasSubtitle() bool {
	var _arg0 *C.GtkHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_has_subtitle(_arg0)
	runtime.KeepAlive(bar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDecorationLayout sets the decoration layout for this header bar,
// overriding the Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header bar
// design that does not allow for buttons to take room on the right, or only
// offers room for a single close button. Split header bars are another example
// for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close, icon (the window icon)
// and menu (a menu button for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu on the left, and
// minimize, maximize and close buttons on the right.
//
// The function takes the following parameters:
//
//    - layout (optional): decoration layout, or NULL to unset the layout.
//
func (bar *HeaderBar) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if layout != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_header_bar_set_decoration_layout(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(layout)
}

// SetHasSubtitle sets whether the header bar should reserve space for a
// subtitle, even if none is currently set.
//
// The function takes the following parameters:
//
//    - setting: TRUE to reserve space for a subtitle.
//
func (bar *HeaderBar) SetHasSubtitle(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_header_bar_set_has_subtitle(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(setting)
}
