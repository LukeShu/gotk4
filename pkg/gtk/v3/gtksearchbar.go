// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_search_bar_get_type()), F: marshalSearchBar},
	})
}

// SearchBar is a container made to have a search entry (possibly with
// additional connex widgets, such as drop-down menus, or buttons) built-in. The
// search bar would appear when a search is started through typing on the
// keyboard, or the application’s search mode is toggled on.
//
// For keyboard presses to start a search, events will need to be forwarded from
// the top-level window that contains the search bar. See
// gtk_search_bar_handle_event() for example code. Common shortcuts such as
// Ctrl+F should be handled as an application action, or through the menu items.
//
// You will also need to tell the search bar about which entry you are using as
// your search entry using gtk_search_bar_connect_entry(). The following example
// shows you how to create a more complex search entry.
//
//
// CSS nodes
//
// GtkSearchBar has a single CSS node with name searchbar.
//
//
// Creating a search bar
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/blob/gtk-3-24/examples/search-bar.c)
type SearchBar interface {
	gextras.Objector

	// ConnectEntry connects the Entry widget passed as the one to be used in
	// this search bar. The entry should be a descendant of the search bar. This
	// is only required if the entry isn’t the direct child of the search bar
	// (as in our main example).
	ConnectEntry(entry Entry)
	// SearchMode returns whether the search mode is on or off.
	SearchMode() bool
	// ShowCloseButton returns whether the close button is shown.
	ShowCloseButton() bool
	// SetSearchMode switches the search mode on or off.
	SetSearchMode(searchMode bool)
	// SetShowCloseButton shows or hides the close button. Applications that
	// already have a “search” toggle button should not show a close button in
	// their search bar, as it duplicates the role of the toggle button.
	SetShowCloseButton(visible bool)
}

// SearchBarClass implements the SearchBar interface.
type SearchBarClass struct {
	*externglib.Object
	BinClass
	BuildableInterface
}

var _ SearchBar = (*SearchBarClass)(nil)

func wrapSearchBar(obj *externglib.Object) SearchBar {
	return &SearchBarClass{
		Object: obj,
		BinClass: BinClass{
			Object: obj,
			ContainerClass: ContainerClass{
				Object: obj,
				WidgetClass: WidgetClass{
					Object:           obj,
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSearchBar(obj), nil
}

// NewSearchBar creates a SearchBar. You will need to tell it about which widget
// is going to be your text entry using gtk_search_bar_connect_entry().
func NewSearchBar() SearchBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_search_bar_new()

	var _searchBar SearchBar // out

	_searchBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(SearchBar)

	return _searchBar
}

// ConnectEntry connects the Entry widget passed as the one to be used in this
// search bar. The entry should be a descendant of the search bar. This is only
// required if the entry isn’t the direct child of the search bar (as in our
// main example).
func (b *SearchBarClass) ConnectEntry(entry Entry) {
	var _arg0 *C.GtkSearchBar // out
	var _arg1 *C.GtkEntry     // out

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_search_bar_connect_entry(_arg0, _arg1)
}

// SearchMode returns whether the search mode is on or off.
func (b *SearchBarClass) SearchMode() bool {
	var _arg0 *C.GtkSearchBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_search_bar_get_search_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether the close button is shown.
func (b *SearchBarClass) ShowCloseButton() bool {
	var _arg0 *C.GtkSearchBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_search_bar_get_show_close_button(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSearchMode switches the search mode on or off.
func (b *SearchBarClass) SetSearchMode(searchMode bool) {
	var _arg0 *C.GtkSearchBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))
	if searchMode {
		_arg1 = C.TRUE
	}

	C.gtk_search_bar_set_search_mode(_arg0, _arg1)
}

// SetShowCloseButton shows or hides the close button. Applications that already
// have a “search” toggle button should not show a close button in their search
// bar, as it duplicates the role of the toggle button.
func (b *SearchBarClass) SetShowCloseButton(visible bool) {
	var _arg0 *C.GtkSearchBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_search_bar_set_show_close_button(_arg0, _arg1)
}
