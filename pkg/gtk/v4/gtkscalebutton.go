// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_ScaleButtonClass_value_changed(GtkScaleButton*, double);
// extern void _gotk4_gtk4_ScaleButton_ConnectPopdown(gpointer, guintptr);
// extern void _gotk4_gtk4_ScaleButton_ConnectPopup(gpointer, guintptr);
// extern void _gotk4_gtk4_ScaleButton_ConnectValueChanged(gpointer, gdouble, guintptr);
import "C"

// GTypeScaleButton returns the GType for the type ScaleButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeScaleButton() coreglib.Type {
	gtype := coreglib.Type(C.gtk_scale_button_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalScaleButton)
	return gtype
}

// ScaleButtonOverrider contains methods that are overridable.
type ScaleButtonOverrider interface {
	// The function takes the following parameters:
	//
	ValueChanged(value float64)
}

// ScaleButton: GtkScaleButton provides a button which pops up a scale widget.
//
// This kind of widget is commonly used for volume controls in multimedia
// applications, and GTK provides a gtk.VolumeButton subclass that is tailored
// for this use case.
//
//
// CSS nodes
//
// GtkScaleButton has a single CSS node with name button. To differentiate it
// from a plain GtkButton, it gets the .scale style class.
type ScaleButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*ScaleButton)(nil)
	_ coreglib.Objector = (*ScaleButton)(nil)
)

func classInitScaleButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkScaleButtonClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ ValueChanged(value float64) }); ok {
		pclass.value_changed = (*[0]byte)(C._gotk4_gtk4_ScaleButtonClass_value_changed)
	}
}

//export _gotk4_gtk4_ScaleButtonClass_value_changed
func _gotk4_gtk4_ScaleButtonClass_value_changed(arg0 *C.GtkScaleButton, arg1 C.double) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ValueChanged(value float64) })

	var _value float64 // out

	_value = float64(arg1)

	iface.ValueChanged(_value)
}

func wrapScaleButton(obj *coreglib.Object) *ScaleButton {
	return &ScaleButton{
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
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalScaleButton(p uintptr) (interface{}, error) {
	return wrapScaleButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ScaleButton_ConnectPopdown
func _gotk4_gtk4_ScaleButton_ConnectPopdown(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectPopdown is emitted to dismiss the popup.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is <kbd>Escape</kbd>.
func (button *ScaleButton) ConnectPopdown(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popdown", false, unsafe.Pointer(C._gotk4_gtk4_ScaleButton_ConnectPopdown), f)
}

//export _gotk4_gtk4_ScaleButton_ConnectPopup
func _gotk4_gtk4_ScaleButton_ConnectPopup(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectPopup is emitted to popup the scale widget.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are <kbd>Space</kbd>, <kbd>Enter</kbd>
// and <kbd>Return</kbd>.
func (button *ScaleButton) ConnectPopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popup", false, unsafe.Pointer(C._gotk4_gtk4_ScaleButton_ConnectPopup), f)
}

//export _gotk4_gtk4_ScaleButton_ConnectValueChanged
func _gotk4_gtk4_ScaleButton_ConnectValueChanged(arg0 C.gpointer, arg1 C.gdouble, arg2 C.guintptr) {
	var f func(value float64)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value float64))
	}

	var _value float64 // out

	_value = float64(arg1)

	f(_value)
}

// ConnectValueChanged is emitted when the value field has changed.
func (button *ScaleButton) ConnectValueChanged(f func(value float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "value-changed", false, unsafe.Pointer(C._gotk4_gtk4_ScaleButton_ConnectValueChanged), f)
}

// NewScaleButton creates a GtkScaleButton.
//
// The new scale button has a range between min and max, with a stepping of
// step.
//
// The function takes the following parameters:
//
//    - min: minimum value of the scale (usually 0).
//    - max: maximum value of the scale (usually 100).
//    - step: stepping of value when a scroll-wheel event, or up/down arrow event
//      occurs (usually 2).
//    - icons (optional): NULL-terminated array of icon names, or NULL if you
//      want to set the list later with gtk_scale_button_set_icons().
//
// The function returns the following values:
//
//    - scaleButton: new GtkScaleButton.
//
func NewScaleButton(min, max, step float64, icons []string) *ScaleButton {
	var _arg1 C.double     // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out
	var _arg4 **C.char     // out
	var _cret *C.GtkWidget // in

	_arg1 = C.double(min)
	_arg2 = C.double(max)
	_arg3 = C.double(step)
	{
		_arg4 = (**C.char)(C.calloc(C.size_t((len(icons) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg4))
		{
			out := unsafe.Slice(_arg4, len(icons)+1)
			var zero *C.char
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(icons[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_scale_button_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)
	runtime.KeepAlive(icons)

	var _scaleButton *ScaleButton // out

	_scaleButton = wrapScaleButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _scaleButton
}

// Adjustment gets the GtkAdjustment associated with the GtkScaleButton’s scale.
//
// See gtk.Range.GetAdjustment() for details.
//
// The function returns the following values:
//
//    - adjustment associated with the scale.
//
func (button *ScaleButton) Adjustment() *Adjustment {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_adjustment(_arg0)
	runtime.KeepAlive(button)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// MinusButton retrieves the minus button of the GtkScaleButton.
//
// The function returns the following values:
//
//    - ret minus button of the GtkScaleButton.
//
func (button *ScaleButton) MinusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_minus_button(_arg0)
	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// PlusButton retrieves the plus button of the GtkScaleButton.
//
// The function returns the following values:
//
//    - ret plus button of the GtkScaleButton.
//
func (button *ScaleButton) PlusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_plus_button(_arg0)
	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// Popup retrieves the popup of the GtkScaleButton.
//
// The function returns the following values:
//
//    - widget: popup of the GtkScaleButton.
//
func (button *ScaleButton) Popup() Widgetter {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_popup(_arg0)
	runtime.KeepAlive(button)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Value gets the current value of the scale button.
//
// The function returns the following values:
//
//    - gdouble: current value of the scale button.
//
func (button *ScaleButton) Value() float64 {
	var _arg0 *C.GtkScaleButton // out
	var _cret C.double          // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_value(_arg0)
	runtime.KeepAlive(button)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the GtkAdjustment to be used as a model for the
// GtkScaleButton’s scale.
//
// See gtk.Range.SetAdjustment() for details.
//
// The function takes the following parameters:
//
//    - adjustment: GtkAdjustment.
//
func (button *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 *C.GtkAdjustment  // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	C.gtk_scale_button_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(adjustment)
}

// SetIcons sets the icons to be used by the scale button.
//
// The function takes the following parameters:
//
//    - icons: NULL-terminated array of icon names.
//
func (button *ScaleButton) SetIcons(icons []string) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 **C.char          // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	{
		_arg1 = (**C.char)(C.calloc(C.size_t((len(icons) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(icons)+1)
			var zero *C.char
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(icons[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_scale_button_set_icons(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(icons)
}

// SetValue sets the current value of the scale.
//
// If the value is outside the minimum or maximum range values, it will be
// clamped to fit inside them.
//
// The scale button emits the gtk.ScaleButton::value-changed signal if the value
// changes.
//
// The function takes the following parameters:
//
//    - value: new value of the scale button.
//
func (button *ScaleButton) SetValue(value float64) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.double          // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = C.double(value)

	C.gtk_scale_button_set_value(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(value)
}
