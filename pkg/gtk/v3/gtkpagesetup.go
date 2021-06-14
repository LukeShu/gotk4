// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
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
		{T: externglib.Type(C.gtk_page_setup_get_type()), F: marshalPageSetup},
	})
}

// PageSetup: a GtkPageSetup object stores the page size, orientation and
// margins. The idea is that you can get one of these from the page setup dialog
// and then pass it to the PrintOperation when printing. The benefit of
// splitting this out of the PrintSettings is that these affect the actual
// layout of the page, and thus need to be set long before user prints.
//
//
// Margins
//
// The margins specified in this object are the “print margins”, i.e. the parts
// of the page that the printer cannot print on. These are different from the
// layout margins that a word processor uses; they are typically used to
// determine the minimal size for the layout margins.
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
type PageSetup interface {
	gextras.Objector

	// Copy copies a PageSetup.
	Copy() PageSetup
	// BottomMargin gets the bottom margin in units of @unit.
	BottomMargin(unit Unit) float64
	// LeftMargin gets the left margin in units of @unit.
	LeftMargin(unit Unit) float64
	// Orientation gets the page orientation of the PageSetup.
	Orientation() PageOrientation
	// PageHeight returns the page height in units of @unit.
	//
	// Note that this function takes orientation and margins into consideration.
	// See gtk_page_setup_get_paper_height().
	PageHeight(unit Unit) float64
	// PageWidth returns the page width in units of @unit.
	//
	// Note that this function takes orientation and margins into consideration.
	// See gtk_page_setup_get_paper_width().
	PageWidth(unit Unit) float64
	// PaperHeight returns the paper height in units of @unit.
	//
	// Note that this function takes orientation, but not margins into
	// consideration. See gtk_page_setup_get_page_height().
	PaperHeight(unit Unit) float64
	// PaperSize gets the paper size of the PageSetup.
	PaperSize() *PaperSize
	// PaperWidth returns the paper width in units of @unit.
	//
	// Note that this function takes orientation, but not margins into
	// consideration. See gtk_page_setup_get_page_width().
	PaperWidth(unit Unit) float64
	// RightMargin gets the right margin in units of @unit.
	RightMargin(unit Unit) float64
	// TopMargin gets the top margin in units of @unit.
	TopMargin(unit Unit) float64
	// LoadFile reads the page setup from the file @file_name. See
	// gtk_page_setup_to_file().
	LoadFile(fileName string) error
	// SetBottomMargin sets the bottom margin of the PageSetup.
	SetBottomMargin(margin float64, unit Unit)
	// SetLeftMargin sets the left margin of the PageSetup.
	SetLeftMargin(margin float64, unit Unit)
	// SetOrientation sets the page orientation of the PageSetup.
	SetOrientation(orientation PageOrientation)
	// SetPaperSize sets the paper size of the PageSetup without changing the
	// margins. See gtk_page_setup_set_paper_size_and_default_margins().
	SetPaperSize(size *PaperSize)
	// SetPaperSizeAndDefaultMargins sets the paper size of the PageSetup and
	// modifies the margins according to the new paper size.
	SetPaperSizeAndDefaultMargins(size *PaperSize)
	// SetRightMargin sets the right margin of the PageSetup.
	SetRightMargin(margin float64, unit Unit)
	// SetTopMargin sets the top margin of the PageSetup.
	SetTopMargin(margin float64, unit Unit)
	// ToFile: this function saves the information from @setup to @file_name.
	ToFile(fileName string) error
}

// pageSetup implements the PageSetup class.
type pageSetup struct {
	gextras.Objector
}

var _ PageSetup = (*pageSetup)(nil)

// WrapPageSetup wraps a GObject to the right type. It is
// primarily used internally.
func WrapPageSetup(obj *externglib.Object) PageSetup {
	return pageSetup{
		Objector: obj,
	}
}

func marshalPageSetup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPageSetup(obj), nil
}

// NewPageSetup constructs a class PageSetup.
func NewPageSetup() PageSetup {
	var _cret C.GtkPageSetup // in

	_cret = C.gtk_page_setup_new()

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(PageSetup)

	return _pageSetup
}

// NewPageSetupFromFile constructs a class PageSetup.
func NewPageSetupFromFile(fileName string) (PageSetup, error) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GtkPageSetup // in
	var _cerr *C.GError      // in

	_cret = C.gtk_page_setup_new_from_file(_arg1, &_cerr)

	var _pageSetup PageSetup // out
	var _goerr error         // out

	_pageSetup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(PageSetup)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pageSetup, _goerr
}

