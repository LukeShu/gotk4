// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeAspectFrame returns the GType for the type AspectFrame.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAspectFrame() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "AspectFrame").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAspectFrame)
	return gtype
}

// AspectFrame: GtkAspectFrame preserves the aspect ratio of its child.
//
// The frame can respect the aspect ratio of the child widget, or use its own
// aspect ratio.
//
//
// CSS nodes
//
// GtkAspectFrame uses a CSS node with name frame.
type AspectFrame struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*AspectFrame)(nil)
)

func wrapAspectFrame(obj *coreglib.Object) *AspectFrame {
	return &AspectFrame{
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

func marshalAspectFrame(p uintptr) (interface{}, error) {
	return wrapAspectFrame(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAspectFrame: create a new GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment of the child within the parent. Ranges from
//      0.0 (left aligned) to 1.0 (right aligned).
//    - yalign: vertical alignment of the child within the parent. Ranges from
//      0.0 (top aligned) to 1.0 (bottom aligned).
//    - ratio: desired aspect ratio.
//    - obeyChild: if TRUE, ratio is ignored, and the aspect ratio is taken from
//      the requistion of the child.
//
// The function returns the following values:
//
//    - aspectFrame: new GtkAspectFrame.
//
func NewAspectFrame(xalign, yalign, ratio float32, obeyChild bool) *AspectFrame {
	var _args [4]girepository.Argument

	*(*C.float)(unsafe.Pointer(&_args[0])) = C.float(xalign)
	*(*C.float)(unsafe.Pointer(&_args[1])) = C.float(yalign)
	*(*C.float)(unsafe.Pointer(&_args[2])) = C.float(ratio)
	if obeyChild {
		*(*C.gboolean)(unsafe.Pointer(&_args[3])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_gret := _info.InvokeClassMethod("new_AspectFrame", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
	runtime.KeepAlive(ratio)
	runtime.KeepAlive(obeyChild)

	var _aspectFrame *AspectFrame // out

	_aspectFrame = wrapAspectFrame(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _aspectFrame
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//    - widget (optional): child widget of self@.
//
func (self *AspectFrame) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_gret := _info.InvokeClassMethod("get_child", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

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

// ObeyChild returns whether the child's size request should override the set
// aspect ratio of the GtkAspectFrame.
//
// The function returns the following values:
//
//    - ok: whether to obey the child's size request.
//
func (self *AspectFrame) ObeyChild() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_gret := _info.InvokeClassMethod("get_obey_child", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Ratio returns the desired aspect ratio of the child.
//
// The function returns the following values:
//
//    - gfloat: desired aspect ratio.
//
func (self *AspectFrame) Ratio() float32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_gret := _info.InvokeClassMethod("get_ratio", _args[:], nil)
	_cret := *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(*(*C.float)(unsafe.Pointer(&_cret)))

	return _gfloat
}

// XAlign returns the horizontal alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function returns the following values:
//
//    - gfloat: horizontal alignment.
//
func (self *AspectFrame) XAlign() float32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_gret := _info.InvokeClassMethod("get_xalign", _args[:], nil)
	_cret := *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(*(*C.float)(unsafe.Pointer(&_cret)))

	return _gfloat
}

// YAlign returns the vertical alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function returns the following values:
//
//    - gfloat: vertical alignment.
//
func (self *AspectFrame) YAlign() float32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_gret := _info.InvokeClassMethod("get_yalign", _args[:], nil)
	_cret := *(*C.float)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gfloat float32 // out

	_gfloat = float32(*(*C.float)(unsafe.Pointer(&_cret)))

	return _gfloat
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (self *AspectFrame) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_info.InvokeClassMethod("set_child", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetObeyChild sets whether the aspect ratio of the child's size request should
// override the set aspect ratio of the GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - obeyChild: if TRUE, ratio is ignored, and the aspect ratio is taken from
//      the requistion of the child.
//
func (self *AspectFrame) SetObeyChild(obeyChild bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if obeyChild {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_info.InvokeClassMethod("set_obey_child", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(obeyChild)
}

// SetRatio sets the desired aspect ratio of the child.
//
// The function takes the following parameters:
//
//    - ratio: aspect ratio of the child.
//
func (self *AspectFrame) SetRatio(ratio float32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.float)(unsafe.Pointer(&_args[1])) = C.float(ratio)

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_info.InvokeClassMethod("set_ratio", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(ratio)
}

// SetXAlign sets the horizontal alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment, from 0.0 (left aligned) to 1.0 (right
//      aligned).
//
func (self *AspectFrame) SetXAlign(xalign float32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.float)(unsafe.Pointer(&_args[1])) = C.float(xalign)

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_info.InvokeClassMethod("set_xalign", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(xalign)
}

// SetYAlign sets the vertical alignment of the child within the allocation of
// the GtkAspectFrame.
//
// The function takes the following parameters:
//
//    - yalign: horizontal alignment, from 0.0 (top aligned) to 1.0 (bottom
//      aligned).
//
func (self *AspectFrame) SetYAlign(yalign float32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.float)(unsafe.Pointer(&_args[1])) = C.float(yalign)

	_info := girepository.MustFind("Gtk", "AspectFrame")
	_info.InvokeClassMethod("set_yalign", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(yalign)
}
