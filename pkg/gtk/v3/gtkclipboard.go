// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
//
// void gotk4_ClipboardTextReceivedFunc(GtkClipboard*,  gchar*, gpointer);
// void gotk4_ClipboardURIReceivedFunc(GtkClipboard*, gchar**, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_clipboard_get_type()), F: marshalClipboard},
	})
}

// ClipboardReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_contents() are received, or when the request fails.
type ClipboardReceivedFunc func(clipboard Clipboard, selectionData *SelectionData)

//export gotk4_ClipboardReceivedFunc
func gotk4_ClipboardReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.GtkSelectionData, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var clipboard Clipboard          // out
	var selectionData *SelectionData // out

	clipboard = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(Clipboard)
	selectionData = WrapSelectionData(unsafe.Pointer(arg1))

	fn := v.(ClipboardReceivedFunc)
	fn(clipboard, selectionData)
}

// ClipboardTextReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_text() are received, or when the request fails.
type ClipboardTextReceivedFunc func(clipboard Clipboard, text string)

//export gotk4_ClipboardTextReceivedFunc
func gotk4_ClipboardTextReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.gchar, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var clipboard Clipboard // out
	var text string         // out

	clipboard = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(Clipboard)
	text = C.GoString(arg1)

	fn := v.(ClipboardTextReceivedFunc)
	fn(clipboard, text)
}

// ClipboardURIReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_uris() are received, or when the request fails.
type ClipboardURIReceivedFunc func(clipboard Clipboard, uris []string)

//export gotk4_ClipboardURIReceivedFunc
func gotk4_ClipboardURIReceivedFunc(arg0 *C.GtkClipboard, arg1 **C.gchar, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var clipboard Clipboard // out
	var uris []string

	clipboard = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(Clipboard)
	{
		var length int
		for p := arg1; *p != nil; p = (**C.gchar)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		src := unsafe.Slice(arg1, length)
		uris = make([]string, length)
		for i := range src {
			uris[i] = C.GoString(src[i])
		}
	}

	fn := v.(ClipboardURIReceivedFunc)
	fn(clipboard, uris)
}

