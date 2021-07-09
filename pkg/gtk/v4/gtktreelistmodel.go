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
		{T: externglib.Type(C.gtk_tree_list_model_get_type()), F: marshalTreeListModel},
		{T: externglib.Type(C.gtk_tree_list_row_get_type()), F: marshalTreeListRow},
	})
}

// TreeListModel: `GtkTreeListModel` is a list model that can create child
// models on demand.
type TreeListModel interface {
	gextras.Objector

	// Autoexpand gets whether the model is set to automatically expand new rows
	// that get added.
	//
	// This can be either rows added by changes to the underlying models or via
	// [method@Gtk.TreeListRow.set_expanded].
	Autoexpand() bool
	// ChildRow gets the row item corresponding to the child at index @position
	// for @self's root model.
	//
	// If @position is greater than the number of children in the root model,
	// nil is returned.
	//
	// Do not confuse this function with [method@Gtk.TreeListModel.get_row].
	ChildRow(position uint) *TreeListRowClass
	// Passthrough gets whether the model is passing through original row items.
	//
	// If this function returns false, the `GListModel` functions for @self
	// return custom `GtkTreeListRow` objects. You need to call
	// [method@Gtk.TreeListRow.get_item] on these objects to get the original
	// item.
	//
	// If true, the values of the child models are passed through in their
	// original state. You then need to call [method@Gtk.TreeListModel.get_row]
	// to get the custom `GtkTreeListRow`s.
	Passthrough() bool
	// Row gets the row object for the given row.
	//
	// If @position is greater than the number of items in @self, nil is
	// returned.
	//
	// The row object can be used to expand and collapse rows as well as to
	// inspect its position in the tree. See its documentation for details.
	//
	// This row object is persistent and will refer to the current item as long
	// as the row is present in @self, independent of other rows being added or
	// removed.
	//
	// If @self is set to not be passthrough, this function is equivalent to
	// calling g_list_model_get_item().
	//
	// Do not confuse this function with
	// [method@Gtk.TreeListModel.get_child_row].
	Row(position uint) *TreeListRowClass
	// SetAutoexpand sets whether the model should autoexpand.
	//
	// If set to true, the model will recursively expand all rows that get added
	// to the model. This can be either rows added by changes to the underlying
	// models or via [method@Gtk.TreeListRow.set_expanded].
	SetAutoexpand(autoexpand bool)
}

// TreeListModelClass implements the TreeListModel interface.
type TreeListModelClass struct {
	*externglib.Object
}

var _ TreeListModel = (*TreeListModelClass)(nil)

func wrapTreeListModel(obj *externglib.Object) TreeListModel {
	return &TreeListModelClass{
		Object: obj,
	}
}

func marshalTreeListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeListModel(obj), nil
}

