// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_frame_clock_phase_get_type()), F: marshalFrameClockPhase},
		{T: externglib.Type(C.gdk_frame_clock_get_type()), F: marshalFrameClock},
	})
}

// FrameClockPhase: used to represent the different paint clock phases that can
// be requested.
//
// The elements of the enumeration correspond to the signals of `GdkFrameClock`.
type FrameClockPhase int

const (
	// FrameClockPhaseNone: no phase
	FrameClockPhaseNone FrameClockPhase = 0b0
	// FrameClockPhaseFlushEvents corresponds to GdkFrameClock::flush-events.
	// Should not be handled by applications.
	FrameClockPhaseFlushEvents FrameClockPhase = 0b1
	// FrameClockPhaseBeforePaint corresponds to GdkFrameClock::before-paint.
	// Should not be handled by applications.
	FrameClockPhaseBeforePaint FrameClockPhase = 0b10
	// FrameClockPhaseUpdate corresponds to GdkFrameClock::update.
	FrameClockPhaseUpdate FrameClockPhase = 0b100
	// FrameClockPhaseLayout corresponds to GdkFrameClock::layout. Should not be
	// handled by applicatiosn.
	FrameClockPhaseLayout FrameClockPhase = 0b1000
	// FrameClockPhasePaint corresponds to GdkFrameClock::paint.
	FrameClockPhasePaint FrameClockPhase = 0b10000
	// FrameClockPhaseResumeEvents corresponds to GdkFrameClock::resume-events.
	// Should not be handled by applications.
	FrameClockPhaseResumeEvents FrameClockPhase = 0b100000
	// FrameClockPhaseAfterPaint corresponds to GdkFrameClock::after-paint.
	// Should not be handled by applications.
	FrameClockPhaseAfterPaint FrameClockPhase = 0b1000000
)

func marshalFrameClockPhase(p uintptr) (interface{}, error) {
	return FrameClockPhase(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FrameClock: `GdkFrameClock` tells the application when to update and repaint
// a surface.
//
// This may be synced to the vertical refresh rate of the monitor, for example.
// Even when the frame clock uses a simple timer rather than a hardware-based
// vertical sync, the frame clock helps because it ensures everything paints at
// the same time (reducing the total number of frames).
//
// The frame clock can also automatically stop painting when it knows the frames
// will not be visible, or scale back animation framerates.
//
// `GdkFrameClock` is designed to be compatible with an OpenGL-based
// implementation or with mozRequestAnimationFrame in Firefox, for example.
//
// A frame clock is idle until someone requests a frame with
// [method@Gdk.FrameClock.request_phase]. At some later point that makes sense
// for the synchronization being implemented, the clock will process a frame and
// emit signals for each phase that has been requested. (See the signals of the
// `GdkFrameClock` class for documentation of the phases.
// GDK_FRAME_CLOCK_PHASE_UPDATE and the [signal@GdkFrameClock::update] signal
// are most interesting for application writers, and are used to update the
// animations, using the frame time given by
// [metohd@Gdk.FrameClock.get_frame_time].
//
// The frame time is reported in microseconds and generally in the same
// timescale as g_get_monotonic_time(), however, it is not the same as
// g_get_monotonic_time(). The frame time does not advance during the time a
// frame is being painted, and outside of a frame, an attempt is made so that
// all calls to [method@Gdk.FrameClock.get_frame_time] that are called at a
// “similar” time get the same value. This means that if different animations
// are timed by looking at the difference in time between an initial value from
// [method@Gdk.FrameClock.get_frame_time] and the value inside the
// [signal@GdkFrameClock::update] signal of the clock, they will stay exactly
// synchronized.
type FrameClock interface {
	gextras.Objector

	// BeginUpdating starts updates for an animation.
	//
	// Until a matching call to [method@Gdk.FrameClock.end_updating] is made,
	// the frame clock will continually request a new frame with the
	// GDK_FRAME_CLOCK_PHASE_UPDATE phase. This function may be called multiple
	// times and frames will be requested until gdk_frame_clock_end_updating()
	// is called the same number of times.
	BeginUpdating()
	// EndUpdating stops updates for an animation.
	//
	// See the documentation for [method@Gdk.FrameClock.begin_updating].
	EndUpdating()
	// CurrentTimings gets the frame timings for the current frame.
	CurrentTimings() *FrameTimings
	// Fps calculates the current frames-per-second, based on the frame timings
	// of @frame_clock.
	Fps() float64
	// FrameCounter: `GdkFrameClock` maintains a 64-bit counter that increments
	// for each frame drawn.
	FrameCounter() int64
	// FrameTime gets the time that should currently be used for animations.
	//
	// Inside the processing of a frame, it’s the time used to compute the
	// animation position of everything in a frame. Outside of a frame, it's the
	// time of the conceptual “previous frame,” which may be either the actual
	// previous frame time, or if that’s too old, an updated time.
	FrameTime() int64
	// HistoryStart returns the frame counter for the oldest frame available in
	// history.
	//
	// `GdkFrameClock` internally keeps a history of `GdkFrameTimings` objects
	// for recent frames that can be retrieved with
	// [method@Gdk.FrameClock.get_timings]. The set of stored frames is the set
	// from the counter values given by
	// [method@Gdk.FrameClock.get_history_start] and
	// [method@Gdk.FrameClock.get_frame_counter], inclusive.
	HistoryStart() int64
	// RefreshInfo predicts a presentation time, based on history.
	//
	// Using the frame history stored in the frame clock, finds the last known
	// presentation time and refresh interval, and assuming that presentation
	// times are separated by the refresh interval, predicts a presentation time
	// that is a multiple of the refresh interval after the last presentation
	// time, and later than @base_time.
	RefreshInfo(baseTime int64) (refreshIntervalReturn int64, presentationTimeReturn int64)
	// Timings retrieves a `GdkFrameTimings` object holding timing information
	// for the current frame or a recent frame.
	//
	// The `GdkFrameTimings` object may not yet be complete: see
	// [method@Gdk.FrameTimings.get_complete].
	Timings(frameCounter int64) *FrameTimings
}

// FrameClockClass implements the FrameClock interface.
type FrameClockClass struct {
	*externglib.Object
}

var _ FrameClock = (*FrameClockClass)(nil)

func wrapFrameClock(obj *externglib.Object) FrameClock {
	return &FrameClockClass{
		Object: obj,
	}
}

func marshalFrameClock(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFrameClock(obj), nil
}

// BeginUpdating starts updates for an animation.
//
// Until a matching call to [method@Gdk.FrameClock.end_updating] is made, the
// frame clock will continually request a new frame with the
// GDK_FRAME_CLOCK_PHASE_UPDATE phase. This function may be called multiple
// times and frames will be requested until gdk_frame_clock_end_updating() is
// called the same number of times.
func (f *FrameClockClass) BeginUpdating() {
	var _arg0 *C.GdkFrameClock // out

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))

	C.gdk_frame_clock_begin_updating(_arg0)
}

