// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeBookmarkList = coreglib.Type(C.gtk_bookmark_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBookmarkList, F: marshalBookmarkList},
	})
}

// BookmarkListOverrider contains methods that are overridable.
type BookmarkListOverrider interface {
}

// BookmarkList: GtkBookmarkList is a list model that wraps GBookmarkFile.
//
// It presents a GListModel and fills it asynchronously with the GFileInfos
// returned from that function.
//
// The GFileInfos in the list have some attributes in the recent namespace
// added: recent::private (boolean) and recent:applications (stringv).
type BookmarkList struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*BookmarkList)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeBookmarkList,
		GoType:        reflect.TypeOf((*BookmarkList)(nil)),
		InitClass:     initClassBookmarkList,
		FinalizeClass: finalizeClassBookmarkList,
	})
}

func initClassBookmarkList(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitBookmarkList(*BookmarkListClass) }); ok {
		klass := (*BookmarkListClass)(gextras.NewStructNative(gclass))
		goval.InitBookmarkList(klass)
	}
}

func finalizeClassBookmarkList(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeBookmarkList(*BookmarkListClass) }); ok {
		klass := (*BookmarkListClass)(gextras.NewStructNative(gclass))
		goval.FinalizeBookmarkList(klass)
	}
}

func wrapBookmarkList(obj *coreglib.Object) *BookmarkList {
	return &BookmarkList{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalBookmarkList(p uintptr) (interface{}, error) {
	return wrapBookmarkList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBookmarkList creates a new GtkBookmarkList with the given attributes.
//
// The function takes the following parameters:
//
//    - filename (optional): bookmark file to load.
//    - attributes (optional) to query.
//
// The function returns the following values:
//
//    - bookmarkList: new GtkBookmarkList.
//
func NewBookmarkList(filename, attributes string) *BookmarkList {
	var _arg1 *C.char            // out
	var _arg2 *C.char            // out
	var _cret *C.GtkBookmarkList // in

	if filename != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if attributes != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_bookmark_list_new(_arg1, _arg2)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(attributes)

	var _bookmarkList *BookmarkList // out

	_bookmarkList = wrapBookmarkList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bookmarkList
}

// Attributes gets the attributes queried on the children.
//
// The function returns the following values:
//
//    - utf8 (optional): queried attributes.
//
func (self *BookmarkList) Attributes() string {
	var _arg0 *C.GtkBookmarkList // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_bookmark_list_get_attributes(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Filename returns the filename of the bookmark file that this list is loading.
//
// The function returns the following values:
//
//    - utf8: filename of the .xbel file.
//
func (self *BookmarkList) Filename() string {
	var _arg0 *C.GtkBookmarkList // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_bookmark_list_get_filename(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IOPriority gets the IO priority to use while loading file.
//
// The function returns the following values:
//
//    - gint: IO priority.
//
func (self *BookmarkList) IOPriority() int {
	var _arg0 *C.GtkBookmarkList // out
	var _cret C.int              // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_bookmark_list_get_io_priority(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsLoading returns TRUE if the files are currently being loaded.
//
// Files will be added to self from time to time while loading is going on. The
// order in which are added is undefined and may change in between runs.
//
// The function returns the following values:
//
//    - ok: TRUE if self is loading.
//
func (self *BookmarkList) IsLoading() bool {
	var _arg0 *C.GtkBookmarkList // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_bookmark_list_is_loading(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAttributes sets the attributes to be enumerated and starts the
// enumeration.
//
// If attributes is NULL, no attributes will be queried, but a list of Infos
// will still be created.
//
// The function takes the following parameters:
//
//    - attributes (optional) to enumerate.
//
func (self *BookmarkList) SetAttributes(attributes string) {
	var _arg0 *C.GtkBookmarkList // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if attributes != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_bookmark_list_set_attributes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(attributes)
}

// SetIOPriority sets the IO priority to use while loading files.
//
// The default IO priority is G_PRIORITY_DEFAULT.
//
// The function takes the following parameters:
//
//    - ioPriority: IO priority to use.
//
func (self *BookmarkList) SetIOPriority(ioPriority int) {
	var _arg0 *C.GtkBookmarkList // out
	var _arg1 C.int              // out

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(ioPriority)

	C.gtk_bookmark_list_set_io_priority(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ioPriority)
}

// BookmarkListClass: instance of this type is always passed by reference.
type BookmarkListClass struct {
	*bookmarkListClass
}

// bookmarkListClass is the struct that's finalized.
type bookmarkListClass struct {
	native *C.GtkBookmarkListClass
}
