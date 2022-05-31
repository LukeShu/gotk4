// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gdk4_FrameClock_ConnectAfterPaint(gpointer, guintptr);
// extern void _gotk4_gdk4_FrameClock_ConnectBeforePaint(gpointer, guintptr);
// extern void _gotk4_gdk4_FrameClock_ConnectFlushEvents(gpointer, guintptr);
// extern void _gotk4_gdk4_FrameClock_ConnectLayout(gpointer, guintptr);
// extern void _gotk4_gdk4_FrameClock_ConnectPaint(gpointer, guintptr);
// extern void _gotk4_gdk4_FrameClock_ConnectResumeEvents(gpointer, guintptr);
// extern void _gotk4_gdk4_FrameClock_ConnectUpdate(gpointer, guintptr);
import "C"

// glib.Type values for gdkframeclock.go.
var (
	GTypeFrameClockPhase = coreglib.Type(C.gdk_frame_clock_phase_get_type())
	GTypeFrameClock      = coreglib.Type(C.gdk_frame_clock_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFrameClockPhase, F: marshalFrameClockPhase},
		{T: GTypeFrameClock, F: marshalFrameClock},
	})
}

// FrameClockPhase: used to represent the different paint clock phases that can
// be requested.
//
// The elements of the enumeration correspond to the signals of GdkFrameClock.
type FrameClockPhase C.guint

