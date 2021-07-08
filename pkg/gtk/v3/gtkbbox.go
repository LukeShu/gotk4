// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_box_style_get_type()), F: marshalButtonBoxStyle},
		{T: externglib.Type(C.gtk_button_box_get_type()), F: marshalButtonBox},
	})
}

// ButtonBoxStyle: used to dictate the style that a ButtonBox uses to layout the
// buttons it contains.
type ButtonBoxStyle int

const (
	// Spread buttons are evenly spread across the box.
	ButtonBoxStyleSpread ButtonBoxStyle = 1
	// Edge buttons are placed at the edges of the box.
	ButtonBoxStyleEdge ButtonBoxStyle = 2
	// Start buttons are grouped towards the start of the box, (on the left for
	// a HBox, or the top for a VBox).
	ButtonBoxStyleStart ButtonBoxStyle = 3
	// End buttons are grouped towards the end of the box, (on the right for a
	// HBox, or the bottom for a VBox).
	ButtonBoxStyleEnd ButtonBoxStyle = 4
	// Center buttons are centered in the box. Since 2.12.
	ButtonBoxStyleCenter ButtonBoxStyle = 5
	// Expand buttons expand to fill the box. This entails giving buttons a
	// "linked" appearance, making button sizes homogeneous, and setting spacing
	// to 0 (same as calling gtk_box_set_homogeneous() and gtk_box_set_spacing()
	// manually). Since 3.12.
	ButtonBoxStyleExpand ButtonBoxStyle = 6
)

func marshalButtonBoxStyle(p uintptr) (interface{}, error) {
	return ButtonBoxStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type ButtonBox interface {
	gextras.Objector

	// ChildNonHomogeneous returns whether the child is exempted from homogenous
	// sizing.
	ChildNonHomogeneous(child Widget) bool
	// ChildSecondary returns whether @child should appear in a secondary group
	// of children.
	ChildSecondary(child Widget) bool
	// Layout retrieves the method being used to arrange the buttons in a button
	// box.
	Layout() ButtonBoxStyle
	// SetChildNonHomogeneous sets whether the child is exempted from homogeous
	// sizing.
	SetChildNonHomogeneous(child Widget, nonHomogeneous bool)
	// SetChildSecondary sets whether @child should appear in a secondary group
	// of children. A typical use of a secondary child is the help button in a
	// dialog.
	//
	// This group appears after the other children if the style is
	// GTK_BUTTONBOX_START, GTK_BUTTONBOX_SPREAD or GTK_BUTTONBOX_EDGE, and
	// before the other children if the style is GTK_BUTTONBOX_END. For
	// horizontal button boxes, the definition of before/after depends on
	// direction of the widget (see gtk_widget_set_direction()). If the style is
	// GTK_BUTTONBOX_START or GTK_BUTTONBOX_END, then the secondary children are
	// aligned at the other end of the button box from the main children. For
	// the other styles, they appear immediately next to the main children.
	SetChildSecondary(child Widget, isSecondary bool)
	// SetLayout changes the way buttons are arranged in their container.
	SetLayout(layoutStyle ButtonBoxStyle)
}

// ButtonBoxClass implements the ButtonBox interface.
type ButtonBoxClass struct {
	*externglib.Object
	BoxClass
	BuildableInterface
	OrientableInterface
}

var _ ButtonBox = (*ButtonBoxClass)(nil)

func wrapButtonBox(obj *externglib.Object) ButtonBox {
	return &ButtonBoxClass{
		Object: obj,
		BoxClass: BoxClass{
			Object: obj,
			ContainerClass: ContainerClass{
				Object: obj,
				WidgetClass: WidgetClass{
					Object:           obj,
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			OrientableInterface: OrientableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalButtonBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapButtonBox(obj), nil
}

// NewButtonBox creates a new ButtonBox.
func NewButtonBox(orientation Orientation) ButtonBox {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_button_box_new(_arg1)

	var _buttonBox ButtonBox // out

	_buttonBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ButtonBox)

	return _buttonBox
}

// ChildNonHomogeneous returns whether the child is exempted from homogenous
// sizing.
func (w *ButtonBoxClass) ChildNonHomogeneous(child Widget) bool {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_button_box_get_child_non_homogeneous(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildSecondary returns whether @child should appear in a secondary group of
// children.
func (w *ButtonBoxClass) ChildSecondary(child Widget) bool {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_button_box_get_child_secondary(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Layout retrieves the method being used to arrange the buttons in a button
// box.
func (w *ButtonBoxClass) Layout() ButtonBoxStyle {
	var _arg0 *C.GtkButtonBox     // out
	var _cret C.GtkButtonBoxStyle // in

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_button_box_get_layout(_arg0)

	var _buttonBoxStyle ButtonBoxStyle // out

	_buttonBoxStyle = ButtonBoxStyle(_cret)

	return _buttonBoxStyle
}

// SetChildNonHomogeneous sets whether the child is exempted from homogeous
// sizing.
func (w *ButtonBoxClass) SetChildNonHomogeneous(child Widget, nonHomogeneous bool) {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if nonHomogeneous {
		_arg2 = C.TRUE
	}

	C.gtk_button_box_set_child_non_homogeneous(_arg0, _arg1, _arg2)
}

// SetChildSecondary sets whether @child should appear in a secondary group of
// children. A typical use of a secondary child is the help button in a dialog.
//
// This group appears after the other children if the style is
// GTK_BUTTONBOX_START, GTK_BUTTONBOX_SPREAD or GTK_BUTTONBOX_EDGE, and before
// the other children if the style is GTK_BUTTONBOX_END. For horizontal button
// boxes, the definition of before/after depends on direction of the widget (see
// gtk_widget_set_direction()). If the style is GTK_BUTTONBOX_START or
// GTK_BUTTONBOX_END, then the secondary children are aligned at the other end
// of the button box from the main children. For the other styles, they appear
// immediately next to the main children.
func (w *ButtonBoxClass) SetChildSecondary(child Widget, isSecondary bool) {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if isSecondary {
		_arg2 = C.TRUE
	}

	C.gtk_button_box_set_child_secondary(_arg0, _arg1, _arg2)
}

// SetLayout changes the way buttons are arranged in their container.
func (w *ButtonBoxClass) SetLayout(layoutStyle ButtonBoxStyle) {
	var _arg0 *C.GtkButtonBox     // out
	var _arg1 C.GtkButtonBoxStyle // out

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(w.Native()))
	_arg1 = C.GtkButtonBoxStyle(layoutStyle)

	C.gtk_button_box_set_layout(_arg0, _arg1)
}
