// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// FilenameDisplayBasename returns the display basename for the particular
// filename, guaranteed to be valid UTF-8. The display name might not be
// identical to the filename, for instance there might be problems converting it
// to UTF-8, and some files can be translated in the display.
//
// If GLib cannot make sense of the encoding of filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if filename was in an invalid
// encoding.
//
// You must pass the whole absolute pathname to this functions so that
// translation of well known locations can be done.
//
// This function is preferred over g_filename_display_name() if you know the
// whole path, as it allows translation.
//
// The function takes the following parameters:
//
//   - filename: absolute pathname in the GLib file name encoding.
//
// The function returns the following values:
//
//   - utf8: newly allocated string containing a rendition of the basename of
//     the filename in valid UTF-8.
//
func FilenameDisplayBasename(filename string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_filename_display_basename(_arg1)
	runtime.KeepAlive(filename)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FilenameDisplayName converts a filename into a valid UTF-8 string.
// The conversion is not necessarily reversible, so you should keep the original
// around and use the return value of this function only for display purposes.
// Unlike g_filename_to_utf8(), the result is guaranteed to be non-NULL even if
// the filename actually isn't in the GLib file name encoding.
//
// If GLib cannot make sense of the encoding of filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if filename was in an invalid
// encoding.
//
// If you know the whole pathname of the file you should use
// g_filename_display_basename(), since that allows location-based translation
// of filenames.
//
// The function takes the following parameters:
//
//   - filename: pathname hopefully in the GLib file name encoding.
//
// The function returns the following values:
//
//   - utf8: newly allocated string containing a rendition of the filename in
//     valid UTF-8.
//
func FilenameDisplayName(filename string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_filename_display_name(_arg1)
	runtime.KeepAlive(filename)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// GetFilenameCharsets determines the preferred character sets used for
// filenames. The first character set from the charsets is the filename
// encoding, the subsequent character sets are used when trying to generate a
// displayable representation of a filename, see g_filename_display_name().
//
// On Unix, the character sets are determined by consulting the environment
// variables G_FILENAME_ENCODING and G_BROKEN_FILENAMES. On Windows,
// the character set used in the GLib API is always UTF-8 and said environment
// variables have no effect.
//
// G_FILENAME_ENCODING may be set to a comma-separated list of character set
// names. The special token "\locale" is taken to mean the character set
// for the [current locale][setlocale]. If G_FILENAME_ENCODING is not set,
// but G_BROKEN_FILENAMES is, the character set of the current locale is taken
// as the filename encoding. If neither environment variable is set, UTF-8 is
// taken as the filename encoding, but the character set of the current locale
// is also put in the list of encodings.
//
// The returned charsets belong to GLib and must not be freed.
//
// Note that on Unix, regardless of the locale character set or
// G_FILENAME_ENCODING value, the actual file names present on a system might be
// in any random encoding or just gibberish.
//
// The function returns the following values:
//
//   - filenameCharsets: return location for the NULL-terminated list of
//     encoding names.
//   - ok: TRUE if the filename encoding is UTF-8.
//
func GetFilenameCharsets() ([]string, bool) {
	var _arg1 **C.gchar  // in
	var _cret C.gboolean // in

	_cret = C.g_get_filename_charsets(&_arg1)

	var _filenameCharsets []string // out
	var _ok bool                   // out

	{
		var i int
		var z *C.gchar
		for p := _arg1; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_arg1, i)
		_filenameCharsets = make([]string, i)
		for i := range src {
			_filenameCharsets[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _filenameCharsets, _ok
}

// URIListExtractURIs splits an URI list conforming to the text/uri-list mime
// type defined in RFC 2483 into individual URIs, discarding any comments.
// The URIs are not validated.
//
// The function takes the following parameters:
//
//   - uriList: URI list.
//
// The function returns the following values:
//
//   - utf8s: newly allocated NULL-terminated list of strings holding the
//     individual URIs. The array should be freed with g_strfreev().
//
func URIListExtractURIs(uriList string) []string {
	var _arg1 *C.gchar  // out
	var _cret **C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uriList)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_uri_list_extract_uris(_arg1)
	runtime.KeepAlive(uriList)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
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
