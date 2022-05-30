// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkcsssection.go.
var (
	GTypeCSSSectionType = coreglib.Type(C.gtk_css_section_type_get_type())
	GTypeCSSSection     = coreglib.Type(C.gtk_css_section_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCSSSectionType, F: marshalCSSSectionType},
		{T: GTypeCSSSection, F: marshalCSSSection},
	})
}

// CSSSectionType: different types of sections indicate parts of a CSS document
// as parsed by GTK’s CSS parser. They are oriented towards the CSS Grammar
// (http://www.w3.org/TR/CSS21/grammar.html), but may contain extensions.
//
// More types might be added in the future as the parser incorporates more
// features.
type CSSSectionType C.gint

const (
	// CSSSectionDocument: section describes a complete document. This section
	// time is the only one where gtk_css_section_get_parent() might return
	// NULL.
	CSSSectionDocument CSSSectionType = iota
	// CSSSectionImport: section defines an import rule.
	CSSSectionImport
	// CSSSectionColorDefinition: section defines a color. This is a GTK
	// extension to CSS.
	CSSSectionColorDefinition
	// CSSSectionBindingSet: section defines a binding set. This is a GTK
	// extension to CSS.
	CSSSectionBindingSet
	// CSSSectionRuleset: section defines a CSS ruleset.
	CSSSectionRuleset
	// CSSSectionSelector: section defines a CSS selector.
	CSSSectionSelector
	// CSSSectionDeclaration: section defines the declaration of a CSS variable.
	CSSSectionDeclaration
	// CSSSectionValue: section defines the value of a CSS declaration.
	CSSSectionValue
	// CSSSectionKeyframes: section defines keyframes. See [CSS
	// Animations](http://dev.w3.org/csswg/css3-animations/#keyframes) for
	// details. Since 3.6.
	CSSSectionKeyframes
)

func marshalCSSSectionType(p uintptr) (interface{}, error) {
	return CSSSectionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CSSSectionType.
func (c CSSSectionType) String() string {
	switch c {
	case CSSSectionDocument:
		return "Document"
	case CSSSectionImport:
		return "Import"
	case CSSSectionColorDefinition:
		return "ColorDefinition"
	case CSSSectionBindingSet:
		return "BindingSet"
	case CSSSectionRuleset:
		return "Ruleset"
	case CSSSectionSelector:
		return "Selector"
	case CSSSectionDeclaration:
		return "Declaration"
	case CSSSectionValue:
		return "Value"
	case CSSSectionKeyframes:
		return "Keyframes"
	default:
		return fmt.Sprintf("CSSSectionType(%d)", c)
	}
}

// CSSSection defines a part of a CSS document. Because sections are nested into
// one another, you can use gtk_css_section_get_parent() to get the containing
// region.
//
// An instance of this type is always passed by reference.
type CSSSection struct {
	*cssSection
}

// cssSection is the struct that's finalized.
type cssSection struct {
	native *C.GtkCssSection
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &CSSSection{&cssSection{(*C.GtkCssSection)(b)}}, nil
}

// EndLine returns the line in the CSS document where this section end. The line
// number is 0-indexed, so the first line of the document will return 0. This
// value may change in future invocations of this function if section is not yet
// parsed completely. This will for example happen in the
// GtkCssProvider::parsing-error signal. The end position and line may be
// identical to the start position and line for sections which failed to parse
// anything successfully.
//
// The function returns the following values:
//
//    - guint: line number.
//
func (section *CSSSection) EndLine() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))
	*(**CSSSection)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EndPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_end_line(). This value may change in future
// invocations of this function if section is not yet parsed completely. This
// will for example happen in the GtkCssProvider::parsing-error signal. The end
// position and line may be identical to the start position and line for
// sections which failed to parse anything successfully.
//
// The function returns the following values:
//
//    - guint: offset in bytes from the start of the line.
//
func (section *CSSSection) EndPosition() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))
	*(**CSSSection)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// File gets the file that section was parsed from. If no such file exists, for
// example because the CSS was loaded via gtk_css_provider_load_from_data(),
// then NULL is returned.
//
// The function returns the following values:
//
//    - file that section was parsed from or NULL if section was parsed from
//      other data.
//
func (section *CSSSection) File() *gio.File {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))
	*(**CSSSection)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _file *gio.File // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_file = &gio.File{
			Object: obj,
		}
	}

	return _file
}

// Parent gets the parent section for the given section. The parent section is
// the section that contains this section. A special case are sections of type
// K_CSS_SECTION_DOCUMENT. Their parent will either be NULL if they are the
// original CSS document that was loaded by gtk_css_provider_load_from_file() or
// a section of type K_CSS_SECTION_IMPORT if it was loaded with an import rule
// from a different file.
//
// The function returns the following values:
//
//    - cssSection (optional): parent section or NULL if none.
//
func (section *CSSSection) Parent() *CSSSection {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))
	*(**CSSSection)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _cssSection *CSSSection // out

	if _cret != nil {
		_cssSection = (*CSSSection)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gtk_css_section_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_cssSection)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_css_section_unref((*C.GtkCssSection)(intern.C))
			},
		)
	}

	return _cssSection
}

// StartLine returns the line in the CSS document where this section starts. The
// line number is 0-indexed, so the first line of the document will return 0.
//
// The function returns the following values:
//
//    - guint: line number.
//
func (section *CSSSection) StartLine() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))
	*(**CSSSection)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// StartPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_start_line().
//
// The function returns the following values:
//
//    - guint: offset in bytes from the start of the line.
//
func (section *CSSSection) StartPosition() uint {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.guint // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(section)))
	*(**CSSSection)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(section)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}
