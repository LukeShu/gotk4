// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_LevelBar_ConnectOffsetChanged(gpointer, gchar*, guintptr);
import "C"

// glib.Type values for gtklevelbar.go.
var GTypeLevelBar = coreglib.Type(C.gtk_level_bar_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeLevelBar, F: marshalLevelBar},
	})
}

// LEVEL_BAR_OFFSET_FULL: name used for the stock full offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_FULL = "full"

// LEVEL_BAR_OFFSET_HIGH: name used for the stock high offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_HIGH = "high"

// LEVEL_BAR_OFFSET_LOW: name used for the stock low offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_LOW = "low"

// LevelBar: GtkLevelBar is a widget that can be used as a level indicator.
//
// Typical use cases are displaying the strength of a password, or showing the
// charge level of a battery.
//
// !An example GtkLevelBar (levelbar.png)
//
// Use gtk.LevelBar.SetValue() to set the current value, and
// gtk.LevelBar.AddOffsetValue() to set the value offsets at which the bar will
// be considered in a different state. GTK will add a few offsets by default on
// the level bar: GTK_LEVEL_BAR_OFFSET_LOW, GTK_LEVEL_BAR_OFFSET_HIGH and
// GTK_LEVEL_BAR_OFFSET_FULL, with values 0.25, 0.75 and 1.0 respectively.
//
// Note that it is your responsibility to update preexisting offsets when
// changing the minimum or maximum value. GTK will simply clamp them to the new
// range.
//
// Adding a custom offset on the bar
//
//    static GtkWidget *
//    create_level_bar (void)
//    {
//      GtkWidget *widget;
//      GtkLevelBar *bar;
//
//      widget = gtk_level_bar_new ();
//      bar = GTK_LEVEL_BAR (widget);
//
//      // This changes the value of the default low offset
//
//      gtk_level_bar_add_offset_value (bar,
//                                      GTK_LEVEL_BAR_OFFSET_LOW,
//                                      0.10);
//
//      // This adds a new offset to the bar; the application will
//      // be able to change its color CSS like this:
//      //
//      // levelbar block.my-offset {
//      //   background-color: magenta;
//      //   border-style: solid;
//      //   border-color: black;
//      //   border-style: 1px;
//      // }
//
//      gtk_level_bar_add_offset_value (bar, "my-offset", 0.60);
//
//      return widget;
//    }
//
//
// The default interval of values is between zero and one, but it’s possible to
// modify the interval using gtk.LevelBar.SetMinValue() and
// gtk.LevelBar.SetMaxValue(). The value will be always drawn in proportion to
// the admissible interval, i.e. a value of 15 with a specified interval between
// 10 and 20 is equivalent to a value of 0.5 with an interval between 0 and 1.
// When K_LEVEL_BAR_MODE_DISCRETE is used, the bar level is rendered as a finite
// number of separated blocks instead of a single one. The number of blocks that
// will be rendered is equal to the number of units specified by the admissible
// interval.
//
// For instance, to build a bar rendered with five blocks, it’s sufficient to
// set the minimum value to 0 and the maximum value to 5 after changing the
// indicator mode to discrete.
//
//
// GtkLevelBar as GtkBuildable
//
// The GtkLevelBar implementation of the GtkBuildable interface supports a
// custom <offsets> element, which can contain any number of <offset> elements,
// each of which must have name and value attributes.
//
// CSS nodes
//
//    levelbar[.discrete]
//    ╰── trough
//        ├── block.filled.level-name
//        ┊
//        ├── block.empty
//        ┊
//
//
// GtkLevelBar has a main CSS node with name levelbar and one of the style
// classes .discrete or .continuous and a subnode with name trough. Below the
// trough node are a number of nodes with name block and style class .filled or
// .empty. In continuous mode, there is exactly one node of each, in discrete
// mode, the number of filled and unfilled nodes corresponds to blocks that are
// drawn. The block.filled nodes also get a style class .level-name
// corresponding to the level for the current value.
//
// In horizontal orientation, the nodes are always arranged from left to right,
// regardless of text direction.
//
//
// Accessibility
//
// GtkLevelBar uses the K_ACCESSIBLE_ROLE_METER role.
type LevelBar struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*LevelBar)(nil)
	_ coreglib.Objector = (*LevelBar)(nil)
)

