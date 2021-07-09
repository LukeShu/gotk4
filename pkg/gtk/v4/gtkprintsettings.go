// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_PrintSettingsFunc(char*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_settings_get_type()), F: marshalPrintSettings},
	})
}

type PrintSettingsFunc func(key string, value string, userData interface{})

//export gotk4_PrintSettingsFunc
func gotk4_PrintSettingsFunc(arg0 *C.char, arg1 *C.char, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var key string           // out
	var value string         // out
	var userData interface{} // out

	key = C.GoString(arg0)
	value = C.GoString(arg1)
	userData = box.Get(uintptr(arg2))

	fn := v.(PrintSettingsFunc)
	fn(key, value, userData)
}

// PrintSettings: `GtkPrintSettings` object represents the settings of a print
// dialog in a system-independent way.
//
// The main use for this object is that once you’ve printed you can get a
// settings object that represents the settings the user chose, and the next
// time you print you can pass that object in so that the user doesn’t have to
// re-set all his settings.
//
// Its also possible to enumerate the settings so that you can easily save the
// settings for the next time your app runs, or even store them in a document.
// The predefined keys try to use shared values as much as possible so that
// moving such a document between systems still works.
type PrintSettings interface {
	gextras.Objector

	// Copy copies a `GtkPrintSettings` object.
	Copy() *PrintSettingsClass
	// Foreach calls @func for each key-value pair of @settings.
	Foreach(fn PrintSettingsFunc)
	// Get looks up the string value associated with @key.
	Get(key string) string
	// Bool returns the boolean represented by the value that is associated with
	// @key.
	//
	// The string “true” represents true, any other string false.
	Bool(key string) bool
	// Collate gets the value of GTK_PRINT_SETTINGS_COLLATE.
	Collate() bool
	// DefaultSource gets the value of GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
	DefaultSource() string
	// Dither gets the value of GTK_PRINT_SETTINGS_DITHER.
	Dither() string
	// Double returns the double value associated with @key, or 0.
	Double(key string) float64
	// DoubleWithDefault returns the floating point number represented by the
	// value that is associated with @key, or @default_val if the value does not
	// represent a floating point number.
	//
	// Floating point numbers are parsed with g_ascii_strtod().
	DoubleWithDefault(key string, def float64) float64
	// Duplex gets the value of GTK_PRINT_SETTINGS_DUPLEX.
	Duplex() PrintDuplex
	// Finishings gets the value of GTK_PRINT_SETTINGS_FINISHINGS.
	Finishings() string
	// Int returns the integer value of @key, or 0.
	Int(key string) int
	// IntWithDefault returns the value of @key, interpreted as an integer, or
	// the default value.
	IntWithDefault(key string, def int) int
	// MediaType gets the value of GTK_PRINT_SETTINGS_MEDIA_TYPE.
	//
	// The set of media types is defined in PWG 5101.1-2002 PWG.
	MediaType() string
	// NCopies gets the value of GTK_PRINT_SETTINGS_N_COPIES.
	NCopies() int
	// NumberUp gets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
	NumberUp() int
	// NumberUpLayout gets the value of GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
	NumberUpLayout() NumberUpLayout
	// Orientation: get the value of GTK_PRINT_SETTINGS_ORIENTATION, converted
	// to a `GtkPageOrientation`.
	Orientation() PageOrientation
	// OutputBin gets the value of GTK_PRINT_SETTINGS_OUTPUT_BIN.
	OutputBin() string
	// PageSet gets the value of GTK_PRINT_SETTINGS_PAGE_SET.
	PageSet() PageSet
	// PaperSize gets the value of GTK_PRINT_SETTINGS_PAPER_FORMAT, converted to
	// a `GtkPaperSize`.
	PaperSize() *PaperSize
	// PrintPages gets the value of GTK_PRINT_SETTINGS_PRINT_PAGES.
	PrintPages() PrintPages
	// Printer: convenience function to obtain the value of
	// GTK_PRINT_SETTINGS_PRINTER.
	Printer() string
	// PrinterLpi gets the value of GTK_PRINT_SETTINGS_PRINTER_LPI.
	PrinterLpi() float64
	// Quality gets the value of GTK_PRINT_SETTINGS_QUALITY.
	Quality() PrintQuality
	// Resolution gets the value of GTK_PRINT_SETTINGS_RESOLUTION.
	Resolution() int
	// ResolutionX gets the value of GTK_PRINT_SETTINGS_RESOLUTION_X.
	ResolutionX() int
	// ResolutionY gets the value of GTK_PRINT_SETTINGS_RESOLUTION_Y.
	ResolutionY() int
	// Reverse gets the value of GTK_PRINT_SETTINGS_REVERSE.
	Reverse() bool
	// Scale gets the value of GTK_PRINT_SETTINGS_SCALE.
	Scale() float64
	// UseColor gets the value of GTK_PRINT_SETTINGS_USE_COLOR.
	UseColor() bool
	// HasKey returns true, if a value is associated with @key.
	HasKey(key string) bool
	// LoadFile reads the print settings from @file_name.
	//
	// If the file could not be loaded then error is set to either a
	// `GFileError` or `GKeyFileError`.
	//
	// See [method@Gtk.PrintSettings.to_file].
	LoadFile(fileName string) error
	// LoadKeyFile reads the print settings from the group @group_name in
	// @key_file.
	//
	// If the file could not be loaded then error is set to either a
	// `GFileError` or `GKeyFileError`.
	LoadKeyFile(keyFile *glib.KeyFile, groupName string) error
	// Set associates @value with @key.
	Set(key string, value string)
	// SetBool sets @key to a boolean value.
	SetBool(key string, value bool)
	// SetCollate sets the value of GTK_PRINT_SETTINGS_COLLATE.
	SetCollate(collate bool)
	// SetDefaultSource sets the value of GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
	SetDefaultSource(defaultSource string)
	// SetDither sets the value of GTK_PRINT_SETTINGS_DITHER.
	SetDither(dither string)
	// SetDouble sets @key to a double value.
	SetDouble(key string, value float64)
	// SetFinishings sets the value of GTK_PRINT_SETTINGS_FINISHINGS.
	SetFinishings(finishings string)
	// SetInt sets @key to an integer value.
	SetInt(key string, value int)
	// SetMediaType sets the value of GTK_PRINT_SETTINGS_MEDIA_TYPE.
	//
	// The set of media types is defined in PWG 5101.1-2002 PWG.
	SetMediaType(mediaType string)
	// SetNCopies sets the value of GTK_PRINT_SETTINGS_N_COPIES.
	SetNCopies(numCopies int)
	// SetNumberUp sets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
	SetNumberUp(numberUp int)
	// SetOutputBin sets the value of GTK_PRINT_SETTINGS_OUTPUT_BIN.
	SetOutputBin(outputBin string)
	// SetPaperSize sets the value of GTK_PRINT_SETTINGS_PAPER_FORMAT,
	// GTK_PRINT_SETTINGS_PAPER_WIDTH and GTK_PRINT_SETTINGS_PAPER_HEIGHT.
	SetPaperSize(paperSize *PaperSize)
	// SetPrinter: convenience function to set GTK_PRINT_SETTINGS_PRINTER to
	// @printer.
	SetPrinter(printer string)
	// SetPrinterLpi sets the value of GTK_PRINT_SETTINGS_PRINTER_LPI.
	SetPrinterLpi(lpi float64)
	// SetResolution sets the values of GTK_PRINT_SETTINGS_RESOLUTION,
	// GTK_PRINT_SETTINGS_RESOLUTION_X and GTK_PRINT_SETTINGS_RESOLUTION_Y.
	SetResolution(resolution int)
	// SetResolutionXY sets the values of GTK_PRINT_SETTINGS_RESOLUTION,
	// GTK_PRINT_SETTINGS_RESOLUTION_X and GTK_PRINT_SETTINGS_RESOLUTION_Y.
	SetResolutionXY(resolutionX int, resolutionY int)
	// SetReverse sets the value of GTK_PRINT_SETTINGS_REVERSE.
	SetReverse(reverse bool)
	// SetScale sets the value of GTK_PRINT_SETTINGS_SCALE.
	SetScale(scale float64)
	// SetUseColor sets the value of GTK_PRINT_SETTINGS_USE_COLOR.
	SetUseColor(useColor bool)
	// ToFile: this function saves the print settings from @settings to
	// @file_name.
	//
	// If the file could not be written then error is set to either a
	// `GFileError` or `GKeyFileError`.
	ToFile(fileName string) error
	// ToGVariant: serialize print settings to an a{sv} variant.
	ToGVariant() *glib.Variant
	// ToKeyFile: this function adds the print settings from @settings to
	// @key_file.
	ToKeyFile(keyFile *glib.KeyFile, groupName string)
	// Unset removes any value associated with @key.
	//
	// This has the same effect as setting the value to nil.
	Unset(key string)
}

