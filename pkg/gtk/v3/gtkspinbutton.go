// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_SpinButton_ConnectWrapped(gpointer, guintptr);
// extern void _gotk4_gtk3_SpinButton_ConnectValueChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_SpinButton_ConnectChangeValue(gpointer, GtkScrollType, guintptr);
// extern void _gotk4_gtk3_SpinButtonClass_wrapped(GtkSpinButton*);
// extern void _gotk4_gtk3_SpinButtonClass_value_changed(GtkSpinButton*);
// extern void _gotk4_gtk3_SpinButtonClass_change_value(GtkSpinButton*, GtkScrollType);
// extern gint _gotk4_gtk3_SpinButtonClass_output(GtkSpinButton*);
// extern gint _gotk4_gtk3_SpinButtonClass_input(GtkSpinButton*, gdouble*);
// extern gboolean _gotk4_gtk3_SpinButton_ConnectOutput(gpointer, guintptr);
// gint _gotk4_gtk3_SpinButton_virtual_input(void* fnptr, GtkSpinButton* arg0, gdouble* arg1) {
//   return ((gint (*)(GtkSpinButton*, gdouble*))(fnptr))(arg0, arg1);
// };
// gint _gotk4_gtk3_SpinButton_virtual_output(void* fnptr, GtkSpinButton* arg0) {
//   return ((gint (*)(GtkSpinButton*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_SpinButton_virtual_change_value(void* fnptr, GtkSpinButton* arg0, GtkScrollType arg1) {
//   ((void (*)(GtkSpinButton*, GtkScrollType))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_SpinButton_virtual_value_changed(void* fnptr, GtkSpinButton* arg0) {
//   ((void (*)(GtkSpinButton*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_SpinButton_virtual_wrapped(void* fnptr, GtkSpinButton* arg0) {
//   ((void (*)(GtkSpinButton*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeSpinButtonUpdatePolicy = coreglib.Type(C.gtk_spin_button_update_policy_get_type())
	GTypeSpinType               = coreglib.Type(C.gtk_spin_type_get_type())
	GTypeSpinButton             = coreglib.Type(C.gtk_spin_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpinButtonUpdatePolicy, F: marshalSpinButtonUpdatePolicy},
		coreglib.TypeMarshaler{T: GTypeSpinType, F: marshalSpinType},
		coreglib.TypeMarshaler{T: GTypeSpinButton, F: marshalSpinButton},
	})
}

// INPUT_ERROR: constant to return from a signal handler for the
// SpinButton::input signal in case of conversion failure.
const INPUT_ERROR = -1

// SpinButtonUpdatePolicy: spin button update policy determines whether the spin
// button displays values even if they are outside the bounds of its adjustment.
// See gtk_spin_button_set_update_policy().
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

// SpinButtonOverrides contains methods that are overridable.
type SpinButtonOverrides struct {
	// The function takes the following parameters:
	//
	ChangeValue func(scroll ScrollType)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Input func(newValue *float64) int
	// The function returns the following values:
	//
	Output       func() int
	ValueChanged func()
	Wrapped      func()
}

func defaultSpinButtonOverrides(v *SpinButton) SpinButtonOverrides {
	return SpinButtonOverrides{
		ChangeValue:  v.changeValue,
		Input:        v.input,
		Output:       v.output,
		ValueChanged: v.valueChanged,
		Wrapped:      v.wrapped,
	}
}

// SpinButton is an ideal way to allow the user to set the value of some
// attribute. Rather than having to directly type a number into a Entry,
// GtkSpinButton allows the user to click on one of two arrows to increment or
// decrement the displayed value. A value can still be typed in, with the bonus
// that it can be checked to ensure it is in a given range.
//
// The main properties of a GtkSpinButton are through an adjustment. See the
// Adjustment section for more details about an adjustment's properties.
// Note that GtkSpinButton will by default make its entry large enough to
// accomodate the lower and upper bounds of the adjustment, which can lead
// to surprising results. Best practice is to set both the Entry:width-chars
// and Entry:max-width-chars poperties to the desired number of characters to
// display in the entry.
//
// CSS nodes
//
//    // Provides a function to retrieve a floating point value from a
//    // GtkSpinButton, and creates a high precision spin button.
//
//    gfloat
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
//      window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      gtk_container_set_border_width (GTK_CONTAINER (window), 5);
//
//      // creates the spinbutton, with three decimal places
//      button = gtk_spin_button_new (adjustment, 0.001, 3);
//      gtk_container_add (GTK_CONTAINER (window), button);
//
//      gtk_widget_show_all (window);
//    }.
type SpinButton struct {
	_ [0]func() // equal guard
	Entry

	*coreglib.Object
	Orientable
}

