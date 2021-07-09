// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_backend_get_type()), F: marshalTLSBackend},
	})
}

// TLSBackendOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSBackendOverrider interface {
	// DefaultDatabase gets the default Database used to verify TLS connections.
	DefaultDatabase() *TLSDatabaseClass
	// SupportsDTLS checks if DTLS is supported. DTLS support may not be
	// available even if TLS support is available, and vice-versa.
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns false for the
	// default Backend, it means no "real" TLS backend is available.
	SupportsTLS() bool
}

// TLSBackend: TLS (Transport Layer Security, aka SSL) and DTLS backend.
type TLSBackend interface {
	gextras.Objector

	// CertificateType gets the #GType of @backend's Certificate implementation.
	CertificateType() externglib.Type
	// ClientConnectionType gets the #GType of @backend's ClientConnection
	// implementation.
	ClientConnectionType() externglib.Type
	// DefaultDatabase gets the default Database used to verify TLS connections.
	DefaultDatabase() *TLSDatabaseClass
	// DTLSClientConnectionType gets the #GType of @backend’s ClientConnection
	// implementation.
	DTLSClientConnectionType() externglib.Type
	// DTLSServerConnectionType gets the #GType of @backend’s ServerConnection
	// implementation.
	DTLSServerConnectionType() externglib.Type
	// FileDatabaseType gets the #GType of @backend's FileDatabase
	// implementation.
	FileDatabaseType() externglib.Type
	// ServerConnectionType gets the #GType of @backend's ServerConnection
	// implementation.
	ServerConnectionType() externglib.Type
	// SetDefaultDatabase: set the default Database used to verify TLS
	// connections
	//
	// Any subsequent call to g_tls_backend_get_default_database() will return
	// the database set in this call. Existing databases and connections are not
	// modified.
	//
	// Setting a nil default database will reset to using the system default
	// database as if g_tls_backend_set_default_database() had never been
	// called.
	SetDefaultDatabase(database TLSDatabase)
	// SupportsDTLS checks if DTLS is supported. DTLS support may not be
	// available even if TLS support is available, and vice-versa.
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns false for the
	// default Backend, it means no "real" TLS backend is available.
	SupportsTLS() bool
}

// TLSBackendInterface implements the TLSBackend interface.
type TLSBackendInterface struct {
	*externglib.Object
}

var _ TLSBackend = (*TLSBackendInterface)(nil)

func wrapTLSBackend(obj *externglib.Object) TLSBackend {
	return &TLSBackendInterface{
		Object: obj,
	}
}

func marshalTLSBackend(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTLSBackend(obj), nil
}

// CertificateType gets the #GType of @backend's Certificate implementation.
func (b *TLSBackendInterface) CertificateType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_get_certificate_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// ClientConnectionType gets the #GType of @backend's ClientConnection
// implementation.
func (b *TLSBackendInterface) ClientConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_get_client_connection_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// DefaultDatabase gets the default Database used to verify TLS connections.
func (b *TLSBackendInterface) DefaultDatabase() *TLSDatabaseClass {
	var _arg0 *C.GTlsBackend  // out
	var _cret *C.GTlsDatabase // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_get_default_database(_arg0)

	var _tlsDatabase *TLSDatabaseClass // out

	_tlsDatabase = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TLSDatabaseClass)

	return _tlsDatabase
}

// DTLSClientConnectionType gets the #GType of @backend’s ClientConnection
// implementation.
func (b *TLSBackendInterface) DTLSClientConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_get_dtls_client_connection_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// DTLSServerConnectionType gets the #GType of @backend’s ServerConnection
// implementation.
func (b *TLSBackendInterface) DTLSServerConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_get_dtls_server_connection_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// FileDatabaseType gets the #GType of @backend's FileDatabase implementation.
func (b *TLSBackendInterface) FileDatabaseType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_get_file_database_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// ServerConnectionType gets the #GType of @backend's ServerConnection
// implementation.
func (b *TLSBackendInterface) ServerConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_get_server_connection_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SetDefaultDatabase: set the default Database used to verify TLS connections
//
// Any subsequent call to g_tls_backend_get_default_database() will return the
// database set in this call. Existing databases and connections are not
// modified.
//
// Setting a nil default database will reset to using the system default
// database as if g_tls_backend_set_default_database() had never been called.
func (b *TLSBackendInterface) SetDefaultDatabase(database TLSDatabase) {
	var _arg0 *C.GTlsBackend  // out
	var _arg1 *C.GTlsDatabase // out

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))
	_arg1 = (*C.GTlsDatabase)(unsafe.Pointer((&database).Native()))

	C.g_tls_backend_set_default_database(_arg0, _arg1)
}

// SupportsDTLS checks if DTLS is supported. DTLS support may not be available
// even if TLS support is available, and vice-versa.
func (b *TLSBackendInterface) SupportsDTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_supports_dtls(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsTLS checks if TLS is supported; if this returns false for the default
// Backend, it means no "real" TLS backend is available.
func (b *TLSBackendInterface) SupportsTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer((&b).Native()))

	_cret = C.g_tls_backend_supports_tls(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
