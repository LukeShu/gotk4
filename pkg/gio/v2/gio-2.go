// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_DBusAuthObserver_ConnectAllowMechanism(gpointer, gchar*, guintptr);
// extern gboolean _gotk4_gio2_DBusAuthObserver_ConnectAuthorizeAuthenticatedPeer(gpointer, GIOStream*, GCredentials*, guintptr);
// extern gboolean _gotk4_gio2_DBusServer_ConnectNewConnection(gpointer, GDBusConnection*, guintptr);
// extern void _gotk4_gio2_AppInfoMonitor_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gio2_DBusConnection_ConnectClosed(gpointer, gboolean, GError*, guintptr);
// extern void _gotk4_gio2_SimpleAction_ConnectActivate(gpointer, GVariant*, guintptr);
// extern void _gotk4_gio2_SimpleAction_ConnectChangeState(gpointer, GVariant*, guintptr);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_app_info_monitor_get_type()), F: marshalAppInfoMonitorrer},
		{T: externglib.Type(C.g_bytes_icon_get_type()), F: marshalBytesIconner},
		{T: externglib.Type(C.g_dbus_action_group_get_type()), F: marshalDBusActionGrouper},
		{T: externglib.Type(C.g_dbus_auth_observer_get_type()), F: marshalDBusAuthObserverer},
		{T: externglib.Type(C.g_dbus_connection_get_type()), F: marshalDBusConnectioner},
		{T: externglib.Type(C.g_dbus_menu_model_get_type()), F: marshalDBusMenuModeller},
		{T: externglib.Type(C.g_dbus_message_get_type()), F: marshalDBusMessager},
		{T: externglib.Type(C.g_dbus_method_invocation_get_type()), F: marshalDBusMethodInvocationer},
		{T: externglib.Type(C.g_dbus_server_get_type()), F: marshalDBusServerer},
		{T: externglib.Type(C.g_menu_get_type()), F: marshalMenuer},
		{T: externglib.Type(C.g_menu_item_get_type()), F: marshalMenuItemmer},
		{T: externglib.Type(C.g_notification_get_type()), F: marshalNotificationer},
		{T: externglib.Type(C.g_property_action_get_type()), F: marshalPropertyActioner},
		{T: externglib.Type(C.g_simple_action_get_type()), F: marshalSimpleActioner},
		{T: externglib.Type(C.g_simple_io_stream_get_type()), F: marshalSimpleIOStreamer},
		{T: externglib.Type(C.g_simple_permission_get_type()), F: marshalSimplePermissioner},
		{T: externglib.Type(C.g_test_dbus_get_type()), F: marshalTestDBusser},
	})
}

// ResolverErrorQuark gets the #GResolver Error Quark.
//
// The function returns the following values:
//
//    - quark: #GQuark.
//
func ResolverErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.g_resolver_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// ResourceErrorQuark gets the #GResource Error Quark.
//
// The function returns the following values:
//
//    - quark: #GQuark.
//
func ResourceErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.g_resource_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// TLSChannelBindingErrorQuark gets the TLS channel binding error quark.
//
// The function returns the following values:
//
//    - quark: #GQuark.
//
func TLSChannelBindingErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.g_tls_channel_binding_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// TLSErrorQuark gets the TLS error quark.
//
// The function returns the following values:
//
//    - quark: #GQuark.
//
func TLSErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.g_tls_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// IOErrorQuark gets the GIO Error Quark.
//
// The function returns the following values:
//
//    - quark: #GQuark.
//
func IOErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.g_io_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// AppInfoMonitor is a very simple object used for monitoring the app info
// database for changes (ie: newly installed or removed applications).
//
// Call g_app_info_monitor_get() to get a InfoMonitor and connect to the
// "changed" signal.
//
// In the usual case, applications should try to make note of the change (doing
// things like invalidating caches) but not act on it. In particular,
// applications should avoid making calls to Info APIs in response to the change
// signal, deferring these until the time that the data is actually required.
// The exception to this case is when application information is actually being
// displayed on the screen (eg: during a search or when the list of all
// applications is shown). The reason for this is that changes to the list of
// installed applications often come in groups (like during system updates) and
// rescanning the list on every change is pointless and expensive.
type AppInfoMonitor struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*AppInfoMonitor)(nil)
)

func wrapAppInfoMonitor(obj *externglib.Object) *AppInfoMonitor {
	return &AppInfoMonitor{
		Object: obj,
	}
}

