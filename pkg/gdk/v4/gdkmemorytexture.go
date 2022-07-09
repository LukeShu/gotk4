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

// GTypeMemoryFormat returns the GType for the type MemoryFormat.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMemoryFormat() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "MemoryFormat").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMemoryFormat)
	return gtype
}

// GTypeMemoryTexture returns the GType for the type MemoryTexture.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMemoryTexture() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "MemoryTexture").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMemoryTexture)
	return gtype
}

// MemoryFormat: GdkMemoryFormat describes a format that bytes can have in
// memory.
//
// It describes formats by listing the contents of the memory passed to it. So
// GDK_MEMORY_A8R8G8B8 will be 1 byte (8 bits) of alpha, followed by a byte each
// of red, green and blue. It is not endian-dependent, so CAIRO_FORMAT_ARGB32 is
// represented by different GdkMemoryFormats on architectures with different
// endiannesses.
//
// Its naming is modelled after VkFormat (see
// https://www.khronos.org/registry/vulkan/specs/1.0/html/vkspec.htmlFormat for
// details).
type MemoryFormat C.gint

const (
	// MemoryB8G8R8A8Premultiplied: 4 bytes; for blue, green, red, alpha. The
	// color values are premultiplied with the alpha value.
	MemoryB8G8R8A8Premultiplied MemoryFormat = iota
	// MemoryA8R8G8B8Premultiplied: 4 bytes; for alpha, red, green, blue. The
	// color values are premultiplied with the alpha value.
	MemoryA8R8G8B8Premultiplied
	// MemoryR8G8B8A8Premultiplied: 4 bytes; for red, green, blue, alpha The
	// color values are premultiplied with the alpha value.
	MemoryR8G8B8A8Premultiplied
	// MemoryB8G8R8A8: 4 bytes; for blue, green, red, alpha.
	MemoryB8G8R8A8
	// MemoryA8R8G8B8: 4 bytes; for alpha, red, green, blue.
	MemoryA8R8G8B8
	// MemoryR8G8B8A8: 4 bytes; for red, green, blue, alpha.
	MemoryR8G8B8A8
	// MemoryA8B8G8R8: 4 bytes; for alpha, blue, green, red.
	MemoryA8B8G8R8
	// MemoryR8G8B8: 3 bytes; for red, green, blue. The data is opaque.
	MemoryR8G8B8
	// MemoryB8G8R8: 3 bytes; for blue, green, red. The data is opaque.
	MemoryB8G8R8
	// MemoryNFormats: number of formats. This value will change as more formats
	// get added, so do not rely on its concrete integer.
	MemoryNFormats
)

func marshalMemoryFormat(p uintptr) (interface{}, error) {
	return MemoryFormat(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for MemoryFormat.
func (m MemoryFormat) String() string {
	switch m {
	case MemoryB8G8R8A8Premultiplied:
		return "B8G8R8A8Premultiplied"
	case MemoryA8R8G8B8Premultiplied:
		return "A8R8G8B8Premultiplied"
	case MemoryR8G8B8A8Premultiplied:
		return "R8G8B8A8Premultiplied"
	case MemoryB8G8R8A8:
		return "B8G8R8A8"
	case MemoryA8R8G8B8:
		return "A8R8G8B8"
	case MemoryR8G8B8A8:
		return "R8G8B8A8"
	case MemoryA8B8G8R8:
		return "A8B8G8R8"
	case MemoryR8G8B8:
		return "R8G8B8"
	case MemoryB8G8R8:
		return "B8G8R8"
	case MemoryNFormats:
		return "NFormats"
	default:
		return fmt.Sprintf("MemoryFormat(%d)", m)
	}
}

// MemoryTexture: GdkTexture representing image data in memory.
type MemoryTexture struct {
	_ [0]func() // equal guard
	Texture
}

var (
	_ Texturer = (*MemoryTexture)(nil)
)

func wrapMemoryTexture(obj *coreglib.Object) *MemoryTexture {
	return &MemoryTexture{
		Texture: Texture{
			Object: obj,
			Paintable: Paintable{
				Object: obj,
			},
		},
	}
}

func marshalMemoryTexture(p uintptr) (interface{}, error) {
	return wrapMemoryTexture(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
