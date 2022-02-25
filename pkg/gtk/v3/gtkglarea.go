// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_GLAreaClass_render(GtkGLArea*, GdkGLContext*);
// extern gboolean _gotk4_gtk3_GLArea_ConnectRender(gpointer, GdkGLContext*, guintptr);
// extern void _gotk4_gtk3_GLAreaClass_resize(GtkGLArea*, int, int);
// extern void _gotk4_gtk3_GLArea_ConnectResize(gpointer, gint, gint, guintptr);
import "C"

// glib.Type values for gtkglarea.go.
var GTypeGLArea = externglib.Type(C.gtk_gl_area_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGLArea, F: marshalGLArea},
	})
}

// GLAreaOverrider contains methods that are overridable.
type GLAreaOverrider interface {
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Render(context gdk.GLContexter) bool
	// The function takes the following parameters:
	//
	//    - width
	//    - height
	//
	Resize(width, height int)
}

// GLArea is a widget that allows drawing with OpenGL.
//
// GLArea sets up its own GLContext for the window it creates, and creates a
// custom GL framebuffer that the widget will do GL rendering onto. It also
// ensures that this framebuffer is the default GL rendering target when
// rendering.
//
// In order to draw, you have to connect to the GLArea::render signal, or
// subclass GLArea and override the GtkGLAreaClass.render() virtual function.
//
// The GLArea widget ensures that the GLContext is associated with the widget's
// drawing area, and it is kept updated when the size and position of the
// drawing area changes.
//
// Drawing with GtkGLArea ##
//
// The simplest way to draw using OpenGL commands in a GLArea is to create a
// widget instance and connect to the GLArea::render signal:
//
//      static void
//      on_realize (GtkGLarea *area)
//      {
//        // We need to make the context current if we want to
//        // call GL API
//        gtk_gl_area_make_current (area);
//
//        // If there were errors during the initialization or
//        // when trying to make the context current, this
//        // function will return a #GError for you to catch
//        if (gtk_gl_area_get_error (area) != NULL)
//          return;
//
//        // You can also use gtk_gl_area_set_error() in order
//        // to show eventual initialization errors on the
//        // GtkGLArea widget itself
//        GError *internal_error = NULL;
//        init_buffer_objects (&error);
//        if (error != NULL)
//          {
//            gtk_gl_area_set_error (area, error);
//            g_error_free (error);
//            return;
//          }
//
//        init_shaders (&error);
//        if (error != NULL)
//          {
//            gtk_gl_area_set_error (area, error);
//            g_error_free (error);
//            return;
//          }
//      }
//
// If you need to change the options for creating the GLContext you should use
// the GLArea::create-context signal.
type GLArea struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*GLArea)(nil)
)

func classInitGLAreaer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkGLAreaClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkGLAreaClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Render(context gdk.GLContexter) bool
	}); ok {
		pclass.render = (*[0]byte)(C._gotk4_gtk3_GLAreaClass_render)
	}

	if _, ok := goval.(interface{ Resize(width, height int) }); ok {
		pclass.resize = (*[0]byte)(C._gotk4_gtk3_GLAreaClass_resize)
	}
}

//export _gotk4_gtk3_GLAreaClass_render
func _gotk4_gtk3_GLAreaClass_render(arg0 *C.GtkGLArea, arg1 *C.GdkGLContext) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Render(context gdk.GLContexter) bool
	})

	var _context gdk.GLContexter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gdk.GLContexter)
			return ok
		})
		rv, ok := casted.(gdk.GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
		}
		_context = rv
	}

	ok := iface.Render(_context)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_GLAreaClass_resize
func _gotk4_gtk3_GLAreaClass_resize(arg0 *C.GtkGLArea, arg1 C.int, arg2 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Resize(width, height int) })

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	iface.Resize(_width, _height)
}

func wrapGLArea(obj *externglib.Object) *GLArea {
	return &GLArea{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
	}
}

func marshalGLArea(p uintptr) (interface{}, error) {
	return wrapGLArea(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_GLArea_ConnectRender
func _gotk4_gtk3_GLArea_ConnectRender(arg0 C.gpointer, arg1 *C.GdkGLContext, arg2 C.guintptr) (cret C.gboolean) {
	var f func(context gdk.GLContexter) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context gdk.GLContexter) (ok bool))
	}

	var _context gdk.GLContexter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gdk.GLContexter)
			return ok
		})
		rv, ok := casted.(gdk.GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
		}
		_context = rv
	}

	ok := f(_context)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectRender signal is emitted every time the contents of the GLArea should
