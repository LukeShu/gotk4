// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_TickCallback
func _gotk4_gtk3_TickCallback(arg1 *C.GtkWidget, arg2 *C.GdkFrameClock, arg3 C.gpointer) (cret C.gboolean) {
	var fn TickCallback
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TickCallback)
	}

	var _widget Widgetter            // out
	var _frameClock gdk.FrameClocker // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gdk.FrameClocker is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.FrameClocker)
			return ok
		})
		rv, ok := casted.(gdk.FrameClocker)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.FrameClocker")
		}
		_frameClock = rv
	}

	ok := fn(_widget, _frameClock)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
