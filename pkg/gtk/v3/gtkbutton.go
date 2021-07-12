// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_get_type()), F: marshalButtoner},
	})
}

// ButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ButtonOverrider interface {
	Activate()
	// Clicked emits a Button::clicked signal to the given Button.
	Clicked()
	// Enter emits a Button::enter signal to the given Button.
	//
	// Deprecated: Use the Widget::enter-notify-event signal.
	Enter()
	// Leave emits a Button::leave signal to the given Button.
	//
	// Deprecated: Use the Widget::leave-notify-event signal.
	Leave()
	// Pressed emits a Button::pressed signal to the given Button.
	//
	// Deprecated: Use the Widget::button-press-event signal.
	Pressed()
	// Released emits a Button::released signal to the given Button.
	//
	// Deprecated: Use the Widget::button-release-event signal.
	Released()
}

// Buttoner describes Button's methods.
type Buttoner interface {
	// Clicked emits a Button::clicked signal to the given Button.
	Clicked()
	// Enter emits a Button::enter signal to the given Button.
	Enter()
	// Alignment gets the alignment of the child in the button.
	Alignment() (xalign float32, yalign float32)
	// AlwaysShowImage returns whether the button will ignore the
	// Settings:gtk-button-images setting and always show the image, if
	// available.
	AlwaysShowImage() bool
	// EventWindow returns the button’s event window if it is realized, nil
	// otherwise.
	EventWindow() *gdk.Window
	// FocusOnClick returns whether the button grabs focus when it is clicked
	// with the mouse.
	FocusOnClick() bool
	// Image gets the widget that is currenty set as the image of @button.
	Image() *Widget
	// ImagePosition gets the position of the image relative to the text inside
	// the button.
	ImagePosition() PositionType
	// Label fetches the text from the label of the button, as set by
	// gtk_button_set_label().
	Label() string
	// Relief returns the current relief style of the given Button.
	Relief() ReliefStyle
	// UseStock returns whether the button label is a stock item.
	UseStock() bool
	// UseUnderline returns whether an embedded underline in the button label
	// indicates a mnemonic.
	UseUnderline() bool
	// Leave emits a Button::leave signal to the given Button.
	Leave()
	// Pressed emits a Button::pressed signal to the given Button.
	Pressed()
	// Released emits a Button::released signal to the given Button.
	Released()
	// SetAlignment sets the alignment of the child.
	SetAlignment(xalign float32, yalign float32)
	// SetAlwaysShowImage: if true, the button will ignore the
	// Settings:gtk-button-images setting and always show the image, if
	// available.
	SetAlwaysShowImage(alwaysShow bool)
	// SetFocusOnClick sets whether the button will grab focus when it is
	// clicked with the mouse.
	SetFocusOnClick(focusOnClick bool)
	// SetImage: set the image of @button to the given widget.
	SetImage(image Widgeter)
	// SetImagePosition sets the position of the image relative to the text
	// inside the button.
	SetImagePosition(position PositionType)
	// SetLabel sets the text of the label of the button to @str.
	SetLabel(label string)
	// SetRelief sets the relief style of the edges of the given Button widget.
	SetRelief(relief ReliefStyle)
	// SetUseStock: if true, the label set on the button is used as a stock id
	// to select the stock item for the button.
	SetUseStock(useStock bool)
	// SetUseUnderline: if true, an underline in the text of the button label
	// indicates the next character should be used for the mnemonic accelerator
	// key.
	SetUseUnderline(useUnderline bool)
}

// Button widget is generally used to trigger a callback function that is called
// when the button is pressed. The various signals and how to use them are
// outlined below.
//
// The Button widget can hold any valid child widget. That is, it can hold
// almost any other standard Widget. The most commonly used child is the Label.
//
//
// CSS nodes
//
// GtkButton has a single CSS node with name button. The node will get the style
// classes .image-button or .text-button, if the content is just an image or
// label, respectively. It may also receive the .flat style class.
//
// Other style classes that are commonly used with GtkButton include
// .suggested-action and .destructive-action. In special cases, buttons can be
// made round by adding the .circular style class.
//
// Button-like widgets like ToggleButton, MenuButton, VolumeButton, LockButton,
// ColorButton, FontButton or FileChooserButton use style classes such as
// .toggle, .popup, .scale, .lock, .color, .font, .file to differentiate
// themselves from a plain GtkButton.
type Button struct {
	Bin

	Actionable
	Activatable
}

var (
	_ Buttoner        = (*Button)(nil)
	_ gextras.Nativer = (*Button)(nil)
)

