// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_widget_get_type()), F: marshalAppChooserWidget},
	})
}

// AppChooserWidget: `GtkAppChooserWidget` is a widget for selecting
// applications.
//
// It is the main building block for [class@Gtk.AppChooserDialog]. Most
// applications only need to use the latter; but you can use this widget as part
// of a larger widget if you have special needs.
//
// `GtkAppChooserWidget` offers detailed control over what applications are
// shown, using the [property@Gtk.AppChooserWidget:show-default],
// [property@Gtk.AppChooserWidget:show-recommended],
// [property@Gtk.AppChooserWidget:show-fallback],
// [property@Gtk.AppChooserWidget:show-other] and
// [property@Gtk.AppChooserWidget:show-all] properties. See the
// [iface@Gtk.AppChooser] documentation for more information about these groups
// of applications.
//
// To keep track of the selected application, use the
// [signal@Gtk.AppChooserWidget::application-selected] and
// [signal@Gtk.AppChooserWidget::application-activated] signals.
//
//
// CSS nodes
//
// `GtkAppChooserWidget` has a single CSS node with name appchooser.
type AppChooserWidget interface {
	Widget
	Accessible
	AppChooser
	Buildable
	ConstraintTarget

	// DefaultText returns the text that is shown if there are not applications
	// that can handle the content type.
	DefaultText() string
	// ShowAll gets whether the app chooser should show all applications in a
	// flat list.
	ShowAll() bool
	// ShowDefault gets whether the app chooser should show the default handler
	// for the content type in a separate section.
	ShowDefault() bool
	// ShowFallback gets whether the app chooser should show related
	// applications for the content type in a separate section.
	ShowFallback() bool
	// ShowOther gets whether the app chooser should show applications which are
	// unrelated to the content type.
	ShowOther() bool
	// ShowRecommended gets whether the app chooser should show recommended
	// applications for the content type in a separate section.
	ShowRecommended() bool
	// SetDefaultText sets the text that is shown if there are not applications
	// that can handle the content type.
	SetDefaultText(text string)
	// SetShowAll sets whether the app chooser should show all applications in a
	// flat list.
	SetShowAll(setting bool)
	// SetShowDefault sets whether the app chooser should show the default
	// handler for the content type in a separate section.
	SetShowDefault(setting bool)
	// SetShowFallback sets whether the app chooser should show related
	// applications for the content type in a separate section.
	SetShowFallback(setting bool)
	// SetShowOther sets whether the app chooser should show applications which
	// are unrelated to the content type.
	SetShowOther(setting bool)
	// SetShowRecommended sets whether the app chooser should show recommended
	// applications for the content type in a separate section.
	SetShowRecommended(setting bool)
}

// appChooserWidget implements the AppChooserWidget class.
type appChooserWidget struct {
	Widget
	Accessible
	AppChooser
	Buildable
	ConstraintTarget
}

var _ AppChooserWidget = (*appChooserWidget)(nil)

// WrapAppChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserWidget(obj *externglib.Object) AppChooserWidget {
	return appChooserWidget{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		AppChooser:       WrapAppChooser(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalAppChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserWidget(obj), nil
}

// NewAppChooserWidget constructs a class AppChooserWidget.
func NewAppChooserWidget(contentType string) AppChooserWidget {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkAppChooserWidget // in

	_cret = C.gtk_app_chooser_widget_new(_arg1)

	var _appChooserWidget AppChooserWidget // out

	_appChooserWidget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(AppChooserWidget)

	return _appChooserWidget
}

// DefaultText returns the text that is shown if there are not applications
// that can handle the content type.
func (s appChooserWidget) DefaultText() string {
	var _arg0 *C.GtkAppChooserWidget // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var _cret *C.char // in

	_cret = C.gtk_app_chooser_widget_get_default_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ShowAll gets whether the app chooser should show all applications in a
// flat list.
func (s appChooserWidget) ShowAll() bool {
	var _arg0 *C.GtkAppChooserWidget // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_app_chooser_widget_get_show_all(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefault gets whether the app chooser should show the default handler
// for the content type in a separate section.
func (s appChooserWidget) ShowDefault() bool {
	var _arg0 *C.GtkAppChooserWidget // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_app_chooser_widget_get_show_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowFallback gets whether the app chooser should show related
// applications for the content type in a separate section.
func (s appChooserWidget) ShowFallback() bool {
	var _arg0 *C.GtkAppChooserWidget // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_app_chooser_widget_get_show_fallback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowOther gets whether the app chooser should show applications which are
// unrelated to the content type.
func (s appChooserWidget) ShowOther() bool {
	var _arg0 *C.GtkAppChooserWidget // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_app_chooser_widget_get_show_other(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRecommended gets whether the app chooser should show recommended
// applications for the content type in a separate section.
func (s appChooserWidget) ShowRecommended() bool {
	var _arg0 *C.GtkAppChooserWidget // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_app_chooser_widget_get_show_recommended(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultText sets the text that is shown if there are not applications
// that can handle the content type.
func (s appChooserWidget) SetDefaultText(text string) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_widget_set_default_text(_arg0, _arg1)
}

// SetShowAll sets whether the app chooser should show all applications in a
// flat list.
func (s appChooserWidget) SetShowAll(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_all(_arg0, _arg1)
}

// SetShowDefault sets whether the app chooser should show the default
// handler for the content type in a separate section.
func (s appChooserWidget) SetShowDefault(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_default(_arg0, _arg1)
}

// SetShowFallback sets whether the app chooser should show related
// applications for the content type in a separate section.
func (s appChooserWidget) SetShowFallback(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_fallback(_arg0, _arg1)
}

// SetShowOther sets whether the app chooser should show applications which
// are unrelated to the content type.
func (s appChooserWidget) SetShowOther(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_other(_arg0, _arg1)
}

// SetShowRecommended sets whether the app chooser should show recommended
// applications for the content type in a separate section.
func (s appChooserWidget) SetShowRecommended(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_recommended(_arg0, _arg1)
}
