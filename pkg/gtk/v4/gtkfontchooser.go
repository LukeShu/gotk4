// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_level_get_type()), F: marshalFontChooserLevel},
		{T: externglib.Type(C.gtk_font_chooser_get_type()), F: marshalFontChooser},
	})
}

// FontChooserLevel specifies the granularity of font selection that is desired
// in a `GtkFontChooser`.
//
// This enumeration may be extended in the future; applications should ignore
// unknown values.
type FontChooserLevel int

const (
	// FontChooserLevelFamily: allow selecting a font family
	FontChooserLevelFamily FontChooserLevel = 0b0
	// FontChooserLevelStyle: allow selecting a specific font face
	FontChooserLevelStyle FontChooserLevel = 0b1
	// FontChooserLevelSize: allow selecting a specific font size
	FontChooserLevelSize FontChooserLevel = 0b10
	// FontChooserLevelVariations: allow changing OpenType font variation axes
	FontChooserLevelVariations FontChooserLevel = 0b100
	// FontChooserLevelFeatures: allow selecting specific OpenType font features
	FontChooserLevelFeatures FontChooserLevel = 0b1000
)

func marshalFontChooserLevel(p uintptr) (interface{}, error) {
	return FontChooserLevel(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FontFilterFunc: the type of function that is used for deciding what fonts get
// shown in a `GtkFontChooser`.
//
// See [method@Gtk.FontChooser.set_filter_func].
type FontFilterFunc func(family *pango.FontFamilyClass, face *pango.FontFaceClass, data interface{}) (ok bool)

//export gotk4_FontFilterFunc
func gotk4_FontFilterFunc(arg0 *C.PangoFontFamily, arg1 *C.PangoFontFace, arg2 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var family *pango.FontFamilyClass // out
	var face *pango.FontFaceClass     // out
	var data interface{}              // out

	family = gextras.CastObject(
		externglib.Take(unsafe.Pointer(arg0))).(*pango.FontFamilyClass)
	face = gextras.CastObject(
		externglib.Take(unsafe.Pointer(arg1))).(*pango.FontFaceClass)
	data = box.Get(uintptr(arg2))

	fn := v.(FontFilterFunc)
	ok := fn(family, face, data)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontChooserOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FontChooserOverrider interface {
	FontActivated(fontname string)
	// FontFace gets the `PangoFontFace` representing the selected font group
	// details (i.e. family, slant, weight, width, etc).
	//
	// If the selected font is not installed, returns nil.
	FontFace() *pango.FontFaceClass
	// FontFamily gets the `PangoFontFamily` representing the selected font
	// family.
	//
	// Font families are a collection of font faces.
	//
	// If the selected font is not installed, returns nil.
	FontFamily() *pango.FontFamilyClass
	// FontMap gets the custom font map of this font chooser widget, or nil if
	// it does not have one.
	FontMap() *pango.FontMapClass
	// FontSize: the selected font size.
	FontSize() int
	// SetFontMap sets a custom font map to use for this font chooser widget.
	//
	// A custom font map can be used to present application-specific fonts
	// instead of or in addition to the normal system fonts.
	//
	// “`c FcConfig *config; PangoFontMap *fontmap;
	//
	// config = FcInitLoadConfigAndFonts (); FcConfigAppFontAddFile (config,
	// my_app_font_file);
	//
	// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
	// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
	//
	// gtk_font_chooser_set_font_map (font_chooser, fontmap); “`
	//
	// Note that other GTK widgets will only be able to use the
	// application-specific font if it is present in the font map they use:
	//
	// “`c context = gtk_widget_get_pango_context (label);
	// pango_context_set_font_map (context, fontmap); “`
	SetFontMap(fontmap pango.FontMap)
}

// FontChooser: `GtkFontChooser` is an interface that can be implemented by
// widgets for choosing fonts.
//
// In GTK, the main objects that implement this interface are
// [class@Gtk.FontChooserWidget], [class@Gtk.FontChooserDialog] and
// [class@Gtk.FontButton].
type FontChooser interface {
	gextras.Objector

	// Font gets the currently-selected font name.
	//
	// Note that this can be a different string than what you set with
	// [method@Gtk.FontChooser.set_font], as the font chooser widget may
	// normalize font names and thus return a string with a different structure.
	// For example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
	// Bold Italic 12”.
	//
	// Use [method@Pango.FontDescription.equal] if you want to compare two font
	// descriptions.
	Font() string
	// FontDesc gets the currently-selected font.
	//
	// Note that this can be a different string than what you set with
	// [method@Gtk.FontChooser.set_font], as the font chooser widget may
	// normalize font names and thus return a string with a different structure.
	// For example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
	// Bold Italic 12”.
	//
	// Use [method@Pango.FontDescription.equal] if you want to compare two font
	// descriptions.
	FontDesc() *pango.FontDescription
	// FontFace gets the `PangoFontFace` representing the selected font group
	// details (i.e. family, slant, weight, width, etc).
	//
	// If the selected font is not installed, returns nil.
	FontFace() *pango.FontFaceClass
	// FontFamily gets the `PangoFontFamily` representing the selected font
	// family.
	//
	// Font families are a collection of font faces.
	//
	// If the selected font is not installed, returns nil.
	FontFamily() *pango.FontFamilyClass
	// FontFeatures gets the currently-selected font features.
	FontFeatures() string
	// FontMap gets the custom font map of this font chooser widget, or nil if
	// it does not have one.
	FontMap() *pango.FontMapClass
	// FontSize: the selected font size.
	FontSize() int
	// Language gets the language that is used for font features.
	Language() string
	// Level returns the current level of granularity for selecting fonts.
	Level() FontChooserLevel
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// ShowPreviewEntry returns whether the preview entry is shown or not.
	ShowPreviewEntry() bool
	// SetFont sets the currently-selected font.
	SetFont(fontname string)
	// SetFontDesc sets the currently-selected font from @font_desc.
	SetFontDesc(fontDesc *pango.FontDescription)
	// SetFontMap sets a custom font map to use for this font chooser widget.
	//
	// A custom font map can be used to present application-specific fonts
	// instead of or in addition to the normal system fonts.
	//
	// “`c FcConfig *config; PangoFontMap *fontmap;
	//
	// config = FcInitLoadConfigAndFonts (); FcConfigAppFontAddFile (config,
	// my_app_font_file);
	//
	// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
	// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
	//
	// gtk_font_chooser_set_font_map (font_chooser, fontmap); “`
	//
	// Note that other GTK widgets will only be able to use the
	// application-specific font if it is present in the font map they use:
	//
	// “`c context = gtk_widget_get_pango_context (label);
	// pango_context_set_font_map (context, fontmap); “`
	SetFontMap(fontmap pango.FontMap)
	// SetLanguage sets the language to use for font features.
	SetLanguage(language string)
	// SetPreviewText sets the text displayed in the preview area.
	//
	// The @text is used to show how the selected font looks.
	SetPreviewText(text string)
	// SetShowPreviewEntry shows or hides the editable preview entry.
	SetShowPreviewEntry(showPreviewEntry bool)
}

// FontChooserInterface implements the FontChooser interface.
type FontChooserInterface struct {
	*externglib.Object
}

var _ FontChooser = (*FontChooserInterface)(nil)

func wrapFontChooser(obj *externglib.Object) FontChooser {
	return &FontChooserInterface{
		Object: obj,
	}
}

func marshalFontChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontChooser(obj), nil
}

// Font gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// [method@Gtk.FontChooser.set_font], as the font chooser widget may normalize
// font names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use [method@Pango.FontDescription.equal] if you want to compare two font
// descriptions.
func (f *FontChooserInterface) Font() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_font(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontDesc gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// [method@Gtk.FontChooser.set_font], as the font chooser widget may normalize
// font names and thus return a string with a different structure. For example,
// “Helvetica Italic Bold 12” could be normalized to “Helvetica Bold Italic 12”.
//
// Use [method@Pango.FontDescription.equal] if you want to compare two font
// descriptions.
func (f *FontChooserInterface) FontDesc() *pango.FontDescription {
	var _arg0 *C.GtkFontChooser       // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_font_desc(_arg0)

	var _fontDescription *pango.FontDescription // out

	_fontDescription = (*pango.FontDescription)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_fontDescription, func(v *pango.FontDescription) {
		C.free(unsafe.Pointer(v))
	})

	return _fontDescription
}

// FontFace gets the `PangoFontFace` representing the selected font group
// details (i.e. family, slant, weight, width, etc).
//
// If the selected font is not installed, returns nil.
func (f *FontChooserInterface) FontFace() *pango.FontFaceClass {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontFace  // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_font_face(_arg0)

	var _fontFace *pango.FontFaceClass // out

	_fontFace = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*pango.FontFaceClass)

	return _fontFace
}

// FontFamily gets the `PangoFontFamily` representing the selected font family.
//
// Font families are a collection of font faces.
//
// If the selected font is not installed, returns nil.
func (f *FontChooserInterface) FontFamily() *pango.FontFamilyClass {
	var _arg0 *C.GtkFontChooser  // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_font_family(_arg0)

	var _fontFamily *pango.FontFamilyClass // out

	_fontFamily = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*pango.FontFamilyClass)

	return _fontFamily
}

