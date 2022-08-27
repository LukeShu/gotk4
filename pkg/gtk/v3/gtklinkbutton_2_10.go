// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewLinkButton creates a new LinkButton with the URI as its text.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
// The function returns the following values:
//
//    - linkButton: new link button widget.
//
func NewLinkButton(uri string) *LinkButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_link_button_new(_arg1)
	runtime.KeepAlive(uri)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// NewLinkButtonWithLabel creates a new LinkButton containing a label.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//    - label (optional): text of the button.
//
// The function returns the following values:
//
//    - linkButton: new link button widget.
//
func NewLinkButtonWithLabel(uri, label string) *LinkButton {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_link_button_new_with_label(_arg1, _arg2)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(label)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// URI retrieves the URI set using gtk_link_button_set_uri().
//
// The function returns the following values:
//
//    - utf8: valid URI. The returned string is owned by the link button and
//      should not be modified or freed.
//
func (linkButton *LinkButton) URI() string {
	var _arg0 *C.GtkLinkButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(coreglib.InternObject(linkButton).Native()))

	_cret = C.gtk_link_button_get_uri(_arg0)
	runtime.KeepAlive(linkButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetURI sets uri as the URI where the LinkButton points. As a side-effect this
// unsets the “visited” state of the button.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
func (linkButton *LinkButton) SetURI(uri string) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(coreglib.InternObject(linkButton).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_link_button_set_uri(_arg0, _arg1)
	runtime.KeepAlive(linkButton)
	runtime.KeepAlive(uri)
}
