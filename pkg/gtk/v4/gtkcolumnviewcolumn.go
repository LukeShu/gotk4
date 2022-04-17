// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkcolumnviewcolumn.go.
var GTypeColumnViewColumn = externglib.Type(C.gtk_column_view_column_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeColumnViewColumn, F: marshalColumnViewColumn},
	})
}

// ColumnViewColumnOverrider contains methods that are overridable.
type ColumnViewColumnOverrider interface {
	externglib.Objector
}

// WrapColumnViewColumnOverrider wraps the ColumnViewColumnOverrider
// interface implementation to access the instance methods.
func WrapColumnViewColumnOverrider(obj ColumnViewColumnOverrider) *ColumnViewColumn {
	return wrapColumnViewColumn(externglib.BaseObject(obj))
}

// ColumnViewColumn: GtkColumnViewColumn represents the columns being added to
// GtkColumnView.
//
// The main ingredient for a GtkColumnViewColumn is the GtkListItemFactory that
// tells the columnview how to create cells for this column from items in the
// model.
//
// Columns have a title, and can optionally have a header menu set with
// gtk.ColumnViewColumn.SetHeaderMenu().
//
// A sorter can be associated with a column using
// gtk.ColumnViewColumn.SetSorter(), to let users influence sorting by clicking
// on the column header.
type ColumnViewColumn struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*ColumnViewColumn)(nil)
)

func classInitColumnViewColumner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapColumnViewColumn(obj *externglib.Object) *ColumnViewColumn {
	return &ColumnViewColumn{
		Object: obj,
	}
}

