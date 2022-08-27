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

// IconName returns the name of the themed icon for the tool button, see
// gtk_tool_button_set_icon_name().
//
// The function returns the following values:
//
//    - utf8 (optional): icon name or NULL if the tool button has no themed icon.
//
func (button *ToolButton) IconName() string {
	var _arg0 *C.GtkToolButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_tool_button_get_icon_name(_arg0)
	runtime.KeepAlive(button)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetIconName sets the icon for the tool button from a named themed icon. See
// the docs for IconTheme for more details. The ToolButton:icon-name property
// only has an effect if not overridden by non-NULL ToolButton:label-widget,
// ToolButton:icon-widget and ToolButton:stock-id properties.
//
// The function takes the following parameters:
//
//    - iconName (optional): name of the themed icon.
//
func (button *ToolButton) SetIconName(iconName string) {
	var _arg0 *C.GtkToolButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_tool_button_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(iconName)
}
