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
// extern void _gotk4_gtk3_CellAccessibleClass_update_cache(GtkCellAccessible*, gboolean);
import "C"

// GType values.
var (
	GTypeCellAccessible = coreglib.Type(C.gtk_cell_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellAccessible, F: marshalCellAccessible},
	})
}

// CellAccessibleOverrider contains methods that are overridable.
type CellAccessibleOverrider interface {
	// The function takes the following parameters:
	//
	UpdateCache(emitSignal bool)
}

type CellAccessible struct {
	_ [0]func() // equal guard
	Accessible

	*coreglib.Object
	atk.Action
	atk.AtkObject
	atk.Component
	atk.TableCell
}

var (
	_ coreglib.Objector = (*CellAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeCellAccessible,
		GoType:        reflect.TypeOf((*CellAccessible)(nil)),
		InitClass:     initClassCellAccessible,
		FinalizeClass: finalizeClassCellAccessible,
	})
}

func initClassCellAccessible(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GtkCellAccessibleClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface{ UpdateCache(emitSignal bool) }); ok {
		pclass.update_cache = (*[0]byte)(C._gotk4_gtk3_CellAccessibleClass_update_cache)
	}
	if goval, ok := goval.(interface{ InitCellAccessible(*CellAccessibleClass) }); ok {
		klass := (*CellAccessibleClass)(gextras.NewStructNative(gclass))
		goval.InitCellAccessible(klass)
	}
}

func finalizeClassCellAccessible(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeCellAccessible(*CellAccessibleClass) }); ok {
		klass := (*CellAccessibleClass)(gextras.NewStructNative(gclass))
		goval.FinalizeCellAccessible(klass)
	}
}

//export _gotk4_gtk3_CellAccessibleClass_update_cache
func _gotk4_gtk3_CellAccessibleClass_update_cache(arg0 *C.GtkCellAccessible, arg1 C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ UpdateCache(emitSignal bool) })

	var _emitSignal bool // out

	if arg1 != 0 {
		_emitSignal = true
	}

	iface.UpdateCache(_emitSignal)
}

func wrapCellAccessible(obj *coreglib.Object) *CellAccessible {
	return &CellAccessible{
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
	}
}

func marshalCellAccessible(p uintptr) (interface{}, error) {
	return wrapCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CellAccessibleClass: instance of this type is always passed by reference.
type CellAccessibleClass struct {
	*cellAccessibleClass
}

// cellAccessibleClass is the struct that's finalized.
type cellAccessibleClass struct {
	native *C.GtkCellAccessibleClass
}

func (c *CellAccessibleClass) ParentClass() *AccessibleClass {
	valptr := &c.native.parent_class
	var _v *AccessibleClass // out
	_v = (*AccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
