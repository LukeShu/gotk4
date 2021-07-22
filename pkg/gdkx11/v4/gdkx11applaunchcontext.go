// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_app_launch_context_get_type()), F: marshalX11AppLaunchContexter},
	})
}

type X11AppLaunchContext struct {
	gdk.AppLaunchContext
}

func wrapX11AppLaunchContext(obj *externglib.Object) *X11AppLaunchContext {
	return &X11AppLaunchContext{
		AppLaunchContext: gdk.AppLaunchContext{
			AppLaunchContext: gio.AppLaunchContext{
				Object: obj,
			},
		},
	}
}

func marshalX11AppLaunchContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11AppLaunchContext(obj), nil
}

func (*X11AppLaunchContext) privateX11AppLaunchContext() {}
