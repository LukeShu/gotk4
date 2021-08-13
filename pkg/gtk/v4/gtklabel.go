// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_label_get_type()), F: marshalLabeller},
	})
}

// Label: GtkLabel widget displays a small amount of text.
//
// As the name implies, most labels are used to label another widget such as a
// button.
//
// !An example GtkLabel (label.png)
//
// CSS nodes
//
//    label
//    ├── [selection]
//    ├── [link]
//    ┊
//    ╰── [link]
//
//
// GtkLabel has a single CSS node with the name label. A wide variety of style
// classes may be applied to labels, such as .title, .subtitle, .dim-label, etc.
// In the GtkShortcutsWindow, labels are used with the .keycap style class.
//
// If the label has a selection, it gets a subnode with name selection.
//
// If the label has links, there is one subnode per link. These subnodes carry
// the link or visited state depending on whether they have been visited. In
// this case, label node also gets a .link style class.
//
//
// GtkLabel as GtkBuildable
//
// The GtkLabel implementation of the GtkBuildable interface supports a custom
// <attributes> element, which supports any number of <attribute> elements. The
// <attribute> element has attributes named “name“, “value“, “start“ and “end“
// and allows you to specify pango.Attribute values for this label.
//
// An example of a UI definition fragment specifying Pango attributes:
//
//    <object class="GtkLabel">
//      <attributes>
//        <attribute name="weight" value="PANGO_WEIGHT_BOLD"/>
//        <attribute name="background" value="red" start="5" end="10"/>
//      </attributes>
//    </object>
//
//
// The start and end attributes specify the range of characters to which the
// Pango attribute applies. If start and end are not specified, the attribute is
// applied to the whole text. Note that specifying ranges does not make much
// sense with translatable attributes. Use markup embedded in the translatable
// content instead.
//
//
// Accessibility
//
// GtkLabel uses the K_ACCESSIBLE_ROLE_LABEL role.
//
//
// Mnemonics
//
// Labels may contain “mnemonics”. Mnemonics are underlined characters in the
// label, used for keyboard navigation. Mnemonics are created by providing a
// string with an underscore before the mnemonic character, such as "_File", to
// the functions gtk.Label.NewWithMnemonic or gtk.Label.SetTextWithMnemonic().
//
// Mnemonics automatically activate any activatable widget the label is inside,
// such as a gtk.Button; if the label is not inside the mnemonic’s target
// widget, you have to tell the label about the target using
// gtk.Label.SetMnemonicWidget. Here’s a simple example where the label is
// inside a button:
//
//    // Pressing Alt+H will activate this button
//    GtkWidget *button = gtk_button_new ();
//    GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
//    gtk_button_set_child (GTK_BUTTON (button), label);
//
//
// There’s a convenience function to create buttons with a mnemonic label
// already inside:
//
//    // Pressing Alt+H will activate this button
//    GtkWidget *button = gtk_button_new_with_mnemonic ("_Hello");
//
//
// To create a mnemonic for a widget alongside the label, such as a gtk.Entry,
// you have to point the label at the entry with gtk.Label.SetMnemonicWidget():
//
//    // Pressing Alt+H will focus the entry
//    GtkWidget *entry = gtk_entry_new ();
//    GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
//    gtk_label_set_mnemonic_widget (GTK_LABEL (label), entry);
//
//
// Markup (styled text)
//
// To make it easy to format text in a label (changing colors, fonts, etc.),
// label text can be provided in a simple markup format:
//
// Here’s how to create a label with a small font:
//
//    GtkWidget *label = gtk_label_new (NULL);
//    gtk_label_set_markup (GTK_LABEL (label), "<small>Small text</small>");
//
//
// (See the Pango manual for complete documentation] of available tags,
// pango.ParseMarkup())
//
// The markup passed to gtk_label_set_markup() must be valid; for example,
// literal <, > and & characters must be escaped as &lt;, &gt;, and &amp;. If
// you pass text obtained from the user, file, or a network to
// gtk.Label.SetMarkup(), you’ll want to escape it with g_markup_escape_text()
// or g_markup_printf_escaped().
//
// Markup strings are just a convenient way to set the pango.AttrList on a
// label; gtk.Label.SetAttributes() may be a simpler way to set attributes in
// some cases. Be careful though; pango.AttrList tends to cause
// internationalization problems, unless you’re applying attributes to the
// entire string (i.e. unless you set the range of each attribute to [0,
// G_MAXINT)). The reason is that specifying the start_index and end_index for a
// pango.Attribute requires knowledge of the exact string being displayed, so
// translations will cause problems.
//
//
// Selectable labels
//
// Labels can be made selectable with gtk.Label.SetSelectable(). Selectable
// labels allow the user to copy the label contents to the clipboard. Only
// labels that contain useful-to-copy information — such as error messages —
// should be made selectable.
//
//
// Text layout
//
// A label can contain any number of paragraphs, but will have performance
// problems if it contains more than a small number. Paragraphs are separated by
// newlines or other paragraph separators understood by Pango.
//
// Labels can automatically wrap text if you call gtk.Label.SetWrap().
//
// gtk.Label.SetJustify() sets how the lines in a label align with one another.
// If you want to set how the label as a whole aligns in its available space,
// see the gtk.Widget:halign and gtk.Widget:valign properties.
//
// The gtk.Label:width-chars and gtk.Label:max-width-chars properties can be
// used to control the size allocation of ellipsized or wrapped labels. For
// ellipsizing labels, if either is specified (and less than the actual text
// size), it is used as the minimum width, and the actual text size is used as
// the natural width of the label. For wrapping labels, width-chars is used as
// the minimum width, if specified, and max-width-chars is used as the natural
// width. Even if max-width-chars specified, wrapping labels will be rewrapped
// to use all of the available width.
//
//
// Links
//
// GTK supports markup for clickable hyperlinks in addition to regular Pango
// markup. The markup for links is borrowed from HTML, using the <a> with
// “href“, “title“ and “class“ attributes. GTK renders links similar to the way
// they appear in web browsers, with colored, underlined text. The “title“
// attribute is displayed as a tooltip on the link. The “class“ attribute is
// used as style class on the CSS node for the link.
//
// An example looks like this:
//
//    const char *text =
//    "Go to the"
//    "<a href=\"http://www.gtk.org title=\"&lt;i&gt;Our&lt;/i&gt; website\">"
//    "GTK website</a> for more...";
//    GtkWidget *label = gtk_label_new (NULL);
//    gtk_label_set_markup (GTK_LABEL (label), text);
//
//
// It is possible to implement custom handling for links and their tooltips with
// the gtk.Label::activate-link signal and the gtk.Label.GetCurrentURI()
// function.
type Label struct {
	Widget
}

