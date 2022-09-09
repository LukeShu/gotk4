// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// GSocketAddressEnumerator* _gotk4_gio2_SocketConnectable_virtual_enumerate(void* fnptr, GSocketConnectable* arg0) {
//   return ((GSocketAddressEnumerator* (*)(GSocketConnectable*))(fnptr))(arg0);
// };
// GSocketAddressEnumerator* _gotk4_gio2_SocketConnectable_virtual_proxy_enumerate(void* fnptr, GSocketConnectable* arg0) {
//   return ((GSocketAddressEnumerator* (*)(GSocketConnectable*))(fnptr))(arg0);
// };
// gchar* _gotk4_gio2_SocketConnectable_virtual_to_string(void* fnptr, GSocketConnectable* arg0) {
//   return ((gchar* (*)(GSocketConnectable*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeSocketConnectable = coreglib.Type(C.g_socket_connectable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSocketConnectable, F: marshalSocketConnectable},
	})
}

// SocketConnectable objects that describe one or more potential socket
// endpoints implement Connectable. Callers can then use
// g_socket_connectable_enumerate() to get a AddressEnumerator to try out each
// socket address in turn until one succeeds, as shown in the sample code below.
//
//    MyConnectionType *
//    connect_to_host (const char    *hostname,
//                     guint16        port,
//                     GCancellable  *cancellable,
//                     GError       **error)
//    {
//      MyConnection *conn = NULL;
//      GSocketConnectable *addr;
//      GSocketAddressEnumerator *enumerator;
//      GSocketAddress *sockaddr;
//      GError *conn_error = NULL;
//
//      addr = g_network_address_new (hostname, port);
//      enumerator = g_socket_connectable_enumerate (addr);
//      g_object_unref (addr);
//
//      // Try each sockaddr until we succeed. Record the first connection error,
//      // but not any further ones (since they'll probably be basically the same
//      // as the first).
//      while (!conn && (sockaddr = g_socket_address_enumerator_next (enumerator, cancellable, error))
//        {
//          conn = connect_to_sockaddr (sockaddr, conn_error ? NULL : &conn_error);
//          g_object_unref (sockaddr);
//        }
//      g_object_unref (enumerator);
//
//      if (conn)
//        {
//          if (conn_error)
//            {
//              // We couldn't connect to the first address, but we succeeded
//              // in connecting to a later address.
//              g_error_free (conn_error);
//            }
//          return conn;
//        }
//      else if (error)
//        {
//          /// Either initial lookup failed, or else the caller cancelled us.
//          if (conn_error)
//            g_error_free (conn_error);
//          return NULL;
//        }
//      else
//        {
//          g_error_propagate (error, conn_error);
//          return NULL;
//        }
//    }.
//
// SocketConnectable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type SocketConnectable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*SocketConnectable)(nil)
)

// SocketConnectabler describes SocketConnectable's interface methods.
type SocketConnectabler interface {
	coreglib.Objector

	// Enumerate creates a AddressEnumerator for connectable.
	Enumerate() SocketAddressEnumeratorrer
	// ProxyEnumerate creates a AddressEnumerator for connectable that will
	// return a Address for each of its addresses that you must connect to via a
	// proxy.
	ProxyEnumerate() SocketAddressEnumeratorrer
	// String: format a Connectable as a string.
	String() string
}

var _ SocketConnectabler = (*SocketConnectable)(nil)

func wrapSocketConnectable(obj *coreglib.Object) *SocketConnectable {
	return &SocketConnectable{
		Object: obj,
	}
}

func marshalSocketConnectable(p uintptr) (interface{}, error) {
	return wrapSocketConnectable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Enumerate creates a AddressEnumerator for connectable.
//
// The function returns the following values:
//
//    - socketAddressEnumerator: new AddressEnumerator.
//
func (connectable *SocketConnectable) Enumerate() SocketAddressEnumeratorrer {
	var _arg0 *C.GSocketConnectable       // out
	var _cret *C.GSocketAddressEnumerator // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	_cret = C.g_socket_connectable_enumerate(_arg0)
	runtime.KeepAlive(connectable)

	var _socketAddressEnumerator SocketAddressEnumeratorrer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddressEnumeratorrer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddressEnumeratorrer)
			return ok
		})
		rv, ok := casted.(SocketAddressEnumeratorrer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddressEnumeratorrer")
		}
		_socketAddressEnumerator = rv
	}

	return _socketAddressEnumerator
}

