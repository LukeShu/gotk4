// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// ShapeFull: convert the characters in text into glyphs.
//
// Given a segment of text and the corresponding PangoAnalysis structure
// returned from itemize, convert the characters into glyphs. You may also pass
// in only a substring of the item from itemize.
//
// This is similar to shape, except it also can optionally take the full
// paragraph text as input, which will then be used to perform certain
// cross-item shaping interactions. If you have access to the broader text of
// which item_text is part of, provide the broader text as paragraph_text. If
// paragraph_text is NULL, item text is used instead.
//
// Note that the extra attributes in the analyis that is returned from itemize
// have indices that are relative to the entire paragraph, so you do not pass
// the full paragraph text as paragraph_text, you need to subtract the item
// offset from their indices before calling shape_full.
//
// The function takes the following parameters:
//
//    - itemText: valid UTF-8 text to shape.
//    - itemLength: length (in bytes) of item_text. -1 means nul-terminated text.
//    - paragraphText (optional): text of the paragraph (see details). May be
//      NULL.
//    - paragraphLength: length (in bytes) of paragraph_text. -1 means
//      nul-terminated text.
//    - analysis: PangoAnalysis structure from itemize.
//    - glyphs: glyph string in which to store results.
//
func ShapeFull(itemText string, itemLength int, paragraphText string, paragraphLength int, analysis *Analysis, glyphs *GlyphString) {
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.char             // out
	var _arg4 C.int               // out
	var _arg5 *C.PangoAnalysis    // out
	var _arg6 *C.PangoGlyphString // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(itemText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(itemLength)
	if paragraphText != "" {
		_arg3 = (*C.char)(unsafe.Pointer(C.CString(paragraphText)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = C.int(paragraphLength)
	_arg5 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg6 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_shape_full(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(itemText)
	runtime.KeepAlive(itemLength)
	runtime.KeepAlive(paragraphText)
	runtime.KeepAlive(paragraphLength)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(glyphs)
}
