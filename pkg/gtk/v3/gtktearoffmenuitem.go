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
		{T: externglib.Type(C.gtk_tearoff_menu_item_get_type()), F: marshalTearoffMenuItem},
	})
}

// TearoffMenuItem is a special MenuItem which is used to tear off and reattach
// its menu.
//
// When its menu is shown normally, the TearoffMenuItem is drawn as a dotted
// line indicating that the menu can be torn off. Activating it causes its menu
// to be torn off and displayed in its own window as a tearoff menu.
//
// When its menu is shown as a tearoff menu, the TearoffMenuItem is drawn as a
// dotted line which has a left pointing arrow graphic indicating that the
// tearoff menu can be reattached. Activating it will erase the tearoff menu
// window.
//
// > TearoffMenuItem is deprecated and should not be used in newly > written
// code. Menus are not meant to be torn around.
type TearoffMenuItem interface {
	gextras.Objector

	privateTearoffMenuItemClass()
}

// TearoffMenuItemClass implements the TearoffMenuItem interface.
type TearoffMenuItemClass struct {
	*externglib.Object
	MenuItemClass
	ActionableInterface
	ActivatableInterface
	BuildableInterface
}

var _ TearoffMenuItem = (*TearoffMenuItemClass)(nil)

func wrapTearoffMenuItem(obj *externglib.Object) TearoffMenuItem {
	return &TearoffMenuItemClass{
		Object: obj,
		MenuItemClass: MenuItemClass{
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
			ActionableInterface: ActionableInterface{
				WidgetClass: WidgetClass{
					Object:           obj,
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
			},
			ActivatableInterface: ActivatableInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		ActionableInterface: ActionableInterface{
			WidgetClass: WidgetClass{
				Object:           obj,
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
		},
		ActivatableInterface: ActivatableInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalTearoffMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTearoffMenuItem(obj), nil
}

// NewTearoffMenuItem creates a new TearoffMenuItem.
//
// Deprecated: since version 3.4.
func NewTearoffMenuItem() TearoffMenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_tearoff_menu_item_new()

	var _tearoffMenuItem TearoffMenuItem // out

	_tearoffMenuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TearoffMenuItem)

	return _tearoffMenuItem
}

func (*TearoffMenuItemClass) privateTearoffMenuItemClass() {}
