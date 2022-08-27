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
//
// The function returns the following values:
//
//    - networkConnectivity: network connectivity state.
//
func (monitor *NetworkMonitor) Connectivity() NetworkConnectivity {
	var _arg0 *C.GNetworkMonitor     // out
	var _cret C.GNetworkConnectivity // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))

	_cret = C.g_network_monitor_get_connectivity(_arg0)
	runtime.KeepAlive(monitor)

	var _networkConnectivity NetworkConnectivity // out

	_networkConnectivity = NetworkConnectivity(_cret)

	return _networkConnectivity
}
