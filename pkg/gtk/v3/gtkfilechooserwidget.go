// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_FileChooserWidget_ConnectDesktopFolder(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectDownFolder(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectHomeFolder(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectLocationPopup(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectLocationPopupOnPaste(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectLocationTogglePopup(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectPlacesShortcut(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectQuickBookmark(gpointer, gint, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectRecentShortcut(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectSearchShortcut(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectShowHidden(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooserWidget_ConnectUpFolder(gpointer, guintptr);
import "C"

// glib.Type values for gtkfilechooserwidget.go.
var GTypeFileChooserWidget = externglib.Type(C.gtk_file_chooser_widget_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeFileChooserWidget, F: marshalFileChooserWidget},
	})
}

// FileChooserWidgetOverrider contains methods that are overridable.
type FileChooserWidgetOverrider interface {
	externglib.Objector
}

// WrapFileChooserWidgetOverrider wraps the FileChooserWidgetOverrider
// interface implementation to access the instance methods.
func WrapFileChooserWidgetOverrider(obj FileChooserWidgetOverrider) *FileChooserWidget {
	return wrapFileChooserWidget(externglib.BaseObject(obj))
}

// FileChooserWidget is a widget for choosing files. It exposes the FileChooser
// interface, and you should use the methods of this interface to interact with
// the widget.
//
//
// CSS nodes
//
// GtkFileChooserWidget has a single CSS node with name filechooser.
type FileChooserWidget struct {
	_ [0]func() // equal guard
	Box

	*externglib.Object
	FileChooser
}

var (
	_ externglib.Objector = (*FileChooserWidget)(nil)
	_ Containerer         = (*FileChooserWidget)(nil)
)

func classInitFileChooserWidgetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFileChooserWidget(obj *externglib.Object) *FileChooserWidget {
	return &FileChooserWidget{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
		Object: obj,
		FileChooser: FileChooser{
			Object: obj,
		},
	}
}

func marshalFileChooserWidget(p uintptr) (interface{}, error) {
	return wrapFileChooserWidget(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_FileChooserWidget_ConnectDesktopFolder
func _gotk4_gtk3_FileChooserWidget_ConnectDesktopFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDesktopFolder signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the user's Desktop folder in the
// file list.
//
// The default binding for this signal is Alt + D.
func (v *FileChooserWidget) ConnectDesktopFolder(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "desktop-folder", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectDesktopFolder), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectDownFolder
func _gotk4_gtk3_FileChooserWidget_ConnectDownFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDownFolder signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser go to a child of the current folder in
// the file hierarchy. The subfolder that will be used is displayed in the path
// bar widget of the file chooser. For example, if the path bar is showing
// "/foo/bar/baz", with bar currently displayed, then this will cause the file
// chooser to switch to the "baz" subfolder.
//
// The default binding for this signal is Alt + Down.
func (v *FileChooserWidget) ConnectDownFolder(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "down-folder", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectDownFolder), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectHomeFolder
func _gotk4_gtk3_FileChooserWidget_ConnectHomeFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectHomeFolder signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the user's home folder in the file
// list.
//
// The default binding for this signal is Alt + Home.
func (v *FileChooserWidget) ConnectHomeFolder(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "home-folder", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectHomeFolder), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectLocationPopup
func _gotk4_gtk3_FileChooserWidget_ConnectLocationPopup(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(path string)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path string))
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_path)
}

