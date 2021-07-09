// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_keymap_get_type()), F: marshalKeymap},
	})
}

// KeyvalConvertCase obtains the upper- and lower-case versions of the keyval
// @symbol. Examples of keyvals are K_KEY_a, K_KEY_Enter, K_KEY_F1, etc.
func KeyvalConvertCase(symbol uint) (lower uint, upper uint) {
	var _arg1 C.guint // out
	var _arg2 C.guint // in
	var _arg3 C.guint // in

	_arg1 = C.guint(symbol)

	C.gdk_keyval_convert_case(_arg1, &_arg2, &_arg3)

	var _lower uint // out
	var _upper uint // out

	_lower = uint(_arg2)
	_upper = uint(_arg3)

	return _lower, _upper
}

// KeyvalFromName converts a key name to a key value.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalFromName(keyvalName string) uint {
	var _arg1 *C.gchar // out
	var _cret C.guint  // in

	_arg1 = (*C.gchar)(C.CString(keyvalName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_keyval_from_name(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// KeyvalIsLower returns true if the given key value is in lower case.
func KeyvalIsLower(keyval uint) bool {
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_is_lower(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyvalIsUpper returns true if the given key value is in upper case.
func KeyvalIsUpper(keyval uint) bool {
	var _arg1 C.guint    // out
	var _cret C.gboolean // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_is_upper(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyvalName converts a key value into a symbolic name.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalName(keyval uint) string {
	var _arg1 C.guint  // out
	var _cret *C.gchar // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_name(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// KeyvalToLower converts a key value to lower case, if applicable.
func KeyvalToLower(keyval uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_lower(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// KeyvalToUnicode: convert from a GDK key symbol to the corresponding ISO10646
// (Unicode) character.
func KeyvalToUnicode(keyval uint) uint32 {
	var _arg1 C.guint   // out
	var _cret C.guint32 // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_unicode(_arg1)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// KeyvalToUpper converts a key value to upper case, if applicable.
func KeyvalToUpper(keyval uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(keyval)

	_cret = C.gdk_keyval_to_upper(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// UnicodeToKeyval: convert from a ISO10646 character to a key symbol.
func UnicodeToKeyval(wc uint32) uint {
	var _arg1 C.guint32 // out
	var _cret C.guint   // in

	_arg1 = C.guint32(wc)

	_cret = C.gdk_unicode_to_keyval(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Keymap defines the translation from keyboard state (including a hardware key,
// a modifier mask, and active keyboard group) to a keyval. This translation has
// two phases. The first phase is to determine the effective keyboard group and
// level for the keyboard state; the second phase is to look up the
// keycode/group/level triplet in the keymap and see what keyval it corresponds
// to.
type Keymap interface {
	gextras.Objector

	// CapsLockState returns whether the Caps Lock modifer is locked.
	CapsLockState() bool
	// Direction returns the direction of effective layout of the keymap.
	Direction() pango.Direction
	// ModifierState returns the current modifier state.
	ModifierState() uint
	// NumLockState returns whether the Num Lock modifer is locked.
	NumLockState() bool
	// ScrollLockState returns whether the Scroll Lock modifer is locked.
	ScrollLockState() bool
	// HaveBidiLayouts determines if keyboard layouts for both right-to-left and
	// left-to-right languages are in use.
	HaveBidiLayouts() bool
	// LookupKey looks up the keyval mapped to a keycode/group/level triplet. If
	// no keyval is bound to @key, returns 0. For normal user input, you want to
	// use gdk_keymap_translate_keyboard_state() instead of this function, since
	// the effective group/level may not be the same as the current keyboard
	// state.
	LookupKey(key *KeymapKey) uint
}

// KeymapClass implements the Keymap interface.
type KeymapClass struct {
	*externglib.Object
}

var _ Keymap = (*KeymapClass)(nil)

func wrapKeymap(obj *externglib.Object) Keymap {
	return &KeymapClass{
		Object: obj,
	}
}

func marshalKeymap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapKeymap(obj), nil
}

// CapsLockState returns whether the Caps Lock modifer is locked.
func (k *KeymapClass) CapsLockState() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer((&k).Native()))

	_cret = C.gdk_keymap_get_caps_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Direction returns the direction of effective layout of the keymap.
func (k *KeymapClass) Direction() pango.Direction {
	var _arg0 *C.GdkKeymap     // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer((&k).Native()))

	_cret = C.gdk_keymap_get_direction(_arg0)

	var _direction pango.Direction // out

	_direction = (pango.Direction)(_cret)

	return _direction
}

// ModifierState returns the current modifier state.
func (k *KeymapClass) ModifierState() uint {
	var _arg0 *C.GdkKeymap // out
	var _cret C.guint      // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer((&k).Native()))

	_cret = C.gdk_keymap_get_modifier_state(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NumLockState returns whether the Num Lock modifer is locked.
func (k *KeymapClass) NumLockState() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer((&k).Native()))

	_cret = C.gdk_keymap_get_num_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ScrollLockState returns whether the Scroll Lock modifer is locked.
func (k *KeymapClass) ScrollLockState() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer((&k).Native()))

	_cret = C.gdk_keymap_get_scroll_lock_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HaveBidiLayouts determines if keyboard layouts for both right-to-left and
// left-to-right languages are in use.
func (k *KeymapClass) HaveBidiLayouts() bool {
	var _arg0 *C.GdkKeymap // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer((&k).Native()))

	_cret = C.gdk_keymap_have_bidi_layouts(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupKey looks up the keyval mapped to a keycode/group/level triplet. If no
// keyval is bound to @key, returns 0. For normal user input, you want to use
// gdk_keymap_translate_keyboard_state() instead of this function, since the
// effective group/level may not be the same as the current keyboard state.
func (k *KeymapClass) LookupKey(key *KeymapKey) uint {
	var _arg0 *C.GdkKeymap    // out
	var _arg1 *C.GdkKeymapKey // out
	var _cret C.guint         // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer((&k).Native()))
	_arg1 = (*C.GdkKeymapKey)(unsafe.Pointer(key))

	_cret = C.gdk_keymap_lookup_key(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// KeymapKey is a hardware key that can be mapped to a keyval.
type KeymapKey struct {
	native C.GdkKeymapKey
}

// WrapKeymapKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapKeymapKey(ptr unsafe.Pointer) *KeymapKey {
	return (*KeymapKey)(ptr)
}

// Native returns the underlying C source pointer.
func (k *KeymapKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&k.native)
}

// Keycode: the hardware keycode. This is an identifying number for a physical
// key.
func (k *KeymapKey) Keycode() uint {
	var v uint // out
	v = uint(k.native.keycode)
	return v
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) Group() int {
	var v int // out
	v = int(k.native.group)
	return v
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) Level() int {
	var v int // out
	v = int(k.native.level)
	return v
}
