// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_fixed_get_type()), F: marshalFixeder},
	})
}

// Fixeder describes Fixed's methods.
type Fixeder interface {
	// Move moves a child of a Fixed container to the given position.
	Move(widget Widgeter, x int, y int)
	// Put adds a widget to a Fixed container at the given position.
	Put(widget Widgeter, x int, y int)
}

// Fixed widget is a container which can place child widgets at fixed positions
// and with fixed sizes, given in pixels. Fixed performs no automatic layout
// management.
//
// For most applications, you should not use this container! It keeps you from
// having to learn about the other GTK+ containers, but it results in broken
// applications. With Fixed, the following things will result in truncated text,
// overlapping widgets, and other display bugs:
//
// - Themes, which may change widget sizes.
//
// - Fonts other than the one you used to write the app will of course change
// the size of widgets containing text; keep in mind that users may use a larger
// font because of difficulty reading the default, or they may be using a
// different OS that provides different fonts.
//
// - Translation of text into other languages changes its size. Also, display of
// non-English text will use a different font in many cases.
//
// In addition, Fixed does not pay attention to text direction and thus may
// produce unwanted results if your app is run under right-to-left languages
// such as Hebrew or Arabic. That is: normally GTK+ will order containers
// appropriately for the text direction, e.g. to put labels to the right of the
// thing they label when using an RTL language, but it can’t do that with Fixed.
// So if you need to reorder widgets depending on the text direction, you would
// need to manually detect it and adjust child positions accordingly.
//
// Finally, fixed positioning makes it kind of annoying to add/remove GUI
// elements, since you have to reposition all the other elements. This is a
// long-term maintenance problem for your application.
//
// If you know none of these things are an issue for your application, and
// prefer the simplicity of Fixed, by all means use the widget. But you should
// be aware of the tradeoffs.
//
// See also Layout, which shares the ability to perform fixed positioning of
// child widgets and additionally adds custom drawing and scrollability.
type Fixed struct {
	Container
}

var (
	_ Fixeder         = (*Fixed)(nil)
	_ gextras.Nativer = (*Fixed)(nil)
)

func wrapFixed(obj *externglib.Object) Fixeder {
	return &Fixed{
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
			},
		},
	}
}

func marshalFixeder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFixed(obj), nil
}

// NewFixed creates a new Fixed.
func NewFixed() *Fixed {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_fixed_new()

	var _fixed *Fixed // out

	_fixed = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Fixed)

	return _fixed
}

// Move moves a child of a Fixed container to the given position.
func (fixed *Fixed) Move(widget Widgeter, x int, y int) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_fixed_move(_arg0, _arg1, _arg2, _arg3)
}

// Put adds a widget to a Fixed container at the given position.
func (fixed *Fixed) Put(widget Widgeter, x int, y int) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_fixed_put(_arg0, _arg1, _arg2, _arg3)
}

type FixedChild struct {
	native C.GtkFixedChild
}

// Native returns the underlying C source pointer.
func (f *FixedChild) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}
