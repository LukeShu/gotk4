// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_FileChooser_ConnectUpdatePreview(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooser_ConnectSelectionChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooser_ConnectFileActivated(gpointer, guintptr);
// extern void _gotk4_gtk3_FileChooser_ConnectCurrentFolderChanged(gpointer, guintptr);
// extern GtkFileChooserConfirmation _gotk4_gtk3_FileChooser_ConnectConfirmOverwrite(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeFileChooserAction = coreglib.Type(C.gtk_file_chooser_action_get_type())
	GTypeFileChooserError  = coreglib.Type(C.gtk_file_chooser_error_get_type())
	GTypeFileChooser       = coreglib.Type(C.gtk_file_chooser_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileChooserAction, F: marshalFileChooserAction},
		coreglib.TypeMarshaler{T: GTypeFileChooserError, F: marshalFileChooserError},
		coreglib.TypeMarshaler{T: GTypeFileChooser, F: marshalFileChooser},
	})
}

// FileChooserAction describes whether a FileChooser is being used to open
// existing files or to save to a possibly new file.
type FileChooserAction C.gint

const (
	// FileChooserActionOpen indicates open mode. The file chooser will only let
	// the user pick an existing file.
	FileChooserActionOpen FileChooserAction = iota
	// FileChooserActionSave indicates save mode. The file chooser will let the
	// user pick an existing file, or type in a new filename.
	FileChooserActionSave
	// FileChooserActionSelectFolder indicates an Open mode for selecting
	// folders. The file chooser will let the user pick an existing folder.
	FileChooserActionSelectFolder
	// FileChooserActionCreateFolder indicates a mode for creating a new folder.
	// The file chooser will let the user name an existing or new folder.
	FileChooserActionCreateFolder
)

