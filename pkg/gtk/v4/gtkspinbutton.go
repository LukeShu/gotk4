// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk4_SpinButton_ConnectOutput(gpointer, guintptr);
// extern void _gotk4_gtk4_SpinButton_ConnectValueChanged(gpointer, guintptr);
// extern void _gotk4_gtk4_SpinButton_ConnectWrapped(gpointer, guintptr);
import "C"

// GTypeSpinButtonUpdatePolicy returns the GType for the type SpinButtonUpdatePolicy.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSpinButtonUpdatePolicy() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "SpinButtonUpdatePolicy").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSpinButtonUpdatePolicy)
	return gtype
}

// GTypeSpinType returns the GType for the type SpinType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSpinType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "SpinType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSpinType)
	return gtype
}

// GTypeSpinButton returns the GType for the type SpinButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSpinButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "SpinButton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSpinButton)
	return gtype
}

// INPUT_ERROR: constant to return from a signal handler for the ::input signal
// in case of conversion failure.
//
// See gtk.SpinButton::input.
const INPUT_ERROR = -1

// SpinButtonUpdatePolicy determines whether the spin button displays values
// outside the adjustment bounds.
//
// See gtk.SpinButton.SetUpdatePolicy().
type SpinButtonUpdatePolicy C.gint

const (
	// UpdateAlways: when refreshing your SpinButton, the value is always
	// displayed.
	UpdateAlways SpinButtonUpdatePolicy = iota
	// UpdateIfValid: when refreshing your SpinButton, the value is only
	// displayed if it is valid within the bounds of the spin button's
	// adjustment.
	UpdateIfValid
)

