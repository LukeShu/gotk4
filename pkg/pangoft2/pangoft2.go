// Code generated by girgen. DO NOT EDIT.

package pangoft2

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/pkg/freetype2"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoft2
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pangoft2.h>
//
// // extern void callbackDelete(gpointer);
import "C"

//export callbackDelete
func callbackDelete(ptr C.gpointer) {
	box.Delete(box.Callback, uintptr(ptr))
}

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{

		// Classes
		{T: externglib.Type(C.pango_ft2_font_map_get_type()), F: marshalFontMap},
	})
}

// FontGetCoverage gets the Coverage for a `PangoFT2Font`. Use
// pango_font_get_coverage() instead.
func FontGetCoverage(font pango.Font, language *pango.Language) pango.Coverage {
	var arg1 *C.PangoFont
	var arg2 *C.PangoLanguage

	arg1 = (*C.PangoFont)(font.Native())
	arg2 = (*C.PangoLanguage)(language.Native())

	ret := C.pango_ft2_font_get_coverage(arg1, arg2)

	var ret0 pango.Coverage

	ret0 = pango.WrapCoverage(externglib.AssumeOwnership(unsafe.Pointer(ret.Native())))

	return ret0
}

// FontGetFace returns the native FreeType2 `FT_Face` structure used for this
// Font. This may be useful if you want to use FreeType2 functions directly.
//
// Use pango_fc_font_lock_face() instead; when you are done with a face from
// pango_fc_font_lock_face() you must call pango_fc_font_unlock_face().
func FontGetFace(font pango.Font) freetype2.Face {
	var arg1 *C.PangoFont

	arg1 = (*C.PangoFont)(font.Native())

	ret := C.pango_ft2_font_get_face(arg1)

	var ret0 freetype2.Face

	ret0 = freetype2.WrapFace(ret)

	return ret0
}

// GetContext retrieves a `PangoContext` for the default PangoFT2 fontmap (see
// pango_ft2_font_map_for_display()) and sets the resolution for the default
// fontmap to @dpi_x by @dpi_y.
func GetContext(dpiX float64, dpiY float64) pango.Context {
	var arg1 C.double
	var arg2 C.double

	arg1 = C.double(dpiX)
	arg2 = C.double(dpiY)

	ret := C.pango_ft2_get_context(arg1, arg2)

	var ret0 pango.Context

	ret0 = pango.WrapContext(externglib.AssumeOwnership(unsafe.Pointer(ret.Native())))

	return ret0
}

// GetUnknownGlyph: return the index of a glyph suitable for drawing unknown
// characters with @font, or PANGO_GLYPH_EMPTY if no suitable glyph found.
//
// If you want to draw an unknown-box for a character that is not covered by the
// font, use PANGO_GET_UNKNOWN_GLYPH() instead.
func GetUnknownGlyph(font pango.Font) pango.Glyph {
	var arg1 *C.PangoFont

	arg1 = (*C.PangoFont)(font.Native())

	ret := C.pango_ft2_get_unknown_glyph(arg1)

	var ret0 pango.Glyph

	{
		var tmp uint32
		tmp = uint32(ret)
		ret0 = pango.Glyph(tmp)
	}

	return ret0
}

// Render renders a GlyphString onto a FreeType2 bitmap.
func Render(bitmap *freetype2.Bitmap, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg1 *C.FT_Bitmap
	var arg2 *C.PangoFont
	var arg3 *C.PangoGlyphString
	var arg4 C.gint
	var arg5 C.gint

	arg1 = (*C.FT_Bitmap)(bitmap.Native())
	arg2 = (*C.PangoFont)(font.Native())
	arg3 = (*C.PangoGlyphString)(glyphs.Native())
	arg4 = C.gint(x)
	arg5 = C.gint(y)

	C.pango_ft2_render(arg1, arg2, arg3, arg4, arg5)
}

// RenderLayout: render a Layout onto a FreeType2 bitmap
func RenderLayout(bitmap *freetype2.Bitmap, layout pango.Layout, x int, y int) {
	var arg1 *C.FT_Bitmap
	var arg2 *C.PangoLayout
	var arg3 C.int
	var arg4 C.int

	arg1 = (*C.FT_Bitmap)(bitmap.Native())
	arg2 = (*C.PangoLayout)(layout.Native())
	arg3 = C.int(x)
	arg4 = C.int(y)

	C.pango_ft2_render_layout(arg1, arg2, arg3, arg4)
}

// RenderLayoutLine: render a LayoutLine onto a FreeType2 bitmap
func RenderLayoutLine(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int, y int) {
	var arg1 *C.FT_Bitmap
	var arg2 *C.PangoLayoutLine
	var arg3 C.int
	var arg4 C.int

	arg1 = (*C.FT_Bitmap)(bitmap.Native())
	arg2 = (*C.PangoLayoutLine)(line.Native())
	arg3 = C.int(x)
	arg4 = C.int(y)

	C.pango_ft2_render_layout_line(arg1, arg2, arg3, arg4)
}

// RenderLayoutLineSubpixel: render a LayoutLine onto a FreeType2 bitmap, with
// he location specified in fixed-point Pango units rather than pixels. (Using
// this will avoid extra inaccuracies from rounding to integer pixels multiple
// times, even if the final glyph positions are integers.)
func RenderLayoutLineSubpixel(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int, y int) {
	var arg1 *C.FT_Bitmap
	var arg2 *C.PangoLayoutLine
	var arg3 C.int
	var arg4 C.int

	arg1 = (*C.FT_Bitmap)(bitmap.Native())
	arg2 = (*C.PangoLayoutLine)(line.Native())
	arg3 = C.int(x)
	arg4 = C.int(y)

	C.pango_ft2_render_layout_line_subpixel(arg1, arg2, arg3, arg4)
}

