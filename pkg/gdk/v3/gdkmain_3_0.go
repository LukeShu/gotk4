// Code generated by girgen. DO NOT EDIT.

package gdk

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// DisableMultidevice disables multidevice support in GDK. This call must
// happen prior to gdk_display_open(), gtk_init(), gtk_init_with_args() or
// gtk_init_check() in order to take effect.
//
// Most common GTK+ applications won’t ever need to call this. Only applications
// that do mixed GDK/Xlib calls could want to disable multidevice support if
// such Xlib code deals with input devices in any way and doesn’t observe the
// presence of XInput 2.
func DisableMultidevice() {
	C.gdk_disable_multidevice()
}

// ErrorTrapPopIgnored removes an error trap pushed with gdk_error_trap_push(),
// but without bothering to wait and see whether an error occurred. If an error
// arrives later asynchronously that was triggered while the trap was pushed,
// that error will be ignored.
func ErrorTrapPopIgnored() {
	C.gdk_error_trap_pop_ignored()
}
