// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_EventControllerKey_ConnectIMUpdate
func _gotk4_gtk4_EventControllerKey_ConnectIMUpdate(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_EventControllerKey_ConnectKeyPressed
func _gotk4_gtk4_EventControllerKey_ConnectKeyPressed(arg0 C.gpointer, arg1 C.guint, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.guintptr) (cret C.gboolean) {
	var f func(keyval, keycode uint, state gdk.ModifierType) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(keyval, keycode uint, state gdk.ModifierType) (ok bool))
	}

	var _keyval uint            // out
	var _keycode uint           // out
	var _state gdk.ModifierType // out

	_keyval = uint(arg1)
	_keycode = uint(arg2)
	_state = gdk.ModifierType(arg3)

	ok := f(_keyval, _keycode, _state)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_EventControllerKey_ConnectKeyReleased
func _gotk4_gtk4_EventControllerKey_ConnectKeyReleased(arg0 C.gpointer, arg1 C.guint, arg2 C.guint, arg3 C.GdkModifierType, arg4 C.guintptr) {
	var f func(keyval, keycode uint, state gdk.ModifierType)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(keyval, keycode uint, state gdk.ModifierType))
	}

	var _keyval uint            // out
	var _keycode uint           // out
	var _state gdk.ModifierType // out

	_keyval = uint(arg1)
	_keycode = uint(arg2)
	_state = gdk.ModifierType(arg3)

	f(_keyval, _keycode, _state)
}

//export _gotk4_gtk4_EventControllerKey_ConnectModifiers
func _gotk4_gtk4_EventControllerKey_ConnectModifiers(arg0 C.gpointer, arg1 C.GdkModifierType, arg2 C.guintptr) (cret C.gboolean) {
	var f func(keyval gdk.ModifierType) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(keyval gdk.ModifierType) (ok bool))
	}

	var _keyval gdk.ModifierType // out

	_keyval = gdk.ModifierType(arg1)

	ok := f(_keyval)

	if ok {
		cret = C.TRUE
	}

	return cret
}
