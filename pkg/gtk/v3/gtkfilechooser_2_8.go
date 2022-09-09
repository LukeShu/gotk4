// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeFileChooserConfirmation = coreglib.Type(C.gtk_file_chooser_confirmation_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileChooserConfirmation, F: marshalFileChooserConfirmation},
	})
}

// FileChooserConfirmation: used as a return value of handlers for the
// FileChooser::confirm-overwrite signal of a FileChooser. This value determines
// whether the file chooser will present the stock confirmation dialog, accept
// the user’s choice of a filename, or let the user choose another filename.
type FileChooserConfirmation C.gint

const (
	// FileChooserConfirmationConfirm: file chooser will present its stock
	// dialog to confirm about overwriting an existing file.
	FileChooserConfirmationConfirm FileChooserConfirmation = iota
	// FileChooserConfirmationAcceptFilename: file chooser will terminate and
	// accept the user’s choice of a file name.
	FileChooserConfirmationAcceptFilename
	// FileChooserConfirmationSelectAgain: file chooser will continue running,
	// so as to let the user select another file name.
	FileChooserConfirmationSelectAgain
)

func marshalFileChooserConfirmation(p uintptr) (interface{}, error) {
	return FileChooserConfirmation(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FileChooserConfirmation.
func (f FileChooserConfirmation) String() string {
	switch f {
	case FileChooserConfirmationConfirm:
		return "Confirm"
	case FileChooserConfirmationAcceptFilename:
		return "AcceptFilename"
	case FileChooserConfirmationSelectAgain:
		return "SelectAgain"
	default:
		return fmt.Sprintf("FileChooserConfirmation(%d)", f)
	}
}
