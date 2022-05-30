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

// glib.Type values for gtkstacksidebar.go.
var GTypeStackSidebar = coreglib.Type(C.gtk_stack_sidebar_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeStackSidebar, F: marshalStackSidebar},
	})
}

// StackSidebar: GtkStackSidebar uses a sidebar to switch between GtkStack
// pages.
//
// In order to use a GtkStackSidebar, you simply use a GtkStack to organize your
// UI flow, and add the sidebar to your sidebar area. You can use
// gtk.StackSidebar.SetStack() to connect the GtkStackSidebar to the GtkStack.
//
//
// CSS nodes
//
// GtkStackSidebar has a single CSS node with name stacksidebar and style class
// .sidebar.
//
// When circumstances require it, GtkStackSidebar adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSidebar struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*StackSidebar)(nil)
)

func wrapStackSidebar(obj *coreglib.Object) *StackSidebar {
	return &StackSidebar{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalStackSidebar(p uintptr) (interface{}, error) {
	return wrapStackSidebar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStackSidebar creates a new GtkStackSidebar.
//
// The function returns the following values:
//
//    - stackSidebar: new GtkStackSidebar.
//
func NewStackSidebar() *StackSidebar {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "StackSidebar").InvokeMethod("new_StackSidebar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _stackSidebar *StackSidebar // out

	_stackSidebar = wrapStackSidebar(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackSidebar
}

// Stack retrieves the stack.
//
// The function returns the following values:
//
//    - stack (optional): associated Stack or NULL if none has been set
//      explicitly.
//
func (self *StackSidebar) Stack() *Stack {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**StackSidebar)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "StackSidebar").InvokeMethod("get_stack", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _stack *Stack // out

	if _cret != nil {
		_stack = wrapStack(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _stack
}

// SetStack: set the GtkStack associated with this GtkStackSidebar.
//
// The sidebar widget will automatically update according to the order and items
// within the given GtkStack.
//
// The function takes the following parameters:
//
//    - stack: GtkStack.
//
func (self *StackSidebar) SetStack(stack *Stack) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**StackSidebar)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "StackSidebar").InvokeMethod("set_stack", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(stack)
}
