// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_stack_sidebar_get_type()), F: marshalStackSidebarrer},
	})
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
	Bin
}

func wrapStackSidebar(obj *externglib.Object) *StackSidebar {
	return &StackSidebar{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
		},
	}
}

func marshalStackSidebarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStackSidebar(obj), nil
}

// NewStackSidebar creates a new sidebar.
func NewStackSidebar() *StackSidebar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_sidebar_new()

	var _stackSidebar *StackSidebar // out

	_stackSidebar = wrapStackSidebar(externglib.Take(unsafe.Pointer(_cret)))

	return _stackSidebar
}

// Stack retrieves the stack. See gtk_stack_sidebar_set_stack().
func (sidebar *StackSidebar) Stack() *Stack {
	var _arg0 *C.GtkStackSidebar // out
	var _cret *C.GtkStack        // in

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_stack_sidebar_get_stack(_arg0)

	runtime.KeepAlive(sidebar)

	var _stack *Stack // out

	if _cret != nil {
		_stack = wrapStack(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _stack
}

// SetStack: set the Stack associated with this StackSidebar.
//
// The sidebar widget will automatically update according to the order (packing)
// and items within the given Stack.
func (sidebar *StackSidebar) SetStack(stack *Stack) {
	var _arg0 *C.GtkStackSidebar // out
	var _arg1 *C.GtkStack        // out

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_sidebar_set_stack(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(stack)
}
