// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_ClipboardURIReceivedFunc
func _gotk4_gtk3_ClipboardURIReceivedFunc(arg1 *C.GtkClipboard, arg2 **C.gchar, arg3 C.gpointer) {
	var fn ClipboardURIReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardURIReceivedFunc)
	}

	var _clipboard *Clipboard // out
	var _uris []string        // out

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	{
		var i int
		var z *C.gchar
		for p := arg2; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(arg2, i)
		_uris = make([]string, i)
		for i := range src {
			_uris[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	fn(_clipboard, _uris)
}
