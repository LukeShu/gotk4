// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
type PathPriorityType int

const (
	// PathPriorityTypeLowest: deprecated
	PathPriorityTypeLowest PathPriorityType = 0
	// PathPriorityTypeGTK: deprecated
	PathPriorityTypeGTK PathPriorityType = 4
	// PathPriorityTypeApplication: deprecated
	PathPriorityTypeApplication PathPriorityType = 8
	// PathPriorityTypeTheme: deprecated
	PathPriorityTypeTheme PathPriorityType = 10
	// PathPriorityTypeRC: deprecated
	PathPriorityTypeRC PathPriorityType = 12
	// PathPriorityTypeHighest: deprecated
	PathPriorityTypeHighest PathPriorityType = 15
)

func marshalPathPriorityType(p uintptr) (interface{}, error) {
	return PathPriorityType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PathType: widget path types. See also gtk_binding_set_add_path().
type PathType int

const (
	// PathTypeWidget: deprecated
	PathTypeWidget PathType = 0
	// PathTypeWidgetClass: deprecated
	PathTypeWidgetClass PathType = 1
	// PathTypeClass: deprecated
	PathTypeClass PathType = 2
)

func marshalPathType(p uintptr) (interface{}, error) {
	return PathType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RCTokenType: the RcTokenType enumeration represents the tokens in the RC
// file. It is exposed so that theme engines can reuse these tokens when parsing
// the theme-engine specific portions of a RC file.
type RCTokenType int

const (
	// RCTokenTypeInvalid: deprecated
	RCTokenTypeInvalid RCTokenType = 270
	// RCTokenTypeInclude: deprecated
	RCTokenTypeInclude RCTokenType = 271
	// RCTokenTypeNormal: deprecated
	RCTokenTypeNormal RCTokenType = 272
	// RCTokenTypeActive: deprecated
	RCTokenTypeActive RCTokenType = 273
	// RCTokenTypePrelight: deprecated
	RCTokenTypePrelight RCTokenType = 274
	// RCTokenTypeSelected: deprecated
	RCTokenTypeSelected RCTokenType = 275
	// RCTokenTypeInsensitive: deprecated
	RCTokenTypeInsensitive RCTokenType = 276
	// RCTokenTypeFg: deprecated
	RCTokenTypeFg RCTokenType = 277
	// RCTokenTypeBg: deprecated
	RCTokenTypeBg RCTokenType = 278
	// RCTokenTypeText: deprecated
	RCTokenTypeText RCTokenType = 279
	// RCTokenTypeBase: deprecated
	RCTokenTypeBase RCTokenType = 280
	// RCTokenTypeXthickness: deprecated
	RCTokenTypeXthickness RCTokenType = 281
	// RCTokenTypeYthickness: deprecated
	RCTokenTypeYthickness RCTokenType = 282
	// RCTokenTypeFont: deprecated
	RCTokenTypeFont RCTokenType = 283
	// RCTokenTypeFontset: deprecated
	RCTokenTypeFontset RCTokenType = 284
	// RCTokenTypeFontName: deprecated
	RCTokenTypeFontName RCTokenType = 285
	// RCTokenTypeBgPixmap: deprecated
	RCTokenTypeBgPixmap RCTokenType = 286
	// RCTokenTypePixmapPath: deprecated
	RCTokenTypePixmapPath RCTokenType = 287
	// RCTokenTypeStyle: deprecated
	RCTokenTypeStyle RCTokenType = 288
	// RCTokenTypeBinding: deprecated
	RCTokenTypeBinding RCTokenType = 289
	// RCTokenTypeBind: deprecated
	RCTokenTypeBind RCTokenType = 290
	// RCTokenTypeWidget: deprecated
	RCTokenTypeWidget RCTokenType = 291
	// RCTokenTypeWidgetClass: deprecated
	RCTokenTypeWidgetClass RCTokenType = 292
	// RCTokenTypeClass: deprecated
	RCTokenTypeClass RCTokenType = 293
	// RCTokenTypeLowest: deprecated
	RCTokenTypeLowest RCTokenType = 294
	// RCTokenTypeGTK: deprecated
	RCTokenTypeGTK RCTokenType = 295
	// RCTokenTypeApplication: deprecated
	RCTokenTypeApplication RCTokenType = 296
	// RCTokenTypeTheme: deprecated
	RCTokenTypeTheme RCTokenType = 297
	// RCTokenTypeRC: deprecated
	RCTokenTypeRC RCTokenType = 298
	// RCTokenTypeHighest: deprecated
	RCTokenTypeHighest RCTokenType = 299
	// RCTokenTypeEngine: deprecated
	RCTokenTypeEngine RCTokenType = 300
	// RCTokenTypeModulePath: deprecated
	RCTokenTypeModulePath RCTokenType = 301
	// RCTokenTypeImModulePath: deprecated
	RCTokenTypeImModulePath RCTokenType = 302
	// RCTokenTypeImModuleFile: deprecated
	RCTokenTypeImModuleFile RCTokenType = 303
	// RCTokenTypeStock: deprecated
	RCTokenTypeStock RCTokenType = 304
	// RCTokenTypeLTR: deprecated
	RCTokenTypeLTR RCTokenType = 305
	// RCTokenTypeRTL: deprecated
	RCTokenTypeRTL RCTokenType = 306
	// RCTokenTypeColor: deprecated
	RCTokenTypeColor RCTokenType = 307
	// RCTokenTypeUnbind: deprecated
	RCTokenTypeUnbind RCTokenType = 308
	// RCTokenTypeLast: deprecated
	RCTokenTypeLast RCTokenType = 309
)

func marshalRCTokenType(p uintptr) (interface{}, error) {
	return RCTokenType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RCFlags: deprecated
type RCFlags int

const (
	// RCFlagsFg: deprecated
	RCFlagsFg RCFlags = 1
	// RCFlagsBg: deprecated
	RCFlagsBg RCFlags = 2
	// RCFlagsText: deprecated
	RCFlagsText RCFlags = 4
	// RCFlagsBase: deprecated
	RCFlagsBase RCFlags = 8
)

func marshalRCFlags(p uintptr) (interface{}, error) {
	return RCFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RCAddDefaultFile adds a file to the list of files to be parsed at the end of
// gtk_init().
func RCAddDefaultFile(filename string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_add_default_file(_arg1)
}

// RCFindModuleInPath searches for a theme engine in the GTK+ search path. This
// function is not useful for applications and should not be used.
func RCFindModuleInPath(moduleFile string) string {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(moduleFile))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.gtk_rc_find_module_in_path(_arg1)

	var _filename string // out

	_filename = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetDefaultFiles retrieves the current list of RC files that will be parsed
// at the end of gtk_init().
func RCGetDefaultFiles() []string {
	var _cret **C.gchar

	_cret = C.gtk_rc_get_default_files()

	var _filenames []string

	{
		var length int
		for p := _cret; *p != nil; p = (**C.gchar)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		src := unsafe.Slice(_cret, length)
		_filenames = make([]string, length)
		for i := range src {
			_filenames[i] = C.GoString(src[i])
		}
	}

	return _filenames
}

// RCGetImModuleFile obtains the path to the IM modules file. See the
// documentation of the `GTK_IM_MODULE_FILE` environment variable for more
// details.
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
func RCGetStyle(widget Widget) Style {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret *C.GtkStyle // in

	_cret = C.gtk_rc_get_style(_arg1)

	var _style Style // out

	_style = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Style)

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
func RCGetStyleByPaths(settings Settings, widgetPath string, classPath string, typ externglib.Type) Style {
	var _arg1 *C.GtkSettings // out
	var _arg2 *C.char        // out
	var _arg3 *C.char        // out
	var _arg4 C.GType        // out

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = (*C.char)(C.CString(widgetPath))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(C.CString(classPath))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(typ)

	var _cret *C.GtkStyle // in

	_cret = C.gtk_rc_get_style_by_paths(_arg1, _arg2, _arg3, _arg4)

	var _style Style // out

	_style = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Style)

	return _style
}

// RCGetThemeDir returns the standard directory in which themes should be
// installed. (GTK+ does not actually use this directory itself.)
func RCGetThemeDir() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_theme_dir()

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RCParse parses a given resource file.
func RCParse(filename string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse(_arg1)
}

// RCParseString parses resource information directly from a string.
func RCParseString(rcString string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(rcString))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse_string(_arg1)
}

// RCReparseAll: if the modification time on any previously read file for the
// default Settings has changed, discard all style information and then reread
// all previously read RC files.
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
func RCReparseAllForSettings(settings Settings, forceLoad bool) bool {
	var _arg1 *C.GtkSettings // out
	var _arg2 C.gboolean     // out

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	if forceLoad {
		_arg2 = C.TRUE
	}

	var _cret C.gboolean // in

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
func RCResetStyles(settings Settings) {
	var _arg1 *C.GtkSettings // out

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))

	C.gtk_rc_reset_styles(_arg1)
}

// RCSetDefaultFiles sets the list of files that GTK+ will read at the end of
// gtk_init().
func RCSetDefaultFiles(filenames []string) {
	var _arg1 **C.gchar

	_arg1 = (**C.gchar)(C.malloc(C.ulong((len(filenames) + 1)) * C.ulong(unsafe.Sizeof(uint(0)))))
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

// RCStyle: the RcStyle-struct is used to represent a set of information about
// the appearance of a widget. This can later be composited together with other
// RcStyle-struct<!-- -->s to form a Style.
type RCStyle interface {
	gextras.Objector

	// Copy makes a copy of the specified RcStyle. This function will correctly
	// copy an RC style that is a member of a class derived from RcStyle.
	Copy() RCStyle
}

// rcStyle implements the RCStyle class.
type rcStyle struct {
	gextras.Objector
}

var _ RCStyle = (*rcStyle)(nil)

// WrapRCStyle wraps a GObject to the right type. It is
// primarily used internally.
func WrapRCStyle(obj *externglib.Object) RCStyle {
	return rcStyle{
		Objector: obj,
	}
}

func marshalRCStyle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRCStyle(obj), nil
}

// NewRCStyle constructs a class RCStyle.
func NewRCStyle() RCStyle {
	var _cret C.GtkRcStyle // in

	_cret = C.gtk_rc_style_new()

	var _rcStyle RCStyle // out

	_rcStyle = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(RCStyle)

	return _rcStyle
}

// Copy makes a copy of the specified RcStyle. This function will correctly
// copy an RC style that is a member of a class derived from RcStyle.
func (o rcStyle) Copy() RCStyle {
	var _arg0 *C.GtkRcStyle // out

	_arg0 = (*C.GtkRcStyle)(unsafe.Pointer(o.Native()))

	var _cret *C.GtkRcStyle // in

	_cret = C.gtk_rc_style_copy(_arg0)

	var _rcStyle RCStyle // out

	_rcStyle = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(RCStyle)

	return _rcStyle
}

type RCContext struct {
	native C.GtkRcContext
}

// WrapRCContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRCContext(ptr unsafe.Pointer) *RCContext {
	if ptr == nil {
		return nil
	}

	return (*RCContext)(ptr)
}

// Native returns the underlying C source pointer.
func (r *RCContext) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// RCProperty: deprecated
type RCProperty struct {
	native C.GtkRcProperty
}

// WrapRCProperty wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRCProperty(ptr unsafe.Pointer) *RCProperty {
	if ptr == nil {
		return nil
	}

	return (*RCProperty)(ptr)
}

// Native returns the underlying C source pointer.
func (r *RCProperty) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Origin gets the field inside the struct.
func (r *RCProperty) Origin() string {
	var v string // out
	v = C.GoString(r.native.origin)
	return v
}
