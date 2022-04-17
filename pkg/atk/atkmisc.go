// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern void _gotk4_atk1_MiscClass_threads_enter(AtkMisc*);
// extern void _gotk4_atk1_MiscClass_threads_leave(AtkMisc*);
import "C"

// glib.Type values for atkmisc.go.
var GTypeMisc = externglib.Type(C.atk_misc_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeMisc, F: marshalMisc},
	})
}

// MiscOverrider contains methods that are overridable.
type MiscOverrider interface {
	externglib.Objector
	// ThreadsEnter: take the thread mutex for the GUI toolkit, if one exists.
	// (This method is implemented by the toolkit ATK implementation layer; for
	// instance, for GTK+, GAIL implements this via GDK_THREADS_ENTER).
	//
	// Deprecated: Since 2.12.
	ThreadsEnter()
	// ThreadsLeave: release the thread mutex for the GUI toolkit, if one
	// exists. This method, and atk_misc_threads_enter, are needed in some
	// situations by threaded application code which services ATK requests,
	// since fulfilling ATK requests often requires calling into the GUI
	// toolkit. If a long-running or potentially blocking call takes place
	// inside such a block, it should be bracketed by
	// atk_misc_threads_leave/atk_misc_threads_enter calls. (This method is
	// implemented by the toolkit ATK implementation layer; for instance, for
	// GTK+, GAIL implements this via GDK_THREADS_LEAVE).
	//
	// Deprecated: Since 2.12.
	ThreadsLeave()
}

// WrapMiscOverrider wraps the MiscOverrider
// interface implementation to access the instance methods.
func WrapMiscOverrider(obj MiscOverrider) *Misc {
	return wrapMisc(externglib.BaseObject(obj))
}

// Misc: set of utility functions for thread locking. This interface and all his
// related methods are deprecated since 2.12.
type Misc struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Misc)(nil)
)

func classInitMiscer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.AtkMiscClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.AtkMiscClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ThreadsEnter() }); ok {
		pclass.threads_enter = (*[0]byte)(C._gotk4_atk1_MiscClass_threads_enter)
	}

	if _, ok := goval.(interface{ ThreadsLeave() }); ok {
		pclass.threads_leave = (*[0]byte)(C._gotk4_atk1_MiscClass_threads_leave)
	}
}

//export _gotk4_atk1_MiscClass_threads_enter
func _gotk4_atk1_MiscClass_threads_enter(arg0 *C.AtkMisc) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ThreadsEnter() })

	iface.ThreadsEnter()
}

//export _gotk4_atk1_MiscClass_threads_leave
func _gotk4_atk1_MiscClass_threads_leave(arg0 *C.AtkMisc) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ThreadsLeave() })

	iface.ThreadsLeave()
}

func wrapMisc(obj *externglib.Object) *Misc {
	return &Misc{
		Object: obj,
	}
}

func marshalMisc(p uintptr) (interface{}, error) {
	return wrapMisc(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ThreadsEnter: take the thread mutex for the GUI toolkit, if one exists. (This
// method is implemented by the toolkit ATK implementation layer; for instance,
// for GTK+, GAIL implements this via GDK_THREADS_ENTER).
//
// Deprecated: Since 2.12.
func (misc *Misc) ThreadsEnter() {
	var _arg0 *C.AtkMisc // out

	_arg0 = (*C.AtkMisc)(unsafe.Pointer(externglib.InternObject(misc).Native()))

	C.atk_misc_threads_enter(_arg0)
	runtime.KeepAlive(misc)
}

// ThreadsLeave: release the thread mutex for the GUI toolkit, if one exists.
// This method, and atk_misc_threads_enter, are needed in some situations by
// threaded application code which services ATK requests, since fulfilling ATK
// requests often requires calling into the GUI toolkit. If a long-running or
// potentially blocking call takes place inside such a block, it should be
// bracketed by atk_misc_threads_leave/atk_misc_threads_enter calls. (This
// method is implemented by the toolkit ATK implementation layer; for instance,
// for GTK+, GAIL implements this via GDK_THREADS_LEAVE).
//
// Deprecated: Since 2.12.
func (misc *Misc) ThreadsLeave() {
	var _arg0 *C.AtkMisc // out

	_arg0 = (*C.AtkMisc)(unsafe.Pointer(externglib.InternObject(misc).Native()))

	C.atk_misc_threads_leave(_arg0)
	runtime.KeepAlive(misc)
}

// MiscGetInstance: obtain the singleton instance of AtkMisc for this
// application.
//
// Deprecated: Since 2.12.
//
// The function returns the following values:
//
//    - misc: singleton instance of AtkMisc for this application.
//
func MiscGetInstance() *Misc {
	var _cret *C.AtkMisc // in

	_cret = C.atk_misc_get_instance()

	var _misc *Misc // out

	_misc = wrapMisc(externglib.Take(unsafe.Pointer(_cret)))

	return _misc
}
