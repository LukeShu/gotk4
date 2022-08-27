// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// AtkHyperlink* _gotk4_atk1_HyperlinkImpl_virtual_get_hyperlink(void* fnptr, AtkHyperlinkImpl* arg0) {
//   return ((AtkHyperlink* (*)(AtkHyperlinkImpl*))(fnptr))(arg0);
// };
import "C"

// Hyperlink gets the hyperlink associated with this object.
//
// The function returns the following values:
//
//    - hyperlink: atkHyperlink object which points to this implementing
//      AtkObject.
//
func (impl *HyperlinkImpl) Hyperlink() *Hyperlink {
	var _arg0 *C.AtkHyperlinkImpl // out
	var _cret *C.AtkHyperlink     // in

	_arg0 = (*C.AtkHyperlinkImpl)(unsafe.Pointer(coreglib.InternObject(impl).Native()))

	_cret = C.atk_hyperlink_impl_get_hyperlink(_arg0)
	runtime.KeepAlive(impl)

	var _hyperlink *Hyperlink // out

	_hyperlink = wrapHyperlink(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _hyperlink
}

// Hyperlink gets the hyperlink associated with this object.
//
// The function returns the following values:
//
//    - hyperlink: atkHyperlink object which points to this implementing
//      AtkObject.
//
func (impl *HyperlinkImpl) hyperlink() *Hyperlink {
	gclass := (*C.AtkHyperlinkImplIface)(coreglib.PeekParentClass(impl))
	fnarg := gclass.get_hyperlink

	var _arg0 *C.AtkHyperlinkImpl // out
	var _cret *C.AtkHyperlink     // in

	_arg0 = (*C.AtkHyperlinkImpl)(unsafe.Pointer(coreglib.InternObject(impl).Native()))

	_cret = C._gotk4_atk1_HyperlinkImpl_virtual_get_hyperlink(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(impl)

	var _hyperlink *Hyperlink // out

	_hyperlink = wrapHyperlink(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _hyperlink
}
