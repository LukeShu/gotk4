// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern PangoFontFace* _gotk4_gtk3_FontChooserIface_get_font_face(GtkFontChooser*);
// extern PangoFontFamily* _gotk4_gtk3_FontChooserIface_get_font_family(GtkFontChooser*);
// extern PangoFontMap* _gotk4_gtk3_FontChooserIface_get_font_map(GtkFontChooser*);
// extern gint _gotk4_gtk3_FontChooserIface_get_font_size(GtkFontChooser*);
// extern void _gotk4_gtk3_FontChooserIface_font_activated(GtkFontChooser*, gchar*);
// extern void _gotk4_gtk3_FontChooserIface_set_font_map(GtkFontChooser*, PangoFontMap*);
// extern void _gotk4_gtk3_FontChooser_ConnectFontActivated(gpointer, gchar*, guintptr);
import "C"

// glib.Type values for gtkfontchooser.go.
var (
	GTypeFontChooserLevel = coreglib.Type(C.gtk_font_chooser_level_get_type())
	GTypeFontChooser      = coreglib.Type(C.gtk_font_chooser_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFontChooserLevel, F: marshalFontChooserLevel},
		{T: GTypeFontChooser, F: marshalFontChooser},
	})
}

// FontChooserLevel: this enumeration specifies the granularity of font
// selection that is desired in a font chooser.
//
// This enumeration may be extended in the future; applications should ignore
// unknown values.
type FontChooserLevel C.guint

const (
	// FontChooserLevelFamily: allow selecting a font family.
	FontChooserLevelFamily FontChooserLevel = 0b0
	// FontChooserLevelStyle: allow selecting a specific font face.
	FontChooserLevelStyle FontChooserLevel = 0b1
	// FontChooserLevelSize: allow selecting a specific font size.
	FontChooserLevelSize       FontChooserLevel = 0b10
	FontChooserLevelVariations FontChooserLevel = 0b100
	// FontChooserLevelFeatures: allow selecting specific OpenType font
	// features.
	FontChooserLevelFeatures FontChooserLevel = 0b1000
)

func marshalFontChooserLevel(p uintptr) (interface{}, error) {
	return FontChooserLevel(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for FontChooserLevel.
func (f FontChooserLevel) String() string {
	if f == 0 {
		return "FontChooserLevel(0)"
	}

	var builder strings.Builder
	builder.Grow(117)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FontChooserLevelFamily:
			builder.WriteString("Family|")
		case FontChooserLevelStyle:
			builder.WriteString("Style|")
		case FontChooserLevelSize:
			builder.WriteString("Size|")
		case FontChooserLevelVariations:
			builder.WriteString("Variations|")
		case FontChooserLevelFeatures:
			builder.WriteString("Features|")
		default:
			builder.WriteString(fmt.Sprintf("FontChooserLevel(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FontChooserLevel) Has(other FontChooserLevel) bool {
	return (f & other) == other
}

// FontFilterFunc: type of function that is used for deciding what fonts get
// shown in a FontChooser. See gtk_font_chooser_set_filter_func().
type FontFilterFunc func(family pango.FontFamilier, face pango.FontFacer) (ok bool)

//export _gotk4_gtk3_FontFilterFunc
func _gotk4_gtk3_FontFilterFunc(arg1 *C.PangoFontFamily, arg2 *C.PangoFontFace, arg3 C.gpointer) (cret C.gboolean) {
	var fn FontFilterFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(FontFilterFunc)
	}

	var _family pango.FontFamilier // out
	var _face pango.FontFacer      // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type pango.FontFamilier is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(pango.FontFamilier)
			return ok
		})
		rv, ok := casted.(pango.FontFamilier)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFamilier")
		}
		_family = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type pango.FontFacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(pango.FontFacer)
			return ok
		})
		rv, ok := casted.(pango.FontFacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFacer")
		}
		_face = rv
	}

	ok := fn(_family, _face)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontChooserOverrider contains methods that are overridable.
