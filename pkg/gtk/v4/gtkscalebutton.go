// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scale_button_get_type()), F: marshalScaleButtonner},
	})
}

// ScaleButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScaleButtonOverrider interface {
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
	Widget

	Orientable
	*externglib.Object
}

func wrapScaleButton(obj *externglib.Object) *ScaleButton {
	return &ScaleButton{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalScaleButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScaleButton(obj), nil
}

// NewScaleButton creates a GtkScaleButton.
//
// The new scale button has a range between min and max, with a stepping of
// step.
func NewScaleButton(min float64, max float64, step float64, icons []string) *ScaleButton {
	var _arg1 C.double     // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out
	var _arg4 **C.char     // out
	var _cret *C.GtkWidget // in

	_arg1 = C.double(min)
	_arg2 = C.double(max)
	_arg3 = C.double(step)
	{
		_arg4 = (**C.char)(C.malloc(C.ulong(len(icons)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
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

	_scaleButton = wrapScaleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _scaleButton
}

// Adjustment gets the GtkAdjustment associated with the GtkScaleButton’s scale.
//
// See gtk.Range.GetAdjustment() for details.
func (button *ScaleButton) Adjustment() *Adjustment {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_adjustment(_arg0)

	runtime.KeepAlive(button)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// MinusButton retrieves the minus button of the GtkScaleButton.
func (button *ScaleButton) MinusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_minus_button(_arg0)

	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// PlusButton retrieves the plus button of the GtkScaleButton.
func (button *ScaleButton) PlusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_plus_button(_arg0)

	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// Popup retrieves the popup of the GtkScaleButton.
func (button *ScaleButton) Popup() Widgetter {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_popup(_arg0)

	runtime.KeepAlive(button)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// Value gets the current value of the scale button.
func (button *ScaleButton) Value() float64 {
	var _arg0 *C.GtkScaleButton // out
	var _cret C.double          // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

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
func (button *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 *C.GtkAdjustment  // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_scale_button_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(adjustment)
}

// SetIcons sets the icons to be used by the scale button.
func (button *ScaleButton) SetIcons(icons []string) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 **C.char          // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	{
		_arg1 = (**C.char)(C.malloc(C.ulong(len(icons)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
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
func (button *ScaleButton) SetValue(value float64) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.double          // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.double(value)

	C.gtk_scale_button_set_value(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(value)
}
