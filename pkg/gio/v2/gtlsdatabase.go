// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
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
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_database_get_type()), F: marshalTLSDatabaser},
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
	CreateCertificateHandle(certificate TLSCertificater) string
	// LookupCertificateForHandle: look up a certificate by its handle.
	//
	// The handle should have been created by calling
	// g_tls_database_create_certificate_handle() on a Database object of the
	// same TLS backend. The handle is designed to remain valid across
	// instantiations of the database.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_for_handle_async() to perform the
	// lookup operation asynchronously.
	LookupCertificateForHandle(handle string, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler) (*TLSCertificate, error)
	// LookupCertificateForHandleAsync: asynchronously look up a certificate by
	// its handle in the database. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	LookupCertificateForHandleAsync(handle string, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback)
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	LookupCertificateForHandleFinish(result AsyncResulter) (*TLSCertificate, error)
	// LookupCertificateIssuer: look up the issuer of @certificate in the
	// database.
	//
	// The Certificate:issuer property of @certificate is not modified, and the
	// two certificates are not hooked into a chain.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_issuer_async() to perform the lookup
	// operation asynchronously.
	LookupCertificateIssuer(certificate TLSCertificater, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler) (*TLSCertificate, error)
	// LookupCertificateIssuerAsync: asynchronously look up the issuer of
	// @certificate in the database. See
	// g_tls_database_lookup_certificate_issuer() for more information.
	LookupCertificateIssuerAsync(certificate TLSCertificater, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback)
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation. See g_tls_database_lookup_certificate_issuer() for more
	// information.
	LookupCertificateIssuerFinish(result AsyncResulter) (*TLSCertificate, error)
	// LookupCertificatesIssuedByAsync: asynchronously look up certificates
	// issued by this issuer in the database. See
	// g_tls_database_lookup_certificates_issued_by() for more information.
	//
	// The database may choose to hold a reference to the issuer byte array for
	// the duration of of this asynchronous operation. The byte array should not
	// be modified during this time.
	LookupCertificatesIssuedByAsync(issuerRawDn []byte, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback)
	// VerifyChain determines the validity of a certificate chain after looking
	// up and adding any missing certificates to the chain.
	//
	// @chain is a chain of Certificate objects each pointing to the next
	// certificate in the chain by its Certificate:issuer property. The chain
	// may initially consist of one or more certificates. After the verification
	// process is complete, @chain may be modified by adding missing
	// certificates, or removing extra certificates. If a certificate anchor was
	// found, then it is added to the @chain.
	//
	// @purpose describes the purpose (or usage) for which the certificate is
	// being used. Typically @purpose will be set to
	// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate
	// is being used to authenticate a server (and we are acting as the client).
	//
	// The @identity is used to ensure the server certificate is valid for the
	// expected peer identity. If the identity does not match the certificate,
	// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If
	// @identity is nil, that bit will never be set in the return value. The
	// peer identity may also be used to check for pinned certificates (trust
	// exceptions) in the database. These may override the normal verification
	// process on a host-by-host basis.
	//
	// Currently there are no @flags, and G_TLS_DATABASE_VERIFY_NONE should be
	// used.
	//
	// If @chain is found to be valid, then the return value will be 0. If
	// @chain is found to be invalid, then the return value will indicate the
	// problems found. If the function is unable to determine whether @chain is
	// valid or not (eg, because @cancellable is triggered before it completes)
	// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
	// will be set accordingly. @error is not set when @chain is successfully
	// analyzed but found to be invalid.
	//
	// This function can block, use g_tls_database_verify_chain_async() to
	// perform the verification operation asynchronously.
	VerifyChain(chain TLSCertificater, purpose string, identity SocketConnectabler, interaction TLSInteractioner, flags TLSDatabaseVerifyFlags, cancellable Cancellabler) (TLSCertificateFlags, error)
	// VerifyChainAsync: asynchronously determines the validity of a certificate
	// chain after looking up and adding any missing certificates to the chain.
	// See g_tls_database_verify_chain() for more information.
	VerifyChainAsync(chain TLSCertificater, purpose string, identity SocketConnectabler, interaction TLSInteractioner, flags TLSDatabaseVerifyFlags, cancellable Cancellabler, callback AsyncReadyCallback)
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
	VerifyChainFinish(result AsyncResulter) (TLSCertificateFlags, error)
}

