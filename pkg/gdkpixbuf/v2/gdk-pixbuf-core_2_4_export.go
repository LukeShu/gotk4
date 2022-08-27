// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

//export _gotk4_gdkpixbuf2_PixbufSaveFunc
func _gotk4_gdkpixbuf2_PixbufSaveFunc(arg1 *C.gchar, arg2 C.gsize, arg3 **C.GError, arg4 C.gpointer) (cret C.gboolean) {
	var fn PixbufSaveFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PixbufSaveFunc)
	}

	var _buf []byte // out

	_buf = make([]byte, arg2)
	copy(_buf, unsafe.Slice((*byte)(unsafe.Pointer(arg1)), arg2))

	err, ok := fn(_buf)

	if err != nil && arg3 != nil {
		*arg3 = (*C.GError)(gerror.New(err))
	}
	if ok {
		cret = C.TRUE
	}

	return cret
}
