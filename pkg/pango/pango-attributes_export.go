// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

//export _gotk4_pango1_AttrDataCopyFunc
func _gotk4_pango1_AttrDataCopyFunc(arg1 C.gconstpointer) (cret C.gpointer) {
	var fn AttrDataCopyFunc
	{
		v := gbox.Get(uintptr(arg1))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AttrDataCopyFunc)
	}

	gpointer := fn()

	var _ unsafe.Pointer

	cret = (C.gpointer)(unsafe.Pointer(gpointer))

	return cret
}

//export _gotk4_pango1_AttrFilterFunc
func _gotk4_pango1_AttrFilterFunc(arg1 *C.PangoAttribute, arg2 C.gpointer) (cret C.gboolean) {
	var fn AttrFilterFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AttrFilterFunc)
	}

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := fn(_attribute)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
