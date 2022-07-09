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

// GTypeMultiSelection returns the GType for the type MultiSelection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMultiSelection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MultiSelection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMultiSelection)
	return gtype
}

// MultiSelectionOverrider contains methods that are overridable.
type MultiSelectionOverrider interface {
}

// MultiSelection: GtkMultiSelection is a GtkSelectionModel that allows
// selecting multiple elements.
type MultiSelection struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SelectionModel
}

var (
	_ coreglib.Objector = (*MultiSelection)(nil)
)

func classInitMultiSelectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMultiSelection(obj *coreglib.Object) *MultiSelection {
	return &MultiSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalMultiSelection(p uintptr) (interface{}, error) {
	return wrapMultiSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMultiSelection creates a new selection to handle model.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to manage, or NULL.
//
// The function returns the following values:
//
//    - multiSelection: new GtkMultiSelection.
//
func NewMultiSelection(model gio.ListModeller) *MultiSelection {
	var _args [1]girepository.Argument

	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}

	_gret := girepository.MustFind("Gtk", "MultiSelection").InvokeMethod("new_MultiSelection", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _multiSelection *MultiSelection // out

	_multiSelection = wrapMultiSelection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _multiSelection
}

// Model returns the underlying model of self.
//
// The function returns the following values:
//
//    - listModel: underlying model.
//
func (self *MultiSelection) Model() *gio.ListModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MultiSelection").InvokeMethod("get_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, self will be empty.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to wrap.
//
func (self *MultiSelection) SetModel(model gio.ListModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	girepository.MustFind("Gtk", "MultiSelection").InvokeMethod("set_model", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
