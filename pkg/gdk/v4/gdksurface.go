// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_surface_get_type()), F: marshalSurface},
	})
}

// Surface: `GdkSurface` is a rectangular region on the screen.
//
// It’s a low-level object, used to implement high-level objects such as
// [class@Gtk.Window] or [class@Gtk.Dialog] in GTK.
//
// The surfaces you see in practice are either [class@Gdk.Toplevel] or
// [class@Gdk.Popup], and those interfaces provide much of the required API to
// interact with these surfaces. Other, more specialized surface types exist,
// but you will rarely interact with them directly.
type Surface interface {
	gextras.Objector

	// Beep emits a short beep associated to @surface.
	//
	// If the display of @surface does not support per-surface beeps, emits a
	// short beep on the display just as [method@Gdk.Display.beep].
	Beep()
	// CreateCairoContext creates a new `GdkCairoContext` for rendering on
	// @surface.
	CreateCairoContext() *CairoContextClass
	// CreateGLContext creates a new `GdkGLContext` for the `GdkSurface`.
	//
	// The context is disconnected from any particular surface or surface. If
	// the creation of the `GdkGLContext` failed, @error will be set. Before
	// using the returned `GdkGLContext`, you will need to call
	// [method@Gdk.GLContext.make_current] or [method@Gdk.GLContext.realize].
	CreateGLContext() (*GLContextClass, error)
	// CreateVulkanContext creates a new `GdkVulkanContext` for rendering on
	// @surface.
	//
	// If the creation of the `GdkVulkanContext` failed, @error will be set.
	CreateVulkanContext() (*VulkanContextClass, error)
	// Destroy destroys the window system resources associated with @surface and
	// decrements @surface's reference count.
	//
	// The window system resources for all children of @surface are also
	// destroyed, but the children’s reference counts are not decremented.
	//
	// Note that a surface will not be destroyed automatically when its
	// reference count reaches zero. You must call this function yourself before
	// that happens.
	Destroy()
	// Cursor retrieves a `GdkCursor` pointer for the cursor currently set on
	// the `GdkSurface`.
	//
	// If the return value is nil then there is no custom cursor set on the
	// surface, and it is using the cursor for its parent surface.
	Cursor() *CursorClass
	// DeviceCursor retrieves a `GdkCursor` pointer for the @device currently
	// set on the specified `GdkSurface`.
	//
	// If the return value is nil then there is no custom cursor set on the
	// specified surface, and it is using the cursor for its parent surface.
	DeviceCursor(device Device) *CursorClass
	// DevicePosition obtains the current device position and modifier state.
	//
	// The position is given in coordinates relative to the upper left corner of
	// @surface.
	DevicePosition(device Device) (x float64, y float64, mask ModifierType, ok bool)
	// Display gets the `GdkDisplay` associated with a `GdkSurface`.
	Display() *DisplayClass
	// FrameClock gets the frame clock for the surface.
	//
	// The frame clock for a surface never changes unless the surface is
	// reparented to a new toplevel surface.
	FrameClock() *FrameClockClass
	// Height returns the height of the given @surface.
	//
	// Surface size is reported in ”application pixels”, not ”device pixels”
	// (see [method@Gdk.Surface.get_scale_factor]).
	Height() int
	// Mapped checks whether the surface has been mapped.
	//
	// A surface is mapped with [method@Gdk.Toplevel.present] or
	// [method@Gdk.Popup.present].
	Mapped() bool
	// ScaleFactor returns the internal scale factor that maps from surface
	// coordinates to the actual device pixels.
	//
	// On traditional systems this is 1, but on very high density outputs this
	// can be a higher value (often 2). A higher value means that drawing is
	// automatically scaled up to a higher resolution, so any code doing drawing
	// will automatically look nicer. However, if you are supplying pixel-based
	// data the scale value can be used to determine whether to use a pixel
	// resource with higher resolution data.
	//
	// The scale of a surface may change during runtime.
	ScaleFactor() int
	// Width returns the width of the given @surface.
	//
	// Surface size is reported in ”application pixels”, not ”device pixels”
	// (see [method@Gdk.Surface.get_scale_factor]).
	Width() int
	// Hide the surface.
	//
	// For toplevel surfaces, withdraws them, so they will no longer be known to
	// the window manager; for all surfaces, unmaps them, so they won’t be
	// displayed. Normally done automatically as part of
	// [method@Gtk.Widget.hide].
	Hide()
	// IsDestroyed: check to see if a surface is destroyed.
	IsDestroyed() bool
	// QueueRender forces a [signal@Gdk.Surface::render] signal emission for
	// @surface to be scheduled.
	//
	// This function is useful for implementations that track invalid regions on
	// their own.
	QueueRender()
	// RequestLayout: request a layout phase from the surface's frame clock.
	//
	// See [method@Gdk.FrameClock.request_phase].
	RequestLayout()
	// SetCursor sets the default mouse pointer for a `GdkSurface`.
	//
	// Passing nil for the @cursor argument means that @surface will use the
	// cursor of its parent surface. Most surfaces should use this default. Note
	// that @cursor must be for the same display as @surface.
	//
	// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
	// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
	SetCursor(cursor Cursor)
	// SetDeviceCursor sets a specific `GdkCursor` for a given device when it
	// gets inside @surface.
	//
	// Passing nil for the @cursor argument means that @surface will use the
	// cursor of its parent surface. Most surfaces should use this default.
	//
	// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
	// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
	SetDeviceCursor(device Device, cursor Cursor)
	// SetInputRegion: apply the region to the surface for the purpose of event
	// handling.
	//
	// Mouse events which happen while the pointer position corresponds to an
	// unset bit in the mask will be passed on the surface below @surface.
	//
	// An input region is typically used with RGBA surfaces. The alpha channel
	// of the surface defines which pixels are invisible and allows for nicely
	// antialiased borders, and the input region controls where the surface is
	// “clickable”.
	//
	// Use [method@Gdk.Display.supports_input_shapes] to find out if a
	// particular backend supports input regions.
	SetInputRegion(region *cairo.Region)
	// SetOpaqueRegion marks a region of the `GdkSurface` as opaque.
	//
	// For optimisation purposes, compositing window managers may like to not
	// draw obscured regions of surfaces, or turn off blending during for these
	// regions. With RGB windows with no transparency, this is just the shape of
	// the window, but with ARGB32 windows, the compositor does not know what
	// regions of the window are transparent or not.
	//
	// This function only works for toplevel surfaces.
	//
	// GTK will update this property automatically if the @surface background is
	// opaque, as we know where the opaque regions are. If your surface
	// background is not opaque, please update this property in your
	// WidgetClass.css_changed() handler.
	SetOpaqueRegion(region *cairo.Region)
}

