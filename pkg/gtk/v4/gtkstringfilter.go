// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeStringFilterMatchMode = coreglib.Type(C.gtk_string_filter_match_mode_get_type())
	GTypeStringFilter          = coreglib.Type(C.gtk_string_filter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStringFilterMatchMode, F: marshalStringFilterMatchMode},
		coreglib.TypeMarshaler{T: GTypeStringFilter, F: marshalStringFilter},
	})
}

// StringFilterMatchMode specifies how search strings are matched inside text.
type StringFilterMatchMode C.gint

const (
	// StringFilterMatchModeExact: search string and text must match exactly.
	StringFilterMatchModeExact StringFilterMatchMode = iota
	// StringFilterMatchModeSubstring: search string must be contained as a
	// substring inside the text.
	StringFilterMatchModeSubstring
	// StringFilterMatchModePrefix: text must begin with the search string.
	StringFilterMatchModePrefix
)

func marshalStringFilterMatchMode(p uintptr) (interface{}, error) {
	return StringFilterMatchMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for StringFilterMatchMode.
func (s StringFilterMatchMode) String() string {
	switch s {
	case StringFilterMatchModeExact:
		return "Exact"
	case StringFilterMatchModeSubstring:
		return "Substring"
	case StringFilterMatchModePrefix:
		return "Prefix"
	default:
		return fmt.Sprintf("StringFilterMatchMode(%d)", s)
	}
}

// StringFilterOverrides contains methods that are overridable.
type StringFilterOverrides struct {
}

func defaultStringFilterOverrides(v *StringFilter) StringFilterOverrides {
	return StringFilterOverrides{}
}

// StringFilter: GtkStringFilter determines whether to include items by
// comparing strings to a fixed search term.
//
// The strings are obtained from the items by evaluating a GtkExpression set
// with gtk.StringFilter.SetExpression(), and they are compared against a search
// term set with gtk.StringFilter.SetSearch().
//
// GtkStringFilter has several different modes of comparison - it
// can match the whole string, just a prefix, or any substring. Use
// gtk.StringFilter.SetMatchMode() choose a mode.
//
// It is also possible to make case-insensitive comparisons, with
// gtk.StringFilter.SetIgnoreCase().
type StringFilter struct {
	_ [0]func() // equal guard
	Filter
}

var (
	_ coreglib.Objector = (*StringFilter)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*StringFilter, *StringFilterClass, StringFilterOverrides](
		GTypeStringFilter,
		initStringFilterClass,
		wrapStringFilter,
		defaultStringFilterOverrides,
	)
}

func initStringFilterClass(gclass unsafe.Pointer, overrides StringFilterOverrides, classInitFunc func(*StringFilterClass)) {
	if classInitFunc != nil {
		class := (*StringFilterClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapStringFilter(obj *coreglib.Object) *StringFilter {
	return &StringFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalStringFilter(p uintptr) (interface{}, error) {
	return wrapStringFilter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewStringFilter creates a new string filter.
//
// You will want to set up the filter by providing a string to search for and by
// providing a property to look up on the item.
//
// The function takes the following parameters:
//
//   - expression (optional) to evaluate or NULL for none.
//
// The function returns the following values:
//
//   - stringFilter: new GtkStringFilter.
//
func NewStringFilter(expression Expressioner) *StringFilter {
	var _arg1 *C.GtkExpression   // out
	var _cret *C.GtkStringFilter // in

	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(expression).Native()))
	}

	_cret = C.gtk_string_filter_new(_arg1)
	runtime.KeepAlive(expression)

	var _stringFilter *StringFilter // out

	_stringFilter = wrapStringFilter(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stringFilter
}

// Expression gets the expression that the string filter uses to obtain strings
// from items.
//
// The function returns the following values:
//
//   - expression (optional): GtkExpression.
//
func (self *StringFilter) Expression() Expressioner {
	var _arg0 *C.GtkStringFilter // out
	var _cret *C.GtkExpression   // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_string_filter_get_expression(_arg0)
	runtime.KeepAlive(self)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Expressioner)
				return ok
			})
			rv, ok := casted.(Expressioner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// IgnoreCase returns whether the filter ignores case differences.
//
// The function returns the following values:
//
//   - ok: TRUE if the filter ignores case.
//
func (self *StringFilter) IgnoreCase() bool {
	var _arg0 *C.GtkStringFilter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_string_filter_get_ignore_case(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MatchMode returns the match mode that the filter is using.
//
// The function returns the following values:
//
//   - stringFilterMatchMode: match mode of the filter.
//
func (self *StringFilter) MatchMode() StringFilterMatchMode {
	var _arg0 *C.GtkStringFilter         // out
	var _cret C.GtkStringFilterMatchMode // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_string_filter_get_match_mode(_arg0)
	runtime.KeepAlive(self)

	var _stringFilterMatchMode StringFilterMatchMode // out

	_stringFilterMatchMode = StringFilterMatchMode(_cret)

	return _stringFilterMatchMode
}

// Search gets the search term.
//
// The function returns the following values:
//
//   - utf8 (optional): search term.
//
func (self *StringFilter) Search() string {
	var _arg0 *C.GtkStringFilter // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_string_filter_get_search(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetExpression sets the expression that the string filter uses to obtain
// strings from items.
//
// The expression must have a value type of G_TYPE_STRING.
//
// The function takes the following parameters:
//
//   - expression (optional): GtkExpression.
//
func (self *StringFilter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 *C.GtkExpression   // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
	}

	C.gtk_string_filter_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetIgnoreCase sets whether the filter ignores case differences.
//
// The function takes the following parameters:
//
//   - ignoreCase: TRUE to ignore case.
//
func (self *StringFilter) SetIgnoreCase(ignoreCase bool) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if ignoreCase {
		_arg1 = C.TRUE
	}

	C.gtk_string_filter_set_ignore_case(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ignoreCase)
}

// SetMatchMode sets the match mode for the filter.
//
// The function takes the following parameters:
//
//   - mode: new match mode.
//
func (self *StringFilter) SetMatchMode(mode StringFilterMatchMode) {
	var _arg0 *C.GtkStringFilter         // out
	var _arg1 C.GtkStringFilterMatchMode // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkStringFilterMatchMode(mode)

	C.gtk_string_filter_set_match_mode(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(mode)
}

// SetSearch sets the string to search for.
//
// The function takes the following parameters:
//
//   - search (optional): string to search for or NULL to clear the search.
//
func (self *StringFilter) SetSearch(search string) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if search != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(search)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_string_filter_set_search(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(search)
}

// StringFilterClass: instance of this type is always passed by reference.
type StringFilterClass struct {
	*stringFilterClass
}

// stringFilterClass is the struct that's finalized.
type stringFilterClass struct {
	native *C.GtkStringFilterClass
}

func (s *StringFilterClass) ParentClass() *FilterClass {
	valptr := &s.native.parent_class
	var _v *FilterClass // out
	_v = (*FilterClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