// TLSDatabaser describes TLSDatabase's methods.
type TLSDatabaser interface {
	// CreateCertificateHandle: create a handle string for the certificate.
	CreateCertificateHandle(certificate TLSCertificater) string
	// LookupCertificateForHandle: look up a certificate by its handle.
	LookupCertificateForHandle(handle string, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler) (*TLSCertificate, error)
	// LookupCertificateForHandleAsync: asynchronously look up a certificate by
	// its handle in the database.
	LookupCertificateForHandleAsync(handle string, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback)
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle.
	LookupCertificateForHandleFinish(result AsyncResulter) (*TLSCertificate, error)
	// LookupCertificateIssuer: look up the issuer of @certificate in the
	// database.
	LookupCertificateIssuer(certificate TLSCertificater, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler) (*TLSCertificate, error)
	// LookupCertificateIssuerAsync: asynchronously look up the issuer of
	// @certificate in the database.
	LookupCertificateIssuerAsync(certificate TLSCertificater, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback)
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation.
	LookupCertificateIssuerFinish(result AsyncResulter) (*TLSCertificate, error)
	// LookupCertificatesIssuedByAsync: asynchronously look up certificates
	// issued by this issuer in the database.
	LookupCertificatesIssuedByAsync(issuerRawDn []byte, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback)
	// VerifyChain determines the validity of a certificate chain after looking
	// up and adding any missing certificates to the chain.
	VerifyChain(chain TLSCertificater, purpose string, identity SocketConnectabler, interaction TLSInteractioner, flags TLSDatabaseVerifyFlags, cancellable Cancellabler) (TLSCertificateFlags, error)
	// VerifyChainAsync: asynchronously determines the validity of a certificate
	// chain after looking up and adding any missing certificates to the chain.
	VerifyChainAsync(chain TLSCertificater, purpose string, identity SocketConnectabler, interaction TLSInteractioner, flags TLSDatabaseVerifyFlags, cancellable Cancellabler, callback AsyncReadyCallback)
	// VerifyChainFinish: finish an asynchronous verify chain operation.
	VerifyChainFinish(result AsyncResulter) (TLSCertificateFlags, error)
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
type TLSDatabase struct {
	*externglib.Object
}

var (
	_ TLSDatabaser    = (*TLSDatabase)(nil)
	_ gextras.Nativer = (*TLSDatabase)(nil)
)

func wrapTLSDatabase(obj *externglib.Object) TLSDatabaser {
	return &TLSDatabase{
		Object: obj,
	}
}

func marshalTLSDatabaser(p uintptr) (interface{}, error) {
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
func (self *TLSDatabase) CreateCertificateHandle(certificate TLSCertificater) string {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GTlsCertificate // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer((certificate).(gextras.Nativer).Native()))

	_cret = C.g_tls_database_create_certificate_handle(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// LookupCertificateForHandle: look up a certificate by its handle.
//
// The handle should have been created by calling
// g_tls_database_create_certificate_handle() on a Database object of the same
// TLS backend. The handle is designed to remain valid across instantiations of
// the database.
//
// If the handle is no longer valid, or does not point to a certificate in this
// database, then nil will be returned.
//
// This function can block, use
// g_tls_database_lookup_certificate_for_handle_async() to perform the lookup
// operation asynchronously.
func (self *TLSDatabase) LookupCertificateForHandle(handle string, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler) (*TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.gchar                  // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _cret *C.GTlsCertificate        // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(handle)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer((interaction).(gextras.Nativer).Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_tls_database_lookup_certificate_for_handle(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificateForHandleAsync: asynchronously look up a certificate by its
// handle in the database. See g_tls_database_lookup_certificate_for_handle()
// for more information.
func (self *TLSDatabase) LookupCertificateForHandleAsync(handle string, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.gchar                  // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(handle)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer((interaction).(gextras.Nativer).Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(gbox.Assign(callback))

	C.g_tls_database_lookup_certificate_for_handle_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
// certificate by its handle. See g_tls_database_lookup_certificate_for_handle()
// for more information.
//
// If the handle is no longer valid, or does not point to a certificate in this
// database, then nil will be returned.
func (self *TLSDatabase) LookupCertificateForHandleFinish(result AsyncResulter) (*TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_tls_database_lookup_certificate_for_handle_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificateIssuer: look up the issuer of @certificate in the database.
//
// The Certificate:issuer property of @certificate is not modified, and the two
// certificates are not hooked into a chain.
//
// This function can block, use g_tls_database_lookup_certificate_issuer_async()
// to perform the lookup operation asynchronously.
func (self *TLSDatabase) LookupCertificateIssuer(certificate TLSCertificater, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler) (*TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _cret *C.GTlsCertificate        // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer((certificate).(gextras.Nativer).Native()))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer((interaction).(gextras.Nativer).Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_tls_database_lookup_certificate_issuer(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificateIssuerAsync: asynchronously look up the issuer of
// @certificate in the database. See g_tls_database_lookup_certificate_issuer()
// for more information.
func (self *TLSDatabase) LookupCertificateIssuerAsync(certificate TLSCertificater, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer((certificate).(gextras.Nativer).Native()))
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer((interaction).(gextras.Nativer).Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(gbox.Assign(callback))

	C.g_tls_database_lookup_certificate_issuer_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
// operation. See g_tls_database_lookup_certificate_issuer() for more
// information.
func (self *TLSDatabase) LookupCertificateIssuerFinish(result AsyncResulter) (*TLSCertificate, error) {
	var _arg0 *C.GTlsDatabase    // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_tls_database_lookup_certificate_issuer_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TLSCertificate)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// LookupCertificatesIssuedByAsync: asynchronously look up certificates issued
// by this issuer in the database. See
// g_tls_database_lookup_certificates_issued_by() for more information.
//
// The database may choose to hold a reference to the issuer byte array for the
// duration of of this asynchronous operation. The byte array should not be
// modified during this time.
func (self *TLSDatabase) LookupCertificatesIssuedByAsync(issuerRawDn []byte, interaction TLSInteractioner, flags TLSDatabaseLookupFlags, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase // out
	var _arg1 *C.GByteArray
	var _arg2 *C.GTlsInteraction        // out
	var _arg3 C.GTlsDatabaseLookupFlags // out
	var _arg4 *C.GCancellable           // out
	var _arg5 C.GAsyncReadyCallback     // out
	var _arg6 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = C.g_byte_array_new_take((*C.guint8)(&issuerRawDn[0]), C.gsize(len(issuerRawDn)))
	defer C.g_byte_array_steal(_arg1, nil)
	_arg2 = (*C.GTlsInteraction)(unsafe.Pointer((interaction).(gextras.Nativer).Native()))
	_arg3 = C.GTlsDatabaseLookupFlags(flags)
	_arg4 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg5 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg6 = C.gpointer(gbox.Assign(callback))

	C.g_tls_database_lookup_certificates_issued_by_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// VerifyChain determines the validity of a certificate chain after looking up
// and adding any missing certificates to the chain.
//
// @chain is a chain of Certificate objects each pointing to the next
// certificate in the chain by its Certificate:issuer property. The chain may
// initially consist of one or more certificates. After the verification process
// is complete, @chain may be modified by adding missing certificates, or
// removing extra certificates. If a certificate anchor was found, then it is
// added to the @chain.
//
// @purpose describes the purpose (or usage) for which the certificate is being
// used. Typically @purpose will be set to
// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate is
// being used to authenticate a server (and we are acting as the client).
//
// The @identity is used to ensure the server certificate is valid for the
// expected peer identity. If the identity does not match the certificate,
// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If @identity
// is nil, that bit will never be set in the return value. The peer identity may
// also be used to check for pinned certificates (trust exceptions) in the
// database. These may override the normal verification process on a
// host-by-host basis.
//
// Currently there are no @flags, and G_TLS_DATABASE_VERIFY_NONE should be used.
//
// If @chain is found to be valid, then the return value will be 0. If @chain is
// found to be invalid, then the return value will indicate the problems found.
// If the function is unable to determine whether @chain is valid or not (eg,
// because @cancellable is triggered before it completes) then the return value
// will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error will be set accordingly.
// @error is not set when @chain is successfully analyzed but found to be
// invalid.
//
// This function can block, use g_tls_database_verify_chain_async() to perform
// the verification operation asynchronously.
func (self *TLSDatabase) VerifyChain(chain TLSCertificater, purpose string, identity SocketConnectabler, interaction TLSInteractioner, flags TLSDatabaseVerifyFlags, cancellable Cancellabler) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.gchar                  // out
	var _arg3 *C.GSocketConnectable     // out
	var _arg4 *C.GTlsInteraction        // out
	var _arg5 C.GTlsDatabaseVerifyFlags // out
	var _arg6 *C.GCancellable           // out
	var _cret C.GTlsCertificateFlags    // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer((chain).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(purpose)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GSocketConnectable)(unsafe.Pointer((identity).(gextras.Nativer).Native()))
	_arg4 = (*C.GTlsInteraction)(unsafe.Pointer((interaction).(gextras.Nativer).Native()))
	_arg5 = C.GTlsDatabaseVerifyFlags(flags)
	_arg6 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.g_tls_database_verify_chain(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_cerr)

	var _tlsCertificateFlags TLSCertificateFlags // out
	var _goerr error                             // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificateFlags, _goerr
}

// VerifyChainAsync: asynchronously determines the validity of a certificate
// chain after looking up and adding any missing certificates to the chain. See
// g_tls_database_verify_chain() for more information.
func (self *TLSDatabase) VerifyChainAsync(chain TLSCertificater, purpose string, identity SocketConnectabler, interaction TLSInteractioner, flags TLSDatabaseVerifyFlags, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsDatabase           // out
	var _arg1 *C.GTlsCertificate        // out
	var _arg2 *C.gchar                  // out
	var _arg3 *C.GSocketConnectable     // out
	var _arg4 *C.GTlsInteraction        // out
	var _arg5 C.GTlsDatabaseVerifyFlags // out
	var _arg6 *C.GCancellable           // out
	var _arg7 C.GAsyncReadyCallback     // out
	var _arg8 C.gpointer

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer((chain).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(purpose)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GSocketConnectable)(unsafe.Pointer((identity).(gextras.Nativer).Native()))
	_arg4 = (*C.GTlsInteraction)(unsafe.Pointer((interaction).(gextras.Nativer).Native()))
	_arg5 = C.GTlsDatabaseVerifyFlags(flags)
	_arg6 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))
	_arg7 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg8 = C.gpointer(gbox.Assign(callback))

	C.g_tls_database_verify_chain_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
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
func (self *TLSDatabase) VerifyChainFinish(result AsyncResulter) (TLSCertificateFlags, error) {
	var _arg0 *C.GTlsDatabase        // out
	var _arg1 *C.GAsyncResult        // out
	var _cret C.GTlsCertificateFlags // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GTlsDatabase)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.g_tls_database_verify_chain_finish(_arg0, _arg1, &_cerr)

	var _tlsCertificateFlags TLSCertificateFlags // out
	var _goerr error                             // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificateFlags, _goerr
}