// be redrawn.
//
// The context is bound to the area prior to emitting this function, and the
// buffers are painted to the window once the emission terminates.
func (area *GLArea) ConnectRender(f func(context gdk.GLContexter) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(area, "render", false, unsafe.Pointer(C._gotk4_gtk3_GLArea_ConnectRender), f)
}

//export _gotk4_gtk3_GLArea_ConnectResize
func _gotk4_gtk3_GLArea_ConnectResize(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(width, height int)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(width, height int))
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	f(_width, _height)
}

// ConnectResize signal is emitted once when the widget is realized, and then
// each time the widget is changed while realized. This is useful in order to
// keep GL state up to date with the widget size, like for instance camera
// properties which may depend on the width/height ratio.
//
// The GL context for the area is guaranteed to be current when this signal is
// emitted.
//
// The default handler sets up the GL viewport.
func (area *GLArea) ConnectResize(f func(width, height int)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(area, "resize", false, unsafe.Pointer(C._gotk4_gtk3_GLArea_ConnectResize), f)
}

// NewGLArea creates a new GLArea widget.
//
// The function returns the following values:
//
//    - glArea: new GLArea.
//
func NewGLArea() *GLArea {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_gl_area_new()

	var _glArea *GLArea // out

	_glArea = wrapGLArea(externglib.Take(unsafe.Pointer(_cret)))

	return _glArea
}

// AttachBuffers ensures that the area framebuffer object is made the current
// draw and read target, and that all the required buffers for the area are
// created and bound to the frambuffer.
//
// This function is automatically called before emitting the GLArea::render
// signal, and doesn't normally need to be called by application code.
func (area *GLArea) AttachBuffers() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	C.gtk_gl_area_attach_buffers(_arg0)
	runtime.KeepAlive(area)
}

// AutoRender returns whether the area is in auto render mode or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the area is auto rendering, FALSE otherwise.
//
func (area *GLArea) AutoRender() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_auto_render(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Context retrieves the GLContext used by area.
//
// The function returns the following values:
//
//    - glContext: GLContext.
//
func (area *GLArea) Context() gdk.GLContexter {
	var _arg0 *C.GtkGLArea    // out
	var _cret *C.GdkGLContext // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_context(_arg0)
	runtime.KeepAlive(area)

	var _glContext gdk.GLContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gdk.GLContexter)
			return ok
		})
		rv, ok := casted.(gdk.GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
		}
		_glContext = rv
	}

	return _glContext
}

// Error gets the current error set on the area.
//
// The function returns the following values:
//
//    - err (optional) or NULL.
//
func (area *GLArea) Error() error {
	var _arg0 *C.GtkGLArea // out
	var _cret *C.GError    // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_error(_arg0)
	runtime.KeepAlive(area)

	var _err error // out

	if _cret != nil {
		_err = gerror.Take(unsafe.Pointer(_cret))
	}

	return _err
}

