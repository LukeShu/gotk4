// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkmenubutton.go.
var GTypeMenuButton = externglib.Type(C.gtk_menu_button_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeMenuButton, F: marshalMenuButton},
	})
}

// MenuButtonOverrider contains methods that are overridable.
type MenuButtonOverrider interface {
	externglib.Objector
}

// WrapMenuButtonOverrider wraps the MenuButtonOverrider
// interface implementation to access the instance methods.
func WrapMenuButtonOverrider(obj MenuButtonOverrider) *MenuButton {
	return wrapMenuButton(externglib.BaseObject(obj))
}

// MenuButton widget is used to display a popup when clicked on. This popup can
// be provided either as a Menu, a Popover or an abstract Model.
//
// The MenuButton widget can hold any valid child widget. That is, it can hold
// almost any other standard Widget. The most commonly used child is Image. If
// no widget is explicitely added to the MenuButton, a Image is automatically
// created, using an arrow image oriented according to MenuButton:direction or
// the generic “open-menu-symbolic” icon if the direction is not set.
//
// The positioning of the popup is determined by the MenuButton:direction
// property of the menu button.
//
// For menus, the Widget:halign and Widget:valign properties of the menu are
// also taken into account. For example, when the direction is GTK_ARROW_DOWN
// and the horizontal alignment is GTK_ALIGN_START, the menu will be positioned
// below the button, with the starting edge (depending on the text direction) of
// the menu aligned with the starting edge of the button. If there is not enough
// space below the button, the menu is popped up above the button instead. If
// the alignment would move part of the menu offscreen, it is “pushed in”.
//
// Direction = Down
//
// - halign = start
//
//    ! (down-start.png)
//
// - halign = center
//
//    ! (down-center.png)
//
// - halign = end
//
//    ! (down-end.png)
//
// Direction = Up
//
// - halign = start
//
//    ! (up-start.png)
//
// - halign = center
//
//    ! (up-center.png)
//
// - halign = end
//
//    ! (up-end.png)
//
// Direction = Left
//
// - valign = start
//
//    ! (left-start.png)
//
// - valign = center
//
//    ! (left-center.png)
//
// - valign = end
//
//    ! (left-end.png)
//
// Direction = Right
//
// - valign = start
//
//    ! (right-start.png)
//
// - valign = center
//
//    ! (right-center.png)
//
// - valign = end
//
//    ! (right-end.png)
//
//
// CSS nodes
//
// GtkMenuButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .popup style class.
type MenuButton struct {
	_ [0]func() // equal guard
	ToggleButton
}

var (
	_ Binner              = (*MenuButton)(nil)
	_ externglib.Objector = (*MenuButton)(nil)
)

func classInitMenuButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMenuButton(obj *externglib.Object) *MenuButton {
	return &MenuButton{
		ToggleButton: ToggleButton{
			Button: Button{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
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
				Object: obj,
				Actionable: Actionable{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
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
				Activatable: Activatable{
					Object: obj,
				},
			},
		},
	}
}

func marshalMenuButton(p uintptr) (interface{}, error) {
	return wrapMenuButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMenuButton creates a new MenuButton widget with downwards-pointing arrow
// as the only child. You can replace the child widget with another Widget
// should you wish to.
//
// The function returns the following values:
//
//    - menuButton: newly created MenuButton widget.
//
func NewMenuButton() *MenuButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_button_new()

	var _menuButton *MenuButton // out

	_menuButton = wrapMenuButton(externglib.Take(unsafe.Pointer(_cret)))

	return _menuButton
}

// AlignWidget returns the parent Widget to use to line up with menu.
//
// The function returns the following values:
//
//    - widget (optional) value or NULL.
//
func (menuButton *MenuButton) AlignWidget() Widgetter {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))

	_cret = C.gtk_menu_button_get_align_widget(_arg0)
	runtime.KeepAlive(menuButton)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Direction returns the direction the popup will be pointing at when popped up.
//
// The function returns the following values:
//
//    - arrowType: ArrowType value.
//
func (menuButton *MenuButton) Direction() ArrowType {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.GtkArrowType   // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))

	_cret = C.gtk_menu_button_get_direction(_arg0)
	runtime.KeepAlive(menuButton)

	var _arrowType ArrowType // out

	_arrowType = ArrowType(_cret)

	return _arrowType
}

