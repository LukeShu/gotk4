// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeVolumeButton = coreglib.Type(C.gtk_volume_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVolumeButton, F: marshalVolumeButton},
	})
}

// VolumeButton: GtkVolumeButton is a GtkScaleButton subclass tailored for
// volume control.
//
// !An example GtkVolumeButton (volumebutton.png).
type VolumeButton struct {
	_ [0]func() // equal guard
	ScaleButton
}

var (
	_ Widgetter         = (*VolumeButton)(nil)
	_ coreglib.Objector = (*VolumeButton)(nil)
)

func wrapVolumeButton(obj *coreglib.Object) *VolumeButton {
	return &VolumeButton{
		ScaleButton: ScaleButton{
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
		},
	}
}

func marshalVolumeButton(p uintptr) (interface{}, error) {
	return wrapVolumeButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVolumeButton creates a GtkVolumeButton.
//
// The button has a range between 0.0 and 1.0, with a stepping of 0.02. Volume
// values can be obtained and modified using the functions from gtk.ScaleButton.
//
// The function returns the following values:
//
//   - volumeButton: new GtkVolumeButton.
//
func NewVolumeButton() *VolumeButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_volume_button_new()

	var _volumeButton *VolumeButton // out

	_volumeButton = wrapVolumeButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _volumeButton
}
