// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeGLContext = coreglib.Type(C.gdk_gl_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLContext, F: marshalGLContext},
	})
}

// GLContext is an object representing the platform-specific OpenGL drawing
// context.
//
// GLContexts are created for a Window using gdk_window_create_gl_context(), and
// the context will match the Visual of the window.
//
// A GLContext is not tied to any particular normal framebuffer. For instance,
// it cannot draw to the Window back buffer. The GDK repaint system is in full
// control of the painting to that. Instead, you can create render buffers or
// textures and use gdk_cairo_draw_from_gl() in the draw function of your widget
// to draw them. Then GDK will handle the integration of your rendering with
// that of other widgets.
//
// Support for GLContext is platform-specific, context creation can fail,
// returning NULL context.
//
// A GLContext has to be made "current" in order to start using it, otherwise
// any OpenGL call will be ignored.
//
// Creating a new OpenGL context ##
//
// In order to create a new GLContext instance you need a Window, which you
// typically get during the realize call of a widget.
//
// A GLContext is not realized until either gdk_gl_context_make_current(), or
// until it is realized using gdk_gl_context_realize(). It is possible to
// specify details of the GL context like the OpenGL version to be used, or
// whether the GL context should have extra state validation enabled after
// calling gdk_window_create_gl_context() by calling gdk_gl_context_realize().
// If the realization fails you have the option to change the settings of the
// GLContext and try again.
//
// Using a GdkGLContext ##
//
// You will need to make the GLContext the current context before issuing OpenGL
// calls; the system sends OpenGL commands to whichever context is current. It
// is possible to have multiple contexts, so you always need to ensure that the
// one which you want to draw with is the current one before issuing commands:
//
//    gdk_gl_context_make_current (context);
//
// You can now perform your drawing using OpenGL commands.
//
// You can check which GLContext is the current one by using
// gdk_gl_context_get_current(); you can also unset any GLContext that is
// currently set by calling gdk_gl_context_clear_current().
type GLContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*GLContext)(nil)
)

// GLContexter describes types inherited from class GLContext.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type GLContexter interface {
	coreglib.Objector
	baseGLContext() *GLContext
}

var _ GLContexter = (*GLContext)(nil)

func wrapGLContext(obj *coreglib.Object) *GLContext {
	return &GLContext{
		Object: obj,
	}
}

func marshalGLContext(p uintptr) (interface{}, error) {
	return wrapGLContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (context *GLContext) baseGLContext() *GLContext {
	return context
}

// BaseGLContext returns the underlying base object.
func BaseGLContext(obj GLContexter) *GLContext {
	return obj.baseGLContext()
}

// DebugEnabled retrieves the value set using
// gdk_gl_context_set_debug_enabled().
//
// The function returns the following values:
//
//    - ok: TRUE if debugging is enabled.
//
func (context *GLContext) DebugEnabled() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_gl_context_get_debug_enabled(_arg0)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Display retrieves the Display the context is created for.
//
// The function returns the following values:
//
//    - display (optional) or NULL.
//
func (context *GLContext) Display() *Display {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_gl_context_get_display(_arg0)
	runtime.KeepAlive(context)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// ForwardCompatible retrieves the value set using
// gdk_gl_context_set_forward_compatible().
//
// The function returns the following values:
//
//    - ok: TRUE if the context should be forward compatible.
//
func (context *GLContext) ForwardCompatible() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_gl_context_get_forward_compatible(_arg0)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RequiredVersion retrieves the major and minor version requested by calling
// gdk_gl_context_set_required_version().
//
// The function returns the following values:
//
//    - major (optional): return location for the major version to request.
//    - minor (optional): return location for the minor version to request.
//
func (context *GLContext) RequiredVersion() (major, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // in
	var _arg2 C.int           // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gdk_gl_context_get_required_version(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

// SharedContext retrieves the GLContext that this context share data with.
//
// The function returns the following values:
//
//    - glContext (optional) or NULL.
//
func (context *GLContext) SharedContext() GLContexter {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkGLContext // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_gl_context_get_shared_context(_arg0)
	runtime.KeepAlive(context)

	var _glContext GLContexter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(GLContexter)
				return ok
			})
			rv, ok := casted.(GLContexter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
			}
			_glContext = rv
		}
	}

	return _glContext
}

