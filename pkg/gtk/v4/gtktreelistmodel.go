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

// glib.Type values for gtktreelistmodel.go.
var (
	GTypeTreeListModel = coreglib.Type(C.gtk_tree_list_model_get_type())
	GTypeTreeListRow   = coreglib.Type(C.gtk_tree_list_row_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTreeListModel, F: marshalTreeListModel},
		{T: GTypeTreeListRow, F: marshalTreeListRow},
	})
}

// TreeListModelOverrider contains methods that are overridable.
type TreeListModelOverrider interface {
}

// TreeListModel: GtkTreeListModel is a list model that can create child models
// on demand.
type TreeListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*TreeListModel)(nil)
)

func classInitTreeListModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeListModel(obj *coreglib.Object) *TreeListModel {
	return &TreeListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalTreeListModel(p uintptr) (interface{}, error) {
	return wrapTreeListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListModel").InvokeMethod("get_autoexpand", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
func (self *TreeListModel) ChildRow(position uint32) *TreeListRow {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)
	*(**TreeListModel)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeListModel").InvokeMethod("get_child_row", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// Model gets the root model that self was created with.
//
// The function returns the following values:
//
//    - listModel: root model.
//
func (self *TreeListModel) Model() *gio.ListModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListModel").InvokeMethod("get_model", args[:], nil)
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
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListModel)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListModel").InvokeMethod("get_passthrough", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
func (self *TreeListModel) Row(position uint32) *TreeListRow {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)
	*(**TreeListModel)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeListModel").InvokeMethod("get_row", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if autoexpand {
		_arg1 = C.TRUE
	}
	*(**TreeListModel)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "TreeListModel").InvokeMethod("set_autoexpand", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(autoexpand)
}

// TreeListRowOverrider contains methods that are overridable.
type TreeListRowOverrider interface {
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
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeListRow)(nil)
)

func classInitTreeListRower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTreeListRow(obj *coreglib.Object) *TreeListRow {
	return &TreeListRow{
		Object: obj,
	}
}

func marshalTreeListRow(p uintptr) (interface{}, error) {
	return wrapTreeListRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
func (self *TreeListRow) ChildRow(position uint32) *TreeListRow {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)
	*(**TreeListRow)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("get_child_row", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
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
func (self *TreeListRow) Children() *gio.ListModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListRow)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("get_children", args[:], nil)
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
func (self *TreeListRow) Depth() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListRow)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("get_depth", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// Expanded gets if a row is currently expanded.
//
// The function returns the following values:
//
//    - ok: TRUE if the row is expanded.
//
func (self *TreeListRow) Expanded() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListRow)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("get_expanded", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListRow)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("get_parent", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
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
func (self *TreeListRow) Position() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListRow)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("get_position", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _guint uint32 // out

	_guint = uint32(_cret)

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
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**TreeListRow)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("is_expandable", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expanded {
		_arg1 = C.TRUE
	}
	*(**TreeListRow)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "TreeListRow").InvokeMethod("set_expanded", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(expanded)
}
