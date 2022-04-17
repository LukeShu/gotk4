// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <pango/pango.h>
// extern void _gotk4_pango1_RendererClass_begin(PangoRenderer*);
// extern void _gotk4_pango1_RendererClass_draw_error_underline(PangoRenderer*, int, int, int, int);
// extern void _gotk4_pango1_RendererClass_draw_glyph(PangoRenderer*, PangoFont*, PangoGlyph, double, double);
// extern void _gotk4_pango1_RendererClass_draw_glyph_item(PangoRenderer*, char*, PangoGlyphItem*, int, int);
// extern void _gotk4_pango1_RendererClass_draw_glyphs(PangoRenderer*, PangoFont*, PangoGlyphString*, int, int);
// extern void _gotk4_pango1_RendererClass_draw_rectangle(PangoRenderer*, PangoRenderPart, int, int, int, int);
// extern void _gotk4_pango1_RendererClass_draw_shape(PangoRenderer*, PangoAttrShape*, int, int);
// extern void _gotk4_pango1_RendererClass_draw_trapezoid(PangoRenderer*, PangoRenderPart, double, double, double, double, double, double);
// extern void _gotk4_pango1_RendererClass_end(PangoRenderer*);
// extern void _gotk4_pango1_RendererClass_part_changed(PangoRenderer*, PangoRenderPart);
// extern void _gotk4_pango1_RendererClass_prepare_run(PangoRenderer*, PangoLayoutRun*);
import "C"

// glib.Type values for pango-renderer.go.
var (
	GTypeRenderPart = externglib.Type(C.pango_render_part_get_type())
	GTypeRenderer   = externglib.Type(C.pango_renderer_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeRenderPart, F: marshalRenderPart},
		{T: GTypeRenderer, F: marshalRenderer},
	})
}

// RenderPart defines different items to render for such purposes as setting
// colors.
type RenderPart C.gint

const (
	// RenderPartForeground: text itself.
	RenderPartForeground RenderPart = iota
	// RenderPartBackground: area behind the text.
	RenderPartBackground
	// RenderPartUnderline: underlines.
	RenderPartUnderline
	// RenderPartStrikethrough: strikethrough lines.
	RenderPartStrikethrough
	// RenderPartOverline: overlines.
	RenderPartOverline
)

