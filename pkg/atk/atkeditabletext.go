// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_editable_text_get_type()), F: marshalEditableTexter},
	})
}

// EditableTextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type EditableTextOverrider interface {
	// CopyText: copy text from start_pos up to, but not including end_pos to
	// the clipboard.
	CopyText(startPos int, endPos int)
	// CutText: copy text from start_pos up to, but not including end_pos to the
	// clipboard and then delete from the widget.
	CutText(startPos int, endPos int)
	// DeleteText: delete text start_pos up to, but not including end_pos.
	DeleteText(startPos int, endPos int)
	// InsertText: insert text at a given position.
	InsertText(_string string, length int, position *int)
	// PasteText: paste text from clipboard to specified position.
	PasteText(position int)
	// SetTextContents: set text contents of text.
	SetTextContents(_string string)
}

// EditableText should be implemented by UI components which contain text which
// the user can edit, via the Object corresponding to that component (see
// Object).
//
// EditableText is a subclass of Text, and as such, an object which implements
// EditableText is by definition an Text implementor as well.
//
// See also: Text
type EditableText struct {
	*externglib.Object
}

// EditableTexter describes EditableText's abstract methods.
type EditableTexter interface {
	externglib.Objector

	// CopyText: copy text from start_pos up to, but not including end_pos to
	// the clipboard.
	CopyText(startPos int, endPos int)
	// CutText: copy text from start_pos up to, but not including end_pos to the
	// clipboard and then delete from the widget.
	CutText(startPos int, endPos int)
	// DeleteText: delete text start_pos up to, but not including end_pos.
	DeleteText(startPos int, endPos int)
	// InsertText: insert text at a given position.
	InsertText(_string string, length int, position *int)
	// PasteText: paste text from clipboard to specified position.
	PasteText(position int)
	// SetTextContents: set text contents of text.
	SetTextContents(_string string)
}

var _ EditableTexter = (*EditableText)(nil)

func wrapEditableText(obj *externglib.Object) *EditableText {
	return &EditableText{
		Object: obj,
	}
}

func marshalEditableTexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEditableText(obj), nil
}

// CopyText: copy text from start_pos up to, but not including end_pos to the
// clipboard.
func (text *EditableText) CopyText(startPos int, endPos int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.atk_editable_text_copy_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// CutText: copy text from start_pos up to, but not including end_pos to the
// clipboard and then delete from the widget.
func (text *EditableText) CutText(startPos int, endPos int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.atk_editable_text_cut_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// DeleteText: delete text start_pos up to, but not including end_pos.
func (text *EditableText) DeleteText(startPos int, endPos int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.atk_editable_text_delete_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// InsertText: insert text at a given position.
func (text *EditableText) InsertText(_string string, length int, position *int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 *C.gchar           // out
	var _arg2 C.gint             // out
	var _arg3 *C.gint            // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(text.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)
	_arg3 = (*C.gint)(unsafe.Pointer(position))

	C.atk_editable_text_insert_text(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(text)
	runtime.KeepAlive(_string)
	runtime.KeepAlive(length)
	runtime.KeepAlive(position)
}

// PasteText: paste text from clipboard to specified position.
func (text *EditableText) PasteText(position int) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 C.gint             // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(position)

	C.atk_editable_text_paste_text(_arg0, _arg1)
	runtime.KeepAlive(text)
	runtime.KeepAlive(position)
}

// SetTextContents: set text contents of text.
func (text *EditableText) SetTextContents(_string string) {
	var _arg0 *C.AtkEditableText // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.AtkEditableText)(unsafe.Pointer(text.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	C.atk_editable_text_set_text_contents(_arg0, _arg1)
	runtime.KeepAlive(text)
	runtime.KeepAlive(_string)
}
