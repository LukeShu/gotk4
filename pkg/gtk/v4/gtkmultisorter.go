// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_multi_sorter_get_type()), F: marshalMultiSorter},
	})
}

// MultiSorter: `GtkMultiSorter` combines multiple sorters by trying them in
// turn.
//
// If the first sorter compares two items as equal, the second is tried next,
// and so on.
type MultiSorter interface {
	gextras.Objector

	// Append: add @sorter to @self to use for sorting at the end.
	//
	// @self will consult all existing sorters before it will sort with the
	// given @sorter.
	Append(sorter Sorter)
	// Remove removes the sorter at the given @position from the list of sorter
	// used by @self.
	//
	// If @position is larger than the number of sorters, nothing happens.
	Remove(position uint)
}

// MultiSorterClass implements the MultiSorter interface.
type MultiSorterClass struct {
	*externglib.Object
	SorterClass
	BuildableInterface
}

var _ MultiSorter = (*MultiSorterClass)(nil)

func wrapMultiSorter(obj *externglib.Object) MultiSorter {
	return &MultiSorterClass{
		Object: obj,
		SorterClass: SorterClass{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalMultiSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMultiSorter(obj), nil
}

// NewMultiSorter creates a new multi sorter.
//
// This sorter compares items by trying each of the sorters in turn, until one
// returns non-zero. In particular, if no sorter has been added to it, it will
// always compare items as equal.
func NewMultiSorter() MultiSorter {
	var _cret *C.GtkMultiSorter // in

	_cret = C.gtk_multi_sorter_new()

	var _multiSorter MultiSorter // out

	_multiSorter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MultiSorter)

	return _multiSorter
}

// Append: add @sorter to @self to use for sorting at the end.
//
// @self will consult all existing sorters before it will sort with the given
// @sorter.
func (s *MultiSorterClass) Append(sorter Sorter) {
	var _arg0 *C.GtkMultiSorter // out
	var _arg1 *C.GtkSorter      // out

	_arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_multi_sorter_append(_arg0, _arg1)
}

// Remove removes the sorter at the given @position from the list of sorter used
// by @self.
//
// If @position is larger than the number of sorters, nothing happens.
func (s *MultiSorterClass) Remove(position uint) {
	var _arg0 *C.GtkMultiSorter // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)

	C.gtk_multi_sorter_remove(_arg0, _arg1)
}
