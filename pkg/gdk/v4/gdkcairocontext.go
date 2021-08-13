// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_cairo_context_get_type()), F: marshalCairoContexter},
	})
}

// CairoContext: GdkCairoContext is an object representing the platform-specific
// draw context.
//
// GdkCairoContexts are created for a surface using
// gdk.Surface.CreateCairoContext(), and the context can then be used to draw on
// that surface.
type CairoContext struct {
	DrawContext
}

// CairoContexter describes CairoContext's abstract methods.
type CairoContexter interface {
	externglib.Objector

	// CairoCreate retrieves a Cairo context to be used to draw on the
	// GdkSurface of context.
	CairoCreate() *cairo.Context
}

var _ CairoContexter = (*CairoContext)(nil)

func wrapCairoContext(obj *externglib.Object) *CairoContext {
	return &CairoContext{
		DrawContext: DrawContext{
			Object: obj,
		},
	}
}

func marshalCairoContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCairoContext(obj), nil
}

// CairoCreate retrieves a Cairo context to be used to draw on the GdkSurface of
// context.
//
// A call to gdk.DrawContext.BeginFrame() with this context must have been done
// or this function will return NULL.
//
// The returned context is guaranteed to be valid until
// gdk.DrawContext.EndFrame() is called.
func (self *CairoContext) CairoCreate() *cairo.Context {
	var _arg0 *C.GdkCairoContext // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GdkCairoContext)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_cairo_context_cairo_create(_arg0)

	runtime.KeepAlive(self)

	var _context *cairo.Context // out

	if _cret != nil {
		_context = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(_context, func(v *cairo.Context) {
			C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
		})
	}

	return _context
}
