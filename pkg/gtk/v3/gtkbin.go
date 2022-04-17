// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkbin.go.
var GTypeBin = externglib.Type(C.gtk_bin_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeBin, F: marshalBin},
	})
}

// BinOverrider contains methods that are overridable.
type BinOverrider interface {
	externglib.Objector
}

// WrapBinOverrider wraps the BinOverrider
// interface implementation to access the instance methods.
func WrapBinOverrider(obj BinOverrider) *Bin {
	return wrapBin(externglib.BaseObject(obj))
}

// Bin widget is a container with just one child. It is not very useful itself,
// but it is useful for deriving subclasses, since it provides common code
// needed for handling a single child widget.
//
// Many GTK+ widgets are subclasses of Bin, including Window, Button, Frame,
// HandleBox or ScrolledWindow.
type Bin struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*Bin)(nil)
)

// Binner describes types inherited from class Bin.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Binner interface {
	externglib.Objector
	baseBin() *Bin
}

var _ Binner = (*Bin)(nil)

func classInitBinner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapBin(obj *externglib.Object) *Bin {
	return &Bin{
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
	}
}

func marshalBin(p uintptr) (interface{}, error) {
	return wrapBin(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (bin *Bin) baseBin() *Bin {
	return bin
}

// BaseBin returns the underlying base object from the
// interface.
func BaseBin(obj Binner) *Bin {
	return obj.baseBin()
}

// Child gets the child of the Bin, or NULL if the bin contains no child widget.
// The returned widget does not have a reference added, so you do not need to
// unref it.
//
// The function returns the following values:
//
//    - widget (optional): child of bin, or NULL if it does not have a child.
//
func (bin *Bin) Child() Widgetter {
	var _arg0 *C.GtkBin    // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkBin)(unsafe.Pointer(externglib.InternObject(bin).Native()))

	_cret = C.gtk_bin_get_child(_arg0)
	runtime.KeepAlive(bin)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _widget
}
