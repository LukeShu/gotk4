// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gio2_ProxyResolverInterface_is_supported(void*);
// extern gchar** _gotk4_gio2_ProxyResolverInterface_lookup(void*, void*, void*, GError**);
// extern gchar** _gotk4_gio2_ProxyResolverInterface_lookup_finish(void*, void*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// GTypeProxyResolver returns the GType for the type ProxyResolver.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeProxyResolver() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "ProxyResolver").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalProxyResolver)
	return gtype
}

// PROXY_RESOLVER_EXTENSION_POINT_NAME: extension point for proxy resolving
// functionality. See [Extending GIO][extending-gio].
const PROXY_RESOLVER_EXTENSION_POINT_NAME = "gio-proxy-resolver"

// ProxyResolver provides synchronous and asynchronous network proxy resolution.
// Resolver is used within Client through the method
// g_socket_connectable_proxy_enumerate().
//
// Implementations of Resolver based on libproxy and GNOME settings can be found
// in glib-networking. GIO comes with an implementation for use inside Flatpak
// portals.
//
// ProxyResolver wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ProxyResolver struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ProxyResolver)(nil)
)

// ProxyResolverer describes ProxyResolver's interface methods.
type ProxyResolverer interface {
	coreglib.Objector

	// IsSupported checks if resolver can be used on this system.
	IsSupported() bool
	// Lookup looks into the system proxy configuration to determine what proxy,
	// if any, to use to connect to uri.
	Lookup(ctx context.Context, uri string) ([]string, error)
	// LookupAsync asynchronous lookup of proxy.
	LookupAsync(ctx context.Context, uri string, callback AsyncReadyCallback)
	// LookupFinish: call this function to obtain the array of proxy URIs when
	// g_proxy_resolver_lookup_async() is complete.
	LookupFinish(result AsyncResulter) ([]string, error)
}

var _ ProxyResolverer = (*ProxyResolver)(nil)

func wrapProxyResolver(obj *coreglib.Object) *ProxyResolver {
	return &ProxyResolver{
		Object: obj,
	}
}

func marshalProxyResolver(p uintptr) (interface{}, error) {
	return wrapProxyResolver(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// IsSupported checks if resolver can be used on this system. (This is used
// internally; g_proxy_resolver_get_default() will only return a proxy resolver
// that returns TRUE for this method.).
//
// The function returns the following values:
//
//    - ok: TRUE if resolver is supported.
//
func (resolver *ProxyResolver) IsSupported() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resolver)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Lookup looks into the system proxy configuration to determine what proxy, if
// any, to use to connect to uri. The returned proxy URIs are of the form
// <protocol>://[user[:password]@]host:port or direct://, where <protocol> could
// be http, rtsp, socks or other proxying protocol.
//
// If you don't know what network protocol is being used on the socket, you
// should use none as the URI protocol. In this case, the resolver might still
// return a generic proxy type (such as SOCKS), but would not return
// protocol-specific proxy types (such as http).
//
// direct:// is used when no proxy is needed. Direct connection should not be
// attempted unless it is part of the returned array of proxies.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - uri: URI representing the destination to connect to.
//
// The function returns the following values:
//
//    - utf8s: a NULL-terminated array of proxy URIs. Must be freed with
//      g_strfreev().
//
func (resolver *ProxyResolver) Lookup(ctx context.Context, uri string) ([]string, error) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_args[1]))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)

	var _utf8s []string // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8s, _goerr
}

// LookupAsync asynchronous lookup of proxy. See g_proxy_resolver_lookup() for
// more details.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - uri: URI representing the destination to connect to.
//    - callback (optional) to call after resolution completes.
//
func (resolver *ProxyResolver) LookupAsync(ctx context.Context, uri string, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_args[1]))
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(callback)
}

// LookupFinish: call this function to obtain the array of proxy URIs when
// g_proxy_resolver_lookup_async() is complete. See g_proxy_resolver_lookup()
// for more details.
//
// The function takes the following parameters:
//
//    - result passed to your ReadyCallback.
//
// The function returns the following values:
//
//    - utf8s: a NULL-terminated array of proxy URIs. Must be freed with
//      g_strfreev().
//
func (resolver *ProxyResolver) LookupFinish(result AsyncResulter) ([]string, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resolver)
	runtime.KeepAlive(result)

	var _utf8s []string // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.void
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8s, _goerr
}

// ProxyResolverGetDefault gets the default Resolver for the system.
//
// The function returns the following values:
//
//    - proxyResolver: default Resolver, which will be a dummy object if no proxy
//      resolver is available.
//
func ProxyResolverGetDefault() *ProxyResolver {
	_gret := girepository.MustFind("Gio", "get_default").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _proxyResolver *ProxyResolver // out

	_proxyResolver = wrapProxyResolver(coreglib.Take(unsafe.Pointer(_cret)))

	return _proxyResolver
}
