// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drag_source_get_type()), F: marshalDragSource},
	})
}

// DragSource: `GtkDragSource` is an event controller to initiate Drag-And-Drop
// operations.
//
// `GtkDragSource` can be set up with the necessary ingredients for a DND
// operation ahead of time. This includes the source for the data that is being
// transferred, in the form of a [class@Gdk.ContentProvider], the desired
// action, and the icon to use during the drag operation. After setting it up,
// the drag source must be added to a widget as an event controller, using
// [method@Gtk.Widget.add_controller].
//
// “`c static void my_widget_init (MyWidget *self) { GtkDragSource *drag_source
// = gtk_drag_source_new ();
//
//    g_signal_connect (drag_source, "prepare", G_CALLBACK (on_drag_prepare), self);
//    g_signal_connect (drag_source, "drag-begin", G_CALLBACK (on_drag_begin), self);
//
//    gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (drag_source));
//
// } “`
//
// Setting up the content provider and icon ahead of time only makes sense when
// the data does not change. More commonly, you will want to set them up just in
// time. To do so, `GtkDragSource` has [signal@Gtk.DragSource::prepare] and
// [signal@Gtk.DragSource::drag-begin] signals.
//
// The ::prepare signal is emitted before a drag is started, and can be used to
// set the content provider and actions that the drag should be started with.
//
// “`c static GdkContentProvider * on_drag_prepare (GtkDragSource *source,
// double x, double y, MyWidget *self) { // This widget supports two types of
// content: GFile objects // and GdkPixbuf objects; GTK will handle the
// serialization // of these types automatically GFile *file =
// my_widget_get_file (self); GdkPixbuf *pixbuf = my_widget_get_pixbuf (self);
//
//    return gdk_content_provider_new_union ((GdkContentProvider *[2]) {
//        gdk_content_provider_new_typed (G_TYPE_FILE, file),
//        gdk_content_provider_new_typed (GDK_TYPE_PIXBUF, pixbuf),
//      }, 2);
//
// } “`
//
// The ::drag-begin signal is emitted after the `GdkDrag` object has been
// created, and can be used to set up the drag icon.
//
// “`c static void on_drag_begin (GtkDragSource *source, GtkDrag *drag, MyWidget
// *self) { // Set the widget as the drag icon GdkPaintable *paintable =
// gtk_widget_paintable_new (GTK_WIDGET (self)); gtk_drag_source_set_icon
// (source, paintable, 0, 0); g_object_unref (paintable); } “`
//
// During the DND operation, `GtkDragSource` emits signals that can be used to
// obtain updates about the status of the operation, but it is not normally
// necessary to connect to any signals, except for one case: when the supported
// actions include GDK_ACTION_MOVE, you need to listen for the
// [signal@Gtk.DragSource::drag-end] signal and delete the data after it has
// been transferred.
type DragSource interface {
	gextras.Objector

	// DragCancel cancels a currently ongoing drag operation.
	DragCancel()
	// Actions gets the actions that are currently set on the `GtkDragSource`.
	Actions() gdk.DragAction
	// Content gets the current content provider of a `GtkDragSource`.
	Content() gdk.ContentProvider
	// Drag returns the underlying `GdkDrag` object for an ongoing drag.
	Drag() gdk.Drag
	// SetActions sets the actions on the `GtkDragSource`.
	//
	// During a DND operation, the actions are offered to potential drop
	// targets. If @actions include GDK_ACTION_MOVE, you need to listen to the
	// [signal@Gtk.DragSource::drag-end] signal and handle @delete_data being
	// true.
	//
	// This function can be called before a drag is started, or in a handler for
	// the [signal@Gtk.DragSource::prepare] signal.
	SetActions(actions gdk.DragAction)
	// SetContent sets a content provider on a `GtkDragSource`.
	//
	// When the data is requested in the cause of a DND operation, it will be
	// obtained from the content provider.
	//
	// This function can be called before a drag is started, or in a handler for
	// the [signal@Gtk.DragSource::prepare] signal.
	//
	// You may consider setting the content provider back to nil in a
	// [signal@Gtk.DragSource::drag-end] signal handler.
	SetContent(content gdk.ContentProvider)
}

// DragSourceClass implements the DragSource interface.
type DragSourceClass struct {
	GestureSingleClass
}

var _ DragSource = (*DragSourceClass)(nil)

func wrapDragSource(obj *externglib.Object) DragSource {
	return &DragSourceClass{
		GestureSingleClass: GestureSingleClass{
			GestureClass: GestureClass{
				EventControllerClass: EventControllerClass{
					Object: obj,
				},
			},
		},
	}
}

func marshalDragSource(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDragSource(obj), nil
}

// NewDragSource creates a new `GtkDragSource` object.
func NewDragSource() DragSource {
	var _cret *C.GtkDragSource // in

	_cret = C.gtk_drag_source_new()

	var _dragSource DragSource // out

	_dragSource = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(DragSource)

	return _dragSource
}

// DragCancel cancels a currently ongoing drag operation.
func (s *DragSourceClass) DragCancel() {
	var _arg0 *C.GtkDragSource // out

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(s.Native()))

	C.gtk_drag_source_drag_cancel(_arg0)
}

// Actions gets the actions that are currently set on the `GtkDragSource`.
func (s *DragSourceClass) Actions() gdk.DragAction {
	var _arg0 *C.GtkDragSource // out
	var _cret C.GdkDragAction  // in

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drag_source_get_actions(_arg0)

	var _dragAction gdk.DragAction // out

	_dragAction = gdk.DragAction(_cret)

	return _dragAction
}

// Content gets the current content provider of a `GtkDragSource`.
func (s *DragSourceClass) Content() gdk.ContentProvider {
	var _arg0 *C.GtkDragSource      // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drag_source_get_content(_arg0)

	var _contentProvider gdk.ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.ContentProvider)

	return _contentProvider
}

// Drag returns the underlying `GdkDrag` object for an ongoing drag.
func (s *DragSourceClass) Drag() gdk.Drag {
	var _arg0 *C.GtkDragSource // out
	var _cret *C.GdkDrag       // in

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drag_source_get_drag(_arg0)

	var _drag gdk.Drag // out

	_drag = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Drag)

	return _drag
}

// SetActions sets the actions on the `GtkDragSource`.
//
// During a DND operation, the actions are offered to potential drop targets. If
// @actions include GDK_ACTION_MOVE, you need to listen to the
// [signal@Gtk.DragSource::drag-end] signal and handle @delete_data being true.
//
// This function can be called before a drag is started, or in a handler for the
// [signal@Gtk.DragSource::prepare] signal.
func (s *DragSourceClass) SetActions(actions gdk.DragAction) {
	var _arg0 *C.GtkDragSource // out
	var _arg1 C.GdkDragAction  // out

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkDragAction(actions)

	C.gtk_drag_source_set_actions(_arg0, _arg1)
}

// SetContent sets a content provider on a `GtkDragSource`.
//
// When the data is requested in the cause of a DND operation, it will be
// obtained from the content provider.
//
// This function can be called before a drag is started, or in a handler for the
// [signal@Gtk.DragSource::prepare] signal.
//
// You may consider setting the content provider back to nil in a
// [signal@Gtk.DragSource::drag-end] signal handler.
func (s *DragSourceClass) SetContent(content gdk.ContentProvider) {
	var _arg0 *C.GtkDragSource      // out
	var _arg1 *C.GdkContentProvider // out

	_arg0 = (*C.GtkDragSource)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkContentProvider)(unsafe.Pointer(content.Native()))

	C.gtk_drag_source_set_content(_arg0, _arg1)
}
