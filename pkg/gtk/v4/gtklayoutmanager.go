// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_layout_manager_get_type()), F: marshalLayoutManagerer},
	})
}

// LayoutManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type LayoutManagerOverrider interface {
	// Allocate assigns the given width, height, and baseline to a widget, and
	// computes the position and sizes of the children of the widget using the
	// layout management policy of manager.
	//
	// The function takes the following parameters:
	//
	//    - widget: GtkWidget using manager.
	//    - width: new width of the widget.
	//    - height: new height of the widget.
	//    - baseline position of the widget, or -1.
	//
	Allocate(widget Widgetter, width, height, baseline int)
	// CreateLayoutChild: create a LayoutChild instance for the given for_child
	// widget.
	//
	// The function takes the following parameters:
	//
	//    - widget using the manager.
	//    - forChild: child of widget.
	//
	// The function returns the following values:
	//
	//    - layoutChild: LayoutChild.
	//
	CreateLayoutChild(widget, forChild Widgetter) LayoutChilder
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	RequestMode(widget Widgetter) SizeRequestMode
	// Measure measures the size of the widget using manager, for the given
	// orientation and size.
	//
	// See the gtk.Widget documentation on layout management for more details.
	//
	// The function takes the following parameters:
	//
	//    - widget: GtkWidget using manager.
	//    - orientation to measure.
	//    - forSize: size for the opposite of orientation; for instance, if the
	//      orientation is GTK_ORIENTATION_HORIZONTAL, this is the height of the
	//      widget; if the orientation is GTK_ORIENTATION_VERTICAL, this is the
	//      width of the widget. This allows to measure the height for the given
	//      width, and the width for the given height. Use -1 if the size is not
	//      known.
	//
	// The function returns the following values:
	//
	//    - minimum (optional) size for the given size and orientation.
	//    - natural (optional): natural, or preferred size for the given size and
	//      orientation.
	//    - minimumBaseline (optional): baseline position for the minimum size.
	//    - naturalBaseline (optional): baseline position for the natural size.
	//
	Measure(widget Widgetter, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int)
	Root()
	Unroot()
}

// LayoutManager: layout managers are delegate classes that handle the preferred
// size and the allocation of a widget.
//
// You typically subclass GtkLayoutManager if you want to implement a layout
// policy for the children of a widget, or if you want to determine the size of
// a widget depending on its contents.
//
// Each GtkWidget can only have a GtkLayoutManager instance associated to it at
// any given time; it is possible, though, to replace the layout manager
// instance using gtk.Widget.SetLayoutManager().
//
//
// Layout properties
//
// A layout manager can expose properties for controlling the layout of each
// child, by creating an object type derived from gtk.LayoutChild and installing
// the properties on it as normal GObject properties.
//
// Each GtkLayoutChild instance storing the layout properties for a specific
// child is created through the gtk.LayoutManager.GetLayoutChild() method; a
// GtkLayoutManager controls the creation of its GtkLayoutChild instances by
// overriding the GtkLayoutManagerClass.create_layout_child() virtual function.
// The typical implementation should look like:
//
//    static GtkLayoutChild *
//    create_layout_child (GtkLayoutManager *manager,
//                         GtkWidget        *container,
//                         GtkWidget        *child)
//    {
//      return g_object_new (your_layout_child_get_type (),
//                           "layout-manager", manager,
//                           "child-widget", child,
//                           NULL);
//    }
//
//
// The gtk.LayoutChild:layout-manager and gtk.LayoutChild:child-widget
// properties on the newly created GtkLayoutChild instance are mandatory. The
// GtkLayoutManager will cache the newly created GtkLayoutChild instance until
// the widget is removed from its parent, or the parent removes the layout
// manager.
//
// Each GtkLayoutManager instance creating a GtkLayoutChild should use
// gtk.LayoutManager.GetLayoutChild() every time it needs to query the layout
// properties; each GtkLayoutChild instance should call
// gtk.LayoutManager.LayoutChanged() every time a property is updated, in order
// to queue a new size measuring and allocation.
type LayoutManager struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*LayoutManager)(nil)
)

// LayoutManagerer describes types inherited from class LayoutManager.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type LayoutManagerer interface {
	externglib.Objector
	baseLayoutManager() *LayoutManager
}

var _ LayoutManagerer = (*LayoutManager)(nil)

func wrapLayoutManager(obj *externglib.Object) *LayoutManager {
	return &LayoutManager{
		Object: obj,
	}
}

