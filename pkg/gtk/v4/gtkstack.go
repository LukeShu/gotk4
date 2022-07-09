// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeStackTransitionType returns the GType for the type StackTransitionType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStackTransitionType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "StackTransitionType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStackTransitionType)
	return gtype
}

// GTypeStack returns the GType for the type Stack.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStack() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Stack").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStack)
	return gtype
}

// GTypeStackPage returns the GType for the type StackPage.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStackPage() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "StackPage").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStackPage)
	return gtype
}

// StackTransitionType: possible transitions between pages in a GtkStack widget.
//
// New values may be added to this enumeration over time.
type StackTransitionType C.gint

const (
	// StackTransitionTypeNone: no transition.
	StackTransitionTypeNone StackTransitionType = iota
	// StackTransitionTypeCrossfade: cross-fade.
	StackTransitionTypeCrossfade
	// StackTransitionTypeSlideRight: slide from left to right.
	StackTransitionTypeSlideRight
	// StackTransitionTypeSlideLeft: slide from right to left.
	StackTransitionTypeSlideLeft
	// StackTransitionTypeSlideUp: slide from bottom up.
	StackTransitionTypeSlideUp
	// StackTransitionTypeSlideDown: slide from top down.
	StackTransitionTypeSlideDown
	// StackTransitionTypeSlideLeftRight: slide from left or right according to
	// the children order.
	StackTransitionTypeSlideLeftRight
	// StackTransitionTypeSlideUpDown: slide from top down or bottom up
	// according to the order.
	StackTransitionTypeSlideUpDown
	// StackTransitionTypeOverUp: cover the old page by sliding up.
	StackTransitionTypeOverUp
	// StackTransitionTypeOverDown: cover the old page by sliding down.
	StackTransitionTypeOverDown
	// StackTransitionTypeOverLeft: cover the old page by sliding to the left.
	StackTransitionTypeOverLeft
	// StackTransitionTypeOverRight: cover the old page by sliding to the right.
	StackTransitionTypeOverRight
	// StackTransitionTypeUnderUp: uncover the new page by sliding up.
	StackTransitionTypeUnderUp
	// StackTransitionTypeUnderDown: uncover the new page by sliding down.
	StackTransitionTypeUnderDown
	// StackTransitionTypeUnderLeft: uncover the new page by sliding to the
	// left.
	StackTransitionTypeUnderLeft
	// StackTransitionTypeUnderRight: uncover the new page by sliding to the
	// right.
	StackTransitionTypeUnderRight
	// StackTransitionTypeOverUpDown: cover the old page sliding up or uncover
	// the new page sliding down, according to order.
	StackTransitionTypeOverUpDown
	// StackTransitionTypeOverDownUp: cover the old page sliding down or uncover
	// the new page sliding up, according to order.
	StackTransitionTypeOverDownUp
	// StackTransitionTypeOverLeftRight: cover the old page sliding left or
	// uncover the new page sliding right, according to order.
	StackTransitionTypeOverLeftRight
	// StackTransitionTypeOverRightLeft: cover the old page sliding right or
	// uncover the new page sliding left, according to order.
	StackTransitionTypeOverRightLeft
	// StackTransitionTypeRotateLeft: pretend the pages are sides of a cube and
	// rotate that cube to the left.
	StackTransitionTypeRotateLeft
	// StackTransitionTypeRotateRight: pretend the pages are sides of a cube and
	// rotate that cube to the right.
	StackTransitionTypeRotateRight
	// StackTransitionTypeRotateLeftRight: pretend the pages are sides of a cube
	// and rotate that cube to the left or right according to the children
	// order.
	StackTransitionTypeRotateLeftRight
)

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	return StackTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for StackTransitionType.
func (s StackTransitionType) String() string {
	switch s {
	case StackTransitionTypeNone:
		return "None"
	case StackTransitionTypeCrossfade:
		return "Crossfade"
	case StackTransitionTypeSlideRight:
		return "SlideRight"
	case StackTransitionTypeSlideLeft:
		return "SlideLeft"
	case StackTransitionTypeSlideUp:
		return "SlideUp"
	case StackTransitionTypeSlideDown:
		return "SlideDown"
	case StackTransitionTypeSlideLeftRight:
		return "SlideLeftRight"
	case StackTransitionTypeSlideUpDown:
		return "SlideUpDown"
	case StackTransitionTypeOverUp:
		return "OverUp"
	case StackTransitionTypeOverDown:
		return "OverDown"
	case StackTransitionTypeOverLeft:
		return "OverLeft"
	case StackTransitionTypeOverRight:
		return "OverRight"
	case StackTransitionTypeUnderUp:
		return "UnderUp"
	case StackTransitionTypeUnderDown:
		return "UnderDown"
	case StackTransitionTypeUnderLeft:
		return "UnderLeft"
	case StackTransitionTypeUnderRight:
		return "UnderRight"
	case StackTransitionTypeOverUpDown:
		return "OverUpDown"
	case StackTransitionTypeOverDownUp:
		return "OverDownUp"
	case StackTransitionTypeOverLeftRight:
		return "OverLeftRight"
	case StackTransitionTypeOverRightLeft:
		return "OverRightLeft"
	case StackTransitionTypeRotateLeft:
		return "RotateLeft"
	case StackTransitionTypeRotateRight:
		return "RotateRight"
	case StackTransitionTypeRotateLeftRight:
		return "RotateLeftRight"
	default:
		return fmt.Sprintf("StackTransitionType(%d)", s)
	}
}

