// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
		{T: externglib.Type(C.g_dbus_proxy_get_type()), F: marshalDBusProxy},
	})
}

// DBusProxyOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusProxyOverrider interface {
	GSignal(senderName string, signalName string, parameters *glib.Variant)
}

// DBusProxy is a base class used for proxies to access a D-Bus interface on a
// remote object. A BusProxy can be constructed for both well-known and unique
// names.
//
// By default, BusProxy will cache all properties (and listen to changes) of the
// remote object, and proxy all signals that get emitted. This behaviour can be
// changed by passing suitable BusProxyFlags when the proxy is created. If the
// proxy is for a well-known name, the property cache is flushed when the name
// owner vanishes and reloaded when a name owner appears.
//
// The unique name owner of the proxy's name is tracked and can be read from
// BusProxy:g-name-owner. Connect to the #GObject::notify signal to get notified
// of changes. Additionally, only signals and property changes emitted from the
// current name owner are considered and calls are always sent to the current
// name owner. This avoids a number of race conditions when the name is lost by
// one owner and claimed by another. However, if no name owner currently exists,
// then calls will be sent to the well-known name which may result in the
// message bus launching an owner (unless G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START
// is set).
//
// The generic BusProxy::g-properties-changed and BusProxy::g-signal signals are
// not very convenient to work with. Therefore, the recommended way of working
// with proxies is to subclass BusProxy, and have more natural properties and
// signals in your derived class. This [example][gdbus-example-gdbus-codegen]
// shows how this can easily be done using the [gdbus-codegen][gdbus-codegen]
// tool.
//
// A BusProxy instance can be used from multiple threads but note that all
// signals (e.g. BusProxy::g-signal, BusProxy::g-properties-changed and
// #GObject::notify) are emitted in the [thread-default main
// context][g-main-context-push-thread-default] of the thread where the instance
// was constructed.
//
// An example using a proxy for a well-known name can be found in
// gdbus-example-watch-proxy.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-watch-proxy.c)
type DBusProxy interface {
	gextras.Objector

	// CallFinish finishes an operation started with g_dbus_proxy_call().
	CallFinish(res AsyncResult) (*glib.Variant, error)
	// CallWithUnixFdListFinish finishes an operation started with
	// g_dbus_proxy_call_with_unix_fd_list().
	CallWithUnixFdListFinish(res AsyncResult) (*UnixFDListClass, *glib.Variant, error)
	// CachedProperty looks up the value for a property from the cache. This
	// call does no blocking IO.
	//
	// If @proxy has an expected interface (see BusProxy:g-interface-info) and
	// @property_name is referenced by it, then @value is checked against the
	// type of the property.
	CachedProperty(propertyName string) *glib.Variant
	// CachedPropertyNames gets the names of all cached properties on @proxy.
	CachedPropertyNames() []string
	// Connection gets the connection @proxy is for.
	Connection() *DBusConnectionClass
	// DefaultTimeout gets the timeout to use if -1 (specifying default timeout)
	// is passed as @timeout_msec in the g_dbus_proxy_call() and
	// g_dbus_proxy_call_sync() functions.
	//
	// See the BusProxy:g-default-timeout property for more details.
	DefaultTimeout() int
	// Flags gets the flags that @proxy was constructed with.
	Flags() DBusProxyFlags
	// InterfaceInfo returns the BusInterfaceInfo, if any, specifying the
	// interface that @proxy conforms to. See the BusProxy:g-interface-info
	// property for more details.
	InterfaceInfo() *DBusInterfaceInfo
	// InterfaceName gets the D-Bus interface name @proxy is for.
	InterfaceName() string
	// Name gets the name that @proxy was constructed for.
	Name() string
	// NameOwner: the unique name that owns the name that @proxy is for or nil
	// if no-one currently owns that name. You may connect to the
	// #GObject::notify signal to track changes to the BusProxy:g-name-owner
	// property.
	NameOwner() string
	// ObjectPath gets the object path @proxy is for.
	ObjectPath() string
	// SetCachedProperty: if @value is not nil, sets the cached value for the
	// property with name @property_name to the value in @value.
	//
	// If @value is nil, then the cached value is removed from the property
	// cache.
	//
	// If @proxy has an expected interface (see BusProxy:g-interface-info) and
	// @property_name is referenced by it, then @value is checked against the
	// type of the property.
	//
	// If the @value #GVariant is floating, it is consumed. This allows
	// convenient 'inline' use of g_variant_new(), e.g.
	//
	//    g_dbus_proxy_set_cached_property (proxy,
	//                                      "SomeProperty",
	//                                      g_variant_new ("(si)",
	//                                                    "A String",
	//                                                    42));
	//
	// Normally you will not need to use this method since @proxy is tracking
	// changes using the `org.freedesktop.DBus.Properties.PropertiesChanged`
	// D-Bus signal. However, for performance reasons an object may decide to
	// not use this signal for some properties and instead use a proprietary
	// out-of-band mechanism to transmit changes.
	//
	// As a concrete example, consider an object with a property
	// `ChatroomParticipants` which is an array of strings. Instead of
	// transmitting the same (long) array every time the property changes, it is
	// more efficient to only transmit the delta using e.g. signals
	// `ChatroomParticipantJoined(String name)` and
	// `ChatroomParticipantParted(String name)`.
	SetCachedProperty(propertyName string, value *glib.Variant)
	// SetDefaultTimeout sets the timeout to use if -1 (specifying default
	// timeout) is passed as @timeout_msec in the g_dbus_proxy_call() and
	// g_dbus_proxy_call_sync() functions.
	//
	// See the BusProxy:g-default-timeout property for more details.
	SetDefaultTimeout(timeoutMsec int)
	// SetInterfaceInfo: ensure that interactions with @proxy conform to the
	// given interface. See the BusProxy:g-interface-info property for more
	// details.
	SetInterfaceInfo(info *DBusInterfaceInfo)
}

