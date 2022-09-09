// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gio2_SettingsBindGetMapping
func _gotk4_gio2_SettingsBindGetMapping(arg1 *C.GValue, arg2 *C.GVariant, arg3 C.gpointer) (cret C.gboolean) {
	var fn SettingsBindGetMapping
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(SettingsBindGetMapping)
	}

	var _value *coreglib.Value // out
	var _variant *glib.Variant // out

	_value = coreglib.ValueFromNative(unsafe.Pointer(arg1))
	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	C.g_variant_ref(arg2)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	ok := fn(_value, _variant)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_SettingsBindSetMapping
func _gotk4_gio2_SettingsBindSetMapping(arg1 *C.GValue, arg2 *C.GVariantType, arg3 C.gpointer) (cret *C.GVariant) {
	var fn SettingsBindSetMapping
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(SettingsBindSetMapping)
	}

	var _value *coreglib.Value          // out
	var _expectedType *glib.VariantType // out

	_value = coreglib.ValueFromNative(unsafe.Pointer(arg1))
	_expectedType = (*glib.VariantType)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	variant := fn(_value, _expectedType)

	var _ *glib.Variant

	cret = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	return cret
}

//export _gotk4_gio2_SettingsGetMapping
func _gotk4_gio2_SettingsGetMapping(arg1 *C.GVariant, arg2 *C.gpointer, arg3 C.gpointer) (cret C.gboolean) {
	var fn SettingsGetMapping
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(SettingsGetMapping)
	}

	var _value *glib.Variant // out

	_value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	C.g_variant_ref(arg1)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_value)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	result, ok := fn(_value)

	var _ unsafe.Pointer
	var _ bool

	*arg2 = (C.gpointer)(unsafe.Pointer(result))
	if ok {
		cret = C.TRUE
	}

	return cret
}
