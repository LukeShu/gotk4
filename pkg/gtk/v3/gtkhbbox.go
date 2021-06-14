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
		{T: externglib.Type(C.gtk_hbutton_box_get_type()), F: marshalHButtonBox},
	})
}

type HButtonBox interface {
	ButtonBox
	Buildable
	Orientable
}

// hButtonBox implements the HButtonBox class.
type hButtonBox struct {
	ButtonBox
	Buildable
	Orientable
}

var _ HButtonBox = (*hButtonBox)(nil)

// WrapHButtonBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapHButtonBox(obj *externglib.Object) HButtonBox {
	return hButtonBox{
		ButtonBox:  WrapButtonBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalHButtonBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHButtonBox(obj), nil
}

// NewHButtonBox constructs a class HButtonBox.
func NewHButtonBox() HButtonBox {
	var _cret C.GtkHButtonBox // in

	_cret = C.gtk_hbutton_box_new()

	var _hButtonBox HButtonBox // out

	_hButtonBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(HButtonBox)

	return _hButtonBox
}
