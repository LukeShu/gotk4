// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_render_part_get_type()), F: marshalRenderPart},
		{T: externglib.Type(C.pango_renderer_get_type()), F: marshalRenderer},
	})
}

// RenderPart defines different items to render for such purposes as setting
// colors.
type RenderPart int

const (
	// Foreground: the text itself
	RenderPartForeground RenderPart = iota
	// Background: the area behind the text
	RenderPartBackground
	// Underline: underlines
	RenderPartUnderline
	// Strikethrough lines
	RenderPartStrikethrough
	// Overline: overlines
	RenderPartOverline
)

func marshalRenderPart(p uintptr) (interface{}, error) {
	return RenderPart(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RendererOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RendererOverrider interface {
	Begin()
	// DrawErrorUnderline: draw a squiggly line that approximately covers the
	// given rectangle in the style of an underline used to indicate a spelling
	// error.
	//
	// The width of the underline is rounded to an integer number of up/down
	// segments and the resulting rectangle is centered in the original
	// rectangle.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawErrorUnderline(x int, y int, width int, height int)
	// DrawGlyphItem draws the glyphs in @glyph_item with the specified
	// `PangoRenderer`, embedding the text associated with the glyphs in the
	// output if the output format supports it.
	//
	// This is useful for rendering text in PDF.
	//
	// Note that @text is the start of the text for layout, which is then
	// indexed by `glyph_item->item->offset`.
	//
	// If @text is nil, this simply calls [method@Pango.Renderer.draw_glyphs].
	//
	// The default implementation of this method simply falls back to
	// [method@Pango.Renderer.draw_glyphs].
	DrawGlyphItem(text string, glyphItem *GlyphItem, x int, y int)
	// DrawGlyphs draws the glyphs in @glyphs with the specified
	// `PangoRenderer`.
	DrawGlyphs(font Font, glyphs *GlyphString, x int, y int)
	// DrawRectangle draws an axis-aligned rectangle in user space coordinates
	// with the specified `PangoRenderer`.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawRectangle(part RenderPart, x int, y int, width int, height int)
	DrawShape(attr *AttrShape, x int, y int)
	// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the
	// X axis using the given `PangoRenderer`; coordinates are in device space.
	DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64)
	End()
	// PartChanged informs Pango that the way that the rendering is done for
	// @part has changed.
	//
	// This should be called if the rendering changes in a way that would
	// prevent multiple pieces being joined together into one drawing call. For
	// instance, if a subclass of `PangoRenderer` was to add a stipple option
	// for drawing underlines, it needs to call
	//
	// “` pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE); “`
	//
	// When the stipple changes or underlines with different stipples might be
	// joined together. Pango automatically calls this for changes to colors.
	// (See [method@Pango.Renderer.set_color])
	PartChanged(part RenderPart)
}

// Renderer: `PangoRenderer` is a base class for objects that can render text
// provided as `PangoGlyphString` or `PangoLayout`.
//
// By subclassing `PangoRenderer` and overriding operations such as @draw_glyphs
// and @draw_rectangle, renderers for particular font backends and destinations
// can be created.
type Renderer interface {
	gextras.Objector

	// Activate does initial setup before rendering operations on @renderer.
	//
	// [method@Pango.Renderer.deactivate] should be called when done drawing.
	// Calls such as [method@Pango.Renderer.draw_layout] automatically activate
	// the layout before drawing on it. Calls to `pango_renderer_activate()` and
	// `pango_renderer_deactivate()` can be nested and the renderer will only be
	// initialized and deinitialized once.
	Activate()
	// Deactivate cleans up after rendering operations on @renderer.
	//
	// See docs for [method@Pango.Renderer.activate].
	Deactivate()
	// DrawErrorUnderline: draw a squiggly line that approximately covers the
	// given rectangle in the style of an underline used to indicate a spelling
	// error.
	//
	// The width of the underline is rounded to an integer number of up/down
	// segments and the resulting rectangle is centered in the original
	// rectangle.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawErrorUnderline(x int, y int, width int, height int)
	// DrawGlyphItem draws the glyphs in @glyph_item with the specified
	// `PangoRenderer`, embedding the text associated with the glyphs in the
	// output if the output format supports it.
	//
	// This is useful for rendering text in PDF.
	//
	// Note that @text is the start of the text for layout, which is then
	// indexed by `glyph_item->item->offset`.
	//
	// If @text is nil, this simply calls [method@Pango.Renderer.draw_glyphs].
	//
	// The default implementation of this method simply falls back to
	// [method@Pango.Renderer.draw_glyphs].
	DrawGlyphItem(text string, glyphItem *GlyphItem, x int, y int)
	// DrawGlyphs draws the glyphs in @glyphs with the specified
	// `PangoRenderer`.
	DrawGlyphs(font Font, glyphs *GlyphString, x int, y int)
	// DrawLayout draws @layout with the specified `PangoRenderer`.
	DrawLayout(layout Layout, x int, y int)
	// DrawRectangle draws an axis-aligned rectangle in user space coordinates
	// with the specified `PangoRenderer`.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawRectangle(part RenderPart, x int, y int, width int, height int)
	// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the
	// X axis using the given `PangoRenderer`; coordinates are in device space.
	DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64)
	// Alpha gets the current alpha for the specified part.
	Alpha(part RenderPart) uint16
	// Color gets the current rendering color for the specified part.
	Color(part RenderPart) *Color
	// Layout gets the layout currently being rendered using @renderer.
	//
	// Calling this function only makes sense from inside a subclass's methods,
	// like in its draw_shape vfunc, for example.
	//
	// The returned layout should not be modified while still being rendered.
	Layout() Layout
	// Matrix gets the transformation matrix that will be applied when
	// rendering.
	//
	// See [method@Pango.Renderer.set_matrix].
	Matrix() *Matrix
	// PartChanged informs Pango that the way that the rendering is done for
	// @part has changed.
	//
	// This should be called if the rendering changes in a way that would
	// prevent multiple pieces being joined together into one drawing call. For
	// instance, if a subclass of `PangoRenderer` was to add a stipple option
	// for drawing underlines, it needs to call
	//
	// “` pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE); “`
	//
	// When the stipple changes or underlines with different stipples might be
	// joined together. Pango automatically calls this for changes to colors.
	// (See [method@Pango.Renderer.set_color])
	PartChanged(part RenderPart)
	// SetAlpha sets the alpha for part of the rendering.
	//
	// Note that the alpha may only be used if a color is specified for @part as
	// well.
	SetAlpha(part RenderPart, alpha uint16)
	// SetColor sets the color for part of the rendering.
	//
	// Also see [method@Pango.Renderer.set_alpha].
	SetColor(part RenderPart, color *Color)
	// SetMatrix sets the transformation matrix that will be applied when
	// rendering.
	SetMatrix(matrix *Matrix)
}

