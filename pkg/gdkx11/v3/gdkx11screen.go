// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screener},
	})
}

// X11GetDefaultScreen gets the default GTK+ screen number.
func X11GetDefaultScreen() int {
	var _cret C.gint // in

	_cret = C.gdk_x11_get_default_screen()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

type X11Screen struct {
	gdk.Screen
}

func wrapX11Screen(obj *externglib.Object) *X11Screen {
	return &X11Screen{
		Screen: gdk.Screen{
			Object: obj,
		},
	}
}

func marshalX11Screener(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Screen(obj), nil
}

// CurrentDesktop returns the current workspace for screen when running under a
// window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
func (screen *X11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkScreen // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// NumberOfDesktops returns the number of workspaces for screen when running
// under a window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
func (screen *X11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkScreen // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// ScreenNumber returns the index of a Screen.
func (screen *X11Screen) ScreenNumber() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.int        // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WindowManagerName returns the name of the window manager for screen.
func (screen *X11Screen) WindowManagerName() string {
	var _arg0 *C.GdkScreen // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
