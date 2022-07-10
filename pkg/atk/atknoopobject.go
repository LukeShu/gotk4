// Code generated by girgen. DO NOT EDIT.

package atk

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

// GTypeNoOpObject returns the GType for the type NoOpObject.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeNoOpObject() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "NoOpObject").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalNoOpObject)
	return gtype
}

// NoOpObjectOverrider contains methods that are overridable.
type NoOpObjectOverrider interface {
}

// NoOpObject is an AtkObject which purports to implement all ATK interfaces. It
// is the type of AtkObject which is created if an accessible object is
// requested for an object type for which no factory type is specified.
type NoOpObject struct {
	_ [0]func() // equal guard
	ObjectClass

	*coreglib.Object
	Action
	Component
	Document
	EditableText
	Hypertext
	Image
	Selection
	Table
	TableCell
	Text
	Value
	Window
}

var (
	_ coreglib.Objector = (*NoOpObject)(nil)
)

func classInitNoOpObjector(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNoOpObject(obj *coreglib.Object) *NoOpObject {
	return &NoOpObject{
		ObjectClass: ObjectClass{
			Object: obj,
		},
		Object: obj,
		Action: Action{
			Object: obj,
		},
		Component: Component{
			Object: obj,
		},
		Document: Document{
			Object: obj,
		},
		EditableText: EditableText{
			Object: obj,
		},
		Hypertext: Hypertext{
			Object: obj,
		},
		Image: Image{
			Object: obj,
		},
		Selection: Selection{
			Object: obj,
		},
		Table: Table{
			Object: obj,
		},
		TableCell: TableCell{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
		Text: Text{
			Object: obj,
		},
		Value: Value{
			Object: obj,
		},
		Window: Window{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
	}
}

func marshalNoOpObject(p uintptr) (interface{}, error) {
	return wrapNoOpObject(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNoOpObject provides a default (non-functioning stub) Object. Application
// maintainers should not use this method.
//
// The function takes the following parameters:
//
//    - obj: #GObject.
//
// The function returns the following values:
//
//    - noOpObject: default (non-functioning stub) Object.
//
func NewNoOpObject(obj *coreglib.Object) *NoOpObject {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(obj.Native()))

	_info := girepository.MustFind("Atk", "NoOpObject")
	_gret := _info.InvokeClassMethod("new_NoOpObject", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(obj)

	var _noOpObject *NoOpObject // out

	_noOpObject = wrapNoOpObject(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _noOpObject
}
