// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_separator_get_type()), F: marshalSeparator},
	})
}

// Separator: `GtkSeparator` is a horizontal or vertical separator widget.
//
// !An example GtkSeparator (separators.png)
//
// A `GtkSeparator` can be used to group the widgets within a window. It
// displays a line with a shadow to make it appear sunken into the interface.
//
//
// CSS nodes
//
// `GtkSeparator` has a single CSS node with name separator. The node gets one
// of the .horizontal or .vertical style classes.
//
//
// Accessibility
//
// `GtkSeparator` uses the K_ACCESSIBLE_ROLE_SEPARATOR role.
type Separator interface {
	gextras.Objector

	privateSeparatorClass()
}

// SeparatorClass implements the Separator interface.
type SeparatorClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
	OrientableInterface
}

var _ Separator = (*SeparatorClass)(nil)

func wrapSeparator(obj *externglib.Object) Separator {
	return &SeparatorClass{
		Object: obj,
		WidgetClass: WidgetClass{
			Object:           obj,
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSeparator(obj), nil
}

// NewSeparator creates a new `GtkSeparator` with the given orientation.
func NewSeparator(orientation Orientation) Separator {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_separator_new(_arg1)

	var _separator Separator // out

	_separator = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Separator)

	return _separator
}

func (*SeparatorClass) privateSeparatorClass() {}
