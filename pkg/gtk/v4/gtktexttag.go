// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeTextTag returns the GType for the type TextTag.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTextTag() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TextTag").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTextTag)
	return gtype
}

// TextTagOverrider contains methods that are overridable.
type TextTagOverrider interface {
}

// TextTag: tag that can be applied to text contained in a GtkTextBuffer.
//
// You may wish to begin by reading the text widget conceptual overview
// (section-text-widget.html), which gives an overview of all the objects and
// data types related to the text widget and how they work together.
//
// Tags should be in the gtk.TextTagTable for a given GtkTextBuffer before using
// them with that buffer.
//
// gtk.TextBuffer.CreateTag() is the best way to create tags. See “gtk4-demo”
// for numerous examples.
//
// For each property of GtkTextTag, there is a “set” property, e.g. “font-set”
// corresponds to “font”. These “set” properties reflect whether a property has
// been set or not.
//
// They are maintained by GTK and you should not set them independently.
type TextTag struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TextTag)(nil)
)

func classInitTextTagger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTextTag(obj *coreglib.Object) *TextTag {
	return &TextTag{
		Object: obj,
	}
}

func marshalTextTag(p uintptr) (interface{}, error) {
	return wrapTextTag(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTextTag creates a GtkTextTag.
//
// The function takes the following parameters:
//
//    - name (optional): tag name, or NULL.
//
// The function returns the following values:
//
//    - textTag: new GtkTextTag.
//
func NewTextTag(name string) *TextTag {
	var _args [1]girepository.Argument

	if name != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_args[0]))
	}

	_gret := girepository.MustFind("Gtk", "TextTag").InvokeMethod("new_TextTag", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)

	var _textTag *TextTag // out

	_textTag = wrapTextTag(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textTag
}

// Changed emits the gtk.TextTagTable::tag-changed signal on the GtkTextTagTable
// where the tag is included.
//
// The signal is already emitted when setting a GtkTextTag property. This
// function is useful for a GtkTextTag subclass.
//
// The function takes the following parameters:
//
//    - sizeChanged: whether the change affects the GtkTextView layout.
//
func (tag *TextTag) Changed(sizeChanged bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(tag).Native()))
	if sizeChanged {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "TextTag").InvokeMethod("changed", _args[:], nil)

	runtime.KeepAlive(tag)
	runtime.KeepAlive(sizeChanged)
}

// Priority: get the tag priority.
//
// The function returns the following values:
//
//    - gint tag’s priority.
//
func (tag *TextTag) Priority() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(tag).Native()))

	_gret := girepository.MustFind("Gtk", "TextTag").InvokeMethod("get_priority", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(tag)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// SetPriority sets the priority of a GtkTextTag.
//
// Valid priorities start at 0 and go to one less than
// gtk.TextTagTable.GetSize(). Each tag in a table has a unique priority;
// setting the priority of one tag shifts the priorities of all the other tags
// in the table to maintain a unique priority for each tag.
//
// Higher priority tags “win” if two tags both set the same text attribute. When
// adding a tag to a tag table, it will be assigned the highest priority in the
// table by default; so normally the precedence of a set of tags is the order in
// which they were added to the table, or created with
// gtk.TextBuffer.CreateTag(), which adds the tag to the buffer’s table
// automatically.
//
// The function takes the following parameters:
//
//    - priority: new priority.
//
func (tag *TextTag) SetPriority(priority int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(tag).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(priority)

	girepository.MustFind("Gtk", "TextTag").InvokeMethod("set_priority", _args[:], nil)

	runtime.KeepAlive(tag)
	runtime.KeepAlive(priority)
}
