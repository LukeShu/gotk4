// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_animation_get_type()), F: marshalPixbufAnimationer},
		{T: externglib.Type(C.gdk_pixbuf_animation_iter_get_type()), F: marshalPixbufAnimationIterer},
	})
}

// PixbufAnimationOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PixbufAnimationOverrider interface {
	// Iter: get an iterator for displaying an animation.
	//
	// The iterator provides the frames that should be displayed at a given
	// time.
	//
	// @start_time would normally come from g_get_current_time(), and marks the
	// beginning of animation playback. After creating an iterator, you should
	// immediately display the pixbuf returned by
	// gdk_pixbuf_animation_iter_get_pixbuf(). Then, you should install a
	// timeout (with g_timeout_add()) or by some other mechanism ensure that
	// you'll update the image after gdk_pixbuf_animation_iter_get_delay_time()
	// milliseconds. Each time the image is updated, you should reinstall the
	// timeout with the new, possibly-changed delay time.
	//
	// As a shortcut, if @start_time is `NULL`, the result of
	// g_get_current_time() will be used automatically.
	//
	// To update the image (i.e. possibly change the result of
	// gdk_pixbuf_animation_iter_get_pixbuf() to a new frame of the animation),
	// call gdk_pixbuf_animation_iter_advance().
	//
	// If you're using PixbufLoader, in addition to updating the image after the
	// delay time, you should also update it whenever you receive the
	// area_updated signal and
	// gdk_pixbuf_animation_iter_on_currently_loading_frame() returns `TRUE`. In
	// this case, the frame currently being fed into the loader has received new
	// data, so needs to be refreshed. The delay time for a frame may also be
	// modified after an area_updated signal, for example if the delay time for
	// a frame is encoded in the data after the frame itself. So your timeout
	// should be reinstalled after any area_updated signal.
	//
	// A delay time of -1 is possible, indicating "infinite".
	Iter(startTime *glib.TimeVal) *PixbufAnimationIter
	Size(width *int, height *int)
	// StaticImage retrieves a static image for the animation.
	//
	// If an animation is really just a plain image (has only one frame), this
	// function returns that image.
	//
	// If the animation is an animation, this function returns a reasonable
	// image to use as a static unanimated image, which might be the first
	// frame, or something more sophisticated depending on the file format.
	//
	// If an animation hasn't loaded any frames yet, this function will return
	// `NULL`.
	StaticImage() *Pixbuf
	// IsStaticImage checks whether the animation is a static image.
	//
	// If you load a file with gdk_pixbuf_animation_new_from_file() and it turns
	// out to be a plain, unanimated image, then this function will return
	// `TRUE`. Use gdk_pixbuf_animation_get_static_image() to retrieve the
	// image.
	IsStaticImage() bool
}

// PixbufAnimationer describes PixbufAnimation's methods.
type PixbufAnimationer interface {
	// Height queries the height of the bounding box of a pixbuf animation.
	Height() int
	// Iter: get an iterator for displaying an animation.
	Iter(startTime *glib.TimeVal) *PixbufAnimationIter
	// StaticImage retrieves a static image for the animation.
	StaticImage() *Pixbuf
	// Width queries the width of the bounding box of a pixbuf animation.
	Width() int
	// IsStaticImage checks whether the animation is a static image.
	IsStaticImage() bool
}

// PixbufAnimation: opaque object representing an animation.
//
// The GdkPixBuf library provides a simple mechanism to load and represent
// animations. An animation is conceptually a series of frames to be displayed
// over time.
//
// The animation may not be represented as a series of frames internally; for
// example, it may be stored as a sprite and instructions for moving the sprite
// around a background.
//
// To display an animation you don't need to understand its representation,
// however; you just ask `GdkPixbuf` what should be displayed at a given point
// in time.
type PixbufAnimation struct {
	*externglib.Object
}

var (
	_ PixbufAnimationer = (*PixbufAnimation)(nil)
	_ gextras.Nativer   = (*PixbufAnimation)(nil)
)

func wrapPixbufAnimation(obj *externglib.Object) PixbufAnimationer {
	return &PixbufAnimation{
		Object: obj,
	}
}

func marshalPixbufAnimationer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPixbufAnimation(obj), nil
}

// NewPixbufAnimationFromFile creates a new animation by loading it from a file.
//
// The file format is detected automatically.
//
// If the file's format does not support multi-frame images, then an animation
// with a single frame will be created.
//
// Possible errors are in the `GDK_PIXBUF_ERROR` and `G_FILE_ERROR` domains.
func NewPixbufAnimationFromFile(filename string) (*PixbufAnimation, error) {
	var _arg1 *C.char               // out
	var _cret *C.GdkPixbufAnimation // in
	var _cerr *C.GError             // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_animation_new_from_file(_arg1, &_cerr)

	var _pixbufAnimation *PixbufAnimation // out
	var _goerr error                      // out

	_pixbufAnimation = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PixbufAnimation)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbufAnimation, _goerr
}

