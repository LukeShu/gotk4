// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// extern void _gotk4_gtk3_ComboBox_ConnectPopup(gpointer, guintptr);
// extern void _gotk4_gtk3_ComboBox_ConnectMoveActive(gpointer, GtkScrollType, guintptr);
// extern void _gotk4_gtk3_ComboBox_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_ComboBoxClass_changed(GtkComboBox*);
// extern gchar* _gotk4_gtk3_ComboBox_ConnectFormatEntryText(gpointer, gchar*, guintptr);
// extern gchar* _gotk4_gtk3_ComboBoxClass_format_entry_text(GtkComboBox*, gchar*);
// extern gboolean _gotk4_gtk3_TreeViewRowSeparatorFunc(GtkTreeModel*, GtkTreeIter*, gpointer);
// extern gboolean _gotk4_gtk3_ComboBox_ConnectPopdown(gpointer, guintptr);
// gchar* _gotk4_gtk3_ComboBox_virtual_format_entry_text(void* fnptr, GtkComboBox* arg0, gchar* arg1) {
//   return ((gchar* (*)(GtkComboBox*, gchar*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_ComboBox_virtual_changed(void* fnptr, GtkComboBox* arg0) {
//   ((void (*)(GtkComboBox*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeComboBox = coreglib.Type(C.gtk_combo_box_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeComboBox, F: marshalComboBox},
	})
}

// ComboBoxOverrides contains methods that are overridable.
type ComboBoxOverrides struct {
	Changed func()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	FormatEntryText func(path string) string
}

func defaultComboBoxOverrides(v *ComboBox) ComboBoxOverrides {
	return ComboBoxOverrides{
		Changed:         v.changed,
		FormatEntryText: v.formatEntryText,
	}
}

// ComboBox is a widget that allows the user to choose from a list of valid
// choices. The GtkComboBox displays the selected choice. When activated, the
// GtkComboBox displays a popup which allows the user to make a new choice. The
// style in which the selected value is displayed, and the style of the popup is
// determined by the current theme. It may be similar to a Windows-style combo
// box.
//
// The GtkComboBox uses the model-view pattern; the list of valid choices is
// specified in the form of a tree model, and the display of the choices can be
// adapted to the data in the model by using cell renderers, as you would in
// a tree view. This is possible since GtkComboBox implements the CellLayout
// interface. The tree model holding the valid choices is not restricted to
// a flat list, it can be a real tree, and the popup will reflect the tree
// structure.
//
// To allow the user to enter values not in the model, the “has-entry” property
// allows the GtkComboBox to contain a Entry. This entry can be accessed by
// calling gtk_bin_get_child() on the combo box.
//
// For a simple list of textual choices, the model-view API of GtkComboBox
// can be a bit overwhelming. In this case, ComboBoxText offers a simple
// alternative. Both GtkComboBox and ComboBoxText can contain an entry.
//
// CSS nodes
//
//    combobox
//    ├── box.linked
//    │   ├── entry.combo
//    │   ╰── button.combo
//    │       ╰── box
//    │           ╰── arrow
//    ╰── window.popup
//
// A GtkComboBox with an entry has a single CSS node with name combobox.
// It contains a box with the .linked class. That box contains an entry and a
// button, both with the .combo class added. The button also contains another
// node with name arrow.
type ComboBox struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	CellEditable
	CellLayout
}

var (
	_ Binner            = (*ComboBox)(nil)
	_ coreglib.Objector = (*ComboBox)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ComboBox, *ComboBoxClass, ComboBoxOverrides](
		GTypeComboBox,
		initComboBoxClass,
		wrapComboBox,
		defaultComboBoxOverrides,
	)
}

