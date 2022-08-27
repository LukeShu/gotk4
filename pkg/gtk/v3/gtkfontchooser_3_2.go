// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// extern gboolean _gotk4_gtk3_FontFilterFunc(PangoFontFamily*, PangoFontFace*, gpointer);
// PangoFontFace* _gotk4_gtk3_FontChooser_virtual_get_font_face(void* fnptr, GtkFontChooser* arg0) {
//   return ((PangoFontFace* (*)(GtkFontChooser*))(fnptr))(arg0);
// };
// PangoFontFamily* _gotk4_gtk3_FontChooser_virtual_get_font_family(void* fnptr, GtkFontChooser* arg0) {
//   return ((PangoFontFamily* (*)(GtkFontChooser*))(fnptr))(arg0);
// };
// gint _gotk4_gtk3_FontChooser_virtual_get_font_size(void* fnptr, GtkFontChooser* arg0) {
//   return ((gint (*)(GtkFontChooser*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_FontChooser_virtual_set_filter_func(void* fnptr, GtkFontChooser* arg0, GtkFontFilterFunc arg1, gpointer arg2, GDestroyNotify arg3) {
//   ((void (*)(GtkFontChooser*, GtkFontFilterFunc, gpointer, GDestroyNotify))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

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
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_font(_arg0)
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
	var _arg0 *C.GtkFontChooser       // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_font_desc(_arg0)
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
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontFace  // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_font_face(_arg0)
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
	var _arg0 *C.GtkFontChooser  // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_font_family(_arg0)
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

// FontSize: selected font size.
//
// The function returns the following values:
//
//    - gint: n integer representing the selected font size, or -1 if no font
//      size is selected.
//
func (fontchooser *FontChooser) FontSize() int {
	var _arg0 *C.GtkFontChooser // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_font_size(_arg0)
	runtime.KeepAlive(fontchooser)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PreviewText gets the text displayed in the preview area.
//
// The function returns the following values:
//
//    - utf8: text displayed in the preview area.
//
func (fontchooser *FontChooser) PreviewText() string {
	var _arg0 *C.GtkFontChooser // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_preview_text(_arg0)
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
	var _arg0 *C.GtkFontChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C.gtk_font_chooser_get_show_preview_entry(_arg0)
	runtime.KeepAlive(fontchooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFilterFunc adds a filter function that decides which fonts to display in
// the font chooser.
//
// The function takes the following parameters:
//
//    - filter (optional) or NULL.
//
func (fontchooser *FontChooser) SetFilterFunc(filter FontFilterFunc) {
	var _arg0 *C.GtkFontChooser   // out
	var _arg1 C.GtkFontFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if filter != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_FontFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(filter))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_font_chooser_set_filter_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(filter)
}

// SetFont sets the currently-selected font.
//
// The function takes the following parameters:
//
//    - fontname: font name like “Helvetica 12” or “Times Bold 18”.
//
func (fontchooser *FontChooser) SetFont(fontname string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_font(_arg0, _arg1)
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
	var _arg0 *C.GtkFontChooser       // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(fontDesc)))

	C.gtk_font_chooser_set_font_desc(_arg0, _arg1)
	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(fontDesc)
}

// SetPreviewText sets the text displayed in the preview area. The text is used
// to show how the selected font looks.
//
// The function takes the following parameters:
//
//    - text to display in the preview area.
//
func (fontchooser *FontChooser) SetPreviewText(text string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_preview_text(_arg0, _arg1)
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
	var _arg0 *C.GtkFontChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if showPreviewEntry {
		_arg1 = C.TRUE
	}

	C.gtk_font_chooser_set_show_preview_entry(_arg0, _arg1)
	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(showPreviewEntry)
}

// fontFace gets the FontFace representing the selected font group details (i.e.
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
func (fontchooser *FontChooser) fontFace() pango.FontFacer {
	gclass := (*C.GtkFontChooserIface)(coreglib.PeekParentClass(fontchooser))
	fnarg := gclass.get_font_face

	var _arg0 *C.GtkFontChooser // out
	var _cret *C.PangoFontFace  // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C._gotk4_gtk3_FontChooser_virtual_get_font_face(unsafe.Pointer(fnarg), _arg0)
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

// fontFamily gets the FontFamily representing the selected font family. Font
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
func (fontchooser *FontChooser) fontFamily() pango.FontFamilier {
	gclass := (*C.GtkFontChooserIface)(coreglib.PeekParentClass(fontchooser))
	fnarg := gclass.get_font_family

	var _arg0 *C.GtkFontChooser  // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C._gotk4_gtk3_FontChooser_virtual_get_font_family(unsafe.Pointer(fnarg), _arg0)
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

// fontSize: selected font size.
//
// The function returns the following values:
//
//    - gint: n integer representing the selected font size, or -1 if no font
//      size is selected.
//
func (fontchooser *FontChooser) fontSize() int {
	gclass := (*C.GtkFontChooserIface)(coreglib.PeekParentClass(fontchooser))
	fnarg := gclass.get_font_size

	var _arg0 *C.GtkFontChooser // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))

	_cret = C._gotk4_gtk3_FontChooser_virtual_get_font_size(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(fontchooser)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// setFilterFunc adds a filter function that decides which fonts to display in
// the font chooser.
//
// The function takes the following parameters:
//
//    - filter (optional) or NULL.
//
func (fontchooser *FontChooser) setFilterFunc(filter FontFilterFunc) {
	gclass := (*C.GtkFontChooserIface)(coreglib.PeekParentClass(fontchooser))
	fnarg := gclass.set_filter_func

	var _arg0 *C.GtkFontChooser   // out
	var _arg1 C.GtkFontFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(coreglib.InternObject(fontchooser).Native()))
	if filter != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_FontFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(filter))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C._gotk4_gtk3_FontChooser_virtual_set_filter_func(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fontchooser)
	runtime.KeepAlive(filter)
}
