// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_slice_list_model_get_type()), F: marshalSliceListModel},
	})
}

// SliceListModel: `GtkSliceListModel` is a list model that presents a slice of
// another model.
//
// This is useful when implementing paging by setting the size to the number of
// elements per page and updating the offset whenever a different page is
// opened.
type SliceListModel interface {
	gextras.Objector

	// Offset gets the offset set via gtk_slice_list_model_set_offset().
	Offset() uint
	// Size gets the size set via gtk_slice_list_model_set_size().
	Size() uint
	// SetOffset sets the offset into the original model for this slice.
	//
	// If the offset is too large for the sliced model, @self will end up empty.
	SetOffset(offset uint)
	// SetSize sets the maximum size. @self will never have more items than
	// @size.
	//
	// It can however have fewer items if the offset is too large or the model
	// sliced from doesn't have enough items.
	SetSize(size uint)
}

// SliceListModelClass implements the SliceListModel interface.
type SliceListModelClass struct {
	*externglib.Object
}

var _ SliceListModel = (*SliceListModelClass)(nil)

func wrapSliceListModel(obj *externglib.Object) SliceListModel {
	return &SliceListModelClass{
		Object: obj,
	}
}

func marshalSliceListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSliceListModel(obj), nil
}

// Offset gets the offset set via gtk_slice_list_model_set_offset().
func (s *SliceListModelClass) Offset() uint {
	var _arg0 *C.GtkSliceListModel // out
	var _cret C.guint              // in

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_slice_list_model_get_offset(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Size gets the size set via gtk_slice_list_model_set_size().
func (s *SliceListModelClass) Size() uint {
	var _arg0 *C.GtkSliceListModel // out
	var _cret C.guint              // in

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_slice_list_model_get_size(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetOffset sets the offset into the original model for this slice.
//
// If the offset is too large for the sliced model, @self will end up empty.
func (s *SliceListModelClass) SetOffset(offset uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer((&s).Native()))
	_arg1 = C.guint(offset)

	C.gtk_slice_list_model_set_offset(_arg0, _arg1)
}

// SetSize sets the maximum size. @self will never have more items than @size.
//
// It can however have fewer items if the offset is too large or the model
// sliced from doesn't have enough items.
func (s *SliceListModelClass) SetSize(size uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer((&s).Native()))
	_arg1 = C.guint(size)

	C.gtk_slice_list_model_set_size(_arg0, _arg1)
}
