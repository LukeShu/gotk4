// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_assistant_page_type_get_type()), F: marshalAssistantPageType},
		{T: externglib.Type(C.gtk_assistant_get_type()), F: marshalAssistanter},
	})
}

// AssistantPageType: enum for determining the page role inside the Assistant.
// It's used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// GTK_ASSISTANT_PAGE_CONFIRM, GTK_ASSISTANT_PAGE_SUMMARY or
// GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”. See
// gtk_assistant_commit() for details.
type AssistantPageType int

const (
	// Content: page has regular contents. Both the Back and forward buttons
	// will be shown.
	AssistantPageTypeContent AssistantPageType = iota
	// Intro: page contains an introduction to the assistant task. Only the
	// Forward button will be shown if there is a next page.
	AssistantPageTypeIntro
	// Confirm: page lets the user confirm or deny the changes. The Back and
	// Apply buttons will be shown.
	AssistantPageTypeConfirm
	// Summary: page informs the user of the changes done. Only the Close button
	// will be shown.
	AssistantPageTypeSummary
	// Progress: used for tasks that take a long time to complete, blocks the
	// assistant until the page is marked as complete. Only the back button will
	// be shown.
	AssistantPageTypeProgress
	// Custom: used for when other page types are not appropriate. No buttons
	// will be shown, and the application must add its own buttons through
	// gtk_assistant_add_action_widget().
	AssistantPageTypeCustom
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	return AssistantPageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AssistantPageFunc: function used by gtk_assistant_set_forward_page_func() to
// know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func(currentPage int, data cgo.Handle) (gint int)

//export gotk4_AssistantPageFunc
func gotk4_AssistantPageFunc(arg0 C.gint, arg1 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var currentPage int // out
	var data cgo.Handle // out

	currentPage = int(arg0)
	data = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(AssistantPageFunc)
	gint := fn(currentPage, data)

	cret = C.gint(gint)

	return cret
}

// AssistantOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AssistantOverrider interface {
	Apply()
	Cancel()
	Close()
	Prepare(page Widgeter)
}

// Assistanter describes Assistant's methods.
type Assistanter interface {
	// AddActionWidget adds a widget to the action area of a Assistant.
	AddActionWidget(child Widgeter)
	// AppendPage appends a page to the @assistant.
	AppendPage(page Widgeter) int
	// Commit erases the visited page history so the back button is not shown on
	// the current page, and removes the cancel button from subsequent pages.
	Commit()
	// CurrentPage returns the page number of the current page.
	CurrentPage() int
	// NPages returns the number of pages in the @assistant
	NPages() int
	// NthPage returns the child widget contained in page number @page_num.
	NthPage(pageNum int) *Widget
	// PageComplete gets whether @page is complete.
	PageComplete(page Widgeter) bool
	// PageHasPadding gets whether page has padding.
	PageHasPadding(page Widgeter) bool
	// PageHeaderImage gets the header image for @page.
	PageHeaderImage(page Widgeter) *gdkpixbuf.Pixbuf
	// PageSideImage gets the side image for @page.
	PageSideImage(page Widgeter) *gdkpixbuf.Pixbuf
	// PageTitle gets the title for @page.
	PageTitle(page Widgeter) string
	// PageType gets the page type of @page.
	PageType(page Widgeter) AssistantPageType
	// InsertPage inserts a page in the @assistant at a given position.
	InsertPage(page Widgeter, position int) int
	// NextPage: navigate to the next page.
	NextPage()
	// PrependPage prepends a page to the @assistant.
	PrependPage(page Widgeter) int
	// PreviousPage: navigate to the previous visited page.
	PreviousPage()
	// RemoveActionWidget removes a widget from the action area of a Assistant.
	RemoveActionWidget(child Widgeter)
	// RemovePage removes the @page_num’s page from @assistant.
	RemovePage(pageNum int)
	// SetCurrentPage switches the page to @page_num.
	SetCurrentPage(pageNum int)
	// SetPageComplete sets whether @page contents are complete.
	SetPageComplete(page Widgeter, complete bool)
	// SetPageHasPadding sets whether the assistant is adding padding around the
	// page.
	SetPageHasPadding(page Widgeter, hasPadding bool)
	// SetPageHeaderImage sets a header image for @page.
	SetPageHeaderImage(page Widgeter, pixbuf gdkpixbuf.Pixbufer)
	// SetPageSideImage sets a side image for @page.
	SetPageSideImage(page Widgeter, pixbuf gdkpixbuf.Pixbufer)
	// SetPageTitle sets a title for @page.
	SetPageTitle(page Widgeter, title string)
	// SetPageType sets the page type for @page.
	SetPageType(page Widgeter, typ AssistantPageType)
	// UpdateButtonsState forces @assistant to recompute the buttons state.
	UpdateButtonsState()
}