// SurfaceClass implements the Surface interface.
type SurfaceClass struct {
	*externglib.Object
}

var _ Surface = (*SurfaceClass)(nil)

func wrapSurface(obj *externglib.Object) Surface {
	return &SurfaceClass{
		Object: obj,
	}
}

func marshalSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSurface(obj), nil
}

// NewSurfacePopup: create a new popup surface.
//
// The surface will be attached to @parent and can be positioned relative to it
// using [method@Gdk.Popup.present].
func NewSurfacePopup(parent Surface, autohide bool) *SurfaceClass {
	var _arg1 *C.GdkSurface // out
	var _arg2 C.gboolean    // out
	var _cret *C.GdkSurface // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer((&parent).Native()))
	if autohide {
		_arg2 = C.TRUE
	}

	_cret = C.gdk_surface_new_popup(_arg1, _arg2)

	var _surface *SurfaceClass // out

	_surface = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SurfaceClass)

	return _surface
}

// NewSurfaceToplevel creates a new toplevel surface.
func NewSurfaceToplevel(display Display) *SurfaceClass {
	var _arg1 *C.GdkDisplay // out
	var _cret *C.GdkSurface // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer((&display).Native()))

	_cret = C.gdk_surface_new_toplevel(_arg1)

	var _surface *SurfaceClass // out

	_surface = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SurfaceClass)

	return _surface
}

// Beep emits a short beep associated to @surface.
//
// If the display of @surface does not support per-surface beeps, emits a short
// beep on the display just as [method@Gdk.Display.beep].
func (s *SurfaceClass) Beep() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	C.gdk_surface_beep(_arg0)
}

// CreateCairoContext creates a new `GdkCairoContext` for rendering on @surface.
func (s *SurfaceClass) CreateCairoContext() *CairoContextClass {
	var _arg0 *C.GdkSurface      // out
	var _cret *C.GdkCairoContext // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_create_cairo_context(_arg0)

	var _cairoContext *CairoContextClass // out

	_cairoContext = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*CairoContextClass)

	return _cairoContext
}