func initComboBoxClass(gclass unsafe.Pointer, overrides ComboBoxOverrides, classInitFunc func(*ComboBoxClass)) {
	pclass := (*C.GtkComboBoxClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeComboBox))))

	if overrides.Changed != nil {
		pclass.changed = (*[0]byte)(C._gotk4_gtk3_ComboBoxClass_changed)
	}

	if overrides.FormatEntryText != nil {
		pclass.format_entry_text = (*[0]byte)(C._gotk4_gtk3_ComboBoxClass_format_entry_text)
	}

	if classInitFunc != nil {
		class := (*ComboBoxClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapComboBox(obj *coreglib.Object) *ComboBox {
	return &ComboBox{
		Bin: Bin{
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
		},
		Object: obj,
		CellEditable: CellEditable{
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
		CellLayout: CellLayout{
			Object: obj,
		},
	}
}

func marshalComboBox(p uintptr) (interface{}, error) {
	return wrapComboBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged signal is emitted when the active item is changed. The can be
// due to the user selecting a different item from the list, or due to a call
// to gtk_combo_box_set_active_iter(). It will also be emitted while typing into
// the entry of a combo box with an entry.
func (comboBox *ComboBox) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "changed", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectChanged), f)
}

// ConnectFormatEntryText: for combo boxes that are created with an entry (See
// GtkComboBox:has-entry).
//
// A signal which allows you to change how the text displayed in a combo box's
// entry is displayed.
//
// Connect a signal handler which returns an allocated string representing path.
// That string will then be used to set the text in the combo box's entry. The
// default signal handler uses the text from the GtkComboBox::entry-text-column
// model column.
//
// Here's an example signal handler which fetches data from the model and
// displays it in the entry.
//
//    static gchar*
//    format_entry_text_callback (GtkComboBox *combo,
//                                const gchar *path,
//                                gpointer     user_data)
//    {
//      GtkTreeIter iter;
//      GtkTreeModel model;
//      gdouble      value;
//
//      model = gtk_combo_box_get_model (combo);
//
//      gtk_tree_model_get_iter_from_string (model, &iter, path);
//      gtk_tree_model_get (model, &iter,
//                          THE_DOUBLE_VALUE_COLUMN, &value,
//                          -1);
//
//      return g_strdup_printf ("g", value);
//    }.
func (comboBox *ComboBox) ConnectFormatEntryText(f func(path string) (utf8 string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "format-entry-text", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectFormatEntryText), f)
}

// ConnectMoveActive signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted to move the active selection.
func (comboBox *ComboBox) ConnectMoveActive(f func(scrollType ScrollType)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "move-active", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectMoveActive), f)
}

// ConnectPopdown signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popdown the combo box list.
//
// The default bindings for this signal are Alt+Up and Escape.
func (comboBox *ComboBox) ConnectPopdown(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "popdown", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectPopdown), f)
}

// ConnectPopup signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popup the combo box list.
//
// The default binding for this signal is Alt+Down.
func (comboBox *ComboBox) ConnectPopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "popup", false, unsafe.Pointer(C._gotk4_gtk3_ComboBox_ConnectPopup), f)
}

// NewComboBox creates a new empty ComboBox.
//
// The function returns the following values:
//
//   - comboBox: new ComboBox.
//
func NewComboBox() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithArea creates a new empty ComboBox using area to layout cells.
//
// The function takes the following parameters:
//
//   - area to use to layout cell renderers.
//
// The function returns the following values:
//
//   - comboBox: new ComboBox.
//
func NewComboBoxWithArea(area CellAreaer) *ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(coreglib.InternObject(area).Native()))

	_cret = C.gtk_combo_box_new_with_area(_arg1)
	runtime.KeepAlive(area)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithAreaAndEntry creates a new empty ComboBox with an entry.
//
// The new combo box will use area to layout cells.
//
// The function takes the following parameters:
//
//   - area to use to layout cell renderers.
//
// The function returns the following values:
//
//   - comboBox: new ComboBox.
//
func NewComboBoxWithAreaAndEntry(area CellAreaer) *ComboBox {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(coreglib.InternObject(area).Native()))

	_cret = C.gtk_combo_box_new_with_area_and_entry(_arg1)
	runtime.KeepAlive(area)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty ComboBox with an entry.
//
// The function returns the following values:
//
//   - comboBox: new ComboBox.
//
func NewComboBoxWithEntry() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new_with_entry()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModel creates a new ComboBox with the model initialized to
// model.
//
// The function takes the following parameters:
//
//   - model: TreeModel.
//
// The function returns the following values:
//
//   - comboBox: new ComboBox.
//
func NewComboBoxWithModel(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))

	_cret = C.gtk_combo_box_new_with_model(_arg1)
	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty ComboBox with an entry and
// with the model initialized to model.
//
// The function takes the following parameters:
//
//   - model: TreeModel.
//
// The function returns the following values:
//
//   - comboBox: new ComboBox.
//
func NewComboBoxWithModelAndEntry(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))

	_cret = C.gtk_combo_box_new_with_model_and_entry(_arg1)
	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// Active returns the index of the currently active item, or -1 if there’s
