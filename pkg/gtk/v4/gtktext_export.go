// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_Text_ConnectActivate
func _gotk4_gtk4_Text_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Text_ConnectBackspace
func _gotk4_gtk4_Text_ConnectBackspace(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Text_ConnectCopyClipboard
func _gotk4_gtk4_Text_ConnectCopyClipboard(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Text_ConnectCutClipboard
func _gotk4_gtk4_Text_ConnectCutClipboard(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Text_ConnectDeleteFromCursor
func _gotk4_gtk4_Text_ConnectDeleteFromCursor(arg0 C.gpointer, arg1 C.GtkDeleteType, arg2 C.gint, arg3 C.guintptr) {
	var f func(typ DeleteType, count int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(typ DeleteType, count int))
	}

	var _typ DeleteType // out
	var _count int      // out

	_typ = DeleteType(arg1)
	_count = int(arg2)

	f(_typ, _count)
}

//export _gotk4_gtk4_Text_ConnectInsertAtCursor
func _gotk4_gtk4_Text_ConnectInsertAtCursor(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(str string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(str string))
	}

	var _str string // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_str)
}

//export _gotk4_gtk4_Text_ConnectInsertEmoji
func _gotk4_gtk4_Text_ConnectInsertEmoji(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Text_ConnectMoveCursor
func _gotk4_gtk4_Text_ConnectMoveCursor(arg0 C.gpointer, arg1 C.GtkMovementStep, arg2 C.gint, arg3 C.gboolean, arg4 C.guintptr) {
	var f func(step MovementStep, count int, extend bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(step MovementStep, count int, extend bool))
	}

	var _step MovementStep // out
	var _count int         // out
	var _extend bool       // out

	_step = MovementStep(arg1)
	_count = int(arg2)
	if arg3 != 0 {
		_extend = true
	}

	f(_step, _count, _extend)
}

//export _gotk4_gtk4_Text_ConnectPasteClipboard
func _gotk4_gtk4_Text_ConnectPasteClipboard(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk4_Text_ConnectPreeditChanged
func _gotk4_gtk4_Text_ConnectPreeditChanged(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(preedit string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(preedit string))
	}

	var _preedit string // out

	_preedit = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_preedit)
}

//export _gotk4_gtk4_Text_ConnectToggleOverwrite
func _gotk4_gtk4_Text_ConnectToggleOverwrite(arg0 C.gpointer, arg1 C.guintptr) {
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
