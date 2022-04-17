// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_HandleBoxClass_child_attached(GtkHandleBox*, GtkWidget*);
// extern void _gotk4_gtk3_HandleBoxClass_child_detached(GtkHandleBox*, GtkWidget*);
// extern void _gotk4_gtk3_HandleBox_ConnectChildAttached(gpointer, GtkWidget*, guintptr);
// extern void _gotk4_gtk3_HandleBox_ConnectChildDetached(gpointer, GtkWidget*, guintptr);
import "C"

// glib.Type values for gtkhandlebox.go.
var GTypeHandleBox = externglib.Type(C.gtk_handle_box_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeHandleBox, F: marshalHandleBox},
	})
}

// HandleBoxOverrider contains methods that are overridable.
type HandleBoxOverrider interface {
	externglib.Objector
	// The function takes the following parameters:
	//
	ChildAttached(child Widgetter)
	// The function takes the following parameters:
	//
	ChildDetached(child Widgetter)
}

// WrapHandleBoxOverrider wraps the HandleBoxOverrider
// interface implementation to access the instance methods.
func WrapHandleBoxOverrider(obj HandleBoxOverrider) *HandleBox {
	return wrapHandleBox(externglib.BaseObject(obj))
}

// HandleBox widget allows a portion of a window to be "torn off". It is a bin
// widget which displays its child and a handle that the user can drag to tear
// off a separate window (the “float window”) containing the child widget. A
// thin “ghost” is drawn in the original location of the handlebox. By dragging
// the separate window back to its original location, it can be reattached.
//
// When reattaching, the ghost and float window, must be aligned along one of
// the edges, the “snap edge”. This either can be specified by the application
// programmer explicitly, or GTK+ will pick a reasonable default based on the
// handle position.
//
// To make detaching and reattaching the handlebox as minimally confusing as
// possible to the user, it is important to set the snap edge so that the snap
// edge does not move when the handlebox is deattached. For instance, if the
// handlebox is packed at the bottom of a VBox, then when the handlebox is
// detached, the bottom edge of the handlebox's allocation will remain fixed as
// the height of the handlebox shrinks, so the snap edge should be set to
// GTK_POS_BOTTOM.
//
// > HandleBox has been deprecated. It is very specialized, lacks features > to
// make it useful and most importantly does not fit well into modern >
// application design. Do not use it. There is no replacement.
type HandleBox struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*HandleBox)(nil)
)

func classInitHandleBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkHandleBoxClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkHandleBoxClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ChildAttached(child Widgetter) }); ok {
		pclass.child_attached = (*[0]byte)(C._gotk4_gtk3_HandleBoxClass_child_attached)
	}

	if _, ok := goval.(interface{ ChildDetached(child Widgetter) }); ok {
		pclass.child_detached = (*[0]byte)(C._gotk4_gtk3_HandleBoxClass_child_detached)
	}
}

//export _gotk4_gtk3_HandleBoxClass_child_attached
func _gotk4_gtk3_HandleBoxClass_child_attached(arg0 *C.GtkHandleBox, arg1 *C.GtkWidget) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ChildAttached(child Widgetter) })

	var _child Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}

	iface.ChildAttached(_child)
}

//export _gotk4_gtk3_HandleBoxClass_child_detached
func _gotk4_gtk3_HandleBoxClass_child_detached(arg0 *C.GtkHandleBox, arg1 *C.GtkWidget) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ChildDetached(child Widgetter) })

	var _child Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}

	iface.ChildDetached(_child)
}

