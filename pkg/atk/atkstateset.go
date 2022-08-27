// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeStateSet = coreglib.Type(C.atk_state_set_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStateSet, F: marshalStateSet},
	})
}

// StateSetOverrides contains methods that are overridable.
type StateSetOverrides struct {
}

func defaultStateSetOverrides(v *StateSet) StateSetOverrides {
	return StateSetOverrides{}
}

// StateSet is a read-only representation of the full set of States that apply
// to an object at a given time. This set is not meant to be modified, but
// rather created when #atk_object_ref_state_set() is called.
type StateSet struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StateSet)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StateSet, *StateSetClass, StateSetOverrides](
		GTypeStateSet,
		initStateSetClass,
		wrapStateSet,
		defaultStateSetOverrides,
	)
}

func initStateSetClass(gclass unsafe.Pointer, overrides StateSetOverrides, classInitFunc func(*StateSetClass)) {
	if classInitFunc != nil {
		class := (*StateSetClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStateSet(obj *coreglib.Object) *StateSet {
	return &StateSet{
		Object: obj,
	}
}

func marshalStateSet(p uintptr) (interface{}, error) {
	return wrapStateSet(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStateSet creates a new empty state set.
//
// The function returns the following values:
//
//    - stateSet: new StateSet.
//
func NewStateSet() *StateSet {
	var _cret *C.AtkStateSet // in

	_cret = C.atk_state_set_new()

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

// AddState adds the state of the specified type to the state set if it is not
// already present.
//
// Note that because an StateSet is a read-only object, this method should be
// used to add a state to a newly-created set which will then be returned by
// #atk_object_ref_state_set. It should not be used to modify the existing state
// of an object. See also #atk_object_notify_state_change.
//
// The function takes the following parameters:
//
//    - typ: StateType.
//
// The function returns the following values:
//
//    - ok: TRUE if the state for type is not already in set.
//
func (set *StateSet) AddState(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_add_state(_arg0, _arg1)
	runtime.KeepAlive(set)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AddStates adds the states of the specified types to the state set.
//
// Note that because an StateSet is a read-only object, this method should be
// used to add states to a newly-created set which will then be returned by
// #atk_object_ref_state_set. It should not be used to modify the existing state
// of an object. See also #atk_object_notify_state_change.
//
// The function takes the following parameters:
//
//    - types: array of StateType.
//
func (set *StateSet) AddStates(types []StateType) {
	var _arg0 *C.AtkStateSet  // out
	var _arg1 *C.AtkStateType // out
	var _arg2 C.gint

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg2 = (C.gint)(len(types))
	if len(types) > 0 {
		_arg1 = (*C.AtkStateType)(unsafe.Pointer(&types[0]))
	}

	C.atk_state_set_add_states(_arg0, _arg1, _arg2)
	runtime.KeepAlive(set)
	runtime.KeepAlive(types)
}

// AndSets constructs the intersection of the two sets, returning NULL if the
// intersection is empty.
//
// The function takes the following parameters:
//
//    - compareSet: another StateSet.
//
// The function returns the following values:
//
//    - stateSet: new StateSet which is the intersection of the two sets.
//
func (set *StateSet) AndSets(compareSet *StateSet) *StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(compareSet).Native()))

	_cret = C.atk_state_set_and_sets(_arg0, _arg1)
	runtime.KeepAlive(set)
	runtime.KeepAlive(compareSet)

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

// ClearStates removes all states from the state set.
func (set *StateSet) ClearStates() {
	var _arg0 *C.AtkStateSet // out

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))

	C.atk_state_set_clear_states(_arg0)
	runtime.KeepAlive(set)
}

// ContainsState checks whether the state for the specified type is in the
// specified set.
//
// The function takes the following parameters:
//
//    - typ: StateType.
//
// The function returns the following values:
//
//    - ok: TRUE if type is the state type is in set.
//
func (set *StateSet) ContainsState(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_contains_state(_arg0, _arg1)
	runtime.KeepAlive(set)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContainsStates checks whether the states for all the specified types are in
// the specified set.
//
// The function takes the following parameters:
//
//    - types: array of StateType.
//
// The function returns the following values:
//
//    - ok: TRUE if all the states for type are in set.
//
func (set *StateSet) ContainsStates(types []StateType) bool {
	var _arg0 *C.AtkStateSet  // out
	var _arg1 *C.AtkStateType // out
	var _arg2 C.gint
	var _cret C.gboolean // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg2 = (C.gint)(len(types))
	if len(types) > 0 {
		_arg1 = (*C.AtkStateType)(unsafe.Pointer(&types[0]))
	}

	_cret = C.atk_state_set_contains_states(_arg0, _arg1, _arg2)
	runtime.KeepAlive(set)
	runtime.KeepAlive(types)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEmpty checks whether the state set is empty, i.e. has no states set.
//
// The function returns the following values:
//
//    - ok: TRUE if set has no states set, otherwise FALSE.
//
func (set *StateSet) IsEmpty() bool {
	var _arg0 *C.AtkStateSet // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))

	_cret = C.atk_state_set_is_empty(_arg0)
	runtime.KeepAlive(set)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// OrSets constructs the union of the two sets.
//
// The function takes the following parameters:
//
//    - compareSet: another StateSet.
//
// The function returns the following values:
//
//    - stateSet (optional): new StateSet which is the union of the two sets,
//      returning NULL is empty.
//
func (set *StateSet) OrSets(compareSet *StateSet) *StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(compareSet).Native()))

	_cret = C.atk_state_set_or_sets(_arg0, _arg1)
	runtime.KeepAlive(set)
	runtime.KeepAlive(compareSet)

	var _stateSet *StateSet // out

	if _cret != nil {
		_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _stateSet
}

// RemoveState removes the state for the specified type from the state set.
//
// Note that because an StateSet is a read-only object, this method should be
// used to remove a state to a newly-created set which will then be returned by
// #atk_object_ref_state_set. It should not be used to modify the existing state
// of an object. See also #atk_object_notify_state_change.
//
// The function takes the following parameters:
//
//    - typ: Type.
//
// The function returns the following values:
//
//    - ok: TRUE if type was the state type is in set.
//
func (set *StateSet) RemoveState(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_remove_state(_arg0, _arg1)
	runtime.KeepAlive(set)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// XorSets constructs the exclusive-or of the two sets, returning NULL is empty.
// The set returned by this operation contains the states in exactly one of the
// two sets.
//
// The function takes the following parameters:
//
//    - compareSet: another StateSet.
//
// The function returns the following values:
//
//    - stateSet: new StateSet which contains the states which are in exactly one
//      of the two sets.
//
func (set *StateSet) XorSets(compareSet *StateSet) *StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(set).Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(compareSet).Native()))

	_cret = C.atk_state_set_xor_sets(_arg0, _arg1)
	runtime.KeepAlive(set)
	runtime.KeepAlive(compareSet)

	var _stateSet *StateSet // out

	_stateSet = wrapStateSet(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

// StateSetClass: instance of this type is always passed by reference.
type StateSetClass struct {
	*stateSetClass
}

// stateSetClass is the struct that's finalized.
type stateSetClass struct {
	native *C.AtkStateSetClass
}