func marshalRenderPart(p uintptr) (interface{}, error) {
	return RenderPart(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RenderPart.
func (r RenderPart) String() string {
	switch r {
	case RenderPartForeground:
		return "Foreground"
	case RenderPartBackground:
		return "Background"
	case RenderPartUnderline:
		return "Underline"
	case RenderPartStrikethrough:
		return "Strikethrough"
	case RenderPartOverline:
		return "Overline"
	default:
		return fmt.Sprintf("RenderPart(%d)", r)
	}
}

// RendererOverrider contains methods that are overridable.
type RendererOverrider interface {
	externglib.Objector
	Begin()
	// DrawErrorUnderline: draw a squiggly line that approximately covers the
	// given rectangle in the style of an underline used to indicate a spelling
	// error.
	//
	// The width of the underline is rounded to an integer number of up/down
	// segments and the resulting rectangle is centered in the original
	// rectangle.
	//
	// This should be called while renderer is already active. Use
	// pango.Renderer.Activate() to activate a renderer.
	//
	// The function takes the following parameters:
	//
	//    - x: x coordinate of underline, in Pango units in user coordinate
	//      system.
	//    - y: y coordinate of underline, in Pango units in user coordinate
	//      system.
	//    - width of underline, in Pango units in user coordinate system.
	//    - height of underline, in Pango units in user coordinate system.
	//
	DrawErrorUnderline(x, y, width, height int)
	// DrawGlyph draws a single glyph with coordinates in device space.
	//
	// The function takes the following parameters:
	//
	//    - font: Font.
	//    - glyph index of a single glyph.
	//    - x: x coordinate of left edge of baseline of glyph.
	//    - y: y coordinate of left edge of baseline of glyph.
	//
	DrawGlyph(font Fonter, glyph Glyph, x, y float64)
	// DrawGlyphItem draws the glyphs in glyph_item with the specified
	// PangoRenderer, embedding the text associated with the glyphs in the
	// output if the output format supports it.
	//
	// This is useful for rendering text in PDF.
	//
	// Note that text is the start of the text for layout, which is then indexed
	// by glyph_item->item->offset.
	//
	// If text is NULL, this simply calls pango.Renderer.DrawGlyphs().
	//
	// The default implementation of this method simply falls back to
	// pango.Renderer.DrawGlyphs().
	//
	// The function takes the following parameters:
	//
	//    - text (optional): UTF-8 text that glyph_item refers to, or NULL.
	//    - glyphItem: PangoGlyphItem.
	//    - x: x position of left edge of baseline, in user space coordinates in
	//      Pango units.
	//    - y: y position of left edge of baseline, in user space coordinates in
	//      Pango units.
	//
	DrawGlyphItem(text string, glyphItem *GlyphItem, x, y int)
	// DrawGlyphs draws the glyphs in glyphs with the specified PangoRenderer.
	//
	// The function takes the following parameters:
	//
	//    - font: PangoFont.
	//    - glyphs: PangoGlyphString.
	//    - x: x position of left edge of baseline, in user space coordinates in
	//      Pango units.
	//    - y: y position of left edge of baseline, in user space coordinates in
	//      Pango units.
	//
	DrawGlyphs(font Fonter, glyphs *GlyphString, x, y int)
	// DrawRectangle draws an axis-aligned rectangle in user space coordinates
	// with the specified PangoRenderer.
	//
	// This should be called while renderer is already active. Use
	// pango.Renderer.Activate() to activate a renderer.
	//
	// The function takes the following parameters:
	//
	//    - part: type of object this rectangle is part of.
	//    - x: x position at which to draw rectangle, in user space coordinates
	//      in Pango units.
	//    - y: y position at which to draw rectangle, in user space coordinates
	//      in Pango units.
	//    - width of rectangle in Pango units.
	//    - height of rectangle in Pango units.
	//
	DrawRectangle(part RenderPart, x, y, width, height int)
	// The function takes the following parameters:
	//
	//    - attr
	//    - x
	//    - y
	//
	DrawShape(attr *AttrShape, x, y int)
	// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the
	// X axis using the given PangoRenderer; coordinates are in device space.
	//
	// The function takes the following parameters:
	//
	//    - part: type of object this trapezoid is part of.
	//    - y1: y coordinate of top of trapezoid.
	//    - x11: x coordinate of left end of top of trapezoid.
	//    - x21: x coordinate of right end of top of trapezoid.
	//    - y2: y coordinate of bottom of trapezoid.
	//    - x12: x coordinate of left end of bottom of trapezoid.
	//    - x22: x coordinate of right end of bottom of trapezoid.
	//
	DrawTrapezoid(part RenderPart, y1, x11, x21, y2, x12, x22 float64)
	End()
	// PartChanged informs Pango that the way that the rendering is done for
	// part has changed.
	//
	// This should be called if the rendering changes in a way that would
	// prevent multiple pieces being joined together into one drawing call. For
	// instance, if a subclass of PangoRenderer was to add a stipple option for
	// drawing underlines, it needs to call
	//
	//    pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE);
	//
	//
	// When the stipple changes or underlines with different stipples might be
	// joined together. Pango automatically calls this for changes to colors.
	// (See pango.Renderer.SetColor()).
	//
	// The function takes the following parameters:
	//
	//    - part for which rendering has changed.
	//
	PartChanged(part RenderPart)
	// The function takes the following parameters:
	//
	PrepareRun(run *LayoutRun)
}

// WrapRendererOverrider wraps the RendererOverrider
// interface implementation to access the instance methods.
func WrapRendererOverrider(obj RendererOverrider) *Renderer {
	return wrapRenderer(externglib.BaseObject(obj))
}

// Renderer: PangoRenderer is a base class for objects that can render text
// provided as PangoGlyphString or PangoLayout.
//
// By subclassing PangoRenderer and overriding operations such as draw_glyphs
// and draw_rectangle, renderers for particular font backends and destinations
// can be created.
type Renderer struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Renderer)(nil)
)

