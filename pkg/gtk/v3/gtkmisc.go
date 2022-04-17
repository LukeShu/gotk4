// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkmisc.go.
var GTypeMisc = externglib.Type(C.gtk_misc_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeMisc, F: marshalMisc},
	})
}

// MiscOverrider contains methods that are overridable.
type MiscOverrider interface {
	externglib.Objector
}

// WrapMiscOverrider wraps the MiscOverrider
// interface implementation to access the instance methods.
func WrapMiscOverrider(obj MiscOverrider) *Misc {
	return wrapMisc(externglib.BaseObject(obj))
}

// Misc widget is an abstract widget which is not useful itself, but is used to
// derive subclasses which have alignment and padding attributes.
//
// The horizontal and vertical padding attributes allows extra space to be added
// around the widget.
//
// The horizontal and vertical alignment attributes enable the widget to be
// positioned within its allocated area. Note that if the widget is added to a
// container in such a way that it expands automatically to fill its allocated
// area, the alignment settings will not alter the widget's position.
//
// Note that the desired effect can in most cases be achieved by using the
// Widget:halign, Widget:valign and Widget:margin properties on the child
// widget, so GtkMisc should not be used in new code. To reflect this fact, all
// Misc API has been deprecated.
type Misc struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Misc)(nil)
)

// Miscer describes types inherited from class Misc.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Miscer interface {
	externglib.Objector
	baseMisc() *Misc
}

var _ Miscer = (*Misc)(nil)

func classInitMiscer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMisc(obj *externglib.Object) *Misc {
	return &Misc{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalMisc(p uintptr) (interface{}, error) {
	return wrapMisc(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (misc *Misc) baseMisc() *Misc {
	return misc
}

// BaseMisc returns the underlying base object from the
// interface.
func BaseMisc(obj Miscer) *Misc {
	return obj.baseMisc()
}

// Alignment gets the X and Y alignment of the widget within its allocation. See
// gtk_misc_set_alignment().
//
// Deprecated: Use Widget alignment and margin properties.
//
// The function returns the following values:
//
//    - xalign (optional): location to store X alignment of misc, or NULL.
//    - yalign (optional): location to store Y alignment of misc, or NULL.
//
func (misc *Misc) Alignment() (xalign float32, yalign float32) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gfloat   // in
	var _arg2 C.gfloat   // in

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(externglib.InternObject(misc).Native()))

	C.gtk_misc_get_alignment(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(misc)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// Padding gets the padding in the X and Y directions of the widget. See
// gtk_misc_set_padding().
//
// Deprecated: Use Widget alignment and margin properties.
//
// The function returns the following values:
//
//    - xpad (optional): location to store padding in the X direction, or NULL.
//    - ypad (optional): location to store padding in the Y direction, or NULL.
//
func (misc *Misc) Padding() (xpad int, ypad int) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gint     // in
	var _arg2 C.gint     // in

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(externglib.InternObject(misc).Native()))

	C.gtk_misc_get_padding(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(misc)

	var _xpad int // out
	var _ypad int // out

	_xpad = int(_arg1)
	_ypad = int(_arg2)

	return _xpad, _ypad
}

// SetAlignment sets the alignment of the widget.
//
// Deprecated: Use Widget's alignment (Widget:halign and Widget:valign) and
// margin properties or Label's Label:xalign and Label:yalign properties.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment, from 0 (left) to 1 (right).
//    - yalign: vertical alignment, from 0 (top) to 1 (bottom).
//
func (misc *Misc) SetAlignment(xalign, yalign float32) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gfloat   // out
	var _arg2 C.gfloat   // out

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(externglib.InternObject(misc).Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)

	C.gtk_misc_set_alignment(_arg0, _arg1, _arg2)
	runtime.KeepAlive(misc)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
}

// SetPadding sets the amount of space to add around the widget.
//
// Deprecated: Use Widget alignment and margin properties.
//
// The function takes the following parameters:
//
//    - xpad: amount of space to add on the left and right of the widget, in
//      pixels.
//    - ypad: amount of space to add on the top and bottom of the widget, in
//      pixels.
//
func (misc *Misc) SetPadding(xpad, ypad int) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gint     // out
	var _arg2 C.gint     // out

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(externglib.InternObject(misc).Native()))
	_arg1 = C.gint(xpad)
	_arg2 = C.gint(ypad)

	C.gtk_misc_set_padding(_arg0, _arg1, _arg2)
	runtime.KeepAlive(misc)
	runtime.KeepAlive(xpad)
	runtime.KeepAlive(ypad)
}
