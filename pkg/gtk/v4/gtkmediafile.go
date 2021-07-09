// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_media_file_get_type()), F: marshalMediaFile},
	})
}

// MediaFileOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MediaFileOverrider interface {
	Close()
	Open()
}

// MediaFile: `GtkMediaFile` implements `GtkMediaStream` for files.
//
// This provides a simple way to play back video files with GTK.
//
// GTK provides a GIO extension point for `GtkMediaFile` implementations to
// allow for external implementations using various media frameworks.
//
// GTK itself includes implementations using GStreamer and ffmpeg.
type MediaFile interface {
	gextras.Objector

	// Clear resets the media file to be empty.
	Clear()
	// InputStream returns the stream that @self is currently playing from.
	//
	// When @self is not playing or not playing from a stream, nil is returned.
	InputStream() *gio.InputStreamClass
	// SetFilename sets the `GtkMediaFile to play the given file.
	//
	// This is a utility function that converts the given @filename to a `GFile`
	// and calls [method@Gtk.MediaFile.set_file].
	SetFilename(filename string)
	// SetInputStream sets the `GtkMediaFile` to play the given stream.
	//
	// If anything is still playing, stop playing it.
	//
	// Full control about the @stream is assumed for the duration of playback.
	// The stream will not be closed.
	SetInputStream(stream gio.InputStream)
	// SetResource sets the `GtkMediaFile to play the given resource.
	//
	// This is a utility function that converts the given @resource_path to a
	// `GFile` and calls [method@Gtk.MediaFile.set_file].
	SetResource(resourcePath string)
}

// MediaFileClass implements the MediaFile interface.
type MediaFileClass struct {
	MediaStreamClass
}

var _ MediaFile = (*MediaFileClass)(nil)

func wrapMediaFile(obj *externglib.Object) MediaFile {
	return &MediaFileClass{
		MediaStreamClass: MediaStreamClass{
			Object: obj,
		},
	}
}

func marshalMediaFile(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMediaFile(obj), nil
}

// NewMediaFile creates a new empty media file.
func NewMediaFile() *MediaFileClass {
	var _cret *C.GtkMediaStream // in

	_cret = C.gtk_media_file_new()

	var _mediaFile *MediaFileClass // out

	_mediaFile = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*MediaFileClass)

	return _mediaFile
}

// NewMediaFileForFilename creates a new media file for the given filename.
//
// This is a utility function that converts the given @filename to a `GFile` and
// calls [ctor@Gtk.MediaFile.new_for_file].
func NewMediaFileForFilename(filename string) *MediaFileClass {
	var _arg1 *C.char           // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_filename(_arg1)

	var _mediaFile *MediaFileClass // out

	_mediaFile = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*MediaFileClass)

	return _mediaFile
}

// NewMediaFileForInputStream creates a new media file to play @stream.
//
// If you want the resulting media to be seekable, the stream should implement
// the `GSeekable` interface.
func NewMediaFileForInputStream(stream gio.InputStream) *MediaFileClass {
	var _arg1 *C.GInputStream   // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer((&stream).Native()))

	_cret = C.gtk_media_file_new_for_input_stream(_arg1)

	var _mediaFile *MediaFileClass // out

	_mediaFile = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*MediaFileClass)

	return _mediaFile
}

// NewMediaFileForResource creates a new new media file for the given resource.
//
// This is a utility function that converts the given @resource to a `GFile` and
// calls [ctor@Gtk.MediaFile.new_for_file].
func NewMediaFileForResource(resourcePath string) *MediaFileClass {
	var _arg1 *C.char           // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_resource(_arg1)

	var _mediaFile *MediaFileClass // out

	_mediaFile = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*MediaFileClass)

	return _mediaFile
}

// Clear resets the media file to be empty.
func (s *MediaFileClass) Clear() {
	var _arg0 *C.GtkMediaFile // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer((&s).Native()))

	C.gtk_media_file_clear(_arg0)
}

// InputStream returns the stream that @self is currently playing from.
//
// When @self is not playing or not playing from a stream, nil is returned.
func (s *MediaFileClass) InputStream() *gio.InputStreamClass {
	var _arg0 *C.GtkMediaFile // out
	var _cret *C.GInputStream // in

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_media_file_get_input_stream(_arg0)

	var _inputStream *gio.InputStreamClass // out

	_inputStream = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*gio.InputStreamClass)

	return _inputStream
}

// SetFilename sets the `GtkMediaFile to play the given file.
//
// This is a utility function that converts the given @filename to a `GFile` and
// calls [method@Gtk.MediaFile.set_file].
func (s *MediaFileClass) SetFilename(filename string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_media_file_set_filename(_arg0, _arg1)
}

// SetInputStream sets the `GtkMediaFile` to play the given stream.
//
// If anything is still playing, stop playing it.
//
// Full control about the @stream is assumed for the duration of playback. The
// stream will not be closed.
func (s *MediaFileClass) SetInputStream(stream gio.InputStream) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.GInputStream // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GInputStream)(unsafe.Pointer((&stream).Native()))

	C.gtk_media_file_set_input_stream(_arg0, _arg1)
}

// SetResource sets the `GtkMediaFile to play the given resource.
//
// This is a utility function that converts the given @resource_path to a
// `GFile` and calls [method@Gtk.MediaFile.set_file].
func (s *MediaFileClass) SetResource(resourcePath string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_media_file_set_resource(_arg0, _arg1)
}
