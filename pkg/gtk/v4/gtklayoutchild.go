// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeLayoutChild returns the GType for the type LayoutChild.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeLayoutChild() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "LayoutChild").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalLayoutChild)
	return gtype
}

// LayoutChildOverrider contains methods that are overridable.
type LayoutChildOverrider interface {
}

// LayoutChild: GtkLayoutChild is the base class for objects that are meant to
// hold layout properties.
//
// If a GtkLayoutManager has per-child properties, like their packing type, or
// the horizontal and vertical span, or the icon name, then the layout manager
// should use a GtkLayoutChild implementation to store those properties.
//
// A GtkLayoutChild instance is only ever valid while a widget is part of a
// layout.
type LayoutChild struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*LayoutChild)(nil)
)

// LayoutChilder describes types inherited from class LayoutChild.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type LayoutChilder interface {
	coreglib.Objector
	baseLayoutChild() *LayoutChild
}

var _ LayoutChilder = (*LayoutChild)(nil)

func classInitLayoutChilder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapLayoutChild(obj *coreglib.Object) *LayoutChild {
	return &LayoutChild{
		Object: obj,
	}
}

func marshalLayoutChild(p uintptr) (interface{}, error) {
	return wrapLayoutChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (layoutChild *LayoutChild) baseLayoutChild() *LayoutChild {
	return layoutChild
}

// BaseLayoutChild returns the underlying base object.
func BaseLayoutChild(obj LayoutChilder) *LayoutChild {
	return obj.baseLayoutChild()
}

// ChildWidget retrieves the GtkWidget associated to the given layout_child.
//
// The function returns the following values:
//
//    - widget: Widget.
//
func (layoutChild *LayoutChild) ChildWidget() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layoutChild).Native()))

	_info := girepository.MustFind("Gtk", "LayoutChild")
	_gret := _info.InvokeClassMethod("get_child_widget", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layoutChild)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// LayoutManager retrieves the GtkLayoutManager instance that created the given
// layout_child.
//
// The function returns the following values:
//
//    - layoutManager: GtkLayoutManager.
//
func (layoutChild *LayoutChild) LayoutManager() LayoutManagerer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layoutChild).Native()))

	_info := girepository.MustFind("Gtk", "LayoutChild")
	_gret := _info.InvokeClassMethod("get_layout_manager", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layoutChild)

	var _layoutManager LayoutManagerer // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gtk.LayoutManagerer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(LayoutManagerer)
			return ok
		})
		rv, ok := casted.(LayoutManagerer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.LayoutManagerer")
		}
		_layoutManager = rv
	}

	return _layoutManager
}
