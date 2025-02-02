// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_ShortcutFunc
func _gotk4_gtk4_ShortcutFunc(arg1 *C.GtkWidget, arg2 *C.GVariant, arg3 C.gpointer) (cret C.gboolean) {
	var fn ShortcutFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ShortcutFunc)
	}

	var _widget Widgetter   // out
	var _args *glib.Variant // out

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
	if arg2 != nil {
		_args = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
		C.g_variant_ref(arg2)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_args)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	ok := fn(_widget, _args)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
