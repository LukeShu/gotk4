// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypePageSetup = coreglib.Type(C.gtk_page_setup_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePageSetup, F: marshalPageSetup},
	})
}

// PageSetup object stores the page size, orientation and margins. The idea
// is that you can get one of these from the page setup dialog and then pass
// it to the PrintOperation when printing. The benefit of splitting this out
// of the PrintSettings is that these affect the actual layout of the page,
// and thus need to be set long before user prints.
//
// # Margins
//
// The margins specified in this object are the “print margins”, i.e.
// the parts of the page that the printer cannot print on. These are different
// from the layout margins that a word processor uses; they are typically used
// to determine the minimal size for the layout margins.
//
// To obtain a PageSetup use gtk_page_setup_new() to get the defaults, or use
// gtk_print_run_page_setup_dialog() to show the page setup dialog and receive
// the resulting page setup.
//
// A page setup dialog
//
//    static GtkPrintSettings *settings = NULL;
//    static GtkPageSetup *page_setup = NULL;
//
//    static void
//    do_page_setup (void)
//    {
//      GtkPageSetup *new_page_setup;
//
//      if (settings == NULL)
//        settings = gtk_print_settings_new ();
//
//      new_page_setup = gtk_print_run_page_setup_dialog (GTK_WINDOW (main_window),
//                                                        page_setup, settings);
//
//      if (page_setup)
//        g_object_unref (page_setup);
//
//      page_setup = new_page_setup;
//    }
//
// Printing support was added in GTK+ 2.10.
type PageSetup struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PageSetup)(nil)
)

func wrapPageSetup(obj *coreglib.Object) *PageSetup {
	return &PageSetup{
		Object: obj,
	}
}

