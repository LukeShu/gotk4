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
		{T: externglib.Type(C.gtk_link_button_get_type()), F: marshalLinkButton},
	})
}

// LinkButton: a GtkLinkButton is a Button with a hyperlink, similar to the one
// used by web browsers, which triggers an action when clicked. It is useful to
// show quick links to resources.
//
// A link button is created by calling either gtk_link_button_new() or
// gtk_link_button_new_with_label(). If using the former, the URI you pass to
// the constructor is used as a label for the widget.
//
// The URI bound to a GtkLinkButton can be set specifically using
// gtk_link_button_set_uri(), and retrieved using gtk_link_button_get_uri().
//
// By default, GtkLinkButton calls gtk_show_uri_on_window() when the button is
// clicked. This behaviour can be overridden by connecting to the
// LinkButton::activate-link signal and returning true from the signal handler.
//
//
// CSS nodes
//
// GtkLinkButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .link style class.
type LinkButton interface {
	Button
	Actionable
	Activatable
	Buildable

	// URI retrieves the URI set using gtk_link_button_set_uri().
	URI() string
	// Visited retrieves the “visited” state of the URI where the LinkButton
	// points. The button becomes visited when it is clicked. If the URI is
	// changed on the button, the “visited” state is unset again.
	//
	// The state may also be changed using gtk_link_button_set_visited().
	Visited() bool
	// SetURI sets @uri as the URI where the LinkButton points. As a side-effect
	// this unsets the “visited” state of the button.
	SetURI(uri string)
	// SetVisited sets the “visited” state of the URI where the LinkButton
	// points. See gtk_link_button_get_visited() for more details.
	SetVisited(visited bool)
}

// linkButton implements the LinkButton class.
type linkButton struct {
	Button
	Actionable
	Activatable
	Buildable
}

var _ LinkButton = (*linkButton)(nil)

// WrapLinkButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapLinkButton(obj *externglib.Object) LinkButton {
	return linkButton{
		Button:      WrapButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalLinkButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLinkButton(obj), nil
}

// NewLinkButton constructs a class LinkButton.
func NewLinkButton(uri string) LinkButton {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkLinkButton // in

	_cret = C.gtk_link_button_new(_arg1)

	var _linkButton LinkButton // out

	_linkButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(LinkButton)

	return _linkButton
}

// NewLinkButtonWithLabel constructs a class LinkButton.
func NewLinkButtonWithLabel(uri string, label string) LinkButton {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret C.GtkLinkButton // in

	_cret = C.gtk_link_button_new_with_label(_arg1, _arg2)

	var _linkButton LinkButton // out

	_linkButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(LinkButton)

	return _linkButton
}

// URI retrieves the URI set using gtk_link_button_set_uri().
func (l linkButton) URI() string {
	var _arg0 *C.GtkLinkButton // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_link_button_get_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Visited retrieves the “visited” state of the URI where the LinkButton
// points. The button becomes visited when it is clicked. If the URI is
// changed on the button, the “visited” state is unset again.
//
// The state may also be changed using gtk_link_button_set_visited().
func (l linkButton) Visited() bool {
	var _arg0 *C.GtkLinkButton // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_link_button_get_visited(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetURI sets @uri as the URI where the LinkButton points. As a side-effect
// this unsets the “visited” state of the button.
func (l linkButton) SetURI(uri string) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_link_button_set_uri(_arg0, _arg1)
}

// SetVisited sets the “visited” state of the URI where the LinkButton
// points. See gtk_link_button_get_visited() for more details.
func (l linkButton) SetVisited(visited bool) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))
	if visited {
		_arg1 = C.TRUE
	}

	C.gtk_link_button_set_visited(_arg0, _arg1)
}
