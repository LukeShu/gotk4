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
import "C"

// NewMenuItemFromModel creates a Item as an exact copy of an existing menu item
// in a Model.
//
// item_index must be valid (ie: be sure to call g_menu_model_get_n_items()
// first).
//
// The function takes the following parameters:
//
//    - model: Model.
//    - itemIndex: index of an item in model.
//
// The function returns the following values:
//
//    - menuItem: new Item.
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

// AttributeValue queries the named attribute on menu_item.
//
// If expected_type is specified and the attribute does not have this type, NULL
// is returned. NULL is also returned if the attribute simply does not exist.
//
// The function takes the following parameters:
//
//    - attribute name to query.
//    - expectedType (optional): expected type of the attribute.
//
// The function returns the following values:
//
//    - variant (optional): attribute value, or NULL.
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
//    - link name to query.
//
// The function returns the following values:
//
//    - menuModel (optional): link, or NULL.
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
