// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_ApplicationClass_window_added(GtkApplication*, GtkWindow*);
// extern void _gotk4_gtk4_ApplicationClass_window_removed(GtkApplication*, GtkWindow*);
// extern void _gotk4_gtk4_Application_ConnectQueryEnd(gpointer, guintptr);
// extern void _gotk4_gtk4_Application_ConnectWindowAdded(gpointer, GtkWindow*, guintptr);
// extern void _gotk4_gtk4_Application_ConnectWindowRemoved(gpointer, GtkWindow*, guintptr);
import "C"

// glib.Type values for gtkapplication.go.
var (
	GTypeApplicationInhibitFlags = coreglib.Type(C.gtk_application_inhibit_flags_get_type())
	GTypeApplication             = coreglib.Type(C.gtk_application_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeApplicationInhibitFlags, F: marshalApplicationInhibitFlags},
		{T: GTypeApplication, F: marshalApplication},
	})
}

// ApplicationInhibitFlags types of user actions that may be blocked by
// GtkApplication.
//
// See gtk.Application.Inhibit().
type ApplicationInhibitFlags C.guint

const (
	// ApplicationInhibitLogout: inhibit ending the user session by logging out
	// or by shutting down the computer.
	ApplicationInhibitLogout ApplicationInhibitFlags = 0b1
	// ApplicationInhibitSwitch: inhibit user switching.
	ApplicationInhibitSwitch ApplicationInhibitFlags = 0b10
	// ApplicationInhibitSuspend: inhibit suspending the session or computer.
	ApplicationInhibitSuspend ApplicationInhibitFlags = 0b100
	// ApplicationInhibitIdle: inhibit the session being marked as idle (and
	// possibly locked).
	ApplicationInhibitIdle ApplicationInhibitFlags = 0b1000
)

func marshalApplicationInhibitFlags(p uintptr) (interface{}, error) {
	return ApplicationInhibitFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ApplicationInhibitFlags.
func (a ApplicationInhibitFlags) String() string {
	if a == 0 {
		return "ApplicationInhibitFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(98)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case ApplicationInhibitLogout:
			builder.WriteString("Logout|")
		case ApplicationInhibitSwitch:
			builder.WriteString("Switch|")
		case ApplicationInhibitSuspend:
			builder.WriteString("Suspend|")
		case ApplicationInhibitIdle:
			builder.WriteString("Idle|")
		default:
			builder.WriteString(fmt.Sprintf("ApplicationInhibitFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a ApplicationInhibitFlags) Has(other ApplicationInhibitFlags) bool {
	return (a & other) == other
}

// ApplicationOverrider contains methods that are overridable.
type ApplicationOverrider interface {
	// The function takes the following parameters:
	//
	WindowAdded(window *Window)
	// The function takes the following parameters:
	//
	WindowRemoved(window *Window)
}

// Application: GtkApplication is a high-level API for writing applications.
//
// It supports many aspects of writing a GTK application in a convenient
// fashion, without enforcing a one-size-fits-all model.
//
// Currently, GtkApplication handles GTK initialization, application uniqueness,
// session management, provides some basic scriptability and desktop shell
// integration by exporting actions and menus and manages a list of toplevel
// windows whose life-cycle is automatically tied to the life-cycle of your
// application.
//
// While GtkApplication works fine with plain gtk.Windows, it is recommended to
// use it together with gtk.ApplicationWindow.
//
//
// Automatic resources
//
// GtkApplication will automatically load menus from the GtkBuilder resource
// located at "gtk/menus.ui", relative to the application's resource base path
// (see g_application_set_resource_base_path()). The menu with the ID "menubar"
// is taken as the application's menubar. Additional menus (most interesting
// submenus) can be named and accessed via gtk.Application.GetMenuByID() which
// allows for dynamic population of a part of the menu structure.
//
// It is also possible to provide the menubar manually using
// gtk.Application.SetMenubar().
//
// GtkApplication will also automatically setup an icon search path for the
// default icon theme by appending "icons" to the resource base path. This
// allows your application to easily store its icons as resources. See
// gtk.IconTheme.AddResourcePath() for more information.
//
// If there is a resource located at "gtk/help-overlay.ui" which defines a
// gtk.ShortcutsWindow with ID "help_overlay" then GtkApplication associates an
// instance of this shortcuts window with each gtk.ApplicationWindow and sets up
// the keyboard accelerator <kbd>Control</kbd>+<kbd>?</kbd> to open it. To
// create a menu item that displays the shortcuts window, associate the item
// with the action win.show-help-overlay.
//
//
// A simple application
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/examples/bp/bloatpad.c) is
// available in the GTK source code repository
//
// GtkApplication optionally registers with a session manager of the users
// session (if you set the gtk.Application:register-session property) and offers
// various functionality related to the session life-cycle.
//
// An application can block various ways to end the session with the
// gtk.Application.Inhibit() function. Typical use cases for this kind of
// inhibiting are long-running, uninterruptible operations, such as burning a CD
// or performing a disk backup. The session manager may not honor the inhibitor,
// but it can be expected to inform the user about the negative consequences of
// ending the session while inhibitors are present.
//
//
// See Also
//
// HowDoI: Using GtkApplication (https://wiki.gnome.org/HowDoI/GtkApplication),
// Getting Started with GTK: Basics (getting_started.html#basics).
type Application struct {
	_ [0]func() // equal guard
	gio.Application
}

var (
	_ coreglib.Objector = (*Application)(nil)
)

func classInitApplicationer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkApplicationClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkApplicationClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ WindowAdded(window *Window) }); ok {
		pclass.window_added = (*[0]byte)(C._gotk4_gtk4_ApplicationClass_window_added)
	}

	if _, ok := goval.(interface{ WindowRemoved(window *Window) }); ok {
		pclass.window_removed = (*[0]byte)(C._gotk4_gtk4_ApplicationClass_window_removed)
	}
}

