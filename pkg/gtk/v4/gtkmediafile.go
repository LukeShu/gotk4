// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_MediaFileClass_close(void*);
// extern void _gotk4_gtk4_MediaFileClass_open(void*);
import "C"

// GTypeMediaFile returns the GType for the type MediaFile.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMediaFile() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MediaFile").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMediaFile)
	return gtype
}

const MEDIA_FILE_EXTENSION_POINT_NAME = "gtk-media-file"

// MediaFileOverrider contains methods that are overridable.
type MediaFileOverrider interface {
	Close()
	Open()
}

// MediaFile: GtkMediaFile implements GtkMediaStream for files.
//
// This provides a simple way to play back video files with GTK.
//
// GTK provides a GIO extension point for GtkMediaFile implementations to allow
// for external implementations using various media frameworks.
//
// GTK itself includes implementations using GStreamer and ffmpeg.
type MediaFile struct {
	_ [0]func() // equal guard
	MediaStream
}

var (
	_ MediaStreamer = (*MediaFile)(nil)
)

// MediaFiler describes types inherited from class MediaFile.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MediaFiler interface {
	coreglib.Objector
	baseMediaFile() *MediaFile
}

var _ MediaFiler = (*MediaFile)(nil)

func classInitMediaFiler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "MediaFileClass")

	if _, ok := goval.(interface{ Close() }); ok {
		o := pclass.StructFieldOffset("close")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaFileClass_close)
	}

	if _, ok := goval.(interface{ Open() }); ok {
		o := pclass.StructFieldOffset("open")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaFileClass_open)
	}
}

//export _gotk4_gtk4_MediaFileClass_close
func _gotk4_gtk4_MediaFileClass_close(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Close() })

	iface.Close()
}

//export _gotk4_gtk4_MediaFileClass_open
func _gotk4_gtk4_MediaFileClass_open(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Open() })

	iface.Open()
}

func wrapMediaFile(obj *coreglib.Object) *MediaFile {
	return &MediaFile{
		MediaStream: MediaStream{
			Object: obj,
			Paintable: gdk.Paintable{
				Object: obj,
			},
		},
	}
}

