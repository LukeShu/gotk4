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
		{T: externglib.Type(C.gtk_app_chooser_button_get_type()), F: marshalAppChooserButton},
	})
}

// AppChooserButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AppChooserButtonOverrider interface {
	CustomItemActivated(itemName string)
}

// AppChooserButton: the AppChooserButton is a widget that lets the user select
// an application. It implements the AppChooser interface.
//
// Initially, a AppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// AppChooserButton:show-default-item is true, the default application.
//
// The list of applications shown in a AppChooserButton includes the recommended
// applications for the given content type. When
// AppChooserButton:show-default-item is set, the default application is also
// included. To let the user chooser other applications, you can set the
// AppChooserButton:show-dialog-item property, which allows to open a full
// AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk_app_chooser_button_append_custom_item(). These items cause the
// AppChooserButton::custom-item-activated signal to be emitted when they are
// selected.
//
// To track changes in the selected application, use the ComboBox::changed
// signal.
type AppChooserButton interface {
	gextras.Objector

	// AppendSeparator appends a separator to the list of applications that is
	// shown in the popup.
	AppendSeparator()
	// Heading returns the text to display at the top of the dialog.
	Heading() string
	// ShowDefaultItem returns the current value of the
	// AppChooserButton:show-default-item property.
	ShowDefaultItem() bool
	// ShowDialogItem returns the current value of the
	// AppChooserButton:show-dialog-item property.
	ShowDialogItem() bool
	// SetActiveCustomItem selects a custom item previously added with
	// gtk_app_chooser_button_append_custom_item().
	//
	// Use gtk_app_chooser_refresh() to bring the selection to its initial
	// state.
	SetActiveCustomItem(name string)
	// SetHeading sets the text to display at the top of the dialog. If the
	// heading is not set, the dialog displays a default text.
	SetHeading(heading string)
	// SetShowDefaultItem sets whether the dropdown menu of this button should
	// show the default application for the given content type at top.
	SetShowDefaultItem(setting bool)
	// SetShowDialogItem sets whether the dropdown menu of this button should
	// show an entry to trigger a AppChooserDialog.
	SetShowDialogItem(setting bool)
}

// AppChooserButtonClass implements the AppChooserButton interface.
type AppChooserButtonClass struct {
	*externglib.Object
	ComboBoxClass
	AppChooserInterface
	BuildableInterface
	CellEditableInterface
	CellLayoutInterface
}

var _ AppChooserButton = (*AppChooserButtonClass)(nil)

func wrapAppChooserButton(obj *externglib.Object) AppChooserButton {
	return &AppChooserButtonClass{
		Object: obj,
		ComboBoxClass: ComboBoxClass{
			Object: obj,
			BinClass: BinClass{
				Object: obj,
				ContainerClass: ContainerClass{
					Object: obj,
					WidgetClass: WidgetClass{
						Object:           obj,
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
			CellEditableInterface: CellEditableInterface{
				WidgetClass: WidgetClass{
					Object:           obj,
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
			},
			CellLayoutInterface: CellLayoutInterface{
				Object: obj,
			},
		},
		AppChooserInterface: AppChooserInterface{
			WidgetClass: WidgetClass{
				Object:           obj,
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		CellEditableInterface: CellEditableInterface{
			WidgetClass: WidgetClass{
				Object:           obj,
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
		},
		CellLayoutInterface: CellLayoutInterface{
			Object: obj,
		},
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppChooserButton(obj), nil
}

// NewAppChooserButton creates a new AppChooserButton for applications that can
// handle content of the given type.
func NewAppChooserButton(contentType string) AppChooserButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_app_chooser_button_new(_arg1)

	var _appChooserButton AppChooserButton // out

	_appChooserButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(AppChooserButton)

	return _appChooserButton
}

// AppendSeparator appends a separator to the list of applications that is shown
// in the popup.
func (s *AppChooserButtonClass) AppendSeparator() {
	var _arg0 *C.GtkAppChooserButton // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_button_append_separator(_arg0)
}

// Heading returns the text to display at the top of the dialog.
func (s *AppChooserButtonClass) Heading() string {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_heading(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ShowDefaultItem returns the current value of the
// AppChooserButton:show-default-item property.
func (s *AppChooserButtonClass) ShowDefaultItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_show_default_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDialogItem returns the current value of the
// AppChooserButton:show-dialog-item property.
func (s *AppChooserButtonClass) ShowDialogItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_show_dialog_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveCustomItem selects a custom item previously added with
// gtk_app_chooser_button_append_custom_item().
//
// Use gtk_app_chooser_refresh() to bring the selection to its initial state.
func (s *AppChooserButtonClass) SetActiveCustomItem(name string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_active_custom_item(_arg0, _arg1)
}

// SetHeading sets the text to display at the top of the dialog. If the heading
// is not set, the dialog displays a default text.
func (s *AppChooserButtonClass) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_heading(_arg0, _arg1)
}

// SetShowDefaultItem sets whether the dropdown menu of this button should show
// the default application for the given content type at top.
func (s *AppChooserButtonClass) SetShowDefaultItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_default_item(_arg0, _arg1)
}

// SetShowDialogItem sets whether the dropdown menu of this button should show
// an entry to trigger a AppChooserDialog.
func (s *AppChooserButtonClass) SetShowDialogItem(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_dialog_item(_arg0, _arg1)
}
