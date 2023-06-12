// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeFilterOutputStream = coreglib.Type(C.g_filter_output_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFilterOutputStream, F: marshalFilterOutputStream},
	})
}

// FilterOutputStreamOverrides contains methods that are overridable.
type FilterOutputStreamOverrides struct {
}

func defaultFilterOutputStreamOverrides(v *FilterOutputStream) FilterOutputStreamOverrides {
	return FilterOutputStreamOverrides{}
}

// FilterOutputStream: base class for output stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterOutputStream struct {
	_ [0]func() // equal guard
	OutputStream
}

var (
	_ OutputStreamer = (*FilterOutputStream)(nil)
)

// FilterOutputStreamer describes types inherited from class FilterOutputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FilterOutputStreamer interface {
	coreglib.Objector
	baseFilterOutputStream() *FilterOutputStream
}

var _ FilterOutputStreamer = (*FilterOutputStream)(nil)

func init() {
	coreglib.RegisterClassInfo[*FilterOutputStream, *FilterOutputStreamClass, FilterOutputStreamOverrides](
		GTypeFilterOutputStream,
		initFilterOutputStreamClass,
		wrapFilterOutputStream,
		defaultFilterOutputStreamOverrides,
	)
}

func initFilterOutputStreamClass(gclass unsafe.Pointer, overrides FilterOutputStreamOverrides, classInitFunc func(*FilterOutputStreamClass)) {
	if classInitFunc != nil {
		class := (*FilterOutputStreamClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFilterOutputStream(obj *coreglib.Object) *FilterOutputStream {
	return &FilterOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
	}
}

func marshalFilterOutputStream(p uintptr) (interface{}, error) {
	return wrapFilterOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (stream *FilterOutputStream) baseFilterOutputStream() *FilterOutputStream {
	return stream
}

// BaseFilterOutputStream returns the underlying base object.
func BaseFilterOutputStream(obj FilterOutputStreamer) *FilterOutputStream {
	return obj.baseFilterOutputStream()
}

// BaseStream gets the base stream for the filter stream.
//
// The function returns the following values:
//
//   - outputStream: Stream.
//
func (stream *FilterOutputStream) BaseStream() OutputStreamer {
	var _arg0 *C.GFilterOutputStream // out
	var _cret *C.GOutputStream       // in

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_filter_output_stream_get_base_stream(_arg0)
	runtime.KeepAlive(stream)

	var _outputStream OutputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.OutputStreamer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(OutputStreamer)
			return ok
		})
		rv, ok := casted.(OutputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.OutputStreamer")
		}
		_outputStream = rv
	}

	return _outputStream
}

// CloseBaseStream returns whether the base stream will be closed when stream is
// closed.
//
// The function returns the following values:
//
//   - ok: TRUE if the base stream will be closed.
//
func (stream *FilterOutputStream) CloseBaseStream() bool {
	var _arg0 *C.GFilterOutputStream // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_cret = C.g_filter_output_stream_get_close_base_stream(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when stream is
// closed.
//
// The function takes the following parameters:
//
//   - closeBase: TRUE to close the base stream.
//
func (stream *FilterOutputStream) SetCloseBaseStream(closeBase bool) {
	var _arg0 *C.GFilterOutputStream // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GFilterOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	if closeBase {
		_arg1 = C.TRUE
	}

	C.g_filter_output_stream_set_close_base_stream(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(closeBase)
}

// FilterOutputStreamClass: instance of this type is always passed by reference.
type FilterOutputStreamClass struct {
	*filterOutputStreamClass
}

// filterOutputStreamClass is the struct that's finalized.
type filterOutputStreamClass struct {
	native *C.GFilterOutputStreamClass
}

func (f *FilterOutputStreamClass) ParentClass() *OutputStreamClass {
	valptr := &f.native.parent_class
	var _v *OutputStreamClass // out
	_v = (*OutputStreamClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
