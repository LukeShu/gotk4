// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_response_type_get_type()), F: marshalResponseType},
		{T: externglib.Type(C.gtk_dialog_flags_get_type()), F: marshalDialogFlags},
		{T: externglib.Type(C.gtk_dialog_get_type()), F: marshalDialog},
	})
}

// ResponseType: predefined values for use as response ids in
// gtk_dialog_add_button(). All predefined values are negative; GTK+ leaves
// values of 0 or greater for application-defined response ids.
type ResponseType int

const (
	// ResponseTypeNone: returned if an action widget has no response id, or if
	// the dialog gets programmatically hidden or destroyed
	ResponseTypeNone ResponseType = -1
	// ResponseTypeReject: generic response id, not used by GTK+ dialogs
	ResponseTypeReject ResponseType = -2
	// ResponseTypeAccept: generic response id, not used by GTK+ dialogs
	ResponseTypeAccept ResponseType = -3
	// ResponseTypeDeleteEvent: returned if the dialog is deleted
	ResponseTypeDeleteEvent ResponseType = -4
	// ResponseTypeOk: returned by OK buttons in GTK+ dialogs
	ResponseTypeOk ResponseType = -5
	// ResponseTypeCancel: returned by Cancel buttons in GTK+ dialogs
	ResponseTypeCancel ResponseType = -6
	// ResponseTypeClose: returned by Close buttons in GTK+ dialogs
	ResponseTypeClose ResponseType = -7
	// ResponseTypeYes: returned by Yes buttons in GTK+ dialogs
	ResponseTypeYes ResponseType = -8
	// ResponseTypeNo: returned by No buttons in GTK+ dialogs
	ResponseTypeNo ResponseType = -9
	// ResponseTypeApply: returned by Apply buttons in GTK+ dialogs
	ResponseTypeApply ResponseType = -10
	// ResponseTypeHelp: returned by Help buttons in GTK+ dialogs
	ResponseTypeHelp ResponseType = -11
)