// no active item. If the model is a non-flat treemodel, and the active item
// is not an immediate child of the root of the tree, this function returns
// gtk_tree_path_get_indices (path)[0], where path is the TreePath of the active
// item.
//
// The function returns the following values:
//
//   - gint: integer which is the index of the currently active item, or -1 if
//     there’s no active item.
//
func (comboBox *ComboBox) Active() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_active(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ActiveID returns the ID of the active row of combo_box. This value is taken
// from the active row and the column specified by the ComboBox:id-column
// property of combo_box (see gtk_combo_box_set_id_column()).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the ComboBox:id-column property of combo_box is not set, or if no row is
// active, or if the active row has a NULL ID value, then NULL is returned.
//
// The function returns the following values:
//
//   - utf8 (optional): ID of the active row, or NULL.
//
func (comboBox *ComboBox) ActiveID() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)
	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ActiveIter sets iter to point to the currently active item, if any item is
// active. Otherwise, iter is left unchanged.
//
// The function returns the following values:
//
//   - iter: TreeIter.
//   - ok: TRUE if iter was set, FALSE otherwise.
//
func (comboBox *ComboBox) ActiveIter() (*TreeIter, bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.GtkTreeIter  // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_active_iter(_arg0, &_arg1)
	runtime.KeepAlive(comboBox)

	var _iter *TreeIter // out
	var _ok bool        // out

	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// AddTearoffs gets the current value of the :add-tearoffs property.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//   - ok: current value of the :add-tearoffs property.
//
func (comboBox *ComboBox) AddTearoffs() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_add_tearoffs(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ButtonSensitivity returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
//
// The function returns the following values:
//
//   - sensitivityType: GTK_SENSITIVITY_ON if the dropdown button is sensitive
//     when the model is empty, GTK_SENSITIVITY_OFF if the button is always
//     insensitive or GTK_SENSITIVITY_AUTO if it is only sensitive as long as
//     the model has one item to be selected.
//
func (comboBox *ComboBox) ButtonSensitivity() SensitivityType {
	var _arg0 *C.GtkComboBox       // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_button_sensitivity(_arg0)
	runtime.KeepAlive(comboBox)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// ColumnSpanColumn returns the column with column span information for
// combo_box.
//
// The function returns the following values:
//
//   - gint: column span column.
//
func (comboBox *ComboBox) ColumnSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_column_span_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// EntryTextColumn returns the column which combo_box is using to get the
// strings from to display in the internal entry.
//
// The function returns the following values:
//
//   - gint: column in the data source model of combo_box.
//
func (comboBox *ComboBox) EntryTextColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_entry_text_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// FocusOnClick returns whether the combo box grabs focus when it is clicked
// with the mouse. See gtk_combo_box_set_focus_on_click().
//
// Deprecated: Use gtk_widget_get_focus_on_click() instead.
//
// The function returns the following values:
//
//   - ok: TRUE if the combo box grabs focus when it is clicked with the mouse.
//
func (combo *ComboBox) FocusOnClick() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(combo).Native()))

	_cret = C.gtk_combo_box_get_focus_on_click(_arg0)
	runtime.KeepAlive(combo)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasEntry returns whether the combo box has an entry.
//
// The function returns the following values:
//
//   - ok: whether there is an entry in combo_box.
//
func (comboBox *ComboBox) HasEntry() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_has_entry(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IDColumn returns the column which combo_box is using to get string IDs for
// values from.
//
// The function returns the following values:
//
//   - gint: column in the data source model of combo_box.
//
func (comboBox *ComboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the TreeModel which is acting as data source for combo_box.
//
// The function returns the following values:
//
//   - treeModel which was passed during construction.
//
func (comboBox *ComboBox) Model() *TreeModel {
	var _arg0 *C.GtkComboBox  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_model(_arg0)
	runtime.KeepAlive(comboBox)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))

	return _treeModel
}

// PopupAccessible gets the accessible object corresponding to the combo box’s
// popup.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// The function returns the following values:
//
//   - object: accessible object corresponding to the combo box’s popup.
//
func (comboBox *ComboBox) PopupAccessible() *atk.AtkObject {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.AtkObject   // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_popup_accessible(_arg0)
	runtime.KeepAlive(comboBox)

	var _object *atk.AtkObject // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_object = &atk.AtkObject{
			Object: obj,
		}
	}

	return _object
}

