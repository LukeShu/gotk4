// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// CurrentFolderFile gets the current folder of chooser as #GFile. See
// gtk_file_chooser_get_current_folder_uri().
//
// The function returns the following values:
//
//    - file for the current folder.
//
func (chooser *FileChooser) CurrentFolderFile() *gio.File {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GFile          // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = C.gtk_file_chooser_get_current_folder_file(_arg0)
	runtime.KeepAlive(chooser)

	var _file *gio.File // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_file = &gio.File{
			Object: obj,
		}
	}

	return _file
}

// File gets the #GFile for the currently selected file in the file selector. If
// multiple files are selected, one of the files will be returned at random.
//
// If the file chooser is in folder mode, this function returns the selected
// folder.
//
// The function returns the following values:
//
//    - file: selected #GFile. You own the returned file; use g_object_unref() to
//      release it.
//
func (chooser *FileChooser) File() *gio.File {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GFile          // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = C.gtk_file_chooser_get_file(_arg0)
	runtime.KeepAlive(chooser)

	var _file *gio.File // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_file = &gio.File{
			Object: obj,
		}
	}

	return _file
}

// Files lists all the selected files and subfolders in the current folder of
// chooser as #GFile. An internal function, see gtk_file_chooser_get_uris().
//
// The function returns the following values:
//
//    - sList: List containing a #GFile for each selected file and subfolder in
//      the current folder. Free the returned list with g_slist_free(), and the
//      files with g_object_unref().
//
func (chooser *FileChooser) Files() []*gio.File {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GSList         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = C.gtk_file_chooser_get_files(_arg0)
	runtime.KeepAlive(chooser)

	var _sList []*gio.File // out

	_sList = make([]*gio.File, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GFile)(v)
		var dst *gio.File // out
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(src))
			dst = &gio.File{
				Object: obj,
			}
		}
		_sList = append(_sList, dst)
	})

	return _sList
}

// PreviewFile gets the #GFile that should be previewed in a custom preview
// Internal function, see gtk_file_chooser_get_preview_uri().
//
// The function returns the following values:
//
//    - file (optional) for the file to preview, or NULL if no file is selected.
//      Free with g_object_unref().
//
func (chooser *FileChooser) PreviewFile() *gio.File {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GFile          // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))

	_cret = C.gtk_file_chooser_get_preview_file(_arg0)
	runtime.KeepAlive(chooser)

	var _file *gio.File // out

	if _cret != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_file = &gio.File{
				Object: obj,
			}
		}
	}

	return _file
}

// SelectFile selects the file referred to by file. An internal function. See
// _gtk_file_chooser_select_uri().
//
// The function takes the following parameters:
//
//    - file to select.
//
func (chooser *FileChooser) SelectFile(file gio.Filer) error {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	C.gtk_file_chooser_select_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(file)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetCurrentFolderFile sets the current folder for chooser from a #GFile.
// Internal function, see gtk_file_chooser_set_current_folder_uri().
//
// The function takes the following parameters:
//
//    - file for the new folder.
//
func (chooser *FileChooser) SetCurrentFolderFile(file gio.Filer) error {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	C.gtk_file_chooser_set_current_folder_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(file)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetFile sets file as the current filename for the file chooser, by changing
// to the file’s parent folder and actually selecting the file in list. If the
// chooser is in GTK_FILE_CHOOSER_ACTION_SAVE mode, the file’s base name will
// also appear in the dialog’s file name entry.
//
// If the file name isn’t in the current folder of chooser, then the current
// folder of chooser will be changed to the folder containing filename. This is
// equivalent to a sequence of gtk_file_chooser_unselect_all() followed by
// gtk_file_chooser_select_filename().
//
// Note that the file must exist, or nothing will be done except for the
// directory change.
//
// If you are implementing a save dialog, you should use this function if you
// already have a file name to which the user may save; for example, when the
// user opens an existing file and then does Save As... If you don’t have a file
// name already — for example, if the user just created a new file and is saving
// it for the first time, do not call this function. Instead, use something
// similar to this:
//
//    if (document_is_new)
//      {
//        // the user just created a new document
//        gtk_file_chooser_set_current_folder_file (chooser, default_file_for_saving);
//        gtk_file_chooser_set_current_name (chooser, "Untitled document");
//      }
//    else
//      {
//        // the user edited an existing document
//        gtk_file_chooser_set_file (chooser, existing_file);
//      }.
//
// The function takes the following parameters:
//
//    - file to set as current.
//
func (chooser *FileChooser) SetFile(file gio.Filer) error {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	C.gtk_file_chooser_set_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(file)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// UnselectFile unselects the file referred to by file. If the file is not in
// the current directory, does not exist, or is otherwise not currently
// selected, does nothing.
//
// The function takes the following parameters:
//
//    - file: #GFile.
//
func (chooser *FileChooser) UnselectFile(file gio.Filer) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(coreglib.InternObject(chooser).Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	C.gtk_file_chooser_unselect_file(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(file)
}