func marshalSpinButtonUpdatePolicy(p uintptr) (interface{}, error) {
	return SpinButtonUpdatePolicy(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SpinButtonUpdatePolicy.
func (s SpinButtonUpdatePolicy) String() string {
	switch s {
	case UpdateAlways:
		return "Always"
	case UpdateIfValid:
		return "IfValid"
	default:
		return fmt.Sprintf("SpinButtonUpdatePolicy(%d)", s)
	}
}

// SpinType values of the GtkSpinType enumeration are used to specify the change
// to make in gtk_spin_button_spin().
type SpinType C.gint

const (
	// SpinStepForward: increment by the adjustments step increment.
	SpinStepForward SpinType = iota
	// SpinStepBackward: decrement by the adjustments step increment.
	SpinStepBackward
	// SpinPageForward: increment by the adjustments page increment.
	SpinPageForward
	// SpinPageBackward: decrement by the adjustments page increment.
	SpinPageBackward
	// SpinHome: go to the adjustments lower bound.
	SpinHome
	// SpinEnd: go to the adjustments upper bound.
	SpinEnd
	// SpinUserDefined: change by a specified amount.
	SpinUserDefined
)

func marshalSpinType(p uintptr) (interface{}, error) {
	return SpinType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SpinType.
func (s SpinType) String() string {
	switch s {
	case SpinStepForward:
		return "StepForward"
	case SpinStepBackward:
		return "StepBackward"
	case SpinPageForward:
		return "PageForward"
	case SpinPageBackward:
		return "PageBackward"
	case SpinHome:
		return "Home"
	case SpinEnd:
		return "End"
	case SpinUserDefined:
		return "UserDefined"
	default:
		return fmt.Sprintf("SpinType(%d)", s)
	}
}

// SpinButton: GtkSpinButton is an ideal way to allow the user to set the value
// of some attribute.
//
// !An example GtkSpinButton (spinbutton.png)
//
// Rather than having to directly type a number into a GtkEntry, GtkSpinButton
// allows the user to click on one of two arrows to increment or decrement the
// displayed value. A value can still be typed in, with the bonus that it can be
// checked to ensure it is in a given range.
//
// The main properties of a GtkSpinButton are through an adjustment. See the
// gtk.Adjustment documentation for more details about an adjustment's
// properties.
//
// Note that GtkSpinButton will by default make its entry large enough to
// accommodate the lower and upper bounds of the adjustment. If this is not
// desired, the automatic sizing can be turned off by explicitly setting
// gtk.Editable:width-chars to a value != -1.
//
// Using a GtkSpinButton to get an integer
//
//    // Provides a function to retrieve an integer value from a GtkSpinButton
//    // and creates a spin button to model percentage values.
//
//    int
//    grab_int_value (GtkSpinButton *button,
//                    gpointer       user_data)
//    {
//      return gtk_spin_button_get_value_as_int (button);
//    }
//
//    void
//    create_integer_spin_button (void)
//    {
//
//      GtkWidget *window, *button;
//      GtkAdjustment *adjustment;
//
//      adjustment = gtk_adjustment_new (50.0, 0.0, 100.0, 1.0, 5.0, 0.0);
//
//      window = gtk_window_new ();
//
//      // creates the spinbutton, with no decimal places
//      button = gtk_spin_button_new (adjustment, 1.0, 0);
//      gtk_window_set_child (GTK_WINDOW (window), button);
//
//      gtk_widget_show (window);
//    }
//
//
// Using a GtkSpinButton to get a floating point value
//
//    // Provides a function to retrieve a floating point value from a
//    // GtkSpinButton, and creates a high precision spin button.
//
//    float
//    grab_float_value (GtkSpinButton *button,
//                      gpointer       user_data)
//    {
//      return gtk_spin_button_get_value (button);
//    }
//
//    void
//    create_floating_spin_button (void)
//    {
//      GtkWidget *window, *button;
//      GtkAdjustment *adjustment;
//
//      adjustment = gtk_adjustment_new (2.500, 0.0, 5.0, 0.001, 0.1, 0.0);
//
//      window = gtk_window_new ();
//
//      // creates the spinbutton, with three decimal places
//      button = gtk_spin_button_new (adjustment, 0.001, 3);
//      gtk_window_set_child (GTK_WINDOW (window), button);
//
//      gtk_widget_show (window);
//    }
//
//
// CSS nodes
//
//    spinbutton.horizontal
//    ├── text
//    │    ├── undershoot.left
//    │    ╰── undershoot.right
//    ├── button.down
//    ╰── button.up
//
//
//
//
//    spinbutton.vertical
//    ├── button.up
//    ├── text
//    │    ├── undershoot.left
//    │    ╰── undershoot.right
//    ╰── button.down
//
//
// GtkSpinButtons main CSS node has the name spinbutton. It creates subnodes for
// the entry and the two buttons, with these names. The button nodes have the
// style classes .up and .down. The GtkText subnodes (if present) are put below
// the text node. The orientation of the spin button is reflected in the
// .vertical or .horizontal style class on the main node.
//
//
// Accessiblity
//
// GtkSpinButton uses the GTK_ACCESSIBLE_ROLE_SPIN_BUTTON role.
type SpinButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	coreglib.InitiallyUnowned
	Accessible
	Buildable
	CellEditable
	ConstraintTarget
	Editable
	Orientable
}

var (
	_ Widgetter         = (*SpinButton)(nil)
	_ coreglib.Objector = (*SpinButton)(nil)
)

func wrapSpinButton(obj *coreglib.Object) *SpinButton {
	return &SpinButton{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		InitiallyUnowned: coreglib.InitiallyUnowned{
			Object: obj,
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		CellEditable: CellEditable{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalSpinButton(p uintptr) (interface{}, error) {
	return wrapSpinButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_SpinButton_ConnectOutput
func _gotk4_gtk4_SpinButton_ConnectOutput(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectOutput is emitted to tweak the formatting of the value for display.
//
//    // show leading zeros
//    static gboolean
//    on_output (GtkSpinButton *spin,
//               gpointer       data)
//    {
//       GtkAdjustment *adjustment;
//       char *text;
//       int value;
//
//       adjustment = gtk_spin_button_get_adjustment (spin);
//       value = (int)gtk_adjustment_get_value (adjustment);
//       text = g_strdup_printf ("02d", value);
//       gtk_spin_button_set_text (spin, text):
//       g_free (text);
//
//       return TRUE;
//    }.
func (spinButton *SpinButton) ConnectOutput(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "output", false, unsafe.Pointer(C._gotk4_gtk4_SpinButton_ConnectOutput), f)
}

//export _gotk4_gtk4_SpinButton_ConnectValueChanged
func _gotk4_gtk4_SpinButton_ConnectValueChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectValueChanged is emitted when the value is changed.
//
// Also see the gtk.SpinButton::output signal.
func (spinButton *SpinButton) ConnectValueChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "value-changed", false, unsafe.Pointer(C._gotk4_gtk4_SpinButton_ConnectValueChanged), f)
}

//export _gotk4_gtk4_SpinButton_ConnectWrapped
func _gotk4_gtk4_SpinButton_ConnectWrapped(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectWrapped is emitted right after the spinbutton wraps from its maximum
// to its minimum value or vice-versa.
func (spinButton *SpinButton) ConnectWrapped(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "wrapped", false, unsafe.Pointer(C._gotk4_gtk4_SpinButton_ConnectWrapped), f)
}

// NewSpinButton creates a new GtkSpinButton.
//
// The function takes the following parameters:
//
//    - adjustment (optional): GtkAdjustment that this spin button should use, or
//      NULL.
//    - climbRate specifies by how much the rate of change in the value will
//      accelerate if you continue to hold down an up/down button or arrow key.
//    - digits: number of decimal places to display.
//
// The function returns the following values:
//
//    - spinButton: new GtkSpinButton.
//
func NewSpinButton(adjustment *Adjustment, climbRate float64, digits uint32) *SpinButton {
	var _args [3]girepository.Argument

	if adjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(climbRate)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(digits)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("new_SpinButton", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)

	var _spinButton *SpinButton // out

	_spinButton = wrapSpinButton(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _spinButton
}

// NewSpinButtonWithRange creates a new GtkSpinButton with the given properties.
//
// This is a convenience constructor that allows creation of a numeric
// GtkSpinButton without manually creating an adjustment. The value is initially
// set to the minimum value and a page increment of 10 * step is the default.
// The precision of the spin button is equivalent to the precision of step.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk.SpinButton.SetDigits() to correct it.
//
// The function takes the following parameters:
//
//    - min: minimum allowable value.
//    - max: maximum allowable value.
//    - step: increment added or subtracted by spinning the widget.
//
// The function returns the following values:
//
//    - spinButton: new GtkSpinButton.
//
func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	var _args [3]girepository.Argument

	*(*C.double)(unsafe.Pointer(&_args[0])) = C.double(min)
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(max)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(step)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("new_SpinButton_with_range", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _spinButton *SpinButton // out

	_spinButton = wrapSpinButton(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _spinButton
}

// Configure changes the properties of an existing spin button.
//
// The adjustment, climb rate, and number of decimal places are updated
// accordingly.
//
// The function takes the following parameters:
//
//    - adjustment (optional): GtkAdjustment to replace the spin button’s
//      existing adjustment, or NULL to leave its current adjustment unchanged.
//    - climbRate: new climb rate.
//    - digits: number of decimal places to display in the spin button.
//
func (spinButton *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if adjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(climbRate)
	*(*C.guint)(unsafe.Pointer(&_args[3])) = C.guint(digits)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("configure", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)
}

// Adjustment: get the adjustment associated with a GtkSpinButton.
//
// The function returns the following values:
//
//    - adjustment: GtkAdjustment of spin_button.
//
func (spinButton *SpinButton) Adjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_adjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// ClimbRate returns the acceleration rate for repeated changes.
//
// The function returns the following values:
//
//    - gdouble: acceleration rate.
//
func (spinButton *SpinButton) ClimbRate() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_climb_rate", _args[:], nil)
	_cret := *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// Digits fetches the precision of spin_button.
//
// The function returns the following values:
//
//    - guint: current precision.
//
func (spinButton *SpinButton) Digits() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_digits", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Increments gets the current step and page the increments used by spin_button.
//
// See gtk.SpinButton.SetIncrements().
//
// The function returns the following values:
//
//    - step (optional): location to store step increment, or NULL.
//    - page (optional): location to store page increment, or NULL.
//
func (spinButton *SpinButton) Increments() (step, page float64) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("get_increments", _args[:], _outs[:])

	runtime.KeepAlive(spinButton)

	var _step float64 // out
	var _page float64 // out

	_step = float64(*(*C.double)(unsafe.Pointer(&_outs[0])))
	_page = float64(*(*C.double)(unsafe.Pointer(&_outs[1])))

	return _step, _page
}

// Numeric returns whether non-numeric text can be typed into the spin button.
//
// The function returns the following values:
//
//    - ok: TRUE if only numeric text can be entered.
//
func (spinButton *SpinButton) Numeric() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_numeric", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Range gets the range allowed for spin_button.
//
// See gtk.SpinButton.SetRange().
//
// The function returns the following values:
//
//    - min (optional): location to store minimum allowed value, or NULL.
//    - max (optional): location to store maximum allowed value, or NULL.
//
func (spinButton *SpinButton) Range() (min, max float64) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("get_range", _args[:], _outs[:])

	runtime.KeepAlive(spinButton)

	var _min float64 // out
	var _max float64 // out

	_min = float64(*(*C.double)(unsafe.Pointer(&_outs[0])))
	_max = float64(*(*C.double)(unsafe.Pointer(&_outs[1])))

	return _min, _max
}

// SnapToTicks returns whether the values are corrected to the nearest step.
//
// The function returns the following values:
//
//    - ok: TRUE if values are snapped to the nearest step.
//
func (spinButton *SpinButton) SnapToTicks() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_snap_to_ticks", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Value: get the value in the spin_button.
//
// The function returns the following values:
//
//    - gdouble: value of spin_button.
//
func (spinButton *SpinButton) Value() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_value", _args[:], nil)
	_cret := *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.double)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// ValueAsInt: get the value spin_button represented as an integer.
//
// The function returns the following values:
//
//    - gint: value of spin_button.
//
func (spinButton *SpinButton) ValueAsInt() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_value_as_int", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Wrap returns whether the spin button’s value wraps around to the opposite
// limit when the upper or lower limit of the range is exceeded.
//
// The function returns the following values:
//
//    - ok: TRUE if the spin button wraps around.
//
func (spinButton *SpinButton) Wrap() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_gret := _info.InvokeClassMethod("get_wrap", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetAdjustment replaces the GtkAdjustment associated with spin_button.
//
// The function takes the following parameters:
//
//    - adjustment: GtkAdjustment to replace the existing adjustment.
//
func (spinButton *SpinButton) SetAdjustment(adjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_adjustment", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
}

// SetClimbRate sets the acceleration rate for repeated changes when you hold
// down a button or key.
//
// The function takes the following parameters:
//
//    - climbRate: rate of acceleration, must be >= 0.
//
func (spinButton *SpinButton) SetClimbRate(climbRate float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(climbRate)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_climb_rate", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(climbRate)
}

// SetDigits: set the precision to be displayed by spin_button.
//
// Up to 20 digit precision is allowed.
//
// The function takes the following parameters:
//
//    - digits: number of digits after the decimal point to be displayed for the
//      spin button’s value.
//
func (spinButton *SpinButton) SetDigits(digits uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(digits)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_digits", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(digits)
}

// SetIncrements sets the step and page increments for spin_button.
//
// This affects how quickly the value changes when the spin button’s arrows are
// activated.
//
// The function takes the following parameters:
//
//    - step: increment applied for a button 1 press.
//    - page: increment applied for a button 2 press.
//
func (spinButton *SpinButton) SetIncrements(step, page float64) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(step)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(page)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_increments", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(step)
	runtime.KeepAlive(page)
}

// SetNumeric sets the flag that determines if non-numeric text can be typed
// into the spin button.
//
// The function takes the following parameters:
//
//    - numeric: flag indicating if only numeric entry is allowed.
//
func (spinButton *SpinButton) SetNumeric(numeric bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if numeric {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_numeric", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(numeric)
}

// SetRange sets the minimum and maximum allowable values for spin_button.
//
// If the current value is outside this range, it will be adjusted to fit within
// the range, otherwise it will remain unchanged.
//
// The function takes the following parameters:
//
//    - min: minimum allowable value.
//    - max: maximum allowable value.
//
func (spinButton *SpinButton) SetRange(min, max float64) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(min)
	*(*C.double)(unsafe.Pointer(&_args[2])) = C.double(max)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_range", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

// SetSnapToTicks sets the policy as to whether values are corrected to the
// nearest step increment when a spin button is activated after providing an
// invalid value.
//
// The function takes the following parameters:
//
//    - snapToTicks: flag indicating if invalid values should be corrected.
//
func (spinButton *SpinButton) SetSnapToTicks(snapToTicks bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if snapToTicks {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_snap_to_ticks", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(snapToTicks)
}

// SetValue sets the value of spin_button.
//
// The function takes the following parameters:
//
//    - value: new value.
//
func (spinButton *SpinButton) SetValue(value float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(*C.double)(unsafe.Pointer(&_args[1])) = C.double(value)

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_value", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(value)
}

// SetWrap sets the flag that determines if a spin button value wraps around to
// the opposite limit when the upper or lower limit of the range is exceeded.
//
// The function takes the following parameters:
//
//    - wrap: flag indicating if wrapping behavior is performed.
//
func (spinButton *SpinButton) SetWrap(wrap bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if wrap {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("set_wrap", _args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(wrap)
}

// Update: manually force an update of the spin button.
func (spinButton *SpinButton) Update() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_info := girepository.MustFind("Gtk", "SpinButton")
	_info.InvokeClassMethod("update", _args[:], nil)

	runtime.KeepAlive(spinButton)
}
