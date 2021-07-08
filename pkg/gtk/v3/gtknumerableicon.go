// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_numerable_icon_get_type()), F: marshalNumerableIcon},
	})
}

// NumerableIcon is a subclass of Icon that can show a number or short string as
// an emblem. The number can be overlayed on top of another emblem, if desired.
//
// It supports theming by taking font and color information from a provided
// StyleContext; see gtk_numerable_icon_set_style_context().
//
// Typical numerable icons: ! (numerableicon.png) ! (numerableicon2.png)
type NumerableIcon interface {
	gextras.Objector

	// BackgroundIconName returns the icon name used as the base background
	// image, or nil if there’s none.
	//
	// Deprecated: since version 3.14.
	BackgroundIconName() string
	// Count returns the value currently displayed by @self.
	//
	// Deprecated: since version 3.14.
	Count() int
	// Label returns the currently displayed label of the icon, or nil.
	//
	// Deprecated: since version 3.14.
	Label() string
	// StyleContext returns the StyleContext used by the icon for theming, or
	// nil if there’s none.
	//
	// Deprecated: since version 3.14.
	StyleContext() StyleContext
	// SetBackgroundIconName updates the icon to use the icon named @icon_name
	// from the current icon theme as the base background image. If @icon_name
	// is nil, @self will go back using style information or default theming for
	// its background image.
	//
	// If this method is called and a #GIcon was already set as background for
	// the icon, @icon_name will be used, i.e. the last method called between
	// gtk_numerable_icon_set_background_icon_name() and
	// gtk_numerable_icon_set_background_gicon() has always priority.
	//
	// Deprecated: since version 3.14.
	SetBackgroundIconName(iconName string)
	// SetCount sets the currently displayed value of @self to @count.
	//
	// The numeric value is always clamped to make it two digits, i.e. between
	// -99 and 99. Setting a count of zero removes the emblem. If this method is
	// called, and a label was already set on the icon, it will automatically be
	// reset to nil before rendering the number, i.e. the last method called
	// between gtk_numerable_icon_set_count() and gtk_numerable_icon_set_label()
	// has always priority.
	//
	// Deprecated: since version 3.14.
	SetCount(count int)
	// SetLabel sets the currently displayed value of @self to the string in
	// @label. Setting an empty label removes the emblem.
	//
	// Note that this is meant for displaying short labels, such as roman
	// numbers, or single letters. For roman numbers, consider using the Unicode
	// characters U+2160 - U+217F. Strings longer than two characters will
	// likely not be rendered very well.
	//
	// If this method is called, and a number was already set on the icon, it
	// will automatically be reset to zero before rendering the label, i.e. the
	// last method called between gtk_numerable_icon_set_label() and
	// gtk_numerable_icon_set_count() has always priority.
	//
	// Deprecated: since version 3.14.
	SetLabel(label string)
	// SetStyleContext updates the icon to fetch theme information from the
	// given StyleContext.
	//
	// Deprecated: since version 3.14.
	SetStyleContext(style StyleContext)
}

// NumerableIconClass implements the NumerableIcon interface.
type NumerableIconClass struct {
	gio.EmblemedIconClass
}

var _ NumerableIcon = (*NumerableIconClass)(nil)

func wrapNumerableIcon(obj *externglib.Object) NumerableIcon {
	return &NumerableIconClass{
		EmblemedIconClass: gio.EmblemedIconClass{
			Object: obj,
		},
	}
}

func marshalNumerableIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNumerableIcon(obj), nil
}

// BackgroundIconName returns the icon name used as the base background image,
// or nil if there’s none.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) BackgroundIconName() string {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_numerable_icon_get_background_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Count returns the value currently displayed by @self.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) Count() int {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_numerable_icon_get_count(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Label returns the currently displayed label of the icon, or nil.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) Label() string {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_numerable_icon_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// StyleContext returns the StyleContext used by the icon for theming, or nil if
// there’s none.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) StyleContext() StyleContext {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret *C.GtkStyleContext  // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_numerable_icon_get_style_context(_arg0)

	var _styleContext StyleContext // out

	_styleContext = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(StyleContext)

	return _styleContext
}

// SetBackgroundIconName updates the icon to use the icon named @icon_name from
// the current icon theme as the base background image. If @icon_name is nil,
// @self will go back using style information or default theming for its
// background image.
//
// If this method is called and a #GIcon was already set as background for the
// icon, @icon_name will be used, i.e. the last method called between
// gtk_numerable_icon_set_background_icon_name() and
// gtk_numerable_icon_set_background_gicon() has always priority.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) SetBackgroundIconName(iconName string) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_numerable_icon_set_background_icon_name(_arg0, _arg1)
}

// SetCount sets the currently displayed value of @self to @count.
//
// The numeric value is always clamped to make it two digits, i.e. between -99
// and 99. Setting a count of zero removes the emblem. If this method is called,
// and a label was already set on the icon, it will automatically be reset to
// nil before rendering the number, i.e. the last method called between
// gtk_numerable_icon_set_count() and gtk_numerable_icon_set_label() has always
// priority.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) SetCount(count int) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 C.gint              // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(count)

	C.gtk_numerable_icon_set_count(_arg0, _arg1)
}

// SetLabel sets the currently displayed value of @self to the string in @label.
// Setting an empty label removes the emblem.
//
// Note that this is meant for displaying short labels, such as roman numbers,
// or single letters. For roman numbers, consider using the Unicode characters
// U+2160 - U+217F. Strings longer than two characters will likely not be
// rendered very well.
//
// If this method is called, and a number was already set on the icon, it will
// automatically be reset to zero before rendering the label, i.e. the last
// method called between gtk_numerable_icon_set_label() and
// gtk_numerable_icon_set_count() has always priority.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) SetLabel(label string) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_numerable_icon_set_label(_arg0, _arg1)
}

// SetStyleContext updates the icon to fetch theme information from the given
// StyleContext.
//
// Deprecated: since version 3.14.
func (s *NumerableIconClass) SetStyleContext(style StyleContext) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 *C.GtkStyleContext  // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(style.Native()))

	C.gtk_numerable_icon_set_style_context(_arg0, _arg1)
}
