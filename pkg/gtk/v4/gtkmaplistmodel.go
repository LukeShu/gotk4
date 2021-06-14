// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_map_list_model_get_type()), F: marshalMapListModel},
	})
}

// MapListModel: a `GtkMapListModel` maps the items in a list model to different
// items.
//
// `GtkMapListModel` uses a [callback@Gtk.MapListModelMapFunc].
//
// Example: Create a list of `GtkEventControllers` “`c static gpointer
// map_to_controllers (gpointer widget, gpointer data) { gpointer result =
// gtk_widget_observe_controllers (widget); g_object_unref (widget); return
// result; }
//
// widgets = gtk_widget_observe_children (widget);
//
// controllers = gtk_map_list_model_new (G_TYPE_LIST_MODEL, widgets,
// map_to_controllers, NULL, NULL);
//
// model = gtk_flatten_list_model_new (GTK_TYPE_EVENT_CONTROLLER, controllers);
// “`
//
// `GtkMapListModel` will attempt to discard the mapped objects as soon as they
// are no longer needed and recreate them if necessary.
type MapListModel interface {
	gextras.Objector

	// HasMap checks if a map function is currently set on @self.
	HasMap() bool
}

// mapListModel implements the MapListModel class.
type mapListModel struct {
	gextras.Objector
}

var _ MapListModel = (*mapListModel)(nil)

// WrapMapListModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapMapListModel(obj *externglib.Object) MapListModel {
	return mapListModel{
		Objector: obj,
	}
}

func marshalMapListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMapListModel(obj), nil
}

// HasMap checks if a map function is currently set on @self.
func (s mapListModel) HasMap() bool {
	var _arg0 *C.GtkMapListModel // out

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_map_list_model_has_map(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
