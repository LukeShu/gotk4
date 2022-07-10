// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gchar* _gotk4_gtk3_TranslateFunc(gchar*, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// STOCK_ABOUT: “About” item. ! (help-about.png)
//
// Deprecated: Use named icon &quot;help-about&quot; or the label
// &quot;_About&quot;.
const STOCK_ABOUT = "gtk-about"

// STOCK_ADD: “Add” item and icon.
//
// Deprecated: Use named icon &quot;list-add&quot; or the label
// &quot;_Add&quot;.
const STOCK_ADD = "gtk-add"

// STOCK_APPLY: “Apply” item and icon.
//
// Deprecated: Do not use an icon. Use label &quot;_Apply&quot;.
const STOCK_APPLY = "gtk-apply"

// STOCK_BOLD: “Bold” item and icon.
//
// Deprecated: Use named icon &quot;format-text-bold&quot;.
const STOCK_BOLD = "gtk-bold"

// STOCK_CANCEL: “Cancel” item and icon.
//
// Deprecated: Do not use an icon. Use label &quot;_Cancel&quot;.
const STOCK_CANCEL = "gtk-cancel"

// STOCK_CAPS_LOCK_WARNING “Caps Lock Warning” icon.
//
// Deprecated: Use named icon &quot;dialog-warning-symbolic&quot;.
const STOCK_CAPS_LOCK_WARNING = "gtk-caps-lock-warning"

// STOCK_CDROM: “CD-Rom” item and icon.
//
// Deprecated: Use named icon &quot;media-optical&quot;.
const STOCK_CDROM = "gtk-cdrom"

// STOCK_CLEAR: “Clear” item and icon.
//
// Deprecated: Use named icon &quot;edit-clear&quot;.
const STOCK_CLEAR = "gtk-clear"

// STOCK_CLOSE: “Close” item and icon.
//
// Deprecated: Use named icon &quot;window-close&quot; or the label
// &quot;_Close&quot;.
const STOCK_CLOSE = "gtk-close"

// STOCK_COLOR_PICKER: “Color Picker” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_COLOR_PICKER = "gtk-color-picker"

// STOCK_CONNECT: “Connect” icon.
//
// Deprecated: since version 3.10.
const STOCK_CONNECT = "gtk-connect"

// STOCK_CONVERT: “Convert” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_CONVERT = "gtk-convert"

// STOCK_COPY: “Copy” item and icon.
//
// Deprecated: Use the named icon &quot;edit-copy&quot; or the label
// &quot;_Copy&quot;.
const STOCK_COPY = "gtk-copy"

// STOCK_CUT: “Cut” item and icon.
//
// Deprecated: Use the named icon &quot;edit-cut&quot; or the label
// &quot;Cu_t&quot;.
const STOCK_CUT = "gtk-cut"

// STOCK_DELETE: “Delete” item and icon.
//
// Deprecated: Use the named icon &quot;edit-delete&quot; or the label
// &quot;_Delete&quot;.
const STOCK_DELETE = "gtk-delete"

// STOCK_DIALOG_AUTHENTICATION: “Authentication” item and icon.
//
// Deprecated: Use named icon &quot;dialog-password&quot;.
const STOCK_DIALOG_AUTHENTICATION = "gtk-dialog-authentication"

// STOCK_DIALOG_ERROR: “Error” item and icon.
//
// Deprecated: Use named icon &quot;dialog-error&quot;.
const STOCK_DIALOG_ERROR = "gtk-dialog-error"

// STOCK_DIALOG_INFO: “Information” item and icon.
//
// Deprecated: Use named icon &quot;dialog-information&quot;.
const STOCK_DIALOG_INFO = "gtk-dialog-info"

// STOCK_DIALOG_QUESTION: “Question” item and icon.
//
// Deprecated: Use named icon &quot;dialog-question&quot;.
const STOCK_DIALOG_QUESTION = "gtk-dialog-question"

// STOCK_DIALOG_WARNING: “Warning” item and icon.
//
// Deprecated: Use named icon &quot;dialog-warning&quot;.
const STOCK_DIALOG_WARNING = "gtk-dialog-warning"

