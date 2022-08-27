// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// MarkupParserFinish finishes parsing markup.
//
// After feeding a Pango markup parser some data with
// g_markup_parse_context_parse(), use this function to get the list of
// attributes and text out of the markup. This function will not free context,
// use g_markup_parse_context_free() to do so.
//
// The function takes the following parameters:
//
//    - context: valid parse context that was returned from markup_parser_new.
//
// The function returns the following values:
//
//    - attrList (optional) address of return location for a PangoAttrList, or
//      NULL.
//    - text (optional) address of return location for text with tags stripped,
//      or NULL.
//    - accelChar (optional) address of return location for accelerator char, or
//      NULL.
//
func MarkupParserFinish(context *glib.MarkupParseContext) (*AttrList, string, uint32, error) {
	var _arg1 *C.GMarkupParseContext // out
	var _arg2 *C.PangoAttrList       // in
	var _arg3 *C.char                // in
	var _arg4 C.gunichar             // in
	var _cerr *C.GError              // in

	_arg1 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	C.pango_markup_parser_finish(_arg1, &_arg2, &_arg3, &_arg4, &_cerr)
	runtime.KeepAlive(context)

	var _attrList *AttrList // out
	var _text string        // out
	var _accelChar uint32   // out
	var _goerr error        // out

	if _arg2 != nil {
		_attrList = (*AttrList)(gextras.NewStructNative(unsafe.Pointer(_arg2)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_attrList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_attr_list_unref((*C.PangoAttrList)(intern.C))
			},
		)
	}
	if _arg3 != nil {
		_text = C.GoString((*C.gchar)(unsafe.Pointer(_arg3)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_accelChar = uint32(_arg4)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _attrList, _text, _accelChar, _goerr
}

// NewMarkupParser: incrementally parses marked-up text to create a plain-text
// string and an attribute list.
//
// See the Pango Markup (pango_markup.html) docs for details about the supported
// markup.
//
// If accel_marker is nonzero, the given character will mark the character
// following it as an accelerator. For example, accel_marker might be an
// ampersand or underscore. All characters marked as an accelerator will receive
// a PANGO_UNDERLINE_LOW attribute, and the first character so marked will be
// returned in accel_char, when calling markup_parser_finish. Two accel_marker
// characters following each other produce a single literal accel_marker
// character.
//
// To feed markup to the parser, use g_markup_parse_context_parse() on the
// returned GMarkupParseContext. When done with feeding markup to the parser,
// use markup_parser_finish to get the data out of it, and then use
// g_markup_parse_context_free() to free it.
//
// This function is designed for applications that read Pango markup from
// streams. To simply parse a string containing Pango markup, the parse_markup
// API is recommended instead.
//
// The function takes the following parameters:
//
//    - accelMarker: character that precedes an accelerator, or 0 for none.
//
// The function returns the following values:
//
//    - markupParseContext: GMarkupParseContext that should be destroyed with
//      g_markup_parse_context_free().
//
func NewMarkupParser(accelMarker uint32) *glib.MarkupParseContext {
	var _arg1 C.gunichar             // out
	var _cret *C.GMarkupParseContext // in

	_arg1 = C.gunichar(accelMarker)

	_cret = C.pango_markup_parser_new(_arg1)
	runtime.KeepAlive(accelMarker)

	var _markupParseContext *glib.MarkupParseContext // out

	_markupParseContext = (*glib.MarkupParseContext)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_markup_parse_context_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_markupParseContext)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_markup_parse_context_unref((*C.GMarkupParseContext)(intern.C))
		},
	)

	return _markupParseContext
}