// Copy copies a PageSetup.
func (o pageSetup) Copy() PageSetup {
	var _arg0 *C.GtkPageSetup // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(o.Native()))

	var _cret *C.GtkPageSetup // in

	_cret = C.gtk_page_setup_copy(_arg0)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(PageSetup)

	return _pageSetup
}

// BottomMargin gets the bottom margin in units of @unit.
func (s pageSetup) BottomMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_bottom_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// LeftMargin gets the left margin in units of @unit.
func (s pageSetup) LeftMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_left_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// Orientation gets the page orientation of the PageSetup.
func (s pageSetup) Orientation() PageOrientation {
	var _arg0 *C.GtkPageSetup // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))

	var _cret C.GtkPageOrientation // in

	_cret = C.gtk_page_setup_get_orientation(_arg0)

	var _pageOrientation PageOrientation // out

	_pageOrientation = PageOrientation(_cret)

	return _pageOrientation
}

// PageHeight returns the page height in units of @unit.
//
// Note that this function takes orientation and margins into consideration.
// See gtk_page_setup_get_paper_height().
func (s pageSetup) PageHeight(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_page_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// PageWidth returns the page width in units of @unit.
//
// Note that this function takes orientation and margins into consideration.
// See gtk_page_setup_get_paper_width().
func (s pageSetup) PageWidth(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_page_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// PaperHeight returns the paper height in units of @unit.
//
// Note that this function takes orientation, but not margins into
// consideration. See gtk_page_setup_get_page_height().
func (s pageSetup) PaperHeight(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_paper_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// PaperSize gets the paper size of the PageSetup.
func (s pageSetup) PaperSize() *PaperSize {
	var _arg0 *C.GtkPageSetup // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkPaperSize // in

	_cret = C.gtk_page_setup_get_paper_size(_arg0)

	var _paperSize *PaperSize // out

	_paperSize = WrapPaperSize(unsafe.Pointer(_cret))

	return _paperSize
}

// PaperWidth returns the paper width in units of @unit.
//
// Note that this function takes orientation, but not margins into
// consideration. See gtk_page_setup_get_page_width().
func (s pageSetup) PaperWidth(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_paper_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// RightMargin gets the right margin in units of @unit.
func (s pageSetup) RightMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_right_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// TopMargin gets the top margin in units of @unit.
func (s pageSetup) TopMargin(unit Unit) float64 {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkUnit)(unit)

	var _cret C.gdouble // in

	_cret = C.gtk_page_setup_get_top_margin(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// LoadFile reads the page setup from the file @file_name. See
// gtk_page_setup_to_file().
func (s pageSetup) LoadFile(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError // in

	C.gtk_page_setup_load_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetBottomMargin sets the bottom margin of the PageSetup.
func (s pageSetup) SetBottomMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = (C.GtkUnit)(unit)

	C.gtk_page_setup_set_bottom_margin(_arg0, _arg1, _arg2)
}

// SetLeftMargin sets the left margin of the PageSetup.
func (s pageSetup) SetLeftMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = (C.GtkUnit)(unit)

	C.gtk_page_setup_set_left_margin(_arg0, _arg1, _arg2)
}

// SetOrientation sets the page orientation of the PageSetup.
func (s pageSetup) SetOrientation(orientation PageOrientation) {
	var _arg0 *C.GtkPageSetup      // out
	var _arg1 C.GtkPageOrientation // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkPageOrientation)(orientation)

	C.gtk_page_setup_set_orientation(_arg0, _arg1)
}

// SetPaperSize sets the paper size of the PageSetup without changing the
// margins. See gtk_page_setup_set_paper_size_and_default_margins().
func (s pageSetup) SetPaperSize(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size.Native()))

	C.gtk_page_setup_set_paper_size(_arg0, _arg1)
}

// SetPaperSizeAndDefaultMargins sets the paper size of the PageSetup and
// modifies the margins according to the new paper size.
func (s pageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size.Native()))

	C.gtk_page_setup_set_paper_size_and_default_margins(_arg0, _arg1)
}

// SetRightMargin sets the right margin of the PageSetup.
func (s pageSetup) SetRightMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = (C.GtkUnit)(unit)

	C.gtk_page_setup_set_right_margin(_arg0, _arg1, _arg2)
}

// SetTopMargin sets the top margin of the PageSetup.
func (s pageSetup) SetTopMargin(margin float64, unit Unit) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 C.gdouble       // out
	var _arg2 C.GtkUnit       // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(margin)
	_arg2 = (C.GtkUnit)(unit)

	C.gtk_page_setup_set_top_margin(_arg0, _arg1, _arg2)
}

// ToFile: this function saves the information from @setup to @file_name.
func (s pageSetup) ToFile(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError // in

	C.gtk_page_setup_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
