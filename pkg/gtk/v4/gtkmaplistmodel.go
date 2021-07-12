// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_map_list_model_get_type()), F: marshalMapListModeler},
	})
}

// MapListModelMapFunc: user function that is called to map an @item of the
// original model to an item expected by the map model.
//
// The returned items must conform to the item type of the model they are used
// with.
type MapListModelMapFunc func(item *externglib.Object, userData cgo.Handle) (object *externglib.Object)

//export gotk4_MapListModelMapFunc
func gotk4_MapListModelMapFunc(arg0 C.gpointer, arg1 C.gpointer) (cret C.gpointer) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out
	var userData cgo.Handle     // out

	item = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(arg0)))).(*externglib.Object)
	userData = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(MapListModelMapFunc)
	object := fn(item, userData)

	cret = C.gpointer(unsafe.Pointer(object.Native()))

	return cret
}

// MapListModeler describes MapListModel's methods.
type MapListModeler interface {
	// Model gets the model that is currently being mapped or nil if none.
	Model() *gio.ListModel
	// HasMap checks if a map function is currently set on @self.
	HasMap() bool
	// SetModel sets the model to be mapped.
	SetModel(model gio.ListModeler)
}

// MapListModel: `GtkMapListModel` maps the items in a list model to different
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
type MapListModel struct {
	*externglib.Object

	gio.ListModel
}

var (
	_ MapListModeler  = (*MapListModel)(nil)
	_ gextras.Nativer = (*MapListModel)(nil)
)

func wrapMapListModel(obj *externglib.Object) MapListModeler {
	return &MapListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalMapListModeler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMapListModel(obj), nil
}

// Model gets the model that is currently being mapped or nil if none.
func (self *MapListModel) Model() *gio.ListModel {
	var _arg0 *C.GtkMapListModel // out
	var _cret *C.GListModel      // in

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_map_list_model_get_model(_arg0)

	var _listModel *gio.ListModel // out

	_listModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.ListModel)

	return _listModel
}

// HasMap checks if a map function is currently set on @self.
func (self *MapListModel) HasMap() bool {
	var _arg0 *C.GtkMapListModel // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_map_list_model_has_map(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetModel sets the model to be mapped.
//
// GTK makes no effort to ensure that @model conforms to the item type expected
// by the map function. It assumes that the caller knows what they are doing and
// have set up an appropriate map function.
func (self *MapListModel) SetModel(model gio.ListModeler) {
	var _arg0 *C.GtkMapListModel // out
	var _arg1 *C.GListModel      // out

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_map_list_model_set_model(_arg0, _arg1)
}
