// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_StyleProvider_ConnectGTKPrivateChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeStyleProvider = coreglib.Type(C.gtk_style_provider_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStyleProvider, F: marshalStyleProvider},
	})
}

// STYLE_PROVIDER_PRIORITY_APPLICATION: priority that can be used when adding a
// GtkStyleProvider for application-specific style information.
const STYLE_PROVIDER_PRIORITY_APPLICATION = 600

// STYLE_PROVIDER_PRIORITY_FALLBACK: priority used for default style information
// that is used in the absence of themes.
//
// Note that this is not very useful for providing default styling for custom
// style classes - themes are likely to override styling provided at this
// priority with catch-all * {...} rules.
const STYLE_PROVIDER_PRIORITY_FALLBACK = 1

// STYLE_PROVIDER_PRIORITY_SETTINGS: priority used for style information
// provided via GtkSettings.
//
// This priority is higher than K_STYLE_PROVIDER_PRIORITY_THEME to let settings
// override themes.
const STYLE_PROVIDER_PRIORITY_SETTINGS = 400

// STYLE_PROVIDER_PRIORITY_THEME: priority used for style information provided
// by themes.
const STYLE_PROVIDER_PRIORITY_THEME = 200

// STYLE_PROVIDER_PRIORITY_USER: priority used for the style information from
// $XDG_CONFIG_HOME/gtk-4.0/gtk.css.
//
// You should not use priorities higher than this, to give the user the last
// word.
const STYLE_PROVIDER_PRIORITY_USER = 800

// StyleProvider: GtkStyleProvider is an interface for style information used by
// GtkStyleContext.
//
// See gtk.StyleContext.AddProvider() and
// gtk.StyleContext().AddProviderForDisplay for adding GtkStyleProviders.
//
// GTK uses the GtkStyleProvider implementation for CSS in gtk.CSSProvider.
//
// StyleProvider wraps an interface. This means the user can get the
// underlying type by calling Cast().
type StyleProvider struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StyleProvider)(nil)
)

// StyleProviderer describes StyleProvider's interface methods.
type StyleProviderer interface {
	coreglib.Objector

	baseStyleProvider() *StyleProvider
}

var _ StyleProviderer = (*StyleProvider)(nil)

func wrapStyleProvider(obj *coreglib.Object) *StyleProvider {
	return &StyleProvider{
		Object: obj,
	}
}

func marshalStyleProvider(p uintptr) (interface{}, error) {
	return wrapStyleProvider(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *StyleProvider) baseStyleProvider() *StyleProvider {
	return v
}

// BaseStyleProvider returns the underlying base object.
func BaseStyleProvider(obj StyleProviderer) *StyleProvider {
	return obj.baseStyleProvider()
}

func (v *StyleProvider) ConnectGTKPrivateChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "gtk-private-changed", false, unsafe.Pointer(C._gotk4_gtk4_StyleProvider_ConnectGTKPrivateChanged), f)
}