// Rendererer describes types inherited from class Renderer.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Rendererer interface {
	externglib.Objector
	baseRenderer() *Renderer
}

var _ Rendererer = (*Renderer)(nil)

func classInitRendererer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.PangoRendererClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.PangoRendererClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Begin() }); ok {
		pclass.begin = (*[0]byte)(C._gotk4_pango1_RendererClass_begin)
	}

	if _, ok := goval.(interface{ DrawErrorUnderline(x, y, width, height int) }); ok {
		pclass.draw_error_underline = (*[0]byte)(C._gotk4_pango1_RendererClass_draw_error_underline)
	}

	if _, ok := goval.(interface {
		DrawGlyph(font Fonter, glyph Glyph, x, y float64)
	}); ok {
		pclass.draw_glyph = (*[0]byte)(C._gotk4_pango1_RendererClass_draw_glyph)
	}

	if _, ok := goval.(interface {
		DrawGlyphItem(text string, glyphItem *GlyphItem, x, y int)
	}); ok {
		pclass.draw_glyph_item = (*[0]byte)(C._gotk4_pango1_RendererClass_draw_glyph_item)
	}

	if _, ok := goval.(interface {
		DrawGlyphs(font Fonter, glyphs *GlyphString, x, y int)
	}); ok {
		pclass.draw_glyphs = (*[0]byte)(C._gotk4_pango1_RendererClass_draw_glyphs)
	}

	if _, ok := goval.(interface {
		DrawRectangle(part RenderPart, x, y, width, height int)
	}); ok {
		pclass.draw_rectangle = (*[0]byte)(C._gotk4_pango1_RendererClass_draw_rectangle)
	}

	if _, ok := goval.(interface {
		DrawShape(attr *AttrShape, x, y int)
	}); ok {
		pclass.draw_shape = (*[0]byte)(C._gotk4_pango1_RendererClass_draw_shape)
	}

	if _, ok := goval.(interface {
		DrawTrapezoid(part RenderPart, y1, x11, x21, y2, x12, x22 float64)
	}); ok {
		pclass.draw_trapezoid = (*[0]byte)(C._gotk4_pango1_RendererClass_draw_trapezoid)
	}

	if _, ok := goval.(interface{ End() }); ok {
		pclass.end = (*[0]byte)(C._gotk4_pango1_RendererClass_end)
	}

	if _, ok := goval.(interface{ PartChanged(part RenderPart) }); ok {
		pclass.part_changed = (*[0]byte)(C._gotk4_pango1_RendererClass_part_changed)
	}

	if _, ok := goval.(interface{ PrepareRun(run *LayoutRun) }); ok {
		pclass.prepare_run = (*[0]byte)(C._gotk4_pango1_RendererClass_prepare_run)
	}
}

//export _gotk4_pango1_RendererClass_begin
func _gotk4_pango1_RendererClass_begin(arg0 *C.PangoRenderer) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Begin() })

	iface.Begin()
}

//export _gotk4_pango1_RendererClass_draw_error_underline
func _gotk4_pango1_RendererClass_draw_error_underline(arg0 *C.PangoRenderer, arg1 C.int, arg2 C.int, arg3 C.int, arg4 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ DrawErrorUnderline(x, y, width, height int) })

	var _x int      // out
	var _y int      // out
	var _width int  // out
	var _height int // out

	_x = int(arg1)
	_y = int(arg2)
	_width = int(arg3)
	_height = int(arg4)

	iface.DrawErrorUnderline(_x, _y, _width, _height)
}

