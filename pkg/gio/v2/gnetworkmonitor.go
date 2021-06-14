// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
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
		{T: externglib.Type(C.g_network_monitor_get_type()), F: marshalNetworkMonitor},
	})
}

// NetworkMonitorOverrider contains methods that are overridable. This
// interface is a subset of the interface NetworkMonitor.
type NetworkMonitorOverrider interface {
	// CanReach attempts to determine whether or not the host pointed to by
	// @connectable can be reached, without actually trying to connect to it.
	//
	// This may return true even when Monitor:network-available is false, if,
	// for example, @monitor can determine that @connectable refers to a host on
	// a local network.
	//
	// If @monitor believes that an attempt to connect to @connectable will
	// succeed, it will return true. Otherwise, it will return false and set
	// @error to an appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
	//
	// Note that although this does not attempt to connect to @connectable, it
	// may still block for a brief period of time (eg, trying to do multicast
	// DNS on the local network), so if you do not want to block, you should use
	// g_network_monitor_can_reach_async().
	CanReach(connectable SocketConnectable, cancellable Cancellable) error
	// CanReachFinish finishes an async network connectivity test. See
	// g_network_monitor_can_reach_async().
	CanReachFinish(result AsyncResult) error

	NetworkChanged(networkAvailable bool)
}

// NetworkMonitor provides an easy-to-use cross-platform API for monitoring
// network connectivity. On Linux, the available implementations are based on
// the kernel's netlink interface and on NetworkManager.
//
// There is also an implementation for use inside Flatpak sandboxes.
type NetworkMonitor interface {
	Initable
	NetworkMonitorOverrider

	// Connectivity gets a more detailed networking state than
	// g_network_monitor_get_network_available().
	//
	// If Monitor:network-available is false, then the connectivity state will
	// be G_NETWORK_CONNECTIVITY_LOCAL.
	//
	// If Monitor:network-available is true, then the connectivity state will be
	// G_NETWORK_CONNECTIVITY_FULL (if there is full Internet connectivity),
	// G_NETWORK_CONNECTIVITY_LIMITED (if the host has a default route, but
	// appears to be unable to actually reach the full Internet), or
	// G_NETWORK_CONNECTIVITY_PORTAL (if the host is trapped behind a "captive
	// portal" that requires some sort of login or acknowledgement before
	// allowing full Internet access).
	//
	// Note that in the case of G_NETWORK_CONNECTIVITY_LIMITED and
	// G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are
	// reachable but others are not. In this case, applications can attempt to
	// connect to remote servers, but should gracefully fall back to their
	// "offline" behavior if the connection attempt fails.
	Connectivity() NetworkConnectivity
	// NetworkAvailable checks if the network is available. "Available" here
	// means that the system has a default route available for at least one of
	// IPv4 or IPv6. It does not necessarily imply that the public Internet is
	// reachable. See Monitor:network-available for more details.
	NetworkAvailable() bool
	// NetworkMetered checks if the network is metered. See
	// Monitor:network-metered for more details.
	NetworkMetered() bool
}

// networkMonitor implements the NetworkMonitor interface.
type networkMonitor struct {
	Initable
}

var _ NetworkMonitor = (*networkMonitor)(nil)

// WrapNetworkMonitor wraps a GObject to a type that implements interface
// NetworkMonitor. It is primarily used internally.
func WrapNetworkMonitor(obj *externglib.Object) NetworkMonitor {
	return NetworkMonitor{
		Initable: WrapInitable(obj),
	}
}

func marshalNetworkMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkMonitor(obj), nil
}

// CanReach attempts to determine whether or not the host pointed to by
// @connectable can be reached, without actually trying to connect to it.
//
// This may return true even when Monitor:network-available is false, if,
// for example, @monitor can determine that @connectable refers to a host on
// a local network.
//
// If @monitor believes that an attempt to connect to @connectable will
// succeed, it will return true. Otherwise, it will return false and set
// @error to an appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
//
// Note that although this does not attempt to connect to @connectable, it
// may still block for a brief period of time (eg, trying to do multicast
// DNS on the local network), so if you do not want to block, you should use
// g_network_monitor_can_reach_async().
func (m networkMonitor) CanReach(connectable SocketConnectable, cancellable Cancellable) error {
	var _arg0 *C.GNetworkMonitor    // out
	var _arg1 *C.GSocketConnectable // out
	var _arg2 *C.GCancellable       // out

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var _cerr *C.GError // in

	C.g_network_monitor_can_reach(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// CanReachFinish finishes an async network connectivity test. See
// g_network_monitor_can_reach_async().
func (m networkMonitor) CanReachFinish(result AsyncResult) error {
	var _arg0 *C.GNetworkMonitor // out
	var _arg1 *C.GAsyncResult    // out

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cerr *C.GError // in

	C.g_network_monitor_can_reach_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Connectivity gets a more detailed networking state than
// g_network_monitor_get_network_available().
//
// If Monitor:network-available is false, then the connectivity state will
// be G_NETWORK_CONNECTIVITY_LOCAL.
//
// If Monitor:network-available is true, then the connectivity state will be
// G_NETWORK_CONNECTIVITY_FULL (if there is full Internet connectivity),
// G_NETWORK_CONNECTIVITY_LIMITED (if the host has a default route, but
// appears to be unable to actually reach the full Internet), or
// G_NETWORK_CONNECTIVITY_PORTAL (if the host is trapped behind a "captive
// portal" that requires some sort of login or acknowledgement before
// allowing full Internet access).
//
// Note that in the case of G_NETWORK_CONNECTIVITY_LIMITED and
// G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are
// reachable but others are not. In this case, applications can attempt to
// connect to remote servers, but should gracefully fall back to their
// "offline" behavior if the connection attempt fails.
func (m networkMonitor) Connectivity() NetworkConnectivity {
	var _arg0 *C.GNetworkMonitor // out

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.GNetworkConnectivity // in

	_cret = C.g_network_monitor_get_connectivity(_arg0)

	var _networkConnectivity NetworkConnectivity // out

	_networkConnectivity = NetworkConnectivity(_cret)

	return _networkConnectivity
}

// NetworkAvailable checks if the network is available. "Available" here
// means that the system has a default route available for at least one of
// IPv4 or IPv6. It does not necessarily imply that the public Internet is
// reachable. See Monitor:network-available for more details.
func (m networkMonitor) NetworkAvailable() bool {
	var _arg0 *C.GNetworkMonitor // out

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.gboolean // in

	_cret = C.g_network_monitor_get_network_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NetworkMetered checks if the network is metered. See
// Monitor:network-metered for more details.
func (m networkMonitor) NetworkMetered() bool {
	var _arg0 *C.GNetworkMonitor // out

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.gboolean // in

	_cret = C.g_network_monitor_get_network_metered(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