// Autoexpand gets whether the model is set to automatically expand new rows
// that get added.
//
// This can be either rows added by changes to the underlying models or via
// [method@Gtk.TreeListRow.set_expanded].
func (s *TreeListModelClass) Autoexpand() bool {
	var _arg0 *C.GtkTreeListModel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_model_get_autoexpand(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildRow gets the row item corresponding to the child at index @position for
// @self's root model.
//
// If @position is greater than the number of children in the root model, nil is
// returned.
//
// Do not confuse this function with [method@Gtk.TreeListModel.get_row].
func (s *TreeListModelClass) ChildRow(position uint) *TreeListRowClass {
	var _arg0 *C.GtkTreeListModel // out
	var _arg1 C.guint             // out
	var _cret *C.GtkTreeListRow   // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer((&s).Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_tree_list_model_get_child_row(_arg0, _arg1)

	var _treeListRow *TreeListRowClass // out

	_treeListRow = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TreeListRowClass)

	return _treeListRow
}

// Passthrough gets whether the model is passing through original row items.
//
// If this function returns false, the `GListModel` functions for @self return
// custom `GtkTreeListRow` objects. You need to call
// [method@Gtk.TreeListRow.get_item] on these objects to get the original item.
//
// If true, the values of the child models are passed through in their original
// state. You then need to call [method@Gtk.TreeListModel.get_row] to get the
// custom `GtkTreeListRow`s.
func (s *TreeListModelClass) Passthrough() bool {
	var _arg0 *C.GtkTreeListModel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_model_get_passthrough(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Row gets the row object for the given row.
//
// If @position is greater than the number of items in @self, nil is returned.
//
// The row object can be used to expand and collapse rows as well as to inspect
// its position in the tree. See its documentation for details.
//
// This row object is persistent and will refer to the current item as long as
// the row is present in @self, independent of other rows being added or
// removed.
//
// If @self is set to not be passthrough, this function is equivalent to calling
// g_list_model_get_item().
//
// Do not confuse this function with [method@Gtk.TreeListModel.get_child_row].
func (s *TreeListModelClass) Row(position uint) *TreeListRowClass {
	var _arg0 *C.GtkTreeListModel // out
	var _arg1 C.guint             // out
	var _cret *C.GtkTreeListRow   // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer((&s).Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_tree_list_model_get_row(_arg0, _arg1)

	var _treeListRow *TreeListRowClass // out

	_treeListRow = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TreeListRowClass)

	return _treeListRow
}

// SetAutoexpand sets whether the model should autoexpand.
//
// If set to true, the model will recursively expand all rows that get added to
// the model. This can be either rows added by changes to the underlying models
// or via [method@Gtk.TreeListRow.set_expanded].
func (s *TreeListModelClass) SetAutoexpand(autoexpand bool) {
	var _arg0 *C.GtkTreeListModel // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer((&s).Native()))
	if autoexpand {
		_arg1 = C.TRUE
	}

	C.gtk_tree_list_model_set_autoexpand(_arg0, _arg1)
}

// TreeListRow: `GtkTreeListRow` is used by `GtkTreeListModel` to represent
// items.
//
// It allows navigating the model as a tree and modify the state of rows.
//
// `GtkTreeListRow` instances are created by a `GtkTreeListModel` only when the
// [property@Gtk.TreeListModel:passthrough] property is not set.
//
// There are various support objects that can make use of `GtkTreeListRow`
// objects, such as the [class@Gtk.TreeExpander] widget that allows displaying
// an icon to expand or collapse a row or [class@Gtk.TreeListRowSorter] that
// makes it possible to sort trees properly.
type TreeListRow interface {
	gextras.Objector

	// ChildRow: if @self is not expanded or @position is greater than the
	// number of children, nil is returned.
	ChildRow(position uint) *TreeListRowClass
	// Depth gets the depth of this row.
	//
	// Rows that correspond to items in the root model have a depth of zero,
	// rows corresponding to items of models of direct children of the root
	// model have a depth of 1 and so on.
	//
	// The depth of a row never changes until the row is destroyed.
	Depth() uint
	// Expanded gets if a row is currently expanded.
	Expanded() bool
	// Item gets the item corresponding to this row,
	//
	// The value returned by this function never changes until the row is
	// destroyed.
	Item() *externglib.Object
	// Parent gets the row representing the parent for @self.
	//
	// That is the row that would need to be collapsed to make this row
	// disappear.
	//
	// If @self is a row corresponding to the root model, nil is returned.
	//
	// The value returned by this function never changes until the row is
	// destroyed.
	Parent() *TreeListRowClass
	// Position returns the position in the `GtkTreeListModel` that @self
	// occupies at the moment.
	Position() uint
	// IsExpandable checks if a row can be expanded.
	//
	// This does not mean that the row is actually expanded, this can be checked
	// with [method@Gtk.TreeListRow.get_expanded].
	//
	// If a row is expandable never changes until the row is destroyed.
	IsExpandable() bool
	// SetExpanded expands or collapses a row.
	//
	// If a row is expanded, the model of calling the
	// [callback@Gtk.TreeListModelCreateModelFunc] for the row's item will be
	// inserted after this row. If a row is collapsed, those items will be
	// removed from the model.
	//
	// If the row is not expandable, this function does nothing.
	SetExpanded(expanded bool)
}

// TreeListRowClass implements the TreeListRow interface.
type TreeListRowClass struct {
	*externglib.Object
}

var _ TreeListRow = (*TreeListRowClass)(nil)

func wrapTreeListRow(obj *externglib.Object) TreeListRow {
	return &TreeListRowClass{
		Object: obj,
	}
}

func marshalTreeListRow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeListRow(obj), nil
}

// ChildRow: if @self is not expanded or @position is greater than the number of
// children, nil is returned.
func (s *TreeListRowClass) ChildRow(position uint) *TreeListRowClass {
	var _arg0 *C.GtkTreeListRow // out
	var _arg1 C.guint           // out
	var _cret *C.GtkTreeListRow // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_tree_list_row_get_child_row(_arg0, _arg1)

	var _treeListRow *TreeListRowClass // out

	_treeListRow = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TreeListRowClass)

	return _treeListRow
}

// Depth gets the depth of this row.
//
// Rows that correspond to items in the root model have a depth of zero, rows
// corresponding to items of models of direct children of the root model have a
// depth of 1 and so on.
//
// The depth of a row never changes until the row is destroyed.
func (s *TreeListRowClass) Depth() uint {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_row_get_depth(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Expanded gets if a row is currently expanded.
func (s *TreeListRowClass) Expanded() bool {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_row_get_expanded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Item gets the item corresponding to this row,
//
// The value returned by this function never changes until the row is destroyed.
func (s *TreeListRowClass) Item() *externglib.Object {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.gpointer        // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_row_get_item(_arg0)

	var _object *externglib.Object // out

	_object = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(&_cret))).(*externglib.Object)

	return _object
}

// Parent gets the row representing the parent for @self.
//
// That is the row that would need to be collapsed to make this row disappear.
//
// If @self is a row corresponding to the root model, nil is returned.
//
// The value returned by this function never changes until the row is destroyed.
func (s *TreeListRowClass) Parent() *TreeListRowClass {
	var _arg0 *C.GtkTreeListRow // out
	var _cret *C.GtkTreeListRow // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_row_get_parent(_arg0)

	var _treeListRow *TreeListRowClass // out

	_treeListRow = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TreeListRowClass)

	return _treeListRow
}

// Position returns the position in the `GtkTreeListModel` that @self occupies
// at the moment.
func (s *TreeListRowClass) Position() uint {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_row_get_position(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IsExpandable checks if a row can be expanded.
//
// This does not mean that the row is actually expanded, this can be checked
// with [method@Gtk.TreeListRow.get_expanded].
//
// If a row is expandable never changes until the row is destroyed.
func (s *TreeListRowClass) IsExpandable() bool {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_tree_list_row_is_expandable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpanded expands or collapses a row.
//
// If a row is expanded, the model of calling the
// [callback@Gtk.TreeListModelCreateModelFunc] for the row's item will be
// inserted after this row. If a row is collapsed, those items will be removed
// from the model.
//
// If the row is not expandable, this function does nothing.
func (s *TreeListRowClass) SetExpanded(expanded bool) {
	var _arg0 *C.GtkTreeListRow // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer((&s).Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.gtk_tree_list_row_set_expanded(_arg0, _arg1)
}
