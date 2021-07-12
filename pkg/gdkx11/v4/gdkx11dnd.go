// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_drag_get_type()), F: marshalX11Drager},
	})
}

// X11Drager describes X11Drag's methods.
type X11Drager interface {
	privateX11Drag()
}

type X11Drag struct {
	gdk.Drag
}

var (
	_ X11Drager       = (*X11Drag)(nil)
	_ gextras.Nativer = (*X11Drag)(nil)
)

func wrapX11Drag(obj *externglib.Object) X11Drager {
	return &X11Drag{
		Drag: gdk.Drag{
			Object: obj,
		},
	}
}

func marshalX11Drager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Drag(obj), nil
}

func (*X11Drag) privateX11Drag() {}
