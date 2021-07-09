// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_style_provider_get_type()), F: marshalStyleProvider},
	})
}

// StyleProviderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type StyleProviderOverrider interface {
	// IconFactory returns the IconFactory defined to be in use for @path, or
	// nil if none is defined.
	//
	// Deprecated: since version 3.8.
	IconFactory(path *WidgetPath) *IconFactoryClass
	// Style returns the style settings affecting a widget defined by @path, or
	// nil if @provider doesn’t contemplate styling @path.
	//
	// Deprecated: since version 3.8.
	Style(path *WidgetPath) *StylePropertiesClass
}

// StyleProvider is an interface used to provide style information to a
// StyleContext. See gtk_style_context_add_provider() and
// gtk_style_context_add_provider_for_screen().
type StyleProvider interface {
	gextras.Objector

	// IconFactory returns the IconFactory defined to be in use for @path, or
	// nil if none is defined.
	//
	// Deprecated: since version 3.8.
	IconFactory(path *WidgetPath) *IconFactoryClass
	// Style returns the style settings affecting a widget defined by @path, or
	// nil if @provider doesn’t contemplate styling @path.
	//
	// Deprecated: since version 3.8.
	Style(path *WidgetPath) *StylePropertiesClass
}

// StyleProviderInterface implements the StyleProvider interface.
type StyleProviderInterface struct {
	*externglib.Object
}

var _ StyleProvider = (*StyleProviderInterface)(nil)

func wrapStyleProvider(obj *externglib.Object) StyleProvider {
	return &StyleProviderInterface{
		Object: obj,
	}
}

func marshalStyleProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStyleProvider(obj), nil
}

// IconFactory returns the IconFactory defined to be in use for @path, or nil if
// none is defined.
//
// Deprecated: since version 3.8.
func (p *StyleProviderInterface) IconFactory(path *WidgetPath) *IconFactoryClass {
	var _arg0 *C.GtkStyleProvider // out
	var _arg1 *C.GtkWidgetPath    // out
	var _cret *C.GtkIconFactory   // in

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer((&p).Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path))

	_cret = C.gtk_style_provider_get_icon_factory(_arg0, _arg1)

	var _iconFactory *IconFactoryClass // out

	_iconFactory = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*IconFactoryClass)

	return _iconFactory
}

// Style returns the style settings affecting a widget defined by @path, or nil
// if @provider doesn’t contemplate styling @path.
//
// Deprecated: since version 3.8.
func (p *StyleProviderInterface) Style(path *WidgetPath) *StylePropertiesClass {
	var _arg0 *C.GtkStyleProvider   // out
	var _arg1 *C.GtkWidgetPath      // out
	var _cret *C.GtkStyleProperties // in

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer((&p).Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path))

	_cret = C.gtk_style_provider_get_style(_arg0, _arg1)

	var _styleProperties *StylePropertiesClass // out

	_styleProperties = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*StylePropertiesClass)

	return _styleProperties
}
