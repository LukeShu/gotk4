// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drag_cancel_reason_get_type()), F: marshalDragCancelReason},
		{T: externglib.Type(C.gdk_drag_get_type()), F: marshalDragger},
	})
}

// DragCancelReason: used in GdkDrag to the reason of a cancelled DND operation.
type DragCancelReason int

const (
	// DragCancelNoTarget: there is no suitable drop target.
	DragCancelNoTarget DragCancelReason = iota
	// DragCancelUserCancelled: drag cancelled by the user.
	DragCancelUserCancelled
	// DragCancelError: unspecified error.
	DragCancelError
)

func marshalDragCancelReason(p uintptr) (interface{}, error) {
	return DragCancelReason(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
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

// DragActionIsUnique checks if action represents a single action or includes
// multiple actions.
//
// When action is 0 - ie no action was given, TRUE is returned.
func DragActionIsUnique(action DragAction) bool {
	var _arg1 C.GdkDragAction // out
	var _cret C.gboolean      // in

	_arg1 = C.GdkDragAction(action)

	_cret = C.gdk_drag_action_is_unique(_arg1)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
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
	*externglib.Object
}

// Dragger describes Drag's abstract methods.
type Dragger interface {
	externglib.Objector

	// DropDone informs GDK that the drop ended.
	DropDone(success bool)
	// Actions determines the bitmask of possible actions proposed by the
	// source.
	Actions() DragAction
	// Content returns the GdkContentProvider associated to the GdkDrag object.
	Content() *ContentProvider
	// Device returns the GdkDevice associated to the GdkDrag object.
	Device() Devicer
	// Display gets the GdkDisplay that the drag object was created for.
	Display() *Display
	// DragSurface returns the surface on which the drag icon should be rendered
	// during the drag operation.
	DragSurface() Surfacer
	// Formats retrieves the formats supported by this GdkDrag object.
	Formats() *ContentFormats
	// SelectedAction determines the action chosen by the drag destination.
	SelectedAction() DragAction
	// Surface returns the GdkSurface where the drag originates.
	Surface() Surfacer
	// SetHotspot sets the position of the drag surface that will be kept under
	// the cursor hotspot.
	SetHotspot(hotX, hotY int)
}

var _ Dragger = (*Drag)(nil)

func wrapDrag(obj *externglib.Object) *Drag {
	return &Drag{
		Object: obj,
	}
}

func marshalDragger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrag(obj), nil
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
func (drag *Drag) DropDone(success bool) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))
	if success {
		_arg1 = C.TRUE
	}

	C.gdk_drag_drop_done(_arg0, _arg1)
	runtime.KeepAlive(drag)
	runtime.KeepAlive(success)
}

// Actions determines the bitmask of possible actions proposed by the source.
func (drag *Drag) Actions() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_actions(_arg0)
	runtime.KeepAlive(drag)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Content returns the GdkContentProvider associated to the GdkDrag object.
func (drag *Drag) Content() *ContentProvider {
	var _arg0 *C.GdkDrag            // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_content(_arg0)
	runtime.KeepAlive(drag)

	var _contentProvider *ContentProvider // out

	_contentProvider = wrapContentProvider(externglib.Take(unsafe.Pointer(_cret)))

	return _contentProvider
}

// Device returns the GdkDevice associated to the GdkDrag object.
func (drag *Drag) Device() Devicer {
	var _arg0 *C.GdkDrag   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_device(_arg0)
	runtime.KeepAlive(drag)

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Devicer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Devicer")
		}
		_device = rv
	}

	return _device
}

// Display gets the GdkDisplay that the drag object was created for.
func (drag *Drag) Display() *Display {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_display(_arg0)
	runtime.KeepAlive(drag)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// DragSurface returns the surface on which the drag icon should be rendered
// during the drag operation.
//
// Note that the surface may not be available until the drag operation has
// begun. GDK will move the surface in accordance with the ongoing drag
// operation. The surface is owned by drag and will be destroyed when the drag
// operation is over.
func (drag *Drag) DragSurface() Surfacer {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_drag_surface(_arg0)
	runtime.KeepAlive(drag)

	var _surface Surfacer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Surfacer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Surfacer")
			}
			_surface = rv
		}
	}

	return _surface
}

// Formats retrieves the formats supported by this GdkDrag object.
func (drag *Drag) Formats() *ContentFormats {
	var _arg0 *C.GdkDrag           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_formats(_arg0)
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

// SelectedAction determines the action chosen by the drag destination.
func (drag *Drag) SelectedAction() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_selected_action(_arg0)
	runtime.KeepAlive(drag)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Surface returns the GdkSurface where the drag originates.
func (drag *Drag) Surface() Surfacer {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))

	_cret = C.gdk_drag_get_surface(_arg0)
	runtime.KeepAlive(drag)

	var _surface Surfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Surfacer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Surfacer")
		}
		_surface = rv
	}

	return _surface
}

// SetHotspot sets the position of the drag surface that will be kept under the
// cursor hotspot.
//
// Initially, the hotspot is at the top left corner of the drag surface.
func (drag *Drag) SetHotspot(hotX, hotY int) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.int      // out
	var _arg2 C.int      // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(drag.Native()))
	_arg1 = C.int(hotX)
	_arg2 = C.int(hotY)

	C.gdk_drag_set_hotspot(_arg0, _arg1, _arg2)
	runtime.KeepAlive(drag)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// ConnectCancel: emitted when the drag operation is cancelled.
func (d *Drag) ConnectCancel(f func(reason DragCancelReason)) glib.SignalHandle {
	return d.Connect("cancel", f)
}

// ConnectDndFinished: emitted when the destination side has finished reading
// all data.
//
// The drag object can now free all miscellaneous data.
func (d *Drag) ConnectDndFinished(f func()) glib.SignalHandle {
	return d.Connect("dnd-finished", f)
}

// ConnectDropPerformed: emitted when the drop operation is performed on an
// accepting client.
func (d *Drag) ConnectDropPerformed(f func()) glib.SignalHandle {
	return d.Connect("drop-performed", f)
}

// DragBegin starts a drag and creates a new drag context for it.
//
// This function is called by the drag source. After this call, you probably
// want to set up the drag icon using the surface returned by
// gdk.Drag.GetDragSurface().
//
// This function returns a reference to the gdk.Drag object, but GTK keeps its
// own reference as well, as long as the DND operation is going on.
//
// Note: if actions include GDK_ACTION_MOVE, you need to listen for the
// gdk.Drag::dnd-finished signal and delete the data at the source if
// gdk.Drag.GetSelectedAction() returns GDK_ACTION_MOVE.
func DragBegin(surface Surfacer, device Devicer, content *ContentProvider, actions DragAction, dx, dy float64) Dragger {
	var _arg1 *C.GdkSurface         // out
	var _arg2 *C.GdkDevice          // out
	var _arg3 *C.GdkContentProvider // out
	var _arg4 C.GdkDragAction       // out
	var _arg5 C.double              // out
	var _arg6 C.double              // out
	var _cret *C.GdkDrag            // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg2 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg3 = (*C.GdkContentProvider)(unsafe.Pointer(content.Native()))
	_arg4 = C.GdkDragAction(actions)
	_arg5 = C.double(dx)
	_arg6 = C.double(dy)

	_cret = C.gdk_drag_begin(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(device)
	runtime.KeepAlive(content)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(dx)
	runtime.KeepAlive(dy)

	var _drag Dragger // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(Dragger)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Dragger")
			}
			_drag = rv
		}
	}

	return _drag
}