// NewPixbufAnimationFromResource creates a new pixbuf animation by loading an
// image from an resource.
//
// The file format is detected automatically. If `NULL` is returned, then @error
// will be set.
func NewPixbufAnimationFromResource(resourcePath string) (*PixbufAnimation, error) {
	var _arg1 *C.char               // out
	var _cret *C.GdkPixbufAnimation // in
	var _cerr *C.GError             // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_animation_new_from_resource(_arg1, &_cerr)

	var _pixbufAnimation *PixbufAnimation // out
	var _goerr error                      // out

	_pixbufAnimation = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PixbufAnimation)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbufAnimation, _goerr
}

// NewPixbufAnimationFromStream creates a new animation by loading it from an
// input stream.
//
// The file format is detected automatically.
//
// If `NULL` is returned, then @error will be set.
//
// The @cancellable can be used to abort the operation from another thread. If
// the operation was cancelled, the error `G_IO_ERROR_CANCELLED` will be
// returned. Other possible errors are in the `GDK_PIXBUF_ERROR` and
// `G_IO_ERROR` domains.
//
// The stream is not closed.
func NewPixbufAnimationFromStream(stream gio.InputStreamer, cancellable gio.Cancellabler) (*PixbufAnimation, error) {
	var _arg1 *C.GInputStream       // out
	var _arg2 *C.GCancellable       // out
	var _cret *C.GdkPixbufAnimation // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer((stream).(gextras.Nativer).Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer((cancellable).(gextras.Nativer).Native()))

	_cret = C.gdk_pixbuf_animation_new_from_stream(_arg1, _arg2, &_cerr)

	var _pixbufAnimation *PixbufAnimation // out
	var _goerr error                      // out

	_pixbufAnimation = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PixbufAnimation)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbufAnimation, _goerr
}

// NewPixbufAnimationFromStreamFinish finishes an asynchronous pixbuf animation
// creation operation started with
// [func@GdkPixbuf.PixbufAnimation.new_from_stream_async].
func NewPixbufAnimationFromStreamFinish(asyncResult gio.AsyncResulter) (*PixbufAnimation, error) {
	var _arg1 *C.GAsyncResult       // out
	var _cret *C.GdkPixbufAnimation // in
	var _cerr *C.GError             // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((asyncResult).(gextras.Nativer).Native()))

	_cret = C.gdk_pixbuf_animation_new_from_stream_finish(_arg1, &_cerr)

	var _pixbufAnimation *PixbufAnimation // out
	var _goerr error                      // out

	_pixbufAnimation = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PixbufAnimation)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbufAnimation, _goerr
}

// Height queries the height of the bounding box of a pixbuf animation.
func (animation *PixbufAnimation) Height() int {
	var _arg0 *C.GdkPixbufAnimation // out
	var _cret C.int                 // in

	_arg0 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	_cret = C.gdk_pixbuf_animation_get_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Iter: get an iterator for displaying an animation.
//
// The iterator provides the frames that should be displayed at a given time.
//
// @start_time would normally come from g_get_current_time(), and marks the
// beginning of animation playback. After creating an iterator, you should
// immediately display the pixbuf returned by
// gdk_pixbuf_animation_iter_get_pixbuf(). Then, you should install a timeout
// (with g_timeout_add()) or by some other mechanism ensure that you'll update
// the image after gdk_pixbuf_animation_iter_get_delay_time() milliseconds. Each
// time the image is updated, you should reinstall the timeout with the new,
// possibly-changed delay time.
//
// As a shortcut, if @start_time is `NULL`, the result of g_get_current_time()
// will be used automatically.
//
// To update the image (i.e. possibly change the result of
// gdk_pixbuf_animation_iter_get_pixbuf() to a new frame of the animation), call
// gdk_pixbuf_animation_iter_advance().
//
// If you're using PixbufLoader, in addition to updating the image after the
// delay time, you should also update it whenever you receive the area_updated
// signal and gdk_pixbuf_animation_iter_on_currently_loading_frame() returns
// `TRUE`. In this case, the frame currently being fed into the loader has
// received new data, so needs to be refreshed. The delay time for a frame may
// also be modified after an area_updated signal, for example if the delay time
// for a frame is encoded in the data after the frame itself. So your timeout
// should be reinstalled after any area_updated signal.
//
// A delay time of -1 is possible, indicating "infinite".
func (animation *PixbufAnimation) Iter(startTime *glib.TimeVal) *PixbufAnimationIter {
	var _arg0 *C.GdkPixbufAnimation     // out
	var _arg1 *C.GTimeVal               // out
	var _cret *C.GdkPixbufAnimationIter // in

	_arg0 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))
	_arg1 = (*C.GTimeVal)(unsafe.Pointer(startTime))

	_cret = C.gdk_pixbuf_animation_get_iter(_arg0, _arg1)

	var _pixbufAnimationIter *PixbufAnimationIter // out

	_pixbufAnimationIter = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PixbufAnimationIter)

	return _pixbufAnimationIter
}