func marshalAppInfoMonitorrer(p uintptr) (interface{}, error) {
	return wrapAppInfoMonitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_AppInfoMonitor_ConnectChanged
func _gotk4_gio2_AppInfoMonitor_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged: signal emitted when the app info database for changes (ie:
// newly installed or removed applications).
func (v *AppInfoMonitor) ConnectChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gio2_AppInfoMonitor_ConnectChanged), f)
}

// BytesIcon specifies an image held in memory in a common format (usually png)
// to be used as icon.
type BytesIcon struct {
	_ [0]func() // equal guard
	*externglib.Object

	LoadableIcon
}

var (
	_ externglib.Objector = (*BytesIcon)(nil)
)

func wrapBytesIcon(obj *externglib.Object) *BytesIcon {
	return &BytesIcon{
		Object: obj,
		LoadableIcon: LoadableIcon{
			Icon: Icon{
				Object: obj,
			},
		},
	}
}

func marshalBytesIconner(p uintptr) (interface{}, error) {
	return wrapBytesIcon(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusActionGroup is an implementation of the Group interface that can be used
// as a proxy for an action group that is exported over D-Bus with
// g_dbus_connection_export_action_group().
type DBusActionGroup struct {
	_ [0]func() // equal guard
	*externglib.Object

	RemoteActionGroup
}

var (
	_ externglib.Objector = (*DBusActionGroup)(nil)
)

func wrapDBusActionGroup(obj *externglib.Object) *DBusActionGroup {
	return &DBusActionGroup{
		Object: obj,
		RemoteActionGroup: RemoteActionGroup{
			ActionGroup: ActionGroup{
				Object: obj,
			},
		},
	}
}

func marshalDBusActionGrouper(p uintptr) (interface{}, error) {
	return wrapDBusActionGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusAuthObserver type provides a mechanism for participating in how a
// BusServer (or a BusConnection) authenticates remote peers. Simply instantiate
// a BusAuthObserver and connect to the signals you are interested in. Note that
// new signals may be added in the future
//
//
// Controlling Authentication Mechanisms
//
// By default, a BusServer or server-side BusConnection will allow any
// authentication mechanism to be used. If you only want to allow D-Bus
// connections with the EXTERNAL mechanism, which makes use of credentials
// passing and is the recommended mechanism for modern Unix platforms such as
// Linux and the BSD family, you would use a signal handler like this:
//
//    static gboolean
//    on_authorize_authenticated_peer (GDBusAuthObserver *observer,
//                                     GIOStream         *stream,
//                                     GCredentials      *credentials,
//                                     gpointer           user_data)
//    {
//      gboolean authorized;
//
//      authorized = FALSE;
//      if (credentials != NULL)
//        {
//          GCredentials *own_credentials;
//          own_credentials = g_credentials_new ();
//          if (g_credentials_is_same_user (credentials, own_credentials, NULL))
//            authorized = TRUE;
//          g_object_unref (own_credentials);
//        }
//
//      return authorized;
//    }.
type DBusAuthObserver struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DBusAuthObserver)(nil)
)

func wrapDBusAuthObserver(obj *externglib.Object) *DBusAuthObserver {
	return &DBusAuthObserver{
		Object: obj,
	}
}

func marshalDBusAuthObserverer(p uintptr) (interface{}, error) {
	return wrapDBusAuthObserver(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusAuthObserver_ConnectAllowMechanism
func _gotk4_gio2_DBusAuthObserver_ConnectAllowMechanism(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) (cret C.gboolean) {
	var f func(mechanism string) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(mechanism string) (ok bool))
	}

	var _mechanism string // out

	_mechanism = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := f(_mechanism)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectAllowMechanism: emitted to check if mechanism is allowed to be used.
func (observer *DBusAuthObserver) ConnectAllowMechanism(f func(mechanism string) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(observer, "allow-mechanism", false, unsafe.Pointer(C._gotk4_gio2_DBusAuthObserver_ConnectAllowMechanism), f)
}

//export _gotk4_gio2_DBusAuthObserver_ConnectAuthorizeAuthenticatedPeer
func _gotk4_gio2_DBusAuthObserver_ConnectAuthorizeAuthenticatedPeer(arg0 C.gpointer, arg1 *C.GIOStream, arg2 *C.GCredentials, arg3 C.guintptr) (cret C.gboolean) {
	var f func(stream IOStreamer, credentials *Credentials) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(stream IOStreamer, credentials *Credentials) (ok bool))
	}

	var _stream IOStreamer        // out
	var _credentials *Credentials // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_stream = rv
	}
	if arg2 != nil {
		_credentials = wrapCredentials(externglib.Take(unsafe.Pointer(arg2)))
	}

	ok := f(_stream, _credentials)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectAuthorizeAuthenticatedPeer: emitted to check if a peer that is
