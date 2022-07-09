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

// GTypeToplevelLayout returns the GType for the type ToplevelLayout.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeToplevelLayout() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "ToplevelLayout").RegisteredGType())
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
	native unsafe.Pointer
}

func marshalToplevelLayout(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ToplevelLayout{&toplevelLayout{(unsafe.Pointer)(b)}}, nil
}

// NewToplevelLayout constructs a struct ToplevelLayout.
func NewToplevelLayout() *ToplevelLayout {
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)
	runtime.KeepAlive(other)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _fullscreen bool // out
	var _ok bool         // out

	if **(**C.void)(unsafe.Pointer(&_outs[0])) != 0 {
		_fullscreen = true
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _monitor *Monitor // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _maximized bool // out
	var _ok bool        // out

	if **(**C.void)(unsafe.Pointer(&_outs[0])) != 0 {
		_maximized = true
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(layout)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	if fullscreen {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}
	if monitor != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(monitor).Native()))
	}

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	if maximized {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(layout)))
	if resizable {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	runtime.KeepAlive(layout)
	runtime.KeepAlive(resizable)
}
