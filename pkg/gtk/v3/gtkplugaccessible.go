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
		{T: externglib.Type(C.gtk_plug_accessible_get_type()), F: marshalPlugAccessibler},
	})
}

// PlugAccessibler describes PlugAccessible's methods.
type PlugAccessibler interface {
	ID() string
}

type PlugAccessible struct {
	WindowAccessible
}

var (
	_ PlugAccessibler = (*PlugAccessible)(nil)
	_ gextras.Nativer = (*PlugAccessible)(nil)
)

func wrapPlugAccessible(obj *externglib.Object) PlugAccessibler {
	return &PlugAccessible{
		WindowAccessible: WindowAccessible{
			ContainerAccessible: ContainerAccessible{
				WidgetAccessible: WidgetAccessible{
					Accessible: Accessible{
						ObjectClass: atk.ObjectClass{
							Object: obj,
						},
					},
					Component: atk.Component{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalPlugAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPlugAccessible(obj), nil
}

func (plug *PlugAccessible) ID() string {
	var _arg0 *C.GtkPlugAccessible // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkPlugAccessible)(unsafe.Pointer(plug.Native()))

	_cret = C.gtk_plug_accessible_get_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