func wrapLevelBar(obj *coreglib.Object) *LevelBar {
	return &LevelBar{
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
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalLevelBar(p uintptr) (interface{}, error) {
	return wrapLevelBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_LevelBar_ConnectOffsetChanged
func _gotk4_gtk4_LevelBar_ConnectOffsetChanged(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(name string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(name string))
	}

	var _name string // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_name)
}

// ConnectOffsetChanged is emitted when an offset specified on the bar changes
// value.
//
// This typically is the result of a gtk.LevelBar.AddOffsetValue() call.
//
// The signal supports detailed connections; you can connect to the detailed
// signal "changed::x" in order to only receive callbacks when the value of
// offset "x" changes.
func (self *LevelBar) ConnectOffsetChanged(f func(name string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "offset-changed", false, unsafe.Pointer(C._gotk4_gtk4_LevelBar_ConnectOffsetChanged), f)
}

// NewLevelBar creates a new GtkLevelBar.
//
// The function returns the following values:
//
//    - levelBar: GtkLevelBar.
//
func NewLevelBar() *LevelBar {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("new_LevelBar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _levelBar *LevelBar // out

	_levelBar = wrapLevelBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _levelBar
}

// NewLevelBarForInterval creates a new GtkLevelBar for the specified interval.
//
// The function takes the following parameters:
//
//    - minValue: positive value.
//    - maxValue: positive value.
//
// The function returns the following values:
//
//    - levelBar: GtkLevelBar.
//
func NewLevelBarForInterval(minValue, maxValue float64) *LevelBar {
	var args [2]girepository.Argument
	var _arg0 C.double // out
	var _arg1 C.double // out
	var _cret *C.void  // in

	_arg0 = C.double(minValue)
	_arg1 = C.double(maxValue)
	*(*float64)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("new_LevelBar_for_interval", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(minValue)
	runtime.KeepAlive(maxValue)

	var _levelBar *LevelBar // out

	_levelBar = wrapLevelBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _levelBar
}

// AddOffsetValue adds a new offset marker on self at the position specified by
// value.
//
// When the bar value is in the interval topped by value (or between value and
// gtk.LevelBar:max-value in case the offset is the last one on the bar) a style
// class named level-name will be applied when rendering the level bar fill.
//
// If another offset marker named name exists, its value will be replaced by
// value.
//
// The function takes the following parameters:
//
//    - name of the new offset.
//    - value for the new offset.
//
func (self *LevelBar) AddOffsetValue(name string, value float64) {
	var args [3]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 *C.void  // out
	var _arg2 C.double // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)
	*(**LevelBar)(unsafe.Pointer(&args[1])) = _arg1
	*(*string)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("add_offset_value", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// Inverted returns whether the levelbar is inverted.
//
// The function returns the following values:
//
//    - ok: TRUE if the level bar is inverted.
//
func (self *LevelBar) Inverted() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**LevelBar)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_inverted", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxValue returns the max-value of the GtkLevelBar.
//
// The function returns the following values:
//
//    - gdouble: positive value.
//
func (self *LevelBar) MaxValue() float64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.double // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**LevelBar)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_max_value", args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinValue returns the min-value of the GtkLevelBar`.
//
// The function returns the following values:
//
//    - gdouble: positive value.
//
func (self *LevelBar) MinValue() float64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.double // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**LevelBar)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_min_value", args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Value returns the value of the GtkLevelBar.
//
// The function returns the following values:
//
//    - gdouble: value in the interval between GtkLevelBar:min-value and
//      GtkLevelBar:max-value.
//
func (self *LevelBar) Value() float64 {
	var args [1]girepository.Argument
	var _arg0 *C.void  // out
	var _cret C.double // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**LevelBar)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_value", args[:], nil)
	_cret = *(*C.double)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RemoveOffsetValue removes an offset marker from a GtkLevelBar.
//
// The marker must have been previously added with
// gtk.LevelBar.AddOffsetValue().
//
// The function takes the following parameters:
//
//    - name (optional) of an offset in the bar.
//
func (self *LevelBar) RemoveOffsetValue(name string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if name != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**LevelBar)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("remove_offset_value", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetInverted sets whether the GtkLevelBar is inverted.
//
// The function takes the following parameters:
//
//    - inverted: TRUE to invert the level bar.
//
func (self *LevelBar) SetInverted(inverted bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if inverted {
		_arg1 = C.TRUE
	}
	*(**LevelBar)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_inverted", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(inverted)
}

// SetMaxValue sets the max-value of the GtkLevelBar.
//
// You probably want to update preexisting level offsets after calling this
// function.
//
// The function takes the following parameters:
//
//    - value: positive value.
//
func (self *LevelBar) SetMaxValue(value float64) {
	var args [2]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 C.double // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(value)
	*(**LevelBar)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_max_value", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetMinValue sets the min-value of the GtkLevelBar.
//
// You probably want to update preexisting level offsets after calling this
// function.
//
// The function takes the following parameters:
//
//    - value: positive value.
//
func (self *LevelBar) SetMinValue(value float64) {
	var args [2]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 C.double // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(value)
	*(**LevelBar)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_min_value", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetValue sets the value of the GtkLevelBar.
//
// The function takes the following parameters:
//
//    - value in the interval between gtk.LevelBar:min-value and
//      gtk.LevelBar:max-value.
//
func (self *LevelBar) SetValue(value float64) {
	var args [2]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 C.double // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(value)
	*(**LevelBar)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_value", args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}
