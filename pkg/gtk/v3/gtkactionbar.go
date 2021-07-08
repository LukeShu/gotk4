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
		{T: externglib.Type(C.gtk_action_bar_get_type()), F: marshalActionBar},
	})
}

// ActionBar is designed to present contextual actions. It is expected to be
// displayed below the content and expand horizontally to fill the area.
//
// It allows placing children at the start or the end. In addition, it contains
// an internal centered box which is centered with respect to the full width of
// the box, even if the children at either side take up different amounts of
// space.
//
//
// CSS nodes
//
// GtkActionBar has a single CSS node with name actionbar.
type ActionBar interface {
	gextras.Objector

	// CenterWidget retrieves the center bar widget of the bar.
	CenterWidget() Widget
	// PackEnd adds @child to @action_bar, packed with reference to the end of
	// the @action_bar.
	PackEnd(child Widget)
	// PackStart adds @child to @action_bar, packed with reference to the start
	// of the @action_bar.
	PackStart(child Widget)
	// SetCenterWidget sets the center widget for the ActionBar.
	SetCenterWidget(centerWidget Widget)
}

// ActionBarClass implements the ActionBar interface.
type ActionBarClass struct {
	*externglib.Object
	BinClass
	BuildableInterface
}

var _ ActionBar = (*ActionBarClass)(nil)

func wrapActionBar(obj *externglib.Object) ActionBar {
	return &ActionBarClass{
		Object: obj,
		BinClass: BinClass{
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
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalActionBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActionBar(obj), nil
}

// NewActionBar creates a new ActionBar widget.
func NewActionBar() ActionBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_action_bar_new()

	var _actionBar ActionBar // out

	_actionBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ActionBar)

	return _actionBar
}

// CenterWidget retrieves the center bar widget of the bar.
func (a *ActionBarClass) CenterWidget() Widget {
	var _arg0 *C.GtkActionBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_bar_get_center_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

// PackEnd adds @child to @action_bar, packed with reference to the end of the
// @action_bar.
func (a *ActionBarClass) PackEnd(child Widget) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_action_bar_pack_end(_arg0, _arg1)
}

// PackStart adds @child to @action_bar, packed with reference to the start of
// the @action_bar.
func (a *ActionBarClass) PackStart(child Widget) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_action_bar_pack_start(_arg0, _arg1)
}

// SetCenterWidget sets the center widget for the ActionBar.
func (a *ActionBarClass) SetCenterWidget(centerWidget Widget) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(centerWidget.Native()))

	C.gtk_action_bar_set_center_widget(_arg0, _arg1)
}
