// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_network_monitor_get_type()), F: marshalNetworkMonitorrer},
	})
}

// NETWORK_MONITOR_EXTENSION_POINT_NAME: extension point for network status
// monitoring functionality. See [Extending GIO][extending-gio].
const NETWORK_MONITOR_EXTENSION_POINT_NAME = "gio-network-monitor"

// NetworkMonitorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type NetworkMonitorOverrider interface {
	// CanReach attempts to determine whether or not the host pointed to by
	// connectable can be reached, without actually trying to connect to it.
	//
	// This may return TRUE even when Monitor:network-available is FALSE, if,
	// for example, monitor can determine that connectable refers to a host on a
	// local network.
	//
	// If monitor believes that an attempt to connect to connectable will
	// succeed, it will return TRUE. Otherwise, it will return FALSE and set
	// error to an appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
	//
	// Note that although this does not attempt to connect to connectable, it
	// may still block for a brief period of time (eg, trying to do multicast
	// DNS on the local network), so if you do not want to block, you should use
	// g_network_monitor_can_reach_async().
	CanReach(ctx context.Context, connectable SocketConnectabler) error
	// CanReachAsync: asynchronously attempts to determine whether or not the
	// host pointed to by connectable can be reached, without actually trying to
	// connect to it.
	//
	// For more details, see g_network_monitor_can_reach().
	//
	// When the operation is finished, callback will be called. You can then
	// call g_network_monitor_can_reach_finish() to get the result of the
	// operation.
	CanReachAsync(ctx context.Context, connectable SocketConnectabler, callback AsyncReadyCallback)
	// CanReachFinish finishes an async network connectivity test. See
	// g_network_monitor_can_reach_async().
	CanReachFinish(result AsyncResulter) error
	NetworkChanged(networkAvailable bool)
}

// NetworkMonitor provides an easy-to-use cross-platform API for monitoring
// network connectivity. On Linux, the available implementations are based on
// the kernel's netlink interface and on NetworkManager.
//
// There is also an implementation for use inside Flatpak sandboxes.
type NetworkMonitor struct {
	Initable
}

// NetworkMonitorrer describes NetworkMonitor's abstract methods.
type NetworkMonitorrer interface {
	externglib.Objector

	// CanReach attempts to determine whether or not the host pointed to by
	// connectable can be reached, without actually trying to connect to it.
	CanReach(ctx context.Context, connectable SocketConnectabler) error
	// CanReachAsync: asynchronously attempts to determine whether or not the
	// host pointed to by connectable can be reached, without actually trying to
	// connect to it.
	CanReachAsync(ctx context.Context, connectable SocketConnectabler, callback AsyncReadyCallback)
	// CanReachFinish finishes an async network connectivity test.
	CanReachFinish(result AsyncResulter) error
	// Connectivity gets a more detailed networking state than
	// g_network_monitor_get_network_available().
	Connectivity() NetworkConnectivity
	// NetworkAvailable checks if the network is available.
	NetworkAvailable() bool
	// NetworkMetered checks if the network is metered.
	NetworkMetered() bool
}

var _ NetworkMonitorrer = (*NetworkMonitor)(nil)

func wrapNetworkMonitor(obj *externglib.Object) *NetworkMonitor {
	return &NetworkMonitor{
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalNetworkMonitorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNetworkMonitor(obj), nil
}

// CanReach attempts to determine whether or not the host pointed to by
// connectable can be reached, without actually trying to connect to it.
//
// This may return TRUE even when Monitor:network-available is FALSE, if, for
// example, monitor can determine that connectable refers to a host on a local
// network.
//
// If monitor believes that an attempt to connect to connectable will succeed,
// it will return TRUE. Otherwise, it will return FALSE and set error to an
// appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
//
// Note that although this does not attempt to connect to connectable, it may
// still block for a brief period of time (eg, trying to do multicast DNS on the
// local network), so if you do not want to block, you should use
// g_network_monitor_can_reach_async().
func (monitor *NetworkMonitor) CanReach(ctx context.Context, connectable SocketConnectabler) error {
	var _arg0 *C.GNetworkMonitor    // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GSocketConnectable // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(monitor.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))

	C.g_network_monitor_can_reach(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CanReachAsync: asynchronously attempts to determine whether or not the host
// pointed to by connectable can be reached, without actually trying to connect
// to it.
//
// For more details, see g_network_monitor_can_reach().
//
// When the operation is finished, callback will be called. You can then call
// g_network_monitor_can_reach_finish() to get the result of the operation.
func (monitor *NetworkMonitor) CanReachAsync(ctx context.Context, connectable SocketConnectabler, callback AsyncReadyCallback) {
	var _arg0 *C.GNetworkMonitor    // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GSocketConnectable // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(monitor.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_network_monitor_can_reach_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connectable)
	runtime.KeepAlive(callback)
}

// CanReachFinish finishes an async network connectivity test. See
// g_network_monitor_can_reach_async().
func (monitor *NetworkMonitor) CanReachFinish(result AsyncResulter) error {
	var _arg0 *C.GNetworkMonitor // out
	var _arg1 *C.GAsyncResult    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(monitor.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_network_monitor_can_reach_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Connectivity gets a more detailed networking state than
// g_network_monitor_get_network_available().
//
// If Monitor:network-available is FALSE, then the connectivity state will be
// G_NETWORK_CONNECTIVITY_LOCAL.
//
// If Monitor:network-available is TRUE, then the connectivity state will be
// G_NETWORK_CONNECTIVITY_FULL (if there is full Internet connectivity),
// G_NETWORK_CONNECTIVITY_LIMITED (if the host has a default route, but appears
// to be unable to actually reach the full Internet), or
// G_NETWORK_CONNECTIVITY_PORTAL (if the host is trapped behind a "captive
// portal" that requires some sort of login or acknowledgement before allowing
// full Internet access).
//
// Note that in the case of G_NETWORK_CONNECTIVITY_LIMITED and
// G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are reachable
// but others are not. In this case, applications can attempt to connect to
// remote servers, but should gracefully fall back to their "offline" behavior
// if the connection attempt fails.
func (monitor *NetworkMonitor) Connectivity() NetworkConnectivity {
	var _arg0 *C.GNetworkMonitor     // out
	var _cret C.GNetworkConnectivity // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_network_monitor_get_connectivity(_arg0)

	runtime.KeepAlive(monitor)

	var _networkConnectivity NetworkConnectivity // out

	_networkConnectivity = NetworkConnectivity(_cret)

	return _networkConnectivity
}

// NetworkAvailable checks if the network is available. "Available" here means
// that the system has a default route available for at least one of IPv4 or
// IPv6. It does not necessarily imply that the public Internet is reachable.
// See Monitor:network-available for more details.
func (monitor *NetworkMonitor) NetworkAvailable() bool {
	var _arg0 *C.GNetworkMonitor // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_network_monitor_get_network_available(_arg0)

	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NetworkMetered checks if the network is metered. See Monitor:network-metered
// for more details.
func (monitor *NetworkMonitor) NetworkMetered() bool {
	var _arg0 *C.GNetworkMonitor // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_network_monitor_get_network_metered(_arg0)

	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NetworkMonitorGetDefault gets the default Monitor for the system.
func NetworkMonitorGetDefault() NetworkMonitorrer {
	var _cret *C.GNetworkMonitor // in

	_cret = C.g_network_monitor_get_default()

	var _networkMonitor NetworkMonitorrer // out

	_networkMonitor = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(NetworkMonitorrer)

	return _networkMonitor
}
