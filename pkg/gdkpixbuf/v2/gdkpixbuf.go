// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	_ "runtime/cgo"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkpixbuf.go.
var GTypePixbufSimpleAnimIter = coreglib.Type(C.gdk_pixbuf_simple_anim_iter_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePixbufSimpleAnimIter, F: marshalPixbufSimpleAnimIter},
	})
}

// The function returns the following values:
//
func PixbufErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gdk_pixbuf_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type PixbufSimpleAnimIter struct {
	_ [0]func() // equal guard
	PixbufAnimationIter
}

var (
	_ coreglib.Objector = (*PixbufSimpleAnimIter)(nil)
)

func wrapPixbufSimpleAnimIter(obj *coreglib.Object) *PixbufSimpleAnimIter {
	return &PixbufSimpleAnimIter{
		PixbufAnimationIter: PixbufAnimationIter{
			Object: obj,
		},
	}
}

func marshalPixbufSimpleAnimIter(p uintptr) (interface{}, error) {
	return wrapPixbufSimpleAnimIter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
