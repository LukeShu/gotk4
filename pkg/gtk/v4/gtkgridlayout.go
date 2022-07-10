// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeGridLayout returns the GType for the type GridLayout.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGridLayout() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "GridLayout").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGridLayout)
	return gtype
}

// GTypeGridLayoutChild returns the GType for the type GridLayoutChild.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGridLayoutChild() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "GridLayoutChild").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGridLayoutChild)
	return gtype
}

// GridLayoutOverrider contains methods that are overridable.
type GridLayoutOverrider interface {
}

// GridLayout: GtkGridLayout is a layout manager which arranges child widgets in
// rows and columns.
//
// Children have an "attach point" defined by the horizontal and vertical index
// of the cell they occupy; children can span multiple rows or columns. The
// layout properties for setting the attach points and spans are set using the
// gtk.GridLayoutChild associated to each child widget.
//
// The behaviour of GtkGridLayout when several children occupy the same grid
// cell is undefined.
//
// GtkGridLayout can be used like a GtkBoxLayout if all children are attached to
// the same row or column; however, if you only ever need a single row or
// column, you should consider using GtkBoxLayout.
type GridLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*GridLayout)(nil)
)

func classInitGridLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGridLayout(obj *coreglib.Object) *GridLayout {
	return &GridLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalGridLayout(p uintptr) (interface{}, error) {
	return wrapGridLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewGridLayout creates a new GtkGridLayout.
//
// The function returns the following values:
//
//    - gridLayout: newly created GtkGridLayout.
//
func NewGridLayout() *GridLayout {
	_info := girepository.MustFind("Gtk", "GridLayout")
	_gret := _info.InvokeClassMethod("new_GridLayout", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _gridLayout *GridLayout // out

	_gridLayout = wrapGridLayout(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _gridLayout
}

// BaselineRow retrieves the row set with gtk_grid_layout_set_baseline_row().
//
// The function returns the following values:
//
//    - gint: global baseline row.
//
func (grid *GridLayout) BaselineRow() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))

	_info := girepository.MustFind("Gtk", "GridLayout")
	_gret := _info.InvokeClassMethod("get_baseline_row", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// ColumnHomogeneous checks whether all columns of grid should have the same
// width.
//
// The function returns the following values:
//
//    - ok: TRUE if the columns are homogeneous, and FALSE otherwise.
//
func (grid *GridLayout) ColumnHomogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))

	_info := girepository.MustFind("Gtk", "GridLayout")
	_gret := _info.InvokeClassMethod("get_column_homogeneous", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ColumnSpacing retrieves the spacing set with
// gtk_grid_layout_set_column_spacing().
//
// The function returns the following values:
//
//    - guint: spacing between consecutive columns.
//
func (grid *GridLayout) ColumnSpacing() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))

	_info := girepository.MustFind("Gtk", "GridLayout")
	_gret := _info.InvokeClassMethod("get_column_spacing", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// RowHomogeneous checks whether all rows of grid should have the same height.
//
// The function returns the following values:
//
//    - ok: TRUE if the rows are homogeneous, and FALSE otherwise.
//
func (grid *GridLayout) RowHomogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))

	_info := girepository.MustFind("Gtk", "GridLayout")
	_gret := _info.InvokeClassMethod("get_row_homogeneous", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing retrieves the spacing set with gtk_grid_layout_set_row_spacing().
//
// The function returns the following values:
//
//    - guint: spacing between consecutive rows.
//
func (grid *GridLayout) RowSpacing() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))

	_info := girepository.MustFind("Gtk", "GridLayout")
	_gret := _info.InvokeClassMethod("get_row_spacing", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(grid)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// SetBaselineRow sets which row defines the global baseline for the entire
// grid.
//
// Each row in the grid can have its own local baseline, but only one of those
// is global, meaning it will be the baseline in the parent of the grid.
//
// The function takes the following parameters:
//
//    - row index.
//
func (grid *GridLayout) SetBaselineRow(row int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(row)

	_info := girepository.MustFind("Gtk", "GridLayout")
	_info.InvokeClassMethod("set_baseline_row", _args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(row)
}

// SetColumnHomogeneous sets whether all columns of grid should have the same
// width.
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE to make columns homogeneous.
//
func (grid *GridLayout) SetColumnHomogeneous(homogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	if homogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "GridLayout")
	_info.InvokeClassMethod("set_column_homogeneous", _args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(homogeneous)
}

// SetColumnSpacing sets the amount of space to insert between consecutive
// columns.
//
// The function takes the following parameters:
//
//    - spacing: amount of space between columns, in pixels.
//
func (grid *GridLayout) SetColumnSpacing(spacing uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(spacing)

	_info := girepository.MustFind("Gtk", "GridLayout")
	_info.InvokeClassMethod("set_column_spacing", _args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(spacing)
}

// SetRowHomogeneous sets whether all rows of grid should have the same height.
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE to make rows homogeneous.
//
func (grid *GridLayout) SetRowHomogeneous(homogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	if homogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "GridLayout")
	_info.InvokeClassMethod("set_row_homogeneous", _args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(homogeneous)
}

// SetRowSpacing sets the amount of space to insert between consecutive rows.
//
// The function takes the following parameters:
//
//    - spacing: amount of space between rows, in pixels.
//
func (grid *GridLayout) SetRowSpacing(spacing uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(grid).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(spacing)

	_info := girepository.MustFind("Gtk", "GridLayout")
	_info.InvokeClassMethod("set_row_spacing", _args[:], nil)

	runtime.KeepAlive(grid)
	runtime.KeepAlive(spacing)
}

// GridLayoutChildOverrider contains methods that are overridable.
type GridLayoutChildOverrider interface {
}

// GridLayoutChild: GtkLayoutChild subclass for children in a GtkGridLayout.
type GridLayoutChild struct {
	_ [0]func() // equal guard
	LayoutChild
}

var (
	_ LayoutChilder = (*GridLayoutChild)(nil)
)

func classInitGridLayoutChilder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGridLayoutChild(obj *coreglib.Object) *GridLayoutChild {
	return &GridLayoutChild{
		LayoutChild: LayoutChild{
			Object: obj,
		},
	}
}

func marshalGridLayoutChild(p uintptr) (interface{}, error) {
	return wrapGridLayoutChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Column retrieves the column number to which child attaches its left side.
//
// The function returns the following values:
//
//    - gint: column number.
//
func (child *GridLayoutChild) Column() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_gret := _info.InvokeClassMethod("get_column", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(child)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// ColumnSpan retrieves the number of columns that child spans to.
//
// The function returns the following values:
//
//    - gint: number of columns.
//
func (child *GridLayoutChild) ColumnSpan() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_gret := _info.InvokeClassMethod("get_column_span", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(child)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Row retrieves the row number to which child attaches its top side.
//
// The function returns the following values:
//
//    - gint: row number.
//
func (child *GridLayoutChild) Row() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_gret := _info.InvokeClassMethod("get_row", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(child)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// RowSpan retrieves the number of rows that child spans to.
//
// The function returns the following values:
//
//    - gint: number of row.
//
func (child *GridLayoutChild) RowSpan() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_gret := _info.InvokeClassMethod("get_row_span", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(child)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// SetColumn sets the column number to attach the left side of child.
//
// The function takes the following parameters:
//
//    - column: attach point for child.
//
func (child *GridLayoutChild) SetColumn(column int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(column)

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_info.InvokeClassMethod("set_column", _args[:], nil)

	runtime.KeepAlive(child)
	runtime.KeepAlive(column)
}

// SetColumnSpan sets the number of columns child spans to.
//
// The function takes the following parameters:
//
//    - span of child.
//
func (child *GridLayoutChild) SetColumnSpan(span int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(span)

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_info.InvokeClassMethod("set_column_span", _args[:], nil)

	runtime.KeepAlive(child)
	runtime.KeepAlive(span)
}

// SetRow sets the row to place child in.
//
// The function takes the following parameters:
//
//    - row for child.
//
func (child *GridLayoutChild) SetRow(row int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(row)

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_info.InvokeClassMethod("set_row", _args[:], nil)

	runtime.KeepAlive(child)
	runtime.KeepAlive(row)
}

// SetRowSpan sets the number of rows child spans to.
//
// The function takes the following parameters:
//
//    - span of child.
//
func (child *GridLayoutChild) SetRowSpan(span int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(span)

	_info := girepository.MustFind("Gtk", "GridLayoutChild")
	_info.InvokeClassMethod("set_row_span", _args[:], nil)

	runtime.KeepAlive(child)
	runtime.KeepAlive(span)
}
