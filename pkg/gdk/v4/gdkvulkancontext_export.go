// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

//export _gotk4_gdk4_VulkanContext_ConnectImagesUpdated
func _gotk4_gdk4_VulkanContext_ConnectImagesUpdated(arg0 C.gpointer, arg1 C.guintptr) {
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
