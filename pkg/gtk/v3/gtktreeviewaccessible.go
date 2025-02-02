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
	GTypeTreeViewAccessible = coreglib.Type(C.gtk_tree_view_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTreeViewAccessible, F: marshalTreeViewAccessible},
	})
}

// TreeViewAccessibleOverrides contains methods that are overridable.
type TreeViewAccessibleOverrides struct {
}

func defaultTreeViewAccessibleOverrides(v *TreeViewAccessible) TreeViewAccessibleOverrides {
	return TreeViewAccessibleOverrides{}
}

type TreeViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*coreglib.Object
	atk.Selection
	atk.Table
	CellAccessibleParent
}

var (
	_ coreglib.Objector = (*TreeViewAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TreeViewAccessible, *TreeViewAccessibleClass, TreeViewAccessibleOverrides](
		GTypeTreeViewAccessible,
		initTreeViewAccessibleClass,
		wrapTreeViewAccessible,
		defaultTreeViewAccessibleOverrides,
	)
}

func initTreeViewAccessibleClass(gclass unsafe.Pointer, overrides TreeViewAccessibleOverrides, classInitFunc func(*TreeViewAccessibleClass)) {
	if classInitFunc != nil {
		class := (*TreeViewAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTreeViewAccessible(obj *coreglib.Object) *TreeViewAccessible {
	return &TreeViewAccessible{
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
		Selection: atk.Selection{
			Object: obj,
		},
		Table: atk.Table{
			Object: obj,
		},
		CellAccessibleParent: CellAccessibleParent{
			Object: obj,
		},
	}
}

func marshalTreeViewAccessible(p uintptr) (interface{}, error) {
	return wrapTreeViewAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TreeViewAccessibleClass: instance of this type is always passed by reference.
type TreeViewAccessibleClass struct {
	*treeViewAccessibleClass
}

// treeViewAccessibleClass is the struct that's finalized.
type treeViewAccessibleClass struct {
	native *C.GtkTreeViewAccessibleClass
}

func (t *TreeViewAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &t.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
