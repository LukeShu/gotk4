// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_sorter_change_get_type()), F: marshalSorterChange},
		{T: externglib.Type(C.gtk_sorter_order_get_type()), F: marshalSorterOrder},
		{T: externglib.Type(C.gtk_sorter_get_type()), F: marshalSorterer},
	})
}

// SorterChange describes changes in a sorter in more detail and allows users to
// optimize resorting.
type SorterChange int

const (
	// SorterChangeDifferent: sorter change cannot be described by any of the
	// other enumeration values
	SorterChangeDifferent SorterChange = iota
	// SorterChangeInverted: sort order was inverted. Comparisons that returned
	// GTK_ORDERING_SMALLER now return GTK_ORDERING_LARGER and vice versa. Other
	// comparisons return the same values as before.
	SorterChangeInverted
	// SorterChangeLessStrict: sorter is less strict: Comparisons may now return
	// GTK_ORDERING_EQUAL that did not do so before.
	SorterChangeLessStrict
	// SorterChangeMoreStrict: sorter is more strict: Comparisons that did
	// return GTK_ORDERING_EQUAL may not do so anymore.
	SorterChangeMoreStrict
)

func marshalSorterChange(p uintptr) (interface{}, error) {
	return SorterChange(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for SorterChange.
func (s SorterChange) String() string {
	switch s {
	case SorterChangeDifferent:
		return "Different"
	case SorterChangeInverted:
		return "Inverted"
	case SorterChangeLessStrict:
		return "LessStrict"
	case SorterChangeMoreStrict:
		return "MoreStrict"
	default:
		return fmt.Sprintf("SorterChange(%d)", s)
	}
}

// SorterOrder describes the type of order that a GtkSorter may produce.
type SorterOrder int

const (
	// SorterOrderPartial order. Any Ordering is possible.
	SorterOrderPartial SorterOrder = iota
	// SorterOrderNone: no order, all elements are considered equal.
	// gtk_sorter_compare() will only return GTK_ORDERING_EQUAL.
	SorterOrderNone
	// SorterOrderTotal order. gtk_sorter_compare() will only return
	// GTK_ORDERING_EQUAL if an item is compared with itself. Two different
	// items will never cause this value to be returned.
	SorterOrderTotal
)

func marshalSorterOrder(p uintptr) (interface{}, error) {
	return SorterOrder(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for SorterOrder.
func (s SorterOrder) String() string {
	switch s {
	case SorterOrderPartial:
		return "Partial"
	case SorterOrderNone:
		return "None"
	case SorterOrderTotal:
		return "Total"
	default:
		return fmt.Sprintf("SorterOrder(%d)", s)
	}
}

// SorterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SorterOverrider interface {
	// Compare compares two given items according to the sort order implemented
	// by the sorter.
	//
	// Sorters implement a partial order:
	//
	// * It is reflexive, ie a = a * It is antisymmetric, ie if a < b and b < a,
	// then a = b * It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
	// then a ≤ c
	//
	// The sorter may signal it conforms to additional constraints via the
	// return value of gtk.Sorter.GetOrder().
	Compare(item1 *externglib.Object, item2 *externglib.Object) Ordering
	// Order gets the order that self conforms to.
	//
	// See gtk.SorterOrder for details of the possible return values.
	//
	// This function is intended to allow optimizations.
	Order() SorterOrder
}

// Sorter: GtkSorter is an object to describe sorting criteria.
//
//
// Its primary user is gtk.SortListModel
//
// The model will use a sorter to determine the order in which its items should
// appear by calling gtk.Sorter.Compare() for pairs of items.
//
// Sorters may change their sorting behavior through their lifetime. In that
// case, they will emit the gtk.Sorter::changed signal to notify that the sort
// order is no longer valid and should be updated by calling
// gtk_sorter_compare() again.
//
// GTK provides various pre-made sorter implementations for common sorting
// operations. gtk.ColumnView has built-in support for sorting lists via the
// gtk.ColumnViewColumn:sorter property, where the user can change the sorting
// by clicking on list headers.
//
// Of course, in particular for large lists, it is also possible to subclass
// GtkSorter and provide one's own sorter.
type Sorter struct {
	*externglib.Object
}

func wrapSorter(obj *externglib.Object) *Sorter {
	return &Sorter{
		Object: obj,
	}
}

func marshalSorterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSorter(obj), nil
}

// Changed emits the gtk.Sorter::changed signal to notify all users of the
// sorter that it has changed.
//
// Users of the sorter should then update the sort order via
// gtk_sorter_compare().
//
// Depending on the change parameter, it may be possible to update the sort
// order without a full resorting. Refer to the gtk.SorterChange documentation
// for details.
//
// This function is intended for implementors of GtkSorter subclasses and should
// not be called from other functions.
func (self *Sorter) Changed(change SorterChange) {
	var _arg0 *C.GtkSorter      // out
	var _arg1 C.GtkSorterChange // out

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkSorterChange(change)

	C.gtk_sorter_changed(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(change)
}

// Compare compares two given items according to the sort order implemented by
// the sorter.
//
// Sorters implement a partial order:
//
// * It is reflexive, ie a = a * It is antisymmetric, ie if a < b and b < a,
// then a = b * It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
// then a ≤ c
//
// The sorter may signal it conforms to additional constraints via the return
// value of gtk.Sorter.GetOrder().
func (self *Sorter) Compare(item1 *externglib.Object, item2 *externglib.Object) Ordering {
	var _arg0 *C.GtkSorter  // out
	var _arg1 C.gpointer    // out
	var _arg2 C.gpointer    // out
	var _cret C.GtkOrdering // in

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(self.Native()))
	_arg1 = C.gpointer(unsafe.Pointer(item1.Native()))
	_arg2 = C.gpointer(unsafe.Pointer(item2.Native()))

	_cret = C.gtk_sorter_compare(_arg0, _arg1, _arg2)

	runtime.KeepAlive(self)
	runtime.KeepAlive(item1)
	runtime.KeepAlive(item2)

	var _ordering Ordering // out

	_ordering = Ordering(_cret)

	return _ordering
}

// Order gets the order that self conforms to.
//
// See gtk.SorterOrder for details of the possible return values.
//
// This function is intended to allow optimizations.
func (self *Sorter) Order() SorterOrder {
	var _arg0 *C.GtkSorter     // out
	var _cret C.GtkSorterOrder // in

	_arg0 = (*C.GtkSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_sorter_get_order(_arg0)

	runtime.KeepAlive(self)

	var _sorterOrder SorterOrder // out

	_sorterOrder = SorterOrder(_cret)

	return _sorterOrder
}
