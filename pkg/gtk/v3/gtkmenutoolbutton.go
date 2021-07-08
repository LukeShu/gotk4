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
		{T: externglib.Type(C.gtk_menu_tool_button_get_type()), F: marshalMenuToolButton},
	})
}

// MenuToolButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MenuToolButtonOverrider interface {
	ShowMenu()
}

// MenuToolButton is a ToolItem that contains a button and a small additional
// button with an arrow. When clicked, the arrow button pops up a dropdown menu.
//
// Use gtk_menu_tool_button_new() to create a new MenuToolButton.
//
//
// GtkMenuToolButton as GtkBuildable
//
// The GtkMenuToolButton implementation of the GtkBuildable interface supports
// adding a menu by specifying “menu” as the “type” attribute of a <child>
// element.
//
// An example for a UI definition fragment with menus:
//
//    <object class="GtkMenuToolButton">
//      <child type="menu">
//        <object class="GtkMenu"/>
//      </child>
//    </object>
type MenuToolButton interface {
	gextras.Objector

	// Menu gets the Menu associated with MenuToolButton.
	Menu() Widget
	// SetArrowTooltipMarkup sets the tooltip markup text to be used as tooltip
	// for the arrow button which pops up the menu. See
	// gtk_tool_item_set_tooltip_text() for setting a tooltip on the whole
	// MenuToolButton.
	SetArrowTooltipMarkup(markup string)
	// SetArrowTooltipText sets the tooltip text to be used as tooltip for the
	// arrow button which pops up the menu. See gtk_tool_item_set_tooltip_text()
	// for setting a tooltip on the whole MenuToolButton.
	SetArrowTooltipText(text string)
	// SetMenu sets the Menu that is popped up when the user clicks on the
	// arrow. If @menu is NULL, the arrow button becomes insensitive.
	SetMenu(menu Widget)
}

// MenuToolButtonClass implements the MenuToolButton interface.
type MenuToolButtonClass struct {
	*externglib.Object
	ToolButtonClass
	ActionableInterface
	ActivatableInterface
	BuildableInterface
}

var _ MenuToolButton = (*MenuToolButtonClass)(nil)

func wrapMenuToolButton(obj *externglib.Object) MenuToolButton {
	return &MenuToolButtonClass{
		Object: obj,
		ToolButtonClass: ToolButtonClass{
			Object: obj,
			ToolItemClass: ToolItemClass{
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

func marshalMenuToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuToolButton(obj), nil
}

// NewMenuToolButton creates a new MenuToolButton using @icon_widget as icon and
// @label as label.
func NewMenuToolButton(iconWidget Widget, label string) MenuToolButton {
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(iconWidget.Native()))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_menu_tool_button_new(_arg1, _arg2)

	var _menuToolButton MenuToolButton // out

	_menuToolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MenuToolButton)

	return _menuToolButton
}

// NewMenuToolButtonFromStock creates a new MenuToolButton. The new
// MenuToolButton will contain an icon and label from the stock item indicated
// by @stock_id.
//
// Deprecated: since version 3.10.
func NewMenuToolButtonFromStock(stockId string) MenuToolButton {
	var _arg1 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_tool_button_new_from_stock(_arg1)

	var _menuToolButton MenuToolButton // out

	_menuToolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MenuToolButton)

	return _menuToolButton
}

// Menu gets the Menu associated with MenuToolButton.
func (b *MenuToolButtonClass) Menu() Widget {
	var _arg0 *C.GtkMenuToolButton // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_menu_tool_button_get_menu(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

// SetArrowTooltipMarkup sets the tooltip markup text to be used as tooltip for
// the arrow button which pops up the menu. See gtk_tool_item_set_tooltip_text()
// for setting a tooltip on the whole MenuToolButton.
func (b *MenuToolButtonClass) SetArrowTooltipMarkup(markup string) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup(_arg0, _arg1)
}

// SetArrowTooltipText sets the tooltip text to be used as tooltip for the arrow
// button which pops up the menu. See gtk_tool_item_set_tooltip_text() for
// setting a tooltip on the whole MenuToolButton.
func (b *MenuToolButtonClass) SetArrowTooltipText(text string) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_tool_button_set_arrow_tooltip_text(_arg0, _arg1)
}

// SetMenu sets the Menu that is popped up when the user clicks on the arrow. If
// @menu is NULL, the arrow button becomes insensitive.
func (b *MenuToolButtonClass) SetMenu(menu Widget) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(menu.Native()))

	C.gtk_menu_tool_button_set_menu(_arg0, _arg1)
}
