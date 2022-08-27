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

// UseES retrieves the value set by gtk_gl_area_set_use_es().
//
// The function returns the following values:
//
//    - ok: TRUE if the GLArea should create an OpenGL ES context and FALSE
//      otherwise.
//
func (area *GLArea) UseES() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(coreglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_use_es(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetUseES sets whether the area should create an OpenGL or an OpenGL ES
// context.
//
// You should check the capabilities of the GLContext before drawing with either
// API.
//
// The function takes the following parameters:
//
//    - useEs: whether to use OpenGL or OpenGL ES.
//
func (area *GLArea) SetUseES(useEs bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(coreglib.InternObject(area).Native()))
	if useEs {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_use_es(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(useEs)
}
