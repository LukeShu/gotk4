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
	GTypeSwitchAccessible = coreglib.Type(C.gtk_switch_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSwitchAccessible, F: marshalSwitchAccessible},
	})
}

// SwitchAccessibleOverrider contains methods that are overridable.
type SwitchAccessibleOverrider interface {
}

type SwitchAccessible struct {
	_ [0]func() // equal guard
	WidgetAccessible

	atk.Action
}

var (
	_ coreglib.Objector = (*SwitchAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeSwitchAccessible,
		GoType:        reflect.TypeOf((*SwitchAccessible)(nil)),
		InitClass:     initClassSwitchAccessible,
		FinalizeClass: finalizeClassSwitchAccessible,
	})
}

func initClassSwitchAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitSwitchAccessible(*SwitchAccessibleClass) }); ok {
		klass := (*SwitchAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitSwitchAccessible(klass)
	}
}

func finalizeClassSwitchAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeSwitchAccessible(*SwitchAccessibleClass) }); ok {
		klass := (*SwitchAccessibleClass)(gextras.NewStructNative(gclass))
		goval.FinalizeSwitchAccessible(klass)
	}
}

func wrapSwitchAccessible(obj *coreglib.Object) *SwitchAccessible {
	return &SwitchAccessible{
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
		Action: atk.Action{
			Object: obj,
		},
	}
}

func marshalSwitchAccessible(p uintptr) (interface{}, error) {
	return wrapSwitchAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SwitchAccessibleClass: instance of this type is always passed by reference.
type SwitchAccessibleClass struct {
	*switchAccessibleClass
}

// switchAccessibleClass is the struct that's finalized.
type switchAccessibleClass struct {
	native *C.GtkSwitchAccessibleClass
}

func (s *SwitchAccessibleClass) ParentClass() *WidgetAccessibleClass {
	valptr := &s.native.parent_class
	var _v *WidgetAccessibleClass // out
	_v = (*WidgetAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