//export _gotk4_pango1_RendererClass_draw_glyph
func _gotk4_pango1_RendererClass_draw_glyph(arg0 *C.PangoRenderer, arg1 *C.PangoFont, arg2 C.PangoGlyph, arg3 C.double, arg4 C.double) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DrawGlyph(font Fonter, glyph Glyph, x, y float64)
	})

	var _font Fonter // out
	var _glyph Glyph // out
	var _x float64   // out
	var _y float64   // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Fonter)
			return ok
		})
		rv, ok := casted.(Fonter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fonter")
		}
		_font = rv
	}
	_glyph = uint32(arg2)
	_x = float64(arg3)
	_y = float64(arg4)

	iface.DrawGlyph(_font, _glyph, _x, _y)
}

//export _gotk4_pango1_RendererClass_draw_glyph_item
func _gotk4_pango1_RendererClass_draw_glyph_item(arg0 *C.PangoRenderer, arg1 *C.char, arg2 *C.PangoGlyphItem, arg3 C.int, arg4 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DrawGlyphItem(text string, glyphItem *GlyphItem, x, y int)
	})

	var _text string          // out
	var _glyphItem *GlyphItem // out
	var _x int                // out
	var _y int                // out

	if arg1 != nil {
		_text = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	}
	_glyphItem = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_x = int(arg3)
	_y = int(arg4)

	iface.DrawGlyphItem(_text, _glyphItem, _x, _y)
}

//export _gotk4_pango1_RendererClass_draw_glyphs
func _gotk4_pango1_RendererClass_draw_glyphs(arg0 *C.PangoRenderer, arg1 *C.PangoFont, arg2 *C.PangoGlyphString, arg3 C.int, arg4 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DrawGlyphs(font Fonter, glyphs *GlyphString, x, y int)
	})

	var _font Fonter         // out
	var _glyphs *GlyphString // out
	var _x int               // out
	var _y int               // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Fonter)
			return ok
		})
		rv, ok := casted.(Fonter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fonter")
		}
		_font = rv
	}
	_glyphs = (*GlyphString)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_x = int(arg3)
	_y = int(arg4)

	iface.DrawGlyphs(_font, _glyphs, _x, _y)
}

//export _gotk4_pango1_RendererClass_draw_rectangle
func _gotk4_pango1_RendererClass_draw_rectangle(arg0 *C.PangoRenderer, arg1 C.PangoRenderPart, arg2 C.int, arg3 C.int, arg4 C.int, arg5 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DrawRectangle(part RenderPart, x, y, width, height int)
	})

	var _part RenderPart // out
	var _x int           // out
	var _y int           // out
	var _width int       // out
	var _height int      // out

	_part = RenderPart(arg1)
	_x = int(arg2)
	_y = int(arg3)
	_width = int(arg4)
	_height = int(arg5)

	iface.DrawRectangle(_part, _x, _y, _width, _height)
}

//export _gotk4_pango1_RendererClass_draw_shape
func _gotk4_pango1_RendererClass_draw_shape(arg0 *C.PangoRenderer, arg1 *C.PangoAttrShape, arg2 C.int, arg3 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DrawShape(attr *AttrShape, x, y int)
	})

	var _attr *AttrShape // out
	var _x int           // out
	var _y int           // out

	_attr = (*AttrShape)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_x = int(arg2)
	_y = int(arg3)

	iface.DrawShape(_attr, _x, _y)
}

//export _gotk4_pango1_RendererClass_draw_trapezoid
func _gotk4_pango1_RendererClass_draw_trapezoid(arg0 *C.PangoRenderer, arg1 C.PangoRenderPart, arg2 C.double, arg3 C.double, arg4 C.double, arg5 C.double, arg6 C.double, arg7 C.double) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		DrawTrapezoid(part RenderPart, y1, x11, x21, y2, x12, x22 float64)
	})

	var _part RenderPart // out
	var _y1 float64      // out
	var _x11 float64     // out
	var _x21 float64     // out
	var _y2 float64      // out
	var _x12 float64     // out
	var _x22 float64     // out

	_part = RenderPart(arg1)
	_y1 = float64(arg2)
	_x11 = float64(arg3)
	_x21 = float64(arg4)
	_y2 = float64(arg5)
	_x12 = float64(arg6)
	_x22 = float64(arg7)

	iface.DrawTrapezoid(_part, _y1, _x11, _x21, _y2, _x12, _x22)
}

