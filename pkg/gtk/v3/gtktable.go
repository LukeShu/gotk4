// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_attach_options_get_type()), F: marshalAttachOptions},
		{T: externglib.Type(C.gtk_table_get_type()), F: marshalTable},
	})
}

// AttachOptions denotes the expansion properties that a widget will have when
// it (or its parent) is resized.
type AttachOptions int

const (
	// AttachOptionsExpand: the widget should expand to take up any extra space
	// in its container that has been allocated.
	AttachOptionsExpand AttachOptions = 0b1
	// AttachOptionsShrink: the widget should shrink as and when possible.
	AttachOptionsShrink AttachOptions = 0b10
	// AttachOptionsFill: the widget should fill the space allocated to it.
	AttachOptionsFill AttachOptions = 0b100
)

func marshalAttachOptions(p uintptr) (interface{}, error) {
	return AttachOptions(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Table: the Table functions allow the programmer to arrange widgets in rows
// and columns, making it easy to align many widgets next to each other,
// horizontally and vertically.
//
// Tables are created with a call to gtk_table_new(), the size of which can
// later be changed with gtk_table_resize().
//
// Widgets can be added to a table using gtk_table_attach() or the more
// convenient (but slightly less flexible) gtk_table_attach_defaults().
//
// To alter the space next to a specific row, use gtk_table_set_row_spacing(),
// and for a column, gtk_table_set_col_spacing(). The gaps between all rows or
// columns can be changed by calling gtk_table_set_row_spacings() or
// gtk_table_set_col_spacings() respectively. Note that spacing is added between
// the children, while padding added by gtk_table_attach() is added on either
// side of the widget it belongs to.
//
// gtk_table_set_homogeneous(), can be used to set whether all cells in the
// table will resize themselves to the size of the largest widget in the table.
//
// > Table has been deprecated. Use Grid instead. It provides the same >
// capabilities as GtkTable for arranging widgets in a rectangular grid, but >
// does support height-for-width geometry management.
type Table interface {
	gextras.Objector

	// Attach adds a widget to a table. The number of “cells” that a widget will
	// occupy is specified by @left_attach, @right_attach, @top_attach and
	// @bottom_attach. These each represent the leftmost, rightmost, uppermost
	// and lowest column and row numbers of the table. (Columns and rows are
	// indexed from zero).
	//
	// To make a button occupy the lower right cell of a 2x2 table, use
	//
	//    gtk_table_attach (table, button,
	//                      1, 2, // left, right attach
	//                      1, 2, // top, bottom attach
	//                      xoptions, yoptions,
	//                      xpadding, ypadding);
	//
	// If you want to make the button span the entire bottom row, use
	// @left_attach == 0 and @right_attach = 2 instead.
	//
	// Deprecated: since version 3.4.
	Attach(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint)
	// AttachDefaults as there are many options associated with
	// gtk_table_attach(), this convenience function provides the programmer
	// with a means to add children to a table with identical padding and
	// expansion options. The values used for the AttachOptions are `GTK_EXPAND
	// | GTK_FILL`, and the padding is set to 0.
	//
	// Deprecated: since version 3.4.
	AttachDefaults(widget Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint)
	// ColSpacing gets the amount of space between column @col, and column @col
	// + 1. See gtk_table_set_col_spacing().
	//
	// Deprecated: since version 3.4.
	ColSpacing(column uint) uint
	// DefaultColSpacing gets the default column spacing for the table. This is
	// the spacing that will be used for newly added columns. (See
	// gtk_table_set_col_spacings())
	//
	// Deprecated: since version 3.4.
	DefaultColSpacing() uint
	// DefaultRowSpacing gets the default row spacing for the table. This is the
	// spacing that will be used for newly added rows. (See
	// gtk_table_set_row_spacings())
	//
	// Deprecated: since version 3.4.
	DefaultRowSpacing() uint
	// Homogeneous returns whether the table cells are all constrained to the
	// same width and height. (See gtk_table_set_homogeneous ())
	//
	// Deprecated: since version 3.4.
	Homogeneous() bool
	// RowSpacing gets the amount of space between row @row, and row @row + 1.
	// See gtk_table_set_row_spacing().
	//
	// Deprecated: since version 3.4.
	RowSpacing(row uint) uint
	// Size gets the number of rows and columns in the table.
	//
	// Deprecated: since version 3.4.
	Size() (rows uint, columns uint)
	// Resize: if you need to change a table’s size after it has been created,
	// this function allows you to do so.
	//
	// Deprecated: since version 3.4.
	Resize(rows uint, columns uint)
	// SetColSpacing alters the amount of space between a given table column and
	// the following column.
	//
	// Deprecated: since version 3.4.
	SetColSpacing(column uint, spacing uint)
	// SetColSpacings sets the space between every column in @table equal to
	// @spacing.
	//
	// Deprecated: since version 3.4.
	SetColSpacings(spacing uint)
	// SetHomogeneous changes the homogenous property of table cells, ie.
	// whether all cells are an equal size or not.
	//
	// Deprecated: since version 3.4.
	SetHomogeneous(homogeneous bool)
	// SetRowSpacing changes the space between a given table row and the
	// subsequent row.
	//
	// Deprecated: since version 3.4.
	SetRowSpacing(row uint, spacing uint)
	// SetRowSpacings sets the space between every row in @table equal to
	// @spacing.
	//
	// Deprecated: since version 3.4.
	SetRowSpacings(spacing uint)
}

// TableClass implements the Table interface.
type TableClass struct {
	*externglib.Object
	ContainerClass
	BuildableInterface
}

var _ Table = (*TableClass)(nil)

func wrapTable(obj *externglib.Object) Table {
	return &TableClass{
		Object: obj,
		ContainerClass: ContainerClass{
			Object: obj,
			WidgetClass: WidgetClass{
				Object:           obj,
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalTable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTable(obj), nil
}

// NewTable: used to create a new table widget. An initial size must be given by
// specifying how many rows and columns the table should have, although this can
// be changed later with gtk_table_resize(). @rows and @columns must both be in
// the range 1 .. 65535. For historical reasons, 0 is accepted as well and is
// silently interpreted as 1.
//
// Deprecated: since version 3.4.
func NewTable(rows uint, columns uint, homogeneous bool) Table {
	var _arg1 C.guint      // out
	var _arg2 C.guint      // out
	var _arg3 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.guint(rows)
	_arg2 = C.guint(columns)
	if homogeneous {
		_arg3 = C.TRUE
	}

	_cret = C.gtk_table_new(_arg1, _arg2, _arg3)

	var _table Table // out

	_table = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Table)

	return _table
}

// Attach adds a widget to a table. The number of “cells” that a widget will
// occupy is specified by @left_attach, @right_attach, @top_attach and
// @bottom_attach. These each represent the leftmost, rightmost, uppermost and
// lowest column and row numbers of the table. (Columns and rows are indexed
// from zero).
//
// To make a button occupy the lower right cell of a 2x2 table, use
//
//    gtk_table_attach (table, button,
//                      1, 2, // left, right attach
//                      1, 2, // top, bottom attach
//                      xoptions, yoptions,
//                      xpadding, ypadding);
//
// If you want to make the button span the entire bottom row, use @left_attach
// == 0 and @right_attach = 2 instead.
//
// Deprecated: since version 3.4.
func (t *TableClass) Attach(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint) {
	var _arg0 *C.GtkTable        // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.guint            // out
	var _arg3 C.guint            // out
	var _arg4 C.guint            // out
	var _arg5 C.guint            // out
	var _arg6 C.GtkAttachOptions // out
	var _arg7 C.GtkAttachOptions // out
	var _arg8 C.guint            // out
	var _arg9 C.guint            // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)
	_arg6 = C.GtkAttachOptions(xoptions)
	_arg7 = C.GtkAttachOptions(yoptions)
	_arg8 = C.guint(xpadding)
	_arg9 = C.guint(ypadding)

	C.gtk_table_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// AttachDefaults as there are many options associated with gtk_table_attach(),
// this convenience function provides the programmer with a means to add
// children to a table with identical padding and expansion options. The values
// used for the AttachOptions are `GTK_EXPAND | GTK_FILL`, and the padding is
// set to 0.
//
// Deprecated: since version 3.4.
func (t *TableClass) AttachDefaults(widget Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint) {
	var _arg0 *C.GtkTable  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.guint      // out
	var _arg3 C.guint      // out
	var _arg4 C.guint      // out
	var _arg5 C.guint      // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)

	C.gtk_table_attach_defaults(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// ColSpacing gets the amount of space between column @col, and column @col + 1.
// See gtk_table_set_col_spacing().
//
// Deprecated: since version 3.4.
func (t *TableClass) ColSpacing(column uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(column)

	_cret = C.gtk_table_get_col_spacing(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// DefaultColSpacing gets the default column spacing for the table. This is the
// spacing that will be used for newly added columns. (See
// gtk_table_set_col_spacings())
//
// Deprecated: since version 3.4.
func (t *TableClass) DefaultColSpacing() uint {
	var _arg0 *C.GtkTable // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_table_get_default_col_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// DefaultRowSpacing gets the default row spacing for the table. This is the
// spacing that will be used for newly added rows. (See
// gtk_table_set_row_spacings())
//
// Deprecated: since version 3.4.
func (t *TableClass) DefaultRowSpacing() uint {
	var _arg0 *C.GtkTable // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_table_get_default_row_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Homogeneous returns whether the table cells are all constrained to the same
// width and height. (See gtk_table_set_homogeneous ())
//
// Deprecated: since version 3.4.
func (t *TableClass) Homogeneous() bool {
	var _arg0 *C.GtkTable // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_table_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing gets the amount of space between row @row, and row @row + 1. See
// gtk_table_set_row_spacing().
//
// Deprecated: since version 3.4.
func (t *TableClass) RowSpacing(row uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(row)

	_cret = C.gtk_table_get_row_spacing(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Size gets the number of rows and columns in the table.
//
// Deprecated: since version 3.4.
func (t *TableClass) Size() (rows uint, columns uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // in
	var _arg2 C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	C.gtk_table_get_size(_arg0, &_arg1, &_arg2)

	var _rows uint    // out
	var _columns uint // out

	_rows = uint(_arg1)
	_columns = uint(_arg2)

	return _rows, _columns
}

// Resize: if you need to change a table’s size after it has been created, this
// function allows you to do so.
//
// Deprecated: since version 3.4.
func (t *TableClass) Resize(rows uint, columns uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(rows)
	_arg2 = C.guint(columns)

	C.gtk_table_resize(_arg0, _arg1, _arg2)
}

// SetColSpacing alters the amount of space between a given table column and the
// following column.
//
// Deprecated: since version 3.4.
func (t *TableClass) SetColSpacing(column uint, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(column)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_col_spacing(_arg0, _arg1, _arg2)
}

// SetColSpacings sets the space between every column in @table equal to
// @spacing.
//
// Deprecated: since version 3.4.
func (t *TableClass) SetColSpacings(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_col_spacings(_arg0, _arg1)
}

// SetHomogeneous changes the homogenous property of table cells, ie. whether
// all cells are an equal size or not.
//
// Deprecated: since version 3.4.
func (t *TableClass) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_table_set_homogeneous(_arg0, _arg1)
}

// SetRowSpacing changes the space between a given table row and the subsequent
// row.
//
// Deprecated: since version 3.4.
func (t *TableClass) SetRowSpacing(row uint, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(row)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_row_spacing(_arg0, _arg1, _arg2)
}

// SetRowSpacings sets the space between every row in @table equal to @spacing.
//
// Deprecated: since version 3.4.
func (t *TableClass) SetRowSpacings(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_row_spacings(_arg0, _arg1)
}

type TableChild struct {
	native C.GtkTableChild
}

// WrapTableChild wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTableChild(ptr unsafe.Pointer) *TableChild {
	return (*TableChild)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TableChild) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

func (t *TableChild) Widget() Widget {
	var v Widget // out
	v = gextras.CastObject(externglib.Take(unsafe.Pointer(t.native.widget))).(Widget)
	return v
}

func (t *TableChild) LeftAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.left_attach)
	return v
}

func (t *TableChild) RightAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.right_attach)
	return v
}

func (t *TableChild) TopAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.top_attach)
	return v
}

func (t *TableChild) BottomAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.bottom_attach)
	return v
}

func (t *TableChild) Xpadding() uint16 {
	var v uint16 // out
	v = uint16(t.native.xpadding)
	return v
}

func (t *TableChild) Ypadding() uint16 {
	var v uint16 // out
	v = uint16(t.native.ypadding)
	return v
}

type TableRowCol struct {
	native C.GtkTableRowCol
}

// WrapTableRowCol wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTableRowCol(ptr unsafe.Pointer) *TableRowCol {
	return (*TableRowCol)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TableRowCol) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

func (t *TableRowCol) Requisition() uint16 {
	var v uint16 // out
	v = uint16(t.native.requisition)
	return v
}

func (t *TableRowCol) Allocation() uint16 {
	var v uint16 // out
	v = uint16(t.native.allocation)
	return v
}

func (t *TableRowCol) Spacing() uint16 {
	var v uint16 // out
	v = uint16(t.native.spacing)
	return v
}
