// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GTypeNumerableIcon returns the GType for the type NumerableIcon.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeNumerableIcon() coreglib.Type {
	gtype := coreglib.Type(C.gtk_numerable_icon_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalNumerableIcon)
	return gtype
}

// NumerableIconOverrider contains methods that are overridable.
type NumerableIconOverrider interface {
}

// NumerableIcon is a subclass of Icon that can show a number or short string as
// an emblem. The number can be overlayed on top of another emblem, if desired.
//
// It supports theming by taking font and color information from a provided
// StyleContext; see gtk_numerable_icon_set_style_context().
//
// Typical numerable icons: ! (numerableicon.png) ! (numerableicon2.png).
type NumerableIcon struct {
	_ [0]func() // equal guard
	gio.EmblemedIcon
}

var (
	_ coreglib.Objector = (*NumerableIcon)(nil)
)

func classInitNumerableIconner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNumerableIcon(obj *coreglib.Object) *NumerableIcon {
	return &NumerableIcon{
		EmblemedIcon: gio.EmblemedIcon{
			Object: obj,
			Icon: gio.Icon{
				Object: obj,
			},
		},
	}
}

func marshalNumerableIcon(p uintptr) (interface{}, error) {
	return wrapNumerableIcon(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BackgroundGIcon returns the #GIcon that was set as the base background image,
// or NULL if there’s none. The caller of this function does not own a reference
// to the returned #GIcon.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - icon (optional) or NULL.
//
func (self *NumerableIcon) BackgroundGIcon() *gio.Icon {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret *C.GIcon            // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_numerable_icon_get_background_gicon(_arg0)
	runtime.KeepAlive(self)

	var _icon *gio.Icon // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_icon = &gio.Icon{
				Object: obj,
			}
		}
	}

	return _icon
}

// BackgroundIconName returns the icon name used as the base background image,
// or NULL if there’s none.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - utf8 (optional): icon name, or NULL.
//
func (self *NumerableIcon) BackgroundIconName() string {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_numerable_icon_get_background_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Count returns the value currently displayed by self.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - gint: currently displayed value.
//
func (self *NumerableIcon) Count() int32 {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_numerable_icon_get_count(_arg0)
	runtime.KeepAlive(self)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Label returns the currently displayed label of the icon, or NULL.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - utf8 (optional): currently displayed label.
//
func (self *NumerableIcon) Label() string {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_numerable_icon_get_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// StyleContext returns the StyleContext used by the icon for theming, or NULL
// if there’s none.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - styleContext (optional) or NULL. This object is internal to GTK+ and
//      should not be unreffed. Use g_object_ref() if you want to keep it around.
//
func (self *NumerableIcon) StyleContext() *StyleContext {
	var _arg0 *C.GtkNumerableIcon // out
	var _cret *C.GtkStyleContext  // in

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_numerable_icon_get_style_context(_arg0)
	runtime.KeepAlive(self)

	var _styleContext *StyleContext // out

	if _cret != nil {
		_styleContext = wrapStyleContext(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _styleContext
}

// SetBackgroundGIcon updates the icon to use icon as the base background image.
// If icon is NULL, self will go back using style information or default theming
// for its background image.
//
// If this method is called and an icon name was already set as background for
// the icon, icon will be used, i.e. the last method called between
// gtk_numerable_icon_set_background_gicon() and
// gtk_numerable_icon_set_background_icon_name() has always priority.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - icon (optional) or NULL.
//
func (self *NumerableIcon) SetBackgroundGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 *C.GIcon            // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if icon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	}

	C.gtk_numerable_icon_set_background_gicon(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(icon)
}

// SetBackgroundIconName updates the icon to use the icon named icon_name from
// the current icon theme as the base background image. If icon_name is NULL,
// self will go back using style information or default theming for its
// background image.
//
// If this method is called and a #GIcon was already set as background for the
// icon, icon_name will be used, i.e. the last method called between
// gtk_numerable_icon_set_background_icon_name() and
// gtk_numerable_icon_set_background_gicon() has always priority.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name, or NULL.
//
func (self *NumerableIcon) SetBackgroundIconName(iconName string) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_numerable_icon_set_background_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetCount sets the currently displayed value of self to count.
//
// The numeric value is always clamped to make it two digits, i.e. between -99
// and 99. Setting a count of zero removes the emblem. If this method is called,
// and a label was already set on the icon, it will automatically be reset to
// NULL before rendering the number, i.e. the last method called between
// gtk_numerable_icon_set_count() and gtk_numerable_icon_set_label() has always
// priority.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - count: number between -99 and 99.
//
func (self *NumerableIcon) SetCount(count int32) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 C.gint              // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gint(count)

	C.gtk_numerable_icon_set_count(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(count)
}

// SetLabel sets the currently displayed value of self to the string in label.
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
//
// The function takes the following parameters:
//
//    - label (optional): short label, or NULL.
//
func (self *NumerableIcon) SetLabel(label string) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_numerable_icon_set_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(label)
}

// SetStyleContext updates the icon to fetch theme information from the given
// StyleContext.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - style: StyleContext.
//
func (self *NumerableIcon) SetStyleContext(style *StyleContext) {
	var _arg0 *C.GtkNumerableIcon // out
	var _arg1 *C.GtkStyleContext  // out

	_arg0 = (*C.GtkNumerableIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(style).Native()))

	C.gtk_numerable_icon_set_style_context(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(style)
}

// NewNumerableIcon creates a new unthemed NumerableIcon.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - baseIcon to overlay on.
//
// The function returns the following values:
//
//    - icon: new #GIcon.
//
func NewNumerableIcon(baseIcon gio.Iconner) *gio.Icon {
	var _arg1 *C.GIcon // out
	var _cret *C.GIcon // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(baseIcon).Native()))

	_cret = C.gtk_numerable_icon_new(_arg1)
	runtime.KeepAlive(baseIcon)

	var _icon *gio.Icon // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_icon = &gio.Icon{
			Object: obj,
		}
	}

	return _icon
}

// NewNumerableIconWithStyleContext creates a new NumerableIcon which will
// themed according to the passed StyleContext. This is a convenience
// constructor that calls gtk_numerable_icon_set_style_context() internally.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - baseIcon to overlay on.
//    - context: StyleContext.
//
// The function returns the following values:
//
//    - icon: new #GIcon.
//
func NewNumerableIconWithStyleContext(baseIcon gio.Iconner, context *StyleContext) *gio.Icon {
	var _arg1 *C.GIcon           // out
	var _arg2 *C.GtkStyleContext // out
	var _cret *C.GIcon           // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(baseIcon).Native()))
	_arg2 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gtk_numerable_icon_new_with_style_context(_arg1, _arg2)
	runtime.KeepAlive(baseIcon)
	runtime.KeepAlive(context)

	var _icon *gio.Icon // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_icon = &gio.Icon{
			Object: obj,
		}
	}

	return _icon
}