// PopupFixedWidth gets whether the popup uses a fixed width matching the
// allocated width of the combo box.
//
// The function returns the following values:
//
//   - ok: TRUE if the popup uses a fixed width.
//
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpanColumn returns the column with row span information for combo_box.
//
// The function returns the following values:
//
//   - gint: row span column.
//
func (comboBox *ComboBox) RowSpanColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_row_span_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Title gets the current title of the menu in tearoff mode. See
// gtk_combo_box_set_add_tearoffs().
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//   - utf8 menu’s title in tearoff mode. This is an internal copy of the string
//     which must not be freed.
//
func (comboBox *ComboBox) Title() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_title(_arg0)
	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WrapWidth returns the wrap width which is used to determine the number of
// columns for the popup menu. If the wrap width is larger than 1, the combo box
// is in table mode.
//
// The function returns the following values:
//
//   - gint: wrap width.
//
func (comboBox *ComboBox) WrapWidth() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	_cret = C.gtk_combo_box_get_wrap_width(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Popdown hides the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) Popdown() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	C.gtk_combo_box_popdown(_arg0)
	runtime.KeepAlive(comboBox)
}

// Popup pops up the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, combo_box must be mapped, or nothing will happen.
func (comboBox *ComboBox) Popup() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	C.gtk_combo_box_popup(_arg0)
	runtime.KeepAlive(comboBox)
}

// PopupForDevice pops up the menu or dropdown list of combo_box, the popup
// window will be grabbed so only device and its associated pointer/keyboard are
// the only Devices able to send events to it.
//
// The function takes the following parameters:
//
//   - device: Device.
//
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(device)
}

