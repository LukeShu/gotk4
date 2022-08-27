// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// GInputStream* _gotk4_gio2_ApplicationCommandLine_virtual_get_stdin(void* fnptr, GApplicationCommandLine* arg0) {
//   return ((GInputStream* (*)(GApplicationCommandLine*))(fnptr))(arg0);
// };
import "C"

// Stdin gets the stdin of the invoking process.
//
// The Stream can be used to read data passed to the standard input of the
// invoking process. This doesn't work on all platforms. Presently, it is only
// available on UNIX when using a D-Bus daemon capable of passing file
// descriptors. If stdin is not available then NULL will be returned. In the
// future, support may be expanded to other platforms.
//
// You must only call this function once per commandline invocation.
//
// The function returns the following values:
//
//    - inputStream (optional) for stdin.
//
func (cmdline *ApplicationCommandLine) Stdin() InputStreamer {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GInputStream            // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))

	_cret = C.g_application_command_line_get_stdin(_arg0)
	runtime.KeepAlive(cmdline)

	var _inputStream InputStreamer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(InputStreamer)
				return ok
			})
			rv, ok := casted.(InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}

	return _inputStream
}

// Stdin gets the stdin of the invoking process.
//
// The Stream can be used to read data passed to the standard input of the
// invoking process. This doesn't work on all platforms. Presently, it is only
// available on UNIX when using a D-Bus daemon capable of passing file
// descriptors. If stdin is not available then NULL will be returned. In the
// future, support may be expanded to other platforms.
//
// You must only call this function once per commandline invocation.
//
// The function returns the following values:
//
//    - inputStream (optional) for stdin.
//
func (cmdline *ApplicationCommandLine) stdin() InputStreamer {
	gclass := (*C.GApplicationCommandLineClass)(coreglib.PeekParentClass(cmdline))
	fnarg := gclass.get_stdin

	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GInputStream            // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(coreglib.InternObject(cmdline).Native()))

	_cret = C._gotk4_gio2_ApplicationCommandLine_virtual_get_stdin(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(cmdline)

	var _inputStream InputStreamer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(InputStreamer)
				return ok
			})
			rv, ok := casted.(InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}

	return _inputStream
}
