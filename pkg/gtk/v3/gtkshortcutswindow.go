// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
// extern void _gotk4_gtk3_ShortcutsWindow_ConnectSearch(gpointer, guintptr);
// extern void _gotk4_gtk3_ShortcutsWindow_ConnectClose(gpointer, guintptr);
// extern void _gotk4_gtk3_ShortcutsWindowClass_search(GtkShortcutsWindow*);
// extern void _gotk4_gtk3_ShortcutsWindowClass_close(GtkShortcutsWindow*);
// void _gotk4_gtk3_ShortcutsWindow_virtual_close(void* fnptr, GtkShortcutsWindow* arg0) {
//   ((void (*)(GtkShortcutsWindow*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_ShortcutsWindow_virtual_search(void* fnptr, GtkShortcutsWindow* arg0) {
//   ((void (*)(GtkShortcutsWindow*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeShortcutsWindow = coreglib.Type(C.gtk_shortcuts_window_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutsWindow, F: marshalShortcutsWindow},
	})
}

// ShortcutsWindowOverrides contains methods that are overridable.
type ShortcutsWindowOverrides struct {
	Close  func()
	Search func()
}

func defaultShortcutsWindowOverrides(v *ShortcutsWindow) ShortcutsWindowOverrides {
	return ShortcutsWindowOverrides{
		Close:  v.close,
		Search: v.search,
	}
}

// ShortcutsWindow shows brief information about the keyboard shortcuts and
// gestures of an application. The shortcuts can be grouped, and you can have
// multiple sections in this window, corresponding to the major modes of your
// application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a GtkShortcutsWindow is with GtkBuilder,
// by populating a ShortcutsWindow with one or more ShortcutsSection objects,
// which contain ShortcutsGroups that in turn contain objects of class
// ShortcutsShortcut.
//
// A simple example:
//
// ! (gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups are
// arranged in columns, and spread across several pages if there are too many to
// find on a single page.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-gedit.ui).
//
// An example with multiple views:
//
// ! (clocks-shortcuts.png)
//
// This example shows a ShortcutsWindow that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-clocks.ui).
//
// An example with multiple sections:
//
// ! (builder-shortcuts.png)
//
// This example shows a ShortcutsWindow with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Binner = (*ShortcutsWindow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ShortcutsWindow, *ShortcutsWindowClass, ShortcutsWindowOverrides](
		GTypeShortcutsWindow,
		initShortcutsWindowClass,
		wrapShortcutsWindow,
		defaultShortcutsWindowOverrides,
	)
}

func initShortcutsWindowClass(gclass unsafe.Pointer, overrides ShortcutsWindowOverrides, classInitFunc func(*ShortcutsWindowClass)) {
	pclass := (*C.GtkShortcutsWindowClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeShortcutsWindow))))

	if overrides.Close != nil {
		pclass.close = (*[0]byte)(C._gotk4_gtk3_ShortcutsWindowClass_close)
	}

	if overrides.Search != nil {
		pclass.search = (*[0]byte)(C._gotk4_gtk3_ShortcutsWindowClass_search)
	}

	if classInitFunc != nil {
		class := (*ShortcutsWindowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapShortcutsWindow(obj *coreglib.Object) *ShortcutsWindow {
	return &ShortcutsWindow{
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
	}
}

func marshalShortcutsWindow(p uintptr) (interface{}, error) {
	return wrapShortcutsWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClose signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to close the window.
//
// The default binding for this signal is the Escape key.
func (v *ShortcutsWindow) ConnectClose(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "close", false, unsafe.Pointer(C._gotk4_gtk3_ShortcutsWindow_ConnectClose), f)
}

// ConnectSearch signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to start a search.
//
// The default binding for this signal is Control-F.
func (v *ShortcutsWindow) ConnectSearch(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "search", false, unsafe.Pointer(C._gotk4_gtk3_ShortcutsWindow_ConnectSearch), f)
}

func (self *ShortcutsWindow) close() {
	gclass := (*C.GtkShortcutsWindowClass)(coreglib.PeekParentClass(self))
	fnarg := gclass.close

	var _arg0 *C.GtkShortcutsWindow // out

	_arg0 = (*C.GtkShortcutsWindow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C._gotk4_gtk3_ShortcutsWindow_virtual_close(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(self)
}

func (self *ShortcutsWindow) search() {
	gclass := (*C.GtkShortcutsWindowClass)(coreglib.PeekParentClass(self))
	fnarg := gclass.search

	var _arg0 *C.GtkShortcutsWindow // out

	_arg0 = (*C.GtkShortcutsWindow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C._gotk4_gtk3_ShortcutsWindow_virtual_search(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(self)
}

// ShortcutsWindowClass: instance of this type is always passed by reference.
type ShortcutsWindowClass struct {
	*shortcutsWindowClass
}

// shortcutsWindowClass is the struct that's finalized.
type shortcutsWindowClass struct {
	native *C.GtkShortcutsWindowClass
}

func (s *ShortcutsWindowClass) ParentClass() *WindowClass {
	valptr := &s.native.parent_class
	var _v *WindowClass // out
	_v = (*WindowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
