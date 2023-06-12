// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// ApplicationGetDefault returns the default #GApplication instance for this
// process.
//
// Normally there is only one #GApplication per process and it becomes the
// default when it is created. You can exercise more control over this by using
// g_application_set_default().
//
// If there is no default application then NULL is returned.
//
// The function returns the following values:
//
//   - application (optional): default application for this process, or NULL.
//
func ApplicationGetDefault() *Application {
	var _cret *C.GApplication // in

	_cret = C.g_application_get_default()

	var _application *Application // out

	if _cret != nil {
		_application = wrapApplication(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _application
}
