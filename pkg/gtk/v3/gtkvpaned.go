// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeVPaned returns the GType for the type VPaned.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeVPaned() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "VPaned").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalVPaned)
	return gtype
}

// VPanedOverrider contains methods that are overridable.
type VPanedOverrider interface {
}

// VPaned widget is a container widget with two children arranged vertically.
// The division between the two panes is adjustable by the user by dragging a
// handle. See Paned for details.
//
// GtkVPaned has been deprecated, use Paned instead.
type VPaned struct {
	_ [0]func() // equal guard
	Paned
}

var (
	_ Containerer       = (*VPaned)(nil)
	_ coreglib.Objector = (*VPaned)(nil)
)

func classInitVPanedder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapVPaned(obj *coreglib.Object) *VPaned {
	return &VPaned{
		Paned: Paned{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalVPaned(p uintptr) (interface{}, error) {
	return wrapVPaned(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVPaned: create a new VPaned
//
// Deprecated: Use gtk_paned_new() with GTK_ORIENTATION_VERTICAL instead.
//
// The function returns the following values:
//
//    - vPaned: new VPaned.
//
func NewVPaned() *VPaned {
	_gret := girepository.MustFind("Gtk", "VPaned").InvokeMethod("new_VPaned", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _vPaned *VPaned // out

	_vPaned = wrapVPaned(coreglib.Take(unsafe.Pointer(_cret)))

	return _vPaned
}