func marshalResponseType(p uintptr) (interface{}, error) {
	return ResponseType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DialogFlags flags used to influence dialog construction.
type DialogFlags int

const (
	// DialogFlagsModal: make the constructed dialog modal, see
	// gtk_window_set_modal()
	DialogFlagsModal DialogFlags = 1
	// DialogFlagsDestroyWithParent: destroy the dialog when its parent is
	// destroyed, see gtk_window_set_destroy_with_parent()
	DialogFlagsDestroyWithParent DialogFlags = 2
	// DialogFlagsUseHeaderBar: create dialog with actions in header bar instead
	// of action area. Since 3.12.
	DialogFlagsUseHeaderBar DialogFlags = 4
)

func marshalDialogFlags(p uintptr) (interface{}, error) {
	return DialogFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Dialog: dialog boxes are a convenient way to prompt the user for a small
// amount of input, e.g. to display a message, ask a question, or anything else
// that does not require extensive effort on the user’s part.
//
// GTK+ treats a dialog as a window split vertically. The top section is a VBox,
// and is where widgets such as a Label or a Entry should be packed. The bottom
// area is known as the “action area”. This is generally used for packing
// buttons into the dialog which may perform functions such as cancel, ok, or
// apply.
//
// Dialog boxes are created with a call to gtk_dialog_new() or
// gtk_dialog_new_with_buttons(). gtk_dialog_new_with_buttons() is recommended;
// it allows you to set the dialog title, some convenient flags, and add simple
// buttons.
//
// If “dialog” is a newly created dialog, the two primary areas of the window
// can be accessed through gtk_dialog_get_content_area() and
// gtk_dialog_get_action_area(), as can be seen from the example below.
//
// A “modal” dialog (that is, one which freezes the rest of the application from
// user input), can be created by calling gtk_window_set_modal() on the dialog.
// Use the GTK_WINDOW() macro to cast the widget returned from gtk_dialog_new()
// into a Window. When using gtk_dialog_new_with_buttons() you can also pass the
// K_DIALOG_MODAL flag to make a dialog modal.
//
// If you add buttons to Dialog using gtk_dialog_new_with_buttons(),
// gtk_dialog_add_button(), gtk_dialog_add_buttons(), or
// gtk_dialog_add_action_widget(), clicking the button will emit a signal called
// Dialog::response with a response ID that you specified. GTK+ will never
// assign a meaning to positive response IDs; these are entirely user-defined.
// But for convenience, you can use the response IDs in the ResponseType
// enumeration (these all have values less than zero). If a dialog receives a
// delete event, the Dialog::response signal will be emitted with a response ID
// of K_RESPONSE_DELETE_EVENT.
//
// If you want to block waiting for a dialog to return before returning control
// flow to your code, you can call gtk_dialog_run(). This function enters a
// recursive main loop and waits for the user to respond to the dialog,
// returning the response ID corresponding to the button the user clicked.
//
// For the simple dialog in the following example, in reality you’d probably use
// MessageDialog to save yourself some effort. But you’d need to create the
// dialog contents manually if you had more than a simple message in the dialog.
//
// An example for simple GtkDialog usage:
//
//    // Function to open a dialog box with a message
//    void
//    quick_message (GtkWindow *parent, gchar *message)
//    {
//     GtkWidget *dialog, *label, *content_area;
//     GtkDialogFlags flags;
//
//     // Create the widgets
//     flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//     dialog = gtk_dialog_new_with_buttons ("Message",
//                                           parent,
//                                           flags,
//                                           _("_OK"),
//                                           GTK_RESPONSE_NONE,
//                                           NULL);
//     content_area = gtk_dialog_get_content_area (GTK_DIALOG (dialog));
//     label = gtk_label_new (message);
//
//     // Ensure that the dialog box is destroyed when the user responds
//
//     g_signal_connect_swapped (dialog,
//                               "response",
//                               G_CALLBACK (gtk_widget_destroy),
//                               dialog);
//
//     // Add the label, and show everything we’ve added
//
//     gtk_container_add (GTK_CONTAINER (content_area), label);
//     gtk_widget_show_all (dialog);
//    }
//
//
// GtkDialog as GtkBuildable
//
// The GtkDialog implementation of the Buildable interface exposes the @vbox and
// @action_area as internal children with the names “vbox” and “action_area”.
//
// GtkDialog supports a custom <action-widgets> element, which can contain
// multiple <action-widget> elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget (which
// should be a child of the dialogs @action_area). To mark a response as
// default, set the “default“ attribute of the <action-widget> element to true.
//
// GtkDialog supports adding action widgets by specifying “action“ as the “type“
// attribute of a <child> element. The widget will be added either to the action
// area or the headerbar of the dialog, depending on the “use-header-bar“
// property. The response id has to be associated with the action widget using
// the <action-widgets> element.
//
// An example of a Dialog UI definition fragment:
//
//    <object class="GtkDialog" id="dialog1">
//      <child type="action">
//        <object class="GtkButton" id="button_cancel"/>
//      </child>
//      <child type="action">
//        <object class="GtkButton" id="button_ok">
//          <property name="can-default">True</property>
//        </object>
//      </child>
//      <action-widgets>
//        <action-widget response="cancel">button_cancel</action-widget>
//        <action-widget response="ok" default="true">button_ok</action-widget>
//      </action-widgets>
//    </object>
type Dialog interface {
	Window
	Buildable

	// AddActionWidget adds an activatable widget to the action area of a
	// Dialog, connecting a signal handler that will emit the Dialog::response
	// signal on the dialog when the widget is activated. The widget is appended
	// to the end of the dialog’s action area. If you want to add a
	// non-activatable widget, simply pack it into the @action_area field of the
	// Dialog struct.
	AddActionWidget(child Widget, responseId int)
	// AddButton adds a button with the given text and sets things up so that
	// clicking the button will emit the Dialog::response signal with the given
	// @response_id. The button is appended to the end of the dialog’s action
	// area. The button widget is returned, but usually you don’t need it.
	AddButton(buttonText string, responseId int) Widget
	// ActionArea returns the action area of @dialog.
	ActionArea() Box
	// ContentArea returns the content area of @dialog.
	ContentArea() Box
	// HeaderBar returns the header bar of @dialog. Note that the headerbar is
	// only used by the dialog if the Dialog:use-header-bar property is true.
	HeaderBar() HeaderBar
	// ResponseForWidget gets the response id of a widget in the action area of
	// a dialog.
	ResponseForWidget(widget Widget) int
	// WidgetForResponse gets the widget button that uses the given response ID
	// in the action area of a dialog.
	WidgetForResponse(responseId int) Widget
	// Response emits the Dialog::response signal with the given response ID.
	// Used to indicate that the user has responded to the dialog in some way;
	// typically either you or gtk_dialog_run() will be monitoring the
	// ::response signal and take appropriate action.
	Response(responseId int)
	// Run blocks in a recursive main loop until the @dialog either emits the
	// Dialog::response signal, or is destroyed. If the dialog is destroyed
	// during the call to gtk_dialog_run(), gtk_dialog_run() returns
	// K_RESPONSE_NONE. Otherwise, it returns the response ID from the
	// ::response signal emission.
	//
	// Before entering the recursive main loop, gtk_dialog_run() calls
	// gtk_widget_show() on the dialog for you. Note that you still need to show
	// any children of the dialog yourself.
	//
	// During gtk_dialog_run(), the default behavior of Widget::delete-event is
	// disabled; if the dialog receives ::delete_event, it will not be destroyed
	// as windows usually are, and gtk_dialog_run() will return
	// K_RESPONSE_DELETE_EVENT. Also, during gtk_dialog_run() the dialog will be
	// modal. You can force gtk_dialog_run() to return at any time by calling
	// gtk_dialog_response() to emit the ::response signal. Destroying the
	// dialog during gtk_dialog_run() is a very bad idea, because your post-run
	// code won’t know whether the dialog was destroyed or not.
	//
	// After gtk_dialog_run() returns, you are responsible for hiding or
	// destroying the dialog if you wish to do so.
	//
	// Typical usage of this function might be:
	//
	//      GtkWidget *dialog = gtk_dialog_new ();
	//      // Set up dialog...
	//
	//      int result = gtk_dialog_run (GTK_DIALOG (dialog));
	//      switch (result)
	//        {
	//          case GTK_RESPONSE_ACCEPT:
	//             // do_application_specific_something ();
	//             break;
	//          default:
	//             // do_nothing_since_dialog_was_cancelled ();
	//             break;
	//        }
	//      gtk_widget_destroy (dialog);
	//
	// Note that even though the recursive main loop gives the effect of a modal
	// dialog (it prevents the user from interacting with other windows in the
	// same window group while the dialog is run), callbacks such as timeouts,
	// IO channel watches, DND drops, etc, will be triggered during a
	// gtk_dialog_run() call.
	Run() int
	// SetAlternativeButtonOrderFromArray sets an alternative button order. If
	// the Settings:gtk-alternative-button-order setting is set to true, the
	// dialog buttons are reordered according to the order of the response ids
	// in @new_order.
	//
	// See gtk_dialog_set_alternative_button_order() for more information.
	//
	// This function is for use by language bindings.
	SetAlternativeButtonOrderFromArray(newOrder []int)
	// SetDefaultResponse sets the last widget in the dialog’s action area with
	// the given @response_id as the default widget for the dialog. Pressing
	// “Enter” normally activates the default widget.
	SetDefaultResponse(responseId int)
	// SetResponseSensitive calls `gtk_widget_set_sensitive (widget, @setting)`
	// for each widget in the dialog’s action area with the given @response_id.
	// A convenient way to sensitize/desensitize dialog buttons.
	SetResponseSensitive(responseId int, setting bool)
}

// dialog implements the Dialog class.
type dialog struct {
	Window
	Buildable
}

var _ Dialog = (*dialog)(nil)

// WrapDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapDialog(obj *externglib.Object) Dialog {
	return dialog{
		Window:    WrapWindow(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDialog(obj), nil
}

// NewDialog constructs a class Dialog.
func NewDialog() Dialog {
	var _cret C.GtkDialog // in

	_cret = C.gtk_dialog_new()

	var _dialog Dialog // out

	_dialog = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Dialog)

	return _dialog
}

// AddActionWidget adds an activatable widget to the action area of a
// Dialog, connecting a signal handler that will emit the Dialog::response
// signal on the dialog when the widget is activated. The widget is appended
// to the end of the dialog’s action area. If you want to add a
// non-activatable widget, simply pack it into the @action_area field of the
// Dialog struct.
func (d dialog) AddActionWidget(child Widget, responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(responseId)

	C.gtk_dialog_add_action_widget(_arg0, _arg1, _arg2)
}

// AddButton adds a button with the given text and sets things up so that
// clicking the button will emit the Dialog::response signal with the given
// @response_id. The button is appended to the end of the dialog’s action
// area. The button widget is returned, but usually you don’t need it.
func (d dialog) AddButton(buttonText string, responseId int) Widget {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(buttonText))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(responseId)

	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_add_button(_arg0, _arg1, _arg2)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// ActionArea returns the action area of @dialog.
func (d dialog) ActionArea() Box {
	var _arg0 *C.GtkDialog // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_get_action_area(_arg0)

	var _box Box // out

	_box = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Box)

	return _box
}

// ContentArea returns the content area of @dialog.
func (d dialog) ContentArea() Box {
	var _arg0 *C.GtkDialog // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_get_content_area(_arg0)

	var _box Box // out

	_box = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Box)

	return _box
}