const (
	// FrameClockPhaseNone: no phase.
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
	return FrameClockPhase(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for FrameClockPhase.
func (f FrameClockPhase) String() string {
	if f == 0 {
		return "FrameClockPhase(0)"
	}

	var builder strings.Builder
	builder.Grow(192)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FrameClockPhaseNone:
			builder.WriteString("None|")
		case FrameClockPhaseFlushEvents:
			builder.WriteString("FlushEvents|")
		case FrameClockPhaseBeforePaint:
			builder.WriteString("BeforePaint|")
		case FrameClockPhaseUpdate:
			builder.WriteString("Update|")
		case FrameClockPhaseLayout:
			builder.WriteString("Layout|")
		case FrameClockPhasePaint:
			builder.WriteString("Paint|")
		case FrameClockPhaseResumeEvents:
			builder.WriteString("ResumeEvents|")
		case FrameClockPhaseAfterPaint:
			builder.WriteString("AfterPaint|")
		default:
			builder.WriteString(fmt.Sprintf("FrameClockPhase(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FrameClockPhase) Has(other FrameClockPhase) bool {
	return (f & other) == other
}

// FrameClockOverrider contains methods that are overridable.
type FrameClockOverrider interface {
}

// FrameClock: GdkFrameClock tells the application when to update and repaint a
// surface.
//
// This may be synced to the vertical refresh rate of the monitor, for example.
// Even when the frame clock uses a simple timer rather than a hardware-based
// vertical sync, the frame clock helps because it ensures everything paints at
// the same time (reducing the total number of frames).
//
// The frame clock can also automatically stop painting when it knows the frames
// will not be visible, or scale back animation framerates.
//
// GdkFrameClock is designed to be compatible with an OpenGL-based
// implementation or with mozRequestAnimationFrame in Firefox, for example.
//
// A frame clock is idle until someone requests a frame with
// gdk.FrameClock.RequestPhase(). At some later point that makes sense for the
// synchronization being implemented, the clock will process a frame and emit
// signals for each phase that has been requested. (See the signals of the
// GdkFrameClock class for documentation of the phases.
// GDK_FRAME_CLOCK_PHASE_UPDATE and the gdkframeclock::update signal are most
// interesting for application writers, and are used to update the animations,
// using the frame time given by gdk.FrameClock.GetFrameTime.
//
// The frame time is reported in microseconds and generally in the same
// timescale as g_get_monotonic_time(), however, it is not the same as
// g_get_monotonic_time(). The frame time does not advance during the time a
// frame is being painted, and outside of a frame, an attempt is made so that
// all calls to gdk.FrameClock.GetFrameTime() that are called at a “similar”
// time get the same value. This means that if different animations are timed by
// looking at the difference in time between an initial value from
// gdk.FrameClock.GetFrameTime() and the value inside the gdkframeclock::update
// signal of the clock, they will stay exactly synchronized.
type FrameClock struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FrameClock)(nil)
)

// FrameClocker describes types inherited from class FrameClock.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FrameClocker interface {
	coreglib.Objector
	baseFrameClock() *FrameClock
}

var _ FrameClocker = (*FrameClock)(nil)

func classInitFrameClocker(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFrameClock(obj *coreglib.Object) *FrameClock {
	return &FrameClock{
		Object: obj,
	}
}

func marshalFrameClock(p uintptr) (interface{}, error) {
	return wrapFrameClock(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (frameClock *FrameClock) baseFrameClock() *FrameClock {
	return frameClock
}

// BaseFrameClock returns the underlying base object.
func BaseFrameClock(obj FrameClocker) *FrameClock {
	return obj.baseFrameClock()
}

//export _gotk4_gdk4_FrameClock_ConnectAfterPaint
func _gotk4_gdk4_FrameClock_ConnectAfterPaint(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectAfterPaint: this signal ends processing of the frame.
//
// Applications should generally not handle this signal.
func (frameClock *FrameClock) ConnectAfterPaint(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(frameClock, "after-paint", false, unsafe.Pointer(C._gotk4_gdk4_FrameClock_ConnectAfterPaint), f)
}

//export _gotk4_gdk4_FrameClock_ConnectBeforePaint
func _gotk4_gdk4_FrameClock_ConnectBeforePaint(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectBeforePaint begins processing of the frame.
//
// Applications should generally not handle this signal.
func (frameClock *FrameClock) ConnectBeforePaint(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(frameClock, "before-paint", false, unsafe.Pointer(C._gotk4_gdk4_FrameClock_ConnectBeforePaint), f)
}

//export _gotk4_gdk4_FrameClock_ConnectFlushEvents
func _gotk4_gdk4_FrameClock_ConnectFlushEvents(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectFlushEvents: used to flush pending motion events that are being
// batched up and compressed together.
//
// Applications should not handle this signal.
func (frameClock *FrameClock) ConnectFlushEvents(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(frameClock, "flush-events", false, unsafe.Pointer(C._gotk4_gdk4_FrameClock_ConnectFlushEvents), f)
}

//export _gotk4_gdk4_FrameClock_ConnectLayout
func _gotk4_gdk4_FrameClock_ConnectLayout(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLayout is emitted as the second step of toolkit and application
// processing of the frame.
//
// Any work to update sizes and positions of application elements should be
// performed. GTK normally handles this internally.
func (frameClock *FrameClock) ConnectLayout(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(frameClock, "layout", false, unsafe.Pointer(C._gotk4_gdk4_FrameClock_ConnectLayout), f)
}

//export _gotk4_gdk4_FrameClock_ConnectPaint
func _gotk4_gdk4_FrameClock_ConnectPaint(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPaint is emitted as the third step of toolkit and application
// processing of the frame.
//
// The frame is repainted. GDK normally handles this internally and emits
// gdk.Surface::render signals which are turned into gtk.Widget::snapshot
// signals by GTK.
func (frameClock *FrameClock) ConnectPaint(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(frameClock, "paint", false, unsafe.Pointer(C._gotk4_gdk4_FrameClock_ConnectPaint), f)
}

//export _gotk4_gdk4_FrameClock_ConnectResumeEvents
func _gotk4_gdk4_FrameClock_ConnectResumeEvents(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectResumeEvents is emitted after processing of the frame is finished.
//
// This signal is handled internally by GTK to resume normal event processing.
// Applications should not handle this signal.
func (frameClock *FrameClock) ConnectResumeEvents(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(frameClock, "resume-events", false, unsafe.Pointer(C._gotk4_gdk4_FrameClock_ConnectResumeEvents), f)
}

//export _gotk4_gdk4_FrameClock_ConnectUpdate
func _gotk4_gdk4_FrameClock_ConnectUpdate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectUpdate is emitted as the first step of toolkit and application
// processing of the frame.
//
// Animations should be updated using gdk.FrameClock.GetFrameTime().
// Applications can connect directly to this signal, or use
// gtk.Widget.AddTickCallback() as a more convenient interface.
func (frameClock *FrameClock) ConnectUpdate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(frameClock, "update", false, unsafe.Pointer(C._gotk4_gdk4_FrameClock_ConnectUpdate), f)
}

// BeginUpdating starts updates for an animation.
//
// Until a matching call to gdk.FrameClock.EndUpdating() is made, the frame
// clock will continually request a new frame with the
// GDK_FRAME_CLOCK_PHASE_UPDATE phase. This function may be called multiple
// times and frames will be requested until gdk_frame_clock_end_updating() is
// called the same number of times.
func (frameClock *FrameClock) BeginUpdating() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	*(**FrameClock)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gdk", "FrameClock").InvokeMethod("begin_updating", args[:], nil)

	runtime.KeepAlive(frameClock)
}

// EndUpdating stops updates for an animation.
//
// See the documentation for gdk.FrameClock.BeginUpdating().
func (frameClock *FrameClock) EndUpdating() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	*(**FrameClock)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gdk", "FrameClock").InvokeMethod("end_updating", args[:], nil)

	runtime.KeepAlive(frameClock)
}

// CurrentTimings gets the frame timings for the current frame.
//
// The function returns the following values:
//
//    - frameTimings (optional): GdkFrameTimings for the frame currently being
//      processed, or even no frame is being processed, for the previous frame.
//      Before any frames have been processed, returns NULL.
//
func (frameClock *FrameClock) CurrentTimings() *FrameTimings {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	*(**FrameClock)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "FrameClock").InvokeMethod("get_current_timings", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frameClock)

	var _frameTimings *FrameTimings // out

	if _cret != nil {
		_frameTimings = (*FrameTimings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gdk_frame_timings_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_frameTimings)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gdk_frame_timings_unref((*C.GdkFrameTimings)(intern.C))
			},
		)
	}

	return _frameTimings
}

// FPS calculates the current frames-per-second, based on the frame timings of
// frame_clock.
//
// The function returns the following values:
//
//    - gdouble: current fps, as a double.
//
func (frameClock *FrameClock) FPS() float64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.double // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	*(**FrameClock)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "FrameClock").InvokeMethod("get_fps", args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frameClock)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// FrameCounter: GdkFrameClock maintains a 64-bit counter that increments for
// each frame drawn.
//
// The function returns the following values:
//
//    - gint64: inside frame processing, the value of the frame counter for the
//      current frame. Outside of frame processing, the frame counter for the
//      last frame.
//
func (frameClock *FrameClock) FrameCounter() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	*(**FrameClock)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "FrameClock").InvokeMethod("get_frame_counter", args[:], nil)
	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frameClock)

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
//
// The function returns the following values:
//
//    - gint64: timestamp in microseconds, in the timescale of of
//      g_get_monotonic_time().
//
func (frameClock *FrameClock) FrameTime() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	*(**FrameClock)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "FrameClock").InvokeMethod("get_frame_time", args[:], nil)
	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frameClock)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// HistoryStart returns the frame counter for the oldest frame available in