//export _gotk4_gtk4_ApplicationClass_window_added
func _gotk4_gtk4_ApplicationClass_window_added(arg0 *C.GtkApplication, arg1 *C.GtkWindow) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ WindowAdded(window *Window) })

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(arg1)))

	iface.WindowAdded(_window)
}

//export _gotk4_gtk4_ApplicationClass_window_removed
func _gotk4_gtk4_ApplicationClass_window_removed(arg0 *C.GtkApplication, arg1 *C.GtkWindow) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ WindowRemoved(window *Window) })

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(arg1)))

	iface.WindowRemoved(_window)
}

func wrapApplication(obj *coreglib.Object) *Application {
	return &Application{
		Application: gio.Application{
			Object: obj,
			ActionGroup: gio.ActionGroup{
				Object: obj,
			},
			ActionMap: gio.ActionMap{
				Object: obj,
			},
		},
	}
}

func marshalApplication(p uintptr) (interface{}, error) {
	return wrapApplication(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Application_ConnectQueryEnd
func _gotk4_gtk4_Application_ConnectQueryEnd(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectQueryEnd is emitted when the session manager is about to end the
// session.
//
// This signal is only emitted if gtk.Application:register-session is TRUE.
// Applications can connect to this signal and call gtk.Application.Inhibit()
// with GTK_APPLICATION_INHIBIT_LOGOUT to delay the end of the session until
// state has been saved.
func (application *Application) ConnectQueryEnd(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(application, "query-end", false, unsafe.Pointer(C._gotk4_gtk4_Application_ConnectQueryEnd), f)
}

//export _gotk4_gtk4_Application_ConnectWindowAdded
func _gotk4_gtk4_Application_ConnectWindowAdded(arg0 C.gpointer, arg1 *C.GtkWindow, arg2 C.guintptr) {
	var f func(window *Window)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(window *Window))
	}

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(arg1)))

	f(_window)
}

// ConnectWindowAdded is emitted when a gtk.Window is added to application
// through gtk.Application.AddWindow().
func (application *Application) ConnectWindowAdded(f func(window *Window)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(application, "window-added", false, unsafe.Pointer(C._gotk4_gtk4_Application_ConnectWindowAdded), f)
}

//export _gotk4_gtk4_Application_ConnectWindowRemoved
func _gotk4_gtk4_Application_ConnectWindowRemoved(arg0 C.gpointer, arg1 *C.GtkWindow, arg2 C.guintptr) {
	var f func(window *Window)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(window *Window))
	}

	var _window *Window // out

	_window = wrapWindow(coreglib.Take(unsafe.Pointer(arg1)))

	f(_window)
}

// ConnectWindowRemoved is emitted when a gtk.Window is removed from
// application.
//
// This can happen as a side-effect of the window being destroyed or explicitly
// through gtk.Application.RemoveWindow().
func (application *Application) ConnectWindowRemoved(f func(window *Window)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(application, "window-removed", false, unsafe.Pointer(C._gotk4_gtk4_Application_ConnectWindowRemoved), f)
}

