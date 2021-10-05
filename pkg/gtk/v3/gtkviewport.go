// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_viewport_get_type()), F: marshalViewporter},
	})
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
	Bin

	Scrollable
	*externglib.Object
}

func wrapViewport(obj *externglib.Object) *Viewport {
	return &Viewport{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
		},
		Scrollable: Scrollable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalViewporter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapViewport(obj), nil
}

// NewViewport creates a new Viewport with the given adjustments, or with
// default adjustments if none are given.
func NewViewport(hadjustment, vadjustment *Adjustment) *Viewport {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	if hadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	}
	if vadjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))
	}

	_cret = C.gtk_viewport_new(_arg1, _arg2)
	runtime.KeepAlive(hadjustment)
	runtime.KeepAlive(vadjustment)

	var _viewport *Viewport // out

	_viewport = wrapViewport(externglib.Take(unsafe.Pointer(_cret)))

	return _viewport
}

// BinWindow gets the bin window of the Viewport.
func (viewport *Viewport) BinWindow() gdk.Windower {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))

	_cret = C.gtk_viewport_get_bin_window(_arg0)
	runtime.KeepAlive(viewport)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gdk.Windower)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// HAdjustment returns the horizontal adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
func (viewport *Viewport) HAdjustment() *Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))

	_cret = C.gtk_viewport_get_hadjustment(_arg0)
	runtime.KeepAlive(viewport)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// ShadowType gets the shadow type of the Viewport. See
// gtk_viewport_set_shadow_type().
func (viewport *Viewport) ShadowType() ShadowType {
	var _arg0 *C.GtkViewport  // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))

	_cret = C.gtk_viewport_get_shadow_type(_arg0)
	runtime.KeepAlive(viewport)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

// VAdjustment returns the vertical adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
func (viewport *Viewport) VAdjustment() *Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))

	_cret = C.gtk_viewport_get_vadjustment(_arg0)
	runtime.KeepAlive(viewport)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// ViewWindow gets the view window of the Viewport.
func (viewport *Viewport) ViewWindow() gdk.Windower {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))

	_cret = C.gtk_viewport_get_view_window(_arg0)
	runtime.KeepAlive(viewport)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(gdk.Windower)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// SetHAdjustment sets the horizontal adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_set_hadjustment().
func (viewport *Viewport) SetHAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	C.gtk_viewport_set_hadjustment(_arg0, _arg1)
	runtime.KeepAlive(viewport)
	runtime.KeepAlive(adjustment)
}

// SetShadowType sets the shadow type of the viewport.
func (viewport *Viewport) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkViewport  // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_viewport_set_shadow_type(_arg0, _arg1)
	runtime.KeepAlive(viewport)
	runtime.KeepAlive(typ)
}

// SetVAdjustment sets the vertical adjustment of the viewport.
//
// Deprecated: Use gtk_scrollable_set_vadjustment().
func (viewport *Viewport) SetVAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	C.gtk_viewport_set_vadjustment(_arg0, _arg1)
	runtime.KeepAlive(viewport)
	runtime.KeepAlive(adjustment)
}