type FontChooserOverrider interface {
	// The function takes the following parameters:
	//
	FontActivated(fontname string)
	// FontFace gets the FontFace representing the selected font group details
	// (i.e. family, slant, weight, width, etc).
	//
	// If the selected font is not installed, returns NULL.
	//
	// The function returns the following values:
	//
	//    - fontFace (optional) representing the selected font group details, or
	//      NULL. The returned object is owned by fontchooser and must not be
	//      modified or freed.
	//
	FontFace() pango.FontFacer
	// FontFamily gets the FontFamily representing the selected font family.
	// Font families are a collection of font faces.
	//
	// If the selected font is not installed, returns NULL.
	//
	// The function returns the following values:
	//
	//    - fontFamily (optional) representing the selected font family, or NULL.
	//      The returned object is owned by fontchooser and must not be modified
	//      or freed.
	//
	FontFamily() pango.FontFamilier
	// FontMap gets the custom font map of this font chooser widget, or NULL if
	// it does not have one.
	//
	// The function returns the following values:
	//
	//    - fontMap (optional) or NULL.
	//
	FontMap() pango.FontMapper
	// FontSize: selected font size.
	//
	// The function returns the following values:
	//
	//    - gint: n integer representing the selected font size, or -1 if no font
	//      size is selected.
	//
	FontSize() int32
	// SetFontMap sets a custom font map to use for this font chooser widget. A
	// custom font map can be used to present application-specific fonts instead
	// of or in addition to the normal system fonts.
	//
	//    FcConfig *config;
	//    PangoFontMap *fontmap;
	//
	//    config = FcInitLoadConfigAndFonts ();
	//    FcConfigAppFontAddFile (config, my_app_font_file);
	//
	//    fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
	//    pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
	//
	//    gtk_font_chooser_set_font_map (font_chooser, fontmap);
	//
	// Note that other GTK+ widgets will only be able to use the
	// application-specific font if it is present in the font map they use:
	//
	//    context = gtk_widget_get_pango_context (label);
	//    pango_context_set_font_map (context, fontmap);.
	//
	// The function takes the following parameters:
	//
	//    - fontmap (optional): FontMap.
	//
	SetFontMap(fontmap pango.FontMapper)
}

// FontChooser is an interface that can be implemented by widgets displaying the
// list of fonts. In GTK+, the main objects that implement this interface are
// FontChooserWidget, FontChooserDialog and FontButton. The GtkFontChooser
// interface has been introducted in GTK+ 3.2.
//
// FontChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type FontChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FontChooser)(nil)
)

// FontChooserer describes FontChooser's interface methods.
type FontChooserer interface {
	coreglib.Objector

	// Font gets the currently-selected font name.
	Font() string
	// FontDesc gets the currently-selected font.
	FontDesc() *pango.FontDescription
	// FontFace gets the FontFace representing the selected font group details
	// (i.e.
	FontFace() pango.FontFacer
	// FontFamily gets the FontFamily representing the selected font family.
	FontFamily() pango.FontFamilier
	// FontFeatures gets the currently-selected font features.
	FontFeatures() string
	// FontMap gets the custom font map of this font chooser widget, or NULL if
	// it does not have one.
	FontMap() pango.FontMapper
	// FontSize: selected font size.
	FontSize() int32
	// Language gets the language that is used for font features.
	Language() string
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// ShowPreviewEntry returns whether the preview entry is shown or not.
	ShowPreviewEntry() bool
	// SetFont sets the currently-selected font.
	SetFont(fontname string)
	// SetFontDesc sets the currently-selected font from font_desc.
	SetFontDesc(fontDesc *pango.FontDescription)
	// SetFontMap sets a custom font map to use for this font chooser widget.
	SetFontMap(fontmap pango.FontMapper)
	// SetLanguage sets the language to use for font features.
	SetLanguage(language string)
	// SetPreviewText sets the text displayed in the preview area.
	SetPreviewText(text string)
	// SetShowPreviewEntry shows or hides the editable preview entry.
	SetShowPreviewEntry(showPreviewEntry bool)

	// Font-activated is emitted when a font is activated.
	ConnectFontActivated(func(fontname string)) coreglib.SignalHandle
}

var _ FontChooserer = (*FontChooser)(nil)

func ifaceInitFontChooserer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkFontChooserIface)(unsafe.Pointer(gifacePtr))
	iface.font_activated = (*[0]byte)(C._gotk4_gtk3_FontChooserIface_font_activated)
	iface.get_font_face = (*[0]byte)(C._gotk4_gtk3_FontChooserIface_get_font_face)
	iface.get_font_family = (*[0]byte)(C._gotk4_gtk3_FontChooserIface_get_font_family)
	iface.get_font_map = (*[0]byte)(C._gotk4_gtk3_FontChooserIface_get_font_map)
	iface.get_font_size = (*[0]byte)(C._gotk4_gtk3_FontChooserIface_get_font_size)
	iface.set_font_map = (*[0]byte)(C._gotk4_gtk3_FontChooserIface_set_font_map)
}