// Assistant is a widget used to represent a generally complex operation
// splitted in several steps, guiding the user through its pages and controlling
// the page flow to collect the necessary data.
//
// The design of GtkAssistant is that it controls what buttons to show and to
// make sensitive, based on what it knows about the page sequence and the
// [type][GtkAssistantPageType] of each page, in addition to state information
// like the page [completion][gtk-assistant-set-page-complete] and
// [committed][gtk-assistant-commit] status.
//
// If you have a case that doesn’t quite fit in Assistants way of handling
// buttons, you can use the K_ASSISTANT_PAGE_CUSTOM page type and handle buttons
// yourself.
//
//
// GtkAssistant as GtkBuildable
//
// The GtkAssistant implementation of the Buildable interface exposes the
// @action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in Builder, simply add it as a child to the
// GtkAssistant object, and set its child properties as necessary.
//
//
// CSS nodes
//
// GtkAssistant has a single CSS node with the name assistant.
type Assistant struct {
	Window
}

var (
	_ Assistanter     = (*Assistant)(nil)
	_ gextras.Nativer = (*Assistant)(nil)
)

func wrapAssistant(obj *externglib.Object) Assistanter {
	return &Assistant{
		Window: Window{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalAssistanter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAssistant(obj), nil
}

// NewAssistant creates a new Assistant.
func NewAssistant() *Assistant {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_assistant_new()

	var _assistant *Assistant // out

	_assistant = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Assistant)

	return _assistant
}

// AddActionWidget adds a widget to the action area of a Assistant.
func (assistant *Assistant) AddActionWidget(child Widgeter) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_assistant_add_action_widget(_arg0, _arg1)
}

// AppendPage appends a page to the @assistant.
func (assistant *Assistant) AppendPage(page Widgeter) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_append_page(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Commit erases the visited page history so the back button is not shown on the
// current page, and removes the cancel button from subsequent pages.
//
// Use this when the information provided up to the current page is hereafter
// deemed permanent and cannot be modified or undone. For example, showing a
// progress page to track a long-running, unreversible operation after the user
// has clicked apply on a confirmation page.
func (assistant *Assistant) Commit() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_commit(_arg0)
}

// CurrentPage returns the page number of the current page.
func (assistant *Assistant) CurrentPage() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	_cret = C.gtk_assistant_get_current_page(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NPages returns the number of pages in the @assistant
func (assistant *Assistant) NPages() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	_cret = C.gtk_assistant_get_n_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number @page_num.
func (assistant *Assistant) NthPage(pageNum int) *Widget {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = C.gint(pageNum)

	_cret = C.gtk_assistant_get_nth_page(_arg0, _arg1)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// PageComplete gets whether @page is complete.
func (assistant *Assistant) PageComplete(page Widgeter) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_get_page_complete(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageHasPadding gets whether page has padding.
func (assistant *Assistant) PageHasPadding(page Widgeter) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_get_page_has_padding(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageHeaderImage gets the header image for @page.
//
// Deprecated: Since GTK+ 3.2, a header is no longer shown; add your header
// decoration to the page content instead.
func (assistant *Assistant) PageHeaderImage(page Widgeter) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_get_page_header_image(_arg0, _arg1)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}

// PageSideImage gets the side image for @page.
//
// Deprecated: Since GTK+ 3.2, sidebar images are not shown anymore.
func (assistant *Assistant) PageSideImage(page Widgeter) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_get_page_side_image(_arg0, _arg1)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}

// PageTitle gets the title for @page.
func (assistant *Assistant) PageTitle(page Widgeter) string {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_get_page_title(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PageType gets the page type of @page.
func (assistant *Assistant) PageType(page Widgeter) AssistantPageType {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _cret C.GtkAssistantPageType // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_get_page_type(_arg0, _arg1)

	var _assistantPageType AssistantPageType // out

	_assistantPageType = AssistantPageType(_cret)

	return _assistantPageType
}

// InsertPage inserts a page in the @assistant at a given position.
func (assistant *Assistant) InsertPage(page Widgeter, position int) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))
	_arg2 = C.gint(position)

	_cret = C.gtk_assistant_insert_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next page.
//
// This function is for use when creating pages of the K_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) NextPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_next_page(_arg0)
}

// PrependPage prepends a page to the @assistant.
func (assistant *Assistant) PrependPage(page Widgeter) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))

	_cret = C.gtk_assistant_prepend_page(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PreviousPage: navigate to the previous visited page.
//
// It is a programming error to call this function when no previous page is
// available.
//
// This function is for use when creating pages of the K_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) PreviousPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_previous_page(_arg0)
}

