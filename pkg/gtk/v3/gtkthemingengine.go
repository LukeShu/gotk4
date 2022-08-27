// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_ThemingEngineClass_render_slider(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble, GtkOrientation);
// extern void _gotk4_gtk3_ThemingEngineClass_render_option(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_line(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_layout(GtkThemingEngine*, cairo_t*, gdouble, gdouble, PangoLayout*);
// extern void _gotk4_gtk3_ThemingEngineClass_render_icon_surface(GtkThemingEngine*, cairo_t*, cairo_surface_t*, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_icon(GtkThemingEngine*, cairo_t*, GdkPixbuf*, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_handle(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_frame_gap(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble, GtkPositionType, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_frame(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_focus(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_extension(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble, GtkPositionType);
// extern void _gotk4_gtk3_ThemingEngineClass_render_expander(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_check(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_background(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_arrow(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_activity(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble);
// void _gotk4_gtk3_ThemingEngine_virtual_render_activity(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_arrow(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_background(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_check(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_expander(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_extension(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5, GtkPositionType arg6) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble, GtkPositionType))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_focus(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_frame(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_frame_gap(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5, GtkPositionType arg6, gdouble arg7, gdouble arg8) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble, GtkPositionType, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_handle(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_icon(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, GdkPixbuf* arg2, gdouble arg3, gdouble arg4) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, GdkPixbuf*, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_icon_surface(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, cairo_surface_t* arg2, gdouble arg3, gdouble arg4) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, cairo_surface_t*, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_layout(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, PangoLayout* arg4) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, PangoLayout*))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_line(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_option(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// void _gotk4_gtk3_ThemingEngine_virtual_render_slider(void* fnptr, GtkThemingEngine* arg0, cairo_t* arg1, gdouble arg2, gdouble arg3, gdouble arg4, gdouble arg5, GtkOrientation arg6) {
//   ((void (*)(GtkThemingEngine*, cairo_t*, gdouble, gdouble, gdouble, gdouble, GtkOrientation))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
// };
import "C"

// GType values.
var (
	GTypeThemingEngine = coreglib.Type(C.gtk_theming_engine_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeThemingEngine, F: marshalThemingEngine},
	})
}