var (
	_ coreglib.Objector = (*SpinButton)(nil)
	_ Widgetter         = (*SpinButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SpinButton, *SpinButtonClass, SpinButtonOverrides](
		GTypeSpinButton,
		initSpinButtonClass,
		wrapSpinButton,
		defaultSpinButtonOverrides,
	)
}

func initSpinButtonClass(gclass unsafe.Pointer, overrides SpinButtonOverrides, classInitFunc func(*SpinButtonClass)) {
	pclass := (*C.GtkSpinButtonClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeSpinButton))))

	if overrides.ChangeValue != nil {
		pclass.change_value = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_change_value)
	}

	if overrides.Input != nil {
		pclass.input = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_input)
	}

	if overrides.Output != nil {
		pclass.output = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_output)
	}

	if overrides.ValueChanged != nil {
		pclass.value_changed = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_value_changed)
	}

	if overrides.Wrapped != nil {
		pclass.wrapped = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_wrapped)
	}

	if classInitFunc != nil {
		class := (*SpinButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSpinButton(obj *coreglib.Object) *SpinButton {
	return &SpinButton{
		Entry: Entry{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
			Object: obj,
			CellEditable: CellEditable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
			Editable: Editable{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalSpinButton(p uintptr) (interface{}, error) {
	return wrapSpinButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChangeValue signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user initiates a value change.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control the cursor programmatically.
//
// The default bindings for this signal are Up/Down and PageUp and/PageDown.
func (spinButton *SpinButton) ConnectChangeValue(f func(scroll ScrollType)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "change-value", false, unsafe.Pointer(C._gotk4_gtk3_SpinButton_ConnectChangeValue), f)
}

// ConnectOutput signal can be used to change to formatting of the value that is
// displayed in the spin buttons entry.
//
//    // show leading zeros
//    static gboolean
//    on_output (GtkSpinButton *spin,
//               gpointer       data)
//    {
//       GtkAdjustment *adjustment;
//       gchar *text;
//       int value;
//
//       adjustment = gtk_spin_button_get_adjustment (spin);
//       value = (int)gtk_adjustment_get_value (adjustment);
//       text = g_strdup_printf ("02d", value);
//       gtk_entry_set_text (GTK_ENTRY (spin), text);
//       g_free (text);
//
//       return TRUE;
//    }.
func (spinButton *SpinButton) ConnectOutput(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "output", false, unsafe.Pointer(C._gotk4_gtk3_SpinButton_ConnectOutput), f)
}

// ConnectValueChanged signal is emitted when the value represented by
// spinbutton changes. Also see the SpinButton::output signal.
func (spinButton *SpinButton) ConnectValueChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "value-changed", false, unsafe.Pointer(C._gotk4_gtk3_SpinButton_ConnectValueChanged), f)
}

// ConnectWrapped signal is emitted right after the spinbutton wraps from its
// maximum to minimum value or vice-versa.
func (spinButton *SpinButton) ConnectWrapped(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "wrapped", false, unsafe.Pointer(C._gotk4_gtk3_SpinButton_ConnectWrapped), f)
}

// NewSpinButton creates a new SpinButton.
//
// The function takes the following parameters:
//
//   - adjustment (optional) object that this spin button should use, or NULL.
//   - climbRate specifies by how much the rate of change in the value will
//     accelerate if you continue to hold down an up/down button or arrow key.
//   - digits: number of decimal places to display.
//
// The function returns the following values:
//
//   - spinButton: new spin button as a Widget.
//
func NewSpinButton(adjustment *Adjustment, climbRate float64, digits uint) *SpinButton {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 C.gdouble        // out
	var _arg3 C.guint          // out
	var _cret *C.GtkWidget     // in

	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}
	_arg2 = C.gdouble(climbRate)
	_arg3 = C.guint(digits)

	_cret = C.gtk_spin_button_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)

	var _spinButton *SpinButton // out

	_spinButton = wrapSpinButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _spinButton
}

