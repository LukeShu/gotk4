// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeActionBar = coreglib.Type(C.gtk_action_bar_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeActionBar, F: marshalActionBar},
	})
}

// ActionBar: GtkActionBar is designed to present contextual actions.
//
// !An example GtkActionBar (action-bar.png)
//
// It is expected to be displayed below the content and expand horizontally to
// fill the area.
//
// It allows placing children at the start or the end. In addition, it contains
// an internal centered box which is centered with respect to the full width
// of the box, even if the children at either side take up different amounts of
// space.
//
// # CSS nodes
//
// GtkActionBar has a single CSS node with name actionbar.
type ActionBar struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*ActionBar)(nil)
)

func wrapActionBar(obj *coreglib.Object) *ActionBar {
	return &ActionBar{
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
	}
}

func marshalActionBar(p uintptr) (interface{}, error) {
	return wrapActionBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewActionBar creates a new GtkActionBar widget.
//
// The function returns the following values:
//
//   - actionBar: new GtkActionBar.
//
func NewActionBar() *ActionBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_action_bar_new()

	var _actionBar *ActionBar // out

	_actionBar = wrapActionBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _actionBar
}

// CenterWidget retrieves the center bar widget of the bar.
//
// The function returns the following values:
//
//   - widget (optional): center GtkWidget or NULL.
//
func (actionBar *ActionBar) CenterWidget() Widgetter {
	var _arg0 *C.GtkActionBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))

	_cret = C.gtk_action_bar_get_center_widget(_arg0)
	runtime.KeepAlive(actionBar)

	var _widget Widgetter // out

	if _cret != nil {
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

// Revealed gets whether the contents of the action bar are revealed.
//
// The function returns the following values:
//
//   - ok: current value of the gtk.ActionBar:revealed property.
//
func (actionBar *ActionBar) Revealed() bool {
	var _arg0 *C.GtkActionBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))

	_cret = C.gtk_action_bar_get_revealed(_arg0)
	runtime.KeepAlive(actionBar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PackEnd adds child to action_bar, packed with reference to the end of the
// action_bar.
//
// The function takes the following parameters:
//
//   - child: GtkWidget to be added to action_bar.
//
func (actionBar *ActionBar) PackEnd(child Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_action_bar_pack_end(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// PackStart adds child to action_bar, packed with reference to the start of the
// action_bar.
//
// The function takes the following parameters:
//
//   - child: GtkWidget to be added to action_bar.
//
func (actionBar *ActionBar) PackStart(child Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_action_bar_pack_start(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// Remove removes a child from action_bar.
//
// The function takes the following parameters:
//
//   - child: GtkWidget to be removed.
//
func (actionBar *ActionBar) Remove(child Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_action_bar_remove(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// SetCenterWidget sets the center widget for the GtkActionBar.
//
// The function takes the following parameters:
//
//   - centerWidget (optional): widget to use for the center.
//
func (actionBar *ActionBar) SetCenterWidget(centerWidget Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	if centerWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(centerWidget).Native()))
	}

	C.gtk_action_bar_set_center_widget(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(centerWidget)
}

// SetRevealed reveals or conceals the content of the action bar.
//
// Note: this does not show or hide action_bar in the gtk.Widget:visible sense,
// so revealing has no effect if the action bar is hidden.
//
// The function takes the following parameters:
//
//   - revealed: new value of the property.
//
func (actionBar *ActionBar) SetRevealed(revealed bool) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(coreglib.InternObject(actionBar).Native()))
	if revealed {
		_arg1 = C.TRUE
	}

	C.gtk_action_bar_set_revealed(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(revealed)
}
