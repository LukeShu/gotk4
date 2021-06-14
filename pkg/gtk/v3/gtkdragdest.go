// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_dest_defaults_get_type()), F: marshalDestDefaults},
	})
}

// DestDefaults: the DestDefaults enumeration specifies the various types of
// action that will be taken on behalf of the user for a drag destination site.
type DestDefaults int

const (
	// DestDefaultsMotion: if set for a widget, GTK+, during a drag over this
	// widget will check if the drag matches this widget’s list of possible
	// targets and actions. GTK+ will then call gdk_drag_status() as
	// appropriate.
	DestDefaultsMotion DestDefaults = 1
	// DestDefaultsHighlight: if set for a widget, GTK+ will draw a highlight on
	// this widget as long as a drag is over this widget and the widget drag
	// format and action are acceptable.
	DestDefaultsHighlight DestDefaults = 2
	// DestDefaultsDrop: if set for a widget, when a drop occurs, GTK+ will will
	// check if the drag matches this widget’s list of possible targets and
	// actions. If so, GTK+ will call gtk_drag_get_data() on behalf of the
	// widget. Whether or not the drop is successful, GTK+ will call
	// gtk_drag_finish(). If the action was a move, then if the drag was
	// successful, then true will be passed for the @delete parameter to
	// gtk_drag_finish().
	DestDefaultsDrop DestDefaults = 4
	// DestDefaultsAll: if set, specifies that all default actions should be
	// taken.
	DestDefaultsAll DestDefaults = 7
)

func marshalDestDefaults(p uintptr) (interface{}, error) {
	return DestDefaults(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}
