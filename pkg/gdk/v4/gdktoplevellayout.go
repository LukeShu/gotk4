// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GTypeToplevelLayout returns the GType for the type ToplevelLayout.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeToplevelLayout() coreglib.Type {
	gtype := coreglib.Type(C.gdk_toplevel_layout_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalToplevelLayout)
	return gtype
}

// ToplevelLayout: GdkToplevelLayout struct contains information that is
// necessary to present a sovereign window on screen.
//
// The GdkToplevelLayout struct is necessary for using gdk.Toplevel.Present().
//
// Toplevel surfaces are sovereign windows that can be presented to the user in
// various states (maximized, on all workspaces, etc).
//
// An instance of this type is always passed by reference.
type ToplevelLayout struct {
	*toplevelLayout
}

// toplevelLayout is the struct that's finalized.
type toplevelLayout struct {
	native *C.GdkToplevelLayout
}

func marshalToplevelLayout(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ToplevelLayout{&toplevelLayout{(*C.GdkToplevelLayout)(b)}}, nil
}

// NewToplevelLayout constructs a struct ToplevelLayout.
func NewToplevelLayout() *ToplevelLayout {
	var _cret *C.GdkToplevelLayout // in

	_cret = C.gdk_toplevel_layout_new()

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_toplevelLayout)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _toplevelLayout
}

// Copy: create a new ToplevelLayout and copy the contents of layout into it.
//
// The function returns the following values:
//
//    - toplevelLayout: copy of layout.
//
func (layout *ToplevelLayout) Copy() *ToplevelLayout {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret *C.GdkToplevelLayout // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_toplevel_layout_copy(_arg0)
	runtime.KeepAlive(layout)

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_toplevelLayout)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _toplevelLayout
}

// Equal: check whether layout and other has identical layout properties.
//
// The function takes the following parameters:
//
//    - other ToplevelLayout.
//
// The function returns the following values:
//
//    - ok: TRUE if layout and other have identical layout properties, otherwise
//      FALSE.
//
func (layout *ToplevelLayout) Equal(other *ToplevelLayout) bool {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 *C.GdkToplevelLayout // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = C.gdk_toplevel_layout_equal(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Fullscreen: if the layout specifies whether to the toplevel should go
// fullscreen, the value pointed to by fullscreen is set to TRUE if it should go
// fullscreen, or FALSE, if it should go unfullscreen.
//
// The function returns the following values:
//
//    - fullscreen: location to store whether the toplevel should be fullscreen.
//    - ok: whether the layout specifies the fullscreen state for the toplevel.
//
func (layout *ToplevelLayout) Fullscreen() (fullscreen bool, ok bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // in
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_toplevel_layout_get_fullscreen(_arg0, &_arg1)
	runtime.KeepAlive(layout)

	var _fullscreen bool // out
	var _ok bool         // out

	if _arg1 != 0 {
		_fullscreen = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _fullscreen, _ok
}

// FullscreenMonitor returns the monitor that the layout is fullscreening the
// surface on.
//
// The function returns the following values:
//
//    - monitor (optional) on which layout fullscreens.
//
func (layout *ToplevelLayout) FullscreenMonitor() *Monitor {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret *C.GdkMonitor        // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_toplevel_layout_get_fullscreen_monitor(_arg0)
	runtime.KeepAlive(layout)

	var _monitor *Monitor // out

	if _cret != nil {
		_monitor = wrapMonitor(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _monitor
}

// Maximized: if the layout specifies whether to the toplevel should go
// maximized, the value pointed to by maximized is set to TRUE if it should go
// fullscreen, or FALSE, if it should go unmaximized.
//
// The function returns the following values:
//
//    - maximized: set to TRUE if the toplevel should be maximized.
//    - ok: whether the layout specifies the maximized state for the toplevel.
//
func (layout *ToplevelLayout) Maximized() (maximized bool, ok bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // in
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_toplevel_layout_get_maximized(_arg0, &_arg1)
	runtime.KeepAlive(layout)

	var _maximized bool // out
	var _ok bool        // out

	if _arg1 != 0 {
		_maximized = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _maximized, _ok
}

// Resizable returns whether the layout should allow the user to resize the
// surface.
//
// The function returns the following values:
//
//    - ok: TRUE if the layout is resizable.
//
func (layout *ToplevelLayout) Resizable() bool {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_toplevel_layout_get_resizable(_arg0)
	runtime.KeepAlive(layout)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFullscreen sets whether the layout should cause the surface to be
// fullscreen when presented.
//
// The function takes the following parameters:
//
//    - fullscreen: TRUE to fullscreen the surface.
//    - monitor (optional) to fullscreen on.
//
func (layout *ToplevelLayout) SetFullscreen(fullscreen bool, monitor *Monitor) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out
	var _arg2 *C.GdkMonitor        // out

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	if fullscreen {
		_arg1 = C.TRUE
	}
	if monitor != nil {
		_arg2 = (*C.GdkMonitor)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	}

	C.gdk_toplevel_layout_set_fullscreen(_arg0, _arg1, _arg2)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(fullscreen)
	runtime.KeepAlive(monitor)
}

// SetMaximized sets whether the layout should cause the surface to be maximized
// when presented.
//
// The function takes the following parameters:
//
//    - maximized: TRUE to maximize.
//
func (layout *ToplevelLayout) SetMaximized(maximized bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	if maximized {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_layout_set_maximized(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(maximized)
}

// SetResizable sets whether the layout should allow the user to resize the
// surface after it has been presented.
//
// The function takes the following parameters:
//
//    - resizable: TRUE to allow resizing.
//
func (layout *ToplevelLayout) SetResizable(resizable bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	if resizable {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_layout_set_resizable(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(resizable)
}