func wrapButton(obj *externglib.Object) Buttoner {
	return &Button{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
		Activatable: Activatable{
			Object: obj,
		},
	}
}

func marshalButtoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapButton(obj), nil
}

// NewButton creates a new Button widget. To add a child widget to the button,
// use gtk_container_add().
func NewButton() *Button {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_button_new()

	var _button *Button // out

	_button = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Button)

	return _button
}

// NewButtonFromIconName creates a new button containing an icon from the
// current icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// This function is a convenience wrapper around gtk_button_new() and
// gtk_button_set_image().
func NewButtonFromIconName(iconName string, size int) *Button {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_button_new_from_icon_name(_arg1, _arg2)

	var _button *Button // out

	_button = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Button)

	return _button
}

// NewButtonFromStock creates a new Button containing the image and text from a
// [stock item][gtkstock]. Some stock ids have preprocessor macros like
// K_STOCK_OK and K_STOCK_APPLY.
//
// If @stock_id is unknown, then it will be treated as a mnemonic label (as for
// gtk_button_new_with_mnemonic()).
//
// Deprecated: Stock items are deprecated. Use gtk_button_new_with_label()
// instead.
func NewButtonFromStock(stockId string) *Button {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_button_new_from_stock(_arg1)

	var _button *Button // out

	_button = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Button)

	return _button
}

// NewButtonWithLabel creates a Button widget with a Label child containing the
// given text.
func NewButtonWithLabel(label string) *Button {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_button_new_with_label(_arg1)

	var _button *Button // out

	_button = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Button)

	return _button
}

// NewButtonWithMnemonic creates a new Button containing a label. If characters
// in @label are preceded by an underscore, they are underlined. If you need a
// literal underscore character in a label, use “__” (two underscores). The
// first underlined character represents a keyboard accelerator called a
// mnemonic. Pressing Alt and that key activates the button.
func NewButtonWithMnemonic(label string) *Button {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_button_new_with_mnemonic(_arg1)

	var _button *Button // out

	_button = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Button)

	return _button
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *Button) Native() uintptr {
	return v.Bin.Container.Widget.InitiallyUnowned.Object.Native()
}

// Clicked emits a Button::clicked signal to the given Button.
func (button *Button) Clicked() {
	var _arg0 *C.GtkButton // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	C.gtk_button_clicked(_arg0)
}

// Enter emits a Button::enter signal to the given Button.
//
// Deprecated: Use the Widget::enter-notify-event signal.
func (button *Button) Enter() {
	var _arg0 *C.GtkButton // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	C.gtk_button_enter(_arg0)
}

// Alignment gets the alignment of the child in the button.
//
// Deprecated: Access the child widget directly if you need to control its
// alignment.
func (button *Button) Alignment() (xalign float32, yalign float32) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gfloat     // in
	var _arg2 C.gfloat     // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	C.gtk_button_get_alignment(_arg0, &_arg1, &_arg2)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// AlwaysShowImage returns whether the button will ignore the
// Settings:gtk-button-images setting and always show the image, if available.
func (button *Button) AlwaysShowImage() bool {
	var _arg0 *C.GtkButton // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_always_show_image(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EventWindow returns the button’s event window if it is realized, nil
// otherwise. This function should be rarely needed.
func (button *Button) EventWindow() *gdk.Window {
	var _arg0 *C.GtkButton // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_event_window(_arg0)

	var _window *gdk.Window // out

	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.Window)

	return _window
}

