// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// GLContextClearCurrent clears the current GLContext.
//
// Any OpenGL call after this function returns will be ignored until
// gdk_gl_context_make_current() is called.
func GLContextClearCurrent() {
	C.gdk_gl_context_clear_current()
}

// GLContextGetCurrent retrieves the current GLContext.
//
// The function returns the following values:
//
//    - glContext (optional): current GLContext, or NULL.
//
func GLContextGetCurrent() GLContexter {
	var _cret *C.GdkGLContext // in

	_cret = C.gdk_gl_context_get_current()

	var _glContext GLContexter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(GLContexter)
				return ok
			})
			rv, ok := casted.(GLContexter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
			}
			_glContext = rv
		}
	}

	return _glContext
}
