// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_draw_context_get_type()), F: marshalDrawContext},
	})
}

// DrawContext: base class for objects implementing different rendering methods.
//
// `GdkDrawContext` is the base object used by contexts implementing different
// rendering methods, such as [class@Gdk.CairoContext] or [class@Gdk.GLContext].
// It provides shared functionality between those contexts.
//
// You will always interact with one of those subclasses.
//
// A `GdkDrawContext` is always associated with a single toplevel surface.
type DrawContext interface {
	gextras.Objector

	// EndFrame ends a drawing operation started with
	// gdk_draw_context_begin_frame().
	//
	// This makes the drawing available on screen. See
	// [method@Gdk.DrawContext.begin_frame] for more details about drawing.
	//
	// When using a [class@Gdk.GLContext], this function may call `glFlush()`
	// implicitly before returning; it is not recommended to call `glFlush()`
	// explicitly before calling this function.
	EndFrame()
	// Display retrieves the `GdkDisplay` the @context is created for
	Display() Display
	// Surface retrieves the surface that @context is bound to.
	Surface() Surface
	// IsInFrame returns true if @context is in the process of drawing to its
	// surface.
	//
	// This is the case between calls to [method@Gdk.DrawContext.begin_frame]
	// and [method@Gdk.DrawContext.end_frame]. In this situation, drawing
	// commands may be effecting the contents of the @context's surface.
	IsInFrame() bool
}

// drawContext implements the DrawContext class.
type drawContext struct {
	gextras.Objector
}

var _ DrawContext = (*drawContext)(nil)

// WrapDrawContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrawContext(obj *externglib.Object) DrawContext {
	return drawContext{
		Objector: obj,
	}
}

func marshalDrawContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrawContext(obj), nil
}

// EndFrame ends a drawing operation started with
// gdk_draw_context_begin_frame().
//
// This makes the drawing available on screen. See
// [method@Gdk.DrawContext.begin_frame] for more details about drawing.
//
// When using a [class@Gdk.GLContext], this function may call `glFlush()`
// implicitly before returning; it is not recommended to call `glFlush()`
// explicitly before calling this function.
func (c drawContext) EndFrame() {
	var _arg0 *C.GdkDrawContext // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	C.gdk_draw_context_end_frame(_arg0)
}

// Display retrieves the `GdkDisplay` the @context is created for
func (c drawContext) Display() Display {
	var _arg0 *C.GdkDrawContext // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	var _cret *C.GdkDisplay // in

	_cret = C.gdk_draw_context_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Display)

	return _display
}

// Surface retrieves the surface that @context is bound to.
func (c drawContext) Surface() Surface {
	var _arg0 *C.GdkDrawContext // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	var _cret *C.GdkSurface // in

	_cret = C.gdk_draw_context_get_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Surface)

	return _surface
}

// IsInFrame returns true if @context is in the process of drawing to its
// surface.
//
// This is the case between calls to [method@Gdk.DrawContext.begin_frame]
// and [method@Gdk.DrawContext.end_frame]. In this situation, drawing
// commands may be effecting the contents of the @context's surface.
func (c drawContext) IsInFrame() bool {
	var _arg0 *C.GdkDrawContext // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_draw_context_is_in_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
