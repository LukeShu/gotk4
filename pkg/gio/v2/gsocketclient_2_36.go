// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// ProxyResolver gets the Resolver being used by client. Normally, this will be
// the resolver returned by g_proxy_resolver_get_default(), but you can override
// it with g_socket_client_set_proxy_resolver().
//
// The function returns the following values:
//
//    - proxyResolver being used by client.
//
func (client *SocketClient) ProxyResolver() *ProxyResolver {
	var _arg0 *C.GSocketClient  // out
	var _cret *C.GProxyResolver // in

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(coreglib.InternObject(client).Native()))

	_cret = C.g_socket_client_get_proxy_resolver(_arg0)
	runtime.KeepAlive(client)

	var _proxyResolver *ProxyResolver // out

	_proxyResolver = wrapProxyResolver(coreglib.Take(unsafe.Pointer(_cret)))

	return _proxyResolver
}

// SetProxyResolver overrides the Resolver used by client. You can call this if
// you want to use specific proxies, rather than using the system default proxy
// settings.
//
// Note that whether or not the proxy resolver is actually used depends on the
// setting of Client:enable-proxy, which is not changed by this function (but
// which is TRUE by default).
//
// The function takes the following parameters:
//
//    - proxyResolver (optional) or NULL for the default.
//
func (client *SocketClient) SetProxyResolver(proxyResolver ProxyResolverer) {
	var _arg0 *C.GSocketClient  // out
	var _arg1 *C.GProxyResolver // out

	_arg0 = (*C.GSocketClient)(unsafe.Pointer(coreglib.InternObject(client).Native()))
	if proxyResolver != nil {
		_arg1 = (*C.GProxyResolver)(unsafe.Pointer(coreglib.InternObject(proxyResolver).Native()))
	}

	C.g_socket_client_set_proxy_resolver(_arg0, _arg1)
	runtime.KeepAlive(client)
	runtime.KeepAlive(proxyResolver)
}