// RemoveActionWidget removes a widget from the action area of a Assistant.
func (assistant *Assistant) RemoveActionWidget(child Widgeter) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_assistant_remove_action_widget(_arg0, _arg1)
}

// RemovePage removes the @page_num’s page from @assistant.
func (assistant *Assistant) RemovePage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_remove_page(_arg0, _arg1)
}

// SetCurrentPage switches the page to @page_num.
//
// Note that this will only be necessary in custom buttons, as the @assistant
// flow can be set with gtk_assistant_set_forward_page_func().
func (assistant *Assistant) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_set_current_page(_arg0, _arg1)
}

// SetPageComplete sets whether @page contents are complete.
//
// This will make @assistant update the buttons state to be able to continue the
// task.
func (assistant *Assistant) SetPageComplete(page Widgeter, complete bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))
	if complete {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_complete(_arg0, _arg1, _arg2)
}

// SetPageHasPadding sets whether the assistant is adding padding around the
// page.
func (assistant *Assistant) SetPageHasPadding(page Widgeter, hasPadding bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))
	if hasPadding {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_has_padding(_arg0, _arg1, _arg2)
}

// SetPageHeaderImage sets a header image for @page.
//
// Deprecated: Since GTK+ 3.2, a header is no longer shown; add your header
// decoration to the page content instead.
func (assistant *Assistant) SetPageHeaderImage(page Widgeter, pixbuf gdkpixbuf.Pixbufer) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer((pixbuf).(gextras.Nativer).Native()))

	C.gtk_assistant_set_page_header_image(_arg0, _arg1, _arg2)
}

// SetPageSideImage sets a side image for @page.
//
// This image used to be displayed in the side area of the assistant when @page
// is the current page.
//
// Deprecated: Since GTK+ 3.2, sidebar images are not shown anymore.
func (assistant *Assistant) SetPageSideImage(page Widgeter, pixbuf gdkpixbuf.Pixbufer) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer((pixbuf).(gextras.Nativer).Native()))

	C.gtk_assistant_set_page_side_image(_arg0, _arg1, _arg2)
}

// SetPageTitle sets a title for @page.
//
// The title is displayed in the header area of the assistant when @page is the
// current page.
func (assistant *Assistant) SetPageTitle(page Widgeter, title string) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_assistant_set_page_title(_arg0, _arg1, _arg2)
}

// SetPageType sets the page type for @page.
//
// The page type determines the page behavior in the @assistant.
func (assistant *Assistant) SetPageType(page Widgeter, typ AssistantPageType) {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkAssistantPageType // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((page).(gextras.Nativer).Native()))
	_arg2 = C.GtkAssistantPageType(typ)

	C.gtk_assistant_set_page_type(_arg0, _arg1, _arg2)
}

// UpdateButtonsState forces @assistant to recompute the buttons state.
//
// GTK+ automatically takes care of this in most situations, e.g. when the user
// goes to a different page, or when the visibility or completeness of a page
// changes.
//
// One situation where it can be necessary to call this function is when
// changing a value on the current page affects the future page flow of the
// assistant.
func (assistant *Assistant) UpdateButtonsState() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_update_buttons_state(_arg0)
}
