// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_dtls_connection_get_type()), F: marshalDTLSConnection},
	})
}

// DTLSConnectionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DTLSConnectionOverrider interface {
	// NegotiatedProtocol gets the name of the application-layer protocol
	// negotiated during the handshake.
	//
	// If the peer did not use the ALPN extension, or did not advertise a
	// protocol that matched one of @conn's protocols, or the TLS backend does
	// not support ALPN, then this will be nil. See
	// g_dtls_connection_set_advertised_protocols().
	NegotiatedProtocol() string
	// Handshake attempts a TLS handshake on @conn.
	//
	// On the client side, it is never necessary to call this method; although
	// the connection needs to perform a handshake after connecting, Connection
	// will handle this for you automatically when you try to send or receive
	// data on the connection. You can call g_dtls_connection_handshake()
	// manually if you want to know whether the initial handshake succeeded or
	// failed (as opposed to just immediately trying to use @conn to read or
	// write, in which case, if it fails, it may not be possible to tell if it
	// failed before or after completing the handshake), but beware that servers
	// may reject client authentication after the handshake has completed, so a
	// successful handshake does not indicate the connection will be usable.
	//
	// Likewise, on the server side, although a handshake is necessary at the
	// beginning of the communication, you do not need to call this function
	// explicitly unless you want clearer error reporting.
	//
	// Previously, calling g_dtls_connection_handshake() after the initial
	// handshake would trigger a rehandshake; however, this usage was deprecated
	// in GLib 2.60 because rehandshaking was removed from the TLS protocol in
	// TLS 1.3. Since GLib 2.64, calling this function after the initial
	// handshake will no longer do anything.
	//
	// Connection::accept_certificate may be emitted during the handshake.
	Handshake(cancellable Cancellable) error
	// HandshakeAsync: asynchronously performs a TLS handshake on @conn. See
	// g_dtls_connection_handshake() for more information.
	HandshakeAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// HandshakeFinish: finish an asynchronous TLS handshake operation. See
	// g_dtls_connection_handshake() for more information.
	HandshakeFinish(result AsyncResult) error
	// SetAdvertisedProtocols sets the list of application-layer protocols to
	// advertise that the caller is willing to speak on this connection. The
	// Application-Layer Protocol Negotiation (ALPN) extension will be used to
	// negotiate a compatible protocol with the peer; use
	// g_dtls_connection_get_negotiated_protocol() to find the negotiated
	// protocol after the handshake. Specifying nil for the the value of
	// @protocols will disable ALPN negotiation.
	//
	// See IANA TLS ALPN Protocol IDs
	// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
	// for a list of registered protocol IDs.
	SetAdvertisedProtocols(protocols []string)
	// Shutdown: shut down part or all of a DTLS connection.
	//
	// If @shutdown_read is true then the receiving side of the connection is
	// shut down, and further reading is disallowed. Subsequent calls to
	// g_datagram_based_receive_messages() will return G_IO_ERROR_CLOSED.
	//
	// If @shutdown_write is true then the sending side of the connection is
	// shut down, and further writing is disallowed. Subsequent calls to
	// g_datagram_based_send_messages() will return G_IO_ERROR_CLOSED.
	//
	// It is allowed for both @shutdown_read and @shutdown_write to be TRUE —
	// this is equivalent to calling g_dtls_connection_close().
	//
	// If @cancellable is cancelled, the Connection may be left partially-closed
	// and any pending untransmitted data may be lost. Call
	// g_dtls_connection_shutdown() again to complete closing the Connection.
	Shutdown(shutdownRead bool, shutdownWrite bool, cancellable Cancellable) error
	// ShutdownAsync: asynchronously shut down part or all of the DTLS
	// connection. See g_dtls_connection_shutdown() for more information.
	ShutdownAsync(shutdownRead bool, shutdownWrite bool, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ShutdownFinish: finish an asynchronous TLS shutdown operation. See
	// g_dtls_connection_shutdown() for more information.
	ShutdownFinish(result AsyncResult) error
}

