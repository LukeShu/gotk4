// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
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
		{T: externglib.Type(C.gtk_menu_bar_get_type()), F: marshalMenuBarer},
	})
}

// MenuBarer describes MenuBar's methods.
type MenuBarer interface {
	// ChildPackDirection retrieves the current child pack direction of the
	// menubar.
	ChildPackDirection() PackDirection
	// PackDirection retrieves the current pack direction of the menubar.
	PackDirection() PackDirection
	// SetChildPackDirection sets how widgets should be packed inside the
	// children of a menubar.
	SetChildPackDirection(childPackDir PackDirection)
	// SetPackDirection sets how items should be packed inside a menubar.
	SetPackDirection(packDir PackDirection)
}

// MenuBar is a subclass of MenuShell which contains one or more MenuItems. The
// result is a standard menu bar which can hold many menu items.
//
//
// CSS nodes
//
// GtkMenuBar has a single CSS node with name menubar.
type MenuBar struct {
	MenuShell
}

var (
	_ MenuBarer       = (*MenuBar)(nil)
	_ gextras.Nativer = (*MenuBar)(nil)
)

func wrapMenuBar(obj *externglib.Object) MenuBarer {
	return &MenuBar{
		MenuShell: MenuShell{
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
				},
			},
		},
	}
}

func marshalMenuBarer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuBar(obj), nil
}

// NewMenuBar creates a new MenuBar
func NewMenuBar() *MenuBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_bar_new()

	var _menuBar *MenuBar // out

	_menuBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MenuBar)

	return _menuBar
}

// NewMenuBarFromModel creates a new MenuBar and populates it with menu items
// and submenus according to @model.
//
// The created menu items are connected to actions found in the
// ApplicationWindow to which the menu bar belongs - typically by means of being
// contained within the ApplicationWindows widget hierarchy.
func NewMenuBarFromModel(model gio.MenuModeler) *MenuBar {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	_cret = C.gtk_menu_bar_new_from_model(_arg1)

	var _menuBar *MenuBar // out

	_menuBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MenuBar)

	return _menuBar
}

// ChildPackDirection retrieves the current child pack direction of the menubar.
// See gtk_menu_bar_set_child_pack_direction().
func (menubar *MenuBar) ChildPackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))

	_cret = C.gtk_menu_bar_get_child_pack_direction(_arg0)

	var _packDirection PackDirection // out

	_packDirection = PackDirection(_cret)

	return _packDirection
}

// PackDirection retrieves the current pack direction of the menubar. See
// gtk_menu_bar_set_pack_direction().
func (menubar *MenuBar) PackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))

	_cret = C.gtk_menu_bar_get_pack_direction(_arg0)

	var _packDirection PackDirection // out

	_packDirection = PackDirection(_cret)

	return _packDirection
}

// SetChildPackDirection sets how widgets should be packed inside the children
// of a menubar.
func (menubar *MenuBar) SetChildPackDirection(childPackDir PackDirection) {
	var _arg0 *C.GtkMenuBar      // out
	var _arg1 C.GtkPackDirection // out

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))
	_arg1 = C.GtkPackDirection(childPackDir)

	C.gtk_menu_bar_set_child_pack_direction(_arg0, _arg1)
}

// SetPackDirection sets how items should be packed inside a menubar.
func (menubar *MenuBar) SetPackDirection(packDir PackDirection) {
	var _arg0 *C.GtkMenuBar      // out
	var _arg1 C.GtkPackDirection // out

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(menubar.Native()))
	_arg1 = C.GtkPackDirection(packDir)

	C.gtk_menu_bar_set_pack_direction(_arg0, _arg1)
}
