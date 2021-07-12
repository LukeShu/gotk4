// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_style_provider_get_type()), F: marshalStyleProviderer},
	})
}

// StyleProviderer describes StyleProvider's methods.
type StyleProviderer interface {
	privateStyleProvider()
}

// StyleProvider: `GtkStyleProvider` is an interface for style information used
// by `GtkStyleContext`.
//
// See [method@Gtk.StyleContext.add_provider] and
// [func@Gtk.StyleContext.add_provider_for_display] for adding
// `GtkStyleProviders`.
//
// GTK uses the `GtkStyleProvider` implementation for CSS in
// [iface@Gtk.CssProvider].
type StyleProvider struct {
	*externglib.Object
}

var (
	_ StyleProviderer = (*StyleProvider)(nil)
	_ gextras.Nativer = (*StyleProvider)(nil)
)

func wrapStyleProvider(obj *externglib.Object) StyleProviderer {
	return &StyleProvider{
		Object: obj,
	}
}

func marshalStyleProviderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStyleProvider(obj), nil
}

func (*StyleProvider) privateStyleProvider() {}