func wrapLabel(obj *externglib.Object) *Label {
	return &Label{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalLabeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLabel(obj), nil
}

// NewLabel creates a new label with the given text inside it.
//
// You can pass NULL to get an empty label widget.
func NewLabel(str string) *Label {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if str != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_label_new(_arg1)

	runtime.KeepAlive(str)

	var _label *Label // out

	_label = wrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

// NewLabelWithMnemonic creates a new GtkLabel, containing the text in str.
//
// If characters in str are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use '__' (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. The mnemonic key can be used to activate
// another widget, chosen automatically, or explicitly using
// gtk.Label.SetMnemonicWidget().
//
// If gtk.Label.SetMnemonicWidget() is not called, then the first activatable
// ancestor of the GtkLabel will be chosen as the mnemonic widget. For instance,
// if the label is inside a button or menu item, the button or menu item will
// automatically become the mnemonic widget and be activated by the mnemonic.
func NewLabelWithMnemonic(str string) *Label {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if str != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_label_new_with_mnemonic(_arg1)

	runtime.KeepAlive(str)

	var _label *Label // out

	_label = wrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

// Attributes gets the labels attribute list.
//
// This is the pango.AttrList that was set on the label using
// gtk.Label.SetAttributes(), if any. This function does not reflect attributes
// that come from the labels markup (see gtk.Label.SetMarkup()). If you want to
// get the effective attributes for the label, use pango_layout_get_attribute
// (gtk_label_get_layout (self)).
func (self *Label) Attributes() *pango.AttrList {
	var _arg0 *C.GtkLabel      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_attributes(_arg0)

	runtime.KeepAlive(self)

	var _attrList *pango.AttrList // out

	if _cret != nil {
		_attrList = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.pango_attr_list_ref(_cret)
		runtime.SetFinalizer(_attrList, func(v *pango.AttrList) {
			C.pango_attr_list_unref((*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(v))))
		})
	}

	return _attrList
}

// CurrentURI returns the URI for the currently active link in the label.
//
// The active link is the one under the mouse pointer or, in a selectable label,
// the link in which the text cursor is currently positioned.
//
// This function is intended for use in a gtk.Label::activate-link handler or
// for use in a gtk.Widget::query-tooltip handler.
func (self *Label) CurrentURI() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_current_uri(_arg0)

	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Ellipsize returns the ellipsizing position of the label.
//
// See gtk.Label.SetEllipsize().
func (self *Label) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkLabel          // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_ellipsize(_arg0)

	runtime.KeepAlive(self)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// ExtraMenu gets the extra menu model of label.
//
// See gtk.Label.SetExtraMenu().
func (self *Label) ExtraMenu() gio.MenuModeller {
	var _arg0 *C.GtkLabel   // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_extra_menu(_arg0)

	runtime.KeepAlive(self)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		_menuModel = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.MenuModeller)
	}

	return _menuModel
}

