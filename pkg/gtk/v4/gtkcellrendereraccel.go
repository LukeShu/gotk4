// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_cell_renderer_accel_mode_get_type()), F: marshalCellRendererAccelMode},
		{T: externglib.Type(C.gtk_cell_renderer_accel_get_type()), F: marshalCellRendererAccel},
	})
}

// CellRendererAccelMode determines if the edited accelerators are GTK
// accelerators. If they are, consumed modifiers are suppressed, only
// accelerators accepted by GTK are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode int

const (
	// GTK: GTK accelerators mode
	CellRendererAccelModeGTK CellRendererAccelMode = iota
	// Other: other accelerator mode
	CellRendererAccelModeOther
)

func marshalCellRendererAccelMode(p uintptr) (interface{}, error) {
	return CellRendererAccelMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CellRendererAccel renders a keyboard accelerator in a cell
//
// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like `Control + a`). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
type CellRendererAccel interface {
	gextras.Objector

	privateCellRendererAccelClass()
}

// CellRendererAccelClass implements the CellRendererAccel interface.
type CellRendererAccelClass struct {
	CellRendererTextClass
}

var _ CellRendererAccel = (*CellRendererAccelClass)(nil)

func wrapCellRendererAccel(obj *externglib.Object) CellRendererAccel {
	return &CellRendererAccelClass{
		CellRendererTextClass: CellRendererTextClass{
			CellRendererClass: CellRendererClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			},
		},
	}
}

func marshalCellRendererAccel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererAccel(obj), nil
}

// NewCellRendererAccel creates a new CellRendererAccel.
func NewCellRendererAccel() CellRendererAccel {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_accel_new()

	var _cellRendererAccel CellRendererAccel // out

	_cellRendererAccel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRendererAccel)

	return _cellRendererAccel
}

func (*CellRendererAccelClass) privateCellRendererAccelClass() {}
