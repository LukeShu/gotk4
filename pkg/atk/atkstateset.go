// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_state_set_get_type()), F: marshalStateSet},
	})
}

// StateSet is a read-only representation of the full set of States that apply
// to an object at a given time. This set is not meant to be modified, but
// rather created when #atk_object_ref_state_set() is called.
type StateSet interface {
	gextras.Objector

	// AndSets constructs the intersection of the two sets, returning nil if the
	// intersection is empty.
	AndSets(compareSet StateSet) *StateSetClass
	// ClearStates removes all states from the state set.
	ClearStates()
	// IsEmpty checks whether the state set is empty, i.e. has no states set.
	IsEmpty() bool
	// OrSets constructs the union of the two sets.
	OrSets(compareSet StateSet) *StateSetClass
	// XorSets constructs the exclusive-or of the two sets, returning nil is
	// empty. The set returned by this operation contains the states in exactly
	// one of the two sets.
	XorSets(compareSet StateSet) *StateSetClass
}

// StateSetClass implements the StateSet interface.
type StateSetClass struct {
	*externglib.Object
}

var _ StateSet = (*StateSetClass)(nil)

func wrapStateSet(obj *externglib.Object) StateSet {
	return &StateSetClass{
		Object: obj,
	}
}

func marshalStateSet(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStateSet(obj), nil
}

// NewStateSet creates a new empty state set.
func NewStateSet() *StateSetClass {
	var _cret *C.AtkStateSet // in

	_cret = C.atk_state_set_new()

	var _stateSet *StateSetClass // out

	_stateSet = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*StateSetClass)

	return _stateSet
}

// AndSets constructs the intersection of the two sets, returning nil if the
// intersection is empty.
func (s *StateSetClass) AndSets(compareSet StateSet) *StateSetClass {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer((&compareSet).Native()))

	_cret = C.atk_state_set_and_sets(_arg0, _arg1)

	var _stateSet *StateSetClass // out

	_stateSet = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*StateSetClass)

	return _stateSet
}

// ClearStates removes all states from the state set.
func (s *StateSetClass) ClearStates() {
	var _arg0 *C.AtkStateSet // out

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer((&s).Native()))

	C.atk_state_set_clear_states(_arg0)
}

// IsEmpty checks whether the state set is empty, i.e. has no states set.
func (s *StateSetClass) IsEmpty() bool {
	var _arg0 *C.AtkStateSet // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer((&s).Native()))

	_cret = C.atk_state_set_is_empty(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OrSets constructs the union of the two sets.
func (s *StateSetClass) OrSets(compareSet StateSet) *StateSetClass {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer((&compareSet).Native()))

	_cret = C.atk_state_set_or_sets(_arg0, _arg1)

	var _stateSet *StateSetClass // out

	_stateSet = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*StateSetClass)

	return _stateSet
}

// XorSets constructs the exclusive-or of the two sets, returning nil is empty.
// The set returned by this operation contains the states in exactly one of the
// two sets.
func (s *StateSetClass) XorSets(compareSet StateSet) *StateSetClass {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer((&compareSet).Native()))

	_cret = C.atk_state_set_xor_sets(_arg0, _arg1)

	var _stateSet *StateSetClass // out

	_stateSet = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*StateSetClass)

	return _stateSet
}
