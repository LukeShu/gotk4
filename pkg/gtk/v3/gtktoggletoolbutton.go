// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_ToggleToolButtonClass_toggled(GtkToggleToolButton*);
// extern void _gotk4_gtk3_ToggleToolButton_ConnectToggled(gpointer, guintptr);
import "C"

// glib.Type values for gtktoggletoolbutton.go.
var GTypeToggleToolButton = externglib.Type(C.gtk_toggle_tool_button_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeToggleToolButton, F: marshalToggleToolButton},
	})
}

// ToggleToolButtonOverrider contains methods that are overridable.
type ToggleToolButtonOverrider interface {
	Toggled()
}

// ToggleToolButton is a ToolItem that contains a toggle button.
//
// Use gtk_toggle_tool_button_new() to create a new GtkToggleToolButton.
//
//
// CSS nodes
//
// GtkToggleToolButton has a single CSS node with name togglebutton.
type ToggleToolButton struct {
	_ [0]func() // equal guard
	ToolButton
}

var (
	_ externglib.Objector = (*ToggleToolButton)(nil)
	_ Binner              = (*ToggleToolButton)(nil)
)

func classInitToggleToolButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkToggleToolButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkToggleToolButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Toggled() }); ok {
		pclass.toggled = (*[0]byte)(C._gotk4_gtk3_ToggleToolButtonClass_toggled)
	}
}

//export _gotk4_gtk3_ToggleToolButtonClass_toggled
func _gotk4_gtk3_ToggleToolButtonClass_toggled(arg0 *C.GtkToggleToolButton) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Toggled() })

	iface.Toggled()
}

func wrapToggleToolButton(obj *externglib.Object) *ToggleToolButton {
	return &ToggleToolButton{
		ToolButton: ToolButton{
			ToolItem: ToolItem{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
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
				Activatable: Activatable{
					Object: obj,
				},
			},
			Object: obj,
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalToggleToolButton(p uintptr) (interface{}, error) {
	return wrapToggleToolButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ToggleToolButton_ConnectToggled
func _gotk4_gtk3_ToggleToolButton_ConnectToggled(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectToggled: emitted whenever the toggle tool button changes state.
func (button *ToggleToolButton) ConnectToggled(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(button, "toggled", false, unsafe.Pointer(C._gotk4_gtk3_ToggleToolButton_ConnectToggled), f)
}

// NewToggleToolButton returns a new ToggleToolButton.
//
// The function returns the following values:
//
//    - toggleToolButton: newly created ToggleToolButton.
//
func NewToggleToolButton() *ToggleToolButton {
	var _cret *C.GtkToolItem // in

	_cret = C.gtk_toggle_tool_button_new()

	var _toggleToolButton *ToggleToolButton // out

	_toggleToolButton = wrapToggleToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleToolButton
}

// NewToggleToolButtonFromStock creates a new ToggleToolButton containing the
// image and text from a stock item. Some stock ids have preprocessor macros
// like K_STOCK_OK and K_STOCK_APPLY.
//
// It is an error if stock_id is not a name of a stock item.
//
// Deprecated: Use gtk_toggle_tool_button_new() instead.
//
// The function takes the following parameters:
//
//    - stockId: name of the stock item.
//
// The function returns the following values:
//
//    - toggleToolButton: new ToggleToolButton.
//
func NewToggleToolButtonFromStock(stockId string) *ToggleToolButton {
	var _arg1 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_tool_button_new_from_stock(_arg1)
	runtime.KeepAlive(stockId)

	var _toggleToolButton *ToggleToolButton // out

	_toggleToolButton = wrapToggleToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleToolButton
}

// Active queries a ToggleToolButton and returns its current state. Returns TRUE
// if the toggle button is pressed in and FALSE if it is raised.
//
// The function returns the following values:
//
//    - ok: TRUE if the toggle tool button is pressed in, FALSE if not.
//
func (button *ToggleToolButton) Active() bool {
	var _arg0 *C.GtkToggleToolButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkToggleToolButton)(unsafe.Pointer(externglib.InternObject(button).Native()))

	_cret = C.gtk_toggle_tool_button_get_active(_arg0)
	runtime.KeepAlive(button)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle tool button. Set to TRUE if you want
// the GtkToggleButton to be “pressed in”, and FALSE to raise it. This action
// causes the toggled signal to be emitted.
//
// The function takes the following parameters:
//
//    - isActive: whether button should be active.
//
func (button *ToggleToolButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleToolButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkToggleToolButton)(unsafe.Pointer(externglib.InternObject(button).Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_tool_button_set_active(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(isActive)
}
