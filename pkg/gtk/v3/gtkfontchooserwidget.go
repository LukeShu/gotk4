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
		{T: externglib.Type(C.gtk_font_chooser_widget_get_type()), F: marshalFontChooserWidgeter},
	})
}

// FontChooserWidgeter describes FontChooserWidget's methods.
type FontChooserWidgeter interface {
	privateFontChooserWidget()
}

// FontChooserWidget widget lists the available fonts, styles and sizes,
// allowing the user to select a font. It is used in the FontChooserDialog
// widget to provide a dialog box for selecting fonts.
//
// To set the font which is initially selected, use gtk_font_chooser_set_font()
// or gtk_font_chooser_set_font_desc().
//
// To get the selected font use gtk_font_chooser_get_font() or
// gtk_font_chooser_get_font_desc().
//
// To change the text which is shown in the preview area, use
// gtk_font_chooser_set_preview_text().
//
//
// CSS nodes
//
// GtkFontChooserWidget has a single CSS node with name fontchooser.
type FontChooserWidget struct {
	Box

	FontChooser
}

var (
	_ FontChooserWidgeter = (*FontChooserWidget)(nil)
	_ gextras.Nativer     = (*FontChooserWidget)(nil)
)

func wrapFontChooserWidget(obj *externglib.Object) FontChooserWidgeter {
	return &FontChooserWidget{
		Box: Box{
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
			Orientable: Orientable{
				Object: obj,
			},
		},
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserWidgeter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontChooserWidget(obj), nil
}

// NewFontChooserWidget creates a new FontChooserWidget.
func NewFontChooserWidget() *FontChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_chooser_widget_new()

	var _fontChooserWidget *FontChooserWidget // out

	_fontChooserWidget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*FontChooserWidget)

	return _fontChooserWidget
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *FontChooserWidget) Native() uintptr {
	return v.Box.Container.Widget.InitiallyUnowned.Object.Native()
}

func (*FontChooserWidget) privateFontChooserWidget() {}
