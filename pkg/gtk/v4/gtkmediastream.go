// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk4_MediaStreamClass_play(void*);
// extern void _gotk4_gtk4_MediaStreamClass_pause(void*);
// extern void _gotk4_gtk4_MediaStreamClass_realize(void*, void*);
// extern void _gotk4_gtk4_MediaStreamClass_seek(void*, gint64);
// extern void _gotk4_gtk4_MediaStreamClass_unrealize(void*, void*);
// extern void _gotk4_gtk4_MediaStreamClass_update_audio(void*, gboolean, double);
import "C"

// GTypeMediaStream returns the GType for the type MediaStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMediaStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MediaStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMediaStream)
	return gtype
}

// MediaStreamOverrider contains methods that are overridable.
type MediaStreamOverrider interface {
	// Pause pauses playback of the stream.
	//
	// If the stream is not playing, do nothing.
	Pause()
	// The function returns the following values:
	//
	Play() bool
	// Realize: called by users to attach the media stream to a GdkSurface they
	// manage.
	//
	// The stream can then access the resources of surface for its rendering
	// purposes. In particular, media streams might want to create a
	// GdkGLContext or sync to the GdkFrameClock.
	//
	// Whoever calls this function is responsible for calling
	// gtk.MediaStream.Unrealize() before either the stream or surface get
	// destroyed.
	//
	// Multiple calls to this function may happen from different users of the
	// video, even with the same surface. Each of these calls must be followed
	// by its own call to gtk.MediaStream.Unrealize().
	//
	// It is not required to call this function to make a media stream work.
	//
	// The function takes the following parameters:
	//
	//    - surface: GdkSurface.
	//
	Realize(surface gdk.Surfacer)
	// Seek: start a seek operation on self to timestamp.
	//
	// If timestamp is out of range, it will be clamped.
	//
	// Seek operations may not finish instantly. While a seek operation is in
	// process, the gtk.MediaStream:seeking property will be set.
	//
	// When calling gtk_media_stream_seek() during an ongoing seek operation,
	// the new seek will override any pending seek.
	//
	// The function takes the following parameters:
	//
	//    - timestamp to seek to.
	//
	Seek(timestamp int64)
	// Unrealize undoes a previous call to gtk_media_stream_realize().
	//
	// This causes the stream to release all resources it had allocated from
	// surface.
	//
	// The function takes the following parameters:
	//
	//    - surface: GdkSurface the stream was realized with.
	//
	Unrealize(surface gdk.Surfacer)
	// The function takes the following parameters:
	//
	//    - muted
	//    - volume
	//
	UpdateAudio(muted bool, volume float64)
}

// MediaStream: GtkMediaStream is the integration point for media playback
// inside GTK.
//
// GTK provides an implementation of the GtkMediaStream interface that is called
// gtk.MediaFile.
//
// Apart from application-facing API for stream playback, GtkMediaStream has a
// number of APIs that are only useful for implementations and should not be
// used in applications: gtk.MediaStream.Prepared(),
// gtk.MediaStream.Unprepared(), gtk.MediaStream.Update(),
// gtk.MediaStream.Ended(), gtk.MediaStream.SeekSuccess(),
// gtk.MediaStream.SeekFailed(), gtk.MediaStream.GError(),
// gtk.MediaStream.Error(), gtk.MediaStream.ErrorValist().
type MediaStream struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gdk.Paintable
}

var (
	_ coreglib.Objector = (*MediaStream)(nil)
)

// MediaStreamer describes types inherited from class MediaStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type MediaStreamer interface {
	coreglib.Objector
	baseMediaStream() *MediaStream
}

var _ MediaStreamer = (*MediaStream)(nil)

func classInitMediaStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "MediaStreamClass")

	if _, ok := goval.(interface{ Pause() }); ok {
		o := pclass.StructFieldOffset("pause")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaStreamClass_pause)
	}

	if _, ok := goval.(interface{ Play() bool }); ok {
		o := pclass.StructFieldOffset("play")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaStreamClass_play)
	}

	if _, ok := goval.(interface{ Realize(surface gdk.Surfacer) }); ok {
		o := pclass.StructFieldOffset("realize")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaStreamClass_realize)
	}

	if _, ok := goval.(interface{ Seek(timestamp int64) }); ok {
		o := pclass.StructFieldOffset("seek")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaStreamClass_seek)
	}

	if _, ok := goval.(interface{ Unrealize(surface gdk.Surfacer) }); ok {
		o := pclass.StructFieldOffset("unrealize")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaStreamClass_unrealize)
	}

	if _, ok := goval.(interface {
		UpdateAudio(muted bool, volume float64)
	}); ok {
		o := pclass.StructFieldOffset("update_audio")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_MediaStreamClass_update_audio)
	}
}

