// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
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
	GTypeRevealerTransitionType = coreglib.Type(C.gtk_revealer_transition_type_get_type())
	GTypeRevealer               = coreglib.Type(C.gtk_revealer_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRevealerTransitionType, F: marshalRevealerTransitionType},
		coreglib.TypeMarshaler{T: GTypeRevealer, F: marshalRevealer},
	})
}

// RevealerTransitionType: these enumeration values describe the possible
// transitions when the child of a GtkRevealer widget is shown or hidden.
type RevealerTransitionType C.gint

const (
	// RevealerTransitionTypeNone: no transition.
	RevealerTransitionTypeNone RevealerTransitionType = iota
	// RevealerTransitionTypeCrossfade: fade in.
	RevealerTransitionTypeCrossfade
	// RevealerTransitionTypeSlideRight: slide in from the left.
	RevealerTransitionTypeSlideRight
	// RevealerTransitionTypeSlideLeft: slide in from the right.
	RevealerTransitionTypeSlideLeft
	// RevealerTransitionTypeSlideUp: slide in from the bottom.
	RevealerTransitionTypeSlideUp
	// RevealerTransitionTypeSlideDown: slide in from the top.
	RevealerTransitionTypeSlideDown
	// RevealerTransitionTypeSwingRight: floop in from the left.
	RevealerTransitionTypeSwingRight
	// RevealerTransitionTypeSwingLeft: floop in from the right.
	RevealerTransitionTypeSwingLeft
	// RevealerTransitionTypeSwingUp: floop in from the bottom.
	RevealerTransitionTypeSwingUp
	// RevealerTransitionTypeSwingDown: floop in from the top.
	RevealerTransitionTypeSwingDown
)

func marshalRevealerTransitionType(p uintptr) (interface{}, error) {
	return RevealerTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RevealerTransitionType.
func (r RevealerTransitionType) String() string {
	switch r {
	case RevealerTransitionTypeNone:
		return "None"
	case RevealerTransitionTypeCrossfade:
		return "Crossfade"
	case RevealerTransitionTypeSlideRight:
		return "SlideRight"
	case RevealerTransitionTypeSlideLeft:
		return "SlideLeft"
	case RevealerTransitionTypeSlideUp:
		return "SlideUp"
	case RevealerTransitionTypeSlideDown:
		return "SlideDown"
	case RevealerTransitionTypeSwingRight:
		return "SwingRight"
	case RevealerTransitionTypeSwingLeft:
		return "SwingLeft"
	case RevealerTransitionTypeSwingUp:
		return "SwingUp"
	case RevealerTransitionTypeSwingDown:
		return "SwingDown"
	default:
		return fmt.Sprintf("RevealerTransitionType(%d)", r)
	}
}

// Revealer: GtkRevealer animates the transition of its child from invisible to
// visible.
//
// The style of transition can be controlled with
// gtk.Revealer.SetTransitionType().
//
// These animations respect the gtk.Settings:gtk-enable-animations setting.
//
// # CSS nodes
//
// GtkRevealer has a single CSS node with name revealer. When styling
// GtkRevealer using CSS, remember that it only hides its contents, not itself.
// That means applied margin, padding and borders will be visible even when the
// gtk.Revealer:reveal-child property is set to FALSE.
//
// # Accessibility
//
// GtkRevealer uses the GTK_ACCESSIBLE_ROLE_GROUP role.
//
// The child of GtkRevealer, if set, is always available in the accessibility
// tree, regardless of the state of the revealer widget.
type Revealer struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Revealer)(nil)
)