// Clipboard: the Clipboard object represents a clipboard of data shared between
// different processes or between different widgets in the same process. Each
// clipboard is identified by a name encoded as a Atom. (Conversion to and from
// strings can be done with gdk_atom_intern() and gdk_atom_name().) The default
// clipboard corresponds to the “CLIPBOARD” atom; another commonly used
// clipboard is the “PRIMARY” clipboard, which, in X, traditionally contains the
// currently selected text.
//
// To support having a number of different formats on the clipboard at the same
// time, the clipboard mechanism allows providing callbacks instead of the
// actual data. When you set the contents of the clipboard, you can either
// supply the data directly (via functions like gtk_clipboard_set_text()), or
// you can supply a callback to be called at a later time when the data is
// needed (via gtk_clipboard_set_with_data() or gtk_clipboard_set_with_owner().)
// Providing a callback also avoids having to make copies of the data when it is
// not needed.
//
// gtk_clipboard_set_with_data() and gtk_clipboard_set_with_owner() are quite
// similar; the choice between the two depends mostly on which is more
// convenient in a particular situation. The former is most useful when you want
// to have a blob of data with callbacks to convert it into the various data
// types that you advertise. When the @clear_func you provided is called, you
// simply free the data blob. The latter is more useful when the contents of
// clipboard reflect the internal state of a #GObject (As an example, for the
// PRIMARY clipboard, when an entry widget provides the clipboard’s contents the
// contents are simply the text within the selected region.) If the contents
// change, the entry widget can call gtk_clipboard_set_with_owner() to update
// the timestamp for clipboard ownership, without having to worry about
// @clear_func being called.
//
// Requesting the data from the clipboard is essentially asynchronous. If the
// contents of the clipboard are provided within the same process, then a direct
// function call will be made to retrieve the data, but if they are provided by
// another process, then the data needs to be retrieved from the other process,
// which may take some time. To avoid blocking the user interface, the call to
// request the selection, gtk_clipboard_request_contents() takes a callback that
// will be called when the contents are received (or when the request fails.) If
// you don’t want to deal with providing a separate callback, you can also use
// gtk_clipboard_wait_for_contents(). What this does is run the GLib main loop
// recursively waiting for the contents. This can simplify the code flow, but
// you still have to be aware that other callbacks in your program can be called
// while this recursive mainloop is running.
//
// Along with the functions to get the clipboard contents as an arbitrary data
// chunk, there are also functions to retrieve it as text,
// gtk_clipboard_request_text() and gtk_clipboard_wait_for_text(). These
// functions take care of determining which formats are advertised by the
// clipboard provider, asking for the clipboard in the best available format and
// converting the results into the UTF-8 encoding. (The standard form for
// representing strings in GTK+.)
type Clipboard interface {
	gextras.Objector

	// Clear clears the contents of the clipboard. Generally this should only be
	// called between the time you call gtk_clipboard_set_with_owner() or
	// gtk_clipboard_set_with_data(), and when the @clear_func you supplied is
	// called. Otherwise, the clipboard may be owned by someone else.
	Clear()
	// RequestText requests the contents of the clipboard as text. When the text
	// is later received, it will be converted to UTF-8 if necessary, and
	// @callback will be called.
	//
	// The @text parameter to @callback will contain the resulting text if the
	// request succeeded, or nil if it failed. This could happen for various
	// reasons, in particular if the clipboard was empty or if the contents of
	// the clipboard could not be converted into text form.
	RequestText(callback ClipboardTextReceivedFunc)
	// RequestUris requests the contents of the clipboard as URIs. When the URIs
	// are later received @callback will be called.
	//
	// The @uris parameter to @callback will contain the resulting array of URIs
	// if the request succeeded, or nil if it failed. This could happen for
	// various reasons, in particular if the clipboard was empty or if the
	// contents of the clipboard could not be converted into URI form.
	RequestUris(callback ClipboardURIReceivedFunc)
	// SetCanStore hints that the clipboard data should be stored somewhere when
	// the application exits or when gtk_clipboard_store () is called.
	//
	// This value is reset when the clipboard owner changes. Where the clipboard
	// data is stored is platform dependent, see gdk_display_store_clipboard ()
	// for more information.
	SetCanStore(targets []TargetEntry)
	// SetText sets the contents of the clipboard to the given UTF-8 string.
	// GTK+ will make a copy of the text and take responsibility for responding
	// for requests for the text, and for converting the text into the requested
	// format.
	SetText(text string, len int)
	// Store stores the current clipboard data somewhere so that it will stay
	// around after the application has quit.
	Store()
	// WaitForText requests the contents of the clipboard as text and converts
	// the result to UTF-8 if necessary. This function waits for the data to be
	// received using the main loop, so events, timeouts, etc, may be dispatched
	// during the wait.
	WaitForText() string
	// WaitForUris requests the contents of the clipboard as URIs. This function
	// waits for the data to be received using the main loop, so events,
	// timeouts, etc, may be dispatched during the wait.
	WaitForUris() []string
	// WaitIsImageAvailable: test to see if there is an image available to be
	// pasted This is done by requesting the TARGETS atom and checking if it
	// contains any of the supported image targets. This function waits for the
	// data to be received using the main loop, so events, timeouts, etc, may be
	// dispatched during the wait.
	//
	// This function is a little faster than calling
	// gtk_clipboard_wait_for_image() since it doesn’t need to retrieve the
	// actual image data.
	WaitIsImageAvailable() bool
	// WaitIsRichTextAvailable: test to see if there is rich text available to
	// be pasted This is done by requesting the TARGETS atom and checking if it
	// contains any of the supported rich text targets. This function waits for
	// the data to be received using the main loop, so events, timeouts, etc,
	// may be dispatched during the wait.
	//
	// This function is a little faster than calling
	// gtk_clipboard_wait_for_rich_text() since it doesn’t need to retrieve the
	// actual text.
	WaitIsRichTextAvailable(buffer TextBuffer) bool
	// WaitIsTextAvailable: test to see if there is text available to be pasted
	// This is done by requesting the TARGETS atom and checking if it contains
	// any of the supported text targets. This function waits for the data to be
	// received using the main loop, so events, timeouts, etc, may be dispatched
	// during the wait.
	//
	// This function is a little faster than calling
	// gtk_clipboard_wait_for_text() since it doesn’t need to retrieve the
	// actual text.
	WaitIsTextAvailable() bool
	// WaitIsUrisAvailable: test to see if there is a list of URIs available to
	// be pasted This is done by requesting the TARGETS atom and checking if it
	// contains the URI targets. This function waits for the data to be received
	// using the main loop, so events, timeouts, etc, may be dispatched during
	// the wait.
	//
	// This function is a little faster than calling
	// gtk_clipboard_wait_for_uris() since it doesn’t need to retrieve the
	// actual URI data.
	WaitIsUrisAvailable() bool
}

// clipboard implements the Clipboard class.
type clipboard struct {
	gextras.Objector
}

var _ Clipboard = (*clipboard)(nil)

// WrapClipboard wraps a GObject to the right type. It is
// primarily used internally.
func WrapClipboard(obj *externglib.Object) Clipboard {
	return clipboard{
		Objector: obj,
	}
}

func marshalClipboard(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapClipboard(obj), nil
}

// Clear clears the contents of the clipboard. Generally this should only be
// called between the time you call gtk_clipboard_set_with_owner() or
// gtk_clipboard_set_with_data(), and when the @clear_func you supplied is
// called. Otherwise, the clipboard may be owned by someone else.
func (c clipboard) Clear() {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))

	C.gtk_clipboard_clear(_arg0)
}