// HasAlpha returns whether the area has an alpha component.
//
// The function returns the following values:
//
//    - ok: TRUE if the area has an alpha component, FALSE otherwise.
//
func (area *GLArea) HasAlpha() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_has_alpha(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasDepthBuffer returns whether the area has a depth buffer.
//
// The function returns the following values:
//
//    - ok: TRUE if the area has a depth buffer, FALSE otherwise.
//
func (area *GLArea) HasDepthBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_has_depth_buffer(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasStencilBuffer returns whether the area has a stencil buffer.
//
// The function returns the following values:
//
//    - ok: TRUE if the area has a stencil buffer, FALSE otherwise.
//
func (area *GLArea) HasStencilBuffer() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_has_stencil_buffer(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RequiredVersion retrieves the required version of OpenGL set using
// gtk_gl_area_set_required_version().
//
// The function returns the following values:
//
//    - major: return location for the required major version.
//    - minor: return location for the required minor version.
//
func (area *GLArea) RequiredVersion() (major int, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gint       // in
	var _arg2 C.gint       // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	C.gtk_gl_area_get_required_version(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(area)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

// UseES retrieves the value set by gtk_gl_area_set_use_es().
//
// The function returns the following values:
//
//    - ok: TRUE if the GLArea should create an OpenGL ES context and FALSE
//      otherwise.
//
func (area *GLArea) UseES() bool {
	var _arg0 *C.GtkGLArea // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	_cret = C.gtk_gl_area_get_use_es(_arg0)
	runtime.KeepAlive(area)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MakeCurrent ensures that the GLContext used by area is associated with the
// GLArea.
//
// This function is automatically called before emitting the GLArea::render
// signal, and doesn't normally need to be called by application code.
func (area *GLArea) MakeCurrent() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	C.gtk_gl_area_make_current(_arg0)
	runtime.KeepAlive(area)
}

// QueueRender marks the currently rendered data (if any) as invalid, and queues
// a redraw of the widget, ensuring that the GLArea::render signal is emitted
// during the draw.
//
// This is only needed when the gtk_gl_area_set_auto_render() has been called
// with a FALSE value. The default behaviour is to emit GLArea::render on each
// draw.
func (area *GLArea) QueueRender() {
	var _arg0 *C.GtkGLArea // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))

	C.gtk_gl_area_queue_render(_arg0)
	runtime.KeepAlive(area)
}

// SetAutoRender: if auto_render is TRUE the GLArea::render signal will be
// emitted every time the widget draws. This is the default and is useful if
// drawing the widget is faster.
//
// If auto_render is FALSE the data from previous rendering is kept around and
// will be used for drawing the widget the next time, unless the window is
// resized. In order to force a rendering gtk_gl_area_queue_render() must be
// called. This mode is useful when the scene changes seldomly, but takes a long
// time to redraw.
//
// The function takes the following parameters:
//
//    - autoRender: boolean.
//
func (area *GLArea) SetAutoRender(autoRender bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))
	if autoRender {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_auto_render(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(autoRender)
}

// SetError sets an error on the area which will be shown instead of the GL
// rendering. This is useful in the GLArea::create-context signal if GL context
// creation fails.
//
// The function takes the following parameters:
//
//    - err (optional): new #GError, or NULL to unset the error.
//
func (area *GLArea) SetError(err error) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 *C.GError    // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))
	if err != nil {
		_arg1 = (*C.GError)(gerror.New(err))
	}

	C.gtk_gl_area_set_error(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(err)
}

// SetHasAlpha: if has_alpha is TRUE the buffer allocated by the widget will
// have an alpha channel component, and when rendering to the window the result
// will be composited over whatever is below the widget.
//
// If has_alpha is FALSE there will be no alpha channel, and the buffer will
// fully replace anything below the widget.
//
// The function takes the following parameters:
//
//    - hasAlpha: TRUE to add an alpha component.
//
func (area *GLArea) SetHasAlpha(hasAlpha bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))
	if hasAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_alpha(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasAlpha)
}

// SetHasDepthBuffer: if has_depth_buffer is TRUE the widget will allocate and
// enable a depth buffer for the target framebuffer. Otherwise there will be
// none.
//
// The function takes the following parameters:
//
//    - hasDepthBuffer: TRUE to add a depth buffer.
//
func (area *GLArea) SetHasDepthBuffer(hasDepthBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))
	if hasDepthBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_depth_buffer(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasDepthBuffer)
}

// SetHasStencilBuffer: if has_stencil_buffer is TRUE the widget will allocate
// and enable a stencil buffer for the target framebuffer. Otherwise there will
// be none.
//
// The function takes the following parameters:
//
//    - hasStencilBuffer: TRUE to add a stencil buffer.
//
func (area *GLArea) SetHasStencilBuffer(hasStencilBuffer bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))
	if hasStencilBuffer {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_has_stencil_buffer(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(hasStencilBuffer)
}

// SetRequiredVersion sets the required version of OpenGL to be used when
// creating the context for the widget.
//
// This function must be called before the area has been realized.
//
// The function takes the following parameters:
//
//    - major version.
//    - minor version.
//
func (area *GLArea) SetRequiredVersion(major, minor int) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))
	_arg1 = C.gint(major)
	_arg2 = C.gint(minor)

	C.gtk_gl_area_set_required_version(_arg0, _arg1, _arg2)
	runtime.KeepAlive(area)
	runtime.KeepAlive(major)
	runtime.KeepAlive(minor)
}

// SetUseES sets whether the area should create an OpenGL or an OpenGL ES
// context.
//
// You should check the capabilities of the GLContext before drawing with either
// API.
//
// The function takes the following parameters:
//
//    - useEs: whether to use OpenGL or OpenGL ES.
//
func (area *GLArea) SetUseES(useEs bool) {
	var _arg0 *C.GtkGLArea // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkGLArea)(unsafe.Pointer(externglib.InternObject(area).Native()))
	if useEs {
		_arg1 = C.TRUE
	}

	C.gtk_gl_area_set_use_es(_arg0, _arg1)
	runtime.KeepAlive(area)
	runtime.KeepAlive(useEs)
}
