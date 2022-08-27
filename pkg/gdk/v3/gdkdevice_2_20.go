// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// AxisUse returns the axis use for index_.
//
// The function takes the following parameters:
//
//    - index_: index of the axis.
//
// The function returns the following values:
//
//    - axisUse specifying how the axis is used.
//
func (device *Device) AxisUse(index_ uint) AxisUse {
	var _arg0 *C.GdkDevice // out
	var _arg1 C.guint      // out
	var _cret C.GdkAxisUse // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_axis_use(_arg0, _arg1)
	runtime.KeepAlive(device)
	runtime.KeepAlive(index_)

	var _axisUse AxisUse // out

	_axisUse = AxisUse(_cret)

	return _axisUse
}

// HasCursor determines whether the pointer follows device motion. This is not
// meaningful for keyboard devices, which don't have a pointer.
//
// The function returns the following values:
//
//    - ok: TRUE if the pointer follows device motion.
//
func (device *Device) HasCursor() bool {
	var _arg0 *C.GdkDevice // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_device_get_has_cursor(_arg0)
	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Key: if index_ has a valid keyval, this function will return TRUE and fill in
// keyval and modifiers with the keyval settings.
//
// The function takes the following parameters:
//
//    - index_: index of the macro button to get.
//
// The function returns the following values:
//
//    - keyval: return value for the keyval.
//    - modifiers: return value for modifiers.
//    - ok: TRUE if keyval is set for index.
//
func (device *Device) Key(index_ uint) (uint, ModifierType, bool) {
	var _arg0 *C.GdkDevice      // out
	var _arg1 C.guint           // out
	var _arg2 C.guint           // in
	var _arg3 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	_arg1 = C.guint(index_)

	_cret = C.gdk_device_get_key(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(device)
	runtime.KeepAlive(index_)

	var _keyval uint            // out
	var _modifiers ModifierType // out
	var _ok bool                // out

	_keyval = uint(_arg2)
	_modifiers = ModifierType(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _modifiers, _ok
}

// Mode determines the mode of the device.
//
// The function returns the following values:
//
//    - inputMode: InputSource.
//
func (device *Device) Mode() InputMode {
	var _arg0 *C.GdkDevice   // out
	var _cret C.GdkInputMode // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_device_get_mode(_arg0)
	runtime.KeepAlive(device)

	var _inputMode InputMode // out

	_inputMode = InputMode(_cret)

	return _inputMode
}

// Name determines the name of the device.
//
// The function returns the following values:
//
//    - utf8: name.
//
func (device *Device) Name() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_device_get_name(_arg0)
	runtime.KeepAlive(device)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Source determines the type of the device.
//
// The function returns the following values:
//
//    - inputSource: InputSource.
//
func (device *Device) Source() InputSource {
	var _arg0 *C.GdkDevice     // out
	var _cret C.GdkInputSource // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_device_get_source(_arg0)
	runtime.KeepAlive(device)

	var _inputSource InputSource // out

	_inputSource = InputSource(_cret)

	return _inputSource
}
