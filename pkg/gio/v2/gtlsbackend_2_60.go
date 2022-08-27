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

// SetDefaultDatabase: set the default Database used to verify TLS connections
//
// Any subsequent call to g_tls_backend_get_default_database() will return the
// database set in this call. Existing databases and connections are not
// modified.
//
// Setting a NULL default database will reset to using the system default
// database as if g_tls_backend_set_default_database() had never been called.
//
// The function takes the following parameters:
//
//    - database (optional): Database.
//
func (backend *TLSBackend) SetDefaultDatabase(database TLSDatabaser) {
	var _arg0 *C.GTlsBackend  // out
	var _arg1 *C.GTlsDatabase // out

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))
	if database != nil {
		_arg1 = (*C.GTlsDatabase)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	}

	C.g_tls_backend_set_default_database(_arg0, _arg1)
	runtime.KeepAlive(backend)
	runtime.KeepAlive(database)
}
