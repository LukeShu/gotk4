// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewFontButton creates a new font picker widget.
//
// The function returns the following values:
//
//    - fontButton: new font picker widget.
//
func NewFontButton() *FontButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_button_new()

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// NewFontButtonWithFont creates a new font picker widget.
//
// The function takes the following parameters:
//
//    - fontname: name of font to display in font chooser dialog.
//
// The function returns the following values:
//
//    - fontButton: new font picker widget.
//
func NewFontButtonWithFont(fontname string) *FontButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_button_new_with_font(_arg1)
	runtime.KeepAlive(fontname)

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// FontName retrieves the name of the currently selected font. This name
// includes style and size information as well. If you want to render something
// with the font, use this string with pango_font_description_from_string() . If
// you’re interested in peeking certain values (family name, style, size,
// weight) just query these properties from the FontDescription object.
//
// Deprecated: Use gtk_font_chooser_get_font() instead.
//
// The function returns the following values:
//
//    - utf8: internal copy of the font name which must not be freed.
//
func (fontButton *FontButton) FontName() string {
	var _arg0 *C.GtkFontButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_cret = C.gtk_font_button_get_font_name(_arg0)
	runtime.KeepAlive(fontButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ShowSize returns whether the font size will be shown in the label.
//
// The function returns the following values:
//
//    - ok: whether the font size will be shown in the label.
//
func (fontButton *FontButton) ShowSize() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_cret = C.gtk_font_button_get_show_size(_arg0)
	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowStyle returns whether the name of the font style will be shown in the
// label.
//
// The function returns the following values:
//
//    - ok: whether the font style will be shown in the label.
//
func (fontButton *FontButton) ShowStyle() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_cret = C.gtk_font_button_get_show_style(_arg0)
	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the font chooser dialog.
//
// The function returns the following values:
//
//    - utf8: internal copy of the title string which must not be freed.
//
func (fontButton *FontButton) Title() string {
	var _arg0 *C.GtkFontButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_cret = C.gtk_font_button_get_title(_arg0)
	runtime.KeepAlive(fontButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseFont returns whether the selected font is used in the label.
//
// The function returns the following values:
//
//    - ok: whether the selected font is used in the label.
//
func (fontButton *FontButton) UseFont() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_cret = C.gtk_font_button_get_use_font(_arg0)
	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseSize returns whether the selected size is used in the label.
//
// The function returns the following values:
//
//    - ok: whether the selected size is used in the label.
//
func (fontButton *FontButton) UseSize() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_cret = C.gtk_font_button_get_use_size(_arg0)
	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFontName sets or updates the currently-displayed font in font picker
// dialog.
//
// Deprecated: Use gtk_font_chooser_set_font() instead.
//
// The function takes the following parameters:
//
//    - fontname: name of font to display in font chooser dialog.
//
// The function returns the following values:
//
//    - ok: TRUE.
//
func (fontButton *FontButton) SetFontName(fontname string) bool {
	var _arg0 *C.GtkFontButton // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_button_set_font_name(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(fontname)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetShowSize: if show_size is TRUE, the font size will be displayed along with
// the name of the selected font.
//
// The function takes the following parameters:
//
//    - showSize: TRUE if font size should be displayed in dialog.
//
func (fontButton *FontButton) SetShowSize(showSize bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	if showSize {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_show_size(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(showSize)
}

// SetShowStyle: if show_style is TRUE, the font style will be displayed along
// with name of the selected font.
//
// The function takes the following parameters:
//
//    - showStyle: TRUE if font style should be displayed in label.
//
func (fontButton *FontButton) SetShowStyle(showStyle bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	if showStyle {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_show_style(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(showStyle)
}

// SetTitle sets the title for the font chooser dialog.
//
// The function takes the following parameters:
//
//    - title: string containing the font chooser dialog title.
//
func (fontButton *FontButton) SetTitle(title string) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_button_set_title(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(title)
}

// SetUseFont: if use_font is TRUE, the font name will be written using the
// selected font.
//
// The function takes the following parameters:
//
//    - useFont: if TRUE, font name will be written using font chosen.
//
func (fontButton *FontButton) SetUseFont(useFont bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	if useFont {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_font(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(useFont)
}

// SetUseSize: if use_size is TRUE, the font name will be written using the
// selected size.
//
// The function takes the following parameters:
//
//    - useSize: if TRUE, font name will be written using the selected size.
//
func (fontButton *FontButton) SetUseSize(useSize bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	if useSize {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_size(_arg0, _arg1)
	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(useSize)
}
