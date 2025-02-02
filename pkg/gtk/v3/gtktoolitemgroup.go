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

// ToolItemGroupClass: instance of this type is always passed by reference.
type ToolItemGroupClass struct {
	*toolItemGroupClass
}

// toolItemGroupClass is the struct that's finalized.
type toolItemGroupClass struct {
	native *C.GtkToolItemGroupClass
}

// ParentClass: parent class.
func (t *ToolItemGroupClass) ParentClass() *ContainerClass {
	valptr := &t.native.parent_class
	var _v *ContainerClass // out
	_v = (*ContainerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
