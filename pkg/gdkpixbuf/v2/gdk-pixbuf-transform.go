// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeInterpType     = coreglib.Type(C.gdk_interp_type_get_type())
	GTypePixbufRotation = coreglib.Type(C.gdk_pixbuf_rotation_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeInterpType, F: marshalInterpType},
		coreglib.TypeMarshaler{T: GTypePixbufRotation, F: marshalPixbufRotation},
	})
}

// InterpType: interpolation modes for scaling functions.
//
// The GDK_INTERP_NEAREST mode is the fastest scaling method, but has horrible
// quality when scaling down; GDK_INTERP_BILINEAR is the best choice if you
// aren't sure what to choose, it has a good speed/quality balance.
//
// **Note**: Cubic filtering is missing from the list; hyperbolic interpolation
// is just as fast and results in higher quality.
type InterpType C.gint

const (
	// InterpNearest: nearest neighbor sampling; this is the fastest and lowest
	// quality mode. Quality is normally unacceptable when scaling down, but may
	// be OK when scaling up.
	InterpNearest InterpType = iota
	// InterpTiles: this is an accurate simulation of the PostScript image
	// operator without any interpolation enabled. Each pixel is rendered as a
	// tiny parallelogram of solid color, the edges of which are implemented
	// with antialiasing. It resembles nearest neighbor for enlargement,
	// and bilinear for reduction.
	InterpTiles
	// InterpBilinear: best quality/speed balance; use this mode by default.
	// Bilinear interpolation. For enlargement, it is equivalent to
	// point-sampling the ideal bilinear-interpolated image. For reduction,
	// it is equivalent to laying down small tiles and integrating over the
	// coverage area.
	InterpBilinear
	// InterpHyper: this is the slowest and highest quality reconstruction
	// function. It is derived from the hyperbolic filters in Wolberg's "Digital
	// Image Warping", and is formally defined as the hyperbolic-filter sampling
	// the ideal hyperbolic-filter interpolated image (the filter is designed to
	// be idempotent for 1:1 pixel mapping). **Deprecated**: this interpolation
	// filter is deprecated, as in reality it has a lower quality than the
	// GDK_INTERP_BILINEAR filter (Since: 2.38).
	InterpHyper
)

func marshalInterpType(p uintptr) (interface{}, error) {
	return InterpType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InterpType.
func (i InterpType) String() string {
	switch i {
	case InterpNearest:
		return "Nearest"
	case InterpTiles:
		return "Tiles"
	case InterpBilinear:
		return "Bilinear"
	case InterpHyper:
		return "Hyper"
	default:
		return fmt.Sprintf("InterpType(%d)", i)
	}
}

// PixbufRotation: possible rotations which can be passed to
// gdk_pixbuf_rotate_simple().
//
// To make them easier to use, their numerical values are the actual degrees.
type PixbufRotation C.gint

const (
	// PixbufRotateNone: no rotation.
	PixbufRotateNone PixbufRotation = 0
	// PixbufRotateCounterclockwise: rotate by 90 degrees.
	PixbufRotateCounterclockwise PixbufRotation = 90
	// PixbufRotateUpsidedown: rotate by 180 degrees.
	PixbufRotateUpsidedown PixbufRotation = 180
	// PixbufRotateClockwise: rotate by 270 degrees.
	PixbufRotateClockwise PixbufRotation = 270
)

func marshalPixbufRotation(p uintptr) (interface{}, error) {
	return PixbufRotation(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PixbufRotation.
func (p PixbufRotation) String() string {
	switch p {
	case PixbufRotateNone:
		return "None"
	case PixbufRotateCounterclockwise:
		return "Counterclockwise"
	case PixbufRotateUpsidedown:
		return "Upsidedown"
	case PixbufRotateClockwise:
		return "Clockwise"
	default:
		return fmt.Sprintf("PixbufRotation(%d)", p)
	}
}
