// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_widget_get_type()), F: marshalFileChooserWidget},
	})
}

// FileChooserWidget: `GtkFileChooserWidget` is a widget for choosing files.
//
// It exposes the [iface@Gtk.FileChooser] interface, and you should use the
// methods of this interface to interact with the widget.
//
//
// CSS nodes
//
// `GtkFileChooserWidget` has a single CSS node with name filechooser.
type FileChooserWidget interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	FileChooser
}

// fileChooserWidget implements the FileChooserWidget class.
type fileChooserWidget struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	FileChooser
}

var _ FileChooserWidget = (*fileChooserWidget)(nil)

// WrapFileChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileChooserWidget(obj *externglib.Object) FileChooserWidget {
	return fileChooserWidget{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		FileChooser:      WrapFileChooser(obj),
	}
}

func marshalFileChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooserWidget(obj), nil
}

// NewFileChooserWidget constructs a class FileChooserWidget.
func NewFileChooserWidget(action FileChooserAction) FileChooserWidget {
	var _arg1 C.GtkFileChooserAction // out

	_arg1 = (C.GtkFileChooserAction)(action)

	var _cret C.GtkFileChooserWidget // in

	_cret = C.gtk_file_chooser_widget_new(_arg1)

	var _fileChooserWidget FileChooserWidget // out

	_fileChooserWidget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(FileChooserWidget)

	return _fileChooserWidget
}