// PrintSettingsClass implements the PrintSettings interface.
type PrintSettingsClass struct {
	*externglib.Object
}

var _ PrintSettings = (*PrintSettingsClass)(nil)

func wrapPrintSettings(obj *externglib.Object) PrintSettings {
	return &PrintSettingsClass{
		Object: obj,
	}
}

func marshalPrintSettings(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPrintSettings(obj), nil
}

// NewPrintSettings creates a new `GtkPrintSettings` object.
func NewPrintSettings() *PrintSettingsClass {
	var _cret *C.GtkPrintSettings // in

	_cret = C.gtk_print_settings_new()

	var _printSettings *PrintSettingsClass // out

	_printSettings = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PrintSettingsClass)

	return _printSettings
}

// NewPrintSettingsFromFile reads the print settings from @file_name.
//
// Returns a new `GtkPrintSettings` object with the restored settings, or nil if
// an error occurred. If the file could not be loaded then error is set to
// either a `GFileError` or `GKeyFileError`.
//
// See [method@Gtk.PrintSettings.to_file].
func NewPrintSettingsFromFile(fileName string) (*PrintSettingsClass, error) {
	var _arg1 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_new_from_file(_arg1, &_cerr)

	var _printSettings *PrintSettingsClass // out
	var _goerr error                       // out

	_printSettings = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PrintSettingsClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _printSettings, _goerr
}

