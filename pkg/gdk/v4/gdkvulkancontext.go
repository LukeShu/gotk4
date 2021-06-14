// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_vulkan_context_get_type()), F: marshalVulkanContext},
	})
}

// VulkanContext: `GdkVulkanContext` is an object representing the
// platform-specific Vulkan draw context.
//
// `GdkVulkanContext`s are created for a surface using
// [method@Gdk.Surface.create_vulkan_context], and the context will match the
// the characteristics of the surface.
//
// Support for `GdkVulkanContext` is platform-specific and context creation can
// fail, returning nil context.
type VulkanContext interface {
	DrawContext
}

// vulkanContext implements the VulkanContext class.
type vulkanContext struct {
	DrawContext
}

var _ VulkanContext = (*vulkanContext)(nil)

// WrapVulkanContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapVulkanContext(obj *externglib.Object) VulkanContext {
	return vulkanContext{
		DrawContext: WrapDrawContext(obj),
	}
}

func marshalVulkanContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVulkanContext(obj), nil
}
