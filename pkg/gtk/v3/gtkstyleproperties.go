// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkstyleproperties.go.
var (
	GTypeStyleProperties = coreglib.Type(C.gtk_style_properties_get_type())
	GTypeGradient        = coreglib.Type(C.gtk_gradient_get_type())
	GTypeSymbolicColor   = coreglib.Type(C.gtk_symbolic_color_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeStyleProperties, F: marshalStyleProperties},
		{T: GTypeGradient, F: marshalGradient},
		{T: GTypeSymbolicColor, F: marshalSymbolicColor},
	})
}

// StylePropertiesOverrider contains methods that are overridable.
type StylePropertiesOverrider interface {
}

// StyleProperties provides the storage for style information that is used by
// StyleContext and other StyleProvider implementations.
//
// Before style properties can be stored in GtkStyleProperties, they must be
// registered with gtk_style_properties_register_property().
//
// Unless you are writing a StyleProvider implementation, you are unlikely to
// use this API directly, as gtk_style_context_get() and its variants are the
// preferred way to access styling information from widget implementations and
// theming engine implementations should use the APIs provided by ThemingEngine
// instead.
//
// StyleProperties has been deprecated in GTK 3.16. The CSS machinery does not
// use it anymore and all users of this object have been deprecated.
type StyleProperties struct {
	_ [0]func() // equal guard
	*coreglib.Object

	StyleProvider
}

var (
	_ coreglib.Objector = (*StyleProperties)(nil)
)

func classInitStylePropertieser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapStyleProperties(obj *coreglib.Object) *StyleProperties {
	return &StyleProperties{
		Object: obj,
		StyleProvider: StyleProvider{
			Object: obj,
		},
	}
}

func marshalStyleProperties(p uintptr) (interface{}, error) {
	return wrapStyleProperties(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStyleProperties returns a newly created StyleProperties
//
// Deprecated: StyleProperties are deprecated.
//
// The function returns the following values:
//
//    - styleProperties: new StyleProperties.
//
func NewStyleProperties() *StyleProperties {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "StyleProperties").InvokeMethod("new_StyleProperties", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _styleProperties *StyleProperties // out

	_styleProperties = wrapStyleProperties(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _styleProperties
}

// Clear clears all style information from props.
//
// Deprecated: StyleProperties are deprecated.
func (props *StyleProperties) Clear() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(props).Native()))
	*(**StyleProperties)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "StyleProperties").InvokeMethod("clear", args[:], nil)

	runtime.KeepAlive(props)
}

// LookupColor returns the symbolic color that is mapped to name.
//
// Deprecated: SymbolicColor is deprecated.
//
// The function takes the following parameters:
//
//    - name: color name to lookup.
//
// The function returns the following values:
//
//    - symbolicColor: mapped color.
//
func (props *StyleProperties) LookupColor(name string) *SymbolicColor {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(props).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**StyleProperties)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "StyleProperties").InvokeMethod("lookup_color", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(props)
	runtime.KeepAlive(name)

	var _symbolicColor *SymbolicColor // out

	_symbolicColor = (*SymbolicColor)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_symbolic_color_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_symbolicColor)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(intern.C))
		},
	)

	return _symbolicColor
}

// MapColor maps color so it can be referenced by name. See
// gtk_style_properties_lookup_color()
//
// Deprecated: SymbolicColor is deprecated.
//
// The function takes the following parameters:
//
//    - name: color name.
//    - color to map name to.
//
func (props *StyleProperties) MapColor(name string, color *SymbolicColor) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(props).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))
	*(**StyleProperties)(unsafe.Pointer(&args[1])) = _arg1
	*(*string)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "StyleProperties").InvokeMethod("map_color", args[:], nil)

	runtime.KeepAlive(props)
	runtime.KeepAlive(name)
	runtime.KeepAlive(color)
}

// Merge merges into props all the style information contained in
// props_to_merge. If replace is TRUE, the values will be overwritten, if it is
// FALSE, the older values will prevail.
//
// Deprecated: StyleProperties are deprecated.
//
// The function takes the following parameters:
//
//    - propsToMerge: second StyleProperties.
//    - replace: whether to replace values or not.
//
func (props *StyleProperties) Merge(propsToMerge *StyleProperties, replace bool) {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(props).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(propsToMerge).Native()))
	if replace {
		_arg2 = C.TRUE
	}
	*(**StyleProperties)(unsafe.Pointer(&args[1])) = _arg1
	*(**StyleProperties)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "StyleProperties").InvokeMethod("merge", args[:], nil)

	runtime.KeepAlive(props)
	runtime.KeepAlive(propsToMerge)
	runtime.KeepAlive(replace)
}

// Gradient is a boxed type that represents a gradient. It is the result of
// parsing a [gradient expression][gtkcssprovider-gradients]. To obtain the
// gradient represented by a GtkGradient, it has to be resolved with
// gtk_gradient_resolve(), which replaces all symbolic color references by the
// colors they refer to (in a given context) and constructs a #cairo_pattern_t
// value.
//
// It is not normally necessary to deal directly with Gradients, since they are
// mostly used behind the scenes by StyleContext and CssProvider.
//
// Gradient is deprecated. It was used internally by GTK’s CSS engine to
// represent gradients. As its handling is not conforming to modern web
// standards, it is not used anymore. If you want to use gradients in your own
// code, please use Cairo directly.
//
// An instance of this type is always passed by reference.
type Gradient struct {
	*gradient
}

