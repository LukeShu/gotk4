// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_section_get_type()), F: marshalCSSSection},
	})
}

// CSSSection defines a part of a CSS document.
//
// Because sections are nested into one another, you can use
// gtk_css_section_get_parent() to get the containing region.
type CSSSection struct {
	nocopy gextras.NoCopy
	native *C.GtkCssSection
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &CSSSection{native: (*C.GtkCssSection)(unsafe.Pointer(b))}, nil
}

// NewCSSSection constructs a struct CSSSection.
func NewCSSSection(file gio.Filer, start *CSSLocation, end *CSSLocation) *CSSSection {
	var _arg1 *C.GFile          // out
	var _arg2 *C.GtkCssLocation // out
	var _arg3 *C.GtkCssLocation // out
	var _cret *C.GtkCssSection  // in

	if file != nil {
		_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))
	}
	_arg2 = (*C.GtkCssLocation)(gextras.StructNative(unsafe.Pointer(start)))
	_arg3 = (*C.GtkCssLocation)(gextras.StructNative(unsafe.Pointer(end)))

	_cret = C.gtk_css_section_new(_arg1, _arg2, _arg3)

	runtime.KeepAlive(file)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.gtk_css_section_unref((*C.GtkCssSection)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _cssSection
}

// EndLocation returns the location in the CSS document where this section ends.
func (section *CSSSection) EndLocation() *CSSLocation {
	var _arg0 *C.GtkCssSection  // out
	var _cret *C.GtkCssLocation // in

	_arg0 = (*C.GtkCssSection)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = C.gtk_css_section_get_end_location(_arg0)

	runtime.KeepAlive(section)

	var _cssLocation *CSSLocation // out

	_cssLocation = (*CSSLocation)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _cssLocation
}

// File gets the file that section was parsed from.
//
// If no such file exists, for example because the CSS was loaded via
// gtk.CSSProvider.LoadFromData(), then NULL is returned.
func (section *CSSSection) File() gio.Filer {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GFile         // in

	_arg0 = (*C.GtkCssSection)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = C.gtk_css_section_get_file(_arg0)

	runtime.KeepAlive(section)

	var _file gio.Filer // out

	_file = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.Filer)

	return _file
}

// Parent gets the parent section for the given section.
//
// The parent section is the section that contains this section. A special case
// are sections of type GTK_CSS_SECTION_DOCUMENT. Their parent will either be
// NULL if they are the original CSS document that was loaded by
// gtk.CSSProvider.LoadFromFile() or a section of type GTK_CSS_SECTION_IMPORT if
// it was loaded with an import rule from a different file.
func (section *CSSSection) Parent() *CSSSection {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GtkCssSection // in

	_arg0 = (*C.GtkCssSection)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = C.gtk_css_section_get_parent(_arg0)

	runtime.KeepAlive(section)

	var _cssSection *CSSSection // out

	if _cret != nil {
		_cssSection = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gtk_css_section_ref(_cret)
		runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
			C.gtk_css_section_unref((*C.GtkCssSection)(gextras.StructNative(unsafe.Pointer(v))))
		})
	}

	return _cssSection
}

// StartLocation returns the location in the CSS document where this section
// starts.
func (section *CSSSection) StartLocation() *CSSLocation {
	var _arg0 *C.GtkCssSection  // out
	var _cret *C.GtkCssLocation // in

	_arg0 = (*C.GtkCssSection)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = C.gtk_css_section_get_start_location(_arg0)

	runtime.KeepAlive(section)

	var _cssLocation *CSSLocation // out

	_cssLocation = (*CSSLocation)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _cssLocation
}

// String prints the section into a human-readable text form using
// gtk.CSSSection.Print().
func (section *CSSSection) String() string {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkCssSection)(gextras.StructNative(unsafe.Pointer(section)))

	_cret = C.gtk_css_section_to_string(_arg0)

	runtime.KeepAlive(section)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
