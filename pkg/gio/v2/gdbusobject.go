// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dbus_object_get_type()), F: marshalDBusObjector},
	})
}

// DBusObjectOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectOverrider interface {
	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	Interface(interfaceName string) DBusInterfacer
	// Interfaces gets the D-Bus interfaces associated with object.
	Interfaces() []DBusInterfacer
	// ObjectPath gets the object path for object.
	ObjectPath() string
	InterfaceAdded(interface_ DBusInterfacer)
	InterfaceRemoved(interface_ DBusInterfacer)
}

// DBusObject type is the base type for D-Bus objects on both the service side
// (see BusObjectSkeleton) and the client side (see BusObjectProxy). It is
// essentially just a container of interfaces.
type DBusObject struct {
	*externglib.Object
}

// DBusObjector describes DBusObject's abstract methods.
type DBusObjector interface {
	externglib.Objector

	// Interface gets the D-Bus interface with name interface_name associated
	// with object, if any.
	Interface(interfaceName string) DBusInterfacer
	// Interfaces gets the D-Bus interfaces associated with object.
	Interfaces() []DBusInterfacer
	// ObjectPath gets the object path for object.
	ObjectPath() string
}

var _ DBusObjector = (*DBusObject)(nil)

func wrapDBusObject(obj *externglib.Object) *DBusObject {
	return &DBusObject{
		Object: obj,
	}
}

func marshalDBusObjector(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObject(obj), nil
}

// Interface gets the D-Bus interface with name interface_name associated with
// object, if any.
func (object *DBusObject) Interface(interfaceName string) DBusInterfacer {
	var _arg0 *C.GDBusObject    // out
	var _arg1 *C.gchar          // out
	var _cret *C.GDBusInterface // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(interfaceName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_get_interface(_arg0, _arg1)

	runtime.KeepAlive(object)
	runtime.KeepAlive(interfaceName)

	var _dBusInterface DBusInterfacer // out

	if _cret != nil {
		_dBusInterface = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(DBusInterfacer)
	}

	return _dBusInterface
}

// Interfaces gets the D-Bus interfaces associated with object.
func (object *DBusObject) Interfaces() []DBusInterfacer {
	var _arg0 *C.GDBusObject // out
	var _cret *C.GList       // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))

	_cret = C.g_dbus_object_get_interfaces(_arg0)

	runtime.KeepAlive(object)

	var _list []DBusInterfacer // out

	_list = make([]DBusInterfacer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GDBusInterface)(v)
		var dst DBusInterfacer // out
		dst = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(DBusInterfacer)
		_list = append(_list, dst)
	})

	return _list
}

// ObjectPath gets the object path for object.
func (object *DBusObject) ObjectPath() string {
	var _arg0 *C.GDBusObject // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GDBusObject)(unsafe.Pointer(object.Native()))

	_cret = C.g_dbus_object_get_object_path(_arg0)

	runtime.KeepAlive(object)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
