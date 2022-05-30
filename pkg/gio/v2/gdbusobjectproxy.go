// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdbusobjectproxy.go.
var GTypeDBusObjectProxy = coreglib.Type(C.g_dbus_object_proxy_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDBusObjectProxy, F: marshalDBusObjectProxy},
	})
}

// DBusObjectProxyOverrider contains methods that are overridable.
type DBusObjectProxyOverrider interface {
}

// DBusObjectProxy is an object used to represent a remote object with one or
// more D-Bus interfaces. Normally, you don't instantiate a BusObjectProxy
// yourself - typically BusObjectManagerClient is used to obtain it.
type DBusObjectProxy struct {
	_ [0]func() // equal guard
	*coreglib.Object

	DBusObject
}

var (
	_ coreglib.Objector = (*DBusObjectProxy)(nil)
)

func classInitDBusObjectProxier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDBusObjectProxy(obj *coreglib.Object) *DBusObjectProxy {
	return &DBusObjectProxy{
		Object: obj,
		DBusObject: DBusObject{
			Object: obj,
		},
	}
}

func marshalDBusObjectProxy(p uintptr) (interface{}, error) {
	return wrapDBusObjectProxy(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewDBusObjectProxy creates a new BusObjectProxy for the given connection and
// object path.
//
// The function takes the following parameters:
//
//    - connection: BusConnection.
//    - objectPath: object path.
//
// The function returns the following values:
//
//    - dBusObjectProxy: new BusObjectProxy.
//
func NewDBusObjectProxy(connection *DBusConnection, objectPath string) *DBusObjectProxy {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**DBusConnection)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gio", "DBusObjectProxy").InvokeMethod("new_DBusObjectProxy", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(connection)
	runtime.KeepAlive(objectPath)

	var _dBusObjectProxy *DBusObjectProxy // out

	_dBusObjectProxy = wrapDBusObjectProxy(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusObjectProxy
}

// Connection gets the connection that proxy is for.
//
// The function returns the following values:
//
//    - dBusConnection Do not free, the object is owned by proxy.
//
func (proxy *DBusObjectProxy) Connection() *DBusConnection {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))
	*(**DBusObjectProxy)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "DBusObjectProxy").InvokeMethod("get_connection", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(proxy)

	var _dBusConnection *DBusConnection // out

	_dBusConnection = wrapDBusConnection(coreglib.Take(unsafe.Pointer(_cret)))

	return _dBusConnection
}
