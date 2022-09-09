// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypePanedAccessible = coreglib.Type(C.gtk_paned_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePanedAccessible, F: marshalPanedAccessible},
	})
}

// PanedAccessibleOverrides contains methods that are overridable.
type PanedAccessibleOverrides struct {
}

func defaultPanedAccessibleOverrides(v *PanedAccessible) PanedAccessibleOverrides {
	return PanedAccessibleOverrides{}
}

type PanedAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Value
}

var (
	_ coreglib.Objector = (*PanedAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*PanedAccessible, *PanedAccessibleClass, PanedAccessibleOverrides](
		GTypePanedAccessible,
		initPanedAccessibleClass,
		wrapPanedAccessible,
		defaultPanedAccessibleOverrides,
	)
}

func initPanedAccessibleClass(gclass unsafe.Pointer, overrides PanedAccessibleOverrides, classInitFunc func(*PanedAccessibleClass)) {
	if classInitFunc != nil {
		class := (*PanedAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPanedAccessible(obj *coreglib.Object) *PanedAccessible {
	return &PanedAccessible{
		ContainerAccessible: ContainerAccessible{
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
		},
		Value: atk.Value{
			Object: obj,
		},
	}
}

func marshalPanedAccessible(p uintptr) (interface{}, error) {
	return wrapPanedAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// PanedAccessibleClass: instance of this type is always passed by reference.
type PanedAccessibleClass struct {
	*panedAccessibleClass
}

// panedAccessibleClass is the struct that's finalized.
type panedAccessibleClass struct {
	native *C.GtkPanedAccessibleClass
}

func (p *PanedAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &p.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
