// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

// glib.Type values for pango-direction.go.
var GTypeDirection = coreglib.Type(C.pango_direction_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDirection, F: marshalDirection},
	})
}

// Direction: PangoDirection represents a direction in the Unicode bidirectional
// algorithm.
//
// Not every value in this enumeration makes sense for every usage of
// PangoDirection; for example, the return value of unichar_direction and
// find_base_dir cannot be PANGO_DIRECTION_WEAK_LTR or PANGO_DIRECTION_WEAK_RTL,
// since every character is either neutral or has a strong direction; on the
// other hand PANGO_DIRECTION_NEUTRAL doesn't make sense to pass to
// itemize_with_base_dir.
//
// The PANGO_DIRECTION_TTB_LTR, PANGO_DIRECTION_TTB_RTL values come from an
// earlier interpretation of this enumeration as the writing direction of a
// block of text and are no longer used; See PangoGravity for how vertical text
// is handled in Pango.
//
// If you are interested in text direction, you should really use fribidi
// directly. PangoDirection is only retained because it is used in some public
// apis.
type Direction C.gint

const (
	// DirectionLTR: strong left-to-right direction.
	DirectionLTR Direction = iota
	// DirectionRTL: strong right-to-left direction.
	DirectionRTL
	// DirectionTtbLTR: deprecated value; treated the same as
	// PANGO_DIRECTION_RTL.
	DirectionTtbLTR
	// DirectionTtbRTL: deprecated value; treated the same as
	// PANGO_DIRECTION_LTR.
	DirectionTtbRTL
	// DirectionWeakLTR: weak left-to-right direction.
	DirectionWeakLTR
	// DirectionWeakRTL: weak right-to-left direction.
	DirectionWeakRTL
	// DirectionNeutral: no direction specified.
	DirectionNeutral
)

func marshalDirection(p uintptr) (interface{}, error) {
	return Direction(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Direction.
func (d Direction) String() string {
	switch d {
	case DirectionLTR:
		return "LTR"
	case DirectionRTL:
		return "RTL"
	case DirectionTtbLTR:
		return "TtbLTR"
	case DirectionTtbRTL:
		return "TtbRTL"
	case DirectionWeakLTR:
		return "WeakLTR"
	case DirectionWeakRTL:
		return "WeakRTL"
	case DirectionNeutral:
		return "Neutral"
	default:
		return fmt.Sprintf("Direction(%d)", d)
	}
}