// Justify returns the justification of the label.
//
// See gtk.Label.SetJustify().
func (self *Label) Justify() Justification {
	var _arg0 *C.GtkLabel        // out
	var _cret C.GtkJustification // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_justify(_arg0)

	runtime.KeepAlive(self)

	var _justification Justification // out

	_justification = Justification(_cret)

	return _justification
}

// Label fetches the text from a label.
//
// The returned text includes any embedded underlines indicating mnemonics and
// Pango markup. (See gtk.Label.GetText()).
func (self *Label) Label() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_label(_arg0)

	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Layout gets the PangoLayout used to display the label.
//
// The layout is useful to e.g. convert text positions to pixel positions, in
// combination with gtk.Label.GetLayoutOffsets(). The returned layout is owned
// by the label so need not be freed by the caller. The label is free to
// recreate its layout at any time, so it should be considered read-only.
func (self *Label) Layout() *pango.Layout {
	var _arg0 *C.GtkLabel    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_layout(_arg0)

	runtime.KeepAlive(self)

	var _layout *pango.Layout // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	return _layout
}

// LayoutOffsets obtains the coordinates where the label will draw its
// PangoLayout.
//
// The coordinates are useful to convert mouse events into coordinates inside
// the pango.Layout, e.g. to take some action if some part of the label is
// clicked. Remember when using the pango.Layout functions you need to convert
// to and from pixels using PANGO_PIXELS() or pango.SCALE.
func (self *Label) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // in
	var _arg2 C.int       // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	C.gtk_label_get_layout_offsets(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(self)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// Lines gets the number of lines to which an ellipsized, wrapping label should
// be limited.
//
// See gtk.Label.SetLines().
func (self *Label) Lines() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.int       // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_lines(_arg0)

	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MaxWidthChars retrieves the desired maximum width of label, in characters.
//
// See gtk.Label.SetWidthChars().
func (self *Label) MaxWidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.int       // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_max_width_chars(_arg0)

	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MnemonicKeyval: return the mnemonic accelerator.
