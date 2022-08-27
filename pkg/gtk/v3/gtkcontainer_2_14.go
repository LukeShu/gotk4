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

// FocusChild returns the current focus child widget inside container. This is
// not the currently focused widget. That can be obtained by calling
// gtk_window_get_focus().
//
// The function returns the following values:
//
//    - widget (optional): child widget which will receive the focus inside
//      container when the container is focused, or NULL if none is set.
//
func (container *Container) FocusChild() Widgetter {
	var _arg0 *C.GtkContainer // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkContainer)(unsafe.Pointer(coreglib.InternObject(container).Native()))

	_cret = C.gtk_container_get_focus_child(_arg0)
	runtime.KeepAlive(container)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}
