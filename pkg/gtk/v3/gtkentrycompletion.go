// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_EntryCompletion_ConnectNoMatches(gpointer, guintptr);
// extern void _gotk4_gtk3_EntryCompletion_ConnectActionActivated(gpointer, gint, guintptr);
// extern void _gotk4_gtk3_EntryCompletionClass_no_matches(GtkEntryCompletion*);
// extern void _gotk4_gtk3_EntryCompletionClass_action_activated(GtkEntryCompletion*, gint);
// extern gboolean _gotk4_gtk3_EntryCompletion_ConnectMatchSelected(gpointer, GtkTreeModel*, GtkTreeIter*, guintptr);
// extern gboolean _gotk4_gtk3_EntryCompletion_ConnectInsertPrefix(gpointer, gchar*, guintptr);
// extern gboolean _gotk4_gtk3_EntryCompletion_ConnectCursorOnMatch(gpointer, GtkTreeModel*, GtkTreeIter*, guintptr);
// extern gboolean _gotk4_gtk3_EntryCompletionClass_match_selected(GtkEntryCompletion*, GtkTreeModel*, GtkTreeIter*);
// extern gboolean _gotk4_gtk3_EntryCompletionClass_insert_prefix(GtkEntryCompletion*, gchar*);
// extern gboolean _gotk4_gtk3_EntryCompletionClass_cursor_on_match(GtkEntryCompletion*, GtkTreeModel*, GtkTreeIter*);
// gboolean _gotk4_gtk3_EntryCompletion_virtual_cursor_on_match(void* fnptr, GtkEntryCompletion* arg0, GtkTreeModel* arg1, GtkTreeIter* arg2) {
//   return ((gboolean (*)(GtkEntryCompletion*, GtkTreeModel*, GtkTreeIter*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gtk3_EntryCompletion_virtual_insert_prefix(void* fnptr, GtkEntryCompletion* arg0, gchar* arg1) {
//   return ((gboolean (*)(GtkEntryCompletion*, gchar*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gtk3_EntryCompletion_virtual_match_selected(void* fnptr, GtkEntryCompletion* arg0, GtkTreeModel* arg1, GtkTreeIter* arg2) {
//   return ((gboolean (*)(GtkEntryCompletion*, GtkTreeModel*, GtkTreeIter*))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gtk3_EntryCompletion_virtual_action_activated(void* fnptr, GtkEntryCompletion* arg0, gint arg1) {
//   ((void (*)(GtkEntryCompletion*, gint))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_EntryCompletion_virtual_no_matches(void* fnptr, GtkEntryCompletion* arg0) {
//   ((void (*)(GtkEntryCompletion*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeEntryCompletion = coreglib.Type(C.gtk_entry_completion_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEntryCompletion, F: marshalEntryCompletion},
	})
}

// EntryCompletionMatchFunc: function which decides whether the row indicated by
// iter matches a given key, and should be displayed as a possible completion
// for key. Note that key is normalized and case-folded (see g_utf8_normalize()
// and g_utf8_casefold()). If this is not appropriate, match functions have
// access to the unmodified key via gtk_entry_get_text (GTK_ENTRY
// (gtk_entry_completion_get_entry ())).
type EntryCompletionMatchFunc func(completion *EntryCompletion, key string, iter *TreeIter) (ok bool)

// EntryCompletionOverrides contains methods that are overridable.
type EntryCompletionOverrides struct {
	// The function takes the following parameters:
	//
	ActionActivated func(index_ int)
	// The function takes the following parameters:
	//
	//    - model
	//    - iter
	//
	// The function returns the following values:
	//
	CursorOnMatch func(model TreeModeller, iter *TreeIter) bool
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	InsertPrefix func(prefix string) bool
	// The function takes the following parameters:
	//
	//    - model
	//    - iter
	//
	// The function returns the following values:
	//
	MatchSelected func(model TreeModeller, iter *TreeIter) bool
	NoMatches     func()
}

func defaultEntryCompletionOverrides(v *EntryCompletion) EntryCompletionOverrides {
	return EntryCompletionOverrides{
		ActionActivated: v.actionActivated,
		CursorOnMatch:   v.cursorOnMatch,
		InsertPrefix:    v.insertPrefix,
		MatchSelected:   v.matchSelected,
		NoMatches:       v.noMatches,
	}
}

// EntryCompletion is an auxiliary object to be used in conjunction with Entry
// to provide the completion functionality. It implements the CellLayout
// interface, to allow the user to add extra cells to the TreeView with
// completion matches.
//
// “Completion functionality” means that when the user modifies the text in the
// entry, EntryCompletion checks which rows in the model match the current
// content of the entry, and displays a list of matches. By default, the
// matching is done by comparing the entry text case-insensitively against the
// text column of the model (see gtk_entry_completion_set_text_column()), but
// this can be overridden with a custom match function (see
// gtk_entry_completion_set_match_func()).
//
// When the user selects a completion, the content of the entry is updated. By
// default, the content of the entry is replaced by the text column of the
// model, but this can be overridden by connecting to the
// EntryCompletion::match-selected signal and updating the entry in the signal
// handler. Note that you should return TRUE from the signal handler to suppress
// the default behaviour.
//
// To add completion functionality to an entry, use gtk_entry_set_completion().
//
// In addition to regular completion matches, which will be inserted into the
// entry when they are selected, EntryCompletion also allows to display
// “actions” in the popup window. Their appearance is similar to menuitems, to
// differentiate them clearly from completion strings. When an action is
// selected, the EntryCompletion::action-activated signal is emitted.
//
// GtkEntryCompletion uses a TreeModelFilter model to represent the subset of
// the entire model that is currently matching. While the GtkEntryCompletion
// signals EntryCompletion::match-selected and EntryCompletion::cursor-on-match
// take the original model and an iter pointing to that model as arguments,
// other callbacks and signals (such as CellLayoutDataFuncs or
// CellArea::apply-attributes) will generally take the filter model as argument.
// As long as you are only calling gtk_tree_model_get(), this will make no
// difference to you. If for some reason, you need the original model, use
// gtk_tree_model_filter_get_model(). Don’t forget to use
// gtk_tree_model_filter_convert_iter_to_child_iter() to obtain a matching iter.
type EntryCompletion struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
	CellLayout
}

