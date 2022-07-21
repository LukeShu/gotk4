// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeBorder returns the GType for the type Border.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBorder() coreglib.Type {
	gtype := coreglib.Type(C.gtk_border_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalBorder)
	return gtype
}

// Border: struct that specifies a border around a rectangular area that can be
// of different width on each side.
//
// An instance of this type is always passed by reference.
type Border struct {
	*border
}

// border is the struct that's finalized.
type border struct {
	native *C.GtkBorder
}

func marshalBorder(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Border{&border{(*C.GtkBorder)(b)}}, nil
}

// NewBorder constructs a struct Border.
func NewBorder() *Border {
	var _cret *C.GtkBorder // in

	_cret = C.gtk_border_new()

	var _border *Border // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_border)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_border_free((*C.GtkBorder)(intern.C))
		},
	)

	return _border
}

// Left: width of the left border.
func (b *Border) Left() int16 {
	valptr := &b.native.left
	var v int16 // out
	v = int16(*valptr)
	return v
}

// Right: width of the right border.
func (b *Border) Right() int16 {
	valptr := &b.native.right
	var v int16 // out
	v = int16(*valptr)
	return v
}

// Top: width of the top border.
func (b *Border) Top() int16 {
	valptr := &b.native.top
	var v int16 // out
	v = int16(*valptr)
	return v
}

// Bottom: width of the bottom border.
func (b *Border) Bottom() int16 {
	valptr := &b.native.bottom
	var v int16 // out
	v = int16(*valptr)
	return v
}

// Left: width of the left border.
func (b *Border) SetLeft(left int16) {
	valptr := &b.native.left
	*valptr = C.gint16(left)
}

// Right: width of the right border.
func (b *Border) SetRight(right int16) {
	valptr := &b.native.right
	*valptr = C.gint16(right)
}

// Top: width of the top border.
func (b *Border) SetTop(top int16) {
	valptr := &b.native.top
	*valptr = C.gint16(top)
}

// Bottom: width of the bottom border.
func (b *Border) SetBottom(bottom int16) {
	valptr := &b.native.bottom
	*valptr = C.gint16(bottom)
}

// Copy copies a Border-struct.
//
// The function returns the following values:
//
//    - border: copy of border_.
//
func (border_ *Border) Copy() *Border {
	var _arg0 *C.GtkBorder // out
	var _cret *C.GtkBorder // in

	_arg0 = (*C.GtkBorder)(gextras.StructNative(unsafe.Pointer(border_)))

	_cret = C.gtk_border_copy(_arg0)
	runtime.KeepAlive(border_)

	var _border *Border // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_border)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_border_free((*C.GtkBorder)(intern.C))
		},
	)

	return _border
}
