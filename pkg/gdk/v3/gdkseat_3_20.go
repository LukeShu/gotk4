// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSeatCapabilities = coreglib.Type(C.gdk_seat_capabilities_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSeatCapabilities, F: marshalSeatCapabilities},
	})
}

// SeatCapabilities flags describing the seat capabilities.
type SeatCapabilities C.guint

const (
	// SeatCapabilityNone: no input capabilities.
	SeatCapabilityNone SeatCapabilities = 0b0
	// SeatCapabilityPointer: seat has a pointer (e.g. mouse).
	SeatCapabilityPointer SeatCapabilities = 0b1
	// SeatCapabilityTouch: seat has touchscreen(s) attached.
	SeatCapabilityTouch SeatCapabilities = 0b10
	// SeatCapabilityTabletStylus: seat has drawing tablet(s) attached.
	SeatCapabilityTabletStylus SeatCapabilities = 0b100
	// SeatCapabilityKeyboard: seat has keyboard(s) attached.
	SeatCapabilityKeyboard SeatCapabilities = 0b1000
	// SeatCapabilityAllPointing: union of all pointing capabilities.
	SeatCapabilityAllPointing SeatCapabilities = 0b111
	// SeatCapabilityAll: union of all capabilities.
	SeatCapabilityAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SeatCapabilities.
func (s SeatCapabilities) String() string {
	if s == 0 {
		return "SeatCapabilities(0)"
	}

	var builder strings.Builder
	builder.Grow(154)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SeatCapabilityNone:
			builder.WriteString("None|")
		case SeatCapabilityPointer:
			builder.WriteString("Pointer|")
		case SeatCapabilityTouch:
			builder.WriteString("Touch|")
		case SeatCapabilityTabletStylus:
			builder.WriteString("TabletStylus|")
		case SeatCapabilityKeyboard:
			builder.WriteString("Keyboard|")
		case SeatCapabilityAllPointing:
			builder.WriteString("AllPointing|")
		case SeatCapabilityAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("SeatCapabilities(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SeatCapabilities) Has(other SeatCapabilities) bool {
	return (s & other) == other
}

// SeatGrabPrepareFunc: type of the callback used to set up window so it can be
// grabbed. A typical action would be ensuring the window is visible, although
// there's room for other initialization actions.
type SeatGrabPrepareFunc func(seat Seater, window Windower)
