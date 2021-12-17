// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_popover_get_type()), F: marshalPopoverer},
	})
}

// PopoverOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PopoverOverrider interface {
	ActivateDefault()
	Closed()
}

// Popover: GtkPopover is a bubble-like context popup.
//
// !An example GtkPopover (popover.png)
//
// It is primarily meant to provide context-dependent information or options.
// Popovers are attached to a parent widget. By default, they point to the whole
// widget area, although this behavior can be changed with
// gtk.Popover.SetPointingTo().
//
// The position of a popover relative to the widget it is attached to can also
// be changed with gtk.Popover.SetPosition()
//
// By default, GtkPopover performs a grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Escape key being
// pressed). If no such modal behavior is desired on a popover,
// gtk.Popover.SetAutohide() may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// GtkPopover is often used to replace menus. The best was to do this is to use
// the gtk.PopoverMenu subclass which supports being populated from a GMenuModel
// with gtk.PopoverMenu.NewFromModel.
//
//    <section>
//      <attribute name="display-hint">horizontal-buttons</attribute>
//      <item>
//        <attribute name="label">Cut</attribute>
//        <attribute name="action">app.cut</attribute>
//        <attribute name="verb-icon">edit-cut-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Copy</attribute>
//        <attribute name="action">app.copy</attribute>
//        <attribute name="verb-icon">edit-copy-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Paste</attribute>
//        <attribute name="action">app.paste</attribute>
//        <attribute name="verb-icon">edit-paste-symbolic</attribute>
//      </item>
//    </section>
//
//
// CSS nodes
//
//    popover[.menu]
//    ├── arrow
//    ╰── contents.background
//        ╰── <child>
//
//
// The contents child node always gets the .background style class and the
// popover itself gets the .menu style class if the popover is menu-like (i.e.
// GtkPopoverMenu).
//
// Particular uses of GtkPopover, such as touch selection popups or magnifiers
// in GtkEntry or GtkTextView get style classes like .touch-selection or
// .magnifier to differentiate from plain popovers.
//
// When styling a popover directly, the popover node should usually not have any
// background. The visible part of the popover can have a shadow. To specify it
// in CSS, set the box-shadow of the contents node.
//
// Note that, in order to accomplish appropriate arrow visuals, GtkPopover uses
// custom drawing for the arrow node. This makes it possible for the arrow to
// change its shape dynamically, but it also limits the possibilities of styling
// it using CSS. In particular, the arrow gets drawn over the content node's
// border and shadow, so they look like one shape, which means that the border
// width of the content node and the arrow node should be the same. The arrow
// also does not support any border shape other than solid, no border-radius,
// only one border width (border-bottom-width is used) and no box-shadow.
type Popover struct {
	Widget

	*externglib.Object
	NativeSurface
	ShortcutManager
}

var (
	_ Widgetter           = (*Popover)(nil)
	_ externglib.Objector = (*Popover)(nil)
)

func wrapPopover(obj *externglib.Object) *Popover {
	return &Popover{
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
		Object: obj,
		NativeSurface: NativeSurface{
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
		},
		ShortcutManager: ShortcutManager{
			Object: obj,
		},
	}
}