// gradient is the struct that's finalized.
type gradient struct {
	native *C.GtkGradient
}

func marshalGradient(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Gradient{&gradient{(*C.GtkGradient)(b)}}, nil
}

// NewGradientLinear constructs a struct Gradient.
func NewGradientLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Gradient {
	var args [4]girepository.Argument
	var _arg0 C.gdouble // out
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out
	var _arg3 C.gdouble // out
	var _cret *C.void   // in

	_arg0 = C.gdouble(x0)
	_arg1 = C.gdouble(y0)
	_arg2 = C.gdouble(x1)
	_arg3 = C.gdouble(y1)
	*(*float64)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2
	*(*float64)(unsafe.Pointer(&args[3])) = _arg3

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(x0)
	runtime.KeepAlive(y0)
	runtime.KeepAlive(x1)
	runtime.KeepAlive(y1)

	var _gradient *Gradient // out

	_gradient = (*Gradient)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_gradient)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_gradient_unref((*C.GtkGradient)(intern.C))
		},
	)

	return _gradient
}

// NewGradientRadial constructs a struct Gradient.
func NewGradientRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) *Gradient {
	var args [6]girepository.Argument
	var _arg0 C.gdouble // out
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out
	var _arg3 C.gdouble // out
	var _arg4 C.gdouble // out
	var _arg5 C.gdouble // out
	var _cret *C.void   // in

	_arg0 = C.gdouble(x0)
	_arg1 = C.gdouble(y0)
	_arg2 = C.gdouble(radius0)
	_arg3 = C.gdouble(x1)
	_arg4 = C.gdouble(y1)
	_arg5 = C.gdouble(radius1)
	*(*float64)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2
	*(*float64)(unsafe.Pointer(&args[3])) = _arg3
	*(*float64)(unsafe.Pointer(&args[4])) = _arg4
	*(*float64)(unsafe.Pointer(&args[5])) = _arg5

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(x0)
	runtime.KeepAlive(y0)
	runtime.KeepAlive(radius0)
	runtime.KeepAlive(x1)
	runtime.KeepAlive(y1)
	runtime.KeepAlive(radius1)

	var _gradient *Gradient // out

	_gradient = (*Gradient)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_gradient)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_gradient_unref((*C.GtkGradient)(intern.C))
		},
	)

	return _gradient
}

// AddColorStop adds a stop color to gradient.
//
// Deprecated: Gradient is deprecated.
//
// The function takes the following parameters:
//
//    - offset for the color stop.
//    - color to use.
//
func (gradient *Gradient) AddColorStop(offset float64, color *SymbolicColor) {
	var args [3]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.gdouble // out
	var _arg2 *C.void   // out

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(gradient)))
	_arg1 = C.gdouble(offset)
	_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))
	*(**Gradient)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2

	runtime.KeepAlive(gradient)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(color)
}

// Resolve: if gradient is resolvable, resolved_gradient will be filled in with
// the resolved gradient as a cairo_pattern_t, and TRUE will be returned.
// Generally, if gradient can’t be resolved, it is due to it being defined on
// top of a named color that doesn't exist in props.
//
// Deprecated: Gradient is deprecated.
//
// The function takes the following parameters:
//
//    - props to use when resolving named colors.
//
// The function returns the following values:
//
//    - resolvedGradient: return location for the resolved pattern.
//    - ok: TRUE if the gradient has been resolved.
//
func (gradient *Gradient) Resolve(props *StyleProperties) (*cairo.Pattern, bool) {
	var args [2]girepository.Argument
	var outs [1]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _out0 *C.void    // in
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(gradient)))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(props).Native()))
	*(**Gradient)(unsafe.Pointer(&args[1])) = _arg1
	*(**StyleProperties)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gradient)
	runtime.KeepAlive(props)

	var _resolvedGradient *cairo.Pattern // out
	var _ok bool                         // out

	{
		_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_out0)}
		_resolvedGradient = (*cairo.Pattern)(unsafe.Pointer(_pp))
	}
	runtime.SetFinalizer(_resolvedGradient, func(v *cairo.Pattern) {
		C.cairo_pattern_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	if _cret != 0 {
		_ok = true
	}

	return _resolvedGradient, _ok
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (gradient *Gradient) ResolveForContext(context *StyleContext) *cairo.Pattern {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(gradient)))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**Gradient)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gradient)
	runtime.KeepAlive(context)

	var _pattern *cairo.Pattern // out

	{
		_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
		_pattern = (*cairo.Pattern)(unsafe.Pointer(_pp))
	}
	runtime.SetFinalizer(_pattern, func(v *cairo.Pattern) {
		C.cairo_pattern_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})

	return _pattern
}

