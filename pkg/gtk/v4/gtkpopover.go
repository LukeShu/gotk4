// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_PopoverClass_activate_default(GtkPopover*);
// extern void _gotk4_gtk4_PopoverClass_closed(GtkPopover*);
// extern void _gotk4_gtk4_Popover_ConnectActivateDefault(gpointer, guintptr);
// extern void _gotk4_gtk4_Popover_ConnectClosed(gpointer, guintptr);
import "C"

// glib.Type values for gtkpopover.go.
var GTypePopover = coreglib.Type(C.gtk_popover_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePopover, F: marshalPopover},
	})
}

// PopoverOverrider contains methods that are overridable.
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
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	NativeSurface
	ShortcutManager
}

var (
	_ Widgetter         = (*Popover)(nil)
	_ coreglib.Objector = (*Popover)(nil)
)

func classInitPopoverer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkPopoverClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkPopoverClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ActivateDefault() }); ok {
		pclass.activate_default = (*[0]byte)(C._gotk4_gtk4_PopoverClass_activate_default)
	}

	if _, ok := goval.(interface{ Closed() }); ok {
		pclass.closed = (*[0]byte)(C._gotk4_gtk4_PopoverClass_closed)
	}
}

//export _gotk4_gtk4_PopoverClass_activate_default
func _gotk4_gtk4_PopoverClass_activate_default(arg0 *C.GtkPopover) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActivateDefault() })

	iface.ActivateDefault()
}

//export _gotk4_gtk4_PopoverClass_closed
func _gotk4_gtk4_PopoverClass_closed(arg0 *C.GtkPopover) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Closed() })

	iface.Closed()
}

func wrapPopover(obj *coreglib.Object) *Popover {
	return &Popover{
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
		ShortcutManager: ShortcutManager{
			Object: obj,
		},
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	return wrapPopover(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Popover_ConnectActivateDefault
func _gotk4_gtk4_Popover_ConnectActivateDefault(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectActivateDefault is emitted whend the user activates the default
// widget.
//
// This is a keybinding signal (class.SignalAction.html).
func (popover *Popover) ConnectActivateDefault(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(popover, "activate-default", false, unsafe.Pointer(C._gotk4_gtk4_Popover_ConnectActivateDefault), f)
}

//export _gotk4_gtk4_Popover_ConnectClosed
func _gotk4_gtk4_Popover_ConnectClosed(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClosed is emitted when the popover is closed.
func (popover *Popover) ConnectClosed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(popover, "closed", false, unsafe.Pointer(C._gotk4_gtk4_Popover_ConnectClosed), f)
}

// NewPopover creates a new GtkPopover.
//
// The function returns the following values:
//
//    - popover: new GtkPopover.
//
func NewPopover() *Popover {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("new_Popover", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _popover *Popover // out

	_popover = wrapPopover(coreglib.Take(unsafe.Pointer(_cret)))

	return _popover
}

// Autohide returns whether the popover is modal.
//
// See gtk.Popover.SetAutohide() for the implications of this.
//
// The function returns the following values:
//
//    - ok: TRUE if popover is modal.
//
func (popover *Popover) Autohide() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_autohide", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CascadePopdown returns whether the popover will close after a modal child is
// closed.
//
// The function returns the following values:
//
//    - ok: TRUE if popover will close after a modal child.
//
func (popover *Popover) CascadePopdown() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_cascade_popdown", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child widget of popover.
//
// The function returns the following values:
//
//    - widget (optional): child widget of popover.
//
func (popover *Popover) Child() Widgetter {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_child", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

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

// HasArrow gets whether this popover is showing an arrow pointing at the widget
// that it is relative to.
//
// The function returns the following values:
//
//    - ok: whether the popover has an arrow.
//
func (popover *Popover) HasArrow() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_has_arrow", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MnemonicsVisible gets whether mnemonics are visible.
//
// The function returns the following values:
//
//    - ok: TRUE if mnemonics are supposed to be visible in this popover.
//
func (popover *Popover) MnemonicsVisible() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_mnemonics_visible", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown pops popover down.
//
// This is different from a gtk.Widget.Hide() call in that it may show the
// popover with a transition. If you want to hide the popover without a
// transition, just use gtk.Widget.Hide().
func (popover *Popover) Popdown() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Popover").InvokeMethod("popdown", args[:], nil)

	runtime.KeepAlive(popover)
}

// Popup pops popover up.
//
// This is different from a gtk.Widget.Show() call in that it may show the
// popover with a transition(). If you want to show the popover without a
// transition, just use [methodGtk.Widget.show.
func (popover *Popover) Popup() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Popover").InvokeMethod("popup", args[:], nil)

	runtime.KeepAlive(popover)
}

// Present presents the popover to the user.
func (popover *Popover) Present() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**Popover)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Popover").InvokeMethod("present", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if autohide {
		_arg1 = C.TRUE
	}
	*(**Popover)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_autohide", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if cascadePopdown {
		_arg1 = C.TRUE
	}
	*(**Popover)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_cascade_popdown", args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(cascadePopdown)
}

// SetChild sets the child widget of popover.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (popover *Popover) SetChild(child Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if child != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}
	*(**Popover)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_child", args[:], nil)

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
//    - widget (optional): child widget of popover to set as the default, or NULL
//      to unset the default widget for the popover.
//
func (popover *Popover) SetDefaultWidget(widget Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if widget != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}
	*(**Popover)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_default_widget", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if hasArrow {
		_arg1 = C.TRUE
	}
	*(**Popover)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_has_arrow", args[:], nil)

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
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if mnemonicsVisible {
		_arg1 = C.TRUE
	}
	*(**Popover)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_mnemonics_visible", args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(mnemonicsVisible)
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
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(rect)))
	*(**Popover)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_pointing_to", args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(rect)
}