// FocusOnClick returns whether the button grabs focus when it is clicked with
// the mouse. See gtk_button_set_focus_on_click().
//
// Deprecated: Use gtk_widget_get_focus_on_click() instead.
func (button *Button) FocusOnClick() bool {
	var _arg0 *C.GtkButton // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_focus_on_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Image gets the widget that is currenty set as the image of @button. This may
// have been explicitly set by gtk_button_set_image() or constructed by
// gtk_button_new_from_stock().
func (button *Button) Image() *Widget {
	var _arg0 *C.GtkButton // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_image(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// ImagePosition gets the position of the image relative to the text inside the
// button.
func (button *Button) ImagePosition() PositionType {
	var _arg0 *C.GtkButton      // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_image_position(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// Label fetches the text from the label of the button, as set by
// gtk_button_set_label(). If the label text has not been set the return value
// will be nil. This will be the case if you create an empty button with
// gtk_button_new() to use as a container.
func (button *Button) Label() string {
	var _arg0 *C.GtkButton // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Relief returns the current relief style of the given Button.
func (button *Button) Relief() ReliefStyle {
	var _arg0 *C.GtkButton     // out
	var _cret C.GtkReliefStyle // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_relief(_arg0)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// UseStock returns whether the button label is a stock item.
//
// Deprecated: since version 3.10.
func (button *Button) UseStock() bool {
	var _arg0 *C.GtkButton // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_use_stock(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an embedded underline in the button label
// indicates a mnemonic. See gtk_button_set_use_underline ().
func (button *Button) UseUnderline() bool {
	var _arg0 *C.GtkButton // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_button_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Leave emits a Button::leave signal to the given Button.
//
// Deprecated: Use the Widget::leave-notify-event signal.
func (button *Button) Leave() {
	var _arg0 *C.GtkButton // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	C.gtk_button_leave(_arg0)
}

// Pressed emits a Button::pressed signal to the given Button.
//
// Deprecated: Use the Widget::button-press-event signal.
func (button *Button) Pressed() {
	var _arg0 *C.GtkButton // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	C.gtk_button_pressed(_arg0)
}

// Released emits a Button::released signal to the given Button.
//
// Deprecated: Use the Widget::button-release-event signal.
func (button *Button) Released() {
	var _arg0 *C.GtkButton // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))

	C.gtk_button_released(_arg0)
}

// SetAlignment sets the alignment of the child. This property has no effect
// unless the child is a Misc or a Alignment.
//
// Deprecated: Access the child widget directly if you need to control its
// alignment.
func (button *Button) SetAlignment(xalign float32, yalign float32) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gfloat     // out
	var _arg2 C.gfloat     // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)

	C.gtk_button_set_alignment(_arg0, _arg1, _arg2)
}

// SetAlwaysShowImage: if true, the button will ignore the
// Settings:gtk-button-images setting and always show the image, if available.
//
// Use this property if the button would be useless or hard to use without the
// image.
func (button *Button) SetAlwaysShowImage(alwaysShow bool) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	if alwaysShow {
		_arg1 = C.TRUE
	}

	C.gtk_button_set_always_show_image(_arg0, _arg1)
}

// SetFocusOnClick sets whether the button will grab focus when it is clicked
// with the mouse. Making mouse clicks not grab focus is useful in places like
// toolbars where you don’t want the keyboard focus removed from the main area
// of the application.
//
// Deprecated: Use gtk_widget_set_focus_on_click() instead.
func (button *Button) SetFocusOnClick(focusOnClick bool) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_button_set_focus_on_click(_arg0, _arg1)
}

// SetImage: set the image of @button to the given widget. The image will be
// displayed if the label text is nil or if Button:always-show-image is true.
// You don’t have to call gtk_widget_show() on @image yourself.
func (button *Button) SetImage(image Widgeter) {
	var _arg0 *C.GtkButton // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((image).(gextras.Nativer).Native()))

	C.gtk_button_set_image(_arg0, _arg1)
}

// SetImagePosition sets the position of the image relative to the text inside
// the button.
func (button *Button) SetImagePosition(position PositionType) {
	var _arg0 *C.GtkButton      // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_button_set_image_position(_arg0, _arg1)
}

// SetLabel sets the text of the label of the button to @str. This text is also
// used to select the stock item if gtk_button_set_use_stock() is used.
//
// This will also clear any previously set labels.
func (button *Button) SetLabel(label string) {
	var _arg0 *C.GtkButton // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_button_set_label(_arg0, _arg1)
}

// SetRelief sets the relief style of the edges of the given Button widget. Two
// styles exist, GTK_RELIEF_NORMAL and GTK_RELIEF_NONE. The default style is, as
// one can guess, GTK_RELIEF_NORMAL. The deprecated value GTK_RELIEF_HALF
// behaves the same as GTK_RELIEF_NORMAL.
func (button *Button) SetRelief(relief ReliefStyle) {
	var _arg0 *C.GtkButton     // out
	var _arg1 C.GtkReliefStyle // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.GtkReliefStyle(relief)

	C.gtk_button_set_relief(_arg0, _arg1)
}

// SetUseStock: if true, the label set on the button is used as a stock id to
// select the stock item for the button.
//
// Deprecated: since version 3.10.
func (button *Button) SetUseStock(useStock bool) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	if useStock {
		_arg1 = C.TRUE
	}

	C.gtk_button_set_use_stock(_arg0, _arg1)
}

// SetUseUnderline: if true, an underline in the text of the button label
// indicates the next character should be used for the mnemonic accelerator key.
func (button *Button) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkButton // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkButton)(unsafe.Pointer(button.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_button_set_use_underline(_arg0, _arg1)
}
