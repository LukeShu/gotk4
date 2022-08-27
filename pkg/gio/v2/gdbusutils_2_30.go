// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// DBusGValueToGVariant converts a #GValue to a #GVariant of the type indicated
// by the type parameter.
//
// The conversion is using the following rules:
//
// - TYPE_STRING: 's', 'o', 'g' or 'ay'
//
// - TYPE_STRV: 'as', 'ao' or 'aay'
//
// - TYPE_BOOLEAN: 'b'
//
// - TYPE_UCHAR: 'y'
//
// - TYPE_INT: 'i', 'n'
//
// - TYPE_UINT: 'u', 'q'
//
// - TYPE_INT64 'x'
//
// - TYPE_UINT64: 't'
//
// - TYPE_DOUBLE: 'd'
//
// - TYPE_VARIANT: Any Type
//
// This can fail if e.g. gvalue is of type TYPE_STRING and type is
// ['i'][G-VARIANT-TYPE-INT32:CAPS]. It will also fail for any #GType (including
// e.g. TYPE_OBJECT and TYPE_BOXED derived-types) not in the table above.
//
// Note that if gvalue is of type TYPE_VARIANT and its value is NULL, the empty
// #GVariant instance (never NULL) for type is returned (e.g. 0 for scalar
// types, the empty string for string types, '/' for object path types, the
// empty array for any array type and so on).
//
// See the g_dbus_gvariant_to_gvalue() function for how to convert a #GVariant
// to a #GValue.
//
// The function takes the following parameters:
//
//    - gvalue to convert to a #GVariant.
//    - typ: Type.
//
// The function returns the following values:
//
//    - variant (never floating) of Type type holding the data from gvalue or an
//      empty #GVariant in case of failure. Free with g_variant_unref().
//
func DBusGValueToGVariant(gvalue *coreglib.Value, typ *glib.VariantType) *glib.Variant {
	var _arg1 *C.GValue       // out
	var _arg2 *C.GVariantType // out
	var _cret *C.GVariant     // in

	_arg1 = (*C.GValue)(unsafe.Pointer(gvalue.Native()))
	_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_dbus_gvalue_to_gvariant(_arg1, _arg2)
	runtime.KeepAlive(gvalue)
	runtime.KeepAlive(typ)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// DBusGVariantToGValue converts a #GVariant to a #GValue. If value is floating,
// it is consumed.
//
// The rules specified in the g_dbus_gvalue_to_gvariant() function are used -
// this function is essentially its reverse form. So, a #GVariant containing any
// basic or string array type will be converted to a #GValue containing a basic
// value or string array. Any other #GVariant (handle, variant, tuple, dict
// entry) will be converted to a #GValue containing that #GVariant.
//
// The conversion never fails - a valid #GValue is always returned in
// out_gvalue.
//
// The function takes the following parameters:
//
//    - value: #GVariant.
//
// The function returns the following values:
//
//    - outGvalue: return location pointing to a zero-filled (uninitialized)
//      #GValue.
//
func DBusGVariantToGValue(value *glib.Variant) coreglib.Value {
	var _arg1 *C.GVariant // out
	var _arg2 C.GValue    // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_dbus_gvariant_to_gvalue(_arg1, &_arg2)
	runtime.KeepAlive(value)

	var _outGvalue coreglib.Value // out

	_outGvalue = *coreglib.ValueFromNative(unsafe.Pointer((&_arg2)))

	return _outGvalue
}