//export _gotk4_gtk4_MediaStreamClass_pause
func _gotk4_gtk4_MediaStreamClass_pause(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Pause() })

	iface.Pause()
}

//export _gotk4_gtk4_MediaStreamClass_play
func _gotk4_gtk4_MediaStreamClass_play(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Play() bool })

	ok := iface.Play()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_MediaStreamClass_realize
func _gotk4_gtk4_MediaStreamClass_realize(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Realize(surface gdk.Surfacer) })

	var _surface gdk.Surfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Surfacer)
			return ok
		})
		rv, ok := casted.(gdk.Surfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
		}
		_surface = rv
	}

	iface.Realize(_surface)
}

//export _gotk4_gtk4_MediaStreamClass_seek
func _gotk4_gtk4_MediaStreamClass_seek(arg0 *C.void, arg1 C.gint64) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Seek(timestamp int64) })

	var _timestamp int64 // out

	_timestamp = int64(arg1)

	iface.Seek(_timestamp)
}

//export _gotk4_gtk4_MediaStreamClass_unrealize
func _gotk4_gtk4_MediaStreamClass_unrealize(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Unrealize(surface gdk.Surfacer) })

	var _surface gdk.Surfacer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Surfacer)
			return ok
		})
		rv, ok := casted.(gdk.Surfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
		}
		_surface = rv
	}

	iface.Unrealize(_surface)
}

//export _gotk4_gtk4_MediaStreamClass_update_audio
func _gotk4_gtk4_MediaStreamClass_update_audio(arg0 *C.void, arg1 C.gboolean, arg2 C.double) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		UpdateAudio(muted bool, volume float64)
	})

	var _muted bool     // out
	var _volume float64 // out

	if arg1 != 0 {
		_muted = true
	}
	_volume = float64(arg2)

	iface.UpdateAudio(_muted, _volume)
}

func wrapMediaStream(obj *coreglib.Object) *MediaStream {
	return &MediaStream{
		Object: obj,
		Paintable: gdk.Paintable{
			Object: obj,
		},
	}
}

func marshalMediaStream(p uintptr) (interface{}, error) {
	return wrapMediaStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *MediaStream) baseMediaStream() *MediaStream {
	return self
}

// BaseMediaStream returns the underlying base object.
func BaseMediaStream(obj MediaStreamer) *MediaStream {
	return obj.baseMediaStream()
}

// Ended pauses the media stream and marks it as ended.
//
// This is a hint only, calls to GtkMediaStream.play() may still happen.
//
// The media stream must be prepared when this function is called.
func (self *MediaStream) Ended() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("ended", _args[:], nil)

	runtime.KeepAlive(self)
}

// GError sets self into an error state.
//
// This will pause the stream (you can check for an error via
// gtk.MediaStream.GetError() in your GtkMediaStream.pause() implementation),
// abort pending seeks and mark the stream as prepared.
//
// if the stream is already in an error state, this call will be ignored and the
// existing error will be retained.
//
// To unset an error, the stream must be reset via a call to
// gtk.MediaStream.Unprepared().
//
// The function takes the following parameters:
//
//    - err: GError to set.
//
func (self *MediaStream) GError(err error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if err != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gerror.New(err))
	}

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("gerror", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(err)
}

// Duration gets the duration of the stream.
//
// If the duration is not known, 0 will be returned.
//
// The function returns the following values:
//
//    - gint64: duration of the stream or 0 if not known.
//
func (self *MediaStream) Duration() int64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_duration", _args[:], nil)
	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gint64 int64 // out

	_gint64 = int64(*(*C.gint64)(unsafe.Pointer(&_cret)))

	return _gint64
}

// GetEnded returns whether the streams playback is finished.
//
// The function returns the following values:
//
//    - ok: TRUE if playback is finished.
//
func (self *MediaStream) GetEnded() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_ended", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Error: if the stream is in an error state, returns the GError explaining that
// state.
//
// Any type of error can be reported here depending on the implementation of the
// media stream.
//
// A media stream in an error cannot be operated on, calls like
// gtk.MediaStream.Play() or gtk.MediaStream.Seek() will not have any effect.
//
// GtkMediaStream itself does not provide a way to unset an error, but
// implementations may provide options. For example, a gtk.MediaFile will unset
// errors when a new source is set, e.g. with gtk.MediaFile.SetFile().
//
// The function returns the following values:
//
//    - err (optional): NULL if not in an error state or the GError of the
//      stream.
//
func (self *MediaStream) Error() error {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_error", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _err error // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_err = gerror.Take(unsafe.Pointer(_cret))
	}

	return _err
}

