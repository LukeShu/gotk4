// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
import "C"

//export _gotk4_atk1_Component_ConnectBoundsChanged
func _gotk4_atk1_Component_ConnectBoundsChanged(arg0 C.gpointer, arg1 *C.AtkRectangle, arg2 C.guintptr) {
	var f func(arg1 *Rectangle)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 *Rectangle))
	}

	var _arg1 *Rectangle // out

	_arg1 = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_arg1)
}
