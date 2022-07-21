// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_FrameClass_compute_child_allocation(GtkFrame*, GtkAllocation*);
import "C"

// GTypeFrame returns the GType for the type Frame.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFrame() coreglib.Type {
	gtype := coreglib.Type(C.gtk_frame_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalFrame)
	return gtype
}

// FrameOverrider contains methods that are overridable.
type FrameOverrider interface {
	// The function takes the following parameters:
	//
	ComputeChildAllocation(allocation *Allocation)
}

// Frame: frame widget is a bin that surrounds its child with a decorative frame
// and an optional label. If present, the label is drawn in a gap in the top
// side of the frame. The position of the label can be controlled with
// gtk_frame_set_label_align().
//
//
// GtkFrame as GtkBuildable
//
// The GtkFrame implementation of the GtkBuildable interface supports placing a
// child in the label position by specifying “label” as the “type” attribute of
// a <child> element. A normal content child can be specified without specifying
// a <child> type attribute.
//
// An example of a UI definition fragment with GtkFrame:
//
//    <object class="GtkFrame">
//      <child type="label">
//        <object class="GtkLabel" id="frame-label"/>
//      </child>
//      <child>
//        <object class="GtkEntry" id="frame-content"/>
//      </child>
//    </object>
//
// CSS nodes
//
//    frame
//    ├── border[.flat]
//    ├── <label widget>
//    ╰── <child>
//
// GtkFrame has a main CSS node named “frame” and a subnode named “border”. The
// “border” node is used to draw the visible border. You can set the appearance
// of the border using CSS properties like “border-style” on the “border” node.
//
// The border node can be given the style class “.flat”, which is used by themes
// to disable drawing of the border. To do this from code, call
// gtk_frame_set_shadow_type() with GTK_SHADOW_NONE to add the “.flat” class or
// any other shadow type to remove it.
type Frame struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*Frame)(nil)
)

func classInitFramer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkFrameClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ ComputeChildAllocation(allocation *Allocation) }); ok {
		pclass.compute_child_allocation = (*[0]byte)(C._gotk4_gtk3_FrameClass_compute_child_allocation)
	}
}

//export _gotk4_gtk3_FrameClass_compute_child_allocation
func _gotk4_gtk3_FrameClass_compute_child_allocation(arg0 *C.GtkFrame, arg1 *C.GtkAllocation) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ComputeChildAllocation(allocation *Allocation) })

	var _allocation *Allocation // out

	_allocation = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.ComputeChildAllocation(_allocation)
}

func wrapFrame(obj *coreglib.Object) *Frame {
	return &Frame{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
			},
		},
	}
}

func marshalFrame(p uintptr) (interface{}, error) {
	return wrapFrame(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFrame creates a new Frame, with optional label label. If label is NULL,
// the label is omitted.
//
// The function takes the following parameters:
//
//    - label (optional): text to use as the label of the frame.
//
// The function returns the following values:
//
//    - frame: new Frame widget.
//
func NewFrame(label string) *Frame {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_frame_new(_arg1)
	runtime.KeepAlive(label)

	var _frame *Frame // out

	_frame = wrapFrame(coreglib.Take(unsafe.Pointer(_cret)))

	return _frame
}

// Label: if the frame’s label widget is a Label, returns the text in the label
// widget. (The frame will have a Label for the label widget if a non-NULL
// argument was passed to gtk_frame_new().).
//
// The function returns the following values:
//
//    - utf8 (optional): text in the label, or NULL if there was no label widget
//      or the lable widget was not a Label. This string is owned by GTK+ and
//      must not be modified or freed.
//
func (frame *Frame) Label() string {
	var _arg0 *C.GtkFrame // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))

	_cret = C.gtk_frame_get_label(_arg0)
	runtime.KeepAlive(frame)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelAlign retrieves the X and Y alignment of the frame’s label. See
// gtk_frame_set_label_align().
//
// The function returns the following values:
//
//    - xalign (optional): location to store X alignment of frame’s label, or
//      NULL.
//    - yalign (optional): location to store X alignment of frame’s label, or
//      NULL.
//
func (frame *Frame) LabelAlign() (xalign, yalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.gfloat    // in
	var _arg2 C.gfloat    // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))

	C.gtk_frame_get_label_align(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(frame)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_frame_set_label_widget().
//
// The function returns the following values:
//
//    - widget (optional): label widget, or NULL if there is none.
//
func (frame *Frame) LabelWidget() Widgetter {
	var _arg0 *C.GtkFrame  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))

	_cret = C.gtk_frame_get_label_widget(_arg0)
	runtime.KeepAlive(frame)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// ShadowType retrieves the shadow type of the frame. See
// gtk_frame_set_shadow_type().
//
// The function returns the following values:
//
//    - shadowType: current shadow type of the frame.
//
func (frame *Frame) ShadowType() ShadowType {
	var _arg0 *C.GtkFrame     // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))

	_cret = C.gtk_frame_get_shadow_type(_arg0)
	runtime.KeepAlive(frame)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

// SetLabel removes the current Frame:label-widget. If label is not NULL,
// creates a new Label with that text and adds it as the Frame:label-widget.
//
// The function takes the following parameters:
//
//    - label (optional): text to use as the label of the frame.
//
func (frame *Frame) SetLabel(label string) {
	var _arg0 *C.GtkFrame // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_frame_set_label(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(label)
}

// SetLabelAlign sets the alignment of the frame widget’s label. The default
// values for a newly created frame are 0.0 and 0.5.
//
// The function takes the following parameters:
//
//    - xalign: position of the label along the top edge of the widget. A value
//      of 0.0 represents left alignment; 1.0 represents right alignment.
//    - yalign: y alignment of the label. A value of 0.0 aligns under the frame;
//      1.0 aligns above the frame. If the values are exactly 0.0 or 1.0 the gap
//      in the frame won’t be painted because the label will be completely above
//      or below the frame.
//
func (frame *Frame) SetLabelAlign(xalign, yalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.gfloat    // out
	var _arg2 C.gfloat    // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)

	C.gtk_frame_set_label_align(_arg0, _arg1, _arg2)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
}

// SetLabelWidget sets the Frame:label-widget for the frame. This is the widget
// that will appear embedded in the top edge of the frame as a title.
//
// The function takes the following parameters:
//
//    - labelWidget (optional): new label widget.
//
func (frame *Frame) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkFrame  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	if labelWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(labelWidget).Native()))
	}

	C.gtk_frame_set_label_widget(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(labelWidget)
}

// SetShadowType sets the Frame:shadow-type for frame, i.e. whether it is drawn
// without (GTK_SHADOW_NONE) or with (other values) a visible border. Values
// other than GTK_SHADOW_NONE are treated identically by GtkFrame. The chosen
// type is applied by removing or adding the .flat class to the CSS node named
// border.
//
// The function takes the following parameters:
//
//    - typ: new ShadowType.
//
func (frame *Frame) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkFrame     // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_frame_set_shadow_type(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(typ)
}
