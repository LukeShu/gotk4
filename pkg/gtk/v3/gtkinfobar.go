// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_info_bar_get_type()), F: marshalInfoBar},
	})
}

// InfoBarOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type InfoBarOverrider interface {
	Close()
	// Response emits the “response” signal with the given @response_id.
	Response(responseId int)
}

// InfoBar is a widget that can be used to show messages to the user without
// showing a dialog. It is often temporarily shown at the top or bottom of a
// document. In contrast to Dialog, which has a action area at the bottom,
// InfoBar has an action area at the side.
//
// The API of InfoBar is very similar to Dialog, allowing you to add buttons to
// the action area with gtk_info_bar_add_button() or
// gtk_info_bar_new_with_buttons(). The sensitivity of action widgets can be
// controlled with gtk_info_bar_set_response_sensitive(). To add widgets to the
// main content area of a InfoBar, use gtk_info_bar_get_content_area() and add
// your widgets to the container.
//
// Similar to MessageDialog, the contents of a InfoBar can by classified as
// error message, warning, informational message, etc, by using
// gtk_info_bar_set_message_type(). GTK+ may use the message type to determine
// how the message is displayed.
//
// A simple example for using a InfoBar:
//
//    GtkWidget *widget, *message_label, *content_area;
//    GtkWidget *grid;
//    GtkInfoBar *bar;
//
//    // set up info bar
//    widget = gtk_info_bar_new ();
//    bar = GTK_INFO_BAR (widget);
//    grid = gtk_grid_new ();
//
//    gtk_widget_set_no_show_all (widget, TRUE);
//    message_label = gtk_label_new ("");
//    content_area = gtk_info_bar_get_content_area (bar);
//    gtk_container_add (GTK_CONTAINER (content_area),
//                       message_label);
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
//    gtk_info_bar_set_message_type (bar,
//                                   GTK_MESSAGE_ERROR);
//    gtk_widget_show (bar);
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
// should be a child of the dialogs @action_area).
//
//
// CSS nodes
//
// GtkInfoBar has a single CSS node with name infobar. The node may get one of
// the style classes .info, .warning, .error or .question, depending on the
// message type.
type InfoBar interface {
	gextras.Objector

	// AddActionWidget: add an activatable widget to the action area of a
	// InfoBar, connecting a signal handler that will emit the InfoBar::response
	// signal on the message area when the widget is activated. The widget is
	// appended to the end of the message areas action area.
	AddActionWidget(child Widget, responseId int)
	// AddButton adds a button with the given text and sets things up so that
	// clicking the button will emit the “response” signal with the given
	// response_id. The button is appended to the end of the info bars's action
	// area. The button widget is returned, but usually you don't need it.
	AddButton(buttonText string, responseId int) *ButtonClass
	// ActionArea returns the action area of @info_bar.
	ActionArea() *BoxClass
	// ContentArea returns the content area of @info_bar.
	ContentArea() *BoxClass
	// MessageType returns the message type of the message area.
	MessageType() MessageType
	Revealed() bool
	// ShowCloseButton returns whether the widget will display a standard close
	// button.
	ShowCloseButton() bool
	// Response emits the “response” signal with the given @response_id.
	Response(responseId int)
	// SetDefaultResponse sets the last widget in the info bar’s action area
	// with the given response_id as the default widget for the dialog. Pressing
	// “Enter” normally activates the default widget.
	//
	// Note that this function currently requires @info_bar to be added to a
	// widget hierarchy.
	SetDefaultResponse(responseId int)
	// SetResponseSensitive calls gtk_widget_set_sensitive (widget, setting) for
	// each widget in the info bars’s action area with the given response_id. A
	// convenient way to sensitize/desensitize dialog buttons.
	SetResponseSensitive(responseId int, setting bool)
	// SetRevealed sets the GtkInfoBar:revealed property to @revealed. This will
	// cause @info_bar to show up with a slide-in transition.
	//
	// Note that this property does not automatically show @info_bar and thus
	// won’t have any effect if it is invisible.
	SetRevealed(revealed bool)
	// SetShowCloseButton: if true, a standard close button is shown. When
	// clicked it emits the response GTK_RESPONSE_CLOSE.
	SetShowCloseButton(setting bool)
}

// InfoBarClass implements the InfoBar interface.
type InfoBarClass struct {
	*externglib.Object
	BoxClass
	BuildableInterface
	OrientableInterface
}

var _ InfoBar = (*InfoBarClass)(nil)

