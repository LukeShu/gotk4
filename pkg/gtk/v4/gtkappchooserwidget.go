// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_AppChooserWidget_ConnectApplicationActivated(gpointer, void*, guintptr);
// extern void _gotk4_gtk4_AppChooserWidget_ConnectApplicationSelected(gpointer, void*, guintptr);
import "C"

// GTypeAppChooserWidget returns the GType for the type AppChooserWidget.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAppChooserWidget() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "AppChooserWidget").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAppChooserWidget)
	return gtype
}

// AppChooserWidget: GtkAppChooserWidget is a widget for selecting applications.
//
// It is the main building block for gtk.AppChooserDialog. Most applications
// only need to use the latter; but you can use this widget as part of a larger
// widget if you have special needs.
//
// GtkAppChooserWidget offers detailed control over what applications are shown,
// using the gtk.AppChooserWidget:show-default,
// gtk.AppChooserWidget:show-recommended, gtk.AppChooserWidget:show-fallback,
// gtk.AppChooserWidget:show-other and gtk.AppChooserWidget:show-all properties.
// See the gtk.AppChooser documentation for more information about these groups
// of applications.
//
// To keep track of the selected application, use the
// gtk.AppChooserWidget::application-selected and
// gtk.AppChooserWidget::application-activated signals.
//
//
// CSS nodes
//
// GtkAppChooserWidget has a single CSS node with name appchooser.
type AppChooserWidget struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	AppChooser
}

var (
	_ Widgetter         = (*AppChooserWidget)(nil)
	_ coreglib.Objector = (*AppChooserWidget)(nil)
)

func wrapAppChooserWidget(obj *coreglib.Object) *AppChooserWidget {
	return &AppChooserWidget{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		AppChooser: AppChooser{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalAppChooserWidget(p uintptr) (interface{}, error) {
	return wrapAppChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_AppChooserWidget_ConnectApplicationActivated
func _gotk4_gtk4_AppChooserWidget_ConnectApplicationActivated(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(application gio.AppInfor)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(application gio.AppInfor))
	}

	var _application gio.AppInfor // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AppInfor is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.AppInfor)
			return ok
		})
		rv, ok := casted.(gio.AppInfor)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AppInfor")
		}
		_application = rv
	}

	f(_application)
}

// ConnectApplicationActivated is emitted when an application item is activated
// from the widget's list.
//
// This usually happens when the user double clicks an item, or an item is
// selected and the user presses one of the keys Space, Shift+Space, Return or
// Enter.
func (self *AppChooserWidget) ConnectApplicationActivated(f func(application gio.AppInfor)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "application-activated", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserWidget_ConnectApplicationActivated), f)
}

//export _gotk4_gtk4_AppChooserWidget_ConnectApplicationSelected
func _gotk4_gtk4_AppChooserWidget_ConnectApplicationSelected(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(application gio.AppInfor)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(application gio.AppInfor))
	}

	var _application gio.AppInfor // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AppInfor is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.AppInfor)
			return ok
		})
		rv, ok := casted.(gio.AppInfor)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AppInfor")
		}
		_application = rv
	}

	f(_application)
}

// ConnectApplicationSelected is emitted when an application item is selected
// from the widget's list.
func (self *AppChooserWidget) ConnectApplicationSelected(f func(application gio.AppInfor)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "application-selected", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserWidget_ConnectApplicationSelected), f)
}

// NewAppChooserWidget creates a new GtkAppChooserWidget for applications that
// can handle content of the given type.
//
// The function takes the following parameters:
//
//    - contentType: content type to show applications for.
//
// The function returns the following values:
//
//    - appChooserWidget: newly created GtkAppChooserWidget.
//
func NewAppChooserWidget(contentType string) *AppChooserWidget {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("new_AppChooserWidget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(contentType)

	var _appChooserWidget *AppChooserWidget // out

	_appChooserWidget = wrapAppChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _appChooserWidget
}

// DefaultText returns the text that is shown if there are not applications that
// can handle the content type.
//
// The function returns the following values:
//
//    - utf8 (optional): value of gtk.AppChooserWidget:default-text.
//
func (self *AppChooserWidget) DefaultText() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("get_default_text", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ShowAll gets whether the app chooser should show all applications in a flat
// list.
//
// The function returns the following values:
//
//    - ok: value of gtk.AppChooserWidget:show-all.
//
func (self *AppChooserWidget) ShowAll() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("get_show_all", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefault gets whether the app chooser should show the default handler for
// the content type in a separate section.
//
// The function returns the following values:
//
//    - ok: value of gtk.AppChooserWidget:show-default.
//
func (self *AppChooserWidget) ShowDefault() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("get_show_default", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowFallback gets whether the app chooser should show related applications
// for the content type in a separate section.
//
// The function returns the following values:
//
//    - ok: value of gtk.AppChooserWidget:show-fallback.
//
func (self *AppChooserWidget) ShowFallback() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("get_show_fallback", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowOther gets whether the app chooser should show applications which are
// unrelated to the content type.
//
// The function returns the following values:
//
//    - ok: value of gtk.AppChooserWidget:show-other.
//
func (self *AppChooserWidget) ShowOther() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("get_show_other", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowRecommended gets whether the app chooser should show recommended
// applications for the content type in a separate section.
//
// The function returns the following values:
//
//    - ok: value of gtk.AppChooserWidget:show-recommended.
//
func (self *AppChooserWidget) ShowRecommended() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("get_show_recommended", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultText sets the text that is shown if there are not applications that
// can handle the content type.
//
// The function takes the following parameters:
//
//    - text: new value for gtk.AppChooserWidget:default-text.
//
func (self *AppChooserWidget) SetDefaultText(text string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("set_default_text", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(text)
}

// SetShowAll sets whether the app chooser should show all applications in a
// flat list.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserWidget:show-all.
//
func (self *AppChooserWidget) SetShowAll(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("set_show_all", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetShowDefault sets whether the app chooser should show the default handler
// for the content type in a separate section.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserWidget:show-default.
//
func (self *AppChooserWidget) SetShowDefault(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("set_show_default", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetShowFallback sets whether the app chooser should show related applications
// for the content type in a separate section.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserWidget:show-fallback.
//
func (self *AppChooserWidget) SetShowFallback(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("set_show_fallback", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetShowOther sets whether the app chooser should show applications which are
// unrelated to the content type.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserWidget:show-other.
//
func (self *AppChooserWidget) SetShowOther(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("set_show_other", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetShowRecommended sets whether the app chooser should show recommended
// applications for the content type in a separate section.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserWidget:show-recommended.
//
func (self *AppChooserWidget) SetShowRecommended(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "AppChooserWidget").InvokeMethod("set_show_recommended", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}
