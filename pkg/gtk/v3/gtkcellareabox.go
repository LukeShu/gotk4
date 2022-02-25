// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkcellareabox.go.
var GTypeCellAreaBox = externglib.Type(C.gtk_cell_area_box_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCellAreaBox, F: marshalCellAreaBox},
	})
}

// CellAreaBoxOverrider contains methods that are overridable.
type CellAreaBoxOverrider interface {
}

// CellAreaBox renders cell renderers into a row or a column depending on its
// Orientation.
//
// GtkCellAreaBox uses a notion of packing. Packing refers to adding cell
// renderers with reference to a particular position in a CellAreaBox. There are
// two reference positions: the start and the end of the box. When the
// CellAreaBox is oriented in the GTK_ORIENTATION_VERTICAL orientation, the
// start is defined as the top of the box and the end is defined as the bottom.
// In the GTK_ORIENTATION_HORIZONTAL orientation start is defined as the left
// side and the end is defined as the right side.
//
// Alignments of CellRenderers rendered in adjacent rows can be configured by
// configuring the CellAreaBox align child cell property with
// gtk_cell_area_cell_set_property() or by specifying the "align" argument to
// gtk_cell_area_box_pack_start() and gtk_cell_area_box_pack_end().
type CellAreaBox struct {
	_ [0]func() // equal guard
	CellArea

	*externglib.Object
	Orientable
}

var (
	_ CellAreaer          = (*CellAreaBox)(nil)
	_ externglib.Objector = (*CellAreaBox)(nil)
)

func classInitCellAreaBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCellAreaBox(obj *externglib.Object) *CellAreaBox {
	return &CellAreaBox{
		CellArea: CellArea{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Buildable: Buildable{
				Object: obj,
			},
			CellLayout: CellLayout{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellAreaBox(p uintptr) (interface{}, error) {
	return wrapCellAreaBox(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellAreaBox creates a new CellAreaBox.
//
// The function returns the following values:
//
//    - cellAreaBox: newly created CellAreaBox.
//
func NewCellAreaBox() *CellAreaBox {
	var _cret *C.GtkCellArea // in

	_cret = C.gtk_cell_area_box_new()

	var _cellAreaBox *CellAreaBox // out

	_cellAreaBox = wrapCellAreaBox(externglib.Take(unsafe.Pointer(_cret)))

	return _cellAreaBox
}

// Spacing gets the spacing added between cell renderers.
//
// The function returns the following values:
//
//    - gint: space added between cell renderers in box.
//
func (box *CellAreaBox) Spacing() int {
	var _arg0 *C.GtkCellAreaBox // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(externglib.InternObject(box).Native()))

	_cret = C.gtk_cell_area_box_get_spacing(_arg0)
	runtime.KeepAlive(box)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PackEnd adds renderer to box, packed with reference to the end of box.
//
// The renderer is packed after (away from end of) any other CellRenderer packed
// with reference to the end of box.
//
// The function takes the following parameters:
//
//    - renderer to add.
//    - expand: whether renderer should receive extra space when the area
//      receives more than its natural size.
//    - align: whether renderer should be aligned in adjacent rows.
//    - fixed: whether renderer should have the same size in all rows.
//
func (box *CellAreaBox) PackEnd(renderer CellRendererer, expand, align, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(externglib.InternObject(box).Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_end(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(box)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(align)
	runtime.KeepAlive(fixed)
}

// PackStart adds renderer to box, packed with reference to the start of box.
//
// The renderer is packed after any other CellRenderer packed with reference to
// the start of box.
//
// The function takes the following parameters:
//
//    - renderer to add.
//    - expand: whether renderer should receive extra space when the area
//      receives more than its natural size.
//    - align: whether renderer should be aligned in adjacent rows.
//    - fixed: whether renderer should have the same size in all rows.
//
func (box *CellAreaBox) PackStart(renderer CellRendererer, expand, align, fixed bool) {
	var _arg0 *C.GtkCellAreaBox  // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out
	var _arg3 C.gboolean         // out
	var _arg4 C.gboolean         // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(externglib.InternObject(box).Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(externglib.InternObject(renderer).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if align {
		_arg3 = C.TRUE
	}
	if fixed {
		_arg4 = C.TRUE
	}

	C.gtk_cell_area_box_pack_start(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(box)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(align)
	runtime.KeepAlive(fixed)
}

// SetSpacing sets the spacing to add between cell renderers in box.
//
// The function takes the following parameters:
//
//    - spacing: space to add between CellRenderers.
//
func (box *CellAreaBox) SetSpacing(spacing int) {
	var _arg0 *C.GtkCellAreaBox // out
	var _arg1 C.gint            // out

	_arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(externglib.InternObject(box).Native()))
	_arg1 = C.gint(spacing)

	C.gtk_cell_area_box_set_spacing(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}
