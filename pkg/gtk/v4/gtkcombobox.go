// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern char* _gotk4_gtk4_ComboBoxClass_format_entry_text(GtkComboBox*, char*);
// extern gboolean _gotk4_gtk4_ComboBox_ConnectPopdown(gpointer, guintptr);
// extern gchar* _gotk4_gtk4_ComboBox_ConnectFormatEntryText(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk4_ComboBoxClass_changed(GtkComboBox*);
// extern void _gotk4_gtk4_ComboBox_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk4_ComboBox_ConnectPopup(gpointer, guintptr);
import "C"

// glib.Type values for gtkcombobox.go.
var GTypeComboBox = coreglib.Type(C.gtk_combo_box_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeComboBox, F: marshalComboBox},
	})
}

// ComboBoxOverrider contains methods that are overridable.
type ComboBoxOverrider interface {
	Changed()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	FormatEntryText(path string) string
}

// ComboBox: GtkComboBox is a widget that allows the user to choose from a list
// of valid choices.
//
// !An example GtkComboBox (combo-box.png)
//
// The GtkComboBox displays the selected choice; when activated, the GtkComboBox
// displays a popup which allows the user to make a new choice.
//
// The GtkComboBox uses the model-view pattern; the list of valid choices is
// specified in the form of a tree model, and the display of the choices can be
// adapted to the data in the model by using cell renderers, as you would in a
// tree view. This is possible since GtkComboBox implements the gtk.CellLayout
// interface. The tree model holding the valid choices is not restricted to a
// flat list, it can be a real tree, and the popup will reflect the tree
// structure.
//
// To allow the user to enter values not in the model, the
// gtk.ComboBox:has-entry property allows the GtkComboBox to contain a
// gtk.Entry. This entry can be accessed by calling gtk.ComboBox.GetChild() on
// the combo box.
//
// For a simple list of textual choices, the model-view API of GtkComboBox can
// be a bit overwhelming. In this case, gtk.ComboBoxText offers a simple
// alternative. Both GtkComboBox and GtkComboBoxText can contain an entry.
//
// CSS nodes
//
//    combobox
//    ├── box.linked
//    │   ╰── button.combo
//    │       ╰── box
//    │           ├── cellview
//    │           ╰── arrow
//    ╰── window.popup
//
//
// A normal combobox contains a box with the .linked class, a button with the
// .combo class and inside those buttons, there are a cellview and an arrow.
//
//    combobox
//    ├── box.linked
//    │   ├── entry.combo
//    │   ╰── button.combo
//    │       ╰── box
//    │           ╰── arrow
//    ╰── window.popup
//
//
// A GtkComboBox with an entry has a single CSS node with name combobox. It
// contains a box with the .linked class. That box contains an entry and a
// button, both with the .combo class added. The button also contains another
// node with name arrow.
//
//
// Accessibility
//
// GtkComboBox uses the GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type ComboBox struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	CellEditable
	CellLayout
}

var (
	_ Widgetter         = (*ComboBox)(nil)
	_ coreglib.Objector = (*ComboBox)(nil)
)

func classInitComboBoxer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkComboBoxClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkComboBoxClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Changed() }); ok {
		pclass.changed = (*[0]byte)(C._gotk4_gtk4_ComboBoxClass_changed)
	}

	if _, ok := goval.(interface{ FormatEntryText(path string) string }); ok {
		pclass.format_entry_text = (*[0]byte)(C._gotk4_gtk4_ComboBoxClass_format_entry_text)
	}
}

//export _gotk4_gtk4_ComboBoxClass_changed
func _gotk4_gtk4_ComboBoxClass_changed(arg0 *C.GtkComboBox) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Changed() })

	iface.Changed()
}

//export _gotk4_gtk4_ComboBoxClass_format_entry_text
func _gotk4_gtk4_ComboBoxClass_format_entry_text(arg0 *C.GtkComboBox, arg1 *C.char) (cret *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ FormatEntryText(path string) string })

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := iface.FormatEntryText(_path)

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

func wrapComboBox(obj *coreglib.Object) *ComboBox {
	return &ComboBox{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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
		Object: obj,
		CellEditable: CellEditable{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
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
		},
		CellLayout: CellLayout{
			Object: obj,
		},
	}
}

func marshalComboBox(p uintptr) (interface{}, error) {
	return wrapComboBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ComboBox_ConnectChanged
func _gotk4_gtk4_ComboBox_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged is emitted when the active item is changed.
//
// The can be due to the user selecting a different item from the list, or due
// to a call to gtk.ComboBox.SetActiveIter(). It will also be emitted while
// typing into the entry of a combo box with an entry.
func (comboBox *ComboBox) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "changed", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectChanged), f)
}

//export _gotk4_gtk4_ComboBox_ConnectFormatEntryText
func _gotk4_gtk4_ComboBox_ConnectFormatEntryText(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) (cret *C.gchar) {
	var f func(path string) (utf8 string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path string) (utf8 string))
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := f(_path)

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

// ConnectFormatEntryText is emitted to allow changing how the text in a combo
// box's entry is displayed.
//
// See gtk.ComboBox:has-entry.
//
// Connect a signal handler which returns an allocated string representing path.
// That string will then be used to set the text in the combo box's entry. The
// default signal handler uses the text from the gtk.ComboBox:entry-text-column
// model column.
//
// Here's an example signal handler which fetches data from the model and
// displays it in the entry.
//
//    static char *
//    format_entry_text_callback (GtkComboBox *combo,
//                                const char *path,
//                                gpointer     user_data)
//    {
//      GtkTreeIter iter;
//      GtkTreeModel model;
//      double       value;
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
	return coreglib.ConnectGeneratedClosure(comboBox, "format-entry-text", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectFormatEntryText), f)
}