//export _gotk4_pango1_RendererClass_end
func _gotk4_pango1_RendererClass_end(arg0 *C.PangoRenderer) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ End() })

	iface.End()
}

//export _gotk4_pango1_RendererClass_part_changed
func _gotk4_pango1_RendererClass_part_changed(arg0 *C.PangoRenderer, arg1 C.PangoRenderPart) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PartChanged(part RenderPart) })

	var _part RenderPart // out

	_part = RenderPart(arg1)

	iface.PartChanged(_part)
}

//export _gotk4_pango1_RendererClass_prepare_run
func _gotk4_pango1_RendererClass_prepare_run(arg0 *C.PangoRenderer, arg1 *C.PangoLayoutRun) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PrepareRun(run *LayoutRun) })

	var _run *LayoutRun // out

	_run = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.PrepareRun(_run)
}

func wrapRenderer(obj *externglib.Object) *Renderer {
	return &Renderer{
		Object: obj,
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	return wrapRenderer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (renderer *Renderer) baseRenderer() *Renderer {
	return renderer
}

// BaseRenderer returns the underlying base object from the
// interface.
func BaseRenderer(obj Rendererer) *Renderer {
	return obj.baseRenderer()
}

// Activate does initial setup before rendering operations on renderer.
//
// pango.Renderer.Deactivate() should be called when done drawing. Calls such as
// pango.Renderer.DrawLayout() automatically activate the layout before drawing
// on it. Calls to pango_renderer_activate() and pango_renderer_deactivate() can
// be nested and the renderer will only be initialized and deinitialized once.
func (renderer *Renderer) Activate() {
	var _arg0 *C.PangoRenderer // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	C.pango_renderer_activate(_arg0)
	runtime.KeepAlive(renderer)
}

// Deactivate cleans up after rendering operations on renderer.
//
// See docs for pango.Renderer.Activate().
func (renderer *Renderer) Deactivate() {
	var _arg0 *C.PangoRenderer // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	C.pango_renderer_deactivate(_arg0)
	runtime.KeepAlive(renderer)
}

// DrawErrorUnderline: draw a squiggly line that approximately covers the given
// rectangle in the style of an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
//
// This should be called while renderer is already active. Use
// pango.Renderer.Activate() to activate a renderer.
//
// The function takes the following parameters:
//
//    - x: x coordinate of underline, in Pango units in user coordinate system.
//    - y: y coordinate of underline, in Pango units in user coordinate system.
//    - width of underline, in Pango units in user coordinate system.
//    - height of underline, in Pango units in user coordinate system.
//
func (renderer *Renderer) DrawErrorUnderline(x, y, width, height int) {
	var _arg0 *C.PangoRenderer // out
	var _arg1 C.int            // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out
	var _arg4 C.int            // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)
	_arg3 = C.int(width)
	_arg4 = C.int(height)

	C.pango_renderer_draw_error_underline(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// DrawGlyph draws a single glyph with coordinates in device space.
//
// The function takes the following parameters:
//
//    - font: Font.
//    - glyph index of a single glyph.
//    - x: x coordinate of left edge of baseline of glyph.
//    - y: y coordinate of left edge of baseline of glyph.
//
func (renderer *Renderer) DrawGlyph(font Fonter, glyph Glyph, x, y float64) {
	var _arg0 *C.PangoRenderer // out
	var _arg1 *C.PangoFont     // out
	var _arg2 C.PangoGlyph     // out
	var _arg3 C.double         // out
	var _arg4 C.double         // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(externglib.InternObject(font).Native()))
	_arg2 = C.guint32(glyph)
	_arg3 = C.double(x)
	_arg4 = C.double(y)

	C.pango_renderer_draw_glyph(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(font)
	runtime.KeepAlive(glyph)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// DrawGlyphItem draws the glyphs in glyph_item with the specified
// PangoRenderer, embedding the text associated with the glyphs in the output if
// the output format supports it.
//
// This is useful for rendering text in PDF.
//
// Note that text is the start of the text for layout, which is then indexed by
// glyph_item->item->offset.
//
// If text is NULL, this simply calls pango.Renderer.DrawGlyphs().
//
// The default implementation of this method simply falls back to
// pango.Renderer.DrawGlyphs().
//
// The function takes the following parameters:
//
//    - text (optional): UTF-8 text that glyph_item refers to, or NULL.
//    - glyphItem: PangoGlyphItem.
//    - x: x position of left edge of baseline, in user space coordinates in
//      Pango units.
//    - y: y position of left edge of baseline, in user space coordinates in
//      Pango units.
//
func (renderer *Renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x, y int) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 *C.char           // out
	var _arg2 *C.PangoGlyphItem // out
	var _arg3 C.int             // out
	var _arg4 C.int             // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if text != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(glyphItem)))
	_arg3 = C.int(x)
	_arg4 = C.int(y)

	C.pango_renderer_draw_glyph_item(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(text)
	runtime.KeepAlive(glyphItem)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// DrawGlyphs draws the glyphs in glyphs with the specified PangoRenderer.
//
// The function takes the following parameters:
//
//    - font: PangoFont.
//    - glyphs: PangoGlyphString.
//    - x: x position of left edge of baseline, in user space coordinates in
//      Pango units.
//    - y: y position of left edge of baseline, in user space coordinates in
//      Pango units.
//
func (renderer *Renderer) DrawGlyphs(font Fonter, glyphs *GlyphString, x, y int) {
	var _arg0 *C.PangoRenderer    // out
	var _arg1 *C.PangoFont        // out
	var _arg2 *C.PangoGlyphString // out
	var _arg3 C.int               // out
	var _arg4 C.int               // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(externglib.InternObject(font).Native()))
	_arg2 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg3 = C.int(x)
	_arg4 = C.int(y)

	C.pango_renderer_draw_glyphs(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(font)
	runtime.KeepAlive(glyphs)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// DrawLayout draws layout with the specified PangoRenderer.
//
// The function takes the following parameters:
//
//    - layout: PangoLayout.
//    - x: x position of left edge of baseline, in user space coordinates in
//      Pango units.
//    - y: y position of left edge of baseline, in user space coordinates in
//      Pango units.
//
func (renderer *Renderer) DrawLayout(layout *Layout, x, y int) {
	var _arg0 *C.PangoRenderer // out
	var _arg1 *C.PangoLayout   // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = (*C.PangoLayout)(unsafe.Pointer(externglib.InternObject(layout).Native()))
	_arg2 = C.int(x)
	_arg3 = C.int(y)

	C.pango_renderer_draw_layout(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// DrawLayoutLine draws line with the specified PangoRenderer.
//
// The function takes the following parameters:
//
//    - line: PangoLayoutLine.
//    - x: x position of left edge of baseline, in user space coordinates in
//      Pango units.
//    - y: y position of left edge of baseline, in user space coordinates in
//      Pango units.
//
func (renderer *Renderer) DrawLayoutLine(line *LayoutLine, x, y int) {
	var _arg0 *C.PangoRenderer   // out
	var _arg1 *C.PangoLayoutLine // out
	var _arg2 C.int              // out
	var _arg3 C.int              // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = (*C.PangoLayoutLine)(gextras.StructNative(unsafe.Pointer(line)))
	_arg2 = C.int(x)
	_arg3 = C.int(y)

	C.pango_renderer_draw_layout_line(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(line)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// DrawRectangle draws an axis-aligned rectangle in user space coordinates with
// the specified PangoRenderer.
//
// This should be called while renderer is already active. Use
// pango.Renderer.Activate() to activate a renderer.
//
// The function takes the following parameters:
//
//    - part: type of object this rectangle is part of.
//    - x: x position at which to draw rectangle, in user space coordinates in
//      Pango units.
//    - y: y position at which to draw rectangle, in user space coordinates in
//      Pango units.
//    - width of rectangle in Pango units.
//    - height of rectangle in Pango units.
//
func (renderer *Renderer) DrawRectangle(part RenderPart, x, y, width, height int) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 C.int             // out
	var _arg3 C.int             // out
	var _arg4 C.int             // out
	var _arg5 C.int             // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.PangoRenderPart(part)
	_arg2 = C.int(x)
	_arg3 = C.int(y)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	C.pango_renderer_draw_rectangle(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(part)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the X
// axis using the given PangoRenderer; coordinates are in device space.
//
// The function takes the following parameters:
//
//    - part: type of object this trapezoid is part of.
//    - y1: y coordinate of top of trapezoid.
//    - x11: x coordinate of left end of top of trapezoid.
//    - x21: x coordinate of right end of top of trapezoid.
//    - y2: y coordinate of bottom of trapezoid.
//    - x12: x coordinate of left end of bottom of trapezoid.
//    - x22: x coordinate of right end of bottom of trapezoid.
//
func (renderer *Renderer) DrawTrapezoid(part RenderPart, y1, x11, x21, y2, x12, x22 float64) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 C.double          // out
	var _arg3 C.double          // out
	var _arg4 C.double          // out
	var _arg5 C.double          // out
	var _arg6 C.double          // out
	var _arg7 C.double          // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.PangoRenderPart(part)
	_arg2 = C.double(y1)
	_arg3 = C.double(x11)
	_arg4 = C.double(x21)
	_arg5 = C.double(y2)
	_arg6 = C.double(x12)
	_arg7 = C.double(x22)

	C.pango_renderer_draw_trapezoid(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(part)
	runtime.KeepAlive(y1)
	runtime.KeepAlive(x11)
	runtime.KeepAlive(x21)
	runtime.KeepAlive(y2)
	runtime.KeepAlive(x12)
	runtime.KeepAlive(x22)
}

// Alpha gets the current alpha for the specified part.
//
// The function takes the following parameters:
//
//    - part to get the alpha for.
//
// The function returns the following values:
//
//    - guint16: alpha for the specified part, or 0 if it hasn't been set and
//      should be inherited from the environment.
//
func (renderer *Renderer) Alpha(part RenderPart) uint16 {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _cret C.guint16         // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.PangoRenderPart(part)

	_cret = C.pango_renderer_get_alpha(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(part)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Color gets the current rendering color for the specified part.
//
// The function takes the following parameters:
//
//    - part to get the color for.
//
// The function returns the following values:
//
//    - color (optional) for the specified part, or NULL if it hasn't been set
//      and should be inherited from the environment.
//
func (renderer *Renderer) Color(part RenderPart) *Color {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _cret *C.PangoColor     // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.PangoRenderPart(part)

	_cret = C.pango_renderer_get_color(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(part)

	var _color *Color // out

	if _cret != nil {
		_color = (*Color)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _color
}

// Layout gets the layout currently being rendered using renderer.
//
// Calling this function only makes sense from inside a subclass's methods, like
// in its draw_shape vfunc, for example.
//
// The returned layout should not be modified while still being rendered.
//
// The function returns the following values:
//
//    - layout (optional): layout, or NULL if no layout is being rendered using
//      renderer at this time.
//
func (renderer *Renderer) Layout() *Layout {
	var _arg0 *C.PangoRenderer // out
	var _cret *C.PangoLayout   // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	_cret = C.pango_renderer_get_layout(_arg0)
	runtime.KeepAlive(renderer)

	var _layout *Layout // out

	if _cret != nil {
		_layout = wrapLayout(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _layout
}

// LayoutLine gets the layout line currently being rendered using renderer.
//
// Calling this function only makes sense from inside a subclass's methods, like
// in its draw_shape vfunc, for example.
//
// The returned layout line should not be modified while still being rendered.
//
// The function returns the following values:
//
//    - layoutLine (optional): layout line, or NULL if no layout line is being
//      rendered using renderer at this time.
//
func (renderer *Renderer) LayoutLine() *LayoutLine {
	var _arg0 *C.PangoRenderer   // out
	var _cret *C.PangoLayoutLine // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	_cret = C.pango_renderer_get_layout_line(_arg0)
	runtime.KeepAlive(renderer)

	var _layoutLine *LayoutLine // out

	if _cret != nil {
		_layoutLine = (*LayoutLine)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.pango_layout_line_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_layoutLine)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_layout_line_unref((*C.PangoLayoutLine)(intern.C))
			},
		)
	}

	return _layoutLine
}

// Matrix gets the transformation matrix that will be applied when rendering.
//
// See pango.Renderer.SetMatrix().
//
// The function returns the following values:
//
//    - matrix (optional): matrix, or NULL if no matrix has been set (which is
//      the same as the identity matrix). The returned matrix is owned by Pango
//      and must not be modified or freed.
//
func (renderer *Renderer) Matrix() *Matrix {
	var _arg0 *C.PangoRenderer // out
	var _cret *C.PangoMatrix   // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))

	_cret = C.pango_renderer_get_matrix(_arg0)
	runtime.KeepAlive(renderer)

	var _matrix *Matrix // out

	if _cret != nil {
		_matrix = (*Matrix)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _matrix
}

// PartChanged informs Pango that the way that the rendering is done for part
// has changed.
//
// This should be called if the rendering changes in a way that would prevent
// multiple pieces being joined together into one drawing call. For instance, if
// a subclass of PangoRenderer was to add a stipple option for drawing
// underlines, it needs to call
//
//    pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE);
//
//
// When the stipple changes or underlines with different stipples might be
// joined together. Pango automatically calls this for changes to colors. (See
// pango.Renderer.SetColor()).
//
// The function takes the following parameters:
//
//    - part for which rendering has changed.
//
func (renderer *Renderer) PartChanged(part RenderPart) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.PangoRenderPart(part)

	C.pango_renderer_part_changed(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(part)
}

// SetAlpha sets the alpha for part of the rendering.
//
// Note that the alpha may only be used if a color is specified for part as
// well.
//
// The function takes the following parameters:
//
//    - part to set the alpha for.
//    - alpha value between 1 and 65536, or 0 to unset the alpha.
//
func (renderer *Renderer) SetAlpha(part RenderPart, alpha uint16) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 C.guint16         // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.PangoRenderPart(part)
	_arg2 = C.guint16(alpha)

	C.pango_renderer_set_alpha(_arg0, _arg1, _arg2)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(part)
	runtime.KeepAlive(alpha)
}

// SetColor sets the color for part of the rendering.
//
// Also see pango.Renderer.SetAlpha().
//
// The function takes the following parameters:
//
//    - part to change the color of.
//    - color (optional): new color or NULL to unset the current color.
//
func (renderer *Renderer) SetColor(part RenderPart, color *Color) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 *C.PangoColor     // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	_arg1 = C.PangoRenderPart(part)
	if color != nil {
		_arg2 = (*C.PangoColor)(gextras.StructNative(unsafe.Pointer(color)))
	}

	C.pango_renderer_set_color(_arg0, _arg1, _arg2)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(part)
	runtime.KeepAlive(color)
}

// SetMatrix sets the transformation matrix that will be applied when rendering.
//
// The function takes the following parameters:
//
//    - matrix (optional): PangoMatrix, or NULL to unset any existing matrix. (No
//      matrix set is the same as setting the identity matrix.).
//
func (renderer *Renderer) SetMatrix(matrix *Matrix) {
	var _arg0 *C.PangoRenderer // out
	var _arg1 *C.PangoMatrix   // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if matrix != nil {
		_arg1 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	}

	C.pango_renderer_set_matrix(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(matrix)
}
