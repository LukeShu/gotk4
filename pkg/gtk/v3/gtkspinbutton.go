// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_SpinButton_ConnectOutput(gpointer, guintptr);
// extern gint _gotk4_gtk3_SpinButtonClass_input(GtkSpinButton*, gdouble*);
// extern gint _gotk4_gtk3_SpinButtonClass_output(GtkSpinButton*);
// extern void _gotk4_gtk3_SpinButtonClass_value_changed(GtkSpinButton*);
// extern void _gotk4_gtk3_SpinButtonClass_wrapped(GtkSpinButton*);
// extern void _gotk4_gtk3_SpinButton_ConnectValueChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_SpinButton_ConnectWrapped(gpointer, guintptr);
import "C"

// glib.Type values for gtkspinbutton.go.
var (
	GTypeSpinButtonUpdatePolicy = coreglib.Type(C.gtk_spin_button_update_policy_get_type())
	GTypeSpinType               = coreglib.Type(C.gtk_spin_type_get_type())
	GTypeSpinButton             = coreglib.Type(C.gtk_spin_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSpinButtonUpdatePolicy, F: marshalSpinButtonUpdatePolicy},
		{T: GTypeSpinType, F: marshalSpinType},
		{T: GTypeSpinButton, F: marshalSpinButton},
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

// SpinButtonOverrider contains methods that are overridable.
type SpinButtonOverrider interface {
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Input(newValue *float64) int32
	// The function returns the following values:
	//
	Output() int32
	ValueChanged()
	Wrapped()
}

// SpinButton is an ideal way to allow the user to set the value of some
// attribute. Rather than having to directly type a number into a Entry,
// GtkSpinButton allows the user to click on one of two arrows to increment or
// decrement the displayed value. A value can still be typed in, with the bonus
// that it can be checked to ensure it is in a given range.
//
// The main properties of a GtkSpinButton are through an adjustment. See the
// Adjustment section for more details about an adjustment's properties. Note
// that GtkSpinButton will by default make its entry large enough to accomodate
// the lower and upper bounds of the adjustment, which can lead to surprising
// results. Best practice is to set both the Entry:width-chars and
// Entry:max-width-chars poperties to the desired number of characters to
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

func classInitSpinButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkSpinButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkSpinButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Input(newValue *float64) int32 }); ok {
		pclass.input = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_input)
	}

	if _, ok := goval.(interface{ Output() int32 }); ok {
		pclass.output = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_output)
	}

	if _, ok := goval.(interface{ ValueChanged() }); ok {
		pclass.value_changed = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_value_changed)
	}

	if _, ok := goval.(interface{ Wrapped() }); ok {
		pclass.wrapped = (*[0]byte)(C._gotk4_gtk3_SpinButtonClass_wrapped)
	}
}

//export _gotk4_gtk3_SpinButtonClass_input
func _gotk4_gtk3_SpinButtonClass_input(arg0 *C.GtkSpinButton, arg1 *C.gdouble) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Input(newValue *float64) int32 })

	var _newValue *float64 // out

	_newValue = (*float64)(unsafe.Pointer(arg1))

	gint := iface.Input(_newValue)

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtk3_SpinButtonClass_output
func _gotk4_gtk3_SpinButtonClass_output(arg0 *C.GtkSpinButton) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Output() int32 })

	gint := iface.Output()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtk3_SpinButtonClass_value_changed
func _gotk4_gtk3_SpinButtonClass_value_changed(arg0 *C.GtkSpinButton) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ValueChanged() })

	iface.ValueChanged()
}

//export _gotk4_gtk3_SpinButtonClass_wrapped
func _gotk4_gtk3_SpinButtonClass_wrapped(arg0 *C.GtkSpinButton) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Wrapped() })

	iface.Wrapped()
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

//export _gotk4_gtk3_SpinButton_ConnectOutput
func _gotk4_gtk3_SpinButton_ConnectOutput(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
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

//export _gotk4_gtk3_SpinButton_ConnectValueChanged
func _gotk4_gtk3_SpinButton_ConnectValueChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectValueChanged signal is emitted when the value represented by
// spinbutton changes. Also see the SpinButton::output signal.
func (spinButton *SpinButton) ConnectValueChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "value-changed", false, unsafe.Pointer(C._gotk4_gtk3_SpinButton_ConnectValueChanged), f)
}

//export _gotk4_gtk3_SpinButton_ConnectWrapped
func _gotk4_gtk3_SpinButton_ConnectWrapped(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectWrapped signal is emitted right after the spinbutton wraps from its
// maximum to minimum value or vice-versa.
func (spinButton *SpinButton) ConnectWrapped(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(spinButton, "wrapped", false, unsafe.Pointer(C._gotk4_gtk3_SpinButton_ConnectWrapped), f)
}

// NewSpinButton creates a new SpinButton.
//
// The function takes the following parameters:
//
//    - adjustment (optional) object that this spin button should use, or NULL.
//    - climbRate specifies by how much the rate of change in the value will
//      accelerate if you continue to hold down an up/down button or arrow key.
//    - digits: number of decimal places to display.
//
// The function returns the following values:
//
//    - spinButton: new spin button as a Widget.
//
func NewSpinButton(adjustment *Adjustment, climbRate float64, digits uint32) *SpinButton {
	var args [3]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.gdouble // out
	var _arg2 C.guint   // out
	var _cret *C.void   // in

	if adjustment != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}
	_arg1 = C.gdouble(climbRate)
	_arg2 = C.guint(digits)
	*(**Adjustment)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1
	*(*uint32)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("new_SpinButton", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)

	var _spinButton *SpinButton // out

	_spinButton = wrapSpinButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _spinButton
}

