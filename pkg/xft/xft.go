// Code generated by girgen. DO NOT EDIT.

package xft

import (
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{

		// Objects/Classes
	})
}

func Init()

type Color struct {
	native *C.XftColor
}

func wrapColor(p *C.XftColor) *Color {
	v := Color{native: p}

	runtime.SetFinalizer(v, nil)
	runtime.SetFinalizer(v, (*Color).free)

	return &v
}

func marshalColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.XftColor)(unsafe.Pointer(b))

	return wrapColor(c)
}

func (c *Color) free() {}

// Native returns the pointer to *C.XftColor. The caller is expected to
// cast.
func (c *Color) Native() unsafe.Pointer {
	return unsafe.Pointer(c.native)
}

type Draw struct {
	native *C.XftDraw
}

func wrapDraw(p *C.XftDraw) *Draw {
	v := Draw{native: p}

	runtime.SetFinalizer(v, nil)
	runtime.SetFinalizer(v, (*Draw).free)

	return &v
}

func marshalDraw(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.XftDraw)(unsafe.Pointer(b))

	return wrapDraw(c)
}

func (d *Draw) free() {}

// Native returns the pointer to *C.XftDraw. The caller is expected to
// cast.
func (d *Draw) Native() unsafe.Pointer {
	return unsafe.Pointer(d.native)
}

type Font struct {
	native *C.XftFont
}

func wrapFont(p *C.XftFont) *Font {
	v := Font{native: p}

	runtime.SetFinalizer(v, nil)
	runtime.SetFinalizer(v, (*Font).free)

	return &v
}

func marshalFont(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.XftFont)(unsafe.Pointer(b))

	return wrapFont(c)
}

func (f *Font) free() {}

// Native returns the pointer to *C.XftFont. The caller is expected to
// cast.
func (f *Font) Native() unsafe.Pointer {
	return unsafe.Pointer(f.native)
}

type GlyphSpec struct {
	native *C.XftGlyphSpec
}

func wrapGlyphSpec(p *C.XftGlyphSpec) *GlyphSpec {
	v := GlyphSpec{native: p}

	runtime.SetFinalizer(v, nil)
	runtime.SetFinalizer(v, (*GlyphSpec).free)

	return &v
}

func marshalGlyphSpec(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.XftGlyphSpec)(unsafe.Pointer(b))

	return wrapGlyphSpec(c)
}

func (g *GlyphSpec) free() {}

// Native returns the pointer to *C.XftGlyphSpec. The caller is expected to
// cast.
func (g *GlyphSpec) Native() unsafe.Pointer {
	return unsafe.Pointer(g.native)
}
