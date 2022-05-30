// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdkframetimings.go.
var GTypeFrameTimings = coreglib.Type(C.gdk_frame_timings_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFrameTimings, F: marshalFrameTimings},
	})
}

// FrameTimings object holds timing information for a single frame of the
// application’s displays. To retrieve FrameTimings objects, use
// gdk_frame_clock_get_timings() or gdk_frame_clock_get_current_timings(). The
// information in FrameTimings is useful for precise synchronization of video
// with the event or audio streams, and for measuring quality metrics for the
// application’s display, such as latency and jitter.
//
// An instance of this type is always passed by reference.
type FrameTimings struct {
	*frameTimings
}

// frameTimings is the struct that's finalized.
type frameTimings struct {
	native *C.GdkFrameTimings
}

func marshalFrameTimings(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &FrameTimings{&frameTimings{(*C.GdkFrameTimings)(b)}}, nil
}

// Complete: timing information in a FrameTimings is filled in incrementally as
// the frame as drawn and passed off to the window system for processing and
// display to the user. The accessor functions for FrameTimings can return 0 to
// indicate an unavailable value for two reasons: either because the information
// is not yet available, or because it isn't available at all. Once
// gdk_frame_timings_get_complete() returns TRUE for a frame, you can be certain
// that no further values will become available and be stored in the
// FrameTimings.
//
// The function returns the following values:
//
//    - ok: TRUE if all information that will be available for the frame has been
//      filled in.
//
func (timings *FrameTimings) Complete() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(timings)))
	*(**FrameTimings)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(timings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FrameCounter gets the frame counter value of the FrameClock when this this
// frame was drawn.
//
// The function returns the following values:
//
//    - gint64: frame counter value for this frame.
//
func (timings *FrameTimings) FrameCounter() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(timings)))
	*(**FrameTimings)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// FrameTime returns the frame time for the frame. This is the time value that
// is typically used to time animations for the frame. See
// gdk_frame_clock_get_frame_time().
//
// The function returns the following values:
//
//    - gint64: frame time for the frame, in the timescale of
//      g_get_monotonic_time().
//
func (timings *FrameTimings) FrameTime() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(timings)))
	*(**FrameTimings)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// PredictedPresentationTime gets the predicted time at which this frame will be
// displayed. Although no predicted time may be available, if one is available,
// it will be available while the frame is being generated, in contrast to
// gdk_frame_timings_get_presentation_time(), which is only available after the
// frame has been presented. In general, if you are simply animating, you should
// use gdk_frame_clock_get_frame_time() rather than this function, but this
// function is useful for applications that want exact control over latency. For
// example, a movie player may want this information for Audio/Video
// synchronization.
//
// The function returns the following values:
//
//    - gint64: predicted time at which the frame will be presented, in the
//      timescale of g_get_monotonic_time(), or 0 if no predicted presentation
//      time is available.
//
func (timings *FrameTimings) PredictedPresentationTime() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(timings)))
	*(**FrameTimings)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// PresentationTime reurns the presentation time. This is the time at which the
// frame became visible to the user.
//
// The function returns the following values:
//
//    - gint64: time the frame was displayed to the user, in the timescale of
//      g_get_monotonic_time(), or 0 if no presentation time is available. See
//      gdk_frame_timings_get_complete().
//
func (timings *FrameTimings) PresentationTime() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(timings)))
	*(**FrameTimings)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// RefreshInterval gets the natural interval between presentation times for the
// display that this frame was displayed on. Frame presentation usually happens
// during the “vertical blanking interval”.
//
// The function returns the following values:
//
//    - gint64: refresh interval of the display, in microseconds, or 0 if the
//      refresh interval is not available. See gdk_frame_timings_get_complete().
//
func (timings *FrameTimings) RefreshInterval() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(timings)))
	*(**FrameTimings)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}