// CreateGLContext creates a new `GdkGLContext` for the `GdkSurface`.
//
// The context is disconnected from any particular surface or surface. If the
// creation of the `GdkGLContext` failed, @error will be set. Before using the
// returned `GdkGLContext`, you will need to call
// [method@Gdk.GLContext.make_current] or [method@Gdk.GLContext.realize].
func (s *SurfaceClass) CreateGLContext() (*GLContextClass, error) {
	var _arg0 *C.GdkSurface   // out
	var _cret *C.GdkGLContext // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_create_gl_context(_arg0, &_cerr)

	var _glContext *GLContextClass // out
	var _goerr error               // out

	_glContext = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*GLContextClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _glContext, _goerr
}

// CreateVulkanContext creates a new `GdkVulkanContext` for rendering on
// @surface.
//
// If the creation of the `GdkVulkanContext` failed, @error will be set.
func (s *SurfaceClass) CreateVulkanContext() (*VulkanContextClass, error) {
	var _arg0 *C.GdkSurface       // out
	var _cret *C.GdkVulkanContext // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_create_vulkan_context(_arg0, &_cerr)

	var _vulkanContext *VulkanContextClass // out
	var _goerr error                       // out

	_vulkanContext = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*VulkanContextClass)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _vulkanContext, _goerr
}

// Destroy destroys the window system resources associated with @surface and
// decrements @surface's reference count.
//
// The window system resources for all children of @surface are also destroyed,
// but the children’s reference counts are not decremented.
//
// Note that a surface will not be destroyed automatically when its reference
// count reaches zero. You must call this function yourself before that happens.
func (s *SurfaceClass) Destroy() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	C.gdk_surface_destroy(_arg0)
}

// Cursor retrieves a `GdkCursor` pointer for the cursor currently set on the
// `GdkSurface`.
//
// If the return value is nil then there is no custom cursor set on the surface,
// and it is using the cursor for its parent surface.
func (s *SurfaceClass) Cursor() *CursorClass {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkCursor  // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_get_cursor(_arg0)

	var _cursor *CursorClass // out

	_cursor = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*CursorClass)

	return _cursor
}

// DeviceCursor retrieves a `GdkCursor` pointer for the @device currently set on
// the specified `GdkSurface`.
//
// If the return value is nil then there is no custom cursor set on the
// specified surface, and it is using the cursor for its parent surface.
func (s *SurfaceClass) DeviceCursor(device Device) *CursorClass {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkDevice  // out
	var _cret *C.GdkCursor  // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((&device).Native()))

	_cret = C.gdk_surface_get_device_cursor(_arg0, _arg1)

	var _cursor *CursorClass // out

	_cursor = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*CursorClass)

	return _cursor
}

// DevicePosition obtains the current device position and modifier state.
//
// The position is given in coordinates relative to the upper left corner of
// @surface.
func (s *SurfaceClass) DevicePosition(device Device) (x float64, y float64, mask ModifierType, ok bool) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.GdkDevice      // out
	var _arg2 C.double          // in
	var _arg3 C.double          // in
	var _arg4 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((&device).Native()))

	_cret = C.gdk_surface_get_device_position(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _x float64         // out
	var _y float64         // out
	var _mask ModifierType // out
	var _ok bool           // out

	_x = float64(_arg2)
	_y = float64(_arg3)
	_mask = (ModifierType)(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _mask, _ok
}

// Display gets the `GdkDisplay` associated with a `GdkSurface`.
func (s *SurfaceClass) Display() *DisplayClass {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_get_display(_arg0)

	var _display *DisplayClass // out

	_display = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*DisplayClass)

	return _display
}

// FrameClock gets the frame clock for the surface.
//
// The frame clock for a surface never changes unless the surface is reparented
// to a new toplevel surface.
func (s *SurfaceClass) FrameClock() *FrameClockClass {
	var _arg0 *C.GdkSurface    // out
	var _cret *C.GdkFrameClock // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_get_frame_clock(_arg0)

	var _frameClock *FrameClockClass // out

	_frameClock = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*FrameClockClass)

	return _frameClock
}