// RenderLayoutSubpixel: render a Layout onto a FreeType2 bitmap, with he
// location specified in fixed-point Pango units rather than pixels. (Using this
// will avoid extra inaccuracies from rounding to integer pixels multiple times,
// even if the final glyph positions are integers.)
func RenderLayoutSubpixel(bitmap *freetype2.Bitmap, layout pango.Layout, x int, y int) {
	var arg1 *C.FT_Bitmap
	var arg2 *C.PangoLayout
	var arg3 C.int
	var arg4 C.int

	arg1 = (*C.FT_Bitmap)(bitmap.Native())
	arg2 = (*C.PangoLayout)(layout.Native())
	arg3 = C.int(x)
	arg4 = C.int(y)

	C.pango_ft2_render_layout_subpixel(arg1, arg2, arg3, arg4)
}

// RenderTransformed renders a GlyphString onto a FreeType2 bitmap, possibly
// transforming the layed-out coordinates through a transformation matrix. Note
// that the transformation matrix for @font is not changed, so to produce
// correct rendering results, the @font must have been loaded using a Context
// with an identical transformation matrix to that passed in to this function.
func RenderTransformed(bitmap *freetype2.Bitmap, matrix *pango.Matrix, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg1 *C.FT_Bitmap
	var arg2 *C.PangoMatrix
	var arg3 *C.PangoFont
	var arg4 *C.PangoGlyphString
	var arg5 C.int
	var arg6 C.int

	arg1 = (*C.FT_Bitmap)(bitmap.Native())
	arg2 = (*C.PangoMatrix)(matrix.Native())
	arg3 = (*C.PangoFont)(font.Native())
	arg4 = (*C.PangoGlyphString)(glyphs.Native())
	arg5 = C.int(x)
	arg6 = C.int(y)

	C.pango_ft2_render_transformed(arg1, arg2, arg3, arg4, arg5, arg6)
}

// ShutdownDisplay: free the global fontmap. (See
// pango_ft2_font_map_for_display()) Use of the global PangoFT2 fontmap is
// deprecated.
func ShutdownDisplay() {

	C.pango_ft2_shutdown_display()
}

// FontMap: the `PangoFT2FontMap` is the `PangoFontMap` implementation for
// FreeType fonts.
type FontMap interface {
	pangofc.FontMap

	// CreateContext: create a `PangoContext` for the given fontmap.
	CreateContext() pango.Context
	// SetDefaultSubstitute sets a function that will be called to do final
	// configuration substitution on a `FcPattern` before it is used to load the
	// font.
	//
	// This function can be used to do things like set hinting and antialiasing
	// options.
	SetDefaultSubstitute(_func SubstituteFunc)
	// SetResolution sets the horizontal and vertical resolutions for the
	// fontmap.
	SetResolution(dpiX float64, dpiY float64)
	// SubstituteChanged: call this function any time the results of the default
	// substitution function set with
	// pango_ft2_font_map_set_default_substitute() change.
	//
	// That is, if your substitution function will return different results for
	// the same input pattern, you must call this function.
	SubstituteChanged()
}

// fontMap implements the FontMap interface.
type fontMap struct {
	pangofc.FontMap
}

// WrapFontMap wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return FontMap{
		pangofc.FontMap: pangofc.WrapFontMap(obj),
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}

// NewFontMap constructs a class FontMap.
func NewFontMap() FontMap {

	ret := C.pango_ft2_font_map_new()

	var ret0 FontMap

	ret0 = WrapFontMap(externglib.AssumeOwnership(unsafe.Pointer(ret.Native())))

	return ret0
}

// CreateContext: create a `PangoContext` for the given fontmap.
func (fontmap fontMap) CreateContext() pango.Context {
	var arg0 *C.PangoFT2FontMap

	arg0 = (*C.PangoFT2FontMap)(fontmap.Native())

	ret := C.pango_ft2_font_map_create_context(arg0)

	var ret0 pango.Context

	ret0 = pango.WrapContext(externglib.AssumeOwnership(unsafe.Pointer(ret.Native())))

	return ret0
}

// SetDefaultSubstitute sets a function that will be called to do final
// configuration substitution on a `FcPattern` before it is used to load the
// font.
//
// This function can be used to do things like set hinting and antialiasing
// options.
func (fontmap fontMap) SetDefaultSubstitute(_func SubstituteFunc) {
	var arg0 *C.PangoFT2FontMap
	var arg1 C.PangoFT2SubstituteFunc
	arg2 := C.gpointer(box.Assign(data))

	arg0 = (*C.PangoFT2FontMap)(fontmap.Native())
	arg1 = (*[0]byte)(C.gotk4_SubstituteFunc)

	C.pango_ft2_font_map_set_default_substitute(arg0, arg1, (*[0]byte)(C.callbackDelete))
}

// SetResolution sets the horizontal and vertical resolutions for the
// fontmap.
func (fontmap fontMap) SetResolution(dpiX float64, dpiY float64) {
	var arg0 *C.PangoFT2FontMap
	var arg1 C.double
	var arg2 C.double

	arg0 = (*C.PangoFT2FontMap)(fontmap.Native())
	arg1 = C.double(dpiX)
	arg2 = C.double(dpiY)

	C.pango_ft2_font_map_set_resolution(arg0, arg1, arg2)
}

// SubstituteChanged: call this function any time the results of the default
// substitution function set with
// pango_ft2_font_map_set_default_substitute() change.
//
// That is, if your substitution function will return different results for
// the same input pattern, you must call this function.
func (fontmap fontMap) SubstituteChanged() {
	var arg0 *C.PangoFT2FontMap

	arg0 = (*C.PangoFT2FontMap)(fontmap.Native())

	C.pango_ft2_font_map_substitute_changed(arg0)
}