func marshalPopoverer(p uintptr) (interface{}, error) {
	return wrapPopover(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPopover creates a new GtkPopover.
func NewPopover() *Popover {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_popover_new()

	var _popover *Popover // out

	_popover = wrapPopover(externglib.Take(unsafe.Pointer(_cret)))

	return _popover
}

// Autohide returns whether the popover is modal.
//
// See gtk.Popover.SetAutohide() for the implications of this.
func (popover *Popover) Autohide() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_autohide(_arg0)
	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CascadePopdown returns whether the popover will close after a modal child is
// closed.
func (popover *Popover) CascadePopdown() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_cascade_popdown(_arg0)
	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child widget of popover.
func (popover *Popover) Child() Widgetter {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_child(_arg0)
	runtime.KeepAlive(popover)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// HasArrow gets whether this popover is showing an arrow pointing at the widget
// that it is relative to.
func (popover *Popover) HasArrow() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_has_arrow(_arg0)
	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MnemonicsVisible gets whether mnemonics are visible.
func (popover *Popover) MnemonicsVisible() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_mnemonics_visible(_arg0)
	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Offset gets the offset previous set with gtk_popover_set_offset().
func (popover *Popover) Offset() (xOffset int, yOffset int) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.int         // in
	var _arg2 C.int         // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	C.gtk_popover_get_offset(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(popover)

	var _xOffset int // out
	var _yOffset int // out

	_xOffset = int(_arg1)
	_yOffset = int(_arg2)

	return _xOffset, _yOffset
}

// PointingTo gets the rectangle that the popover points to.
//
// If a rectangle to point to has been set, this function will return TRUE and
// fill in rect with such rectangle, otherwise it will return FALSE and fill in
// rect with the parent widget coordinates.
func (popover *Popover) PointingTo() (*gdk.Rectangle, bool) {
	var _arg0 *C.GtkPopover  // out
	var _arg1 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_pointing_to(_arg0, &_arg1)
	runtime.KeepAlive(popover)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// Position returns the preferred position of popover.
func (popover *Popover) Position() PositionType {
	var _arg0 *C.GtkPopover     // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_position(_arg0)
	runtime.KeepAlive(popover)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// Popdown pops popover down.
//
// This is different from a gtk.Widget.Hide() call in that it may show the
// popover with a transition. If you want to hide the popover without a
// transition, just use gtk.Widget.Hide().
func (popover *Popover) Popdown() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	C.gtk_popover_popdown(_arg0)
	runtime.KeepAlive(popover)
}

// Popup pops popover up.
//
// This is different from a gtk.Widget.Show() call in that it may show the
// popover with a transition(). If you want to show the popover without a
// transition, just use [methodGtk.Widget.show.
func (popover *Popover) Popup() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	C.gtk_popover_popup(_arg0)
	runtime.KeepAlive(popover)
}

// Present presents the popover to the user.
func (popover *Popover) Present() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	C.gtk_popover_present(_arg0)
	runtime.KeepAlive(popover)
}

// SetAutohide sets whether popover is modal.
//
// A modal popover will grab the keyboard focus on it when being displayed.
// Clicking outside the popover area or pressing Esc will dismiss the popover.
//
// Called this function on an already showing popup with a new autohide value
// different from the current one, will cause the popup to be hidden.
//
// The function takes the following parameters:
//
//    - autohide: TRUE to dismiss the popover on outside clicks.
//
func (popover *Popover) SetAutohide(autohide bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if autohide {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_autohide(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(autohide)
}

// SetCascadePopdown: if cascade_popdown is TRUE, the popover will be closed
// when a child modal popover is closed.
//
// If FALSE, popover will stay visible.
//
// The function takes the following parameters:
//
//    - cascadePopdown: TRUE if the popover should follow a child closing.
//
func (popover *Popover) SetCascadePopdown(cascadePopdown bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if cascadePopdown {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_cascade_popdown(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(cascadePopdown)
}

// SetChild sets the child widget of popover.
//
// The function takes the following parameters:
//
//    - child widget.
//
func (popover *Popover) SetChild(child Widgetter) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_popover_set_child(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(child)
}

// SetDefaultWidget sets the default widget of a GtkPopover.
//
// The default widget is the widget that’s activated when the user presses Enter
// in a dialog (for example). This function sets or unsets the default widget
// for a GtkPopover.
//
// The function takes the following parameters:
//
//    - widget: child widget of popover to set as the default, or NULL to unset
//    the default widget for the popover.
//
func (popover *Popover) SetDefaultWidget(widget Widgetter) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_popover_set_default_widget(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(widget)
}

// SetHasArrow sets whether this popover should draw an arrow pointing at the
// widget it is relative to.
//
// The function takes the following parameters:
//
//    - hasArrow: TRUE to draw an arrow.
//
func (popover *Popover) SetHasArrow(hasArrow bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if hasArrow {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_has_arrow(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(hasArrow)
}

// SetMnemonicsVisible sets whether mnemonics should be visible.
//
// The function takes the following parameters:
//
//    - mnemonicsVisible: new value.
//
func (popover *Popover) SetMnemonicsVisible(mnemonicsVisible bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if mnemonicsVisible {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_mnemonics_visible(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(mnemonicsVisible)
}

// SetOffset sets the offset to use when calculating the position of the
// popover.
//
// These values are used when preparing the gdk.PopupLayout for positioning the
// popover.
//
// The function takes the following parameters:
//
//    - xOffset: x offset to adjust the position by.
//    - yOffset: y offset to adjust the position by.
//
func (popover *Popover) SetOffset(xOffset, yOffset int) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.int         // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	_arg1 = C.int(xOffset)
	_arg2 = C.int(yOffset)

	C.gtk_popover_set_offset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(xOffset)
	runtime.KeepAlive(yOffset)
}

// SetPointingTo sets the rectangle that popover points to.
//
// This is in the coordinate space of the popover parent.
//
// The function takes the following parameters:
//
//    - rect: rectangle to point to.
//
func (popover *Popover) SetPointingTo(rect *gdk.Rectangle) {
	var _arg0 *C.GtkPopover   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect)))

	C.gtk_popover_set_pointing_to(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(rect)
}

// SetPosition sets the preferred position for popover to appear.
//
// If the popover is currently visible, it will be immediately updated.
//
// This preference will be respected where possible, although on lack of space
// (eg. if close to the window edges), the GtkPopover may choose to appear on
// the opposite side.
//
// The function takes the following parameters:
//
//    - position: preferred popover position.
//
func (popover *Popover) SetPosition(position PositionType) {
	var _arg0 *C.GtkPopover     // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_popover_set_position(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(position)
}

// ConnectActivateDefault: emitted whend the user activates the default widget.
//
// This is a keybinding signal (class.SignalAction.html).
func (popover *Popover) ConnectActivateDefault(f func()) externglib.SignalHandle {
	return popover.Connect("activate-default", f)
}

// ConnectClosed: emitted when the popover is closed.
func (popover *Popover) ConnectClosed(f func()) externglib.SignalHandle {
	return popover.Connect("closed", f)
}
