// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeSeparator = coreglib.Type(C.gtk_separator_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSeparator, F: marshalSeparator},
	})
}

// SeparatorOverrides contains methods that are overridable.
type SeparatorOverrides struct {
}

func defaultSeparatorOverrides(v *Separator) SeparatorOverrides {
	return SeparatorOverrides{}
}

// Separator is a horizontal or vertical separator widget, depending on the
// value of the Orientable:orientation property, used to group the widgets
// within a window. It displays a line with a shadow to make it appear sunken
// into the interface.
//
//
// CSS nodes
//
// GtkSeparator has a single CSS node with name separator. The node gets one of
// the .horizontal or .vertical style classes.
type Separator struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*Separator)(nil)
	_ coreglib.Objector = (*Separator)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Separator, *SeparatorClass, SeparatorOverrides](
		GTypeSeparator,
		initSeparatorClass,
		wrapSeparator,
		defaultSeparatorOverrides,
	)
}

func initSeparatorClass(gclass unsafe.Pointer, overrides SeparatorOverrides, classInitFunc func(*SeparatorClass)) {
	if classInitFunc != nil {
		class := (*SeparatorClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSeparator(obj *coreglib.Object) *Separator {
	return &Separator{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalSeparator(p uintptr) (interface{}, error) {
	return wrapSeparator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SeparatorClass: instance of this type is always passed by reference.
type SeparatorClass struct {
	*separatorClass
}

// separatorClass is the struct that's finalized.
type separatorClass struct {
	native *C.GtkSeparatorClass
}

func (s *SeparatorClass) ParentClass() *WidgetClass {
	valptr := &s.native.parent_class
	var _v *WidgetClass // out
	_v = (*WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