func marshalMediaFile(p uintptr) (interface{}, error) {
	return wrapMediaFile(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *MediaFile) baseMediaFile() *MediaFile {
	return self
}

// BaseMediaFile returns the underlying base object.
func BaseMediaFile(obj MediaFiler) *MediaFile {
	return obj.baseMediaFile()
}

// NewMediaFile creates a new empty media file.
//
// The function returns the following values:
//
//    - mediaFile: new GtkMediaFile.
//
func NewMediaFile() *MediaFile {
	_info := girepository.MustFind("Gtk", "MediaFile")
	_gret := _info.InvokeClassMethod("new_MediaFile", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _mediaFile *MediaFile // out

	_mediaFile = wrapMediaFile(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _mediaFile
}

// NewMediaFileForFile creates a new media file to play file.
//
// The function takes the following parameters:
//
//    - file to play.
//
// The function returns the following values:
//
//    - mediaFile: new GtkMediaFile playing file.
//
func NewMediaFileForFile(file gio.Filer) *MediaFile {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_info := girepository.MustFind("Gtk", "MediaFile")
	_gret := _info.InvokeClassMethod("new_MediaFile_for_file", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(file)

	var _mediaFile *MediaFile // out

	_mediaFile = wrapMediaFile(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _mediaFile
}

// NewMediaFileForFilename creates a new media file for the given filename.
//
// This is a utility function that converts the given filename to a GFile and
// calls gtk.MediaFile.NewForFile.
//
// The function takes the following parameters:
//
//    - filename to open.
//
// The function returns the following values:
//
//    - mediaFile: new GtkMediaFile playing filename.
//
func NewMediaFileForFilename(filename string) *MediaFile {
	var _args [1]girepository.Argument

	*(**C.char)(unsafe.Pointer(&_args[0])) = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "MediaFile")
	_gret := _info.InvokeClassMethod("new_MediaFile_for_filename", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(filename)

	var _mediaFile *MediaFile // out

	_mediaFile = wrapMediaFile(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _mediaFile
}

// NewMediaFileForInputStream creates a new media file to play stream.
//
// If you want the resulting media to be seekable, the stream should implement
// the GSeekable interface.
//
// The function takes the following parameters:
//
//    - stream to play.
//
// The function returns the following values:
//
//    - mediaFile: new GtkMediaFile.
//
func NewMediaFileForInputStream(stream gio.InputStreamer) *MediaFile {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gtk", "MediaFile")
	_gret := _info.InvokeClassMethod("new_MediaFile_for_input_stream", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _mediaFile *MediaFile // out

	_mediaFile = wrapMediaFile(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _mediaFile
}

// NewMediaFileForResource creates a new new media file for the given resource.
//
// This is a utility function that converts the given resource to a GFile and
// calls gtk.MediaFile.NewForFile.
//
// The function takes the following parameters:
//
//    - resourcePath: resource path to open.
//
// The function returns the following values:
//
//    - mediaFile: new GtkMediaFile playing resource_path.
//
func NewMediaFileForResource(resourcePath string) *MediaFile {
	var _args [1]girepository.Argument

	*(**C.char)(unsafe.Pointer(&_args[0])) = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "MediaFile")
	_gret := _info.InvokeClassMethod("new_MediaFile_for_resource", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(resourcePath)

	var _mediaFile *MediaFile // out

	_mediaFile = wrapMediaFile(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _mediaFile
}

// Clear resets the media file to be empty.
func (self *MediaFile) Clear() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "MediaFile")
	_info.InvokeClassMethod("clear", _args[:], nil)

	runtime.KeepAlive(self)
}

// File returns the file that self is currently playing from.
//
// When self is not playing or not playing from a file, NULL is returned.
//
// The function returns the following values:
//
//    - file (optional): currently playing file or NULL if not playing from a
//      file.
//
func (self *MediaFile) File() *gio.File {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "MediaFile")
	_gret := _info.InvokeClassMethod("get_file", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _file *gio.File // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
			_file = &gio.File{
				Object: obj,
			}
		}
	}

	return _file
}

// InputStream returns the stream that self is currently playing from.
//
// When self is not playing or not playing from a stream, NULL is returned.
//
// The function returns the following values:
//
//    - inputStream (optional): currently playing stream or NULL if not playing
//      from a stream.
//
func (self *MediaFile) InputStream() gio.InputStreamer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "MediaFile")
	_gret := _info.InvokeClassMethod("get_input_stream", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _inputStream gio.InputStreamer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.InputStreamer)
				return ok
			})
			rv, ok := casted.(gio.InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}

	return _inputStream
}

// SetFile sets the GtkMediaFile to play the given file.
//
// If any file is still playing, stop playing it.
//
// The function takes the following parameters:
//
//    - file (optional) to play.
//
func (self *MediaFile) SetFile(file gio.Filer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}

	_info := girepository.MustFind("Gtk", "MediaFile")
	_info.InvokeClassMethod("set_file", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(file)
}

// SetFilename sets the `GtkMediaFile to play the given file.
//
// This is a utility function that converts the given filename to a GFile and
// calls gtk.MediaFile.SetFile().
//
// The function takes the following parameters:
//
//    - filename (optional): name of file to play.
//
func (self *MediaFile) SetFilename(filename string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if filename != "" {
		*(**C.char)(unsafe.Pointer(&_args[1])) = (*C.char)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[1]))))
	}

	_info := girepository.MustFind("Gtk", "MediaFile")
	_info.InvokeClassMethod("set_filename", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(filename)
}

// SetInputStream sets the GtkMediaFile to play the given stream.
//
// If anything is still playing, stop playing it.
//
// Full control about the stream is assumed for the duration of playback. The
// stream will not be closed.
//
// The function takes the following parameters:
//
//    - stream (optional) to play from.
//
func (self *MediaFile) SetInputStream(stream gio.InputStreamer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if stream != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	}

	_info := girepository.MustFind("Gtk", "MediaFile")
	_info.InvokeClassMethod("set_input_stream", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(stream)
}

// SetResource sets the `GtkMediaFile to play the given resource.
//
// This is a utility function that converts the given resource_path to a GFile
// and calls gtk.MediaFile.SetFile().
//
// The function takes the following parameters:
//
//    - resourcePath (optional): path to resource to play.
//
func (self *MediaFile) SetResource(resourcePath string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if resourcePath != "" {
		*(**C.char)(unsafe.Pointer(&_args[1])) = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[1]))))
	}

	_info := girepository.MustFind("Gtk", "MediaFile")
	_info.InvokeClassMethod("set_resource", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(resourcePath)
}
