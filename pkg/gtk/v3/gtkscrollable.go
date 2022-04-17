// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_ScrollableInterface_get_border(GtkScrollable*, GtkBorder*);
import "C"

// glib.Type values for gtkscrollable.go.
var GTypeScrollable = externglib.Type(C.gtk_scrollable_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeScrollable, F: marshalScrollable},
	})
}

// ScrollableOverrider contains methods that are overridable.
type ScrollableOverrider interface {
	externglib.Objector
	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable. An example for this would be treeview headers. GTK+ can
	// use this information to display overlayed graphics, like the overshoot
	// indication, at the right position.
	//
	// The function returns the following values:
	//
	//    - border: return location for the results.
	//    - ok: TRUE if border has been set.
	//
	Border() (*Border, bool)
}

// WrapScrollableOverrider wraps the ScrollableOverrider
// interface implementation to access the instance methods.
func WrapScrollableOverrider(obj ScrollableOverrider) *Scrollable {
	return wrapScrollable(externglib.BaseObject(obj))
}

// Scrollable is an interface that is implemented by widgets with native
// scrolling ability.
//
// To implement this interface you should override the Scrollable:hadjustment
// and Scrollable:vadjustment properties.
//
//
// Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments, the
// widget should populate the adjustments’ Adjustment:lower, Adjustment:upper,
// Adjustment:step-increment, Adjustment:page-increment and Adjustment:page-size
// properties and connect to the Adjustment::value-changed signal.
//
// - Because its preferred size is the size for a fully expanded widget, the
// scrollable widget must be able to cope with underallocations. This means that
// it must accept any value passed to its WidgetClass.size_allocate() function.
//
// - When the parent allocates space to the scrollable child widget, the widget
// should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the Adjustment::value-changed signal, the
// scrollable widget should scroll its contents.
type Scrollable struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Scrollable)(nil)
)

// Scrollabler describes Scrollable's interface methods.
type Scrollabler interface {
	externglib.Objector

	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable.
	Border() (*Border, bool)
	// HAdjustment retrieves the Adjustment used for horizontal scrolling.
	HAdjustment() *Adjustment
	// HScrollPolicy gets the horizontal ScrollablePolicy.
	HScrollPolicy() ScrollablePolicy
	// VAdjustment retrieves the Adjustment used for vertical scrolling.
	VAdjustment() *Adjustment
	// VScrollPolicy gets the vertical ScrollablePolicy.
	VScrollPolicy() ScrollablePolicy
	// SetHAdjustment sets the horizontal adjustment of the Scrollable.
	SetHAdjustment(hadjustment *Adjustment)
	// SetHScrollPolicy sets the ScrollablePolicy to determine whether
	// horizontal scrolling should start below the minimum width or below the
	// natural width.
	SetHScrollPolicy(policy ScrollablePolicy)
	// SetVAdjustment sets the vertical adjustment of the Scrollable.
	SetVAdjustment(vadjustment *Adjustment)
	// SetVScrollPolicy sets the ScrollablePolicy to determine whether vertical
	// scrolling should start below the minimum height or below the natural
	// height.
	SetVScrollPolicy(policy ScrollablePolicy)
}

var _ Scrollabler = (*Scrollable)(nil)

func ifaceInitScrollabler(gifacePtr, data C.gpointer) {
	iface := (*C.GtkScrollableInterface)(unsafe.Pointer(gifacePtr))
	iface.get_border = (*[0]byte)(C._gotk4_gtk3_ScrollableInterface_get_border)
}

//export _gotk4_gtk3_ScrollableInterface_get_border
func _gotk4_gtk3_ScrollableInterface_get_border(arg0 *C.GtkScrollable, arg1 *C.GtkBorder) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ScrollableOverrider)

	border, ok := iface.Border()

	*arg1 = *(*C.GtkBorder)(gextras.StructNative(unsafe.Pointer(border)))
	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapScrollable(obj *externglib.Object) *Scrollable {
	return &Scrollable{
		Object: obj,
	}
}

func marshalScrollable(p uintptr) (interface{}, error) {
	return wrapScrollable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Border returns the size of a non-scrolling border around the outside of the
// scrollable. An example for this would be treeview headers. GTK+ can use this
// information to display overlayed graphics, like the overshoot indication, at
// the right position.
//
// The function returns the following values:
//
//    - border: return location for the results.
//    - ok: TRUE if border has been set.
//
func (scrollable *Scrollable) Border() (*Border, bool) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 C.GtkBorder      // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))

	_cret = C.gtk_scrollable_get_border(_arg0, &_arg1)
	runtime.KeepAlive(scrollable)

	var _border *Border // out
	var _ok bool        // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _border, _ok
}

// HAdjustment retrieves the Adjustment used for horizontal scrolling.
//
// The function returns the following values:
//
//    - adjustment: horizontal Adjustment.
//
func (scrollable *Scrollable) HAdjustment() *Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))

	_cret = C.gtk_scrollable_get_hadjustment(_arg0)
	runtime.KeepAlive(scrollable)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// HScrollPolicy gets the horizontal ScrollablePolicy.
//
// The function returns the following values:
//
//    - scrollablePolicy: horizontal ScrollablePolicy.
//
func (scrollable *Scrollable) HScrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))

	_cret = C.gtk_scrollable_get_hscroll_policy(_arg0)
	runtime.KeepAlive(scrollable)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

// VAdjustment retrieves the Adjustment used for vertical scrolling.
//
// The function returns the following values:
//
//    - adjustment: vertical Adjustment.
//
func (scrollable *Scrollable) VAdjustment() *Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))

	_cret = C.gtk_scrollable_get_vadjustment(_arg0)
	runtime.KeepAlive(scrollable)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// VScrollPolicy gets the vertical ScrollablePolicy.
//
// The function returns the following values:
//
//    - scrollablePolicy: vertical ScrollablePolicy.
//
func (scrollable *Scrollable) VScrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))

	_cret = C.gtk_scrollable_get_vscroll_policy(_arg0)
	runtime.KeepAlive(scrollable)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

// SetHAdjustment sets the horizontal adjustment of the Scrollable.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): Adjustment.
//
func (scrollable *Scrollable) SetHAdjustment(hadjustment *Adjustment) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))
	if hadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(hadjustment).Native()))
	}

	C.gtk_scrollable_set_hadjustment(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(hadjustment)
}

// SetHScrollPolicy sets the ScrollablePolicy to determine whether horizontal
// scrolling should start below the minimum width or below the natural width.
//
// The function takes the following parameters:
//
//    - policy: horizontal ScrollablePolicy.
//
func (scrollable *Scrollable) SetHScrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_hscroll_policy(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(policy)
}

// SetVAdjustment sets the vertical adjustment of the Scrollable.
//
// The function takes the following parameters:
//
//    - vadjustment (optional): Adjustment.
//
func (scrollable *Scrollable) SetVAdjustment(vadjustment *Adjustment) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))
	if vadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(externglib.InternObject(vadjustment).Native()))
	}

	C.gtk_scrollable_set_vadjustment(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(vadjustment)
}

// SetVScrollPolicy sets the ScrollablePolicy to determine whether vertical
// scrolling should start below the minimum height or below the natural height.
//
// The function takes the following parameters:
//
//    - policy: vertical ScrollablePolicy.
//
func (scrollable *Scrollable) SetVScrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(externglib.InternObject(scrollable).Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_vscroll_policy(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(policy)
}
