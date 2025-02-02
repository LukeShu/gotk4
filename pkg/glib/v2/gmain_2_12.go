// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// MainCurrentSource returns the currently firing source for this thread.
//
// The function returns the following values:
//
//   - source (optional): currently firing source or NULL.
//
func MainCurrentSource() *Source {
	var _cret *C.GSource // in

	_cret = C.g_main_current_source()

	var _source *Source // out

	if _cret != nil {
		_source = (*Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_source_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_source)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_source_unref((*C.GSource)(intern.C))
			},
		)
	}

	return _source
}
