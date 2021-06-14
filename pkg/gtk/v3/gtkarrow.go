// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_arrow_get_type()), F: marshalArrow},
	})
}

// Arrow: gtkArrow should be used to draw simple arrows that need to point in
// one of the four cardinal directions (up, down, left, or right). The style of
// the arrow can be one of shadow in, shadow out, etched in, or etched out. Note
// that these directions and style types may be amended in versions of GTK+ to
// come.
//
// GtkArrow will fill any space alloted to it, but since it is inherited from
// Misc, it can be padded and/or aligned, to fill exactly the space the
// programmer desires.
//
// Arrows are created with a call to gtk_arrow_new(). The direction or style of
// an arrow can be changed after creation by using gtk_arrow_set().
//
// GtkArrow has been deprecated; you can simply use a Image with a suitable icon
// name, such as “pan-down-symbolic“. When replacing GtkArrow by an image, pay
// attention to the fact that GtkArrow is doing automatic flipping between
// K_ARROW_LEFT and K_ARROW_RIGHT, depending on the text direction. To get the
// same effect with an image, use the icon names “pan-start-symbolic“ and
// “pan-end-symbolic“, which react to the text direction.
type Arrow interface {
	Misc
	Buildable

	// Set sets the direction and style of the Arrow, @arrow.
	Set(arrowType ArrowType, shadowType ShadowType)
}

// arrow implements the Arrow class.
type arrow struct {
	Misc
	Buildable
}

var _ Arrow = (*arrow)(nil)

// WrapArrow wraps a GObject to the right type. It is
// primarily used internally.
func WrapArrow(obj *externglib.Object) Arrow {
	return arrow{
		Misc:      WrapMisc(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalArrow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapArrow(obj), nil
}

// NewArrow constructs a class Arrow.
func NewArrow(arrowType ArrowType, shadowType ShadowType) Arrow {
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out

	_arg1 = (C.GtkArrowType)(arrowType)
	_arg2 = (C.GtkShadowType)(shadowType)

	var _cret C.GtkArrow // in

	_cret = C.gtk_arrow_new(_arg1, _arg2)

	var _arrow Arrow // out

	_arrow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Arrow)

	return _arrow
}

// Set sets the direction and style of the Arrow, @arrow.
func (a arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	var _arg0 *C.GtkArrow     // out
	var _arg1 C.GtkArrowType  // out
	var _arg2 C.GtkShadowType // out

	_arg0 = (*C.GtkArrow)(unsafe.Pointer(a.Native()))
	_arg1 = (C.GtkArrowType)(arrowType)
	_arg2 = (C.GtkShadowType)(shadowType)

	C.gtk_arrow_set(_arg0, _arg1, _arg2)
}