// Loop returns whether the stream is set to loop.
//
// See gtk.MediaStream.SetLoop() for details.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream should loop.
//
func (self *MediaStream) Loop() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_loop", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Muted returns whether the audio for the stream is muted.
//
// See gtk.MediaStream.SetMuted() for details.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream is muted.
//
func (self *MediaStream) Muted() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_muted", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Playing: return whether the stream is currently playing.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream is playing.
//
func (self *MediaStream) Playing() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_playing", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Timestamp returns the current presentation timestamp in microseconds.
//
// The function returns the following values:
//
//    - gint64: timestamp in microseconds.
//
func (self *MediaStream) Timestamp() int64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_timestamp", _args[:], nil)
	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gint64 int64 // out

	_gint64 = int64(*(*C.gint64)(unsafe.Pointer(&_cret)))

	return _gint64
}

// Volume returns the volume of the audio for the stream.
//
// See gtk.MediaStream.SetVolume() for details.
//
// The function returns the following values:
//
//    - gdouble: volume of the stream from 0.0 to 1.0.
//
func (self *MediaStream) Volume() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("get_volume", _args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// HasAudio returns whether the stream has audio.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream has audio.
//
func (self *MediaStream) HasAudio() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("has_audio", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// HasVideo returns whether the stream has video.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream has video.
//
func (self *MediaStream) HasVideo() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("has_video", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsPrepared returns whether the stream has finished initializing.
//
// At this point the existence of audio and video is known.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream is prepared.
//
func (self *MediaStream) IsPrepared() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("is_prepared", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsSeekable checks if a stream may be seekable.
//
// This is meant to be a hint. Streams may not allow seeking even if this
// function returns TRUE. However, if this function returns FALSE, streams are
// guaranteed to not be seekable and user interfaces may hide controls that
// allow seeking.
//
// It is allowed to call gtk.MediaStream.Seek() on a non-seekable stream, though
// it will not do anything.
//
// The function returns the following values:
//
//    - ok: TRUE if the stream may support seeking.
//
func (self *MediaStream) IsSeekable() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("is_seekable", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsSeeking checks if there is currently a seek operation going on.
//
// The function returns the following values:
//
//    - ok: TRUE if a seek operation is ongoing.
//
func (self *MediaStream) IsSeeking() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "MediaStream").InvokeMethod("is_seeking", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Pause pauses playback of the stream.
//
// If the stream is not playing, do nothing.
func (self *MediaStream) Pause() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("pause", _args[:], nil)

	runtime.KeepAlive(self)
}

// Play starts playing the stream.
//
// If the stream is in error or already playing, do nothing.
func (self *MediaStream) Play() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("play", _args[:], nil)

	runtime.KeepAlive(self)
}

// Prepared: called by GtkMediaStream implementations to advertise the stream
// being ready to play and providing details about the stream.
//
// Note that the arguments are hints. If the stream implementation cannot
// determine the correct values, it is better to err on the side of caution and
// return TRUE. User interfaces will use those values to determine what controls
// to show.
//
// This function may not be called again until the stream has been reset via
// gtk.MediaStream.Unprepared().
//
// The function takes the following parameters:
//
//    - hasAudio: TRUE if the stream should advertise audio support.
//    - hasVideo: TRUE if the stream should advertise video support.
//    - seekable: TRUE if the stream should advertise seekability.
//    - duration of the stream or 0 if unknown.
//
func (self *MediaStream) Prepared(hasAudio, hasVideo, seekable bool, duration int64) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if hasAudio {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}
	if hasVideo {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}
	if seekable {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}
	*(*C.gint64)(unsafe.Pointer(&_args[4])) = C.gint64(duration)

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("prepared", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(hasAudio)
	runtime.KeepAlive(hasVideo)
	runtime.KeepAlive(seekable)
	runtime.KeepAlive(duration)
}

// Realize: called by users to attach the media stream to a GdkSurface they
// manage.
//
// The stream can then access the resources of surface for its rendering
// purposes. In particular, media streams might want to create a GdkGLContext or
// sync to the GdkFrameClock.
//
// Whoever calls this function is responsible for calling
// gtk.MediaStream.Unrealize() before either the stream or surface get
// destroyed.
//
// Multiple calls to this function may happen from different users of the video,
// even with the same surface. Each of these calls must be followed by its own
// call to gtk.MediaStream.Unrealize().
//
// It is not required to call this function to make a media stream work.
//
// The function takes the following parameters:
//
//    - surface: GdkSurface.
//
func (self *MediaStream) Realize(surface gdk.Surfacer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("realize", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(surface)
}

// Seek: start a seek operation on self to timestamp.
//
// If timestamp is out of range, it will be clamped.
//
// Seek operations may not finish instantly. While a seek operation is in
// process, the gtk.MediaStream:seeking property will be set.
//
// When calling gtk_media_stream_seek() during an ongoing seek operation, the
// new seek will override any pending seek.
//
// The function takes the following parameters:
//
//    - timestamp to seek to.
//
func (self *MediaStream) Seek(timestamp int64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.gint64)(unsafe.Pointer(&_args[1])) = C.gint64(timestamp)

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("seek", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(timestamp)
}

// SeekFailed ends a seek operation started via GtkMediaStream.seek() as a
// failure.
//
// This will not cause an error on the stream and will assume that playback
// continues as if no seek had happened.
//
// See gtk.MediaStream.SeekSuccess() for the other way of ending a seek.
func (self *MediaStream) SeekFailed() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("seek_failed", _args[:], nil)

	runtime.KeepAlive(self)
}

// SeekSuccess ends a seek operation started via GtkMediaStream.seek()
// successfully.
//
// This function will unset the GtkMediaStream:ended property if it was set.
//
// See gtk.MediaStream.SeekFailed() for the other way of ending a seek.
func (self *MediaStream) SeekSuccess() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("seek_success", _args[:], nil)

	runtime.KeepAlive(self)
}

// SetLoop sets whether the stream should loop.
//
// In this case, it will attempt to restart playback from the beginning instead
// of stopping at the end.
//
// Not all streams may support looping, in particular non-seekable streams.
// Those streams will ignore the loop setting and just end.
//
// The function takes the following parameters:
//
//    - loop: TRUE if the stream should loop.
//
func (self *MediaStream) SetLoop(loop bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if loop {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("set_loop", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(loop)
}

// SetMuted sets whether the audio stream should be muted.
//
// Muting a stream will cause no audio to be played, but it does not modify the
// volume. This means that muting and then unmuting the stream will restore the
// volume settings.
//
// If the stream has no audio, calling this function will still work but it will
// not have an audible effect.
//
// The function takes the following parameters:
//
//    - muted: TRUE if the stream should be muted.
//
func (self *MediaStream) SetMuted(muted bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if muted {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("set_muted", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(muted)
}

// SetPlaying starts or pauses playback of the stream.
//
// The function takes the following parameters:
//
//    - playing: whether to start or pause playback.
//
func (self *MediaStream) SetPlaying(playing bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if playing {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("set_playing", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(playing)
}

// SetVolume sets the volume of the audio stream.
//
// This function call will work even if the stream is muted.
//
// The given volume should range from 0.0 for silence to 1.0 for as loud as
// possible. Values outside of this range will be clamped to the nearest value.
//
// If the stream has no audio or is muted, calling this function will still work
// but it will not have an immediate audible effect. When the stream is unmuted,
// the new volume setting will take effect.
//
// The function takes the following parameters:
//
//    - volume: new volume of the stream from 0.0 to 1.0.
//
func (self *MediaStream) SetVolume(volume float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(volume)

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("set_volume", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(volume)
}

// Unprepared resets a given media stream implementation.
//
// gtk.MediaStream.Prepared() can then be called again.
//
// This function will also reset any error state the stream was in.
func (self *MediaStream) Unprepared() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("unprepared", _args[:], nil)

	runtime.KeepAlive(self)
}

// Unrealize undoes a previous call to gtk_media_stream_realize().
//
// This causes the stream to release all resources it had allocated from
// surface.
//
// The function takes the following parameters:
//
//    - surface: GdkSurface the stream was realized with.
//
func (self *MediaStream) Unrealize(surface gdk.Surfacer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("unrealize", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(surface)
}

// Update: media stream implementations should regularly call this function to
// update the timestamp reported by the stream.
//
// It is up to implementations to call this at the frequency they deem
// appropriate.
//
// The media stream must be prepared when this function is called.
//
// The function takes the following parameters:
//
//    - timestamp: new timestamp.
//
func (self *MediaStream) Update(timestamp int64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.gint64)(unsafe.Pointer(&_args[1])) = C.gint64(timestamp)

	girepository.MustFind("Gtk", "MediaStream").InvokeMethod("update", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(timestamp)
}