// DTLSConnection is the base DTLS connection class type, which wraps a Based
// and provides DTLS encryption on top of it. Its subclasses, ClientConnection
// and ServerConnection, implement client-side and server-side DTLS,
// respectively.
//
// For TLS support, see Connection.
//
// As DTLS is datagram based, Connection implements Based, presenting a
// datagram-socket-like API for the encrypted connection. This operates over a
// base datagram connection, which is also a Based (Connection:base-socket).
//
// To close a DTLS connection, use g_dtls_connection_close().
//
// Neither ServerConnection or ClientConnection set the peer address on their
// base Based if it is a #GSocket — it is up to the caller to do that if they
// wish. If they do not, and g_socket_close() is called on the base socket, the
// Connection will not raise a G_IO_ERROR_NOT_CONNECTED error on further I/O.
type DTLSConnection interface {
	gextras.Objector

	// Close the DTLS connection. This is equivalent to calling
	// g_dtls_connection_shutdown() to shut down both sides of the connection.
	//
	// Closing a Connection waits for all buffered but untransmitted data to be
	// sent before it completes. It then sends a `close_notify` DTLS alert to
	// the peer and may wait for a `close_notify` to be received from the peer.
	// It does not close the underlying Connection:base-socket; that must be
	// closed separately.
	//
	// Once @conn is closed, all other operations will return G_IO_ERROR_CLOSED.
	// Closing a Connection multiple times will not return an error.
	//
	// Connections will be automatically closed when the last reference is
	// dropped, but you might want to call this function to make sure resources
	// are released as early as possible.
	//
	// If @cancellable is cancelled, the Connection may be left partially-closed
	// and any pending untransmitted data may be lost. Call
	// g_dtls_connection_close() again to complete closing the Connection.
	Close(cancellable Cancellable) error
	// CloseAsync: asynchronously close the DTLS connection. See
	// g_dtls_connection_close() for more information.
	CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// CloseFinish: finish an asynchronous TLS close operation. See
	// g_dtls_connection_close() for more information.
	CloseFinish(result AsyncResult) error
	// Certificate gets @conn's certificate, as set by
	// g_dtls_connection_set_certificate().
	Certificate() *TLSCertificateClass
	// Database gets the certificate database that @conn uses to verify peer
	// certificates. See g_dtls_connection_set_database().
	Database() *TLSDatabaseClass
	// Interaction: get the object that will be used to interact with the user.
	// It will be used for things like prompting the user for passwords. If nil
	// is returned, then no user interaction will occur for this connection.
	Interaction() *TLSInteractionClass
	// NegotiatedProtocol gets the name of the application-layer protocol
	// negotiated during the handshake.
	//
	// If the peer did not use the ALPN extension, or did not advertise a
	// protocol that matched one of @conn's protocols, or the TLS backend does
	// not support ALPN, then this will be nil. See
	// g_dtls_connection_set_advertised_protocols().
	NegotiatedProtocol() string
	// PeerCertificate gets @conn's peer's certificate after the handshake has
	// completed or failed. (It is not set during the emission of
	// Connection::accept-certificate.)
	PeerCertificate() *TLSCertificateClass
	// PeerCertificateErrors gets the errors associated with validating @conn's
	// peer's certificate, after the handshake has completed or failed. (It is
	// not set during the emission of Connection::accept-certificate.)
	PeerCertificateErrors() TLSCertificateFlags
	// RehandshakeMode gets @conn rehandshaking mode. See
	// g_dtls_connection_set_rehandshake_mode() for details.
	//
	// Deprecated: since version 2.64.
	RehandshakeMode() TLSRehandshakeMode
	// RequireCloseNotify tests whether or not @conn expects a proper TLS close
	// notification when the connection is closed. See
	// g_dtls_connection_set_require_close_notify() for details.
	RequireCloseNotify() bool
	// Handshake attempts a TLS handshake on @conn.
	//
	// On the client side, it is never necessary to call this method; although
	// the connection needs to perform a handshake after connecting, Connection
	// will handle this for you automatically when you try to send or receive
	// data on the connection. You can call g_dtls_connection_handshake()
	// manually if you want to know whether the initial handshake succeeded or
	// failed (as opposed to just immediately trying to use @conn to read or
	// write, in which case, if it fails, it may not be possible to tell if it
	// failed before or after completing the handshake), but beware that servers
	// may reject client authentication after the handshake has completed, so a
	// successful handshake does not indicate the connection will be usable.
	//
	// Likewise, on the server side, although a handshake is necessary at the
	// beginning of the communication, you do not need to call this function
	// explicitly unless you want clearer error reporting.
	//
	// Previously, calling g_dtls_connection_handshake() after the initial
	// handshake would trigger a rehandshake; however, this usage was deprecated
	// in GLib 2.60 because rehandshaking was removed from the TLS protocol in
	// TLS 1.3. Since GLib 2.64, calling this function after the initial
	// handshake will no longer do anything.
	//
	// Connection::accept_certificate may be emitted during the handshake.
	Handshake(cancellable Cancellable) error
	// HandshakeAsync: asynchronously performs a TLS handshake on @conn. See
	// g_dtls_connection_handshake() for more information.
	HandshakeAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// HandshakeFinish: finish an asynchronous TLS handshake operation. See
	// g_dtls_connection_handshake() for more information.
	HandshakeFinish(result AsyncResult) error
	// SetAdvertisedProtocols sets the list of application-layer protocols to
	// advertise that the caller is willing to speak on this connection. The
	// Application-Layer Protocol Negotiation (ALPN) extension will be used to
	// negotiate a compatible protocol with the peer; use
	// g_dtls_connection_get_negotiated_protocol() to find the negotiated
	// protocol after the handshake. Specifying nil for the the value of
	// @protocols will disable ALPN negotiation.
	//
	// See IANA TLS ALPN Protocol IDs
	// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
	// for a list of registered protocol IDs.
	SetAdvertisedProtocols(protocols []string)
	// SetCertificate: this sets the certificate that @conn will present to its
	// peer during the TLS handshake. For a ServerConnection, it is mandatory to
	// set this, and that will normally be done at construct time.
	//
	// For a ClientConnection, this is optional. If a handshake fails with
	// G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server requires a
	// certificate, and if you try connecting again, you should call this method
	// first. You can call g_dtls_client_connection_get_accepted_cas() on the
	// failed connection to get a list of Certificate Authorities that the
	// server will accept certificates from.
	//
	// (It is also possible that a server will allow the connection with or
	// without a certificate; in that case, if you don't provide a certificate,
	// you can tell that the server requested one by the fact that
	// g_dtls_client_connection_get_accepted_cas() will return non-nil.)
	SetCertificate(certificate TLSCertificate)
	// SetDatabase sets the certificate database that is used to verify peer
	// certificates. This is set to the default database by default. See
	// g_tls_backend_get_default_database(). If set to nil, then peer
	// certificate validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA
	// error (meaning Connection::accept-certificate will always be emitted on
	// client-side connections, unless that bit is not set in
	// ClientConnection:validation-flags).
	SetDatabase(database TLSDatabase)
	// SetInteraction: set the object that will be used to interact with the
	// user. It will be used for things like prompting the user for passwords.
	//
	// The @interaction argument will normally be a derived subclass of
	// Interaction. nil can also be provided if no user interaction should occur
	// for this connection.
	SetInteraction(interaction TLSInteraction)
	// SetRequireCloseNotify sets whether or not @conn expects a proper TLS
	// close notification before the connection is closed. If this is true (the
	// default), then @conn will expect to receive a TLS close notification from
	// its peer before the connection is closed, and will return a
	// G_TLS_ERROR_EOF error if the connection is closed without proper
	// notification (since this may indicate a network error, or
	// man-in-the-middle attack).
	//
	// In some protocols, the application will know whether or not the
	// connection was closed cleanly based on application-level data (because
	// the application-level data includes a length field, or is somehow
	// self-delimiting); in this case, the close notify is redundant and may be
	// omitted. You can use g_dtls_connection_set_require_close_notify() to tell
	// @conn to allow an "unannounced" connection close, in which case the close
	// will show up as a 0-length read, as in a non-TLS Based, and it is up to
	// the application to check that the data has been fully received.
	//
	// Note that this only affects the behavior when the peer closes the
	// connection; when the application calls g_dtls_connection_close_async() on
	// @conn itself, this will send a close notification regardless of the
	// setting of this property. If you explicitly want to do an unclean close,
	// you can close @conn's Connection:base-socket rather than closing @conn
	// itself.
	SetRequireCloseNotify(requireCloseNotify bool)
	// Shutdown: shut down part or all of a DTLS connection.
	//
	// If @shutdown_read is true then the receiving side of the connection is
	// shut down, and further reading is disallowed. Subsequent calls to
	// g_datagram_based_receive_messages() will return G_IO_ERROR_CLOSED.
	//
	// If @shutdown_write is true then the sending side of the connection is
	// shut down, and further writing is disallowed. Subsequent calls to
	// g_datagram_based_send_messages() will return G_IO_ERROR_CLOSED.
	//
	// It is allowed for both @shutdown_read and @shutdown_write to be TRUE —
	// this is equivalent to calling g_dtls_connection_close().
	//
	// If @cancellable is cancelled, the Connection may be left partially-closed
	// and any pending untransmitted data may be lost. Call
	// g_dtls_connection_shutdown() again to complete closing the Connection.
	Shutdown(shutdownRead bool, shutdownWrite bool, cancellable Cancellable) error
	// ShutdownAsync: asynchronously shut down part or all of the DTLS
	// connection. See g_dtls_connection_shutdown() for more information.
	ShutdownAsync(shutdownRead bool, shutdownWrite bool, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ShutdownFinish: finish an asynchronous TLS shutdown operation. See
	// g_dtls_connection_shutdown() for more information.
	ShutdownFinish(result AsyncResult) error
}

// DTLSConnectionInterface implements the DTLSConnection interface.
type DTLSConnectionInterface struct {
	DatagramBasedInterface
}

var _ DTLSConnection = (*DTLSConnectionInterface)(nil)

func wrapDTLSConnection(obj *externglib.Object) DTLSConnection {
	return &DTLSConnectionInterface{
		DatagramBasedInterface: DatagramBasedInterface{
			Object: obj,
		},
	}
}

func marshalDTLSConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDTLSConnection(obj), nil
}

