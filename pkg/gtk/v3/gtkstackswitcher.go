// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
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

// StackSwitcherOverrides contains methods that are overridable.
type StackSwitcherOverrides struct {
}

func defaultStackSwitcherOverrides(v *StackSwitcher) StackSwitcherOverrides {
	return StackSwitcherOverrides{}
}

// StackSwitcher widget acts as a controller for a Stack; it shows a row of
// buttons to switch between the various pages of the associated stack widget.
//
// All the content for the buttons comes from the child properties of the Stack;
// the button visibility in a StackSwitcher widget is controlled by the
// visibility of the child in the Stack.
//
// It is possible to associate multiple StackSwitcher widgets with the same
// Stack widget.
//
// The GtkStackSwitcher widget was added in 3.10.
//
//
// CSS nodes
//
// GtkStackSwitcher has a single CSS node named stackswitcher and style class
// .stack-switcher.
//
// When circumstances require it, GtkStackSwitcher adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSwitcher struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*StackSwitcher)(nil)
	_ coreglib.Objector = (*StackSwitcher)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StackSwitcher, *StackSwitcherClass, StackSwitcherOverrides](
		GTypeStackSwitcher,
		initStackSwitcherClass,
		wrapStackSwitcher,
		defaultStackSwitcherOverrides,
	)
}

func initStackSwitcherClass(gclass unsafe.Pointer, overrides StackSwitcherOverrides, classInitFunc func(*StackSwitcherClass)) {
	if classInitFunc != nil {
		class := (*StackSwitcherClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStackSwitcher(obj *coreglib.Object) *StackSwitcher {
	return &StackSwitcher{
		Box: Box{
			Container: Container{
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
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalStackSwitcher(p uintptr) (interface{}, error) {
	return wrapStackSwitcher(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStackSwitcher: create a new StackSwitcher.
//
// The function returns the following values:
//
//    - stackSwitcher: new StackSwitcher.
//
func NewStackSwitcher() *StackSwitcher {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_switcher_new()

	var _stackSwitcher *StackSwitcher // out

	_stackSwitcher = wrapStackSwitcher(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackSwitcher
}

// Stack retrieves the stack. See gtk_stack_switcher_set_stack().
//
// The function returns the following values:
//
//    - stack (optional): stack, or NULL if none has been set explicitly.
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
//    - stack (optional): Stack.
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

// StackSwitcherClass: instance of this type is always passed by reference.
type StackSwitcherClass struct {
	*stackSwitcherClass
}

// stackSwitcherClass is the struct that's finalized.
type stackSwitcherClass struct {
	native *C.GtkStackSwitcherClass
}

func (s *StackSwitcherClass) ParentClass() *BoxClass {
	valptr := &s.native.parent_class
	var _v *BoxClass // out
	_v = (*BoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
