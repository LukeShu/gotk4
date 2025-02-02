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
	GTypeRadioButtonAccessible = coreglib.Type(C.gtk_radio_button_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRadioButtonAccessible, F: marshalRadioButtonAccessible},
	})
}

// RadioButtonAccessibleOverrides contains methods that are overridable.
type RadioButtonAccessibleOverrides struct {
}

func defaultRadioButtonAccessibleOverrides(v *RadioButtonAccessible) RadioButtonAccessibleOverrides {
	return RadioButtonAccessibleOverrides{}
}

type RadioButtonAccessible struct {
	_ [0]func() // equal guard
	ToggleButtonAccessible
}

var (
	_ coreglib.Objector = (*RadioButtonAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*RadioButtonAccessible, *RadioButtonAccessibleClass, RadioButtonAccessibleOverrides](
		GTypeRadioButtonAccessible,
		initRadioButtonAccessibleClass,
		wrapRadioButtonAccessible,
		defaultRadioButtonAccessibleOverrides,
	)
}

func initRadioButtonAccessibleClass(gclass unsafe.Pointer, overrides RadioButtonAccessibleOverrides, classInitFunc func(*RadioButtonAccessibleClass)) {
	if classInitFunc != nil {
		class := (*RadioButtonAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapRadioButtonAccessible(obj *coreglib.Object) *RadioButtonAccessible {
	return &RadioButtonAccessible{
		ToggleButtonAccessible: ToggleButtonAccessible{
			ButtonAccessible: ButtonAccessible{
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
				Object: obj,
				Action: atk.Action{
					Object: obj,
				},
				Image: atk.Image{
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioButtonAccessible(p uintptr) (interface{}, error) {
	return wrapRadioButtonAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RadioButtonAccessibleClass: instance of this type is always passed by
// reference.
type RadioButtonAccessibleClass struct {
	*radioButtonAccessibleClass
}

// radioButtonAccessibleClass is the struct that's finalized.
type radioButtonAccessibleClass struct {
	native *C.GtkRadioButtonAccessibleClass
}

func (r *RadioButtonAccessibleClass) ParentClass() *ToggleButtonAccessibleClass {
	valptr := &r.native.parent_class
	var _v *ToggleButtonAccessibleClass // out
	_v = (*ToggleButtonAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
