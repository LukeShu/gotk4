// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// TransformParse parses the given string into a transform and puts it in
// out_transform.
//
// Strings printed via gsk.Transform.ToString() can be read in again
// successfully using this function.
//
// If string does not describe a valid transform, FALSE is returned and NULL is
// put in out_transform.
//
// The function takes the following parameters:
//
//    - str: string to parse.
//
// The function returns the following values:
//
//    - outTransform: location to put the transform in.
//    - ok: TRUE if string described a valid transform.
//
func TransformParse(str string) (*Transform, bool) {
	var args [1]girepository.Argument
	var outs [1]girepository.Argument
	var _arg0 *C.void    // out
	var _out0 *C.void    // in
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg0))
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gsk", "parse").Invoke(args[:], outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(str)

	var _outTransform *Transform // out
	var _ok bool                 // out
	_out0 = *(**Transform)(unsafe.Pointer(&outs[0]))

	_outTransform = (*Transform)(gextras.NewStructNative(unsafe.Pointer(_out0)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_outTransform)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gsk_transform_unref((*C.GskTransform)(intern.C))
		},
	)
	if _cret != 0 {
		_ok = true
	}

	return _outTransform, _ok
}
