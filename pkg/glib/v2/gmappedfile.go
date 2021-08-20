// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_mapped_file_get_type()), F: marshalMappedFile},
	})
}

// MappedFile represents a file mapping created with g_mapped_file_new(). It has
// only private members and should not be accessed directly.
//
// An instance of this type is always passed by reference.
type MappedFile struct {
	*mappedFile
}

// mappedFile is the struct that's finalized.
type mappedFile struct {
	native *C.GMappedFile
}

func marshalMappedFile(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &MappedFile{&mappedFile{(*C.GMappedFile)(unsafe.Pointer(b))}}, nil
}

// NewMappedFile constructs a struct MappedFile.
func NewMappedFile(filename string, writable bool) (*MappedFile, error) {
	var _arg1 *C.gchar       // out
	var _arg2 C.gboolean     // out
	var _cret *C.GMappedFile // in
	var _cerr *C.GError      // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	if writable {
		_arg2 = C.TRUE
	}

	_cret = C.g_mapped_file_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(writable)

	var _mappedFile *MappedFile // out
	var _goerr error            // out

	_mappedFile = (*MappedFile)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_mappedFile)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_mapped_file_unref((*C.GMappedFile)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _mappedFile, _goerr
}

// NewMappedFileFromFd constructs a struct MappedFile.
func NewMappedFileFromFd(fd int, writable bool) (*MappedFile, error) {
	var _arg1 C.gint         // out
	var _arg2 C.gboolean     // out
	var _cret *C.GMappedFile // in
	var _cerr *C.GError      // in

	_arg1 = C.gint(fd)
	if writable {
		_arg2 = C.TRUE
	}

	_cret = C.g_mapped_file_new_from_fd(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(fd)
	runtime.KeepAlive(writable)

	var _mappedFile *MappedFile // out
	var _goerr error            // out

	_mappedFile = (*MappedFile)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_mappedFile)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_mapped_file_unref((*C.GMappedFile)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _mappedFile, _goerr
}

// Bytes creates a new #GBytes which references the data mapped from file. The
// mapped contents of the file must not be modified after creating this bytes
// object, because a #GBytes should be immutable.
func (file *MappedFile) Bytes() *Bytes {
	var _arg0 *C.GMappedFile // out
	var _cret *C.GBytes      // in

	_arg0 = (*C.GMappedFile)(gextras.StructNative(unsafe.Pointer(file)))

	_cret = C.g_mapped_file_get_bytes(_arg0)
	runtime.KeepAlive(file)

	var _bytes *Bytes // out

	_bytes = (*Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}

// Contents returns the contents of a File.
//
// Note that the contents may not be zero-terminated, even if the File is backed
// by a text file.
//
// If the file is empty then NULL is returned.
func (file *MappedFile) Contents() string {
	var _arg0 *C.GMappedFile // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GMappedFile)(gextras.StructNative(unsafe.Pointer(file)))

	_cret = C.g_mapped_file_get_contents(_arg0)
	runtime.KeepAlive(file)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Length returns the length of the contents of a File.
func (file *MappedFile) Length() uint {
	var _arg0 *C.GMappedFile // out
	var _cret C.gsize        // in

	_arg0 = (*C.GMappedFile)(gextras.StructNative(unsafe.Pointer(file)))

	_cret = C.g_mapped_file_get_length(_arg0)
	runtime.KeepAlive(file)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}
