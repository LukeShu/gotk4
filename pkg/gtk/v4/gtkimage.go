// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_image_type_get_type()), F: marshalImageType},
		{T: externglib.Type(C.gtk_image_get_type()), F: marshalImager},
	})
}

// ImageType describes the image data representation used by a gtk.Image.
//
// If you want to get the image from the widget, you can only get the
// currently-stored representation; for instance, if the
// gtk_image_get_storage_type() returns GTK_IMAGE_PAINTABLE, then you can call
// gtk_image_get_paintable().
//
// For empty images, you can request any storage type (call any of the "get"
// functions), but they will all return NULL values.
type ImageType C.gint

const (
	// ImageEmpty: there is no image displayed by the widget.
	ImageEmpty ImageType = iota
	// ImageIconName: widget contains a named icon.
	ImageIconName
	// ImageGIcon: widget contains a #GIcon.
	ImageGIcon
	// ImagePaintable: widget contains a Paintable.
	ImagePaintable
)

func marshalImageType(p uintptr) (interface{}, error) {
	return ImageType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ImageType.
func (i ImageType) String() string {
	switch i {
	case ImageEmpty:
		return "Empty"
	case ImageIconName:
		return "IconName"
	case ImageGIcon:
		return "GIcon"
	case ImagePaintable:
		return "Paintable"
	default:
		return fmt.Sprintf("ImageType(%d)", i)
	}
}

// Image: GtkImage widget displays an image.
//
// !An example GtkImage (image.png)
//
// Various kinds of object can be displayed as an image; most typically, you
// would load a GdkTexture from a file, using the convenience function
// gtk.Image.NewFromFile, for instance:
//
//    GtkWidget *image = gtk_image_new_from_file ("myfile.png");
//
//
// If the file isn’t loaded successfully, the image will contain a “broken
// image” icon similar to that used in many web browsers.
//
// If you want to handle errors in loading the file yourself, for example by
// displaying an error message, then load the image with
// gdk.Texture.NewFromFile, then create the GtkImage with
// gtk.Image.NewFromPaintable.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of GResource inside GIO, for
// details. In this case, gtk.Image:resource, gtk.Image.NewFromResource, and
// gtk.Image.SetFromResource() should be used.
//
// GtkImage displays its image as an icon, with a size that is determined by the
// application. See gtk.Picture if you want to show an image at is actual size.
//
//
// CSS nodes
//
// GtkImage has a single CSS node with the name image. The style classes
// .normal-icons or .large-icons may appear, depending on the
// gtk.Image:icon-size property.
//
//
// Accessibility
//
// GtkImage uses the GTK_ACCESSIBLE_ROLE_IMG role.
type Image struct {
	Widget
}

var (
	_ Widgetter = (*Image)(nil)
)

func wrapImage(obj *externglib.Object) *Image {
	return &Image{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
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

func marshalImager(p uintptr) (interface{}, error) {
	return wrapImage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewImage creates a new empty GtkImage widget.
func NewImage() *Image {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_image_new()

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromFile creates a new GtkImage displaying the file filename.
//
// If the file isn’t found or can’t be loaded, the resulting GtkImage will
// display a “broken image” icon. This function never returns NULL, it always
// returns a valid GtkImage widget.
//
// If you need to detect failures to load the file, use gdk.Texture.NewFromFile
// to load the file yourself, then create the GtkImage from the texture.
//
// The storage type (see gtk.Image.GetStorageType()) of the returned image is
// not defined, it will be whatever is appropriate for displaying the file.
//
// The function takes the following parameters:
//
//    - filename: filename.
//
func NewImageFromFile(filename string) *Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_file(_arg1)
	runtime.KeepAlive(filename)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromGIcon creates a GtkImage displaying an icon from the current icon
// theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// The function takes the following parameters:
//
//    - icon: icon.
//
func NewImageFromGIcon(icon gio.Iconner) *Image {
	var _arg1 *C.GIcon     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.gtk_image_new_from_gicon(_arg1)
	runtime.KeepAlive(icon)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromIconName creates a GtkImage displaying an icon from the current
// icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// The function takes the following parameters:
//
//    - iconName: icon name or NULL.
//
func NewImageFromIconName(iconName string) *Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_image_new_from_icon_name(_arg1)
	runtime.KeepAlive(iconName)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPaintable creates a new GtkImage displaying paintable.
//
// The GtkImage does not assume a reference to the paintable; you still need to
// unref it if you own references. GtkImage will add its own reference rather
// than adopting yours.
//
// The GtkImage will track changes to the paintable and update its size and
// contents in response to it.
//
// The function takes the following parameters:
//
//    - paintable: GdkPaintable, or NULL.
//
func NewImageFromPaintable(paintable gdk.Paintabler) *Image {
	var _arg1 *C.GdkPaintable // out
	var _cret *C.GtkWidget    // in

	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))
	}

	_cret = C.gtk_image_new_from_paintable(_arg1)
	runtime.KeepAlive(paintable)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPixbuf creates a new GtkImage displaying pixbuf.
//
// The GtkImage does not assume a reference to the pixbuf; you still need to
// unref it if you own references. GtkImage will add its own reference rather
// than adopting yours.
//
// This is a helper for gtk.Image.NewFromPaintable, and you can't get back the
// exact pixbuf once this is called, only a texture.
//
// Note that this function just creates an GtkImage from the pixbuf. The
// GtkImage created will not react to state changes. Should you want that, you
// should use gtk.Image.NewFromIconName.
//
// The function takes the following parameters:
//
//    - pixbuf: GdkPixbuf, or NULL.
//
func NewImageFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Image {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	}

	_cret = C.gtk_image_new_from_pixbuf(_arg1)
	runtime.KeepAlive(pixbuf)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromResource creates a new GtkImage displaying the resource file
// resource_path.
//
// If the file isn’t found or can’t be loaded, the resulting GtkImage will
// display a “broken image” icon. This function never returns NULL, it always
// returns a valid GtkImage widget.
//
// If you need to detect failures to load the file, use
// gdkpixbuf.Pixbuf.NewFromFile to load the file yourself, then create the
// GtkImage from the pixbuf.
//
// The storage type (see gtk.Image.GetStorageType()) of the returned image is
// not defined, it will be whatever is appropriate for displaying the file.
//
// The function takes the following parameters:
//
//    - resourcePath: resource path.
//
func NewImageFromResource(resourcePath string) *Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_resource(_arg1)
	runtime.KeepAlive(resourcePath)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// Clear resets the image to be empty.
func (image *Image) Clear() {
	var _arg0 *C.GtkImage // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	C.gtk_image_clear(_arg0)
	runtime.KeepAlive(image)
}

// GIcon gets the GIcon being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk.Image.GetStorageType()). The caller of this function does not own a
// reference to the returned GIcon.
func (image *Image) GIcon() gio.Iconner {
	var _arg0 *C.GtkImage // out
	var _cret *C.GIcon    // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_gicon(_arg0)
	runtime.KeepAlive(image)

	var _icon gio.Iconner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(gio.Iconner)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.Iconner")
			}
			_icon = rv
		}
	}

	return _icon
}

// IconName gets the icon name and size being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME
// (see gtk.Image.GetStorageType()). The returned string is owned by the
// GtkImage and should not be freed.
func (image *Image) IconName() string {
	var _arg0 *C.GtkImage // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_icon_name(_arg0)
	runtime.KeepAlive(image)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// IconSize gets the icon size used by the image when rendering icons.
func (image *Image) IconSize() IconSize {
	var _arg0 *C.GtkImage   // out
	var _cret C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_icon_size(_arg0)
	runtime.KeepAlive(image)

	var _iconSize IconSize // out

	_iconSize = IconSize(_cret)

	return _iconSize
}

// Paintable gets the image GdkPaintable being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PAINTABLE
// (see gtk.Image.GetStorageType()). The caller of this function does not own a
// reference to the returned paintable.
func (image *Image) Paintable() gdk.Paintabler {
	var _arg0 *C.GtkImage     // out
	var _cret *C.GdkPaintable // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_paintable(_arg0)
	runtime.KeepAlive(image)

	var _paintable gdk.Paintabler // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(gdk.Paintabler)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gdk.Paintabler")
			}
			_paintable = rv
		}
	}

	return _paintable
}