// ThemingEngineOverrides contains methods that are overridable.
type ThemingEngineOverrides struct {
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderActivity func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - angle
	//    - x
	//    - y
	//    - size
	//
	RenderArrow func(cr *cairo.Context, angle, x, y, size float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderBackground func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderCheck func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderExpander func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//    - gapSide
	//
	RenderExtension func(cr *cairo.Context, x, y, width, height float64, gapSide PositionType)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderFocus func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderFrame func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//    - gapSide
	//    - xy0Gap
	//    - xy1Gap
	//
	RenderFrameGap func(cr *cairo.Context, x, y, width, height float64, gapSide PositionType, xy0Gap, xy1Gap float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderHandle func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - pixbuf
	//    - x
	//    - y
	//
	RenderIcon func(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x, y float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - surface
	//    - x
	//    - y
	//
	RenderIconSurface func(cr *cairo.Context, surface *cairo.Surface, x, y float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - layout
	//
	RenderLayout func(cr *cairo.Context, x, y float64, layout *pango.Layout)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x0
	//    - y0
	//    - x1
	//    - y1
	//
	RenderLine func(cr *cairo.Context, x0, y0, x1, y1 float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderOption func(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//    - orientation
	//
	RenderSlider func(cr *cairo.Context, x, y, width, height float64, orientation Orientation)
}

func defaultThemingEngineOverrides(v *ThemingEngine) ThemingEngineOverrides {
	return ThemingEngineOverrides{
		RenderActivity:    v.renderActivity,
		RenderArrow:       v.renderArrow,
		RenderBackground:  v.renderBackground,
		RenderCheck:       v.renderCheck,
		RenderExpander:    v.renderExpander,
		RenderExtension:   v.renderExtension,
		RenderFocus:       v.renderFocus,
		RenderFrame:       v.renderFrame,
		RenderFrameGap:    v.renderFrameGap,
		RenderHandle:      v.renderHandle,
		RenderIcon:        v.renderIcon,
		RenderIconSurface: v.renderIconSurface,
		RenderLayout:      v.renderLayout,
		RenderLine:        v.renderLine,
		RenderOption:      v.renderOption,
		RenderSlider:      v.renderSlider,
	}
}

// ThemingEngine was the object used for rendering themed content in GTK+
// widgets. It used to allow overriding GTK+'s default implementation of
// rendering functions by allowing engines to be loaded as modules.
//
// ThemingEngine has been deprecated in GTK+ 3.14 and will be ignored for
// rendering. The advancements in CSS theming are good enough to allow themers
// to achieve their goals without the need to modify source code.
type ThemingEngine struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ThemingEngine)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ThemingEngine, *ThemingEngineClass, ThemingEngineOverrides](
		GTypeThemingEngine,
		initThemingEngineClass,
		wrapThemingEngine,
		defaultThemingEngineOverrides,
	)
}

func initThemingEngineClass(gclass unsafe.Pointer, overrides ThemingEngineOverrides, classInitFunc func(*ThemingEngineClass)) {
	pclass := (*C.GtkThemingEngineClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeThemingEngine))))

	if overrides.RenderActivity != nil {
		pclass.render_activity = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_activity)
	}

	if overrides.RenderArrow != nil {
		pclass.render_arrow = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_arrow)
	}

	if overrides.RenderBackground != nil {
		pclass.render_background = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_background)
	}

	if overrides.RenderCheck != nil {
		pclass.render_check = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_check)
	}

	if overrides.RenderExpander != nil {
		pclass.render_expander = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_expander)
	}

	if overrides.RenderExtension != nil {
		pclass.render_extension = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_extension)
	}

	if overrides.RenderFocus != nil {
		pclass.render_focus = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_focus)
	}

	if overrides.RenderFrame != nil {
		pclass.render_frame = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_frame)
	}

	if overrides.RenderFrameGap != nil {
		pclass.render_frame_gap = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_frame_gap)
	}

	if overrides.RenderHandle != nil {
		pclass.render_handle = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_handle)
	}

	if overrides.RenderIcon != nil {
		pclass.render_icon = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_icon)
	}

	if overrides.RenderIconSurface != nil {
		pclass.render_icon_surface = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_icon_surface)
	}

	if overrides.RenderLayout != nil {
		pclass.render_layout = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_layout)
	}

	if overrides.RenderLine != nil {
		pclass.render_line = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_line)
	}

	if overrides.RenderOption != nil {
		pclass.render_option = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_option)
	}

	if overrides.RenderSlider != nil {
		pclass.render_slider = (*[0]byte)(C._gotk4_gtk3_ThemingEngineClass_render_slider)
	}

	if classInitFunc != nil {
		class := (*ThemingEngineClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapThemingEngine(obj *coreglib.Object) *ThemingEngine {
	return &ThemingEngine{
		Object: obj,
	}
}

func marshalThemingEngine(p uintptr) (interface{}, error) {
	return wrapThemingEngine(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Screen returns the Screen to which engine currently rendering to.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - screen (optional) or NULL.
//
func (engine *ThemingEngine) Screen() *gdk.Screen {
	var _arg0 *C.GtkThemingEngine // out
	var _cret *C.GdkScreen        // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))

	_cret = C.gtk_theming_engine_get_screen(_arg0)
	runtime.KeepAlive(engine)

	var _screen *gdk.Screen // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_screen = &gdk.Screen{
				Object: obj,
			}
		}
	}

	return _screen
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderActivity(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_activity

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_activity(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - angle
//    - x
//    - y
//    - size
//
func (engine *ThemingEngine) renderArrow(cr *cairo.Context, angle, x, y, size float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_arrow

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(angle)
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(size)

	C._gotk4_gtk3_ThemingEngine_virtual_render_arrow(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(angle)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(size)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderBackground(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_background

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_background(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderCheck(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_check

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_check(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderExpander(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_expander

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_expander(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//    - gapSide
//
func (engine *ThemingEngine) renderExtension(cr *cairo.Context, x, y, width, height float64, gapSide PositionType) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_extension

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out
	var _arg6 C.GtkPositionType   // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)
	_arg6 = C.GtkPositionType(gapSide)

	C._gotk4_gtk3_ThemingEngine_virtual_render_extension(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(gapSide)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderFocus(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_focus

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_focus(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderFrame(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_frame

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_frame(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//    - gapSide
//    - xy0Gap
//    - xy1Gap
//
func (engine *ThemingEngine) renderFrameGap(cr *cairo.Context, x, y, width, height float64, gapSide PositionType, xy0Gap, xy1Gap float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_frame_gap

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out
	var _arg6 C.GtkPositionType   // out
	var _arg7 C.gdouble           // out
	var _arg8 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)
	_arg6 = C.GtkPositionType(gapSide)
	_arg7 = C.gdouble(xy0Gap)
	_arg8 = C.gdouble(xy1Gap)

	C._gotk4_gtk3_ThemingEngine_virtual_render_frame_gap(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(gapSide)
	runtime.KeepAlive(xy0Gap)
	runtime.KeepAlive(xy1Gap)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderHandle(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_handle

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_handle(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - pixbuf
//    - x
//    - y
//
func (engine *ThemingEngine) renderIcon(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x, y float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_icon

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.GdkPixbuf        // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)

	C._gotk4_gtk3_ThemingEngine_virtual_render_icon(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(pixbuf)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// The function takes the following parameters:
//
//    - cr
//    - surface
//    - x
//    - y
//
func (engine *ThemingEngine) renderIconSurface(cr *cairo.Context, surface *cairo.Surface, x, y float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_icon_surface

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.cairo_surface_t  // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)

	C._gotk4_gtk3_ThemingEngine_virtual_render_icon_surface(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - layout
//
func (engine *ThemingEngine) renderLayout(cr *cairo.Context, x, y float64, layout *pango.Layout) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_layout

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 *C.PangoLayout      // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	C._gotk4_gtk3_ThemingEngine_virtual_render_layout(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(layout)
}

// The function takes the following parameters:
//
//    - cr
//    - x0
//    - y0
//    - x1
//    - y1
//
func (engine *ThemingEngine) renderLine(cr *cairo.Context, x0, y0, x1, y1 float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_line

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x0)
	_arg3 = C.gdouble(y0)
	_arg4 = C.gdouble(x1)
	_arg5 = C.gdouble(y1)

	C._gotk4_gtk3_ThemingEngine_virtual_render_line(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x0)
	runtime.KeepAlive(y0)
	runtime.KeepAlive(x1)
	runtime.KeepAlive(y1)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//
func (engine *ThemingEngine) renderOption(cr *cairo.Context, x, y, width, height float64) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_option

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C._gotk4_gtk3_ThemingEngine_virtual_render_option(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// The function takes the following parameters:
//
//    - cr
//    - x
//    - y
//    - width
//    - height
//    - orientation
//
func (engine *ThemingEngine) renderSlider(cr *cairo.Context, x, y, width, height float64, orientation Orientation) {
	gclass := (*C.GtkThemingEngineClass)(coreglib.PeekParentClass(engine))
	fnarg := gclass.render_slider

	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.cairo_t          // out
	var _arg2 C.gdouble           // out
	var _arg3 C.gdouble           // out
	var _arg4 C.gdouble           // out
	var _arg5 C.gdouble           // out
	var _arg6 C.GtkOrientation    // out

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)
	_arg6 = C.GtkOrientation(orientation)

	C._gotk4_gtk3_ThemingEngine_virtual_render_slider(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(orientation)
}

// ThemingEngineLoad loads and initializes a theming engine module from the
// standard directories.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - name: theme engine name to load.
//
// The function returns the following values:
//
//    - themingEngine (optional): theming engine, or NULL if the engine name
//      doesn’t exist.
//
func ThemingEngineLoad(name string) *ThemingEngine {
	var _arg1 *C.gchar            // out
	var _cret *C.GtkThemingEngine // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_load(_arg1)
	runtime.KeepAlive(name)

	var _themingEngine *ThemingEngine // out

	if _cret != nil {
		_themingEngine = wrapThemingEngine(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _themingEngine
}

// ThemingEngineClass: base class for theming engines.
//
// An instance of this type is always passed by reference.
type ThemingEngineClass struct {
	*themingEngineClass
}

// themingEngineClass is the struct that's finalized.
type themingEngineClass struct {
	native *C.GtkThemingEngineClass
}