// STOCK_DIRECTORY: “Directory” icon.
//
// Deprecated: Use named icon &quot;folder&quot;.
const STOCK_DIRECTORY = "gtk-directory"

// STOCK_DISCARD: “Discard” item.
//
// Deprecated: since version 3.10.
const STOCK_DISCARD = "gtk-discard"

// STOCK_DISCONNECT: “Disconnect” icon.
//
// Deprecated: since version 3.10.
const STOCK_DISCONNECT = "gtk-disconnect"

// STOCK_DND: “Drag-And-Drop” icon.
//
// Deprecated: since version 3.10.
const STOCK_DND = "gtk-dnd"

// STOCK_DND_MULTIPLE: “Drag-And-Drop multiple” icon.
//
// Deprecated: since version 3.10.
const STOCK_DND_MULTIPLE = "gtk-dnd-multiple"

// STOCK_EDIT: “Edit” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_EDIT = "gtk-edit"

// STOCK_EXECUTE: “Execute” item and icon.
//
// Deprecated: Use named icon &quot;system-run&quot;.
const STOCK_EXECUTE = "gtk-execute"

// STOCK_FILE: “File” item and icon.
//
// Since 3.0, this item has a label, before it only had an icon.
//
// Deprecated: Use named icon &quot;text-x-generic&quot;.
const STOCK_FILE = "gtk-file"

// STOCK_FIND: “Find” item and icon.
//
// Deprecated: Use named icon &quot;edit-find&quot;.
const STOCK_FIND = "gtk-find"

// STOCK_FIND_AND_REPLACE: “Find and Replace” item and icon.
//
// Deprecated: Use named icon &quot;edit-find-replace&quot;.
const STOCK_FIND_AND_REPLACE = "gtk-find-and-replace"

// STOCK_FLOPPY: “Floppy” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_FLOPPY = "gtk-floppy"

// STOCK_FULLSCREEN: “Fullscreen” item and icon.
//
// Deprecated: Use named icon &quot;view-fullscreen&quot;.
const STOCK_FULLSCREEN = "gtk-fullscreen"

// STOCK_GOTO_BOTTOM: “Bottom” item and icon.
//
// Deprecated: Use named icon &quot;go-bottom&quot;.
const STOCK_GOTO_BOTTOM = "gtk-goto-bottom"

// STOCK_GOTO_FIRST: “First” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-first&quot;.
const STOCK_GOTO_FIRST = "gtk-goto-first"

// STOCK_GOTO_LAST: “Last” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-last&quot;.
const STOCK_GOTO_LAST = "gtk-goto-last"

// STOCK_GOTO_TOP: “Top” item and icon.
//
// Deprecated: Use named icon &quot;go-top&quot;.
const STOCK_GOTO_TOP = "gtk-goto-top"

// STOCK_GO_BACK: “Back” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-previous&quot;.
const STOCK_GO_BACK = "gtk-go-back"

// STOCK_GO_DOWN: “Down” item and icon.
//
// Deprecated: Use named icon &quot;go-down&quot;.
const STOCK_GO_DOWN = "gtk-go-down"

// STOCK_GO_FORWARD: “Forward” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-next&quot;.
const STOCK_GO_FORWARD = "gtk-go-forward"

// STOCK_GO_UP: “Up” item and icon.
//
// Deprecated: Use named icon &quot;go-up&quot;.
const STOCK_GO_UP = "gtk-go-up"

// STOCK_HARDDISK: “Harddisk” item and icon.
//
// Deprecated: Use named icon &quot;drive-harddisk&quot;.
const STOCK_HARDDISK = "gtk-harddisk"

// STOCK_HELP: “Help” item and icon.
//
// Deprecated: Use named icon &quot;help-browser&quot;.
const STOCK_HELP = "gtk-help"

// STOCK_HOME: “Home” item and icon.
//
// Deprecated: Use named icon &quot;go-home&quot;.
const STOCK_HOME = "gtk-home"

// STOCK_INDENT: “Indent” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;format-indent-more&quot;.
const STOCK_INDENT = "gtk-indent"

