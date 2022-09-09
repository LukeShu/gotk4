// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeTextViewAccessible = coreglib.Type(C.gtk_text_view_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTextViewAccessible, F: marshalTextViewAccessible},
	})
}

// TextViewAccessibleOverrides contains methods that are overridable.
type TextViewAccessibleOverrides struct {
}

func defaultTextViewAccessibleOverrides(v *TextViewAccessible) TextViewAccessibleOverrides {
	return TextViewAccessibleOverrides{}
}

type TextViewAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	*coreglib.Object
	atk.EditableText
	atk.StreamableContent
	atk.Text
}

var (
	_ coreglib.Objector = (*TextViewAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TextViewAccessible, *TextViewAccessibleClass, TextViewAccessibleOverrides](
		GTypeTextViewAccessible,
		initTextViewAccessibleClass,
		wrapTextViewAccessible,
		defaultTextViewAccessibleOverrides,
	)
}

func initTextViewAccessibleClass(gclass unsafe.Pointer, overrides TextViewAccessibleOverrides, classInitFunc func(*TextViewAccessibleClass)) {
	if classInitFunc != nil {
		class := (*TextViewAccessibleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTextViewAccessible(obj *coreglib.Object) *TextViewAccessible {
	return &TextViewAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					AtkObject: atk.AtkObject{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
		Object: obj,
		EditableText: atk.EditableText{
			Object: obj,
		},
		StreamableContent: atk.StreamableContent{
			Object: obj,
		},
		Text: atk.Text{
			Object: obj,
		},
	}
}

func marshalTextViewAccessible(p uintptr) (interface{}, error) {
	return wrapTextViewAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// TextViewAccessibleClass: instance of this type is always passed by reference.
type TextViewAccessibleClass struct {
	*textViewAccessibleClass
}

// textViewAccessibleClass is the struct that's finalized.
type textViewAccessibleClass struct {
	native *C.GtkTextViewAccessibleClass
}

func (t *TextViewAccessibleClass) ParentClass() *ContainerAccessibleClass {
	valptr := &t.native.parent_class
	var _v *ContainerAccessibleClass // out
	_v = (*ContainerAccessibleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