// NewPrintSettingsFromGVariant: deserialize print settings from an a{sv}
// variant.
//
// The variant must be in the format produced by
// [method@Gtk.PrintSettings.to_gvariant].
func NewPrintSettingsFromGVariant(variant *glib.Variant) *PrintSettingsClass {
	var _arg1 *C.GVariant         // out
	var _cret *C.GtkPrintSettings // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant))

	_cret = C.gtk_print_settings_new_from_gvariant(_arg1)

	var _printSettings *PrintSettingsClass // out

	_printSettings = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PrintSettingsClass)

	return _printSettings
}

// NewPrintSettingsFromKeyFile reads the print settings from the group
// @group_name in @key_file.
//
// Returns a new `GtkPrintSettings` object with the restored settings, or nil if
// an error occurred. If the file could not be loaded then error is set to
// either `GFileError` or `GKeyFileError`.
func NewPrintSettingsFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PrintSettingsClass, error) {
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_print_settings_new_from_key_file(_arg1, _arg2, &_cerr)

	var _printSettings *PrintSettingsClass // out
	var _goerr error                       // out

	_printSettings = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PrintSettingsClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _printSettings, _goerr
}

// Copy copies a `GtkPrintSettings` object.
func (o *PrintSettingsClass) Copy() *PrintSettingsClass {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPrintSettings // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_settings_copy(_arg0)

	var _printSettings *PrintSettingsClass // out

	_printSettings = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PrintSettingsClass)

	return _printSettings
}

// Foreach calls @func for each key-value pair of @settings.
func (s *PrintSettingsClass) Foreach(fn PrintSettingsFunc) {
	var _arg0 *C.GtkPrintSettings    // out
	var _arg1 C.GtkPrintSettingsFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*[0]byte)(C.gotk4_PrintSettingsFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_print_settings_foreach(_arg0, _arg1, _arg2)
}

