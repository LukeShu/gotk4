// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_vbutton_box_get_type()), F: marshalVButtonBoxer},
	})
}

// VButtonBoxer describes VButtonBox's methods.
type VButtonBoxer interface {
	privateVButtonBox()
}

type VButtonBox struct {
	ButtonBox
}

var (
	_ VButtonBoxer    = (*VButtonBox)(nil)
	_ gextras.Nativer = (*VButtonBox)(nil)
)

func wrapVButtonBox(obj *externglib.Object) VButtonBoxer {
	return &VButtonBox{
		ButtonBox: ButtonBox{
			Box: Box{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalVButtonBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVButtonBox(obj), nil
}

// NewVButtonBox creates a new vertical button box.
//
// Deprecated: Use gtk_button_box_new() with GTK_ORIENTATION_VERTICAL instead.
func NewVButtonBox() *VButtonBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vbutton_box_new()

	var _vButtonBox *VButtonBox // out

	_vButtonBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*VButtonBox)

	return _vButtonBox
}

func (*VButtonBox) privateVButtonBox() {}