var (
	_ coreglib.Objector = (*EntryCompletion)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EntryCompletion, *EntryCompletionClass, EntryCompletionOverrides](
		GTypeEntryCompletion,
		initEntryCompletionClass,
		wrapEntryCompletion,
		defaultEntryCompletionOverrides,
	)
}

func initEntryCompletionClass(gclass unsafe.Pointer, overrides EntryCompletionOverrides, classInitFunc func(*EntryCompletionClass)) {
	pclass := (*C.GtkEntryCompletionClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeEntryCompletion))))

	if overrides.ActionActivated != nil {
		pclass.action_activated = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_action_activated)
	}

	if overrides.CursorOnMatch != nil {
		pclass.cursor_on_match = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_cursor_on_match)
	}

	if overrides.InsertPrefix != nil {
		pclass.insert_prefix = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_insert_prefix)
	}

	if overrides.MatchSelected != nil {
		pclass.match_selected = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_match_selected)
	}

	if overrides.NoMatches != nil {
		pclass.no_matches = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_no_matches)
	}

	if classInitFunc != nil {
		class := (*EntryCompletionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEntryCompletion(obj *coreglib.Object) *EntryCompletion {
	return &EntryCompletion{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
		CellLayout: CellLayout{
			Object: obj,
		},
	}
}

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	return wrapEntryCompletion(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActionActivated gets emitted when an action is activated.
func (completion *EntryCompletion) ConnectActionActivated(f func(index int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "action-activated", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectActionActivated), f)
}

// ConnectCursorOnMatch gets emitted when a match from the cursor is on a match
// of the list. The default behaviour is to replace the contents of the entry
// with the contents of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk_entry_completion_set_model().
func (completion *EntryCompletion) ConnectCursorOnMatch(f func(model TreeModeller, iter *TreeIter) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "cursor-on-match", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectCursorOnMatch), f)
}

// ConnectInsertPrefix gets emitted when the inline autocompletion is triggered.
// The default behaviour is to make the entry display the whole prefix and
// select the newly inserted part.
//
// Applications may connect to this signal in order to insert only a smaller
// part of the prefix into the entry - e.g. the entry used in the FileChooser
// inserts only the part of the prefix up to the next '/'.
func (completion *EntryCompletion) ConnectInsertPrefix(f func(prefix string) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "insert-prefix", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectInsertPrefix), f)
}

// ConnectMatchSelected gets emitted when a match from the list is selected. The
// default behaviour is to replace the contents of the entry with the contents
// of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk_entry_completion_set_model().
func (completion *EntryCompletion) ConnectMatchSelected(f func(model TreeModeller, iter *TreeIter) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "match-selected", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectMatchSelected), f)
}

// ConnectNoMatches gets emitted when the filter model has zero number of rows
// in completion_complete method. (In other words when GtkEntryCompletion is out
// of suggestions).
func (completion *EntryCompletion) ConnectNoMatches(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "no-matches", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectNoMatches), f)
}

// The function takes the following parameters:
//
func (completion *EntryCompletion) actionActivated(index_ int) {
	gclass := (*C.GtkEntryCompletionClass)(coreglib.PeekParentClass(completion))
	fnarg := gclass.action_activated

	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	_arg1 = C.gint(index_)

	C._gotk4_gtk3_EntryCompletion_virtual_action_activated(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(index_)
}

// The function takes the following parameters:
//
//    - model
//    - iter
//
// The function returns the following values:
//
func (completion *EntryCompletion) cursorOnMatch(model TreeModeller, iter *TreeIter) bool {
	gclass := (*C.GtkEntryCompletionClass)(coreglib.PeekParentClass(completion))
	fnarg := gclass.cursor_on_match

	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.GtkTreeModel       // out
	var _arg2 *C.GtkTreeIter        // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C._gotk4_gtk3_EntryCompletion_virtual_cursor_on_match(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(model)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (completion *EntryCompletion) insertPrefix(prefix string) bool {
	gclass := (*C.GtkEntryCompletionClass)(coreglib.PeekParentClass(completion))
	fnarg := gclass.insert_prefix

	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.gchar              // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(prefix)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gtk3_EntryCompletion_virtual_insert_prefix(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(prefix)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - model
//    - iter
//
// The function returns the following values:
//
func (completion *EntryCompletion) matchSelected(model TreeModeller, iter *TreeIter) bool {
	gclass := (*C.GtkEntryCompletionClass)(coreglib.PeekParentClass(completion))
	fnarg := gclass.match_selected

	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.GtkTreeModel       // out
	var _arg2 *C.GtkTreeIter        // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	_arg2 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C._gotk4_gtk3_EntryCompletion_virtual_match_selected(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(model)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (completion *EntryCompletion) noMatches() {
	gclass := (*C.GtkEntryCompletionClass)(coreglib.PeekParentClass(completion))
	fnarg := gclass.no_matches

	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	C._gotk4_gtk3_EntryCompletion_virtual_no_matches(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(completion)
}

// EntryCompletionClass: instance of this type is always passed by reference.
type EntryCompletionClass struct {
	*entryCompletionClass
}

// entryCompletionClass is the struct that's finalized.
type entryCompletionClass struct {
	native *C.GtkEntryCompletionClass
}