// successfully authenticated is authorized.
func (observer *DBusAuthObserver) ConnectAuthorizeAuthenticatedPeer(f func(stream IOStreamer, credentials *Credentials) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(observer, "authorize-authenticated-peer", false, unsafe.Pointer(C._gotk4_gio2_DBusAuthObserver_ConnectAuthorizeAuthenticatedPeer), f)
}

// DBusConnection type is used for D-Bus connections to remote peers such as a
// message buses. It is a low-level API that offers a lot of flexibility. For
// instance, it lets you establish a connection over any transport that can by
// represented as a OStream.
//
// This class is rarely used directly in D-Bus clients. If you are writing a
// D-Bus client, it is often easier to use the g_bus_own_name(),
// g_bus_watch_name() or g_dbus_proxy_new_for_bus() APIs.
//
// As an exception to the usual GLib rule that a particular object must not be
// used by two threads at the same time, BusConnection's methods may be called
// from any thread. This is so that g_bus_get() and g_bus_get_sync() can safely
// return the same BusConnection when called from any thread.
//
// Most of the ways to obtain a BusConnection automatically initialize it (i.e.
// connect to D-Bus): for instance, g_dbus_connection_new() and g_bus_get(), and
// the synchronous versions of those methods, give you an initialized
// connection. Language bindings for GIO should use g_initable_new() or
// g_async_initable_new_async(), which also initialize the connection.
//
// If you construct an uninitialized BusConnection, such as via g_object_new(),
// you must initialize it via g_initable_init() or g_async_initable_init_async()
// before using its methods or properties. Calling methods or accessing
// properties on a BusConnection that has not completed initialization
// successfully is considered to be invalid, and leads to undefined behaviour.
// In particular, if initialization fails with a #GError, the only valid thing
// you can do with that BusConnection is to free it with g_object_unref().
//
//
// An example D-Bus server
//
// Here is an example for a D-Bus server: gdbus-example-server.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-server.c)
//
//
// An example for exporting a subtree
//
// Here is an example for exporting a subtree: gdbus-example-subtree.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-subtree.c)
//
//
// An example for file descriptor passing
//
// Here is an example for passing UNIX file descriptors: gdbus-unix-fd-client.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-unix-fd-client.c)
//
//
// An example for exporting a GObject
//
// Here is an example for exporting a #GObject: gdbus-example-export.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-export.c).
type DBusConnection struct {
	_ [0]func() // equal guard
	*externglib.Object

	AsyncInitable
	Initable
}

var (
	_ externglib.Objector = (*DBusConnection)(nil)
)

func wrapDBusConnection(obj *externglib.Object) *DBusConnection {
	return &DBusConnection{
		Object: obj,
		AsyncInitable: AsyncInitable{
			Object: obj,
		},
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalDBusConnectioner(p uintptr) (interface{}, error) {
	return wrapDBusConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusConnection_ConnectClosed
func _gotk4_gio2_DBusConnection_ConnectClosed(arg0 C.gpointer, arg1 C.gboolean, arg2 *C.GError, arg3 C.guintptr) {
	var f func(remotePeerVanished bool, err error)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(remotePeerVanished bool, err error))
	}

	var _remotePeerVanished bool // out
	var _err error               // out

	if arg1 != 0 {
		_remotePeerVanished = true
	}
	if arg2 != nil {
		_err = gerror.Take(unsafe.Pointer(arg2))
	}

	f(_remotePeerVanished, _err)
}

// ConnectClosed: emitted when the connection is closed.
//
//
// The cause of this event can be
//
// - If g_dbus_connection_close() is called. In this case remote_peer_vanished
// is set to FALSE and error is NULL.
//
// - If the remote peer closes the connection. In this case remote_peer_vanished
// is set to TRUE and error is set.
//
// - If the remote peer sends invalid or malformed data. In this case
// remote_peer_vanished is set to FALSE and error is set.
//
// Upon receiving this signal, you should give up your reference to connection.
// You are guaranteed that this signal is emitted only once.
func (connection *DBusConnection) ConnectClosed(f func(remotePeerVanished bool, err error)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(connection, "closed", false, unsafe.Pointer(C._gotk4_gio2_DBusConnection_ConnectClosed), f)
}

// DBusMenuModel is an implementation of Model that can be used as a proxy for a
// menu model that is exported over D-Bus with
// g_dbus_connection_export_menu_model().
type DBusMenuModel struct {
	_ [0]func() // equal guard
	MenuModel
}

var (
	_ MenuModeller = (*DBusMenuModel)(nil)
)

func wrapDBusMenuModel(obj *externglib.Object) *DBusMenuModel {
	return &DBusMenuModel{
		MenuModel: MenuModel{
			Object: obj,
		},
	}
}

func marshalDBusMenuModeller(p uintptr) (interface{}, error) {
	return wrapDBusMenuModel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusMessage: type for representing D-Bus messages that can be sent or
// received on a BusConnection.
type DBusMessage struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DBusMessage)(nil)
)

