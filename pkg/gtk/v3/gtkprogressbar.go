// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeProgressBar returns the GType for the type ProgressBar.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeProgressBar() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ProgressBar").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalProgressBar)
	return gtype
}

// ProgressBarOverrider contains methods that are overridable.
type ProgressBarOverrider interface {
}

// ProgressBar is typically used to display the progress of a long running
// operation. It provides a visual clue that processing is underway. The
// GtkProgressBar can be used in two different modes: percentage mode and
// activity mode.
//
// When an application can determine how much work needs to take place (e.g.
// read a fixed number of bytes from a file) and can monitor its progress, it
// can use the GtkProgressBar in percentage mode and the user sees a growing bar
// indicating the percentage of the work that has been completed. In this mode,
// the application is required to call gtk_progress_bar_set_fraction()
// periodically to update the progress bar.
//
// When an application has no accurate way of knowing the amount of work to do,
// it can use the ProgressBar in activity mode, which shows activity by a block
// moving back and forth within the progress area. In this mode, the application
// is required to call gtk_progress_bar_pulse() periodically to update the
// progress bar.
//
// There is quite a bit of flexibility provided to control the appearance of the
// ProgressBar. Functions are provided to control the orientation of the bar,
// optional text can be displayed along with the bar, and the step size used in
// activity mode can be set.
//
// CSS nodes
//
//    progressbar[.osd]
//    ├── [text]
//    ╰── trough[.empty][.full]
//        ╰── progress[.pulse]
//
// GtkProgressBar has a main CSS node with name progressbar and subnodes with
// names text and trough, of which the latter has a subnode named progress. The
// text subnode is only present if text is shown. The progress subnode has the
// style class .pulse when in activity mode. It gets the style classes .left,
// .right, .top or .bottom added when the progress 'touches' the corresponding
// end of the GtkProgressBar. The .osd class on the progressbar node is for use
// in overlays like the one Epiphany has for page loading progress.
type ProgressBar struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*ProgressBar)(nil)
	_ coreglib.Objector = (*ProgressBar)(nil)
)

func classInitProgressBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapProgressBar(obj *coreglib.Object) *ProgressBar {
	return &ProgressBar{
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
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	return wrapProgressBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewProgressBar creates a new ProgressBar.
//
// The function returns the following values:
//
//    - progressBar: ProgressBar.
//
func NewProgressBar() *ProgressBar {
	_gret := girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("new_ProgressBar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _progressBar *ProgressBar // out

	_progressBar = wrapProgressBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _progressBar
}

// Fraction returns the current fraction of the task that’s been completed.
//
// The function returns the following values:
//
//    - gdouble: fraction from 0.0 to 1.0.
//
func (pbar *ProgressBar) Fraction() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))

	_gret := girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("get_fraction", _args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pbar)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// Inverted gets the value set by gtk_progress_bar_set_inverted().
//
// The function returns the following values:
//
//    - ok: TRUE if the progress bar is inverted.
//
func (pbar *ProgressBar) Inverted() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))

	_gret := girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("get_inverted", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pbar)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// PulseStep retrieves the pulse step set with
// gtk_progress_bar_set_pulse_step().
//
// The function returns the following values:
//
//    - gdouble: fraction from 0.0 to 1.0.
//
func (pbar *ProgressBar) PulseStep() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))

	_gret := girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("get_pulse_step", _args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pbar)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// ShowText gets the value of the ProgressBar:show-text property. See
// gtk_progress_bar_set_show_text().
//
// The function returns the following values:
//
//    - ok: TRUE if text is shown in the progress bar.
//
func (pbar *ProgressBar) ShowText() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))

	_gret := girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("get_show_text", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pbar)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Text retrieves the text that is displayed with the progress bar, if any,
// otherwise NULL. The return value is a reference to the text, not a copy of
// it, so will become invalid if you change the text in the progress bar.
//
// The function returns the following values:
//
//    - utf8 (optional): text, or NULL; this string is owned by the widget and
//      should not be modified or freed.
//
func (pbar *ProgressBar) Text() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))

	_gret := girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("get_text", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pbar)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Pulse indicates that some progress has been made, but you don’t know how
// much. Causes the progress bar to enter “activity mode,” where a block bounces
// back and forth. Each call to gtk_progress_bar_pulse() causes the block to
// move by a little bit (the amount of movement per pulse is determined by
// gtk_progress_bar_set_pulse_step()).
func (pbar *ProgressBar) Pulse() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))

	girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("pulse", _args[:], nil)

	runtime.KeepAlive(pbar)
}

// SetFraction causes the progress bar to “fill in” the given fraction of the
// bar. The fraction should be between 0.0 and 1.0, inclusive.
//
// The function takes the following parameters:
//
//    - fraction of the task that’s been completed.
//
func (pbar *ProgressBar) SetFraction(fraction float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(fraction)

	girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("set_fraction", _args[:], nil)

	runtime.KeepAlive(pbar)
	runtime.KeepAlive(fraction)
}

// SetInverted progress bars normally grow from top to bottom or left to right.
// Inverted progress bars grow in the opposite direction.
//
// The function takes the following parameters:
//
//    - inverted: TRUE to invert the progress bar.
//
func (pbar *ProgressBar) SetInverted(inverted bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))
	if inverted {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("set_inverted", _args[:], nil)

	runtime.KeepAlive(pbar)
	runtime.KeepAlive(inverted)
}

// SetPulseStep sets the fraction of total progress bar length to move the
// bouncing block for each call to gtk_progress_bar_pulse().
//
// The function takes the following parameters:
//
//    - fraction between 0.0 and 1.0.
//
func (pbar *ProgressBar) SetPulseStep(fraction float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(fraction)

	girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("set_pulse_step", _args[:], nil)

	runtime.KeepAlive(pbar)
	runtime.KeepAlive(fraction)
}

// SetShowText sets whether the progress bar will show text next to the bar. The
// shown text is either the value of the ProgressBar:text property or, if that
// is NULL, the ProgressBar:fraction value, as a percentage.
//
// To make a progress bar that is styled and sized suitably for containing text
// (even if the actual text is blank), set ProgressBar:show-text to TRUE and
// ProgressBar:text to the empty string (not NULL).
//
// The function takes the following parameters:
//
//    - showText: whether to show text.
//
func (pbar *ProgressBar) SetShowText(showText bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))
	if showText {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("set_show_text", _args[:], nil)

	runtime.KeepAlive(pbar)
	runtime.KeepAlive(showText)
}

// SetText causes the given text to appear next to the progress bar.
//
// If text is NULL and ProgressBar:show-text is TRUE, the current value of
// ProgressBar:fraction will be displayed as a percentage.
//
// If text is non-NULL and ProgressBar:show-text is TRUE, the text will be
// displayed. In this case, it will not display the progress percentage. If text
// is the empty string, the progress bar will still be styled and sized suitably
// for containing text, as long as ProgressBar:show-text is TRUE.
//
// The function takes the following parameters:
//
//    - text (optional): UTF-8 string, or NULL.
//
func (pbar *ProgressBar) SetText(text string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pbar).Native()))
	if text != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "ProgressBar").InvokeMethod("set_text", _args[:], nil)

	runtime.KeepAlive(pbar)
	runtime.KeepAlive(text)
}
