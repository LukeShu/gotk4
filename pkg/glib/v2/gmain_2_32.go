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

// SOURCE_CONTINUE: use this macro as the return value of a Func to leave the
// #GSource in the main loop.
const SOURCE_CONTINUE = true

// SOURCE_REMOVE: use this macro as the return value of a Func to remove the
// #GSource from the main loop.
const SOURCE_REMOVE = false

// MainContextRefThreadDefault gets the thread-default Context for this
// thread, as with g_main_context_get_thread_default(), but also adds
// a reference to it with g_main_context_ref(). In addition, unlike
// g_main_context_get_thread_default(), if the thread-default context is the
// global default context, this will return that Context (with a ref added to
// it) rather than returning NULL.
//
// The function returns the following values:
//
//   - mainContext: thread-default Context. Unref with g_main_context_unref()
//     when you are done with it.
//
func MainContextRefThreadDefault() *MainContext {
	var _cret *C.GMainContext // in

	_cret = C.g_main_context_ref_thread_default()

	var _mainContext *MainContext // out

	_mainContext = (*MainContext)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_mainContext)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_main_context_unref((*C.GMainContext)(intern.C))
		},
	)

	return _mainContext
}
