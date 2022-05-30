// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkwaylanddevice.go.
var GTypeWaylandDevice = coreglib.Type(C.gdk_wayland_device_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeWaylandDevice, F: marshalWaylandDevice},
	})
}

// WaylandDeviceOverrider contains methods that are overridable.
type WaylandDeviceOverrider interface {
}

// WaylandDevice: wayland implementation of GdkDevice.
//
// Beyond the regular gdk.Device API, the Wayland implementation provides access
// to Wayland objects such as the wl_seat with
// gdkwayland.WaylandDevice.GetWlSeat(), the wl_keyboard with
// gdkwayland.WaylandDevice.GetWlKeyboard() and the wl_pointer with
// gdkwayland.WaylandDevice.GetWlPointer().
type WaylandDevice struct {
	_ [0]func() // equal guard
	gdk.Device
}

var (
	_ gdk.Devicer = (*WaylandDevice)(nil)
)

func classInitWaylandDevicer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWaylandDevice(obj *coreglib.Object) *WaylandDevice {
	return &WaylandDevice{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalWaylandDevice(p uintptr) (interface{}, error) {
	return wrapWaylandDevice(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NodePath returns the /dev/input/event* path of this device.
//
// For GdkDevices that possibly coalesce multiple hardware devices (eg. mouse,
// keyboard, touch,...), this function will return NULL.
//
// This is most notably implemented for devices of type GDK_SOURCE_PEN,
// GDK_SOURCE_TABLET_PAD.
//
// The function returns the following values:
//
//    - utf8 (optional): /dev/input/event* path of this device.
//
func (device *WaylandDevice) NodePath() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**WaylandDevice)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("GdkWayland", "WaylandDevice").InvokeMethod("get_node_path", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}
