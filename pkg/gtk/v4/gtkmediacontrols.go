// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeMediaControls = coreglib.Type(C.gtk_media_controls_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMediaControls, F: marshalMediaControls},
	})
}

// MediaControlsOverrider contains methods that are overridable.
type MediaControlsOverrider interface {
}

// MediaControls: GtkMediaControls is a widget to show controls for a video.
//
// !An example GtkMediaControls (media-controls.png)
//
// Usually, GtkMediaControls is used as part of gtk.Video.
type MediaControls struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*MediaControls)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:         GTypeMediaControls,
		GoType:        reflect.TypeOf((*MediaControls)(nil)),
		InitClass:     initClassMediaControls,
		FinalizeClass: finalizeClassMediaControls,
	})
}

func initClassMediaControls(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitMediaControls(*MediaControlsClass) }); ok {
		klass := (*MediaControlsClass)(gextras.NewStructNative(gclass))
		goval.InitMediaControls(klass)
	}
}

func finalizeClassMediaControls(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ FinalizeMediaControls(*MediaControlsClass) }); ok {
		klass := (*MediaControlsClass)(gextras.NewStructNative(gclass))
		goval.FinalizeMediaControls(klass)
	}
}

func wrapMediaControls(obj *coreglib.Object) *MediaControls {
	return &MediaControls{
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

func marshalMediaControls(p uintptr) (interface{}, error) {
	return wrapMediaControls(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMediaControls creates a new GtkMediaControls managing the stream passed to
// it.
//
// The function takes the following parameters:
//
//    - stream (optional) to manage or NULL for none.
//
// The function returns the following values:
//
//    - mediaControls: new GtkMediaControls.
//
func NewMediaControls(stream MediaStreamer) *MediaControls {
	var _arg1 *C.GtkMediaStream // out
	var _cret *C.GtkWidget      // in

	if stream != nil {
		_arg1 = (*C.GtkMediaStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	}

	_cret = C.gtk_media_controls_new(_arg1)
	runtime.KeepAlive(stream)

	var _mediaControls *MediaControls // out

	_mediaControls = wrapMediaControls(coreglib.Take(unsafe.Pointer(_cret)))

	return _mediaControls
}

// MediaStream gets the media stream managed by controls or NULL if none.
//
// The function returns the following values:
//
//    - mediaStream (optional): media stream managed by controls.
//
func (controls *MediaControls) MediaStream() MediaStreamer {
	var _arg0 *C.GtkMediaControls // out
	var _cret *C.GtkMediaStream   // in

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(coreglib.InternObject(controls).Native()))

	_cret = C.gtk_media_controls_get_media_stream(_arg0)
	runtime.KeepAlive(controls)

	var _mediaStream MediaStreamer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(MediaStreamer)
				return ok
			})
			rv, ok := casted.(MediaStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.MediaStreamer")
			}
			_mediaStream = rv
		}
	}

	return _mediaStream
}

// SetMediaStream sets the stream that is controlled by controls.
//
// The function takes the following parameters:
//
//    - stream (optional): GtkMediaStream, or NULL.
//
func (controls *MediaControls) SetMediaStream(stream MediaStreamer) {
	var _arg0 *C.GtkMediaControls // out
	var _arg1 *C.GtkMediaStream   // out

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(coreglib.InternObject(controls).Native()))
	if stream != nil {
		_arg1 = (*C.GtkMediaStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	}

	C.gtk_media_controls_set_media_stream(_arg0, _arg1)
	runtime.KeepAlive(controls)
	runtime.KeepAlive(stream)
}

// MediaControlsClass: instance of this type is always passed by reference.
type MediaControlsClass struct {
	*mediaControlsClass
}

// mediaControlsClass is the struct that's finalized.
type mediaControlsClass struct {
	native *C.GtkMediaControlsClass
}

func (m *MediaControlsClass) ParentClass() *WidgetClass {
	valptr := &m.native.parent_class
	var _v *WidgetClass // out
	_v = (*WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
