// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
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
	GTypeContainerCellAccessible = coreglib.Type(C.gtk_container_cell_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContainerCellAccessible, F: marshalContainerCellAccessible},
	})
}

// ContainerCellAccessibleOverrides contains methods that are overridable.
type ContainerCellAccessibleOverrides struct {
}

func defaultContainerCellAccessibleOverrides(v *ContainerCellAccessible) ContainerCellAccessibleOverrides {
	return ContainerCellAccessibleOverrides{}
}

type ContainerCellAccessible struct {
	_ [0]func() // equal guard
	CellAccessible
}

var (
	_ coreglib.Objector = (*ContainerCellAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ContainerCellAccessible, *ContainerCellAccessibleClass, ContainerCellAccessibleOverrides](
		GTypeContainerCellAccessible,
		initContainerCellAccessibleClass,
		wrapContainerCellAccessible,
		defaultContainerCellAccessibleOverrides,
	)
}

func initContainerCellAccessibleClass(gclass unsafe.Pointer, overrides ContainerCellAccessibleOverrides, classInitFunc func(*ContainerCellAccessibleClass)) {
	if classInitFunc != nil {
		class := (*ContainerCellAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapContainerCellAccessible(obj *coreglib.Object) *ContainerCellAccessible {
	return &ContainerCellAccessible{
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
	}
}

func marshalContainerCellAccessible(p uintptr) (interface{}, error) {
	return wrapContainerCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
func NewContainerCellAccessible() *ContainerCellAccessible {
	var _cret *C.GtkContainerCellAccessible // in

	_cret = C.gtk_container_cell_accessible_new()

	var _containerCellAccessible *ContainerCellAccessible // out

	_containerCellAccessible = wrapContainerCellAccessible(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _containerCellAccessible
}

// The function takes the following parameters:
//
func (container *ContainerCellAccessible) AddChild(child *CellAccessible) {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _arg1 *C.GtkCellAccessible          // out

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(coreglib.InternObject(container).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_container_cell_accessible_add_child(_arg0, _arg1)
	runtime.KeepAlive(container)
	runtime.KeepAlive(child)
}

// Children: get a list of children.
//
// The function returns the following values:
//
func (container *ContainerCellAccessible) Children() []*CellAccessible {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _cret *C.GList                      // in

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(coreglib.InternObject(container).Native()))

	_cret = C.gtk_container_cell_accessible_get_children(_arg0)
	runtime.KeepAlive(container)

	var _list []*CellAccessible // out

	_list = make([]*CellAccessible, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkCellAccessible)(v)
		var dst *CellAccessible // out
		dst = wrapCellAccessible(coreglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// The function takes the following parameters:
//
func (container *ContainerCellAccessible) RemoveChild(child *CellAccessible) {
	var _arg0 *C.GtkContainerCellAccessible // out
	var _arg1 *C.GtkCellAccessible          // out

	_arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(coreglib.InternObject(container).Native()))
	_arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_container_cell_accessible_remove_child(_arg0, _arg1)
	runtime.KeepAlive(container)
	runtime.KeepAlive(child)
}

// ContainerCellAccessibleClass: instance of this type is always passed by
// reference.
type ContainerCellAccessibleClass struct {
	*containerCellAccessibleClass
}

// containerCellAccessibleClass is the struct that's finalized.
type containerCellAccessibleClass struct {
	native *C.GtkContainerCellAccessibleClass
}

func (c *ContainerCellAccessibleClass) ParentClass() *CellAccessibleClass {
	valptr := &c.native.parent_class
	var _v *CellAccessibleClass // out
	_v = (*CellAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
