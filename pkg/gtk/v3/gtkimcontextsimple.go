// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeIMContextSimple = coreglib.Type(C.gtk_im_context_simple_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeIMContextSimple, F: marshalIMContextSimple},
	})
}

// MAX_COMPOSE_LEN: maximum length of sequences in compose tables.
const MAX_COMPOSE_LEN = 7

// IMContextSimpleOverrides contains methods that are overridable.
type IMContextSimpleOverrides struct {
}

func defaultIMContextSimpleOverrides(v *IMContextSimple) IMContextSimpleOverrides {
	return IMContextSimpleOverrides{}
}

// IMContextSimple is a simple input method context supporting table-based input
// methods. It has a built-in table of compose sequences that is derived from
// the X11 Compose files.
//
// GtkIMContextSimple reads additional compose sequences from the first of the
// following files that is found: ~/.config/gtk-3.0/Compose, ~/.XCompose,
// /usr/share/X11/locale/$locale/Compose (for locales that have a nontrivial
// Compose file). The syntax of these files is described in the Compose(5)
// manual page.
//
//
// Unicode characters
//
// GtkIMContextSimple also supports numeric entry of Unicode characters by
// typing Ctrl-Shift-u, followed by a hexadecimal Unicode codepoint. For
// example, Ctrl-Shift-u 1 2 3 Enter yields U+0123 LATIN SMALL LETTER G WITH
// CEDILLA, i.e. ģ.
type IMContextSimple struct {
	_ [0]func() // equal guard
	IMContext
}

var (
	_ IMContexter = (*IMContextSimple)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*IMContextSimple, *IMContextSimpleClass, IMContextSimpleOverrides](
		GTypeIMContextSimple,
		initIMContextSimpleClass,
		wrapIMContextSimple,
		defaultIMContextSimpleOverrides,
	)
}

func initIMContextSimpleClass(gclass unsafe.Pointer, overrides IMContextSimpleOverrides, classInitFunc func(*IMContextSimpleClass)) {
	if classInitFunc != nil {
		class := (*IMContextSimpleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapIMContextSimple(obj *coreglib.Object) *IMContextSimple {
	return &IMContextSimple{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMContextSimple(p uintptr) (interface{}, error) {
	return wrapIMContextSimple(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIMContextSimple creates a new IMContextSimple.
//
// The function returns the following values:
//
//    - imContextSimple: new IMContextSimple.
//
func NewIMContextSimple() *IMContextSimple {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_im_context_simple_new()

	var _imContextSimple *IMContextSimple // out

	_imContextSimple = wrapIMContextSimple(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imContextSimple
}

// AddComposeFile adds an additional table from the X11 compose file.
//
// The function takes the following parameters:
//
//    - composeFile: path of compose file.
//
func (contextSimple *IMContextSimple) AddComposeFile(composeFile string) {
	var _arg0 *C.GtkIMContextSimple // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.GtkIMContextSimple)(unsafe.Pointer(coreglib.InternObject(contextSimple).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(composeFile)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_context_simple_add_compose_file(_arg0, _arg1)
	runtime.KeepAlive(contextSimple)
	runtime.KeepAlive(composeFile)
}

// IMContextSimpleClass: instance of this type is always passed by reference.
type IMContextSimpleClass struct {
	*imContextSimpleClass
}

// imContextSimpleClass is the struct that's finalized.
type imContextSimpleClass struct {
	native *C.GtkIMContextSimpleClass
}

func (i *IMContextSimpleClass) ParentClass() *IMContextClass {
	valptr := &i.native.parent_class
	var _v *IMContextClass // out
	_v = (*IMContextClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
