// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeEmblem returns the GType for the type Emblem.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEmblem() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "Emblem").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalEmblem)
	return gtype
}

// Emblem is an implementation of #GIcon that supports having an emblem, which
// is an icon with additional properties. It can than be added to a Icon.
//
// Currently, only metainformation about the emblem's origin is supported. More
// may be added in the future.
type Emblem struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Icon
}

var (
	_ coreglib.Objector = (*Emblem)(nil)
)

func wrapEmblem(obj *coreglib.Object) *Emblem {
	return &Emblem{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblem(p uintptr) (interface{}, error) {
	return wrapEmblem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEmblem creates a new emblem for icon.
//
// The function takes the following parameters:
//
//    - icon: GIcon containing the icon.
//
// The function returns the following values:
//
//    - emblem: new #GEmblem.
//
func NewEmblem(icon Iconner) *Emblem {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_gret := girepository.MustFind("Gio", "Emblem").InvokeMethod("new_Emblem", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(icon)

	var _emblem *Emblem // out

	_emblem = wrapEmblem(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _emblem
}

// GetIcon gives back the icon from emblem.
//
// The function returns the following values:
//
//    - icon The returned object belongs to the emblem and should not be modified
//      or freed.
//
func (emblem *Emblem) GetIcon() *Icon {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(emblem).Native()))

	_gret := girepository.MustFind("Gio", "Emblem").InvokeMethod("get_icon", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(emblem)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.Take(unsafe.Pointer(_cret)))

	return _icon
}