func wrapInfoBar(obj *externglib.Object) InfoBar {
	return &InfoBarClass{
		Object: obj,
		BoxClass: BoxClass{
			Object: obj,
			ContainerClass: ContainerClass{
				Object: obj,
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			OrientableInterface: OrientableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalInfoBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInfoBar(obj), nil
}

// NewInfoBar creates a new InfoBar object.
func NewInfoBar() *InfoBarClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_info_bar_new()

	var _infoBar *InfoBarClass // out

	_infoBar = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*InfoBarClass)

	return _infoBar
}

// AddActionWidget: add an activatable widget to the action area of a InfoBar,
// connecting a signal handler that will emit the InfoBar::response signal on
// the message area when the widget is activated. The widget is appended to the
// end of the message areas action area.
func (i *InfoBarClass) AddActionWidget(child Widget, responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&child).Native()))
	_arg2 = C.gint(responseId)

	C.gtk_info_bar_add_action_widget(_arg0, _arg1, _arg2)
}

// AddButton adds a button with the given text and sets things up so that
// clicking the button will emit the “response” signal with the given
// response_id. The button is appended to the end of the info bars's action
// area. The button widget is returned, but usually you don't need it.
func (i *InfoBarClass) AddButton(buttonText string, responseId int) *ButtonClass {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gint        // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))
	_arg1 = (*C.gchar)(C.CString(buttonText))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(responseId)

	_cret = C.gtk_info_bar_add_button(_arg0, _arg1, _arg2)

	var _button *ButtonClass // out

	_button = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ButtonClass)

	return _button
}

// ActionArea returns the action area of @info_bar.
func (i *InfoBarClass) ActionArea() *BoxClass {
	var _arg0 *C.GtkInfoBar // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))

	_cret = C.gtk_info_bar_get_action_area(_arg0)

	var _box *BoxClass // out

	_box = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*BoxClass)

	return _box
}

// ContentArea returns the content area of @info_bar.
func (i *InfoBarClass) ContentArea() *BoxClass {
	var _arg0 *C.GtkInfoBar // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))

	_cret = C.gtk_info_bar_get_content_area(_arg0)

	var _box *BoxClass // out

	_box = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*BoxClass)

	return _box
}

// MessageType returns the message type of the message area.
func (i *InfoBarClass) MessageType() MessageType {
	var _arg0 *C.GtkInfoBar    // out
	var _cret C.GtkMessageType // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))

	_cret = C.gtk_info_bar_get_message_type(_arg0)

	var _messageType MessageType // out

	_messageType = (MessageType)(_cret)

	return _messageType
}

func (i *InfoBarClass) Revealed() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))

	_cret = C.gtk_info_bar_get_revealed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether the widget will display a standard close
// button.
func (i *InfoBarClass) ShowCloseButton() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))

	_cret = C.gtk_info_bar_get_show_close_button(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Response emits the “response” signal with the given @response_id.
func (i *InfoBarClass) Response(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))
	_arg1 = C.gint(responseId)

	C.gtk_info_bar_response(_arg0, _arg1)
}

// SetDefaultResponse sets the last widget in the info bar’s action area with
// the given response_id as the default widget for the dialog. Pressing “Enter”
// normally activates the default widget.
//
// Note that this function currently requires @info_bar to be added to a widget
// hierarchy.
func (i *InfoBarClass) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))
	_arg1 = C.gint(responseId)

	C.gtk_info_bar_set_default_response(_arg0, _arg1)
}

// SetResponseSensitive calls gtk_widget_set_sensitive (widget, setting) for
// each widget in the info bars’s action area with the given response_id. A
// convenient way to sensitize/desensitize dialog buttons.
func (i *InfoBarClass) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gint        // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))
	_arg1 = C.gint(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_info_bar_set_response_sensitive(_arg0, _arg1, _arg2)
}

// SetRevealed sets the GtkInfoBar:revealed property to @revealed. This will
// cause @info_bar to show up with a slide-in transition.
//
// Note that this property does not automatically show @info_bar and thus won’t
// have any effect if it is invisible.
func (i *InfoBarClass) SetRevealed(revealed bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))
	if revealed {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_revealed(_arg0, _arg1)
}

// SetShowCloseButton: if true, a standard close button is shown. When clicked
// it emits the response GTK_RESPONSE_CLOSE.
func (i *InfoBarClass) SetShowCloseButton(setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&i).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_show_close_button(_arg0, _arg1)
}
