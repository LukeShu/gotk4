// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_search_bar_get_type()), F: marshalSearchBar},
	})
}

// SearchBar: `GtkSearchBar` is a container made to have a search entry.
//
// !An example GtkSearchBar (search-bar.png)
//
// It can also contain additional widgets, such as drop-down menus, or buttons.
// The search bar would appear when a search is started through typing on the
// keyboard, or the application’s search mode is toggled on.
//
// For keyboard presses to start a search, the search bar must be told of a
// widget to capture key events from through
// [method@Gtk.SearchBar.set_key_capture_widget]. This widget will typically be
// the top-level window, or a parent container of the search bar. Common
// shortcuts such as Ctrl+F should be handled as an application action, or
// through the menu items.
//
// You will also need to tell the search bar about which entry you are using as
// your search entry using [method@Gtk.SearchBar.connect_entry].
//
//
// Creating a search bar
//
// The following example shows you how to create a more complex search entry.
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/examples/search-bar.c)
//
//
// CSS nodes
//
// “` searchbar ╰── revealer ╰── box ├── [child] ╰── [button.close] “`
//
// `GtkSearchBar` has a main CSS node with name searchbar. It has a child node
// with name revealer that contains a node with name box. The box node contains
// both the CSS node of the child widget as well as an optional button node
// which gets the .close style class applied.
//
//
// Accessibility
//
// `GtkSearchBar` uses the GTK_ACCESSIBLE_ROLE_SEARCH role.
type SearchBar interface {
	gextras.Objector

	// ConnectEntry connects the `GtkEditable widget passed as the one to be
	// used in this search bar.
	//
	// The entry should be a descendant of the search bar. Calling this function
	// manually is only required if the entry isn’t the direct child of the
	// search bar (as in our main example).
	ConnectEntry(entry Editable)
	// Child gets the child widget of @bar.
	Child() Widget
	// KeyCaptureWidget gets the widget that @bar is capturing key events from.
	KeyCaptureWidget() Widget
	// SearchMode returns whether the search mode is on or off.
	SearchMode() bool
	// ShowCloseButton returns whether the close button is shown.
	ShowCloseButton() bool
	// SetChild sets the child widget of @bar.
	SetChild(child Widget)
	// SetKeyCaptureWidget sets @widget as the widget that @bar will capture key
	// events from.
	//
	// If key events are handled by the search bar, the bar will be shown, and
	// the entry populated with the entered text.
	//
	// Note that despite the name of this function, the events are only
	// 'captured' in the bubble phase, which means that editable child widgets
	// of @widget will receive text input before it gets captured. If that is
	// not desired, you can capture and forward the events yourself with
	// [method@Gtk.EventControllerKey.forward].
	SetKeyCaptureWidget(widget Widget)
	// SetSearchMode switches the search mode on or off.
	SetSearchMode(searchMode bool)
	// SetShowCloseButton shows or hides the close button.
	//
	// Applications that already have a “search” toggle button should not show a
	// close button in their search bar, as it duplicates the role of the toggle
	// button.
	SetShowCloseButton(visible bool)
}

// SearchBarClass implements the SearchBar interface.
type SearchBarClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
}

var _ SearchBar = (*SearchBarClass)(nil)

func wrapSearchBar(obj *externglib.Object) SearchBar {
	return &SearchBarClass{
		Object: obj,
		WidgetClass: WidgetClass{
			Object:           obj,
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
	}
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSearchBar(obj), nil
}

// NewSearchBar creates a `GtkSearchBar`.
//
// You will need to tell it about which widget is going to be your text entry
// using [method@Gtk.SearchBar.connect_entry].
func NewSearchBar() SearchBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_search_bar_new()

	var _searchBar SearchBar // out

	_searchBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(SearchBar)

	return _searchBar
}

// ConnectEntry connects the `GtkEditable widget passed as the one to be used in
// this search bar.
//
// The entry should be a descendant of the search bar. Calling this function
// manually is only required if the entry isn’t the direct child of the search
// bar (as in our main example).
func (b *SearchBarClass) ConnectEntry(entry Editable) {
	var _arg0 *C.GtkSearchBar // out
	var _arg1 *C.GtkEditable  // out

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkEditable)(unsafe.Pointer(entry.Native()))

	C.gtk_search_bar_connect_entry(_arg0, _arg1)
}

// Child gets the child widget of @bar.
func (b *SearchBarClass) Child() Widget {
	var _arg0 *C.GtkSearchBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_search_bar_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

// KeyCaptureWidget gets the widget that @bar is capturing key events from.
func (b *SearchBarClass) KeyCaptureWidget() Widget {
	var _arg0 *C.GtkSearchBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_search_bar_get_key_capture_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
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

// SetChild sets the child widget of @bar.
func (b *SearchBarClass) SetChild(child Widget) {
	var _arg0 *C.GtkSearchBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_search_bar_set_child(_arg0, _arg1)
}

// SetKeyCaptureWidget sets @widget as the widget that @bar will capture key
// events from.
//
// If key events are handled by the search bar, the bar will be shown, and the
// entry populated with the entered text.
//
// Note that despite the name of this function, the events are only 'captured'
// in the bubble phase, which means that editable child widgets of @widget will
// receive text input before it gets captured. If that is not desired, you can
// capture and forward the events yourself with
// [method@Gtk.EventControllerKey.forward].
func (b *SearchBarClass) SetKeyCaptureWidget(widget Widget) {
	var _arg0 *C.GtkSearchBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_search_bar_set_key_capture_widget(_arg0, _arg1)
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

// SetShowCloseButton shows or hides the close button.
//
// Applications that already have a “search” toggle button should not show a
// close button in their search bar, as it duplicates the role of the toggle
// button.
func (b *SearchBarClass) SetShowCloseButton(visible bool) {
	var _arg0 *C.GtkSearchBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkSearchBar)(unsafe.Pointer(b.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_search_bar_set_show_close_button(_arg0, _arg1)
}
