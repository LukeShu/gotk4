// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtknotebookpageaccessible.go.
var GTypeNotebookPageAccessible = coreglib.Type(C.gtk_notebook_page_accessible_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeNotebookPageAccessible, F: marshalNotebookPageAccessible},
	})
}

// NotebookPageAccessibleOverrider contains methods that are overridable.
type NotebookPageAccessibleOverrider interface {
}

type NotebookPageAccessible struct {
	_ [0]func() // equal guard
	atk.ObjectClass

	*coreglib.Object
	atk.Component
}

var (
	_ coreglib.Objector = (*NotebookPageAccessible)(nil)
)

func classInitNotebookPageAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNotebookPageAccessible(obj *coreglib.Object) *NotebookPageAccessible {
	return &NotebookPageAccessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		Object: obj,
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalNotebookPageAccessible(p uintptr) (interface{}, error) {
	return wrapNotebookPageAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
//    - notebook
//    - child
//
// The function returns the following values:
//
func NewNotebookPageAccessible(notebook *NotebookAccessible, child Widgetter) *NotebookPageAccessible {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(notebook).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(**NotebookAccessible)(unsafe.Pointer(&args[0])) = _arg0
	*(*Widgetter)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "NotebookPageAccessible").InvokeMethod("new_NotebookPageAccessible", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(notebook)
	runtime.KeepAlive(child)

	var _notebookPageAccessible *NotebookPageAccessible // out

	_notebookPageAccessible = wrapNotebookPageAccessible(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _notebookPageAccessible
}

func (page *NotebookPageAccessible) Invalidate() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**NotebookPageAccessible)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "NotebookPageAccessible").InvokeMethod("invalidate", args[:], nil)

	runtime.KeepAlive(page)
}
