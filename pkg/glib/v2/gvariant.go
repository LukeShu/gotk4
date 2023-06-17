// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// GType values.
var (
	GTypeVariantBuilder = coreglib.Type(C.g_variant_builder_get_type())
	GTypeVariant        = coreglib.Type(coreglib.TypeVariant)
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVariantBuilder, F: marshalVariantBuilder},
		coreglib.TypeMarshaler{T: GTypeVariant, F: marshalVariant},
	})
}

// VariantParseError: error codes returned by parsing text-format GVariants.
type VariantParseError C.gint

const (
	// VariantParseErrorFailed: generic error (unused).
	VariantParseErrorFailed VariantParseError = iota
	// VariantParseErrorBasicTypeExpected: non-basic Type was given where a
	// basic type was expected.
	VariantParseErrorBasicTypeExpected
	// VariantParseErrorCannotInferType: cannot infer the Type.
	VariantParseErrorCannotInferType
	// VariantParseErrorDefiniteTypeExpected: indefinite Type was given where a
	// definite type was expected.
	VariantParseErrorDefiniteTypeExpected
	// VariantParseErrorInputNotAtEnd: extra data after parsing finished.
	VariantParseErrorInputNotAtEnd
	// VariantParseErrorInvalidCharacter: invalid character in number or unicode
	// escape.
	VariantParseErrorInvalidCharacter
	// VariantParseErrorInvalidFormatString: not a valid #GVariant format
	// string.
	VariantParseErrorInvalidFormatString
	// VariantParseErrorInvalidObjectPath: not a valid object path.
	VariantParseErrorInvalidObjectPath
	// VariantParseErrorInvalidSignature: not a valid type signature.
	VariantParseErrorInvalidSignature
	// VariantParseErrorInvalidTypeString: not a valid #GVariant type string.
	VariantParseErrorInvalidTypeString
	// VariantParseErrorNoCommonType: could not find a common type for array
	// entries.
	VariantParseErrorNoCommonType
	// VariantParseErrorNumberOutOfRange: numerical value is out of range of the
	// given type.
	VariantParseErrorNumberOutOfRange
	// VariantParseErrorNumberTooBig: numerical value is out of range for any
	// type.
	VariantParseErrorNumberTooBig
	// VariantParseErrorTypeError: cannot parse as variant of the specified
	// type.
	VariantParseErrorTypeError
	// VariantParseErrorUnexpectedToken: unexpected token was encountered.
	VariantParseErrorUnexpectedToken
	// VariantParseErrorUnknownKeyword: unknown keyword was encountered.
	VariantParseErrorUnknownKeyword
	// VariantParseErrorUnterminatedStringConstant: unterminated string
	// constant.
	VariantParseErrorUnterminatedStringConstant
	// VariantParseErrorValueExpected: no value given.
	VariantParseErrorValueExpected
	// VariantParseErrorRecursion: variant was too deeply nested; #GVariant is
	// only guaranteed to handle nesting up to 64 levels (Since: 2.64).
	VariantParseErrorRecursion
)

// String returns the name in string for VariantParseError.
func (v VariantParseError) String() string {
	switch v {
	case VariantParseErrorFailed:
		return "Failed"
	case VariantParseErrorBasicTypeExpected:
		return "BasicTypeExpected"
	case VariantParseErrorCannotInferType:
		return "CannotInferType"
	case VariantParseErrorDefiniteTypeExpected:
		return "DefiniteTypeExpected"
	case VariantParseErrorInputNotAtEnd:
		return "InputNotAtEnd"
	case VariantParseErrorInvalidCharacter:
		return "InvalidCharacter"
	case VariantParseErrorInvalidFormatString:
		return "InvalidFormatString"
	case VariantParseErrorInvalidObjectPath:
		return "InvalidObjectPath"
	case VariantParseErrorInvalidSignature:
		return "InvalidSignature"
	case VariantParseErrorInvalidTypeString:
		return "InvalidTypeString"
	case VariantParseErrorNoCommonType:
		return "NoCommonType"
	case VariantParseErrorNumberOutOfRange:
		return "NumberOutOfRange"
	case VariantParseErrorNumberTooBig:
		return "NumberTooBig"
	case VariantParseErrorTypeError:
		return "TypeError"
	case VariantParseErrorUnexpectedToken:
		return "UnexpectedToken"
	case VariantParseErrorUnknownKeyword:
		return "UnknownKeyword"
	case VariantParseErrorUnterminatedStringConstant:
		return "UnterminatedStringConstant"
	case VariantParseErrorValueExpected:
		return "ValueExpected"
	case VariantParseErrorRecursion:
		return "Recursion"
	default:
		return fmt.Sprintf("VariantParseError(%d)", v)
	}
}

// VariantBuilder: utility type for constructing container-type #GVariant
// instances.
//
// This is an opaque structure and may only be accessed using the following
// functions.
//
// Builder is not threadsafe in any way. Do not attempt to access it from more
// than one thread.
//
// An instance of this type is always passed by reference.
type VariantBuilder struct {
	*variantBuilder
}

// variantBuilder is the struct that's finalized.
type variantBuilder struct {
	native *C.GVariantBuilder
}

func marshalVariantBuilder(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &VariantBuilder{&variantBuilder{(*C.GVariantBuilder)(b)}}, nil
}