func marshalLayoutManagerer(p uintptr) (interface{}, error) {
	return wrapLayoutManager(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (manager *LayoutManager) baseLayoutManager() *LayoutManager {
	return manager
}

// BaseLayoutManager returns the underlying base object.
func BaseLayoutManager(obj LayoutManagerer) *LayoutManager {
	return obj.baseLayoutManager()
}

// Allocate assigns the given width, height, and baseline to a widget, and
// computes the position and sizes of the children of the widget using the
// layout management policy of manager.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget using manager.
//    - width: new width of the widget.
//    - height: new height of the widget.
//    - baseline position of the widget, or -1.
//
func (manager *LayoutManager) Allocate(widget Widgetter, width, height, baseline int) {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _arg4 C.int               // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(width)
	_arg3 = C.int(height)
	_arg4 = C.int(baseline)

	C.gtk_layout_manager_allocate(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(baseline)
}

// LayoutChild retrieves a GtkLayoutChild instance for the GtkLayoutManager,
// creating one if necessary.
//
// The child widget must be a child of the widget using manager.
//
// The GtkLayoutChild instance is owned by the GtkLayoutManager, and is
// guaranteed to exist as long as child is a child of the GtkWidget using the
// given GtkLayoutManager.
//
// The function takes the following parameters:
//
//    - child: GtkWidget.
//
// The function returns the following values:
//
//    - layoutChild: GtkLayoutChild.
//
func (manager *LayoutManager) LayoutChild(child Widgetter) LayoutChilder {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.GtkLayoutChild   // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_layout_manager_get_layout_child(_arg0, _arg1)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(child)

	var _layoutChild LayoutChilder // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.LayoutChilder is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(LayoutChilder)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.LayoutChilder")
		}
		_layoutChild = rv
	}

	return _layoutChild
}

// RequestMode retrieves the request mode of manager.
//
// The function returns the following values:
//
//    - sizeRequestMode: GtkSizeRequestMode.
//
func (manager *LayoutManager) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkLayoutManager  // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_layout_manager_get_request_mode(_arg0)
	runtime.KeepAlive(manager)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

// Widget retrieves the GtkWidget using the given GtkLayoutManager.
//
// The function returns the following values:
//
//    - widget (optional): GtkWidget.
//
func (manager *LayoutManager) Widget() Widgetter {
	var _arg0 *C.GtkLayoutManager // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_layout_manager_get_widget(_arg0)
	runtime.KeepAlive(manager)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// LayoutChanged queues a resize on the GtkWidget using manager, if any.
//
// This function should be called by subclasses of GtkLayoutManager in response
// to changes to their layout management policies.
func (manager *LayoutManager) LayoutChanged() {
	var _arg0 *C.GtkLayoutManager // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))

	C.gtk_layout_manager_layout_changed(_arg0)
	runtime.KeepAlive(manager)
}

// Measure measures the size of the widget using manager, for the given
// orientation and size.
//
// See the gtk.Widget documentation on layout management for more details.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget using manager.
//    - orientation to measure.
//    - forSize: size for the opposite of orientation; for instance, if the
//      orientation is GTK_ORIENTATION_HORIZONTAL, this is the height of the
//      widget; if the orientation is GTK_ORIENTATION_VERTICAL, this is the width
//      of the widget. This allows to measure the height for the given width, and
//      the width for the given height. Use -1 if the size is not known.
//
// The function returns the following values:
//
//    - minimum (optional) size for the given size and orientation.
//    - natural (optional): natural, or preferred size for the given size and
//      orientation.
//    - minimumBaseline (optional): baseline position for the minimum size.
//    - naturalBaseline (optional): baseline position for the natural size.
//
func (manager *LayoutManager) Measure(widget Widgetter, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int) {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.GtkOrientation    // out
	var _arg3 C.int               // out
	var _arg4 C.int               // in
	var _arg5 C.int               // in
	var _arg6 C.int               // in
	var _arg7 C.int               // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkOrientation(orientation)
	_arg3 = C.int(forSize)

	C.gtk_layout_manager_measure(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)
	runtime.KeepAlive(manager)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(forSize)

	var _minimum int         // out
	var _natural int         // out
	var _minimumBaseline int // out
	var _naturalBaseline int // out

	_minimum = int(_arg4)
	_natural = int(_arg5)
	_minimumBaseline = int(_arg6)
	_naturalBaseline = int(_arg7)

	return _minimum, _natural, _minimumBaseline, _naturalBaseline
}
