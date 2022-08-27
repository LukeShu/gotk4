// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewColorButton creates a new color button.
//
// This returns a widget in the form of a small button containing a swatch
// representing the current selected color. When the button is clicked, a
// color-selection dialog will open, allowing the user to select a color. The
// swatch will be updated to reflect the new color when the user finishes.
//
// The function returns the following values:
//
//    - colorButton: new color button.
//
func NewColorButton() *ColorButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_button_new()

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// NewColorButtonWithColor creates a new color button.
//
// Deprecated: Use gtk_color_button_new_with_rgba() instead.
//
// The function takes the following parameters:
//
//    - color to set the current color with.
//
// The function returns the following values:
//
//    - colorButton: new color button.
//
func NewColorButtonWithColor(color *gdk.Color) *ColorButton {
	var _arg1 *C.GdkColor  // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	_cret = C.gtk_color_button_new_with_color(_arg1)
	runtime.KeepAlive(color)

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// Alpha returns the current alpha value.
//
// Deprecated: Use gtk_color_chooser_get_rgba() instead.
//
// The function returns the following values:
//
//    - guint16: integer between 0 and 65535.
//
func (button *ColorButton) Alpha() uint16 {
	var _arg0 *C.GtkColorButton // out
	var _cret C.guint16         // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_color_button_get_alpha(_arg0)
	runtime.KeepAlive(button)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Color sets color to be the current color in the ColorButton widget.
//
// Deprecated: Use gtk_color_chooser_get_rgba() instead.
//
// The function returns the following values:
//
//    - color to fill in with the current color.
//
func (button *ColorButton) Color() *gdk.Color {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.GdkColor        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	C.gtk_color_button_get_color(_arg0, &_arg1)
	runtime.KeepAlive(button)

	var _color *gdk.Color // out

	_color = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// Title gets the title of the color selection dialog.
//
// The function returns the following values:
//
//    - utf8: internal string, do not free the return value.
//
func (button *ColorButton) Title() string {
	var _arg0 *C.GtkColorButton // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_color_button_get_title(_arg0)
	runtime.KeepAlive(button)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseAlpha does the color selection dialog use the alpha channel ?
//
// Deprecated: Use gtk_color_chooser_get_use_alpha() instead.
//
// The function returns the following values:
//
//    - ok: TRUE if the color sample uses alpha channel, FALSE if not.
//
func (button *ColorButton) UseAlpha() bool {
	var _arg0 *C.GtkColorButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_color_button_get_use_alpha(_arg0)
	runtime.KeepAlive(button)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAlpha sets the current opacity to be alpha.
//
// Deprecated: Use gtk_color_chooser_set_rgba() instead.
//
// The function takes the following parameters:
//
//    - alpha: integer between 0 and 65535.
//
func (button *ColorButton) SetAlpha(alpha uint16) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.guint16         // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_button_set_alpha(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(alpha)
}

// SetColor sets the current color to be color.
//
// Deprecated: Use gtk_color_chooser_set_rgba() instead.
//
// The function takes the following parameters:
//
//    - color to set the current color with.
//
func (button *ColorButton) SetColor(color *gdk.Color) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.GdkColor       // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_color_button_set_color(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(color)
}

// SetTitle sets the title for the color selection dialog.
//
// The function takes the following parameters:
//
//    - title: string containing new window title.
//
func (button *ColorButton) SetTitle(title string) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_color_button_set_title(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(title)
}

// SetUseAlpha sets whether or not the color button should use the alpha
// channel.
//
// Deprecated: Use gtk_color_chooser_set_use_alpha() instead.
//
// The function takes the following parameters:
//
//    - useAlpha: TRUE if color button should use alpha channel, FALSE if not.
//
func (button *ColorButton) SetUseAlpha(useAlpha bool) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_button_set_use_alpha(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(useAlpha)
}
