// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeColorChooserDialog returns the GType for the type ColorChooserDialog.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeColorChooserDialog() coreglib.Type {
	gtype := coreglib.Type(C.gtk_color_chooser_dialog_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalColorChooserDialog)
	return gtype
}

// ColorChooserDialogOverrider contains methods that are overridable.
type ColorChooserDialogOverrider interface {
}

// ColorChooserDialog widget is a dialog for choosing a color. It implements the
// ColorChooser interface.
type ColorChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	ColorChooser
}

var (
	_ coreglib.Objector = (*ColorChooserDialog)(nil)
	_ Binner            = (*ColorChooserDialog)(nil)
)

func classInitColorChooserDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapColorChooserDialog(obj *coreglib.Object) *ColorChooserDialog {
	return &ColorChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
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
				},
			},
		},
		Object: obj,
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	return wrapColorChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserDialog creates a new ColorChooserDialog.
//
// The function takes the following parameters:
//
//    - title (optional): title of the dialog, or NULL.
//    - parent (optional): transient parent of the dialog, or NULL.
//
// The function returns the following values:
//
//    - colorChooserDialog: new ColorChooserDialog.
//
func NewColorChooserDialog(title string, parent *Window) *ColorChooserDialog {
	var _arg1 *C.gchar     // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if parent != nil {
		_arg2 = (*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_cret = C.gtk_color_chooser_dialog_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)

	var _colorChooserDialog *ColorChooserDialog // out

	_colorChooserDialog = wrapColorChooserDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserDialog
}
