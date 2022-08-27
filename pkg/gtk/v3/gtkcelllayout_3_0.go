// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// GtkCellArea* _gotk4_gtk3_CellLayout_virtual_get_area(void* fnptr, GtkCellLayout* arg0) {
//   return ((GtkCellArea* (*)(GtkCellLayout*))(fnptr))(arg0);
// };
import "C"

// Area returns the underlying CellArea which might be cell_layout if called on
// a CellArea or might be NULL if no CellArea is used by cell_layout.
//
// The function returns the following values:
//
//    - cellArea (optional): cell area used by cell_layout, or NULL in case no
//      cell area is used.
//
func (cellLayout *CellLayout) Area() CellAreaer {
	var _arg0 *C.GtkCellLayout // out
	var _cret *C.GtkCellArea   // in

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))

	_cret = C.gtk_cell_layout_get_area(_arg0)
	runtime.KeepAlive(cellLayout)

	var _cellArea CellAreaer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(CellAreaer)
				return ok
			})
			rv, ok := casted.(CellAreaer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellAreaer")
			}
			_cellArea = rv
		}
	}

	return _cellArea
}

// Area returns the underlying CellArea which might be cell_layout if called on
// a CellArea or might be NULL if no CellArea is used by cell_layout.
//
// The function returns the following values:
//
//    - cellArea (optional): cell area used by cell_layout, or NULL in case no
//      cell area is used.
//
func (cellLayout *CellLayout) area() CellAreaer {
	gclass := (*C.GtkCellLayoutIface)(coreglib.PeekParentClass(cellLayout))
	fnarg := gclass.get_area

	var _arg0 *C.GtkCellLayout // out
	var _cret *C.GtkCellArea   // in

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))

	_cret = C._gotk4_gtk3_CellLayout_virtual_get_area(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(cellLayout)

	var _cellArea CellAreaer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(CellAreaer)
				return ok
			})
			rv, ok := casted.(CellAreaer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellAreaer")
			}
			_cellArea = rv
		}
	}

	return _cellArea
}
