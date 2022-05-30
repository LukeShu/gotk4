// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkwaylandseat.go.
var GTypeWaylandSeat = coreglib.Type(C.gdk_wayland_seat_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeWaylandSeat, F: marshalWaylandSeat},
	})
}

// WaylandSeatOverrider contains methods that are overridable.
type WaylandSeatOverrider interface {
}

// WaylandSeat: wayland implementation of GdkSeat.
//
// Beyond the regular gdk.Seat API, the Wayland implementation provides access
// to the Wayland wl_seat object with gdkwayland.WaylandSeat.GetWlSeat().
type WaylandSeat struct {
	_ [0]func() // equal guard
	gdk.Seat
}

var (
	_ gdk.Seater = (*WaylandSeat)(nil)
)

func classInitWaylandSeater(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWaylandSeat(obj *coreglib.Object) *WaylandSeat {
	return &WaylandSeat{
		Seat: gdk.Seat{
			Object: obj,
		},
	}
}

func marshalWaylandSeat(p uintptr) (interface{}, error) {
	return wrapWaylandSeat(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