// NewSpinButtonWithRange: this is a convenience constructor that allows
// creation of a numeric SpinButton without manually creating an adjustment.
// The value is initially set to the minimum value and a page increment of 10
// * step is the default. The precision of the spin button is equivalent to the
// precision of step.
//
// Note that the way in which the precision is derived works best if step is
// a power of ten. If the resulting precision is not suitable for your needs,
// use gtk_spin_button_set_digits() to correct it.
//
// The function takes the following parameters:
//
//   - min: minimum allowable value.
//   - max: maximum allowable value.
//   - step: increment added or subtracted by spinning the widget.
//
// The function returns the following values:
//
//   - spinButton: new spin button as a Widget.
//
func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _arg3 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)
	_arg3 = C.gdouble(step)

	_cret = C.gtk_spin_button_new_with_range(_arg1, _arg2, _arg3)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _spinButton *SpinButton // out

	_spinButton = wrapSpinButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _spinButton
}

// Configure changes the properties of an existing spin button. The adjustment,
// climb rate, and number of decimal places are updated accordingly.
//
// The function takes the following parameters:
//
//   - adjustment (optional) to replace the spin button’s existing adjustment,
//     or NULL to leave its current adjustment unchanged.
//   - climbRate: new climb rate.
//   - digits: number of decimal places to display in the spin button.
//
func (spinButton *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 *C.GtkAdjustment // out
	var _arg2 C.gdouble        // out
	var _arg3 C.guint          // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}
	_arg2 = C.gdouble(climbRate)
	_arg3 = C.guint(digits)

	C.gtk_spin_button_configure(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)
}