// AddWindow adds a window to application.
//
// This call can only happen after the application has started; typically, you
// should add new application windows in response to the emission of the
// GApplication::activate signal.
//
// This call is equivalent to setting the gtk.Window:application property of
// window to application.
//
// Normally, the connection between the application and the window will remain
// until the window is destroyed, but you can explicitly remove it with
// gtk.Application.RemoveWindow().
//
// GTK will keep the application running as long as it has any windows.
//
// The function takes the following parameters:
//
//    - window: GtkWindow.
//
func (application *Application) AddWindow(window *Window) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Application").InvokeMethod("add_window", args[:], nil)

	runtime.KeepAlive(application)
	runtime.KeepAlive(window)
}

// AccelsForAction gets the accelerators that are currently associated with the
// given action.
//
// The function takes the following parameters:
//
//    - detailedActionName: detailed action name, specifying an action and target
//      to obtain accelerators for.
//
// The function returns the following values:
//
//    - utf8s: accelerators for detailed_action_name.
//
func (application *Application) AccelsForAction(detailedActionName string) []string {
	var args [2]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 *C.void  // out
	var _cret **C.char // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(detailedActionName)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("get_accels_for_action", args[:], nil)
	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)
	runtime.KeepAlive(detailedActionName)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ActionsForAccel returns the list of actions (possibly empty) that accel maps
// to.
//
// Each item in the list is a detailed action name in the usual form.
//
// This might be useful to discover if an accel already exists in order to
// prevent installation of a conflicting accelerator (from an accelerator editor
// or a plugin system, for example). Note that having more than one action per
// accelerator may not be a bad thing and might make sense in cases where the
// actions never appear in the same context.
//
// In case there are no actions for a given accelerator, an empty array is
// returned. NULL is never returned.
//
// It is a programmer error to pass an invalid accelerator string.
//
// If you are unsure, check it with gtk.AcceleratorParse() first.
//
// The function takes the following parameters:
//
//    - accel: accelerator that can be parsed by gtk.AcceleratorParse().
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of actions for accel.
//
func (application *Application) ActionsForAccel(accel string) []string {
	var args [2]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 *C.void  // out
	var _cret **C.char // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(accel)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("get_actions_for_accel", args[:], nil)
	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)
	runtime.KeepAlive(accel)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ActiveWindow gets the “active” window for the application.
//
// The active window is the one that was most recently focused (within the
// application). This window may not have the focus at the moment if another
// application has it — this is just the most recently-focused window within
// this application.
//
// The function returns the following values:
//
//    - window (optional): active window.
//
func (application *Application) ActiveWindow() *Window {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	*(**Application)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("get_active_window", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)

	var _window *Window // out

	if _cret != nil {
		_window = wrapWindow(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _window
}

// MenuByID gets a menu from automatically loaded resources.
//
// See the section on Automatic resources
// (class.Application.html#automatic-resources) for more information.
//
// The function takes the following parameters:
//
//    - id of the menu to look up.
//
// The function returns the following values:
//
//    - menu (optional) gets the menu with the given id from the automatically
//      loaded resources.
//
func (application *Application) MenuByID(id string) *gio.Menu {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("get_menu_by_id", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)
	runtime.KeepAlive(id)

	var _menu *gio.Menu // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_menu = &gio.Menu{
				MenuModel: gio.MenuModel{
					Object: obj,
				},
			}
		}
	}

	return _menu
}