// StaticImage retrieves a static image for the animation.
//
// If an animation is really just a plain image (has only one frame), this
// function returns that image.
//
// If the animation is an animation, this function returns a reasonable image to
// use as a static unanimated image, which might be the first frame, or
// something more sophisticated depending on the file format.
//
// If an animation hasn't loaded any frames yet, this function will return
// `NULL`.
func (animation *PixbufAnimation) StaticImage() *Pixbuf {
	var _arg0 *C.GdkPixbufAnimation // out
	var _cret *C.GdkPixbuf          // in

	_arg0 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	_cret = C.gdk_pixbuf_animation_get_static_image(_arg0)

	var _pixbuf *Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Pixbuf)

	return _pixbuf
}

// Width queries the width of the bounding box of a pixbuf animation.
func (animation *PixbufAnimation) Width() int {
	var _arg0 *C.GdkPixbufAnimation // out
	var _cret C.int                 // in

	_arg0 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	_cret = C.gdk_pixbuf_animation_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsStaticImage checks whether the animation is a static image.
//
// If you load a file with gdk_pixbuf_animation_new_from_file() and it turns out
// to be a plain, unanimated image, then this function will return `TRUE`. Use
// gdk_pixbuf_animation_get_static_image() to retrieve the image.
func (animation *PixbufAnimation) IsStaticImage() bool {
	var _arg0 *C.GdkPixbufAnimation // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	_cret = C.gdk_pixbuf_animation_is_static_image(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PixbufAnimationIterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PixbufAnimationIterOverrider interface {
	// Advance: possibly advances an animation to a new frame.
	//
	// Chooses the frame based on the start time passed to
	// gdk_pixbuf_animation_get_iter().
	//
	// @current_time would normally come from g_get_current_time(), and must be
	// greater than or equal to the time passed to
	// gdk_pixbuf_animation_get_iter(), and must increase or remain unchanged
	// each time gdk_pixbuf_animation_iter_get_pixbuf() is called. That is, you
	// can't go backward in time; animations only play forward.
	//
	// As a shortcut, pass `NULL` for the current time and g_get_current_time()
	// will be invoked on your behalf. So you only need to explicitly pass
	// @current_time if you're doing something odd like playing the animation at
	// double speed.
	//
	// If this function returns `FALSE`, there's no need to update the animation
	// display, assuming the display had been rendered prior to advancing; if
	// `TRUE`, you need to call gdk_pixbuf_animation_iter_get_pixbuf() and
	// update the display with the new pixbuf.
	Advance(currentTime *glib.TimeVal) bool
	// DelayTime gets the number of milliseconds the current pixbuf should be
	// displayed, or -1 if the current pixbuf should be displayed forever.
	//
	// The `g_timeout_add()` function conveniently takes a timeout in
	// milliseconds, so you can use a timeout to schedule the next update.
	//
	// Note that some formats, like GIF, might clamp the timeout values in the
	// image file to avoid updates that are just too quick. The minimum timeout
	// for GIF images is currently 20 milliseconds.
	DelayTime() int
	// Pixbuf gets the current pixbuf which should be displayed.
	//
	// The pixbuf might not be the same size as the animation itself
	// (gdk_pixbuf_animation_get_width(), gdk_pixbuf_animation_get_height()).
	//
	// This pixbuf should be displayed for
	// gdk_pixbuf_animation_iter_get_delay_time() milliseconds.
	//
	// The caller of this function does not own a reference to the returned
	// pixbuf; the returned pixbuf will become invalid when the iterator
	// advances to the next frame, which may happen anytime you call
	// gdk_pixbuf_animation_iter_advance().
	//
	// Copy the pixbuf to keep it (don't just add a reference), as it may get
	// recycled as you advance the iterator.
	Pixbuf() *Pixbuf
	// OnCurrentlyLoadingFrame: used to determine how to respond to the
	// area_updated signal on PixbufLoader when loading an animation.
	//
	// The `::area_updated` signal is emitted for an area of the frame currently
	// streaming in to the loader. So if you're on the currently loading frame,
	// you will need to redraw the screen for the updated area.
	OnCurrentlyLoadingFrame() bool
}

// PixbufAnimationIterer describes PixbufAnimationIter's methods.
type PixbufAnimationIterer interface {
	// Advance: possibly advances an animation to a new frame.
	Advance(currentTime *glib.TimeVal) bool
	// DelayTime gets the number of milliseconds the current pixbuf should be
	// displayed, or -1 if the current pixbuf should be displayed forever.
	DelayTime() int
	// Pixbuf gets the current pixbuf which should be displayed.
	Pixbuf() *Pixbuf
	// OnCurrentlyLoadingFrame: used to determine how to respond to the
	// area_updated signal on PixbufLoader when loading an animation.
	OnCurrentlyLoadingFrame() bool
}

// PixbufAnimationIter: opaque object representing an iterator which points to a
// certain position in an animation.
type PixbufAnimationIter struct {
	*externglib.Object
}

var (
	_ PixbufAnimationIterer = (*PixbufAnimationIter)(nil)
	_ gextras.Nativer       = (*PixbufAnimationIter)(nil)
)

func wrapPixbufAnimationIter(obj *externglib.Object) PixbufAnimationIterer {
	return &PixbufAnimationIter{
		Object: obj,
	}
}

func marshalPixbufAnimationIterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPixbufAnimationIter(obj), nil
}

