// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.g_inet_address_get_type()), F: marshalInetAddresser},
	})
}

// InetAddressOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type InetAddressOverrider interface {
	// String converts address to string form.
	String() string
}

// InetAddress represents an IPv4 or IPv6 internet address. Use
// g_resolver_lookup_by_name() or g_resolver_lookup_by_name_async() to look up
// the Address for a hostname. Use g_resolver_lookup_by_address() or
// g_resolver_lookup_by_address_async() to look up the hostname for a Address.
//
// To actually connect to a remote host, you will need a SocketAddress (which
// includes a Address as well as a port number).
type InetAddress struct {
	*externglib.Object
}

func wrapInetAddress(obj *externglib.Object) *InetAddress {
	return &InetAddress{
		Object: obj,
	}
}

func marshalInetAddresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInetAddress(obj), nil
}

// NewInetAddressAny creates a Address for the "any" address (unassigned/"don't
// care") for family.
func NewInetAddressAny(family SocketFamily) *InetAddress {
	var _arg1 C.GSocketFamily // out
	var _cret *C.GInetAddress // in

	_arg1 = C.GSocketFamily(family)

	_cret = C.g_inet_address_new_any(_arg1)

	runtime.KeepAlive(family)

	var _inetAddress *InetAddress // out

	_inetAddress = wrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetAddress
}

// NewInetAddressFromString parses string as an IP address and creates a new
// Address.
func NewInetAddressFromString(_string string) *InetAddress {
	var _arg1 *C.gchar        // out
	var _cret *C.GInetAddress // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_inet_address_new_from_string(_arg1)

	runtime.KeepAlive(_string)

	var _inetAddress *InetAddress // out

	if _cret != nil {
		_inetAddress = wrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _inetAddress
}

// NewInetAddressLoopback creates a Address for the loopback address for family.
func NewInetAddressLoopback(family SocketFamily) *InetAddress {
	var _arg1 C.GSocketFamily // out
	var _cret *C.GInetAddress // in

	_arg1 = C.GSocketFamily(family)

	_cret = C.g_inet_address_new_loopback(_arg1)

	runtime.KeepAlive(family)

	var _inetAddress *InetAddress // out

	_inetAddress = wrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetAddress
}

// Equal checks if two Address instances are equal, e.g. the same address.
func (address *InetAddress) Equal(otherAddress *InetAddress) bool {
	var _arg0 *C.GInetAddress // out
	var _arg1 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(otherAddress.Native()))

	_cret = C.g_inet_address_equal(_arg0, _arg1)

	runtime.KeepAlive(address)
	runtime.KeepAlive(otherAddress)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Family gets address's family
func (address *InetAddress) Family() SocketFamily {
	var _arg0 *C.GInetAddress // out
	var _cret C.GSocketFamily // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_family(_arg0)

	runtime.KeepAlive(address)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// IsAny tests whether address is the "any" address for its family.
func (address *InetAddress) IsAny() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_any(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLinkLocal tests whether address is a link-local address (that is, if it
// identifies a host on a local network that is not connected to the Internet).
func (address *InetAddress) IsLinkLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_link_local(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsLoopback tests whether address is the loopback address for its family.
func (address *InetAddress) IsLoopback() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_loopback(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcGlobal tests whether address is a global multicast address.
func (address *InetAddress) IsMcGlobal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_mc_global(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcLinkLocal tests whether address is a link-local multicast address.
func (address *InetAddress) IsMcLinkLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_mc_link_local(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcNodeLocal tests whether address is a node-local multicast address.
func (address *InetAddress) IsMcNodeLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_mc_node_local(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcOrgLocal tests whether address is an organization-local multicast
// address.
func (address *InetAddress) IsMcOrgLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_mc_org_local(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMcSiteLocal tests whether address is a site-local multicast address.
func (address *InetAddress) IsMcSiteLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_mc_site_local(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMulticast tests whether address is a multicast address.
func (address *InetAddress) IsMulticast() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_multicast(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSiteLocal tests whether address is a site-local address such as 10.0.0.1
// (that is, the address identifies a host on a local network that can not be
// reached directly from the Internet, but which may have outgoing Internet
// connectivity via a NAT or firewall).
func (address *InetAddress) IsSiteLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_is_site_local(_arg0)

	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NativeSize gets the size of the native raw binary address for address. This
// is the size of the data that you get from g_inet_address_to_bytes().
func (address *InetAddress) NativeSize() uint {
	var _arg0 *C.GInetAddress // out
	var _cret C.gsize         // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_get_native_size(_arg0)

	runtime.KeepAlive(address)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// String converts address to string form.
func (address *InetAddress) String() string {
	var _arg0 *C.GInetAddress // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_to_string(_arg0)

	runtime.KeepAlive(address)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
