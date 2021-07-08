// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_core_get_type()), F: marshalX11DeviceCore},
	})
}

type X11DeviceCore interface {
	gextras.Objector

	privateX11DeviceCoreClass()
}

// X11DeviceCoreClass implements the X11DeviceCore interface.
type X11DeviceCoreClass struct {
	gdk.DeviceClass
}

var _ X11DeviceCore = (*X11DeviceCoreClass)(nil)

func wrapX11DeviceCore(obj *externglib.Object) X11DeviceCore {
	return &X11DeviceCoreClass{
		DeviceClass: gdk.DeviceClass{
			Object: obj,
		},
	}
}

func marshalX11DeviceCore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11DeviceCore(obj), nil
}

func (*X11DeviceCoreClass) privateX11DeviceCoreClass() {}