// RequestText requests the contents of the clipboard as text. When the text
// is later received, it will be converted to UTF-8 if necessary, and
// @callback will be called.
//
// The @text parameter to @callback will contain the resulting text if the
// request succeeded, or nil if it failed. This could happen for various
// reasons, in particular if the clipboard was empty or if the contents of
// the clipboard could not be converted into text form.
func (c clipboard) RequestText(callback ClipboardTextReceivedFunc) {
	var _arg0 *C.GtkClipboard                // out
	var _arg1 C.GtkClipboardTextReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*[0]byte)(C.gotk4_ClipboardTextReceivedFunc)
	_arg2 = C.gpointer(box.Assign(callback))

	C.gtk_clipboard_request_text(_arg0, _arg1, _arg2)
}

// RequestUris requests the contents of the clipboard as URIs. When the URIs
// are later received @callback will be called.
//
// The @uris parameter to @callback will contain the resulting array of URIs
// if the request succeeded, or nil if it failed. This could happen for
// various reasons, in particular if the clipboard was empty or if the
// contents of the clipboard could not be converted into URI form.
func (c clipboard) RequestUris(callback ClipboardURIReceivedFunc) {
	var _arg0 *C.GtkClipboard               // out
	var _arg1 C.GtkClipboardURIReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*[0]byte)(C.gotk4_ClipboardURIReceivedFunc)
	_arg2 = C.gpointer(box.Assign(callback))

	C.gtk_clipboard_request_uris(_arg0, _arg1, _arg2)
}

// SetCanStore hints that the clipboard data should be stored somewhere when
// the application exits or when gtk_clipboard_store () is called.
//
// This value is reset when the clipboard owner changes. Where the clipboard
// data is stored is platform dependent, see gdk_display_store_clipboard ()
// for more information.
func (c clipboard) SetCanStore(targets []TargetEntry) {
	var _arg0 *C.GtkClipboard // out
	var _arg1 *C.GtkTargetEntry
	var _arg2 C.gint

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))
	_arg2 = C.gint(len(targets))
	_arg1 = (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	C.gtk_clipboard_set_can_store(_arg0, _arg1, _arg2)
}

// SetText sets the contents of the clipboard to the given UTF-8 string.
// GTK+ will make a copy of the text and take responsibility for responding
// for requests for the text, and for converting the text into the requested
// format.
func (c clipboard) SetText(text string, len int) {
	var _arg0 *C.GtkClipboard // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gint          // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(len)

	C.gtk_clipboard_set_text(_arg0, _arg1, _arg2)
}

// Store stores the current clipboard data somewhere so that it will stay
// around after the application has quit.
func (c clipboard) Store() {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))

	C.gtk_clipboard_store(_arg0)
}

// WaitForText requests the contents of the clipboard as text and converts
// the result to UTF-8 if necessary. This function waits for the data to be
// received using the main loop, so events, timeouts, etc, may be dispatched
// during the wait.
func (c clipboard) WaitForText() string {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_clipboard_wait_for_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// WaitForUris requests the contents of the clipboard as URIs. This function
// waits for the data to be received using the main loop, so events,
// timeouts, etc, may be dispatched during the wait.
func (c clipboard) WaitForUris() []string {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))

	var _cret **C.gchar

	_cret = C.gtk_clipboard_wait_for_uris(_arg0)

	var _utf8s []string

	{
		var length int
		for p := _cret; *p != nil; p = (**C.gchar)(unsafe.Add(unsafe.Pointer(p), unsafe.Sizeof(uint(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		src := unsafe.Slice(_cret, length)
		_utf8s = make([]string, length)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// WaitIsImageAvailable: test to see if there is an image available to be
// pasted This is done by requesting the TARGETS atom and checking if it
// contains any of the supported image targets. This function waits for the
// data to be received using the main loop, so events, timeouts, etc, may be
// dispatched during the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_image() since it doesn’t need to retrieve the
// actual image data.
func (c clipboard) WaitIsImageAvailable() bool {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_clipboard_wait_is_image_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsRichTextAvailable: test to see if there is rich text available to
// be pasted This is done by requesting the TARGETS atom and checking if it
// contains any of the supported rich text targets. This function waits for
// the data to be received using the main loop, so events, timeouts, etc,
// may be dispatched during the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_rich_text() since it doesn’t need to retrieve the
// actual text.
func (c clipboard) WaitIsRichTextAvailable(buffer TextBuffer) bool {
	var _arg0 *C.GtkClipboard  // out
	var _arg1 *C.GtkTextBuffer // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_clipboard_wait_is_rich_text_available(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsTextAvailable: test to see if there is text available to be pasted
// This is done by requesting the TARGETS atom and checking if it contains
// any of the supported text targets. This function waits for the data to be
// received using the main loop, so events, timeouts, etc, may be dispatched
// during the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_text() since it doesn’t need to retrieve the
// actual text.
func (c clipboard) WaitIsTextAvailable() bool {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_clipboard_wait_is_text_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsUrisAvailable: test to see if there is a list of URIs available to
// be pasted This is done by requesting the TARGETS atom and checking if it
// contains the URI targets. This function waits for the data to be received
// using the main loop, so events, timeouts, etc, may be dispatched during
// the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_uris() since it doesn’t need to retrieve the
// actual URI data.
func (c clipboard) WaitIsUrisAvailable() bool {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_clipboard_wait_is_uris_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
