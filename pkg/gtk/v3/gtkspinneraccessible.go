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
	GTypeSpinnerAccessible = coreglib.Type(C.gtk_spinner_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpinnerAccessible, F: marshalSpinnerAccessible},
	})
}

// SpinnerAccessibleOverrider contains methods that are overridable.
type SpinnerAccessibleOverrider interface {
}

type SpinnerAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Image
}

var (
	_ coreglib.Objector = (*SpinnerAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeSpinnerAccessible,
		GoType:        reflect.TypeOf((*SpinnerAccessible)(nil)),
		InitClass:     initClassSpinnerAccessible,
		FinalizeClass: finalizeClassSpinnerAccessible,
	})
}

func initClassSpinnerAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitSpinnerAccessible(*SpinnerAccessibleClass) }); ok {
		klass := (*SpinnerAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitSpinnerAccessible(klass)
	}
}

func finalizeClassSpinnerAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeSpinnerAccessible(*SpinnerAccessibleClass) }); ok {
		klass := (*SpinnerAccessibleClass)(gextras.NewStructNative(gclass))
		goval.FinalizeSpinnerAccessible(klass)
	}
}

func wrapSpinnerAccessible(obj *coreglib.Object) *SpinnerAccessible {
	return &SpinnerAccessible{
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
		Image: atk.Image{
			Object: obj,
		},
	}
}

func marshalSpinnerAccessible(p uintptr) (interface{}, error) {
	return wrapSpinnerAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SpinnerAccessibleClass: instance of this type is always passed by reference.
type SpinnerAccessibleClass struct {
	*spinnerAccessibleClass
}

// spinnerAccessibleClass is the struct that's finalized.
type spinnerAccessibleClass struct {
	native *C.GtkSpinnerAccessibleClass
}

func (s *SpinnerAccessibleClass) ParentClass() *WidgetAccessibleClass {
	valptr := &s.native.parent_class
	var _v *WidgetAccessibleClass // out
	_v = (*WidgetAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
