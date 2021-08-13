// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drag_cancel_reason_get_type()), F: marshalDragCancelReason},
		{T: externglib.Type(C.gdk_drag_protocol_get_type()), F: marshalDragProtocol},
		{T: externglib.Type(C.gdk_drag_action_get_type()), F: marshalDragAction},
	})
}

// DragCancelReason: used in DragContext to the reason of a cancelled DND
// operation.
type DragCancelReason int

const (
	// DragCancelNoTarget: there is no suitable drop target.
	DragCancelNoTarget DragCancelReason = iota
	// DragCancelUserCancelled: drag cancelled by the user
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

// DragProtocol: used in DragContext to indicate the protocol according to which
// DND is done.
type DragProtocol int

const (
	// DragProtoNone: no protocol.
	DragProtoNone DragProtocol = iota
	// DragProtoMotif: motif DND protocol. No longer supported
	DragProtoMotif
	// DragProtoXdnd: xdnd protocol.
	DragProtoXdnd
	// DragProtoRootwin: extension to the Xdnd protocol for unclaimed root
	// window drops.
	DragProtoRootwin
	// DragProtoWin32Dropfiles: simple WM_DROPFILES protocol.
	DragProtoWin32Dropfiles
	// DragProtoOle2: complex OLE2 DND protocol (not implemented).
	DragProtoOle2
	// DragProtoLocal: intra-application DND.
	DragProtoLocal
	// DragProtoWayland: wayland DND protocol.
	DragProtoWayland
)

func marshalDragProtocol(p uintptr) (interface{}, error) {
	return DragProtocol(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for DragProtocol.
func (d DragProtocol) String() string {
	switch d {
	case DragProtoNone:
		return "None"
	case DragProtoMotif:
		return "Motif"
	case DragProtoXdnd:
		return "Xdnd"
	case DragProtoRootwin:
		return "Rootwin"
	case DragProtoWin32Dropfiles:
		return "Win32Dropfiles"
	case DragProtoOle2:
		return "Ole2"
	case DragProtoLocal:
		return "Local"
	case DragProtoWayland:
		return "Wayland"
	default:
		return fmt.Sprintf("DragProtocol(%d)", d)
	}
}

// DragAction: used in DragContext to indicate what the destination should do
// with the dropped data.
type DragAction int

const (
	// ActionDefault means nothing, and should not be used.
	ActionDefault DragAction = 0b1
	// ActionCopy: copy the data.
	ActionCopy DragAction = 0b10
	// ActionMove: move the data, i.e. first copy it, then delete it from the
	// source using the DELETE target of the X selection protocol.
	ActionMove DragAction = 0b100
	// ActionLink: add a link to the data. Note that this is only useful if
	// source and destination agree on what it means.
	ActionLink DragAction = 0b1000
	// ActionPrivate: special action which tells the source that the destination
	// will do something that the source doesn’t understand.
	ActionPrivate DragAction = 0b10000
	// ActionAsk: ask the user what to do with the data.
	ActionAsk DragAction = 0b100000
)

func marshalDragAction(p uintptr) (interface{}, error) {
	return DragAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for DragAction.
func (d DragAction) String() string {
	if d == 0 {
		return "DragAction(0)"
	}

	var builder strings.Builder
	builder.Grow(70)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case ActionDefault:
			builder.WriteString("Default|")
		case ActionCopy:
			builder.WriteString("Copy|")
		case ActionMove:
			builder.WriteString("Move|")
		case ActionLink:
			builder.WriteString("Link|")
		case ActionPrivate:
			builder.WriteString("Private|")
		case ActionAsk:
			builder.WriteString("Ask|")
		default:
			builder.WriteString(fmt.Sprintf("DragAction(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// DragAbort aborts a drag without dropping.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragAbort(context *DragContext, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.guint32(time_)

	C.gdk_drag_abort(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(time_)
}

// DragDrop drops on the current destination.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragDrop(context *DragContext, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.guint32(time_)

	C.gdk_drag_drop(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(time_)
}

// DragDropDone: inform GDK if the drop ended successfully. Passing FALSE for
// success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should be the last call
// before dropping the reference to the context.
//
// The DragContext will only take the first gdk_drag_drop_done() call as
// effective, if this function is called multiple times, all subsequent calls
// will be ignored.
func DragDropDone(context *DragContext, success bool) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		_arg2 = C.TRUE
	}

	C.gdk_drag_drop_done(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(success)
}

// DragDropSucceeded returns whether the dropped data has been successfully
// transferred. This function is intended to be used while handling a
// GDK_DROP_FINISHED event, its return value is meaningless at other times.
func DragDropSucceeded(context *DragContext) bool {
	var _arg1 *C.GdkDragContext // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_drop_succeeded(_arg1)

	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragFindWindowForScreen finds the destination window and DND protocol to use
// at the given pointer position.
//
// This function is called by the drag source to obtain the dest_window and
// protocol parameters for gdk_drag_motion().
func DragFindWindowForScreen(context *DragContext, dragWindow Windower, screen *Screen, xRoot int, yRoot int) (Windower, DragProtocol) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GdkWindow      // out
	var _arg3 *C.GdkScreen      // out
	var _arg4 C.gint            // out
	var _arg5 C.gint            // out
	var _arg6 *C.GdkWindow      // in
	var _arg7 C.GdkDragProtocol // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(dragWindow.Native()))
	_arg3 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg4 = C.gint(xRoot)
	_arg5 = C.gint(yRoot)

	C.gdk_drag_find_window_for_screen(_arg1, _arg2, _arg3, _arg4, _arg5, &_arg6, &_arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(dragWindow)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(xRoot)
	runtime.KeepAlive(yRoot)

	var _destWindow Windower   // out
	var _protocol DragProtocol // out

	_destWindow = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_arg6)))).(Windower)
	_protocol = DragProtocol(_arg7)

	return _destWindow, _protocol
}

// DragMotion updates the drag context when the pointer moves or the set of
// actions changes.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragMotion(context *DragContext, destWindow Windower, protocol DragProtocol, xRoot int, yRoot int, suggestedAction DragAction, possibleActions DragAction, time_ uint32) bool {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GdkWindow      // out
	var _arg3 C.GdkDragProtocol // out
	var _arg4 C.gint            // out
	var _arg5 C.gint            // out
	var _arg6 C.GdkDragAction   // out
	var _arg7 C.GdkDragAction   // out
	var _arg8 C.guint32         // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(destWindow.Native()))
	_arg3 = C.GdkDragProtocol(protocol)
	_arg4 = C.gint(xRoot)
	_arg5 = C.gint(yRoot)
	_arg6 = C.GdkDragAction(suggestedAction)
	_arg7 = C.GdkDragAction(possibleActions)
	_arg8 = C.guint32(time_)

	_cret = C.gdk_drag_motion(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)

	runtime.KeepAlive(context)
	runtime.KeepAlive(destWindow)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(xRoot)
	runtime.KeepAlive(yRoot)
	runtime.KeepAlive(suggestedAction)
	runtime.KeepAlive(possibleActions)
	runtime.KeepAlive(time_)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragStatus selects one of the actions offered by the drag source.
//
// This function is called by the drag destination in response to
// gdk_drag_motion() called by the drag source.
func DragStatus(context *DragContext, action DragAction, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.GdkDragAction   // out
	var _arg3 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.GdkDragAction(action)
	_arg3 = C.guint32(time_)

	C.gdk_drag_status(_arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(action)
	runtime.KeepAlive(time_)
}

// DropFinish ends the drag operation after a drop.
//
// This function is called by the drag destination.
func DropFinish(context *DragContext, success bool, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out
	var _arg3 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		_arg2 = C.TRUE
	}
	_arg3 = C.guint32(time_)

	C.gdk_drop_finish(_arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(success)
	runtime.KeepAlive(time_)
}

// DropReply accepts or rejects a drop.
//
// This function is called by the drag destination in response to a drop
// initiated by the drag source.
func DropReply(context *DragContext, accepted bool, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out
	var _arg3 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if accepted {
		_arg2 = C.TRUE
	}
	_arg3 = C.guint32(time_)

	C.gdk_drop_reply(_arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(accepted)
	runtime.KeepAlive(time_)
}
