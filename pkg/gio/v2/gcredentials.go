// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeCredentials returns the GType for the type Credentials.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCredentials() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "Credentials").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCredentials)
	return gtype
}

// Credentials type is a reference-counted wrapper for native credentials. This
// information is typically used for identifying, authenticating and authorizing
// other processes.
//
// Some operating systems supports looking up the credentials of the remote peer
// of a communication endpoint - see e.g. g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving credentials
// over a Unix Domain Socket, see CredentialsMessage,
// g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a struct ucred - see the unix(7) man
// page for details. This corresponds to G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On Apple operating systems (including iOS, tvOS, and macOS), the native
// credential type is a struct xucred. This corresponds to
// G_CREDENTIALS_TYPE_APPLE_XUCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native credential type is
// a struct cmsgcred. This corresponds to G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a struct unpcbid. This corresponds
// to G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a struct sockpeercred. This
// corresponds to G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native credential
// type is a ucred_t. This corresponds to G_CREDENTIALS_TYPE_SOLARIS_UCRED.
type Credentials struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Credentials)(nil)
)

func wrapCredentials(obj *coreglib.Object) *Credentials {
	return &Credentials{
		Object: obj,
	}
}

func marshalCredentials(p uintptr) (interface{}, error) {
	return wrapCredentials(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCredentials creates a new #GCredentials object with credentials matching
// the the current process.
//
// The function returns the following values:
//
//    - credentials Free with g_object_unref().
//
func NewCredentials() *Credentials {
	_gret := girepository.MustFind("Gio", "Credentials").InvokeMethod("new_Credentials", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _credentials *Credentials // out

	_credentials = wrapCredentials(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _credentials
}

// IsSameUser checks if credentials and other_credentials is the same user.
//
// This operation can fail if #GCredentials is not supported on the the OS.
//
// The function takes the following parameters:
//
//    - otherCredentials: #GCredentials.
//
func (credentials *Credentials) IsSameUser(otherCredentials *Credentials) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(credentials).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(otherCredentials).Native()))

	girepository.MustFind("Gio", "Credentials").InvokeMethod("is_same_user", _args[:], nil)

	runtime.KeepAlive(credentials)
	runtime.KeepAlive(otherCredentials)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// String creates a human-readable textual representation of credentials that
// can be used in logging and debug messages. The format of the returned string
// may change in future GLib release.
//
// The function returns the following values:
//
//    - utf8: string that should be freed with g_free().
//
func (credentials *Credentials) String() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(credentials).Native()))

	_gret := girepository.MustFind("Gio", "Credentials").InvokeMethod("to_string", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(credentials)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
