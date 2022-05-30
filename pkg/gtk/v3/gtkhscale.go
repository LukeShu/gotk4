// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkhscale.go.
var GTypeHScale = coreglib.Type(C.gtk_hscale_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeHScale, F: marshalHScale},
	})
}

// HScaleOverrider contains methods that are overridable.
type HScaleOverrider interface {
}

// HScale widget is used to allow the user to select a value using a horizontal
// slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkHScale has been deprecated, use Scale instead.
type HScale struct {
	_ [0]func() // equal guard
	Scale
}

var (
	_ Ranger = (*HScale)(nil)
)

func classInitHScaler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapHScale(obj *coreglib.Object) *HScale {
	return &HScale{
		Scale: Scale{
			Range: Range{
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
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHScale(p uintptr) (interface{}, error) {
	return wrapHScale(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHScale creates a new HScale.
//
// Deprecated: Use gtk_scale_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function takes the following parameters:
//
//    - adjustment (optional) which sets the range of the scale.
//
// The function returns the following values:
//
//    - hScale: new HScale.
//
func NewHScale(adjustment *Adjustment) *HScale {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	if adjustment != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}
	*(**Adjustment)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "HScale").InvokeMethod("new_HScale", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(adjustment)

	var _hScale *HScale // out

	_hScale = wrapHScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _hScale
}

// NewHScaleWithRange creates a new horizontal scale widget that lets the user
// input a number between min and max (including min and max) with the increment
// step. step must be nonzero; it’s the distance the slider moves when using the
// arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_scale_set_digits() to correct it.
//
// Deprecated: Use gtk_scale_new_with_range() with GTK_ORIENTATION_HORIZONTAL
// instead.
//
// The function takes the following parameters:
//
//    - min: minimum value.
//    - max: maximum value.
//    - step increment (tick size) used with keyboard shortcuts.
//
// The function returns the following values:
//
//    - hScale: new HScale.
//
func NewHScaleWithRange(min, max, step float64) *HScale {
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

	_gret := girepository.MustFind("Gtk", "HScale").InvokeMethod("new_HScale_with_range", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _hScale *HScale // out

	_hScale = wrapHScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _hScale
}
