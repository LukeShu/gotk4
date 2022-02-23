// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern GIOChannel* _gotk4_atk1_StreamableContentIface_get_stream(AtkStreamableContent*, gchar*);
// extern gchar* _gotk4_atk1_StreamableContentIface_get_mime_type(AtkStreamableContent*, gint);
// extern gchar* _gotk4_atk1_StreamableContentIface_get_uri(AtkStreamableContent*, gchar*);
// extern gint _gotk4_atk1_StreamableContentIface_get_n_mime_types(AtkStreamableContent*);
import "C"

// glib.Type values for atkstreamablecontent.go.
var GTypeStreamableContent = externglib.Type(C.atk_streamable_content_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeStreamableContent, F: marshalStreamableContent},
	})
}

// StreamableContentOverrider contains methods that are overridable.
type StreamableContentOverrider interface {
	// MIMEType gets the character string of the specified mime type. The first
	// mime type is at position 0, the second at position 1, and so on.
	//
	// The function takes the following parameters:
	//
	//    - i: gint representing the position of the mime type starting from 0.
	//
	// The function returns the following values:
	//
	//    - utf8: gchar* representing the specified mime type; the caller should
	//      not free the character string.
	//
	MIMEType(i int) string
	// NMIMETypes gets the number of mime types supported by this object.
	//
	// The function returns the following values:
	//
	//    - gint which is the number of mime types supported by the object.
	//
	NMIMETypes() int
	// Stream gets the content in the specified mime type.
	//
	// The function takes the following parameters:
	//
	//    - mimeType: gchar* representing the mime type.
	//
	// The function returns the following values:
	//
	//    - ioChannel which contains the content in the specified mime type.
	//
	Stream(mimeType string) *glib.IOChannel
	// URI: get a string representing a URI in IETF standard format (see
	// http://www.ietf.org/rfc/rfc2396.txt) from which the object's content may
	// be streamed in the specified mime-type, if one is available. If mime_type
	// is NULL, the URI for the default (and possibly only) mime-type is
	// returned.
	//
	// Note that it is possible for get_uri to return NULL but for get_stream to
	// work nonetheless, since not all GIOChannels connect to URIs.
	//
	// The function takes the following parameters:
	//
	//    - mimeType: gchar* representing the mime type, or NULL to request a URI
	//      for the default mime type.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional) returns a string representing a URI, or NULL if no
	//      corresponding URI can be constructed.
	//
	URI(mimeType string) string
}

// StreamableContent: interface whereby an object allows its backing content to
// be streamed to clients. Typical implementors would be images or icons, HTML
// content, or multimedia display/rendering widgets.
//
// Negotiation of content type is allowed. Clients may examine the backing data
// and transform, convert, or parse the content in order to present it in an
// alternate form to end-users.
//
// The AtkStreamableContent interface is particularly useful for saving,
// printing, or post-processing entire documents, or for persisting alternate
// views of a document. If document content itself is being serialized, stored,
// or converted, then use of the AtkStreamableContent interface can help address
// performance issues. Unlike most ATK interfaces, this interface is not
// strongly tied to the current user-agent view of the a particular document,
// but may in some cases give access to the underlying model data.
type StreamableContent struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*StreamableContent)(nil)
)

// StreamableContenter describes StreamableContent's interface methods.
type StreamableContenter interface {
	externglib.Objector

	// MIMEType gets the character string of the specified mime type.
	MIMEType(i int) string
	// NMIMETypes gets the number of mime types supported by this object.
	NMIMETypes() int
	// Stream gets the content in the specified mime type.
	Stream(mimeType string) *glib.IOChannel
	// URI: get a string representing a URI in IETF standard format (see
	// http://www.ietf.org/rfc/rfc2396.txt) from which the object's content may
	// be streamed in the specified mime-type, if one is available.
	URI(mimeType string) string
}

var _ StreamableContenter = (*StreamableContent)(nil)

func ifaceInitStreamableContenter(gifacePtr, data C.gpointer) {
	iface := (*C.AtkStreamableContentIface)(unsafe.Pointer(gifacePtr))
	iface.get_mime_type = (*[0]byte)(C._gotk4_atk1_StreamableContentIface_get_mime_type)
	iface.get_n_mime_types = (*[0]byte)(C._gotk4_atk1_StreamableContentIface_get_n_mime_types)
	iface.get_stream = (*[0]byte)(C._gotk4_atk1_StreamableContentIface_get_stream)
	iface.get_uri = (*[0]byte)(C._gotk4_atk1_StreamableContentIface_get_uri)
}

