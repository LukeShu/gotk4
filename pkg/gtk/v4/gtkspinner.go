// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkspinner.go.
var GTypeSpinner = externglib.Type(C.gtk_spinner_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSpinner, F: marshalSpinner},
	})
}

// Spinner: GtkSpinner widget displays an icon-size spinning animation.
//
// It is often used as an alternative to a gtk.ProgressBar for displaying
// indefinite activity, instead of actual progress.
//
// !An example GtkSpinner (spinner.png)
//
// To start the animation, use gtk.Spinner.Start(), to stop it use
// gtk.Spinner.Stop().
//
//
// CSS nodes
//
// GtkSpinner has a single CSS node with the name spinner. When the animation is
// active, the :checked pseudoclass is added to this node.
type Spinner struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Spinner)(nil)
)

func wrapSpinner(obj *externglib.Object) *Spinner {
	return &Spinner{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalSpinner(p uintptr) (interface{}, error) {
	return wrapSpinner(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSpinner returns a new spinner widget. Not yet started.
//
// The function returns the following values:
//
//    - spinner: new GtkSpinner.
//
func NewSpinner() *Spinner {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_spinner_new()

	var _spinner *Spinner // out

	_spinner = wrapSpinner(externglib.Take(unsafe.Pointer(_cret)))

	return _spinner
}

// Spinning returns whether the spinner is spinning.
//
// The function returns the following values:
//
//    - ok: TRUE if the spinner is active.
//
func (spinner *Spinner) Spinning() bool {
	var _arg0 *C.GtkSpinner // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	_cret = C.gtk_spinner_get_spinning(_arg0)
	runtime.KeepAlive(spinner)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSpinning sets the activity of the spinner.
//
// The function takes the following parameters:
//
//    - spinning: whether the spinner should be spinning.
//
func (spinner *Spinner) SetSpinning(spinning bool) {
	var _arg0 *C.GtkSpinner // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))
	if spinning {
		_arg1 = C.TRUE
	}

	C.gtk_spinner_set_spinning(_arg0, _arg1)
	runtime.KeepAlive(spinner)
	runtime.KeepAlive(spinning)
}

// Start starts the animation of the spinner.
func (spinner *Spinner) Start() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	C.gtk_spinner_start(_arg0)
	runtime.KeepAlive(spinner)
}

// Stop stops the animation of the spinner.
func (spinner *Spinner) Stop() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	C.gtk_spinner_stop(_arg0)
	runtime.KeepAlive(spinner)
}
