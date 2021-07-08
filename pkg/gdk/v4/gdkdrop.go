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
		{T: externglib.Type(C.gdk_drop_get_type()), F: marshalDrop},
	})
}

// Drop: the `GdkDrop` object represents the target of an ongoing DND operation.
//
// Possible drop sites get informed about the status of the ongoing drag
// operation with events of type GDK_DRAG_ENTER, GDK_DRAG_LEAVE, GDK_DRAG_MOTION
// and GDK_DROP_START. The `GdkDrop` object can be obtained from these
// [class@Gdk.Event] types using [method@Gdk.DNDEvent.get_drop].
//
// The actual data transfer is initiated from the target side via an async read,
// using one of the `GdkDrop` methods for this purpose:
// [method@Gdk.Drop.read_async] or [method@Gdk.Drop.read_value_async].
//
// GTK provides a higher level abstraction based on top of these functions, and
// so they are not normally needed in GTK applications. See the "Drag and Drop"
// section of the GTK documentation for more information.
type Drop interface {
	gextras.Objector

	// Finish ends the drag operation after a drop.
	//
	// The @action must be a single action selected from the actions available
	// via [method@Gdk.Drop.get_actions].
	Finish(action DragAction)
	// Actions returns the possible actions for this `GdkDrop`.
	//
	// If this value contains multiple actions - i.e.
	// [func@Gdk.DragAction.is_unique] returns false for the result -
	// [method@Gdk.Drop.finish] must choose the action to use when accepting the
	// drop. This will only happen if you passed GDK_ACTION_ASK as one of the
	// possible actions in [method@Gdk.Drop.status]. GDK_ACTION_ASK itself will
	// not be included in the actions returned by this function.
	//
	// This value may change over the lifetime of the [class@Gdk.Drop] both as a
	// response to source side actions as well as to calls to
	// [method@Gdk.Drop.status] or [method@Gdk.Drop.finish]. The source side
	// will not change this value anymore once a drop has started.
	Actions() DragAction
	// Device returns the `GdkDevice` performing the drop.
	Device() Device
	// Display gets the `GdkDisplay` that @self was created for.
	Display() Display
	// Drag: if this is an in-app drag-and-drop operation, returns the `GdkDrag`
	// that corresponds to this drop.
	//
	// If it is not, nil is returned.
	Drag() Drag
	// Formats returns the `GdkContentFormats` that the drop offers the data to
	// be read in.
	Formats() *ContentFormats
	// Surface returns the `GdkSurface` performing the drop.
	Surface() Surface
	// Status selects all actions that are potentially supported by the
	// destination.
	//
	// When calling this function, do not restrict the passed in actions to the
	// ones provided by [method@Gdk.Drop.get_actions]. Those actions may change
	// in the future, even depending on the actions you provide here.
	//
	// The @preferred action is a hint to the drag'n'drop mechanism about which
	// action to use when multiple actions are possible.
	//
	// This function should be called by drag destinations in response to
	// GDK_DRAG_ENTER or GDK_DRAG_MOTION events. If the destination does not yet
	// know the exact actions it supports, it should set any possible actions
	// first and then later call this function again.
	Status(actions DragAction, preferred DragAction)
}

// DropClass implements the Drop interface.
type DropClass struct {
	*externglib.Object
}

var _ Drop = (*DropClass)(nil)

func wrapDrop(obj *externglib.Object) Drop {
	return &DropClass{
		Object: obj,
	}
}

func marshalDrop(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrop(obj), nil
}

// Finish ends the drag operation after a drop.
//
// The @action must be a single action selected from the actions available via
// [method@Gdk.Drop.get_actions].
func (s *DropClass) Finish(action DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkDragAction(action)

	C.gdk_drop_finish(_arg0, _arg1)
}

// Actions returns the possible actions for this `GdkDrop`.
//
// If this value contains multiple actions - i.e.
// [func@Gdk.DragAction.is_unique] returns false for the result -
// [method@Gdk.Drop.finish] must choose the action to use when accepting the
// drop. This will only happen if you passed GDK_ACTION_ASK as one of the
// possible actions in [method@Gdk.Drop.status]. GDK_ACTION_ASK itself will not
// be included in the actions returned by this function.
//
// This value may change over the lifetime of the [class@Gdk.Drop] both as a
// response to source side actions as well as to calls to
// [method@Gdk.Drop.status] or [method@Gdk.Drop.finish]. The source side will
// not change this value anymore once a drop has started.
func (s *DropClass) Actions() DragAction {
	var _arg0 *C.GdkDrop      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_actions(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Device returns the `GdkDevice` performing the drop.
func (s *DropClass) Device() Device {
	var _arg0 *C.GdkDrop   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_device(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

// Display gets the `GdkDisplay` that @self was created for.
func (s *DropClass) Display() Display {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

// Drag: if this is an in-app drag-and-drop operation, returns the `GdkDrag`
// that corresponds to this drop.
//
// If it is not, nil is returned.
func (s *DropClass) Drag() Drag {
	var _arg0 *C.GdkDrop // out
	var _cret *C.GdkDrag // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_drag(_arg0)

	var _drag Drag // out

	_drag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Drag)

	return _drag
}

// Formats returns the `GdkContentFormats` that the drop offers the data to be
// read in.
func (s *DropClass) Formats() *ContentFormats {
	var _arg0 *C.GdkDrop           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(unsafe.Pointer(v)))
	})

	return _contentFormats
}

// Surface returns the `GdkSurface` performing the drop.
func (s *DropClass) Surface() Surface {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_drop_get_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _surface
}

// Status selects all actions that are potentially supported by the destination.
//
// When calling this function, do not restrict the passed in actions to the ones
// provided by [method@Gdk.Drop.get_actions]. Those actions may change in the
// future, even depending on the actions you provide here.
//
// The @preferred action is a hint to the drag'n'drop mechanism about which
// action to use when multiple actions are possible.
//
// This function should be called by drag destinations in response to
// GDK_DRAG_ENTER or GDK_DRAG_MOTION events. If the destination does not yet
// know the exact actions it supports, it should set any possible actions first
// and then later call this function again.
func (s *DropClass) Status(actions DragAction, preferred DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out
	var _arg2 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkDragAction(actions)
	_arg2 = C.GdkDragAction(preferred)

	C.gdk_drop_status(_arg0, _arg1, _arg2)
}
