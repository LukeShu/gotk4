// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gdk4_VulkanContext_ConnectImagesUpdated(gpointer, guintptr);
import "C"

// glib.Type values for gdkvulkancontext.go.
var GTypeVulkanContext = coreglib.Type(C.gdk_vulkan_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeVulkanContext, F: marshalVulkanContext},
	})
}

// VulkanContext: GdkVulkanContext is an object representing the
// platform-specific Vulkan draw context.
//
// GdkVulkanContexts are created for a surface using
// gdk.Surface.CreateVulkanContext(), and the context will match the the
// characteristics of the surface.
//
// Support for GdkVulkanContext is platform-specific and context creation can
// fail, returning NULL context.
type VulkanContext struct {
	_ [0]func() // equal guard
	DrawContext

	*coreglib.Object
	gio.Initable
}

var (
	_ DrawContexter     = (*VulkanContext)(nil)
	_ coreglib.Objector = (*VulkanContext)(nil)
)

// VulkanContexter describes types inherited from class VulkanContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type VulkanContexter interface {
	coreglib.Objector
	baseVulkanContext() *VulkanContext
}

var _ VulkanContexter = (*VulkanContext)(nil)

func wrapVulkanContext(obj *coreglib.Object) *VulkanContext {
	return &VulkanContext{
		DrawContext: DrawContext{
			Object: obj,
		},
		Object: obj,
		Initable: gio.Initable{
			Object: obj,
		},
	}
}

func marshalVulkanContext(p uintptr) (interface{}, error) {
	return wrapVulkanContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *VulkanContext) baseVulkanContext() *VulkanContext {
	return v
}

// BaseVulkanContext returns the underlying base object.
func BaseVulkanContext(obj VulkanContexter) *VulkanContext {
	return obj.baseVulkanContext()
}

//export _gotk4_gdk4_VulkanContext_ConnectImagesUpdated
func _gotk4_gdk4_VulkanContext_ConnectImagesUpdated(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectImagesUpdated is emitted when the images managed by this context have
// changed.
//
// Usually this means that the swapchain had to be recreated, for example in
// response to a change of the surface size.
func (v *VulkanContext) ConnectImagesUpdated(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "images-updated", false, unsafe.Pointer(C._gotk4_gdk4_VulkanContext_ConnectImagesUpdated), f)
}