// EndUpdating stops updates for an animation.
//
// See the documentation for [method@Gdk.FrameClock.begin_updating].
func (f *FrameClockClass) EndUpdating() {
	var _arg0 *C.GdkFrameClock // out

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))

	C.gdk_frame_clock_end_updating(_arg0)
}

// CurrentTimings gets the frame timings for the current frame.
func (f *FrameClockClass) CurrentTimings() *FrameTimings {
	var _arg0 *C.GdkFrameClock   // out
	var _cret *C.GdkFrameTimings // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))

	_cret = C.gdk_frame_clock_get_current_timings(_arg0)

	var _frameTimings *FrameTimings // out

	_frameTimings = (*FrameTimings)(unsafe.Pointer(_cret))
	C.gdk_frame_timings_ref(_cret)
	runtime.SetFinalizer(_frameTimings, func(v *FrameTimings) {
		C.gdk_frame_timings_unref((*C.GdkFrameTimings)(unsafe.Pointer(v)))
	})

	return _frameTimings
}

// Fps calculates the current frames-per-second, based on the frame timings of
// @frame_clock.
func (f *FrameClockClass) Fps() float64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.double         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))

	_cret = C.gdk_frame_clock_get_fps(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// FrameCounter: `GdkFrameClock` maintains a 64-bit counter that increments for
// each frame drawn.
func (f *FrameClockClass) FrameCounter() int64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))

	_cret = C.gdk_frame_clock_get_frame_counter(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// FrameTime gets the time that should currently be used for animations.
//
// Inside the processing of a frame, it’s the time used to compute the animation
// position of everything in a frame. Outside of a frame, it's the time of the
// conceptual “previous frame,” which may be either the actual previous frame
// time, or if that’s too old, an updated time.
func (f *FrameClockClass) FrameTime() int64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))

	_cret = C.gdk_frame_clock_get_frame_time(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// HistoryStart returns the frame counter for the oldest frame available in
// history.
//
// `GdkFrameClock` internally keeps a history of `GdkFrameTimings` objects for
// recent frames that can be retrieved with [method@Gdk.FrameClock.get_timings].
// The set of stored frames is the set from the counter values given by
// [method@Gdk.FrameClock.get_history_start] and
// [method@Gdk.FrameClock.get_frame_counter], inclusive.
func (f *FrameClockClass) HistoryStart() int64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))

	_cret = C.gdk_frame_clock_get_history_start(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// RefreshInfo predicts a presentation time, based on history.
//
// Using the frame history stored in the frame clock, finds the last known
// presentation time and refresh interval, and assuming that presentation times
// are separated by the refresh interval, predicts a presentation time that is a
// multiple of the refresh interval after the last presentation time, and later
// than @base_time.
func (f *FrameClockClass) RefreshInfo(baseTime int64) (refreshIntervalReturn int64, presentationTimeReturn int64) {
	var _arg0 *C.GdkFrameClock // out
	var _arg1 C.gint64         // out
	var _arg2 C.gint64         // in
	var _arg3 C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))
	_arg1 = C.gint64(baseTime)

	C.gdk_frame_clock_get_refresh_info(_arg0, _arg1, &_arg2, &_arg3)

	var _refreshIntervalReturn int64  // out
	var _presentationTimeReturn int64 // out

	_refreshIntervalReturn = int64(_arg2)
	_presentationTimeReturn = int64(_arg3)

	return _refreshIntervalReturn, _presentationTimeReturn
}

// Timings retrieves a `GdkFrameTimings` object holding timing information for
// the current frame or a recent frame.
//
// The `GdkFrameTimings` object may not yet be complete: see
// [method@Gdk.FrameTimings.get_complete].
func (f *FrameClockClass) Timings(frameCounter int64) *FrameTimings {
	var _arg0 *C.GdkFrameClock   // out
	var _arg1 C.gint64           // out
	var _cret *C.GdkFrameTimings // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer((&f).Native()))
	_arg1 = C.gint64(frameCounter)

	_cret = C.gdk_frame_clock_get_timings(_arg0, _arg1)

	var _frameTimings *FrameTimings // out

	_frameTimings = (*FrameTimings)(unsafe.Pointer(_cret))
	C.gdk_frame_timings_ref(_cret)
	runtime.SetFinalizer(_frameTimings, func(v *FrameTimings) {
		C.gdk_frame_timings_unref((*C.GdkFrameTimings)(unsafe.Pointer(v)))
	})

	return _frameTimings
}
