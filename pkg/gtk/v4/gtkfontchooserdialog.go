// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeFontChooserDialog returns the GType for the type FontChooserDialog.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontChooserDialog() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "FontChooserDialog").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontChooserDialog)
	return gtype
}

// FontChooserDialog: GtkFontChooserDialog widget is a dialog for selecting a
// font.
//
// !An example GtkFontChooserDialog (fontchooser.png)
//
// GtkFontChooserDialog implements the gtk.FontChooser interface and does not
// provide much API of its own.
//
// To create a GtkFontChooserDialog, use gtk.FontChooserDialog.New.
//
//
// GtkFontChooserDialog as GtkBuildable
//
// The GtkFontChooserDialog implementation of the GtkBuildable interface exposes
// the buttons with the names “select_button” and “cancel_button”.
type FontChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	FontChooser
}

var (
	_ coreglib.Objector = (*FontChooserDialog)(nil)
	_ Widgetter         = (*FontChooserDialog)(nil)
)

func wrapFontChooserDialog(obj *coreglib.Object) *FontChooserDialog {
	return &FontChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
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
				Object: obj,
				Root: Root{
					NativeSurface: NativeSurface{
						Widget: Widget{
							InitiallyUnowned: coreglib.InitiallyUnowned{
								Object: obj,
							},
							Object: obj,
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
				ShortcutManager: ShortcutManager{
					Object: obj,
				},
			},
		},
		Object: obj,
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserDialog(p uintptr) (interface{}, error) {
	return wrapFontChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontChooserDialog creates a new GtkFontChooserDialog.
//
// The function takes the following parameters:
//
//    - title (optional): title of the dialog, or NULL.
//    - parent (optional): transient parent of the dialog, or NULL.
//
// The function returns the following values:
//
//    - fontChooserDialog: new GtkFontChooserDialog.
//
func NewFontChooserDialog(title string, parent *Window) *FontChooserDialog {
	var _args [2]girepository.Argument

	if title != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_args[0]))
	}
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	_gret := girepository.MustFind("Gtk", "FontChooserDialog").InvokeMethod("new_FontChooserDialog", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)

	var _fontChooserDialog *FontChooserDialog // out

	_fontChooserDialog = wrapFontChooserDialog(coreglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserDialog
}
