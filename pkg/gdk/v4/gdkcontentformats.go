// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_formats_builder_get_type()), F: marshalContentFormatsBuilder},
	})
}

// InternMIMEType canonicalizes the given mime type and interns the result.
//
// If string is not a valid mime type, NULL is returned instead. See RFC 2048
// for the syntax if mime types.
func InternMIMEType(_string string) string {
	var _arg1 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_intern_mime_type(_arg1)

	runtime.KeepAlive(_string)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ContentFormatsBuilder: GdkContentFormatsBuilder is an auxiliary struct used
// to create new GdkContentFormats, and should not be kept around.
type ContentFormatsBuilder struct {
	nocopy gextras.NoCopy
	native *C.GdkContentFormatsBuilder
}

func marshalContentFormatsBuilder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &ContentFormatsBuilder{native: (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b))}, nil
}

// NewContentFormatsBuilder constructs a struct ContentFormatsBuilder.
func NewContentFormatsBuilder() *ContentFormatsBuilder {
	var _cret *C.GdkContentFormatsBuilder // in

	_cret = C.gdk_content_formats_builder_new()

	var _contentFormatsBuilder *ContentFormatsBuilder // out

	_contentFormatsBuilder = (*ContentFormatsBuilder)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_contentFormatsBuilder, func(v *ContentFormatsBuilder) {
		C.gdk_content_formats_builder_unref((*C.GdkContentFormatsBuilder)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormatsBuilder
}

// AddFormats appends all formats from formats to builder, skipping those that
// already exist.
func (builder *ContentFormatsBuilder) AddFormats(formats *ContentFormats) {
	var _arg0 *C.GdkContentFormatsBuilder // out
	var _arg1 *C.GdkContentFormats        // out

	_arg0 = (*C.GdkContentFormatsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	C.gdk_content_formats_builder_add_formats(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(formats)
}

// AddGType appends type to builder if it has not already been added.
func (builder *ContentFormatsBuilder) AddGType(typ externglib.Type) {
	var _arg0 *C.GdkContentFormatsBuilder // out
	var _arg1 C.GType                     // out

	_arg0 = (*C.GdkContentFormatsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.GType(typ)

	C.gdk_content_formats_builder_add_gtype(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(typ)
}

// AddMIMEType appends mime_type to builder if it has not already been added.
func (builder *ContentFormatsBuilder) AddMIMEType(mimeType string) {
	var _arg0 *C.GdkContentFormatsBuilder // out
	var _arg1 *C.char                     // out

	_arg0 = (*C.GdkContentFormatsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_content_formats_builder_add_mime_type(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(mimeType)
}

// ToFormats creates a new GdkContentFormats from the given builder.
//
// The given GdkContentFormatsBuilder is reset once this function returns; you
// cannot call this function multiple times on the same builder instance.
//
// This function is intended primarily for bindings. C code should use
// gdk.ContentFormatsBuilder.FreeToFormats().
func (builder *ContentFormatsBuilder) ToFormats() *ContentFormats {
	var _arg0 *C.GdkContentFormatsBuilder // out
	var _cret *C.GdkContentFormats        // in

	_arg0 = (*C.GdkContentFormatsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))

	_cret = C.gdk_content_formats_builder_to_formats(_arg0)

	runtime.KeepAlive(builder)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}
