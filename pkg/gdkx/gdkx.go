// Code generated by girgen. DO NOT EDIT.

package gdkx

import (
	"github.com/diamondburned/gotk4/cairo"
	"github.com/diamondburned/gotk4/gdk"
	"github.com/diamondburned/gotk4/xlib"
	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{

		// Objects/Classes
	})
}

// X11AtomToXatom: converts from a Atom to the X atom for the default GDK
// display with the same string value.
func X11AtomToXatom(atom gdk.Atom) xlib.Atom

// X11AtomToXatomForDisplay: converts from a Atom to the X atom for a Display
// with the same string value. The special value GDK_NONE is converted to None.
func X11AtomToXatomForDisplay(display *X11Display, atom gdk.Atom) xlib.Atom

// X11DeviceGetID: returns the device ID as seen by XInput2.
//
// > If gdk_disable_multidevice() has been called, this function > will
// respectively return 2/3 for the core pointer and keyboard, > (matching the
// IDs for the Virtual Core Pointer and Keyboard in > XInput 2), but calling
// this function on any slave devices (i.e. > those managed via XInput 1.x),
// will return 0.
func X11DeviceGetID(device *X11DeviceCore) int

// X11DeviceManagerLookup: returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager *X11DeviceManagerCore, deviceID int) *X11DeviceCore

// X11GetDefaultRootXwindow: gets the root window of the default screen (see
// gdk_x11_get_default_screen()).
func X11GetDefaultRootXwindow() xlib.Window

// X11GetDefaultScreen: gets the default GTK+ screen number.
func X11GetDefaultScreen() int

// X11GetDefaultXdisplay: gets the default GTK+ display.
func X11GetDefaultXdisplay() *xlib.Display

// X11GetParentRelativePattern: used with gdk_window_set_background_pattern() to
// inherit background from parent window. Useful for imitating transparency when
// compositing is not available. Otherwise behaves like a transparent pattern.
func X11GetParentRelativePattern() *cairo.Pattern

// X11GetServerTime: routine to get the current X server time stamp.
func X11GetServerTime(window *X11Window) uint32

// X11GetXatomByName: returns the X atom for GDK’s default display corresponding
// to @atom_name. This function caches the result, so if called repeatedly it is
// much faster than XInternAtom(), which is a round trip to the server each
// time.
func X11GetXatomByName(atomName string) xlib.Atom

// X11GetXatomByNameForDisplay: returns the X atom for a Display corresponding
// to @atom_name. This function caches the result, so if called repeatedly it is
// much faster than XInternAtom(), which is a round trip to the server each
// time.
func X11GetXatomByNameForDisplay(display *X11Display, atomName string) xlib.Atom

// X11GetXatomName: returns the name of an X atom for GDK’s default display.
// This function is meant mainly for debugging, so for convenience, unlike
// XAtomName() and gdk_atom_name(), the result doesn’t need to be freed. Also,
// this function will never return nil, even if @xatom is invalid.
func X11GetXatomName(xatom xlib.Atom) string

// X11GetXatomNameForDisplay: returns the name of an X atom for its display.
// This function is meant mainly for debugging, so for convenience, unlike
// XAtomName() and gdk_atom_name(), the result doesn’t need to be freed.
func X11GetXatomNameForDisplay(display *X11Display, xatom xlib.Atom) string

// X11LookupXdisplay: find the Display corresponding to @xdisplay, if any
// exists.
func X11LookupXdisplay(xdisplay *xlib.Display) *X11Display

// X11XatomToAtom: convert from an X atom for the default display to the
// corresponding Atom.
func X11XatomToAtom(xatom xlib.Atom) gdk.Atom

// X11XatomToAtomForDisplay: convert from an X atom for a Display to the
// corresponding Atom.
func X11XatomToAtomForDisplay(display *X11Display, xatom xlib.Atom) gdk.Atom
