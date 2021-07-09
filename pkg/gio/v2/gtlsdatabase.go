// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_database_get_type()), F: marshalTLSDatabase},
	})
}

// TLSDatabaseOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSDatabaseOverrider interface {
	// CreateCertificateHandle: create a handle string for the certificate. The
	// database will only be able to create a handle for certificates that
	// originate from the database. In cases where the database cannot create a
	// handle for a certificate, nil will be returned.
	//
	// This handle should be stable across various instances of the application,
	// and between applications. If a certificate is modified in the database,
	// then it is not guaranteed that this handle will continue to point to it.
	CreateCertificateHandle(certificate TLSCertificate) string
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	LookupCertificateForHandleFinish(result AsyncResult) (*TLSCertificateClass, error)
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation. See g_tls_database_lookup_certificate_issuer() for more
	// information.
	LookupCertificateIssuerFinish(result AsyncResult) (*TLSCertificateClass, error)
	// VerifyChainFinish: finish an asynchronous verify chain operation. See
	// g_tls_database_verify_chain() for more information.
	//
	// If @chain is found to be valid, then the return value will be 0. If
	// @chain is found to be invalid, then the return value will indicate the
	// problems found. If the function is unable to determine whether @chain is
	// valid or not (eg, because @cancellable is triggered before it completes)
	// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
	// will be set accordingly. @error is not set when @chain is successfully
	// analyzed but found to be invalid.
	VerifyChainFinish(result AsyncResult) (TLSCertificateFlags, error)
}

// TLSDatabase is used to look up certificates and other information from a
// certificate or key store. It is an abstract base class which TLS library
// specific subtypes override.
//
// A Database may be accessed from multiple threads by the TLS backend. All
// implementations are required to be fully thread-safe.
//
// Most common client applications will not directly interact with Database. It
// is used internally by Connection.
type TLSDatabase interface {
	gextras.Objector

	// CreateCertificateHandle: create a handle string for the certificate. The
	// database will only be able to create a handle for certificates that
	// originate from the database. In cases where the database cannot create a
	// handle for a certificate, nil will be returned.
	//
	// This handle should be stable across various instances of the application,
	// and between applications. If a certificate is modified in the database,
	// then it is not guaranteed that this handle will continue to point to it.
	CreateCertificateHandle(certificate TLSCertificate) string
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	LookupCertificateForHandleFinish(result AsyncResult) (*TLSCertificateClass, error)
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation. See g_tls_database_lookup_certificate_issuer() for more
	// information.
	LookupCertificateIssuerFinish(result AsyncResult) (*TLSCertificateClass, error)
	// VerifyChainFinish: finish an asynchronous verify chain operation. See
	// g_tls_database_verify_chain() for more information.
	//
	// If @chain is found to be valid, then the return value will be 0. If
	// @chain is found to be invalid, then the return value will indicate the
	// problems found. If the function is unable to determine whether @chain is
	// valid or not (eg, because @cancellable is triggered before it completes)
	// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
	// will be set accordingly. @error is not set when @chain is successfully
	// analyzed but found to be invalid.
	VerifyChainFinish(result AsyncResult) (TLSCertificateFlags, error)
}

// TLSDatabaseClass implements the TLSDatabase interface.
type TLSDatabaseClass struct {
	*externglib.Object
}

var _ TLSDatabase = (*TLSDatabaseClass)(nil)

func wrapTLSDatabase(obj *externglib.Object) TLSDatabase {
	return &TLSDatabaseClass{
		Object: obj,
	}
}

func marshalTLSDatabase(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTLSDatabase(obj), nil
}

// CreateCertificateHandle: create a handle string for the certificate. The
// database will only be able to create a handle for certificates that originate
// from the database. In cases where the database cannot create a handle for a
// certificate, nil will be returned.
//
// This handle should be stable across various instances of the application, and
// between applications. If a certificate is modified in the database, then it
// is not guaranteed that this handle will continue to point to it.
func (s *TLSDatabaseClass) CreateCertificateHandle(certificate TLSCertificate) string {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GTlsCertificate // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer((&certificate).Native()))

	_cret = C.g_tls_database_create_certificate_handle(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
// certificate by its handle. See g_tls_database_lookup_certificate_for_handle()
// for more information.
//
// If the handle is no longer valid, or does not point to a certificate in this
// database, then nil will be returned.
func (s *TLSDatabaseClass) LookupCertificateForHandleFinish(result AsyncResult) (*TLSCertificateClass, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&result).Native()))

	_cret = C.g_tls_database_lookup_certificate_for_handle_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificate *TLSCertificateClass // out
	var _goerr error                         // out

	_tlsCertificate = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TLSCertificateClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
// operation. See g_tls_database_lookup_certificate_issuer() for more
// information.
func (s *TLSDatabaseClass) LookupCertificateIssuerFinish(result AsyncResult) (*TLSCertificateClass, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&result).Native()))

	_cret = C.g_tls_database_lookup_certificate_issuer_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificate *TLSCertificateClass // out
	var _goerr error                         // out

	_tlsCertificate = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*TLSCertificateClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// VerifyChainFinish: finish an asynchronous verify chain operation. See
// g_tls_database_verify_chain() for more information.
//
// If @chain is found to be valid, then the return value will be 0. If @chain is
// found to be invalid, then the return value will indicate the problems found.
// If the function is unable to determine whether @chain is valid or not (eg,
// because @cancellable is triggered before it completes) then the return value
// will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error will be set accordingly.
// @error is not set when @chain is successfully analyzed but found to be
// invalid.
func (s *TLSDatabaseClass) VerifyChainFinish(result AsyncResult) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase        // out
	var _arg1 *C.GAsyncResult        // out
	var _cret C.GTlsCertificateFlags // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((&result).Native()))

	_cret = C.g_tls_database_verify_chain_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificateFlags TLSCertificateFlags // out
	var _goerr error                             // out

	_tlsCertificateFlags = (TLSCertificateFlags)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificateFlags, _goerr
}
