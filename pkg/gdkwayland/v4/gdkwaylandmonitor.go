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

// GTypeWaylandMonitor returns the GType for the type WaylandMonitor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeWaylandMonitor() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("GdkWayland", "WaylandMonitor").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalWaylandMonitor)
	return gtype
}

// WaylandMonitor: wayland implementation of GdkMonitor.
//
// Beyond the gdk.Monitor API, the Wayland implementation offers access to the
// Wayland wl_output object with gdkwayland.WaylandMonitor.GetWlOutput().
type WaylandMonitor struct {
	_ [0]func() // equal guard
	gdk.Monitor
}

var (
	_ coreglib.Objector = (*WaylandMonitor)(nil)
)

func wrapWaylandMonitor(obj *coreglib.Object) *WaylandMonitor {
	return &WaylandMonitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalWaylandMonitor(p uintptr) (interface{}, error) {
	return wrapWaylandMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
