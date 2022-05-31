// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_atk1_ComponentIface_grab_focus(AtkComponent*);
// extern gboolean _gotk4_atk1_ComponentIface_set_size(AtkComponent*, gint, gint);
// extern gdouble _gotk4_atk1_ComponentIface_get_alpha(AtkComponent*);
// extern gint _gotk4_atk1_ComponentIface_get_mdi_zorder(AtkComponent*);
// extern void _gotk4_atk1_ComponentIface_bounds_changed(AtkComponent*, AtkRectangle*);
// extern void _gotk4_atk1_ComponentIface_remove_focus_handler(AtkComponent*, guint);
// extern void _gotk4_atk1_Component_ConnectBoundsChanged(gpointer, AtkRectangle*, guintptr);
import "C"

// glib.Type values for atkcomponent.go.
var (
	GTypeScrollType = coreglib.Type(C.atk_scroll_type_get_type())
	GTypeComponent  = coreglib.Type(C.atk_component_get_type())
	GTypeRectangle  = coreglib.Type(C.atk_rectangle_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeScrollType, F: marshalScrollType},
		{T: GTypeComponent, F: marshalComponent},
		{T: GTypeRectangle, F: marshalRectangle},
	})
}

// ScrollType specifies where an object should be placed on the screen when
// using scroll_to.
type ScrollType C.gint

const (
	// ScrollTopLeft: scroll the object vertically and horizontally to bring its
	// top left corner to the top left corner of the window.
	ScrollTopLeft ScrollType = iota
	// ScrollBottomRight: scroll the object vertically and horizontally to bring
	// its bottom right corner to the bottom right corner of the window.
	ScrollBottomRight
	// ScrollTopEdge: scroll the object vertically to bring its top edge to the
	// top edge of the window.
	ScrollTopEdge
	// ScrollBottomEdge: scroll the object vertically to bring its bottom edge
	// to the bottom edge of the window.
	ScrollBottomEdge
	// ScrollLeftEdge: scroll the object vertically and horizontally to bring
	// its left edge to the left edge of the window.
	ScrollLeftEdge
	// ScrollRightEdge: scroll the object vertically and horizontally to bring
	// its right edge to the right edge of the window.
	ScrollRightEdge
	// ScrollAnywhere: scroll the object vertically and horizontally so that as
	// much as possible of the object becomes visible. The exact placement is
	// determined by the application.
	ScrollAnywhere
)

func marshalScrollType(p uintptr) (interface{}, error) {
	return ScrollType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ScrollType.
func (s ScrollType) String() string {
	switch s {
	case ScrollTopLeft:
		return "TopLeft"
	case ScrollBottomRight:
		return "BottomRight"
	case ScrollTopEdge:
		return "TopEdge"
	case ScrollBottomEdge:
		return "BottomEdge"
	case ScrollLeftEdge:
		return "LeftEdge"
	case ScrollRightEdge:
		return "RightEdge"
	case ScrollAnywhere:
		return "Anywhere"
	default:
		return fmt.Sprintf("ScrollType(%d)", s)
	}
}

// ComponentOverrider contains methods that are overridable.
type ComponentOverrider interface {
	// The function takes the following parameters:
	//
	BoundsChanged(bounds *Rectangle)
	// Alpha returns the alpha value (i.e. the opacity) for this component, on a
	// scale from 0 (fully transparent) to 1.0 (fully opaque).
	//
	// The function returns the following values:
	//
	//    - gdouble: alpha value from 0 to 1.0, inclusive.
	//
	Alpha() float64
	// MDIZOrder gets the zorder of the component. The value G_MININT will be
	// returned if the layer of the component is not ATK_LAYER_MDI or
	// ATK_LAYER_WINDOW.
	//
	// The function returns the following values:
	//
	//    - gint which is the zorder of the component, i.e. the depth at which
	//      the component is shown in relation to other components in the same
	//      container.
	//
	MDIZOrder() int32
	// GrabFocus grabs focus for this component.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if successful, FALSE otherwise.
	//
	GrabFocus() bool
	// RemoveFocusHandler: remove the handler specified by handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	//
	// Deprecated: If you need to track when an object gains or lose the focus,
	// use the Object::state-change "focused" notification instead.
	//
	// The function takes the following parameters:
	//
	//    - handlerId: handler id of the focus handler to be removed from
	//      component.
	//
	RemoveFocusHandler(handlerId uint32)
	// SetSize: set the size of the component in terms of width and height.
	//
	// The function takes the following parameters:
	//
	//    - width to set for component.
	//    - height to set for component.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE or FALSE whether the size was set or not.
	//
	SetSize(width, height int32) bool
}

// Component should be implemented by most if not all UI elements with an actual
// on-screen presence, i.e. components which can be said to have a
// screen-coordinate bounding box. Virtually all widgets will need to have
// Component implementations provided for their corresponding Object class. In
// short, only UI elements which are *not* GUI elements will omit this ATK
// interface.
//
// A possible exception might be textual information with a transparent
// background, in which case text glyph bounding box information is provided by
// Text.
//
// Component wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Component struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Component)(nil)
)

