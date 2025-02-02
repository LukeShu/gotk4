// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeMenu     = coreglib.Type(C.g_menu_get_type())
	GTypeMenuItem = coreglib.Type(C.g_menu_item_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMenu, F: marshalMenu},
		coreglib.TypeMarshaler{T: GTypeMenuItem, F: marshalMenuItem},
	})
}

// ResourceErrorQuark gets the #GResource Error Quark.
//
// The function returns the following values:
//
//   - quark: #GQuark.
//
func ResourceErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.g_resource_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

// Menu is a simple implementation of Model. You populate a #GMenu by adding
// Item instances to it.
//
// There are some convenience functions to allow you to directly add
// items (avoiding Item) for the common cases. To add a regular item,
// use g_menu_insert(). To add a section, use g_menu_insert_section(). To add a
// submenu, use g_menu_insert_submenu().
type Menu struct {
	_ [0]func() // equal guard
	MenuModel
}

var (
	_ MenuModeller = (*Menu)(nil)
)

func wrapMenu(obj *coreglib.Object) *Menu {
	return &Menu{
		MenuModel: MenuModel{
			Object: obj,
		},
	}
}

func marshalMenu(p uintptr) (interface{}, error) {
	return wrapMenu(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMenu creates a new #GMenu.
//
// The new menu has no items.
//
// The function returns the following values:
//
//   - menu: new #GMenu.
//
func NewMenu() *Menu {
	var _cret *C.GMenu // in

	_cret = C.g_menu_new()

	var _menu *Menu // out

	_menu = wrapMenu(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _menu
}

// Append: convenience function for appending a normal menu item to the end of
// menu. Combine g_menu_item_new() and g_menu_insert_item() for a more flexible
// alternative.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - detailedAction (optional): detailed action string, or NULL.
//
func (menu *Menu) Append(label, detailedAction string) {
	var _arg0 *C.GMenu // out
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if detailedAction != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.g_menu_append(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(label)
	runtime.KeepAlive(detailedAction)
}

// AppendItem appends item to the end of menu.
//
// See g_menu_insert_item() for more information.
//
// The function takes the following parameters:
//
//   - item to append.
//
func (menu *Menu) AppendItem(item *MenuItem) {
	var _arg0 *C.GMenu     // out
	var _arg1 *C.GMenuItem // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	C.g_menu_append_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(item)
}

// AppendSection: convenience function for appending a section menu item to the
// end of menu. Combine g_menu_item_new_section() and g_menu_insert_item() for a
// more flexible alternative.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - section with the items of the section.
//
func (menu *Menu) AppendSection(label string, section MenuModeller) {
	var _arg0 *C.GMenu      // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GMenuModel // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(section).Native()))

	C.g_menu_append_section(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(label)
	runtime.KeepAlive(section)
}

// AppendSubmenu: convenience function for appending a submenu menu item to the
// end of menu. Combine g_menu_item_new_submenu() and g_menu_insert_item() for a
// more flexible alternative.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - submenu with the items of the submenu.
//
func (menu *Menu) AppendSubmenu(label string, submenu MenuModeller) {
	var _arg0 *C.GMenu      // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GMenuModel // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))

	C.g_menu_append_submenu(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(label)
	runtime.KeepAlive(submenu)
}

// Freeze marks menu as frozen.
//
// After the menu is frozen, it is an error to attempt to make any changes to
// it. In effect this means that the #GMenu API must no longer be used.
//
// This function causes g_menu_model_is_mutable() to begin returning FALSE,
// which has some positive performance implications.
func (menu *Menu) Freeze() {
	var _arg0 *C.GMenu // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	C.g_menu_freeze(_arg0)
	runtime.KeepAlive(menu)
}

// Insert: convenience function for inserting a normal menu item into menu.
// Combine g_menu_item_new() and g_menu_insert_item() for a more flexible
// alternative.
//
// The function takes the following parameters:
//
//   - position at which to insert the item.
//   - label (optional): section label, or NULL.
//   - detailedAction (optional): detailed action string, or NULL.
//
func (menu *Menu) Insert(position int, label, detailedAction string) {
	var _arg0 *C.GMenu // out
	var _arg1 C.gint   // out
	var _arg2 *C.gchar // out
	var _arg3 *C.gchar // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.gint(position)
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if detailedAction != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	C.g_menu_insert(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(position)
	runtime.KeepAlive(label)
	runtime.KeepAlive(detailedAction)
}

// InsertItem inserts item into menu.
//
// The "insertion" is actually done by copying all of the attribute and link
// values of item and using them to form a new item within menu. As such,
// item itself is not really inserted, but rather, a menu item that is exactly
// the same as the one presently described by item.
//
// This means that item is essentially useless after the insertion occurs.
// Any changes you make to it are ignored unless it is inserted again (at which
// point its updated values will be copied).
//
// You should probably just free item once you're done.
//
// There are many convenience functions to take care of common cases. See
// g_menu_insert(), g_menu_insert_section() and g_menu_insert_submenu() as well
// as "prepend" and "append" variants of each of these functions.
//
// The function takes the following parameters:
//
//   - position at which to insert the item.
//   - item to insert.
//
func (menu *Menu) InsertItem(position int, item *MenuItem) {
	var _arg0 *C.GMenu     // out
	var _arg1 C.gint       // out
	var _arg2 *C.GMenuItem // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.gint(position)
	_arg2 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	C.g_menu_insert_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(position)
	runtime.KeepAlive(item)
}

// InsertSection: convenience function for inserting a section menu item into
// menu. Combine g_menu_item_new_section() and g_menu_insert_item() for a more
// flexible alternative.
//
// The function takes the following parameters:
//
//   - position at which to insert the item.
//   - label (optional): section label, or NULL.
//   - section with the items of the section.
//
func (menu *Menu) InsertSection(position int, label string, section MenuModeller) {
	var _arg0 *C.GMenu      // out
	var _arg1 C.gint        // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.GMenuModel // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.gint(position)
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(section).Native()))

	C.g_menu_insert_section(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(position)
	runtime.KeepAlive(label)
	runtime.KeepAlive(section)
}

// InsertSubmenu: convenience function for inserting a submenu menu item into
// menu. Combine g_menu_item_new_submenu() and g_menu_insert_item() for a more
// flexible alternative.
//
// The function takes the following parameters:
//
//   - position at which to insert the item.
//   - label (optional): section label, or NULL.
//   - submenu with the items of the submenu.
//
func (menu *Menu) InsertSubmenu(position int, label string, submenu MenuModeller) {
	var _arg0 *C.GMenu      // out
	var _arg1 C.gint        // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.GMenuModel // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.gint(position)
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))

	C.g_menu_insert_submenu(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(position)
	runtime.KeepAlive(label)
	runtime.KeepAlive(submenu)
}

// Prepend: convenience function for prepending a normal menu item to the
// start of menu. Combine g_menu_item_new() and g_menu_insert_item() for a more
// flexible alternative.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - detailedAction (optional): detailed action string, or NULL.
//
func (menu *Menu) Prepend(label, detailedAction string) {
	var _arg0 *C.GMenu // out
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if detailedAction != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.g_menu_prepend(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(label)
	runtime.KeepAlive(detailedAction)
}

// PrependItem prepends item to the start of menu.
//
// See g_menu_insert_item() for more information.
//
// The function takes the following parameters:
//
//   - item to prepend.
//
func (menu *Menu) PrependItem(item *MenuItem) {
	var _arg0 *C.GMenu     // out
	var _arg1 *C.GMenuItem // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(item).Native()))

	C.g_menu_prepend_item(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(item)
}

// PrependSection: convenience function for prepending a section menu item to
// the start of menu. Combine g_menu_item_new_section() and g_menu_insert_item()
// for a more flexible alternative.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - section with the items of the section.
//
func (menu *Menu) PrependSection(label string, section MenuModeller) {
	var _arg0 *C.GMenu      // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GMenuModel // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(section).Native()))

	C.g_menu_prepend_section(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(label)
	runtime.KeepAlive(section)
}

// PrependSubmenu: convenience function for prepending a submenu menu item to
// the start of menu. Combine g_menu_item_new_submenu() and g_menu_insert_item()
// for a more flexible alternative.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - submenu with the items of the submenu.
//
func (menu *Menu) PrependSubmenu(label string, submenu MenuModeller) {
	var _arg0 *C.GMenu      // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GMenuModel // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))

	C.g_menu_prepend_submenu(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(label)
	runtime.KeepAlive(submenu)
}

// Remove removes an item from the menu.
//
// position gives the index of the item to remove.
//
// It is an error if position is not in range the range from 0 to one less than
// the number of items in the menu.
//
// It is not possible to remove items by identity since items are added to the
// menu simply by copying their links and attributes (ie: identity of the item
// itself is not preserved).
//
// The function takes the following parameters:
//
//   - position of the item to remove.
//
func (menu *Menu) Remove(position int) {
	var _arg0 *C.GMenu // out
	var _arg1 C.gint   // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))
	_arg1 = C.gint(position)

	C.g_menu_remove(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(position)
}

// RemoveAll removes all items in the menu.
func (menu *Menu) RemoveAll() {
	var _arg0 *C.GMenu // out

	_arg0 = (*C.GMenu)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	C.g_menu_remove_all(_arg0)
	runtime.KeepAlive(menu)
}

// MenuItem is an opaque structure type. You must access it using the functions
// below.
type MenuItem struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*MenuItem)(nil)
)

func wrapMenuItem(obj *coreglib.Object) *MenuItem {
	return &MenuItem{
		Object: obj,
	}
}

func marshalMenuItem(p uintptr) (interface{}, error) {
	return wrapMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMenuItem creates a new Item.
//
// If label is non-NULL it is used to set the "label" attribute of the new item.
//
// If detailed_action is non-NULL it is used to set the "action" and possibly
// the "target" attribute of the new item. See g_menu_item_set_detailed_action()
// for more information.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - detailedAction (optional): detailed action string, or NULL.
//
// The function returns the following values:
//
//   - menuItem: new Item.
//
func NewMenuItem(label, detailedAction string) *MenuItem {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret *C.GMenuItem // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if detailedAction != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.g_menu_item_new(_arg1, _arg2)
	runtime.KeepAlive(label)
	runtime.KeepAlive(detailedAction)

	var _menuItem *MenuItem // out

	_menuItem = wrapMenuItem(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _menuItem
}

// NewMenuItemFromModel creates a Item as an exact copy of an existing menu item
// in a Model.
//
// item_index must be valid (ie: be sure to call g_menu_model_get_n_items()
// first).
//
// The function takes the following parameters:
//
//   - model: Model.
//   - itemIndex: index of an item in model.
//
// The function returns the following values:
//
//   - menuItem: new Item.
//
func NewMenuItemFromModel(model MenuModeller, itemIndex int) *MenuItem {
	var _arg1 *C.GMenuModel // out
	var _arg2 C.gint        // out
	var _cret *C.GMenuItem  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	_arg2 = C.gint(itemIndex)

	_cret = C.g_menu_item_new_from_model(_arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(itemIndex)

	var _menuItem *MenuItem // out

	_menuItem = wrapMenuItem(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _menuItem
}

// NewMenuItemSection creates a new Item representing a section.
//
// This is a convenience API around g_menu_item_new() and
// g_menu_item_set_section().
//
// The effect of having one menu appear as a section of another is exactly
// as it sounds: the items from section become a direct part of the menu that
// menu_item is added to.
//
// Visual separation is typically displayed between two non-empty sections. If
// label is non-NULL then it will be encorporated into this visual indication.
// This allows for labeled subsections of a menu.
//
// As a simple example, consider a typical "Edit" menu from a simple program.
// It probably contains an "Undo" and "Redo" item, followed by a separator,
// followed by "Cut", "Copy" and "Paste".
//
// This would be accomplished by creating three #GMenu instances. The first
// would be populated with the "Undo" and "Redo" items, and the second with the
// "Cut", "Copy" and "Paste" items. The first and second menus would then be
// added as submenus of the third. In XML format, this would look something like
// the following:
//
//    <menu id='edit-menu'>
//      <section>
//        <item label='Undo'/>
//        <item label='Redo'/>
//      </section>
//      <section>
//        <item label='Cut'/>
//        <item label='Copy'/>
//        <item label='Paste'/>
//      </section>
//    </menu>
//
// The following example is exactly equivalent. It is more illustrative of the
// exact relationship between the menus and items (keeping in mind that the
// 'link' element defines a new menu that is linked to the containing one).
// The style of the second example is more verbose and difficult to read (and
// therefore not recommended except for the purpose of understanding what is
// really going on).
//
//    <menu id='edit-menu'>
//      <item>
//        <link name='section'>
//          <item label='Undo'/>
//          <item label='Redo'/>
//        </link>
//      </item>
//      <item>
//        <link name='section'>
//          <item label='Cut'/>
//          <item label='Copy'/>
//          <item label='Paste'/>
//        </link>
//      </item>
//    </menu>.
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - section with the items of the section.
//
// The function returns the following values:
//
//   - menuItem: new Item.
//
func NewMenuItemSection(label string, section MenuModeller) *MenuItem {
	var _arg1 *C.gchar      // out
	var _arg2 *C.GMenuModel // out
	var _cret *C.GMenuItem  // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(section).Native()))

	_cret = C.g_menu_item_new_section(_arg1, _arg2)
	runtime.KeepAlive(label)
	runtime.KeepAlive(section)

	var _menuItem *MenuItem // out

	_menuItem = wrapMenuItem(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _menuItem
}

// NewMenuItemSubmenu creates a new Item representing a submenu.
//
// This is a convenience API around g_menu_item_new() and
// g_menu_item_set_submenu().
//
// The function takes the following parameters:
//
//   - label (optional): section label, or NULL.
//   - submenu with the items of the submenu.
//
// The function returns the following values:
//
//   - menuItem: new Item.
//
func NewMenuItemSubmenu(label string, submenu MenuModeller) *MenuItem {
	var _arg1 *C.gchar      // out
	var _arg2 *C.GMenuModel // out
	var _cret *C.GMenuItem  // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))

	_cret = C.g_menu_item_new_submenu(_arg1, _arg2)
	runtime.KeepAlive(label)
	runtime.KeepAlive(submenu)

	var _menuItem *MenuItem // out

	_menuItem = wrapMenuItem(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _menuItem
}

// AttributeValue queries the named attribute on menu_item.
//
// If expected_type is specified and the attribute does not have this type, NULL
// is returned. NULL is also returned if the attribute simply does not exist.
//
// The function takes the following parameters:
//
//   - attribute name to query.
//   - expectedType (optional): expected type of the attribute.
//
// The function returns the following values:
//
//   - variant (optional): attribute value, or NULL.
//
func (menuItem *MenuItem) AttributeValue(attribute string, expectedType *glib.VariantType) *glib.Variant {
	var _arg0 *C.GMenuItem    // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariantType // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	defer C.free(unsafe.Pointer(_arg1))
	if expectedType != nil {
		_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(expectedType)))
	}

	_cret = C.g_menu_item_get_attribute_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(attribute)
	runtime.KeepAlive(expectedType)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// Link queries the named link on menu_item.
//
// The function takes the following parameters:
//
//   - link name to query.
//
// The function returns the following values:
//
//   - menuModel (optional): link, or NULL.
//
func (menuItem *MenuItem) Link(link string) MenuModeller {
	var _arg0 *C.GMenuItem  // out
	var _arg1 *C.gchar      // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(link)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_menu_item_get_link(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(link)

	var _menuModel MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(MenuModeller)
				return ok
			})
			rv, ok := casted.(MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// SetActionAndTargetValue sets or unsets the "action" and "target" attributes
// of menu_item.
//
// If action is NULL then both the "action" and "target" attributes are unset
// (and target_value is ignored).
//
// If action is non-NULL then the "action" attribute is set. The "target"
// attribute is then set to the value of target_value if it is non-NULL or unset
// otherwise.
//
// Normal menu items (ie: not submenu, section or other custom item types) are
// expected to have the "action" attribute set to identify the action that
// they are associated with. The state type of the action help to determine
// the disposition of the menu item. See #GAction and Group for an overview of
// actions.
//
// In general, clicking on the menu item will result in activation of the named
// action with the "target" attribute given as the parameter to the action
// invocation. If the "target" attribute is not set then the action is invoked
// with no parameter.
//
// If the action has no state then the menu item is usually drawn as a plain
// menu item (ie: with no additional decoration).
//
// If the action has a boolean state then the menu item is usually drawn as a
// toggle menu item (ie: with a checkmark or equivalent indication). The item
// should be marked as 'toggled' or 'checked' when the boolean state is TRUE.
//
// If the action has a string state then the menu item is usually drawn as a
// radio menu item (ie: with a radio bullet or equivalent indication). The item
// should be marked as 'selected' when the string state is equal to the value of
// the target property.
//
// See g_menu_item_set_action_and_target() or g_menu_item_set_detailed_action()
// for two equivalent calls that are probably more convenient for most uses.
//
// The function takes the following parameters:
//
//   - action (optional): name of the action for this item.
//   - targetValue (optional) to use as the action target.
//
func (menuItem *MenuItem) SetActionAndTargetValue(action string, targetValue *glib.Variant) {
	var _arg0 *C.GMenuItem // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.GVariant  // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if action != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(action)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if targetValue != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(targetValue)))
	}

	C.g_menu_item_set_action_and_target_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(action)
	runtime.KeepAlive(targetValue)
}

// SetAttributeValue sets or unsets an attribute on menu_item.
//
// The attribute to set or unset is specified by attribute. This can be one of
// the standard attribute names G_MENU_ATTRIBUTE_LABEL, G_MENU_ATTRIBUTE_ACTION,
// G_MENU_ATTRIBUTE_TARGET, or a custom attribute name. Attribute names are
// restricted to lowercase characters, numbers and '-'. Furthermore, the names
// must begin with a lowercase character, must not end with a '-', and must not
// contain consecutive dashes.
//
// must consist only of lowercase ASCII characters, digits and '-'.
//
// If value is non-NULL then it is used as the new value for the attribute.
// If value is NULL then the attribute is unset. If the value #GVariant is
// floating, it is consumed.
//
// See also g_menu_item_set_attribute() for a more convenient way to do the
// same.
//
// The function takes the following parameters:
//
//   - attribute to set.
//   - value (optional) to use as the value, or NULL.
//
func (menuItem *MenuItem) SetAttributeValue(attribute string, value *glib.Variant) {
	var _arg0 *C.GMenuItem // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.GVariant  // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	defer C.free(unsafe.Pointer(_arg1))
	if value != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))
	}

	C.g_menu_item_set_attribute_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(attribute)
	runtime.KeepAlive(value)
}

// SetDetailedAction sets the "action" and possibly the "target" attribute of
// menu_item.
//
// The format of detailed_action is the same format parsed by
// g_action_parse_detailed_name().
//
// See g_menu_item_set_action_and_target() or
// g_menu_item_set_action_and_target_value() for more flexible (but slightly
// less convenient) alternatives.
//
// See also g_menu_item_set_action_and_target_value() for a description of the
// semantics of the action and target attributes.
//
// The function takes the following parameters:
//
//   - detailedAction: "detailed" action string.
//
func (menuItem *MenuItem) SetDetailedAction(detailedAction string) {
	var _arg0 *C.GMenuItem // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_menu_item_set_detailed_action(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(detailedAction)
}

// SetIcon sets (or unsets) the icon on menu_item.
//
// This call is the same as calling g_icon_serialize() and using the result as
// the value to g_menu_item_set_attribute_value() for G_MENU_ATTRIBUTE_ICON.
//
// This API is only intended for use with "noun" menu items; things like
// bookmarks or applications in an "Open With" menu. Don't use it on menu items
// corresponding to verbs (eg: stock icons for 'Save' or 'Quit').
//
// If icon is NULL then the icon is unset.
//
// The function takes the following parameters:
//
//   - icon or NULL.
//
func (menuItem *MenuItem) SetIcon(icon Iconner) {
	var _arg0 *C.GMenuItem // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	C.g_menu_item_set_icon(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(icon)
}

// SetLabel sets or unsets the "label" attribute of menu_item.
//
// If label is non-NULL it is used as the label for the menu item. If it is NULL
// then the label attribute is unset.
//
// The function takes the following parameters:
//
//   - label (optional) to set, or NULL to unset.
//
func (menuItem *MenuItem) SetLabel(label string) {
	var _arg0 *C.GMenuItem // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_menu_item_set_label(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(label)
}

// SetLink creates a link from menu_item to model if non-NULL, or unsets it.
//
// Links are used to establish a relationship between a particular menu item
// and another menu. For example, G_MENU_LINK_SUBMENU is used to associate
// a submenu with a particular menu item, and G_MENU_LINK_SECTION is used to
// create a section. Other types of link can be used, but there is no guarantee
// that clients will be able to make sense of them. Link types are restricted
// to lowercase characters, numbers and '-'. Furthermore, the names must begin
// with a lowercase character, must not end with a '-', and must not contain
// consecutive dashes.
//
// The function takes the following parameters:
//
//   - link: type of link to establish or unset.
//   - model (optional) to link to (or NULL to unset).
//
func (menuItem *MenuItem) SetLink(link string, model MenuModeller) {
	var _arg0 *C.GMenuItem  // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GMenuModel // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(link)))
	defer C.free(unsafe.Pointer(_arg1))
	if model != nil {
		_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.g_menu_item_set_link(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(link)
	runtime.KeepAlive(model)
}

// SetSection sets or unsets the "section" link of menu_item to section.
//
// The effect of having one menu appear as a section of another is exactly
// as it sounds: the items from section become a direct part of the menu that
// menu_item is added to. See g_menu_item_new_section() for more information
// about what it means for a menu item to be a section.
//
// The function takes the following parameters:
//
//   - section (optional) or NULL.
//
func (menuItem *MenuItem) SetSection(section MenuModeller) {
	var _arg0 *C.GMenuItem  // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if section != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(section).Native()))
	}

	C.g_menu_item_set_section(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(section)
}

// SetSubmenu sets or unsets the "submenu" link of menu_item to submenu.
//
// If submenu is non-NULL, it is linked to. If it is NULL then the link is
// unset.
//
// The effect of having one menu appear as a submenu of another is exactly as it
// sounds.
//
// The function takes the following parameters:
//
//   - submenu (optional) or NULL.
//
func (menuItem *MenuItem) SetSubmenu(submenu MenuModeller) {
	var _arg0 *C.GMenuItem  // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GMenuItem)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	if submenu != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(submenu).Native()))
	}

	C.g_menu_item_set_submenu(_arg0, _arg1)
	runtime.KeepAlive(menuItem)
	runtime.KeepAlive(submenu)
}
