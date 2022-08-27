// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// CreateGLContext creates a new GLContext matching the framebuffer format to
// the visual of the Window. The context is disconnected from any particular
// window or surface.
//
// If the creation of the GLContext failed, error will be set.
//
// Before using the returned GLContext, you will need to call
// gdk_gl_context_make_current() or gdk_gl_context_realize().
//
// The function returns the following values:
//
//    - glContext: newly created GLContext, or NULL on error.
//
func (window *Window) CreateGLContext() (GLContexter, error) {
	var _arg0 *C.GdkWindow    // out
	var _cret *C.GdkGLContext // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gdk_window_create_gl_context(_arg0, &_cerr)
	runtime.KeepAlive(window)

	var _glContext GLContexter // out
	var _goerr error           // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(GLContexter)
			return ok
		})
		rv, ok := casted.(GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
		}
		_glContext = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _glContext, _goerr
}

// MarkPaintFromClip: if you call this during a paint (e.g. between
// gdk_window_begin_paint_region() and gdk_window_end_paint() then GDK will mark
// the current clip region of the window as being drawn. This is required when
// mixing GL rendering via gdk_cairo_draw_from_gl() and cairo rendering, as
// otherwise GDK has no way of knowing when something paints over the GL-drawn
// regions.
//
// This is typically called automatically by GTK+ and you don't need to care
// about this.
//
// The function takes the following parameters:
//
//    - cr: #cairo_t.
//
func (window *Window) MarkPaintFromClip(cr *cairo.Context) {
	var _arg0 *C.GdkWindow // out
	var _arg1 *C.cairo_t   // out

	_arg0 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	C.gdk_window_mark_paint_from_clip(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(cr)
}
