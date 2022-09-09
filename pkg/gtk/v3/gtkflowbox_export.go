// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_FlowBox_ConnectActivateCursorChild
func _gotk4_gtk3_FlowBox_ConnectActivateCursorChild(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk3_FlowBox_ConnectChildActivated
func _gotk4_gtk3_FlowBox_ConnectChildActivated(arg0 C.gpointer, arg1 *C.GtkFlowBoxChild, arg2 C.guintptr) {
	var f func(child *FlowBoxChild)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(child *FlowBoxChild))
	}

	var _child *FlowBoxChild // out

	_child = wrapFlowBoxChild(coreglib.Take(unsafe.Pointer(arg1)))

	f(_child)
}

//export _gotk4_gtk3_FlowBox_ConnectMoveCursor
func _gotk4_gtk3_FlowBox_ConnectMoveCursor(arg0 C.gpointer, arg1 C.GtkMovementStep, arg2 C.gint, arg3 C.guintptr) (cret C.gboolean) {
	var f func(step MovementStep, count int) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(step MovementStep, count int) (ok bool))
	}

	var _step MovementStep // out
	var _count int         // out

	_step = MovementStep(arg1)
	_count = int(arg2)

	ok := f(_step, _count)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_FlowBox_ConnectSelectAll
func _gotk4_gtk3_FlowBox_ConnectSelectAll(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk3_FlowBox_ConnectSelectedChildrenChanged
func _gotk4_gtk3_FlowBox_ConnectSelectedChildrenChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk3_FlowBox_ConnectToggleCursorChild
func _gotk4_gtk3_FlowBox_ConnectToggleCursorChild(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk3_FlowBox_ConnectUnselectAll
func _gotk4_gtk3_FlowBox_ConnectUnselectAll(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk3_FlowBoxChild_ConnectActivate
func _gotk4_gtk3_FlowBoxChild_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
