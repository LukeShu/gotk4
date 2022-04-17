// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkselectionfiltermodel.go.
var GTypeSelectionFilterModel = externglib.Type(C.gtk_selection_filter_model_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSelectionFilterModel, F: marshalSelectionFilterModel},
	})
}

// SelectionFilterModelOverrider contains methods that are overridable.
type SelectionFilterModelOverrider interface {
	externglib.Objector
}

// WrapSelectionFilterModelOverrider wraps the SelectionFilterModelOverrider
// interface implementation to access the instance methods.
func WrapSelectionFilterModelOverrider(obj SelectionFilterModelOverrider) *SelectionFilterModel {
	return wrapSelectionFilterModel(externglib.BaseObject(obj))
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

func classInitSelectionFilterModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSelectionFilterModel(obj *externglib.Object) *SelectionFilterModel {
	return &SelectionFilterModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSelectionFilterModel(p uintptr) (interface{}, error) {
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
func NewSelectionFilterModel(model SelectionModelOverrider) *SelectionFilterModel {
	var _arg1 *C.GtkSelectionModel       // out
	var _cret *C.GtkSelectionFilterModel // in

	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(externglib.InternObject(model).Native()))
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
func (self *SelectionFilterModel) Model() SelectionModelOverrider {
	var _arg0 *C.GtkSelectionFilterModel // out
	var _cret *C.GtkSelectionModel       // in

	_arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_selection_filter_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel SelectionModelOverrider // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(SelectionModelOverrider)
				return ok
			})
			rv, ok := casted.(SelectionModelOverrider)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.SelectionModeller")
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
func (self *SelectionFilterModel) SetModel(model SelectionModelOverrider) {
	var _arg0 *C.GtkSelectionFilterModel // out
	var _arg1 *C.GtkSelectionModel       // out

	_arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(externglib.InternObject(model).Native()))
	}

	C.gtk_selection_filter_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
