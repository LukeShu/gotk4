// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_seat_capabilities_get_type()), F: marshalSeatCapabilities},
		{T: externglib.Type(C.gdk_seat_get_type()), F: marshalSeater},
	})
}

// SeatCapabilities flags describing the seat capabilities.
type SeatCapabilities int

const (
	// SeatCapabilityNone: no input capabilities
	SeatCapabilityNone SeatCapabilities = 0b0
	// SeatCapabilityPointer: seat has a pointer (e.g. mouse)
	SeatCapabilityPointer SeatCapabilities = 0b1
	// SeatCapabilityTouch: seat has touchscreen(s) attached
	SeatCapabilityTouch SeatCapabilities = 0b10
	// SeatCapabilityTabletStylus: seat has drawing tablet(s) attached
	SeatCapabilityTabletStylus SeatCapabilities = 0b100
	// SeatCapabilityKeyboard: seat has keyboard(s) attached
	SeatCapabilityKeyboard SeatCapabilities = 0b1000
	// SeatCapabilityTabletPad: seat has drawing tablet pad(s) attached
	SeatCapabilityTabletPad SeatCapabilities = 0b10000
	// SeatCapabilityAllPointing: union of all pointing capabilities
	SeatCapabilityAllPointing SeatCapabilities = 0b111
	// SeatCapabilityAll: union of all capabilities
	SeatCapabilityAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for SeatCapabilities.
func (s SeatCapabilities) String() string {
	if s == 0 {
		return "SeatCapabilities(0)"
	}

	var builder strings.Builder
	builder.Grow(178)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SeatCapabilityNone:
			builder.WriteString("None|")
		case SeatCapabilityPointer:
			builder.WriteString("Pointer|")
		case SeatCapabilityTouch:
			builder.WriteString("Touch|")
		case SeatCapabilityTabletStylus:
			builder.WriteString("TabletStylus|")
		case SeatCapabilityKeyboard:
			builder.WriteString("Keyboard|")
		case SeatCapabilityTabletPad:
			builder.WriteString("TabletPad|")
		case SeatCapabilityAllPointing:
			builder.WriteString("AllPointing|")
		case SeatCapabilityAll:
			builder.WriteString("All|")
		default:
			builder.WriteString(fmt.Sprintf("SeatCapabilities(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Seat: GdkSeat object represents a collection of input devices that belong to
// a user.
type Seat struct {
	*externglib.Object
}

// Seater describes Seat's abstract methods.
type Seater interface {
	externglib.Objector

	// Capabilities returns the capabilities this GdkSeat currently has.
	Capabilities() SeatCapabilities
	// Devices returns the devices that match the given capabilities.
	Devices(capabilities SeatCapabilities) []Devicer
	// Display returns the GdkDisplay this seat belongs to.
	Display() *Display
	// Keyboard returns the device that routes keyboard events.
	Keyboard() Devicer
	// Pointer returns the device that routes pointer events.
	Pointer() Devicer
	// Tools returns all GdkDeviceTools that are known to the application.
	Tools() []DeviceTool
}

var _ Seater = (*Seat)(nil)

func wrapSeat(obj *externglib.Object) *Seat {
	return &Seat{
		Object: obj,
	}
}

func marshalSeater(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSeat(obj), nil
}

// Capabilities returns the capabilities this GdkSeat currently has.
func (seat *Seat) Capabilities() SeatCapabilities {
	var _arg0 *C.GdkSeat            // out
	var _cret C.GdkSeatCapabilities // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_capabilities(_arg0)

	runtime.KeepAlive(seat)

	var _seatCapabilities SeatCapabilities // out

	_seatCapabilities = SeatCapabilities(_cret)

	return _seatCapabilities
}

// Devices returns the devices that match the given capabilities.
func (seat *Seat) Devices(capabilities SeatCapabilities) []Devicer {
	var _arg0 *C.GdkSeat            // out
	var _arg1 C.GdkSeatCapabilities // out
	var _cret *C.GList              // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))
	_arg1 = C.GdkSeatCapabilities(capabilities)

	_cret = C.gdk_seat_get_devices(_arg0, _arg1)

	runtime.KeepAlive(seat)
	runtime.KeepAlive(capabilities)

	var _list []Devicer // out

	_list = make([]Devicer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkDevice)(v)
		var dst Devicer // out
		dst = (externglib.CastObject(externglib.Take(unsafe.Pointer(src)))).(Devicer)
		_list = append(_list, dst)
	})

	return _list
}

// Display returns the GdkDisplay this seat belongs to.
func (seat *Seat) Display() *Display {
	var _arg0 *C.GdkSeat    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_display(_arg0)

	runtime.KeepAlive(seat)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Keyboard returns the device that routes keyboard events.
func (seat *Seat) Keyboard() Devicer {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)

	runtime.KeepAlive(seat)

	var _device Devicer // out

	if _cret != nil {
		_device = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Devicer)
	}

	return _device
}

// Pointer returns the device that routes pointer events.
func (seat *Seat) Pointer() Devicer {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)

	runtime.KeepAlive(seat)

	var _device Devicer // out

	if _cret != nil {
		_device = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Devicer)
	}

	return _device
}

// Tools returns all GdkDeviceTools that are known to the application.
func (seat *Seat) Tools() []DeviceTool {
	var _arg0 *C.GdkSeat // out
	var _cret *C.GList   // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_tools(_arg0)

	runtime.KeepAlive(seat)

	var _list []DeviceTool // out

	_list = make([]DeviceTool, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkDeviceTool)(v)
		var dst DeviceTool // out
		dst = *wrapDeviceTool(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}
