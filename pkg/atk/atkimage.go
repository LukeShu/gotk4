// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_image_get_type()), F: marshalImager},
	})
}

// ImageOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ImageOverrider interface {
	// ImageDescription: get a textual description of this image.
	ImageDescription() string
	// ImageLocale retrieves the locale identifier associated to the Image.
	ImageLocale() string
	// ImagePosition gets the position of the image in the form of a point
	// specifying the images top-left corner.
	//
	// If the position can not be obtained (e.g. missing support), x and y are
	// set to -1.
	ImagePosition(coordType CoordType) (x int32, y int32)
	// ImageSize: get the width and height in pixels for the specified image.
	// The values of width and height are returned as -1 if the values cannot be
	// obtained (for instance, if the object is not onscreen).
	//
	// If the size can not be obtained (e.g. missing support), x and y are set
	// to -1.
	ImageSize() (width int32, height int32)
	// SetImageDescription sets the textual description for this image.
	SetImageDescription(description string) bool
}

// Image should be implemented by Object subtypes on behalf of components which
// display image/pixmap information onscreen, and which provide information
// (other than just widget borders, etc.) via that image content. For instance,
// icons, buttons with icons, toolbar elements, and image viewing panes
// typically should implement Image.
//
// Image primarily provides two types of information: coordinate information
// (useful for screen review mode of screenreaders, and for use by onscreen
// magnifiers), and descriptive information. The descriptive information is
// provided for alternative, text-only presentation of the most significant
// information present in the image.
type Image struct {
	*externglib.Object
}

// Imager describes Image's abstract methods.
type Imager interface {
	externglib.Objector

	// ImageDescription: get a textual description of this image.
	ImageDescription() string
	// ImageLocale retrieves the locale identifier associated to the Image.
	ImageLocale() string
	// ImagePosition gets the position of the image in the form of a point
	// specifying the images top-left corner.
	ImagePosition(coordType CoordType) (x int32, y int32)
	// ImageSize: get the width and height in pixels for the specified image.
	ImageSize() (width int32, height int32)
	// SetImageDescription sets the textual description for this image.
	SetImageDescription(description string) bool
}

var _ Imager = (*Image)(nil)

func wrapImage(obj *externglib.Object) *Image {
	return &Image{
		Object: obj,
	}
}

func marshalImager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapImage(obj), nil
}

// ImageDescription: get a textual description of this image.
func (image *Image) ImageDescription() string {
	var _arg0 *C.AtkImage // out
	var _cret *C.gchar    // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.atk_image_get_image_description(_arg0)
	runtime.KeepAlive(image)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ImageLocale retrieves the locale identifier associated to the Image.
func (image *Image) ImageLocale() string {
	var _arg0 *C.AtkImage // out
	var _cret *C.gchar    // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.atk_image_get_image_locale(_arg0)
	runtime.KeepAlive(image)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ImagePosition gets the position of the image in the form of a point
// specifying the images top-left corner.
//
// If the position can not be obtained (e.g. missing support), x and y are set
// to -1.
func (image *Image) ImagePosition(coordType CoordType) (x int32, y int32) {
	var _arg0 *C.AtkImage    // out
	var _arg1 C.gint         // in
	var _arg2 C.gint         // in
	var _arg3 C.AtkCoordType // out

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))
	_arg3 = C.AtkCoordType(coordType)

	C.atk_image_get_image_position(_arg0, &_arg1, &_arg2, _arg3)
	runtime.KeepAlive(image)
	runtime.KeepAlive(coordType)

	var _x int32 // out
	var _y int32 // out

	_x = int32(_arg1)
	_y = int32(_arg2)

	return _x, _y
}

// ImageSize: get the width and height in pixels for the specified image. The
// values of width and height are returned as -1 if the values cannot be
// obtained (for instance, if the object is not onscreen).
//
// If the size can not be obtained (e.g. missing support), x and y are set to
// -1.
func (image *Image) ImageSize() (width int32, height int32) {
	var _arg0 *C.AtkImage // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))

	C.atk_image_get_image_size(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(image)

	var _width int32  // out
	var _height int32 // out

	_width = int32(_arg1)
	_height = int32(_arg2)

	return _width, _height
}

// SetImageDescription sets the textual description for this image.
func (image *Image) SetImageDescription(description string) bool {
	var _arg0 *C.AtkImage // out
	var _arg1 *C.gchar    // out
	var _cret C.gboolean  // in

	_arg0 = (*C.AtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_image_set_image_description(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(description)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
