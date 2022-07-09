// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GDBusInterfaceInfo* _gotk4_gio2_DBusInterfaceIface_get_info(void*);
// extern GDBusObject* _gotk4_gio2_DBusInterfaceIface_dup_object(void*);
// extern void _gotk4_gio2_DBusInterfaceIface_set_object(void*, void*);
import "C"

// GTypeDBusInterface returns the GType for the type DBusInterface.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusInterface() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusInterface").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusInterface)
	return gtype
}

// DBusInterfaceOverrider contains methods that are overridable.
type DBusInterfaceOverrider interface {
	// DupObject gets the BusObject that interface_ belongs to, if any.
	//
	// The function returns the following values:
	//
	//    - dBusObject (optional) or NULL. The returned reference should be freed
	//      with g_object_unref().
	//
	DupObject() *DBusObject
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by interface_.
	//
	// The function returns the following values:
	//
	//    - dBusInterfaceInfo Do not free.
	//
	Info() *DBusInterfaceInfo
	// SetObject sets the BusObject for interface_ to object.
	//
	// Note that interface_ will hold a weak reference to object.
	//
	// The function takes the following parameters:
	//
	//    - object (optional) or NULL.
	//
	SetObject(object DBusObjector)
}

// DBusInterface type is the base type for D-Bus interfaces both on the service
// side (see BusInterfaceSkeleton) and client side (see BusProxy).
//
// DBusInterface wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DBusInterface struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DBusInterface)(nil)
)

// DBusInterfacer describes DBusInterface's interface methods.
type DBusInterfacer interface {
	coreglib.Objector

	// GetObject gets the BusObject that interface_ belongs to, if any.
	GetObject() *DBusObject
	// Info gets D-Bus introspection information for the D-Bus interface
	// implemented by interface_.
	Info() *DBusInterfaceInfo
	// SetObject sets the BusObject for interface_ to object.
	SetObject(object DBusObjector)
}

var _ DBusInterfacer = (*DBusInterface)(nil)

func ifaceInitDBusInterfacer(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gio", "DBusInterfaceIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("dup_object"))) = unsafe.Pointer(C._gotk4_gio2_DBusInterfaceIface_dup_object)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_info"))) = unsafe.Pointer(C._gotk4_gio2_DBusInterfaceIface_get_info)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("set_object"))) = unsafe.Pointer(C._gotk4_gio2_DBusInterfaceIface_set_object)
}

//export _gotk4_gio2_DBusInterfaceIface_dup_object
func _gotk4_gio2_DBusInterfaceIface_dup_object(arg0 *C.void) (cret *C.GDBusObject) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusInterfaceOverrider)

	dBusObject := iface.DupObject()

	if dBusObject != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(dBusObject).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(dBusObject).Native()))
	}

	return cret
}

//export _gotk4_gio2_DBusInterfaceIface_get_info
func _gotk4_gio2_DBusInterfaceIface_get_info(arg0 *C.void) (cret *C.GDBusInterfaceInfo) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusInterfaceOverrider)

	dBusInterfaceInfo := iface.Info()

	cret = (*C.void)(gextras.StructNative(unsafe.Pointer(dBusInterfaceInfo)))

	return cret
}

//export _gotk4_gio2_DBusInterfaceIface_set_object
func _gotk4_gio2_DBusInterfaceIface_set_object(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(DBusInterfaceOverrider)

	var _object DBusObjector // out

	if arg1 != nil {
		{
			objptr := unsafe.Pointer(arg1)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(DBusObjector)
				return ok
			})
			rv, ok := casted.(DBusObjector)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.DBusObjector")
			}
			_object = rv
		}
	}

	iface.SetObject(_object)
}

func wrapDBusInterface(obj *coreglib.Object) *DBusInterface {
	return &DBusInterface{
		Object: obj,
	}
}

func marshalDBusInterface(p uintptr) (interface{}, error) {
	return wrapDBusInterface(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// GetObject gets the BusObject that interface_ belongs to, if any.
//
// The function returns the following values:
//
//    - dBusObject (optional) or NULL. The returned reference should be freed
//      with g_object_unref().
//
func (interface_ *DBusInterface) GetObject() *DBusObject {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)

	var _dBusObject *DBusObject // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusObject = wrapDBusObject(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _dBusObject
}

// Info gets D-Bus introspection information for the D-Bus interface implemented
// by interface_.
//
// The function returns the following values:
//
//    - dBusInterfaceInfo Do not free.
//
func (interface_ *DBusInterface) Info() *DBusInterfaceInfo {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(interface_)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dBusInterfaceInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _dBusInterfaceInfo
}

// SetObject sets the BusObject for interface_ to object.
//
// Note that interface_ will hold a weak reference to object.
//
// The function takes the following parameters:
//
//    - object (optional) or NULL.
//
func (interface_ *DBusInterface) SetObject(object DBusObjector) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(interface_).Native()))
	if object != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	}

	runtime.KeepAlive(interface_)
	runtime.KeepAlive(object)
}