// DBusProxyClass implements the DBusProxy interface.
type DBusProxyClass struct {
	*externglib.Object
	AsyncInitableInterface
	DBusInterfaceInterface
	InitableInterface
}

var _ DBusProxy = (*DBusProxyClass)(nil)

func wrapDBusProxy(obj *externglib.Object) DBusProxy {
	return &DBusProxyClass{
		Object: obj,
		AsyncInitableInterface: AsyncInitableInterface{
			Object: obj,
		},
		DBusInterfaceInterface: DBusInterfaceInterface{
			Object: obj,
		},
		InitableInterface: InitableInterface{
			Object: obj,
		},
	}
}

func marshalDBusProxy(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusProxy(obj), nil
}

// NewDBusProxyFinish finishes creating a BusProxy.
func NewDBusProxyFinish(res AsyncResult) (*DBusProxyClass, error) {
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GDBusProxy   // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&res).Native()))

	_cret = C.g_dbus_proxy_new_finish(_arg1, &_cerr)

	var _dBusProxy *DBusProxyClass // out
	var _goerr error               // out

	_dBusProxy = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*DBusProxyClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusProxy, _goerr
}

// NewDBusProxyForBusFinish finishes creating a BusProxy.
func NewDBusProxyForBusFinish(res AsyncResult) (*DBusProxyClass, error) {
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GDBusProxy   // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&res).Native()))

	_cret = C.g_dbus_proxy_new_for_bus_finish(_arg1, &_cerr)

	var _dBusProxy *DBusProxyClass // out
	var _goerr error               // out

	_dBusProxy = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*DBusProxyClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusProxy, _goerr
}

// CallFinish finishes an operation started with g_dbus_proxy_call().
func (p *DBusProxyClass) CallFinish(res AsyncResult) (*glib.Variant, error) {
	var _arg0 *C.GDBusProxy   // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GVariant     // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&res).Native()))

	_cret = C.g_dbus_proxy_call_finish(_arg0, _arg1, &_cerr)

	var _variant *glib.Variant // out
	var _goerr error           // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _variant, _goerr
}

// CallWithUnixFdListFinish finishes an operation started with
// g_dbus_proxy_call_with_unix_fd_list().
func (p *DBusProxyClass) CallWithUnixFdListFinish(res AsyncResult) (*UnixFDListClass, *glib.Variant, error) {
	var _arg0 *C.GDBusProxy   // out
	var _arg1 *C.GUnixFDList  // in
	var _arg2 *C.GAsyncResult // out
	var _cret *C.GVariant     // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))
	_arg2 = (*C.GAsyncResult)(unsafe.Pointer((&res).Native()))

	_cret = C.g_dbus_proxy_call_with_unix_fd_list_finish(_arg0, &_arg1, _arg2, &_cerr)

	var _outFdList *UnixFDListClass // out
	var _variant *glib.Variant      // out
	var _goerr error                // out

	_outFdList = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_arg1))).(*UnixFDListClass)
	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _outFdList, _variant, _goerr
}

