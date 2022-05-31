// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkflattenlistmodel.go.
var GTypeFlattenListModel = coreglib.Type(C.gtk_flatten_list_model_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFlattenListModel, F: marshalFlattenListModel},
	})
}

// FlattenListModelOverrider contains methods that are overridable.
type FlattenListModelOverrider interface {
}

// FlattenListModel: GtkFlattenListModel is a list model that concatenates other
// list models.
//
// GtkFlattenListModel takes a list model containing list models, and flattens
// it into a single model.
type FlattenListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*FlattenListModel)(nil)
)

func classInitFlattenListModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFlattenListModel(obj *coreglib.Object) *FlattenListModel {
	return &FlattenListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalFlattenListModel(p uintptr) (interface{}, error) {
	return wrapFlattenListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFlattenListModel creates a new GtkFlattenListModel that flattens list.
//
// The function takes the following parameters:
//
//    - model (optional) to be flattened.
//
// The function returns the following values:
//
//    - flattenListModel: new GtkFlattenListModel.
//
func NewFlattenListModel(model gio.ListModeller) *FlattenListModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	if model != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}
	*(*gio.ListModeller)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FlattenListModel").InvokeMethod("new_FlattenListModel", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _flattenListModel *FlattenListModel // out

	_flattenListModel = wrapFlattenListModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _flattenListModel
}

// Model gets the model set via gtk_flatten_list_model_set_model().
//
// The function returns the following values:
//
//    - listModel (optional): model flattened by self.
//
func (self *FlattenListModel) Model() *gio.ListModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**FlattenListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "FlattenListModel").InvokeMethod("get_model", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listModel = &gio.ListModel{
				Object: obj,
			}
		}
	}

	return _listModel
}

// ModelForItem returns the model containing the item at the given position.
//
// The function takes the following parameters:
//
//    - position: position.
//
// The function returns the following values:
//
//    - listModel: model containing the item at position.
//
func (self *FlattenListModel) ModelForItem(position uint32) *gio.ListModel {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)
	*(**FlattenListModel)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "FlattenListModel").InvokeMethod("get_model_for_item", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// SetModel sets a new model to be flattened.
//
// The function takes the following parameters:
//
//    - model (optional): new model or NULL.
//
func (self *FlattenListModel) SetModel(model gio.ListModeller) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	*(**FlattenListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "FlattenListModel").InvokeMethod("set_model", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
