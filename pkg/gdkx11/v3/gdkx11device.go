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
import "C"

// X11DeviceGetID returns the device ID as seen by XInput2.
//
// > If gdk_disable_multidevice() has been called, this function > will
// respectively return 2/3 for the core pointer and keyboard, > (matching the
// IDs for the Virtual Core Pointer and Keyboard in > XInput 2), but calling
// this function on any slave devices (i.e. > those managed via XInput 1.x),
// will return 0.
//
// The function takes the following parameters:
//
//    - device: Device.
//
// The function returns the following values:
//
//    - gint: XInput2 device ID.
//
func X11DeviceGetID(device *X11DeviceCore) int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**X11DeviceCore)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("GdkX11", "x11_device_get_id").Invoke(args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(device)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}
