// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// GIcon* _gotk4_gio2_Mount_virtual_get_symbolic_icon(void* fnptr, GMount* arg0) {
//   return ((GIcon* (*)(GMount*))(fnptr))(arg0);
// };
import "C"

// SymbolicIcon gets the symbolic icon for mount.
//
// The function returns the following values:
//
//    - icon: #GIcon. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) SymbolicIcon() *Icon {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C.g_mount_get_symbolic_icon(_arg0)
	runtime.KeepAlive(mount)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}

// symbolicIcon gets the symbolic icon for mount.
//
// The function returns the following values:
//
//    - icon: #GIcon. The returned object should be unreffed with
//      g_object_unref() when no longer needed.
//
func (mount *Mount) symbolicIcon() *Icon {
	gclass := (*C.GMountIface)(coreglib.PeekParentClass(mount))
	fnarg := gclass.get_symbolic_icon

	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(coreglib.InternObject(mount).Native()))

	_cret = C._gotk4_gio2_Mount_virtual_get_symbolic_icon(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(mount)

	var _icon *Icon // out

	_icon = wrapIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _icon
}
