// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_StyleContextClass_changed(GtkStyleContext*);
// void _gotk4_gtk4_StyleContext_virtual_changed(void* fnptr, GtkStyleContext* arg0) {
//   ((void (*)(GtkStyleContext*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeStyleContextPrintFlags = coreglib.Type(C.gtk_style_context_print_flags_get_type())
	GTypeStyleContext           = coreglib.Type(C.gtk_style_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStyleContextPrintFlags, F: marshalStyleContextPrintFlags},
		coreglib.TypeMarshaler{T: GTypeStyleContext, F: marshalStyleContext},
	})
}

// StyleContextPrintFlags flags that modify the behavior of
// gtk_style_context_to_string().
//
// New values may be added to this enumeration.
type StyleContextPrintFlags C.guint

const (
	// StyleContextPrintNone: default value.
	StyleContextPrintNone StyleContextPrintFlags = 0b0
	// StyleContextPrintRecurse: print the entire tree of CSS nodes starting at
	// the style context's node.
	StyleContextPrintRecurse StyleContextPrintFlags = 0b1
	// StyleContextPrintShowStyle: show the values of the CSS properties for
	// each node.
	StyleContextPrintShowStyle StyleContextPrintFlags = 0b10
	// StyleContextPrintShowChange: show information about what changes affect
	// the styles.
	StyleContextPrintShowChange StyleContextPrintFlags = 0b100
)

