// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeX11Visual returns the GType for the type X11Visual.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeX11Visual() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("GdkX11", "X11Visual").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalX11Visual)
	return gtype
}

type X11Visual struct {
	_ [0]func() // equal guard
	gdk.Visual
}

var (
	_ coreglib.Objector = (*X11Visual)(nil)
)

func wrapX11Visual(obj *coreglib.Object) *X11Visual {
	return &X11Visual{
		Visual: gdk.Visual{
			Object: obj,
		},
	}
}

func marshalX11Visual(p uintptr) (interface{}, error) {
	return wrapX11Visual(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
