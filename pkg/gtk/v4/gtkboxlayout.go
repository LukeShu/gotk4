// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeBoxLayout returns the GType for the type BoxLayout.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBoxLayout() coreglib.Type {
	gtype := coreglib.Type(C.gtk_box_layout_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalBoxLayout)
	return gtype
}

// BoxLayoutOverrider contains methods that are overridable.
type BoxLayoutOverrider interface {
}

// BoxLayout: GtkBoxLayout is a layout manager that arranges children in a
// single row or column.
//
// Whether it is a row or column depends on the value of its
// gtk.Orientable:orientation property. Within the other dimension all children
// all allocated the same size. The GtkBoxLayout will respect the
// gtk.Widget:halign and gtk.Widget:valign properties of each child widget.
//
// If you want all children to be assigned the same size, you can use the
// gtk.BoxLayout:homogeneous property.
//
// If you want to specify the amount of space placed between each child, you can
// use the gtk.BoxLayout:spacing property.
type BoxLayout struct {
	_ [0]func() // equal guard
	LayoutManager

	*coreglib.Object
	Orientable
}

var (
	_ LayoutManagerer   = (*BoxLayout)(nil)
	_ coreglib.Objector = (*BoxLayout)(nil)
)

func classInitBoxLayouter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBoxLayout(obj *coreglib.Object) *BoxLayout {
	return &BoxLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalBoxLayout(p uintptr) (interface{}, error) {
	return wrapBoxLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBoxLayout creates a new GtkBoxLayout.
//
// The function takes the following parameters:
//
//    - orientation for the new layout.
//
// The function returns the following values:
//
//    - boxLayout: new box layout.
//
func NewBoxLayout(orientation Orientation) *BoxLayout {
	var _arg1 C.GtkOrientation    // out
	var _cret *C.GtkLayoutManager // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_box_layout_new(_arg1)
	runtime.KeepAlive(orientation)

	var _boxLayout *BoxLayout // out

	_boxLayout = wrapBoxLayout(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _boxLayout
}

// BaselinePosition gets the value set by
// gtk_box_layout_set_baseline_position().
//
// The function returns the following values:
//
//    - baselinePosition: baseline position.
//
func (boxLayout *BoxLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBoxLayout       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(coreglib.InternObject(boxLayout).Native()))

	_cret = C.gtk_box_layout_get_baseline_position(_arg0)
	runtime.KeepAlive(boxLayout)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// Homogeneous returns whether the layout is set to be homogeneous.
//
// The function returns the following values:
//
//    - ok: TRUE if the layout is homogeneous.
//
func (boxLayout *BoxLayout) Homogeneous() bool {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(coreglib.InternObject(boxLayout).Native()))

	_cret = C.gtk_box_layout_get_homogeneous(_arg0)
	runtime.KeepAlive(boxLayout)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Spacing returns the space that box_layout puts between children.
//
// The function returns the following values:
//
//    - guint: spacing of the layout.
//
func (boxLayout *BoxLayout) Spacing() uint32 {
	var _arg0 *C.GtkBoxLayout // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(coreglib.InternObject(boxLayout).Native()))

	_cret = C.gtk_box_layout_get_spacing(_arg0)
	runtime.KeepAlive(boxLayout)

	var _guint uint32 // out

	_guint = uint32(_cret)

	return _guint
}

// SetBaselinePosition sets the baseline position of a box layout.
//
// The baseline position affects only horizontal boxes with at least one
// baseline aligned child. If there is more vertical space available than
// requested, and the baseline is not allocated by the parent then the given
// position is used to allocate the baseline within the extra space available.
//
// The function takes the following parameters:
//
//    - position: GtkBaselinePosition.
//
func (boxLayout *BoxLayout) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkBoxLayout       // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(coreglib.InternObject(boxLayout).Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_box_layout_set_baseline_position(_arg0, _arg1)
	runtime.KeepAlive(boxLayout)
	runtime.KeepAlive(position)
}

// SetHomogeneous sets whether the box layout will allocate the same size to all
// children.
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE to set the box layout as homogeneous.
//
func (boxLayout *BoxLayout) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(coreglib.InternObject(boxLayout).Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_layout_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(boxLayout)
	runtime.KeepAlive(homogeneous)
}

// SetSpacing sets how much spacing to put between children.
//
// The function takes the following parameters:
//
//    - spacing to apply between children.
//
func (boxLayout *BoxLayout) SetSpacing(spacing uint32) {
	var _arg0 *C.GtkBoxLayout // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkBoxLayout)(unsafe.Pointer(coreglib.InternObject(boxLayout).Native()))
	_arg1 = C.guint(spacing)

	C.gtk_box_layout_set_spacing(_arg0, _arg1)
	runtime.KeepAlive(boxLayout)
	runtime.KeepAlive(spacing)
}
