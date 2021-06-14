// Code generated by girgen. DO NOT EDIT.

package pangoxft

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 pango pangoxft
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangoxft.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_xft_renderer_get_type()), F: marshalRenderer},
	})
}

// Renderer is a subclass of Renderer used for rendering with Pango's Xft
// backend. It can be used directly, or it can be further subclassed to modify
// exactly how drawing of individual elements occurs.
type Renderer interface {
	Renderer
}

// renderer implements the Renderer class.
type renderer struct {
	Renderer
}

var _ Renderer = (*renderer)(nil)

// WrapRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapRenderer(obj *externglib.Object) Renderer {
	return renderer{
		Renderer: WrapRenderer(obj),
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRenderer(obj), nil
}
