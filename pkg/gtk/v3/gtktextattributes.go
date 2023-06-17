// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeTextAttributes = coreglib.Type(C.gtk_text_attributes_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTextAttributes, F: marshalTextAttributes},
	})
}

// TextAppearance: instance of this type is always passed by reference.
type TextAppearance struct {
	*textAppearance
}

// textAppearance is the struct that's finalized.
type textAppearance struct {
	native *C.GtkTextAppearance
}

// BgColor: background Color.
func (t *TextAppearance) BgColor() *gdk.Color {
	valptr := &t.native.bg_color
	var _v *gdk.Color // out
	_v = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// FgColor: foreground Color.
func (t *TextAppearance) FgColor() *gdk.Color {
	valptr := &t.native.fg_color
	var _v *gdk.Color // out
	_v = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// Rise: super/subscript rise, can be negative.
func (t *TextAppearance) Rise() int {
	valptr := &t.native.rise
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Rise: super/subscript rise, can be negative.
func (t *TextAppearance) SetRise(rise int) {
	valptr := &t.native.rise
	*valptr = C.gint(rise)
}

// TextAttributes: using TextAttributes directly should rarely be necessary.
// It’s primarily useful with gtk_text_iter_get_attributes(). As with most
// GTK+ structs, the fields in this struct should only be read, never modified
// directly.
//
// An instance of this type is always passed by reference.
type TextAttributes struct {
	*textAttributes
}

// textAttributes is the struct that's finalized.
type textAttributes struct {
	native *C.GtkTextAttributes
}

func marshalTextAttributes(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &TextAttributes{&textAttributes{(*C.GtkTextAttributes)(b)}}, nil
}

// NewTextAttributes constructs a struct TextAttributes.
func NewTextAttributes() *TextAttributes {
	var _cret *C.GtkTextAttributes // in

	_cret = C.gtk_text_attributes_new()

	var _textAttributes *TextAttributes // out

	_textAttributes = (*TextAttributes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_textAttributes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_text_attributes_unref((*C.GtkTextAttributes)(intern.C))
		},
	)

	return _textAttributes
}

// Appearance for text.
func (t *TextAttributes) Appearance() *TextAppearance {
	valptr := &t.native.appearance
	var _v *TextAppearance // out
	_v = (*TextAppearance)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// Justification for text.
func (t *TextAttributes) Justification() Justification {
	valptr := &t.native.justification
	var _v Justification // out
	_v = Justification(*valptr)
	return _v
}

// Direction for text.
func (t *TextAttributes) Direction() TextDirection {
	valptr := &t.native.direction
	var _v TextDirection // out
	_v = TextDirection(*valptr)
	return _v
}

// Font for text.
func (t *TextAttributes) Font() *pango.FontDescription {
	valptr := &t.native.font
	var _v *pango.FontDescription // out
	_v = (*pango.FontDescription)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// FontScale: font scale factor.
func (t *TextAttributes) FontScale() float64 {
	valptr := &t.native.font_scale
	var _v float64 // out
	_v = float64(*valptr)
	return _v
}

// LeftMargin: width of the left margin in pixels.
func (t *TextAttributes) LeftMargin() int {
	valptr := &t.native.left_margin
	var _v int // out
	_v = int(*valptr)
	return _v
}

// RightMargin: width of the right margin in pixels.
func (t *TextAttributes) RightMargin() int {
	valptr := &t.native.right_margin
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Indent: amount to indent the paragraph, in pixels.
func (t *TextAttributes) Indent() int {
	valptr := &t.native.indent
	var _v int // out
	_v = int(*valptr)
	return _v
}

// PixelsAboveLines pixels of blank space above paragraphs.
func (t *TextAttributes) PixelsAboveLines() int {
	valptr := &t.native.pixels_above_lines
	var _v int // out
	_v = int(*valptr)
	return _v
}

// PixelsBelowLines pixels of blank space below paragraphs.
func (t *TextAttributes) PixelsBelowLines() int {
	valptr := &t.native.pixels_below_lines
	var _v int // out
	_v = int(*valptr)
	return _v
}

// PixelsInsideWrap pixels of blank space between wrapped lines in a paragraph.
func (t *TextAttributes) PixelsInsideWrap() int {
	valptr := &t.native.pixels_inside_wrap
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Tabs: custom TabArray for this text.
func (t *TextAttributes) Tabs() *pango.TabArray {
	valptr := &t.native.tabs
	var _v *pango.TabArray // out
	_v = (*pango.TabArray)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// WrapMode for text.
func (t *TextAttributes) WrapMode() WrapMode {
	valptr := &t.native.wrap_mode
	var _v WrapMode // out
	_v = WrapMode(*valptr)
	return _v
}

// Language for text.
func (t *TextAttributes) Language() *pango.Language {
	valptr := &t.native.language
	var _v *pango.Language // out
	_v = (*pango.Language)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// LetterSpacing: extra space to insert between graphemes, in Pango units.
func (t *TextAttributes) LetterSpacing() int {
	valptr := &t.native.letter_spacing
	var _v int // out
	_v = int(*valptr)
	return _v
}

// FontScale: font scale factor.
func (t *TextAttributes) SetFontScale(fontScale float64) {
	valptr := &t.native.font_scale
	*valptr = C.gdouble(fontScale)
}

// LeftMargin: width of the left margin in pixels.
func (t *TextAttributes) SetLeftMargin(leftMargin int) {
	valptr := &t.native.left_margin
	*valptr = C.gint(leftMargin)
}

// RightMargin: width of the right margin in pixels.
func (t *TextAttributes) SetRightMargin(rightMargin int) {
	valptr := &t.native.right_margin
	*valptr = C.gint(rightMargin)
}

// Indent: amount to indent the paragraph, in pixels.
func (t *TextAttributes) SetIndent(indent int) {
	valptr := &t.native.indent
	*valptr = C.gint(indent)
}

// PixelsAboveLines pixels of blank space above paragraphs.
func (t *TextAttributes) SetPixelsAboveLines(pixelsAboveLines int) {
	valptr := &t.native.pixels_above_lines
	*valptr = C.gint(pixelsAboveLines)
}

// PixelsBelowLines pixels of blank space below paragraphs.
func (t *TextAttributes) SetPixelsBelowLines(pixelsBelowLines int) {
	valptr := &t.native.pixels_below_lines
	*valptr = C.gint(pixelsBelowLines)
}

// PixelsInsideWrap pixels of blank space between wrapped lines in a paragraph.
func (t *TextAttributes) SetPixelsInsideWrap(pixelsInsideWrap int) {
	valptr := &t.native.pixels_inside_wrap
	*valptr = C.gint(pixelsInsideWrap)
}

// LetterSpacing: extra space to insert between graphemes, in Pango units.
func (t *TextAttributes) SetLetterSpacing(letterSpacing int) {
	valptr := &t.native.letter_spacing
	*valptr = C.gint(letterSpacing)
}

// Copy copies src and returns a new TextAttributes.
//
// The function returns the following values:
//
//   - textAttributes: copy of src, free with gtk_text_attributes_unref().
//
func (src *TextAttributes) Copy() *TextAttributes {
	var _arg0 *C.GtkTextAttributes // out
	var _cret *C.GtkTextAttributes // in

	_arg0 = (*C.GtkTextAttributes)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.gtk_text_attributes_copy(_arg0)
	runtime.KeepAlive(src)

	var _textAttributes *TextAttributes // out

	_textAttributes = (*TextAttributes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_textAttributes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_text_attributes_unref((*C.GtkTextAttributes)(intern.C))
		},
	)

	return _textAttributes
}

// CopyValues copies the values from src to dest so that dest has the same
// values as src. Frees existing values in dest.
//
// The function takes the following parameters:
//
//   - dest: another TextAttributes.
//
func (src *TextAttributes) CopyValues(dest *TextAttributes) {
	var _arg0 *C.GtkTextAttributes // out
	var _arg1 *C.GtkTextAttributes // out

	_arg0 = (*C.GtkTextAttributes)(gextras.StructNative(unsafe.Pointer(src)))
	_arg1 = (*C.GtkTextAttributes)(gextras.StructNative(unsafe.Pointer(dest)))

	C.gtk_text_attributes_copy_values(_arg0, _arg1)
	runtime.KeepAlive(src)
	runtime.KeepAlive(dest)
}
