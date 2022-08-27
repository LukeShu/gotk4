// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// SettingsGetForScreen gets the Settings object for screen, creating it if
// necessary.
//
// The function takes the following parameters:
//
//    - screen: Screen.
//
// The function returns the following values:
//
//    - settings Settings object.
//
func SettingsGetForScreen(screen *gdk.Screen) *Settings {
	var _arg1 *C.GdkScreen   // out
	var _cret *C.GtkSettings // in

	_arg1 = (*C.GdkScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_cret = C.gtk_settings_get_for_screen(_arg1)
	runtime.KeepAlive(screen)

	var _settings *Settings // out

	_settings = wrapSettings(coreglib.Take(unsafe.Pointer(_cret)))

	return _settings
}