// Get looks up the string value associated with @key.
func (s *PrintSettingsClass) Get(key string) string {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Bool returns the boolean represented by the value that is associated with
// @key.
//
// The string “true” represents true, any other string false.
func (s *PrintSettingsClass) Bool(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_bool(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Collate gets the value of GTK_PRINT_SETTINGS_COLLATE.
func (s *PrintSettingsClass) Collate() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_collate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DefaultSource gets the value of GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
func (s *PrintSettingsClass) DefaultSource() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_default_source(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Dither gets the value of GTK_PRINT_SETTINGS_DITHER.
func (s *PrintSettingsClass) Dither() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_dither(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Double returns the double value associated with @key, or 0.
func (s *PrintSettingsClass) Double(key string) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_double(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DoubleWithDefault returns the floating point number represented by the value
// that is associated with @key, or @default_val if the value does not represent
// a floating point number.
//
// Floating point numbers are parsed with g_ascii_strtod().
func (s *PrintSettingsClass) DoubleWithDefault(key string, def float64) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(def)

	_cret = C.gtk_print_settings_get_double_with_default(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Duplex gets the value of GTK_PRINT_SETTINGS_DUPLEX.
func (s *PrintSettingsClass) Duplex() PrintDuplex {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintDuplex    // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_duplex(_arg0)

	var _printDuplex PrintDuplex // out

	_printDuplex = (PrintDuplex)(_cret)

	return _printDuplex
}

// Finishings gets the value of GTK_PRINT_SETTINGS_FINISHINGS.
func (s *PrintSettingsClass) Finishings() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_finishings(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Int returns the integer value of @key, or 0.
func (s *PrintSettingsClass) Int(key string) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_int(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IntWithDefault returns the value of @key, interpreted as an integer, or the
// default value.
func (s *PrintSettingsClass) IntWithDefault(key string, def int) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(def)

	_cret = C.gtk_print_settings_get_int_with_default(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MediaType gets the value of GTK_PRINT_SETTINGS_MEDIA_TYPE.
//
// The set of media types is defined in PWG 5101.1-2002 PWG.
func (s *PrintSettingsClass) MediaType() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_media_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// NCopies gets the value of GTK_PRINT_SETTINGS_N_COPIES.
func (s *PrintSettingsClass) NCopies() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_n_copies(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NumberUp gets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
func (s *PrintSettingsClass) NumberUp() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_number_up(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NumberUpLayout gets the value of GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
func (s *PrintSettingsClass) NumberUpLayout() NumberUpLayout {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkNumberUpLayout // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_number_up_layout(_arg0)

	var _numberUpLayout NumberUpLayout // out

	_numberUpLayout = (NumberUpLayout)(_cret)

	return _numberUpLayout
}

// Orientation: get the value of GTK_PRINT_SETTINGS_ORIENTATION, converted to a
// `GtkPageOrientation`.
func (s *PrintSettingsClass) Orientation() PageOrientation {
	var _arg0 *C.GtkPrintSettings  // out
	var _cret C.GtkPageOrientation // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_orientation(_arg0)

	var _pageOrientation PageOrientation // out

	_pageOrientation = (PageOrientation)(_cret)

	return _pageOrientation
}

// OutputBin gets the value of GTK_PRINT_SETTINGS_OUTPUT_BIN.
func (s *PrintSettingsClass) OutputBin() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_output_bin(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PageSet gets the value of GTK_PRINT_SETTINGS_PAGE_SET.
func (s *PrintSettingsClass) PageSet() PageSet {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPageSet        // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_page_set(_arg0)

	var _pageSet PageSet // out

	_pageSet = (PageSet)(_cret)

	return _pageSet
}

// PaperSize gets the value of GTK_PRINT_SETTINGS_PAPER_FORMAT, converted to a
// `GtkPaperSize`.
func (s *PrintSettingsClass) PaperSize() *PaperSize {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPaperSize     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_paper_size(_arg0)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.free(unsafe.Pointer(v))
	})

	return _paperSize
}

// PrintPages gets the value of GTK_PRINT_SETTINGS_PRINT_PAGES.
func (s *PrintSettingsClass) PrintPages() PrintPages {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintPages     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_print_pages(_arg0)

	var _printPages PrintPages // out

	_printPages = (PrintPages)(_cret)

	return _printPages
}

// Printer: convenience function to obtain the value of
// GTK_PRINT_SETTINGS_PRINTER.
func (s *PrintSettingsClass) Printer() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_printer(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PrinterLpi gets the value of GTK_PRINT_SETTINGS_PRINTER_LPI.
func (s *PrintSettingsClass) PrinterLpi() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_printer_lpi(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Quality gets the value of GTK_PRINT_SETTINGS_QUALITY.
func (s *PrintSettingsClass) Quality() PrintQuality {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintQuality   // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_quality(_arg0)

	var _printQuality PrintQuality // out

	_printQuality = (PrintQuality)(_cret)

	return _printQuality
}

// Resolution gets the value of GTK_PRINT_SETTINGS_RESOLUTION.
func (s *PrintSettingsClass) Resolution() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_resolution(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ResolutionX gets the value of GTK_PRINT_SETTINGS_RESOLUTION_X.
func (s *PrintSettingsClass) ResolutionX() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_resolution_x(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ResolutionY gets the value of GTK_PRINT_SETTINGS_RESOLUTION_Y.
func (s *PrintSettingsClass) ResolutionY() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_resolution_y(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Reverse gets the value of GTK_PRINT_SETTINGS_REVERSE.
func (s *PrintSettingsClass) Reverse() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_reverse(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Scale gets the value of GTK_PRINT_SETTINGS_SCALE.
func (s *PrintSettingsClass) Scale() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_scale(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// UseColor gets the value of GTK_PRINT_SETTINGS_USE_COLOR.
func (s *PrintSettingsClass) UseColor() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_get_use_color(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasKey returns true, if a value is associated with @key.
func (s *PrintSettingsClass) HasKey(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LoadFile reads the print settings from @file_name.
//
// If the file could not be loaded then error is set to either a `GFileError` or
// `GKeyFileError`.
//
// See [method@Gtk.PrintSettings.to_file].
func (s *PrintSettingsClass) LoadFile(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_load_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// LoadKeyFile reads the print settings from the group @group_name in @key_file.
//
// If the file could not be loaded then error is set to either a `GFileError` or
// `GKeyFileError`.
func (s *PrintSettingsClass) LoadKeyFile(keyFile *glib.KeyFile, groupName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_load_key_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Set associates @value with @key.
func (s *PrintSettingsClass) Set(key string, value string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_set(_arg0, _arg1, _arg2)
}

// SetBool sets @key to a boolean value.
func (s *PrintSettingsClass) SetBool(key string, value bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	if value {
		_arg2 = C.TRUE
	}

	C.gtk_print_settings_set_bool(_arg0, _arg1, _arg2)
}

// SetCollate sets the value of GTK_PRINT_SETTINGS_COLLATE.
func (s *PrintSettingsClass) SetCollate(collate bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	if collate {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_collate(_arg0, _arg1)
}

// SetDefaultSource sets the value of GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
func (s *PrintSettingsClass) SetDefaultSource(defaultSource string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(defaultSource))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_default_source(_arg0, _arg1)
}

// SetDither sets the value of GTK_PRINT_SETTINGS_DITHER.
func (s *PrintSettingsClass) SetDither(dither string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(dither))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_dither(_arg0, _arg1)
}

// SetDouble sets @key to a double value.
func (s *PrintSettingsClass) SetDouble(key string, value float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)

	C.gtk_print_settings_set_double(_arg0, _arg1, _arg2)
}

// SetFinishings sets the value of GTK_PRINT_SETTINGS_FINISHINGS.
func (s *PrintSettingsClass) SetFinishings(finishings string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(finishings))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_finishings(_arg0, _arg1)
}

// SetInt sets @key to an integer value.
func (s *PrintSettingsClass) SetInt(key string, value int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(value)

	C.gtk_print_settings_set_int(_arg0, _arg1, _arg2)
}

// SetMediaType sets the value of GTK_PRINT_SETTINGS_MEDIA_TYPE.
//
// The set of media types is defined in PWG 5101.1-2002 PWG.
func (s *PrintSettingsClass) SetMediaType(mediaType string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(mediaType))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_media_type(_arg0, _arg1)
}

// SetNCopies sets the value of GTK_PRINT_SETTINGS_N_COPIES.
func (s *PrintSettingsClass) SetNCopies(numCopies int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = C.int(numCopies)

	C.gtk_print_settings_set_n_copies(_arg0, _arg1)
}

// SetNumberUp sets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
func (s *PrintSettingsClass) SetNumberUp(numberUp int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = C.int(numberUp)

	C.gtk_print_settings_set_number_up(_arg0, _arg1)
}

// SetOutputBin sets the value of GTK_PRINT_SETTINGS_OUTPUT_BIN.
func (s *PrintSettingsClass) SetOutputBin(outputBin string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(outputBin))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_output_bin(_arg0, _arg1)
}

// SetPaperSize sets the value of GTK_PRINT_SETTINGS_PAPER_FORMAT,
// GTK_PRINT_SETTINGS_PAPER_WIDTH and GTK_PRINT_SETTINGS_PAPER_HEIGHT.
func (s *PrintSettingsClass) SetPaperSize(paperSize *PaperSize) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GtkPaperSize     // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(paperSize))

	C.gtk_print_settings_set_paper_size(_arg0, _arg1)
}

// SetPrinter: convenience function to set GTK_PRINT_SETTINGS_PRINTER to
// @printer.
func (s *PrintSettingsClass) SetPrinter(printer string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(printer))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_printer(_arg0, _arg1)
}

// SetPrinterLpi sets the value of GTK_PRINT_SETTINGS_PRINTER_LPI.
func (s *PrintSettingsClass) SetPrinterLpi(lpi float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = C.double(lpi)

	C.gtk_print_settings_set_printer_lpi(_arg0, _arg1)
}

// SetResolution sets the values of GTK_PRINT_SETTINGS_RESOLUTION,
// GTK_PRINT_SETTINGS_RESOLUTION_X and GTK_PRINT_SETTINGS_RESOLUTION_Y.
func (s *PrintSettingsClass) SetResolution(resolution int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = C.int(resolution)

	C.gtk_print_settings_set_resolution(_arg0, _arg1)
}

// SetResolutionXY sets the values of GTK_PRINT_SETTINGS_RESOLUTION,
// GTK_PRINT_SETTINGS_RESOLUTION_X and GTK_PRINT_SETTINGS_RESOLUTION_Y.
func (s *PrintSettingsClass) SetResolutionXY(resolutionX int, resolutionY int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = C.int(resolutionX)
	_arg2 = C.int(resolutionY)

	C.gtk_print_settings_set_resolution_xy(_arg0, _arg1, _arg2)
}

// SetReverse sets the value of GTK_PRINT_SETTINGS_REVERSE.
func (s *PrintSettingsClass) SetReverse(reverse bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	if reverse {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_reverse(_arg0, _arg1)
}

// SetScale sets the value of GTK_PRINT_SETTINGS_SCALE.
func (s *PrintSettingsClass) SetScale(scale float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = C.double(scale)

	C.gtk_print_settings_set_scale(_arg0, _arg1)
}

// SetUseColor sets the value of GTK_PRINT_SETTINGS_USE_COLOR.
func (s *PrintSettingsClass) SetUseColor(useColor bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	if useColor {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_use_color(_arg0, _arg1)
}

// ToFile: this function saves the print settings from @settings to @file_name.
//
// If the file could not be written then error is set to either a `GFileError`
// or `GKeyFileError`.
func (s *PrintSettingsClass) ToFile(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ToGVariant: serialize print settings to an a{sv} variant.
func (s *PrintSettingsClass) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GVariant         // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))

	_cret = C.gtk_print_settings_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// ToKeyFile: this function adds the print settings from @settings to @key_file.
func (s *PrintSettingsClass) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_to_key_file(_arg0, _arg1, _arg2)
}

// Unset removes any value associated with @key.
//
// This has the same effect as setting the value to nil.
func (s *PrintSettingsClass) Unset(key string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_unset(_arg0, _arg1)
}

// PageRange: range of pages to print.
//
// See also [method@Gtk.PrintSettings.set_page_ranges].
type PageRange struct {
	native C.GtkPageRange
}

// WrapPageRange wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPageRange(ptr unsafe.Pointer) *PageRange {
	return (*PageRange)(ptr)
}

// Native returns the underlying C source pointer.
func (p *PageRange) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Start: start of page range.
func (p *PageRange) Start() int {
	var v int // out
	v = int(p.native.start)
	return v
}

// End: end of page range.
func (p *PageRange) End() int {
	var v int // out
	v = int(p.native.end)
	return v
}
