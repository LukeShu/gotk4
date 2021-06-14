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
		{T: externglib.Type(C.gtk_toggle_tool_button_get_type()), F: marshalToggleToolButton},
	})
}

// ToggleToolButton: a ToggleToolButton is a ToolItem that contains a toggle
// button.
//
// Use gtk_toggle_tool_button_new() to create a new GtkToggleToolButton.
//
//
// CSS nodes
//
// GtkToggleToolButton has a single CSS node with name togglebutton.
type ToggleToolButton interface {
	ToolButton
	Actionable
	Activatable
	Buildable

	// Active queries a ToggleToolButton and returns its current state. Returns
	// true if the toggle button is pressed in and false if it is raised.
	Active() bool
	// SetActive sets the status of the toggle tool button. Set to true if you
	// want the GtkToggleButton to be “pressed in”, and false to raise it. This
	// action causes the toggled signal to be emitted.
	SetActive(isActive bool)
}

// toggleToolButton implements the ToggleToolButton class.
type toggleToolButton struct {
	ToolButton
	Actionable
	Activatable
	Buildable
}

var _ ToggleToolButton = (*toggleToolButton)(nil)

// WrapToggleToolButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapToggleToolButton(obj *externglib.Object) ToggleToolButton {
	return toggleToolButton{
		ToolButton:  WrapToolButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalToggleToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleToolButton(obj), nil
}

// NewToggleToolButton constructs a class ToggleToolButton.
func NewToggleToolButton() ToggleToolButton {
	var _cret C.GtkToggleToolButton // in

	_cret = C.gtk_toggle_tool_button_new()

	var _toggleToolButton ToggleToolButton // out

	_toggleToolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToggleToolButton)

	return _toggleToolButton
}

// NewToggleToolButtonFromStock constructs a class ToggleToolButton.
func NewToggleToolButtonFromStock(stockId string) ToggleToolButton {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkToggleToolButton // in

	_cret = C.gtk_toggle_tool_button_new_from_stock(_arg1)

	var _toggleToolButton ToggleToolButton // out

	_toggleToolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToggleToolButton)

	return _toggleToolButton
}

// Active queries a ToggleToolButton and returns its current state. Returns
// true if the toggle button is pressed in and false if it is raised.
func (b toggleToolButton) Active() bool {
	var _arg0 *C.GtkToggleToolButton // out

	_arg0 = (*C.GtkToggleToolButton)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_toggle_tool_button_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle tool button. Set to true if you
// want the GtkToggleButton to be “pressed in”, and false to raise it. This
// action causes the toggled signal to be emitted.
func (b toggleToolButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleToolButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkToggleToolButton)(unsafe.Pointer(b.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_tool_button_set_active(_arg0, _arg1)
}
