// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// X11DeviceGetID returns the device ID as seen by XInput2.
//
// The function takes the following parameters:
//
//    - device: Device.
//
// The function returns the following values:
//
//    - gint: XInput2 device ID.
//
func X11DeviceGetID(device *X11DeviceXI2) int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_info := girepository.MustFind("GdkX11", "x11_device_get_id")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}
