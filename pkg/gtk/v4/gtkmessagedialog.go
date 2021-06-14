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
		{T: externglib.Type(C.gtk_buttons_type_get_type()), F: marshalButtonsType},
		{T: externglib.Type(C.gtk_message_dialog_get_type()), F: marshalMessageDialog},
	})
}

// ButtonsType: prebuilt sets of buttons for `GtkDialog`.
//
// If none of these choices are appropriate, simply use GTK_BUTTONS_NONE and
// call [method@Gtk.Dialog.add_buttons].
//
// > Please note that GTK_BUTTONS_OK, GTK_BUTTONS_YES_NO > and
// GTK_BUTTONS_OK_CANCEL are discouraged by the > GNOME Human Interface
// Guidelines (http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType int

const (
	// ButtonsTypeNone: no buttons at all
	ButtonsTypeNone ButtonsType = 0
	// ButtonsTypeOk: an OK button
	ButtonsTypeOk ButtonsType = 1
	// ButtonsTypeClose: a Close button
	ButtonsTypeClose ButtonsType = 2
	// ButtonsTypeCancel: a Cancel button
	ButtonsTypeCancel ButtonsType = 3
	// ButtonsTypeYesNo yes and No buttons
	ButtonsTypeYesNo ButtonsType = 4
	// ButtonsTypeOkCancel: OK and Cancel buttons
	ButtonsTypeOkCancel ButtonsType = 5
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	return ButtonsType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// MessageDialog: `GtkMessageDialog` presents a dialog with some message text.
//
// !An example GtkMessageDialog (messagedialog.png)
//
// It’s simply a convenience widget; you could construct the equivalent of
// `GtkMessageDialog` from `GtkDialog` without too much effort, but
// `GtkMessageDialog` saves typing.
//
// The easiest way to do a modal message dialog is to use the GTK_DIALOG_MODAL
// flag, which will call [method@Gtk.Window.set_modal] internally. The dialog
// will prevent interaction with the parent window until it's hidden or
// destroyed. You can use the [signal@Gtk.Dialog::response] signal to know when
// the user dismissed the dialog.
//
// An example for using a modal dialog: “`c GtkDialogFlags flags =
// GTK_DIALOG_DESTROY_WITH_PARENT | GTK_DIALOG_MODAL; dialog =
// gtk_message_dialog_new (parent_window, flags, GTK_MESSAGE_ERROR,
// GTK_BUTTONS_CLOSE, "Error reading “s”: s", filename, g_strerror (errno)); //
// Destroy the dialog when the user responds to it // (e.g. clicks a button)
//
// g_signal_connect (dialog, "response", G_CALLBACK (gtk_window_destroy), NULL);
// “`
//
// You might do a non-modal `GtkMessageDialog` simply by omitting the
// GTK_DIALOG_MODAL flag:
//
// “`c GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT; dialog =
// gtk_message_dialog_new (parent_window, flags, GTK_MESSAGE_ERROR,
// GTK_BUTTONS_CLOSE, "Error reading “s”: s", filename, g_strerror (errno));
//
// // Destroy the dialog when the user responds to it // (e.g. clicks a button)
// // g_signal_connect (dialog, "response", G_CALLBACK (gtk_window_destroy), NULL);
// // “`
//
//
// GtkMessageDialog as GtkBuildable
//
// The `GtkMessageDialog` implementation of the `GtkBuildable` interface exposes
// the message area as an internal child with the name “message_area”.
type MessageDialog interface {
	Dialog
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager

	// MessageArea returns the message area of the dialog.
	//
	// This is the box where the dialog’s primary and secondary labels are
	// packed. You can add your own extra content to that box and it will appear
	// below those labels. See [method@Gtk.Dialog.get_content_area] for the
	// corresponding function in the parent [class@Gtk.Dialog].
	MessageArea() Widget
	// SetMarkup sets the text of the message dialog.
	SetMarkup(str string)
}

// messageDialog implements the MessageDialog class.
type messageDialog struct {
	Dialog
	Accessible
	Buildable
	ConstraintTarget
	Native
	Root
	ShortcutManager
}

var _ MessageDialog = (*messageDialog)(nil)

// WrapMessageDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapMessageDialog(obj *externglib.Object) MessageDialog {
	return messageDialog{
		Dialog:           WrapDialog(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Native:           WrapNative(obj),
		Root:             WrapRoot(obj),
		ShortcutManager:  WrapShortcutManager(obj),
	}
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMessageDialog(obj), nil
}

// MessageArea returns the message area of the dialog.
//
// This is the box where the dialog’s primary and secondary labels are
// packed. You can add your own extra content to that box and it will appear
// below those labels. See [method@Gtk.Dialog.get_content_area] for the
// corresponding function in the parent [class@Gtk.Dialog].
func (m messageDialog) MessageArea() Widget {
	var _arg0 *C.GtkMessageDialog // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(m.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_message_dialog_get_message_area(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// SetMarkup sets the text of the message dialog.
func (m messageDialog) SetMarkup(str string) {
	var _arg0 *C.GtkMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_message_dialog_set_markup(_arg0, _arg1)
}
