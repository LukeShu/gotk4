// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkrecentchooserwidget.go.
var GTypeRecentChooserWidget = externglib.Type(C.gtk_recent_chooser_widget_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeRecentChooserWidget, F: marshalRecentChooserWidget},
	})
}

// RecentChooserWidgetOverrider contains methods that are overridable.
type RecentChooserWidgetOverrider interface {
}

// RecentChooserWidget is a widget suitable for selecting recently used files.
// It is the main building block of a RecentChooserDialog. Most applications
// will only need to use the latter; you can use RecentChooserWidget as part of
// a larger window if you have special needs.
//
// Note that RecentChooserWidget does not have any methods of its own. Instead,
// you should use the functions that work on a RecentChooser.
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooserWidget struct {
	_ [0]func() // equal guard
	Box

	*externglib.Object
	RecentChooser
}

var (
	_ externglib.Objector = (*RecentChooserWidget)(nil)
	_ Containerer         = (*RecentChooserWidget)(nil)
)

func classInitRecentChooserWidgetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRecentChooserWidget(obj *externglib.Object) *RecentChooserWidget {
	return &RecentChooserWidget{
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
		Object: obj,
		RecentChooser: RecentChooser{
			Object: obj,
		},
	}
}

func marshalRecentChooserWidget(p uintptr) (interface{}, error) {
	return wrapRecentChooserWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRecentChooserWidget creates a new RecentChooserWidget object. This is an
// embeddable widget used to access the recently used resources list.
//
// The function returns the following values:
//
//    - recentChooserWidget: new RecentChooserWidget.
//
func NewRecentChooserWidget() *RecentChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_recent_chooser_widget_new()

	var _recentChooserWidget *RecentChooserWidget // out

	_recentChooserWidget = wrapRecentChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserWidget
}

// NewRecentChooserWidgetForManager creates a new RecentChooserWidget with a
// specified recent manager.
//
// This is useful if you have implemented your own recent manager, or if you
// have a customized instance of a RecentManager object.
//
// The function takes the following parameters:
//
//    - manager: RecentManager.
//
// The function returns the following values:
//
//    - recentChooserWidget: new RecentChooserWidget.
//
func NewRecentChooserWidgetForManager(manager *RecentManager) *RecentChooserWidget {
	var _arg1 *C.GtkRecentManager // out
	var _cret *C.GtkWidget        // in

	_arg1 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_chooser_widget_new_for_manager(_arg1)
	runtime.KeepAlive(manager)

	var _recentChooserWidget *RecentChooserWidget // out

	_recentChooserWidget = wrapRecentChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserWidget
}