func wrapDBusMessage(obj *externglib.Object) *DBusMessage {
	return &DBusMessage{
		Object: obj,
	}
}

func marshalDBusMessager(p uintptr) (interface{}, error) {
	return wrapDBusMessage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusMethodInvocation instances of the BusMethodInvocation class are used when
// handling D-Bus method calls. It provides a way to asynchronously return
// results and errors.
//
// The normal way to obtain a BusMethodInvocation object is to receive it as an
// argument to the handle_method_call() function in a BusInterfaceVTable that
// was passed to g_dbus_connection_register_object().
type DBusMethodInvocation struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DBusMethodInvocation)(nil)
)

func wrapDBusMethodInvocation(obj *externglib.Object) *DBusMethodInvocation {
	return &DBusMethodInvocation{
		Object: obj,
	}
}

func marshalDBusMethodInvocationer(p uintptr) (interface{}, error) {
	return wrapDBusMethodInvocation(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DBusServer is a helper for listening to and accepting D-Bus connections. This
// can be used to create a new D-Bus server, allowing two peers to use the D-Bus
// protocol for their own specialized communication. A server instance provided
// in this way will not perform message routing or implement the
// org.freedesktop.DBus interface.
//
// To just export an object on a well-known name on a message bus, such as the
// session or system bus, you should instead use g_bus_own_name().
//
// An example of peer-to-peer communication with GDBus can be found in
// gdbus-example-peer.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-peer.c).
//
// Note that a minimal BusServer will accept connections from any peer. In many
// use-cases it will be necessary to add a BusAuthObserver that only accepts
// connections that have successfully authenticated as the same user that is
// running the BusServer. Since GLib 2.68 this can be achieved more simply by
// passing the G_DBUS_SERVER_FLAGS_AUTHENTICATION_REQUIRE_SAME_USER flag to the
// server.
type DBusServer struct {
	_ [0]func() // equal guard
	*externglib.Object

	Initable
}

var (
	_ externglib.Objector = (*DBusServer)(nil)
)

func wrapDBusServer(obj *externglib.Object) *DBusServer {
	return &DBusServer{
		Object: obj,
		Initable: Initable{
			Object: obj,
		},
	}
}

func marshalDBusServerer(p uintptr) (interface{}, error) {
	return wrapDBusServer(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_DBusServer_ConnectNewConnection
func _gotk4_gio2_DBusServer_ConnectNewConnection(arg0 C.gpointer, arg1 *C.GDBusConnection, arg2 C.guintptr) (cret C.gboolean) {
	var f func(connection *DBusConnection) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(connection *DBusConnection) (ok bool))
	}

	var _connection *DBusConnection // out

	_connection = wrapDBusConnection(externglib.Take(unsafe.Pointer(arg1)))

	ok := f(_connection)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectNewConnection: emitted when a new authenticated connection has been
// made. Use g_dbus_connection_get_peer_credentials() to figure out what
// identity (if any), was authenticated.
//
// If you want to accept the connection, take a reference to the connection
// object and return TRUE. When you are done with the connection call
// g_dbus_connection_close() and give up your reference. Note that the other
// peer may disconnect at any time - a typical thing to do when accepting a
// connection is to listen to the BusConnection::closed signal.
//
// If BusServer:flags contains G_DBUS_SERVER_FLAGS_RUN_IN_THREAD then the signal
// is emitted in a new thread dedicated to the connection. Otherwise the signal
// is emitted in the [thread-default main
// context][g-main-context-push-thread-default] of the thread that server was
// constructed in.
//
// You are guaranteed that signal handlers for this signal runs before incoming
// messages on connection are processed. This means that it's suitable to call
// g_dbus_connection_register_object() or similar from the signal handler.
func (server *DBusServer) ConnectNewConnection(f func(connection *DBusConnection) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(server, "new-connection", false, unsafe.Pointer(C._gotk4_gio2_DBusServer_ConnectNewConnection), f)
}

// Menu is a simple implementation of Model. You populate a #GMenu by adding
// Item instances to it.
//
// There are some convenience functions to allow you to directly add items
// (avoiding Item) for the common cases. To add a regular item, use
// g_menu_insert(). To add a section, use g_menu_insert_section(). To add a
// submenu, use g_menu_insert_submenu().
type Menu struct {
	_ [0]func() // equal guard
	MenuModel
}

var (
	_ MenuModeller = (*Menu)(nil)
)

func wrapMenu(obj *externglib.Object) *Menu {
	return &Menu{
		MenuModel: MenuModel{
			Object: obj,
		},
	}
}

func marshalMenuer(p uintptr) (interface{}, error) {
	return wrapMenu(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MenuItem is an opaque structure type. You must access it using the functions
// below.
type MenuItem struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*MenuItem)(nil)
)

func wrapMenuItem(obj *externglib.Object) *MenuItem {
	return &MenuItem{
		Object: obj,
	}
}

func marshalMenuItemmer(p uintptr) (interface{}, error) {
	return wrapMenuItem(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Notification is a mechanism for creating a notification to be shown to the
// user -- typically as a pop-up notification presented by the desktop
// environment shell.
//
// The key difference between #GNotification and other similar APIs is that, if
// supported by the desktop environment, notifications sent with #GNotification
// will persist after the application has exited, and even across system
// reboots.
//
// Since the user may click on a notification while the application is not
// running, applications using #GNotification should be able to be started as a
// D-Bus service, using #GApplication.
//
// User interaction with a notification (either the default action, or buttons)
// must be associated with actions on the application (ie: "app." actions). It
// is not possible to route user interaction through the notification itself,
// because the object will not exist if the application is autostarted as a
// result of a notification being clicked.
//
// A notification can be sent with g_application_send_notification().
type Notification struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Notification)(nil)
)

func wrapNotification(obj *externglib.Object) *Notification {
	return &Notification{
		Object: obj,
	}
}

func marshalNotificationer(p uintptr) (interface{}, error) {
	return wrapNotification(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// PropertyAction is a way to get a #GAction with a state value reflecting and
// controlling the value of a #GObject property.
//
// The state of the action will correspond to the value of the property.
// Changing it will change the property (assuming the requested value matches
// the requirements as specified in the Spec).
//
// Only the most common types are presently supported. Booleans are mapped to
// booleans, strings to strings, signed/unsigned integers to int32/uint32 and
// floats and doubles to doubles.
//
// If the property is an enum then the state will be string-typed and conversion
// will automatically be performed between the enum value and "nick" string as
// per the Value table.
//
// Flags types are not currently supported.
//
// Properties of object types, boxed types and pointer types are not supported
// and probably never will be.
//
// Properties of #GVariant types are not currently supported.
//
// If the property is boolean-valued then the action will have a NULL parameter
// type, and activating the action (with no parameter) will toggle the value of
// the property.
//
// In all other cases, the parameter type will correspond to the type of the
// property.
//
// The general idea here is to reduce the number of locations where a particular
// piece of state is kept (and therefore has to be synchronised between). Action
// does not have a separate state that is kept in sync with the property value
// -- its state is the property value.
//
// For example, it might be useful to create a #GAction corresponding to the
// "visible-child-name" property of a Stack so that the current page can be
// switched from a menu. The active radio indication in the menu is then
// directly determined from the active page of the Stack.
//
// An anti-example would be binding the "active-id" property on a ComboBox. This
// is because the state of the combobox itself is probably uninteresting and is
// actually being used to control something else.
//
// Another anti-example would be to bind to the "visible-child-name" property of
// a Stack if this value is actually stored in #GSettings. In that case, the
// real source of the value is #GSettings. If you want a #GAction to control a
// setting stored in #GSettings, see g_settings_create_action() instead, and
// possibly combine its use with g_settings_bind().
type PropertyAction struct {
	_ [0]func() // equal guard
	*externglib.Object

	Action
}

var (
	_ externglib.Objector = (*PropertyAction)(nil)
)

func wrapPropertyAction(obj *externglib.Object) *PropertyAction {
	return &PropertyAction{
		Object: obj,
		Action: Action{
			Object: obj,
		},
	}
}

func marshalPropertyActioner(p uintptr) (interface{}, error) {
	return wrapPropertyAction(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SimpleAction is the obvious simple implementation of the #GAction interface.
// This is the easiest way to create an action for purposes of adding it to a
// ActionGroup.
//
// See also Action.
type SimpleAction struct {
	_ [0]func() // equal guard
	*externglib.Object

	Action
}

var (
	_ externglib.Objector = (*SimpleAction)(nil)
)

func wrapSimpleAction(obj *externglib.Object) *SimpleAction {
	return &SimpleAction{
		Object: obj,
		Action: Action{
			Object: obj,
		},
	}
}

func marshalSimpleActioner(p uintptr) (interface{}, error) {
	return wrapSimpleAction(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gio2_SimpleAction_ConnectActivate
func _gotk4_gio2_SimpleAction_ConnectActivate(arg0 C.gpointer, arg1 *C.GVariant, arg2 C.guintptr) {
	var f func(parameter *glib.Variant)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(parameter *glib.Variant))
	}

	var _parameter *glib.Variant // out

	if arg1 != nil {
		_parameter = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg1)))
		C.g_variant_ref(arg1)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_parameter)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	f(_parameter)
}

// ConnectActivate indicates that the action was just activated.
//
// parameter will always be of the expected type, i.e. the parameter type
// specified when the action was created. If an incorrect type is given when
// activating the action, this signal is not emitted.
//
// Since GLib 2.40, if no handler is connected to this signal then the default
// behaviour for boolean-stated actions with a NULL parameter type is to toggle
// them via the Action::change-state signal. For stateful actions where the
// state type is equal to the parameter type, the default is to forward them
// directly to Action::change-state. This should allow almost all users of
// Action to connect only one handler or the other.
func (simple *SimpleAction) ConnectActivate(f func(parameter *glib.Variant)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(simple, "activate", false, unsafe.Pointer(C._gotk4_gio2_SimpleAction_ConnectActivate), f)
}

//export _gotk4_gio2_SimpleAction_ConnectChangeState
func _gotk4_gio2_SimpleAction_ConnectChangeState(arg0 C.gpointer, arg1 *C.GVariant, arg2 C.guintptr) {
	var f func(value *glib.Variant)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value *glib.Variant))
	}

	var _value *glib.Variant // out

	if arg1 != nil {
		_value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg1)))
		C.g_variant_ref(arg1)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_value)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	f(_value)
}

