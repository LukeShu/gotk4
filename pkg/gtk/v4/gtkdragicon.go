// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drag_icon_get_type()), F: marshalDragIcon},
	})
}

// DragIcon: `GtkDragIcon` is a `GtkRoot` implementation for drag icons.
//
// A drag icon moves with the pointer during a Drag-and-Drop operation and is
// destroyed when the drag ends.
//
// To set up a drag icon and associate it with an ongoing drag operation, use
// [func@Gtk.DragIcon.get_for_drag] to get the icon for a drag. You can then use
// it like any other widget and use [method@Gtk.DragIcon.set_child] to set
// whatever widget should be used for the drag icon.
//
// Keep in mind that drag icons do not allow user input.
type DragIcon interface {
	gextras.Objector

	// Child gets the widget currently used as drag icon.
	Child() *WidgetClass
	// SetChild sets the widget to display as the drag icon.
	SetChild(child Widget)
}

// DragIconClass implements the DragIcon interface.
type DragIconClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
	NativeInterface
	RootInterface
}

var _ DragIcon = (*DragIconClass)(nil)

func wrapDragIcon(obj *externglib.Object) DragIcon {
	return &DragIconClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
		NativeInterface: NativeInterface{
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
		},
		RootInterface: RootInterface{
			Object: obj,
			NativeInterface: NativeInterface{
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					AccessibleInterface: AccessibleInterface{
						Object: obj,
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
					ConstraintTargetInterface: ConstraintTargetInterface{
						Object: obj,
					},
				},
			},
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
		},
	}
}

func marshalDragIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDragIcon(obj), nil
}

// Child gets the widget currently used as drag icon.
func (s *DragIconClass) Child() *WidgetClass {
	var _arg0 *C.GtkDragIcon // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkDragIcon)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_drag_icon_get_child(_arg0)

	var _widget *WidgetClass // out

	_widget = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*WidgetClass)

	return _widget
}

// SetChild sets the widget to display as the drag icon.
func (s *DragIconClass) SetChild(child Widget) {
	var _arg0 *C.GtkDragIcon // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkDragIcon)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&child).Native()))

	C.gtk_drag_icon_set_child(_arg0, _arg1)
}
