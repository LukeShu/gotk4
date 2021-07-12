// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_expander_get_type()), F: marshalTreeExpanderer},
	})
}

// TreeExpanderer describes TreeExpander's methods.
type TreeExpanderer interface {
	// Child gets the child widget displayed by @self.
	Child() *Widget
	// Item forwards the item set on the `GtkTreeListRow` that @self is
	// managing.
	Item() *externglib.Object
	// ListRow gets the list row managed by @self.
	ListRow() *TreeListRow
	// SetChild sets the content widget to display.
	SetChild(child Widgeter)
	// SetListRow sets the tree list row that this expander should manage.
	SetListRow(listRow TreeListRower)
}

// TreeExpander: `GtkTreeExpander` is a widget that provides an expander for a
// list.
//
// It is typically placed as a bottommost child into a `GtkListView` to allow
// users to expand and collapse children in a list with a
// [class@Gtk.TreeListModel]. `GtkTreeExpander` provides the common UI elements,
// gestures and keybindings for this purpose.
//
// On top of this, the "listitem.expand", "listitem.collapse" and
// "listitem.toggle-expand" actions are provided to allow adding custom UI for
// managing expanded state.
//
// The `GtkTreeListModel` must be set to not be passthrough. Then it will
// provide [class@Gtk.TreeListRow] items which can be set via
// [method@Gtk.TreeExpander.set_list_row] on the expander. The expander will
// then watch that row item automatically. [method@Gtk.TreeExpander.set_child]
// sets the widget that displays the actual row contents.
//
//
// CSS nodes
//
// “` treeexpander ├── [indent]* ├── [expander] ╰── <child> “`
//
// `GtkTreeExpander` has zero or one CSS nodes with the name "expander" that
// should display the expander icon. The node will be `:checked` when it is
// expanded. If the node is not expandable, an "indent" node will be displayed
// instead.
//
// For every level of depth, another "indent" node is prepended.
//
//
// Accessibility
//
// `GtkTreeExpander` uses the GTK_ACCESSIBLE_ROLE_GROUP role. The expander icon
// is represented as a GTK_ACCESSIBLE_ROLE_BUTTON, labelled by the expander's
// child, and toggling it will change the GTK_ACCESSIBLE_STATE_EXPANDED state.
type TreeExpander struct {
	Widget
}

var (
	_ TreeExpanderer  = (*TreeExpander)(nil)
	_ gextras.Nativer = (*TreeExpander)(nil)
)

func wrapTreeExpander(obj *externglib.Object) TreeExpanderer {
	return &TreeExpander{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalTreeExpanderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeExpander(obj), nil
}

// NewTreeExpander creates a new `GtkTreeExpander`
func NewTreeExpander() *TreeExpander {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_tree_expander_new()

	var _treeExpander *TreeExpander // out

	_treeExpander = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TreeExpander)

	return _treeExpander
}

// Child gets the child widget displayed by @self.
func (self *TreeExpander) Child() *Widget {
	var _arg0 *C.GtkTreeExpander // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_expander_get_child(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// Item forwards the item set on the `GtkTreeListRow` that @self is managing.
//
// This call is essentially equivalent to calling:
//
// “`c gtk_tree_list_row_get_item (gtk_tree_expander_get_list_row (@self)); “`
func (self *TreeExpander) Item() *externglib.Object {
	var _arg0 *C.GtkTreeExpander // out
	var _cret C.gpointer         // in

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_expander_get_item(_arg0)

	var _object *externglib.Object // out

	_object = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*externglib.Object)

	return _object
}

// ListRow gets the list row managed by @self.
func (self *TreeExpander) ListRow() *TreeListRow {
	var _arg0 *C.GtkTreeExpander // out
	var _cret *C.GtkTreeListRow  // in

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_expander_get_list_row(_arg0)

	var _treeListRow *TreeListRow // out

	_treeListRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TreeListRow)

	return _treeListRow
}

// SetChild sets the content widget to display.
func (self *TreeExpander) SetChild(child Widgeter) {
	var _arg0 *C.GtkTreeExpander // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_tree_expander_set_child(_arg0, _arg1)
}

// SetListRow sets the tree list row that this expander should manage.
func (self *TreeExpander) SetListRow(listRow TreeListRower) {
	var _arg0 *C.GtkTreeExpander // out
	var _arg1 *C.GtkTreeListRow  // out

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkTreeListRow)(unsafe.Pointer((listRow).(gextras.Nativer).Native()))

	C.gtk_tree_expander_set_list_row(_arg0, _arg1)
}
