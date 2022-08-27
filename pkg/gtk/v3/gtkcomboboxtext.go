// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
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
import "C"

// GType values.
var (
	GTypeComboBoxText = coreglib.Type(C.gtk_combo_box_text_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeComboBoxText, F: marshalComboBoxText},
	})
}

// ComboBoxTextOverrides contains methods that are overridable.
type ComboBoxTextOverrides struct {
}

func defaultComboBoxTextOverrides(v *ComboBoxText) ComboBoxTextOverrides {
	return ComboBoxTextOverrides{}
}

// ComboBoxText is a simple variant of ComboBox that hides the model-view
// complexity for simple text-only use cases.
//
// To create a GtkComboBoxText, use gtk_combo_box_text_new() or
// gtk_combo_box_text_new_with_entry().
//
// You can add items to a GtkComboBoxText with gtk_combo_box_text_append_text(),
// gtk_combo_box_text_insert_text() or gtk_combo_box_text_prepend_text() and
// remove options with gtk_combo_box_text_remove().
//
// If the GtkComboBoxText contains an entry (via the “has-entry” property), its
// contents can be retrieved using gtk_combo_box_text_get_active_text(). The
// entry itself can be accessed by calling gtk_bin_get_child() on the combo box.
//
// You should not call gtk_combo_box_set_model() or attempt to pack more cells
// into this combo box via its GtkCellLayout interface.
//
//
// GtkComboBoxText as GtkBuildable
//
// The GtkComboBoxText implementation of the GtkBuildable interface supports
// adding items directly using the <items> element and specifying <item>
// elements for each item. Each <item> element can specify the “id”
// corresponding to the appended text and also supports the regular translation
// attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying GtkComboBoxText items:
//
//    <object class="GtkComboBoxText">
//      <items>
//        <item translatable="yes" id="factory">Factory</item>
//        <item translatable="yes" id="home">Home</item>
//        <item translatable="yes" id="subway">Subway</item>
//      </items>
//    </object>
//
// CSS nodes
//
//    combobox
//    ╰── box.linked
//        ├── entry.combo
//        ├── button.combo
//        ╰── window.popup
//
// GtkComboBoxText has a single CSS node with name combobox. It adds the style
// class .combo to the main CSS nodes of its entry and button children, and the
// .linked class to the node of its internal box.
type ComboBoxText struct {
	_ [0]func() // equal guard
	ComboBox
}

var (
	_ Binner            = (*ComboBoxText)(nil)
	_ coreglib.Objector = (*ComboBoxText)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ComboBoxText, *ComboBoxTextClass, ComboBoxTextOverrides](
		GTypeComboBoxText,
		initComboBoxTextClass,
		wrapComboBoxText,
		defaultComboBoxTextOverrides,
	)
}

func initComboBoxTextClass(gclass unsafe.Pointer, overrides ComboBoxTextOverrides, classInitFunc func(*ComboBoxTextClass)) {
	if classInitFunc != nil {
		class := (*ComboBoxTextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapComboBoxText(obj *coreglib.Object) *ComboBoxText {
	return &ComboBoxText{
		ComboBox: ComboBox{
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
			Object: obj,
			CellEditable: CellEditable{
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
			CellLayout: CellLayout{
				Object: obj,
			},
		},
	}
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	return wrapComboBoxText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ComboBoxTextClass: instance of this type is always passed by reference.
type ComboBoxTextClass struct {
	*comboBoxTextClass
}

// comboBoxTextClass is the struct that's finalized.
type comboBoxTextClass struct {
	native *C.GtkComboBoxTextClass
}

func (c *ComboBoxTextClass) ParentClass() *ComboBoxClass {
	valptr := &c.native.parent_class
	var _v *ComboBoxClass // out
	_v = (*ComboBoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