// SetActive sets the active item of combo_box to be the item at index.
//
// The function takes the following parameters:
//
//   - index_: index in the model passed during construction, or -1 to have no
//     active item.
//
func (comboBox *ComboBox) SetActive(index_ int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(index_)

	C.gtk_combo_box_set_active(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(index_)
}

// SetActiveID changes the active row of combo_box to the one that has an
// ID equal to active_id, or unsets the active row if active_id is NULL.
// Rows having a NULL ID string cannot be made active by this function.
//
// If the ComboBox:id-column property of combo_box is unset or if no row has the
// given ID then the function does nothing and returns FALSE.
//
// The function takes the following parameters:
//
//   - activeId (optional): ID of the row to select, or NULL.
//
// The function returns the following values:
//
//   - ok: TRUE if a row with a matching ID was found. If a NULL active_id was
//     given to unset the active row, the function always returns TRUE.
//
func (comboBox *ComboBox) SetActiveID(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if activeId != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(activeId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(activeId)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by iter,
// or unsets the active item if iter is NULL.
//
// The function takes the following parameters:
//
//   - iter (optional) or NULL.
//
func (comboBox *ComboBox) SetActiveIter(iter *TreeIter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkTreeIter // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if iter != nil {
		_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}

	C.gtk_combo_box_set_active_iter(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(iter)
}

// SetAddTearoffs sets whether the popup menu should have a tearoff menu item.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//   - addTearoffs: TRUE to add tearoff menu items.
//
func (comboBox *ComboBox) SetAddTearoffs(addTearoffs bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if addTearoffs {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_add_tearoffs(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(addTearoffs)
}

// SetButtonSensitivity sets whether the dropdown button of the combo
// box should be always sensitive (GTK_SENSITIVITY_ON), never sensitive
// (GTK_SENSITIVITY_OFF) or only if there is at least one item to display
// (GTK_SENSITIVITY_AUTO).
//
// The function takes the following parameters:
//
//   - sensitivity: specify the sensitivity of the dropdown button.
//
func (comboBox *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	var _arg0 *C.GtkComboBox       // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_combo_box_set_button_sensitivity(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(sensitivity)
}

// SetColumnSpanColumn sets the column with column span information for
// combo_box to be column_span. The column span column contains integers which
// indicate how many columns an item should span.
//
// The function takes the following parameters:
//
//   - columnSpan: column in the model passed during construction.
//
func (comboBox *ComboBox) SetColumnSpanColumn(columnSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(columnSpan)

	C.gtk_combo_box_set_column_span_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(columnSpan)
}

// SetEntryTextColumn sets the model column which combo_box should use to
// get strings from to be text_column. The column text_column in the model of
// combo_box must be of type G_TYPE_STRING.
//
// This is only relevant if combo_box has been created with ComboBox:has-entry
// as TRUE.
//
// The function takes the following parameters:
//
//   - textColumn: column in model to get the strings from for the internal
//     entry.
//
func (comboBox *ComboBox) SetEntryTextColumn(textColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(textColumn)

	C.gtk_combo_box_set_entry_text_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(textColumn)
}

// SetFocusOnClick sets whether the combo box will grab focus when it is clicked
// with the mouse. Making mouse clicks not grab focus is useful in places like
// toolbars where you don’t want the keyboard focus removed from the main area
// of the application.
//
// Deprecated: Use gtk_widget_set_focus_on_click() instead.
//
// The function takes the following parameters:
//
//   - focusOnClick: whether the combo box grabs focus when clicked with the
//     mouse.
//
func (combo *ComboBox) SetFocusOnClick(focusOnClick bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(combo).Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_focus_on_click(_arg0, _arg1)
	runtime.KeepAlive(combo)
	runtime.KeepAlive(focusOnClick)
}

// SetIDColumn sets the model column which combo_box should use to get string
// IDs for values from. The column id_column in the model of combo_box must be
// of type G_TYPE_STRING.
//
// The function takes the following parameters:
//
//   - idColumn: column in model to get string IDs for values from.
//
func (comboBox *ComboBox) SetIDColumn(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(idColumn)
}

// SetModel sets the model used by combo_box to be model. Will unset a
// previously set model (if applicable). If model is NULL, then it will unset
// the model.
//
// Note that this function does not clear the cell renderers, you have to
// call gtk_cell_layout_clear() yourself if you need to set up different cell
// renderers for the new model.
//
// The function takes the following parameters:
//
//   - model (optional): TreeModel.
//
func (comboBox *ComboBox) SetModel(model TreeModeller) {
	var _arg0 *C.GtkComboBox  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if model != nil {
		_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_combo_box_set_model(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(model)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width matching the allocated width of the combo box.
//
// The function takes the following parameters:
//
//   - fixed: whether to use a fixed popup width.
//
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fixed)
}

// SetRowSeparatorFunc sets the row separator function, which is used to
// determine whether a row should be drawn as a separator. If the row separator
// function is NULL, no separators are drawn. This is the default value.
//
// The function takes the following parameters:
//
//   - fn: TreeViewRowSeparatorFunc.
//
func (comboBox *ComboBox) SetRowSeparatorFunc(fn TreeViewRowSeparatorFunc) {
	var _arg0 *C.GtkComboBox                // out
	var _arg1 C.GtkTreeViewRowSeparatorFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_TreeViewRowSeparatorFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_combo_box_set_row_separator_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fn)
}

// SetRowSpanColumn sets the column with row span information for combo_box to
// be row_span. The row span column contains integers which indicate how many
// rows an item should span.
//
// The function takes the following parameters:
//
//   - rowSpan: column in the model passed during construction.
//
func (comboBox *ComboBox) SetRowSpanColumn(rowSpan int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(rowSpan)

	C.gtk_combo_box_set_row_span_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(rowSpan)
}

// SetTitle sets the menu’s title in tearoff mode.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//   - title for the menu in tearoff mode.
//
func (comboBox *ComboBox) SetTitle(title string) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_set_title(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(title)
}

// SetWrapWidth sets the wrap width of combo_box to be width. The wrap width is
// basically the preferred number of columns when you want the popup to be layed
// out in a table.
//
// The function takes the following parameters:
//
//   - width: preferred number of columns.
//
func (comboBox *ComboBox) SetWrapWidth(width int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = C.gint(width)

	C.gtk_combo_box_set_wrap_width(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(width)
}

func (comboBox *ComboBox) changed() {
	gclass := (*C.GtkComboBoxClass)(coreglib.PeekParentClass(comboBox))
	fnarg := gclass.changed

	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))

	C._gotk4_gtk3_ComboBox_virtual_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(comboBox)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (comboBox *ComboBox) formatEntryText(path string) string {
	gclass := (*C.GtkComboBoxClass)(coreglib.PeekParentClass(comboBox))
	fnarg := gclass.format_entry_text

	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.gchar       // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gtk3_ComboBox_virtual_format_entry_text(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(path)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ComboBoxClass: instance of this type is always passed by reference.
type ComboBoxClass struct {
	*comboBoxClass
}

// comboBoxClass is the struct that's finalized.
type comboBoxClass struct {
	native *C.GtkComboBoxClass
}

// ParentClass: parent class.
func (c *ComboBoxClass) ParentClass() *BinClass {
	valptr := &c.native.parent_class
	var _v *BinClass // out
	_v = (*BinClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
