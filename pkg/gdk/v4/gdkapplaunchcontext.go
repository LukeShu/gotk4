// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_app_launch_context_get_type()), F: marshalAppLaunchContexter},
	})
}

// AppLaunchContext: GdkAppLaunchContext handles launching an application in a
// graphical context.
//
// It is an implementation of GAppLaunchContext that provides startup
// notification and allows to launch applications on a specific screen or
// workspace.
//
// Launching an application
//
//    GdkAppLaunchContext *context;
//
//    context = gdk_display_get_app_launch_context (display);
//
//    gdk_app_launch_context_set_display (display);
//    gdk_app_launch_context_set_timestamp (gdk_event_get_time (event));
//
//    if (!g_app_info_launch_default_for_uri ("http://www.gtk.org", context, &error))
//      g_warning ("Launching failed: s\n", error->message);
//
//    g_object_unref (context);.
type AppLaunchContext struct {
	_ [0]func() // equal guard
	gio.AppLaunchContext
}

var (
	_ externglib.Objector = (*AppLaunchContext)(nil)
)

func wrapAppLaunchContext(obj *externglib.Object) *AppLaunchContext {
	return &AppLaunchContext{
		AppLaunchContext: gio.AppLaunchContext{
			Object: obj,
		},
	}
}

func marshalAppLaunchContexter(p uintptr) (interface{}, error) {
	return wrapAppLaunchContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Display gets the GdkDisplay that context is for.
//
// The function returns the following values:
//
//    - display of context.
//
func (context *AppLaunchContext) Display() *Display {
	var _arg0 *C.GdkAppLaunchContext // out
	var _cret *C.GdkDisplay          // in

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_app_launch_context_get_display(_arg0)
	runtime.KeepAlive(context)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// SetDesktop sets the workspace on which applications will be launched.
//
// This only works when running under a window manager that supports multiple
// workspaces, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec).
//
// When the workspace is not specified or desktop is set to -1, it is up to the
// window manager to pick one, typically it will be the current workspace.
//
// The function takes the following parameters:
//
//    - desktop: number of a workspace, or -1.
//
func (context *AppLaunchContext) SetDesktop(desktop int) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 C.int                  // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.int(desktop)

	C.gdk_app_launch_context_set_desktop(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(desktop)
}

// SetIcon sets the icon for applications that are launched with this context.
//
// Window Managers can use this information when displaying startup
// notification.
//
// See also gdk.AppLaunchContext.SetIconName().
//
// The function takes the following parameters:
//
//    - icon (optional) or NULL.
//
func (context *AppLaunchContext) SetIcon(icon gio.Iconner) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 *C.GIcon               // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	if icon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	}

	C.gdk_app_launch_context_set_icon(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(icon)
}

// SetIconName sets the icon for applications that are launched with this
// context.
//
// The icon_name will be interpreted in the same way as the Icon field in
// desktop files. See also gdk.AppLaunchContext.SetIcon()().
//
// If both icon and icon_name are set, the icon_name takes priority. If neither
// icon or icon_name is set, the icon is taken from either the file that is
// passed to launched application or from the GAppInfo for the launched
// application itself.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name, or NULL.
//
func (context *AppLaunchContext) SetIconName(iconName string) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gdk_app_launch_context_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(iconName)
}

// SetTimestamp sets the timestamp of context.
//
// The timestamp should ideally be taken from the event that triggered the
// launch.
//
// Window managers can use this information to avoid moving the focus to the
// newly launched application when the user is busy typing in another window.
// This is also known as 'focus stealing prevention'.
//
// The function takes the following parameters:
//
//    - timestamp: timestamp.
//
func (context *AppLaunchContext) SetTimestamp(timestamp uint32) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 C.guint32              // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.guint32(timestamp)

	C.gdk_app_launch_context_set_timestamp(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(timestamp)
}
