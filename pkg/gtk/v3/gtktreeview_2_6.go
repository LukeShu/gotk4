// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// extern gboolean _gotk4_gtk3_TreeViewRowSeparatorFunc(GtkTreeModel*, GtkTreeIter*, gpointer);
import "C"

// FixedHeightMode returns whether fixed height mode is turned on for tree_view.
//
// The function returns the following values:
//
//    - ok: TRUE if tree_view is in fixed height mode.
//
func (treeView *TreeView) FixedHeightMode() bool {
	var _arg0 *C.GtkTreeView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTreeView)(unsafe.Pointer(coreglib.InternObject(treeView).Native()))

	_cret = C.gtk_tree_view_get_fixed_height_mode(_arg0)
	runtime.KeepAlive(treeView)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HoverExpand returns whether hover expansion mode is turned on for tree_view.
//
// The function returns the following values:
//
//    - ok: TRUE if tree_view is in hover expansion mode.
//
func (treeView *TreeView) HoverExpand() bool {
	var _arg0 *C.GtkTreeView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTreeView)(unsafe.Pointer(coreglib.InternObject(treeView).Native()))

	_cret = C.gtk_tree_view_get_hover_expand(_arg0)
	runtime.KeepAlive(treeView)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HoverSelection returns whether hover selection mode is turned on for
// tree_view.
//
// The function returns the following values:
//
//    - ok: TRUE if tree_view is in hover selection mode.
//
func (treeView *TreeView) HoverSelection() bool {
	var _arg0 *C.GtkTreeView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTreeView)(unsafe.Pointer(coreglib.InternObject(treeView).Native()))

	_cret = C.gtk_tree_view_get_hover_selection(_arg0)
	runtime.KeepAlive(treeView)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFixedHeightMode enables or disables the fixed height mode of tree_view.
// Fixed height mode speeds up TreeView by assuming that all rows have the same
// height. Only enable this option if all rows are the same height and all
// columns are of type GTK_TREE_VIEW_COLUMN_FIXED.
//
// The function takes the following parameters:
//
//    - enable: TRUE to enable fixed height mode.
//
func (treeView *TreeView) SetFixedHeightMode(enable bool) {
	var _arg0 *C.GtkTreeView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkTreeView)(unsafe.Pointer(coreglib.InternObject(treeView).Native()))
	if enable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_set_fixed_height_mode(_arg0, _arg1)
	runtime.KeepAlive(treeView)
	runtime.KeepAlive(enable)
}

// SetHoverExpand enables or disables the hover expansion mode of tree_view.
// Hover expansion makes rows expand or collapse if the pointer moves over them.
//
// The function takes the following parameters:
//
//    - expand: TRUE to enable hover selection mode.
//
func (treeView *TreeView) SetHoverExpand(expand bool) {
	var _arg0 *C.GtkTreeView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkTreeView)(unsafe.Pointer(coreglib.InternObject(treeView).Native()))
	if expand {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_set_hover_expand(_arg0, _arg1)
	runtime.KeepAlive(treeView)
	runtime.KeepAlive(expand)
}

// SetHoverSelection enables or disables the hover selection mode of tree_view.
// Hover selection makes the selected row follow the pointer. Currently, this
// works only for the selection modes GTK_SELECTION_SINGLE and
// GTK_SELECTION_BROWSE.
//
// The function takes the following parameters:
//
//    - hover: TRUE to enable hover selection mode.
//
func (treeView *TreeView) SetHoverSelection(hover bool) {
	var _arg0 *C.GtkTreeView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkTreeView)(unsafe.Pointer(coreglib.InternObject(treeView).Native()))
	if hover {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_set_hover_selection(_arg0, _arg1)
	runtime.KeepAlive(treeView)
	runtime.KeepAlive(hover)
}

// SetRowSeparatorFunc sets the row separator function, which is used to
// determine whether a row should be drawn as a separator. If the row separator
// function is NULL, no separators are drawn. This is the default value.
//
// The function takes the following parameters:
//
//    - fn (optional): TreeViewRowSeparatorFunc.
//
func (treeView *TreeView) SetRowSeparatorFunc(fn TreeViewRowSeparatorFunc) {
	var _arg0 *C.GtkTreeView                // out
	var _arg1 C.GtkTreeViewRowSeparatorFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkTreeView)(unsafe.Pointer(coreglib.InternObject(treeView).Native()))
	if fn != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_TreeViewRowSeparatorFunc)
		_arg2 = C.gpointer(gbox.Assign(fn))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_tree_view_set_row_separator_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(treeView)
	runtime.KeepAlive(fn)
}
