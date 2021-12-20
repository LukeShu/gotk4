// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_cairo_renderer_get_type()), F: marshalCairoRendererer},
	})
}

// CairoRenderer: GSK renderer that is using cairo.
//
// Since it is using cairo, this renderer cannot support 3D transformations.
type CairoRenderer struct {
	_ [0]func() // equal guard
	Renderer
}

var (
	_ Rendererer = (*CairoRenderer)(nil)
)

func wrapCairoRenderer(obj *externglib.Object) *CairoRenderer {
	return &CairoRenderer{
		Renderer: Renderer{
			Object: obj,
		},
	}
}

func marshalCairoRendererer(p uintptr) (interface{}, error) {
	return wrapCairoRenderer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCairoRenderer creates a new Cairo renderer.
//
// The Cairo renderer is the fallback renderer drawing in ways similar to how
// GTK 3 drew its content. Its primary use is as comparison tool.
//
// The Cairo renderer is incomplete. It cannot render 3D transformed content and
// will instead render an error marker. Its usage should be avoided.
//
// The function returns the following values:
//
//    - cairoRenderer: new Cairo renderer.
//
func NewCairoRenderer() *CairoRenderer {
	var _cret *C.GskRenderer // in

	_cret = C.gsk_cairo_renderer_new()

	var _cairoRenderer *CairoRenderer // out

	_cairoRenderer = wrapCairoRenderer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cairoRenderer
}
