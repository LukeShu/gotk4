// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_selection_dialog_get_type()), F: marshalColorSelectionDialog},
	})
}

type ColorSelectionDialog interface {
	gextras.Objector

	// ColorSelection retrieves the ColorSelection widget embedded in the
	// dialog.
	ColorSelection() *WidgetClass
}

// ColorSelectionDialogClass implements the ColorSelectionDialog interface.
type ColorSelectionDialogClass struct {
	*externglib.Object
	DialogClass
	BuildableInterface
}

var _ ColorSelectionDialog = (*ColorSelectionDialogClass)(nil)

func wrapColorSelectionDialog(obj *externglib.Object) ColorSelectionDialog {
	return &ColorSelectionDialogClass{
		Object: obj,
		DialogClass: DialogClass{
			Object: obj,
			WindowClass: WindowClass{
				Object: obj,
				BinClass: BinClass{
					Object: obj,
					ContainerClass: ContainerClass{
						Object: obj,
						WidgetClass: WidgetClass{
							InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
							BuildableInterface: BuildableInterface{
								Object: obj,
							},
						},
						BuildableInterface: BuildableInterface{
							Object: obj,
						},
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalColorSelectionDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorSelectionDialog(obj), nil
}

// NewColorSelectionDialog creates a new ColorSelectionDialog.
func NewColorSelectionDialog(title string) *ColorSelectionDialogClass {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_color_selection_dialog_new(_arg1)

	var _colorSelectionDialog *ColorSelectionDialogClass // out

	_colorSelectionDialog = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ColorSelectionDialogClass)

	return _colorSelectionDialog
}

// ColorSelection retrieves the ColorSelection widget embedded in the dialog.
func (c *ColorSelectionDialogClass) ColorSelection() *WidgetClass {
	var _arg0 *C.GtkColorSelectionDialog // out
	var _cret *C.GtkWidget               // in

	_arg0 = (*C.GtkColorSelectionDialog)(unsafe.Pointer((&c).Native()))

	_cret = C.gtk_color_selection_dialog_get_color_selection(_arg0)

	var _widget *WidgetClass // out

	_widget = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*WidgetClass)

	return _widget
}
