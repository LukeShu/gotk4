// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/gdk"
	"github.com/diamondburned/gotk4/xlib"
	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{
		// Enums
		// Skipped X11DeviceType.

		// Objects/Classes
	})
}

type X11DeviceType int

const (
	X11DeviceTypeLogical X11DeviceType = 0

	X11DeviceTypePhysical X11DeviceType = 1

	X11DeviceTypeFloating X11DeviceType = 2
)

// X11DeviceGetID: returns the device ID as seen by XInput2.
func X11DeviceGetID(device *X11DeviceXI2) int

// X11DeviceManagerLookup: returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager *X11DeviceManagerXI2, deviceID int) *X11DeviceXI2

// X11FreeCompoundText: frees the data returned from
// gdk_x11_display_string_to_compound_text().
func X11FreeCompoundText(ctext uint8)

// X11FreeTextList: frees the array of strings created by
// gdk_x11_display_text_property_to_text_list().
func X11FreeTextList(list string)

// X11GetServerTime: routine to get the current X server time stamp.
func X11GetServerTime(surface *X11Surface) uint32

// X11GetXatomByNameForDisplay: returns the X atom for a Display corresponding
// to @atom_name. This function caches the result, so if called repeatedly it is
// much faster than XInternAtom(), which is a round trip to the server each
// time.
func X11GetXatomByNameForDisplay(display *X11Display, atomName string) xlib.Atom

// X11GetXatomNameForDisplay: returns the name of an X atom for its display.
// This function is meant mainly for debugging, so for convenience, unlike
// XAtomName() and the result doesn’t need to be freed.
func X11GetXatomNameForDisplay(display *X11Display, xatom xlib.Atom) string

// X11LookupXdisplay: find the Display corresponding to @xdisplay, if any
// exists.
func X11LookupXdisplay(xdisplay *xlib.Display) *X11Display

// X11SetSmClientID: sets the `SM_CLIENT_ID` property on the application’s
// leader window so that the window manager can save the application’s state
// using the X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientID string)

type X11AppLaunchContext struct {
	gdk.AppLaunchContext
}

type X11DeviceManagerXI2 struct {
	*glib.Object
}

type X11DeviceXI2 struct {
	gdk.Device
}

type X11Display struct {
	gdk.Display
}

type X11Drag struct {
	gdk.Drag
}

type X11Monitor struct {
	gdk.Monitor
}

type X11Screen struct {
	*glib.Object
}

type X11Surface struct {
	gdk.Surface
}