// ProxyEnumerate creates a AddressEnumerator for connectable that will return a
// Address for each of its addresses that you must connect to via a proxy.
//
// If connectable does not implement g_socket_connectable_proxy_enumerate(),
// this will fall back to calling g_socket_connectable_enumerate().
//
// The function returns the following values:
//
//    - socketAddressEnumerator: new AddressEnumerator.
//
func (connectable *SocketConnectable) ProxyEnumerate() SocketAddressEnumeratorrer {
	var _arg0 *C.GSocketConnectable       // out
	var _cret *C.GSocketAddressEnumerator // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	_cret = C.g_socket_connectable_proxy_enumerate(_arg0)
	runtime.KeepAlive(connectable)

	var _socketAddressEnumerator SocketAddressEnumeratorrer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddressEnumeratorrer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddressEnumeratorrer)
			return ok
		})
		rv, ok := casted.(SocketAddressEnumeratorrer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddressEnumeratorrer")
		}
		_socketAddressEnumerator = rv
	}

	return _socketAddressEnumerator
}

// String: format a Connectable as a string. This is a human-readable format for
// use in debugging output, and is not a stable serialization format. It is not
// suitable for use in user interfaces as it exposes too much information for a
// user.
//
// If the Connectable implementation does not support string formatting, the
// implementation’s type name will be returned as a fallback.
//
// The function returns the following values:
//
//    - utf8: formatted string.
//
func (connectable *SocketConnectable) String() string {
	var _arg0 *C.GSocketConnectable // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	_cret = C.g_socket_connectable_to_string(_arg0)
	runtime.KeepAlive(connectable)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Enumerate creates a AddressEnumerator for connectable.
//
// The function returns the following values:
//
//    - socketAddressEnumerator: new AddressEnumerator.
//
func (connectable *SocketConnectable) enumerate() SocketAddressEnumeratorrer {
	gclass := (*C.GSocketConnectableIface)(coreglib.PeekParentClass(connectable))
	fnarg := gclass.enumerate

	var _arg0 *C.GSocketConnectable       // out
	var _cret *C.GSocketAddressEnumerator // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	_cret = C._gotk4_gio2_SocketConnectable_virtual_enumerate(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(connectable)

	var _socketAddressEnumerator SocketAddressEnumeratorrer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddressEnumeratorrer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddressEnumeratorrer)
			return ok
		})
		rv, ok := casted.(SocketAddressEnumeratorrer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddressEnumeratorrer")
		}
		_socketAddressEnumerator = rv
	}

	return _socketAddressEnumerator
}

// proXYEnumerate creates a AddressEnumerator for connectable that will return a
// Address for each of its addresses that you must connect to via a proxy.
//
// If connectable does not implement g_socket_connectable_proxy_enumerate(),
// this will fall back to calling g_socket_connectable_enumerate().
//
// The function returns the following values:
//
//    - socketAddressEnumerator: new AddressEnumerator.
//
func (connectable *SocketConnectable) proxyEnumerate() SocketAddressEnumeratorrer {
	gclass := (*C.GSocketConnectableIface)(coreglib.PeekParentClass(connectable))
	fnarg := gclass.proxy_enumerate

	var _arg0 *C.GSocketConnectable       // out
	var _cret *C.GSocketAddressEnumerator // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	_cret = C._gotk4_gio2_SocketConnectable_virtual_proxy_enumerate(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(connectable)

	var _socketAddressEnumerator SocketAddressEnumeratorrer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.SocketAddressEnumeratorrer is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddressEnumeratorrer)
			return ok
		})
		rv, ok := casted.(SocketAddressEnumeratorrer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddressEnumeratorrer")
		}
		_socketAddressEnumerator = rv
	}

	return _socketAddressEnumerator
}

// Str: format a Connectable as a string. This is a human-readable format for
// use in debugging output, and is not a stable serialization format. It is not
// suitable for use in user interfaces as it exposes too much information for a
// user.
//
// If the Connectable implementation does not support string formatting, the
// implementation’s type name will be returned as a fallback.
//
// The function returns the following values:
//
//    - utf8: formatted string.
//
func (connectable *SocketConnectable) str() string {
	gclass := (*C.GSocketConnectableIface)(coreglib.PeekParentClass(connectable))
	fnarg := gclass.to_string

	var _arg0 *C.GSocketConnectable // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(connectable).Native()))

	_cret = C._gotk4_gio2_SocketConnectable_virtual_to_string(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(connectable)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SocketConnectableIface provides an interface for returning a
// AddressEnumerator and AddressEnumerator
//
// An instance of this type is always passed by reference.
type SocketConnectableIface struct {
	*socketConnectableIface
}

// socketConnectableIface is the struct that's finalized.
type socketConnectableIface struct {
	native *C.GSocketConnectableIface
}