// MenuModel returns the Model used to generate the popup.
//
// The function returns the following values:
//
//    - menuModel (optional) or NULL.
//
func (menuButton *MenuButton) MenuModel() gio.MenuModeller {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GMenuModel    // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))

	_cret = C.gtk_menu_button_get_menu_model(_arg0)
	runtime.KeepAlive(menuButton)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// Popover returns the Popover that pops out of the button. If the button is not
// using a Popover, this function returns NULL.
//
// The function returns the following values:
//
//    - popover (optional) or NULL.
//
func (menuButton *MenuButton) Popover() *Popover {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GtkPopover    // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))

	_cret = C.gtk_menu_button_get_popover(_arg0)
	runtime.KeepAlive(menuButton)

	var _popover *Popover // out

	if _cret != nil {
		_popover = wrapPopover(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _popover
}

// Popup returns the Menu that pops out of the button. If the button does not
// use a Menu, this function returns NULL.
//
// The function returns the following values:
//
//    - menu (optional) or NULL.
//
func (menuButton *MenuButton) Popup() *Menu {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GtkMenu       // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))

	_cret = C.gtk_menu_button_get_popup(_arg0)
	runtime.KeepAlive(menuButton)

	var _menu *Menu // out

	if _cret != nil {
		_menu = wrapMenu(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _menu
}

// UsePopover returns whether a Popover or a Menu will be constructed from the
// menu model.
//
// The function returns the following values:
//
//    - ok: TRUE if using a Popover.
//
func (menuButton *MenuButton) UsePopover() bool {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))

	_cret = C.gtk_menu_button_get_use_popover(_arg0)
	runtime.KeepAlive(menuButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAlignWidget sets the Widget to use to line the menu with when popped up.
// Note that the align_widget must contain the MenuButton itself.
//
// Setting it to NULL means that the menu will be aligned with the button
// itself.
//
// Note that this property is only used with menus currently, and not for
// popovers.
//
// The function takes the following parameters:
//
//    - alignWidget (optional): Widget.
//
func (menuButton *MenuButton) SetAlignWidget(alignWidget Widgetter) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))
	if alignWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(alignWidget).Native()))
	}

	C.gtk_menu_button_set_align_widget(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(alignWidget)
}

// SetDirection sets the direction in which the popup will be popped up, as well
// as changing the arrow’s direction. The child will not be changed to an arrow
// if it was customized.
//
// If the does not fit in the available space in the given direction, GTK+ will
// its best to keep it inside the screen and fully visible.
//
// If you pass GTK_ARROW_NONE for a direction, the popup will behave as if you
// passed GTK_ARROW_DOWN (although you won’t see any arrows).
//
// The function takes the following parameters:
//
//    - direction: ArrowType.
//
func (menuButton *MenuButton) SetDirection(direction ArrowType) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 C.GtkArrowType   // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))
	_arg1 = C.GtkArrowType(direction)

	C.gtk_menu_button_set_direction(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(direction)
}

// SetMenuModel sets the Model from which the popup will be constructed, or NULL
// to dissociate any existing menu model and disable the button.
//
// Depending on the value of MenuButton:use-popover, either a Menu will be
// created with gtk_menu_new_from_model(), or a Popover with
// gtk_popover_new_from_model(). In either case, actions will be connected as
// documented for these functions.
//
// If MenuButton:popup or MenuButton:popover are already set, those widgets are
// dissociated from the menu_button, and those properties are set to NULL.
//
// The function takes the following parameters:
//
//    - menuModel (optional) or NULL to unset and disable the button.
//
func (menuButton *MenuButton) SetMenuModel(menuModel gio.MenuModeller) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GMenuModel    // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))
	if menuModel != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(externglib.InternObject(menuModel).Native()))
	}

	C.gtk_menu_button_set_menu_model(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(menuModel)
}

// SetPopover sets the Popover that will be popped up when the menu_button is
// clicked, or NULL to dissociate any existing popover and disable the button.
//
// If MenuButton:menu-model or MenuButton:popup are set, those objects are
// dissociated from the menu_button, and those properties are set to NULL.
//
// The function takes the following parameters:
//
//    - popover (optional) or NULL to unset and disable the button.
//
func (menuButton *MenuButton) SetPopover(popover Widgetter) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))
	if popover != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(popover).Native()))
	}

	C.gtk_menu_button_set_popover(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(popover)
}

// SetPopup sets the Menu that will be popped up when the menu_button is
// clicked, or NULL to dissociate any existing menu and disable the button.
//
// If MenuButton:menu-model or MenuButton:popover are set, those objects are
// dissociated from the menu_button, and those properties are set to NULL.
//
// The function takes the following parameters:
//
//    - menu (optional) or NULL to unset and disable the button.
//
func (menuButton *MenuButton) SetPopup(menu Widgetter) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))
	if menu != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(menu).Native()))
	}

	C.gtk_menu_button_set_popup(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(menu)
}

// SetUsePopover sets whether to construct a Popover instead of Menu when
// gtk_menu_button_set_menu_model() is called. Note that this property is only
// consulted when a new menu model is set.
//
// The function takes the following parameters:
//
//    - usePopover: TRUE to construct a popover from the menu model.
//
func (menuButton *MenuButton) SetUsePopover(usePopover bool) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(externglib.InternObject(menuButton).Native()))
	if usePopover {
		_arg1 = C.TRUE
	}

	C.gtk_menu_button_set_use_popover(_arg0, _arg1)
	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(usePopover)
}