func marshalFileChooserAction(p uintptr) (interface{}, error) {
	return FileChooserAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FileChooserAction.
func (f FileChooserAction) String() string {
	switch f {
	case FileChooserActionOpen:
		return "Open"
	case FileChooserActionSave:
		return "Save"
	case FileChooserActionSelectFolder:
		return "SelectFolder"
	case FileChooserActionCreateFolder:
		return "CreateFolder"
	default:
		return fmt.Sprintf("FileChooserAction(%d)", f)
	}
}

// FileChooserError: these identify the various errors that can occur while
// calling FileChooser functions.
type FileChooserError C.gint

const (
	// FileChooserErrorNonexistent indicates that a file does not exist.
	FileChooserErrorNonexistent FileChooserError = iota
	// FileChooserErrorBadFilename indicates a malformed filename.
	FileChooserErrorBadFilename
	// FileChooserErrorAlreadyExists indicates a duplicate path (e.g. when
	// adding a bookmark).
	FileChooserErrorAlreadyExists
	// FileChooserErrorIncompleteHostname indicates an incomplete hostname (e.g.
	// "http://foo" without a slash after that).
	FileChooserErrorIncompleteHostname
)

func marshalFileChooserError(p uintptr) (interface{}, error) {
	return FileChooserError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FileChooserError.
func (f FileChooserError) String() string {
	switch f {
	case FileChooserErrorNonexistent:
		return "Nonexistent"
	case FileChooserErrorBadFilename:
		return "BadFilename"
	case FileChooserErrorAlreadyExists:
		return "AlreadyExists"
	case FileChooserErrorIncompleteHostname:
		return "IncompleteHostname"
	default:
		return fmt.Sprintf("FileChooserError(%d)", f)
	}
}

// FileChooser is an interface that can be implemented by file selection
// widgets. In GTK+, the main objects that implement this interface are
// FileChooserWidget, FileChooserDialog, and FileChooserButton. You do not need
// to write an object that implements the FileChooser interface unless you are
// trying to adapt an existing file selector to expose a standard programming
// interface.
//
// FileChooser allows for shortcuts to various places in the filesystem. In the
// default implementation these are displayed in the left pane. It may be a bit
// confusing at first that these shortcuts come from various sources and in
// various flavours, so lets explain the terminology here:
//
// - Bookmarks: are created by the user, by dragging folders from the right pane
// to the left pane, or by using the “Add”. Bookmarks can be renamed and deleted
// by the user.
//
// - Shortcuts: can be provided by the application. For example, a Paint program
// may want to add a shortcut for a Clipart folder. Shortcuts cannot be modified
// by the user.
//
// - Volumes: are provided by the underlying filesystem abstraction. They are
// the “roots” of the filesystem.
//
//
// File Names and Encodings
//
// When the user is finished selecting files in a FileChooser, your program can
// get the selected names either as filenames or as URIs. For URIs, the normal
// escaping rules are applied if the URI contains non-ASCII characters. However,
// filenames are always returned in the character set specified by the
// G_FILENAME_ENCODING environment variable. Please see the GLib documentation
// for more details about this variable.
//
// This means that while you can pass the result of
// gtk_file_chooser_get_filename() to g_open() or g_fopen(), you may not be able
// to directly set it as the text of a Label widget unless you convert it first
// to UTF-8, which all GTK+ widgets expect. You should use g_filename_to_utf8()
// to convert filenames into strings that can be passed to GTK+ widgets.
//
//
// Adding a Preview Widget
//
// You can add a custom preview widget to a file chooser and then get
// notification about when the preview needs to be updated. To install a preview
// widget, use gtk_file_chooser_set_preview_widget(). Then, connect to the
// FileChooser::update-preview signal to get notified when you need to update
// the contents of the preview.
//
// Your callback should use gtk_file_chooser_get_preview_filename() to see what
// needs previewing. Once you have generated the preview for the corresponding
// file, you must call gtk_file_chooser_set_preview_widget_active() with a
// boolean flag that indicates whether your callback could successfully generate
// a preview.
//
// Example: Using a Preview Widget
//
//
//      GtkWidget *toggle;
//
//      ...
//
//      toggle = gtk_check_button_new_with_label ("Open file read-only");
//      gtk_widget_show (toggle);
//      gtk_file_chooser_set_extra_widget (my_file_chooser, toggle);
//    }
//
// If you want to set more than one extra widget in the file chooser, you can a
// container such as a Box or a Grid and include your widgets in it. Then, set
// the container as the whole extra widget.
//
// FileChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type FileChooser struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*FileChooser)(nil)
)

// FileChooserer describes FileChooser's interface methods.
type FileChooserer interface {
	coreglib.Objector

	// AddChoice adds a 'choice' to the file chooser.
	AddChoice(id, label string, options, optionLabels []string)
	// AddFilter adds filter to the list of filters that the user can select
	// between.
	AddFilter(filter *FileFilter)
	// AddShortcutFolder adds a folder to be displayed with the shortcut folders
	// in a file chooser.
	AddShortcutFolder(folder string) error
	// AddShortcutFolderURI adds a folder URI to be displayed with the shortcut
	// folders in a file chooser.
	AddShortcutFolderURI(uri string) error
	// Action gets the type of operation that the file chooser is performing;
	// see gtk_file_chooser_set_action().
	Action() FileChooserAction
	// Choice gets the currently selected option in the 'choice' with the given
	// ID.
	Choice(id string) string
	// CreateFolders gets whether file choser will offer to create new folders.
	CreateFolders() bool
	// CurrentFolder gets the current folder of chooser as a local filename.
	CurrentFolder() string
	// CurrentFolderFile gets the current folder of chooser as #GFile.
	CurrentFolderFile() *gio.File
	// CurrentFolderURI gets the current folder of chooser as an URI.
	CurrentFolderURI() string
	// CurrentName gets the current name in the file selector, as entered by the
	// user in the text entry for “Name”.
	CurrentName() string
	// DoOverwriteConfirmation queries whether a file chooser is set to confirm
	// for overwriting when the user types a file name that already exists.
	DoOverwriteConfirmation() bool
	// ExtraWidget gets the current extra widget; see
	// gtk_file_chooser_set_extra_widget().
	ExtraWidget() Widgetter
	// File gets the #GFile for the currently selected file in the file
	// selector.
	File() *gio.File
	// Filename gets the filename for the currently selected file in the file
	// selector.
	Filename() string
	// Filenames lists all the selected files and subfolders in the current
	// folder of chooser.
	Filenames() []string
	// Files lists all the selected files and subfolders in the current folder
	// of chooser as #GFile.
	Files() []*gio.File
	// Filter gets the current filter; see gtk_file_chooser_set_filter().
	Filter() *FileFilter
	// LocalOnly gets whether only local files can be selected in the file
	// selector.
	LocalOnly() bool
	// PreviewFile gets the #GFile that should be previewed in a custom preview
	// Internal function, see gtk_file_chooser_get_preview_uri().
	PreviewFile() *gio.File
	// PreviewFilename gets the filename that should be previewed in a custom
	// preview widget.
	PreviewFilename() string
	// PreviewURI gets the URI that should be previewed in a custom preview
	// widget.
	PreviewURI() string
	// PreviewWidget gets the current preview widget; see
	// gtk_file_chooser_set_preview_widget().
	PreviewWidget() Widgetter
	// PreviewWidgetActive gets whether the preview widget set by
	// gtk_file_chooser_set_preview_widget() should be shown for the current
	// filename.
	PreviewWidgetActive() bool
	// SelectMultiple gets whether multiple files can be selected in the file
	// selector.
	SelectMultiple() bool
	// ShowHidden gets whether hidden files and folders are displayed in the
	// file selector.
	ShowHidden() bool
	// URI gets the URI for the currently selected file in the file selector.
	URI() string
	// URIs lists all the selected files and subfolders in the current folder of
	// chooser.
	URIs() []string
	// UsePreviewLabel gets whether a stock label should be drawn with the name
	// of the previewed file.
	UsePreviewLabel() bool
	// ListFilters lists the current set of user-selectable filters; see
	// gtk_file_chooser_add_filter(), gtk_file_chooser_remove_filter().
	ListFilters() []*FileFilter
	// ListShortcutFolderURIs queries the list of shortcut folders in the file
	// chooser, as set by gtk_file_chooser_add_shortcut_folder_uri().
	ListShortcutFolderURIs() []string
	// ListShortcutFolders queries the list of shortcut folders in the file
	// chooser, as set by gtk_file_chooser_add_shortcut_folder().
	ListShortcutFolders() []string
	// RemoveChoice removes a 'choice' that has been added with
	// gtk_file_chooser_add_choice().
	RemoveChoice(id string)
	// RemoveFilter removes filter from the list of filters that the user can
	// select between.
	RemoveFilter(filter *FileFilter)
	// RemoveShortcutFolder removes a folder from a file chooser’s list of
	// shortcut folders.
	RemoveShortcutFolder(folder string) error
	// RemoveShortcutFolderURI removes a folder URI from a file chooser’s list
	// of shortcut folders.
	RemoveShortcutFolderURI(uri string) error
	// SelectAll selects all the files in the current folder of a file chooser.
	SelectAll()
	// SelectFile selects the file referred to by file.
	SelectFile(file gio.Filer) error
	// SelectFilename selects a filename.
	SelectFilename(filename string) bool
	// SelectURI selects the file to by uri.
	SelectURI(uri string) bool
	// SetAction sets the type of operation that the chooser is performing; the
	// user interface is adapted to suit the selected action.
	SetAction(action FileChooserAction)
	// SetChoice selects an option in a 'choice' that has been added with
	// gtk_file_chooser_add_choice().
	SetChoice(id, option string)
	// SetCreateFolders sets whether file choser will offer to create new
	// folders.
	SetCreateFolders(createFolders bool)
	// SetCurrentFolder sets the current folder for chooser from a local
	// filename.
	SetCurrentFolder(filename string) bool
	// SetCurrentFolderFile sets the current folder for chooser from a #GFile.
	SetCurrentFolderFile(file gio.Filer) error
	// SetCurrentFolderURI sets the current folder for chooser from an URI.
	SetCurrentFolderURI(uri string) bool
	// SetCurrentName sets the current name in the file selector, as if entered
	// by the user.
	SetCurrentName(name string)
	// SetDoOverwriteConfirmation sets whether a file chooser in
	// GTK_FILE_CHOOSER_ACTION_SAVE mode will present a confirmation dialog if
	// the user types a file name that already exists.
	SetDoOverwriteConfirmation(doOverwriteConfirmation bool)
	// SetExtraWidget sets an application-supplied widget to provide extra
	// options to the user.
	SetExtraWidget(extraWidget Widgetter)
	// SetFile sets file as the current filename for the file chooser, by
	// changing to the file’s parent folder and actually selecting the file in
	// list.
	SetFile(file gio.Filer) error
	// SetFilename sets filename as the current filename for the file chooser,
	// by changing to the file’s parent folder and actually selecting the file
	// in list; all other files will be unselected.
	SetFilename(filename string) bool
	// SetFilter sets the current filter; only the files that pass the filter
	// will be displayed.
	SetFilter(filter *FileFilter)
	// SetLocalOnly sets whether only local files can be selected in the file
	// selector.
	SetLocalOnly(localOnly bool)
	// SetPreviewWidget sets an application-supplied widget to use to display a
	// custom preview of the currently selected file.
	SetPreviewWidget(previewWidget Widgetter)
	// SetPreviewWidgetActive sets whether the preview widget set by
	// gtk_file_chooser_set_preview_widget() should be shown for the current
	// filename.
	SetPreviewWidgetActive(active bool)
	// SetSelectMultiple sets whether multiple files can be selected in the file
	// selector.
	SetSelectMultiple(selectMultiple bool)
	// SetShowHidden sets whether hidden files and folders are displayed in the
	// file selector.
	SetShowHidden(showHidden bool)
	// SetURI sets the file referred to by uri as the current file for the file
	// chooser, by changing to the URI’s parent folder and actually selecting
	// the URI in the list.
	SetURI(uri string) bool
	// SetUsePreviewLabel sets whether the file chooser should display a stock
	// label with the name of the file that is being previewed; the default is
	// TRUE.
	SetUsePreviewLabel(useLabel bool)
	// UnselectAll unselects all the files in the current folder of a file
	// chooser.
	UnselectAll()
	// UnselectFile unselects the file referred to by file.
	UnselectFile(file gio.Filer)
	// UnselectFilename unselects a currently selected filename.
	UnselectFilename(filename string)
	// UnselectURI unselects the file referred to by uri.
	UnselectURI(uri string)

	// Confirm-overwrite: this signal gets emitted whenever it is appropriate to
	// present a confirmation dialog when the user has selected a file name that
	// already exists.
	ConnectConfirmOverwrite(func() (fileChooserConfirmation FileChooserConfirmation)) coreglib.SignalHandle
	// Current-folder-changed: this signal is emitted when the current folder in
	// a FileChooser changes.
	ConnectCurrentFolderChanged(func()) coreglib.SignalHandle
	// File-activated: this signal is emitted when the user "activates" a file
	// in the file chooser.
	ConnectFileActivated(func()) coreglib.SignalHandle
	// Selection-changed: this signal is emitted when there is a change in the
	// set of selected files in a FileChooser.
	ConnectSelectionChanged(func()) coreglib.SignalHandle
	// Update-preview: this signal is emitted when the preview in a file chooser
	// should be regenerated.
	ConnectUpdatePreview(func()) coreglib.SignalHandle
}

var _ FileChooserer = (*FileChooser)(nil)

func wrapFileChooser(obj *coreglib.Object) *FileChooser {
	return &FileChooser{
		Object: obj,
	}
}

func marshalFileChooser(p uintptr) (interface{}, error) {
	return wrapFileChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectConfirmOverwrite: this signal gets emitted whenever it is appropriate
// to present a confirmation dialog when the user has selected a file name that
// already exists. The signal only gets emitted when the file chooser is in
// GTK_FILE_CHOOSER_ACTION_SAVE mode.
//
// Most applications just need to turn on the
// FileChooser:do-overwrite-confirmation property (or call the
// gtk_file_chooser_set_do_overwrite_confirmation() function), and they will
// automatically get a stock confirmation dialog. Applications which need to
// customize this behavior should do that, and also connect to the
// FileChooser::confirm-overwrite signal.
//
// A signal handler for this signal must return a FileChooserConfirmation value,
// which indicates the action to take. If the handler determines that the user
// wants to select a different filename, it should return
// GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN. If it determines that the user is
// satisfied with his choice of file name, it should return
// GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME. On the other hand, if it
// determines that the stock confirmation dialog should be used, it should
// return GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM. The following example
// illustrates this.
//
// Custom confirmation
//
//    static GtkFileChooserConfirmation
//    confirm_overwrite_callback (GtkFileChooser *chooser, gpointer data)
//    {
//      char *uri;
//
//      uri = gtk_file_chooser_get_uri (chooser);
//
//      if (is_uri_read_only (uri))
//        {
//          if (user_wants_to_replace_read_only_file (uri))
//            return GTK_FILE_CHOOSER_CONFIRMATION_ACCEPT_FILENAME;
//          else
//            return GTK_FILE_CHOOSER_CONFIRMATION_SELECT_AGAIN;
//        } else
//          return GTK_FILE_CHOOSER_CONFIRMATION_CONFIRM; // fall back to the default dialog
//    }
//
//    ...
//
//    chooser = gtk_file_chooser_dialog_new (...);
//
//    gtk_file_chooser_set_do_overwrite_confirmation (GTK_FILE_CHOOSER (dialog), TRUE);
//    g_signal_connect (chooser, "confirm-overwrite",
//                      G_CALLBACK (confirm_overwrite_callback), NULL);
//
//    if (gtk_dialog_run (chooser) == GTK_RESPONSE_ACCEPT)
//            save_to_file (gtk_file_chooser_get_filename (GTK_FILE_CHOOSER (chooser));
//
//    gtk_widget_destroy (chooser);.
func (chooser *FileChooser) ConnectConfirmOverwrite(f func() (fileChooserConfirmation FileChooserConfirmation)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "confirm-overwrite", false, unsafe.Pointer(C._gotk4_gtk3_FileChooser_ConnectConfirmOverwrite), f)
}

// ConnectCurrentFolderChanged: this signal is emitted when the current folder
// in a FileChooser changes. This can happen due to the user performing some
// action that changes folders, such as selecting a bookmark or visiting a
// folder on the file list. It can also happen as a result of calling a function
// to explicitly change the current folder in a file chooser.
//
// Normally you do not need to connect to this signal, unless you need to keep
// track of which folder a file chooser is showing.
//
// See also: gtk_file_chooser_set_current_folder(),
// gtk_file_chooser_get_current_folder(),
// gtk_file_chooser_set_current_folder_uri(),
// gtk_file_chooser_get_current_folder_uri().
func (chooser *FileChooser) ConnectCurrentFolderChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "current-folder-changed", false, unsafe.Pointer(C._gotk4_gtk3_FileChooser_ConnectCurrentFolderChanged), f)
}

