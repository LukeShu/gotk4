// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// DistributeNaturalAllocation distributes extra_space to child sizes by
// bringing smaller children up to natural size first.
//
// The remaining space will be added to the minimum_size member of the
// GtkRequestedSize struct. If all sizes reach their natural size then the
// remaining space is returned.
func DistributeNaturalAllocation(extraSpace int, nRequestedSizes uint, sizes *RequestedSize) int {
	var _arg1 C.gint              // out
	var _arg2 C.guint             // out
	var _arg3 *C.GtkRequestedSize // out
	var _cret C.gint              // in

	_arg1 = C.gint(extraSpace)
	_arg2 = C.guint(nRequestedSizes)
	_arg3 = (*C.GtkRequestedSize)(gextras.StructNative(unsafe.Pointer(sizes)))

	_cret = C.gtk_distribute_natural_allocation(_arg1, _arg2, _arg3)

	runtime.KeepAlive(extraSpace)
	runtime.KeepAlive(nRequestedSizes)
	runtime.KeepAlive(sizes)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RequestedSize represents a request of a screen object in a given orientation.
// These are primarily used in container implementations when allocating a
// natural size for children calling. See gtk_distribute_natural_allocation().
type RequestedSize struct {
	nocopy gextras.NoCopy
	native *C.GtkRequestedSize
}

// Data: client pointer
func (r *RequestedSize) Data() cgo.Handle {
	var v cgo.Handle // out
	v = (cgo.Handle)(unsafe.Pointer(r.native.data))
	return v
}

// MinimumSize: minimum size needed for allocation in a given orientation
func (r *RequestedSize) MinimumSize() int {
	var v int // out
	v = int(r.native.minimum_size)
	return v
}

// NaturalSize: natural size for allocation in a given orientation
func (r *RequestedSize) NaturalSize() int {
	var v int // out
	v = int(r.native.natural_size)
	return v
}
