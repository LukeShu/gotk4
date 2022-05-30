// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkdevicepad.go.
var (
	GTypeDevicePadFeature = coreglib.Type(C.gdk_device_pad_feature_get_type())
	GTypeDevicePad        = coreglib.Type(C.gdk_device_pad_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDevicePadFeature, F: marshalDevicePadFeature},
		{T: GTypeDevicePad, F: marshalDevicePad},
	})
}

// DevicePadFeature: pad feature.
type DevicePadFeature C.gint

const (
	// DevicePadFeatureButton: button.
	DevicePadFeatureButton DevicePadFeature = iota
	// DevicePadFeatureRing: ring-shaped interactive area.
	DevicePadFeatureRing
	// DevicePadFeatureStrip: straight interactive area.
	DevicePadFeatureStrip
)

func marshalDevicePadFeature(p uintptr) (interface{}, error) {
	return DevicePadFeature(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DevicePadFeature.
func (d DevicePadFeature) String() string {
	switch d {
	case DevicePadFeatureButton:
		return "Button"
	case DevicePadFeatureRing:
		return "Ring"
	case DevicePadFeatureStrip:
		return "Strip"
	default:
		return fmt.Sprintf("DevicePadFeature(%d)", d)
	}
}

// DevicePadOverrider contains methods that are overridable.
type DevicePadOverrider interface {
}

// DevicePad: GdkDevicePad is an interface implemented by devices of type
// GDK_SOURCE_TABLET_PAD
//
// It allows querying the features provided by the pad device.
//
// Tablet pads may contain one or more groups, each containing a subset of the
// buttons/rings/strips available. gdk.DevicePad.GetNGroups() can be used to
// obtain the number of groups, gdk.DevicePad.GetNFeatures() and
// gdk.DevicePad.GetFeatureGroup() can be combined to find out the number of
// buttons/rings/strips the device has, and how are they grouped.
//
// Each of those groups have different modes, which may be used to map each
// individual pad feature to multiple actions. Only one mode is effective
// (current) for each given group, different groups may have different current
// modes. The number of available modes in a group can be found out through
// gdk.DevicePad.GetGroupNModes(), and the current mode for a given group will
// be notified through events of type K_PAD_GROUP_MODE.
//
// DevicePad wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DevicePad struct {
	_ [0]func() // equal guard
	Device
}

var (
	_ Devicer = (*DevicePad)(nil)
)

// DevicePadder describes DevicePad's interface methods.
type DevicePadder interface {
	coreglib.Objector

	baseDevicePad() *DevicePad
}

var _ DevicePadder = (*DevicePad)(nil)

func ifaceInitDevicePadder(gifacePtr, data C.gpointer) {
}

func wrapDevicePad(obj *coreglib.Object) *DevicePad {
	return &DevicePad{
		Device: Device{
			Object: obj,
		},
	}
}

func marshalDevicePad(p uintptr) (interface{}, error) {
	return wrapDevicePad(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *DevicePad) baseDevicePad() *DevicePad {
	return v
}

// BaseDevicePad returns the underlying base object.
func BaseDevicePad(obj DevicePadder) *DevicePad {
	return obj.baseDevicePad()
}
