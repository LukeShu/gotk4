// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

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
	GTypeApplicationInhibitFlags = coreglib.Type(C.gtk_application_inhibit_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeApplicationInhibitFlags, F: marshalApplicationInhibitFlags},
	})
}

// ApplicationInhibitFlags types of user actions that may be blocked by
// gtk_application_inhibit().
type ApplicationInhibitFlags C.guint

const (
	// ApplicationInhibitLogout: inhibit ending the user session by logging out
	// or by shutting down the computer.
	ApplicationInhibitLogout ApplicationInhibitFlags = 0b1
	// ApplicationInhibitSwitch: inhibit user switching.
	ApplicationInhibitSwitch ApplicationInhibitFlags = 0b10
	// ApplicationInhibitSuspend: inhibit suspending the session or computer.
	ApplicationInhibitSuspend ApplicationInhibitFlags = 0b100
	// ApplicationInhibitIdle: inhibit the session being marked as idle (and
	// possibly locked).
	ApplicationInhibitIdle ApplicationInhibitFlags = 0b1000
)

func marshalApplicationInhibitFlags(p uintptr) (interface{}, error) {
	return ApplicationInhibitFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ApplicationInhibitFlags.
func (a ApplicationInhibitFlags) String() string {
	if a == 0 {
		return "ApplicationInhibitFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(98)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case ApplicationInhibitLogout:
			builder.WriteString("Logout|")
		case ApplicationInhibitSwitch:
			builder.WriteString("Switch|")
		case ApplicationInhibitSuspend:
			builder.WriteString("Suspend|")
		case ApplicationInhibitIdle:
			builder.WriteString("Idle|")
		default:
			builder.WriteString(fmt.Sprintf("ApplicationInhibitFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a ApplicationInhibitFlags) Has(other ApplicationInhibitFlags) bool {
	return (a & other) == other
}