//export _gotk4_gtk4_ComboBox_ConnectPopdown
func _gotk4_gtk4_ComboBox_ConnectPopdown(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectPopdown is emitted to popdown the combo box list.
//
// This is an keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are Alt+Up and Escape.
func (comboBox *ComboBox) ConnectPopdown(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "popdown", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectPopdown), f)
}

//export _gotk4_gtk4_ComboBox_ConnectPopup
func _gotk4_gtk4_ComboBox_ConnectPopup(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPopup is emitted to popup the combo box list.
//
// This is an keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Alt+Down.
func (comboBox *ComboBox) ConnectPopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(comboBox, "popup", false, unsafe.Pointer(C._gotk4_gtk4_ComboBox_ConnectPopup), f)
}

// NewComboBox creates a new empty GtkComboBox.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBox() *ComboBox {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty GtkComboBox with an entry.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBoxWithEntry() *ComboBox {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_entry", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModel creates a new GtkComboBox with a model.
//
// The function takes the following parameters:
//
//    - model: GtkTreeModel.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBoxWithModel(model TreeModeller) *ComboBox {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	*(*TreeModeller)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_model", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty GtkComboBox with an entry
// and a model.
//
// The function takes the following parameters:
//
//    - model: GtkTreeModel.
//
// The function returns the following values:
//
//    - comboBox: new GtkComboBox.
//
func NewComboBoxWithModelAndEntry(model TreeModeller) *ComboBox {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	*(*TreeModeller)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("new_ComboBox_with_model_and_entry", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// ActiveID returns the ID of the active row of combo_box.
//
// This value is taken from the active row and the column specified by the
// gtk.ComboBox:id-column property of combo_box (see
// gtk.ComboBox.SetIDColumn()).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the gtk.ComboBox:id-column property of combo_box is not set, or if no row
// is active, or if the active row has a NULL ID value, then NULL is returned.
//
// The function returns the following values:
//
//    - utf8 (optional): ID of the active row, or NULL.
//
func (comboBox *ComboBox) ActiveID() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_active_id", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Child gets the child widget of combo_box.
//
// The function returns the following values:
//
//    - widget (optional): child widget of combo_box.
//
func (comboBox *ComboBox) Child() Widgetter {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_child", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
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

// HasEntry returns whether the combo box has an entry.
//
// The function returns the following values:
//
//    - ok: whether there is an entry in combo_box.
//
func (comboBox *ComboBox) HasEntry() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_has_entry", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model returns the GtkTreeModel of combo_box.
//
// The function returns the following values:
//
//    - treeModel (optional): GtkTreeModel which was passed during construction.
//
func (comboBox *ComboBox) Model() *TreeModel {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_model", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _treeModel *TreeModel // out

	if _cret != nil {
		_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _treeModel
}

// PopupFixedWidth gets whether the popup uses a fixed width.
//
// The function returns the following values:
//
//    - ok: TRUE if the popup uses a fixed width.
//
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("get_popup_fixed_width", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown hides the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) Popdown() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("popdown", args[:], nil)

	runtime.KeepAlive(comboBox)
}

// Popup pops up the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, combo_box must be mapped, or nothing will happen.
func (comboBox *ComboBox) Popup() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("popup", args[:], nil)

	runtime.KeepAlive(comboBox)
}

// PopupForDevice pops up the menu of combo_box.
//
// Note that currently this does not do anything with the device, as it was
// previously only used for list-mode combo boxes, and those were removed in GTK
// 4. However, it is retained in case similar functionality is added back later.
//
// The function takes the following parameters:
//
//    - device: GdkDevice.
//
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("popup_for_device", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(device)
}

// SetActiveID changes the active row of combo_box to the one that has an ID
// equal to active_id.
//
// If active_id is NULL, the active row is unset. Rows having a NULL ID string
// cannot be made active by this function.
//
// If the gtk.ComboBox:id-column property of combo_box is unset or if no row has
// the given ID then the function does nothing and returns FALSE.
//
// The function takes the following parameters:
//
//    - activeId (optional): ID of the row to select, or NULL.
//
// The function returns the following values:
//
//    - ok: TRUE if a row with a matching ID was found. If a NULL active_id was
//      given to unset the active row, the function always returns TRUE.
//
func (comboBox *ComboBox) SetActiveID(activeId string) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if activeId != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(activeId)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_active_id", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(activeId)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by iter.
//
// If iter is NULL, the active item is unset.
//
// The function takes the following parameters:
//
//    - iter (optional): GtkTreeIter, or NULL.
//
func (comboBox *ComboBox) SetActiveIter(iter *TreeIter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if iter != nil {
		_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_active_iter", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(iter)
}

// SetChild sets the child widget of combo_box.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (comboBox *ComboBox) SetChild(child Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if child != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_child", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(child)
}

// SetModel sets the model used by combo_box to be model.
//
// Will unset a previously set model (if applicable). If model is NULL, then it
// will unset the model.
//
// Note that this function does not clear the cell renderers, you have to call
// gtk.CellLayout.Clear() yourself if you need to set up different cell
// renderers for the new model.
//
// The function takes the following parameters:
//
//    - model (optional): GtkTreeModel.
//
func (comboBox *ComboBox) SetModel(model TreeModeller) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if model != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_model", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(model)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width.
//
// If fixed is TRUE, the popup's width is set to match the allocated width of
// the combo box.
//
// The function takes the following parameters:
//
//    - fixed: whether to use a fixed popup width.
//
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(comboBox).Native()))
	if fixed {
		_arg1 = C.TRUE
	}
	*(**ComboBox)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "ComboBox").InvokeMethod("set_popup_fixed_width", args[:], nil)

	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fixed)
}
