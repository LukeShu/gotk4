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

// Application gets the Application associated with the window (if any).
//
// The function returns the following values:
//
//    - application (optional) or NULL.
//
func (window *Window) Application() *Application {
	var _arg0 *C.GtkWindow      // out
	var _cret *C.GtkApplication // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gtk_window_get_application(_arg0)
	runtime.KeepAlive(window)

	var _application *Application // out

	if _cret != nil {
		_application = wrapApplication(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _application
}

// HasResizeGrip determines whether the window may have a resize grip.
//
// Deprecated: Resize grips have been removed.
//
// The function returns the following values:
//
//    - ok: TRUE if the window has a resize grip.
//
func (window *Window) HasResizeGrip() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gtk_window_get_has_resize_grip(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResizeGripArea: if a window has a resize grip, this will retrieve the grip
// position, width and height into the specified Rectangle.
//
// Deprecated: Resize grips have been removed.
//
// The function returns the following values:
//
//    - rect: pointer to a Rectangle which we should store the resize grip area.
//    - ok: TRUE if the resize grip’s area was retrieved.
//
func (window *Window) ResizeGripArea() (*gdk.Rectangle, bool) {
	var _arg0 *C.GtkWindow   // out
	var _arg1 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gtk_window_get_resize_grip_area(_arg0, &_arg1)
	runtime.KeepAlive(window)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// ResizeGripIsVisible determines whether a resize grip is visible for the
// specified window.
//
// Deprecated: Resize grips have been removed.
//
// The function returns the following values:
//
//    - ok: TRUE if a resize grip exists and is visible.
//
func (window *Window) ResizeGripIsVisible() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gtk_window_resize_grip_is_visible(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResizeToGeometry: like gtk_window_resize(), but width and height are
// interpreted in terms of the base size and increment set with
// gtk_window_set_geometry_hints.
//
// Deprecated: This function does nothing. Use gtk_window_resize() and compute
// the geometry yourself.
//
// The function takes the following parameters:
//
//    - width in resize increments to resize the window to.
//    - height in resize increments to resize the window to.
//
func (window *Window) ResizeToGeometry(width, height int) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)

	C.gtk_window_resize_to_geometry(_arg0, _arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetApplication sets or unsets the Application associated with the window.
//
// The application will be kept alive for at least as long as it has any windows
// associated with it (see g_application_hold() for a way to keep it alive
// without windows).
//
// Normally, the connection between the application and the window will remain
// until the window is destroyed, but you can explicitly remove it by setting
// the application to NULL.
//
// This is equivalent to calling gtk_application_remove_window() and/or
// gtk_application_add_window() on the old/new applications as relevant.
//
// The function takes the following parameters:
//
//    - application (optional) or NULL to unset.
//
func (window *Window) SetApplication(application *Application) {
	var _arg0 *C.GtkWindow      // out
	var _arg1 *C.GtkApplication // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if application != nil {
		_arg1 = (*C.GtkApplication)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	}

	C.gtk_window_set_application(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(application)
}

// SetDefaultGeometry: like gtk_window_set_default_size(), but width and height
// are interpreted in terms of the base size and increment set with
// gtk_window_set_geometry_hints.
//
// Deprecated: This function does nothing. If you want to set a default size,
// use gtk_window_set_default_size() instead.
//
// The function takes the following parameters:
//
//    - width in resize increments, or -1 to unset the default width.
//    - height in resize increments, or -1 to unset the default height.
//
func (window *Window) SetDefaultGeometry(width, height int) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)

	C.gtk_window_set_default_geometry(_arg0, _arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetHasResizeGrip sets whether window has a corner resize grip.
//
// Note that the resize grip is only shown if the window is actually resizable
// and not maximized. Use gtk_window_resize_grip_is_visible() to find out if the
// resize grip is currently shown.
//
// Deprecated: Resize grips have been removed.
//
// The function takes the following parameters:
//
//    - value: TRUE to allow a resize grip.
//
func (window *Window) SetHasResizeGrip(value bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if value {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_has_resize_grip(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(value)
}

// SetHasUserRefCount tells GTK+ whether to drop its extra reference to the
// window when gtk_widget_destroy() is called.
//
// This function is only exported for the benefit of language bindings which may
// need to keep the window alive until their wrapper object is garbage
// collected. There is no justification for ever calling this function in an
// application.
//
// The function takes the following parameters:
//
//    - setting: new value.
//
func (window *Window) SetHasUserRefCount(setting bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_has_user_ref_count(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(setting)
}
