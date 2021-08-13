// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_revealer_transition_type_get_type()), F: marshalRevealerTransitionType},
		{T: externglib.Type(C.gtk_revealer_get_type()), F: marshalRevealerer},
	})
}

// RevealerTransitionType: these enumeration values describe the possible
// transitions when the child of a Revealer widget is shown or hidden.
type RevealerTransitionType int

const (
	// RevealerTransitionTypeNone: no transition
	RevealerTransitionTypeNone RevealerTransitionType = iota
	// RevealerTransitionTypeCrossfade: fade in
	RevealerTransitionTypeCrossfade
	// RevealerTransitionTypeSlideRight: slide in from the left
	RevealerTransitionTypeSlideRight
	// RevealerTransitionTypeSlideLeft: slide in from the right
	RevealerTransitionTypeSlideLeft
	// RevealerTransitionTypeSlideUp: slide in from the bottom
	RevealerTransitionTypeSlideUp
	// RevealerTransitionTypeSlideDown: slide in from the top
	RevealerTransitionTypeSlideDown
)

func marshalRevealerTransitionType(p uintptr) (interface{}, error) {
	return RevealerTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
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
	default:
		return fmt.Sprintf("RevealerTransitionType(%d)", r)
	}
}

// Revealer widget is a container which animates the transition of its child
// from invisible to visible.
//
// The style of transition can be controlled with
// gtk_revealer_set_transition_type().
//
// These animations respect the Settings:gtk-enable-animations setting.
//
//
// CSS nodes
//
// GtkRevealer has a single CSS node with name revealer.
//
// The GtkRevealer widget was added in GTK+ 3.10.
type Revealer struct {
	Bin
}

func wrapRevealer(obj *externglib.Object) *Revealer {
	return &Revealer{
		Bin: Bin{
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
					Object: obj,
				},
			},
		},
	}
}

func marshalRevealerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRevealer(obj), nil
}

// NewRevealer creates a new Revealer.
func NewRevealer() *Revealer {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_revealer_new()

	var _revealer *Revealer // out

	_revealer = wrapRevealer(externglib.Take(unsafe.Pointer(_cret)))

	return _revealer
}

// ChildRevealed returns whether the child is fully revealed, in other words
// whether the transition to the revealed state is completed.
func (revealer *Revealer) ChildRevealed() bool {
	var _arg0 *C.GtkRevealer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(revealer.Native()))

	_cret = C.gtk_revealer_get_child_revealed(_arg0)

	runtime.KeepAlive(revealer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RevealChild returns whether the child is currently revealed. See
// gtk_revealer_set_reveal_child().
//
// This function returns TRUE as soon as the transition is to the revealed state
// is started. To learn whether the child is fully revealed (ie the transition
// is completed), use gtk_revealer_get_child_revealed().
func (revealer *Revealer) RevealChild() bool {
	var _arg0 *C.GtkRevealer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(revealer.Native()))

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
func (revealer *Revealer) TransitionDuration() uint {
	var _arg0 *C.GtkRevealer // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(revealer.Native()))

	_cret = C.gtk_revealer_get_transition_duration(_arg0)

	runtime.KeepAlive(revealer)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionType gets the type of animation that will be used for transitions
// in revealer.
func (revealer *Revealer) TransitionType() RevealerTransitionType {
	var _arg0 *C.GtkRevealer              // out
	var _cret C.GtkRevealerTransitionType // in

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(revealer.Native()))

	_cret = C.gtk_revealer_get_transition_type(_arg0)

	runtime.KeepAlive(revealer)

	var _revealerTransitionType RevealerTransitionType // out

	_revealerTransitionType = RevealerTransitionType(_cret)

	return _revealerTransitionType
}

// SetRevealChild tells the Revealer to reveal or conceal its child.
//
// The transition will be animated with the current transition type of revealer.
func (revealer *Revealer) SetRevealChild(revealChild bool) {
	var _arg0 *C.GtkRevealer // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(revealer.Native()))
	if revealChild {
		_arg1 = C.TRUE
	}

	C.gtk_revealer_set_reveal_child(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(revealChild)
}

// SetTransitionDuration sets the duration that transitions will take.
func (revealer *Revealer) SetTransitionDuration(duration uint) {
	var _arg0 *C.GtkRevealer // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(revealer.Native()))
	_arg1 = C.guint(duration)

	C.gtk_revealer_set_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(duration)
}

// SetTransitionType sets the type of animation that will be used for
// transitions in revealer. Available types include various kinds of fades and
// slides.
func (revealer *Revealer) SetTransitionType(transition RevealerTransitionType) {
	var _arg0 *C.GtkRevealer              // out
	var _arg1 C.GtkRevealerTransitionType // out

	_arg0 = (*C.GtkRevealer)(unsafe.Pointer(revealer.Native()))
	_arg1 = C.GtkRevealerTransitionType(transition)

	C.gtk_revealer_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(revealer)
	runtime.KeepAlive(transition)
}