// Close the DTLS connection. This is equivalent to calling
// g_dtls_connection_shutdown() to shut down both sides of the connection.
//
// Closing a Connection waits for all buffered but untransmitted data to be sent
// before it completes. It then sends a `close_notify` DTLS alert to the peer
// and may wait for a `close_notify` to be received from the peer. It does not
// close the underlying Connection:base-socket; that must be closed separately.
//
// Once @conn is closed, all other operations will return G_IO_ERROR_CLOSED.
// Closing a Connection multiple times will not return an error.
//
// Connections will be automatically closed when the last reference is dropped,
// but you might want to call this function to make sure resources are released
// as early as possible.
//
// If @cancellable is cancelled, the Connection may be left partially-closed and
// any pending untransmitted data may be lost. Call g_dtls_connection_close()
// again to complete closing the Connection.
func (c *DTLSConnectionInterface) Close(cancellable Cancellable) error {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GCancellable    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((&cancellable).Native()))

	C.g_dtls_connection_close(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// CloseAsync: asynchronously close the DTLS connection. See
// g_dtls_connection_close() for more information.
func (c *DTLSConnectionInterface) CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GDtlsConnection    // out
	var _arg1 C.int                 // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = C.int(ioPriority)
	_arg2 = (*C.GCancellable)(unsafe.Pointer((&cancellable).Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_dtls_connection_close_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// CloseFinish: finish an asynchronous TLS close operation. See
// g_dtls_connection_close() for more information.
func (c *DTLSConnectionInterface) CloseFinish(result AsyncResult) error {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GAsyncResult    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&result).Native()))

	C.g_dtls_connection_close_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Certificate gets @conn's certificate, as set by
// g_dtls_connection_set_certificate().
func (c *DTLSConnectionInterface) Certificate() *TLSCertificateClass {
	var _arg0 *C.GDtlsConnection // out
	var _cret *C.GTlsCertificate // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_certificate(_arg0)

	var _tlsCertificate *TLSCertificateClass // out

	_tlsCertificate = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*TLSCertificateClass)

	return _tlsCertificate
}