// ConnectFileActivated: this signal is emitted when the user "activates" a file
// in the file chooser. This can happen by double-clicking on a file in the file
// list, or by pressing Enter.
//
// Normally you do not need to connect to this signal. It is used internally by
// FileChooserDialog to know when to activate the default button in the dialog.
//
// See also: gtk_file_chooser_get_filename(), gtk_file_chooser_get_filenames(),
// gtk_file_chooser_get_uri(), gtk_file_chooser_get_uris().
func (chooser *FileChooser) ConnectFileActivated(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "file-activated", false, unsafe.Pointer(C._gotk4_gtk3_FileChooser_ConnectFileActivated), f)
}

// ConnectSelectionChanged: this signal is emitted when there is a change in the
// set of selected files in a FileChooser. This can happen when the user
// modifies the selection with the mouse or the keyboard, or when explicitly
// calling functions to change the selection.
//
// Normally you do not need to connect to this signal, as it is easier to wait
// for the file chooser to finish running, and then to get the list of selected
// files using the functions mentioned below.
//
// See also: gtk_file_chooser_select_filename(),
// gtk_file_chooser_unselect_filename(), gtk_file_chooser_get_filename(),
// gtk_file_chooser_get_filenames(), gtk_file_chooser_select_uri(),
// gtk_file_chooser_unselect_uri(), gtk_file_chooser_get_uri(),
// gtk_file_chooser_get_uris().
func (chooser *FileChooser) ConnectSelectionChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "selection-changed", false, unsafe.Pointer(C._gotk4_gtk3_FileChooser_ConnectSelectionChanged), f)
}

