// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_dialog_get_type()), F: marshalAppChooserDialogger},
	})
}

// AppChooserDialog shows a AppChooserWidget inside a Dialog.
//
// Note that AppChooserDialog does not have any interesting methods of its own.
// Instead, you should get the embedded AppChooserWidget using
// gtk_app_chooser_dialog_get_widget() and call its methods if the generic
// AppChooser interface is not sufficient for your needs.
//
// To set the heading that is shown above the AppChooserWidget, use
// gtk_app_chooser_dialog_set_heading().
type AppChooserDialog struct {
	Dialog

	*externglib.Object
	AppChooser
}

var (
	_ externglib.Objector = (*AppChooserDialog)(nil)
	_ Binner              = (*AppChooserDialog)(nil)
)

func wrapAppChooserDialog(obj *externglib.Object) *AppChooserDialog {
	return &AppChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
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
		AppChooser: AppChooser{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalAppChooserDialogger(p uintptr) (interface{}, error) {
	return wrapAppChooserDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAppChooserDialog creates a new AppChooserDialog for the provided #GFile,
// to allow the user to select an application for it.
//
// The function takes the following parameters:
//
//    - parent or NULL.
//    - flags for this dialog.
//    - file: #GFile.
//
func NewAppChooserDialog(parent *Window, flags DialogFlags, file gio.Filer) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.GFile         // out
	var _cret *C.GtkWidget     // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.GFile)(unsafe.Pointer(file.Native()))

	_cret = C.gtk_app_chooser_dialog_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(file)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserDialog
}

// NewAppChooserDialogForContentType creates a new AppChooserDialog for the
// provided content type, to allow the user to select an application for it.
//
// The function takes the following parameters:
//
//    - parent or NULL.
//    - flags for this dialog.
//    - contentType: content type string.
//
func NewAppChooserDialogForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	var _arg1 *C.GtkWindow     // out
	var _arg2 C.GtkDialogFlags // out
	var _arg3 *C.gchar         // out
	var _cret *C.GtkWidget     // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}
	_arg2 = C.GtkDialogFlags(flags)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_app_chooser_dialog_new_for_content_type(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(contentType)

	var _appChooserDialog *AppChooserDialog // out

	_appChooserDialog = wrapAppChooserDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserDialog
}

// Heading returns the text to display at the top of the dialog.
func (self *AppChooserDialog) Heading() string {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_dialog_get_heading(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Widget returns the AppChooserWidget of this dialog.
func (self *AppChooserDialog) Widget() Widgetter {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_dialog_get_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SetHeading sets the text to display at the top of the dialog. If the heading
// is not set, the dialog displays a default text.
//
// The function takes the following parameters:
//
//    - heading: string containing Pango markup.
//
func (self *AppChooserDialog) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserDialog // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(heading)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_dialog_set_heading(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(heading)
}
