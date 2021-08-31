// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_application_window_get_type()), F: marshalApplicationWindower},
	})
}

// ApplicationWindow is a Window subclass that offers some extra functionality
// for better integration with Application features. Notably, it can handle both
// the application menu as well as the menubar. See
// gtk_application_set_app_menu() and gtk_application_set_menubar().
//
// This class implements the Group and Map interfaces, to let you add
// window-specific actions that will be exported by the associated Application,
// together with its application-wide actions. Window-specific actions are
// prefixed with the “win.” prefix and application-wide actions are prefixed
// with the “app.” prefix. Actions must be addressed with the prefixed name when
// referring to them from a Model.
//
// Note that widgets that are placed inside a ApplicationWindow can also
// activate these actions, if they implement the Actionable interface.
//
// As with Application, the GDK lock will be acquired when processing actions
// arriving from other processes and should therefore be held when activating
// actions locally (if GDK threads are enabled).
//
// The settings Settings:gtk-shell-shows-app-menu and
// Settings:gtk-shell-shows-menubar tell GTK+ whether the desktop environment is
// showing the application menu and menubar models outside the application as
// part of the desktop shell. For instance, on OS X, both menus will be
// displayed remotely; on Windows neither will be. gnome-shell (starting with
// version 3.4) will display the application menu, but not the menubar.
//
// If the desktop environment does not display the menubar, then
// ApplicationWindow will automatically show a MenuBar for it. This behaviour
// can be overridden with the ApplicationWindow:show-menubar property. If the
// desktop environment does not display the application menu, then it will
// automatically be included in the menubar or in the windows client-side
// decorations.
//
// A GtkApplicationWindow with a menubar
//
//    GtkApplication *app = gtk_application_new ("org.gtk.test", 0);
//
//    GtkBuilder *builder = gtk_builder_new_from_string (
//        "<interface>"
//        "  <menu id='menubar'>"
//        "    <submenu label='_Edit'>"
//        "      <item label='_Copy' action='win.copy'/>"
//        "      <item label='_Paste' action='win.paste'/>"
//        "    </submenu>"
//        "  </menu>"
//        "</interface>",
//        -1);
//
//    GMenuModel *menubar = G_MENU_MODEL (gtk_builder_get_object (builder,
//                                                                "menubar"));
//    gtk_application_set_menubar (GTK_APPLICATION (app), menubar);
//    g_object_unref (builder);
//
//    // ...
//
//    GtkWidget *window = gtk_application_window_new (app);
//
//
// Handling fallback yourself
//
// A simple example (https://git.gnome.org/browse/gtk+/tree/examples/sunny.c)
//
// The XML format understood by Builder for Model consists of a toplevel <menu>
// element, which contains one or more <item> elements. Each <item> element
// contains <attribute> and <link> elements with a mandatory name attribute.
// <link> elements have the same content model as <menu>. Instead of <link
// name="submenu> or <link name="section">, you can use <submenu> or <section>
// elements.
//
// Attribute values can be translated using gettext, like other Builder content.
// <attribute> elements can be marked for translation with a translatable="yes"
// attribute. It is also possible to specify message context and translator
// comments, using the context and comments attributes. To make use of this, the
// Builder must have been given the gettext domain to use.
//
// The following attributes are used when constructing menu items:
//
// - "label": a user-visible string to display
//
// - "action": the prefixed name of the action to trigger
//
// - "target": the parameter to use when activating the action
//
// - "icon" and "verb-icon": names of icons that may be displayed
//
// - "submenu-action": name of an action that may be used to determine if a
// submenu can be opened
//
// - "hidden-when": a string used to determine when the item will be hidden.
// Possible values include "action-disabled", "action-missing", "macos-menubar".
//
// The following attributes are used when constructing sections:
//
// - "label": a user-visible string to use as section heading
//
// - "display-hint": a string used to determine special formatting for the
// section. Possible values include "horizontal-buttons".
//
// - "text-direction": a string used to determine the TextDirection to use when
// "display-hint" is set to "horizontal-buttons". Possible values include "rtl",
// "ltr", and "none".
//
// The following attributes are used when constructing submenus:
//
// - "label": a user-visible string to display
//
// - "icon": icon name to display
type ApplicationWindow struct {
	Window

	gio.ActionGroup
	gio.ActionMap
	*externglib.Object
}

func wrapApplicationWindow(obj *externglib.Object) *ApplicationWindow {
	return &ApplicationWindow{
		Window: Window{
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
						Object: obj,
					},
				},
			},
		},
		ActionGroup: gio.ActionGroup{
			Object: obj,
		},
		ActionMap: gio.ActionMap{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalApplicationWindower(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapApplicationWindow(obj), nil
}

// NewApplicationWindow creates a new ApplicationWindow.
func NewApplicationWindow(application *Application) *ApplicationWindow {
	var _arg1 *C.GtkApplication // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	_cret = C.gtk_application_window_new(_arg1)
	runtime.KeepAlive(application)

	var _applicationWindow *ApplicationWindow // out

	_applicationWindow = wrapApplicationWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _applicationWindow
}

// HelpOverlay gets the ShortcutsWindow that has been set up with a prior call
// to gtk_application_window_set_help_overlay().
func (window *ApplicationWindow) HelpOverlay() *ShortcutsWindow {
	var _arg0 *C.GtkApplicationWindow // out
	var _cret *C.GtkShortcutsWindow   // in

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_application_window_get_help_overlay(_arg0)
	runtime.KeepAlive(window)

	var _shortcutsWindow *ShortcutsWindow // out

	if _cret != nil {
		_shortcutsWindow = wrapShortcutsWindow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _shortcutsWindow
}

// ID returns the unique ID of the window. If the window has not yet been added
// to a Application, returns 0.
func (window *ApplicationWindow) ID() uint32 {
	var _arg0 *C.GtkApplicationWindow // out
	var _cret C.guint                 // in

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_application_window_get_id(_arg0)
	runtime.KeepAlive(window)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// ShowMenubar returns whether the window will display a menubar for the app
// menu and menubar as needed.
func (window *ApplicationWindow) ShowMenubar() bool {
	var _arg0 *C.GtkApplicationWindow // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gtk_application_window_get_show_menubar(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetHelpOverlay associates a shortcuts window with the application window, and
// sets up an action with the name win.show-help-overlay to present it.
//
// window takes resposibility for destroying help_overlay.
func (window *ApplicationWindow) SetHelpOverlay(helpOverlay *ShortcutsWindow) {
	var _arg0 *C.GtkApplicationWindow // out
	var _arg1 *C.GtkShortcutsWindow   // out

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))
	if helpOverlay != nil {
		_arg1 = (*C.GtkShortcutsWindow)(unsafe.Pointer(helpOverlay.Native()))
	}

	C.gtk_application_window_set_help_overlay(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(helpOverlay)
}

// SetShowMenubar sets whether the window will display a menubar for the app
// menu and menubar as needed.
func (window *ApplicationWindow) SetShowMenubar(showMenubar bool) {
	var _arg0 *C.GtkApplicationWindow // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkApplicationWindow)(unsafe.Pointer(window.Native()))
	if showMenubar {
		_arg1 = C.TRUE
	}

	C.gtk_application_window_set_show_menubar(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(showMenubar)
}
