// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeDragIcon returns the GType for the type DragIcon.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDragIcon() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "DragIcon").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDragIcon)
	return gtype
}

// DragIconOverrider contains methods that are overridable.
type DragIconOverrider interface {
}

// DragIcon: GtkDragIcon is a GtkRoot implementation for drag icons.
//
// A drag icon moves with the pointer during a Drag-and-Drop operation and is
// destroyed when the drag ends.
//
// To set up a drag icon and associate it with an ongoing drag operation, use
// gtk.DragIcon().GetForDrag to get the icon for a drag. You can then use it
// like any other widget and use gtk.DragIcon.SetChild() to set whatever widget
// should be used for the drag icon.
//
// Keep in mind that drag icons do not allow user input.
type DragIcon struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Root
}

var (
	_ Widgetter         = (*DragIcon)(nil)
	_ coreglib.Objector = (*DragIcon)(nil)
)

func classInitDragIconner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDragIcon(obj *coreglib.Object) *DragIcon {
	return &DragIcon{
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
	}
}

func marshalDragIcon(p uintptr) (interface{}, error) {
	return wrapDragIcon(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child gets the widget currently used as drag icon.
//
// The function returns the following values:
//
//    - widget (optional): drag icon or NULL if none.
//
func (self *DragIcon) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "DragIcon").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetChild sets the widget to display as the drag icon.
//
// The function takes the following parameters:
//
//    - child (optional): GtkWidget or NULL.
//
func (self *DragIcon) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "DragIcon").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// DragIconCreateWidgetForValue creates a widget that can be used as a drag icon
// for the given value.
//
// Supported types include strings, GdkRGBA and GtkTextBuffer. If GTK does not
// know how to create a widget for a given value, it will return NULL.
//
// This method is used to set the default drag icon on drag'n'drop operations
// started by GtkDragSource, so you don't need to set a drag icon using this
// function there.
//
// The function takes the following parameters:
//
//    - value: GValue.
//
// The function returns the following values:
//
//    - widget (optional): new GtkWidget for displaying value as a drag icon.
//
func DragIconCreateWidgetForValue(value *coreglib.Value) Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(value.Native()))

	_gret := girepository.MustFind("Gtk", "create_widget_for_value").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(value)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// DragIconGetForDrag gets the GtkDragIcon in use with drag.
//
// If no drag icon exists yet, a new one will be created and shown.
//
// The function takes the following parameters:
//
//    - drag: GdkDrag.
//
// The function returns the following values:
//
//    - widget: GtkDragIcon.
//
func DragIconGetForDrag(drag gdk.Dragger) Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_gret := girepository.MustFind("Gtk", "get_for_drag").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(drag)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// DragIconSetFromPaintable creates a GtkDragIcon that shows paintable, and
// associates it with the drag operation.
//
// The hotspot position on the paintable is aligned with the hotspot of the
// cursor.
//
// The function takes the following parameters:
//
//    - drag: GdkDrag.
//    - paintable: GdkPaintable to display.
//    - hotX: x coordinate of the hotspot.
//    - hotY: y coordinate of the hotspot.
//
func DragIconSetFromPaintable(drag gdk.Dragger, paintable gdk.Paintabler, hotX, hotY int32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(hotX)
	*(*C.int)(unsafe.Pointer(&_args[3])) = C.int(hotY)

	girepository.MustFind("Gtk", "set_from_paintable").Invoke(_args[:], nil)

	runtime.KeepAlive(drag)
	runtime.KeepAlive(paintable)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}
