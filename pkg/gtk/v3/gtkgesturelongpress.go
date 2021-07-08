// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_long_press_get_type()), F: marshalGestureLongPress},
	})
}

// GestureLongPress is a Gesture implementation able to recognize long presses,
// triggering the GestureLongPress::pressed after the timeout is exceeded.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the GestureLongPress::cancelled signal will
// be emitted.
type GestureLongPress interface {
	gextras.Objector

	privateGestureLongPressClass()
}

// GestureLongPressClass implements the GestureLongPress interface.
type GestureLongPressClass struct {
	GestureSingleClass
}

var _ GestureLongPress = (*GestureLongPressClass)(nil)

func wrapGestureLongPress(obj *externglib.Object) GestureLongPress {
	return &GestureLongPressClass{
		GestureSingleClass: GestureSingleClass{
			GestureClass: GestureClass{
				EventControllerClass: EventControllerClass{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureLongPress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureLongPress(obj), nil
}

// NewGestureLongPress returns a newly created Gesture that recognizes long
// presses.
func NewGestureLongPress(widget Widget) GestureLongPress {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_long_press_new(_arg1)

	var _gestureLongPress GestureLongPress // out

	_gestureLongPress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(GestureLongPress)

	return _gestureLongPress
}

func (*GestureLongPressClass) privateGestureLongPressClass() {}
