// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_Overlay_ConnectGetChildPosition(gpointer, GtkWidget*, GdkRectangle*, guintptr);
import "C"

// glib.Type values for gtkoverlay.go.
var GTypeOverlay = externglib.Type(C.gtk_overlay_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeOverlay, F: marshalOverlay},
	})
}

// Overlay: GtkOverlay is a container which contains a single main child, on top
// of which it can place “overlay” widgets.
//
// !An example GtkOverlay (overlay.png)
//
// The position of each overlay widget is determined by its gtk.Widget:halign
// and gtk.Widget:valign properties. E.g. a widget with both alignments set to
// GTK_ALIGN_START will be placed at the top left corner of the GtkOverlay
// container, whereas an overlay with halign set to GTK_ALIGN_CENTER and valign
// set to GTK_ALIGN_END will be placed a the bottom edge of the GtkOverlay,
// horizontally centered. The position can be adjusted by setting the margin
// properties of the child to non-zero values.
//
// More complicated placement of overlays is possible by connecting to the
// gtk.Overlay::get-child-position signal.
//
// An overlay’s minimum and natural sizes are those of its main child. The sizes
// of overlay children are not considered when measuring these preferred sizes.
//
//
// GtkOverlay as GtkBuildable
//
// The GtkOverlay implementation of the GtkBuildable interface supports placing
// a child as an overlay by specifying “overlay” as the “type” attribute of a
// <child> element.
//
//
// CSS nodes
//
// GtkOverlay has a single CSS node with the name “overlay”. Overlay children
// whose alignments cause them to be positioned at an edge get the style classes
// “.left”, “.right”, “.top”, and/or “.bottom” according to their position.
type Overlay struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Overlay)(nil)
)

func wrapOverlay(obj *externglib.Object) *Overlay {
	return &Overlay{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalOverlay(p uintptr) (interface{}, error) {
	return wrapOverlay(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Overlay_ConnectGetChildPosition
func _gotk4_gtk4_Overlay_ConnectGetChildPosition(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 *C.GdkRectangle, arg3 C.guintptr) (cret C.gboolean) {
	var f func(widget Widgetter) (allocation *gdk.Rectangle, ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(widget Widgetter) (allocation *gdk.Rectangle, ok bool))
	}

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	allocation, ok := f(_widget)

	*arg2 = *(*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(allocation)))
	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectGetChildPosition: emitted to determine the position and size of any
// overlay child widgets.
//
// A handler for this signal should fill allocation with the desired position
// and size for widget, relative to the 'main' child of overlay.
//
// The default handler for this signal uses the widget's halign and valign
// properties to determine the position and gives the widget its natural size
// (except that an alignment of GTK_ALIGN_FILL will cause the overlay to be
// full-width/height). If the main child is a GtkScrolledWindow, the overlays
// are placed relative to its contents.
func (overlay *Overlay) ConnectGetChildPosition(f func(widget Widgetter) (allocation *gdk.Rectangle, ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(overlay, "get-child-position", false, unsafe.Pointer(C._gotk4_gtk4_Overlay_ConnectGetChildPosition), f)
}

// NewOverlay creates a new GtkOverlay.
//
// The function returns the following values:
//
//    - overlay: new GtkOverlay object.
//
func NewOverlay() *Overlay {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_overlay_new()

	var _overlay *Overlay // out

	_overlay = wrapOverlay(externglib.Take(unsafe.Pointer(_cret)))

	return _overlay
}

// AddOverlay adds widget to overlay.
//
// The widget will be stacked on top of the main widget added with
// gtk.Overlay.SetChild().
//
// The position at which widget is placed is determined from its
// gtk.Widget:halign and gtk.Widget:valign properties.
//
// The function takes the following parameters:
//
//    - widget: GtkWidget to be added to the container.
//
func (overlay *Overlay) AddOverlay(widget Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_overlay_add_overlay(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
}

// Child gets the child widget of overlay.
//
// The function returns the following values:
//
//    - widget (optional): child widget of overlay.
//
func (overlay *Overlay) Child() Widgetter {
	var _arg0 *C.GtkOverlay // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))

	_cret = C.gtk_overlay_get_child(_arg0)
	runtime.KeepAlive(overlay)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// ClipOverlay gets whether widget should be clipped within the parent.
//
// The function takes the following parameters:
//
//    - widget: overlay child of GtkOverlay.
//
// The function returns the following values:
//
//    - ok: whether the widget is clipped within the parent.
//
func (overlay *Overlay) ClipOverlay(widget Widgetter) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	_cret = C.gtk_overlay_get_clip_overlay(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MeasureOverlay gets whether widget's size is included in the measurement of
// overlay.
//
// The function takes the following parameters:
//
//    - widget: overlay child of GtkOverlay.
//
// The function returns the following values:
//
//    - ok: whether the widget is measured.
//
func (overlay *Overlay) MeasureOverlay(widget Widgetter) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	_cret = C.gtk_overlay_get_measure_overlay(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveOverlay removes an overlay that was added with
// gtk_overlay_add_overlay().
//
// The function takes the following parameters:
//
//    - widget: GtkWidget to be removed.
//
func (overlay *Overlay) RemoveOverlay(widget Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.gtk_overlay_remove_overlay(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
}

// SetChild sets the child widget of overlay.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (overlay *Overlay) SetChild(child Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))
	}

	C.gtk_overlay_set_child(_arg0, _arg1)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(child)
}

// SetClipOverlay sets whether widget should be clipped within the parent.
//
// The function takes the following parameters:
//
//    - widget: overlay child of GtkOverlay.
//    - clipOverlay: whether the child should be clipped.
//
func (overlay *Overlay) SetClipOverlay(widget Widgetter, clipOverlay bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	if clipOverlay {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_clip_overlay(_arg0, _arg1, _arg2)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(clipOverlay)
}

// SetMeasureOverlay sets whether widget is included in the measured size of
// overlay.
//
// The overlay will request the size of the largest child that has this property
// set to TRUE. Children who are not included may be drawn outside of overlay's
// allocation if they are too large.
//
// The function takes the following parameters:
//
//    - widget: overlay child of GtkOverlay.
//    - measure: whether the child should be measured.
//
func (overlay *Overlay) SetMeasureOverlay(widget Widgetter, measure bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(externglib.InternObject(overlay).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	if measure {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_measure_overlay(_arg0, _arg1, _arg2)
	runtime.KeepAlive(overlay)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(measure)
}