// Height returns the height of the given @surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels” (see
// [method@Gdk.Surface.get_scale_factor]).
func (s *SurfaceClass) Height() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_get_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Mapped checks whether the surface has been mapped.
//
// A surface is mapped with [method@Gdk.Toplevel.present] or
// [method@Gdk.Popup.present].
func (s *SurfaceClass) Mapped() bool {
	var _arg0 *C.GdkSurface // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_get_mapped(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ScaleFactor returns the internal scale factor that maps from surface
// coordinates to the actual device pixels.
//
// On traditional systems this is 1, but on very high density outputs this can
// be a higher value (often 2). A higher value means that drawing is
// automatically scaled up to a higher resolution, so any code doing drawing
// will automatically look nicer. However, if you are supplying pixel-based data
// the scale value can be used to determine whether to use a pixel resource with
// higher resolution data.
//
// The scale of a surface may change during runtime.
func (s *SurfaceClass) ScaleFactor() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_get_scale_factor(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Width returns the width of the given @surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels” (see
// [method@Gdk.Surface.get_scale_factor]).
func (s *SurfaceClass) Width() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Hide the surface.
//
// For toplevel surfaces, withdraws them, so they will no longer be known to the
// window manager; for all surfaces, unmaps them, so they won’t be displayed.
// Normally done automatically as part of [method@Gtk.Widget.hide].
func (s *SurfaceClass) Hide() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	C.gdk_surface_hide(_arg0)
}

// IsDestroyed: check to see if a surface is destroyed.
func (s *SurfaceClass) IsDestroyed() bool {
	var _arg0 *C.GdkSurface // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	_cret = C.gdk_surface_is_destroyed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// QueueRender forces a [signal@Gdk.Surface::render] signal emission for
// @surface to be scheduled.
//
// This function is useful for implementations that track invalid regions on
// their own.
func (s *SurfaceClass) QueueRender() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	C.gdk_surface_queue_render(_arg0)
}

// RequestLayout: request a layout phase from the surface's frame clock.
//
// See [method@Gdk.FrameClock.request_phase].
func (s *SurfaceClass) RequestLayout() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))

	C.gdk_surface_request_layout(_arg0)
}

// SetCursor sets the default mouse pointer for a `GdkSurface`.
//
// Passing nil for the @cursor argument means that @surface will use the cursor
// of its parent surface. Most surfaces should use this default. Note that
// @cursor must be for the same display as @surface.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture] to
// create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
func (s *SurfaceClass) SetCursor(cursor Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GdkCursor)(unsafe.Pointer((&cursor).Native()))

	C.gdk_surface_set_cursor(_arg0, _arg1)
}

// SetDeviceCursor sets a specific `GdkCursor` for a given device when it gets
// inside @surface.
//
// Passing nil for the @cursor argument means that @surface will use the cursor
// of its parent surface. Most surfaces should use this default.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture] to
// create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
func (s *SurfaceClass) SetDeviceCursor(device Device, cursor Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkDevice  // out
	var _arg2 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((&device).Native()))
	_arg2 = (*C.GdkCursor)(unsafe.Pointer((&cursor).Native()))

	C.gdk_surface_set_device_cursor(_arg0, _arg1, _arg2)
}

// SetInputRegion: apply the region to the surface for the purpose of event
// handling.
//
// Mouse events which happen while the pointer position corresponds to an unset
// bit in the mask will be passed on the surface below @surface.
//
// An input region is typically used with RGBA surfaces. The alpha channel of
// the surface defines which pixels are invisible and allows for nicely
// antialiased borders, and the input region controls where the surface is
// “clickable”.
//
// Use [method@Gdk.Display.supports_input_shapes] to find out if a particular
// backend supports input regions.
func (s *SurfaceClass) SetInputRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region))

	C.gdk_surface_set_input_region(_arg0, _arg1)
}

// SetOpaqueRegion marks a region of the `GdkSurface` as opaque.
//
// For optimisation purposes, compositing window managers may like to not draw
// obscured regions of surfaces, or turn off blending during for these regions.
// With RGB windows with no transparency, this is just the shape of the window,
// but with ARGB32 windows, the compositor does not know what regions of the
// window are transparent or not.
//
// This function only works for toplevel surfaces.
//
// GTK will update this property automatically if the @surface background is
// opaque, as we know where the opaque regions are. If your surface background
// is not opaque, please update this property in your WidgetClass.css_changed()
// handler.
func (s *SurfaceClass) SetOpaqueRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer((&s).Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region))

	C.gdk_surface_set_opaque_region(_arg0, _arg1)
}