// String creates a string representation for gradient that is suitable for
// using in GTK CSS files.
//
// Deprecated: Gradient is deprecated.
//
// The function returns the following values:
//
//    - utf8: string representation for gradient.
//
func (gradient *Gradient) String() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(gradient)))
	*(**Gradient)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(gradient)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SymbolicColor is a boxed type that represents a symbolic color. It is the
// result of parsing a [color expression][gtkcssprovider-symbolic-colors]. To
// obtain the color represented by a GtkSymbolicColor, it has to be resolved
// with gtk_symbolic_color_resolve(), which replaces all symbolic color
// references by the colors they refer to (in a given context) and evaluates
// mix, shade and other expressions, resulting in a RGBA value.
//
// It is not normally necessary to deal directly with SymbolicColors, since they
// are mostly used behind the scenes by StyleContext and CssProvider.
//
// SymbolicColor is deprecated. Symbolic colors are considered an implementation
// detail of GTK+.
//
// An instance of this type is always passed by reference.
type SymbolicColor struct {
	*symbolicColor
}

// symbolicColor is the struct that's finalized.
type symbolicColor struct {
	native *C.GtkSymbolicColor
}

func marshalSymbolicColor(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SymbolicColor{&symbolicColor{(*C.GtkSymbolicColor)(b)}}, nil
}

// NewSymbolicColorAlpha constructs a struct SymbolicColor.
func NewSymbolicColorAlpha(color *SymbolicColor, factor float64) *SymbolicColor {
	var args [2]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.gdouble // out
	var _cret *C.void   // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))
	_arg1 = C.gdouble(factor)
	*(**SymbolicColor)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(color)
	runtime.KeepAlive(factor)

	var _symbolicColor *SymbolicColor // out

	_symbolicColor = (*SymbolicColor)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_symbolicColor)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(intern.C))
		},
	)

	return _symbolicColor
}

// NewSymbolicColorLiteral constructs a struct SymbolicColor.
func NewSymbolicColorLiteral(color *gdk.RGBA) *SymbolicColor {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))
	*(**gdk.RGBA)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(color)

	var _symbolicColor *SymbolicColor // out

	_symbolicColor = (*SymbolicColor)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_symbolicColor)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(intern.C))
		},
	)

	return _symbolicColor
}

// NewSymbolicColorMix constructs a struct SymbolicColor.
func NewSymbolicColorMix(color1 *SymbolicColor, color2 *SymbolicColor, factor float64) *SymbolicColor {
	var args [3]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 *C.void   // out
	var _arg2 C.gdouble // out
	var _cret *C.void   // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(color1)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(color2)))
	_arg2 = C.gdouble(factor)
	*(**SymbolicColor)(unsafe.Pointer(&args[0])) = _arg0
	*(**SymbolicColor)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(color1)
	runtime.KeepAlive(color2)
	runtime.KeepAlive(factor)

	var _symbolicColor *SymbolicColor // out

	_symbolicColor = (*SymbolicColor)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_symbolicColor)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(intern.C))
		},
	)

	return _symbolicColor
}

// NewSymbolicColorName constructs a struct SymbolicColor.
func NewSymbolicColorName(name string) *SymbolicColor {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg0))
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)

	var _symbolicColor *SymbolicColor // out

	_symbolicColor = (*SymbolicColor)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_symbolicColor)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(intern.C))
		},
	)

	return _symbolicColor
}

// NewSymbolicColorShade constructs a struct SymbolicColor.
func NewSymbolicColorShade(color *SymbolicColor, factor float64) *SymbolicColor {
	var args [2]girepository.Argument
	var _arg0 *C.void   // out
	var _arg1 C.gdouble // out
	var _cret *C.void   // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))
	_arg1 = C.gdouble(factor)
	*(**SymbolicColor)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(color)
	runtime.KeepAlive(factor)

	var _symbolicColor *SymbolicColor // out

	_symbolicColor = (*SymbolicColor)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_symbolicColor)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(intern.C))
		},
	)

	return _symbolicColor
}

// NewSymbolicColorWin32 constructs a struct SymbolicColor.
func NewSymbolicColorWin32(themeClass string, id int) *SymbolicColor {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(themeClass)))
	defer C.free(unsafe.Pointer(_arg0))
	_arg1 = C.gint(id)
	*(*string)(unsafe.Pointer(&args[0])) = _arg0
	*(*int)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(themeClass)
	runtime.KeepAlive(id)

	var _symbolicColor *SymbolicColor // out

	_symbolicColor = (*SymbolicColor)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_symbolicColor)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(intern.C))
		},
	)

	return _symbolicColor
}

// String converts the given color to a string representation. This is useful
// both for debugging and for serialization of strings. The format of the string
// may change between different versions of GTK, but it is guaranteed that the
// GTK css parser is able to read the string and create the same symbolic color
// from it.
//
// Deprecated: SymbolicColor is deprecated.
//
// The function returns the following values:
//
//    - utf8: new string representing color.
//
func (color *SymbolicColor) String() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(color)))
	*(**SymbolicColor)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(color)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
