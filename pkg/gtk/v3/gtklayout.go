// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeLayout returns the GType for the type Layout.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeLayout() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Layout").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalLayout)
	return gtype
}

// LayoutOverrider contains methods that are overridable.
type LayoutOverrider interface {
}

// Layout is similar to DrawingArea in that it’s a “blank slate” and doesn’t do
// anything except paint a blank background by default. It’s different in that
// it supports scrolling natively due to implementing Scrollable, and can
// contain child widgets since it’s a Container.
//
// If you just want to draw, a DrawingArea is a better choice since it has lower
// overhead. If you just need to position child widgets at specific points, then
// Fixed provides that functionality on its own.
//
// When handling expose events on a Layout, you must draw to the Window returned
// by gtk_layout_get_bin_window(), rather than to the one returned by
// gtk_widget_get_window() as you would for a DrawingArea.
type Layout struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Scrollable
}

var (
	_ Containerer       = (*Layout)(nil)
	_ coreglib.Objector = (*Layout)(nil)
)

func classInitLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapLayout(obj *coreglib.Object) *Layout {
	return &Layout{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
		Object: obj,
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalLayout(p uintptr) (interface{}, error) {
	return wrapLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLayout creates a new Layout. Unless you have a specific adjustment you’d
// like the layout to use for scrolling, pass NULL for hadjustment and
// vadjustment.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): horizontal scroll adjustment, or NULL.
//    - vadjustment (optional): vertical scroll adjustment, or NULL.
//
// The function returns the following values:
//
//    - layout: new Layout.
//
func NewLayout(hadjustment, vadjustment *Adjustment) *Layout {
	var _args [2]girepository.Argument

	if hadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}
	if vadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "Layout")
	_gret := _info.InvokeClassMethod("new_Layout", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(hadjustment)
	runtime.KeepAlive(vadjustment)

	var _layout *Layout // out

	_layout = wrapLayout(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _layout
}

// BinWindow: retrieve the bin window of the layout used for drawing operations.
//
// The function returns the following values:
//
//    - window: Window.
//
func (layout *Layout) BinWindow() gdk.Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_info := girepository.MustFind("Gtk", "Layout")
	_gret := _info.InvokeClassMethod("get_bin_window", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// HAdjustment: this function should only be called after the layout has been
// placed in a ScrolledWindow or otherwise configured for scrolling. It returns
// the Adjustment used for communication between the horizontal scrollbar and
// layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
//
// The function returns the following values:
//
//    - adjustment: horizontal scroll adjustment.
//
func (layout *Layout) HAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_info := girepository.MustFind("Gtk", "Layout")
	_gret := _info.InvokeClassMethod("get_hadjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// Size gets the size that has been set on the layout, and that determines the
// total extents of the layout’s scrollbar area. See gtk_layout_set_size ().
//
// The function returns the following values:
//
//    - width (optional): location to store the width set on layout, or NULL.
//    - height (optional): location to store the height set on layout, or NULL.
//
func (layout *Layout) Size() (width, height uint32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_info := girepository.MustFind("Gtk", "Layout")
	_info.InvokeClassMethod("get_size", _args[:], _outs[:])

	runtime.KeepAlive(layout)

	var _width uint32  // out
	var _height uint32 // out

	_width = uint32(*(*C.guint)(unsafe.Pointer(&_outs[0])))
	_height = uint32(*(*C.guint)(unsafe.Pointer(&_outs[1])))

	return _width, _height
}

// VAdjustment: this function should only be called after the layout has been
// placed in a ScrolledWindow or otherwise configured for scrolling. It returns
// the Adjustment used for communication between the vertical scrollbar and
// layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
//
// The function returns the following values:
//
//    - adjustment: vertical scroll adjustment.
//
func (layout *Layout) VAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_info := girepository.MustFind("Gtk", "Layout")
	_gret := _info.InvokeClassMethod("get_vadjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// Move moves a current child of layout to a new position.
//
// The function takes the following parameters:
//
//    - childWidget: current child of layout.
//    - x: x position to move to.
//    - y: y position to move to.
//
func (layout *Layout) Move(childWidget Widgetter, x, y int32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(childWidget).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(x)
	*(*C.gint)(unsafe.Pointer(&_args[3])) = C.gint(y)

	_info := girepository.MustFind("Gtk", "Layout")
	_info.InvokeClassMethod("move", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(childWidget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// Put adds child_widget to layout, at position (x,y). layout becomes the new
// parent container of child_widget.
//
// The function takes the following parameters:
//
//    - childWidget: child widget.
//    - x: x position of child widget.
//    - y: y position of child widget.
//
func (layout *Layout) Put(childWidget Widgetter, x, y int32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(childWidget).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(x)
	*(*C.gint)(unsafe.Pointer(&_args[3])) = C.gint(y)

	_info := girepository.MustFind("Gtk", "Layout")
	_info.InvokeClassMethod("put", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(childWidget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// SetHAdjustment sets the horizontal scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_set_hadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): new scroll adjustment.
//
func (layout *Layout) SetHAdjustment(adjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	if adjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "Layout")
	_info.InvokeClassMethod("set_hadjustment", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(adjustment)
}

// SetSize sets the size of the scrollable area of the layout.
//
// The function takes the following parameters:
//
//    - width of entire scrollable area.
//    - height of entire scrollable area.
//
func (layout *Layout) SetSize(width, height uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(width)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(height)

	_info := girepository.MustFind("Gtk", "Layout")
	_info.InvokeClassMethod("set_size", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetVAdjustment sets the vertical scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_set_vadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): new scroll adjustment.
//
func (layout *Layout) SetVAdjustment(adjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	if adjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "Layout")
	_info.InvokeClassMethod("set_vadjustment", _args[:], nil)

	runtime.KeepAlive(layout)
	runtime.KeepAlive(adjustment)
}
