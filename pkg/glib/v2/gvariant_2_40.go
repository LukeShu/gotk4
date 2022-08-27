// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// GType values.
var (
	GTypeVariantDict = coreglib.Type(C.g_variant_dict_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVariantDict, F: marshalVariantDict},
	})
}

// VariantParseErrorPrintContext pretty-prints a message showing the context of
// a #GVariant parse error within the string for which parsing was attempted.
//
// The resulting string is suitable for output to the console or other monospace
// media where newlines are treated in the usual way.
//
// The message will typically look something like one of the following:
//
//    unterminated string constant:
//      (1, 2, 3, 'abc
//                ^^^^
//
// or
//
//    unable to find a common type:
//      [1, 2, 3, 'str']
//       ^        ^^^^^
//
// The format of the message may change in a future version.
//
// error must have come from a failed attempt to g_variant_parse() and
// source_str must be exactly the same string that caused the error. If
// source_str was not nul-terminated when you passed it to g_variant_parse()
// then you must add nul termination before using this function.
//
// The function takes the following parameters:
//
//    - err from the ParseError domain.
//    - sourceStr: string that was given to the parser.
//
// The function returns the following values:
//
//    - utf8: printed message.
//
func VariantParseErrorPrintContext(err error, sourceStr string) string {
	var _arg1 *C.GError // out
	var _arg2 *C.gchar  // out
	var _cret *C.gchar  // in

	if err != nil {
		_arg1 = (*C.GError)(gerror.New(err))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(sourceStr)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_variant_parse_error_print_context(_arg1, _arg2)
	runtime.KeepAlive(err)
	runtime.KeepAlive(sourceStr)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// VariantDict is a mutable interface to #GVariant dictionaries.
//
// It can be used for doing a sequence of dictionary lookups in an efficient way
// on an existing #GVariant dictionary or it can be used to construct new
// dictionaries with a hashtable-like interface. It can also be used for taking
// existing dictionaries and modifying them in order to create new ones.
//
// Dict can only be used with G_VARIANT_TYPE_VARDICT dictionaries.
//
// It is possible to use Dict allocated on the stack or on the heap. When using
// a stack-allocated Dict, you begin with a call to g_variant_dict_init() and
// free the resources with a call to g_variant_dict_clear().
//
// Heap-allocated Dict follows normal refcounting rules: you allocate it with
// g_variant_dict_new() and use g_variant_dict_ref() and g_variant_dict_unref().
//
// g_variant_dict_end() is used to convert the Dict back into a dictionary-type
// #GVariant. When used with stack-allocated instances, this also implicitly
// frees all associated memory, but for heap-allocated instances, you must still
// call g_variant_dict_unref() afterwards.
//
// You will typically want to use a heap-allocated Dict when you expose it as
// part of an API. For most other uses, the stack-allocated form will be more
// convenient.
//
// Consider the following two examples that do the same thing in each style:
// take an existing dictionary and look up the "count" uint32 key, adding 1 to
// it if it is found, or returning an error if the key is not found. Each
// returns the new dictionary as a floating #GVariant.
//
// Using a stack-allocated GVariantDict
//
//      GVariant *
//      add_to_count (GVariant  *orig,
//                    GError   **error)
//      {
//        GVariantDict *dict;
//        GVariant *result;
//        guint32 count;
//
//        dict = g_variant_dict_new (orig);
//
//        if (g_variant_dict_lookup (dict, "count", "u", &count))
//          {
//            g_variant_dict_insert (dict, "count", "u", count + 1);
//            result = g_variant_dict_end (dict);
//          }
//        else
//          {
//            g_set_error (...);
//            result = NULL;
//          }
//
//        g_variant_dict_unref (dict);
//
//        return result;
//      }
//
// An instance of this type is always passed by reference.
type VariantDict struct {
	*variantDict
}

// variantDict is the struct that's finalized.
type variantDict struct {
	native *C.GVariantDict
}

func marshalVariantDict(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &VariantDict{&variantDict{(*C.GVariantDict)(b)}}, nil
}

// NewVariantDict constructs a struct VariantDict.
func NewVariantDict(fromAsv *Variant) *VariantDict {
	var _arg1 *C.GVariant     // out
	var _cret *C.GVariantDict // in

	if fromAsv != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(fromAsv)))
	}

	_cret = C.g_variant_dict_new(_arg1)
	runtime.KeepAlive(fromAsv)

	var _variantDict *VariantDict // out

	_variantDict = (*VariantDict)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variantDict)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_dict_unref((*C.GVariantDict)(intern.C))
		},
	)

	return _variantDict
}