// STOCK_INDEX: “Index” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_INDEX = "gtk-index"

// STOCK_INFO: “Info” item and icon.
//
// Deprecated: Use named icon &quot;dialog-information&quot;.
const STOCK_INFO = "gtk-info"

// STOCK_ITALIC: “Italic” item and icon.
//
// Deprecated: Use named icon &quot;format-text-italic&quot;.
const STOCK_ITALIC = "gtk-italic"

// STOCK_JUMP_TO: “Jump to” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;go-jump&quot;.
const STOCK_JUMP_TO = "gtk-jump-to"

// STOCK_JUSTIFY_CENTER: “Center” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-center&quot;.
const STOCK_JUSTIFY_CENTER = "gtk-justify-center"

// STOCK_JUSTIFY_FILL: “Fill” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-fill&quot;.
const STOCK_JUSTIFY_FILL = "gtk-justify-fill"

// STOCK_JUSTIFY_LEFT: “Left” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-left&quot;.
const STOCK_JUSTIFY_LEFT = "gtk-justify-left"

// STOCK_JUSTIFY_RIGHT: “Right” item and icon.
//
// Deprecated: Use named icon &quot;format-justify-right&quot;.
const STOCK_JUSTIFY_RIGHT = "gtk-justify-right"

// STOCK_LEAVE_FULLSCREEN: “Leave Fullscreen” item and icon.
//
// Deprecated: Use named icon &quot;view-restore&quot;.
const STOCK_LEAVE_FULLSCREEN = "gtk-leave-fullscreen"

// STOCK_MEDIA_FORWARD: “Media Forward” item and icon. The icon has an RTL
// variant.
//
// Deprecated: Use named icon &quot;media-seek-forward&quot; or the label
// &quot;_Forward&quot;.
const STOCK_MEDIA_FORWARD = "gtk-media-forward"

// STOCK_MEDIA_NEXT: “Media Next” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;media-skip-forward&quot; or the label
// &quot;_Next&quot;.
const STOCK_MEDIA_NEXT = "gtk-media-next"

// STOCK_MEDIA_PAUSE: “Media Pause” item and icon.
//
// Deprecated: Use named icon &quot;media-playback-pause&quot; or the label
// &quot;P_ause&quot;.
const STOCK_MEDIA_PAUSE = "gtk-media-pause"

// STOCK_MEDIA_PLAY: “Media Play” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;media-playback-start&quot; or the label
// &quot;_Play&quot;.
const STOCK_MEDIA_PLAY = "gtk-media-play"

// STOCK_MEDIA_PREVIOUS: “Media Previous” item and icon. The icon has an RTL
// variant.
//
// Deprecated: Use named icon &quot;media-skip-backward&quot; or the label
// &quot;Pre_vious&quot;.
const STOCK_MEDIA_PREVIOUS = "gtk-media-previous"

// STOCK_MEDIA_RECORD: “Media Record” item and icon.
//
// Deprecated: Use named icon &quot;media-record&quot; or the label
// &quot;_Record&quot;.
const STOCK_MEDIA_RECORD = "gtk-media-record"

// STOCK_MEDIA_REWIND: “Media Rewind” item and icon. The icon has an RTL
// variant.
//
// Deprecated: Use named icon &quot;media-seek-backward&quot; or the label
// &quot;R_ewind&quot;.
const STOCK_MEDIA_REWIND = "gtk-media-rewind"

// STOCK_MEDIA_STOP: “Media Stop” item and icon.
//
// Deprecated: Use named icon &quot;media-playback-stop&quot; or the label
// &quot;_Stop&quot;.
const STOCK_MEDIA_STOP = "gtk-media-stop"

// STOCK_MISSING_IMAGE: “Missing image” icon.
//
// Deprecated: Use named icon &quot;image-missing&quot;.
const STOCK_MISSING_IMAGE = "gtk-missing-image"

// STOCK_NETWORK: “Network” item and icon.
//
// Deprecated: Use named icon &quot;network-workgroup&quot;.
const STOCK_NETWORK = "gtk-network"

