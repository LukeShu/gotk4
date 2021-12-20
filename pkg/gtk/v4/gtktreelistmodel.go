// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// GListModel* _gotk4_gtk4_TreeListModelCreateModelFunc(gpointer, gpointer);
// extern void callbackDelete(gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_list_model_get_type()), F: marshalTreeListModeller},
		{T: externglib.Type(C.gtk_tree_list_row_get_type()), F: marshalTreeListRower},
	})
}

// TreeListModelCreateModelFunc: prototype of the function called to create new
// child models when gtk_tree_list_row_set_expanded() is called.
//
// This function can return NULL to indicate that item is guaranteed to be a
// leaf node and will never have children. If it does not have children but may
// get children later, it should return an empty model that is filled once
// children arrive.
type TreeListModelCreateModelFunc func(item *externglib.Object) (listModel gio.ListModeller)

//export _gotk4_gtk4_TreeListModelCreateModelFunc
func _gotk4_gtk4_TreeListModelCreateModelFunc(arg0 C.gpointer, arg1 C.gpointer) (cret *C.GListModel) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out

	item = externglib.Take(unsafe.Pointer(arg0))

	fn := v.(TreeListModelCreateModelFunc)
	listModel := fn(item)

	if listModel != nil {
		cret = (*C.GListModel)(unsafe.Pointer(listModel.Native()))
		C.g_object_ref(C.gpointer(listModel.Native()))
	}

	return cret
}

// TreeListModel: GtkTreeListModel is a list model that can create child models
// on demand.
type TreeListModel struct {
	_ [0]func() // equal guard
	*externglib.Object

	gio.ListModel
}

var (
	_ externglib.Objector = (*TreeListModel)(nil)
)

