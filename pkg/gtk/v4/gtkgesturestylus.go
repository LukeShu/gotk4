// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_GestureStylus_ConnectDown(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureStylus_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureStylus_ConnectProximity(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk4_GestureStylus_ConnectUp(gpointer, gdouble, gdouble, guintptr);
import "C"

// GTypeGestureStylus returns the GType for the type GestureStylus.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeGestureStylus() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "GestureStylus").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalGestureStylus)
	return gtype
}

// GestureStylus: GtkGestureStylus is a GtkGesture specific to stylus input.
//
// The provided signals just relay the basic information of the stylus events.
type GestureStylus struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureStylus)(nil)
)

func wrapGestureStylus(obj *coreglib.Object) *GestureStylus {
	return &GestureStylus{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureStylus(p uintptr) (interface{}, error) {
	return wrapGestureStylus(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_GestureStylus_ConnectDown
func _gotk4_gtk4_GestureStylus_ConnectDown(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectDown is emitted when the stylus touches the device.
func (gesture *GestureStylus) ConnectDown(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "down", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectDown), f)
}

//export _gotk4_gtk4_GestureStylus_ConnectMotion
func _gotk4_gtk4_GestureStylus_ConnectMotion(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectMotion is emitted when the stylus moves while touching the device.
func (gesture *GestureStylus) ConnectMotion(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "motion", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectMotion), f)
}

//export _gotk4_gtk4_GestureStylus_ConnectProximity
func _gotk4_gtk4_GestureStylus_ConnectProximity(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectProximity is emitted when the stylus is in proximity of the device.
func (gesture *GestureStylus) ConnectProximity(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "proximity", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectProximity), f)
}

//export _gotk4_gtk4_GestureStylus_ConnectUp
func _gotk4_gtk4_GestureStylus_ConnectUp(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(x, y float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y float64))
	}

	var _x float64 // out
	var _y float64 // out

	_x = float64(arg1)
	_y = float64(arg2)

	f(_x, _y)
}

// ConnectUp is emitted when the stylus no longer touches the device.
func (gesture *GestureStylus) ConnectUp(f func(x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(gesture, "up", false, unsafe.Pointer(C._gotk4_gtk4_GestureStylus_ConnectUp), f)
}

// NewGestureStylus creates a new GtkGestureStylus.
//
// The function returns the following values:
//
//    - gestureStylus: newly created stylus gesture.
//
func NewGestureStylus() *GestureStylus {
	_info := girepository.MustFind("Gtk", "GestureStylus")
	_gret := _info.InvokeClassMethod("new_GestureStylus", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _gestureStylus *GestureStylus // out

	_gestureStylus = wrapGestureStylus(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _gestureStylus
}

// Backlog returns the accumulated backlog of tracking information.
//
// By default, GTK will limit rate of input events. On stylus input where
// accuracy of strokes is paramount, this function returns the accumulated
// coordinate/timing state before the emission of the current
// [Gtk.GestureStylus::motion] signal.
//
// This function may only be called within a gtk.GestureStylus::motion signal
// handler, the state given in this signal and obtainable through
// gtk.GestureStylus.GetAxis() express the latest (most up-to-date) state in
// motion history.
//
// The backlog is provided in chronological order.
//
// The function returns the following values:
//
//    - backlog coordinates and times for the backlog events.
//    - ok: TRUE if there is a backlog to unfold in the current state.
//
func (gesture *GestureStylus) Backlog() ([]gdk.TimeCoord, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_info := girepository.MustFind("Gtk", "GestureStylus")
	_gret := _info.InvokeClassMethod("get_backlog", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _backlog []gdk.TimeCoord // out
	var _ok bool                 // out

	defer C.free(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0]))))
	{
		src := unsafe.Slice((**C.void)(*(**C.void)(unsafe.Pointer(&_outs[0]))), *(*C.guint)(unsafe.Pointer(&_outs[1])))
		_backlog = make([]gdk.TimeCoord, *(*C.guint)(unsafe.Pointer(&_outs[1])))
		for i := 0; i < int(*(*C.guint)(unsafe.Pointer(&_outs[1]))); i++ {
			_backlog[i] = *(*gdk.TimeCoord)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_backlog[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _backlog, _ok
}

// DeviceTool returns the GdkDeviceTool currently driving input through this
// gesture.
//
// This function must be called from the handler of one of the
// gtk.GestureStylus::down, gtk.GestureStylus::motion, gtk.GestureStylus::up or
// gtk.GestureStylus::proximity signals.
//
// The function returns the following values:
//
//    - deviceTool (optional): current stylus tool.
//
func (gesture *GestureStylus) DeviceTool() *gdk.DeviceTool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(gesture).Native()))

	_info := girepository.MustFind("Gtk", "GestureStylus")
	_gret := _info.InvokeClassMethod("get_device_tool", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gesture)

	var _deviceTool *gdk.DeviceTool // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
			_deviceTool = &gdk.DeviceTool{
				Object: obj,
			}
		}
	}

	return _deviceTool
}
