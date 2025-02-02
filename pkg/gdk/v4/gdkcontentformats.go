// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeContentFormatsBuilder = coreglib.Type(C.gdk_content_formats_builder_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeContentFormatsBuilder, F: marshalContentFormatsBuilder},
	})
}

// InternMIMEType canonicalizes the given mime type and interns the result.
//
// If string is not a valid mime type, NULL is returned instead. See RFC 2048
// for the syntax if mime types.
//
// The function takes the following parameters:
//
//   - str: string of a potential mime type.
//
// The function returns the following values:
//
//   - utf8: interned string for the canonicalized mime type or NULL if the
//     string wasn't a valid mime type.
//
func InternMIMEType(str string) string {
	var _arg1 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_intern_mime_type(_arg1)
	runtime.KeepAlive(str)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ContentFormatsBuilder: GdkContentFormatsBuilder is an auxiliary struct used
// to create new GdkContentFormats, and should not be kept around.
//
// An instance of this type is always passed by reference.
type ContentFormatsBuilder struct {
	*contentFormatsBuilder
}

// contentFormatsBuilder is the struct that's finalized.
type contentFormatsBuilder struct {
	native *C.GdkContentFormatsBuilder
}

func marshalContentFormatsBuilder(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ContentFormatsBuilder{&contentFormatsBuilder{(*C.GdkContentFormatsBuilder)(b)}}, nil
}

// NewContentFormatsBuilder constructs a struct ContentFormatsBuilder.
func NewContentFormatsBuilder() *ContentFormatsBuilder {
	var _cret *C.GdkContentFormatsBuilder // in

	_cret = C.gdk_content_formats_builder_new()

	var _contentFormatsBuilder *ContentFormatsBuilder // out

	_contentFormatsBuilder = (*ContentFormatsBuilder)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormatsBuilder)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_builder_unref((*C.GdkContentFormatsBuilder)(intern.C))
		},
	)

	return _contentFormatsBuilder
}

// AddFormats appends all formats from formats to builder, skipping those that
// already exist.
//
// The function takes the following parameters:
//
//   - formats to add.
//
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
//
// The function takes the following parameters:
//
//   - typ: GType.
//
func (builder *ContentFormatsBuilder) AddGType(typ coreglib.Type) {
	var _arg0 *C.GdkContentFormatsBuilder // out
	var _arg1 C.GType                     // out

	_arg0 = (*C.GdkContentFormatsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = C.GType(typ)

	C.gdk_content_formats_builder_add_gtype(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(typ)
}

// AddMIMEType appends mime_type to builder if it has not already been added.
//
// The function takes the following parameters:
//
//   - mimeType: mime type.
//
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
// The given GdkContentFormatsBuilder is reset once this function returns;
// you cannot call this function multiple times on the same builder instance.
//
// This function is intended primarily for bindings. C code should use
// gdk.ContentFormatsBuilder.FreeToFormats().
//
// The function returns the following values:
//
//   - contentFormats: newly created GdkContentFormats with all the formats
//     added to builder.
//
func (builder *ContentFormatsBuilder) ToFormats() *ContentFormats {
	var _arg0 *C.GdkContentFormatsBuilder // out
	var _cret *C.GdkContentFormats        // in

	_arg0 = (*C.GdkContentFormatsBuilder)(gextras.StructNative(unsafe.Pointer(builder)))

	_cret = C.gdk_content_formats_builder_to_formats(_arg0)
	runtime.KeepAlive(builder)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}
