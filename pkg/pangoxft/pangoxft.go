// Code generated by girgen. DO NOT EDIT.

package pangoxft

import (
	"github.com/diamondburned/gotk4/pango"
	"github.com/diamondburned/gotk4/xlib"
	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoxft
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pangoxft.h>
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{

		// Objects/Classes
	})
}

// GetContext: retrieves a Context appropriate for rendering with Xft fonts on
// the given screen of the given display.
func GetContext(display *xlib.Display, screen int) *pango.Context

// GetFontMap: returns the XftFontMap for the given display and screen. The
// fontmap is owned by Pango and will be valid until the display is closed.
func GetFontMap(display *xlib.Display, screen int) *pango.FontMap
