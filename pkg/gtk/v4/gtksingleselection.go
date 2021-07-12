// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_single_selection_get_type()), F: marshalSingleSelectioner},
	})
}

// SingleSelectioner describes SingleSelection's methods.
type SingleSelectioner interface {
	// Autoselect checks if autoselect has been enabled or disabled via
	// gtk_single_selection_set_autoselect().
	Autoselect() bool
	// CanUnselect: if true, gtk_selection_model_unselect_item() is supported
	// and allows unselecting the selected item.
	CanUnselect() bool
	// Model gets the model that @self is wrapping.
	Model() *gio.ListModel
	// Selected gets the position of the selected item.
	Selected() uint
	// SelectedItem gets the selected item.
	SelectedItem() *externglib.Object
	// SetAutoselect enables or disables autoselect.
	SetAutoselect(autoselect bool)
	// SetCanUnselect: if true, unselecting the current item via
	// gtk_selection_model_unselect_item() is supported.
	SetCanUnselect(canUnselect bool)
	// SetModel sets the model that @self should wrap.
	SetModel(model gio.ListModeler)
	// SetSelected selects the item at the given position.
	SetSelected(position uint)
}

// SingleSelection: `GtkSingleSelection` is a `GtkSelectionModel` that allows
// selecting a single item.
//
// Note that the selection is *persistent* -- if the selected item is removed
// and re-added in the same ::items-changed emission, it stays selected. In
// particular, this means that changing the sort order of an underlying sort
// model will preserve the selection.
type SingleSelection struct {
	*externglib.Object

	SelectionModel
}

var (
	_ SingleSelectioner = (*SingleSelection)(nil)
	_ gextras.Nativer   = (*SingleSelection)(nil)
)

func wrapSingleSelection(obj *externglib.Object) SingleSelectioner {
	return &SingleSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalSingleSelectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSingleSelection(obj), nil
}

// NewSingleSelection creates a new selection to handle @model.
func NewSingleSelection(model gio.ListModeler) *SingleSelection {
	var _arg1 *C.GListModel         // out
	var _cret *C.GtkSingleSelection // in

	_arg1 = (*C.GListModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	_cret = C.gtk_single_selection_new(_arg1)

	var _singleSelection *SingleSelection // out

	_singleSelection = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*SingleSelection)

	return _singleSelection
}

// Autoselect checks if autoselect has been enabled or disabled via
// gtk_single_selection_set_autoselect().
func (self *SingleSelection) Autoselect() bool {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_autoselect(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUnselect: if true, gtk_selection_model_unselect_item() is supported and
// allows unselecting the selected item.
func (self *SingleSelection) CanUnselect() bool {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_can_unselect(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model gets the model that @self is wrapping.
func (self *SingleSelection) Model() *gio.ListModel {
	var _arg0 *C.GtkSingleSelection // out
	var _cret *C.GListModel         // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_model(_arg0)

	var _listModel *gio.ListModel // out

	_listModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.ListModel)

	return _listModel
}

// Selected gets the position of the selected item.
//
// If no item is selected, GTK_INVALID_LIST_POSITION is returned.
func (self *SingleSelection) Selected() uint {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.guint               // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_selected(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SelectedItem gets the selected item.
//
// If no item is selected, nil is returned.
func (self *SingleSelection) SelectedItem() *externglib.Object {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.gpointer            // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_selected_item(_arg0)

	var _object *externglib.Object // out

	_object = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*externglib.Object)

	return _object
}

// SetAutoselect enables or disables autoselect.
//
// If @autoselect is true, @self will enforce that an item is always selected.
// It will select a new item when the currently selected item is deleted and it
// will disallow unselecting the current item.
func (self *SingleSelection) SetAutoselect(autoselect bool) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	if autoselect {
		_arg1 = C.TRUE
	}

	C.gtk_single_selection_set_autoselect(_arg0, _arg1)
}

// SetCanUnselect: if true, unselecting the current item via
// gtk_selection_model_unselect_item() is supported.
//
// Note that setting [property@Gtk.SingleSelection:autoselect] will cause
// unselecting to not work, so it practically makes no sense to set both at the
// same time the same time.
func (self *SingleSelection) SetCanUnselect(canUnselect bool) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	if canUnselect {
		_arg1 = C.TRUE
	}

	C.gtk_single_selection_set_can_unselect(_arg0, _arg1)
}

// SetModel sets the model that @self should wrap.
//
// If @model is nil, @self will be empty.
func (self *SingleSelection) SetModel(model gio.ListModeler) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 *C.GListModel         // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_single_selection_set_model(_arg0, _arg1)
}

// SetSelected selects the item at the given position.
//
// If the list does not have an item at @position or GTK_INVALID_LIST_POSITION
// is given, the behavior depends on the value of the
// [property@Gtk.SingleSelection:autoselect] property: If it is set, no change
// will occur and the old item will stay selected. If it is unset, the selection
// will be unset and no item will be selected.
func (self *SingleSelection) SetSelected(position uint) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 C.guint               // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	C.gtk_single_selection_set_selected(_arg0, _arg1)
}
