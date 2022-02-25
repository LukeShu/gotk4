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
// extern void _gotk4_gtk4_PasswordEntry_ConnectActivate(gpointer, guintptr);
import "C"

// glib.Type values for gtkpasswordentry.go.
var GTypePasswordEntry = externglib.Type(C.gtk_password_entry_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePasswordEntry, F: marshalPasswordEntry},
	})
}

// PasswordEntryOverrider contains methods that are overridable.
type PasswordEntryOverrider interface {
}

// PasswordEntry: GtkPasswordEntry is an entry that has been tailored for
// entering secrets.
//
// !An example GtkPasswordEntry (password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, GtkPasswordEntry will also place the text in a
// non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// GtkPasswordEntry provides only minimal API and should be used with the
// gtk.Editable API.
//
// CSS Nodes
//
//    entry.password
//    ╰── text
//        ├── image.caps-lock-indicator
//        ┊
//
//
// GtkPasswordEntry has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// GtkPasswordEntry uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry struct {
	_ [0]func() // equal guard
	Widget

	*externglib.Object
	Editable
}

var (
	_ Widgetter           = (*PasswordEntry)(nil)
	_ externglib.Objector = (*PasswordEntry)(nil)
)

func classInitPasswordEntrier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPasswordEntry(obj *externglib.Object) *PasswordEntry {
	return &PasswordEntry{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
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
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalPasswordEntry(p uintptr) (interface{}, error) {
	return wrapPasswordEntry(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_PasswordEntry_ConnectActivate
func _gotk4_gtk4_PasswordEntry_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectActivate: emitted when the entry is activated.
//
// The keybindings for this signal are all forms of the Enter key.
func (entry *PasswordEntry) ConnectActivate(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(entry, "activate", false, unsafe.Pointer(C._gotk4_gtk4_PasswordEntry_ConnectActivate), f)
}

// NewPasswordEntry creates a GtkPasswordEntry.
//
// The function returns the following values:
//
//    - passwordEntry: new GtkPasswordEntry.
//
func NewPasswordEntry() *PasswordEntry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_password_entry_new()

	var _passwordEntry *PasswordEntry // out

	_passwordEntry = wrapPasswordEntry(externglib.Take(unsafe.Pointer(_cret)))

	return _passwordEntry
}

// ExtraMenu gets the menu model set with gtk_password_entry_set_extra_menu().
//
// The function returns the following values:
//
//    - menuModel: (nullable): the menu model.
//
func (entry *PasswordEntry) ExtraMenu() gio.MenuModeller {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret *C.GMenuModel       // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(externglib.InternObject(entry).Native()))

	_cret = C.gtk_password_entry_get_extra_menu(_arg0)
	runtime.KeepAlive(entry)

	var _menuModel gio.MenuModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.MenuModeller is nil")
		}

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

	return _menuModel
}

// ShowPeekIcon returns whether the entry is showing an icon to reveal the
// contents.
//
// The function returns the following values:
//
//    - ok: TRUE if an icon is shown.
//
func (entry *PasswordEntry) ShowPeekIcon() bool {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(externglib.InternObject(entry).Native()))

	_cret = C.gtk_password_entry_get_show_peek_icon(_arg0)
	runtime.KeepAlive(entry)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// entry.
//
// The function takes the following parameters:
//
//    - model (optional): GMenuModel.
//
func (entry *PasswordEntry) SetExtraMenu(model gio.MenuModeller) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 *C.GMenuModel       // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(externglib.InternObject(entry).Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(externglib.InternObject(model).Native()))
	}

	C.gtk_password_entry_set_extra_menu(_arg0, _arg1)
	runtime.KeepAlive(entry)
	runtime.KeepAlive(model)
}

// SetShowPeekIcon sets whether the entry should have a clickable icon to reveal
// the contents.
//
// Setting this to FALSE also hides the text again.
//
// The function takes the following parameters:
//
//    - showPeekIcon: whether to show the peek icon.
//
func (entry *PasswordEntry) SetShowPeekIcon(showPeekIcon bool) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(externglib.InternObject(entry).Native()))
	if showPeekIcon {
		_arg1 = C.TRUE
	}

	C.gtk_password_entry_set_show_peek_icon(_arg0, _arg1)
	runtime.KeepAlive(entry)
	runtime.KeepAlive(showPeekIcon)
}