// PixelSize gets the pixel size used for named icons.
func (image *Image) PixelSize() int {
	var _arg0 *C.GtkImage // out
	var _cret C.int       // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_pixel_size(_arg0)
	runtime.KeepAlive(image)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// StorageType gets the type of representation being used by the GtkImage to
// store image data.
//
// If the GtkImage has no image data, the return value will be GTK_IMAGE_EMPTY.
func (image *Image) StorageType() ImageType {
	var _arg0 *C.GtkImage    // out
	var _cret C.GtkImageType // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_storage_type(_arg0)
	runtime.KeepAlive(image)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

// SetFromFile sets a GtkImage to show a file.
//
// See gtk.Image.NewFromFile for details.
//
// The function takes the following parameters:
//
//    - filename or NULL.
//
func (image *Image) SetFromFile(filename string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	if filename != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_image_set_from_file(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(filename)
}

// SetFromGIcon sets a GtkImage to show a GIcon.
//
// See gtk.Image.NewFromGIcon for details.
//
// The function takes the following parameters:
//
//    - icon: icon.
//
func (image *Image) SetFromGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.GIcon    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_image_set_from_gicon(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(icon)
}

// SetFromIconName sets a GtkImage to show a named icon.
//
// See gtk.Image.NewFromIconName for details.
//
// The function takes the following parameters:
//
//    - iconName: icon name or NULL.
//
func (image *Image) SetFromIconName(iconName string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_image_set_from_icon_name(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(iconName)
}

// SetFromPaintable sets a GtkImage to show a GdkPaintable.
//
// See gtk.Image.NewFromPaintable for details.
//
// The function takes the following parameters:
//
//    - paintable: GdkPaintable or NULL.
//
func (image *Image) SetFromPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.GtkImage     // out
	var _arg1 *C.GdkPaintable // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))
	}

	C.gtk_image_set_from_paintable(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(paintable)
}

// SetFromPixbuf sets a GtkImage to show a GdkPixbuf.
//
// See gtk.Image.NewFromPixbuf for details.
//
// Note: This is a helper for gtk.Image.SetFromPaintable(), and you can't get
// back the exact pixbuf once this is called, only a paintable.
//
// The function takes the following parameters:
//
//    - pixbuf: GdkPixbuf or NULL.
//
func (image *Image) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkImage  // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	}

	C.gtk_image_set_from_pixbuf(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(pixbuf)
}

// SetFromResource sets a GtkImage to show a resource.
//
// See gtk.Image.NewFromResource for details.
//
// The function takes the following parameters:
//
//    - resourcePath: resource path or NULL.
//
func (image *Image) SetFromResource(resourcePath string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	if resourcePath != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_image_set_from_resource(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(resourcePath)
}

// SetIconSize suggests an icon size to the theme for named icons.
//
// The function takes the following parameters:
//
//    - iconSize: new icon size.
//
func (image *Image) SetIconSize(iconSize IconSize) {
	var _arg0 *C.GtkImage   // out
	var _arg1 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_image_set_icon_size(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(iconSize)
}

// SetPixelSize sets the pixel size to use for named icons.
//
// If the pixel size is set to a value != -1, it is used instead of the icon
// size set by gtk.Image.SetFromIconName().
//
// The function takes the following parameters:
//
//    - pixelSize: new pixel size.
//
func (image *Image) SetPixelSize(pixelSize int) {
	var _arg0 *C.GtkImage // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = C.int(pixelSize)

	C.gtk_image_set_pixel_size(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(pixelSize)
}
