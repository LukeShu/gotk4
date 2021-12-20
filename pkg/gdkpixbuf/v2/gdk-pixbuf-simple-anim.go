// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_get_type()), F: marshalPixbufSimpleAnimmer},
	})
}

// PixbufSimpleAnim: opaque struct representing a simple animation.
type PixbufSimpleAnim struct {
	_ [0]func() // equal guard
	PixbufAnimation
}

var (
	_ externglib.Objector = (*PixbufSimpleAnim)(nil)
)

func wrapPixbufSimpleAnim(obj *externglib.Object) *PixbufSimpleAnim {
	return &PixbufSimpleAnim{
		PixbufAnimation: PixbufAnimation{
			Object: obj,
		},
	}
}

func marshalPixbufSimpleAnimmer(p uintptr) (interface{}, error) {
	return wrapPixbufSimpleAnim(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPixbufSimpleAnim creates a new, empty animation.
//
// The function takes the following parameters:
//
//    - width of the animation.
//    - height of the animation.
//    - rate: speed of the animation, in frames per second.
//
// The function returns the following values:
//
//    - pixbufSimpleAnim: newly allocated PixbufSimpleAnim.
//
func NewPixbufSimpleAnim(width, height int, rate float32) *PixbufSimpleAnim {
	var _arg1 C.gint                 // out
	var _arg2 C.gint                 // out
	var _arg3 C.gfloat               // out
	var _cret *C.GdkPixbufSimpleAnim // in

	_arg1 = C.gint(width)
	_arg2 = C.gint(height)
	_arg3 = C.gfloat(rate)

	_cret = C.gdk_pixbuf_simple_anim_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(rate)

	var _pixbufSimpleAnim *PixbufSimpleAnim // out

	_pixbufSimpleAnim = wrapPixbufSimpleAnim(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pixbufSimpleAnim
}

// AddFrame adds a new frame to animation. The pixbuf must have the dimensions
// specified when the animation was constructed.
//
// The function takes the following parameters:
//
//    - pixbuf to add.
//
func (animation *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) {
	var _arg0 *C.GdkPixbufSimpleAnim // out
	var _arg1 *C.GdkPixbuf           // out

	_arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(animation.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gdk_pixbuf_simple_anim_add_frame(_arg0, _arg1)
	runtime.KeepAlive(animation)
	runtime.KeepAlive(pixbuf)
}

// Loop gets whether animation should loop indefinitely when it reaches the end.
//
// The function returns the following values:
//
//    - ok: TRUE if the animation loops forever, FALSE otherwise.
//
func (animation *PixbufSimpleAnim) Loop() bool {
	var _arg0 *C.GdkPixbufSimpleAnim // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(animation.Native()))

	_cret = C.gdk_pixbuf_simple_anim_get_loop(_arg0)
	runtime.KeepAlive(animation)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetLoop sets whether animation should loop indefinitely when it reaches the
// end.
//
// The function takes the following parameters:
//
//    - loop: whether to loop the animation.
//
func (animation *PixbufSimpleAnim) SetLoop(loop bool) {
	var _arg0 *C.GdkPixbufSimpleAnim // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(animation.Native()))
	if loop {
		_arg1 = C.TRUE
	}

	C.gdk_pixbuf_simple_anim_set_loop(_arg0, _arg1)
	runtime.KeepAlive(animation)
	runtime.KeepAlive(loop)
}
