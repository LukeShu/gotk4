// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeOffscreenWindow = coreglib.Type(C.gtk_offscreen_window_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeOffscreenWindow, F: marshalOffscreenWindow},
	})
}

// OffscreenWindowOverrides contains methods that are overridable.
type OffscreenWindowOverrides struct {
}

func defaultOffscreenWindowOverrides(v *OffscreenWindow) OffscreenWindowOverrides {
	return OffscreenWindowOverrides{}
}

// OffscreenWindow is strictly intended to be used for obtaining snapshots of
// widgets that are not part of a normal widget hierarchy. Since OffscreenWindow
// is a toplevel widget you cannot obtain snapshots of a full window with it
// since you cannot pack a toplevel widget in another toplevel.
//
// The idea is to take a widget and manually set the state of it, add it to a
// GtkOffscreenWindow and then retrieve the snapshot as a #cairo_surface_t or
// Pixbuf.
//
// GtkOffscreenWindow derives from Window only as an implementation detail.
// Applications should not use any API specific to Window to operate on this
// object. It should be treated as a Bin that has no parent widget.
//
// When contained offscreen widgets are redrawn, GtkOffscreenWindow will emit a
// Widget::damage-event signal.
type OffscreenWindow struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Binner = (*OffscreenWindow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*OffscreenWindow, *OffscreenWindowClass, OffscreenWindowOverrides](
		GTypeOffscreenWindow,
		initOffscreenWindowClass,
		wrapOffscreenWindow,
		defaultOffscreenWindowOverrides,
	)
}

func initOffscreenWindowClass(gclass unsafe.Pointer, overrides OffscreenWindowOverrides, classInitFunc func(*OffscreenWindowClass)) {
	if classInitFunc != nil {
		class := (*OffscreenWindowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapOffscreenWindow(obj *coreglib.Object) *OffscreenWindow {
	return &OffscreenWindow{
		Window: Window{
			Bin: Bin{
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
			},
		},
	}
}

func marshalOffscreenWindow(p uintptr) (interface{}, error) {
	return wrapOffscreenWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewOffscreenWindow creates a toplevel container widget that is used to
// retrieve snapshots of widgets without showing them on the screen.
//
// The function returns the following values:
//
//   - offscreenWindow: pointer to a Widget.
//
func NewOffscreenWindow() *OffscreenWindow {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_offscreen_window_new()

	var _offscreenWindow *OffscreenWindow // out

	_offscreenWindow = wrapOffscreenWindow(coreglib.Take(unsafe.Pointer(_cret)))

	return _offscreenWindow
}

// Pixbuf retrieves a snapshot of the contained widget in the form of a Pixbuf.
// This is a new pixbuf with a reference count of 1, and the application should
// unreference it once it is no longer needed.
//
// The function returns the following values:
//
//   - pixbuf (optional) pointer, or NULL.
//
func (offscreen *OffscreenWindow) Pixbuf() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkOffscreenWindow // out
	var _cret *C.GdkPixbuf          // in

	_arg0 = (*C.GtkOffscreenWindow)(unsafe.Pointer(coreglib.InternObject(offscreen).Native()))

	_cret = C.gtk_offscreen_window_get_pixbuf(_arg0)
	runtime.KeepAlive(offscreen)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// Surface retrieves a snapshot of the contained widget in the form of a
// #cairo_surface_t. If you need to keep this around over window resizes then
// you should add a reference to it.
//
// The function returns the following values:
//
//   - surface (optional) pointer to the offscreen surface, or NULL.
//
func (offscreen *OffscreenWindow) Surface() *cairo.Surface {
	var _arg0 *C.GtkOffscreenWindow // out
	var _cret *C.cairo_surface_t    // in

	_arg0 = (*C.GtkOffscreenWindow)(unsafe.Pointer(coreglib.InternObject(offscreen).Native()))

	_cret = C.gtk_offscreen_window_get_surface(_arg0)
	runtime.KeepAlive(offscreen)

	var _surface *cairo.Surface // out

	if _cret != nil {
		_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
		C.cairo_surface_reference(_cret)
		runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
			C.cairo_surface_destroy((*C.cairo_surface_t)(unsafe.Pointer(v.Native())))
		})
	}

	return _surface
}

// OffscreenWindowClass: instance of this type is always passed by reference.
type OffscreenWindowClass struct {
	*offscreenWindowClass
}

// offscreenWindowClass is the struct that's finalized.
type offscreenWindowClass struct {
	native *C.GtkOffscreenWindowClass
}

// ParentClass: parent class.
func (o *OffscreenWindowClass) ParentClass() *WindowClass {
	valptr := &o.native.parent_class
	var _v *WindowClass // out
	_v = (*WindowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
