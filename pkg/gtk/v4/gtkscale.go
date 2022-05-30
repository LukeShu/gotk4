// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkscale.go.
var GTypeScale = coreglib.Type(C.gtk_scale_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeScale, F: marshalScale},
	})
}

// ScaleOverrider contains methods that are overridable.
type ScaleOverrider interface {
}

// Scale: GtkScale is a slider control used to select a numeric value.
//
// !An example GtkScale (scales.png)
//
// To use it, you’ll probably want to investigate the methods on its base class,
// gtkrange, in addition to the methods for GtkScale itself. To set the value of
// a scale, you would normally use gtk.Range.SetValue(). To detect changes to
// the value, you would normally use the gtk.Range::value-changed signal.
//
// Note that using the same upper and lower bounds for the GtkScale (through the
// GtkRange methods) will hide the slider itself. This is useful for
// applications that want to show an undeterminate value on the scale, without
// changing the layout of the application (such as movie or music players).
//
//
// GtkScale as GtkBuildable
//
// GtkScale supports a custom <marks> element, which can contain multiple
// <mark\> elements. The “value” and “position” attributes have the same meaning
// as gtk.Scale.AddMark() parameters of the same name. If the element is not
// empty, its content is taken as the markup to show at the mark. It can be
// translated with the usual ”translatable” and “context” attributes.
//
// CSS nodes
//
//    scale[.fine-tune][.marks-before][.marks-after]
//    ├── [value][.top][.right][.bottom][.left]
//    ├── marks.top
//    │   ├── mark
//    │   ┊    ├── [label]
//    │   ┊    ╰── indicator
//    ┊   ┊
//    │   ╰── mark
//    ├── marks.bottom
//    │   ├── mark
//    │   ┊    ├── indicator
//    │   ┊    ╰── [label]
//    ┊   ┊
//    │   ╰── mark
//    ╰── trough
//        ├── [fill]
//        ├── [highlight]
//        ╰── slider
//
//
// GtkScale has a main CSS node with name scale and a subnode for its contents,
// with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scale is in
// 'fine-tuning' mode.
//
// If the scale has an origin (see gtk.Scale.SetHasOrigin()), there is a subnode
// with name highlight below the trough node that is used for rendering the
// highlighted part of the trough.
//
// If the scale is showing a fill level (see gtk.Range.SetShowFillLevel()),
// there is a subnode with name fill below the trough node that is used for
// rendering the filled in part of the trough.
//
// If marks are present, there is a marks subnode before or after the trough
// node, below which each mark gets a node with name mark. The marks nodes get
// either the .top or .bottom style class.
//
// The mark node has a subnode named indicator. If the mark has text, it also
// has a subnode named label. When the mark is either above or left of the
// scale, the label subnode is the first when present. Otherwise, the indicator
// subnode is the first.
//
// The main CSS node gets the 'marks-before' and/or 'marks-after' style classes
// added depending on what marks are present.
//
// If the scale is displaying the value (see gtk.Scale:draw-value), there is
// subnode with name value. This node will get the .top or .bottom style classes
// similar to the marks node.
//
//
// Accessibility
//
// GtkScale uses the GTK_ACCESSIBLE_ROLE_SLIDER role.
type Scale struct {
	_ [0]func() // equal guard
	Range
}

var (
	_ Widgetter         = (*Scale)(nil)
	_ coreglib.Objector = (*Scale)(nil)
)

func classInitScaler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapScale(obj *coreglib.Object) *Scale {
	return &Scale{
		Range: Range{
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
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalScale(p uintptr) (interface{}, error) {
	return wrapScale(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ClearMarks removes any marks that have been added.
func (scale *Scale) ClearMarks() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	*(**Scale)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Scale").InvokeMethod("clear_marks", args[:], nil)

	runtime.KeepAlive(scale)
}

// DrawValue returns whether the current value is displayed as a string next to
// the slider.
//
// The function returns the following values:
//
//    - ok: whether the current value is displayed as a string.
//
func (scale *Scale) DrawValue() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	*(**Scale)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Scale").InvokeMethod("get_draw_value", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scale)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasOrigin returns whether the scale has an origin.
//
// The function returns the following values:
//
//    - ok: TRUE if the scale has an origin.
//
func (scale *Scale) HasOrigin() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	*(**Scale)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Scale").InvokeMethod("get_has_origin", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scale)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Layout gets the PangoLayout used to display the scale.
//
// The returned object is owned by the scale so does not need to be freed by the
// caller.
//
// The function returns the following values:
//
//    - layout (optional): pango.Layout for this scale, or NULL if the
//      gtkscale:draw-value property is FALSE.
//
func (scale *Scale) Layout() *pango.Layout {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	*(**Scale)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Scale").InvokeMethod("get_layout", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scale)

	var _layout *pango.Layout // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_layout = &pango.Layout{
				Object: obj,
			}
		}
	}

	return _layout
}

// SetDrawValue specifies whether the current value is displayed as a string
// next to the slider.
//
// The function takes the following parameters:
//
//    - drawValue: TRUE to draw the value.
//
func (scale *Scale) SetDrawValue(drawValue bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	if drawValue {
		_arg1 = C.TRUE
	}
	*(**Scale)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Scale").InvokeMethod("set_draw_value", args[:], nil)

	runtime.KeepAlive(scale)
	runtime.KeepAlive(drawValue)
}

// SetHasOrigin sets whether the scale has an origin.
//
// If gtkscale:has-origin is set to TRUE (the default), the scale will highlight
// the part of the trough between the origin (bottom or left side) and the
// current value.
//
// The function takes the following parameters:
//
//    - hasOrigin: TRUE if the scale has an origin.
//
func (scale *Scale) SetHasOrigin(hasOrigin bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	if hasOrigin {
		_arg1 = C.TRUE
	}
	*(**Scale)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Scale").InvokeMethod("set_has_origin", args[:], nil)

	runtime.KeepAlive(scale)
	runtime.KeepAlive(hasOrigin)
}