//export _gotk4_gtk3_FontChooserIface_font_activated
func _gotk4_gtk3_FontChooserIface_font_activated(arg0 *C.GtkFontChooser, arg1 *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(FontChooserOverrider)

	var _fontname string // out

	_fontname = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.FontActivated(_fontname)
}

//export _gotk4_gtk3_FontChooserIface_get_font_face
func _gotk4_gtk3_FontChooserIface_get_font_face(arg0 *C.GtkFontChooser) (cret *C.PangoFontFace) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(FontChooserOverrider)

	fontFace := iface.FontFace()

	if fontFace != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontFace).Native()))
	}

	return cret
}

//export _gotk4_gtk3_FontChooserIface_get_font_family
func _gotk4_gtk3_FontChooserIface_get_font_family(arg0 *C.GtkFontChooser) (cret *C.PangoFontFamily) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(FontChooserOverrider)

	fontFamily := iface.FontFamily()

	if fontFamily != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontFamily).Native()))
	}

	return cret
}

//export _gotk4_gtk3_FontChooserIface_get_font_map
func _gotk4_gtk3_FontChooserIface_get_font_map(arg0 *C.GtkFontChooser) (cret *C.PangoFontMap) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(FontChooserOverrider)

	fontMap := iface.FontMap()

	if fontMap != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontMap).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(fontMap).Native()))
	}

	return cret
}

//export _gotk4_gtk3_FontChooserIface_get_font_size
func _gotk4_gtk3_FontChooserIface_get_font_size(arg0 *C.GtkFontChooser) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(FontChooserOverrider)

	gint := iface.FontSize()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtk3_FontChooserIface_set_font_map
func _gotk4_gtk3_FontChooserIface_set_font_map(arg0 *C.GtkFontChooser, arg1 *C.PangoFontMap) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(FontChooserOverrider)

	var _fontmap pango.FontMapper // out

	if arg1 != nil {
		{
			objptr := unsafe.Pointer(arg1)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontMapper)
				return ok
			})
			rv, ok := casted.(pango.FontMapper)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontMapper")
			}
			_fontmap = rv
		}
	}

	iface.SetFontMap(_fontmap)
}

func wrapFontChooser(obj *coreglib.Object) *FontChooser {
	return &FontChooser{
		Object: obj,
	}
}

func marshalFontChooser(p uintptr) (interface{}, error) {
	return wrapFontChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_FontChooser_ConnectFontActivated
func _gotk4_gtk3_FontChooser_ConnectFontActivated(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(fontname string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(fontname string))
	}

	var _fontname string // out

	_fontname = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_fontname)
}

// ConnectFontActivated is emitted when a font is activated. This usually
// happens when the user double clicks an item, or an item is selected and the
// user presses one of the keys Space, Shift+Space, Return or Enter.
func (fontchooser *FontChooser) ConnectFontActivated(f func(fontname string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(fontchooser, "font-activated", false, unsafe.Pointer(C._gotk4_gtk3_FontChooser_ConnectFontActivated), f)
}

// Font gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may normalize font
// names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two font
// descriptions.
//
// The function returns the following values:
//
//    - utf8 (optional): string with the name of the current font, or NULL if no
//      font is selected. You must free this string with g_free().
//
func (fontchooser *FontChooser) Font() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// FontDesc gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may normalize font
// names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two font
// descriptions.
//
// The function returns the following values:
//
//    - fontDescription (optional) for the current font, or NULL if no font is
//      selected.
//
func (fontchooser *FontChooser) FontDesc() *pango.FontDescription {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontDescription *pango.FontDescription // out

	if _cret != nil {
		_fontDescription = (*pango.FontDescription)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_fontDescription)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_font_description_free((*C.PangoFontDescription)(intern.C))
			},
		)
	}

	return _fontDescription
}

// FontFace gets the FontFace representing the selected font group details (i.e.
// family, slant, weight, width, etc).
//
// If the selected font is not installed, returns NULL.
//
// The function returns the following values:
//
//    - fontFace (optional) representing the selected font group details, or
//      NULL. The returned object is owned by fontchooser and must not be
//      modified or freed.
//
func (fontchooser *FontChooser) FontFace() pango.FontFacer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontFace pango.FontFacer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontFacer)
				return ok
			})
			rv, ok := casted.(pango.FontFacer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFacer")
			}
			_fontFace = rv
		}
	}

	return _fontFace
}