func wrapRevealer(obj *coreglib.Object) *Revealer {
	return &Revealer{
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

func marshalRevealer(p uintptr) (interface{}, error) {
	return wrapRevealer(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRevealer creates a new GtkRevealer.
//
// The function returns the following values:
//
//   - revealer: newly created GtkRevealer.
//
func NewRevealer() *Revealer {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_revealer_new()

	var _revealer *Revealer // out

	_revealer = wrapRevealer(coreglib.Take(unsafe.Pointer(_cret)))

	return _revealer
}

// Child gets the child widget of revealer.
//
// The function returns the following values:
//
//   - widget (optional): child widget of revealer.
//
func (revealer *Revealer) Child() Widgetter {
	var _arg0 *C.GtkRevealer // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_child(_arg0)
	runtime.KeepAlive(revealer)

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

// ChildRevealed returns whether the child is fully revealed.
//
// In other words, this returns whether the transition to the revealed state is
// completed.
//
// The function returns the following values:
//
//   - ok: TRUE if the child is fully revealed.
//
func (revealer *Revealer) ChildRevealed() bool {
	var _arg0 *C.GtkRevealer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_child_revealed(_arg0)
	runtime.KeepAlive(revealer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RevealChild returns whether the child is currently revealed.
//
// This function returns TRUE as soon as the transition is to the revealed state
// is started. To learn whether the child is fully revealed (ie the transition
// is completed), use gtk.Revealer.GetChildRevealed().
//
// The function returns the following values:
//
//   - ok: TRUE if the child is revealed.
//
func (revealer *Revealer) RevealChild() bool {
	var _arg0 *C.GtkRevealer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_reveal_child(_arg0)
	runtime.KeepAlive(revealer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions will take.
//
// The function returns the following values:
//
//   - guint: transition duration.
//
func (revealer *Revealer) TransitionDuration() uint {
	var _arg0 *C.GtkRevealer // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_transition_duration(_arg0)
	runtime.KeepAlive(revealer)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionType gets the type of animation that will be used for transitions
// in revealer.
//
// The function returns the following values:
//
//   - revealerTransitionType: current transition type of revealer.
//
func (revealer *Revealer) TransitionType() RevealerTransitionType {
	var _arg0 *C.GtkRevealer              // out
	var _cret C.GtkRevealerTransitionType // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))

	_cret = C.gtk_revealer_get_transition_type(_arg0)
	runtime.KeepAlive(revealer)

	var _revealerTransitionType RevealerTransitionType // out

	_revealerTransitionType = RevealerTransitionType(_cret)

	return _revealerTransitionType
}

// SetChild sets the child widget of revealer.
//
// The function takes the following parameters:
//
//   - child (optional) widget.
//
func (revealer *Revealer) SetChild(child Widgetter) {
	var _arg0 *C.GtkRevealer // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.gtk_revealer_set_child(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(child)
}

// SetRevealChild tells the GtkRevealer to reveal or conceal its child.
//
// The transition will be animated with the current transition type of revealer.
//
// The function takes the following parameters:
//
//   - revealChild: TRUE to reveal the child.
//
func (revealer *Revealer) SetRevealChild(revealChild bool) {
	var _arg0 *C.GtkRevealer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))
	if revealChild {
		_arg1 = C.TRUE
	}

	C.gtk_revealer_set_reveal_child(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(revealChild)
}

// SetTransitionDuration sets the duration that transitions will take.
//
// The function takes the following parameters:
//
//   - duration: new duration, in milliseconds.
//
func (revealer *Revealer) SetTransitionDuration(duration uint) {
	var _arg0 *C.GtkRevealer // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))
	_arg1 = C.guint(duration)

	C.gtk_revealer_set_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(duration)
}

// SetTransitionType sets the type of animation that will be used for
// transitions in revealer.
//
// Available types include various kinds of fades and slides.
//
// The function takes the following parameters:
//
//   - transition: new transition type.
//
func (revealer *Revealer) SetTransitionType(transition RevealerTransitionType) {
	var _arg0 *C.GtkRevealer              // out
	var _arg1 C.GtkRevealerTransitionType // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(coreglib.InternObject(revealer).Native()))
	_arg1 = C.GtkRevealerTransitionType(transition)

	C.gtk_revealer_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(transition)
}
