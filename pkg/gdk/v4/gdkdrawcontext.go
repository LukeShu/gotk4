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

// GTypeDrawContext returns the GType for the type DrawContext.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDrawContext() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "DrawContext").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDrawContext)
	return gtype
}

// DrawContext: base class for objects implementing different rendering methods.
//
// GdkDrawContext is the base object used by contexts implementing different
// rendering methods, such as gdk.CairoContext or gdk.GLContext. It provides
// shared functionality between those contexts.
//
// You will always interact with one of those subclasses.
//
// A GdkDrawContext is always associated with a single toplevel surface.
type DrawContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DrawContext)(nil)
)

// DrawContexter describes types inherited from class DrawContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DrawContexter interface {
	coreglib.Objector
	baseDrawContext() *DrawContext
}

var _ DrawContexter = (*DrawContext)(nil)

func wrapDrawContext(obj *coreglib.Object) *DrawContext {
	return &DrawContext{
		Object: obj,
	}
}

func marshalDrawContext(p uintptr) (interface{}, error) {
	return wrapDrawContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (context *DrawContext) baseDrawContext() *DrawContext {
	return context
}

// BaseDrawContext returns the underlying base object.
func BaseDrawContext(obj DrawContexter) *DrawContext {
	return obj.baseDrawContext()
}

// BeginFrame indicates that you are beginning the process of redrawing region
// on the context's surface.
//
// Calling this function begins a drawing operation using context on the surface
// that context was created from. The actual requirements and guarantees for the
// drawing operation vary for different implementations of drawing, so a
// gdk.CairoContext and a gdk.GLContext need to be treated differently.
//
// A call to this function is a requirement for drawing and must be followed by
// a call to gdk.DrawContext.EndFrame(), which will complete the drawing
// operation and ensure the contents become visible on screen.
//
// Note that the region passed to this function is the minimum region that needs
// to be drawn and depending on implementation, windowing system and hardware in
// use, it might be necessary to draw a larger region. Drawing implementation
// must use [methodGdk.DrawContext.get_frame_region() to query the region that
// must be drawn.
//
// When using GTK, the widget system automatically places calls to
// gdk_draw_context_begin_frame() and gdk_draw_context_end_frame() via the use
// of gsk.Renderers, so application code does not need to call these functions
// explicitly.
//
// The function takes the following parameters:
//
//    - region: minimum region that should be drawn.
//
func (context *DrawContext) BeginFrame(region *cairo.Region) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(region.Native()))

	girepository.MustFind("Gdk", "DrawContext").InvokeMethod("begin_frame", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(region)
}

// EndFrame ends a drawing operation started with
// gdk_draw_context_begin_frame().
//
// This makes the drawing available on screen. See gdk.DrawContext.BeginFrame()
// for more details about drawing.
//
// When using a gdk.GLContext, this function may call glFlush() implicitly
// before returning; it is not recommended to call glFlush() explicitly before
// calling this function.
func (context *DrawContext) EndFrame() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gdk", "DrawContext").InvokeMethod("end_frame", _args[:], nil)

	runtime.KeepAlive(context)
}

// Display retrieves the GdkDisplay the context is created for.
//
// The function returns the following values:
//
//    - display (optional): GdkDisplay or NULL.
//
func (context *DrawContext) Display() *Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawContext").InvokeMethod("get_display", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _display *Display // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// FrameRegion retrieves the region that is currently being repainted.
//
// After a call to gdk.DrawContext.BeginFrame() this function will return a
// union of the region passed to that function and the area of the surface that
// the context determined needs to be repainted.
//
// If context is not in between calls to gdk.DrawContext.BeginFrame() and
// gdk.DrawContext.EndFrame(), NULL will be returned.
//
// The function returns the following values:
//
//    - region (optional): cairo region or NULL if not drawing a frame.
//
func (context *DrawContext) FrameRegion() *cairo.Region {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawContext").InvokeMethod("get_frame_region", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _region *cairo.Region // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
			_region = (*cairo.Region)(unsafe.Pointer(_pp))
		}
		C.cairo_region_reference(_cret)
		runtime.SetFinalizer(_region, func(v *cairo.Region) {
			C.cairo_region_destroy((*C.void)(unsafe.Pointer(v.Native())))
		})
	}

	return _region
}

// Surface retrieves the surface that context is bound to.
//
// The function returns the following values:
//
//    - surface (optional) or NULL.
//
func (context *DrawContext) Surface() Surfacer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawContext").InvokeMethod("get_surface", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _surface Surfacer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Surfacer)
				return ok
			})
			rv, ok := casted.(Surfacer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
			}
			_surface = rv
		}
	}

	return _surface
}

// IsInFrame returns TRUE if context is in the process of drawing to its
// surface.
//
// This is the case between calls to gdk.DrawContext.BeginFrame() and
// gdk.DrawContext.EndFrame(). In this situation, drawing commands may be
// effecting the contents of the context's surface.
//
// The function returns the following values:
//
//    - ok: TRUE if the context is between gdk.DrawContext.BeginFrame() and
//      gdk.DrawContext.EndFrame() calls.
//
func (context *DrawContext) IsInFrame() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gdk", "DrawContext").InvokeMethod("is_in_frame", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
