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
	GTypeListBoxRowAccessible = coreglib.Type(C.gtk_list_box_row_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeListBoxRowAccessible, F: marshalListBoxRowAccessible},
	})
}

// ListBoxRowAccessibleOverrides contains methods that are overridable.
type ListBoxRowAccessibleOverrides struct {
}

func defaultListBoxRowAccessibleOverrides(v *ListBoxRowAccessible) ListBoxRowAccessibleOverrides {
	return ListBoxRowAccessibleOverrides{}
}

type ListBoxRowAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ coreglib.Objector = (*ListBoxRowAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ListBoxRowAccessible, *ListBoxRowAccessibleClass, ListBoxRowAccessibleOverrides](
		GTypeListBoxRowAccessible,
		initListBoxRowAccessibleClass,
		wrapListBoxRowAccessible,
		defaultListBoxRowAccessibleOverrides,
	)
}

func initListBoxRowAccessibleClass(gclass unsafe.Pointer, overrides ListBoxRowAccessibleOverrides, classInitFunc func(*ListBoxRowAccessibleClass)) {
	if classInitFunc != nil {
		class := (*ListBoxRowAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapListBoxRowAccessible(obj *coreglib.Object) *ListBoxRowAccessible {
	return &ListBoxRowAccessible{
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
	}
}

func marshalListBoxRowAccessible(p uintptr) (interface{}, error) {
	return wrapListBoxRowAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ListBoxRowAccessibleClass: instance of this type is always passed by
// reference.
type ListBoxRowAccessibleClass struct {
	*listBoxRowAccessibleClass
}

// listBoxRowAccessibleClass is the struct that's finalized.
type listBoxRowAccessibleClass struct {
	native *C.GtkListBoxRowAccessibleClass
}

func (l *ListBoxRowAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &l.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
