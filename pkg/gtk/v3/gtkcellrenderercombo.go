// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_cell_renderer_combo_get_type()), F: marshalCellRendererCombor},
	})
}

// CellRendererCombo renders text in a cell like CellRendererText from which it
// is derived. But while CellRendererText offers a simple entry to edit the
// text, CellRendererCombo offers a ComboBox widget to edit the text. The values
// to display in the combo box are taken from the tree model specified in the
// CellRendererCombo:model property.
//
// The combo cell renderer takes care of adding a text cell renderer to the
// combo box and sets it to display the column specified by its
// CellRendererCombo:text-column property. Further properties of the combo box
// can be set in a handler for the CellRenderer::editing-started signal.
//
// The CellRendererCombo cell renderer was added in GTK+ 2.6.
type CellRendererCombo struct {
	CellRendererText
}

func wrapCellRendererCombo(obj *externglib.Object) *CellRendererCombo {
	return &CellRendererCombo{
		CellRendererText: CellRendererText{
			CellRenderer: CellRenderer{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererCombor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererCombo(obj), nil
}

// NewCellRendererCombo creates a new CellRendererCombo. Adjust how text is
// drawn using object properties. Object properties can be set globally (with
// g_object_set()). Also, with TreeViewColumn, you can bind a property to a
// value in a TreeModel. For example, you can bind the “text” property on the
// cell renderer to a string value in the model, thus rendering a different
// string in each row of the TreeView.
func NewCellRendererCombo() *CellRendererCombo {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_combo_new()

	var _cellRendererCombo *CellRendererCombo // out

	_cellRendererCombo = wrapCellRendererCombo(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererCombo
}

func (*CellRendererCombo) privateCellRendererCombo() {}
