// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_PageSetupDoneFunc(GtkPageSetup*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_error_get_type()), F: marshalPrintError},
		{T: externglib.Type(C.gtk_print_operation_action_get_type()), F: marshalPrintOperationAction},
		{T: externglib.Type(C.gtk_print_operation_result_get_type()), F: marshalPrintOperationResult},
		{T: externglib.Type(C.gtk_print_status_get_type()), F: marshalPrintStatus},
		{T: externglib.Type(C.gtk_print_operation_get_type()), F: marshalPrintOperation},
	})
}

// PrintError: error codes that identify various errors that can occur while
// using the GTK printing support.
type PrintError int

const (
	// General: unspecified error occurred.
	PrintErrorGeneral PrintError = iota
	// InternalError: internal error occurred.
	PrintErrorInternalError
	// NOMEM: memory allocation failed.
	PrintErrorNOMEM
	// InvalidFile: error occurred while loading a page setup or paper size from
	// a key file.
	PrintErrorInvalidFile
)

func marshalPrintError(p uintptr) (interface{}, error) {
	return PrintError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintOperationAction determines what action the print operation should
// perform.
//
// A parameter of this typs is passed to [method@Gtk.PrintOperation.run].
type PrintOperationAction int

const (
	// PrintDialog: show the print dialog.
	PrintOperationActionPrintDialog PrintOperationAction = iota
	// Print: start to print without showing the print dialog, based on the
	// current print settings.
	PrintOperationActionPrint
	// Preview: show the print preview.
	PrintOperationActionPreview
	// Export: export to a file. This requires the export-filename property to
	// be set.
	PrintOperationActionExport
)

func marshalPrintOperationAction(p uintptr) (interface{}, error) {
	return PrintOperationAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintOperationResult: the result of a print operation.
//
// A value of this type is returned by [method@Gtk.PrintOperation.run].
type PrintOperationResult int

const (
	// Error has occurred.
	PrintOperationResultError PrintOperationResult = iota
	// Apply: the print settings should be stored.
	PrintOperationResultApply
	// Cancel: the print operation has been canceled, the print settings should
	// not be stored.
	PrintOperationResultCancel
	// InProgress: the print operation is not complete yet. This value will only
	// be returned when running asynchronously.
	PrintOperationResultInProgress
)

func marshalPrintOperationResult(p uintptr) (interface{}, error) {
	return PrintOperationResult(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintStatus: the status gives a rough indication of the completion of a
// running print operation.
type PrintStatus int

const (
	// Initial: the printing has not started yet; this status is set initially,
	// and while the print dialog is shown.
	PrintStatusInitial PrintStatus = iota
	// Preparing: this status is set while the begin-print signal is emitted and
	// during pagination.
	PrintStatusPreparing
	// GeneratingData: this status is set while the pages are being rendered.
	PrintStatusGeneratingData
	// SendingData: the print job is being sent off to the printer.
	PrintStatusSendingData
	// Pending: the print job has been sent to the printer, but is not printed
	// for some reason, e.g. the printer may be stopped.
	PrintStatusPending
	// PendingIssue: some problem has occurred during printing, e.g. a paper
	// jam.
	PrintStatusPendingIssue
	// Printing: the printer is processing the print job.
	PrintStatusPrinting
	// Finished: the printing has been completed successfully.
	PrintStatusFinished
	// FinishedAborted: the printing has been aborted.
	PrintStatusFinishedAborted
)

func marshalPrintStatus(p uintptr) (interface{}, error) {
	return PrintStatus(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PageSetupDoneFunc: the type of function that is passed to
// gtk_print_run_page_setup_dialog_async().
//
// This function will be called when the page setup dialog is dismissed, and
// also serves as destroy notify for @data.
type PageSetupDoneFunc func(pageSetup *PageSetupClass, data interface{})

//export gotk4_PageSetupDoneFunc
func gotk4_PageSetupDoneFunc(arg0 *C.GtkPageSetup, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var pageSetup *PageSetupClass // out
	var data interface{}          // out

	pageSetup = gextras.CastObject(
		externglib.Take(unsafe.Pointer(arg0))).(*PageSetupClass)
	data = box.Get(uintptr(arg1))

	fn := v.(PageSetupDoneFunc)
	fn(pageSetup, data)
}

// PrintRunPageSetupDialog runs a page setup dialog, letting the user modify the
// values from @page_setup. If the user cancels the dialog, the returned
// PageSetup is identical to the passed in @page_setup, otherwise it contains
// the modifications done in the dialog.
//
// Note that this function may use a recursive mainloop to show the page setup
// dialog. See gtk_print_run_page_setup_dialog_async() if this is a problem.
func PrintRunPageSetupDialog(parent Window, pageSetup PageSetup, settings PrintSettings) *PageSetupClass {
	var _arg1 *C.GtkWindow        // out
	var _arg2 *C.GtkPageSetup     // out
	var _arg3 *C.GtkPrintSettings // out
	var _cret *C.GtkPageSetup     // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer((&parent).Native()))
	_arg2 = (*C.GtkPageSetup)(unsafe.Pointer((&pageSetup).Native()))
	_arg3 = (*C.GtkPrintSettings)(unsafe.Pointer((&settings).Native()))

	_cret = C.gtk_print_run_page_setup_dialog(_arg1, _arg2, _arg3)

	var _pageSetup *PageSetupClass // out

	_pageSetup = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PageSetupClass)

	return _pageSetup
}

// PrintRunPageSetupDialogAsync runs a page setup dialog, letting the user
// modify the values from @page_setup.
//
// In contrast to gtk_print_run_page_setup_dialog(), this function returns after
// showing the page setup dialog on platforms that support this, and calls
// @done_cb from a signal handler for the ::response signal of the dialog.
func PrintRunPageSetupDialogAsync(parent Window, pageSetup PageSetup, settings PrintSettings, doneCb PageSetupDoneFunc) {
	var _arg1 *C.GtkWindow           // out
	var _arg2 *C.GtkPageSetup        // out
	var _arg3 *C.GtkPrintSettings    // out
	var _arg4 C.GtkPageSetupDoneFunc // out
	var _arg5 C.gpointer

	_arg1 = (*C.GtkWindow)(unsafe.Pointer((&parent).Native()))
	_arg2 = (*C.GtkPageSetup)(unsafe.Pointer((&pageSetup).Native()))
	_arg3 = (*C.GtkPrintSettings)(unsafe.Pointer((&settings).Native()))
	_arg4 = (*[0]byte)(C.gotk4_PageSetupDoneFunc)
	_arg5 = C.gpointer(box.Assign(doneCb))

	C.gtk_print_run_page_setup_dialog_async(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// PrintOperationOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PrintOperationOverrider interface {
	BeginPrint(context PrintContext)
	CustomWidgetApply(widget Widget)
	DrawPage(context PrintContext, pageNr int)
	EndPrint(context PrintContext)
	Paginate(context PrintContext) bool
	Preview(preview PrintOperationPreview, context PrintContext, parent Window) bool
	RequestPageSetup(context PrintContext, pageNr int, setup PageSetup)
	StatusChanged()
	UpdateCustomWidget(widget Widget, setup PageSetup, settings PrintSettings)
}

// PrintOperation: `GtkPrintOperation` is the high-level, portable printing API.
//
// It looks a bit different than other GTK dialogs such as the `GtkFileChooser`,
// since some platforms don’t expose enough infrastructure to implement a good
// print dialog. On such platforms, `GtkPrintOperation` uses the native print
// dialog. On platforms which do not provide a native print dialog, GTK uses its
// own, see [class@Gtk.PrintUnixDialog].
//
// The typical way to use the high-level printing API is to create a
// `GtkPrintOperation` object with [ctor@Gtk.PrintOperation.new] when the user
// selects to print. Then you set some properties on it, e.g. the page size, any
// [class@Gtk.PrintSettings] from previous print operations, the number of
// pages, the current page, etc.
//
// Then you start the print operation by calling
// [method@Gtk.PrintOperation.run]. It will then show a dialog, let the user
// select a printer and options. When the user finished the dialog, various
// signals will be emitted on the `GtkPrintOperation`, the main one being
// [signal@Gtk.PrintOperation::draw-page], which you are supposed to handle and
// render the page on the provided [class@Gtk.PrintContext] using Cairo.
//
//
// The high-level printing API
//
// “`c static GtkPrintSettings *settings = NULL;
//
// static void do_print (void) { GtkPrintOperation *print;
// GtkPrintOperationResult res;
//
//    print = gtk_print_operation_new ();
//
//    if (settings != NULL)
//      gtk_print_operation_set_print_settings (print, settings);
//
//    g_signal_connect (print, "begin_print", G_CALLBACK (begin_print), NULL);
//    g_signal_connect (print, "draw_page", G_CALLBACK (draw_page), NULL);
//
//    res = gtk_print_operation_run (print, GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG,
//                                   GTK_WINDOW (main_window), NULL);
//
//    if (res == GTK_PRINT_OPERATION_RESULT_APPLY)
//      {
//        if (settings != NULL)
//          g_object_unref (settings);
//        settings = g_object_ref (gtk_print_operation_get_print_settings (print));
//      }
//
//    g_object_unref (print);
//
// } “`
//
// By default `GtkPrintOperation` uses an external application to do print
// preview. To implement a custom print preview, an application must connect to
// the preview signal. The functions
// [method@Gtk.PrintOperationPreview.render_page],
// [method@Gtk.PrintOperationPreview.end_preview] and
// [method@Gtk.PrintOperationPreview.is_selected] are useful when implementing a
// print preview.
type PrintOperation interface {
	gextras.Objector

	// Cancel cancels a running print operation.
	//
	// This function may be called from a
	// [signal@Gtk.PrintOperation::begin-print],
	// [signal@Gtk.PrintOperation::paginate] or
	// [signal@Gtk.PrintOperation::draw-page] signal handler to stop the
	// currently running print operation.
	Cancel()
	// DrawPageFinish: signal that drawing of particular page is complete.
	//
	// It is called after completion of page drawing (e.g. drawing in another
	// thread). If [method@Gtk.PrintOperation.set_defer_drawing] was called
	// before, then this function has to be called by application. Otherwise it
	// is called by GTK itself.
	DrawPageFinish()
	// DefaultPageSetup returns the default page setup.
	DefaultPageSetup() *PageSetupClass
	// EmbedPageSetup gets whether page setup selection combos are embedded
	EmbedPageSetup() bool
	// Error: call this when the result of a print operation is
	// GTK_PRINT_OPERATION_RESULT_ERROR.
	//
	// It can be called either after [method@Gtk.PrintOperation.run] returns, or
	// in the [signal@Gtk.PrintOperation::done] signal handler.
	//
	// The returned `GError` will contain more details on what went wrong.
	Error() error
	// HasSelection gets whether there is a selection.
	HasSelection() bool
	// NPagesToPrint returns the number of pages that will be printed.
	//
	// Note that this value is set during print preparation phase
	// (GTK_PRINT_STATUS_PREPARING), so this function should never be called
	// before the data generation phase (GTK_PRINT_STATUS_GENERATING_DATA). You
	// can connect to the [signal@Gtk.PrintOperation::status-changed] signal and
	// call gtk_print_operation_get_n_pages_to_print() when print status is
	// GTK_PRINT_STATUS_GENERATING_DATA.
	//
	// This is typically used to track the progress of print operation.
	NPagesToPrint() int
	// PrintSettings returns the current print settings.
	//
	// Note that the return value is nil until either
	// [method@Gtk.PrintOperation.set_print_settings] or
	// [method@Gtk.PrintOperation.run] have been called.
	PrintSettings() *PrintSettingsClass
	// Status returns the status of the print operation.
	//
	// Also see [method@Gtk.PrintOperation.get_status_string].
	Status() PrintStatus
	// StatusString returns a string representation of the status of the print
	// operation.
	//
	// The string is translated and suitable for displaying the print status
	// e.g. in a `GtkStatusbar`.
	//
	// Use [method@Gtk.PrintOperation.get_status] to obtain a status value that
	// is suitable for programmatic use.
	StatusString() string
	// SupportSelection gets whether the application supports print of selection
	SupportSelection() bool
	// IsFinished: convenience function to find out if the print operation is
	// finished.
	//
	// a print operation is finished if its status is either
	// GTK_PRINT_STATUS_FINISHED or GTK_PRINT_STATUS_FINISHED_ABORTED.
	//
	// Note: when you enable print status tracking the print operation can be in
	// a non-finished state even after done has been called, as the operation
	// status then tracks the print job status on the printer.
	IsFinished() bool
	// SetAllowAsync sets whether gtk_print_operation_run() may return before
	// the print operation is completed.
	//
	// Note that some platforms may not allow asynchronous operation.
	SetAllowAsync(allowAsync bool)
	// SetCurrentPage sets the current page.
	//
	// If this is called before [method@Gtk.PrintOperation.run], the user will
	// be able to select to print only the current page.
	//
	// Note that this only makes sense for pre-paginated documents.
	SetCurrentPage(currentPage int)
	// SetCustomTabLabel sets the label for the tab holding custom widgets.
	SetCustomTabLabel(label string)
	// SetDefaultPageSetup makes @default_page_setup the default page setup for
	// @op.
	//
	// This page setup will be used by [method@Gtk.PrintOperation.run], but it
	// can be overridden on a per-page basis by connecting to the
	// [signal@Gtk.PrintOperation::request-page-setup] signal.
	SetDefaultPageSetup(defaultPageSetup PageSetup)
	// SetDeferDrawing sets up the `GtkPrintOperation` to wait for calling of
	// [method@Gtk.PrintOperation.draw_page_finish from application.
	//
	// This can be used for drawing page in another thread.
	//
	// This function must be called in the callback of the
	// [signal@Gtk.PrintOperation::draw-page] signal.
	SetDeferDrawing()
	// SetEmbedPageSetup: embed page size combo box and orientation combo box
	// into page setup page.
	//
	// Selected page setup is stored as default page setup in
	// `GtkPrintOperation`.
	SetEmbedPageSetup(embed bool)
	// SetExportFilename sets up the `GtkPrintOperation` to generate a file
	// instead of showing the print dialog.
	//
	// The intended use of this function is for implementing “Export to PDF”
	// actions. Currently, PDF is the only supported format.
	//
	// “Print to PDF” support is independent of this and is done by letting the
	// user pick the “Print to PDF” item from the list of printers in the print
	// dialog.
	SetExportFilename(filename string)
	// SetHasSelection sets whether there is a selection to print.
	//
	// Application has to set number of pages to which the selection will draw
	// by [method@Gtk.PrintOperation.set_n_pages] in a handler for the
	// [signal@Gtk.PrintOperation::begin-print] signal.
	SetHasSelection(hasSelection bool)
	// SetJobName sets the name of the print job.
	//
	// The name is used to identify the job (e.g. in monitoring applications
	// like eggcups).
	//
	// If you don’t set a job name, GTK picks a default one by numbering
	// successive print jobs.
	SetJobName(jobName string)
	// SetNPages sets the number of pages in the document.
	//
	// This must be set to a positive number before the rendering starts. It may
	// be set in a [signal@Gtk.PrintOperation::begin-print] signal handler.
	//
	// Note that the page numbers passed to the
	// [signal@Gtk.PrintOperation::request-page-setup] and
	// [signal@Gtk.PrintOperation::draw-page] signals are 0-based, i.e. if the
	// user chooses to print all pages, the last ::draw-page signal will be for
	// page @n_pages - 1.
	SetNPages(nPages int)
	// SetPrintSettings sets the print settings for @op.
	//
	// This is typically used to re-establish print settings from a previous
	// print operation, see [method@Gtk.PrintOperation.run].
	SetPrintSettings(printSettings PrintSettings)
	// SetShowProgress: if @show_progress is true, the print operation will show
	// a progress dialog during the print operation.
	SetShowProgress(showProgress bool)
	// SetSupportSelection sets whether selection is supported by
	// `GtkPrintOperation`.
	SetSupportSelection(supportSelection bool)
	// SetTrackPrintStatus: if track_status is true, the print operation will
	// try to continue report on the status of the print job in the printer
	// queues and printer.
	//
	// This can allow your application to show things like “out of paper”
	// issues, and when the print job actually reaches the printer.
	//
	// This function is often implemented using some form of polling, so it
	// should not be enabled unless needed.
	SetTrackPrintStatus(trackStatus bool)
	// SetUseFullPage: if @full_page is true, the transformation for the cairo
	// context obtained from `GtkPrintContext` puts the origin at the top left
	// corner of the page.
	//
	// This may not be the top left corner of the sheet, depending on page
	// orientation and the number of pages per sheet). Otherwise, the origin is
	// at the top left corner of the imageable area (i.e. inside the margins).
	SetUseFullPage(fullPage bool)
}

// PrintOperationClass implements the PrintOperation interface.
type PrintOperationClass struct {
	*externglib.Object
	PrintOperationPreviewInterface
}

var _ PrintOperation = (*PrintOperationClass)(nil)

func wrapPrintOperation(obj *externglib.Object) PrintOperation {
	return &PrintOperationClass{
		Object: obj,
		PrintOperationPreviewInterface: PrintOperationPreviewInterface{
			Object: obj,
		},
	}
}

func marshalPrintOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPrintOperation(obj), nil
}

// NewPrintOperation creates a new `GtkPrintOperation`.
func NewPrintOperation() *PrintOperationClass {
	var _cret *C.GtkPrintOperation // in

	_cret = C.gtk_print_operation_new()

	var _printOperation *PrintOperationClass // out

	_printOperation = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PrintOperationClass)

	return _printOperation
}

// Cancel cancels a running print operation.
//
// This function may be called from a [signal@Gtk.PrintOperation::begin-print],
// [signal@Gtk.PrintOperation::paginate] or
// [signal@Gtk.PrintOperation::draw-page] signal handler to stop the currently
// running print operation.
func (o *PrintOperationClass) Cancel() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	C.gtk_print_operation_cancel(_arg0)
}

// DrawPageFinish: signal that drawing of particular page is complete.
//
// It is called after completion of page drawing (e.g. drawing in another
// thread). If [method@Gtk.PrintOperation.set_defer_drawing] was called before,
// then this function has to be called by application. Otherwise it is called by
// GTK itself.
func (o *PrintOperationClass) DrawPageFinish() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	C.gtk_print_operation_draw_page_finish(_arg0)
}

// DefaultPageSetup returns the default page setup.
func (o *PrintOperationClass) DefaultPageSetup() *PageSetupClass {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.GtkPageSetup      // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_default_page_setup(_arg0)

	var _pageSetup *PageSetupClass // out

	_pageSetup = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*PageSetupClass)

	return _pageSetup
}

// EmbedPageSetup gets whether page setup selection combos are embedded
func (o *PrintOperationClass) EmbedPageSetup() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_embed_page_setup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Error: call this when the result of a print operation is
// GTK_PRINT_OPERATION_RESULT_ERROR.
//
// It can be called either after [method@Gtk.PrintOperation.run] returns, or in
// the [signal@Gtk.PrintOperation::done] signal handler.
//
// The returned `GError` will contain more details on what went wrong.
func (o *PrintOperationClass) Error() error {
	var _arg0 *C.GtkPrintOperation // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	C.gtk_print_operation_get_error(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// HasSelection gets whether there is a selection.
func (o *PrintOperationClass) HasSelection() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_has_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NPagesToPrint returns the number of pages that will be printed.
//
// Note that this value is set during print preparation phase
// (GTK_PRINT_STATUS_PREPARING), so this function should never be called before
// the data generation phase (GTK_PRINT_STATUS_GENERATING_DATA). You can connect
// to the [signal@Gtk.PrintOperation::status-changed] signal and call
// gtk_print_operation_get_n_pages_to_print() when print status is
// GTK_PRINT_STATUS_GENERATING_DATA.
//
// This is typically used to track the progress of print operation.
func (o *PrintOperationClass) NPagesToPrint() int {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.int                // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_n_pages_to_print(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrintSettings returns the current print settings.
//
// Note that the return value is nil until either
// [method@Gtk.PrintOperation.set_print_settings] or
// [method@Gtk.PrintOperation.run] have been called.
func (o *PrintOperationClass) PrintSettings() *PrintSettingsClass {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.GtkPrintSettings  // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_print_settings(_arg0)

	var _printSettings *PrintSettingsClass // out

	_printSettings = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*PrintSettingsClass)

	return _printSettings
}

// Status returns the status of the print operation.
//
// Also see [method@Gtk.PrintOperation.get_status_string].
func (o *PrintOperationClass) Status() PrintStatus {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.GtkPrintStatus     // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_status(_arg0)

	var _printStatus PrintStatus // out

	_printStatus = (PrintStatus)(_cret)

	return _printStatus
}

// StatusString returns a string representation of the status of the print
// operation.
//
// The string is translated and suitable for displaying the print status e.g. in
// a `GtkStatusbar`.
//
// Use [method@Gtk.PrintOperation.get_status] to obtain a status value that is
// suitable for programmatic use.
func (o *PrintOperationClass) StatusString() string {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_status_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SupportSelection gets whether the application supports print of selection
func (o *PrintOperationClass) SupportSelection() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_get_support_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsFinished: convenience function to find out if the print operation is
// finished.
//
// a print operation is finished if its status is either
// GTK_PRINT_STATUS_FINISHED or GTK_PRINT_STATUS_FINISHED_ABORTED.
//
// Note: when you enable print status tracking the print operation can be in a
// non-finished state even after done has been called, as the operation status
// then tracks the print job status on the printer.
func (o *PrintOperationClass) IsFinished() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	_cret = C.gtk_print_operation_is_finished(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAllowAsync sets whether gtk_print_operation_run() may return before the
// print operation is completed.
//
// Note that some platforms may not allow asynchronous operation.
func (o *PrintOperationClass) SetAllowAsync(allowAsync bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	if allowAsync {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_allow_async(_arg0, _arg1)
}

// SetCurrentPage sets the current page.
//
// If this is called before [method@Gtk.PrintOperation.run], the user will be
// able to select to print only the current page.
//
// Note that this only makes sense for pre-paginated documents.
func (o *PrintOperationClass) SetCurrentPage(currentPage int) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	_arg1 = C.int(currentPage)

	C.gtk_print_operation_set_current_page(_arg0, _arg1)
}

// SetCustomTabLabel sets the label for the tab holding custom widgets.
func (o *PrintOperationClass) SetCustomTabLabel(label string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_custom_tab_label(_arg0, _arg1)
}

// SetDefaultPageSetup makes @default_page_setup the default page setup for @op.
//
// This page setup will be used by [method@Gtk.PrintOperation.run], but it can
// be overridden on a per-page basis by connecting to the
// [signal@Gtk.PrintOperation::request-page-setup] signal.
func (o *PrintOperationClass) SetDefaultPageSetup(defaultPageSetup PageSetup) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.GtkPageSetup      // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer((&defaultPageSetup).Native()))

	C.gtk_print_operation_set_default_page_setup(_arg0, _arg1)
}

// SetDeferDrawing sets up the `GtkPrintOperation` to wait for calling of
// [method@Gtk.PrintOperation.draw_page_finish from application.
//
// This can be used for drawing page in another thread.
//
// This function must be called in the callback of the
// [signal@Gtk.PrintOperation::draw-page] signal.
func (o *PrintOperationClass) SetDeferDrawing() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))

	C.gtk_print_operation_set_defer_drawing(_arg0)
}