// ConnectLocationPopup signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default bindings for this signal are Control + L with a path string of ""
// (the empty string). It is also bound to / with a path string of "/" (a
// slash): this lets you type / and immediately type a path name. On Unix
// systems, this is bound to ~ (tilde) with a path string of "~" itself for
// access to home directories.
func (v *FileChooserWidget) ConnectLocationPopup(f func(path string)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "location-popup", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectLocationPopup), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectLocationPopupOnPaste
func _gotk4_gtk3_FileChooserWidget_ConnectLocationPopupOnPaste(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLocationPopupOnPaste signal is a [keybinding signal][GtkBindingSignal]
// which gets emitted when the user asks for it.
//
// This is used to make the file chooser show a "Location" prompt when the user
// pastes into a FileChooserWidget.
//
// The default binding for this signal is Control + V.
func (v *FileChooserWidget) ConnectLocationPopupOnPaste(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "location-popup-on-paste", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectLocationPopupOnPaste), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectLocationTogglePopup
func _gotk4_gtk3_FileChooserWidget_ConnectLocationTogglePopup(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLocationTogglePopup signal is a [keybinding signal][GtkBindingSignal]
// which gets emitted when the user asks for it.
//
// This is used to toggle the visibility of a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default binding for this signal is Control + L.
func (v *FileChooserWidget) ConnectLocationTogglePopup(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "location-toggle-popup", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectLocationTogglePopup), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectPlacesShortcut
func _gotk4_gtk3_FileChooserWidget_ConnectPlacesShortcut(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPlacesShortcut signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to move the focus to the places sidebar.
//
// The default binding for this signal is Alt + P.
func (v *FileChooserWidget) ConnectPlacesShortcut(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "places-shortcut", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectPlacesShortcut), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectQuickBookmark
func _gotk4_gtk3_FileChooserWidget_ConnectQuickBookmark(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(bookmarkIndex int)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(bookmarkIndex int))
	}

	var _bookmarkIndex int // out

	_bookmarkIndex = int(arg1)

	f(_bookmarkIndex)
}

// ConnectQuickBookmark signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser switch to the bookmark specified in the
// bookmark_index parameter. For example, if you have three bookmarks, you can
// pass 0, 1, 2 to this signal to switch to each of them, respectively.
//
// The default binding for this signal is Alt + 1, Alt + 2, etc. until Alt + 0.
// Note that in the default binding, that Alt + 1 is actually defined to switch
// to the bookmark at index 0, and so on successively; Alt + 0 is defined to
// switch to the bookmark at index 10.
func (v *FileChooserWidget) ConnectQuickBookmark(f func(bookmarkIndex int)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "quick-bookmark", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectQuickBookmark), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectRecentShortcut
func _gotk4_gtk3_FileChooserWidget_ConnectRecentShortcut(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectRecentShortcut signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the Recent location.
//
// The default binding for this signal is Alt + R.
func (v *FileChooserWidget) ConnectRecentShortcut(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "recent-shortcut", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectRecentShortcut), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectSearchShortcut
func _gotk4_gtk3_FileChooserWidget_ConnectSearchShortcut(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectSearchShortcut signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser show the search entry.
//
// The default binding for this signal is Alt + S.
func (v *FileChooserWidget) ConnectSearchShortcut(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "search-shortcut", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectSearchShortcut), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectShowHidden
func _gotk4_gtk3_FileChooserWidget_ConnectShowHidden(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectShowHidden signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user asks for it.
//
// This is used to make the file chooser display hidden files.
//
// The default binding for this signal is Control + H.
func (v *FileChooserWidget) ConnectShowHidden(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "show-hidden", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectShowHidden), f)
}

//export _gotk4_gtk3_FileChooserWidget_ConnectUpFolder
func _gotk4_gtk3_FileChooserWidget_ConnectUpFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectUpFolder signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user asks for it.
//
// This is used to make the file chooser go to the parent of the current folder
// in the file hierarchy.
//
// The default binding for this signal is Alt + Up.
func (v *FileChooserWidget) ConnectUpFolder(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "up-folder", false, unsafe.Pointer(C._gotk4_gtk3_FileChooserWidget_ConnectUpFolder), f)
}

// NewFileChooserWidget creates a new FileChooserWidget. This is a file chooser
// widget that can be embedded in custom windows, and it is the same widget that
// is used by FileChooserDialog.
//
// The function takes the following parameters:
//
//    - action: open or save mode for the widget.
//
// The function returns the following values:
//
//    - fileChooserWidget: new FileChooserWidget.
//
func NewFileChooserWidget(action FileChooserAction) *FileChooserWidget {
	var _arg1 C.GtkFileChooserAction // out
	var _cret *C.GtkWidget           // in

	_arg1 = C.GtkFileChooserAction(action)

	_cret = C.gtk_file_chooser_widget_new(_arg1)
	runtime.KeepAlive(action)

	var _fileChooserWidget *FileChooserWidget // out

	_fileChooserWidget = wrapFileChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fileChooserWidget
}
