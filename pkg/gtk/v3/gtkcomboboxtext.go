// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeComboBoxText returns the GType for the type ComboBoxText.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeComboBoxText() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ComboBoxText").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalComboBoxText)
	return gtype
}

// ComboBoxTextOverrider contains methods that are overridable.
type ComboBoxTextOverrider interface {
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

func classInitComboBoxTexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

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

// NewComboBoxText creates a new ComboBoxText, which is a ComboBox just
// displaying strings.
//
// The function returns the following values:
//
//    - comboBoxText: new ComboBoxText.
//
func NewComboBoxText() *ComboBoxText {
	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_gret := _info.InvokeClassMethod("new_ComboBoxText", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _comboBoxText *ComboBoxText // out

	_comboBoxText = wrapComboBoxText(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _comboBoxText
}

// NewComboBoxTextWithEntry creates a new ComboBoxText, which is a ComboBox just
// displaying strings. The combo box created by this function has an entry.
//
// The function returns the following values:
//
//    - comboBoxText: new ComboBoxText.
//
func NewComboBoxTextWithEntry() *ComboBoxText {
	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_gret := _info.InvokeClassMethod("new_ComboBoxText_with_entry", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _comboBoxText *ComboBoxText // out

	_comboBoxText = wrapComboBoxText(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _comboBoxText
}

// Append appends text to the list of strings stored in combo_box. If id is
// non-NULL then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a position of
// -1.
//
// The function takes the following parameters:
//
//    - id (optional): string ID for this value, or NULL.
//    - text: string.
//
func (comboBox *ComboBoxText) Append(id, text string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if id != "" {
		*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))
	}
	*(**C.gchar)(unsafe.Pointer(&_args[2])) = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[2]))))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("append", _args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// AppendText appends text to the list of strings stored in combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a position
// of -1.
//
// The function takes the following parameters:
//
//    - text: string.
//
func (comboBox *ComboBoxText) AppendText(text string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("append_text", _args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(text)
}

// ActiveText returns the currently active string in combo_box, or NULL if none
// is selected. If combo_box contains an entry, this function will return its
// contents (which will not necessarily be an item from the list).
//
// The function returns the following values:
//
//    - utf8: newly allocated string containing the currently active text. Must
//      be freed with g_free().
//
func (comboBox *ComboBoxText) ActiveText() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_gret := _info.InvokeClassMethod("get_active_text", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret))))

	return _utf8
}

// Insert inserts text at position in the list of strings stored in combo_box.
// If id is non-NULL then it is used as the ID of the row. See
// ComboBox:id-column.
//
// If position is negative then text is appended.
//
// The function takes the following parameters:
//
//    - position: index to insert text.
//    - id (optional): string ID for this value, or NULL.
//    - text: string to display.
//
func (comboBox *ComboBoxText) Insert(position int32, id, text string) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(position)
	if id != "" {
		*(**C.gchar)(unsafe.Pointer(&_args[2])) = (*C.gchar)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[2]))))
	}
	*(**C.gchar)(unsafe.Pointer(&_args[3])) = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[3]))))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("insert", _args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// InsertText inserts text at position in the list of strings stored in
// combo_box.
//
// If position is negative then text is appended.
//
// This is the same as calling gtk_combo_box_text_insert() with a NULL ID
// string.
//
// The function takes the following parameters:
//
//    - position: index to insert text.
//    - text: string.
//
func (comboBox *ComboBoxText) InsertText(position int32, text string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(position)
	*(**C.gchar)(unsafe.Pointer(&_args[2])) = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[2]))))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("insert_text", _args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
	runtime.KeepAlive(text)
}

// Prepend prepends text to the list of strings stored in combo_box. If id is
// non-NULL then it is used as the ID of the row.
//
// This is the same as calling gtk_combo_box_text_insert() with a position of 0.
//
// The function takes the following parameters:
//
//    - id (optional): string ID for this value, or NULL.
//    - text: string.
//
func (comboBox *ComboBoxText) Prepend(id, text string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if id != "" {
		*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))
	}
	*(**C.gchar)(unsafe.Pointer(&_args[2])) = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[2]))))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("prepend", _args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// PrependText prepends text to the list of strings stored in combo_box.
//
// This is the same as calling gtk_combo_box_text_insert_text() with a position
// of 0.
//
// The function takes the following parameters:
//
//    - text: string.
//
func (comboBox *ComboBoxText) PrependText(text string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("prepend_text", _args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(text)
}

// Remove removes the string at position from combo_box.
//
// The function takes the following parameters:
//
//    - position: index of the item to remove.
//
func (comboBox *ComboBoxText) Remove(position int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(position)

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("remove", _args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
}

// RemoveAll removes all the text entries from the combo box.
func (comboBox *ComboBoxText) RemoveAll() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_info := girepository.MustFind("Gtk", "ComboBoxText")
	_info.InvokeClassMethod("remove_all", _args[:], nil)

	runtime.KeepAlive(comboBox)
}