// Stack: GtkStack is a container which only shows one of its children at a
// time.
//
// In contrast to GtkNotebook, GtkStack does not provide a means for users to
// change the visible child. Instead, a separate widget such as
// gtk.StackSwitcher or gtk.StackSidebar can be used with GtkStack to provide
// this functionality.
//
// Transitions between pages can be animated as slides or fades. This can be
// controlled with gtk.Stack.SetTransitionType(). These animations respect the
// gtk.Settings:gtk-enable-animations setting.
//
// GtkStack maintains a gtk.StackPage object for each added child, which holds
// additional per-child properties. You obtain the GtkStackPage for a child with
// gtk.Stack.GetPage() and you can obtain a GtkSelectionModel containing all the
// pages with gtk.Stack.GetPages().
//
//
// GtkStack as GtkBuildable
//
// To set child-specific properties in a .ui file, create GtkStackPage objects
// explicitly, and set the child widget as a property on it:
//
//      <object class="GtkStack" id="stack">
//        <child>
//          <object class="GtkStackPage">
//            <property name="name">page1</property>
//            <property name="title">In the beginning…</property>
//            <property name="child">
//              <object class="GtkLabel">
//                <property name="label">It was dark</property>
//              </object>
//            </property>
//          </object>
//        </child>
//
//
//
// CSS nodes
//
// GtkStack has a single CSS node named stack.
//
//
// Accessibility
//
// GtkStack uses the GTK_ACCESSIBLE_ROLE_TAB_PANEL for the stack pages, which
// are the accessible parent objects of the child widgets.
type Stack struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Stack)(nil)
)

