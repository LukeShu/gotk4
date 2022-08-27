// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// PlugWindow retrieves the window of the plug. Use this to check if the plug
// has been created inside of the socket.
//
// The function returns the following values:
//
//    - window (optional) of the plug if available, or NULL.
//
func (socket_ *Socket) PlugWindow() gdk.Windower {
	var _arg0 *C.GtkSocket // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkSocket)(unsafe.Pointer(coreglib.InternObject(socket_).Native()))

	_cret = C.gtk_socket_get_plug_window(_arg0)
	runtime.KeepAlive(socket_)

	var _window gdk.Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Windower)
				return ok
			})
			rv, ok := casted.(gdk.Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}
