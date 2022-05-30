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

// glib.Type values for gtkslicelistmodel.go.
var GTypeSliceListModel = coreglib.Type(C.gtk_slice_list_model_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSliceListModel, F: marshalSliceListModel},
	})
}

// SliceListModelOverrider contains methods that are overridable.
type SliceListModelOverrider interface {
}

// SliceListModel: GtkSliceListModel is a list model that presents a slice of
// another model.
//
// This is useful when implementing paging by setting the size to the number of
// elements per page and updating the offset whenever a different page is
// opened.
type SliceListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*SliceListModel)(nil)
)

func classInitSliceListModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSliceListModel(obj *coreglib.Object) *SliceListModel {
	return &SliceListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSliceListModel(p uintptr) (interface{}, error) {
	return wrapSliceListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSliceListModel creates a new slice model.
//
// It presents the slice from offset to offset + size of the given model.
//
// The function takes the following parameters:
//
//    - model (optional) to use, or NULL.
//    - offset of the slice.
//    - size: maximum size of the slice.
//
// The function returns the following values:
//
//    - sliceListModel: new GtkSliceListModel.
//
func NewSliceListModel(model gio.ListModeller, offset, size uint) *SliceListModel {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _arg2 C.guint // out
	var _cret *C.void // in

	if model != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(model).Native()))
	}
	_arg1 = C.guint(offset)
	_arg2 = C.guint(size)
	*(*gio.ListModeller)(unsafe.Pointer(&args[0])) = _arg0
	*(*uint)(unsafe.Pointer(&args[1])) = _arg1
	*(*uint)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "SliceListModel").InvokeMethod("new_SliceListModel", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(size)

	var _sliceListModel *SliceListModel // out

	_sliceListModel = wrapSliceListModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sliceListModel
}

// Model gets the model that is currently being used or NULL if none.
//
// The function returns the following values:
//
//    - listModel (optional): model in use.
//
func (self *SliceListModel) Model() *gio.ListModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**SliceListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SliceListModel").InvokeMethod("get_model", args[:], nil)
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

// Offset gets the offset set via gtk_slice_list_model_set_offset().
//
// The function returns the following values:
//
//    - guint: offset.
//
func (self *SliceListModel) Offset() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**SliceListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SliceListModel").InvokeMethod("get_offset", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Size gets the size set via gtk_slice_list_model_set_size().
//
// The function returns the following values:
//
//    - guint: size.
//
func (self *SliceListModel) Size() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**SliceListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SliceListModel").InvokeMethod("get_size", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetModel sets the model to show a slice of.
//
// The model's item type must conform to self's item type.
//
// The function takes the following parameters:
//
//    - model (optional) to be sliced.
//
func (self *SliceListModel) SetModel(model gio.ListModeller) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	*(**SliceListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SliceListModel").InvokeMethod("set_model", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetOffset sets the offset into the original model for this slice.
//
// If the offset is too large for the sliced model, self will end up empty.
//
// The function takes the following parameters:
//
//    - offset: new offset to use.
//
func (self *SliceListModel) SetOffset(offset uint) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(offset)
	*(**SliceListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SliceListModel").InvokeMethod("set_offset", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(offset)
}

// SetSize sets the maximum size. self will never have more items than size.
//
// It can however have fewer items if the offset is too large or the model
// sliced from doesn't have enough items.
//
// The function takes the following parameters:
//
//    - size: maximum size.
//
func (self *SliceListModel) SetSize(size uint) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(size)
	*(**SliceListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SliceListModel").InvokeMethod("set_size", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(size)
}
