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
		{T: externglib.Type(C.atk_registry_get_type()), F: marshalRegistry},
	})
}

// GetDefaultRegistry gets a default implementation of the ObjectFactory/type
// registry. Note: For most toolkit maintainers, this will be the correct
// registry for registering new Object factories. Following a call to this
// function, maintainers may call atk_registry_set_factory_type() to associate
// an ObjectFactory subclass with the GType of objects for whom accessibility
// information will be provided.
func GetDefaultRegistry() Registry {
	var _cret *C.AtkRegistry // in

	_cret = C.atk_get_default_registry()

	var _registry Registry // out

	_registry = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Registry)

	return _registry
}

// Registry: the AtkRegistry is normally used to create appropriate ATK "peers"
// for user interface components. Application developers usually need only
// interact with the AtkRegistry by associating appropriate ATK implementation
// classes with GObject classes via the atk_registry_set_factory_type call,
// passing the appropriate GType for application custom widget classes.
type Registry interface {
	gextras.Objector

	// Factory gets an ObjectFactory appropriate for creating Objects
	// appropriate for @type.
	Factory(typ externglib.Type) ObjectFactory
	// FactoryType provides a #GType indicating the ObjectFactory subclass
	// associated with @type.
	FactoryType(typ externglib.Type) externglib.Type
	// SetFactoryType: associate an ObjectFactory subclass with a #GType. Note:
	// The associated @factory_type will thereafter be responsible for the
	// creation of new Object implementations for instances appropriate for
	// @type.
	SetFactoryType(typ externglib.Type, factoryType externglib.Type)
}

// RegistryClass implements the Registry interface.
type RegistryClass struct {
	*externglib.Object
}

var _ Registry = (*RegistryClass)(nil)

func wrapRegistry(obj *externglib.Object) Registry {
	return &RegistryClass{
		Object: obj,
	}
}

func marshalRegistry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRegistry(obj), nil
}

// Factory gets an ObjectFactory appropriate for creating Objects appropriate
// for @type.
func (r *RegistryClass) Factory(typ externglib.Type) ObjectFactory {
	var _arg0 *C.AtkRegistry      // out
	var _arg1 C.GType             // out
	var _cret *C.AtkObjectFactory // in

	_arg0 = (*C.AtkRegistry)(unsafe.Pointer(r.Native()))
	_arg1 = (C.GType)(typ)

	_cret = C.atk_registry_get_factory(_arg0, _arg1)

	var _objectFactory ObjectFactory // out

	_objectFactory = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ObjectFactory)

	return _objectFactory
}

// FactoryType provides a #GType indicating the ObjectFactory subclass
// associated with @type.
func (r *RegistryClass) FactoryType(typ externglib.Type) externglib.Type {
	var _arg0 *C.AtkRegistry // out
	var _arg1 C.GType        // out
	var _cret C.GType        // in

	_arg0 = (*C.AtkRegistry)(unsafe.Pointer(r.Native()))
	_arg1 = (C.GType)(typ)

	_cret = C.atk_registry_get_factory_type(_arg0, _arg1)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SetFactoryType: associate an ObjectFactory subclass with a #GType. Note: The
// associated @factory_type will thereafter be responsible for the creation of
// new Object implementations for instances appropriate for @type.
func (r *RegistryClass) SetFactoryType(typ externglib.Type, factoryType externglib.Type) {
	var _arg0 *C.AtkRegistry // out
	var _arg1 C.GType        // out
	var _arg2 C.GType        // out

	_arg0 = (*C.AtkRegistry)(unsafe.Pointer(r.Native()))
	_arg1 = (C.GType)(typ)
	_arg2 = (C.GType)(factoryType)

	C.atk_registry_set_factory_type(_arg0, _arg1, _arg2)
}
