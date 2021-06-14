// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_clipboard_get_type()), F: marshalClipboard},
	})
}

// Clipboard: the `GdkClipboard` object represents data shared between
// applications or inside an application.
//
// To get a `GdkClipboard` object, use [method@Gdk.Display.get_clipboard] or
// [method@Gdk.Display.get_primary_clipboard]. You can find out about the data
// that is currently available in a clipboard using
// [method@Gdk.Clipboard.get_formats].
//
// To make text or image data available in a clipboard, use
// [method@Gdk.Clipboard.set_text] or [method@Gdk.Clipboard.set_texture]. For
// other data, you can use [method@Gdk.Clipboard.set_content], which takes a
// [class@Gdk.ContentProvider] object.
//
// To read textual or image data from a clipboard, use
// [method@Gdk.Clipboard.read_text_async] or
// [method@Gdk.Clipboard.read_texture_async]. For other data, use
// [method@Gdk.Clipboard.read_async], which provides a `GInputStream` object.
type Clipboard interface {
	gextras.Objector

	// Content returns the `GdkContentProvider` currently set on @clipboard.
	//
	// If the @clipboard is empty or its contents are not owned by the current
	// process, nil will be returned.
	Content() ContentProvider
	// Display gets the `GdkDisplay` that the clipboard was created for.
	Display() Display
	// Formats gets the formats that the clipboard can provide its current
	// contents in.
	Formats() *ContentFormats
	// IsLocal returns if the clipboard is local.
	//
	// A clipboard is considered local if it was last claimed by the running
	// application.
	//
	// Note that [method@Gdk.Clipboard.get_content] may return nil even on a
	// local clipboard. In this case the clipboard is empty.
	IsLocal() bool
	// SetContent sets a new content provider on @clipboard.
	//
	// The clipboard will claim the `GdkDisplay`'s resources and advertise these
	// new contents to other applications.
	//
	// In the rare case of a failure, this function will return false. The
	// clipboard will then continue reporting its old contents and ignore
	// @provider.
	//
	// If the contents are read by either an external application or the
	// @clipboard's read functions, @clipboard will select the best format to
	// transfer the contents and then request that format from @provider.
	SetContent(provider ContentProvider) bool
}

// clipboard implements the Clipboard class.
type clipboard struct {
	gextras.Objector
}

var _ Clipboard = (*clipboard)(nil)

// WrapClipboard wraps a GObject to the right type. It is
// primarily used internally.
func WrapClipboard(obj *externglib.Object) Clipboard {
	return clipboard{
		Objector: obj,
	}
}

func marshalClipboard(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapClipboard(obj), nil
}

// Content returns the `GdkContentProvider` currently set on @clipboard.
//
// If the @clipboard is empty or its contents are not owned by the current
// process, nil will be returned.
func (c clipboard) Content() ContentProvider {
	var _arg0 *C.GdkClipboard // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	var _cret *C.GdkContentProvider // in

	_cret = C.gdk_clipboard_get_content(_arg0)

	var _contentProvider ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ContentProvider)

	return _contentProvider
}

// Display gets the `GdkDisplay` that the clipboard was created for.
func (c clipboard) Display() Display {
	var _arg0 *C.GdkClipboard // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	var _cret *C.GdkDisplay // in

	_cret = C.gdk_clipboard_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Display)

	return _display
}

// Formats gets the formats that the clipboard can provide its current
// contents in.
func (c clipboard) Formats() *ContentFormats {
	var _arg0 *C.GdkClipboard // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	var _cret *C.GdkContentFormats // in

	_cret = C.gdk_clipboard_get_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = WrapContentFormats(unsafe.Pointer(_cret))

	return _contentFormats
}

// IsLocal returns if the clipboard is local.
//
// A clipboard is considered local if it was last claimed by the running
// application.
//
// Note that [method@Gdk.Clipboard.get_content] may return nil even on a
// local clipboard. In this case the clipboard is empty.
func (c clipboard) IsLocal() bool {
	var _arg0 *C.GdkClipboard // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_clipboard_is_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetContent sets a new content provider on @clipboard.
//
// The clipboard will claim the `GdkDisplay`'s resources and advertise these
// new contents to other applications.
//
// In the rare case of a failure, this function will return false. The
// clipboard will then continue reporting its old contents and ignore
// @provider.
//
// If the contents are read by either an external application or the
// @clipboard's read functions, @clipboard will select the best format to
// transfer the contents and then request that format from @provider.
func (c clipboard) SetContent(provider ContentProvider) bool {
	var _arg0 *C.GdkClipboard       // out
	var _arg1 *C.GdkContentProvider // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_clipboard_set_content(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