// SetEmbedPageSetup: embed page size combo box and orientation combo box into
// page setup page.
//
// Selected page setup is stored as default page setup in `GtkPrintOperation`.
func (o *PrintOperationClass) SetEmbedPageSetup(embed bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	if embed {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_embed_page_setup(_arg0, _arg1)
}

// SetExportFilename sets up the `GtkPrintOperation` to generate a file instead
// of showing the print dialog.
//
// The intended use of this function is for implementing “Export to PDF”
// actions. Currently, PDF is the only supported format.
//
// “Print to PDF” support is independent of this and is done by letting the user
// pick the “Print to PDF” item from the list of printers in the print dialog.
func (o *PrintOperationClass) SetExportFilename(filename string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_export_filename(_arg0, _arg1)
}

// SetHasSelection sets whether there is a selection to print.
//
// Application has to set number of pages to which the selection will draw by
// [method@Gtk.PrintOperation.set_n_pages] in a handler for the
// [signal@Gtk.PrintOperation::begin-print] signal.
func (o *PrintOperationClass) SetHasSelection(hasSelection bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	if hasSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_has_selection(_arg0, _arg1)
}

// SetJobName sets the name of the print job.
//
// The name is used to identify the job (e.g. in monitoring applications like
// eggcups).
//
// If you don’t set a job name, GTK picks a default one by numbering successive
// print jobs.
func (o *PrintOperationClass) SetJobName(jobName string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	_arg1 = (*C.char)(C.CString(jobName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_job_name(_arg0, _arg1)
}

// SetNPages sets the number of pages in the document.
//
// This must be set to a positive number before the rendering starts. It may be
// set in a [signal@Gtk.PrintOperation::begin-print] signal handler.
//
// Note that the page numbers passed to the
// [signal@Gtk.PrintOperation::request-page-setup] and
// [signal@Gtk.PrintOperation::draw-page] signals are 0-based, i.e. if the user
// chooses to print all pages, the last ::draw-page signal will be for page
// @n_pages - 1.
func (o *PrintOperationClass) SetNPages(nPages int) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	_arg1 = C.int(nPages)

	C.gtk_print_operation_set_n_pages(_arg0, _arg1)
}

// SetPrintSettings sets the print settings for @op.
//
// This is typically used to re-establish print settings from a previous print
// operation, see [method@Gtk.PrintOperation.run].
func (o *PrintOperationClass) SetPrintSettings(printSettings PrintSettings) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.GtkPrintSettings  // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer((&printSettings).Native()))

	C.gtk_print_operation_set_print_settings(_arg0, _arg1)
}

// SetShowProgress: if @show_progress is true, the print operation will show a
// progress dialog during the print operation.
func (o *PrintOperationClass) SetShowProgress(showProgress bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	if showProgress {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_show_progress(_arg0, _arg1)
}

// SetSupportSelection sets whether selection is supported by
// `GtkPrintOperation`.
func (o *PrintOperationClass) SetSupportSelection(supportSelection bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	if supportSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_support_selection(_arg0, _arg1)
}

// SetTrackPrintStatus: if track_status is true, the print operation will try to
// continue report on the status of the print job in the printer queues and
// printer.
//
// This can allow your application to show things like “out of paper” issues,
// and when the print job actually reaches the printer.
//
// This function is often implemented using some form of polling, so it should
// not be enabled unless needed.
func (o *PrintOperationClass) SetTrackPrintStatus(trackStatus bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	if trackStatus {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_track_print_status(_arg0, _arg1)
}

// SetUseFullPage: if @full_page is true, the transformation for the cairo
// context obtained from `GtkPrintContext` puts the origin at the top left
// corner of the page.
//
// This may not be the top left corner of the sheet, depending on page
// orientation and the number of pages per sheet). Otherwise, the origin is at
// the top left corner of the imageable area (i.e. inside the margins).
func (o *PrintOperationClass) SetUseFullPage(fullPage bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer((&o).Native()))
	if fullPage {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_use_full_page(_arg0, _arg1)
}
