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
		{T: externglib.Type(C.gtk_link_button_get_type()), F: marshalLinkButtoner},
	})
}

// LinkButtoner describes LinkButton's methods.
type LinkButtoner interface {
	// URI retrieves the URI of the `GtkLinkButton`.
	URI() string
	// Visited retrieves the “visited” state of the `GtkLinkButton`.
	Visited() bool
	// SetURI sets @uri as the URI where the `GtkLinkButton` points.
	SetURI(uri string)
	// SetVisited sets the “visited” state of the `GtkLinkButton`.
	SetVisited(visited bool)
}

// LinkButton: `GtkLinkButton` is a button with a hyperlink.
//
// !An example GtkLinkButton (link-button.png)
//
// It is useful to show quick links to resources.
//
// A link button is created by calling either [ctor@Gtk.LinkButton.new] or
// [ctor@Gtk.LinkButton.new_with_label]. If using the former, the URI you pass
// to the constructor is used as a label for the widget.
//
// The URI bound to a `GtkLinkButton` can be set specifically using
// [method@Gtk.LinkButton.set_uri].
//
// By default, `GtkLinkButton` calls [func@Gtk.show_uri] when the button is
// clicked. This behaviour can be overridden by connecting to the
// [signal@Gtk.LinkButton::activate-link] signal and returning true from the
// signal handler.
//
//
// CSS nodes
//
// `GtkLinkButton` has a single CSS node with name button. To differentiate it
// from a plain `GtkButton`, it gets the .link style class.
//
//
// Accessibility
//
// `GtkLinkButton` uses the K_ACCESSIBLE_ROLE_LINK role.
type LinkButton struct {
	Button
}

var (
	_ LinkButtoner    = (*LinkButton)(nil)
	_ gextras.Nativer = (*LinkButton)(nil)
)

func wrapLinkButton(obj *externglib.Object) LinkButtoner {
	return &LinkButton{
		Button: Button{
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
			Actionable: Actionable{
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
			},
		},
	}
}

func marshalLinkButtoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLinkButton(obj), nil
}

// NewLinkButton creates a new `GtkLinkButton` with the URI as its text.
func NewLinkButton(uri string) *LinkButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_link_button_new(_arg1)

	var _linkButton *LinkButton // out

	_linkButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LinkButton)

	return _linkButton
}

// NewLinkButtonWithLabel creates a new `GtkLinkButton` containing a label.
func NewLinkButtonWithLabel(uri string, label string) *LinkButton {
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_link_button_new_with_label(_arg1, _arg2)

	var _linkButton *LinkButton // out

	_linkButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LinkButton)

	return _linkButton
}

// URI retrieves the URI of the `GtkLinkButton`.
func (linkButton *LinkButton) URI() string {
	var _arg0 *C.GtkLinkButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))

	_cret = C.gtk_link_button_get_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Visited retrieves the “visited” state of the `GtkLinkButton`.
//
// The button becomes visited when it is clicked. If the URI is changed on the
// button, the “visited” state is unset again.
//
// The state may also be changed using [method@Gtk.LinkButton.set_visited].
func (linkButton *LinkButton) Visited() bool {
	var _arg0 *C.GtkLinkButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))

	_cret = C.gtk_link_button_get_visited(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetURI sets @uri as the URI where the `GtkLinkButton` points.
//
// As a side-effect this unsets the “visited” state of the button.
func (linkButton *LinkButton) SetURI(uri string) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_link_button_set_uri(_arg0, _arg1)
}

// SetVisited sets the “visited” state of the `GtkLinkButton`.
//
// See [method@Gtk.LinkButton.get_visited] for more details.
func (linkButton *LinkButton) SetVisited(visited bool) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))
	if visited {
		_arg1 = C.TRUE
	}

	C.gtk_link_button_set_visited(_arg0, _arg1)
}
