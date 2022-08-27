// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// GFile* _gotk4_gio2_Volume_virtual_get_activation_root(void* fnptr, GVolume* arg0) {
//   return ((GFile* (*)(GVolume*))(fnptr))(arg0);
// };
import "C"

// ActivationRoot gets the activation root for a #GVolume if it is known ahead
// of mount time. Returns NULL otherwise. If not NULL and if volume is mounted,
// then the result of g_mount_get_root() on the #GMount object obtained from
// g_volume_get_mount() will always either be equal or a prefix of what this
// function returns. In other words, in code
//
//    (g_file_has_prefix (volume_activation_root, mount_root) ||
//     g_file_equal (volume_activation_root, mount_root))
//
// will always be TRUE.
//
// Activation roots are typically used in Monitor implementations to find the
// underlying mount to shadow, see g_mount_is_shadowed() for more details.
//
// The function returns the following values:
//
//    - file (optional): activation root of volume or NULL. Use g_object_unref()
//      to free.
//
func (volume *Volume) ActivationRoot() *File {
	var _arg0 *C.GVolume // out
	var _cret *C.GFile   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = C.g_volume_get_activation_root(_arg0)
	runtime.KeepAlive(volume)

	var _file *File // out

	if _cret != nil {
		_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _file
}

// activationRoot gets the activation root for a #GVolume if it is known ahead
// of mount time. Returns NULL otherwise. If not NULL and if volume is mounted,
// then the result of g_mount_get_root() on the #GMount object obtained from
// g_volume_get_mount() will always either be equal or a prefix of what this
// function returns. In other words, in code
//
//    (g_file_has_prefix (volume_activation_root, mount_root) ||
//     g_file_equal (volume_activation_root, mount_root))
//
// will always be TRUE.
//
// Activation roots are typically used in Monitor implementations to find the
// underlying mount to shadow, see g_mount_is_shadowed() for more details.
//
// The function returns the following values:
//
//    - file (optional): activation root of volume or NULL. Use g_object_unref()
//      to free.
//
func (volume *Volume) activationRoot() *File {
	gclass := (*C.GVolumeIface)(coreglib.PeekParentClass(volume))
	fnarg := gclass.get_activation_root

	var _arg0 *C.GVolume // out
	var _cret *C.GFile   // in

	_arg0 = (*C.GVolume)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = C._gotk4_gio2_Volume_virtual_get_activation_root(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(volume)

	var _file *File // out

	if _cret != nil {
		_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _file
}
