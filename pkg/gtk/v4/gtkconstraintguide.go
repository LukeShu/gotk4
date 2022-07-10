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
// #include <glib-object.h>
import "C"

// GTypeConstraintGuide returns the GType for the type ConstraintGuide.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeConstraintGuide() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ConstraintGuide").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalConstraintGuide)
	return gtype
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
	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_gret := _info.InvokeClassMethod("new_ConstraintGuide", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _constraintGuide *ConstraintGuide // out

	_constraintGuide = wrapConstraintGuide(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _constraintGuide
}

// MaxSize gets the maximum size of guide.
//
// The function takes the following parameters:
//
//    - width (optional): return location for the maximum width, or NULL.
//    - height (optional): return location for the maximum height, or NULL.
//
func (guide *ConstraintGuide) MaxSize(width, height *int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if width != nil {
		*(**C.int)(unsafe.Pointer(&_args[1])) = (*C.int)(unsafe.Pointer(width))
	}
	if height != nil {
		*(**C.int)(unsafe.Pointer(&_args[2])) = (*C.int)(unsafe.Pointer(height))
	}

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_info.InvokeClassMethod("get_max_size", _args[:], nil)

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
func (guide *ConstraintGuide) MinSize(width, height *int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if width != nil {
		*(**C.int)(unsafe.Pointer(&_args[1])) = (*C.int)(unsafe.Pointer(width))
	}
	if height != nil {
		*(**C.int)(unsafe.Pointer(&_args[2])) = (*C.int)(unsafe.Pointer(height))
	}

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_info.InvokeClassMethod("get_min_size", _args[:], nil)

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_gret := _info.InvokeClassMethod("get_name", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(guide)

	var _utf8 string // out

	if *(**C.char)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))
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
func (guide *ConstraintGuide) NatSize(width, height *int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if width != nil {
		*(**C.int)(unsafe.Pointer(&_args[1])) = (*C.int)(unsafe.Pointer(width))
	}
	if height != nil {
		*(**C.int)(unsafe.Pointer(&_args[2])) = (*C.int)(unsafe.Pointer(height))
	}

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_info.InvokeClassMethod("get_nat_size", _args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetMaxSize sets the maximum size of guide.
//
// If guide is attached to a GtkConstraintLayout, the constraints will be
// updated to reflect the new size.
//
// The function takes the following parameters:
//
//    - width: new maximum width, or -1 to not change it.
//    - height: new maximum height, or -1 to not change it.
//
func (guide *ConstraintGuide) SetMaxSize(width, height int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(height)

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_info.InvokeClassMethod("set_max_size", _args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetMinSize sets the minimum size of guide.
//
// If guide is attached to a GtkConstraintLayout, the constraints will be
// updated to reflect the new size.
//
// The function takes the following parameters:
//
//    - width: new minimum width, or -1 to not change it.
//    - height: new minimum height, or -1 to not change it.
//
func (guide *ConstraintGuide) SetMinSize(width, height int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(height)

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_info.InvokeClassMethod("set_min_size", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	if name != "" {
		*(**C.char)(unsafe.Pointer(&_args[1])) = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[1]))))
	}

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_info.InvokeClassMethod("set_name", _args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(name)
}

// SetNatSize sets the natural size of guide.
//
// If guide is attached to a GtkConstraintLayout, the constraints will be
// updated to reflect the new size.
//
// The function takes the following parameters:
//
//    - width: new natural width, or -1 to not change it.
//    - height: new natural height, or -1 to not change it.
//
func (guide *ConstraintGuide) SetNatSize(width, height int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(guide).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(height)

	_info := girepository.MustFind("Gtk", "ConstraintGuide")
	_info.InvokeClassMethod("set_nat_size", _args[:], nil)

	runtime.KeepAlive(guide)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}
