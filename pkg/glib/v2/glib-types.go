// Code generated by girgen. DO NOT EDIT.

package glib

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// The function returns the following values:
//
func StrvGetType() coreglib.Type {
	var _cret C.GType // in

	_cret = C.g_strv_get_type()

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// The function returns the following values:
//
func VariantGetGType() coreglib.Type {
	var _cret C.GType // in

	_cret = C.g_variant_get_gtype()

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}
