// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_path_priority_type_get_type()), F: marshalPathPriorityType},
		{T: externglib.Type(C.gtk_path_type_get_type()), F: marshalPathType},
		{T: externglib.Type(C.gtk_rc_token_type_get_type()), F: marshalRCTokenType},
		{T: externglib.Type(C.gtk_rc_flags_get_type()), F: marshalRCFlags},
		{T: externglib.Type(C.gtk_rc_style_get_type()), F: marshalRCStyle},
	})
}

// PathPriorityType priorities for path lookups. See also
// gtk_binding_set_add_path().
//
// Deprecated: since version 3.0.
type PathPriorityType int

const (
	// Lowest: deprecated
	PathPriorityTypeLowest PathPriorityType = 0
	// GTK: deprecated
	PathPriorityTypeGTK PathPriorityType = 4
	// Application: deprecated
	PathPriorityTypeApplication PathPriorityType = 8
	// Theme: deprecated
	PathPriorityTypeTheme PathPriorityType = 10
	// RC: deprecated
	PathPriorityTypeRC PathPriorityType = 12
	// Highest: deprecated
	PathPriorityTypeHighest PathPriorityType = 15
)

func marshalPathPriorityType(p uintptr) (interface{}, error) {
	return PathPriorityType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PathType: widget path types. See also gtk_binding_set_add_path().
//
// Deprecated: since version 3.0.
type PathType int

const (
	// Widget: deprecated
	PathTypeWidget PathType = iota
	// WidgetClass: deprecated
	PathTypeWidgetClass
	// Class: deprecated
	PathTypeClass
)

func marshalPathType(p uintptr) (interface{}, error) {
	return PathType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RCTokenType: the RcTokenType enumeration represents the tokens in the RC
// file. It is exposed so that theme engines can reuse these tokens when parsing
// the theme-engine specific portions of a RC file.
//
// Deprecated: since version 3.0.
type RCTokenType int

const (
	// Invalid: deprecated
	RCTokenTypeInvalid RCTokenType = 270
	// Include: deprecated
	RCTokenTypeInclude RCTokenType = 271
	// Normal: deprecated
	RCTokenTypeNormal RCTokenType = 272
	// Active: deprecated
	RCTokenTypeActive RCTokenType = 273
	// Prelight: deprecated
	RCTokenTypePrelight RCTokenType = 274
	// Selected: deprecated
	RCTokenTypeSelected RCTokenType = 275
	// Insensitive: deprecated
	RCTokenTypeInsensitive RCTokenType = 276
	// Fg: deprecated
	RCTokenTypeFg RCTokenType = 277
	// Bg: deprecated
	RCTokenTypeBg RCTokenType = 278
	// Text: deprecated
	RCTokenTypeText RCTokenType = 279
	// Base: deprecated
	RCTokenTypeBase RCTokenType = 280
	// Xthickness: deprecated
	RCTokenTypeXthickness RCTokenType = 281
	// Ythickness: deprecated
	RCTokenTypeYthickness RCTokenType = 282
	// Font: deprecated
	RCTokenTypeFont RCTokenType = 283
	// Fontset: deprecated
	RCTokenTypeFontset RCTokenType = 284
	// FontName: deprecated
	RCTokenTypeFontName RCTokenType = 285
	// BgPixmap: deprecated
	RCTokenTypeBgPixmap RCTokenType = 286
	// PixmapPath: deprecated
	RCTokenTypePixmapPath RCTokenType = 287
	// Style: deprecated
	RCTokenTypeStyle RCTokenType = 288
	// Binding: deprecated
	RCTokenTypeBinding RCTokenType = 289
	// Bind: deprecated
	RCTokenTypeBind RCTokenType = 290
	// Widget: deprecated
	RCTokenTypeWidget RCTokenType = 291
	// WidgetClass: deprecated
	RCTokenTypeWidgetClass RCTokenType = 292
	// Class: deprecated
	RCTokenTypeClass RCTokenType = 293
	// Lowest: deprecated
	RCTokenTypeLowest RCTokenType = 294
	// GTK: deprecated
	RCTokenTypeGTK RCTokenType = 295
	// Application: deprecated
	RCTokenTypeApplication RCTokenType = 296
	// Theme: deprecated
	RCTokenTypeTheme RCTokenType = 297
	// RC: deprecated
	RCTokenTypeRC RCTokenType = 298
	// Highest: deprecated
	RCTokenTypeHighest RCTokenType = 299
	// Engine: deprecated
	RCTokenTypeEngine RCTokenType = 300
	// ModulePath: deprecated
	RCTokenTypeModulePath RCTokenType = 301
	// ImModulePath: deprecated
	RCTokenTypeImModulePath RCTokenType = 302
	// ImModuleFile: deprecated
	RCTokenTypeImModuleFile RCTokenType = 303
	// Stock: deprecated
	RCTokenTypeStock RCTokenType = 304
	// LTR: deprecated
	RCTokenTypeLTR RCTokenType = 305
	// RTL: deprecated
	RCTokenTypeRTL RCTokenType = 306
	// Color: deprecated
	RCTokenTypeColor RCTokenType = 307
	// Unbind: deprecated
	RCTokenTypeUnbind RCTokenType = 308
	// Last: deprecated
	RCTokenTypeLast RCTokenType = 309
)

func marshalRCTokenType(p uintptr) (interface{}, error) {
	return RCTokenType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RCFlags: deprecated
type RCFlags int

const (
	// RCFlagsFg: deprecated
	RCFlagsFg RCFlags = 0b1
	// RCFlagsBg: deprecated
	RCFlagsBg RCFlags = 0b10
	// RCFlagsText: deprecated
	RCFlagsText RCFlags = 0b100
	// RCFlagsBase: deprecated
	RCFlagsBase RCFlags = 0b1000
)

func marshalRCFlags(p uintptr) (interface{}, error) {
	return RCFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RCAddDefaultFile adds a file to the list of files to be parsed at the end of
// gtk_init().
//
// Deprecated: since version 3.0.
func RCAddDefaultFile(filename string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_add_default_file(_arg1)
}

// RCFindModuleInPath searches for a theme engine in the GTK+ search path. This
// function is not useful for applications and should not be used.
//
// Deprecated: since version 3.0.
func RCFindModuleInPath(moduleFile string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(C.CString(moduleFile))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_rc_find_module_in_path(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCFindPixmapInPath looks up a file in pixmap path for the specified Settings.
// If the file is not found, it outputs a warning message using g_warning() and
// returns nil.
//
// Deprecated: since version 3.0.
func RCFindPixmapInPath(settings Settings, scanner *glib.Scanner, pixmapFile string) string {
	var _arg1 *C.GtkSettings // out
	var _arg2 *C.GScanner    // out
	var _arg3 *C.gchar       // out
	var _cret *C.gchar       // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg3 = (*C.gchar)(C.CString(pixmapFile))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_rc_find_pixmap_in_path(_arg1, _arg2, _arg3)

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetDefaultFiles retrieves the current list of RC files that will be parsed
// at the end of gtk_init().
//
// Deprecated: since version 3.0.
func RCGetDefaultFiles() []string {
	var _cret **C.gchar

	_cret = C.gtk_rc_get_default_files()

	var _filenames []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString(src[i])
		}
	}

	return _filenames
}

// RCGetImModuleFile obtains the path to the IM modules file. See the
// documentation of the `GTK_IM_MODULE_FILE` environment variable for more
// details.
//
// Deprecated: since version 3.0.
func RCGetImModuleFile() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_im_module_file()

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetImModulePath obtains the path in which to look for IM modules. See the
// documentation of the `GTK_PATH` environment variable for more details about
// looking up modules. This function is useful solely for utilities supplied
// with GTK+ and should not be used by applications under normal circumstances.
//
// Deprecated: since version 3.0.
func RCGetImModulePath() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_im_module_path()

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetModuleDir returns a directory in which GTK+ looks for theme engines. For
// full information about the search for theme engines, see the docs for
// `GTK_PATH` in [Running GTK+ Applications][gtk-running].
//
// Deprecated: since version 3.0.
func RCGetModuleDir() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_module_dir()

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetStyle finds all matching RC styles for a given widget, composites them
// together, and then creates a Style representing the composite appearance.
// (GTK+ actually keeps a cache of previously created styles, so a new style may
// not be created.)
//
// Deprecated: since version 3.0.
func RCGetStyle(widget Widget) Style {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkStyle  // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_rc_get_style(_arg1)

	var _style Style // out

	_style = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Style)

	return _style
}

// RCGetStyleByPaths creates up a Style from styles defined in a RC file by
// providing the raw components used in matching. This function may be useful
// when creating pseudo-widgets that should be themed like widgets but don’t
// actually have corresponding GTK+ widgets. An example of this would be items
// inside a GNOME canvas widget.
//
// The action of gtk_rc_get_style() is similar to:
//
//    gtk_widget_path (widget, NULL, &path, NULL);
//    gtk_widget_class_path (widget, NULL, &class_path, NULL);
//    gtk_rc_get_style_by_paths (gtk_widget_get_settings (widget),
//                               path, class_path,
//                               G_OBJECT_TYPE (widget));
//
// Deprecated: since version 3.0.
func RCGetStyleByPaths(settings Settings, widgetPath string, classPath string, typ externglib.Type) Style {
	var _arg1 *C.GtkSettings // out
	var _arg2 *C.char        // out
	var _arg3 *C.char        // out
	var _arg4 C.GType        // out
	var _cret *C.GtkStyle    // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = (*C.char)(C.CString(widgetPath))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(C.CString(classPath))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (C.GType)(typ)

	_cret = C.gtk_rc_get_style_by_paths(_arg1, _arg2, _arg3, _arg4)

	var _style Style // out

	_style = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Style)

	return _style
}

// RCGetThemeDir returns the standard directory in which themes should be
// installed. (GTK+ does not actually use this directory itself.)
//
// Deprecated: since version 3.0.
func RCGetThemeDir() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_theme_dir()

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RCParse parses a given resource file.
//
// Deprecated: since version 3.0.
func RCParse(filename string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse(_arg1)
}

// RCParsePriority parses a PathPriorityType variable from the format expected
// in a RC file.
//
// Deprecated: since version 3.0.
func RCParsePriority(scanner *glib.Scanner, priority *PathPriorityType) uint {
	var _arg1 *C.GScanner            // out
	var _arg2 *C.GtkPathPriorityType // out
	var _cret C.guint                // in

	_arg1 = (*C.GScanner)(unsafe.Pointer(scanner))
	_arg2 = (*C.GtkPathPriorityType)(unsafe.Pointer(priority))

	_cret = C.gtk_rc_parse_priority(_arg1, _arg2)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RCParseState parses a StateType variable from the format expected in a RC
// file.
//
// Deprecated: since version 3.0.
func RCParseState(scanner *glib.Scanner) (StateType, uint) {
	var _arg1 *C.GScanner    // out
	var _arg2 C.GtkStateType // in
	var _cret C.guint        // in

	_arg1 = (*C.GScanner)(unsafe.Pointer(scanner))

	_cret = C.gtk_rc_parse_state(_arg1, &_arg2)

	var _state StateType // out
	var _guint uint      // out

	_state = StateType(_arg2)
	_guint = uint(_cret)

	return _state, _guint
}

// RCParseString parses resource information directly from a string.
//
// Deprecated: since version 3.0.
func RCParseString(rcString string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(rcString))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse_string(_arg1)
}

// RCReparseAll: if the modification time on any previously read file for the
// default Settings has changed, discard all style information and then reread
// all previously read RC files.
//
// Deprecated: since version 3.0.
func RCReparseAll() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_rc_reparse_all()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RCReparseAllForSettings: if the modification time on any previously read file
// for the given Settings has changed, discard all style information and then
// reread all previously read RC files.
//
// Deprecated: since version 3.0.
func RCReparseAllForSettings(settings Settings, forceLoad bool) bool {
	var _arg1 *C.GtkSettings // out
	var _arg2 C.gboolean     // out
	var _cret C.gboolean     // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	if forceLoad {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_rc_reparse_all_for_settings(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RCResetStyles: this function recomputes the styles for all widgets that use a
// particular Settings object. (There is one Settings object per Screen, see
// gtk_settings_get_for_screen()); It is useful when some global parameter has
// changed that affects the appearance of all widgets, because when a widget
// gets a new style, it will both redraw and recompute any cached information
// about its appearance. As an example, it is used when the default font size
// set by the operating system changes. Note that this function doesn’t affect
// widgets that have a style set explicitly on them with gtk_widget_set_style().
//
// Deprecated: since version 3.0.
func RCResetStyles(settings Settings) {
	var _arg1 *C.GtkSettings // out

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))

	C.gtk_rc_reset_styles(_arg1)
}

// RCSetDefaultFiles sets the list of files that GTK+ will read at the end of
// gtk_init().
//
// Deprecated: since version 3.0.
func RCSetDefaultFiles(filenames []string) {
	var _arg1 **C.gchar

	_arg1 = (**C.gchar)(C.malloc(C.ulong(len(filenames)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(filenames))
		for i := range filenames {
			out[i] = (*C.gchar)(C.CString(filenames[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_rc_set_default_files(_arg1)
}

// RCStyleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RCStyleOverrider interface {
	Merge(src RCStyle)
	Parse(settings Settings, scanner *glib.Scanner) uint
}

// RCStyle: the RcStyle-struct is used to represent a set of information about
// the appearance of a widget. This can later be composited together with other
// RcStyle-struct<!-- -->s to form a Style.
type RCStyle interface {
	gextras.Objector

	// Copy makes a copy of the specified RcStyle. This function will correctly
	// copy an RC style that is a member of a class derived from RcStyle.
	//
	// Deprecated: since version 3.0.
	Copy() RCStyle
}

// RCStyleClass implements the RCStyle interface.
type RCStyleClass struct {
	*externglib.Object
}

var _ RCStyle = (*RCStyleClass)(nil)

func wrapRCStyle(obj *externglib.Object) RCStyle {
	return &RCStyleClass{
		Object: obj,
	}
}

func marshalRCStyle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRCStyle(obj), nil
}

// NewRCStyle creates a new RcStyle with no fields set and a reference count of
// 1.
//
// Deprecated: since version 3.0.
func NewRCStyle() RCStyle {
	var _cret *C.GtkRcStyle // in

	_cret = C.gtk_rc_style_new()

	var _rcStyle RCStyle // out

	_rcStyle = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RCStyle)

	return _rcStyle
}

// Copy makes a copy of the specified RcStyle. This function will correctly copy
// an RC style that is a member of a class derived from RcStyle.
//
// Deprecated: since version 3.0.
func (o *RCStyleClass) Copy() RCStyle {
	var _arg0 *C.GtkRcStyle // out
	var _cret *C.GtkRcStyle // in

	_arg0 = (*C.GtkRcStyle)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_rc_style_copy(_arg0)

	var _rcStyle RCStyle // out

	_rcStyle = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RCStyle)

	return _rcStyle
}