// CachedProperty looks up the value for a property from the cache. This call
// does no blocking IO.
//
// If @proxy has an expected interface (see BusProxy:g-interface-info) and
// @property_name is referenced by it, then @value is checked against the type
// of the property.
func (p *DBusProxyClass) CachedProperty(propertyName string) *glib.Variant {
	var _arg0 *C.GDBusProxy // out
	var _arg1 *C.gchar      // out
	var _cret *C.GVariant   // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))
	_arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_proxy_get_cached_property(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// CachedPropertyNames gets the names of all cached properties on @proxy.
func (p *DBusProxyClass) CachedPropertyNames() []string {
	var _arg0 *C.GDBusProxy // out
	var _cret **C.gchar

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_cached_property_names(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Connection gets the connection @proxy is for.
func (p *DBusProxyClass) Connection() *DBusConnectionClass {
	var _arg0 *C.GDBusProxy      // out
	var _cret *C.GDBusConnection // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_connection(_arg0)

	var _dBusConnection *DBusConnectionClass // out

	_dBusConnection = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*DBusConnectionClass)

	return _dBusConnection
}

// DefaultTimeout gets the timeout to use if -1 (specifying default timeout) is
// passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
func (p *DBusProxyClass) DefaultTimeout() int {
	var _arg0 *C.GDBusProxy // out
	var _cret C.gint        // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_default_timeout(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Flags gets the flags that @proxy was constructed with.
func (p *DBusProxyClass) Flags() DBusProxyFlags {
	var _arg0 *C.GDBusProxy     // out
	var _cret C.GDBusProxyFlags // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_flags(_arg0)

	var _dBusProxyFlags DBusProxyFlags // out

	_dBusProxyFlags = (DBusProxyFlags)(_cret)

	return _dBusProxyFlags
}

// InterfaceInfo returns the BusInterfaceInfo, if any, specifying the interface
// that @proxy conforms to. See the BusProxy:g-interface-info property for more
// details.
func (p *DBusProxyClass) InterfaceInfo() *DBusInterfaceInfo {
	var _arg0 *C.GDBusProxy         // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_interface_info(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// InterfaceName gets the D-Bus interface name @proxy is for.
func (p *DBusProxyClass) InterfaceName() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_interface_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Name gets the name that @proxy was constructed for.
func (p *DBusProxyClass) Name() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// NameOwner: the unique name that owns the name that @proxy is for or nil if
// no-one currently owns that name. You may connect to the #GObject::notify
// signal to track changes to the BusProxy:g-name-owner property.
func (p *DBusProxyClass) NameOwner() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_name_owner(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ObjectPath gets the object path @proxy is for.
func (p *DBusProxyClass) ObjectPath() string {
	var _arg0 *C.GDBusProxy // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))

	_cret = C.g_dbus_proxy_get_object_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SetCachedProperty: if @value is not nil, sets the cached value for the
// property with name @property_name to the value in @value.
//
// If @value is nil, then the cached value is removed from the property cache.
//
// If @proxy has an expected interface (see BusProxy:g-interface-info) and
// @property_name is referenced by it, then @value is checked against the type
// of the property.
//
// If the @value #GVariant is floating, it is consumed. This allows convenient
// 'inline' use of g_variant_new(), e.g.
//
//    g_dbus_proxy_set_cached_property (proxy,
//                                      "SomeProperty",
//                                      g_variant_new ("(si)",
//                                                    "A String",
//                                                    42));
//
// Normally you will not need to use this method since @proxy is tracking
// changes using the `org.freedesktop.DBus.Properties.PropertiesChanged` D-Bus
// signal. However, for performance reasons an object may decide to not use this
// signal for some properties and instead use a proprietary out-of-band
// mechanism to transmit changes.
//
// As a concrete example, consider an object with a property
// `ChatroomParticipants` which is an array of strings. Instead of transmitting
// the same (long) array every time the property changes, it is more efficient
// to only transmit the delta using e.g. signals
// `ChatroomParticipantJoined(String name)` and
// `ChatroomParticipantParted(String name)`.
func (p *DBusProxyClass) SetCachedProperty(propertyName string, value *glib.Variant) {
	var _arg0 *C.GDBusProxy // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GVariant   // out

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))
	_arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(unsafe.Pointer(value))

	C.g_dbus_proxy_set_cached_property(_arg0, _arg1, _arg2)
}

// SetDefaultTimeout sets the timeout to use if -1 (specifying default timeout)
// is passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
func (p *DBusProxyClass) SetDefaultTimeout(timeoutMsec int) {
	var _arg0 *C.GDBusProxy // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))
	_arg1 = C.gint(timeoutMsec)

	C.g_dbus_proxy_set_default_timeout(_arg0, _arg1)
}

// SetInterfaceInfo: ensure that interactions with @proxy conform to the given
// interface. See the BusProxy:g-interface-info property for more details.
func (p *DBusProxyClass) SetInterfaceInfo(info *DBusInterfaceInfo) {
	var _arg0 *C.GDBusProxy         // out
	var _arg1 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusProxy)(unsafe.Pointer((&p).Native()))
	_arg1 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	C.g_dbus_proxy_set_interface_info(_arg0, _arg1)
}