// Componenter describes Component's interface methods.
type Componenter interface {
	coreglib.Objector

	// Alpha returns the alpha value (i.e.
	Alpha() float64
	// MDIZOrder gets the zorder of the component.
	MDIZOrder() int32
	// GrabFocus grabs focus for this component.
	GrabFocus() bool
	// RemoveFocusHandler: remove the handler specified by handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	RemoveFocusHandler(handlerId uint32)
	// SetSize: set the size of the component in terms of width and height.
	SetSize(width, height int32) bool

	// Bounds-changed: 'bounds-changed" signal is emitted when the bposition or
	// size of the component changes.
	ConnectBoundsChanged(func(arg1 *Rectangle)) coreglib.SignalHandle
}

var _ Componenter = (*Component)(nil)

func ifaceInitComponenter(gifacePtr, data C.gpointer) {
	iface := (*C.AtkComponentIface)(unsafe.Pointer(gifacePtr))
	iface.bounds_changed = (*[0]byte)(C._gotk4_atk1_ComponentIface_bounds_changed)
	iface.get_alpha = (*[0]byte)(C._gotk4_atk1_ComponentIface_get_alpha)
	iface.get_mdi_zorder = (*[0]byte)(C._gotk4_atk1_ComponentIface_get_mdi_zorder)
	iface.grab_focus = (*[0]byte)(C._gotk4_atk1_ComponentIface_grab_focus)
	iface.remove_focus_handler = (*[0]byte)(C._gotk4_atk1_ComponentIface_remove_focus_handler)
	iface.set_size = (*[0]byte)(C._gotk4_atk1_ComponentIface_set_size)
}

//export _gotk4_atk1_ComponentIface_bounds_changed
func _gotk4_atk1_ComponentIface_bounds_changed(arg0 *C.AtkComponent, arg1 *C.AtkRectangle) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	var _bounds *Rectangle // out

	_bounds = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.BoundsChanged(_bounds)
}

//export _gotk4_atk1_ComponentIface_get_alpha
func _gotk4_atk1_ComponentIface_get_alpha(arg0 *C.AtkComponent) (cret C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	gdouble := iface.Alpha()

	cret = C.gdouble(gdouble)

	return cret
}

//export _gotk4_atk1_ComponentIface_get_mdi_zorder
func _gotk4_atk1_ComponentIface_get_mdi_zorder(arg0 *C.AtkComponent) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	gint := iface.MDIZOrder()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_ComponentIface_grab_focus
func _gotk4_atk1_ComponentIface_grab_focus(arg0 *C.AtkComponent) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	ok := iface.GrabFocus()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_ComponentIface_remove_focus_handler
func _gotk4_atk1_ComponentIface_remove_focus_handler(arg0 *C.AtkComponent, arg1 C.guint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	var _handlerId uint32 // out

	_handlerId = uint32(arg1)

	iface.RemoveFocusHandler(_handlerId)
}

//export _gotk4_atk1_ComponentIface_set_size
func _gotk4_atk1_ComponentIface_set_size(arg0 *C.AtkComponent, arg1 C.gint, arg2 C.gint) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ComponentOverrider)

	var _width int32  // out
	var _height int32 // out

	_width = int32(arg1)
	_height = int32(arg2)

	ok := iface.SetSize(_width, _height)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapComponent(obj *coreglib.Object) *Component {
	return &Component{
		Object: obj,
	}
}

