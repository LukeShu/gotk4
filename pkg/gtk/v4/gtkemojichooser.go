// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_EmojiChooser_ConnectEmojiPicked(gpointer, gchar*, guintptr);
import "C"

// glib.Type values for gtkemojichooser.go.
var GTypeEmojiChooser = externglib.Type(C.gtk_emoji_chooser_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeEmojiChooser, F: marshalEmojiChooser},
	})
}

// EmojiChooserOverrider contains methods that are overridable.
type EmojiChooserOverrider interface {
	externglib.Objector
}

// WrapEmojiChooserOverrider wraps the EmojiChooserOverrider
// interface implementation to access the instance methods.
func WrapEmojiChooserOverrider(obj EmojiChooserOverrider) *EmojiChooser {
	return wrapEmojiChooser(externglib.BaseObject(obj))
}

// EmojiChooser: GtkEmojiChooser is used by text widgets such as GtkEntry or
// GtkTextView to let users insert Emoji characters.
//
// !An example GtkEmojiChooser (emojichooser.png)
//
// GtkEmojiChooser emits the gtk.EmojiChooser::emoji-picked signal when an Emoji
// is selected.
//
// CSS nodes
//
//    popover
//    ├── box.emoji-searchbar
//    │   ╰── entry.search
//    ╰── box.emoji-toolbar
//        ├── button.image-button.emoji-section
//        ├── ...
//        ╰── button.image-button.emoji-section
//
//
// Every GtkEmojiChooser consists of a main node called popover. The contents of
// the popover are largely implementation defined and supposed to inherit
// general styles. The top searchbar used to search emoji and gets the
// .emoji-searchbar style class itself. The bottom toolbar used to switch
// between different emoji categories consists of buttons with the
// .emoji-section style class and gets the .emoji-toolbar style class itself.
type EmojiChooser struct {
	_ [0]func() // equal guard
	Popover
}

var (
	_ Widgetter           = (*EmojiChooser)(nil)
	_ externglib.Objector = (*EmojiChooser)(nil)
)

func classInitEmojiChooserer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapEmojiChooser(obj *externglib.Object) *EmojiChooser {
	return &EmojiChooser{
		Popover: Popover{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			NativeSurface: NativeSurface{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					Accessible: Accessible{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					ConstraintTarget: ConstraintTarget{
						Object: obj,
					},
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalEmojiChooser(p uintptr) (interface{}, error) {
	return wrapEmojiChooser(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_EmojiChooser_ConnectEmojiPicked
func _gotk4_gtk4_EmojiChooser_ConnectEmojiPicked(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(text string)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(text string))
	}

	var _text string // out

	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_text)
}

// ConnectEmojiPicked is emitted when the user selects an Emoji.
func (v *EmojiChooser) ConnectEmojiPicked(f func(text string)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(v, "emoji-picked", false, unsafe.Pointer(C._gotk4_gtk4_EmojiChooser_ConnectEmojiPicked), f)
}

// NewEmojiChooser creates a new GtkEmojiChooser.
//
// The function returns the following values:
//
//    - emojiChooser: new GtkEmojiChooser.
//
func NewEmojiChooser() *EmojiChooser {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_emoji_chooser_new()

	var _emojiChooser *EmojiChooser // out

	_emojiChooser = wrapEmojiChooser(externglib.Take(unsafe.Pointer(_cret)))

	return _emojiChooser
}