func marshalPageSetup(p uintptr) (interface{}, error) {
	return wrapPageSetup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPageSetup creates a new PageSetup.
//
// The function returns the following values:
//
//   - pageSetup: new PageSetup.
//
func NewPageSetup() *PageSetup {
	var _cret *C.GtkPageSetup // in

	_cret = C.gtk_page_setup_new()

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// NewPageSetupFromFile reads the page setup from the file file_name. Returns
// a new PageSetup object with the restored page setup, or NULL if an error
// occurred. See gtk_page_setup_to_file().
//
// The function takes the following parameters:
//
//   - fileName: filename to read the page setup from.
//
// The function returns the following values:
//
//   - pageSetup: restored PageSetup.
//
func NewPageSetupFromFile(fileName string) (*PageSetup, error) {
	var _arg1 *C.gchar        // out
	var _cret *C.GtkPageSetup // in
	var _cerr *C.GError       // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_page_setup_new_from_file(_arg1, &_cerr)
	runtime.KeepAlive(fileName)

	var _pageSetup *PageSetup // out
	var _goerr error          // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _pageSetup, _goerr
}

// NewPageSetupFromGVariant: desrialize a page setup from an a{sv} variant in
// the format produced by gtk_page_setup_to_gvariant().
//
// The function takes the following parameters:
//
//   - variant: a{sv} #GVariant.
//
// The function returns the following values:
//
//   - pageSetup: new PageSetup object.
//
func NewPageSetupFromGVariant(variant *glib.Variant) *PageSetup {
	var _arg1 *C.GVariant     // out
	var _cret *C.GtkPageSetup // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	_cret = C.gtk_page_setup_new_from_gvariant(_arg1)
	runtime.KeepAlive(variant)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// NewPageSetupFromKeyFile reads the page setup from the group group_name in
// the key file key_file. Returns a new PageSetup object with the restored page
// setup, or NULL if an error occurred.
//
// The function takes the following parameters:
//
//   - keyFile to retrieve the page_setup from.
//   - groupName (optional): name of the group in the key_file to read, or NULL
//     to use the default name “Page Setup”.
//
// The function returns the following values:
//
//   - pageSetup: restored PageSetup.
//
func NewPageSetupFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PageSetup, error) {
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out
	var _cret *C.GtkPageSetup // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_page_setup_new_from_key_file(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)

	var _pageSetup *PageSetup // out
	var _goerr error          // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _pageSetup, _goerr
}

// Copy copies a PageSetup.
//
// The function returns the following values:
//
//   - pageSetup: copy of other.
//
func (other *PageSetup) Copy() *PageSetup {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GtkPageSetup // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_cret = C.gtk_page_setup_copy(_arg0)
	runtime.KeepAlive(other)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// BottomMargin gets the bottom margin in units of unit.
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: bottom margin.
//
func (setup *PageSetup) BottomMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_bottom_margin(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// LeftMargin gets the left margin in units of unit.
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: left margin.
//
func (setup *PageSetup) LeftMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_left_margin(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Orientation gets the page orientation of the PageSetup.
//
// The function returns the following values:
//
//   - pageOrientation: page orientation.
//
func (setup *PageSetup) Orientation() PageOrientation {
	var _arg0 *C.GtkPageSetup      // out
	var _cret C.GtkPageOrientation // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))

	_cret = C.gtk_page_setup_get_orientation(_arg0)
	runtime.KeepAlive(setup)

	var _pageOrientation PageOrientation // out

	_pageOrientation = PageOrientation(_cret)

	return _pageOrientation
}

// PageHeight returns the page height in units of unit.
//
// Note that this function takes orientation and margins into consideration.
// See gtk_page_setup_get_paper_height().
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: page height.
//
func (setup *PageSetup) PageHeight(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_page_height(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageWidth returns the page width in units of unit.
//
// Note that this function takes orientation and margins into consideration.
// See gtk_page_setup_get_paper_width().
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: page width.
//
func (setup *PageSetup) PageWidth(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_page_width(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PaperHeight returns the paper height in units of unit.
//
// Note that this function takes orientation, but not margins into
// consideration. See gtk_page_setup_get_page_height().
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: paper height.
//
func (setup *PageSetup) PaperHeight(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_paper_height(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PaperSize gets the paper size of the PageSetup.
//
// The function returns the following values:
//
//   - paperSize: paper size.
//
func (setup *PageSetup) PaperSize() *PaperSize {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GtkPaperSize // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))

	_cret = C.gtk_page_setup_get_paper_size(_arg0)
	runtime.KeepAlive(setup)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _paperSize
}

// PaperWidth returns the paper width in units of unit.
//
// Note that this function takes orientation, but not margins into
// consideration. See gtk_page_setup_get_page_width().
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: paper width.
//
func (setup *PageSetup) PaperWidth(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_paper_width(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RightMargin gets the right margin in units of unit.
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: right margin.
//
func (setup *PageSetup) RightMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_right_margin(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// TopMargin gets the top margin in units of unit.
//
// The function takes the following parameters:
//
//   - unit for the return value.
//
// The function returns the following values:
//
//   - gdouble: top margin.
//
func (setup *PageSetup) TopMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_page_setup_get_top_margin(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// LoadFile reads the page setup from the file file_name. See
// gtk_page_setup_to_file().
//
// The function takes the following parameters:
//
//   - fileName: filename to read the page setup from.
//
func (setup *PageSetup) LoadFile(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_page_setup_load_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(fileName)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadKeyFile reads the page setup from the group group_name in the key file
// key_file.
//
// The function takes the following parameters:
//
//   - keyFile to retrieve the page_setup from.
//   - groupName (optional): name of the group in the key_file to read, or NULL
//     to use the default name “Page Setup”.
//
func (setup *PageSetup) LoadKeyFile(keyFile *glib.KeyFile, groupName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_page_setup_load_key_file(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetBottomMargin sets the bottom margin of the PageSetup.
//
// The function takes the following parameters:
//
//   - margin: new bottom margin in units of unit.
//   - unit units for margin.
//
func (setup *PageSetup) SetBottomMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_bottom_margin(_arg0, _arg1, _arg2)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(margin)
	runtime.KeepAlive(unit)
}

// SetLeftMargin sets the left margin of the PageSetup.
//
// The function takes the following parameters:
//
//   - margin: new left margin in units of unit.
//   - unit units for margin.
//
func (setup *PageSetup) SetLeftMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_left_margin(_arg0, _arg1, _arg2)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(margin)
	runtime.KeepAlive(unit)
}

// SetOrientation sets the page orientation of the PageSetup.
//
// The function takes the following parameters:
//
//   - orientation: PageOrientation value.
//
func (setup *PageSetup) SetOrientation(orientation PageOrientation) {
	var _arg0 *C.GtkPageSetup      // out
	var _arg1 C.GtkPageOrientation // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.GtkPageOrientation(orientation)

	C.gtk_page_setup_set_orientation(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(orientation)
}

// SetPaperSize sets the paper size of the PageSetup without changing the
// margins. See gtk_page_setup_set_paper_size_and_default_margins().
//
// The function takes the following parameters:
//
//   - size: PaperSize.
//
func (setup *PageSetup) SetPaperSize(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	C.gtk_page_setup_set_paper_size(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(size)
}

// SetPaperSizeAndDefaultMargins sets the paper size of the PageSetup and
// modifies the margins according to the new paper size.
//
// The function takes the following parameters:
//
//   - size: PaperSize.
//
func (setup *PageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	C.gtk_page_setup_set_paper_size_and_default_margins(_arg0, _arg1)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(size)
}

// SetRightMargin sets the right margin of the PageSetup.
//
// The function takes the following parameters:
//
//   - margin: new right margin in units of unit.
//   - unit units for margin.
//
func (setup *PageSetup) SetRightMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_right_margin(_arg0, _arg1, _arg2)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(margin)
	runtime.KeepAlive(unit)
}

// SetTopMargin sets the top margin of the PageSetup.
//
// The function takes the following parameters:
//
//   - margin: new top margin in units of unit.
//   - unit units for margin.
//
func (setup *PageSetup) SetTopMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = C.GtkUnit(unit)

	C.gtk_page_setup_set_top_margin(_arg0, _arg1, _arg2)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(margin)
	runtime.KeepAlive(unit)
}

// ToFile: this function saves the information from setup to file_name.
//
// The function takes the following parameters:
//
//   - fileName: file to save to.
//
func (setup *PageSetup) ToFile(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_page_setup_to_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(fileName)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ToGVariant: serialize page setup to an a{sv} variant.
//
// The function returns the following values:
//
//   - variant: new, floating, #GVariant.
//
func (setup *PageSetup) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))

	_cret = C.gtk_page_setup_to_gvariant(_arg0)
	runtime.KeepAlive(setup)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// ToKeyFile: this function adds the page setup from setup to key_file.
//
// The function takes the following parameters:
//
//   - keyFile to save the page setup to.
//   - groupName (optional): group to add the settings to in key_file, or NULL
//     to use the default name “Page Setup”.
//
func (setup *PageSetup) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(setup).Native()))
	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_page_setup_to_key_file(_arg0, _arg1, _arg2)
	runtime.KeepAlive(setup)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)
}