func marshalColumnViewColumn(p uintptr) (interface{}, error) {
	return wrapColumnViewColumn(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColumnViewColumn creates a new GtkColumnViewColumn that uses the given
// factory for mapping items to widgets.
//
// You most likely want to call gtk.ColumnView.AppendColumn() next.
//
// The function takes ownership of the argument, so you can write code like:
//
//    column = gtk_column_view_column_new (_("Name"),
//      gtk_builder_list_item_factory_new_from_resource ("/name.ui"));.
//
// The function takes the following parameters:
//
//    - title (optional): title to use for this column.
//    - factory (optional) to populate items with.
//
// The function returns the following values:
//
//    - columnViewColumn: new GtkColumnViewColumn using the given factory.
//
func NewColumnViewColumn(title string, factory *ListItemFactory) *ColumnViewColumn {
	var _arg1 *C.char                // out
	var _arg2 *C.GtkListItemFactory  // out
	var _cret *C.GtkColumnViewColumn // in

	if title != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if factory != nil {
		_arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(externglib.InternObject(factory).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(factory).Native()))
	}

	_cret = C.gtk_column_view_column_new(_arg1, _arg2)
	runtime.KeepAlive(title)
	runtime.KeepAlive(factory)

	var _columnViewColumn *ColumnViewColumn // out

	_columnViewColumn = wrapColumnViewColumn(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _columnViewColumn
}

// ColumnView gets the column view that's currently displaying this column.
//
// If self has not been added to a column view yet, NULL is returned.
//
// The function returns the following values:
//
//    - columnView (optional): column view displaying self.
//
func (self *ColumnViewColumn) ColumnView() *ColumnView {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret *C.GtkColumnView       // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_column_view(_arg0)
	runtime.KeepAlive(self)

	var _columnView *ColumnView // out

	if _cret != nil {
		_columnView = wrapColumnView(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _columnView
}

// Expand returns whether this column should expand.
//
// The function returns the following values:
//
//    - ok: TRUE if this column expands.
//
func (self *ColumnViewColumn) Expand() bool {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_expand(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Factory gets the factory that's currently used to populate list items for
// this column.
//
// The function returns the following values:
//
//    - listItemFactory (optional): factory in use.
//
func (self *ColumnViewColumn) Factory() *ListItemFactory {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret *C.GtkListItemFactory  // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_factory(_arg0)
	runtime.KeepAlive(self)

	var _listItemFactory *ListItemFactory // out

	if _cret != nil {
		_listItemFactory = wrapListItemFactory(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listItemFactory
}

// FixedWidth gets the fixed width of the column.
//
// The function returns the following values:
//
//    - gint: fixed with of the column.
//
func (self *ColumnViewColumn) FixedWidth() int {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret C.int                  // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_fixed_width(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HeaderMenu gets the menu model that is used to create the context menu for
// the column header.
//
// The function returns the following values:
//
//    - menuModel (optional) or NULL.
//
func (self *ColumnViewColumn) HeaderMenu() gio.MenuModeller {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret *C.GMenuModel          // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_header_menu(_arg0)
	runtime.KeepAlive(self)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.MenuModeller)
				return ok
			})
			rv, ok := casted.(gio.MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// Resizable returns whether this column is resizable.
//
// The function returns the following values:
//
//    - ok: TRUE if this column is resizable.
//
func (self *ColumnViewColumn) Resizable() bool {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_resizable(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sorter returns the sorter that is associated with the column.
//
// The function returns the following values:
//
//    - sorter (optional): GtkSorter of self.
//
func (self *ColumnViewColumn) Sorter() *Sorter {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret *C.GtkSorter           // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_sorter(_arg0)
	runtime.KeepAlive(self)

	var _sorter *Sorter // out

	if _cret != nil {
		_sorter = wrapSorter(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _sorter
}

// Title returns the title set with gtk_column_view_column_set_title().
//
// The function returns the following values:
//
//    - utf8 (optional) column's title.
//
func (self *ColumnViewColumn) Title() string {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Visible returns whether this column is visible.
//
// The function returns the following values:
//
//    - ok: TRUE if this column is visible.
//
func (self *ColumnViewColumn) Visible() bool {
	var _arg0 *C.GtkColumnViewColumn // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_column_view_column_get_visible(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpand sets the column to take available extra space.
//
// The extra space is shared equally amongst all columns that have the expand
// set to TRUE.
//
// The function takes the following parameters:
//
//    - expand: TRUE if this column should expand to fill available sace.
//
func (self *ColumnViewColumn) SetExpand(expand bool) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if expand {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_column_set_expand(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expand)
}

// SetFactory sets the GtkListItemFactory to use for populating list items for
// this column.
//
// The function takes the following parameters:
//
//    - factory (optional) to use or NULL for none.
//
func (self *ColumnViewColumn) SetFactory(factory *ListItemFactory) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.GtkListItemFactory  // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(externglib.InternObject(factory).Native()))
	}

	C.gtk_column_view_column_set_factory(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetFixedWidth: if fixed_width is not -1, sets the fixed width of column;
// otherwise unsets it.
//
// Setting a fixed width overrides the automatically calculated width.
// Interactive resizing also sets the “fixed-width” property.
//
// The function takes the following parameters:
//
//    - fixedWidth: new fixed width, or -1.
//
func (self *ColumnViewColumn) SetFixedWidth(fixedWidth int) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.int                  // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.int(fixedWidth)

	C.gtk_column_view_column_set_fixed_width(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(fixedWidth)
}

// SetHeaderMenu sets the menu model that is used to create the context menu for
// the column header.
//
// The function takes the following parameters:
//
//    - menu (optional): GMenuModel, or NULL.
//
func (self *ColumnViewColumn) SetHeaderMenu(menu gio.MenuModeller) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.GMenuModel          // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if menu != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(externglib.InternObject(menu).Native()))
	}

	C.gtk_column_view_column_set_header_menu(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(menu)
}

// SetResizable sets whether this column should be resizable by dragging.
//
// The function takes the following parameters:
//
//    - resizable: whether this column should be resizable.
//
func (self *ColumnViewColumn) SetResizable(resizable bool) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if resizable {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_column_set_resizable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(resizable)
}

// SetSorter associates a sorter with the column.
//
// If sorter is NULL, the column will not let users change the sorting by
// clicking on its header.
//
// This sorter can be made active by clicking on the column header, or by
// calling gtk.ColumnView.SortByColumn().
//
// See gtk.ColumnView.GetSorter() for the necessary steps for setting up
// customizable sorting for gtk.ColumnView.
//
// The function takes the following parameters:
//
//    - sorter (optional): GtkSorter to associate with column.
//
func (self *ColumnViewColumn) SetSorter(sorter *Sorter) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.GtkSorter           // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if sorter != nil {
		_arg1 = (*C.GtkSorter)(unsafe.Pointer(externglib.InternObject(sorter).Native()))
	}

	C.gtk_column_view_column_set_sorter(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sorter)
}

// SetTitle sets the title of this column.
//
// The title is displayed in the header of a GtkColumnView for this column and
// is therefore user-facing text that should be translated.
//
// The function takes the following parameters:
//
//    - title (optional): title to use for this column.
//
func (self *ColumnViewColumn) SetTitle(title string) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if title != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_column_view_column_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetVisible sets whether this column should be visible in views.
//
// The function takes the following parameters:
//
//    - visible: whether this column should be visible.
//
func (self *ColumnViewColumn) SetVisible(visible bool) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_column_set_visible(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(visible)
}
