// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"
)

// #cgo pkg-config: gtk4 gtk4-x11
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11SetSmClientID sets the `SM_CLIENT_ID` property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientId string) {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(smClientId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_x11_set_sm_client_id(_arg1)
}
