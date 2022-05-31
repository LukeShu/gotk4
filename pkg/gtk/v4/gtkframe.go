// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_FrameClass_compute_child_allocation(GtkFrame*, GtkAllocation*);
import "C"

// glib.Type values for gtkframe.go.
var GTypeFrame = coreglib.Type(C.gtk_frame_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeFrame, F: marshalFrame},
	})
}

// FrameOverrider contains methods that are overridable.
type FrameOverrider interface {
	// The function takes the following parameters:
	//
	ComputeChildAllocation(allocation *Allocation)
}

// Frame: GtkFrame is a widget that surrounds its child with a decorative frame
// and an optional label.
//
// !An example GtkFrame (frame.png)
//
// If present, the label is drawn inside the top edge of the frame. The
// horizontal position of the label can be controlled with
// gtk.Frame.SetLabelAlign().
//
// GtkFrame clips its child. You can use this to add rounded corners to widgets,
// but be aware that it also cuts off shadows.
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
//        <object class="GtkLabel" id="frame_label"/>
//      </child>
//      <child>
//        <object class="GtkEntry" id="frame_content"/>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    frame
//    ├── <label widget>
//    ╰── <child>
//
//
// GtkFrame has a main CSS node with name “frame”, which is used to draw the
// visible border. You can set the appearance of the border using CSS properties
// like “border-style” on this node.
type Frame struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Frame)(nil)
)

func classInitFramer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkFrameClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkFrameClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ComputeChildAllocation(allocation *Allocation) }); ok {
		pclass.compute_child_allocation = (*[0]byte)(C._gotk4_gtk4_FrameClass_compute_child_allocation)
	}
}

//export _gotk4_gtk4_FrameClass_compute_child_allocation
func _gotk4_gtk4_FrameClass_compute_child_allocation(arg0 *C.GtkFrame, arg1 *C.GtkAllocation) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ComputeChildAllocation(allocation *Allocation) })

	var _allocation *Allocation // out

	_allocation = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	iface.ComputeChildAllocation(_allocation)
}

func wrapFrame(obj *coreglib.Object) *Frame {
	return &Frame{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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
	}
}

func marshalFrame(p uintptr) (interface{}, error) {
	return wrapFrame(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFrame creates a new GtkFrame, with optional label label.
//
// If label is NULL, the label is omitted.
//
// The function takes the following parameters:
//
//    - label (optional): text to use as the label of the frame.
//
// The function returns the following values:
//
//    - frame: new GtkFrame widget.
//
func NewFrame(label string) *Frame {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	if label != "" {
		_arg0 = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg0))
	}
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Frame").InvokeMethod("new_Frame", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(label)

	var _frame *Frame // out

	_frame = wrapFrame(coreglib.Take(unsafe.Pointer(_cret)))

	return _frame
}

// Child gets the child widget of frame.
//
// The function returns the following values:
//
//    - widget (optional): child widget of frame.
//
func (frame *Frame) Child() Widgetter {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	*(**Frame)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Frame").InvokeMethod("get_child", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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

// Label returns the frame labels text.
//
// If the frame's label widget is not a GtkLabel, NULL is returned.
//
// The function returns the following values:
//
//    - utf8 (optional): text in the label, or NULL if there was no label widget
//      or the label widget was not a GtkLabel. This string is owned by GTK and
//      must not be modified or freed.
//
func (frame *Frame) Label() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	*(**Frame)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Frame").InvokeMethod("get_label", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frame)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelAlign retrieves the X alignment of the frame’s label.
//
// The function returns the following values:
//
//    - gfloat frames X alignment.
//
func (frame *Frame) LabelAlign() float32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.float // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	*(**Frame)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Frame").InvokeMethod("get_label_align", args[:], nil)
	_cret = *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(frame)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// LabelWidget retrieves the label widget for the frame.
//
// The function returns the following values:
//
//    - widget (optional): label widget, or NULL if there is none.
//
func (frame *Frame) LabelWidget() Widgetter {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	*(**Frame)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Frame").InvokeMethod("get_label_widget", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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

// SetChild sets the child widget of frame.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (frame *Frame) SetChild(child Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	if child != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}
	*(**Frame)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Frame").InvokeMethod("set_child", args[:], nil)

	runtime.KeepAlive(frame)
	runtime.KeepAlive(child)
}

// SetLabel creates a new GtkLabel with the label and sets it as the frame's
// label widget.
//
// The function takes the following parameters:
//
//    - label (optional): text to use as the label of the frame.
//
func (frame *Frame) SetLabel(label string) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	if label != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**Frame)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Frame").InvokeMethod("set_label", args[:], nil)

	runtime.KeepAlive(frame)
	runtime.KeepAlive(label)
}

// SetLabelAlign sets the X alignment of the frame widget’s label.
//
// The default value for a newly created frame is 0.0.
//
// The function takes the following parameters:
//
//    - xalign: position of the label along the top edge of the widget. A value
//      of 0.0 represents left alignment; 1.0 represents right alignment.
//
func (frame *Frame) SetLabelAlign(xalign float32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.float // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	_arg1 = C.float(xalign)
	*(**Frame)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Frame").InvokeMethod("set_label_align", args[:], nil)

	runtime.KeepAlive(frame)
	runtime.KeepAlive(xalign)
}

// SetLabelWidget sets the label widget for the frame.
//
// This is the widget that will appear embedded in the top edge of the frame as
// a title.
//
// The function takes the following parameters:
//
//    - labelWidget (optional): new label widget.
//
func (frame *Frame) SetLabelWidget(labelWidget Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(frame).Native()))
	if labelWidget != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(labelWidget).Native()))
	}
	*(**Frame)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Frame").InvokeMethod("set_label_widget", args[:], nil)

	runtime.KeepAlive(frame)
	runtime.KeepAlive(labelWidget)
}
