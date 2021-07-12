// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_button_get_type()), F: marshalFontButtoner},
	})
}

// FontButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FontButtonOverrider interface {
	FontSet()
}

// FontButtoner describes FontButton's methods.
type FontButtoner interface {
	// FontName retrieves the name of the currently selected font.
	FontName() string
	// ShowSize returns whether the font size will be shown in the label.
	ShowSize() bool
	// ShowStyle returns whether the name of the font style will be shown in the
	// label.
	ShowStyle() bool
	// Title retrieves the title of the font chooser dialog.
	Title() string
	// UseFont returns whether the selected font is used in the label.
	UseFont() bool
	// UseSize returns whether the selected size is used in the label.
	UseSize() bool
	// SetFontName sets or updates the currently-displayed font in font picker
	// dialog.
	SetFontName(fontname string) bool
	// SetShowSize: if @show_size is true, the font size will be displayed along
	// with the name of the selected font.
	SetShowSize(showSize bool)
	// SetShowStyle: if @show_style is true, the font style will be displayed
	// along with name of the selected font.
	SetShowStyle(showStyle bool)
	// SetTitle sets the title for the font chooser dialog.
	SetTitle(title string)
	// SetUseFont: if @use_font is true, the font name will be written using the
	// selected font.
	SetUseFont(useFont bool)
	// SetUseSize: if @use_size is true, the font name will be written using the
	// selected size.
	SetUseSize(useSize bool)
}

// FontButton is a button which displays the currently selected font an allows
// to open a font chooser dialog to change the font. It is suitable widget for
// selecting a font in a preference dialog.
//
//
// CSS nodes
//
// GtkFontButton has a single CSS node with name button and style class .font.
type FontButton struct {
	Button

	FontChooser
}

var (
	_ FontButtoner    = (*FontButton)(nil)
	_ gextras.Nativer = (*FontButton)(nil)
)

func wrapFontButton(obj *externglib.Object) FontButtoner {
	return &FontButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
		},
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontButtoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontButton(obj), nil
}

// NewFontButton creates a new font picker widget.
func NewFontButton() *FontButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_button_new()

	var _fontButton *FontButton // out

	_fontButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*FontButton)

	return _fontButton
}

// NewFontButtonWithFont creates a new font picker widget.
func NewFontButtonWithFont(fontname string) *FontButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_button_new_with_font(_arg1)

	var _fontButton *FontButton // out

	_fontButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*FontButton)

	return _fontButton
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *FontButton) Native() uintptr {
	return v.Button.Bin.Container.Widget.InitiallyUnowned.Object.Native()
}

// FontName retrieves the name of the currently selected font. This name
// includes style and size information as well. If you want to render something
// with the font, use this string with pango_font_description_from_string() . If
// you’re interested in peeking certain values (family name, style, size,
// weight) just query these properties from the FontDescription object.
//
// Deprecated: Use gtk_font_chooser_get_font() instead.
func (fontButton *FontButton) FontName() string {
	var _arg0 *C.GtkFontButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_font_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ShowSize returns whether the font size will be shown in the label.
func (fontButton *FontButton) ShowSize() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_show_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowStyle returns whether the name of the font style will be shown in the
// label.
func (fontButton *FontButton) ShowStyle() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_show_style(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the font chooser dialog.
func (fontButton *FontButton) Title() string {
	var _arg0 *C.GtkFontButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseFont returns whether the selected font is used in the label.
func (fontButton *FontButton) UseFont() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_use_font(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseSize returns whether the selected size is used in the label.
func (fontButton *FontButton) UseSize() bool {
	var _arg0 *C.GtkFontButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))

	_cret = C.gtk_font_button_get_use_size(_arg0)

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
func (fontButton *FontButton) SetFontName(fontname string) bool {
	var _arg0 *C.GtkFontButton // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_button_set_font_name(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetShowSize: if @show_size is true, the font size will be displayed along
// with the name of the selected font.
func (fontButton *FontButton) SetShowSize(showSize bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if showSize {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_show_size(_arg0, _arg1)
}

// SetShowStyle: if @show_style is true, the font style will be displayed along
// with name of the selected font.
func (fontButton *FontButton) SetShowStyle(showStyle bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if showStyle {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_show_style(_arg0, _arg1)
}

// SetTitle sets the title for the font chooser dialog.
func (fontButton *FontButton) SetTitle(title string) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_button_set_title(_arg0, _arg1)
}

// SetUseFont: if @use_font is true, the font name will be written using the
// selected font.
func (fontButton *FontButton) SetUseFont(useFont bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if useFont {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_font(_arg0, _arg1)
}

// SetUseSize: if @use_size is true, the font name will be written using the
// selected size.
func (fontButton *FontButton) SetUseSize(useSize bool) {
	var _arg0 *C.GtkFontButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkFontButton)(unsafe.Pointer(fontButton.Native()))
	if useSize {
		_arg1 = C.TRUE
	}

	C.gtk_font_button_set_use_size(_arg0, _arg1)
}
