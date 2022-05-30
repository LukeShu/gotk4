// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_GestureDrag_ConnectDragBegin(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureDrag_ConnectDragEnd(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_gtk3_GestureDrag_ConnectDragUpdate(gpointer, gdouble, gdouble, guintptr);
import "C"

// glib.Type values for gtkgesturedrag.go.
var GTypeGestureDrag = coreglib.Type(C.gtk_gesture_drag_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeGestureDrag, F: marshalGestureDrag},
	})
}

// GestureDragOverrider contains methods that are overridable.
type GestureDragOverrider interface {
}

// GestureDrag is a Gesture implementation that recognizes drag operations. The
// drag operation itself can be tracked throught the GestureDrag::drag-begin,
// GestureDrag::drag-update and GestureDrag::drag-end signals, or the relevant
// coordinates be extracted through gtk_gesture_drag_get_offset() and
// gtk_gesture_drag_get_start_point().
type GestureDrag struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureDrag)(nil)
)

func classInitGestureDragger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGestureDrag(obj *coreglib.Object) *GestureDrag {
	return &GestureDrag{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureDrag(p uintptr) (interface{}, error) {
	return wrapGestureDrag(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_GestureDrag_ConnectDragBegin
func _gotk4_gtk3_GestureDrag_ConnectDragBegin(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(startX, startY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(startX, startY float64))
	}

	var _startX float64 // out
	var _startY float64 // out

	_startX = float64(arg1)
	_startY = float64(arg2)

	f(_startX, _startY)
}

// ConnectDragBegin: this signal is emitted whenever dragging starts.
func (v *GestureDrag) ConnectDragBegin(f func(startX, startY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "drag-begin", false, unsafe.Pointer(C._gotk4_gtk3_GestureDrag_ConnectDragBegin), f)
}

//export _gotk4_gtk3_GestureDrag_ConnectDragEnd
func _gotk4_gtk3_GestureDrag_ConnectDragEnd(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(offsetX, offsetY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offsetX, offsetY float64))
	}

	var _offsetX float64 // out
	var _offsetY float64 // out

	_offsetX = float64(arg1)
	_offsetY = float64(arg2)

	f(_offsetX, _offsetY)
}

// ConnectDragEnd: this signal is emitted whenever the dragging is finished.
func (v *GestureDrag) ConnectDragEnd(f func(offsetX, offsetY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "drag-end", false, unsafe.Pointer(C._gotk4_gtk3_GestureDrag_ConnectDragEnd), f)
}

//export _gotk4_gtk3_GestureDrag_ConnectDragUpdate
func _gotk4_gtk3_GestureDrag_ConnectDragUpdate(arg0 C.gpointer, arg1 C.gdouble, arg2 C.gdouble, arg3 C.guintptr) {
	var f func(offsetX, offsetY float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(offsetX, offsetY float64))
	}

	var _offsetX float64 // out
	var _offsetY float64 // out

	_offsetX = float64(arg1)
	_offsetY = float64(arg2)

	f(_offsetX, _offsetY)
}

// ConnectDragUpdate: this signal is emitted whenever the dragging point moves.
func (v *GestureDrag) ConnectDragUpdate(f func(offsetX, offsetY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "drag-update", false, unsafe.Pointer(C._gotk4_gtk3_GestureDrag_ConnectDragUpdate), f)
}

// NewGestureDrag returns a newly created Gesture that recognizes drags.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureDrag: newly created GestureDrag.
//
func NewGestureDrag(widget Widgetter) *GestureDrag {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	*(*Widgetter)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "GestureDrag").InvokeMethod("new_GestureDrag", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _gestureDrag *GestureDrag // out

	_gestureDrag = wrapGestureDrag(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureDrag
}
