// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkapplaunchcontext.go.
var GTypeAppLaunchContext = coreglib.Type(C.gdk_app_launch_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAppLaunchContext, F: marshalAppLaunchContext},
	})
}

// AppLaunchContext is an implementation of LaunchContext that handles launching
// an application in a graphical context. It provides startup notification and
// allows to launch applications on a specific screen or workspace.
//
// Launching an application
//
//    GdkAppLaunchContext *context;
//
//    context = gdk_display_get_app_launch_context (display);
//
//    gdk_app_launch_context_set_screen (screen);
//    gdk_app_launch_context_set_timestamp (event->time);
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
	_ coreglib.Objector = (*AppLaunchContext)(nil)
)

func wrapAppLaunchContext(obj *coreglib.Object) *AppLaunchContext {
	return &AppLaunchContext{
		AppLaunchContext: gio.AppLaunchContext{
			Object: obj,
		},
	}
}

func marshalAppLaunchContext(p uintptr) (interface{}, error) {
	return wrapAppLaunchContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAppLaunchContext creates a new AppLaunchContext.
//
// Deprecated: Use gdk_display_get_app_launch_context() instead.
//
// The function returns the following values:
//
//    - appLaunchContext: new AppLaunchContext.
//
func NewAppLaunchContext() *AppLaunchContext {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("new_AppLaunchContext", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _appLaunchContext *AppLaunchContext // out

	_appLaunchContext = wrapAppLaunchContext(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appLaunchContext
}

// SetDesktop sets the workspace on which applications will be launched when
// using this context when running under a window manager that supports multiple
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
func (context *AppLaunchContext) SetDesktop(desktop int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.gint(desktop)
	*(**AppLaunchContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_desktop", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(desktop)
}

// SetDisplay sets the display on which applications will be launched when using
// this context. See also gdk_app_launch_context_set_screen().
//
// Deprecated: Use gdk_display_get_app_launch_context() instead.
//
// The function takes the following parameters:
//
//    - display: Display.
//
func (context *AppLaunchContext) SetDisplay(display *Display) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	*(**AppLaunchContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_display", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(display)
}

// SetIcon sets the icon for applications that are launched with this context.
//
// Window Managers can use this information when displaying startup
// notification.
//
// See also gdk_app_launch_context_set_icon_name().
//
// The function takes the following parameters:
//
//    - icon (optional) or NULL.
//
func (context *AppLaunchContext) SetIcon(icon gio.Iconner) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if icon != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	}
	*(**AppLaunchContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_icon", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(icon)
}

// SetIconName sets the icon for applications that are launched with this
// context. The icon_name will be interpreted in the same way as the Icon field
// in desktop files. See also gdk_app_launch_context_set_icon().
//
// If both icon and icon_name are set, the icon_name takes priority. If neither
// icon or icon_name is set, the icon is taken from either the file that is
// passed to launched application or from the Info for the launched application
// itself.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name, or NULL.
//
func (context *AppLaunchContext) SetIconName(iconName string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if iconName != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**AppLaunchContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_icon_name", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(iconName)
}

// SetScreen sets the screen on which applications will be launched when using
// this context. See also gdk_app_launch_context_set_display().
//
// If both screen and display are set, the screen takes priority. If neither
// screen or display are set, the default screen and display are used.
//
// The function takes the following parameters:
//
//    - screen: Screen.
//
func (context *AppLaunchContext) SetScreen(screen *Screen) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	*(**AppLaunchContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_screen", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(screen)
}

// SetTimestamp sets the timestamp of context. The timestamp should ideally be
// taken from the event that triggered the launch.
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
	var args [2]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.guint32 // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.guint32(timestamp)
	*(**AppLaunchContext)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "AppLaunchContext").InvokeMethod("set_timestamp", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(timestamp)
}