// UseES checks whether the context is using an OpenGL or OpenGL ES profile.
//
// The function returns the following values:
//
//    - ok: TRUE if the GLContext is using an OpenGL ES profile.
//
func (context *GLContext) UseES() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_gl_context_get_use_es(_arg0)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Version retrieves the OpenGL version of the context.
//
// The context must be realized prior to calling this function.
//
// The function returns the following values:
//
//    - major: return location for the major version.
//    - minor: return location for the minor version.
//
func (context *GLContext) Version() (major, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // in
	var _arg2 C.int           // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gdk_gl_context_get_version(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

// Window retrieves the Window used by the context.
//
// The function returns the following values:
//
//    - window (optional) or NULL.
//
func (context *GLContext) Window() Windower {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkWindow    // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_gl_context_get_window(_arg0)
	runtime.KeepAlive(context)

	var _window Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Windower)
				return ok
			})
			rv, ok := casted.(Windower)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// IsLegacy: whether the GLContext is in legacy mode or not.
//
// The GLContext must be realized before calling this function.
//
// When realizing a GL context, GDK will try to use the OpenGL 3.2 core profile;
// this profile removes all the OpenGL API that was deprecated prior to the 3.2
// version of the specification. If the realization is successful, this function
// will return FALSE.
//
// If the underlying OpenGL implementation does not support core profiles, GDK
// will fall back to a pre-3.2 compatibility profile, and this function will
// return TRUE.
//
// You can use the value returned by this function to decide which kind of
// OpenGL API to use, or whether to do extension discovery, or what kind of
// shader programs to load.
//
// The function returns the following values:
//
//    - ok: TRUE if the GL context is in legacy mode.
//
func (context *GLContext) IsLegacy() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gdk_gl_context_is_legacy(_arg0)
	runtime.KeepAlive(context)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MakeCurrent makes the context the current one.
func (context *GLContext) MakeCurrent() {
	var _arg0 *C.GdkGLContext // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gdk_gl_context_make_current(_arg0)
	runtime.KeepAlive(context)
}

// Realize realizes the given GLContext.
//
// It is safe to call this function on a realized GLContext.
func (context *GLContext) Realize() error {
	var _arg0 *C.GdkGLContext // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C.gdk_gl_context_realize(_arg0, &_cerr)
	runtime.KeepAlive(context)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetDebugEnabled sets whether the GLContext should perform extra validations
// and run time checking. This is useful during development, but has additional
// overhead.
//
// The GLContext must not be realized or made current prior to calling this
// function.
//
// The function takes the following parameters:
//
//    - enabled: whether to enable debugging in the context.
//
func (context *GLContext) SetDebugEnabled(enabled bool) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.gdk_gl_context_set_debug_enabled(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(enabled)
}

// SetForwardCompatible sets whether the GLContext should be forward compatible.
//
// Forward compatibile contexts must not support OpenGL functionality that has
// been marked as deprecated in the requested version; non-forward compatible
// contexts, on the other hand, must support both deprecated and non deprecated
// functionality.
//
// The GLContext must not be realized or made current prior to calling this
// function.
//
// The function takes the following parameters:
//
//    - compatible: whether the context should be forward compatible.
//
func (context *GLContext) SetForwardCompatible(compatible bool) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	if compatible {
		_arg1 = C.TRUE
	}

	C.gdk_gl_context_set_forward_compatible(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(compatible)
}

// SetRequiredVersion sets the major and minor version of OpenGL to request.
//
// Setting major and minor to zero will use the default values.
//
// The GLContext must not be realized or made current prior to calling this
// function.
//
// The function takes the following parameters:
//
//    - major version to request.
//    - minor version to request.
//
func (context *GLContext) SetRequiredVersion(major, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.int(major)
	_arg2 = C.int(minor)

	C.gdk_gl_context_set_required_version(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(major)
	runtime.KeepAlive(minor)
}

// SetUseES requests that GDK create a OpenGL ES context instead of an OpenGL
// one, if the platform and windowing system allows it.
//
// The context must not have been realized.
//
// By default, GDK will attempt to automatically detect whether the underlying
// GL implementation is OpenGL or OpenGL ES once the context is realized.
//
// You should check the return value of gdk_gl_context_get_use_es() after
// calling gdk_gl_context_realize() to decide whether to use the OpenGL or
// OpenGL ES API, extensions, or shaders.
//
// The function takes the following parameters:
//
//    - useEs: whether the context should use OpenGL ES instead of OpenGL, or -1
//      to allow auto-detection.
//
func (context *GLContext) SetUseES(useEs int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.int(useEs)

	C.gdk_gl_context_set_use_es(_arg0, _arg1)
	runtime.KeepAlive(context)
	runtime.KeepAlive(useEs)
}