// STOCK_NEW: “New” item and icon.
//
// Deprecated: Use named icon &quot;document-new&quot; or the label
// &quot;_New&quot;.
const STOCK_NEW = "gtk-new"

// STOCK_NO: “No” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_NO = "gtk-no"

// STOCK_OK: “OK” item and icon.
//
// Deprecated: Do not use an icon. Use label &quot;_OK&quot;.
const STOCK_OK = "gtk-ok"

// STOCK_OPEN: “Open” item and icon.
//
// Deprecated: Use named icon &quot;document-open&quot; or the label
// &quot;_Open&quot;.
const STOCK_OPEN = "gtk-open"

// STOCK_ORIENTATION_LANDSCAPE: “Landscape Orientation” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_ORIENTATION_LANDSCAPE = "gtk-orientation-landscape"

// STOCK_ORIENTATION_PORTRAIT: “Portrait Orientation” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_ORIENTATION_PORTRAIT = "gtk-orientation-portrait"

// STOCK_ORIENTATION_REVERSE_LANDSCAPE: “Reverse Landscape Orientation” item and
// icon.
//
// Deprecated: since version 3.10.
const STOCK_ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"

// STOCK_ORIENTATION_REVERSE_PORTRAIT: “Reverse Portrait Orientation” item and
// icon.
//
// Deprecated: since version 3.10.
const STOCK_ORIENTATION_REVERSE_PORTRAIT = "gtk-orientation-reverse-portrait"

// STOCK_PAGE_SETUP: “Page Setup” item and icon.
//
// Deprecated: Use named icon &quot;document-page-setup&quot; or the label
// &quot;Page Set_up&quot;.
const STOCK_PAGE_SETUP = "gtk-page-setup"

// STOCK_PASTE: “Paste” item and icon.
//
// Deprecated: Use named icon &quot;edit-paste&quot; or the label
// &quot;_Paste&quot;.
const STOCK_PASTE = "gtk-paste"

// STOCK_PREFERENCES: “Preferences” item and icon.
//
// Deprecated: Use named icon &quot;preferences-system&quot; or the label
// &quot;_Preferences&quot;.
const STOCK_PREFERENCES = "gtk-preferences"

// STOCK_PRINT: “Print” item and icon.
//
// Deprecated: Use named icon &quot;document-print&quot; or the label
// &quot;_Print&quot;.
const STOCK_PRINT = "gtk-print"

// STOCK_PRINT_ERROR: “Print Error” icon.
//
// Deprecated: Use named icon &quot;printer-error&quot;.
const STOCK_PRINT_ERROR = "gtk-print-error"

// STOCK_PRINT_PAUSED: “Print Paused” icon.
//
// Deprecated: since version 3.10.
const STOCK_PRINT_PAUSED = "gtk-print-paused"

// STOCK_PRINT_PREVIEW: “Print Preview” item and icon.
//
// Deprecated: Use label &quot;Pre_view&quot;.
const STOCK_PRINT_PREVIEW = "gtk-print-preview"

// STOCK_PRINT_REPORT: “Print Report” icon.
//
// Deprecated: since version 3.10.
const STOCK_PRINT_REPORT = "gtk-print-report"

// STOCK_PRINT_WARNING: “Print Warning” icon.
//
// Deprecated: since version 3.10.
const STOCK_PRINT_WARNING = "gtk-print-warning"

// STOCK_PROPERTIES: “Properties” item and icon.
//
// Deprecated: Use named icon &quot;document-properties&quot; or the label
// &quot;_Properties&quot;.
const STOCK_PROPERTIES = "gtk-properties"

// STOCK_QUIT: “Quit” item and icon.
//
// Deprecated: Use named icon &quot;application-exit&quot; or the label
// &quot;_Quit&quot;.
const STOCK_QUIT = "gtk-quit"

// STOCK_REDO: “Redo” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;edit-redo&quot; or the label
// &quot;_Redo&quot;.
const STOCK_REDO = "gtk-redo"

// STOCK_REFRESH: “Refresh” item and icon.
//
// Deprecated: Use named icon &quot;view-refresh&quot; or the label
// &quot;_Refresh&quot;.
const STOCK_REFRESH = "gtk-refresh"