func wrapHandleBox(obj *externglib.Object) *HandleBox {
	return &HandleBox{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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

func marshalHandleBox(p uintptr) (interface{}, error) {
	return wrapHandleBox(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_HandleBox_ConnectChildAttached
func _gotk4_gtk3_HandleBox_ConnectChildAttached(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guintptr) {
	var f func(widget Widgetter)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	f(_widget)
}

// ConnectChildAttached: this signal is emitted when the contents of the
// handlebox are reattached to the main window.
func (handleBox *HandleBox) ConnectChildAttached(f func(widget Widgetter)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(handleBox, "child-attached", false, unsafe.Pointer(C._gotk4_gtk3_HandleBox_ConnectChildAttached), f)
}

//export _gotk4_gtk3_HandleBox_ConnectChildDetached
func _gotk4_gtk3_HandleBox_ConnectChildDetached(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guintptr) {
	var f func(widget Widgetter)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	f(_widget)
}

// ConnectChildDetached: this signal is emitted when the contents of the
// handlebox are detached from the main window.
func (handleBox *HandleBox) ConnectChildDetached(f func(widget Widgetter)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(handleBox, "child-detached", false, unsafe.Pointer(C._gotk4_gtk3_HandleBox_ConnectChildDetached), f)
}

// NewHandleBox: create a new handle box.
//
// Deprecated: HandleBox has been deprecated.
//
// The function returns the following values:
//
//    - handleBox: new HandleBox.
//
func NewHandleBox() *HandleBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_handle_box_new()

	var _handleBox *HandleBox // out

	_handleBox = wrapHandleBox(externglib.Take(unsafe.Pointer(_cret)))

	return _handleBox
}

// ChildDetached: whether the handlebox’s child is currently detached.
//
// Deprecated: HandleBox has been deprecated.
//
// The function returns the following values:
//
//    - ok: TRUE if the child is currently detached, otherwise FALSE.
//
func (handleBox *HandleBox) ChildDetached() bool {
	var _arg0 *C.GtkHandleBox // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(externglib.InternObject(handleBox).Native()))

	_cret = C.gtk_handle_box_get_child_detached(_arg0)
	runtime.KeepAlive(handleBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HandlePosition gets the handle position of the handle box. See
// gtk_handle_box_set_handle_position().
//
// Deprecated: HandleBox has been deprecated.
//
// The function returns the following values:
//
//    - positionType: current handle position.
//
func (handleBox *HandleBox) HandlePosition() PositionType {
	var _arg0 *C.GtkHandleBox   // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(externglib.InternObject(handleBox).Native()))

	_cret = C.gtk_handle_box_get_handle_position(_arg0)
	runtime.KeepAlive(handleBox)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// ShadowType gets the type of shadow drawn around the handle box. See
// gtk_handle_box_set_shadow_type().
//
// Deprecated: HandleBox has been deprecated.
//
// The function returns the following values:
//
//    - shadowType: type of shadow currently drawn around the handle box.
//
func (handleBox *HandleBox) ShadowType() ShadowType {
	var _arg0 *C.GtkHandleBox // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(externglib.InternObject(handleBox).Native()))

	_cret = C.gtk_handle_box_get_shadow_type(_arg0)
	runtime.KeepAlive(handleBox)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

// SnapEdge gets the edge used for determining reattachment of the handle box.
// See gtk_handle_box_set_snap_edge().
//
// Deprecated: HandleBox has been deprecated.
//
// The function returns the following values:
//
//    - positionType: edge used for determining reattachment, or
//      (GtkPositionType)-1 if this is determined (as per default) from the
//      handle position.
//
func (handleBox *HandleBox) SnapEdge() PositionType {
	var _arg0 *C.GtkHandleBox   // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(externglib.InternObject(handleBox).Native()))

	_cret = C.gtk_handle_box_get_snap_edge(_arg0)
	runtime.KeepAlive(handleBox)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// SetHandlePosition sets the side of the handlebox where the handle is drawn.
//
// Deprecated: HandleBox has been deprecated.
//
// The function takes the following parameters:
//
//    - position: side of the handlebox where the handle should be drawn.
//
func (handleBox *HandleBox) SetHandlePosition(position PositionType) {
	var _arg0 *C.GtkHandleBox   // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(externglib.InternObject(handleBox).Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_handle_box_set_handle_position(_arg0, _arg1)
	runtime.KeepAlive(handleBox)
	runtime.KeepAlive(position)
}

// SetShadowType sets the type of shadow to be drawn around the border of the
// handle box.
//
// Deprecated: HandleBox has been deprecated.
//
// The function takes the following parameters:
//
//    - typ: shadow type.
//
func (handleBox *HandleBox) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkHandleBox // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(externglib.InternObject(handleBox).Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_handle_box_set_shadow_type(_arg0, _arg1)
	runtime.KeepAlive(handleBox)
	runtime.KeepAlive(typ)
}

// SetSnapEdge sets the snap edge of a handlebox. The snap edge is the edge of
// the detached child that must be aligned with the corresponding edge of the
// “ghost” left behind when the child was detached to reattach the torn-off
// window. Usually, the snap edge should be chosen so that it stays in the same
// place on the screen when the handlebox is torn off.
//
// If the snap edge is not set, then an appropriate value will be guessed from
// the handle position. If the handle position is GTK_POS_RIGHT or GTK_POS_LEFT,
// then the snap edge will be GTK_POS_TOP, otherwise it will be GTK_POS_LEFT.
//
// Deprecated: HandleBox has been deprecated.
//
// The function takes the following parameters:
//
//    - edge: snap edge, or -1 to unset the value; in which case GTK+ will try to
//      guess an appropriate value in the future.
//
func (handleBox *HandleBox) SetSnapEdge(edge PositionType) {
	var _arg0 *C.GtkHandleBox   // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(externglib.InternObject(handleBox).Native()))
	_arg1 = C.GtkPositionType(edge)

	C.gtk_handle_box_set_snap_edge(_arg0, _arg1)
	runtime.KeepAlive(handleBox)
	runtime.KeepAlive(edge)
}
