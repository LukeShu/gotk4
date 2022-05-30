// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// BindingsActivateEvent looks up key bindings for object to find one matching
// event, and if one was found, activate it.
//
// The function takes the following parameters:
//
//    - object (generally must be a widget).
//    - event: EventKey.
//
// The function returns the following values:
//
//    - ok: TRUE if a matching key binding was found.
//
func BindingsActivateEvent(object *coreglib.Object, event *gdk.EventKey) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(object.Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(event)))
	*(**coreglib.Object)(unsafe.Pointer(&args[0])) = _arg0
	*(**gdk.EventKey)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "bindings_activate_event").Invoke(args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(object)
	runtime.KeepAlive(event)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingArg holds the data associated with an argument for a key binding
// signal emission as stored in BindingSignal.
//
// An instance of this type is always passed by reference.
type BindingArg struct {
	*bindingArg
}

// bindingArg is the struct that's finalized.
type bindingArg struct {
	native *C.GtkBindingArg
}

// BindingEntry: each key binding element of a binding sets binding list is
// represented by a GtkBindingEntry.
//
// An instance of this type is always passed by reference.
type BindingEntry struct {
	*bindingEntry
}

// bindingEntry is the struct that's finalized.
type bindingEntry struct {
	native *C.GtkBindingEntry
}

// BindingSet: binding set maintains a list of activatable key bindings. A
// single binding set can match multiple types of widgets. Similar to style
// contexts, can be matched by any information contained in a widgets
// WidgetPath. When a binding within a set is matched upon activation, an action
// signal is emitted on the target widget to carry out the actual activation.
//
// An instance of this type is always passed by reference.
type BindingSet struct {
	*bindingSet
}

// bindingSet is the struct that's finalized.
type bindingSet struct {
	native *C.GtkBindingSet
}

// SetName: unique name of this binding set.
func (b *BindingSet) SetName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(b.native.set_name)))
	return v
}

// Priority: unused.
func (b *BindingSet) Priority() int {
	var v int // out
	v = int(b.native.priority)
	return v
}

// Entries: key binding entries in this binding set.
func (b *BindingSet) Entries() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(gextras.NewStructNative(unsafe.Pointer(b.native.entries)))
	return v
}

// Current: implementation detail.
func (b *BindingSet) Current() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(gextras.NewStructNative(unsafe.Pointer(b.native.current)))
	return v
}

// BindingSetFind: find a binding set by its globally unique name.
//
// The set_name can either be a name used for gtk_binding_set_new() or the type
// name of a class used in gtk_binding_set_by_class().
//
// The function takes the following parameters:
//
//    - setName: unique binding set name.
//
// The function returns the following values:
//
//    - bindingSet (optional): NULL or the specified binding set.
//
func BindingSetFind(setName string) *BindingSet {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(setName)))
	defer C.free(unsafe.Pointer(_arg0))
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "find").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(setName)

	var _bindingSet *BindingSet // out

	if _cret != nil {
		_bindingSet = (*BindingSet)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _bindingSet
}

// BindingSignal stores the necessary information to activate a widget in
// response to a key press via a signal emission.
//
// An instance of this type is always passed by reference.
type BindingSignal struct {
	*bindingSignal
}

// bindingSignal is the struct that's finalized.
type bindingSignal struct {
	native *C.GtkBindingSignal
}
