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

// GTypeViewport returns the GType for the type Viewport.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeViewport() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Viewport").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalViewport)
	return gtype
}

// ViewportOverrider contains methods that are overridable.
type ViewportOverrider interface {
}

// Viewport widget acts as an adaptor class, implementing scrollability for
// child widgets that lack their own scrolling capabilities. Use GtkViewport to
// scroll child widgets such as Grid, Box, and so on.
//
// If a widget has native scrolling abilities, such as TextView, TreeView or
// IconView, it can be added to a ScrolledWindow with gtk_container_add(). If a
// widget does not, you must first add the widget to a Viewport, then add the
// viewport to the scrolled window. gtk_container_add() does this automatically
// if a child that does not implement Scrollable is added to a ScrolledWindow,
// so you can ignore the presence of the viewport.
//
// The GtkViewport will start scrolling content only if allocated less than the
// child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// GtkViewport has a single CSS node with name viewport.
type Viewport struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	Scrollable
}

var (
	_ Binner            = (*Viewport)(nil)
	_ coreglib.Objector = (*Viewport)(nil)
)

func classInitViewporter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapViewport(obj *coreglib.Object) *Viewport {
	return &Viewport{
		Bin: Bin{
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
		},
		Object: obj,
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	return wrapViewport(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewViewport creates a new Viewport with the given adjustments, or with
// default adjustments if none are given.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): horizontal adjustment.
//    - vadjustment (optional): vertical adjustment.
//
// The function returns the following values:
//
//    - viewport: new Viewport.
//
func NewViewport(hadjustment, vadjustment *Adjustment) *Viewport {
	var _args [2]girepository.Argument

	if hadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}
	if vadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "Viewport")
	_gret := _info.InvokeClassMethod("new_Viewport", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(hadjustment)
	runtime.KeepAlive(vadjustment)

	var _viewport *Viewport // out

	_viewport = wrapViewport(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _viewport
}

// BinWindow gets the bin window of the Viewport.
//
// The function returns the following values:
//
//    - window: Window.
//
func (viewport *Viewport) BinWindow() gdk.Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_info := girepository.MustFind("Gtk", "Viewport")
	_gret := _info.InvokeClassMethod("get_bin_window", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(viewport)

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

// HAdjustment returns the horizontal adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
//
// The function returns the following values:
//
//    - adjustment: horizontal adjustment of viewport.
//
func (viewport *Viewport) HAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_info := girepository.MustFind("Gtk", "Viewport")
	_gret := _info.InvokeClassMethod("get_hadjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(viewport)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// VAdjustment returns the vertical adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
//
// The function returns the following values:
//
//    - adjustment: vertical adjustment of viewport.
//
func (viewport *Viewport) VAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_info := girepository.MustFind("Gtk", "Viewport")
	_gret := _info.InvokeClassMethod("get_vadjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(viewport)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// ViewWindow gets the view window of the Viewport.
//
// The function returns the following values:
//
//    - window: Window.
//
func (viewport *Viewport) ViewWindow() gdk.Windower {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))

	_info := girepository.MustFind("Gtk", "Viewport")
	_gret := _info.InvokeClassMethod("get_view_window", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(viewport)

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

// SetHAdjustment sets the horizontal adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_set_hadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): Adjustment.
//
func (viewport *Viewport) SetHAdjustment(adjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))
	if adjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "Viewport")
	_info.InvokeClassMethod("set_hadjustment", _args[:], nil)

	runtime.KeepAlive(viewport)
	runtime.KeepAlive(adjustment)
}

// SetVAdjustment sets the vertical adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_set_vadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): Adjustment.
//
func (viewport *Viewport) SetVAdjustment(adjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(viewport).Native()))
	if adjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "Viewport")
	_info.InvokeClassMethod("set_vadjustment", _args[:], nil)

	runtime.KeepAlive(viewport)
	runtime.KeepAlive(adjustment)
}