// Menubar returns the menu model that has been set with
// gtk.Application.SetMenubar().
//
// The function returns the following values:
//
//    - menuModel (optional): menubar for windows of application.
//
func (application *Application) Menubar() gio.MenuModeller {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	*(**Application)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("get_menubar", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.MenuModeller)
				return ok
			})
			rv, ok := casted.(gio.MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// WindowByID returns the gtk.ApplicationWindow with the given ID.
//
// The ID of a GtkApplicationWindow can be retrieved with
// gtk.ApplicationWindow.GetID().
//
// The function takes the following parameters:
//
//    - id: identifier number.
//
// The function returns the following values:
//
//    - window (optional) for the given id.
//
func (application *Application) WindowByID(id uint32) *Window {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = C.guint(id)
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("get_window_by_id", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)
	runtime.KeepAlive(id)

	var _window *Window // out

	if _cret != nil {
		_window = wrapWindow(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _window
}

// Windows gets a list of the gtk.Window instances associated with application.
//
// The list is sorted by most recently focused window, such that the first
// element is the currently focused window. (Useful for choosing a parent for a
// transient window.)
//
// The list that is returned should not be modified in any way. It will only
// remain valid until the next focus change or window creation or deletion.
//
// The function returns the following values:
//
//    - list: GList of GtkWindow instances.
//
func (application *Application) Windows() []*Window {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	*(**Application)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("get_windows", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)

	var _list []*Window // out

	_list = make([]*Window, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *Window // out
		dst = wrapWindow(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// ListActionDescriptions lists the detailed action names which have associated
// accelerators.
//
// See gtk.Application.SetAccelsForAction().
//
// The function returns the following values:
//
//    - utf8s: detailed action names.
//
func (application *Application) ListActionDescriptions() []string {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret **C.char // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	*(**Application)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Application").InvokeMethod("list_action_descriptions", args[:], nil)
	_cret = *(***C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(application)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// RemoveWindow: remove a window from application.
//
// If window belongs to application then this call is equivalent to setting the
// gtk.Window:application property of window to NULL.
//
// The application may stop running as a result of a call to this function, if
// window was the last window of the application.
//
// The function takes the following parameters:
//
//    - window: GtkWindow.
//
func (application *Application) RemoveWindow(window *Window) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Application").InvokeMethod("remove_window", args[:], nil)

	runtime.KeepAlive(application)
	runtime.KeepAlive(window)
}

// SetAccelsForAction sets zero or more keyboard accelerators that will trigger
// the given action.
//
// The first item in accels will be the primary accelerator, which may be
// displayed in the UI.
//
// To remove all accelerators for an action, use an empty, zero-terminated array
// for accels.
//
// For the detailed_action_name, see g_action_parse_detailed_name() and
// g_action_print_detailed_name().
//
// The function takes the following parameters:
//
//    - detailedActionName: detailed action name, specifying an action and target
//      to associate accelerators with.
//    - accels: list of accelerators in the format understood by
//      gtk.AcceleratorParse().
//
func (application *Application) SetAccelsForAction(detailedActionName string, accels []string) {
	var args [3]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 *C.void  // out
	var _arg2 **C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(detailedActionName)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.void)(C.calloc(C.size_t((len(accels) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(accels)+1)
			var zero *C.void
			out[len(accels)] = zero
			for i := range accels {
				out[i] = (*C.void)(unsafe.Pointer(C.CString(accels[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1
	*(*string)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "Application").InvokeMethod("set_accels_for_action", args[:], nil)

	runtime.KeepAlive(application)
	runtime.KeepAlive(detailedActionName)
	runtime.KeepAlive(accels)
}

// SetMenubar sets or unsets the menubar for windows of application.
//
// This is a menubar in the traditional sense.
//
// This can only be done in the primary instance of the application, after it
// has been registered. GApplication::startup is a good place to call this.
//
// Depending on the desktop environment, this may appear at the top of each
// window, or at the top of the screen. In some environments, if both the
// application menu and the menubar are set, the application menu will be
// presented as if it were the first item of the menubar. Other environments
// treat the two as completely separate — for example, the application menu may
// be rendered by the desktop shell while the menubar (if set) remains in each
// individual window.
//
// Use the base GActionMap interface to add actions, to respond to the user
// selecting these menu items.
//
// The function takes the following parameters:
//
//    - menubar (optional): GMenuModel.
//
func (application *Application) SetMenubar(menubar gio.MenuModeller) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	if menubar != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(menubar).Native()))
	}
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Application").InvokeMethod("set_menubar", args[:], nil)

	runtime.KeepAlive(application)
	runtime.KeepAlive(menubar)
}

// Uninhibit removes an inhibitor that has been previously established.
//
// See gtk.Application.Inhibit().
//
// Inhibitors are also cleared when the application exits.
//
// The function takes the following parameters:
//
//    - cookie that was returned by gtk.Application.Inhibit().
//
func (application *Application) Uninhibit(cookie uint32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(application).Native()))
	_arg1 = C.guint(cookie)
	*(**Application)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Application").InvokeMethod("uninhibit", args[:], nil)

	runtime.KeepAlive(application)
	runtime.KeepAlive(cookie)
}
