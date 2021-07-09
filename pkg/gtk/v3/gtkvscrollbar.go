// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_vscrollbar_get_type()), F: marshalVScrollbar},
	})
}

// VScrollbar: the VScrollbar widget is a widget arranged vertically creating a
// scrollbar. See Scrollbar for details on scrollbars. Adjustment pointers may
// be added to handle the adjustment of the scrollbar or it may be left nil in
// which case one will be created for you. See Scrollbar for a description of
// what the fields in an adjustment represent for a scrollbar.
//
// GtkVScrollbar has been deprecated, use Scrollbar instead.
type VScrollbar interface {
	gextras.Objector

	privateVScrollbarClass()
}

// VScrollbarClass implements the VScrollbar interface.
type VScrollbarClass struct {
	*externglib.Object
	ScrollbarClass
	BuildableInterface
	OrientableInterface
}

var _ VScrollbar = (*VScrollbarClass)(nil)

func wrapVScrollbar(obj *externglib.Object) VScrollbar {
	return &VScrollbarClass{
		Object: obj,
		ScrollbarClass: ScrollbarClass{
			Object: obj,
			RangeClass: RangeClass{
				Object: obj,
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				OrientableInterface: OrientableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			OrientableInterface: OrientableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalVScrollbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVScrollbar(obj), nil
}

// NewVScrollbar creates a new vertical scrollbar.
//
// Deprecated: since version 3.2.
func NewVScrollbar(adjustment Adjustment) *VScrollbarClass {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((&adjustment).Native()))

	_cret = C.gtk_vscrollbar_new(_arg1)

	var _vScrollbar *VScrollbarClass // out

	_vScrollbar = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*VScrollbarClass)

	return _vScrollbar
}

func (*VScrollbarClass) privateVScrollbarClass() {}
