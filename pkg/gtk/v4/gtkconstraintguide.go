// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkconstraintguide.go.
var GTypeConstraintGuide = coreglib.Type(C.gtk_constraint_guide_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeConstraintGuide, F: marshalConstraintGuide},
	})
}

// ConstraintGuideOverrider contains methods that are overridable.
type ConstraintGuideOverrider interface {
}

// ConstraintGuide: GtkConstraintGuide is an invisible layout element in a
// GtkConstraintLayout.
//
// The GtkConstraintLayout treats guides like widgets. They can be used as the
// source or target of a GtkConstraint.
//
// Guides have a minimum, maximum and natural size. Depending on the constraints
// that are applied, they can act like a guideline that widgets can be aligned
// to, or like *flexible space*.
//
// Unlike a GtkWidget, a GtkConstraintGuide will not be drawn.
type ConstraintGuide struct {
	_ [0]func() // equal guard
	*coreglib.Object

	ConstraintTarget
}

var (
	_ coreglib.Objector = (*ConstraintGuide)(nil)
)

func classInitConstraintGuider(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapConstraintGuide(obj *coreglib.Object) *ConstraintGuide {
	return &ConstraintGuide{
		Object: obj,
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
	}
}

func marshalConstraintGuide(p uintptr) (interface{}, error) {
	return wrapConstraintGuide(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConstraintGuide creates a new GtkConstraintGuide object.
//
// The function returns the following values:
//
//    - constraintGuide: new GtkConstraintGuide object.
//
func NewConstraintGuide() *ConstraintGuide {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ConstraintGuide").InvokeMethod("new_ConstraintGuide", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _constraintGuide *ConstraintGuide // out

	_constraintGuide = wrapConstraintGuide(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constraintGuide
}

// MaxSize gets the maximum size of guide.
//
// The function takes the following parameters:
//
//    - width (optional): return location for the maximum width, or NULL.
//    - height (optional): return location for the maximum height, or NULL.
//
func (guide *ConstraintGuide) MaxSize(width, height *int) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if width != nil {
		_arg1 = (*C.void)(unsafe.Pointer(width))
	}
	if height != nil {
		_arg2 = (*C.void)(unsafe.Pointer(height))
	}
	*(**ConstraintGuide)(unsafe.Pointer(&args[1])) = _arg1
	*(**int)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "ConstraintGuide").InvokeMethod("get_max_size", args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// MinSize gets the minimum size of guide.
//
// The function takes the following parameters:
//
//    - width (optional): return location for the minimum width, or NULL.
//    - height (optional): return location for the minimum height, or NULL.
//
func (guide *ConstraintGuide) MinSize(width, height *int) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if width != nil {
		_arg1 = (*C.void)(unsafe.Pointer(width))
	}
	if height != nil {
		_arg2 = (*C.void)(unsafe.Pointer(height))
	}
	*(**ConstraintGuide)(unsafe.Pointer(&args[1])) = _arg1
	*(**int)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "ConstraintGuide").InvokeMethod("get_min_size", args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// Name retrieves the name set using gtk_constraint_guide_set_name().
//
// The function returns the following values:
//
//    - utf8 (optional): name of the guide.
//
func (guide *ConstraintGuide) Name() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	*(**ConstraintGuide)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ConstraintGuide").InvokeMethod("get_name", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(guide)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NatSize gets the natural size of guide.
//
// The function takes the following parameters:
//
//    - width (optional): return location for the natural width, or NULL.
//    - height (optional): return location for the natural height, or NULL.
//
func (guide *ConstraintGuide) NatSize(width, height *int) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if width != nil {
		_arg1 = (*C.void)(unsafe.Pointer(width))
	}
	if height != nil {
		_arg2 = (*C.void)(unsafe.Pointer(height))
	}
	*(**ConstraintGuide)(unsafe.Pointer(&args[1])) = _arg1
	*(**int)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "ConstraintGuide").InvokeMethod("get_nat_size", args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetName sets a name for the given GtkConstraintGuide.
//
// The name is useful for debugging purposes.
//
// The function takes the following parameters:
//
//    - name (optional) for the guide.
//
func (guide *ConstraintGuide) SetName(name string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if name != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**ConstraintGuide)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ConstraintGuide").InvokeMethod("set_name", args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(name)
}
