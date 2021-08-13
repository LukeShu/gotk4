// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_label_get_type()), F: marshalEditableLabeller},
	})
}

// EditableLabel: GtkEditableLabel is a label that allows users to edit the text
// by switching to an “edit mode”.
//
// !An example GtkEditableLabel (editable-label.png)
//
// GtkEditableLabel does not have API of its own, but it implements the
// gtk.Editable interface.
//
// The default bindings for activating the edit mode is to click or press the
// Enter key. The default bindings for leaving the edit mode are the Enter key
// (to save the results) or the Escape key (to cancel the editing).
//
// CSS nodes
//
//    editablelabel[.editing]
//    ╰── stack
//        ├── label
//        ╰── text
//
//
// GtkEditableLabel has a main node with the name editablelabel. When the entry
// is in editing mode, it gets the .editing style class.
//
// For all the subnodes added to the text node in various situations, see
// gtk.Text.
type EditableLabel struct {
	Widget

	Editable
	*externglib.Object
}

func wrapEditableLabel(obj *externglib.Object) *EditableLabel {
	return &EditableLabel{
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
		Editable: Editable{
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
		},
		Object: obj,
	}
}

func marshalEditableLabeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEditableLabel(obj), nil
}

// NewEditableLabel creates a new GtkEditableLabel widget.
func NewEditableLabel(str string) *EditableLabel {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_editable_label_new(_arg1)

	runtime.KeepAlive(str)

	var _editableLabel *EditableLabel // out

	_editableLabel = wrapEditableLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _editableLabel
}

// Editing returns whether the label is currently in “editing mode”.
func (self *EditableLabel) Editing() bool {
	var _arg0 *C.GtkEditableLabel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_editable_label_get_editing(_arg0)

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartEditing switches the label into “editing mode”.
func (self *EditableLabel) StartEditing() {
	var _arg0 *C.GtkEditableLabel // out

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(self.Native()))

	C.gtk_editable_label_start_editing(_arg0)
	runtime.KeepAlive(self)
}

// StopEditing switches the label out of “editing mode”.
//
// If commit is TRUE, the resulting text is kept as the gtk.Editable:text
// property value, otherwise the resulting text is discarded and the label will
// keep its previous gtk.Editable:text property value.
func (self *EditableLabel) StopEditing(commit bool) {
	var _arg0 *C.GtkEditableLabel // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkEditableLabel)(unsafe.Pointer(self.Native()))
	if commit {
		_arg1 = C.TRUE
	}

	C.gtk_editable_label_stop_editing(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(commit)
}
