// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scrollable_get_type()), F: marshalScrollabler},
	})
}

// ScrollableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScrollableOverrider interface {
	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable.
	//
	// An example for this would be treeview headers. GTK can use this
	// information to display overlaid graphics, like the overshoot indication,
	// at the right position.
	Border() (Border, bool)
}

// Scrollabler describes Scrollable's methods.
type Scrollabler interface {
	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable.
	Border() (Border, bool)
	// HAdjustment retrieves the `GtkAdjustment` used for horizontal scrolling.
	HAdjustment() *Adjustment
	// HscrollPolicy gets the horizontal `GtkScrollablePolicy`.
	HscrollPolicy() ScrollablePolicy
	// VAdjustment retrieves the `GtkAdjustment` used for vertical scrolling.
	VAdjustment() *Adjustment
	// VscrollPolicy gets the vertical `GtkScrollablePolicy`.
	VscrollPolicy() ScrollablePolicy
	// SetHAdjustment sets the horizontal adjustment of the `GtkScrollable`.
	SetHAdjustment(hadjustment Adjustmenter)
	// SetHscrollPolicy sets the `GtkScrollablePolicy`.
	SetHscrollPolicy(policy ScrollablePolicy)
	// SetVAdjustment sets the vertical adjustment of the `GtkScrollable`.
	SetVAdjustment(vadjustment Adjustmenter)
	// SetVscrollPolicy sets the `GtkScrollablePolicy`.
	SetVscrollPolicy(policy ScrollablePolicy)
}

// Scrollable: `GtkScrollable` is an interface for widgets with native scrolling
// ability.
//
// To implement this interface you should override the
// [property@Gtk.Scrollable:hadjustment] and
// [property@Gtk.Scrollable:vadjustment] properties.
//
//
// Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments, the
// widget should populate the adjustments’ [property@Gtk.Adjustment:lower],
// [property@Gtk.Adjustment:upper], [property@Gtk.Adjustment:step-increment],
// [property@Gtk.Adjustment:page-increment] and
// [property@Gtk.Adjustment:page-size] properties and connect to the
// [signal@Gtk.Adjustment::value-changed] signal.
//
// - Because its preferred size is the size for a fully expanded widget, the
// scrollable widget must be able to cope with underallocations. This means that
// it must accept any value passed to its GtkWidgetClass.size_allocate()
// function.
//
// - When the parent allocates space to the scrollable child widget, the widget
// should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the
// [signal@Gtk.Adjustment::value-changed] signal, the scrollable widget should
// scroll its contents.
type Scrollable struct {
	*externglib.Object
}

var (
	_ Scrollabler     = (*Scrollable)(nil)
	_ gextras.Nativer = (*Scrollable)(nil)
)

func wrapScrollable(obj *externglib.Object) Scrollabler {
	return &Scrollable{
		Object: obj,
	}
}

func marshalScrollabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScrollable(obj), nil
}

// Border returns the size of a non-scrolling border around the outside of the
// scrollable.
//
// An example for this would be treeview headers. GTK can use this information
// to display overlaid graphics, like the overshoot indication, at the right
// position.
func (scrollable *Scrollable) Border() (Border, bool) {
	var _arg0 *C.GtkScrollable // out
	var _border Border
	var _cret C.gboolean // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_border(_arg0, (*C.GtkBorder)(unsafe.Pointer(&_border)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _border, _ok
}

// HAdjustment retrieves the `GtkAdjustment` used for horizontal scrolling.
func (scrollable *Scrollable) HAdjustment() *Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_hadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// HscrollPolicy gets the horizontal `GtkScrollablePolicy`.
func (scrollable *Scrollable) HscrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_hscroll_policy(_arg0)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

// VAdjustment retrieves the `GtkAdjustment` used for vertical scrolling.
func (scrollable *Scrollable) VAdjustment() *Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_vadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// VscrollPolicy gets the vertical `GtkScrollablePolicy`.
func (scrollable *Scrollable) VscrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_vscroll_policy(_arg0)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

// SetHAdjustment sets the horizontal adjustment of the `GtkScrollable`.
func (scrollable *Scrollable) SetHAdjustment(hadjustment Adjustmenter) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((hadjustment).(gextras.Nativer).Native()))

	C.gtk_scrollable_set_hadjustment(_arg0, _arg1)
}

// SetHscrollPolicy sets the `GtkScrollablePolicy`.
//
// The policy determines whether horizontal scrolling should start below the
// minimum width or below the natural width.
func (scrollable *Scrollable) SetHscrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_hscroll_policy(_arg0, _arg1)
}

// SetVAdjustment sets the vertical adjustment of the `GtkScrollable`.
func (scrollable *Scrollable) SetVAdjustment(vadjustment Adjustmenter) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((vadjustment).(gextras.Nativer).Native()))

	C.gtk_scrollable_set_vadjustment(_arg0, _arg1)
}

// SetVscrollPolicy sets the `GtkScrollablePolicy`.
//
// The policy determines whether vertical scrolling should start below the
// minimum height or below the natural height.
func (scrollable *Scrollable) SetVscrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_vscroll_policy(_arg0, _arg1)
}