// RendererClass implements the Renderer interface.
type RendererClass struct {
	*externglib.Object
}

var _ Renderer = (*RendererClass)(nil)

func wrapRenderer(obj *externglib.Object) Renderer {
	return &RendererClass{
		Object: obj,
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRenderer(obj), nil
}

// Activate does initial setup before rendering operations on @renderer.
//
// [method@Pango.Renderer.deactivate] should be called when done drawing. Calls
// such as [method@Pango.Renderer.draw_layout] automatically activate the layout
// before drawing on it. Calls to `pango_renderer_activate()` and
// `pango_renderer_deactivate()` can be nested and the renderer will only be
// initialized and deinitialized once.
func (r *RendererClass) Activate() {
	var _arg0 *C.PangoRenderer // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	C.pango_renderer_activate(_arg0)
}

// Deactivate cleans up after rendering operations on @renderer.
//
// See docs for [method@Pango.Renderer.activate].
func (r *RendererClass) Deactivate() {
	var _arg0 *C.PangoRenderer // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	C.pango_renderer_deactivate(_arg0)
}

// DrawErrorUnderline: draw a squiggly line that approximately covers the given
// rectangle in the style of an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
//
// This should be called while @renderer is already active. Use
// [method@Pango.Renderer.activate] to activate a renderer.
func (r *RendererClass) DrawErrorUnderline(x int, y int, width int, height int) {
	var _arg0 *C.PangoRenderer // out
	var _arg1 C.int            // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out
	var _arg4 C.int            // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)
	_arg3 = C.int(width)
	_arg4 = C.int(height)

	C.pango_renderer_draw_error_underline(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// DrawGlyphItem draws the glyphs in @glyph_item with the specified
// `PangoRenderer`, embedding the text associated with the glyphs in the output
// if the output format supports it.
//
// This is useful for rendering text in PDF.
//
// Note that @text is the start of the text for layout, which is then indexed by
// `glyph_item->item->offset`.
//
// If @text is nil, this simply calls [method@Pango.Renderer.draw_glyphs].
//
// The default implementation of this method simply falls back to
// [method@Pango.Renderer.draw_glyphs].
func (r *RendererClass) DrawGlyphItem(text string, glyphItem *GlyphItem, x int, y int) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 *C.char           // out
	var _arg2 *C.PangoGlyphItem // out
	var _arg3 C.int             // out
	var _arg4 C.int             // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem))
	_arg3 = C.int(x)
	_arg4 = C.int(y)

	C.pango_renderer_draw_glyph_item(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// DrawGlyphs draws the glyphs in @glyphs with the specified `PangoRenderer`.
func (r *RendererClass) DrawGlyphs(font Font, glyphs *GlyphString, x int, y int) {
	var _arg0 *C.PangoRenderer    // out
	var _arg1 *C.PangoFont        // out
	var _arg2 *C.PangoGlyphString // out
	var _arg3 C.int               // out
	var _arg4 C.int               // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	_arg2 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs))
	_arg3 = C.int(x)
	_arg4 = C.int(y)

	C.pango_renderer_draw_glyphs(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// DrawLayout draws @layout with the specified `PangoRenderer`.
func (r *RendererClass) DrawLayout(layout Layout, x int, y int) {
	var _arg0 *C.PangoRenderer // out
	var _arg1 *C.PangoLayout   // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	_arg2 = C.int(x)
	_arg3 = C.int(y)

	C.pango_renderer_draw_layout(_arg0, _arg1, _arg2, _arg3)
}

// DrawRectangle draws an axis-aligned rectangle in user space coordinates with
// the specified `PangoRenderer`.
//
// This should be called while @renderer is already active. Use
// [method@Pango.Renderer.activate] to activate a renderer.
func (r *RendererClass) DrawRectangle(part RenderPart, x int, y int, width int, height int) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 C.int             // out
	var _arg3 C.int             // out
	var _arg4 C.int             // out
	var _arg5 C.int             // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.PangoRenderPart(part)
	_arg2 = C.int(x)
	_arg3 = C.int(y)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	C.pango_renderer_draw_rectangle(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the X
// axis using the given `PangoRenderer`; coordinates are in device space.
func (r *RendererClass) DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 C.double          // out
	var _arg3 C.double          // out
	var _arg4 C.double          // out
	var _arg5 C.double          // out
	var _arg6 C.double          // out
	var _arg7 C.double          // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.PangoRenderPart(part)
	_arg2 = C.double(y1)
	_arg3 = C.double(x11)
	_arg4 = C.double(x21)
	_arg5 = C.double(y2)
	_arg6 = C.double(x12)
	_arg7 = C.double(x22)

	C.pango_renderer_draw_trapezoid(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// Alpha gets the current alpha for the specified part.
func (r *RendererClass) Alpha(part RenderPart) uint16 {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _cret C.guint16         // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.PangoRenderPart(part)

	_cret = C.pango_renderer_get_alpha(_arg0, _arg1)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Color gets the current rendering color for the specified part.
func (r *RendererClass) Color(part RenderPart) *Color {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _cret *C.PangoColor     // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.PangoRenderPart(part)

	_cret = C.pango_renderer_get_color(_arg0, _arg1)

	var _color *Color // out

	_color = (*Color)(unsafe.Pointer(_cret))

	return _color
}

// Layout gets the layout currently being rendered using @renderer.
//
// Calling this function only makes sense from inside a subclass's methods, like
// in its draw_shape vfunc, for example.
//
// The returned layout should not be modified while still being rendered.
func (r *RendererClass) Layout() Layout {
	var _arg0 *C.PangoRenderer // out
	var _cret *C.PangoLayout   // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	_cret = C.pango_renderer_get_layout(_arg0)

	var _layout Layout // out

	_layout = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Layout)

	return _layout
}

// Matrix gets the transformation matrix that will be applied when rendering.
//
// See [method@Pango.Renderer.set_matrix].
func (r *RendererClass) Matrix() *Matrix {
	var _arg0 *C.PangoRenderer // out
	var _cret *C.PangoMatrix   // in

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	_cret = C.pango_renderer_get_matrix(_arg0)

	var _matrix *Matrix // out

	_matrix = (*Matrix)(unsafe.Pointer(_cret))

	return _matrix
}

// PartChanged informs Pango that the way that the rendering is done for @part
// has changed.
//
// This should be called if the rendering changes in a way that would prevent
// multiple pieces being joined together into one drawing call. For instance, if
// a subclass of `PangoRenderer` was to add a stipple option for drawing
// underlines, it needs to call
//
// “` pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE); “`
//
// When the stipple changes or underlines with different stipples might be
// joined together. Pango automatically calls this for changes to colors. (See
// [method@Pango.Renderer.set_color])
func (r *RendererClass) PartChanged(part RenderPart) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.PangoRenderPart(part)

	C.pango_renderer_part_changed(_arg0, _arg1)
}

// SetAlpha sets the alpha for part of the rendering.
//
// Note that the alpha may only be used if a color is specified for @part as
// well.
func (r *RendererClass) SetAlpha(part RenderPart, alpha uint16) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 C.guint16         // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.PangoRenderPart(part)
	_arg2 = C.guint16(alpha)

	C.pango_renderer_set_alpha(_arg0, _arg1, _arg2)
}

// SetColor sets the color for part of the rendering.
//
// Also see [method@Pango.Renderer.set_alpha].
func (r *RendererClass) SetColor(part RenderPart, color *Color) {
	var _arg0 *C.PangoRenderer  // out
	var _arg1 C.PangoRenderPart // out
	var _arg2 *C.PangoColor     // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = C.PangoRenderPart(part)
	_arg2 = (*C.PangoColor)(unsafe.Pointer(color))

	C.pango_renderer_set_color(_arg0, _arg1, _arg2)
}

// SetMatrix sets the transformation matrix that will be applied when rendering.
func (r *RendererClass) SetMatrix(matrix *Matrix) {
	var _arg0 *C.PangoRenderer // out
	var _arg1 *C.PangoMatrix   // out

	_arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.PangoMatrix)(unsafe.Pointer(matrix))

	C.pango_renderer_set_matrix(_arg0, _arg1)
}
