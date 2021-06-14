// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_app_launch_context_get_type()), F: marshalAppLaunchContext},
	})
}

// AppLaunchContext: `GdkAppLaunchContext` handles launching an application in a
// graphical context.
//
// It is an implementation of `GAppLaunchContext` that provides startup
// notification and allows to launch applications on a specific screen or
// workspace.
//
//
// Launching an application
//
// “`c GdkAppLaunchContext *context;
//
// context = gdk_display_get_app_launch_context (display);
//
// gdk_app_launch_context_set_display (display);
// gdk_app_launch_context_set_timestamp (gdk_event_get_time (event));
//
// if (!g_app_info_launch_default_for_uri ("http://www.gtk.org", context,
// &error)) g_warning ("Launching failed: s\n", error->message);
//
// g_object_unref (context); “`
type AppLaunchContext interface {
	AppLaunchContext

	// Display gets the `GdkDisplay` that @context is for.
	Display() Display
	// SetDesktop sets the workspace on which applications will be launched.
	//
	// This only works when running under a window manager that supports
	// multiple workspaces, as described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec).
	//
	// When the workspace is not specified or @desktop is set to -1, it is up to
	// the window manager to pick one, typically it will be the current
	// workspace.
	SetDesktop(desktop int)
	// SetIconName sets the icon for applications that are launched with this
	// context.
	//
	// The @icon_name will be interpreted in the same way as the Icon field in
	// desktop files. See also [method@Gdk.AppLaunchContext.set_icon()].
	//
	// If both @icon and @icon_name are set, the @icon_name takes priority. If
	// neither @icon or @icon_name is set, the icon is taken from either the
	// file that is passed to launched application or from the `GAppInfo` for
	// the launched application itself.
	SetIconName(iconName string)
	// SetTimestamp sets the timestamp of @context.
	//
	// The timestamp should ideally be taken from the event that triggered the
	// launch.
	//
	// Window managers can use this information to avoid moving the focus to the
	// newly launched application when the user is busy typing in another
	// window. This is also known as 'focus stealing prevention'.
	SetTimestamp(timestamp uint32)
}

// appLaunchContext implements the AppLaunchContext class.
type appLaunchContext struct {
	AppLaunchContext
}

var _ AppLaunchContext = (*appLaunchContext)(nil)

// WrapAppLaunchContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppLaunchContext(obj *externglib.Object) AppLaunchContext {
	return appLaunchContext{
		AppLaunchContext: WrapAppLaunchContext(obj),
	}
}

func marshalAppLaunchContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppLaunchContext(obj), nil
}

// Display gets the `GdkDisplay` that @context is for.
func (c appLaunchContext) Display() Display {
	var _arg0 *C.GdkAppLaunchContext // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(c.Native()))

	var _cret *C.GdkDisplay // in

	_cret = C.gdk_app_launch_context_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Display)

	return _display
}

// SetDesktop sets the workspace on which applications will be launched.
//
// This only works when running under a window manager that supports
// multiple workspaces, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec).
//
// When the workspace is not specified or @desktop is set to -1, it is up to
// the window manager to pick one, typically it will be the current
// workspace.
func (c appLaunchContext) SetDesktop(desktop int) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 C.int                  // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(desktop)

	C.gdk_app_launch_context_set_desktop(_arg0, _arg1)
}

// SetIconName sets the icon for applications that are launched with this
// context.
//
// The @icon_name will be interpreted in the same way as the Icon field in
// desktop files. See also [method@Gdk.AppLaunchContext.set_icon()].
//
// If both @icon and @icon_name are set, the @icon_name takes priority. If
// neither @icon or @icon_name is set, the icon is taken from either the
// file that is passed to launched application or from the `GAppInfo` for
// the launched application itself.
func (c appLaunchContext) SetIconName(iconName string) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_app_launch_context_set_icon_name(_arg0, _arg1)
}

// SetTimestamp sets the timestamp of @context.
//
// The timestamp should ideally be taken from the event that triggered the
// launch.
//
// Window managers can use this information to avoid moving the focus to the
// newly launched application when the user is busy typing in another
// window. This is also known as 'focus stealing prevention'.
func (c appLaunchContext) SetTimestamp(timestamp uint32) {
	var _arg0 *C.GdkAppLaunchContext // out
	var _arg1 C.guint32              // out

	_arg0 = (*C.GdkAppLaunchContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint32(timestamp)

	C.gdk_app_launch_context_set_timestamp(_arg0, _arg1)
}
