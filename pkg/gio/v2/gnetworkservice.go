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
// #include <glib-object.h>
import "C"

// GTypeNetworkService returns the GType for the type NetworkService.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeNetworkService() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "NetworkService").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalNetworkService)
	return gtype
}

// NetworkServiceOverrider contains methods that are overridable.
type NetworkServiceOverrider interface {
}

// NetworkService: like Address does with hostnames, Service provides an easy
// way to resolve a SRV record, and then attempt to connect to one of the hosts
// that implements that service, handling service priority/weighting, multiple
// IP addresses, and multiple address families.
//
// See Target for more information about SRV records, and see Connectable for an
// example of using the connectable interface.
type NetworkService struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SocketConnectable
}

var (
	_ coreglib.Objector = (*NetworkService)(nil)
)

func classInitNetworkServicer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNetworkService(obj *coreglib.Object) *NetworkService {
	return &NetworkService{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalNetworkService(p uintptr) (interface{}, error) {
	return wrapNetworkService(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNetworkService creates a new Service representing the given service,
// protocol, and domain. This will initially be unresolved; use the Connectable
// interface to resolve it.
//
// The function takes the following parameters:
//
//    - service type to look up (eg, "ldap").
//    - protocol: networking protocol to use for service (eg, "tcp").
//    - domain: DNS domain to look up the service in.
//
// The function returns the following values:
//
//    - networkService: new Service.
//
func NewNetworkService(service, protocol, domain string) *NetworkService {
	var _args [3]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))
	*(**C.gchar)(unsafe.Pointer(&_args[2])) = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[2]))))

	_info := girepository.MustFind("Gio", "NetworkService")
	_gret := _info.InvokeClassMethod("new_NetworkService", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(service)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(domain)

	var _networkService *NetworkService // out

	_networkService = wrapNetworkService(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _networkService
}

// Domain gets the domain that srv serves. This might be either UTF-8 or
// ASCII-encoded, depending on what srv was created with.
//
// The function returns the following values:
//
//    - utf8 srv's domain name.
//
func (srv *NetworkService) Domain() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_info := girepository.MustFind("Gio", "NetworkService")
	_gret := _info.InvokeClassMethod("get_domain", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))

	return _utf8
}

// Protocol gets srv's protocol name (eg, "tcp").
//
// The function returns the following values:
//
//    - utf8 srv's protocol name.
//
func (srv *NetworkService) Protocol() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_info := girepository.MustFind("Gio", "NetworkService")
	_gret := _info.InvokeClassMethod("get_protocol", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))

	return _utf8
}

// Scheme gets the URI scheme used to resolve proxies. By default, the service
// name is used as scheme.
//
// The function returns the following values:
//
//    - utf8 srv's scheme name.
//
func (srv *NetworkService) Scheme() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_info := girepository.MustFind("Gio", "NetworkService")
	_gret := _info.InvokeClassMethod("get_scheme", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))

	return _utf8
}

// Service gets srv's service name (eg, "ldap").
//
// The function returns the following values:
//
//    - utf8 srv's service name.
//
func (srv *NetworkService) Service() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(srv).Native()))

	_info := girepository.MustFind("Gio", "NetworkService")
	_gret := _info.InvokeClassMethod("get_service", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(srv)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))

	return _utf8
}

// SetScheme set's the URI scheme used to resolve proxies. By default, the
// service name is used as scheme.
//
// The function takes the following parameters:
//
//    - scheme: URI scheme.
//
func (srv *NetworkService) SetScheme(scheme string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(srv).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gio", "NetworkService")
	_info.InvokeClassMethod("set_scheme", _args[:], nil)

	runtime.KeepAlive(srv)
	runtime.KeepAlive(scheme)
}
