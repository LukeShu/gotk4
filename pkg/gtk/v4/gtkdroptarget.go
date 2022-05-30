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
// extern gboolean _gotk4_gtk4_DropTarget_ConnectAccept(gpointer, GdkDrop*, guintptr);
// extern void _gotk4_gtk4_DropTarget_ConnectLeave(gpointer, guintptr);
import "C"

// glib.Type values for gtkdroptarget.go.
var GTypeDropTarget = coreglib.Type(C.gtk_drop_target_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeDropTarget, F: marshalDropTarget},
	})
}

// DropTargetOverrider contains methods that are overridable.
type DropTargetOverrider interface {
}

// DropTarget: GtkDropTarget is an event controller to receive Drag-and-Drop
// operations.
//
// The most basic way to use a GtkDropTarget to receive drops on a widget is to
// create it via gtk.DropTarget.New, passing in the GType of the data you want
// to receive and connect to the gtk.DropTarget::drop signal to receive the
// data:
//
//    static gboolean
//    on_drop (GtkDropTarget *target,
//             const GValue  *value,
//             double         x,
//             double         y,
//             gpointer       data)
//    {
//      MyWidget *self = data;
//
//      // Call the appropriate setter depending on the type of data
//      // that we received
//      if (G_VALUE_HOLDS (value, G_TYPE_FILE))
//        my_widget_set_file (self, g_value_get_object (value));
//      else if (G_VALUE_HOLDS (value, GDK_TYPE_PIXBUF))
//        my_widget_set_pixbuf (self, g_value_get_object (value));
//      else
//        return FALSE;
//
//      return TRUE;
//    }
//
//    static void
//    my_widget_init (MyWidget *self)
//    {
//      GtkDropTarget *target =
//        gtk_drop_target_new (G_TYPE_INVALID, GDK_ACTION_COPY);
//
//      // This widget accepts two types of drop types: GFile objects
//      // and GdkPixbuf objects
//      gtk_drop_target_set_gtypes (target, (GTypes [2]) {
//        G_TYPE_FILE,
//        GDK_TYPE_PIXBUF,
//      }, 2);
//
//      gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (target));
//    }
//
//
// GtkDropTarget supports more options, such as:
//
//    * rejecting potential drops via the gtk.DropTarget::accept signal
//      and the gtk.DropTarget.Reject() function to let other drop
//      targets handle the drop
//    * tracking an ongoing drag operation before the drop via the
//      gtk.DropTarget::enter, gtk.DropTarget::motion and
//      gtk.DropTarget::leave signals
//    * configuring how to receive data by setting the
//      gtk.DropTarget:preload property and listening for its
//      availability via the gtk.DropTarget:value property
//
// However, GtkDropTarget is ultimately modeled in a synchronous way and only
// supports data transferred via GType. If you want full control over an ongoing
// drop, the gtk.DropTargetAsync object gives you this ability.
//
// While a pointer is dragged over the drop target's widget and the drop has not
// been rejected, that widget will receive the GTK_STATE_FLAG_DROP_ACTIVE state,
// which can be used to style the widget.
//
// If you are not interested in receiving the drop, but just want to update UI
// state during a Drag-and-Drop operation (e.g. switching tabs), you can use
// gtk.DropControllerMotion.
type DropTarget struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*DropTarget)(nil)
)

func classInitDropTargetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDropTarget(obj *coreglib.Object) *DropTarget {
	return &DropTarget{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalDropTarget(p uintptr) (interface{}, error) {
	return wrapDropTarget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_DropTarget_ConnectAccept
func _gotk4_gtk4_DropTarget_ConnectAccept(arg0 C.gpointer, arg1 *C.GdkDrop, arg2 C.guintptr) (cret C.gboolean) {
	var f func(drop gdk.Dropper) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(drop gdk.Dropper) (ok bool))
	}

	var _drop gdk.Dropper // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Dropper is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Dropper)
			return ok
		})
		rv, ok := casted.(gdk.Dropper)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dropper")
		}
		_drop = rv
	}

	ok := f(_drop)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectAccept is emitted on the drop site when a drop operation is about to
