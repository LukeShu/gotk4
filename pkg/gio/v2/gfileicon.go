// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gfileicon.go.
var GTypeFileIcon = externglib.Type(C.g_file_icon_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFileIcon, F: marshalFileIcon},
	})
}

// FileIconOverrider contains methods that are overridable.
type FileIconOverrider interface {
	externglib.Objector
}

// WrapFileIconOverrider wraps the FileIconOverrider
// interface implementation to access the instance methods.
func WrapFileIconOverrider(obj FileIconOverrider) *FileIcon {
	return wrapFileIcon(externglib.BaseObject(obj))
}

// FileIcon specifies an icon by pointing to an image file to be used as icon.
type FileIcon struct {
	_ [0]func() // equal guard
	*externglib.Object

	LoadableIcon
}

var (
	_ externglib.Objector = (*FileIcon)(nil)
)

func classInitFileIconner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFileIcon(obj *externglib.Object) *FileIcon {
	return &FileIcon{
		Object: obj,
		LoadableIcon: LoadableIcon{
			Icon: Icon{
				Object: obj,
			},
		},
	}
}

func marshalFileIcon(p uintptr) (interface{}, error) {
	return wrapFileIcon(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFileIcon creates a new icon for a file.
//
// The function takes the following parameters:
//
//    - file: #GFile.
//
// The function returns the following values:
//
//    - fileIcon for the given file, or NULL on error.
//
func NewFileIcon(file FileOverrider) *FileIcon {
	var _arg1 *C.GFile // out
	var _cret *C.GIcon // in

	_arg1 = (*C.GFile)(unsafe.Pointer(externglib.InternObject(file).Native()))

	_cret = C.g_file_icon_new(_arg1)
	runtime.KeepAlive(file)

	var _fileIcon *FileIcon // out

	_fileIcon = wrapFileIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileIcon
}

// File gets the #GFile associated with the given icon.
//
// The function returns the following values:
//
//    - file: #GFile.
//
func (icon *FileIcon) File() FileOverrider {
	var _arg0 *C.GFileIcon // out
	var _cret *C.GFile     // in

	_arg0 = (*C.GFileIcon)(unsafe.Pointer(externglib.InternObject(icon).Native()))

	_cret = C.g_file_icon_get_file(_arg0)
	runtime.KeepAlive(icon)

	var _file FileOverrider // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(FileOverrider)
			return ok
		})
		rv, ok := casted.(FileOverrider)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
		}
		_file = rv
	}

	return _file
}
