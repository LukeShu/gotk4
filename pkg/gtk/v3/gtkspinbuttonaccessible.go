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
	GTypeSpinButtonAccessible = coreglib.Type(C.gtk_spin_button_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpinButtonAccessible, F: marshalSpinButtonAccessible},
	})
}

// SpinButtonAccessibleOverrider contains methods that are overridable.
type SpinButtonAccessibleOverrider interface {
}

type SpinButtonAccessible struct {
	_ [0]func() // equal guard
	EntryAccessible

	*coreglib.Object
	atk.Value
}

var (
	_ coreglib.Objector = (*SpinButtonAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeSpinButtonAccessible,
		GoType:        reflect.TypeOf((*SpinButtonAccessible)(nil)),
		InitClass:     initClassSpinButtonAccessible,
		FinalizeClass: finalizeClassSpinButtonAccessible,
	})
}

func initClassSpinButtonAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitSpinButtonAccessible(*SpinButtonAccessibleClass)
	}); ok {
		klass := (*SpinButtonAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitSpinButtonAccessible(klass)
	}
}

func finalizeClassSpinButtonAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		FinalizeSpinButtonAccessible(*SpinButtonAccessibleClass)
	}); ok {
		klass := (*SpinButtonAccessibleClass)(gextras.NewStructNative(gclass))
		goval.FinalizeSpinButtonAccessible(klass)
	}
}

func wrapSpinButtonAccessible(obj *coreglib.Object) *SpinButtonAccessible {
	return &SpinButtonAccessible{
		EntryAccessible: EntryAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
			Object: obj,
			Action: atk.Action{
				Object: obj,
			},
			EditableText: atk.EditableText{
				Object: obj,
			},
			Text: atk.Text{
				Object: obj,
			},
		},
		Object: obj,
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalSpinButtonAccessible(p uintptr) (interface{}, error) {
	return wrapSpinButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SpinButtonAccessibleClass: instance of this type is always passed by
// reference.
type SpinButtonAccessibleClass struct {
	*spinButtonAccessibleClass
}

// spinButtonAccessibleClass is the struct that's finalized.
type spinButtonAccessibleClass struct {
	native *C.GtkSpinButtonAccessibleClass
}

func (s *SpinButtonAccessibleClass) ParentClass() *EntryAccessibleClass {
	valptr := &s.native.parent_class
	var _v *EntryAccessibleClass // out
	_v = (*EntryAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
