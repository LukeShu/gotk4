// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// CairoDrawFromGL: main way to draw GL content in GTK.
//
// It takes a render buffer ID (source_type == RENDERBUFFER) or a texture id
// (source_type == TEXTURE) and draws it onto cr with an OVER operation,
// respecting the current clip. The top left corner of the rectangle specified
// by x, y, width and height will be drawn at the current (0,0) position of the
// cairo_t.
//
// This will work for *all* cairo_t, as long as surface is realized, but the
// fallback implementation that reads back the pixels from the buffer may be
// used in the general case. In the case of direct drawing to a surface with no
// special effects applied to cr it will however use a more efficient approach.
//
// For RENDERBUFFER the code will always fall back to software for buffers with
// alpha components, so make sure you use TEXTURE if using alpha.
//
// Calling this may change the current GL context.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - surface we're rendering for (not necessarily into).
//    - source: GL ID of the source buffer.
//    - sourceType: type of the source.
//    - bufferScale: scale-factor that the source buffer is allocated for.
//    - x: source x position in source to start copying from in GL coordinates.
//    - y: source y position in source to start copying from in GL coordinates.
//    - width of the region to draw.
//    - height of the region to draw.
//
func CairoDrawFromGL(cr *cairo.Context, surface Surfacer, source, sourceType, bufferScale, x, y, width, height int) {
	var _arg1 *C.cairo_t    // out
	var _arg2 *C.GdkSurface // out
	var _arg3 C.int         // out
	var _arg4 C.int         // out
	var _arg5 C.int         // out
	var _arg6 C.int         // out
	var _arg7 C.int         // out
	var _arg8 C.int         // out
	var _arg9 C.int         // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkSurface)(unsafe.Pointer(externglib.InternObject(surface).Native()))
	_arg3 = C.int(source)
	_arg4 = C.int(sourceType)
	_arg5 = C.int(bufferScale)
	_arg6 = C.int(x)
	_arg7 = C.int(y)
	_arg8 = C.int(width)
	_arg9 = C.int(height)

	C.gdk_cairo_draw_from_gl(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(source)
	runtime.KeepAlive(sourceType)
	runtime.KeepAlive(bufferScale)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// CairoRectangle adds the given rectangle to the current path of cr.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - rectangle: Rectangle.
//
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	var _arg1 *C.cairo_t      // out
	var _arg2 *C.GdkRectangle // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	C.gdk_cairo_rectangle(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(rectangle)
}

// CairoRegion adds the given region to the current path of cr.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - region: #cairo_region_t.
//
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.cairo_region_t // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_cairo_region(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(region)
}

// CairoRegionCreateFromSurface creates region that covers the area where the
// given surface is more than 50% opaque.
//
// This function takes into account device offsets that might be set with
// cairo_surface_set_device_offset().
//
// The function takes the following parameters:
//
//    - surface: cairo surface.
//
// The function returns the following values:
//
//    - region: cairo_region_t; must be freed with cairo_region_destroy().
//
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	var _arg1 *C.cairo_surface_t // out
	var _cret *C.cairo_region_t  // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_cairo_region_create_from_surface(_arg1)
	runtime.KeepAlive(surface)

	var _region *cairo.Region // out

	{
		_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
		_region = (*cairo.Region)(unsafe.Pointer(_pp))
	}
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.cairo_region_destroy((*C.cairo_region_t)(unsafe.Pointer(v.Native())))
	})

	return _region
}

// CairoSetSourcePixbuf sets the given pixbuf as the source pattern for cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of pixbuf is pixbuf_x, pixbuf_y.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - pixbuf: Pixbuf.
//    - pixbufX: x coordinate of location to place upper left corner of pixbuf.
//    - pixbufY: y coordinate of location to place upper left corner of pixbuf.
//
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, pixbufX, pixbufY float64) {
	var _arg1 *C.cairo_t   // out
	var _arg2 *C.GdkPixbuf // out
	var _arg3 C.double     // out
	var _arg4 C.double     // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(externglib.InternObject(pixbuf).Native()))
	_arg3 = C.double(pixbufX)
	_arg4 = C.double(pixbufY)

	C.gdk_cairo_set_source_pixbuf(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(pixbufX)
	runtime.KeepAlive(pixbufY)
}

// CairoSetSourceRGBA sets the specified RGBA as the source color of cr.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - rgba: RGBA.
//
func CairoSetSourceRGBA(cr *cairo.Context, rgba *RGBA) {
	var _arg1 *C.cairo_t // out
	var _arg2 *C.GdkRGBA // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	C.gdk_cairo_set_source_rgba(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(rgba)
}
