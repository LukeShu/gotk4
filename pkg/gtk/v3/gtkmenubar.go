// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_menu_bar_get_type()), F: marshalMenuBar},
	})
}

// MenuBar: the MenuBar is a subclass of MenuShell which contains one or more
// MenuItems. The result is a standard menu bar which can hold many menu items.
//
//
// CSS nodes
//
// GtkMenuBar has a single CSS node with name menubar.
type MenuBar interface {
	gextras.Objector

	// ChildPackDirection retrieves the current child pack direction of the
	// menubar. See gtk_menu_bar_set_child_pack_direction().
	ChildPackDirection() PackDirection
	// PackDirection retrieves the current pack direction of the menubar. See
	// gtk_menu_bar_set_pack_direction().
	PackDirection() PackDirection
}

// MenuBarClass implements the MenuBar interface.
type MenuBarClass struct {
	*externglib.Object
	MenuShellClass
	BuildableInterface
}

var _ MenuBar = (*MenuBarClass)(nil)

func wrapMenuBar(obj *externglib.Object) MenuBar {
	return &MenuBarClass{
		Object: obj,
		MenuShellClass: MenuShellClass{
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
	}
}

func marshalMenuBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuBar(obj), nil
}

// NewMenuBar creates a new MenuBar
func NewMenuBar() *MenuBarClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_bar_new()

	var _menuBar *MenuBarClass // out

	_menuBar = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*MenuBarClass)

	return _menuBar
}

// NewMenuBarFromModel creates a new MenuBar and populates it with menu items
// and submenus according to @model.
//
// The created menu items are connected to actions found in the
// ApplicationWindow to which the menu bar belongs - typically by means of being
// contained within the ApplicationWindows widget hierarchy.
func NewMenuBarFromModel(model gio.MenuModel) *MenuBarClass {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer((&model).Native()))

	_cret = C.gtk_menu_bar_new_from_model(_arg1)

	var _menuBar *MenuBarClass // out

	_menuBar = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*MenuBarClass)

	return _menuBar
}

// ChildPackDirection retrieves the current child pack direction of the menubar.
// See gtk_menu_bar_set_child_pack_direction().
func (m *MenuBarClass) ChildPackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer((&m).Native()))

	_cret = C.gtk_menu_bar_get_child_pack_direction(_arg0)

	var _packDirection PackDirection // out

	_packDirection = (PackDirection)(_cret)

	return _packDirection
}

// PackDirection retrieves the current pack direction of the menubar. See
// gtk_menu_bar_set_pack_direction().
func (m *MenuBarClass) PackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer((&m).Native()))

	_cret = C.gtk_menu_bar_get_pack_direction(_arg0)

	var _packDirection PackDirection // out

	_packDirection = (PackDirection)(_cret)

	return _packDirection
}
