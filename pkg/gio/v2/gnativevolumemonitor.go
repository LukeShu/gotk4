// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeNativeVolumeMonitor = coreglib.Type(C.g_native_volume_monitor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNativeVolumeMonitor, F: marshalNativeVolumeMonitor},
	})
}

const NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME = "gio-native-volume-monitor"

// NativeVolumeMonitorOverrider contains methods that are overridable.
type NativeVolumeMonitorOverrider interface {
}

type NativeVolumeMonitor struct {
	_ [0]func() // equal guard
	VolumeMonitor
}

var (
	_ coreglib.Objector = (*NativeVolumeMonitor)(nil)
)

// NativeVolumeMonitorrer describes types inherited from class NativeVolumeMonitor.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type NativeVolumeMonitorrer interface {
	coreglib.Objector
	baseNativeVolumeMonitor() *NativeVolumeMonitor
}

var _ NativeVolumeMonitorrer = (*NativeVolumeMonitor)(nil)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeNativeVolumeMonitor,
		GoType:        reflect.TypeOf((*NativeVolumeMonitor)(nil)),
		InitClass:     initClassNativeVolumeMonitor,
		FinalizeClass: finalizeClassNativeVolumeMonitor,
	})
}

func initClassNativeVolumeMonitor(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitNativeVolumeMonitor(*NativeVolumeMonitorClass)
	}); ok {
		klass := (*NativeVolumeMonitorClass)(gextras.NewStructNative(gclass))
		goval.InitNativeVolumeMonitor(klass)
	}
}

func finalizeClassNativeVolumeMonitor(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		FinalizeNativeVolumeMonitor(*NativeVolumeMonitorClass)
	}); ok {
		klass := (*NativeVolumeMonitorClass)(gextras.NewStructNative(gclass))
		goval.FinalizeNativeVolumeMonitor(klass)
	}
}

func wrapNativeVolumeMonitor(obj *coreglib.Object) *NativeVolumeMonitor {
	return &NativeVolumeMonitor{
		VolumeMonitor: VolumeMonitor{
			Object: obj,
		},
	}
}

func marshalNativeVolumeMonitor(p uintptr) (interface{}, error) {
	return wrapNativeVolumeMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *NativeVolumeMonitor) baseNativeVolumeMonitor() *NativeVolumeMonitor {
	return v
}

// BaseNativeVolumeMonitor returns the underlying base object.
func BaseNativeVolumeMonitor(obj NativeVolumeMonitorrer) *NativeVolumeMonitor {
	return obj.baseNativeVolumeMonitor()
}

// NativeVolumeMonitorClass: instance of this type is always passed by
// reference.
type NativeVolumeMonitorClass struct {
	*nativeVolumeMonitorClass
}

// nativeVolumeMonitorClass is the struct that's finalized.
type nativeVolumeMonitorClass struct {
	native *C.GNativeVolumeMonitorClass
}

func (n *NativeVolumeMonitorClass) ParentClass() *VolumeMonitorClass {
	valptr := &n.native.parent_class
	var _v *VolumeMonitorClass // out
	_v = (*VolumeMonitorClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