// Database gets the certificate database that @conn uses to verify peer
// certificates. See g_dtls_connection_set_database().
func (c *DTLSConnectionInterface) Database() *TLSDatabaseClass {
	var _arg0 *C.GDtlsConnection // out
	var _cret *C.GTlsDatabase    // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_database(_arg0)

	var _tlsDatabase *TLSDatabaseClass // out

	_tlsDatabase = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*TLSDatabaseClass)

	return _tlsDatabase
}

// Interaction: get the object that will be used to interact with the user. It
// will be used for things like prompting the user for passwords. If nil is
// returned, then no user interaction will occur for this connection.
func (c *DTLSConnectionInterface) Interaction() *TLSInteractionClass {
	var _arg0 *C.GDtlsConnection // out
	var _cret *C.GTlsInteraction // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_interaction(_arg0)

	var _tlsInteraction *TLSInteractionClass // out

	_tlsInteraction = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*TLSInteractionClass)

	return _tlsInteraction
}

// NegotiatedProtocol gets the name of the application-layer protocol negotiated
// during the handshake.
//
// If the peer did not use the ALPN extension, or did not advertise a protocol
// that matched one of @conn's protocols, or the TLS backend does not support
// ALPN, then this will be nil. See
// g_dtls_connection_set_advertised_protocols().
func (c *DTLSConnectionInterface) NegotiatedProtocol() string {
	var _arg0 *C.GDtlsConnection // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_negotiated_protocol(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PeerCertificate gets @conn's peer's certificate after the handshake has
// completed or failed. (It is not set during the emission of
// Connection::accept-certificate.)
func (c *DTLSConnectionInterface) PeerCertificate() *TLSCertificateClass {
	var _arg0 *C.GDtlsConnection // out
	var _cret *C.GTlsCertificate // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_peer_certificate(_arg0)

	var _tlsCertificate *TLSCertificateClass // out

	_tlsCertificate = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*TLSCertificateClass)

	return _tlsCertificate
}

