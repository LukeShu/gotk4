// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hsv_get_type()), F: marshalHSVer},
	})
}

// HSVOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type HSVOverrider interface {
	Changed()
	Move(typ DirectionType)
}

// HSV is the “color wheel” part of a complete color selector widget. It allows
// to select a color by determining its HSV components in an intuitive way.
// Moving the selection around the outer ring changes the hue, and moving the
// selection point inside the inner triangle changes value and saturation.
//
// HSV has been deprecated together with ColorSelection, where it was used.
type HSV struct {
	Widget
}

func wrapHSV(obj *externglib.Object) *HSV {
	return &HSV{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalHSVer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHSV(obj), nil
}

// NewHSV creates a new HSV color selector.
func NewHSV() *HSV {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hsv_new()

	var _hsV *HSV // out

	_hsV = wrapHSV(externglib.Take(unsafe.Pointer(_cret)))

	return _hsV
}

// Color queries the current color in an HSV color selector. Returned values
// will be in the [0.0, 1.0] range.
func (hsv *HSV) Color() (h float64, s float64, v float64) {
	var _arg0 *C.GtkHSV // out
	var _arg1 C.gdouble // in
	var _arg2 C.gdouble // in
	var _arg3 C.gdouble // in

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(hsv.Native()))

	C.gtk_hsv_get_color(_arg0, &_arg1, &_arg2, &_arg3)

	var _h float64 // out
	var _s float64 // out
	var _v float64 // out

	_h = float64(_arg1)
	_s = float64(_arg2)
	_v = float64(_arg3)

	return _h, _s, _v
}

// Metrics queries the size and ring width of an HSV color selector.
func (hsv *HSV) Metrics() (size int, ringWidth int) {
	var _arg0 *C.GtkHSV // out
	var _arg1 C.gint    // in
	var _arg2 C.gint    // in

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(hsv.Native()))

	C.gtk_hsv_get_metrics(_arg0, &_arg1, &_arg2)

	var _size int      // out
	var _ringWidth int // out

	_size = int(_arg1)
	_ringWidth = int(_arg2)

	return _size, _ringWidth
}

// IsAdjusting: HSV color selector can be said to be adjusting if multiple rapid
// changes are being made to its value, for example, when the user is adjusting
// the value with the mouse. This function queries whether the HSV color
// selector is being adjusted or not.
func (hsv *HSV) IsAdjusting() bool {
	var _arg0 *C.GtkHSV  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(hsv.Native()))

	_cret = C.gtk_hsv_is_adjusting(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetColor sets the current color in an HSV color selector. Color component
// values must be in the [0.0, 1.0] range.
func (hsv *HSV) SetColor(h float64, s float64, v float64) {
	var _arg0 *C.GtkHSV // out
	var _arg1 C.double  // out
	var _arg2 C.double  // out
	var _arg3 C.double  // out

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(hsv.Native()))
	_arg1 = C.double(h)
	_arg2 = C.double(s)
	_arg3 = C.double(v)

	C.gtk_hsv_set_color(_arg0, _arg1, _arg2, _arg3)
}

// SetMetrics sets the size and ring width of an HSV color selector.
func (hsv *HSV) SetMetrics(size int, ringWidth int) {
	var _arg0 *C.GtkHSV // out
	var _arg1 C.gint    // out
	var _arg2 C.gint    // out

	_arg0 = (*C.GtkHSV)(unsafe.Pointer(hsv.Native()))
	_arg1 = C.gint(size)
	_arg2 = C.gint(ringWidth)

	C.gtk_hsv_set_metrics(_arg0, _arg1, _arg2)
}
