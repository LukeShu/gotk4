// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gnetworkaddress.go.
var GTypeNetworkAddress = externglib.Type(C.g_network_address_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeNetworkAddress, F: marshalNetworkAddress},
	})
}

// NetworkAddressOverrider contains methods that are overridable.
type NetworkAddressOverrider interface {
}

// NetworkAddress provides an easy way to resolve a hostname and then attempt to
// connect to that host, handling the possibility of multiple IP addresses and
// multiple address families.
//
// The enumeration results of resolved addresses *may* be cached as long as this
// object is kept alive which may have unexpected results if alive for too long.
//
// See Connectable for an example of using the connectable interface.
type NetworkAddress struct {
	_ [0]func() // equal guard
	*externglib.Object

	SocketConnectable
}

var (
	_ externglib.Objector = (*NetworkAddress)(nil)
)

func classInitNetworkAddresser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNetworkAddress(obj *externglib.Object) *NetworkAddress {
	return &NetworkAddress{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalNetworkAddress(p uintptr) (interface{}, error) {
	return wrapNetworkAddress(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNetworkAddress creates a new Connectable for connecting to the given
// hostname and port.
//
// Note that depending on the configuration of the machine, a hostname of
// localhost may refer to the IPv4 loopback address only, or to both IPv4 and
// IPv6; use g_network_address_new_loopback() to create a Address that is
// guaranteed to resolve to both addresses.
//
// The function takes the following parameters:
//
//    - hostname: hostname.
//    - port: port.
//
// The function returns the following values:
//
//    - networkAddress: new Address.
//
func NewNetworkAddress(hostname string, port uint16) *NetworkAddress {
	var _arg1 *C.gchar              // out
	var _arg2 C.guint16             // out
	var _cret *C.GSocketConnectable // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(port)

	_cret = C.g_network_address_new(_arg1, _arg2)
	runtime.KeepAlive(hostname)
	runtime.KeepAlive(port)

	var _networkAddress *NetworkAddress // out

	_networkAddress = wrapNetworkAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _networkAddress
}

// NewNetworkAddressLoopback creates a new Connectable for connecting to the
// local host over a loopback connection to the given port. This is intended for
// use in connecting to local services which may be running on IPv4 or IPv6.
//
// The connectable will return IPv4 and IPv6 loopback addresses, regardless of
// how the host resolves localhost. By contrast, g_network_address_new() will
// often only return an IPv4 address when resolving localhost, and an IPv6
// address for localhost6.
//
// g_network_address_get_hostname() will always return localhost for a Address
// created with this constructor.
//
// The function takes the following parameters:
//
//    - port: port.
//
// The function returns the following values:
//
//    - networkAddress: new Address.
//
func NewNetworkAddressLoopback(port uint16) *NetworkAddress {
	var _arg1 C.guint16             // out
	var _cret *C.GSocketConnectable // in

	_arg1 = C.guint16(port)

	_cret = C.g_network_address_new_loopback(_arg1)
	runtime.KeepAlive(port)

	var _networkAddress *NetworkAddress // out

	_networkAddress = wrapNetworkAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _networkAddress
}

// Hostname gets addr's hostname. This might be either UTF-8 or ASCII-encoded,
// depending on what addr was created with.
//
// The function returns the following values:
//
//    - utf8 addr's hostname.
//
func (addr *NetworkAddress) Hostname() string {
	var _arg0 *C.GNetworkAddress // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.g_network_address_get_hostname(_arg0)
	runtime.KeepAlive(addr)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Port gets addr's port number.
//
// The function returns the following values:
//
//    - guint16 addr's port (which may be 0).
//
func (addr *NetworkAddress) Port() uint16 {
	var _arg0 *C.GNetworkAddress // out
	var _cret C.guint16          // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.g_network_address_get_port(_arg0)
	runtime.KeepAlive(addr)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Scheme gets addr's scheme.
//
// The function returns the following values:
//
//    - utf8 (optional) addr's scheme (NULL if not built from URI).
//
func (addr *NetworkAddress) Scheme() string {
	var _arg0 *C.GNetworkAddress // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkAddress)(unsafe.Pointer(addr.Native()))

	_cret = C.g_network_address_get_scheme(_arg0)
	runtime.KeepAlive(addr)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NetworkAddressParse creates a new Connectable for connecting to the given
// hostname and port. May fail and return NULL in case parsing host_and_port
// fails.
//
// host_and_port may be in any of a number of recognised formats; an IPv6
// address, an IPv4 address, or a domain name (in which case a DNS lookup is
// performed). Quoting with [] is supported for all address types. A port
// override may be specified in the usual way with a colon.
//
// If no port is specified in host_and_port then default_port will be used as
// the port number to connect to.
//
// In general, host_and_port is expected to be provided by the user (allowing
// them to give the hostname, and a port override if necessary) and default_port
// is expected to be provided by the application.
//
// (The port component of host_and_port can also be specified as a service name
// rather than as a numeric port, but this functionality is deprecated, because
// it depends on the contents of /etc/services, which is generally quite sparse
// on platforms other than Linux.).
//
// The function takes the following parameters:
//
//    - hostAndPort: hostname and optionally a port.
//    - defaultPort: default port if not in host_and_port.
//
// The function returns the following values:
//
//    - networkAddress: new Address, or NULL on error.
//
func NetworkAddressParse(hostAndPort string, defaultPort uint16) (*NetworkAddress, error) {
	var _arg1 *C.gchar              // out
	var _arg2 C.guint16             // out
	var _cret *C.GSocketConnectable // in
	var _cerr *C.GError             // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostAndPort)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)

	_cret = C.g_network_address_parse(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(hostAndPort)
	runtime.KeepAlive(defaultPort)

	var _networkAddress *NetworkAddress // out
	var _goerr error                    // out

	_networkAddress = wrapNetworkAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _networkAddress, _goerr
}

// NetworkAddressParseURI creates a new Connectable for connecting to the given
// uri. May fail and return NULL in case parsing uri fails.
//
// Using this rather than g_network_address_new() or g_network_address_parse()
// allows Client to determine when to use application-specific proxy protocols.
//
// The function takes the following parameters:
//
//    - uri: hostname and optionally a port.
//    - defaultPort: default port if none is found in the URI.
//
// The function returns the following values:
//
//    - networkAddress: new Address, or NULL on error.
//
func NetworkAddressParseURI(uri string, defaultPort uint16) (*NetworkAddress, error) {
	var _arg1 *C.gchar              // out
	var _arg2 C.guint16             // out
	var _cret *C.GSocketConnectable // in
	var _cerr *C.GError             // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(defaultPort)

	_cret = C.g_network_address_parse_uri(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(defaultPort)

	var _networkAddress *NetworkAddress // out
	var _goerr error                    // out

	_networkAddress = wrapNetworkAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _networkAddress, _goerr
}
