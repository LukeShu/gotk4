// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for atkregistry.go.
var GTypeRegistry = coreglib.Type(C.atk_registry_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeRegistry, F: marshalRegistry},
	})
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
	var _cret *C.void // in

	_gret := girepository.MustFind("Atk", "get_default_registry").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
