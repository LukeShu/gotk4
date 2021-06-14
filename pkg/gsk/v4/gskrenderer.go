// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_renderer_get_type()), F: marshalRenderer},
	})
}

// Renderer: `GskRenderer` is a class that renders a scene graph defined via a
// tree of [class@Gsk.RenderNode] instances.
//
// Typically you will use a `GskRenderer` instance to repeatedly call
// [method@Gsk.Renderer.render] to update the contents of its associated
// [class@Gdk.Surface].
//
// It is necessary to realize a `GskRenderer` instance using
// [method@Gsk.Renderer.realize] before calling [method@Gsk.Renderer.render], in
// order to create the appropriate windowing system resources needed to render
// the scene.
type Renderer interface {
	gextras.Objector

	// IsRealized checks whether the @renderer is realized or not.
	IsRealized() bool
	// Unrealize releases all the resources created by gsk_renderer_realize().
	Unrealize()
}

// renderer implements the Renderer class.
type renderer struct {
	gextras.Objector
}

var _ Renderer = (*renderer)(nil)

// WrapRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapRenderer(obj *externglib.Object) Renderer {
	return renderer{
		Objector: obj,
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRenderer(obj), nil
}

// IsRealized checks whether the @renderer is realized or not.
func (r renderer) IsRealized() bool {
	var _arg0 *C.GskRenderer // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	var _cret C.gboolean // in

	_cret = C.gsk_renderer_is_realized(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Unrealize releases all the resources created by gsk_renderer_realize().
func (r renderer) Unrealize() {
	var _arg0 *C.GskRenderer // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(r.Native()))

	C.gsk_renderer_unrealize(_arg0)
}
