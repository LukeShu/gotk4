// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_selection_filter_model_get_type()), F: marshalSelectionFilterModeller},
	})
}

// SelectionFilterModel: GtkSelectionFilterModel is a list model that presents
// the selection from a GtkSelectionModel.
type SelectionFilterModel struct {
	_ [0]func() // equal guard
	*externglib.Object

	gio.ListModel
}

var (
	_ externglib.Objector = (*SelectionFilterModel)(nil)
)

func wrapSelectionFilterModel(obj *externglib.Object) *SelectionFilterModel {
	return &SelectionFilterModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSelectionFilterModeller(p uintptr) (interface{}, error) {
	return wrapSelectionFilterModel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSelectionFilterModel creates a new GtkSelectionFilterModel that will
// include the selected items from the underlying selection model.
//
// The function takes the following parameters:
//
//    - model (optional): selection model to filter, or NULL.
//
// The function returns the following values:
//
//    - selectionFilterModel: new GtkSelectionFilterModel.
//
func NewSelectionFilterModel(model SelectionModeller) *SelectionFilterModel {
	var _arg1 *C.GtkSelectionModel       // out
	var _cret *C.GtkSelectionFilterModel // in

	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	}

	_cret = C.gtk_selection_filter_model_new(_arg1)
	runtime.KeepAlive(model)

	var _selectionFilterModel *SelectionFilterModel // out

	_selectionFilterModel = wrapSelectionFilterModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _selectionFilterModel
}

// Model gets the model currently filtered or NULL if none.
//
// The function returns the following values:
//
//    - selectionModel (optional): model that gets filtered.
//
func (self *SelectionFilterModel) Model() SelectionModeller {
	var _arg0 *C.GtkSelectionFilterModel // out
	var _cret *C.GtkSelectionModel       // in

	_arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_selection_filter_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel SelectionModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(SelectionModeller)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.SelectionModeller")
			}
			_selectionModel = rv
		}
	}

	return _selectionModel
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that model conforms to the item type
// of self. It assumes that the caller knows what they are doing and have set up
// an appropriate filter to ensure that item types match.
//
// The function takes the following parameters:
//
//    - model (optional) to be filtered.
//
func (self *SelectionFilterModel) SetModel(model SelectionModeller) {
	var _arg0 *C.GtkSelectionFilterModel // out
	var _arg1 *C.GtkSelectionModel       // out

	_arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_selection_filter_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
