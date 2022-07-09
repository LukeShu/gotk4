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

// GTypeMisc returns the GType for the type Misc.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMisc() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Misc").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMisc)
	return gtype
}

// MiscOverrider contains methods that are overridable.
type MiscOverrider interface {
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
	coreglib.Objector
	baseMisc() *Misc
}

var _ Miscer = (*Misc)(nil)

func classInitMiscer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapMisc(obj *coreglib.Object) *Misc {
	return &Misc{
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
	}
}

func marshalMisc(p uintptr) (interface{}, error) {
	return wrapMisc(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (misc *Misc) baseMisc() *Misc {
	return misc
}

// BaseMisc returns the underlying base object.
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
func (misc *Misc) Alignment() (xalign, yalign float32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(misc).Native()))

	girepository.MustFind("Gtk", "Misc").InvokeMethod("get_alignment", _args[:], _outs[:])

	runtime.KeepAlive(misc)

	var _xalign float32 // out
	var _yalign float32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_xalign = *(*float32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_yalign = *(*float32)(unsafe.Pointer(_outs[1]))
	}

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
func (misc *Misc) Padding() (xpad, ypad int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(misc).Native()))

	girepository.MustFind("Gtk", "Misc").InvokeMethod("get_padding", _args[:], _outs[:])

	runtime.KeepAlive(misc)

	var _xpad int32 // out
	var _ypad int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_xpad = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_ypad = *(*int32)(unsafe.Pointer(_outs[1]))
	}

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(misc).Native()))
	*(*C.gfloat)(unsafe.Pointer(&_args[1])) = C.gfloat(xalign)
	*(*C.gfloat)(unsafe.Pointer(&_args[2])) = C.gfloat(yalign)

	girepository.MustFind("Gtk", "Misc").InvokeMethod("set_alignment", _args[:], nil)

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
func (misc *Misc) SetPadding(xpad, ypad int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(misc).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(xpad)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(ypad)

	girepository.MustFind("Gtk", "Misc").InvokeMethod("set_padding", _args[:], nil)

	runtime.KeepAlive(misc)
	runtime.KeepAlive(xpad)
	runtime.KeepAlive(ypad)
}
