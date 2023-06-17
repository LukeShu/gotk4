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
// extern void _gotk4_gtk3_AccessibleClass_widget_unset(GtkAccessible*);
// extern void _gotk4_gtk3_AccessibleClass_widget_set(GtkAccessible*);
// extern void _gotk4_gtk3_AccessibleClass_connect_widget_destroyed(GtkAccessible*);
// void _gotk4_gtk3_Accessible_virtual_connect_widget_destroyed(void* fnptr, GtkAccessible* arg0) {
//   ((void (*)(GtkAccessible*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_Accessible_virtual_widget_set(void* fnptr, GtkAccessible* arg0) {
//   ((void (*)(GtkAccessible*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_Accessible_virtual_widget_unset(void* fnptr, GtkAccessible* arg0) {
//   ((void (*)(GtkAccessible*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeAccessible = coreglib.Type(C.gtk_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAccessible, F: marshalAccessible},
	})
}

// AccessibleOverrides contains methods that are overridable.
type AccessibleOverrides struct {
	// ConnectWidgetDestroyed: this function specifies the callback function to
	// be called when the widget corresponding to a GtkAccessible is destroyed.
	//
	// Deprecated: Use gtk_accessible_set_widget() and its vfuncs.
	ConnectWidgetDestroyed func()
	WidgetSet              func()
	WidgetUnset            func()
}

func defaultAccessibleOverrides(v *Accessible) AccessibleOverrides {
	return AccessibleOverrides{
		ConnectWidgetDestroyed: v.connectWidgetDestroyed,
		WidgetSet:              v.widgetSet,
		WidgetUnset:            v.widgetUnset,
	}
}

// Accessible class is the base class for accessible implementations for Widget
// subclasses. It is a thin wrapper around Object, which adds facilities for
// associating a widget with its accessible object.
//
// An accessible implementation for a third-party widget should derive from
// Accessible and implement the suitable interfaces from ATK, such as Text
// or Selection. To establish the connection between the widget class and its
// corresponding acccessible implementation, override the get_accessible vfunc
// in WidgetClass.
type Accessible struct {
	_ [0]func() // equal guard
	atk.AtkObject
}

var (
	_ coreglib.Objector = (*Accessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Accessible, *AccessibleClass, AccessibleOverrides](
		GTypeAccessible,
		initAccessibleClass,
		wrapAccessible,
		defaultAccessibleOverrides,
	)
}

func initAccessibleClass(gclass unsafe.Pointer, overrides AccessibleOverrides, classInitFunc func(*AccessibleClass)) {
	pclass := (*C.GtkAccessibleClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeAccessible))))

	if overrides.ConnectWidgetDestroyed != nil {
		pclass.connect_widget_destroyed = (*[0]byte)(C._gotk4_gtk3_AccessibleClass_connect_widget_destroyed)
	}

	if overrides.WidgetSet != nil {
		pclass.widget_set = (*[0]byte)(C._gotk4_gtk3_AccessibleClass_widget_set)
	}

	if overrides.WidgetUnset != nil {
		pclass.widget_unset = (*[0]byte)(C._gotk4_gtk3_AccessibleClass_widget_unset)
	}

	if classInitFunc != nil {
		class := (*AccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAccessible(obj *coreglib.Object) *Accessible {
	return &Accessible{
		AtkObject: atk.AtkObject{
			Object: obj,
		},
	}
}

func marshalAccessible(p uintptr) (interface{}, error) {
	return wrapAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectWidgetDestroyed: this function specifies the callback function to be
// called when the widget corresponding to a GtkAccessible is destroyed.
//
// Deprecated: Use gtk_accessible_set_widget() and its vfuncs.
func (accessible *Accessible) ConnectWidgetDestroyed() {
	var _arg0 *C.GtkAccessible // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))

	C.gtk_accessible_connect_widget_destroyed(_arg0)
	runtime.KeepAlive(accessible)
}

// Widget gets the Widget corresponding to the Accessible. The returned widget
// does not have a reference added, so you do not need to unref it.
//
// The function returns the following values:
//
//   - widget (optional): pointer to the Widget corresponding to the Accessible,
//     or NULL.
//
func (accessible *Accessible) Widget() Widgetter {
	var _arg0 *C.GtkAccessible // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))

	_cret = C.gtk_accessible_get_widget(_arg0)
	runtime.KeepAlive(accessible)

	var _widget Widgetter // out

	if _cret != nil {
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

// SetWidget sets the Widget corresponding to the Accessible.
//
// accessible will not hold a reference to widget. It is the caller’s
// responsibility to ensure that when widget is destroyed, the widget is unset
// by calling this function again with widget set to NULL.
//
// The function takes the following parameters:
//
//   - widget (optional) or NULL to unset.
//
func (accessible *Accessible) SetWidget(widget Widgetter) {
	var _arg0 *C.GtkAccessible // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	C.gtk_accessible_set_widget(_arg0, _arg1)
	runtime.KeepAlive(accessible)
	runtime.KeepAlive(widget)
}

// connectWidgetDestroyed: this function specifies the callback function to be
// called when the widget corresponding to a GtkAccessible is destroyed.
//
// Deprecated: Use gtk_accessible_set_widget() and its vfuncs.
func (accessible *Accessible) connectWidgetDestroyed() {
	gclass := (*C.GtkAccessibleClass)(coreglib.PeekParentClass(accessible))
	fnarg := gclass.connect_widget_destroyed

	var _arg0 *C.GtkAccessible // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))

	C._gotk4_gtk3_Accessible_virtual_connect_widget_destroyed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(accessible)
}

func (accessible *Accessible) widgetSet() {
	gclass := (*C.GtkAccessibleClass)(coreglib.PeekParentClass(accessible))
	fnarg := gclass.widget_set

	var _arg0 *C.GtkAccessible // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))

	C._gotk4_gtk3_Accessible_virtual_widget_set(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(accessible)
}

func (accessible *Accessible) widgetUnset() {
	gclass := (*C.GtkAccessibleClass)(coreglib.PeekParentClass(accessible))
	fnarg := gclass.widget_unset

	var _arg0 *C.GtkAccessible // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))

	C._gotk4_gtk3_Accessible_virtual_widget_unset(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(accessible)
}

// AccessibleClass: instance of this type is always passed by reference.
type AccessibleClass struct {
	*accessibleClass
}

// accessibleClass is the struct that's finalized.
type accessibleClass struct {
	native *C.GtkAccessibleClass
}

func (a *AccessibleClass) ParentClass() *atk.ObjectClass {
	valptr := &a.native.parent_class
	var _v *atk.ObjectClass // out
	_v = (*atk.ObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
