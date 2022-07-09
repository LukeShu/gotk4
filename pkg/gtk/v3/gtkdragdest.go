// Code generated by girgen. DO NOT EDIT.

package gtk

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
import "C"

// GTypeDestDefaults returns the GType for the type DestDefaults.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDestDefaults() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "DestDefaults").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDestDefaults)
	return gtype
}

// DestDefaults enumeration specifies the various types of action that will be
// taken on behalf of the user for a drag destination site.
type DestDefaults C.guint

const (
	// DestDefaultMotion: if set for a widget, GTK+, during a drag over this
	// widget will check if the drag matches this widget’s list of possible
	// targets and actions. GTK+ will then call gdk_drag_status() as
	// appropriate.
	DestDefaultMotion DestDefaults = 0b1
	// DestDefaultHighlight: if set for a widget, GTK+ will draw a highlight on
	// this widget as long as a drag is over this widget and the widget drag
	// format and action are acceptable.
	DestDefaultHighlight DestDefaults = 0b10
	// DestDefaultDrop: if set for a widget, when a drop occurs, GTK+ will will
	// check if the drag matches this widget’s list of possible targets and
	// actions. If so, GTK+ will call gtk_drag_get_data() on behalf of the
	// widget. Whether or not the drop is successful, GTK+ will call
	// gtk_drag_finish(). If the action was a move, then if the drag was
	// successful, then TRUE will be passed for the delete parameter to
	// gtk_drag_finish().
	DestDefaultDrop DestDefaults = 0b100
	// DestDefaultAll: if set, specifies that all default actions should be
	// taken.
	DestDefaultAll DestDefaults = 0b111
)

func marshalDestDefaults(p uintptr) (interface{}, error) {
	return DestDefaults(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DestDefaults.
func (d DestDefaults) String() string {
	if d == 0 {
		return "DestDefaults(0)"
	}

	var builder strings.Builder
	builder.Grow(69)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case DestDefaultMotion:
			builder.WriteString("Motion|")
		case DestDefaultHighlight:
			builder.WriteString("Highlight|")
		case DestDefaultDrop:
			builder.WriteString("Drop|")
		case DestDefaultAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("DestDefaults(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DestDefaults) Has(other DestDefaults) bool {
	return (d & other) == other
}

// DragDestAddImageTargets: add the image targets supported by SelectionData to
// the target list of the drag destination. The targets are added with info = 0.
// If you need another value, use gtk_target_list_add_image_targets() and
// gtk_drag_dest_set_target_list().
func (widget *Widget) DragDestAddImageTargets() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_add_image_targets", _args[:], nil)

	runtime.KeepAlive(widget)
}

// DragDestAddTextTargets: add the text targets supported by SelectionData to
// the target list of the drag destination. The targets are added with info = 0.
// If you need another value, use gtk_target_list_add_text_targets() and
// gtk_drag_dest_set_target_list().
func (widget *Widget) DragDestAddTextTargets() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_add_text_targets", _args[:], nil)

	runtime.KeepAlive(widget)
}

// DragDestAddURITargets: add the URI targets supported by SelectionData to the
// target list of the drag destination. The targets are added with info = 0. If
// you need another value, use gtk_target_list_add_uri_targets() and
// gtk_drag_dest_set_target_list().
func (widget *Widget) DragDestAddURITargets() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_add_uri_targets", _args[:], nil)

	runtime.KeepAlive(widget)
}

// DragDestGetTargetList returns the list of targets this widget can accept from
// drag-and-drop.
//
// The function returns the following values:
//
//    - targetList (optional) or NULL if none.
//
func (widget *Widget) DragDestGetTargetList() *TargetList {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_gret := girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_get_target_list", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _targetList *TargetList // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_targetList = (*TargetList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gtk_target_list_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_targetList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _targetList
}

// DragDestGetTrackMotion returns whether the widget has been configured to
// always emit Widget::drag-motion signals.
//
// The function returns the following values:
//
//    - ok: TRUE if the widget always emits Widget::drag-motion events.
//
func (widget *Widget) DragDestGetTrackMotion() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_gret := girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_get_track_motion", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(widget)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// DragDestSetTargetList sets the target types that this widget can accept from
// drag-and-drop. The widget must first be made into a drag destination with
// gtk_drag_dest_set().
//
// The function takes the following parameters:
//
//    - targetList (optional): list of droppable targets, or NULL for none.
//
func (widget *Widget) DragDestSetTargetList(targetList *TargetList) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if targetList != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(targetList)))
	}

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_set_target_list", _args[:], nil)

	runtime.KeepAlive(widget)
	runtime.KeepAlive(targetList)
}

// DragDestSetTrackMotion tells the widget to emit Widget::drag-motion and
// Widget::drag-leave events regardless of the targets and the
// GTK_DEST_DEFAULT_MOTION flag.
//
// This may be used when a widget wants to do generic actions regardless of the
// targets that the source offers.
//
// The function takes the following parameters:
//
//    - trackMotion: whether to accept all targets.
//
func (widget *Widget) DragDestSetTrackMotion(trackMotion bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if trackMotion {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_set_track_motion", _args[:], nil)

	runtime.KeepAlive(widget)
	runtime.KeepAlive(trackMotion)
}

// DragDestUnset clears information about a drop destination set with
// gtk_drag_dest_set(). The widget will no longer receive notification of drags.
func (widget *Widget) DragDestUnset() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	girepository.MustFind("Gtk", "Widget").InvokeMethod("drag_dest_unset", _args[:], nil)

	runtime.KeepAlive(widget)
}
