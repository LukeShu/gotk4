// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_visual_get_type()), F: marshalX11Visual},
	})
}

type X11Visual interface {
	gextras.Objector

	privateX11VisualClass()
}

// X11VisualClass implements the X11Visual interface.
type X11VisualClass struct {
	gdk.VisualClass
}

var _ X11Visual = (*X11VisualClass)(nil)

func wrapX11Visual(obj *externglib.Object) X11Visual {
	return &X11VisualClass{
		VisualClass: gdk.VisualClass{
			Object: obj,
		},
	}
}

func marshalX11Visual(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Visual(obj), nil
}

func (*X11VisualClass) privateX11VisualClass() {}
