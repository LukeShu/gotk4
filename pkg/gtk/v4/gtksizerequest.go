// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// RequestedSize represents a request of a screen object in a given orientation.
// These are primarily used in container implementations when allocating a
// natural size for children calling. See gtk_distribute_natural_allocation().
//
// An instance of this type is always passed by reference.
type RequestedSize struct {
	*requestedSize
}

// requestedSize is the struct that's finalized.
type requestedSize struct {
	native unsafe.Pointer
}

// Data: client pointer.
func (r *RequestedSize) Data() unsafe.Pointer {
	offset := girepository.MustFind("Gtk", "RequestedSize").StructFieldOffset("data")
	valptr := unsafe.Add(unsafe.Pointer(r), offset)
	var v unsafe.Pointer // out
	v = (unsafe.Pointer)(unsafe.Pointer(valptr))
	return v
}

// MinimumSize: minimum size needed for allocation in a given orientation.
func (r *RequestedSize) MinimumSize() int32 {
	offset := girepository.MustFind("Gtk", "RequestedSize").StructFieldOffset("minimum_size")
	valptr := unsafe.Add(unsafe.Pointer(r), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}

// NaturalSize: natural size for allocation in a given orientation.
func (r *RequestedSize) NaturalSize() int32 {
	offset := girepository.MustFind("Gtk", "RequestedSize").StructFieldOffset("natural_size")
	valptr := unsafe.Add(unsafe.Pointer(r), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}
