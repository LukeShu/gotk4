// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
	GTypeStringList   = coreglib.Type(C.gtk_string_list_get_type())
	GTypeStringObject = coreglib.Type(C.gtk_string_object_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStringList, F: marshalStringList},
		coreglib.TypeMarshaler{T: GTypeStringObject, F: marshalStringObject},
	})
}

// StringListOverrides contains methods that are overridable.
type StringListOverrides struct {
}

func defaultStringListOverrides(v *StringList) StringListOverrides {
	return StringListOverrides{}
}

// StringList: GtkStringList is a list model that wraps an array of strings.
//
// The objects in the model have a "string" property.
//
// GtkStringList is well-suited for any place where you would typically use a
// char*[], but need a list model.
//
// # GtkStringList as GtkBuildable
//
// The GtkStringList implementation of the GtkBuildable interface supports
// adding items directly using the <items> element and specifying <item>
// elements for each item. Each <item> element supports the regular translation
// attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying a GtkStringList
//
//    <object class="GtkStringList">
//      <items>
//        <item translatable="yes">Factory</item>
//        <item translatable="yes">Home</item>
//        <item translatable="yes">Subway</item>
//      </items>
//    </object>.
type StringList struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
	Buildable
}

var (
	_ coreglib.Objector = (*StringList)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StringList, *StringListClass, StringListOverrides](
		GTypeStringList,
		initStringListClass,
		wrapStringList,
		defaultStringListOverrides,
	)
}

func initStringListClass(gclass unsafe.Pointer, overrides StringListOverrides, classInitFunc func(*StringListClass)) {
	if classInitFunc != nil {
		class := (*StringListClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStringList(obj *coreglib.Object) *StringList {
	return &StringList{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalStringList(p uintptr) (interface{}, error) {
	return wrapStringList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStringList creates a new GtkStringList with the given strings.
//
// The function takes the following parameters:
//
//   - strings (optional) to put in the model.
//
// The function returns the following values:
//
//   - stringList: new GtkStringList.
//
func NewStringList(strings []string) *StringList {
	var _arg1 **C.char         // out
	var _cret *C.GtkStringList // in

	{
		_arg1 = (**C.char)(C.calloc(C.size_t((len(strings) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(strings)+1)
			var zero *C.char
			out[len(strings)] = zero
			for i := range strings {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(strings[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_string_list_new(_arg1)
	runtime.KeepAlive(strings)

	var _stringList *StringList // out

	_stringList = wrapStringList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stringList
}

// Append appends string to self.
//
// The string will be copied. See gtk.StringList.Take() for a way to avoid that.
//
// The function takes the following parameters:
//
//   - str: string to insert.
//
func (self *StringList) Append(str string) {
	var _arg0 *C.GtkStringList // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkStringList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_string_list_append(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(str)
}

// String gets the string that is at position in self.
//
// If self does not contain position items, NULL is returned.
//
// This function returns the const char *. To get the object wrapping it,
// use g_list_model_get_item().
//
// The function takes the following parameters:
//
//   - position to get the string for.
//
// The function returns the following values:
//
//   - utf8 (optional): string at the given position.
//
func (self *StringList) String(position uint) string {
	var _arg0 *C.GtkStringList // out
	var _arg1 C.guint          // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkStringList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)

	_cret = C.gtk_string_list_get_string(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Remove removes the string at position from self.
//
// position must be smaller than the current length of the list.
//
// The function takes the following parameters:
//
//   - position of the string that is to be removed.
//
func (self *StringList) Remove(position uint) {
	var _arg0 *C.GtkStringList // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkStringList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)

	C.gtk_string_list_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}

// Splice changes self by removing n_removals strings and adding additions to
// it.
//
// This function is more efficient than gtk.StringList.Append() and
// gtk.StringList.Remove(), because it only emits the ::items-changed signal
// once for the change.
//
// This function copies the strings in additions.
//
// The parameters position and n_removals must be correct (ie: position +
// n_removals must be less than or equal to the length of the list at the time
// this function is called).
//
// The function takes the following parameters:
//
//   - position at which to make the change.
//   - nRemovals: number of strings to remove.
//   - additions (optional) strings to add.
//
func (self *StringList) Splice(position, nRemovals uint, additions []string) {
	var _arg0 *C.GtkStringList // out
	var _arg1 C.guint          // out
	var _arg2 C.guint          // out
	var _arg3 **C.char         // out

	_arg0 = (*C.GtkStringList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(nRemovals)
	{
		_arg3 = (**C.char)(C.calloc(C.size_t((len(additions) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(additions)+1)
			var zero *C.char
			out[len(additions)] = zero
			for i := range additions {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(additions[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_string_list_splice(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
	runtime.KeepAlive(nRemovals)
	runtime.KeepAlive(additions)
}

// Take adds string to self at the end, and takes ownership of it.
//
// This variant of gtk.StringList.Append() is convenient for formatting strings:
//
//    gtk_string_list_take (self, g_strdup_print ("d dollars", lots));.
//
// The function takes the following parameters:
//
//   - str: string to insert.
//
func (self *StringList) Take(str string) {
	var _arg0 *C.GtkStringList // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkStringList)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))

	C.gtk_string_list_take(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(str)
}

// StringObjectOverrides contains methods that are overridable.
type StringObjectOverrides struct {
}

func defaultStringObjectOverrides(v *StringObject) StringObjectOverrides {
	return StringObjectOverrides{}
}

// StringObject: GtkStringObject is the type of items in a GtkStringList.
//
// A GtkStringObject is a wrapper around a const char*; it has a
// gtk.StringObject:string property.
type StringObject struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StringObject)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StringObject, *StringObjectClass, StringObjectOverrides](
		GTypeStringObject,
		initStringObjectClass,
		wrapStringObject,
		defaultStringObjectOverrides,
	)
}

func initStringObjectClass(gclass unsafe.Pointer, overrides StringObjectOverrides, classInitFunc func(*StringObjectClass)) {
	if classInitFunc != nil {
		class := (*StringObjectClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStringObject(obj *coreglib.Object) *StringObject {
	return &StringObject{
		Object: obj,
	}
}

func marshalStringObject(p uintptr) (interface{}, error) {
	return wrapStringObject(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStringObject wraps a string in an object for use with GListModel.
//
// The function takes the following parameters:
//
//   - str: string to wrap.
//
// The function returns the following values:
//
//   - stringObject: new GtkStringObject.
//
func NewStringObject(str string) *StringObject {
	var _arg1 *C.char            // out
	var _cret *C.GtkStringObject // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_string_object_new(_arg1)
	runtime.KeepAlive(str)

	var _stringObject *StringObject // out

	_stringObject = wrapStringObject(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stringObject
}

// String returns the string contained in a GtkStringObject.
//
// The function returns the following values:
//
//   - utf8: string of self.
//
func (self *StringObject) String() string {
	var _arg0 *C.GtkStringObject // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkStringObject)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_string_object_get_string(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// StringListClass: instance of this type is always passed by reference.
type StringListClass struct {
	*stringListClass
}

// stringListClass is the struct that's finalized.
type stringListClass struct {
	native *C.GtkStringListClass
}

// StringObjectClass: instance of this type is always passed by reference.
type StringObjectClass struct {
	*stringObjectClass
}

// stringObjectClass is the struct that's finalized.
type stringObjectClass struct {
	native *C.GtkStringObjectClass
}
