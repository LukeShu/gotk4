// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

// GTypeRegistry returns the GType for the type Registry.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRegistry() coreglib.Type {
	gtype := coreglib.Type(C.atk_registry_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalRegistry)
	return gtype
}

// GetDefaultRegistry gets a default implementation of the ObjectFactory/type
// registry. Note: For most toolkit maintainers, this will be the correct
// registry for registering new Object factories. Following a call to this
// function, maintainers may call atk_registry_set_factory_type() to associate
// an ObjectFactory subclass with the GType of objects for whom accessibility
// information will be provided.
//
// The function returns the following values:
//
//    - registry: default implementation of the ObjectFactory/type registry.
//
func GetDefaultRegistry() *Registry {
	var _cret *C.AtkRegistry // in

	_cret = C.atk_get_default_registry()

	var _registry *Registry // out

	_registry = wrapRegistry(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _registry
}

// RegistryOverrider contains methods that are overridable.
type RegistryOverrider interface {
}

// Registry is normally used to create appropriate ATK "peers" for user
// interface components. Application developers usually need only interact with
// the AtkRegistry by associating appropriate ATK implementation classes with
// GObject classes via the atk_registry_set_factory_type call, passing the
// appropriate GType for application custom widget classes.
type Registry struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Registry)(nil)
)

func classInitRegistrier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRegistry(obj *coreglib.Object) *Registry {
	return &Registry{
		Object: obj,
	}
}

func marshalRegistry(p uintptr) (interface{}, error) {
	return wrapRegistry(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Factory gets an ObjectFactory appropriate for creating Objects appropriate
// for type.
//
// The function takes the following parameters:
//
//    - typ with which to look up the associated ObjectFactory.
//
// The function returns the following values:
//
//    - objectFactory appropriate for creating Objects appropriate for type.
//
func (registry *Registry) Factory(typ coreglib.Type) *ObjectFactory {
	var _arg0 *C.AtkRegistry      // out
	var _arg1 C.GType             // out
	var _cret *C.AtkObjectFactory // in

	_arg0 = (*C.AtkRegistry)(unsafe.Pointer(coreglib.InternObject(registry).Native()))
	_arg1 = C.GType(typ)

	_cret = C.atk_registry_get_factory(_arg0, _arg1)
	runtime.KeepAlive(registry)
	runtime.KeepAlive(typ)

	var _objectFactory *ObjectFactory // out

	_objectFactory = wrapObjectFactory(coreglib.Take(unsafe.Pointer(_cret)))

	return _objectFactory
}

// FactoryType provides a #GType indicating the ObjectFactory subclass
// associated with type.
//
// The function takes the following parameters:
//
//    - typ with which to look up the associated ObjectFactory subclass.
//
// The function returns the following values:
//
//    - gType associated with type type.
//
func (registry *Registry) FactoryType(typ coreglib.Type) coreglib.Type {
	var _arg0 *C.AtkRegistry // out
	var _arg1 C.GType        // out
	var _cret C.GType        // in

	_arg0 = (*C.AtkRegistry)(unsafe.Pointer(coreglib.InternObject(registry).Native()))
	_arg1 = C.GType(typ)

	_cret = C.atk_registry_get_factory_type(_arg0, _arg1)
	runtime.KeepAlive(registry)
	runtime.KeepAlive(typ)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// SetFactoryType: associate an ObjectFactory subclass with a #GType. Note: The
// associated factory_type will thereafter be responsible for the creation of
// new Object implementations for instances appropriate for type.
//
// The function takes the following parameters:
//
//    - typ: Object type.
//    - factoryType type to associate with type. Must implement AtkObject
//      appropriate for type.
//
func (registry *Registry) SetFactoryType(typ, factoryType coreglib.Type) {
	var _arg0 *C.AtkRegistry // out
	var _arg1 C.GType        // out
	var _arg2 C.GType        // out

	_arg0 = (*C.AtkRegistry)(unsafe.Pointer(coreglib.InternObject(registry).Native()))
	_arg1 = C.GType(typ)
	_arg2 = C.GType(factoryType)

	C.atk_registry_set_factory_type(_arg0, _arg1, _arg2)
	runtime.KeepAlive(registry)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(factoryType)
}