func wrapTreeListModel(obj *externglib.Object) *TreeListModel {
	return &TreeListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalTreeListModeller(p uintptr) (interface{}, error) {
	return wrapTreeListModel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTreeListModel creates a new empty GtkTreeListModel displaying root with
// all rows collapsed.
//
// The function takes the following parameters:
//
//    - root: GListModel to use as root.
//    - passthrough: TRUE to pass through items from the models.
//    - autoexpand: TRUE to set the autoexpand property and expand the root
//      model.
//    - createFunc: function to call to create the GListModel for the children of
//      an item.
//
// The function returns the following values:
//
//    - treeListModel: newly created GtkTreeListModel.
//
func NewTreeListModel(root gio.ListModeller, passthrough, autoexpand bool, createFunc TreeListModelCreateModelFunc) *TreeListModel {
	var _arg1 *C.GListModel                     // out
	var _arg2 C.gboolean                        // out
	var _arg3 C.gboolean                        // out
	var _arg4 C.GtkTreeListModelCreateModelFunc // out
	var _arg5 C.gpointer
	var _arg6 C.GDestroyNotify
	var _cret *C.GtkTreeListModel // in

	_arg1 = (*C.GListModel)(unsafe.Pointer(root.Native()))
	C.g_object_ref(C.gpointer(root.Native()))
	if passthrough {
		_arg2 = C.TRUE
	}
	if autoexpand {
		_arg3 = C.TRUE
	}
	_arg4 = (*[0]byte)(C._gotk4_gtk4_TreeListModelCreateModelFunc)
	_arg5 = C.gpointer(gbox.Assign(createFunc))
	_arg6 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_cret = C.gtk_tree_list_model_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(root)
	runtime.KeepAlive(passthrough)
	runtime.KeepAlive(autoexpand)
	runtime.KeepAlive(createFunc)

	var _treeListModel *TreeListModel // out

	_treeListModel = wrapTreeListModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _treeListModel
}

// Autoexpand gets whether the model is set to automatically expand new rows
// that get added.
//
// This can be either rows added by changes to the underlying models or via
// gtk.TreeListRow.SetExpanded().
//
// The function returns the following values:
//
//    - ok: TRUE if the model is set to autoexpand.
//
func (self *TreeListModel) Autoexpand() bool {
	var _arg0 *C.GtkTreeListModel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_model_get_autoexpand(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildRow gets the row item corresponding to the child at index position for
// self's root model.
//
// If position is greater than the number of children in the root model, NULL is
// returned.
//
// Do not confuse this function with gtk.TreeListModel.GetRow().
//
// The function takes the following parameters:
//
//    - position of the child to get.
//
// The function returns the following values:
//
//    - treeListRow (optional): child in position.
//
func (self *TreeListModel) ChildRow(position uint) *TreeListRow {
	var _arg0 *C.GtkTreeListModel // out
	var _arg1 C.guint             // out
	var _cret *C.GtkTreeListRow   // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_tree_list_model_get_child_row(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// Model gets the root model that self was created with.
//
// The function returns the following values:
//
//    - listModel: root model.
//
func (self *TreeListModel) Model() gio.ListModeller {
	var _arg0 *C.GtkTreeListModel // out
	var _cret *C.GListModel       // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.ListModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(gio.ListModeller)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// Passthrough gets whether the model is passing through original row items.
//
// If this function returns FALSE, the GListModel functions for self return
// custom GtkTreeListRow objects. You need to call gtk.TreeListRow.GetItem() on
// these objects to get the original item.
//
// If TRUE, the values of the child models are passed through in their original
// state. You then need to call gtk.TreeListModel.GetRow() to get the custom
// GtkTreeListRows.
//
// The function returns the following values:
//
//    - ok: TRUE if the model is passing through original row items.
//
func (self *TreeListModel) Passthrough() bool {
	var _arg0 *C.GtkTreeListModel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_model_get_passthrough(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Row gets the row object for the given row.
//
// If position is greater than the number of items in self, NULL is returned.
//
// The row object can be used to expand and collapse rows as well as to inspect
// its position in the tree. See its documentation for details.
//
// This row object is persistent and will refer to the current item as long as
// the row is present in self, independent of other rows being added or removed.
//
// If self is set to not be passthrough, this function is equivalent to calling
// g_list_model_get_item().
//
// Do not confuse this function with gtk.TreeListModel.GetChildRow().
//
// The function takes the following parameters:
//
//    - position of the row to fetch.
//
// The function returns the following values:
//
//    - treeListRow (optional): row item.
//
func (self *TreeListModel) Row(position uint) *TreeListRow {
	var _arg0 *C.GtkTreeListModel // out
	var _arg1 C.guint             // out
	var _cret *C.GtkTreeListRow   // in

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_tree_list_model_get_row(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// SetAutoexpand sets whether the model should autoexpand.
//
// If set to TRUE, the model will recursively expand all rows that get added to
// the model. This can be either rows added by changes to the underlying models
// or via gtk.TreeListRow.SetExpanded().
//
// The function takes the following parameters:
//
//    - autoexpand: TRUE to make the model autoexpand its rows.
//
func (self *TreeListModel) SetAutoexpand(autoexpand bool) {
	var _arg0 *C.GtkTreeListModel // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkTreeListModel)(unsafe.Pointer(self.Native()))
	if autoexpand {
		_arg1 = C.TRUE
	}

	C.gtk_tree_list_model_set_autoexpand(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(autoexpand)
}

// TreeListRow: GtkTreeListRow is used by GtkTreeListModel to represent items.
//
// It allows navigating the model as a tree and modify the state of rows.
//
// GtkTreeListRow instances are created by a GtkTreeListModel only when the
// gtk.TreeListModel:passthrough property is not set.
//
// There are various support objects that can make use of GtkTreeListRow
// objects, such as the gtk.TreeExpander widget that allows displaying an icon
// to expand or collapse a row or gtk.TreeListRowSorter that makes it possible
// to sort trees properly.
type TreeListRow struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*TreeListRow)(nil)
)

func wrapTreeListRow(obj *externglib.Object) *TreeListRow {
	return &TreeListRow{
		Object: obj,
	}
}

func marshalTreeListRower(p uintptr) (interface{}, error) {
	return wrapTreeListRow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ChildRow: if self is not expanded or position is greater than the number of
// children, NULL is returned.
//
// The function takes the following parameters:
//
//    - position of the child to get.
//
// The function returns the following values:
//
//    - treeListRow (optional): child in position.
//
func (self *TreeListRow) ChildRow(position uint) *TreeListRow {
	var _arg0 *C.GtkTreeListRow // out
	var _arg1 C.guint           // out
	var _cret *C.GtkTreeListRow // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_tree_list_row_get_child_row(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// Children: if the row is expanded, gets the model holding the children of
// self.
//
// This model is the model created by the gtk.TreeListModelCreateModelFunc and
// contains the original items, no matter what value
// gtk.TreeListModel:passthrough is set to.
//
// The function returns the following values:
//
//    - listModel (optional): model containing the children.
//
func (self *TreeListRow) Children() gio.ListModeller {
	var _arg0 *C.GtkTreeListRow // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_get_children(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(gio.ListModeller)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.ListModeller")
			}
			_listModel = rv
		}
	}

	return _listModel
}

// Depth gets the depth of this row.
//
// Rows that correspond to items in the root model have a depth of zero, rows
// corresponding to items of models of direct children of the root model have a
// depth of 1 and so on.
//
// The depth of a row never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - guint: depth of this row.
//
func (self *TreeListRow) Depth() uint {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_get_depth(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Expanded gets if a row is currently expanded.
//
// The function returns the following values:
//
//    - ok: TRUE if the row is expanded.
//
func (self *TreeListRow) Expanded() bool {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_get_expanded(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Item gets the item corresponding to this row,
//
// The value returned by this function never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - object (optional): item of this row or NULL when the row was destroyed.
//
func (self *TreeListRow) Item() *externglib.Object {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.gpointer        // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_get_item(_arg0)
	runtime.KeepAlive(self)

	var _object *externglib.Object // out

	_object = externglib.AssumeOwnership(unsafe.Pointer(_cret))

	return _object
}

// Parent gets the row representing the parent for self.
//
// That is the row that would need to be collapsed to make this row disappear.
//
// If self is a row corresponding to the root model, NULL is returned.
//
// The value returned by this function never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - treeListRow (optional): parent of self.
//
func (self *TreeListRow) Parent() *TreeListRow {
	var _arg0 *C.GtkTreeListRow // out
	var _cret *C.GtkTreeListRow // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_get_parent(_arg0)
	runtime.KeepAlive(self)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// Position returns the position in the GtkTreeListModel that self occupies at
// the moment.
//
// The function returns the following values:
//
//    - guint: position in the model.
//
func (self *TreeListRow) Position() uint {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.guint           // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_get_position(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IsExpandable checks if a row can be expanded.
//
// This does not mean that the row is actually expanded, this can be checked
// with gtk.TreeListRow.GetExpanded().
//
// If a row is expandable never changes until the row is destroyed.
//
// The function returns the following values:
//
//    - ok: TRUE if the row is expandable.
//
func (self *TreeListRow) IsExpandable() bool {
	var _arg0 *C.GtkTreeListRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_list_row_is_expandable(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpanded expands or collapses a row.
//
// If a row is expanded, the model of calling the
// gtk.TreeListModelCreateModelFunc for the row's item will be inserted after
// this row. If a row is collapsed, those items will be removed from the model.
//
// If the row is not expandable, this function does nothing.
//
// The function takes the following parameters:
//
//    - expanded: TRUE if the row should be expanded.
//
func (self *TreeListRow) SetExpanded(expanded bool) {
	var _arg0 *C.GtkTreeListRow // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkTreeListRow)(unsafe.Pointer(self.Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.gtk_tree_list_row_set_expanded(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expanded)
}
