// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_FlowBoxFilterFunc
func _gotk4_gtk3_FlowBoxFilterFunc(arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) (cret C.gboolean) {
	var fn FlowBoxFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FlowBoxFilterFunc)
	}

	var _child *FlowBoxChild // out

	_child = wrapFlowBoxChild(coreglib.Take(unsafe.Pointer(arg1)))

	ok := fn(_child)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_FlowBoxForEachFunc
func _gotk4_gtk3_FlowBoxForEachFunc(arg1 *C.GtkFlowBox, arg2 *C.GtkFlowBoxChild, arg3 C.gpointer) {
	var fn FlowBoxForEachFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FlowBoxForEachFunc)
	}

	var _box *FlowBox        // out
	var _child *FlowBoxChild // out

	_box = wrapFlowBox(coreglib.Take(unsafe.Pointer(arg1)))
	_child = wrapFlowBoxChild(coreglib.Take(unsafe.Pointer(arg2)))

	fn(_box, _child)
}

//export _gotk4_gtk3_FlowBoxSortFunc
func _gotk4_gtk3_FlowBoxSortFunc(arg1 *C.GtkFlowBoxChild, arg2 *C.GtkFlowBoxChild, arg3 C.gpointer) (cret C.gint) {
	var fn FlowBoxSortFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FlowBoxSortFunc)
	}

	var _child1 *FlowBoxChild // out
	var _child2 *FlowBoxChild // out

	_child1 = wrapFlowBoxChild(coreglib.Take(unsafe.Pointer(arg1)))
	_child2 = wrapFlowBoxChild(coreglib.Take(unsafe.Pointer(arg2)))

	gint := fn(_child1, _child2)

	cret = C.gint(gint)

	return cret
}
