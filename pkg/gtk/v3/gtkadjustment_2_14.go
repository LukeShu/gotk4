// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// Configure sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the Adjustment::changed
// signal. See gtk_adjustment_set_lower() for an alternative way of compressing
// multiple emissions of Adjustment::changed into one.
//
// The function takes the following parameters:
//
//    - value: new value.
//    - lower: new minimum value.
//    - upper: new maximum value.
//    - stepIncrement: new step increment.
//    - pageIncrement: new page increment.
//    - pageSize: new page size.
//
func (adjustment *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _arg5 C.gdouble        // out
	var _arg6 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(value)
	_arg2 = C.gdouble(lower)
	_arg3 = C.gdouble(upper)
	_arg4 = C.gdouble(stepIncrement)
	_arg5 = C.gdouble(pageIncrement)
	_arg6 = C.gdouble(pageSize)

	C.gtk_adjustment_configure(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(value)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
	runtime.KeepAlive(stepIncrement)
	runtime.KeepAlive(pageIncrement)
	runtime.KeepAlive(pageSize)
}

// Lower retrieves the minimum value of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current minimum value of the adjustment.
//
func (adjustment *Adjustment) Lower() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_lower(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageIncrement retrieves the page increment of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current page increment of the adjustment.
//
func (adjustment *Adjustment) PageIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_page_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageSize retrieves the page size of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current page size of the adjustment.
//
func (adjustment *Adjustment) PageSize() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_page_size(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// StepIncrement retrieves the step increment of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current step increment of the adjustment.
//
func (adjustment *Adjustment) StepIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_step_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Upper retrieves the maximum value of the adjustment.
//
// The function returns the following values:
//
//    - gdouble: current maximum value of the adjustment.
//
func (adjustment *Adjustment) Upper() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_cret = C.gtk_adjustment_get_upper(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetLower sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual setters,
// multiple Adjustment::changed signals will be emitted. However, since the
// emission of the Adjustment::changed signal is tied to the emission of the
// #GObject::notify signals of the changed properties, it’s possible to compress
// the Adjustment::changed signals into one by calling g_object_freeze_notify()
// and g_object_thaw_notify() around the calls to the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties to
// change, or using gtk_adjustment_configure() has the same effect of
// compressing Adjustment::changed emissions.
//
// The function takes the following parameters:
//
//    - lower: new minimum value.
//
func (adjustment *Adjustment) SetLower(lower float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(lower)

	C.gtk_adjustment_set_lower(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(lower)
}

// SetPageIncrement sets the page increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - pageIncrement: new page increment.
//
func (adjustment *Adjustment) SetPageIncrement(pageIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(pageIncrement)

	C.gtk_adjustment_set_page_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageIncrement)
}

// SetPageSize sets the page size of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the GtkAdjustment::changed signal when setting multiple adjustment
// properties.
//
// The function takes the following parameters:
//
//    - pageSize: new page size.
//
func (adjustment *Adjustment) SetPageSize(pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(pageSize)

	C.gtk_adjustment_set_page_size(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageSize)
}

// SetStepIncrement sets the step increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - stepIncrement: new step increment.
//
func (adjustment *Adjustment) SetStepIncrement(stepIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(stepIncrement)

	C.gtk_adjustment_set_step_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(stepIncrement)
}

// SetUpper sets the maximum value of the adjustment.
//
// Note that values will be restricted by upper - page-size if the page-size
// property is nonzero.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
//
// The function takes the following parameters:
//
//    - upper: new maximum value.
//
func (adjustment *Adjustment) SetUpper(upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	_arg1 = C.gdouble(upper)

	C.gtk_adjustment_set_upper(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(upper)
}
