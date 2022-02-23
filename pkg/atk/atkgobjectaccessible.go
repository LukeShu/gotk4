// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

// glib.Type values for atkgobjectaccessible.go.
var GTypeGObjectAccessible = externglib.Type(C.atk_gobject_accessible_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGObjectAccessible, F: marshalGObjectAccessible},
	})
}

// GObjectAccessibleOverrider contains methods that are overridable.
type GObjectAccessibleOverrider interface {
}

// GObjectAccessible: this object class is derived from AtkObject. It can be
// used as a basis for implementing accessible objects for GObjects which are
// not derived from GtkWidget. One example of its use is in providing an
// accessible object for GnomeCanvasItem in the GAIL library.
type GObjectAccessible struct {
	_ [0]func() // equal guard
	ObjectClass
}

var (
	_ externglib.Objector = (*GObjectAccessible)(nil)
)

func classInitGObjectAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGObjectAccessible(obj *externglib.Object) *GObjectAccessible {
	return &GObjectAccessible{
		ObjectClass: ObjectClass{
			Object: obj,
		},
	}
}

func marshalGObjectAccessible(p uintptr) (interface{}, error) {
	return wrapGObjectAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Object gets the GObject for which obj is the accessible object.
//
// The function returns the following values:
//
//    - object which is the object for which obj is the accessible object.
//
func (obj *GObjectAccessible) Object() *externglib.Object {
	var _arg0 *C.AtkGObjectAccessible // out
	var _cret *C.GObject              // in

	_arg0 = (*C.AtkGObjectAccessible)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_gobject_accessible_get_object(_arg0)
	runtime.KeepAlive(obj)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// GObjectAccessibleForObject gets the accessible object for the specified obj.
//
// The function takes the following parameters:
//
//    - obj: #GObject.
//
// The function returns the following values:
//
//    - object which is the accessible object for the obj.
//
func GObjectAccessibleForObject(obj *externglib.Object) *ObjectClass {
	var _arg1 *C.GObject   // out
	var _cret *C.AtkObject // in

	_arg1 = (*C.GObject)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_gobject_accessible_for_object(_arg1)
	runtime.KeepAlive(obj)

	var _object *ObjectClass // out

	_object = wrapObject(externglib.Take(unsafe.Pointer(_cret)))

	return _object
}
