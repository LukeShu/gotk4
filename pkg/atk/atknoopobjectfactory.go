// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_no_op_object_factory_get_type()), F: marshalNoOpObjectFactory},
	})
}

// NoOpObjectFactory: the AtkObjectFactory which creates an AtkNoOpObject. An
// instance of this is created by an AtkRegistry if no factory type has not been
// specified to create an accessible object of a particular type.
type NoOpObjectFactory interface {
	gextras.Objector

	privateNoOpObjectFactoryClass()
}

// NoOpObjectFactoryClass implements the NoOpObjectFactory interface.
type NoOpObjectFactoryClass struct {
	ObjectFactoryClass
}

var _ NoOpObjectFactory = (*NoOpObjectFactoryClass)(nil)

func wrapNoOpObjectFactory(obj *externglib.Object) NoOpObjectFactory {
	return &NoOpObjectFactoryClass{
		ObjectFactoryClass: ObjectFactoryClass{
			Object: obj,
		},
	}
}

func marshalNoOpObjectFactory(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNoOpObjectFactory(obj), nil
}

// NewNoOpObjectFactory creates an instance of an ObjectFactory which generates
// primitive (non-functioning) Objects.
func NewNoOpObjectFactory() NoOpObjectFactory {
	var _cret *C.AtkObjectFactory // in

	_cret = C.atk_no_op_object_factory_new()

	var _noOpObjectFactory NoOpObjectFactory // out

	_noOpObjectFactory = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(NoOpObjectFactory)

	return _noOpObjectFactory
}

func (*NoOpObjectFactoryClass) privateNoOpObjectFactoryClass() {}
