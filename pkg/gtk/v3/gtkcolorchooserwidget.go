// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeColorChooserWidget = coreglib.Type(C.gtk_color_chooser_widget_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorChooserWidget, F: marshalColorChooserWidget},
	})
}

// ColorChooserWidgetOverrider contains methods that are overridable.
type ColorChooserWidgetOverrider interface {
}

// ColorChooserWidget widget lets the user select a color. By default, the
// chooser presents a predefined palette of colors, plus a small number of
// settable custom colors. It is also possible to select a different color with
// the single-color editor. To enter the single-color editing mode, use the
// context menu of any color of the palette, or use the '+' button to add a new
// custom color.
//
// The chooser automatically remembers the last selection, as well as custom
// colors.
//
// To change the initially selected color, use gtk_color_chooser_set_rgba(). To
// get the selected color use gtk_color_chooser_get_rgba().
//
// The ColorChooserWidget is used in the ColorChooserDialog to provide a dialog
// for selecting colors.
//
//
// CSS names
//
// GtkColorChooserWidget has a single CSS node with name colorchooser.
type ColorChooserWidget struct {
	_ [0]func() // equal guard
	Box

	*coreglib.Object
	ColorChooser
}

var (
	_ coreglib.Objector = (*ColorChooserWidget)(nil)
	_ Containerer       = (*ColorChooserWidget)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeColorChooserWidget,
		GoType:        reflect.TypeOf((*ColorChooserWidget)(nil)),
		InitClass:     initClassColorChooserWidget,
		FinalizeClass: finalizeClassColorChooserWidget,
	})
}

func initClassColorChooserWidget(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitColorChooserWidget(*ColorChooserWidgetClass)
	}); ok {
		klass := (*ColorChooserWidgetClass)(gextras.NewStructNative(gclass))
		goval.InitColorChooserWidget(klass)
	}
}

func finalizeClassColorChooserWidget(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		FinalizeColorChooserWidget(*ColorChooserWidgetClass)
	}); ok {
		klass := (*ColorChooserWidgetClass)(gextras.NewStructNative(gclass))
		goval.FinalizeColorChooserWidget(klass)
	}
}

func wrapColorChooserWidget(obj *coreglib.Object) *ColorChooserWidget {
	return &ColorChooserWidget{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
		Object: obj,
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserWidget(p uintptr) (interface{}, error) {
	return wrapColorChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserWidget creates a new ColorChooserWidget.
//
// The function returns the following values:
//
//    - colorChooserWidget: new ColorChooserWidget.
//
func NewColorChooserWidget() *ColorChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_chooser_widget_new()

	var _colorChooserWidget *ColorChooserWidget // out

	_colorChooserWidget = wrapColorChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserWidget
}

// ColorChooserWidgetClass: instance of this type is always passed by reference.
type ColorChooserWidgetClass struct {
	*colorChooserWidgetClass
}

// colorChooserWidgetClass is the struct that's finalized.
type colorChooserWidgetClass struct {
	native *C.GtkColorChooserWidgetClass
}

// ParentClass: parent class.
func (c *ColorChooserWidgetClass) ParentClass() *BoxClass {
	valptr := &c.native.parent_class
	var _v *BoxClass // out
	_v = (*BoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
