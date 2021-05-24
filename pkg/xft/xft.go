// Code generated by girgen. DO NOT EDIT.

package xft

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
//
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{

		// Records
		// Skipped Color.
		// Skipped Draw.
		// Skipped Font.
		// Skipped GlyphSpec.

	})
}

func Init() {
	C.XftInit()
}

type Color struct {
	native *C.XftColor
}

// WrapColor wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColor(ptr unsafe.Pointer) *Color {
	p := (*C.XftColor)(ptr)
	v := Color{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Color).free)

	return &v
}

func marshalColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapColor(unsafe.Pointer(b))
}

func (c *Color) free() {
	C.free(unsafe.Pointer(c.native))
}

// Native returns the pointer to *C.XftColor. The caller is expected to
// cast.
func (c *Color) Native() unsafe.Pointer {
	return unsafe.Pointer(c.native)
}

type Draw struct {
	native *C.XftDraw
}

// WrapDraw wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDraw(ptr unsafe.Pointer) *Draw {
	p := (*C.XftDraw)(ptr)
	v := Draw{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Draw).free)

	return &v
}

func marshalDraw(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDraw(unsafe.Pointer(b))
}

func (d *Draw) free() {
	C.free(unsafe.Pointer(d.native))
}

// Native returns the pointer to *C.XftDraw. The caller is expected to
// cast.
func (d *Draw) Native() unsafe.Pointer {
	return unsafe.Pointer(d.native)
}

type Font struct {
	native *C.XftFont
}

// WrapFont wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFont(ptr unsafe.Pointer) *Font {
	p := (*C.XftFont)(ptr)
	v := Font{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Font).free)

	return &v
}

func marshalFont(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFont(unsafe.Pointer(b))
}

func (f *Font) free() {
	C.free(unsafe.Pointer(f.native))
}

// Native returns the pointer to *C.XftFont. The caller is expected to
// cast.
func (f *Font) Native() unsafe.Pointer {
	return unsafe.Pointer(f.native)
}

type GlyphSpec struct {
	native *C.XftGlyphSpec
}

// WrapGlyphSpec wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGlyphSpec(ptr unsafe.Pointer) *GlyphSpec {
	p := (*C.XftGlyphSpec)(ptr)
	v := GlyphSpec{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*GlyphSpec).free)

	return &v
}

func marshalGlyphSpec(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGlyphSpec(unsafe.Pointer(b))
}

func (g *GlyphSpec) free() {
	C.free(unsafe.Pointer(g.native))
}

// Native returns the pointer to *C.XftGlyphSpec. The caller is expected to
// cast.
func (g *GlyphSpec) Native() unsafe.Pointer {
	return unsafe.Pointer(g.native)
}
