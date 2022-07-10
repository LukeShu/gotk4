// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeMenuBar returns the GType for the type MenuBar.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMenuBar() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MenuBar").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMenuBar)
	return gtype
}

// MenuBarOverrider contains methods that are overridable.
type MenuBarOverrider interface {
}

// MenuBar is a subclass of MenuShell which contains one or more MenuItems. The
// result is a standard menu bar which can hold many menu items.
//
//
// CSS nodes
//
// GtkMenuBar has a single CSS node with name menubar.
type MenuBar struct {
	_ [0]func() // equal guard
	MenuShell
}

var (
	_ MenuSheller = (*MenuBar)(nil)
)

func classInitMenuBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMenuBar(obj *coreglib.Object) *MenuBar {
	return &MenuBar{
		MenuShell: MenuShell{
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
	}
}

func marshalMenuBar(p uintptr) (interface{}, error) {
	return wrapMenuBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMenuBar creates a new MenuBar.
//
// The function returns the following values:
//
//    - menuBar: new menu bar, as a Widget.
//
func NewMenuBar() *MenuBar {
	_info := girepository.MustFind("Gtk", "MenuBar")
	_gret := _info.InvokeClassMethod("new_MenuBar", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _menuBar *MenuBar // out

	_menuBar = wrapMenuBar(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _menuBar
}

// NewMenuBarFromModel creates a new MenuBar and populates it with menu items
// and submenus according to model.
//
// The created menu items are connected to actions found in the
// ApplicationWindow to which the menu bar belongs - typically by means of being
// contained within the ApplicationWindows widget hierarchy.
//
// The function takes the following parameters:
//
//    - model: Model.
//
// The function returns the following values:
//
//    - menuBar: new MenuBar.
//
func NewMenuBarFromModel(model gio.MenuModeller) *MenuBar {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))

	_info := girepository.MustFind("Gtk", "MenuBar")
	_gret := _info.InvokeClassMethod("new_MenuBar_from_model", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _menuBar *MenuBar // out

	_menuBar = wrapMenuBar(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _menuBar
}