// PeerCertificateErrors gets the errors associated with validating @conn's
// peer's certificate, after the handshake has completed or failed. (It is not
// set during the emission of Connection::accept-certificate.)
func (c *DTLSConnectionInterface) PeerCertificateErrors() TLSCertificateFlags {
	var _arg0 *C.GDtlsConnection     // out
	var _cret C.GTlsCertificateFlags // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_peer_certificate_errors(_arg0)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = (TLSCertificateFlags)(_cret)

	return _tlsCertificateFlags
}

// RehandshakeMode gets @conn rehandshaking mode. See
// g_dtls_connection_set_rehandshake_mode() for details.
//
// Deprecated: since version 2.64.
func (c *DTLSConnectionInterface) RehandshakeMode() TLSRehandshakeMode {
	var _arg0 *C.GDtlsConnection    // out
	var _cret C.GTlsRehandshakeMode // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_rehandshake_mode(_arg0)

	var _tlsRehandshakeMode TLSRehandshakeMode // out

	_tlsRehandshakeMode = (TLSRehandshakeMode)(_cret)

	return _tlsRehandshakeMode
}

// RequireCloseNotify tests whether or not @conn expects a proper TLS close
// notification when the connection is closed. See
// g_dtls_connection_set_require_close_notify() for details.
func (c *DTLSConnectionInterface) RequireCloseNotify() bool {
	var _arg0 *C.GDtlsConnection // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))

	_cret = C.g_dtls_connection_get_require_close_notify(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Handshake attempts a TLS handshake on @conn.
//
// On the client side, it is never necessary to call this method; although the
// connection needs to perform a handshake after connecting, Connection will
// handle this for you automatically when you try to send or receive data on the
// connection. You can call g_dtls_connection_handshake() manually if you want
// to know whether the initial handshake succeeded or failed (as opposed to just
// immediately trying to use @conn to read or write, in which case, if it fails,
// it may not be possible to tell if it failed before or after completing the
// handshake), but beware that servers may reject client authentication after
// the handshake has completed, so a successful handshake does not indicate the
// connection will be usable.
//
// Likewise, on the server side, although a handshake is necessary at the
// beginning of the communication, you do not need to call this function
// explicitly unless you want clearer error reporting.
//
// Previously, calling g_dtls_connection_handshake() after the initial handshake
// would trigger a rehandshake; however, this usage was deprecated in GLib 2.60
// because rehandshaking was removed from the TLS protocol in TLS 1.3. Since
// GLib 2.64, calling this function after the initial handshake will no longer
// do anything.
//
// Connection::accept_certificate may be emitted during the handshake.
func (c *DTLSConnectionInterface) Handshake(cancellable Cancellable) error {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GCancellable    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer((&cancellable).Native()))

	C.g_dtls_connection_handshake(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// HandshakeAsync: asynchronously performs a TLS handshake on @conn. See
// g_dtls_connection_handshake() for more information.
func (c *DTLSConnectionInterface) HandshakeAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GDtlsConnection    // out
	var _arg1 C.int                 // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = C.int(ioPriority)
	_arg2 = (*C.GCancellable)(unsafe.Pointer((&cancellable).Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_dtls_connection_handshake_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// HandshakeFinish: finish an asynchronous TLS handshake operation. See
// g_dtls_connection_handshake() for more information.
func (c *DTLSConnectionInterface) HandshakeFinish(result AsyncResult) error {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GAsyncResult    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&result).Native()))

	C.g_dtls_connection_handshake_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetAdvertisedProtocols sets the list of application-layer protocols to
// advertise that the caller is willing to speak on this connection. The
// Application-Layer Protocol Negotiation (ALPN) extension will be used to
// negotiate a compatible protocol with the peer; use
// g_dtls_connection_get_negotiated_protocol() to find the negotiated protocol
// after the handshake. Specifying nil for the the value of @protocols will
// disable ALPN negotiation.
//
// See IANA TLS ALPN Protocol IDs
// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
// for a list of registered protocol IDs.
func (c *DTLSConnectionInterface) SetAdvertisedProtocols(protocols []string) {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 **C.gchar

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (**C.gchar)(C.malloc(C.ulong(len(protocols)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(protocols))
		for i := range protocols {
			out[i] = (*C.gchar)(C.CString(protocols[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_dtls_connection_set_advertised_protocols(_arg0, _arg1)
}

// SetCertificate: this sets the certificate that @conn will present to its peer
// during the TLS handshake. For a ServerConnection, it is mandatory to set
// this, and that will normally be done at construct time.
//
// For a ClientConnection, this is optional. If a handshake fails with
// G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server requires a
// certificate, and if you try connecting again, you should call this method
// first. You can call g_dtls_client_connection_get_accepted_cas() on the failed
// connection to get a list of Certificate Authorities that the server will
// accept certificates from.
//
// (It is also possible that a server will allow the connection with or without
// a certificate; in that case, if you don't provide a certificate, you can tell
// that the server requested one by the fact that
// g_dtls_client_connection_get_accepted_cas() will return non-nil.)
func (c *DTLSConnectionInterface) SetCertificate(certificate TLSCertificate) {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GTlsCertificate // out

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer((&certificate).Native()))

	C.g_dtls_connection_set_certificate(_arg0, _arg1)
}

// SetDatabase sets the certificate database that is used to verify peer
// certificates. This is set to the default database by default. See
// g_tls_backend_get_default_database(). If set to nil, then peer certificate
// validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA error (meaning
// Connection::accept-certificate will always be emitted on client-side
// connections, unless that bit is not set in
// ClientConnection:validation-flags).
func (c *DTLSConnectionInterface) SetDatabase(database TLSDatabase) {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GTlsDatabase    // out

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GTlsDatabase)(unsafe.Pointer((&database).Native()))

	C.g_dtls_connection_set_database(_arg0, _arg1)
}

// SetInteraction: set the object that will be used to interact with the user.
// It will be used for things like prompting the user for passwords.
//
// The @interaction argument will normally be a derived subclass of Interaction.
// nil can also be provided if no user interaction should occur for this
// connection.
func (c *DTLSConnectionInterface) SetInteraction(interaction TLSInteraction) {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GTlsInteraction // out

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GTlsInteraction)(unsafe.Pointer((&interaction).Native()))

	C.g_dtls_connection_set_interaction(_arg0, _arg1)
}

// SetRequireCloseNotify sets whether or not @conn expects a proper TLS close
// notification before the connection is closed. If this is true (the default),
// then @conn will expect to receive a TLS close notification from its peer
// before the connection is closed, and will return a G_TLS_ERROR_EOF error if
// the connection is closed without proper notification (since this may indicate
// a network error, or man-in-the-middle attack).
//
// In some protocols, the application will know whether or not the connection
// was closed cleanly based on application-level data (because the
// application-level data includes a length field, or is somehow
// self-delimiting); in this case, the close notify is redundant and may be
// omitted. You can use g_dtls_connection_set_require_close_notify() to tell
// @conn to allow an "unannounced" connection close, in which case the close
// will show up as a 0-length read, as in a non-TLS Based, and it is up to the
// application to check that the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the connection;
// when the application calls g_dtls_connection_close_async() on @conn itself,
// this will send a close notification regardless of the setting of this
// property. If you explicitly want to do an unclean close, you can close
// @conn's Connection:base-socket rather than closing @conn itself.
func (c *DTLSConnectionInterface) SetRequireCloseNotify(requireCloseNotify bool) {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	if requireCloseNotify {
		_arg1 = C.TRUE
	}

	C.g_dtls_connection_set_require_close_notify(_arg0, _arg1)
}

// Shutdown: shut down part or all of a DTLS connection.
//
// If @shutdown_read is true then the receiving side of the connection is shut
// down, and further reading is disallowed. Subsequent calls to
// g_datagram_based_receive_messages() will return G_IO_ERROR_CLOSED.
//
// If @shutdown_write is true then the sending side of the connection is shut
// down, and further writing is disallowed. Subsequent calls to
// g_datagram_based_send_messages() will return G_IO_ERROR_CLOSED.
//
// It is allowed for both @shutdown_read and @shutdown_write to be TRUE — this
// is equivalent to calling g_dtls_connection_close().
//
// If @cancellable is cancelled, the Connection may be left partially-closed and
// any pending untransmitted data may be lost. Call g_dtls_connection_shutdown()
// again to complete closing the Connection.
func (c *DTLSConnectionInterface) Shutdown(shutdownRead bool, shutdownWrite bool, cancellable Cancellable) error {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 C.gboolean         // out
	var _arg2 C.gboolean         // out
	var _arg3 *C.GCancellable    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	if shutdownRead {
		_arg1 = C.TRUE
	}
	if shutdownWrite {
		_arg2 = C.TRUE
	}
	_arg3 = (*C.GCancellable)(unsafe.Pointer((&cancellable).Native()))

	C.g_dtls_connection_shutdown(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ShutdownAsync: asynchronously shut down part or all of the DTLS connection.
// See g_dtls_connection_shutdown() for more information.
func (c *DTLSConnectionInterface) ShutdownAsync(shutdownRead bool, shutdownWrite bool, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GDtlsConnection    // out
	var _arg1 C.gboolean            // out
	var _arg2 C.gboolean            // out
	var _arg3 C.int                 // out
	var _arg4 *C.GCancellable       // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	if shutdownRead {
		_arg1 = C.TRUE
	}
	if shutdownWrite {
		_arg2 = C.TRUE
	}
	_arg3 = C.int(ioPriority)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((&cancellable).Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(box.Assign(callback))

	C.g_dtls_connection_shutdown_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// ShutdownFinish: finish an asynchronous TLS shutdown operation. See
// g_dtls_connection_shutdown() for more information.
func (c *DTLSConnectionInterface) ShutdownFinish(result AsyncResult) error {
	var _arg0 *C.GDtlsConnection // out
	var _arg1 *C.GAsyncResult    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GDtlsConnection)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&result).Native()))

	C.g_dtls_connection_shutdown_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
