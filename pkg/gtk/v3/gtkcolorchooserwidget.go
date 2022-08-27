// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ColorChooserWidgetClass: instance of this type is always passed by reference.
type ColorChooserWidgetClass struct {
	*colorChooserWidgetClass
}

// colorChooserWidgetClass is the struct that's finalized.
type colorChooserWidgetClass struct {
	native *C.GtkColorChooserWidgetClass
}

// ParentClass: parent class.
func (c *ColorChooserWidgetClass) ParentClass() *BoxClass {
	valptr := &c.native.parent_class
	var _v *BoxClass // out
	_v = (*BoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
