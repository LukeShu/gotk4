// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_switch_get_type()), F: marshalSwitch},
	})
}

// Switch is a widget that has two states: on or off. The user can control which
// state should be active by clicking the empty area, or by dragging the handle.
//
// GtkSwitch can also handle situations where the underlying state changes with
// a delay. See Switch::state-set for details.
//
// CSS nodes
//
//    switch
//    ╰── slider
//
// GtkSwitch has two css nodes, the main node with the name switch and a subnode
// named slider. Neither of them is using any style classes.
type Switch interface {
	Widget
	Actionable
	Activatable
	Buildable

	// Active gets whether the Switch is in its “on” or “off” state.
	Active() bool
	// State gets the underlying state of the Switch.
	State() bool
	// SetActive changes the state of @sw to the desired one.
	SetActive(isActive bool)
	// SetState sets the underlying state of the Switch.
	//
	// Normally, this is the same as Switch:active, unless the switch is set up
	// for delayed state changes. This function is typically called from a
	// Switch::state-set signal handler.
	//
	// See Switch::state-set for details.
	SetState(state bool)
}

// _switch implements the Switch class.
type _switch struct {
	Widget
	Actionable
	Activatable
	Buildable
}

var _ Switch = (*_switch)(nil)

// WrapSwitch wraps a GObject to the right type. It is
// primarily used internally.
func WrapSwitch(obj *externglib.Object) Switch {
	return _switch{
		Widget:      WrapWidget(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalSwitch(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSwitch(obj), nil
}

// NewSwitch constructs a class Switch.
func NewSwitch() Switch {
	var _cret C.GtkSwitch // in

	_cret = C.gtk_switch_new()

	var __switch Switch // out

	__switch = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Switch)

	return __switch
}

// Active gets whether the Switch is in its “on” or “off” state.
func (s _switch) Active() bool {
	var _arg0 *C.GtkSwitch // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_switch_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// State gets the underlying state of the Switch.
func (s _switch) State() bool {
	var _arg0 *C.GtkSwitch // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_switch_get_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive changes the state of @sw to the desired one.
func (s _switch) SetActive(isActive bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_active(_arg0, _arg1)
}

// SetState sets the underlying state of the Switch.
//
// Normally, this is the same as Switch:active, unless the switch is set up
// for delayed state changes. This function is typically called from a
// Switch::state-set signal handler.
//
// See Switch::state-set for details.
func (s _switch) SetState(state bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))
	if state {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_state(_arg0, _arg1)
}
