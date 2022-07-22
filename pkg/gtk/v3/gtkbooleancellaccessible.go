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
	GTypeBooleanCellAccessible = coreglib.Type(C.gtk_boolean_cell_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBooleanCellAccessible, F: marshalBooleanCellAccessible},
	})
}

// BooleanCellAccessibleOverrider contains methods that are overridable.
type BooleanCellAccessibleOverrider interface {
}

type BooleanCellAccessible struct {
	_ [0]func() // equal guard
	RendererCellAccessible
}

var (
	_ coreglib.Objector = (*BooleanCellAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeBooleanCellAccessible,
		GoType:        reflect.TypeOf((*BooleanCellAccessible)(nil)),
		InitClass:     initClassBooleanCellAccessible,
		FinalizeClass: finalizeClassBooleanCellAccessible,
	})
}

func initClassBooleanCellAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		InitBooleanCellAccessible(*BooleanCellAccessibleClass)
	}); ok {
		klass := (*BooleanCellAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitBooleanCellAccessible(klass)
	}
}

func finalizeClassBooleanCellAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface {
		FinalizeBooleanCellAccessible(*BooleanCellAccessibleClass)
	}); ok {
		klass := (*BooleanCellAccessibleClass)(gextras.NewStructNative(gclass))
		goval.FinalizeBooleanCellAccessible(klass)
	}
}

func wrapBooleanCellAccessible(obj *coreglib.Object) *BooleanCellAccessible {
	return &BooleanCellAccessible{
		RendererCellAccessible: RendererCellAccessible{
			CellAccessible: CellAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				AtkObject: atk.AtkObject{
					Object: obj,
				},
				Component: atk.Component{
					Object: obj,
				},
				TableCell: atk.TableCell{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalBooleanCellAccessible(p uintptr) (interface{}, error) {
	return wrapBooleanCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BooleanCellAccessibleClass: instance of this type is always passed by
// reference.
type BooleanCellAccessibleClass struct {
	*booleanCellAccessibleClass
}

// booleanCellAccessibleClass is the struct that's finalized.
type booleanCellAccessibleClass struct {
	native *C.GtkBooleanCellAccessibleClass
}

func (b *BooleanCellAccessibleClass) ParentClass() *RendererCellAccessibleClass {
	valptr := &b.native.parent_class
	var _v *RendererCellAccessibleClass // out
	_v = (*RendererCellAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