// FontFeatures gets the currently-selected font features.
func (f *FontChooserInterface) FontFeatures() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_font_features(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontMap gets the custom font map of this font chooser widget, or nil if it
// does not have one.
func (f *FontChooserInterface) FontMap() *pango.FontMapClass {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontMap   // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_font_map(_arg0)

	var _fontMap *pango.FontMapClass // out

	_fontMap = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*pango.FontMapClass)

	return _fontMap
}

// FontSize: the selected font size.
func (f *FontChooserInterface) FontSize() int {
	var _arg0 *C.GtkFontChooser // out
	var _cret C.int             // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_font_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Language gets the language that is used for font features.
func (f *FontChooserInterface) Language() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_language(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Level returns the current level of granularity for selecting fonts.
func (f *FontChooserInterface) Level() FontChooserLevel {
	var _arg0 *C.GtkFontChooser     // out
	var _cret C.GtkFontChooserLevel // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_level(_arg0)

	var _fontChooserLevel FontChooserLevel // out

	_fontChooserLevel = (FontChooserLevel)(_cret)

	return _fontChooserLevel
}

// PreviewText gets the text displayed in the preview area.
func (f *FontChooserInterface) PreviewText() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_preview_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ShowPreviewEntry returns whether the preview entry is shown or not.
func (f *FontChooserInterface) ShowPreviewEntry() bool {
	var _arg0 *C.GtkFontChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))

	_cret = C.gtk_font_chooser_get_show_preview_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFont sets the currently-selected font.
func (f *FontChooserInterface) SetFont(fontname string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))
	_arg1 = (*C.char)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_font(_arg0, _arg1)
}

