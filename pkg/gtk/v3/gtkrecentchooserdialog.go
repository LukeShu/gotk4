// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_chooser_dialog_get_type()), F: marshalRecentChooserDialogger},
	})
}

// RecentChooserDialog is a dialog box suitable for displaying the recently used
// documents. This widgets works by putting a RecentChooserWidget inside a
// Dialog. It exposes the RecentChooserIface interface, so you can use all the
// RecentChooser functions on the recent chooser dialog as well as those for
// Dialog.
//
// Note that RecentChooserDialog does not have any methods of its own. Instead,
// you should use the functions that work on a RecentChooser.
//
//
// Typical usage
//
// In the simplest of cases, you can use the following code to use a
// RecentChooserDialog to select a recently used file:
//
//    GtkWidget *dialog;
//    gint res;
//
//    dialog = gtk_recent_chooser_dialog_new ("Recent Documents",
//                                            parent_window,
//                                            _("_Cancel"),
//                                            GTK_RESPONSE_CANCEL,
//                                            _("_Open"),
//                                            GTK_RESPONSE_ACCEPT,
//                                            NULL);
//
//    res = gtk_dialog_run (GTK_DIALOG (dialog));
//    if (res == GTK_RESPONSE_ACCEPT)
//      {
//        GtkRecentInfo *info;
//        GtkRecentChooser *chooser = GTK_RECENT_CHOOSER (dialog);
//
//        info = gtk_recent_chooser_get_current_item (chooser);
//        open_file (gtk_recent_info_get_uri (info));
//        gtk_recent_info_unref (info);
//      }
//
//    gtk_widget_destroy (dialog);
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooserDialog struct {
	Dialog

	RecentChooser
	*externglib.Object
}

func wrapRecentChooserDialog(obj *externglib.Object) *RecentChooserDialog {
	return &RecentChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							Object: obj,
						},
					},
				},
			},
		},
		RecentChooser: RecentChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalRecentChooserDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRecentChooserDialog(obj), nil
}

func (*RecentChooserDialog) privateRecentChooserDialog() {}