func marshalComponent(p uintptr) (interface{}, error) {
	return wrapComponent(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_atk1_Component_ConnectBoundsChanged
func _gotk4_atk1_Component_ConnectBoundsChanged(arg0 C.gpointer, arg1 *C.AtkRectangle, arg2 C.guintptr) {
	var f func(arg1 *Rectangle)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 *Rectangle))
	}

	var _arg1 *Rectangle // out

	_arg1 = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_arg1)
}

// ConnectBoundsChanged: 'bounds-changed" signal is emitted when the bposition
// or size of the component changes.
func (component *Component) ConnectBoundsChanged(f func(arg1 *Rectangle)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(component, "bounds-changed", false, unsafe.Pointer(C._gotk4_atk1_Component_ConnectBoundsChanged), f)
}

// Alpha returns the alpha value (i.e. the opacity) for this component, on a
// scale from 0 (fully transparent) to 1.0 (fully opaque).
//
// The function returns the following values:
//
//    - gdouble: alpha value from 0 to 1.0, inclusive.
//
func (component *Component) Alpha() float64 {
	var args [1]girepository.Argument
	var _arg0 *C.void   // out
	var _cret C.gdouble // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	*(**Component)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MDIZOrder gets the zorder of the component. The value G_MININT will be
// returned if the layer of the component is not ATK_LAYER_MDI or
// ATK_LAYER_WINDOW.
//
// The function returns the following values:
//
//    - gint which is the zorder of the component, i.e. the depth at which the
//      component is shown in relation to other components in the same container.
//
func (component *Component) MDIZOrder() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	*(**Component)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// GrabFocus grabs focus for this component.
//
// The function returns the following values:
//
//    - ok: TRUE if successful, FALSE otherwise.
//
func (component *Component) GrabFocus() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	*(**Component)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveFocusHandler: remove the handler specified by handler_id from the list
// of functions to be executed when this object receives focus events (in or
// out).
//
// Deprecated: If you need to track when an object gains or lose the focus, use
// the Object::state-change "focused" notification instead.
//
// The function takes the following parameters:
//
//    - handlerId: handler id of the focus handler to be removed from component.
//
func (component *Component) RemoveFocusHandler(handlerId uint32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.guint // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	_arg1 = C.guint(handlerId)
	*(**Component)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(component)
	runtime.KeepAlive(handlerId)
}

// SetSize: set the size of the component in terms of width and height.
//
// The function takes the following parameters:
//
//    - width to set for component.
//    - height to set for component.
//
// The function returns the following values:
//
//    - ok: TRUE or FALSE whether the size was set or not.
//
func (component *Component) SetSize(width, height int32) bool {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gint     // out
	var _arg2 C.gint     // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(component).Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)
	*(**Component)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(component)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Rectangle: data structure for holding a rectangle. Those coordinates are
// relative to the component top-level parent.
//
// An instance of this type is always passed by reference.
type Rectangle struct {
	*rectangle
}

// rectangle is the struct that's finalized.
type rectangle struct {
	native *C.AtkRectangle
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Rectangle{&rectangle{(*C.AtkRectangle)(b)}}, nil
}

// NewRectangle creates a new Rectangle instance from the given
// fields.
func NewRectangle(x, y, width, height int32) Rectangle {
	var f0 C.gint // out
	f0 = C.gint(x)
	var f1 C.gint // out
	f1 = C.gint(y)
	var f2 C.gint // out
	f2 = C.gint(width)
	var f3 C.gint // out
	f3 = C.gint(height)

	v := C.AtkRectangle{
		x:      f0,
		y:      f1,
		width:  f2,
		height: f3,
	}

	return *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// X coordinate of the left side of the rectangle.
func (r *Rectangle) X() int32 {
	var v int32 // out
	v = int32(r.native.x)
	return v
}

// Y coordinate of the top side of the rectangle.
func (r *Rectangle) Y() int32 {
	var v int32 // out
	v = int32(r.native.y)
	return v
}

// Width: width of the rectangle.
func (r *Rectangle) Width() int32 {
	var v int32 // out
	v = int32(r.native.width)
	return v
}

// Height: height of the rectangle.
func (r *Rectangle) Height() int32 {
	var v int32 // out
	v = int32(r.native.height)
	return v
}
