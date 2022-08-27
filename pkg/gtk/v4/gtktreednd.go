// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// GdkContentProvider* _gotk4_gtk4_TreeDragSource_virtual_drag_data_get(void* fnptr, GtkTreeDragSource* arg0, GtkTreePath* arg1) {
//   return ((GdkContentProvider* (*)(GtkTreeDragSource*, GtkTreePath*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gtk4_TreeDragDest_virtual_drag_data_received(void* fnptr, GtkTreeDragDest* arg0, GtkTreePath* arg1, GValue* arg2) {
//   return ((gboolean (*)(GtkTreeDragDest*, GtkTreePath*, GValue*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gtk4_TreeDragDest_virtual_row_drop_possible(void* fnptr, GtkTreeDragDest* arg0, GtkTreePath* arg1, GValue* arg2) {
//   return ((gboolean (*)(GtkTreeDragDest*, GtkTreePath*, GValue*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gtk4_TreeDragSource_virtual_drag_data_delete(void* fnptr, GtkTreeDragSource* arg0, GtkTreePath* arg1) {
//   return ((gboolean (*)(GtkTreeDragSource*, GtkTreePath*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gtk4_TreeDragSource_virtual_row_draggable(void* fnptr, GtkTreeDragSource* arg0, GtkTreePath* arg1) {
//   return ((gboolean (*)(GtkTreeDragSource*, GtkTreePath*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeTreeDragDest   = coreglib.Type(C.gtk_tree_drag_dest_get_type())
	GTypeTreeDragSource = coreglib.Type(C.gtk_tree_drag_source_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTreeDragDest, F: marshalTreeDragDest},
		coreglib.TypeMarshaler{T: GTypeTreeDragSource, F: marshalTreeDragSource},
	})
}

// TreeCreateRowDragContent creates a content provider for dragging path from
// tree_model.
//
// The function takes the following parameters:
//
//    - treeModel: TreeModel.
//    - path: row in tree_model.
//
// The function returns the following values:
//
//    - contentProvider: new ContentProvider.
//
func TreeCreateRowDragContent(treeModel TreeModeller, path *TreePath) *gdk.ContentProvider {
	var _arg1 *C.GtkTreeModel       // out
	var _arg2 *C.GtkTreePath        // out
	var _cret *C.GdkContentProvider // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(treeModel).Native()))
	_arg2 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_create_row_drag_content(_arg1, _arg2)
	runtime.KeepAlive(treeModel)
	runtime.KeepAlive(path)

	var _contentProvider *gdk.ContentProvider // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_contentProvider = &gdk.ContentProvider{
			Object: obj,
		}
	}

	return _contentProvider
}

// TreeGetRowDragData obtains a tree_model and path from value of target type
// GTK_TYPE_TREE_ROW_DATA.
//
// The returned path must be freed with gtk_tree_path_free().
//
// The function takes the following parameters:
//
//    - value: #GValue.
//
// The function returns the following values:
//
//    - treeModel (optional): TreeModel.
//    - path (optional): row in tree_model.
//    - ok: TRUE if selection_data had target type GTK_TYPE_TREE_ROW_DATA is
//      otherwise valid.
//
func TreeGetRowDragData(value *coreglib.Value) (*TreeModel, *TreePath, bool) {
	var _arg1 *C.GValue       // out
	var _arg2 *C.GtkTreeModel // in
	var _arg3 *C.GtkTreePath  // in
	var _cret C.gboolean      // in

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_tree_get_row_drag_data(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(value)

	var _treeModel *TreeModel // out
	var _path *TreePath       // out
	var _ok bool              // out

	if _arg2 != nil {
		_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_arg2)))
	}
	if _arg3 != nil {
		_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_arg3)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_path)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_tree_path_free((*C.GtkTreePath)(intern.C))
			},
		)
	}
	if _cret != 0 {
		_ok = true
	}

	return _treeModel, _path, _ok
}

// TreeDragDest: interface for Drag-and-Drop destinations in GtkTreeView.
//
// TreeDragDest wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TreeDragDest struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeDragDest)(nil)
)

// TreeDragDester describes TreeDragDest's interface methods.
type TreeDragDester interface {
	coreglib.Objector

	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from value.
	DragDataReceived(dest *TreePath, value *coreglib.Value) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path.
	RowDropPossible(destPath *TreePath, value *coreglib.Value) bool
}

