// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_accessible_get_type()), F: marshalAccessibler},
	})
}

// AccessibleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AccessibleOverrider interface {
	// ConnectWidgetDestroyed: this function specifies the callback function to
	// be called when the widget corresponding to a GtkAccessible is destroyed.
	//
	// Deprecated: Use gtk_accessible_set_widget() and its vfuncs.
	ConnectWidgetDestroyed()
	WidgetSet()
	WidgetUnset()
}

// Accessible class is the base class for accessible implementations for Widget
// subclasses. It is a thin wrapper around Object, which adds facilities for
// associating a widget with its accessible object.
//
// An accessible implementation for a third-party widget should derive from
// Accessible and implement the suitable interfaces from ATK, such as Text or
// Selection. To establish the connection between the widget class and its
// corresponding acccessible implementation, override the get_accessible vfunc
// in WidgetClass.
type Accessible struct {
	_ [0]func() // equal guard
	atk.ObjectClass
}

var (
	_ externglib.Objector = (*Accessible)(nil)
)

func wrapAccessible(obj *externglib.Object) *Accessible {
	return &Accessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
	}
}

func marshalAccessibler(p uintptr) (interface{}, error) {
	return wrapAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectWidgetDestroyed: this function specifies the callback function to be
// called when the widget corresponding to a GtkAccessible is destroyed.
//
// Deprecated: Use gtk_accessible_set_widget() and its vfuncs.
func (accessible *Accessible) ConnectWidgetDestroyed() {
	var _arg0 *C.GtkAccessible // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))

	C.gtk_accessible_connect_widget_destroyed(_arg0)
	runtime.KeepAlive(accessible)
}

// Widget gets the Widget corresponding to the Accessible. The returned widget
// does not have a reference added, so you do not need to unref it.
//
// The function returns the following values:
//
//    - widget (optional): pointer to the Widget corresponding to the Accessible,
//      or NULL.
//
func (accessible *Accessible) Widget() Widgetter {
	var _arg0 *C.GtkAccessible // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))

	_cret = C.gtk_accessible_get_widget(_arg0)
	runtime.KeepAlive(accessible)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
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
//    - widget (optional) or NULL to unset.
//
func (accessible *Accessible) SetWidget(widget Widgetter) {
	var _arg0 *C.GtkAccessible // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_accessible_set_widget(_arg0, _arg1)
	runtime.KeepAlive(accessible)
	runtime.KeepAlive(widget)
}