// ConnectChangeState indicates that the action just received a request to
// change its state.
//
// value will always be of the correct state type, i.e. the type of the initial
// state passed to g_simple_action_new_stateful(). If an incorrect type is given
// when requesting to change the state, this signal is not emitted.
//
// If no handler is connected to this signal then the default behaviour is to
// call g_simple_action_set_state() to set the state to the requested value. If
// you connect a signal handler then no default action is taken. If the state
// should change then you must call g_simple_action_set_state() from the
// handler.
//
// An example of a 'change-state' handler:
//
//    static void
//    change_volume_state (GSimpleAction *action,
//                         GVariant      *value,
//                         gpointer       user_data)
//    {
//      gint requested;
//
//      requested = g_variant_get_int32 (value);
//
//      // Volume only goes from 0 to 10
//      if (0 <= requested && requested <= 10)
//        g_simple_action_set_state (action, value);
//    }
//
// The handler need not set the state to the requested value. It could set it to
// any value at all, or take some other action.
func (simple *SimpleAction) ConnectChangeState(f func(value *glib.Variant)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(simple, "change-state", false, unsafe.Pointer(C._gotk4_gio2_SimpleAction_ConnectChangeState), f)
}

// SimpleIOStream creates a OStream from an arbitrary Stream and Stream. This
// allows any pair of input and output streams to be used with OStream methods.
//
// This is useful when you obtained a Stream and a Stream by other means, for
// instance creating them with platform specific methods as
// g_unix_input_stream_new() or g_win32_input_stream_new(), and you want to take
// advantage of the methods provided by OStream.
type SimpleIOStream struct {
	_ [0]func() // equal guard
	IOStream
}

