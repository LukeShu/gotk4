// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_toggle_button_get_type()), F: marshalToggleButton},
	})
}

// ToggleButton: a `GtkToggleButton` is a button which remains “pressed-in” when
// clicked.
//
// Clicking again will cause the toggle button to return to its normal state.
//
// A toggle button is created by calling either [ctor@Gtk.ToggleButton.new] or
// [ctor@Gtk.ToggleButton.new_with_label]. If using the former, it is advisable
// to pack a widget, (such as a `GtkLabel` and/or a `GtkImage`), into the toggle
// button’s container. (See [class@Gtk.Button] for more information).
//
// The state of a `GtkToggleButton` can be set specifically using
// [method@Gtk.ToggleButton.set_active], and retrieved using
// [method@Gtk.ToggleButton.get_active].
//
// To simply switch the state of a toggle button, use
// [method@Gtk.ToggleButton.toggled].
//
//
// Grouping
//
// Toggle buttons can be grouped together, to form mutually exclusive groups -
// only one of the buttons can be toggled at a time, and toggling another one
// will switch the currently toggled one off.
//
// To add a `GtkToggleButton` to a group, use
// [method@Gtk.ToggleButton.set_group].
//
//
// CSS nodes
//
// `GtkToggleButton` has a single CSS node with name button. To differentiate it
// from a plain `GtkButton`, it gets the .toggle style class.
//
// Creating two `GtkToggleButton` widgets.
//
// “`c static void output_state (GtkToggleButton *source, gpointer user_data) {
// printf ("Active: d\n", gtk_toggle_button_get_active (source)); }
//
// void make_toggles (void) { GtkWidget *window, *toggle1, *toggle2; GtkWidget
// *box; const char *text;
//
//    window = gtk_window_new ();
//    box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//    text = "Hi, I’m a toggle button.";
//    toggle1 = gtk_toggle_button_new_with_label (text);
//
//    g_signal_connect (toggle1, "toggled",
//                      G_CALLBACK (output_state),
//                      NULL);
//    gtk_box_append (GTK_BOX (box), toggle1);
//
//    text = "Hi, I’m a toggle button.";
//    toggle2 = gtk_toggle_button_new_with_label (text);
//    g_signal_connect (toggle2, "toggled",
//                      G_CALLBACK (output_state),
//                      NULL);
//    gtk_box_append (GTK_BOX (box), toggle2);
//
//    gtk_window_set_child (GTK_WINDOW (window), box);
//    gtk_widget_show (window);
//
// } “`
type ToggleButton interface {
	Button
	Accessible
	Actionable
	Buildable
	ConstraintTarget

	// Active queries a `GtkToggleButton` and returns its current state.
	//
	// Returns true if the toggle button is pressed in and false if it is
	// raised.
	Active() bool
	// SetActive sets the status of the toggle button.
	//
	// Set to true if you want the `GtkToggleButton` to be “pressed in”, and
	// false to raise it.
	//
	// If the status of the button changes, this action causes the
	// [signal@GtkToggleButton::toggled] signal to be emitted.
	SetActive(isActive bool)
	// SetGroup adds @self to the group of @group.
	//
	// In a group of multiple toggle buttons, only one button can be active at a
	// time.
	//
	// Setting up groups in a cycle leads to undefined behavior.
	//
	// Note that the same effect can be achieved via the
	// [interface@Gtk.Actionable] API, by using the same action with parameter
	// type and state type 's' for all buttons in the group, and giving each
	// button its own target value.
	SetGroup(group ToggleButton)
	// Toggled emits the ::toggled signal on the `GtkToggleButton`.
	//
	// There is no good reason for an application ever to call this function.
	Toggled()
}

// toggleButton implements the ToggleButton class.
type toggleButton struct {
	Button
	Accessible
	Actionable
	Buildable
	ConstraintTarget
}

var _ ToggleButton = (*toggleButton)(nil)

// WrapToggleButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapToggleButton(obj *externglib.Object) ToggleButton {
	return toggleButton{
		Button:           WrapButton(obj),
		Accessible:       WrapAccessible(obj),
		Actionable:       WrapActionable(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleButton(obj), nil
}

// NewToggleButton constructs a class ToggleButton.
func NewToggleButton() ToggleButton {
	var _cret C.GtkToggleButton // in

	_cret = C.gtk_toggle_button_new()

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToggleButton)

	return _toggleButton
}

// NewToggleButtonWithLabel constructs a class ToggleButton.
func NewToggleButtonWithLabel(label string) ToggleButton {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkToggleButton // in

	_cret = C.gtk_toggle_button_new_with_label(_arg1)

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToggleButton)

	return _toggleButton
}

// NewToggleButtonWithMnemonic constructs a class ToggleButton.
func NewToggleButtonWithMnemonic(label string) ToggleButton {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkToggleButton // in

	_cret = C.gtk_toggle_button_new_with_mnemonic(_arg1)

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ToggleButton)

	return _toggleButton
}

// Active queries a `GtkToggleButton` and returns its current state.
//
// Returns true if the toggle button is pressed in and false if it is
// raised.
func (t toggleButton) Active() bool {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_toggle_button_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle button.
//
// Set to true if you want the `GtkToggleButton` to be “pressed in”, and
// false to raise it.
//
// If the status of the button changes, this action causes the
// [signal@GtkToggleButton::toggled] signal to be emitted.
func (t toggleButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_active(_arg0, _arg1)
}

// SetGroup adds @self to the group of @group.
//
// In a group of multiple toggle buttons, only one button can be active at a
// time.
//
// Setting up groups in a cycle leads to undefined behavior.
//
// Note that the same effect can be achieved via the
// [interface@Gtk.Actionable] API, by using the same action with parameter
// type and state type 's' for all buttons in the group, and giving each
// button its own target value.
func (t toggleButton) SetGroup(group ToggleButton) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkToggleButton)(unsafe.Pointer(group.Native()))

	C.gtk_toggle_button_set_group(_arg0, _arg1)
}

// Toggled emits the ::toggled signal on the `GtkToggleButton`.
//
// There is no good reason for an application ever to call this function.
func (t toggleButton) Toggled() {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	C.gtk_toggle_button_toggled(_arg0)
}
