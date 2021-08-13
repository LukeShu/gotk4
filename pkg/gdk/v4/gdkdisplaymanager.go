// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_manager_get_type()), F: marshalDisplayManagerer},
	})
}

// SetAllowedBackends sets a list of backends that GDK should try to use.
//
// This can be useful if your application does not work with certain GDK
// backends.
//
// By default, GDK tries all included backends.
//
// For example:
//
//    gdk_set_allowed_backends ("wayland,macos,*");
//
//
// instructs GDK to try the Wayland backend first, followed by the MacOs
// backend, and then all others.
//
// If the GDK_BACKEND environment variable is set, it determines what backends
// are tried in what order, while still respecting the set of allowed backends
// that are specified by this function.
//
// The possible backend names are:
//
//    - broadway
//    - macos
//    - wayland.
//    - win32
//    - x11
//
// You can also include a * in the list to try all remaining backends.
//
// This call must happen prior to functions that open a display, such as
// gdk.Display().Open, gtk_init(), or gtk_init_check() in order to take effect.
func SetAllowedBackends(backends string) {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(backends)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_set_allowed_backends(_arg1)
	runtime.KeepAlive(backends)
}

// DisplayManager: singleton object that offers notification when displays
// appear or disappear.
//
// You can use gdk.DisplayManager().Get to obtain the GdkDisplayManager
// singleton, but that should be rarely necessary. Typically, initializing GTK
// opens a display that you can work with without ever accessing the
// GdkDisplayManager.
//
// The GDK library can be built with support for multiple backends. The
// GdkDisplayManager object determines which backend is used at runtime.
//
// In the rare case that you need to influence which of the backends is being
// used, you can use gdk.SetAllowedBackends(). Note that you need to call this
// function before initializing GTK.
//
//
// Backend-specific code
//
// When writing backend-specific code that is supposed to work with multiple GDK
// backends, you have to consider both compile time and runtime. At compile
// time, use the K_WINDOWING_X11, K_WINDOWING_WIN32 macros, etc. to find out
// which backends are present in the GDK library you are building your
// application against. At runtime, use type-check macros like
// GDK_IS_X11_DISPLAY() to find out which backend is in use:
//
//    #ifdef GDK_WINDOWING_X11
//      if (GDK_IS_X11_DISPLAY (display))
//        {
//          // make X11-specific calls here
//        }
//      else
//    #endif
//    #ifdef GDK_WINDOWING_MACOS
//      if (GDK_IS_MACOS_DISPLAY (display))
//        {
//          // make Quartz-specific calls here
//        }
//      else
//    #endif
//      g_error ("Unsupported GDK backend");
type DisplayManager struct {
	*externglib.Object
}

func wrapDisplayManager(obj *externglib.Object) *DisplayManager {
	return &DisplayManager{
		Object: obj,
	}
}

func marshalDisplayManagerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDisplayManager(obj), nil
}

// DefaultDisplay gets the default GdkDisplay.
func (manager *DisplayManager) DefaultDisplay() *Display {
	var _arg0 *C.GdkDisplayManager // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gdk_display_manager_get_default_display(_arg0)

	runtime.KeepAlive(manager)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// OpenDisplay opens a display.
func (manager *DisplayManager) OpenDisplay(name string) *Display {
	var _arg0 *C.GdkDisplayManager // out
	var _arg1 *C.char              // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_display_manager_open_display(_arg0, _arg1)

	runtime.KeepAlive(manager)
	runtime.KeepAlive(name)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// SetDefaultDisplay sets display as the default display.
func (manager *DisplayManager) SetDefaultDisplay(display *Display) {
	var _arg0 *C.GdkDisplayManager // out
	var _arg1 *C.GdkDisplay        // out

	_arg0 = (*C.GdkDisplayManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_manager_set_default_display(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(display)
}

// DisplayManagerGet gets the singleton GdkDisplayManager object.
//
// When called for the first time, this function consults the GDK_BACKEND
// environment variable to find out which of the supported GDK backends to use
// (in case GDK has been compiled with multiple backends).
//
// Applications can use set_allowed_backends to limit what backends wil be used.
func DisplayManagerGet() *DisplayManager {
	var _cret *C.GdkDisplayManager // in

	_cret = C.gdk_display_manager_get()

	var _displayManager *DisplayManager // out

	_displayManager = wrapDisplayManager(externglib.Take(unsafe.Pointer(_cret)))

	return _displayManager
}
