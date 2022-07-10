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
// #include <glib-object.h>
import "C"

// NewSimplePermission creates a new #GPermission instance that represents an
// action that is either always or never allowed.
//
// The function takes the following parameters:
//
//    - allowed: TRUE if the action is allowed.
//
// The function returns the following values:
//
//    - simplePermission as a #GPermission.
//
func NewSimplePermission(allowed bool) *SimplePermission {
	var _args [1]girepository.Argument

	if allowed {
		*(*C.gboolean)(unsafe.Pointer(&_args[0])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "SimplePermission")
	_gret := _info.InvokeClassMethod("new_SimplePermission", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(allowed)

	var _simplePermission *SimplePermission // out

	_simplePermission = wrapSimplePermission(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _simplePermission
}