// history.
//
// GdkFrameClock internally keeps a history of GdkFrameTimings objects for
// recent frames that can be retrieved with gdk.FrameClock.GetTimings(). The set
// of stored frames is the set from the counter values given by
// gdk.FrameClock.GetHistoryStart() and gdk.FrameClock.GetFrameCounter(),
// inclusive.
//
// The function returns the following values:
//
//    - gint64: frame counter value for the oldest frame that is available in the
//      internal frame history of the GdkFrameClock.
//
func (frameClock *FrameClock) HistoryStart() int64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.gint64 // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	*(**FrameClock)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "FrameClock").InvokeMethod("get_history_start", args[:], nil)
	_cret = *(*C.gint64)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frameClock)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// Timings retrieves a GdkFrameTimings object holding timing information for the
// current frame or a recent frame.
//
// The GdkFrameTimings object may not yet be complete: see
// gdk.FrameTimings.GetComplete().
//
// The function takes the following parameters:
//
//    - frameCounter: frame counter value identifying the frame to be received.
//
// The function returns the following values:
//
//    - frameTimings (optional): GdkFrameTimings object for the specified frame,
//      or NULL if it is not available. See gdk.FrameClock.GetHistoryStart().
//
func (frameClock *FrameClock) Timings(frameCounter int64) *FrameTimings {
	var args [2]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 C.gint64 // out
	var _cret *C.void  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frameClock).Native()))
	_arg1 = C.gint64(frameCounter)
	*(**FrameClock)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gdk", "FrameClock").InvokeMethod("get_timings", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frameClock)
	runtime.KeepAlive(frameCounter)

	var _frameTimings *FrameTimings // out

	if _cret != nil {
		_frameTimings = (*FrameTimings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gdk_frame_timings_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_frameTimings)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gdk_frame_timings_unref((*C.GdkFrameTimings)(intern.C))
			},
		)
	}

	return _frameTimings
}
