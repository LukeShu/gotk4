// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeOverlayLayout      = coreglib.Type(C.gtk_overlay_layout_get_type())
	GTypeOverlayLayoutChild = coreglib.Type(C.gtk_overlay_layout_child_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOverlayLayout, F: marshalOverlayLayout},
		coreglib.TypeMarshaler{T: GTypeOverlayLayoutChild, F: marshalOverlayLayoutChild},
	})
}

// OverlayLayoutOverrides contains methods that are overridable.
type OverlayLayoutOverrides struct {
}

func defaultOverlayLayoutOverrides(v *OverlayLayout) OverlayLayoutOverrides {
	return OverlayLayoutOverrides{}
}

// OverlayLayout: GtkOverlayLayout is the layout manager used by GtkOverlay.
//
// It places widgets as overlays on top of the main child.
//
// This is not a reusable layout manager, since it expects its widget to be a
// GtkOverlay. It only listed here so that its layout properties get documented.
type OverlayLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*OverlayLayout)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*OverlayLayout, *OverlayLayoutClass, OverlayLayoutOverrides](
		GTypeOverlayLayout,
		initOverlayLayoutClass,
		wrapOverlayLayout,
		defaultOverlayLayoutOverrides,
	)
}

func initOverlayLayoutClass(gclass unsafe.Pointer, overrides OverlayLayoutOverrides, classInitFunc func(*OverlayLayoutClass)) {
	if classInitFunc != nil {
		class := (*OverlayLayoutClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapOverlayLayout(obj *coreglib.Object) *OverlayLayout {
	return &OverlayLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalOverlayLayout(p uintptr) (interface{}, error) {
	return wrapOverlayLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewOverlayLayout creates a new GtkOverlayLayout instance.
//
// The function returns the following values:
//
//    - overlayLayout: newly created instance.
//
func NewOverlayLayout() *OverlayLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_overlay_layout_new()

	var _overlayLayout *OverlayLayout // out

	_overlayLayout = wrapOverlayLayout(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _overlayLayout
}

// OverlayLayoutChildOverrides contains methods that are overridable.
type OverlayLayoutChildOverrides struct {
}

func defaultOverlayLayoutChildOverrides(v *OverlayLayoutChild) OverlayLayoutChildOverrides {
	return OverlayLayoutChildOverrides{}
}

// OverlayLayoutChild: GtkLayoutChild subclass for children in a
// GtkOverlayLayout.
type OverlayLayoutChild struct {
	_ [0]func() // equal guard
	LayoutChild
}

var (
	_ LayoutChilder = (*OverlayLayoutChild)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*OverlayLayoutChild, *OverlayLayoutChildClass, OverlayLayoutChildOverrides](
		GTypeOverlayLayoutChild,
		initOverlayLayoutChildClass,
		wrapOverlayLayoutChild,
		defaultOverlayLayoutChildOverrides,
	)
}

func initOverlayLayoutChildClass(gclass unsafe.Pointer, overrides OverlayLayoutChildOverrides, classInitFunc func(*OverlayLayoutChildClass)) {
	if classInitFunc != nil {
		class := (*OverlayLayoutChildClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapOverlayLayoutChild(obj *coreglib.Object) *OverlayLayoutChild {
	return &OverlayLayoutChild{
		LayoutChild: LayoutChild{
			Object: obj,
		},
	}
}

func marshalOverlayLayoutChild(p uintptr) (interface{}, error) {
	return wrapOverlayLayoutChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClipOverlay retrieves whether the child is clipped.
//
// The function returns the following values:
//
//    - ok: whether the child is clipped.
//
func (child *OverlayLayoutChild) ClipOverlay() bool {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.gtk_overlay_layout_child_get_clip_overlay(_arg0)
	runtime.KeepAlive(child)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Measure retrieves whether the child is measured.
//
// The function returns the following values:
//
//    - ok: whether the child is measured.
//
func (child *OverlayLayoutChild) Measure() bool {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_cret = C.gtk_overlay_layout_child_get_measure(_arg0)
	runtime.KeepAlive(child)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetClipOverlay sets whether to clip this child.
//
// The function takes the following parameters:
//
//    - clipOverlay: whether to clip this child.
//
func (child *OverlayLayoutChild) SetClipOverlay(clipOverlay bool) {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if clipOverlay {
		_arg1 = C.TRUE
	}

	C.gtk_overlay_layout_child_set_clip_overlay(_arg0, _arg1)
	runtime.KeepAlive(child)
	runtime.KeepAlive(clipOverlay)
}

// SetMeasure sets whether to measure this child.
//
// The function takes the following parameters:
//
//    - measure: whether to measure this child.
//
func (child *OverlayLayoutChild) SetMeasure(measure bool) {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if measure {
		_arg1 = C.TRUE
	}

	C.gtk_overlay_layout_child_set_measure(_arg0, _arg1)
	runtime.KeepAlive(child)
	runtime.KeepAlive(measure)
}

// OverlayLayoutChildClass: instance of this type is always passed by reference.
type OverlayLayoutChildClass struct {
	*overlayLayoutChildClass
}

// overlayLayoutChildClass is the struct that's finalized.
type overlayLayoutChildClass struct {
	native *C.GtkOverlayLayoutChildClass
}

func (o *OverlayLayoutChildClass) ParentClass() *LayoutChildClass {
	valptr := &o.native.parent_class
	var _v *LayoutChildClass // out
	_v = (*LayoutChildClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// OverlayLayoutClass: instance of this type is always passed by reference.
type OverlayLayoutClass struct {
	*overlayLayoutClass
}

// overlayLayoutClass is the struct that's finalized.
type overlayLayoutClass struct {
	native *C.GtkOverlayLayoutClass
}

func (o *OverlayLayoutClass) ParentClass() *LayoutManagerClass {
	valptr := &o.native.parent_class
	var _v *LayoutManagerClass // out
	_v = (*LayoutManagerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
