// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_at_context_get_type()), F: marshalATContext},
	})
}

// ATContext: `GtkATContext` is an abstract class provided by GTK to communicate
// to platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a `GtkATContext` subclass, and is
// responsible for updating the accessible state in response to state changes in
// `GtkAccessible`.
type ATContext interface {
	gextras.Objector

	// Accessible retrieves the `GtkAccessible` using this context.
	Accessible() Accessible
	// AccessibleRole retrieves the accessible role of this context.
	AccessibleRole() AccessibleRole
}

// ATContextClass implements the ATContext interface.
type ATContextClass struct {
	*externglib.Object
}

var _ ATContext = (*ATContextClass)(nil)

func wrapATContext(obj *externglib.Object) ATContext {
	return &ATContextClass{
		Object: obj,
	}
}

func marshalATContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapATContext(obj), nil
}

// NewATContextCreate creates a new `GtkATContext` instance for the given
// accessible role, accessible instance, and display connection.
//
// The `GtkATContext` implementation being instantiated will depend on the
// platform.
func NewATContextCreate(accessibleRole AccessibleRole, accessible Accessible, display gdk.Display) ATContext {
	var _arg1 C.GtkAccessibleRole // out
	var _arg2 *C.GtkAccessible    // out
	var _arg3 *C.GdkDisplay       // out
	var _cret *C.GtkATContext     // in

	_arg1 = C.GtkAccessibleRole(accessibleRole)
	_arg2 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	_arg3 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gtk_at_context_create(_arg1, _arg2, _arg3)

	var _atContext ATContext // out

	_atContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(ATContext)

	return _atContext
}

// Accessible retrieves the `GtkAccessible` using this context.
func (s *ATContextClass) Accessible() Accessible {
	var _arg0 *C.GtkATContext  // out
	var _cret *C.GtkAccessible // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_at_context_get_accessible(_arg0)

	var _accessible Accessible // out

	_accessible = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Accessible)

	return _accessible
}

// AccessibleRole retrieves the accessible role of this context.
func (s *ATContextClass) AccessibleRole() AccessibleRole {
	var _arg0 *C.GtkATContext     // out
	var _cret C.GtkAccessibleRole // in

	_arg0 = (*C.GtkATContext)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_at_context_get_accessible_role(_arg0)

	var _accessibleRole AccessibleRole // out

	_accessibleRole = AccessibleRole(_cret)

	return _accessibleRole
}
