// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// FontMap gets the font map that has been set with gtk_widget_set_font_map().
//
// The function returns the following values:
//
//    - fontMap (optional) or NULL.
//
func (widget *Widget) FontMap() pango.FontMapper {
	var _arg0 *C.GtkWidget    // out
	var _cret *C.PangoFontMap // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_widget_get_font_map(_arg0)
	runtime.KeepAlive(widget)

	var _fontMap pango.FontMapper // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontMapper)
				return ok
			})
			rv, ok := casted.(pango.FontMapper)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontMapper")
			}
			_fontMap = rv
		}
	}

	return _fontMap
}

// FontOptions returns the #cairo_font_options_t used for Pango rendering. When
// not set, the defaults font options for the Screen will be used.
//
// The function returns the following values:
//
//    - fontOptions (optional) or NULL if not set.
//
func (widget *Widget) FontOptions() *cairo.FontOptions {
	var _arg0 *C.GtkWidget            // out
	var _cret *C.cairo_font_options_t // in

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.gtk_widget_get_font_options(_arg0)
	runtime.KeepAlive(widget)

	var _fontOptions *cairo.FontOptions // out

	if _cret != nil {
		_fontOptions = (*cairo.FontOptions)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _fontOptions
}

// SetFontMap sets the font map to use for Pango rendering. When not set, the
// widget will inherit the font map from its parent.
//
// The function takes the following parameters:
//
//    - fontMap (optional) or NULL to unset any previously set font map.
//
func (widget *Widget) SetFontMap(fontMap pango.FontMapper) {
	var _arg0 *C.GtkWidget    // out
	var _arg1 *C.PangoFontMap // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if fontMap != nil {
		_arg1 = (*C.PangoFontMap)(unsafe.Pointer(coreglib.InternObject(fontMap).Native()))
	}

	C.gtk_widget_set_font_map(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(fontMap)
}

// SetFontOptions sets the #cairo_font_options_t used for Pango rendering in
// this widget. When not set, the default font options for the Screen will be
// used.
//
// The function takes the following parameters:
//
//    - options (optional) or NULL to unset any previously set default font
//      options.
//
func (widget *Widget) SetFontOptions(options *cairo.FontOptions) {
	var _arg0 *C.GtkWidget            // out
	var _arg1 *C.cairo_font_options_t // out

	_arg0 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if options != nil {
		_arg1 = (*C.cairo_font_options_t)(gextras.StructNative(unsafe.Pointer(options)))
	}

	C.gtk_widget_set_font_options(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(options)
}
