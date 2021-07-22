// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_info_bar_get_type()), F: marshalInfoBarrer},
	})
}

// InfoBar: GtkInfoBar can be show messages to the user without a dialog.
//
// !An example GtkInfoBar (info-bar.png)
//
// It is often temporarily shown at the top or bottom of a document. In contrast
// to gtk.Dialog, which has an action area at the bottom, GtkInfoBar has an
// action area at the side.
//
// The API of GtkInfoBar is very similar to GtkDialog, allowing you to add
// buttons to the action area with gtk.InfoBar.AddButton() or
// gtk.InfoBar.NewWithButtons. The sensitivity of action widgets can be
// controlled with gtk.InfoBar.SetResponseSensitive().
//
// To add widgets to the main content area of a GtkInfoBar, use
// gtk.InfoBar.AddChild().
//
// Similar to gtk.MessageDialog, the contents of a GtkInfoBar can by classified
// as error message, warning, informational message, etc, by using
// gtk.InfoBar.SetMessageType(). GTK may use the message type to determine how
// the message is displayed.
//
// A simple example for using a GtkInfoBar:
//
//    GtkWidget *message_label;
//    GtkWidget *widget;
//    GtkWidget *grid;
//    GtkInfoBar *bar;
//
//    // set up info bar
//    widget = gtk_info_bar_new ();
//    bar = GTK_INFO_BAR (widget);
//    grid = gtk_grid_new ();
//
//    message_label = gtk_label_new ("");
//    gtk_info_bar_add_child (bar, message_label);
//    gtk_info_bar_add_button (bar,
//                             _("_OK"),
//                             GTK_RESPONSE_OK);
//    g_signal_connect (bar,
//                      "response",
//                      G_CALLBACK (gtk_widget_hide),
//                      NULL);
//    gtk_grid_attach (GTK_GRID (grid),
//                     widget,
//                     0, 2, 1, 1);
//
//    // ...
//
//    // show an error message
//    gtk_label_set_text (GTK_LABEL (message_label), "An error occurred!");
//    gtk_info_bar_set_message_type (bar, GTK_MESSAGE_ERROR);
//    gtk_widget_show (bar);
//
//
//
// GtkInfoBar as GtkBuildable
//
// The GtkInfoBar implementation of the GtkBuildable interface exposes the
// content area and action area as internal children with the names
// “content_area” and “action_area”.
//
// GtkInfoBar supports a custom <action-widgets> element, which can contain
// multiple <action-widget> elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget (which
// should be a child of the dialogs action_area).
//
//
// CSS nodes
//
// GtkInfoBar has a single CSS node with name infobar. The node may get one of
// the style classes .info, .warning, .error or .question, depending on the
// message type. If the info bar shows a close button, that button will have the
// .close style class applied.
type InfoBar struct {
	Widget
}

func wrapInfoBar(obj *externglib.Object) *InfoBar {
	return &InfoBar{
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
			Object: obj,
		},
	}
}

func marshalInfoBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInfoBar(obj), nil
}

// NewInfoBar creates a new GtkInfoBar object.
func NewInfoBar() *InfoBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_info_bar_new()

	var _infoBar *InfoBar // out

	_infoBar = wrapInfoBar(externglib.Take(unsafe.Pointer(_cret)))

	return _infoBar
}

// AddActionWidget: add an activatable widget to the action area of a
// GtkInfoBar.
//
// This also connects a signal handler that will emit the gtk.InfoBar::response
// signal on the message area when the widget is activated. The widget is
// appended to the end of the message areas action area.
func (infoBar *InfoBar) AddActionWidget(child Widgetter, responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(responseId)

	C.gtk_info_bar_add_action_widget(_arg0, _arg1, _arg2)
}

// AddButton adds a button with the given text.
//
// Clicking the button will emit the gtk.InfoBar::response signal with the given
// response_id. The button is appended to the end of the info bars's action
// area. The button widget is returned, but usually you don't need it.
func (infoBar *InfoBar) AddButton(buttonText string, responseId int) *Button {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(buttonText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(responseId)

	_cret = C.gtk_info_bar_add_button(_arg0, _arg1, _arg2)

	var _button *Button // out

	_button = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// AddChild adds a widget to the content area of the info bar.
func (infoBar *InfoBar) AddChild(widget Widgetter) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_info_bar_add_child(_arg0, _arg1)
}

// MessageType returns the message type of the message area.
func (infoBar *InfoBar) MessageType() MessageType {
	var _arg0 *C.GtkInfoBar    // out
	var _cret C.GtkMessageType // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))

	_cret = C.gtk_info_bar_get_message_type(_arg0)

	var _messageType MessageType // out

	_messageType = MessageType(_cret)

	return _messageType
}

// Revealed returns whether the info bar is currently revealed.
func (infoBar *InfoBar) Revealed() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))

	_cret = C.gtk_info_bar_get_revealed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether the widget will display a standard close
// button.
func (infoBar *InfoBar) ShowCloseButton() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))

	_cret = C.gtk_info_bar_get_show_close_button(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveActionWidget removes a widget from the action area of info_bar.
//
// The widget must have been put there by a call to
// gtk.InfoBar.AddActionWidget() or gtk.InfoBar.AddButton().
func (infoBar *InfoBar) RemoveActionWidget(widget Widgetter) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_info_bar_remove_action_widget(_arg0, _arg1)
}

// RemoveChild removes a widget from the content area of the info bar.
func (infoBar *InfoBar) RemoveChild(widget Widgetter) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_info_bar_remove_child(_arg0, _arg1)
}

// Response emits the “response” signal with the given response_id.
func (infoBar *InfoBar) Response(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.int         // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = C.int(responseId)

	C.gtk_info_bar_response(_arg0, _arg1)
}

// SetDefaultResponse sets the last widget in the info bar’s action area with
// the given response_id as the default widget for the dialog.
//
// Pressing “Enter” normally activates the default widget.
//
// Note that this function currently requires info_bar to be added to a widget
// hierarchy.
func (infoBar *InfoBar) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.int         // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = C.int(responseId)

	C.gtk_info_bar_set_default_response(_arg0, _arg1)
}

// SetMessageType sets the message type of the message area.
//
// GTK uses this type to determine how the message is displayed.
func (infoBar *InfoBar) SetMessageType(messageType MessageType) {
	var _arg0 *C.GtkInfoBar    // out
	var _arg1 C.GtkMessageType // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = C.GtkMessageType(messageType)

	C.gtk_info_bar_set_message_type(_arg0, _arg1)
}

// SetResponseSensitive sets the sensitivity of action widgets for response_id.
//
// Calls gtk_widget_set_sensitive (widget, setting) for each widget in the info
// bars’s action area with the given response_id. A convenient way to
// sensitize/desensitize buttons.
func (infoBar *InfoBar) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.int         // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	_arg1 = C.int(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_info_bar_set_response_sensitive(_arg0, _arg1, _arg2)
}

// SetRevealed sets whether the GtkInfoBar is revealed.
//
// Changing this will make info_bar reveal or conceal itself via a sliding
// transition.
//
// Note: this does not show or hide info_bar in the gtk.Widget:visible sense, so
// revealing has no effect if gtk.Widget:visible is FALSE.
func (infoBar *InfoBar) SetRevealed(revealed bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	if revealed {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_revealed(_arg0, _arg1)
}

// SetShowCloseButton: if true, a standard close button is shown.
//
// When clicked it emits the response GTK_RESPONSE_CLOSE.
func (infoBar *InfoBar) SetShowCloseButton(setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(infoBar.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_show_close_button(_arg0, _arg1)
}
