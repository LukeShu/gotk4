// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// NewAttrLetterSpacing: create a new letter-spacing attribute.
//
// The function takes the following parameters:
//
//    - letterSpacing: amount of extra space to add between graphemes of the
//      text, in Pango units.
//
// The function returns the following values:
//
//    - attribute: newly allocated PangoAttribute, which should be freed with
//      pango.Attribute.Destroy().
//
func NewAttrLetterSpacing(letterSpacing int) *Attribute {
	var _arg1 C.int             // out
	var _cret *C.PangoAttribute // in

	_arg1 = C.int(letterSpacing)

	_cret = C.pango_attr_letter_spacing_new(_arg1)
	runtime.KeepAlive(letterSpacing)

	var _attribute *Attribute // out

	_attribute = (*Attribute)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_attribute)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_attribute_destroy((*C.PangoAttribute)(intern.C))
		},
	)

	return _attribute
}