// Adjustment: get the adjustment associated with a SpinButton.
//
// The function returns the following values:
//
//   - adjustment of spin_button.
//
func (spinButton *SpinButton) Adjustment() *Adjustment {
	var _arg0 *C.GtkSpinButton // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_adjustment(_arg0)
	runtime.KeepAlive(spinButton)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// Digits fetches the precision of spin_button. See
// gtk_spin_button_set_digits().
//
// The function returns the following values:
//
//   - guint: current precision.
//
func (spinButton *SpinButton) Digits() uint {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_digits(_arg0)
	runtime.KeepAlive(spinButton)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Increments gets the current step and page the increments used by spin_button.
// See gtk_spin_button_set_increments().
//
// The function returns the following values:
//
//   - step (optional): location to store step increment, or NULL.
//   - page (optional): location to store page increment, or NULL.
//
func (spinButton *SpinButton) Increments() (step, page float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gdouble        // in
	var _arg2 C.gdouble        // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	C.gtk_spin_button_get_increments(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(spinButton)

	var _step float64 // out
	var _page float64 // out

	_step = float64(_arg1)
	_page = float64(_arg2)

	return _step, _page
}

// Numeric returns whether non-numeric text can be typed into the spin button.
// See gtk_spin_button_set_numeric().
//
// The function returns the following values:
//
//   - ok: TRUE if only numeric text can be entered.
//
func (spinButton *SpinButton) Numeric() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_numeric(_arg0)
	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Range gets the range allowed for spin_button. See
// gtk_spin_button_set_range().
//
// The function returns the following values:
//
//   - min (optional): location to store minimum allowed value, or NULL.
//   - max (optional): location to store maximum allowed value, or NULL.
//
func (spinButton *SpinButton) Range() (min, max float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gdouble        // in
	var _arg2 C.gdouble        // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	C.gtk_spin_button_get_range(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(spinButton)

	var _min float64 // out
	var _max float64 // out

	_min = float64(_arg1)
	_max = float64(_arg2)

	return _min, _max
}

// SnapToTicks returns whether the values are corrected to the nearest step.
// See gtk_spin_button_set_snap_to_ticks().
//
// The function returns the following values:
//
//   - ok: TRUE if values are snapped to the nearest step.
//
func (spinButton *SpinButton) SnapToTicks() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_snap_to_ticks(_arg0)
	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UpdatePolicy gets the update behavior of a spin button. See
// gtk_spin_button_set_update_policy().
//
// The function returns the following values:
//
//   - spinButtonUpdatePolicy: current update policy.
//
func (spinButton *SpinButton) UpdatePolicy() SpinButtonUpdatePolicy {
	var _arg0 *C.GtkSpinButton            // out
	var _cret C.GtkSpinButtonUpdatePolicy // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_update_policy(_arg0)
	runtime.KeepAlive(spinButton)

	var _spinButtonUpdatePolicy SpinButtonUpdatePolicy // out

	_spinButtonUpdatePolicy = SpinButtonUpdatePolicy(_cret)

	return _spinButtonUpdatePolicy
}

// Value: get the value in the spin_button.
//
// The function returns the following values:
//
//   - gdouble: value of spin_button.
//
func (spinButton *SpinButton) Value() float64 {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_value(_arg0)
	runtime.KeepAlive(spinButton)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ValueAsInt: get the value spin_button represented as an integer.
//
// The function returns the following values:
//
//   - gint: value of spin_button.
//
func (spinButton *SpinButton) ValueAsInt() int {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_value_as_int(_arg0)
	runtime.KeepAlive(spinButton)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Wrap returns whether the spin button’s value wraps around to the
// opposite limit when the upper or lower limit of the range is exceeded.
// See gtk_spin_button_set_wrap().
//
// The function returns the following values:
//
//   - ok: TRUE if the spin button wraps around.
//
func (spinButton *SpinButton) Wrap() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C.gtk_spin_button_get_wrap(_arg0)
	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAdjustment replaces the Adjustment associated with spin_button.
//
// The function takes the following parameters:
//
//   - adjustment to replace the existing adjustment.
//
func (spinButton *SpinButton) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	C.gtk_spin_button_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
}

// SetDigits: set the precision to be displayed by spin_button. Up to 20 digit
// precision is allowed.
//
// The function takes the following parameters:
//
//   - digits: number of digits after the decimal point to be displayed for the
//     spin button’s value.
//
func (spinButton *SpinButton) SetDigits(digits uint) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.guint(digits)

	C.gtk_spin_button_set_digits(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(digits)
}

// SetIncrements sets the step and page increments for spin_button. This affects
// how quickly the value changes when the spin button’s arrows are activated.
//
// The function takes the following parameters:
//
//   - step: increment applied for a button 1 press.
//   - page: increment applied for a button 2 press.
//
func (spinButton *SpinButton) SetIncrements(step, page float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.gdouble(step)
	_arg2 = C.gdouble(page)

	C.gtk_spin_button_set_increments(_arg0, _arg1, _arg2)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(step)
	runtime.KeepAlive(page)
}

// SetNumeric sets the flag that determines if non-numeric text can be typed
// into the spin button.
//
// The function takes the following parameters:
//
//   - numeric: flag indicating if only numeric entry is allowed.
//
func (spinButton *SpinButton) SetNumeric(numeric bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if numeric {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_numeric(_arg0, _arg1)
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
//   - min: minimum allowable value.
//   - max: maximum allowable value.
//
func (spinButton *SpinButton) SetRange(min, max float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)

	C.gtk_spin_button_set_range(_arg0, _arg1, _arg2)
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
//   - snapToTicks: flag indicating if invalid values should be corrected.
//
func (spinButton *SpinButton) SetSnapToTicks(snapToTicks bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if snapToTicks {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_snap_to_ticks(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(snapToTicks)
}

// SetUpdatePolicy sets the update behavior of a spin button. This determines
// whether the spin button is always updated or only when a valid value is set.
//
// The function takes the following parameters:
//
//   - policy: SpinButtonUpdatePolicy value.
//
func (spinButton *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	var _arg0 *C.GtkSpinButton            // out
	var _arg1 C.GtkSpinButtonUpdatePolicy // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.GtkSpinButtonUpdatePolicy(policy)

	C.gtk_spin_button_set_update_policy(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(policy)
}

// SetValue sets the value of spin_button.
//
// The function takes the following parameters:
//
//   - value: new value.
//
func (spinButton *SpinButton) SetValue(value float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.gdouble(value)

	C.gtk_spin_button_set_value(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(value)
}

// SetWrap sets the flag that determines if a spin button value wraps around to
// the opposite limit when the upper or lower limit of the range is exceeded.
//
// The function takes the following parameters:
//
//   - wrap: flag indicating if wrapping behavior is performed.
//
func (spinButton *SpinButton) SetWrap(wrap bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if wrap {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_wrap(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(wrap)
}

// Spin: increment or decrement a spin button’s value in a specified direction
// by a specified amount.
//
// The function takes the following parameters:
//
//   - direction indicating the direction to spin.
//   - increment: step increment to apply in the specified direction.
//
func (spinButton *SpinButton) Spin(direction SpinType, increment float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.GtkSpinType    // out
	var _arg2 C.gdouble        // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.GtkSpinType(direction)
	_arg2 = C.gdouble(increment)

	C.gtk_spin_button_spin(_arg0, _arg1, _arg2)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(direction)
	runtime.KeepAlive(increment)
}

// Update: manually force an update of the spin button.
func (spinButton *SpinButton) Update() {
	var _arg0 *C.GtkSpinButton // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	C.gtk_spin_button_update(_arg0)
	runtime.KeepAlive(spinButton)
}

// The function takes the following parameters:
//
func (spinButton *SpinButton) changeValue(scroll ScrollType) {
	gclass := (*C.GtkSpinButtonClass)(coreglib.PeekParentClass(spinButton))
	fnarg := gclass.change_value

	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.GtkScrollType  // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.GtkScrollType(scroll)

	C._gotk4_gtk3_SpinButton_virtual_change_value(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(scroll)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (spinButton *SpinButton) input(newValue *float64) int {
	gclass := (*C.GtkSpinButtonClass)(coreglib.PeekParentClass(spinButton))
	fnarg := gclass.input

	var _arg0 *C.GtkSpinButton // out
	var _arg1 *C.gdouble       // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = (*C.gdouble)(unsafe.Pointer(newValue))

	_cret = C._gotk4_gtk3_SpinButton_virtual_input(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(newValue)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// The function returns the following values:
//
func (spinButton *SpinButton) output() int {
	gclass := (*C.GtkSpinButtonClass)(coreglib.PeekParentClass(spinButton))
	fnarg := gclass.output

	var _arg0 *C.GtkSpinButton // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	_cret = C._gotk4_gtk3_SpinButton_virtual_output(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(spinButton)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (spinButton *SpinButton) valueChanged() {
	gclass := (*C.GtkSpinButtonClass)(coreglib.PeekParentClass(spinButton))
	fnarg := gclass.value_changed

	var _arg0 *C.GtkSpinButton // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	C._gotk4_gtk3_SpinButton_virtual_value_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(spinButton)
}

func (spinButton *SpinButton) wrapped() {
	gclass := (*C.GtkSpinButtonClass)(coreglib.PeekParentClass(spinButton))
	fnarg := gclass.wrapped

	var _arg0 *C.GtkSpinButton // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))

	C._gotk4_gtk3_SpinButton_virtual_wrapped(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(spinButton)
}

// SpinButtonClass: instance of this type is always passed by reference.
type SpinButtonClass struct {
	*spinButtonClass
}

// spinButtonClass is the struct that's finalized.
type spinButtonClass struct {
	native *C.GtkSpinButtonClass
}

func (s *SpinButtonClass) ParentClass() *EntryClass {
	valptr := &s.native.parent_class
	var _v *EntryClass // out
	_v = (*EntryClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
