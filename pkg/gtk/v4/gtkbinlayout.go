// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeBinLayout = coreglib.Type(C.gtk_bin_layout_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBinLayout, F: marshalBinLayout},
	})
}

// BinLayoutOverrides contains methods that are overridable.
type BinLayoutOverrides struct {
}

func defaultBinLayoutOverrides(v *BinLayout) BinLayoutOverrides {
	return BinLayoutOverrides{}
}

// BinLayout: GtkBinLayout is a GtkLayoutManager subclass useful for create
// "bins" of widgets.
//
// GtkBinLayout will stack each child of a widget on top of each other,
// using the gtk.Widget:hexpand, gtk.Widget:vexpand, gtk.Widget:halign,
// and gtk.Widget:valign properties of each child to determine where they should
// be positioned.
type BinLayout struct {
	_ [0]func() // equal guard
	LayoutManager
}

var (
	_ LayoutManagerer = (*BinLayout)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*BinLayout, *BinLayoutClass, BinLayoutOverrides](
		GTypeBinLayout,
		initBinLayoutClass,
		wrapBinLayout,
		defaultBinLayoutOverrides,
	)
}

func initBinLayoutClass(gclass unsafe.Pointer, overrides BinLayoutOverrides, classInitFunc func(*BinLayoutClass)) {
	if classInitFunc != nil {
		class := (*BinLayoutClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBinLayout(obj *coreglib.Object) *BinLayout {
	return &BinLayout{
		LayoutManager: LayoutManager{
			Object: obj,
		},
	}
}

func marshalBinLayout(p uintptr) (interface{}, error) {
	return wrapBinLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBinLayout creates a new GtkBinLayout instance.
//
// The function returns the following values:
//
//   - binLayout: newly created GtkBinLayout.
//
func NewBinLayout() *BinLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_bin_layout_new()

	var _binLayout *BinLayout // out

	_binLayout = wrapBinLayout(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _binLayout
}

// BinLayoutClass: instance of this type is always passed by reference.
type BinLayoutClass struct {
	*binLayoutClass
}

// binLayoutClass is the struct that's finalized.
type binLayoutClass struct {
	native *C.GtkBinLayoutClass
}

func (b *BinLayoutClass) ParentClass() *LayoutManagerClass {
	valptr := &b.native.parent_class
	var _v *LayoutManagerClass // out
	_v = (*LayoutManagerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