func wrapStack(obj *coreglib.Object) *Stack {
	return &Stack{
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

func marshalStack(p uintptr) (interface{}, error) {
	return wrapStack(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStack creates a new GtkStack.
//
// The function returns the following values:
//
//    - stack: new GtkStack.
//
func NewStack() *Stack {
	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("new_Stack", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _stack *Stack // out

	_stack = wrapStack(coreglib.Take(unsafe.Pointer(_cret)))

	return _stack
}

// AddChild adds a child to stack.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
// The function returns the following values:
//
//    - stackPage: GtkStackPage for child.
//
func (stack *Stack) AddChild(child Widgetter) *StackPage {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("add_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// AddNamed adds a child to stack.
//
// The child is identified by the name.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - name (optional) for child or NULL.
//
// The function returns the following values:
//
//    - stackPage: GtkStackPage for child.
//
func (stack *Stack) AddNamed(child Widgetter, name string) *StackPage {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if name != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_args[2]))
	}

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("add_named", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// AddTitled adds a child to stack.
//
// The child is identified by the name. The title will be used by
// GtkStackSwitcher to represent child in a tab bar, so it should be short.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - name (optional) for child.
//    - title: human-readable title for child.
//
// The function returns the following values:
//
//    - stackPage: GtkStackPage for child.
//
func (stack *Stack) AddTitled(child Widgetter, name, title string) *StackPage {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	if name != "" {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_args[2]))
	}
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_args[3]))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("add_titled", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
	runtime.KeepAlive(title)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// ChildByName finds the child with the name given as the argument.
//
// Returns NULL if there is no child with this name.
//
// The function takes the following parameters:
//
//    - name of the child to find.
//
// The function returns the following values:
//
//    - widget (optional): requested child of the GtkStack.
//
func (stack *Stack) ChildByName(name string) Widgetter {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_child_by_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)
	runtime.KeepAlive(name)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Hhomogeneous gets whether stack is horizontally homogeneous.
//
// The function returns the following values:
//
//    - ok: whether stack is horizontally homogeneous.
//
func (stack *Stack) Hhomogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_hhomogeneous", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize returns whether the Stack is set up to interpolate between
// the sizes of children on page switch.
//
// The function returns the following values:
//
//    - ok: TRUE if child sizes are interpolated.
//
func (stack *Stack) InterpolateSize() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_interpolate_size", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Page returns the GtkStackPage object for child.
//
// The function takes the following parameters:
//
//    - child of stack.
//
// The function returns the following values:
//
//    - stackPage: GtkStackPage for child.
//
func (stack *Stack) Page(child Widgetter) *StackPage {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_page", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)

	var _stackPage *StackPage // out

	_stackPage = wrapStackPage(coreglib.Take(unsafe.Pointer(_cret)))

	return _stackPage
}

// Pages returns a GListModel that contains the pages of the stack.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track and modify the visible page.
//
// The function returns the following values:
//
//    - selectionModel: GtkSelectionModel for the stack's children.
//
func (stack *Stack) Pages() *SelectionModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_pages", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _selectionModel *SelectionModel // out

	_selectionModel = wrapSelectionModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _selectionModel
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in stack will take.
//
// The function returns the following values:
//
//    - guint: transition duration.
//
func (stack *Stack) TransitionDuration() uint32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_transition_duration", _args[:], nil)
	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// TransitionRunning returns whether the stack is currently in a transition from
// one page to another.
//
// The function returns the following values:
//
//    - ok: TRUE if the transition is currently running, FALSE otherwise.
//
func (stack *Stack) TransitionRunning() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_transition_running", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Vhomogeneous gets whether stack is vertically homogeneous.
//
// The function returns the following values:
//
//    - ok: whether stack is vertically homogeneous.
//
func (stack *Stack) Vhomogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_vhomogeneous", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// VisibleChild gets the currently visible child of stack.
//
// Returns NULL if there are no visible children.
//
// The function returns the following values:
//
//    - widget (optional): visible child of the GtkStack.
//
func (stack *Stack) VisibleChild() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_visible_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// VisibleChildName returns the name of the currently visible child of stack.
//
// Returns NULL if there is no visible child.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the visible child of the GtkStack.
//
func (stack *Stack) VisibleChildName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))

	_gret := girepository.MustFind("Gtk", "Stack").InvokeMethod("get_visible_child_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stack)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Remove removes a child widget from stack.
//
// The function takes the following parameters:
//
//    - child to remove.
//
func (stack *Stack) Remove(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "Stack").InvokeMethod("remove", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)
}

// SetHhomogeneous sets the GtkStack to be horizontally homogeneous or not.
//
// If it is homogeneous, the GtkStack will request the same width for all its
// children. If it isn't, the stack may change width when a different child
// becomes visible.
//
// The function takes the following parameters:
//
//    - hhomogeneous: TRUE to make stack horizontally homogeneous.
//
func (stack *Stack) SetHhomogeneous(hhomogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	if hhomogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_hhomogeneous", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(hhomogeneous)
}

// SetInterpolateSize sets whether or not stack will interpolate its size when
// changing the visible child.
//
// If the gtk.Stack:interpolate-size property is set to TRUE, stack will
// interpolate its size between the current one and the one it'll take after
// changing the visible child, according to the set transition duration.
//
// The function takes the following parameters:
//
//    - interpolateSize: new value.
//
func (stack *Stack) SetInterpolateSize(interpolateSize bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	if interpolateSize {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_interpolate_size", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(interpolateSize)
}

// SetTransitionDuration sets the duration that transitions between pages in
// stack will take.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (stack *Stack) SetTransitionDuration(duration uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(duration)

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_transition_duration", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(duration)
}

// SetVhomogeneous sets the Stack to be vertically homogeneous or not.
//
// If it is homogeneous, the GtkStack will request the same height for all its
// children. If it isn't, the stack may change height when a different child
// becomes visible.
//
// The function takes the following parameters:
//
//    - vhomogeneous: TRUE to make stack vertically homogeneous.
//
func (stack *Stack) SetVhomogeneous(vhomogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	if vhomogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_vhomogeneous", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(vhomogeneous)
}

// SetVisibleChild makes child the visible child of stack.
//
// If child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of stack.
//
// Note that the child widget has to be visible itself (see gtk.Widget.Show())
// in order to become the visible child of stack.
//
// The function takes the following parameters:
//
//    - child of stack.
//
func (stack *Stack) SetVisibleChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_visible_child", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(child)
}

// SetVisibleChildName makes the child with the given name visible.
//
// If child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of stack.
//
// Note that the child widget has to be visible itself (see gtk.Widget.Show())
// in order to become the visible child of stack.
//
// The function takes the following parameters:
//
//    - name of the child to make visible.
//
func (stack *Stack) SetVisibleChildName(name string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stack).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "Stack").InvokeMethod("set_visible_child_name", _args[:], nil)

	runtime.KeepAlive(stack)
	runtime.KeepAlive(name)
}

