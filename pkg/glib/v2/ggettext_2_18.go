// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// Dgettext: this function is a wrapper of dgettext() which does not
// translate the message if the default domain as set with textdomain() has no
// translations for the current locale.
//
// The advantage of using this function over dgettext() proper is that libraries
// using this function (like GTK+) will not use translations if the application
// using the library does not have translations for the current locale.
// This results in a consistent English-only interface instead of one having
// partial translations. For this feature to work, the call to textdomain() and
// setlocale() should precede any g_dgettext() invocations. For GTK+, it means
// calling textdomain() before gtk_init or its variants.
//
// This function disables translations if and only if upon its first call all
// the following conditions hold:
//
// - domain is not NULL
//
// - textdomain() has been called to set a default text domain
//
// - there is no translations available for the default text domain and the
// current locale
//
// - current locale is not "C" or any English locales (those starting with
// "en_")
//
// Note that this behavior may not be desired for example if an application has
// its untranslated messages in a language other than English. In those cases
// the application should call textdomain() after initializing GTK+.
//
// Applications should normally not use this function directly, but use the _()
// macro for translations.
//
// The function takes the following parameters:
//
//   - domain (optional): translation domain to use, or NULL to use the domain
//     set with textdomain().
//   - msgid: message to translate.
//
// The function returns the following values:
//
//   - utf8: translated string.
//
func Dgettext(domain, msgid string) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _cret *C.gchar // in

	if domain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dgettext(_arg1, _arg2)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msgid)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Dngettext: this function is a wrapper of dngettext() which does not
// translate the message if the default domain as set with textdomain() has no
// translations for the current locale.
//
// See g_dgettext() for details of how this differs from dngettext() proper.
//
// The function takes the following parameters:
//
//   - domain (optional): translation domain to use, or NULL to use the domain
//     set with textdomain().
//   - msgid: message to translate.
//   - msgidPlural: plural form of the message.
//   - n: quantity for which translation is needed.
//
// The function returns the following values:
//
//   - utf8: translated string.
//
func Dngettext(domain, msgid, msgidPlural string, n uint32) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _arg3 *C.gchar // out
	var _arg4 C.gulong // out
	var _cret *C.gchar // in

	if domain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(msgidPlural)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gulong(n)

	_cret = C.g_dngettext(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msgid)
	runtime.KeepAlive(msgidPlural)
	runtime.KeepAlive(n)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Dpgettext2: this function is a variant of g_dgettext() which supports a
// disambiguating message context. GNU gettext uses the '\004' character to
// separate the message context and message id in msgctxtid.
//
// This uses g_dgettext() internally. See that functions for differences with
// dgettext() proper.
//
// This function differs from C_() in that it is not a macro and thus you may
// use non-string-literals as context and msgid arguments.
//
// The function takes the following parameters:
//
//   - domain (optional): translation domain to use, or NULL to use the domain
//     set with textdomain().
//   - context: message context.
//   - msgid: message.
//
// The function returns the following values:
//
//   - utf8: translated string.
//
func Dpgettext2(domain, context, msgid string) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _arg3 *C.gchar // out
	var _cret *C.gchar // in

	if domain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(context)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(msgid)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_dpgettext2(_arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(context)
	runtime.KeepAlive(msgid)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
