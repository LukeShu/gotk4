// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

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
