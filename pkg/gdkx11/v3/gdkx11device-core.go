// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkx11device-core.go.
var GTypeX11DeviceCore = externglib.Type(C.gdk_x11_device_core_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeX11DeviceCore, F: marshalX11DeviceCore},
	})
}

// X11DeviceCoreOverrider contains methods that are overridable.
type X11DeviceCoreOverrider interface {
	externglib.Objector
}

// WrapX11DeviceCoreOverrider wraps the X11DeviceCoreOverrider
// interface implementation to access the instance methods.
func WrapX11DeviceCoreOverrider(obj X11DeviceCoreOverrider) *X11DeviceCore {
	return wrapX11DeviceCore(externglib.BaseObject(obj))
}

type X11DeviceCore struct {
	_ [0]func() // equal guard
	gdk.Device
}

var (
	_ gdk.Devicer = (*X11DeviceCore)(nil)
)

func classInitX11DeviceCorer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11DeviceCore(obj *externglib.Object) *X11DeviceCore {
	return &X11DeviceCore{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalX11DeviceCore(p uintptr) (interface{}, error) {
	return wrapX11DeviceCore(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