// Clear releases all memory associated with a Dict without freeing the Dict
// structure itself.
//
// It typically only makes sense to do this on a stack-allocated Dict if you
// want to abort building the value part-way through. This function need not be
// called if you call g_variant_dict_end() and it also doesn't need to be called
// on dicts allocated with g_variant_dict_new (see g_variant_dict_unref() for
// that).
//
// It is valid to call this function on either an initialised Dict or one that
// was previously cleared by an earlier call to g_variant_dict_clear() but it is
// not valid to call this function on uninitialised memory.
func (dict *VariantDict) Clear() {
	var _arg0 *C.GVariantDict // out

	_arg0 = (*C.GVariantDict)(gextras.StructNative(unsafe.Pointer(dict)))

	C.g_variant_dict_clear(_arg0)
	runtime.KeepAlive(dict)
}

// Contains checks if key exists in dict.
//
// The function takes the following parameters:
//
//    - key to look up in the dictionary.
//
// The function returns the following values:
//
//    - ok: TRUE if key is in dict.
//
func (dict *VariantDict) Contains(key string) bool {
	var _arg0 *C.GVariantDict // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantDict)(gextras.StructNative(unsafe.Pointer(dict)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_variant_dict_contains(_arg0, _arg1)
	runtime.KeepAlive(dict)
	runtime.KeepAlive(key)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// End returns the current value of dict as a #GVariant of type
// G_VARIANT_TYPE_VARDICT, clearing it in the process.
//
// It is not permissible to use dict in any way after this call except for
// reference counting operations (in the case of a heap-allocated Dict) or by
// reinitialising it with g_variant_dict_init() (in the case of
// stack-allocated).
//
// The function returns the following values:
//
//    - variant: new, floating, #GVariant.
//
func (dict *VariantDict) End() *Variant {
	var _arg0 *C.GVariantDict // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GVariantDict)(gextras.StructNative(unsafe.Pointer(dict)))

	_cret = C.g_variant_dict_end(_arg0)
	runtime.KeepAlive(dict)

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

// InsertValue inserts (or replaces) a key in a Dict.
//
// value is consumed if it is floating.
//
// The function takes the following parameters:
//
//    - key to insert a value for.
//    - value to insert.
//
func (dict *VariantDict) InsertValue(key string, value *Variant) {
	var _arg0 *C.GVariantDict // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariant     // out

	_arg0 = (*C.GVariantDict)(gextras.StructNative(unsafe.Pointer(dict)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_variant_dict_insert_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dict)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)
}

// LookupValue looks up a value in a Dict.
//
// If key is not found in dictionary, NULL is returned.
//
// The expected_type string specifies what type of value is expected. If the
// value associated with key has a different type then NULL is returned.
//
// If the key is found and the value has the correct type, it is returned. If
// expected_type was specified then any non-NULL return value will have this
// type.
//
// The function takes the following parameters:
//
//    - key to look up in the dictionary.
//    - expectedType (optional) or NULL.
//
// The function returns the following values:
//
//    - variant: value of the dictionary key, or NULL.
//
func (dict *VariantDict) LookupValue(key string, expectedType *VariantType) *Variant {
	var _arg0 *C.GVariantDict // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.GVariantType // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GVariantDict)(gextras.StructNative(unsafe.Pointer(dict)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	if expectedType != nil {
		_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(expectedType)))
	}

	_cret = C.g_variant_dict_lookup_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dict)
	runtime.KeepAlive(key)
	runtime.KeepAlive(expectedType)

	var _variant *Variant // out

	_variant = (*Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// Remove removes a key and its associated value from a Dict.
//
// The function takes the following parameters:
//
//    - key to remove.
//
// The function returns the following values:
//
//    - ok: TRUE if the key was found and removed.
//
func (dict *VariantDict) Remove(key string) bool {
	var _arg0 *C.GVariantDict // out
	var _arg1 *C.gchar        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantDict)(gextras.StructNative(unsafe.Pointer(dict)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_variant_dict_remove(_arg0, _arg1)
	runtime.KeepAlive(dict)
	runtime.KeepAlive(key)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