//export _gotk4_atk1_StreamableContentIface_get_mime_type
func _gotk4_atk1_StreamableContentIface_get_mime_type(arg0 *C.AtkStreamableContent, arg1 C.gint) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	var _i int // out

	_i = int(arg1)

	utf8 := iface.MIMEType(_i)

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_atk1_StreamableContentIface_get_n_mime_types
func _gotk4_atk1_StreamableContentIface_get_n_mime_types(arg0 *C.AtkStreamableContent) (cret C.gint) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	gint := iface.NMIMETypes()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_StreamableContentIface_get_stream
func _gotk4_atk1_StreamableContentIface_get_stream(arg0 *C.AtkStreamableContent, arg1 *C.gchar) (cret *C.GIOChannel) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	var _mimeType string // out

	_mimeType = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ioChannel := iface.Stream(_mimeType)

	cret = (*C.GIOChannel)(gextras.StructNative(unsafe.Pointer(ioChannel)))

	return cret
}

//export _gotk4_atk1_StreamableContentIface_get_uri
func _gotk4_atk1_StreamableContentIface_get_uri(arg0 *C.AtkStreamableContent, arg1 *C.gchar) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(StreamableContentOverrider)

	var _mimeType string // out

	_mimeType = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := iface.URI(_mimeType)

	if utf8 != "" {
		cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
		defer C.free(unsafe.Pointer(cret))
	}

	return cret
}

func wrapStreamableContent(obj *externglib.Object) *StreamableContent {
	return &StreamableContent{
		Object: obj,
	}
}

func marshalStreamableContent(p uintptr) (interface{}, error) {
	return wrapStreamableContent(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MIMEType gets the character string of the specified mime type. The first mime
// type is at position 0, the second at position 1, and so on.
//
// The function takes the following parameters:
//
//    - i: gint representing the position of the mime type starting from 0.
//
// The function returns the following values:
//
//    - utf8: gchar* representing the specified mime type; the caller should not
//      free the character string.
//
func (streamable *StreamableContent) MIMEType(i int) string {
	var _arg0 *C.AtkStreamableContent // out
	var _arg1 C.gint                  // out
	var _cret *C.gchar                // in

	_arg0 = (*C.AtkStreamableContent)(unsafe.Pointer(streamable.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_streamable_content_get_mime_type(_arg0, _arg1)
	runtime.KeepAlive(streamable)
	runtime.KeepAlive(i)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NMIMETypes gets the number of mime types supported by this object.
//
// The function returns the following values:
//
//    - gint which is the number of mime types supported by the object.
//
func (streamable *StreamableContent) NMIMETypes() int {
	var _arg0 *C.AtkStreamableContent // out
	var _cret C.gint                  // in

	_arg0 = (*C.AtkStreamableContent)(unsafe.Pointer(streamable.Native()))

	_cret = C.atk_streamable_content_get_n_mime_types(_arg0)
	runtime.KeepAlive(streamable)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Stream gets the content in the specified mime type.
//
// The function takes the following parameters:
//
//    - mimeType: gchar* representing the mime type.
//
// The function returns the following values:
//
//    - ioChannel which contains the content in the specified mime type.
//
func (streamable *StreamableContent) Stream(mimeType string) *glib.IOChannel {
	var _arg0 *C.AtkStreamableContent // out
	var _arg1 *C.gchar                // out
	var _cret *C.GIOChannel           // in

	_arg0 = (*C.AtkStreamableContent)(unsafe.Pointer(streamable.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_streamable_content_get_stream(_arg0, _arg1)
	runtime.KeepAlive(streamable)
	runtime.KeepAlive(mimeType)

	var _ioChannel *glib.IOChannel // out

	_ioChannel = (*glib.IOChannel)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ioChannel)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_io_channel_unref((*C.GIOChannel)(intern.C))
		},
	)

	return _ioChannel
}

// URI: get a string representing a URI in IETF standard format (see
// http://www.ietf.org/rfc/rfc2396.txt) from which the object's content may be
// streamed in the specified mime-type, if one is available. If mime_type is
// NULL, the URI for the default (and possibly only) mime-type is returned.
//
// Note that it is possible for get_uri to return NULL but for get_stream to
// work nonetheless, since not all GIOChannels connect to URIs.
//
// The function takes the following parameters:
//
//    - mimeType: gchar* representing the mime type, or NULL to request a URI for
//      the default mime type.
//
// The function returns the following values:
//
//    - utf8 (optional) returns a string representing a URI, or NULL if no
//      corresponding URI can be constructed.
//
func (streamable *StreamableContent) URI(mimeType string) string {
	var _arg0 *C.AtkStreamableContent // out
	var _arg1 *C.gchar                // out
	var _cret *C.gchar                // in

	_arg0 = (*C.AtkStreamableContent)(unsafe.Pointer(streamable.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_streamable_content_get_uri(_arg0, _arg1)
	runtime.KeepAlive(streamable)
	runtime.KeepAlive(mimeType)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}