var _ TreeDragDester = (*TreeDragDest)(nil)

func wrapTreeDragDest(obj *coreglib.Object) *TreeDragDest {
	return &TreeDragDest{
		Object: obj,
	}
}

func marshalTreeDragDest(p uintptr) (interface{}, error) {
	return wrapTreeDragDest(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataReceived asks the TreeDragDest to insert a row before the path dest,
// deriving the contents of the row from value. If dest is outside the tree so
// that inserting before it is impossible, FALSE will be returned. Also, FALSE
// may be returned if the new row is not created for some model-specific reason.
// Should robustly handle a dest no longer found in the model!.
//
// The function takes the following parameters:
//
//    - dest: row to drop in front of.
//    - value: data to drop.
//
// The function returns the following values:
//
//    - ok: whether a new row was created before position dest.
//
func (dragDest *TreeDragDest) DragDataReceived(dest *TreePath, value *coreglib.Value) bool {
	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(coreglib.InternObject(dragDest).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(dest)))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_tree_drag_dest_drag_data_received(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowDropPossible determines whether a drop is possible before the given
// dest_path, at the same depth as dest_path. i.e., can we drop the data in
// value at that location. dest_path does not have to exist; the return value
// will almost certainly be FALSE if the parent of dest_path doesn’t exist,
// though.
//
// The function takes the following parameters:
//
//    - destPath: destination row.
//    - value: data being dropped.
//
// The function returns the following values:
//
//    - ok: TRUE if a drop is possible before dest_path.
//
func (dragDest *TreeDragDest) RowDropPossible(destPath *TreePath, value *coreglib.Value) bool {
	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(coreglib.InternObject(dragDest).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(destPath)))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_tree_drag_dest_row_drop_possible(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(destPath)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// dragDataReceived asks the TreeDragDest to insert a row before the path dest,
// deriving the contents of the row from value. If dest is outside the tree so
// that inserting before it is impossible, FALSE will be returned. Also, FALSE
// may be returned if the new row is not created for some model-specific reason.
// Should robustly handle a dest no longer found in the model!.
//
// The function takes the following parameters:
//
//    - dest: row to drop in front of.
//    - value: data to drop.
//
// The function returns the following values:
//
//    - ok: whether a new row was created before position dest.
//
func (dragDest *TreeDragDest) dragDataReceived(dest *TreePath, value *coreglib.Value) bool {
	gclass := (*C.GtkTreeDragDestIface)(coreglib.PeekParentClass(dragDest))
	fnarg := gclass.drag_data_received

	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(coreglib.InternObject(dragDest).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(dest)))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C._gotk4_gtk4_TreeDragDest_virtual_drag_data_received(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// rowDropPossible determines whether a drop is possible before the given
// dest_path, at the same depth as dest_path. i.e., can we drop the data in
// value at that location. dest_path does not have to exist; the return value
// will almost certainly be FALSE if the parent of dest_path doesn’t exist,
// though.
//
// The function takes the following parameters:
//
//    - destPath: destination row.
//    - value: data being dropped.
//
// The function returns the following values:
//
//    - ok: TRUE if a drop is possible before dest_path.
//
func (dragDest *TreeDragDest) rowDropPossible(destPath *TreePath, value *coreglib.Value) bool {
	gclass := (*C.GtkTreeDragDestIface)(coreglib.PeekParentClass(dragDest))
	fnarg := gclass.row_drop_possible

	var _arg0 *C.GtkTreeDragDest // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GValue          // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(coreglib.InternObject(dragDest).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(destPath)))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C._gotk4_gtk4_TreeDragDest_virtual_row_drop_possible(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(destPath)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragSource: interface for Drag-and-Drop destinations in GtkTreeView.
//
// TreeDragSource wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TreeDragSource struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TreeDragSource)(nil)
)

// TreeDragSourcer describes TreeDragSource's interface methods.
type TreeDragSourcer interface {
	coreglib.Objector

	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop.
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to return a ContentProvider
	// representing the row at path.
	DragDataGet(path *TreePath) *gdk.ContentProvider
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation.
	RowDraggable(path *TreePath) bool
}

var _ TreeDragSourcer = (*TreeDragSource)(nil)

func wrapTreeDragSource(obj *coreglib.Object) *TreeDragSource {
	return &TreeDragSource{
		Object: obj,
	}
}

func marshalTreeDragSource(p uintptr) (interface{}, error) {
	return wrapTreeDragSource(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataDelete asks the TreeDragSource to delete the row at path, because it
// was moved somewhere else via drag-and-drop. Returns FALSE if the deletion
// fails because path no longer exists, or for some model-specific reason.
// Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was being dragged.
//
// The function returns the following values:
//
//    - ok: TRUE if the row was successfully deleted.
//
func (dragSource *TreeDragSource) DragDataDelete(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_drag_data_delete(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragDataGet asks the TreeDragSource to return a ContentProvider representing
// the row at path. Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was dragged.
//
// The function returns the following values:
//
//    - contentProvider (optional) for the given path or NULL if none exists.
//
func (dragSource *TreeDragSource) DragDataGet(path *TreePath) *gdk.ContentProvider {
	var _arg0 *C.GtkTreeDragSource  // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_drag_data_get(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _contentProvider *gdk.ContentProvider // out

	if _cret != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_contentProvider = &gdk.ContentProvider{
				Object: obj,
			}
		}
	}

	return _contentProvider
}

// RowDraggable asks the TreeDragSource whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement this
// interface, the row is assumed draggable.
//
// The function takes the following parameters:
//
//    - path: row on which user is initiating a drag.
//
// The function returns the following values:
//
//    - ok: TRUE if the row can be dragged.
//
func (dragSource *TreeDragSource) RowDraggable(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_row_draggable(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// dragDataDelete asks the TreeDragSource to delete the row at path, because it
// was moved somewhere else via drag-and-drop. Returns FALSE if the deletion
// fails because path no longer exists, or for some model-specific reason.
// Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was being dragged.
//
// The function returns the following values:
//
//    - ok: TRUE if the row was successfully deleted.
//
func (dragSource *TreeDragSource) dragDataDelete(path *TreePath) bool {
	gclass := (*C.GtkTreeDragSourceIface)(coreglib.PeekParentClass(dragSource))
	fnarg := gclass.drag_data_delete

	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C._gotk4_gtk4_TreeDragSource_virtual_drag_data_delete(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// dragDataGet asks the TreeDragSource to return a ContentProvider representing
// the row at path. Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was dragged.
//
// The function returns the following values:
//
//    - contentProvider (optional) for the given path or NULL if none exists.
//
func (dragSource *TreeDragSource) dragDataGet(path *TreePath) *gdk.ContentProvider {
	gclass := (*C.GtkTreeDragSourceIface)(coreglib.PeekParentClass(dragSource))
	fnarg := gclass.drag_data_get

	var _arg0 *C.GtkTreeDragSource  // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C._gotk4_gtk4_TreeDragSource_virtual_drag_data_get(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _contentProvider *gdk.ContentProvider // out

	if _cret != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_contentProvider = &gdk.ContentProvider{
				Object: obj,
			}
		}
	}

	return _contentProvider
}

// rowDraggable asks the TreeDragSource whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement this
// interface, the row is assumed draggable.
//
// The function takes the following parameters:
//
//    - path: row on which user is initiating a drag.
//
// The function returns the following values:
//
//    - ok: TRUE if the row can be dragged.
//
func (dragSource *TreeDragSource) rowDraggable(path *TreePath) bool {
	gclass := (*C.GtkTreeDragSourceIface)(coreglib.PeekParentClass(dragSource))
	fnarg := gclass.row_draggable

	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(coreglib.InternObject(dragSource).Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C._gotk4_gtk4_TreeDragSource_virtual_row_draggable(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragDestIface: instance of this type is always passed by reference.
type TreeDragDestIface struct {
	*treeDragDestIface
}

// treeDragDestIface is the struct that's finalized.
type treeDragDestIface struct {
	native *C.GtkTreeDragDestIface
}

// TreeDragSourceIface: instance of this type is always passed by reference.
type TreeDragSourceIface struct {
	*treeDragSourceIface
}

// treeDragSourceIface is the struct that's finalized.
type treeDragSourceIface struct {
	native *C.GtkTreeDragSourceIface
}