func marshalStyleContextPrintFlags(p uintptr) (interface{}, error) {
	return StyleContextPrintFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for StyleContextPrintFlags.
func (s StyleContextPrintFlags) String() string {
	if s == 0 {
		return "StyleContextPrintFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(101)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case StyleContextPrintNone:
			builder.WriteString("None|")
		case StyleContextPrintRecurse:
			builder.WriteString("Recurse|")
		case StyleContextPrintShowStyle:
			builder.WriteString("ShowStyle|")
		case StyleContextPrintShowChange:
			builder.WriteString("ShowChange|")
		default:
			builder.WriteString(fmt.Sprintf("StyleContextPrintFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s StyleContextPrintFlags) Has(other StyleContextPrintFlags) bool {
	return (s & other) == other
}

// StyleContextOverrides contains methods that are overridable.
type StyleContextOverrides struct {
	Changed func()
}

func defaultStyleContextOverrides(v *StyleContext) StyleContextOverrides {
	return StyleContextOverrides{
		Changed: v.changed,
	}
}

// StyleContext: GtkStyleContext stores styling information affecting a widget.
//
// In order to construct the final style information, GtkStyleContext queries
// information from all attached GtkStyleProviders. Style providers can be
// either attached explicitly to the context through
// gtk.StyleContext.AddProvider(), or to the display through
// gtk.StyleContext().AddProviderForDisplay. The resulting style is a
// combination of all providers’ information in priority order.
//
// For GTK widgets, any GtkStyleContext returned by gtk.Widget.GetStyleContext()
// will already have a GdkDisplay and RTL/LTR information set. The style context
// will also be updated automatically if any of these settings change on the
// widget.
//
//
// Style Classes
//
// Widgets can add style classes to their context, which can be used to
// associate different styles by class. The documentation for individual widgets
// lists which style classes it uses itself, and which style classes may be
// added by applications to affect their appearance.
//
//
// Custom styling in UI libraries and applications
//
// If you are developing a library with custom widgets that render differently
// than standard components, you may need to add a GtkStyleProvider yourself
// with the GTK_STYLE_PROVIDER_PRIORITY_FALLBACK priority, either a
// GtkCssProvider or a custom object implementing the GtkStyleProvider
// interface. This way themes may still attempt to style your UI elements in a
// different way if needed so.
//
// If you are using custom styling on an applications, you probably want then to
// make your style information prevail to the theme’s, so you must use a
// GtkStyleProvider with the GTK_STYLE_PROVIDER_PRIORITY_APPLICATION priority,
// keep in mind that the user settings in XDG_CONFIG_HOME/gtk-4.0/gtk.css will
// still take precedence over your changes, as it uses the
// GTK_STYLE_PROVIDER_PRIORITY_USER priority.
type StyleContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StyleContext)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StyleContext, *StyleContextClass, StyleContextOverrides](
		GTypeStyleContext,
		initStyleContextClass,
		wrapStyleContext,
		defaultStyleContextOverrides,
	)
}

func initStyleContextClass(gclass unsafe.Pointer, overrides StyleContextOverrides, classInitFunc func(*StyleContextClass)) {
	pclass := (*C.GtkStyleContextClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeStyleContext))))

	if overrides.Changed != nil {
		pclass.changed = (*[0]byte)(C._gotk4_gtk4_StyleContextClass_changed)
	}

	if classInitFunc != nil {
		class := (*StyleContextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStyleContext(obj *coreglib.Object) *StyleContext {
	return &StyleContext{
		Object: obj,
	}
}

func marshalStyleContext(p uintptr) (interface{}, error) {
	return wrapStyleContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddClass adds a style class to context, so later uses of the style context
// will make use of this new class for styling.
//
// In the CSS file format, a GtkEntry defining a “search” class, would be
// matched by:
//
//    entry.search { ... }
//
//
// While any widget defining a “search” class would be matched by:
//
//    .search { ... }.
//
// The function takes the following parameters:
//
//    - className class name to use in styling.
//
func (context *StyleContext) AddClass(className string) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(className)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_style_context_add_class(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(className)
}

// AddProvider adds a style provider to context, to be used in style
// construction.
//
// Note that a style provider added by this function only affects the style of
// the widget to which context belongs. If you want to affect the style of all
// widgets, use gtk.StyleContext().AddProviderForDisplay.
//
// Note: If both priorities are the same, a GtkStyleProvider added through this
// function takes precedence over another added through
// gtk.StyleContext().AddProviderForDisplay.
//
// The function takes the following parameters:
//
//    - provider: GtkStyleProvider.
//    - priority of the style provider. The lower it is, the earlier it will be
//      used in the style construction. Typically this will be in the range
//      between GTK_STYLE_PROVIDER_PRIORITY_FALLBACK and
//      GTK_STYLE_PROVIDER_PRIORITY_USER.
//
func (context *StyleContext) AddProvider(provider StyleProviderer, priority uint) {
	var _arg0 *C.GtkStyleContext  // out
	var _arg1 *C.GtkStyleProvider // out
	var _arg2 C.guint             // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg2 = C.guint(priority)

	C.gtk_style_context_add_provider(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(priority)
}

// Border gets the border for a given state as a GtkBorder.
//
// The function returns the following values:
//
//    - border: return value for the border settings.
//
func (context *StyleContext) Border() *Border {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.GtkBorder        // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_style_context_get_border(_arg0, &_arg1)
	runtime.KeepAlive(context)

	var _border *Border // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _border
}

// Color gets the foreground color for a given state.
//
// The function returns the following values:
//
//    - color: return value for the foreground color.
//
func (context *StyleContext) Color() *gdk.RGBA {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.GdkRGBA          // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_style_context_get_color(_arg0, &_arg1)
	runtime.KeepAlive(context)

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// Display returns the GdkDisplay to which context is attached.
//
// The function returns the following values:
//
//    - display: GdkDisplay.
//
func (context *StyleContext) Display() *gdk.Display {
	var _arg0 *C.GtkStyleContext // out
	var _cret *C.GdkDisplay      // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gtk_style_context_get_display(_arg0)
	runtime.KeepAlive(context)

	var _display *gdk.Display // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// Margin gets the margin for a given state as a GtkBorder.
//
// The function returns the following values:
//
//    - margin: return value for the margin settings.
//
func (context *StyleContext) Margin() *Border {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.GtkBorder        // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_style_context_get_margin(_arg0, &_arg1)
	runtime.KeepAlive(context)

	var _margin *Border // out

	_margin = (*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _margin
}

// Padding gets the padding for a given state as a GtkBorder.
//
// The function returns the following values:
//
//    - padding: return value for the padding settings.
//
func (context *StyleContext) Padding() *Border {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.GtkBorder        // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_style_context_get_padding(_arg0, &_arg1)
	runtime.KeepAlive(context)

	var _padding *Border // out

	_padding = (*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _padding
}

// Scale returns the scale used for assets.
//
// The function returns the following values:
//
//    - gint: scale.
//
func (context *StyleContext) Scale() int {
	var _arg0 *C.GtkStyleContext // out
	var _cret C.int              // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gtk_style_context_get_scale(_arg0)
	runtime.KeepAlive(context)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// State returns the state used for style matching.
//
// This method should only be used to retrieve the GtkStateFlags to pass to
// GtkStyleContext methods, like gtk.StyleContext.GetPadding(). If you need to
// retrieve the current state of a GtkWidget, use gtk.Widget.GetStateFlags().
//
// The function returns the following values:
//
//    - stateFlags: state flags.
//
func (context *StyleContext) State() StateFlags {
	var _arg0 *C.GtkStyleContext // out
	var _cret C.GtkStateFlags    // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gtk_style_context_get_state(_arg0)
	runtime.KeepAlive(context)

	var _stateFlags StateFlags // out

	_stateFlags = StateFlags(_cret)

	return _stateFlags
}

// HasClass returns TRUE if context currently has defined the given class name.
//
// The function takes the following parameters:
//
//    - className class name.
//
// The function returns the following values:
//
//    - ok: TRUE if context has class_name defined.
//
func (context *StyleContext) HasClass(className string) bool {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.char            // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(className)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_context_has_class(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(className)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupColor looks up and resolves a color name in the context color map.
//
// The function takes the following parameters:
//
//    - colorName: color name to lookup.
//
// The function returns the following values:
//
//    - color: return location for the looked up color.
//    - ok: TRUE if color_name was found and resolved, FALSE otherwise.
//
func (context *StyleContext) LookupColor(colorName string) (*gdk.RGBA, bool) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.char            // out
	var _arg2 C.GdkRGBA          // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(colorName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_context_lookup_color(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(colorName)

	var _color *gdk.RGBA // out
	var _ok bool         // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

// RemoveClass removes class_name from context.
//
// The function takes the following parameters:
//
//    - className class name to remove.
//
func (context *StyleContext) RemoveClass(className string) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(className)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_style_context_remove_class(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(className)
}

// RemoveProvider removes provider from the style providers list in context.
//
// The function takes the following parameters:
//
//    - provider: GtkStyleProvider.
//
func (context *StyleContext) RemoveProvider(provider StyleProviderer) {
	var _arg0 *C.GtkStyleContext  // out
	var _arg1 *C.GtkStyleProvider // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	C.gtk_style_context_remove_provider(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(provider)
}

// Restore restores context state to a previous stage.
//
// See gtk.StyleContext.Save().
func (context *StyleContext) Restore() {
	var _arg0 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_style_context_restore(_arg0)
	runtime.KeepAlive(context)
}

// Save saves the context state.
//
// This allows temporary modifications done through gtk.StyleContext.AddClass(),
// gtk.StyleContext.RemoveClass(), gtk.StyleContext.SetState() to be quickly
// reverted in one go through gtk.StyleContext.Restore().
//
// The matching call to gtk.StyleContext.Restore() must be done before GTK
// returns to the main loop.
func (context *StyleContext) Save() {
	var _arg0 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gtk_style_context_save(_arg0)
	runtime.KeepAlive(context)
}

// SetDisplay attaches context to the given display.
//
// The display is used to add style information from “global” style providers,
// such as the display's GtkSettings instance.
//
// If you are using a GtkStyleContext returned from
// gtk.Widget.GetStyleContext(), you do not need to call this yourself.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//
func (context *StyleContext) SetDisplay(display *gdk.Display) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.GdkDisplay      // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	C.gtk_style_context_set_display(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(display)
}

// SetScale sets the scale to use when getting image assets for the style.
//
// The function takes the following parameters:
//
//    - scale: scale.
//
func (context *StyleContext) SetScale(scale int) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.int              // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.int(scale)

	C.gtk_style_context_set_scale(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(scale)
}

// SetState sets the state to be used for style matching.
//
// The function takes the following parameters:
//
//    - flags: state to represent.
//
func (context *StyleContext) SetState(flags StateFlags) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.GtkStateFlags    // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.GtkStateFlags(flags)

	C.gtk_style_context_set_state(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(flags)
}

// String converts the style context into a string representation.
//
// The string representation always includes information about the name, state,
// id, visibility and style classes of the CSS node that is backing context.
// Depending on the flags, more information may be included.
//
// This function is intended for testing and debugging of the CSS implementation
// in GTK. There are no guarantees about the format of the returned string, it
// may change.
//
// The function takes the following parameters:
//
//    - flags flags that determine what to print.
//
// The function returns the following values:
//
//    - utf8: newly allocated string representing context.
//
func (context *StyleContext) String(flags StyleContextPrintFlags) string {
	var _arg0 *C.GtkStyleContext          // out
	var _arg1 C.GtkStyleContextPrintFlags // out
	var _cret *C.char                     // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.GtkStyleContextPrintFlags(flags)

	_cret = C.gtk_style_context_to_string(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(flags)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (context *StyleContext) changed() {
	gclass := (*C.GtkStyleContextClass)(coreglib.PeekParentClass(context))
	fnarg := gclass.changed

	var _arg0 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C._gotk4_gtk4_StyleContext_virtual_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(context)
}

// StyleContextAddProviderForDisplay adds a global style provider to display,
// which will be used in style construction for all GtkStyleContexts under
// display.
//
// GTK uses this to make styling information from GtkSettings available.
//
// Note: If both priorities are the same, A GtkStyleProvider added through
// gtk.StyleContext.AddProvider() takes precedence over another added through
// this function.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//    - provider: GtkStyleProvider.
//    - priority of the style provider. The lower it is, the earlier it will be
//      used in the style construction. Typically this will be in the range
//      between GTK_STYLE_PROVIDER_PRIORITY_FALLBACK and
//      GTK_STYLE_PROVIDER_PRIORITY_USER.
//
func StyleContextAddProviderForDisplay(display *gdk.Display, provider StyleProviderer, priority uint) {
	var _arg1 *C.GdkDisplay       // out
	var _arg2 *C.GtkStyleProvider // out
	var _arg3 C.guint             // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg2 = (*C.GtkStyleProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))
	_arg3 = C.guint(priority)

	C.gtk_style_context_add_provider_for_display(_arg1, _arg2, _arg3)
	runtime.KeepAlive(display)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(priority)
}

// StyleContextRemoveProviderForDisplay removes provider from the global style
// providers list in display.
//
// The function takes the following parameters:
//
//    - display: GdkDisplay.
//    - provider: GtkStyleProvider.
//
func StyleContextRemoveProviderForDisplay(display *gdk.Display, provider StyleProviderer) {
	var _arg1 *C.GdkDisplay       // out
	var _arg2 *C.GtkStyleProvider // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))
	_arg2 = (*C.GtkStyleProvider)(unsafe.Pointer(coreglib.InternObject(provider).Native()))

	C.gtk_style_context_remove_provider_for_display(_arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(provider)
}

// StyleContextClass: instance of this type is always passed by reference.
type StyleContextClass struct {
	*styleContextClass
}

// styleContextClass is the struct that's finalized.
type styleContextClass struct {
	native *C.GtkStyleContextClass
}
