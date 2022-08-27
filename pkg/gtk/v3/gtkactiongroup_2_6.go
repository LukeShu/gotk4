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

// TranslateString translates a string using the function set with
// gtk_action_group_set_translate_func(). This is mainly intended for language
// bindings.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - str: string.
//
// The function returns the following values:
//
//    - utf8: translation of string.
//
func (actionGroup *ActionGroup) TranslateString(str string) string {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(coreglib.InternObject(actionGroup).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_action_group_translate_string(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(str)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
