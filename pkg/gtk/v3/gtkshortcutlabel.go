// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_label_get_type()), F: marshalShortcutLabeller},
	})
}

// ShortcutLabel is a widget that represents a single keyboard shortcut or
// gesture in the user interface.
type ShortcutLabel struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer         = (*ShortcutLabel)(nil)
	_ externglib.Objector = (*ShortcutLabel)(nil)
)

func wrapShortcutLabel(obj *externglib.Object) *ShortcutLabel {
	return &ShortcutLabel{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutLabeller(p uintptr) (interface{}, error) {
	return wrapShortcutLabel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewShortcutLabel creates a new ShortcutLabel with accelerator set.
//
// The function takes the following parameters:
//
//    - accelerator: initial accelerator.
//
// The function returns the following values:
//
//    - shortcutLabel: newly-allocated ShortcutLabel.
//
func NewShortcutLabel(accelerator string) *ShortcutLabel {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_shortcut_label_new(_arg1)
	runtime.KeepAlive(accelerator)

	var _shortcutLabel *ShortcutLabel // out

	_shortcutLabel = wrapShortcutLabel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcutLabel
}

// Accelerator retrieves the current accelerator of self.
//
// The function returns the following values:
//
//    - utf8 (optional): current accelerator.
//
func (self *ShortcutLabel) Accelerator() string {
	var _arg0 *C.GtkShortcutLabel // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_label_get_accelerator(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// DisabledText retrieves the text that is displayed when no accelerator is set.
//
// The function returns the following values:
//
//    - utf8 (optional): current text displayed when no accelerator is set.
//
func (self *ShortcutLabel) DisabledText() string {
	var _arg0 *C.GtkShortcutLabel // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_label_get_disabled_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetAccelerator sets the accelerator to be displayed by self.
//
// The function takes the following parameters:
//
//    - accelerator: new accelerator.
//
func (self *ShortcutLabel) SetAccelerator(accelerator string) {
	var _arg0 *C.GtkShortcutLabel // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_shortcut_label_set_accelerator(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(accelerator)
}

// SetDisabledText sets the text to be displayed by self when no accelerator is
// set.
//
// The function takes the following parameters:
//
//    - disabledText: text to be displayed when no accelerator is set.
//
func (self *ShortcutLabel) SetDisabledText(disabledText string) {
	var _arg0 *C.GtkShortcutLabel // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(disabledText)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_shortcut_label_set_disabled_text(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(disabledText)
}
