// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperation: this should not be accessed directly. Use the accessor
// functions below.
type MountOperation interface {
	MountOperation

	// Parent gets the transient parent used by the MountOperation
	Parent() Window
	// IsShowing returns whether the MountOperation is currently displaying a
	// window.
	IsShowing() bool
	// SetParent sets the transient parent for windows shown by the
	// MountOperation.
	SetParent(parent Window)
}

// mountOperation implements the MountOperation class.
type mountOperation struct {
	MountOperation
}

var _ MountOperation = (*mountOperation)(nil)

// WrapMountOperation wraps a GObject to the right type. It is
// primarily used internally.
func WrapMountOperation(obj *externglib.Object) MountOperation {
	return mountOperation{
		MountOperation: WrapMountOperation(obj),
	}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMountOperation(obj), nil
}

// NewMountOperation constructs a class MountOperation.
func NewMountOperation(parent Window) MountOperation {
	var _arg1 *C.GtkWindow // out

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	var _cret C.GtkMountOperation // in

	_cret = C.gtk_mount_operation_new(_arg1)

	var _mountOperation MountOperation // out

	_mountOperation = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MountOperation)

	return _mountOperation
}

// Parent gets the transient parent used by the MountOperation
func (o mountOperation) Parent() Window {
	var _arg0 *C.GtkMountOperation // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var _cret *C.GtkWindow // in

	_cret = C.gtk_mount_operation_get_parent(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Window)

	return _window
}

// IsShowing returns whether the MountOperation is currently displaying a
// window.
func (o mountOperation) IsShowing() bool {
	var _arg0 *C.GtkMountOperation // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_mount_operation_is_showing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetParent sets the transient parent for windows shown by the
// MountOperation.
func (o mountOperation) SetParent(parent Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
}
