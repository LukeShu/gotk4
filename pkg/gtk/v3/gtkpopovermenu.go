// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
		{T: externglib.Type(C.gtk_popover_menu_get_type()), F: marshalPopoverMenuer},
	})
}

// PopoverMenu is a subclass of Popover that treats its children like menus and
// allows switching between them. It is meant to be used primarily together with
// ModelButton, but any widget can be used, such as SpinButton or Scale. In this
// respect, GtkPopoverMenu is more flexible than popovers that are created from
// a Model with gtk_popover_new_from_model().
//
// To add a child as a submenu, set the PopoverMenu:submenu child property to
// the name of the submenu. To let the user open this submenu, add a ModelButton
// whose ModelButton:menu-name property is set to the name you've given to the
// submenu.
//
// By convention, the first child of a submenu should be a ModelButton to switch
// back to the parent menu. Such a button should use the ModelButton:inverted
// and ModelButton:centered properties to achieve a title-like appearance and
// place the submenu indicator at the opposite side. To switch back to the main
// menu, use "main" as the menu name.
//
// Example
//
//    <object class="GtkPopoverMenu">
//      <child>
//        <object class="GtkBox">
//          <property name="visible">True</property>
//          <property name="margin">10</property>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">win.frob</property>
//              <property name="text" translatable="yes">Frob</property>
//            </object>
//          </child>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="menu-name">more</property>
//              <property name="text" translatable="yes">More</property>
//            </object>
//          </child>
//        </object>
//      </child>
//      <child>
//        <object class="GtkBox">
//          <property name="visible">True</property>
//          <property name="margin">10</property>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">win.foo</property>
//              <property name="text" translatable="yes">Foo</property>
//            </object>
//          </child>
//          <child>
//            <object class="GtkModelButton">
//              <property name="visible">True</property>
//              <property name="action-name">win.bar</property>
//              <property name="text" translatable="yes">Bar</property>
//            </object>
//          </child>
//        </object>
//        <packing>
//          <property name="submenu">more</property>
//        </packing>
//      </child>
//    </object>
//
// Just like normal popovers created using gtk_popover_new_from_model,
// PopoverMenu instances have a single css node called "popover" and get the
// .menu style class.
type PopoverMenu struct {
	Popover
}

func wrapPopoverMenu(obj *externglib.Object) *PopoverMenu {
	return &PopoverMenu{
		Popover: Popover{
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
	}
}

func marshalPopoverMenuer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPopoverMenu(obj), nil
}

// NewPopoverMenu creates a new popover menu.
func NewPopoverMenu() *PopoverMenu {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_popover_menu_new()

	var _popoverMenu *PopoverMenu // out

	_popoverMenu = wrapPopoverMenu(externglib.Take(unsafe.Pointer(_cret)))

	return _popoverMenu
}

// OpenSubmenu opens a submenu of the popover. The name must be one of the names
// given to the submenus of popover with PopoverMenu:submenu, or "main" to
// switch back to the main menu.
//
// ModelButton will open submenus automatically when the ModelButton:menu-name
// property is set, so this function is only needed when you are using other
// kinds of widgets to initiate menu changes.
func (popover *PopoverMenu) OpenSubmenu(name string) {
	var _arg0 *C.GtkPopoverMenu // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkPopoverMenu)(unsafe.Pointer(popover.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_popover_menu_open_submenu(_arg0, _arg1)
}