// Advance: possibly advances an animation to a new frame.
//
// Chooses the frame based on the start time passed to
// gdk_pixbuf_animation_get_iter().
//
// @current_time would normally come from g_get_current_time(), and must be
// greater than or equal to the time passed to gdk_pixbuf_animation_get_iter(),
// and must increase or remain unchanged each time
// gdk_pixbuf_animation_iter_get_pixbuf() is called. That is, you can't go
// backward in time; animations only play forward.
//
// As a shortcut, pass `NULL` for the current time and g_get_current_time() will
// be invoked on your behalf. So you only need to explicitly pass @current_time
// if you're doing something odd like playing the animation at double speed.
//
// If this function returns `FALSE`, there's no need to update the animation
// display, assuming the display had been rendered prior to advancing; if
// `TRUE`, you need to call gdk_pixbuf_animation_iter_get_pixbuf() and update
// the display with the new pixbuf.
func (iter *PixbufAnimationIter) Advance(currentTime *glib.TimeVal) bool {
	var _arg0 *C.GdkPixbufAnimationIter // out
	var _arg1 *C.GTimeVal               // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter.Native()))
	_arg1 = (*C.GTimeVal)(unsafe.Pointer(currentTime))

	_cret = C.gdk_pixbuf_animation_iter_advance(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DelayTime gets the number of milliseconds the current pixbuf should be
// displayed, or -1 if the current pixbuf should be displayed forever.
//
// The `g_timeout_add()` function conveniently takes a timeout in milliseconds,
// so you can use a timeout to schedule the next update.
//
// Note that some formats, like GIF, might clamp the timeout values in the image
// file to avoid updates that are just too quick. The minimum timeout for GIF
// images is currently 20 milliseconds.
func (iter *PixbufAnimationIter) DelayTime() int {
	var _arg0 *C.GdkPixbufAnimationIter // out
	var _cret C.int                     // in

	_arg0 = (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter.Native()))

	_cret = C.gdk_pixbuf_animation_iter_get_delay_time(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Pixbuf gets the current pixbuf which should be displayed.
//
// The pixbuf might not be the same size as the animation itself
// (gdk_pixbuf_animation_get_width(), gdk_pixbuf_animation_get_height()).
//
// This pixbuf should be displayed for
// gdk_pixbuf_animation_iter_get_delay_time() milliseconds.
//
// The caller of this function does not own a reference to the returned pixbuf;
// the returned pixbuf will become invalid when the iterator advances to the
// next frame, which may happen anytime you call
// gdk_pixbuf_animation_iter_advance().
//
// Copy the pixbuf to keep it (don't just add a reference), as it may get
// recycled as you advance the iterator.
func (iter *PixbufAnimationIter) Pixbuf() *Pixbuf {
	var _arg0 *C.GdkPixbufAnimationIter // out
	var _cret *C.GdkPixbuf              // in

	_arg0 = (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter.Native()))

	_cret = C.gdk_pixbuf_animation_iter_get_pixbuf(_arg0)

	var _pixbuf *Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Pixbuf)

	return _pixbuf
}

// OnCurrentlyLoadingFrame: used to determine how to respond to the area_updated
// signal on PixbufLoader when loading an animation.
//
// The `::area_updated` signal is emitted for an area of the frame currently
// streaming in to the loader. So if you're on the currently loading frame, you
// will need to redraw the screen for the updated area.
func (iter *PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {
	var _arg0 *C.GdkPixbufAnimationIter // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter.Native()))

	_cret = C.gdk_pixbuf_animation_iter_on_currently_loading_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
