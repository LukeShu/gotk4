// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeDrawingContext returns the GType for the type DrawingContext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDrawingContext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "DrawingContext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDrawingContext)
	return gtype
}

// DrawingContext is an object that represents the current drawing state of a
// Window.
//
// It's possible to use a DrawingContext to draw on a Window via rendering API
// like Cairo or OpenGL.
//
// A DrawingContext can only be created by calling gdk_window_begin_draw_frame()
// and will be valid until a call to gdk_window_end_draw_frame().
//
// DrawingContext is available since GDK 3.22.
type DrawingContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DrawingContext)(nil)
)

func wrapDrawingContext(obj *coreglib.Object) *DrawingContext {
	return &DrawingContext{
		Object: obj,
	}
}

func marshalDrawingContext(p uintptr) (interface{}, error) {
	return wrapDrawingContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CairoContext retrieves a Cairo context to be used to draw on the Window that
// created the DrawingContext.
//
// The returned context is guaranteed to be valid as long as the DrawingContext
// is valid, that is between a call to gdk_window_begin_draw_frame() and
// gdk_window_end_draw_frame().
//
// The function returns the following values:
//
//    - ret: cairo context to be used to draw the contents of the Window. The
//      context is owned by the DrawingContext and should not be destroyed.
//
func (context *DrawingContext) CairoContext() *cairo.Context {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawingContext").InvokeMethod("get_cairo_context", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ret *cairo.Context // out

	_ret = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	C.cairo_reference(_cret)
	runtime.SetFinalizer(_ret, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})

	return _ret
}

// Clip retrieves a copy of the clip region used when creating the context.
//
// The function returns the following values:
//
//    - region (optional): cairo region.
//
func (context *DrawingContext) Clip() *cairo.Region {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawingContext").InvokeMethod("get_clip", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _region *cairo.Region // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
			_region = (*cairo.Region)(unsafe.Pointer(_pp))
		}
		runtime.SetFinalizer(_region, func(v *cairo.Region) {
			C.cairo_region_destroy((*C.void)(unsafe.Pointer(v.Native())))
		})
	}

	return _region
}

// Window retrieves the window that created the drawing context.
//
// The function returns the following values:
//
//    - window: Window.
//
func (context *DrawingContext) Window() Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawingContext").InvokeMethod("get_window", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _window Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// IsValid checks whether the given DrawingContext is valid.
//
// The function returns the following values:
//
//    - ok: TRUE if the context is valid.
//
func (context *DrawingContext) IsValid() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawingContext").InvokeMethod("is_valid", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
