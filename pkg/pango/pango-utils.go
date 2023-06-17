// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// ParseStretch parses a font stretch.
//
// The allowed values are "ultra_condensed", "extra_condensed", "condensed",
// "semi_condensed", "normal", "semi_expanded", "expanded", "extra_expanded" and
// "ultra_expanded". Case variations are ignored and the '_' characters may be
// omitted.
//
// The function takes the following parameters:
//
//   - str: string to parse.
//   - warn: if TRUE, issue a g_warning() on bad input.
//
// The function returns the following values:
//
//   - stretch: PangoStretch to store the result in.
//   - ok: TRUE if str was successfully parsed.
//
func ParseStretch(str string, warn bool) (Stretch, bool) {
	var _arg1 *C.char        // out
	var _arg2 C.PangoStretch // in
	var _arg3 C.gboolean     // out
	var _cret C.gboolean     // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_stretch(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _stretch Stretch // out
	var _ok bool         // out

	_stretch = Stretch(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _stretch, _ok
}

// ParseStyle parses a font style.
//
// The allowed values are "normal", "italic" and "oblique", case variations
// being ignored.
//
// The function takes the following parameters:
//
//   - str: string to parse.
//   - warn: if TRUE, issue a g_warning() on bad input.
//
// The function returns the following values:
//
//   - style: PangoStyle to store the result in.
//   - ok: TRUE if str was successfully parsed.
//
func ParseStyle(str string, warn bool) (Style, bool) {
	var _arg1 *C.char      // out
	var _arg2 C.PangoStyle // in
	var _arg3 C.gboolean   // out
	var _cret C.gboolean   // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_style(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _style Style // out
	var _ok bool     // out

	_style = Style(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _style, _ok
}

// ParseVariant parses a font variant.
//
// The allowed values are "normal" and "smallcaps" or "small_caps", case
// variations being ignored.
//
// The function takes the following parameters:
//
//   - str: string to parse.
//   - warn: if TRUE, issue a g_warning() on bad input.
//
// The function returns the following values:
//
//   - variant: PangoVariant to store the result in.
//   - ok: TRUE if str was successfully parsed.
//
func ParseVariant(str string, warn bool) (Variant, bool) {
	var _arg1 *C.char        // out
	var _arg2 C.PangoVariant // in
	var _arg3 C.gboolean     // out
	var _cret C.gboolean     // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_variant(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _variant Variant // out
	var _ok bool         // out

	_variant = Variant(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _variant, _ok
}

// ParseWeight parses a font weight.
//
// The allowed values are "heavy", "ultrabold", "bold", "normal", "light",
// "ultraleight" and integers. Case variations are ignored.
//
// The function takes the following parameters:
//
//   - str: string to parse.
//   - warn: if TRUE, issue a g_warning() on bad input.
//
// The function returns the following values:
//
//   - weight: PangoWeight to store the result in.
//   - ok: TRUE if str was successfully parsed.
//
func ParseWeight(str string, warn bool) (Weight, bool) {
	var _arg1 *C.char       // out
	var _arg2 C.PangoWeight // in
	var _arg3 C.gboolean    // out
	var _cret C.gboolean    // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_weight(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _weight Weight // out
	var _ok bool       // out

	_weight = Weight(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _weight, _ok
}

// SplitFileList splits a G_SEARCHPATH_SEPARATOR-separated list of files,
// stripping white space and substituting ~/ with $HOME/.
//
// Deprecated: since version 1.38.
//
// The function takes the following parameters:
//
//   - str: G_SEARCHPATH_SEPARATOR separated list of filenames.
//
// The function returns the following values:
//
//   - utf8s: list of strings to be freed with g_strfreev().
//
func SplitFileList(str string) []string {
	var _arg1 *C.char  // out
	var _cret **C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_split_file_list(_arg1)
	runtime.KeepAlive(str)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// TrimString trims leading and trailing whitespace from a string.
//
// Deprecated: since version 1.38.
//
// The function takes the following parameters:
//
//   - str: string.
//
// The function returns the following values:
//
//   - utf8: newly-allocated string that must be freed with g_free().
//
func TrimString(str string) string {
	var _arg1 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_trim_string(_arg1)
	runtime.KeepAlive(str)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
