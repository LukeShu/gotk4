// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_anchor_hints_get_type()), F: marshalAnchorHints},
		{T: externglib.Type(C.gdk_popup_layout_get_type()), F: marshalPopupLayout},
	})
}

// AnchorHints: positioning hints for aligning a surface relative to a
// rectangle.
//
// These hints determine how the surface should be positioned in the case that
// the surface would fall off-screen if placed in its ideal position.
//
// For example, GDK_ANCHOR_FLIP_X will replace GDK_GRAVITY_NORTH_WEST with
// GDK_GRAVITY_NORTH_EAST and vice versa if the surface extends beyond the left
// or right edges of the monitor.
//
// If GDK_ANCHOR_SLIDE_X is set, the surface can be shifted horizontally to fit
// on-screen. If GDK_ANCHOR_RESIZE_X is set, the surface can be shrunken
// horizontally to fit.
//
// In general, when multiple flags are set, flipping should take precedence over
// sliding, which should take precedence over resizing.
type AnchorHints int

const (
	// AnchorFlipX: allow flipping anchors horizontally
	AnchorFlipX AnchorHints = 0b1
	// AnchorFlipY: allow flipping anchors vertically
	AnchorFlipY AnchorHints = 0b10
	// AnchorSlideX: allow sliding surface horizontally
	AnchorSlideX AnchorHints = 0b100
	// AnchorSlideY: allow sliding surface vertically
	AnchorSlideY AnchorHints = 0b1000
	// AnchorResizeX: allow resizing surface horizontally
	AnchorResizeX AnchorHints = 0b10000
	// AnchorResizeY: allow resizing surface vertically
	AnchorResizeY AnchorHints = 0b100000
	// AnchorFlip: allow flipping anchors on both axes
	AnchorFlip AnchorHints = 0b11
	// AnchorSlide: allow sliding surface on both axes
	AnchorSlide AnchorHints = 0b1100
	// AnchorResize: allow resizing surface on both axes
	AnchorResize AnchorHints = 0b110000
)

func marshalAnchorHints(p uintptr) (interface{}, error) {
	return AnchorHints(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for AnchorHints.
func (a AnchorHints) String() string {
	if a == 0 {
		return "AnchorHints(0)"
	}

	var builder strings.Builder
	builder.Grow(113)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AnchorFlipX:
			builder.WriteString("FlipX|")
		case AnchorFlipY:
			builder.WriteString("FlipY|")
		case AnchorSlideX:
			builder.WriteString("SlideX|")
		case AnchorSlideY:
			builder.WriteString("SlideY|")
		case AnchorResizeX:
			builder.WriteString("ResizeX|")
		case AnchorResizeY:
			builder.WriteString("ResizeY|")
		case AnchorFlip:
			builder.WriteString("Flip|")
		case AnchorSlide:
			builder.WriteString("Slide|")
		case AnchorResize:
			builder.WriteString("Resize|")
		default:
			builder.WriteString(fmt.Sprintf("AnchorHints(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AnchorHints) Has(other AnchorHints) bool {
	return (a & other) == other
}

// PopupLayout: GdkPopupLayout struct contains information that is necessary
// position a gdk.Popup relative to its parent.
//
// The positioning requires a negotiation with the windowing system, since it
// depends on external constraints, such as the position of the parent surface,
// and the screen dimensions.
//
// The basic ingredients are a rectangle on the parent surface, and the anchor
// on both that rectangle and the popup. The anchors specify a side or corner to
// place next to each other.
//
// !Popup anchors (popup-anchors.png)
//
// For cases where placing the anchors next to each other would make the popup
// extend offscreen, the layout includes some hints for how to resolve this
// problem. The hints may suggest to flip the anchor position to the other side,
// or to 'slide' the popup along a side, or to resize it.
//
// !Flipping popups (popup-flip.png)
//
// !Sliding popups (popup-slide.png)
//
// These hints may be combined.
//
// Ultimatively, it is up to the windowing system to determine the position and
// size of the popup. You can learn about the result by calling
// gdk.Popup.GetPositionX(), gdk.Popup.GetPositionY(), gdk.Popup.GetRectAnchor()
// and gdk.Popup.GetSurfaceAnchor() after the popup has been presented. This can
// be used to adjust the rendering. For example, gtk.Popover changes its arrow
// position accordingly. But you have to be careful avoid changing the size of
// the popover, or it has to be presented again.
//
// An instance of this type is always passed by reference.
type PopupLayout struct {
	*popupLayout
}

// popupLayout is the struct that's finalized.
type popupLayout struct {
	native *C.GdkPopupLayout
}

func marshalPopupLayout(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &PopupLayout{&popupLayout{(*C.GdkPopupLayout)(unsafe.Pointer(b))}}, nil
}

// NewPopupLayout constructs a struct PopupLayout.
func NewPopupLayout(anchorRect *Rectangle, rectAnchor Gravity, surfaceAnchor Gravity) *PopupLayout {
	var _arg1 *C.GdkRectangle   // out
	var _arg2 C.GdkGravity      // out
	var _arg3 C.GdkGravity      // out
	var _cret *C.GdkPopupLayout // in

	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(anchorRect)))
	_arg2 = C.GdkGravity(rectAnchor)
	_arg3 = C.GdkGravity(surfaceAnchor)

	_cret = C.gdk_popup_layout_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(anchorRect)
	runtime.KeepAlive(rectAnchor)
	runtime.KeepAlive(surfaceAnchor)

	var _popupLayout *PopupLayout // out

	_popupLayout = (*PopupLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_popupLayout)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_popup_layout_unref((*C.GdkPopupLayout)(intern.C))
		},
	)

	return _popupLayout
}

