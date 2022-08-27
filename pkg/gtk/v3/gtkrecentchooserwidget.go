// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
	GTypeRecentChooserWidget = coreglib.Type(C.gtk_recent_chooser_widget_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRecentChooserWidget, F: marshalRecentChooserWidget},
	})
}

// RecentChooserWidgetOverrides contains methods that are overridable.
type RecentChooserWidgetOverrides struct {
}

func defaultRecentChooserWidgetOverrides(v *RecentChooserWidget) RecentChooserWidgetOverrides {
	return RecentChooserWidgetOverrides{}
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

	*coreglib.Object
	RecentChooser
}

var (
	_ coreglib.Objector = (*RecentChooserWidget)(nil)
	_ Containerer       = (*RecentChooserWidget)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*RecentChooserWidget, *RecentChooserWidgetClass, RecentChooserWidgetOverrides](
		GTypeRecentChooserWidget,
		initRecentChooserWidgetClass,
		wrapRecentChooserWidget,
		defaultRecentChooserWidgetOverrides,
	)
}

func initRecentChooserWidgetClass(gclass unsafe.Pointer, overrides RecentChooserWidgetOverrides, classInitFunc func(*RecentChooserWidgetClass)) {
	if classInitFunc != nil {
		class := (*RecentChooserWidgetClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapRecentChooserWidget(obj *coreglib.Object) *RecentChooserWidget {
	return &RecentChooserWidget{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
	return wrapRecentChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// RecentChooserWidgetClass: instance of this type is always passed by
// reference.
type RecentChooserWidgetClass struct {
	*recentChooserWidgetClass
}

// recentChooserWidgetClass is the struct that's finalized.
type recentChooserWidgetClass struct {
	native *C.GtkRecentChooserWidgetClass
}

func (r *RecentChooserWidgetClass) ParentClass() *BoxClass {
	valptr := &r.native.parent_class
	var _v *BoxClass // out
	_v = (*BoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
