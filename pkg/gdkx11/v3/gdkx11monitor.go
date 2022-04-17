// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gdkx11monitor.go.
var GTypeX11Monitor = externglib.Type(C.gdk_x11_monitor_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeX11Monitor, F: marshalX11Monitor},
	})
}

// X11MonitorOverrider contains methods that are overridable.
type X11MonitorOverrider interface {
	externglib.Objector
}

// WrapX11MonitorOverrider wraps the X11MonitorOverrider
// interface implementation to access the instance methods.
func WrapX11MonitorOverrider(obj X11MonitorOverrider) *X11Monitor {
	return wrapX11Monitor(externglib.BaseObject(obj))
}

type X11Monitor struct {
	_ [0]func() // equal guard
	gdk.Monitor
}

var (
	_ externglib.Objector = (*X11Monitor)(nil)
)

func classInitX11Monitorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11Monitor(obj *externglib.Object) *X11Monitor {
	return &X11Monitor{
		Monitor: gdk.Monitor{
			Object: obj,
		},
	}
}

func marshalX11Monitor(p uintptr) (interface{}, error) {
	return wrapX11Monitor(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
