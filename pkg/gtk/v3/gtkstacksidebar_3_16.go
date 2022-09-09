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
	GTypeStackSidebar = coreglib.Type(C.gtk_stack_sidebar_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStackSidebar, F: marshalStackSidebar},
	})
}

// StackSidebarOverrides contains methods that are overridable.
type StackSidebarOverrides struct {
}

func defaultStackSidebarOverrides(v *StackSidebar) StackSidebarOverrides {
	return StackSidebarOverrides{}
}

// StackSidebar enables you to quickly and easily provide a consistent "sidebar"
// object for your user interface.
//
// In order to use a GtkStackSidebar, you simply use a GtkStack to organize your
// UI flow, and add the sidebar to your sidebar area. You can use
// gtk_stack_sidebar_set_stack() to connect the StackSidebar to the Stack.
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
	Bin
}

var (
	_ Binner = (*StackSidebar)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StackSidebar, *StackSidebarClass, StackSidebarOverrides](
		GTypeStackSidebar,
		initStackSidebarClass,
		wrapStackSidebar,
		defaultStackSidebarOverrides,
	)
}

func initStackSidebarClass(gclass unsafe.Pointer, overrides StackSidebarOverrides, classInitFunc func(*StackSidebarClass)) {
	if classInitFunc != nil {
		class := (*StackSidebarClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStackSidebar(obj *coreglib.Object) *StackSidebar {
	return &StackSidebar{
		Bin: Bin{
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
		},
	}
}

func marshalStackSidebar(p uintptr) (interface{}, error) {
	return wrapStackSidebar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStackSidebar creates a new sidebar.
//
// The function returns the following values:
//
//    - stackSidebar: new StackSidebar.
//
func NewStackSidebar() *StackSidebar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_sidebar_new()

	var _stackSidebar *StackSidebar // out

	_stackSidebar = wrapStackSidebar(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackSidebar
}

// Stack retrieves the stack. See gtk_stack_sidebar_set_stack().
//
// The function returns the following values:
//
//    - stack (optional): associated Stack or NULL if none has been set
//      explicitly.
//
func (sidebar *StackSidebar) Stack() *Stack {
	var _arg0 *C.GtkStackSidebar // out
	var _cret *C.GtkStack        // in

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(coreglib.InternObject(sidebar).Native()))

	_cret = C.gtk_stack_sidebar_get_stack(_arg0)
	runtime.KeepAlive(sidebar)

	var _stack *Stack // out

	if _cret != nil {
		_stack = wrapStack(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _stack
}

// SetStack: set the Stack associated with this StackSidebar.
//
// The sidebar widget will automatically update according to the order (packing)
// and items within the given Stack.
//
// The function takes the following parameters:
//
//    - stack: Stack.
//
func (sidebar *StackSidebar) SetStack(stack *Stack) {
	var _arg0 *C.GtkStackSidebar // out
	var _arg1 *C.GtkStack        // out

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(coreglib.InternObject(sidebar).Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	C.gtk_stack_sidebar_set_stack(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(stack)
}
