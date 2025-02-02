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

// GType values.
var (
	GTypeStackSwitcher = coreglib.Type(C.gtk_stack_switcher_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStackSwitcher, F: marshalStackSwitcher},
	})
}

// StackSwitcher: GtkStackSwitcher shows a row of buttons to switch between
// GtkStack pages.
//
// !An example GtkStackSwitcher (stackswitcher.png)
//
// It acts as a controller for the associated GtkStack.
//
// All the content for the buttons comes from the properties of the stacks
// gtk.StackPage objects; the button visibility in a GtkStackSwitcher widget is
// controlled by the visibility of the child in the GtkStack.
//
// It is possible to associate multiple GtkStackSwitcher widgets with the same
// GtkStack widget.
//
// # CSS nodes
//
// GtkStackSwitcher has a single CSS node named stackswitcher and style class
// .stack-switcher.
//
// When circumstances require it, GtkStackSwitcher adds the .needs-attention
// style class to the widgets representing the stack pages.
//
// # Accessibility
//
// GtkStackSwitcher uses the GTK_ACCESSIBLE_ROLE_TAB_LIST role and uses the
// GTK_ACCESSIBLE_ROLE_TAB for its buttons.
type StackSwitcher struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*StackSwitcher)(nil)
)

func wrapStackSwitcher(obj *coreglib.Object) *StackSwitcher {
	return &StackSwitcher{
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

func marshalStackSwitcher(p uintptr) (interface{}, error) {
	return wrapStackSwitcher(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStackSwitcher: create a new GtkStackSwitcher.
//
// The function returns the following values:
//
//   - stackSwitcher: new GtkStackSwitcher.
//
func NewStackSwitcher() *StackSwitcher {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_switcher_new()

	var _stackSwitcher *StackSwitcher // out

	_stackSwitcher = wrapStackSwitcher(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackSwitcher
}

// Stack retrieves the stack.
//
// The function returns the following values:
//
//   - stack (optional): stack, or NULL if none has been set explicitly.
//
func (switcher *StackSwitcher) Stack() *Stack {
	var _arg0 *C.GtkStackSwitcher // out
	var _cret *C.GtkStack         // in

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(coreglib.InternObject(switcher).Native()))

	_cret = C.gtk_stack_switcher_get_stack(_arg0)
	runtime.KeepAlive(switcher)

	var _stack *Stack // out

	if _cret != nil {
		_stack = wrapStack(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _stack
}

// SetStack sets the stack to control.
//
// The function takes the following parameters:
//
//   - stack (optional): GtkStack.
//
func (switcher *StackSwitcher) SetStack(stack *Stack) {
	var _arg0 *C.GtkStackSwitcher // out
	var _arg1 *C.GtkStack         // out

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(coreglib.InternObject(switcher).Native()))
	if stack != nil {
		_arg1 = (*C.GtkStack)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	}

	C.gtk_stack_switcher_set_stack(_arg0, _arg1)
	runtime.KeepAlive(switcher)
	runtime.KeepAlive(stack)
}