// NewSpinButtonWithRange: this is a convenience constructor that allows
// creation of a numeric SpinButton without manually creating an adjustment. The
// value is initially set to the minimum value and a page increment of 10 * step
// is the default. The precision of the spin button is equivalent to the
// precision of step.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_spin_button_set_digits() to correct it.
//
// The function takes the following parameters:
//
//    - min: minimum allowable value.
//    - max: maximum allowable value.
//    - step: increment added or subtracted by spinning the widget.
//
// The function returns the following values:
//
//    - spinButton: new spin button as a Widget.
//
func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	var args [3]girepository.Argument
	var _arg0 C.gdouble // out
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out
	var _cret *C.void   // in

	_arg0 = C.gdouble(min)
	_arg1 = C.gdouble(max)
	_arg2 = C.gdouble(step)
	*(*float64)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("new_SpinButton_with_range", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
//    - adjustment (optional) to replace the spin button’s existing adjustment,
//      or NULL to leave its current adjustment unchanged.
//    - climbRate: new climb rate.
//    - digits: number of decimal places to display in the spin button.
//
func (spinButton *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint32) {
	var args [4]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 *C.void   // out
	var _arg2 C.gdouble // out
	var _arg3 C.guint   // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if adjustment != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}
	_arg2 = C.gdouble(climbRate)
	_arg3 = C.guint(digits)
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1
	*(**Adjustment)(unsafe.Pointer(&args[2])) = _arg2
	*(*float64)(unsafe.Pointer(&args[3])) = _arg3

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("configure", args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)
}

// Adjustment: get the adjustment associated with a SpinButton.
//
// The function returns the following values:
//
//    - adjustment of spin_button.
//
func (spinButton *SpinButton) Adjustment() *Adjustment {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("get_adjustment", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
//    - guint: current precision.
//
func (spinButton *SpinButton) Digits() uint32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("get_digits", args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// Numeric returns whether non-numeric text can be typed into the spin button.
// See gtk_spin_button_set_numeric().
//
// The function returns the following values:
//
//    - ok: TRUE if only numeric text can be entered.
//
func (spinButton *SpinButton) Numeric() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("get_numeric", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SnapToTicks returns whether the values are corrected to the nearest step. See
// gtk_spin_button_set_snap_to_ticks().
//
// The function returns the following values:
//
//    - ok: TRUE if values are snapped to the nearest step.
//
func (spinButton *SpinButton) SnapToTicks() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("get_snap_to_ticks", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
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
	var args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret C.gdouble // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("get_value", args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ValueAsInt: get the value spin_button represented as an integer.
//
// The function returns the following values:
//
//    - gint: value of spin_button.
//
func (spinButton *SpinButton) ValueAsInt() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("get_value_as_int", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(spinButton)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Wrap returns whether the spin button’s value wraps around to the opposite
// limit when the upper or lower limit of the range is exceeded. See
// gtk_spin_button_set_wrap().
//
// The function returns the following values:
//
//    - ok: TRUE if the spin button wraps around.
//
func (spinButton *SpinButton) Wrap() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "SpinButton").InvokeMethod("get_wrap", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
//    - adjustment to replace the existing adjustment.
//
func (spinButton *SpinButton) SetAdjustment(adjustment *Adjustment) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_adjustment", args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
}

// SetDigits: set the precision to be displayed by spin_button. Up to 20 digit
// precision is allowed.
//
// The function takes the following parameters:
//
//    - digits: number of digits after the decimal point to be displayed for the
//      spin button’s value.
//
func (spinButton *SpinButton) SetDigits(digits uint32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.guint(digits)
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_digits", args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(digits)
}

// SetIncrements sets the step and page increments for spin_button. This affects
// how quickly the value changes when the spin button’s arrows are activated.
//
// The function takes the following parameters:
//
//    - step: increment applied for a button 1 press.
//    - page: increment applied for a button 2 press.
//
func (spinButton *SpinButton) SetIncrements(step, page float64) {
	var args [3]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.gdouble(step)
	_arg2 = C.gdouble(page)
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_increments", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if numeric {
		_arg1 = C.TRUE
	}
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_numeric", args[:], nil)

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
	var args [3]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_range", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if snapToTicks {
		_arg1 = C.TRUE
	}
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_snap_to_ticks", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.gdouble // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	_arg1 = C.gdouble(value)
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_value", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	if wrap {
		_arg1 = C.TRUE
	}
	*(**SpinButton)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("set_wrap", args[:], nil)

	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(wrap)
}

// Update: manually force an update of the spin button.
func (spinButton *SpinButton) Update() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(spinButton).Native()))
	*(**SpinButton)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "SpinButton").InvokeMethod("update", args[:], nil)

	runtime.KeepAlive(spinButton)
}
