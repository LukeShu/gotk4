// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
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
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperationer},
	})
}

// MountOperation: this should not be accessed directly. Use the accessor
// functions below.
type MountOperation struct {
	gio.MountOperation
}

func wrapMountOperation(obj *externglib.Object) *MountOperation {
	return &MountOperation{
		MountOperation: gio.MountOperation{
			Object: obj,
		},
	}
}

func marshalMountOperationer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMountOperation(obj), nil
}

// NewMountOperation creates a new MountOperation
func NewMountOperation(parent *Window) *MountOperation {
	var _arg1 *C.GtkWindow       // out
	var _cret *C.GMountOperation // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_mount_operation_new(_arg1)

	var _mountOperation *MountOperation // out

	_mountOperation = wrapMountOperation(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mountOperation
}

// Parent gets the transient parent used by the MountOperation
func (op *MountOperation) Parent() *Window {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GtkWindow         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_mount_operation_get_parent(_arg0)

	var _window *Window // out

	_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _window
}

// Screen gets the screen on which windows of the MountOperation will be shown.
func (op *MountOperation) Screen() *gdk.Screen {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GdkScreen         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_mount_operation_get_screen(_arg0)

	var _screen *gdk.Screen // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_screen = &gdk.Screen{
			Object: obj,
		}
	}

	return _screen
}

// IsShowing returns whether the MountOperation is currently displaying a
// window.
func (op *MountOperation) IsShowing() bool {
	var _arg0 *C.GtkMountOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_mount_operation_is_showing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetParent sets the transient parent for windows shown by the MountOperation.
func (op *MountOperation) SetParent(parent *Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
}

// SetScreen sets the screen to show windows of the MountOperation on.
func (op *MountOperation) SetScreen(screen *gdk.Screen) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GdkScreen         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_mount_operation_set_screen(_arg0, _arg1)
}