// STOCK_REMOVE: “Remove” item and icon.
//
// Deprecated: Use named icon &quot;list-remove&quot; or the label
// &quot;_Remove&quot;.
const STOCK_REMOVE = "gtk-remove"

// STOCK_REVERT_TO_SAVED: “Revert” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;document-revert&quot; or the label
// &quot;_Revert&quot;.
const STOCK_REVERT_TO_SAVED = "gtk-revert-to-saved"

// STOCK_SAVE: “Save” item and icon.
//
// Deprecated: Use named icon &quot;document-save&quot; or the label
// &quot;_Save&quot;.
const STOCK_SAVE = "gtk-save"

// STOCK_SAVE_AS: “Save As” item and icon.
//
// Deprecated: Use named icon &quot;document-save-as&quot; or the label
// &quot;Save _As&quot;.
const STOCK_SAVE_AS = "gtk-save-as"

// STOCK_SELECT_ALL: “Select All” item and icon.
//
// Deprecated: Use named icon &quot;edit-select-all&quot; or the label
// &quot;Select _All&quot;.
const STOCK_SELECT_ALL = "gtk-select-all"

// STOCK_SELECT_COLOR: “Color” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_SELECT_COLOR = "gtk-select-color"

// STOCK_SELECT_FONT: “Font” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_SELECT_FONT = "gtk-select-font"

// STOCK_SORT_ASCENDING: “Ascending” item and icon.
//
// Deprecated: Use named icon &quot;view-sort-ascending&quot;.
const STOCK_SORT_ASCENDING = "gtk-sort-ascending"

// STOCK_SORT_DESCENDING: “Descending” item and icon.
//
// Deprecated: Use named icon &quot;view-sort-descending&quot;.
const STOCK_SORT_DESCENDING = "gtk-sort-descending"

// STOCK_SPELL_CHECK: “Spell Check” item and icon.
//
// Deprecated: Use named icon &quot;tools-check-spelling&quot;.
const STOCK_SPELL_CHECK = "gtk-spell-check"

// STOCK_STOP: “Stop” item and icon.
//
// Deprecated: Use named icon &quot;process-stop&quot; or the label
// &quot;_Stop&quot;.
const STOCK_STOP = "gtk-stop"

// STOCK_STRIKETHROUGH: “Strikethrough” item and icon.
//
// Deprecated: Use named icon &quot;format-text-strikethrough&quot; or the label
// &quot;_Strikethrough&quot;.
const STOCK_STRIKETHROUGH = "gtk-strikethrough"

// STOCK_UNDELETE: “Undelete” item and icon. The icon has an RTL variant.
//
// Deprecated: since version 3.10.
const STOCK_UNDELETE = "gtk-undelete"

// STOCK_UNDERLINE: “Underline” item and icon.
//
// Deprecated: Use named icon &quot;format-text-underline&quot; or the label
// &quot;_Underline&quot;.
const STOCK_UNDERLINE = "gtk-underline"

// STOCK_UNDO: “Undo” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;edit-undo&quot; or the label
// &quot;_Undo&quot;.
const STOCK_UNDO = "gtk-undo"

// STOCK_UNINDENT: “Unindent” item and icon. The icon has an RTL variant.
//
// Deprecated: Use named icon &quot;format-indent-less&quot;.
const STOCK_UNINDENT = "gtk-unindent"

// STOCK_YES: “Yes” item and icon.
//
// Deprecated: since version 3.10.
const STOCK_YES = "gtk-yes"

// STOCK_ZOOM_100: “Zoom 100%” item and icon.
//
// Deprecated: Use named icon &quot;zoom-original&quot; or the label
// &quot;_Normal Size&quot;.
const STOCK_ZOOM_100 = "gtk-zoom-100"

// STOCK_ZOOM_FIT: “Zoom to Fit” item and icon.
//
// Deprecated: Use named icon &quot;zoom-fit-best&quot; or the label &quot;Best
// _Fit&quot;.
const STOCK_ZOOM_FIT = "gtk-zoom-fit"

