// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_toplevel_layout_get_type()), F: marshalToplevelLayout},
	})
}

// ToplevelLayout: GdkToplevelLayout struct contains information that is
// necessary to present a sovereign window on screen.
//
// The GdkToplevelLayout struct is necessary for using gdk.Toplevel.Present().
//
// Toplevel surfaces are sovereign windows that can be presented to the user in
// various states (maximized, on all workspaces, etc).
type ToplevelLayout struct {
	nocopy gextras.NoCopy
	native *C.GdkToplevelLayout
}

func marshalToplevelLayout(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &ToplevelLayout{native: (*C.GdkToplevelLayout)(unsafe.Pointer(b))}, nil
}

// NewToplevelLayout constructs a struct ToplevelLayout.
func NewToplevelLayout() *ToplevelLayout {
	var _cret *C.GdkToplevelLayout // in

	_cret = C.gdk_toplevel_layout_new()

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_toplevelLayout, func(v *ToplevelLayout) {
		C.gdk_toplevel_layout_unref((*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _toplevelLayout
}

// Copy: create a new ToplevelLayout and copy the contents of layout into it.
func (layout *ToplevelLayout) Copy() *ToplevelLayout {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret *C.GdkToplevelLayout // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_toplevel_layout_copy(_arg0)

	runtime.KeepAlive(layout)

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_toplevelLayout, func(v *ToplevelLayout) {
		C.gdk_toplevel_layout_unref((*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _toplevelLayout
}

// Equal: check whether layout and other has identical layout properties.
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
func (layout *ToplevelLayout) FullscreenMonitor() *Monitor {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret *C.GdkMonitor        // in

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_toplevel_layout_get_fullscreen_monitor(_arg0)

	runtime.KeepAlive(layout)

	var _monitor *Monitor // out

	if _cret != nil {
		_monitor = wrapMonitor(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _monitor
}

// Maximized: if the layout specifies whether to the toplevel should go
// maximized, the value pointed to by maximized is set to TRUE if it should go
// fullscreen, or FALSE, if it should go unmaximized.
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
func (layout *ToplevelLayout) SetFullscreen(fullscreen bool, monitor *Monitor) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out
	var _arg2 *C.GdkMonitor        // out

	_arg0 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	if fullscreen {
		_arg1 = C.TRUE
	}
	if monitor != nil {
		_arg2 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))
	}

	C.gdk_toplevel_layout_set_fullscreen(_arg0, _arg1, _arg2)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(fullscreen)
	runtime.KeepAlive(monitor)
}

// SetMaximized sets whether the layout should cause the surface to be maximized
// when presented.
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