//
// If the label has been set so that it has a mnemonic key this function returns
// the keyval used for the mnemonic accelerator. If there is no mnemonic set up
// it returns GDK_KEY_VoidSymbol.
func (self *Label) MnemonicKeyval() uint {
	var _arg0 *C.GtkLabel // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_mnemonic_keyval(_arg0)

	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MnemonicWidget retrieves the target of the mnemonic (keyboard shortcut) of
// this label.
//
// See gtk.Label.SetMnemonicWidget().
func (self *Label) MnemonicWidget() Widgetter {
	var _arg0 *C.GtkLabel  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_mnemonic_widget(_arg0)

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Selectable returns whether the label is selectable.
func (self *Label) Selectable() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_selectable(_arg0)

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionBounds gets the selected range of characters in the label.
func (self *Label) SelectionBounds() (start int, end int, ok bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // in
	var _arg2 C.int       // in
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_selection_bounds(_arg0, &_arg1, &_arg2)

	runtime.KeepAlive(self)

	var _start int // out
	var _end int   // out
	var _ok bool   // out

	_start = int(_arg1)
	_end = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _start, _end, _ok
}

// SingleLineMode returns whether the label is in single line mode.
func (self *Label) SingleLineMode() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_single_line_mode(_arg0)

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Text fetches the text from a label.
//
// The returned text is as it appears on screen. This does not include any
// embedded underlines indicating mnemonics or Pango markup. (See
// gtk.Label.GetLabel())
func (self *Label) Text() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_text(_arg0)

	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseMarkup returns whether the label’s text is interpreted as Pango markup.
//
// See gtk.Label.SetUseMarkup().
func (self *Label) UseMarkup() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_use_markup(_arg0)

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an embedded underlines in the label indicate
// mnemonics.
//
// See gtk.Label.SetUseUnderline().
func (self *Label) UseUnderline() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_use_underline(_arg0)

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WidthChars retrieves the desired width of label, in characters.
//
// See gtk.Label.SetWidthChars().
func (self *Label) WidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.int       // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_width_chars(_arg0)

	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Wrap returns whether lines in the label are automatically wrapped.
//
// See gtk.Label.SetWrap().
func (self *Label) Wrap() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_wrap(_arg0)

	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WrapMode returns line wrap mode used by the label.
//
// See gtk.Label.SetWrapMode().
func (self *Label) WrapMode() pango.WrapMode {
	var _arg0 *C.GtkLabel     // out
	var _cret C.PangoWrapMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_wrap_mode(_arg0)

	runtime.KeepAlive(self)

	var _wrapMode pango.WrapMode // out

	_wrapMode = pango.WrapMode(_cret)

	return _wrapMode
}

// XAlign gets the xalign of the label.
//
// See the gtk.Label:xalign property.
func (self *Label) XAlign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.float     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_xalign(_arg0)

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// YAlign gets the yalign of the label.
//
// See the gtk.Label:yalign property.
func (self *Label) YAlign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.float     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_label_get_yalign(_arg0)

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// SelectRegion selects a range of characters in the label, if the label is
// selectable.
//
// See gtk.Label.SetSelectable(). If the label is not selectable, this function
// has no effect. If start_offset or end_offset are -1, then the end of the
// label will be substituted.
func (self *Label) SelectRegion(startOffset int, endOffset int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out
	var _arg2 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(startOffset)
	_arg2 = C.int(endOffset)

	C.gtk_label_select_region(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(startOffset)
	runtime.KeepAlive(endOffset)
}

// SetAttributes: apply attributes to the label text.
//
// The attributes set with this function will be applied and merged with any
// other attributes previously effected by way of the gtk.Label:use-underline or
// gtk.Label:use-markup properties. While it is not recommended to mix markup
// strings with manually set attributes, if you must; know that the attributes
// will be applied to the label after the markup string is parsed.
func (self *Label) SetAttributes(attrs *pango.AttrList) {
	var _arg0 *C.GtkLabel      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if attrs != nil {
		_arg1 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	}

	C.gtk_label_set_attributes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(attrs)
}

// SetEllipsize sets the mode used to ellipsizei the text.
//
// The text will be ellipsized if there is not enough space to render the entire
// string.
func (self *Label) SetEllipsize(mode pango.EllipsizeMode) {
	var _arg0 *C.GtkLabel          // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.gtk_label_set_ellipsize(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(mode)
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// label.
func (self *Label) SetExtraMenu(model gio.MenuModeller) {
	var _arg0 *C.GtkLabel   // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_label_set_extra_menu(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetJustify sets the alignment of the lines in the text of the label relative
// to each other.
//
// GTK_JUSTIFY_LEFT is the default value when the widget is first created with
// gtk.Label.New. If you instead want to set the alignment of the label as a
// whole, use gtk.Widget.SetHAlign() instead. gtk.Label.SetJustify() has no
// effect on labels containing only a single line.
func (self *Label) SetJustify(jtype Justification) {
	var _arg0 *C.GtkLabel        // out
	var _arg1 C.GtkJustification // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkJustification(jtype)

	C.gtk_label_set_justify(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(jtype)
}

// SetLabel sets the text of the label.
//
// The label is interpreted as including embedded underlines and/or Pango markup
// depending on the values of the gtk.Label:use-underline and
// gtk.Label:use-markup properties.
func (self *Label) SetLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(str)
}

// SetLines sets the number of lines to which an ellipsized, wrapping label
// should be limited.
//
// This has no effect if the label is not wrapping or ellipsized. Set this to -1
// if you don’t want to limit the number of lines.
func (self *Label) SetLines(lines int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(lines)

	C.gtk_label_set_lines(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(lines)
}

// SetMarkup sets the labels text and attributes from markup.
//
// The string must be marked up with Pango markup (see pango.ParseMarkup()).
//
// If the str is external data, you may need to escape it with
// g_markup_escape_text() or g_markup_printf_escaped():
//
//    GtkWidget *self = gtk_label_new (NULL);
//    const char *str = "...";
//    const char *format = "<span style=\"italic\">\s</span>";
//    char *markup;
//
//    markup = g_markup_printf_escaped (format, str);
//    gtk_label_set_markup (GTK_LABEL (self), markup);
//    g_free (markup);
//
//
// This function will set the gtk.Label:use-markup property to TRUE as a side
// effect.
//
// If you set the label contents using the gtk.Label:label property you should
// also ensure that you set the gtk.Label:use-markup property accordingly.
//
// See also: gtk.Label.SetText()
func (self *Label) SetMarkup(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(str)
}

// SetMarkupWithMnemonic sets the labels text, attributes and mnemonic from
// markup.
//
// Parses str which is marked up with Pango markup (see pango.ParseMarkup()),
// setting the label’s text and attribute list based on the parse results. If
// characters in str are preceded by an underscore, they are underlined
// indicating that they represent a keyboard accelerator called a mnemonic.
//
// The mnemonic key can be used to activate another widget, chosen
// automatically, or explicitly using methodGtk.Label.set_mnemonic_widget].
func (self *Label) SetMarkupWithMnemonic(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup_with_mnemonic(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(str)
}

// SetMaxWidthChars sets the desired maximum width in characters of label to
// n_chars.
func (self *Label) SetMaxWidthChars(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(nChars)

	C.gtk_label_set_max_width_chars(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(nChars)
}

// SetMnemonicWidget: associate the label with its mnemonic target.
//
// If the label has been set so that it has a mnemonic key (using i.e.
// gtk.Label.SetMarkupWithMnemonic(), gtk.Label.SetTextWithMnemonic(),
// gtk.Label.NewWithMnemonic or the gtk.Label:useUnderline property) the label
// can be associated with a widget that is the target of the mnemonic. When the
// label is inside a widget (like a gtk.Button or a gtk.Notebook tab) it is
// automatically associated with the correct widget, but sometimes (i.e. when
// the target is a gtk.Entry next to the label) you need to set it explicitly
// using this function.
//
// The target widget will be accelerated by emitting the
// gtkwidget::mnemonic-activate signal on it. The default handler for this
// signal will activate the widget if there are no mnemonic collisions and
// toggle focus between the colliding widgets otherwise.
func (self *Label) SetMnemonicWidget(widget Widgetter) {
	var _arg0 *C.GtkLabel  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_label_set_mnemonic_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetSelectable makes text in the label selectable.
//
// Selectable labels allow the user to select text from the label, for
// copy-and-paste.
func (self *Label) SetSelectable(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_selectable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetSingleLineMode sets whether the label is in single line mode.
func (self *Label) SetSingleLineMode(singleLineMode bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if singleLineMode {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_single_line_mode(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(singleLineMode)
}

// SetText sets the text within the GtkLabel widget.
//
// It overwrites any text that was there before.
//
// This function will clear any previously set mnemonic accelerators, and set
// the gtk.Label:use-underline property to FALSE as a side effect.
//
// This function will set the gtk.Label:use-markup property to FALSE as a side
// effect.
//
// See also: gtk.Label.SetMarkup()
func (self *Label) SetText(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(str)
}

// SetTextWithMnemonic sets the label’s text from the string str.
//
// If characters in str are preceded by an underscore, they are underlined
// indicating that they represent a keyboard accelerator called a mnemonic. The
// mnemonic key can be used to activate another widget, chosen automatically, or
// explicitly using gtk.Label.SetMnemonicWidget().
func (self *Label) SetTextWithMnemonic(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text_with_mnemonic(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(str)
}

// SetUseMarkup sets whether the text of the label contains markup.
//
// See gtk.Label.SetMarkup().
func (self *Label) SetUseMarkup(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_markup(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetUseUnderline sets whether underlines in the text indicate mnemonics.
func (self *Label) SetUseUnderline(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetWidthChars sets the desired width in characters of label to n_chars.
func (self *Label) SetWidthChars(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(nChars)

	C.gtk_label_set_width_chars(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(nChars)
}

// SetWrap toggles line wrapping within the GtkLabel widget.
//
// TRUE makes it break lines if text exceeds the widget’s size. FALSE lets the
// text get cut off by the edge of the widget if it exceeds the widget size.
//
// Note that setting line wrapping to TRUE does not make the label wrap at its
// parent container’s width, because GTK widgets conceptually can’t make their
// requisition depend on the parent container’s size. For a label that wraps at
// a specific position, set the label’s width using gtk.Widget.SetSizeRequest().
func (self *Label) SetWrap(wrap bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	if wrap {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_wrap(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(wrap)
}

// SetWrapMode controls how line wrapping is done.
//
// This only affects the label if line wrapping is on. (See gtk.Label.SetWrap())
// The default is PANGO_WRAP_WORD which means wrap on word boundaries.
func (self *Label) SetWrapMode(wrapMode pango.WrapMode) {
	var _arg0 *C.GtkLabel     // out
	var _arg1 C.PangoWrapMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.PangoWrapMode(wrapMode)

	C.gtk_label_set_wrap_mode(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(wrapMode)
}

// SetXAlign sets the xalign of the label.
//
// See the gtk.Label:xalign property.
func (self *Label) SetXAlign(xalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.float(xalign)

	C.gtk_label_set_xalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(xalign)
}

// SetYAlign sets the yalign of the label.
//
// See the gtk.Label:yalign property.
func (self *Label) SetYAlign(yalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(self.Native()))
	_arg1 = C.float(yalign)

	C.gtk_label_set_yalign(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(yalign)
}
