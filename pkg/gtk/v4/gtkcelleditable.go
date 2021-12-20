// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_editable_get_type()), F: marshalCellEditabler},
	})
}

// CellEditableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellEditableOverrider interface {
	// EditingDone emits the CellEditable::editing-done signal.
	EditingDone()
	// RemoveWidget emits the CellEditable::remove-widget signal.
	RemoveWidget()
	// StartEditing begins editing on a cell_editable.
	//
	// The CellRenderer for the cell creates and returns a CellEditable from
	// gtk_cell_renderer_start_editing(), configured for the CellRenderer type.
	//
	// gtk_cell_editable_start_editing() can then set up cell_editable suitably
	// for editing a cell, e.g. making the Esc key emit
	// CellEditable::editing-done.
	//
	// Note that the cell_editable is created on-demand for the current edit;
	// its lifetime is temporary and does not persist across other edits and/or
	// cells.
	//
	// The function takes the following parameters:
	//
	//    - event (optional) that began the editing process, or NULL if editing
	//      was initiated programmatically.
	//
	StartEditing(event gdk.Eventer)
}

// CellEditable: interface for widgets that can be used for editing cells
//
// The CellEditable interface must be implemented for widgets to be usable to
// edit the contents of a TreeView cell. It provides a way to specify how
// temporary widgets should be configured for editing, get the new value, etc.
type CellEditable struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*CellEditable)(nil)
)

// CellEditabler describes CellEditable's interface methods.
type CellEditabler interface {
	externglib.Objector

	// EditingDone emits the CellEditable::editing-done signal.
	EditingDone()
	// RemoveWidget emits the CellEditable::remove-widget signal.
	RemoveWidget()
	// StartEditing begins editing on a cell_editable.
	StartEditing(event gdk.Eventer)
}

var _ CellEditabler = (*CellEditable)(nil)

func wrapCellEditable(obj *externglib.Object) *CellEditable {
	return &CellEditable{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalCellEditabler(p uintptr) (interface{}, error) {
	return wrapCellEditable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEditingDone: this signal is a sign for the cell renderer to update its
// value from the cell_editable.
//
// Implementations of CellEditable are responsible for emitting this signal when
// they are done editing, e.g. Entry emits this signal when the user presses
// Enter. Typical things to do in a handler for ::editing-done are to capture
// the edited value, disconnect the cell_editable from signals on the
// CellRenderer, etc.
//
// gtk_cell_editable_editing_done() is a convenience method for emitting
// CellEditable::editing-done.
func (cellEditable *CellEditable) ConnectEditingDone(f func()) externglib.SignalHandle {
	return cellEditable.Connect("editing-done", f)
}

// ConnectRemoveWidget: this signal is meant to indicate that the cell is
// finished editing, and the cell_editable widget is being removed and may
// subsequently be destroyed.
//
// Implementations of CellEditable are responsible for emitting this signal when
// they are done editing. It must be emitted after the
// CellEditable::editing-done signal, to give the cell renderer a chance to
// update the cell's value before the widget is removed.
//
// gtk_cell_editable_remove_widget() is a convenience method for emitting
// CellEditable::remove-widget.
func (cellEditable *CellEditable) ConnectRemoveWidget(f func()) externglib.SignalHandle {
	return cellEditable.Connect("remove-widget", f)
}

// EditingDone emits the CellEditable::editing-done signal.
func (cellEditable *CellEditable) EditingDone() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(cellEditable.Native()))

	C.gtk_cell_editable_editing_done(_arg0)
	runtime.KeepAlive(cellEditable)
}

// RemoveWidget emits the CellEditable::remove-widget signal.
func (cellEditable *CellEditable) RemoveWidget() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(cellEditable.Native()))

	C.gtk_cell_editable_remove_widget(_arg0)
	runtime.KeepAlive(cellEditable)
}

// StartEditing begins editing on a cell_editable.
//
// The CellRenderer for the cell creates and returns a CellEditable from
// gtk_cell_renderer_start_editing(), configured for the CellRenderer type.
//
// gtk_cell_editable_start_editing() can then set up cell_editable suitably for
// editing a cell, e.g. making the Esc key emit CellEditable::editing-done.
//
// Note that the cell_editable is created on-demand for the current edit; its
// lifetime is temporary and does not persist across other edits and/or cells.
//
// The function takes the following parameters:
//
//    - event (optional) that began the editing process, or NULL if editing was
//      initiated programmatically.
//
func (cellEditable *CellEditable) StartEditing(event gdk.Eventer) {
	var _arg0 *C.GtkCellEditable // out
	var _arg1 *C.GdkEvent        // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(cellEditable.Native()))
	if event != nil {
		_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	}

	C.gtk_cell_editable_start_editing(_arg0, _arg1)
	runtime.KeepAlive(cellEditable)
	runtime.KeepAlive(event)
}