// StackPage: GtkStackPage is an auxiliary class used by GtkStack.
type StackPage struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Accessible
}

var (
	_ coreglib.Objector = (*StackPage)(nil)
)

func wrapStackPage(obj *coreglib.Object) *StackPage {
	return &StackPage{
		Object: obj,
		Accessible: Accessible{
			Object: obj,
		},
	}
}

func marshalStackPage(p uintptr) (interface{}, error) {
	return wrapStackPage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child returns the stack child to which self belongs.
//
// The function returns the following values:
//
//    - widget: child to which self belongs.
//
func (self *StackPage) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StackPage").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// IconName returns the icon name of the page.
//
// The function returns the following values:
//
//    - utf8 (optional): value of the gtk.StackPage:icon-name property.
//
func (self *StackPage) IconName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StackPage").InvokeMethod("get_icon_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Name returns the name of the page.
//
// The function returns the following values:
//
//    - utf8 (optional): value of the gtk.StackPage:name property.
//
func (self *StackPage) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StackPage").InvokeMethod("get_name", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NeedsAttention returns whether the page is marked as “needs attention”.
//
// The function returns the following values:
//
//    - ok: value of the gtk.StackPage:needs-attention property.
//
func (self *StackPage) NeedsAttention() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StackPage").InvokeMethod("get_needs_attention", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the page title.
//
// The function returns the following values:
//
//    - utf8 (optional): value of the gtk.StackPage:title property.
//
func (self *StackPage) Title() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StackPage").InvokeMethod("get_title", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UseUnderline gets whether underlines in the page title indicate mnemonics.
//
// The function returns the following values:
//
//    - ok: value of the gtk.StackPage:use-underline property.
//
func (self *StackPage) UseUnderline() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StackPage").InvokeMethod("get_use_underline", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Visible returns whether page is visible in its GtkStack.
//
// This is independent from the gtk.Widget:visible property of its widget.
//
// The function returns the following values:
//
//    - ok: TRUE if page is visible.
//
func (self *StackPage) Visible() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "StackPage").InvokeMethod("get_visible", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetIconName sets the icon name of the page.
//
// The function takes the following parameters:
//
//    - setting: new value to set.
//
func (self *StackPage) SetIconName(setting string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(setting)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "StackPage").InvokeMethod("set_icon_name", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetName sets the name of the page.
//
// The function takes the following parameters:
//
//    - setting: new value to set.
//
func (self *StackPage) SetName(setting string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(setting)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "StackPage").InvokeMethod("set_name", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetNeedsAttention sets whether the page is marked as “needs attention”.
//
// The function takes the following parameters:
//
//    - setting: new value to set.
//
func (self *StackPage) SetNeedsAttention(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "StackPage").InvokeMethod("set_needs_attention", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetTitle sets the page title.
//
// The function takes the following parameters:
//
//    - setting: new value to set.
//
func (self *StackPage) SetTitle(setting string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(setting)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "StackPage").InvokeMethod("set_title", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetUseUnderline sets whether underlines in the page title indicate mnemonics.
//
// The function takes the following parameters:
//
//    - setting: new value to set.
//
func (self *StackPage) SetUseUnderline(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "StackPage").InvokeMethod("set_use_underline", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetVisible sets whether page is visible in its GtkStack.
//
// The function takes the following parameters:
//
//    - visible: new property value.
//
func (self *StackPage) SetVisible(visible bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if visible {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "StackPage").InvokeMethod("set_visible", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(visible)
}