// Copy makes a copy of layout.
func (layout *PopupLayout) Copy() *PopupLayout {
	var _arg0 *C.GdkPopupLayout // out
	var _cret *C.GdkPopupLayout // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_popup_layout_copy(_arg0)
	runtime.KeepAlive(layout)

	var _popupLayout *PopupLayout // out

	_popupLayout = (*PopupLayout)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_popupLayout)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_popup_layout_unref((*C.GdkPopupLayout)(intern.C))
		},
	)

	return _popupLayout
}

// Equal: check whether layout and other has identical layout properties.
func (layout *PopupLayout) Equal(other *PopupLayout) bool {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 *C.GdkPopupLayout // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = C.gdk_popup_layout_equal(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AnchorHints: get the GdkAnchorHints.
func (layout *PopupLayout) AnchorHints() AnchorHints {
	var _arg0 *C.GdkPopupLayout // out
	var _cret C.GdkAnchorHints  // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_popup_layout_get_anchor_hints(_arg0)
	runtime.KeepAlive(layout)

	var _anchorHints AnchorHints // out

	_anchorHints = AnchorHints(_cret)

	return _anchorHints
}

// AnchorRect: get the anchor rectangle.
func (layout *PopupLayout) AnchorRect() *Rectangle {
	var _arg0 *C.GdkPopupLayout // out
	var _cret *C.GdkRectangle   // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_popup_layout_get_anchor_rect(_arg0)
	runtime.KeepAlive(layout)

	var _rectangle *Rectangle // out

	_rectangle = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _rectangle
}

// Offset retrieves the offset for the anchor rectangle.
func (layout *PopupLayout) Offset() (dx int32, dy int32) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.int             // in
	var _arg2 C.int             // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	C.gdk_popup_layout_get_offset(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(layout)

	var _dx int32 // out
	var _dy int32 // out

	_dx = int32(_arg1)
	_dy = int32(_arg2)

	return _dx, _dy
}

// RectAnchor returns the anchor position on the anchor rectangle.
func (layout *PopupLayout) RectAnchor() Gravity {
	var _arg0 *C.GdkPopupLayout // out
	var _cret C.GdkGravity      // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_popup_layout_get_rect_anchor(_arg0)
	runtime.KeepAlive(layout)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// ShadowWidth obtains the shadow widths of this layout.
func (layout *PopupLayout) ShadowWidth() (left int32, right int32, top int32, bottom int32) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.int             // in
	var _arg2 C.int             // in
	var _arg3 C.int             // in
	var _arg4 C.int             // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	C.gdk_popup_layout_get_shadow_width(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(layout)

	var _left int32   // out
	var _right int32  // out
	var _top int32    // out
	var _bottom int32 // out

	_left = int32(_arg1)
	_right = int32(_arg2)
	_top = int32(_arg3)
	_bottom = int32(_arg4)

	return _left, _right, _top, _bottom
}

// SurfaceAnchor returns the anchor position on the popup surface.
func (layout *PopupLayout) SurfaceAnchor() Gravity {
	var _arg0 *C.GdkPopupLayout // out
	var _cret C.GdkGravity      // in

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	_cret = C.gdk_popup_layout_get_surface_anchor(_arg0)
	runtime.KeepAlive(layout)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// SetAnchorHints: set new anchor hints.
//
// The set anchor_hints determines how surface will be moved if the anchor
// points cause it to move off-screen. For example, GDK_ANCHOR_FLIP_X will
// replace GDK_GRAVITY_NORTH_WEST with GDK_GRAVITY_NORTH_EAST and vice versa if
// surface extends beyond the left or right edges of the monitor.
func (layout *PopupLayout) SetAnchorHints(anchorHints AnchorHints) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.GdkAnchorHints  // out

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = C.GdkAnchorHints(anchorHints)

	C.gdk_popup_layout_set_anchor_hints(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(anchorHints)
}

// SetAnchorRect: set the anchor rectangle.
func (layout *PopupLayout) SetAnchorRect(anchorRect *Rectangle) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 *C.GdkRectangle   // out

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(anchorRect)))

	C.gdk_popup_layout_set_anchor_rect(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(anchorRect)
}

// SetOffset: offset the position of the anchor rectangle with the given delta.
func (layout *PopupLayout) SetOffset(dx int32, dy int32) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = C.int(dx)
	_arg2 = C.int(dy)

	C.gdk_popup_layout_set_offset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(dx)
	runtime.KeepAlive(dy)
}

// SetRectAnchor: set the anchor on the anchor rectangle.
func (layout *PopupLayout) SetRectAnchor(anchor Gravity) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.GdkGravity      // out

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = C.GdkGravity(anchor)

	C.gdk_popup_layout_set_rect_anchor(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(anchor)
}

// SetShadowWidth sets the shadow width of the popup.
//
// The shadow width corresponds to the part of the computed surface size that
// would consist of the shadow margin surrounding the window, would there be
// any.
func (layout *PopupLayout) SetShadowWidth(left int32, right int32, top int32, bottom int32) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out
	var _arg3 C.int             // out
	var _arg4 C.int             // out

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = C.int(left)
	_arg2 = C.int(right)
	_arg3 = C.int(top)
	_arg4 = C.int(bottom)

	C.gdk_popup_layout_set_shadow_width(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(left)
	runtime.KeepAlive(right)
	runtime.KeepAlive(top)
	runtime.KeepAlive(bottom)
}

// SetSurfaceAnchor: set the anchor on the popup surface.
func (layout *PopupLayout) SetSurfaceAnchor(anchor Gravity) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.GdkGravity      // out

	_arg0 = (*C.GdkPopupLayout)(gextras.StructNative(unsafe.Pointer(layout)))
	_arg1 = C.GdkGravity(anchor)

	C.gdk_popup_layout_set_surface_anchor(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(anchor)
}
