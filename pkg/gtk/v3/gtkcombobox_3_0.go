// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ActiveID returns the ID of the active row of combo_box. This value is taken
// from the active row and the column specified by the ComboBox:id-column
// property of combo_box (see gtk_combo_box_set_id_column()).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the ComboBox:id-column property of combo_box is not set, or if no row is
// active, or if the active row has a NULL ID value, then NULL is returned.
//
// The function returns the following values:
//
//    - utf8 (optional): ID of the active row, or NULL.
//
func (comboBox *ComboBox) ActiveID() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)
	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IDColumn returns the column which combo_box is using to get string IDs for
// values from.
//
// The function returns the following values:
//
//    - gint: column in the data source model of combo_box.
//
func (comboBox *ComboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PopupFixedWidth gets whether the popup uses a fixed width matching the
// allocated width of the combo box.
//
// The function returns the following values:
//
//    - ok: TRUE if the popup uses a fixed width.
//
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupForDevice pops up the menu or dropdown list of combo_box, the popup
// window will be grabbed so only device and its associated pointer/keyboard are
// the only Devices able to send events to it.
//
// The function takes the following parameters:
//
//    - device: Device.
//
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(device)
}

// SetActiveID changes the active row of combo_box to the one that has an ID
// equal to active_id, or unsets the active row if active_id is NULL. Rows
// having a NULL ID string cannot be made active by this function.
//
// If the ComboBox:id-column property of combo_box is unset or if no row has the
// given ID then the function does nothing and returns FALSE.
//
// The function takes the following parameters:
//
//    - activeId (optional): ID of the row to select, or NULL.
//
// The function returns the following values:
//
//    - ok: TRUE if a row with a matching ID was found. If a NULL active_id was
//      given to unset the active row, the function always returns TRUE.
//
func (comboBox *ComboBox) SetActiveID(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if activeId != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(activeId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(activeId)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetIDColumn sets the model column which combo_box should use to get string
// IDs for values from. The column id_column in the model of combo_box must be
// of type G_TYPE_STRING.
//
// The function takes the following parameters:
//
//    - idColumn: column in model to get string IDs for values from.
//
func (comboBox *ComboBox) SetIDColumn(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(idColumn)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width matching the allocated width of the combo box.
//
// The function takes the following parameters:
//
//    - fixed: whether to use a fixed popup width.
//
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fixed)
}