// ConnectUpdatePreview: this signal is emitted when the preview in a file
// chooser should be regenerated. For example, this can happen when the
// currently selected file changes. You should use this signal if you want your
// file chooser to have a preview widget.
//
// Once you have installed a preview widget with
// gtk_file_chooser_set_preview_widget(), you should update it when this signal
// is emitted. You can use the functions gtk_file_chooser_get_preview_filename()
// or gtk_file_chooser_get_preview_uri() to get the name of the file to preview.
// Your widget may not be able to preview all kinds of files; your callback must
// call gtk_file_chooser_set_preview_widget_active() to inform the file chooser
// about whether the preview was generated successfully or not.
//
// Please see the example code in [Using a Preview
// Widget][gtkfilechooser-preview].
//
// See also: gtk_file_chooser_set_preview_widget(),
// gtk_file_chooser_set_preview_widget_active(),
// gtk_file_chooser_set_use_preview_label(),
// gtk_file_chooser_get_preview_filename(), gtk_file_chooser_get_preview_uri().
func (chooser *FileChooser) ConnectUpdatePreview(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(chooser, "update-preview", false, unsafe.Pointer(C._gotk4_gtk3_FileChooser_ConnectUpdatePreview), f)
}

// UsePreviewLabel gets whether a stock label should be drawn with the name of
// the previewed file. See gtk_file_chooser_set_use_preview_label().
//
// The function returns the following values:
//
//    - ok: TRUE if the file chooser is set to display a label with the name of
//      the previewed file, FALSE otherwise.
//
func (chooser *FileChooser) UsePreviewLabel() bool {
	var _arg0 *C.GtkFileChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = C.gtk_file_chooser_get_use_preview_label(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
