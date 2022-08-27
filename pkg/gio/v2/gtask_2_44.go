// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// Completed gets the value of #GTask:completed. This changes from FALSE to TRUE
// after the task’s callback is invoked, and will return FALSE if called from
// inside the callback.
//
// The function returns the following values:
//
//    - ok: TRUE if the task has completed, FALSE otherwise.
//
func (task *Task) Completed() bool {
	var _arg0 *C.GTask   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GTask)(unsafe.Pointer(coreglib.InternObject(task).Native()))

	_cret = C.g_task_get_completed(_arg0)
	runtime.KeepAlive(task)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