// SetFontDesc sets the currently-selected font from @font_desc.
func (f *FontChooserInterface) SetFontDesc(fontDesc *pango.FontDescription) {
	var _arg0 *C.GtkFontChooser       // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(fontDesc))

	C.gtk_font_chooser_set_font_desc(_arg0, _arg1)
}

// SetFontMap sets a custom font map to use for this font chooser widget.
//
// A custom font map can be used to present application-specific fonts instead
// of or in addition to the normal system fonts.
//
// “`c FcConfig *config; PangoFontMap *fontmap;
//
// config = FcInitLoadConfigAndFonts (); FcConfigAppFontAddFile (config,
// my_app_font_file);
//
// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
// gtk_font_chooser_set_font_map (font_chooser, fontmap); “`
//
// Note that other GTK widgets will only be able to use the application-specific
// font if it is present in the font map they use:
//
// “`c context = gtk_widget_get_pango_context (label);
// pango_context_set_font_map (context, fontmap); “`
func (f *FontChooserInterface) SetFontMap(fontmap pango.FontMap) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.PangoFontMap   // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))
	_arg1 = (*C.PangoFontMap)(unsafe.Pointer((&fontmap).Native()))

	C.gtk_font_chooser_set_font_map(_arg0, _arg1)
}

// SetLanguage sets the language to use for font features.
func (f *FontChooserInterface) SetLanguage(language string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))
	_arg1 = (*C.char)(C.CString(language))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_language(_arg0, _arg1)
}

// SetPreviewText sets the text displayed in the preview area.
//
// The @text is used to show how the selected font looks.
func (f *FontChooserInterface) SetPreviewText(text string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_preview_text(_arg0, _arg1)
}

// SetShowPreviewEntry shows or hides the editable preview entry.
func (f *FontChooserInterface) SetShowPreviewEntry(showPreviewEntry bool) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer((&f).Native()))
	if showPreviewEntry {
		_arg1 = C.TRUE
	}

	C.gtk_font_chooser_set_show_preview_entry(_arg0, _arg1)
}
