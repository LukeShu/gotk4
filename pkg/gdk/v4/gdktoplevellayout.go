// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gdktoplevellayout.go.
var GTypeToplevelLayout = coreglib.Type(C.gdk_toplevel_layout_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeToplevelLayout, F: marshalToplevelLayout},
	})
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
	var _cret *C.void // in

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_toplevelLayout)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_toplevel_layout_unref((*C.GdkToplevelLayout)(intern.C))
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
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(**ToplevelLayout)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_toplevelLayout)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_toplevel_layout_unref((*C.GdkToplevelLayout)(intern.C))
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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))
	*(**ToplevelLayout)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FullscreenMonitor returns the monitor that the layout is fullscreening the
// surface on.
//
// The function returns the following values:
//
//    - monitor (optional) on which layout fullscreens.
//
func (layout *ToplevelLayout) FullscreenMonitor() *Monitor {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(**ToplevelLayout)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _monitor *Monitor // out

	if _cret != nil {
		_monitor = wrapMonitor(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _monitor
}

// Resizable returns whether the layout should allow the user to resize the
// surface.
//
// The function returns the following values:
//
//    - ok: TRUE if the layout is resizable.
//
func (layout *ToplevelLayout) Resizable() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(**ToplevelLayout)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out
	var _arg2 *C.void    // out

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	if fullscreen {
		_arg1 = C.TRUE
	}
	if monitor != nil {
		_arg2 = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	}
	*(**ToplevelLayout)(unsafe.Pointer(&args[1])) = _arg1
	*(*bool)(unsafe.Pointer(&args[2])) = _arg2

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	if maximized {
		_arg1 = C.TRUE
	}
	*(**ToplevelLayout)(unsafe.Pointer(&args[1])) = _arg1

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	if resizable {
		_arg1 = C.TRUE
	}
	*(**ToplevelLayout)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(layout)
	runtime.KeepAlive(resizable)
}
