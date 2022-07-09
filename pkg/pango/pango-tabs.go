// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeTabAlign returns the GType for the type TabAlign.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTabAlign() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "TabAlign").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTabAlign)
	return gtype
}

// GTypeTabArray returns the GType for the type TabArray.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTabArray() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "TabArray").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTabArray)
	return gtype
}

// TabAlign: PangoTabAlign specifies where a tab stop appears relative to the
// text.
type TabAlign C.gint

const (
	// TabLeft: tab stop appears to the left of the text.
	TabLeft TabAlign = iota
)

func marshalTabAlign(p uintptr) (interface{}, error) {
	return TabAlign(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for TabAlign.
func (t TabAlign) String() string {
	switch t {
	case TabLeft:
		return "Left"
	default:
		return fmt.Sprintf("TabAlign(%d)", t)
	}
}

// TabArray: PangoTabArray contains an array of tab stops.
//
// PangoTabArray can be used to set tab stops in a PangoLayout. Each tab stop
// has an alignment and a position.
//
// An instance of this type is always passed by reference.
type TabArray struct {
	*tabArray
}

// tabArray is the struct that's finalized.
type tabArray struct {
	native unsafe.Pointer
}

func marshalTabArray(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &TabArray{&tabArray{(unsafe.Pointer)(b)}}, nil
}

// NewTabArray constructs a struct TabArray.
func NewTabArray(initialSize int32, positionsInPixels bool) *TabArray {
	var _args [2]girepository.Argument

	*(*C.gint)(unsafe.Pointer(&_args[0])) = C.gint(initialSize)
	if positionsInPixels {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(initialSize)
	runtime.KeepAlive(positionsInPixels)

	var _tabArray *TabArray // out

	_tabArray = (*TabArray)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_tabArray)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				args := [1]girepository.Argument{(*C.void)(intern.C)}
				girepository.MustFind("Pango", "TabArray").InvokeMethod("free", args[:], nil)
			}
		},
	)

	return _tabArray
}

// Copy copies a PangoTabArray.
//
// The function returns the following values:
//
//    - tabArray: newly allocated PangoTabArray, which should be freed with
//      pango.TabArray.Free().
//
func (src *TabArray) Copy() *TabArray {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(src)

	var _tabArray *TabArray // out

	_tabArray = (*TabArray)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_tabArray)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				args := [1]girepository.Argument{(*C.void)(intern.C)}
				girepository.MustFind("Pango", "TabArray").InvokeMethod("free", args[:], nil)
			}
		},
	)

	return _tabArray
}

// PositionsInPixels returns TRUE if the tab positions are in pixels, FALSE if
// they are in Pango units.
//
// The function returns the following values:
//
//    - ok: whether positions are in pixels.
//
func (tabArray *TabArray) PositionsInPixels() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(tabArray)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(tabArray)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Size gets the number of tab stops in tab_array.
//
// The function returns the following values:
//
//    - gint: number of tab stops in the array.
//
func (tabArray *TabArray) Size() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(tabArray)))

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(tabArray)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Tab gets the alignment and position of a tab stop.
//
// The function takes the following parameters:
//
//    - tabIndex: tab stop index.
//
// The function returns the following values:
//
//    - alignment (optional): location to store alignment, or NULL.
//    - location (optional) to store tab position, or NULL.
//
func (tabArray *TabArray) Tab(tabIndex int32) (TabAlign, int32) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(tabArray)))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(tabIndex)

	runtime.KeepAlive(tabArray)
	runtime.KeepAlive(tabIndex)

	var _alignment TabAlign // out
	var _location int32     // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_alignment = *(*TabAlign)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_location = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _alignment, _location
}

// Resize resizes a tab array.
//
// You must subsequently initialize any tabs that were added as a result of
// growing the array.
//
// The function takes the following parameters:
//
//    - newSize: new size of the array.
//
func (tabArray *TabArray) Resize(newSize int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(tabArray)))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(newSize)

	runtime.KeepAlive(tabArray)
	runtime.KeepAlive(newSize)
}
