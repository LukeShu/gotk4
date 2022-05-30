// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gdk4_Drag_ConnectDNDFinished(gpointer, guintptr);
// extern void _gotk4_gdk4_Drag_ConnectDropPerformed(gpointer, guintptr);
import "C"

// glib.Type values for gdkdrag.go.
var (
	GTypeDragCancelReason = coreglib.Type(C.gdk_drag_cancel_reason_get_type())
	GTypeDrag             = coreglib.Type(C.gdk_drag_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDragCancelReason, F: marshalDragCancelReason},
		{T: GTypeDrag, F: marshalDrag},
	})
}

// DragCancelReason: used in GdkDrag to the reason of a cancelled DND operation.
type DragCancelReason C.gint

const (
	// DragCancelNoTarget: there is no suitable drop target.
	DragCancelNoTarget DragCancelReason = iota
	// DragCancelUserCancelled: drag cancelled by the user.
	DragCancelUserCancelled
	// DragCancelError: unspecified error.
	DragCancelError
)

func marshalDragCancelReason(p uintptr) (interface{}, error) {
	return DragCancelReason(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DragCancelReason.
func (d DragCancelReason) String() string {
	switch d {
	case DragCancelNoTarget:
		return "NoTarget"
	case DragCancelUserCancelled:
		return "UserCancelled"
	case DragCancelError:
		return "Error"
	default:
		return fmt.Sprintf("DragCancelReason(%d)", d)
	}
}

// Drag: GdkDrag object represents the source of an ongoing DND operation.
//
// A GdkDrag is created when a drag is started, and stays alive for duration of
// the DND operation. After a drag has been started with gdk.Drag().Begin, the
// caller gets informed about the status of the ongoing drag operation with
// signals on the GdkDrag object.
//
// GTK provides a higher level abstraction based on top of these functions, and
// so they are not normally needed in GTK applications. See the "Drag and Drop"
// section of the GTK documentation for more information.
type Drag struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Drag)(nil)
)

// Dragger describes types inherited from class Drag.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Dragger interface {
	coreglib.Objector
	baseDrag() *Drag
}

var _ Dragger = (*Drag)(nil)

func wrapDrag(obj *coreglib.Object) *Drag {
	return &Drag{
		Object: obj,
	}
}

func marshalDrag(p uintptr) (interface{}, error) {
	return wrapDrag(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (drag *Drag) baseDrag() *Drag {
	return drag
}

// BaseDrag returns the underlying base object.
func BaseDrag(obj Dragger) *Drag {
	return obj.baseDrag()
}

//export _gotk4_gdk4_Drag_ConnectDNDFinished
func _gotk4_gdk4_Drag_ConnectDNDFinished(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectDNDFinished is emitted when the destination side has finished reading
// all data.
//
// The drag object can now free all miscellaneous data.
func (drag *Drag) ConnectDNDFinished(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drag, "dnd-finished", false, unsafe.Pointer(C._gotk4_gdk4_Drag_ConnectDNDFinished), f)
}

//export _gotk4_gdk4_Drag_ConnectDropPerformed
func _gotk4_gdk4_Drag_ConnectDropPerformed(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectDropPerformed is emitted when the drop operation is performed on an
// accepting client.
func (drag *Drag) ConnectDropPerformed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drag, "drop-performed", false, unsafe.Pointer(C._gotk4_gdk4_Drag_ConnectDropPerformed), f)
}

// DropDone informs GDK that the drop ended.
//
// Passing FALSE for success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should be the last call
// before dropping the reference to the drag.
//
// The GdkDrag will only take the first gdk.Drag.DropDone() call as effective,
// if this function is called multiple times, all subsequent calls will be
// ignored.
//
// The function takes the following parameters:
//
//    - success: whether the drag was ultimatively successful.
//
func (drag *Drag) DropDone(success bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	if success {
		_arg1 = C.TRUE
	}
	*(**Drag)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gdk", "Drag").InvokeMethod("drop_done", args[:], nil)

	runtime.KeepAlive(drag)
	runtime.KeepAlive(success)
}

// Content returns the GdkContentProvider associated to the GdkDrag object.
//
// The function returns the following values:
//
//    - contentProvider: GdkContentProvider associated to drag.
//
func (drag *Drag) Content() *ContentProvider {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	*(**Drag)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Drag").InvokeMethod("get_content", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(drag)

	var _contentProvider *ContentProvider // out

	_contentProvider = wrapContentProvider(coreglib.Take(unsafe.Pointer(_cret)))

	return _contentProvider
}

// Device returns the GdkDevice associated to the GdkDrag object.
//
// The function returns the following values:
//
//    - device: GdkDevice associated to drag.
//
func (drag *Drag) Device() Devicer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	*(**Drag)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Drag").InvokeMethod("get_device", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(drag)

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	return _device
}

// Display gets the GdkDisplay that the drag object was created for.
//
// The function returns the following values:
//
//    - display: GdkDisplay.
//
func (drag *Drag) Display() *Display {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	*(**Drag)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Drag").InvokeMethod("get_display", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(drag)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// DragSurface returns the surface on which the drag icon should be rendered
// during the drag operation.
//
// Note that the surface may not be available until the drag operation has
// begun. GDK will move the surface in accordance with the ongoing drag
// operation. The surface is owned by drag and will be destroyed when the drag
// operation is over.
//
// The function returns the following values:
//
//    - surface (optional): drag surface, or NULL.
//
func (drag *Drag) DragSurface() Surfacer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	*(**Drag)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Drag").InvokeMethod("get_drag_surface", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(drag)

	var _surface Surfacer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Surfacer)
				return ok
			})
			rv, ok := casted.(Surfacer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
			}
			_surface = rv
		}
	}

	return _surface
}

// Formats retrieves the formats supported by this GdkDrag object.
//
// The function returns the following values:
//
//    - contentFormats: GdkContentFormats.
//
func (drag *Drag) Formats() *ContentFormats {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	*(**Drag)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Drag").InvokeMethod("get_formats", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(drag)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// Surface returns the GdkSurface where the drag originates.
//
// The function returns the following values:
//
//    - surface: GdkSurface where the drag originates.
//
func (drag *Drag) Surface() Surfacer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	*(**Drag)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gdk", "Drag").InvokeMethod("get_surface", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(drag)

	var _surface Surfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Surfacer)
			return ok
		})
		rv, ok := casted.(Surfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
		}
		_surface = rv
	}

	return _surface
}