// HeaderBar returns the header bar of @dialog. Note that the headerbar is
// only used by the dialog if the Dialog:use-header-bar property is true.
func (d dialog) HeaderBar() HeaderBar {
	var _arg0 *C.GtkDialog // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_get_header_bar(_arg0)

	var _headerBar HeaderBar // out

	_headerBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(HeaderBar)

	return _headerBar
}

// ResponseForWidget gets the response id of a widget in the action area of
// a dialog.
func (d dialog) ResponseForWidget(widget Widget) int {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret C.gint // in

	_cret = C.gtk_dialog_get_response_for_widget(_arg0, _arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// WidgetForResponse gets the widget button that uses the given response ID
// in the action area of a dialog.
func (d dialog) WidgetForResponse(responseId int) Widget {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.gint(responseId)

	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_get_widget_for_response(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Response emits the Dialog::response signal with the given response ID.
// Used to indicate that the user has responded to the dialog in some way;
// typically either you or gtk_dialog_run() will be monitoring the
// ::response signal and take appropriate action.
func (d dialog) Response(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.gint(responseId)

	C.gtk_dialog_response(_arg0, _arg1)
}

// Run blocks in a recursive main loop until the @dialog either emits the
// Dialog::response signal, or is destroyed. If the dialog is destroyed
// during the call to gtk_dialog_run(), gtk_dialog_run() returns
// K_RESPONSE_NONE. Otherwise, it returns the response ID from the
// ::response signal emission.
//
// Before entering the recursive main loop, gtk_dialog_run() calls
// gtk_widget_show() on the dialog for you. Note that you still need to show
// any children of the dialog yourself.
//
// During gtk_dialog_run(), the default behavior of Widget::delete-event is
// disabled; if the dialog receives ::delete_event, it will not be destroyed
// as windows usually are, and gtk_dialog_run() will return
// K_RESPONSE_DELETE_EVENT. Also, during gtk_dialog_run() the dialog will be
// modal. You can force gtk_dialog_run() to return at any time by calling
// gtk_dialog_response() to emit the ::response signal. Destroying the
// dialog during gtk_dialog_run() is a very bad idea, because your post-run
// code won’t know whether the dialog was destroyed or not.
//
// After gtk_dialog_run() returns, you are responsible for hiding or
// destroying the dialog if you wish to do so.
//
// Typical usage of this function might be:
//
//      GtkWidget *dialog = gtk_dialog_new ();
//      // Set up dialog...
//
//      int result = gtk_dialog_run (GTK_DIALOG (dialog));
//      switch (result)
//        {
//          case GTK_RESPONSE_ACCEPT:
//             // do_application_specific_something ();
//             break;
//          default:
//             // do_nothing_since_dialog_was_cancelled ();
//             break;
//        }
//      gtk_widget_destroy (dialog);
//
// Note that even though the recursive main loop gives the effect of a modal
// dialog (it prevents the user from interacting with other windows in the
// same window group while the dialog is run), callbacks such as timeouts,
// IO channel watches, DND drops, etc, will be triggered during a
// gtk_dialog_run() call.
func (d dialog) Run() int {
	var _arg0 *C.GtkDialog // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))

	var _cret C.gint // in

	_cret = C.gtk_dialog_run(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// SetAlternativeButtonOrderFromArray sets an alternative button order. If
// the Settings:gtk-alternative-button-order setting is set to true, the
// dialog buttons are reordered according to the order of the response ids
// in @new_order.
//
// See gtk_dialog_set_alternative_button_order() for more information.
//
// This function is for use by language bindings.
func (d dialog) SetAlternativeButtonOrderFromArray(newOrder []int) {
	var _arg0 *C.GtkDialog // out
	var _arg2 *C.gint
	var _arg1 C.gint

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.gint(len(newOrder))
	_arg2 = (*C.gint)(unsafe.Pointer(&newOrder[0]))

	C.gtk_dialog_set_alternative_button_order_from_array(_arg0, _arg1, _arg2)
}

// SetDefaultResponse sets the last widget in the dialog’s action area with
// the given @response_id as the default widget for the dialog. Pressing
// “Enter” normally activates the default widget.
func (d dialog) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.gint(responseId)

	C.gtk_dialog_set_default_response(_arg0, _arg1)
}

// SetResponseSensitive calls `gtk_widget_set_sensitive (widget, @setting)`
// for each widget in the dialog’s action area with the given @response_id.
// A convenient way to sensitize/desensitize dialog buttons.
func (d dialog) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out
	var _arg2 C.gboolean   // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.gint(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_dialog_set_response_sensitive(_arg0, _arg1, _arg2)
}
