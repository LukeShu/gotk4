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
// #include <glib-object.h>
import "C"

// GTypeFlattenListModel returns the GType for the type FlattenListModel.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFlattenListModel() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "FlattenListModel").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFlattenListModel)
	return gtype
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
	var _args [1]girepository.Argument

	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}

	_info := girepository.MustFind("Gtk", "FlattenListModel")
	_gret := _info.InvokeClassMethod("new_FlattenListModel", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _flattenListModel *FlattenListModel // out

	_flattenListModel = wrapFlattenListModel(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _flattenListModel
}

// Model gets the model set via gtk_flatten_list_model_set_model().
//
// The function returns the following values:
//
//    - listModel (optional): model flattened by self.
//
func (self *FlattenListModel) Model() *gio.ListModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "FlattenListModel")
	_gret := _info.InvokeClassMethod("get_model", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	_info := girepository.MustFind("Gtk", "FlattenListModel")
	_gret := _info.InvokeClassMethod("get_model_for_item", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	_info := girepository.MustFind("Gtk", "FlattenListModel")
	_info.InvokeClassMethod("set_model", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
