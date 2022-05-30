// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// NewPollableSource: utility method for InputStream and OutputStream
// implementations. Creates a new #GSource that expects a callback of type
// SourceFunc. The new source does not actually do anything on its own; use
// g_source_add_child_source() to add other sources to it to cause it to
// trigger.
//
// The function takes the following parameters:
//
//    - pollableStream: stream associated with the new source.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func NewPollableSource(pollableStream *coreglib.Object) *glib.Source {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(pollableStream.Native()))
	*(**coreglib.Object)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gio", "pollable_source_new").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pollableStream)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_source_unref((*C.GSource)(intern.C))
		},
	)

	return _source
}