// begin.
//
// If the drop is not accepted, FALSE will be returned and the drop target will
// ignore the drop. If TRUE is returned, the drop is accepted for now but may be
// rejected later via a call to gtk.DropTarget.Reject() or ultimately by
// returning FALSE from a gtk.DropTarget::drop handler.
//
// The default handler for this signal decides whether to accept the drop based
// on the formats provided by the drop.
//
// If the decision whether the drop will be accepted or rejected depends on the
// data, this function should return TRUE, the gtk.DropTarget:preload property
// should be set and the value should be inspected via the ::notify:value
// signal, calling gtk.DropTarget.Reject() if required.
func (self *DropTarget) ConnectAccept(f func(drop gdk.Dropper) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "accept", false, unsafe.Pointer(C._gotk4_gtk4_DropTarget_ConnectAccept), f)
}

//export _gotk4_gtk4_DropTarget_ConnectLeave
func _gotk4_gtk4_DropTarget_ConnectLeave(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectLeave is emitted on the drop site when the pointer leaves the widget.
//
// Its main purpose it to undo things done in gtk.DropTarget::enter.
func (self *DropTarget) ConnectLeave(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "leave", false, unsafe.Pointer(C._gotk4_gtk4_DropTarget_ConnectLeave), f)
}

// Drop gets the currently handled drop operation.
//
// If no drop operation is going on, NULL is returned.
//
// The function returns the following values:
//
//    - drop (optional): current drop.
//
func (self *DropTarget) Drop() gdk.Dropper {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**DropTarget)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "DropTarget").InvokeMethod("get_drop", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _drop gdk.Dropper // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Dropper)
				return ok
			})
			rv, ok := casted.(gdk.Dropper)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dropper")
			}
			_drop = rv
		}
	}

	return _drop
}

// Formats gets the data formats that this drop target accepts.
//
// If the result is NULL, all formats are expected to be supported.
//
// The function returns the following values:
//
//    - contentFormats (optional): supported data formats.
//
func (self *DropTarget) Formats() *gdk.ContentFormats {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**DropTarget)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "DropTarget").InvokeMethod("get_formats", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _contentFormats *gdk.ContentFormats // out

	if _cret != nil {
		_contentFormats = (*gdk.ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_contentFormats)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
			},
		)
	}

	return _contentFormats
}

// Preload gets whether data should be preloaded on hover.
//
// The function returns the following values:
//
//    - ok: TRUE if drop data should be preloaded.
//
func (self *DropTarget) Preload() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**DropTarget)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "DropTarget").InvokeMethod("get_preload", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Value gets the current drop data, as a GValue.
//
// The function returns the following values:
//
//    - value (optional): current drop data.
//
func (self *DropTarget) Value() *coreglib.Value {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**DropTarget)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "DropTarget").InvokeMethod("get_value", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _value *coreglib.Value // out

	if _cret != nil {
		_value = coreglib.ValueFromNative(unsafe.Pointer(_cret))
	}

	return _value
}

// Reject rejects the ongoing drop operation.
//
// If no drop operation is ongoing, i.e when gtk.DropTarget:drop is NULL, this
// function does nothing.
//
// This function should be used when delaying the decision on whether to accept
// a drag or not until after reading the data.
func (self *DropTarget) Reject() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**DropTarget)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "DropTarget").InvokeMethod("reject", args[:], nil)

	runtime.KeepAlive(self)
}

// SetPreload sets whether data should be preloaded on hover.
//
// The function takes the following parameters:
//
//    - preload: TRUE to preload drop data.
//
func (self *DropTarget) SetPreload(preload bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if preload {
		_arg1 = C.TRUE
	}
	*(**DropTarget)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "DropTarget").InvokeMethod("set_preload", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(preload)
}
