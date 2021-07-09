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
		{T: externglib.Type(C.gdk_drag_cancel_reason_get_type()), F: marshalDragCancelReason},
		{T: externglib.Type(C.gdk_drag_get_type()), F: marshalDrag},
	})
}

// DragCancelReason: used in `GdkDrag` to the reason of a cancelled DND
// operation.
type DragCancelReason int

const (
	// NoTarget: there is no suitable drop target.
	DragCancelReasonNoTarget DragCancelReason = iota
	// UserCancelled: drag cancelled by the user
	DragCancelReasonUserCancelled
	// Error: unspecified error.
	DragCancelReasonError
)

func marshalDragCancelReason(p uintptr) (interface{}, error) {
	return DragCancelReason(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Drag: the `GdkDrag` object represents the source of an ongoing DND operation.
//
// A `GdkDrag` is created when a drag is started, and stays alive for duration
// of the DND operation. After a drag has been started with
// [func@Gdk.Drag.begin], the caller gets informed about the status of the
// ongoing drag operation with signals on the `GdkDrag` object.
//
// GTK provides a higher level abstraction based on top of these functions, and
// so they are not normally needed in GTK applications. See the "Drag and Drop"
// section of the GTK documentation for more information.
type Drag interface {
	gextras.Objector

	// DropDone informs GDK that the drop ended.
	//
	// Passing false for @success may trigger a drag cancellation animation.
	//
	// This function is called by the drag source, and should be the last call
	// before dropping the reference to the @drag.
	//
	// The `GdkDrag` will only take the first [method@Gdk.Drag.drop_done] call
	// as effective, if this function is called multiple times, all subsequent
	// calls will be ignored.
	DropDone(success bool)
	// Actions determines the bitmask of possible actions proposed by the
	// source.
	Actions() DragAction
	// Content returns the `GdkContentProvider` associated to the `GdkDrag`
	// object.
	Content() *ContentProviderClass
	// Device returns the `GdkDevice` associated to the `GdkDrag` object.
	Device() *DeviceClass
	// Display gets the `GdkDisplay` that the drag object was created for.
	Display() *DisplayClass
	// DragSurface returns the surface on which the drag icon should be rendered
	// during the drag operation.
	//
	// Note that the surface may not be available until the drag operation has
	// begun. GDK will move the surface in accordance with the ongoing drag
	// operation. The surface is owned by @drag and will be destroyed when the
	// drag operation is over.
	DragSurface() *SurfaceClass
	// Formats retrieves the formats supported by this `GdkDrag` object.
	Formats() *ContentFormats
	// SelectedAction determines the action chosen by the drag destination.
	SelectedAction() DragAction
	// Surface returns the `GdkSurface` where the drag originates.
	Surface() *SurfaceClass
	// SetHotspot sets the position of the drag surface that will be kept under
	// the cursor hotspot.
	//
	// Initially, the hotspot is at the top left corner of the drag surface.
	SetHotspot(hotX int, hotY int)
}

// DragClass implements the Drag interface.
type DragClass struct {
	*externglib.Object
}

var _ Drag = (*DragClass)(nil)

func wrapDrag(obj *externglib.Object) Drag {
	return &DragClass{
		Object: obj,
	}
}

func marshalDrag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrag(obj), nil
}

// DropDone informs GDK that the drop ended.
//
// Passing false for @success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should be the last call
// before dropping the reference to the @drag.
//
// The `GdkDrag` will only take the first [method@Gdk.Drag.drop_done] call as
// effective, if this function is called multiple times, all subsequent calls
// will be ignored.
func (d *DragClass) DropDone(success bool) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))
	if success {
		_arg1 = C.TRUE
	}

	C.gdk_drag_drop_done(_arg0, _arg1)
}

// Actions determines the bitmask of possible actions proposed by the source.
func (d *DragClass) Actions() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_actions(_arg0)

	var _dragAction DragAction // out

	_dragAction = (DragAction)(_cret)

	return _dragAction
}

// Content returns the `GdkContentProvider` associated to the `GdkDrag` object.
func (d *DragClass) Content() *ContentProviderClass {
	var _arg0 *C.GdkDrag            // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_content(_arg0)

	var _contentProvider *ContentProviderClass // out

	_contentProvider = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ContentProviderClass)

	return _contentProvider
}

// Device returns the `GdkDevice` associated to the `GdkDrag` object.
func (d *DragClass) Device() *DeviceClass {
	var _arg0 *C.GdkDrag   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_device(_arg0)

	var _device *DeviceClass // out

	_device = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*DeviceClass)

	return _device
}

// Display gets the `GdkDisplay` that the drag object was created for.
func (d *DragClass) Display() *DisplayClass {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_display(_arg0)

	var _display *DisplayClass // out

	_display = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*DisplayClass)

	return _display
}

// DragSurface returns the surface on which the drag icon should be rendered
// during the drag operation.
//
// Note that the surface may not be available until the drag operation has
// begun. GDK will move the surface in accordance with the ongoing drag
// operation. The surface is owned by @drag and will be destroyed when the drag
// operation is over.
func (d *DragClass) DragSurface() *SurfaceClass {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_drag_surface(_arg0)

	var _surface *SurfaceClass // out

	_surface = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*SurfaceClass)

	return _surface
}

// Formats retrieves the formats supported by this `GdkDrag` object.
func (d *DragClass) Formats() *ContentFormats {
	var _arg0 *C.GdkDrag           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(unsafe.Pointer(v)))
	})

	return _contentFormats
}

// SelectedAction determines the action chosen by the drag destination.
func (d *DragClass) SelectedAction() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_selected_action(_arg0)

	var _dragAction DragAction // out

	_dragAction = (DragAction)(_cret)

	return _dragAction
}

// Surface returns the `GdkSurface` where the drag originates.
func (d *DragClass) Surface() *SurfaceClass {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))

	_cret = C.gdk_drag_get_surface(_arg0)

	var _surface *SurfaceClass // out

	_surface = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*SurfaceClass)

	return _surface
}

// SetHotspot sets the position of the drag surface that will be kept under the
// cursor hotspot.
//
// Initially, the hotspot is at the top left corner of the drag surface.
func (d *DragClass) SetHotspot(hotX int, hotY int) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.int      // out
	var _arg2 C.int      // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer((&d).Native()))
	_arg1 = C.int(hotX)
	_arg2 = C.int(hotY)

	C.gdk_drag_set_hotspot(_arg0, _arg1, _arg2)
}
