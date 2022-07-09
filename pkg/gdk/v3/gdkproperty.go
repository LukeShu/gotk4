// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypePropMode returns the GType for the type PropMode.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePropMode() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "PropMode").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPropMode)
	return gtype
}

// PropMode describes how existing data is combined with new data when using
// gdk_property_change().
type PropMode C.gint

const (
	// PropModeReplace: new data replaces the existing data.
	PropModeReplace PropMode = iota
	// PropModePrepend: new data is prepended to the existing data.
	PropModePrepend
	// PropModeAppend: new data is appended to the existing data.
	PropModeAppend
)

func marshalPropMode(p uintptr) (interface{}, error) {
	return PropMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PropMode.
func (p PropMode) String() string {
	switch p {
	case PropModeReplace:
		return "Replace"
	case PropModePrepend:
		return "Prepend"
	case PropModeAppend:
		return "Append"
	default:
		return fmt.Sprintf("PropMode(%d)", p)
	}
}

// UTF8ToStringTarget converts an UTF-8 string into the best possible
// representation as a STRING. The representation of characters not in STRING is
// not specified; it may be as pseudo-escape sequences \x{ABCD}, or it may be in
// some other form of approximation.
//
// The function takes the following parameters:
//
//    - str: UTF-8 string.
//
// The function returns the following values:
//
//    - utf8 (optional): newly-allocated string, or NULL if the conversion
//      failed. (It should not fail for any properly formed UTF-8 string unless
//      system limits like memory or file descriptors are exceeded.).
//
func UTF8ToStringTarget(str string) string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[0]))

	_gret := girepository.MustFind("Gdk", "utf8_to_string_target").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}
