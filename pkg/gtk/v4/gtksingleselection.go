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

// GTypeSingleSelection returns the GType for the type SingleSelection.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSingleSelection() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "SingleSelection").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSingleSelection)
	return gtype
}

// SingleSelectionOverrider contains methods that are overridable.
type SingleSelectionOverrider interface {
}

// SingleSelection: GtkSingleSelection is a GtkSelectionModel that allows
// selecting a single item.
//
// Note that the selection is *persistent* -- if the selected item is removed
// and re-added in the same ::items-changed emission, it stays selected. In
// particular, this means that changing the sort order of an underlying sort
// model will preserve the selection.
type SingleSelection struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SelectionModel
}

var (
	_ coreglib.Objector = (*SingleSelection)(nil)
)

func classInitSingleSelectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSingleSelection(obj *coreglib.Object) *SingleSelection {
	return &SingleSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalSingleSelection(p uintptr) (interface{}, error) {
	return wrapSingleSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSingleSelection creates a new selection to handle model.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to manage, or NULL.
//
// The function returns the following values:
//
//    - singleSelection: new GtkSingleSelection.
//
func NewSingleSelection(model gio.ListModeller) *SingleSelection {
	var _args [1]girepository.Argument

	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_gret := _info.InvokeClassMethod("new_SingleSelection", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _singleSelection *SingleSelection // out

	_singleSelection = wrapSingleSelection(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _singleSelection
}

// Autoselect checks if autoselect has been enabled or disabled via
// gtk_single_selection_set_autoselect().
//
// The function returns the following values:
//
//    - ok: TRUE if autoselect is enabled.
//
func (self *SingleSelection) Autoselect() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_gret := _info.InvokeClassMethod("get_autoselect", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// CanUnselect: if TRUE, gtk_selection_model_unselect_item() is supported and
// allows unselecting the selected item.
//
// The function returns the following values:
//
//    - ok: TRUE to support unselecting.
//
func (self *SingleSelection) CanUnselect() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_gret := _info.InvokeClassMethod("get_can_unselect", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Model gets the model that self is wrapping.
//
// The function returns the following values:
//
//    - listModel: model being wrapped.
//
func (self *SingleSelection) Model() *gio.ListModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_gret := _info.InvokeClassMethod("get_model", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	{
		obj := coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
		_listModel = &gio.ListModel{
			Object: obj,
		}
	}

	return _listModel
}

// Selected gets the position of the selected item.
//
// If no item is selected, GTK_INVALID_LIST_POSITION is returned.
//
// The function returns the following values:
//
//    - guint: position of the selected item.
//
func (self *SingleSelection) Selected() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_gret := _info.InvokeClassMethod("get_selected", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// SelectedItem gets the selected item.
//
// If no item is selected, NULL is returned.
//
// The function returns the following values:
//
//    - object (optional): selected item.
//
func (self *SingleSelection) SelectedItem() *coreglib.Object {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_gret := _info.InvokeClassMethod("get_selected_item", _args[:], nil)
	_cret := *(*C.gpointer)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _object *coreglib.Object // out

	_object = coreglib.Take(unsafe.Pointer(*(*C.gpointer)(unsafe.Pointer(&_cret))))

	return _object
}

// SetAutoselect enables or disables autoselect.
//
// If autoselect is TRUE, self will enforce that an item is always selected. It
// will select a new item when the currently selected item is deleted and it
// will disallow unselecting the current item.
//
// The function takes the following parameters:
//
//    - autoselect: TRUE to always select an item.
//
func (self *SingleSelection) SetAutoselect(autoselect bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if autoselect {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_info.InvokeClassMethod("set_autoselect", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(autoselect)
}

// SetCanUnselect: if TRUE, unselecting the current item via
// gtk_selection_model_unselect_item() is supported.
//
// Note that setting gtk.SingleSelection:autoselect will cause unselecting to
// not work, so it practically makes no sense to set both at the same time the
// same time.
//
// The function takes the following parameters:
//
//    - canUnselect: TRUE to allow unselecting.
//
func (self *SingleSelection) SetCanUnselect(canUnselect bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if canUnselect {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_info.InvokeClassMethod("set_can_unselect", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(canUnselect)
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, self will be empty.
//
// The function takes the following parameters:
//
//    - model (optional): GListModel to wrap.
//
func (self *SingleSelection) SetModel(model gio.ListModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_info.InvokeClassMethod("set_model", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetSelected selects the item at the given position.
//
// If the list does not have an item at position or GTK_INVALID_LIST_POSITION is
// given, the behavior depends on the value of the
// gtk.SingleSelection:autoselect property: If it is set, no change will occur
// and the old item will stay selected. If it is unset, the selection will be
// unset and no item will be selected.
//
// The function takes the following parameters:
//
//    - position: item to select or GTK_INVALID_LIST_POSITION.
//
func (self *SingleSelection) SetSelected(position uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(position)

	_info := girepository.MustFind("Gtk", "SingleSelection")
	_info.InvokeClassMethod("set_selected", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}