// FontFamily gets the FontFamily representing the selected font family. Font
// families are a collection of font faces.
//
// If the selected font is not installed, returns NULL.
//
// The function returns the following values:
//
//    - fontFamily (optional) representing the selected font family, or NULL. The
//      returned object is owned by fontchooser and must not be modified or
//      freed.
//
func (fontchooser *FontChooser) FontFamily() pango.FontFamilier {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontFamily pango.FontFamilier // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontFamilier)
				return ok
			})
			rv, ok := casted.(pango.FontFamilier)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontFamilier")
			}
			_fontFamily = rv
		}
	}

	return _fontFamily
}

// FontFeatures gets the currently-selected font features.
//
// The function returns the following values:
//
//    - utf8: currently selected font features.
//
func (fontchooser *FontChooser) FontFeatures() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontMap gets the custom font map of this font chooser widget, or NULL if it
// does not have one.
//
// The function returns the following values:
//
//    - fontMap (optional) or NULL.
//
func (fontchooser *FontChooser) FontMap() pango.FontMapper {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _fontMap pango.FontMapper // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(pango.FontMapper)
				return ok
			})
			rv, ok := casted.(pango.FontMapper)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.FontMapper")
			}
			_fontMap = rv
		}
	}

	return _fontMap
}

// FontSize: selected font size.
//
// The function returns the following values:
//
//    - gint: n integer representing the selected font size, or -1 if no font
//      size is selected.
//
func (fontchooser *FontChooser) FontSize() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Language gets the language that is used for font features.
//
// The function returns the following values:
//
//    - utf8: currently selected language.
//
func (fontchooser *FontChooser) Language() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PreviewText gets the text displayed in the preview area.
//
// The function returns the following values:
//
//    - utf8: text displayed in the preview area.
//
func (fontchooser *FontChooser) PreviewText() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ShowPreviewEntry returns whether the preview entry is shown or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the preview entry is shown or FALSE if it is hidden.
//
func (fontchooser *FontChooser) ShowPreviewEntry() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	*(**FontChooser)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontchooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFont sets the currently-selected font.
//
// The function takes the following parameters:
//
//    - fontname: font name like “Helvetica 12” or “Times Bold 18”.
//
func (fontchooser *FontChooser) SetFont(fontname string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**FontChooser)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontname)
}

// SetFontDesc sets the currently-selected font from font_desc.
//
// The function takes the following parameters:
//
//    - fontDesc: FontDescription.
//
func (fontchooser *FontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(fontDesc)))
	*(**FontChooser)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontDesc)
}

// SetFontMap sets a custom font map to use for this font chooser widget. A
// custom font map can be used to present application-specific fonts instead of
// or in addition to the normal system fonts.
//
//    FcConfig *config;
//    PangoFontMap *fontmap;
//
//    config = FcInitLoadConfigAndFonts ();
//    FcConfigAppFontAddFile (config, my_app_font_file);
//
//    fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
//    pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
//    gtk_font_chooser_set_font_map (font_chooser, fontmap);
//
// Note that other GTK+ widgets will only be able to use the
// application-specific font if it is present in the font map they use:
//
//    context = gtk_widget_get_pango_context (label);
//    pango_context_set_font_map (context, fontmap);.
//
// The function takes the following parameters:
//
//    - fontmap (optional): FontMap.
//
func (fontchooser *FontChooser) SetFontMap(fontmap pango.FontMapper) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if fontmap != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontmap).Native()))
	}
	*(**FontChooser)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontmap)
}

// SetLanguage sets the language to use for font features.
//
// The function takes the following parameters:
//
//    - language: language.
//
func (fontchooser *FontChooser) SetLanguage(language string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(language)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**FontChooser)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(language)
}

// SetPreviewText sets the text displayed in the preview area. The text is used
// to show how the selected font looks.
//
// The function takes the following parameters:
//
//    - text to display in the preview area.
//
func (fontchooser *FontChooser) SetPreviewText(text string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	*(**FontChooser)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(text)
}

// SetShowPreviewEntry shows or hides the editable preview entry.
//
// The function takes the following parameters:
//
//    - showPreviewEntry: whether to show the editable preview entry or not.
//
func (fontchooser *FontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if showPreviewEntry {
		_arg1 = C.TRUE
	}
	*(**FontChooser)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(showPreviewEntry)
}