// NewVariantBuilder constructs a struct VariantBuilder.
func NewVariantBuilder(typ *VariantType) *VariantBuilder {
	var _arg1 *C.GVariantType    // out
	var _cret *C.GVariantBuilder // in

	_arg1 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_builder_new(_arg1)
	runtime.KeepAlive(typ)

	var _variantBuilder *VariantBuilder // out

	_variantBuilder = (*VariantBuilder)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variantBuilder)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_builder_unref((*C.GVariantBuilder)(intern.C))
		},
	)

	return _variantBuilder
}

// AddValue adds value to builder.
//
// It is an error to call this function in any way that would create an
// inconsistent value to be constructed. Some examples of this are putting
// different types of items into an array, putting the wrong types or number of
// items in a tuple, putting more than one value into a variant, etc.
//
// If value is a floating reference (see g_variant_ref_sink()), the builder
// instance takes ownership of value.
//
// The function takes the following parameters:
//
//   - value: #GVariant.
//
func (builder *VariantBuilder) AddValue(value *Variant) {
	var _arg0 *C.GVariantBuilder // out
	var _arg1 *C.GVariant        // out

	_arg0 = (*C.GVariantBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_variant_builder_add_value(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(value)
}

// Close closes the subcontainer inside the given builder that was opened by the
// most recent call to g_variant_builder_open().
//
// It is an error to call this function in any way that would create an
// inconsistent value to be constructed (ie: too few values added to the
// subcontainer).
func (builder *VariantBuilder) Close() {
	var _arg0 *C.GVariantBuilder // out

	_arg0 = (*C.GVariantBuilder)(gextras.StructNative(unsafe.Pointer(builder)))

	C.g_variant_builder_close(_arg0)
	runtime.KeepAlive(builder)
}

// End ends the builder process and returns the constructed value.
//
// It is not permissible to use builder in any way after this call except
// for reference counting operations (in the case of a heap-allocated
// Builder) or by reinitialising it with g_variant_builder_init() (in the
// case of stack-allocated). This means that for the stack-allocated builders
// there is no need to call g_variant_builder_clear() after the call to
// g_variant_builder_end().
//
// It is an error to call this function in any way that would create an
// inconsistent value to be constructed (ie: insufficient number of items added
// to a container with a specific number of children required). It is also an
// error to call this function if the builder was created with an indefinite
// array or maybe type and no children have been added; in this case it is
// impossible to infer the type of the empty array.
//
// The function returns the following values:
//
//   - variant: new, floating, #GVariant.
//
func (builder *VariantBuilder) End() *Variant {
	var _arg0 *C.GVariantBuilder // out
	var _cret *C.GVariant        // in

	_arg0 = (*C.GVariantBuilder)(gextras.StructNative(unsafe.Pointer(builder)))

	_cret = C.g_variant_builder_end(_arg0)
	runtime.KeepAlive(builder)

	var _variant *Variant // out

	_variant = (*Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// Open opens a subcontainer inside the given builder. When done adding items to
// the subcontainer, g_variant_builder_close() must be called. type is the type
// of the container: so to build a tuple of several values, type must include
// the tuple itself.
//
// It is an error to call this function in any way that would cause an
// inconsistent value to be constructed (ie: adding too many values or a value
// of an incorrect type).
//
// Example of building a nested variant:
//
//    GVariantBuilder builder;
//    guint32 some_number = get_number ();
//    g_autoptr (GHashTable) some_dict = get_dict ();
//    GHashTableIter iter;
//    const gchar *key;
//    const GVariant *value;
//    g_autoptr (GVariant) output = NULL;
//
//    g_variant_builder_init (&builder, G_VARIANT_TYPE ("(ua{sv})"));
//    g_variant_builder_add (&builder, "u", some_number);
//    g_variant_builder_open (&builder, G_VARIANT_TYPE ("a{sv}"));
//
//    g_hash_table_iter_init (&iter, some_dict);
//    while (g_hash_table_iter_next (&iter, (gpointer *) &key, (gpointer *) &value))
//      {
//        g_variant_builder_open (&builder, G_VARIANT_TYPE ("{sv}"));
//        g_variant_builder_add (&builder, "s", key);
//        g_variant_builder_add (&builder, "v", value);
//        g_variant_builder_close (&builder);
//      }
//
//    g_variant_builder_close (&builder);
//
//    output = g_variant_builder_end (&builder);.
//
// The function takes the following parameters:
//
//   - typ of the container.
//
func (builder *VariantBuilder) Open(typ *VariantType) {
	var _arg0 *C.GVariantBuilder // out
	var _arg1 *C.GVariantType    // out

	_arg0 = (*C.GVariantBuilder)(gextras.StructNative(unsafe.Pointer(builder)))
	_arg1 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	C.g_variant_builder_open(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(typ)
}

func marshalVariant(p uintptr) (interface{}, error) {
	_cret := C.g_value_dup_variant((*C.GValue)(unsafe.Pointer(p)))
	if _cret == nil {
		return (*Variant)(nil), nil
	}

	_variant := (*Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)
	return _variant, nil
}

// ForEach iterates over items in value. The iteration breaks out once f
// returns true. This method wraps around g_variant_iter_new.
func (value *Variant) ForEach(f func(*Variant) (stop bool)) {
	valueNative := (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	var iter C.GVariantIter
	C.g_variant_iter_init(&iter, valueNative)

	next := func() *Variant {
		item := C.g_variant_iter_next_value(&iter)
		if item == nil {
			return nil
		}

		variant := (*Variant)(gextras.NewStructNative(unsafe.Pointer(item)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)

		return variant
	}

	for item := next(); item != nil; item = next() {
		if f(item) {
			break
		}
	}

	runtime.KeepAlive(value)
}
