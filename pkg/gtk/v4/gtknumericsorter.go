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
		{T: externglib.Type(C.gtk_numeric_sorter_get_type()), F: marshalNumericSorter},
	})
}

// NumericSorter: `GtkNumericSorter` is a `GtkSorter` that compares numbers.
//
// To obtain the numbers to compare, this sorter evaluates a
// [class@Gtk.Expression].
type NumericSorter interface {
	gextras.Objector

	// Expression gets the expression that is evaluated to obtain numbers from
	// items.
	Expression() Expression
	// SortOrder gets whether this sorter will sort smaller numbers first.
	SortOrder() SortType
	// SetExpression sets the expression that is evaluated to obtain numbers
	// from items.
	//
	// Unless an expression is set on @self, the sorter will always compare
	// items as invalid.
	//
	// The expression must have a return type that can be compared numerically,
	// such as G_TYPE_INT or G_TYPE_DOUBLE.
	SetExpression(expression Expression)
	// SetSortOrder sets whether to sort smaller numbers before larger ones.
	SetSortOrder(sortOrder SortType)
}

// NumericSorterClass implements the NumericSorter interface.
type NumericSorterClass struct {
	SorterClass
}

var _ NumericSorter = (*NumericSorterClass)(nil)

func wrapNumericSorter(obj *externglib.Object) NumericSorter {
	return &NumericSorterClass{
		SorterClass: SorterClass{
			Object: obj,
		},
	}
}

func marshalNumericSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNumericSorter(obj), nil
}

// NewNumericSorter creates a new numeric sorter using the given @expression.
//
// Smaller numbers will be sorted first. You can call
// [method@Gtk.NumericSorter.set_sort_order] to change this.
func NewNumericSorter(expression Expression) NumericSorter {
	var _arg1 *C.GtkExpression    // out
	var _cret *C.GtkNumericSorter // in

	_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_numeric_sorter_new(_arg1)

	var _numericSorter NumericSorter // out

	_numericSorter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(NumericSorter)

	return _numericSorter
}

// Expression gets the expression that is evaluated to obtain numbers from
// items.
func (s *NumericSorterClass) Expression() Expression {
	var _arg0 *C.GtkNumericSorter // out
	var _cret *C.GtkExpression    // in

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_numeric_sorter_get_expression(_arg0)

	var _expression Expression // out

	_expression = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Expression)

	return _expression
}

// SortOrder gets whether this sorter will sort smaller numbers first.
func (s *NumericSorterClass) SortOrder() SortType {
	var _arg0 *C.GtkNumericSorter // out
	var _cret C.GtkSortType       // in

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_numeric_sorter_get_sort_order(_arg0)

	var _sortType SortType // out

	_sortType = SortType(_cret)

	return _sortType
}

// SetExpression sets the expression that is evaluated to obtain numbers from
// items.
//
// Unless an expression is set on @self, the sorter will always compare items as
// invalid.
//
// The expression must have a return type that can be compared numerically, such
// as G_TYPE_INT or G_TYPE_DOUBLE.
func (s *NumericSorterClass) SetExpression(expression Expression) {
	var _arg0 *C.GtkNumericSorter // out
	var _arg1 *C.GtkExpression    // out

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_numeric_sorter_set_expression(_arg0, _arg1)
}

// SetSortOrder sets whether to sort smaller numbers before larger ones.
func (s *NumericSorterClass) SetSortOrder(sortOrder SortType) {
	var _arg0 *C.GtkNumericSorter // out
	var _arg1 C.GtkSortType       // out

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkSortType(sortOrder)

	C.gtk_numeric_sorter_set_sort_order(_arg0, _arg1)
}