var (
	_ IOStreamer = (*SimpleIOStream)(nil)
)

func wrapSimpleIOStream(obj *externglib.Object) *SimpleIOStream {
	return &SimpleIOStream{
		IOStream: IOStream{
			Object: obj,
		},
	}
}

func marshalSimpleIOStreamer(p uintptr) (interface{}, error) {
	return wrapSimpleIOStream(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SimplePermission is a trivial implementation of #GPermission that represents
// a permission that is either always or never allowed. The value is given at
// construction and doesn't change.
//
// Calling request or release will result in errors.
type SimplePermission struct {
	_ [0]func() // equal guard
	Permission
}

var (
	_ Permissioner = (*SimplePermission)(nil)
)

func wrapSimplePermission(obj *externglib.Object) *SimplePermission {
	return &SimplePermission{
		Permission: Permission{
			Object: obj,
		},
	}
}

func marshalSimplePermissioner(p uintptr) (interface{}, error) {
	return wrapSimplePermission(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TestDBus: helper class for testing code which uses D-Bus without touching the
// user's session bus.
//
// Note that DBus modifies the user’s environment, calling setenv(). This is not
// thread-safe, so all DBus calls should be completed before threads are
// spawned, or should have appropriate locking to ensure no access conflicts to
// environment variables shared between DBus and other threads.
//
//
// Creating unit tests using GTestDBus
//
// Testing of D-Bus services can be tricky because normally we only ever run
// D-Bus services over an existing instance of the D-Bus daemon thus we usually
// don't activate D-Bus services that are not yet installed into the target
// system. The DBus object makes this easier for us by taking care of the lower
// level tasks such as running a private D-Bus daemon and looking up uninstalled
// services in customizable locations, typically in your source code tree.
//
// The first thing you will need is a separate service description file for the
// D-Bus daemon. Typically a services subdirectory of your tests directory is a
// good place to put this file.
//
// The service file should list your service along with an absolute path to the
// uninstalled service executable in your source tree. Using autotools we would
// achieve this by adding a file such as my-server.service.in in the services
// directory and have it processed by configure.
//
//    [D-BUS Service]
//    Name=org.gtk.GDBus.Examples.ObjectManager
//    Exec=abs_top_builddir@/gio/tests/gdbus-example-objectmanager-server
//
// You will also need to indicate this service directory in your test fixtures,
// so you will need to pass the path while compiling your test cases. Typically
// this is done with autotools with an added preprocessor flag specified to
// compile your tests such as:
//
//       -DTEST_SERVICES=\""$(abs_top_builddir)/tests/services"\"
//
//    Once you have a service definition file which is local to your source tree,
//
// you can proceed to set up a GTest fixture using the DBus scaffolding.
//
// An example of a test fixture for D-Bus services can be found here:
// gdbus-test-fixture.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-test-fixture.c)
//
// Note that these examples only deal with isolating the D-Bus aspect of your
// service. To successfully run isolated unit tests on your service you may need
// some additional modifications to your test case fixture. For example; if your
// service uses GSettings and installs a schema then it is important that your
// test service not load the schema in the ordinary installed location (chances
// are that your service and schema files are not yet installed, or worse; there
// is an older version of the schema file sitting in the install location).
//
// Most of the time we can work around these obstacles using the environment.
// Since the environment is inherited by the D-Bus daemon created by DBus and
// then in turn inherited by any services the D-Bus daemon activates, using the
// setup routine for your fixture is a practical place to help sandbox your
// runtime environment. For the rather typical GSettings case we can work around
// this by setting GSETTINGS_SCHEMA_DIR to the in tree directory holding your
// schemas in the above fixture_setup() routine.
//
// The GSettings schemas need to be locally pre-compiled for this to work. This
// can be achieved by compiling the schemas locally as a step before running
// test cases, an autotools setup might do the following in the directory
// holding schemas:
//
//        all-am:
//                $(GLIB_COMPILE_SCHEMAS) .
//
//        CLEANFILES += gschemas.compiled.
type TestDBus struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*TestDBus)(nil)
)

func wrapTestDBus(obj *externglib.Object) *TestDBus {
	return &TestDBus{
		Object: obj,
	}
}

func marshalTestDBusser(p uintptr) (interface{}, error) {
	return wrapTestDBus(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