// STOCK_ZOOM_IN: “Zoom In” item and icon.
//
// Deprecated: Use named icon &quot;zoom-in&quot; or the label &quot;Zoom
// _In&quot;.
const STOCK_ZOOM_IN = "gtk-zoom-in"

// STOCK_ZOOM_OUT: “Zoom Out” item and icon.
//
// Deprecated: Use named icon &quot;zoom-out&quot; or the label &quot;Zoom
// _Out&quot;.
const STOCK_ZOOM_OUT = "gtk-zoom-out"

type Stock = string

// TranslateFunc: function used to translate messages in e.g. IconFactory and
// ActionGroup.
//
// Deprecated: since version 3.10.
type TranslateFunc func(path string) (utf8 string)

//export _gotk4_gtk3_TranslateFunc
func _gotk4_gtk3_TranslateFunc(arg1 *C.gchar, arg2 C.gpointer) (cret *C.gchar) {
	var fn TranslateFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TranslateFunc)
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	utf8 := fn(_path)

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

// StockListIDs retrieves a list of all known stock IDs added to a IconFactory
// or registered with gtk_stock_add(). The list must be freed with
// g_slist_free(), and each string in the list must be freed with g_free().
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - sList: list of known stock IDs.
//
func StockListIDs() []string {
	_info := girepository.MustFind("Gtk", "stock_list_ids")
	_gret := _info.InvokeFunction(nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _sList []string // out

	_sList = make([]string, 0, gextras.SListSize(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	gextras.MoveSList(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))), true, func(v unsafe.Pointer) {
		src := (*C.gchar)(v)
		var dst string // out
		dst = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&src)))))
		defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&src))))
		_sList = append(_sList, dst)
	})

	return _sList
}

// StockLookup fills item with the registered values for stock_id, returning
// TRUE if stock_id was known.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - stockId: stock item name.
//
// The function returns the following values:
//
//    - item: stock item to initialize with values.
//    - ok: TRUE if item was initialized.
//
func StockLookup(stockId string) (*StockItem, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "stock_lookup")
	_gret := _info.InvokeFunction(_args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stockId)

	var _item *StockItem // out
	var _ok bool         // out

	_item = (*StockItem)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _item, _ok
}

// StockSetTranslateFunc sets a function to be used for translating the label of
// a stock item.
//
// If no function is registered for a translation domain, g_dgettext() is used.
//
// The function is used for all stock items whose translation_domain matches
// domain. Note that it is possible to use strings different from the actual
// gettext translation domain of your application for this, as long as your
// TranslateFunc uses the correct domain when calling dgettext(). This can be
// useful, e.g. when dealing with message contexts:
//
//    GtkStockItem items[] = {
//     { MY_ITEM1, NC_("odd items", "Item 1"), 0, 0, "odd-item-domain" },
//     { MY_ITEM2, NC_("even items", "Item 2"), 0, 0, "even-item-domain" },
//    };
//
//    gchar *
//    my_translate_func (const gchar *msgid,
//                       gpointer     data)
//    {
//      gchar *msgctxt = data;
//
//      return (gchar*)g_dpgettext2 (GETTEXT_PACKAGE, msgctxt, msgid);
//    }
//
//    ...
//
//    gtk_stock_add (items, G_N_ELEMENTS (items));
//    gtk_stock_set_translate_func ("odd-item-domain", my_translate_func, "odd items");
//    gtk_stock_set_translate_func ("even-item-domain", my_translate_func, "even items");
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - domain: translation domain for which func shall be used.
//    - fn: TranslateFunc.
//
func StockSetTranslateFunc(domain string, fn TranslateFunc) {
	var _args [4]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_TranslateFunc)
	_args[2] = C.gpointer(gbox.Assign(fn))
	_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	_info := girepository.MustFind("Gtk", "stock_set_translate_func")
	_info.InvokeFunction(_args[:], nil)

	runtime.KeepAlive(domain)
	runtime.KeepAlive(fn)
}

// StockItem: deprecated: since version 3.10.
//
// An instance of this type is always passed by reference.
type StockItem struct {
	*stockItem
}

// stockItem is the struct that's finalized.
type stockItem struct {
	native unsafe.Pointer
}
