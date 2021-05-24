// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio"
	"github.com/diamondburned/gotk4/pkg/glib"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
//
// extern gboolean gotk4_PixbufSaveFunc(const gchar*, gsize, GError**, gpointer)
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		// Enums
		{T: externglib.Type(C.gdk_colorspace_get_type()), F: marshalColorspace},
		{T: externglib.Type(C.gdk_interp_type_get_type()), F: marshalInterpType},
		{T: externglib.Type(C.gdk_pixbuf_alpha_mode_get_type()), F: marshalPixbufAlphaMode},
		{T: externglib.Type(C.gdk_pixbuf_error_get_type()), F: marshalPixbufError},
		{T: externglib.Type(C.gdk_pixbuf_rotation_get_type()), F: marshalPixbufRotation},

		// Records
		{T: externglib.Type(C.gdk_pixbuf_format_get_type()), F: marshalPixbufFormat},
		// Skipped PixbufLoaderClass.
		// Skipped PixbufSimpleAnimClass.

		// Classes
		{T: externglib.Type(C.gdk_pixbuf_get_type()), F: marshalPixbuf},
		{T: externglib.Type(C.gdk_pixbuf_animation_get_type()), F: marshalPixbufAnimation},
		{T: externglib.Type(C.gdk_pixbuf_animation_iter_get_type()), F: marshalPixbufAnimationIter},
		{T: externglib.Type(C.gdk_pixbuf_loader_get_type()), F: marshalPixbufLoader},
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_get_type()), F: marshalPixbufSimpleAnim},
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_iter_get_type()), F: marshalPixbufSimpleAnimIter},
	})
}

// Colorspace: this enumeration defines the color spaces that are supported by
// the gdk-pixbuf library. Currently only RGB is supported.
type Colorspace int

const (
	// ColorspaceRGB indicates a red/green/blue additive color space.
	ColorspaceRGB Colorspace = 0
)

func marshalColorspace(p uintptr) (interface{}, error) {
	return Colorspace(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// InterpType: this enumeration describes the different interpolation modes that
// can be used with the scaling functions. @GDK_INTERP_NEAREST is the fastest
// scaling method, but has horrible quality when scaling down.
// @GDK_INTERP_BILINEAR is the best choice if you aren't sure what to choose, it
// has a good speed/quality balance.
//
// **Note**: Cubic filtering is missing from the list; hyperbolic interpolation
// is just as fast and results in higher quality.
type InterpType int

const (
	// InterpTypeNearest: nearest neighbor sampling; this is the fastest and
	// lowest quality mode. Quality is normally unacceptable when scaling down,
	// but may be OK when scaling up.
	InterpTypeNearest InterpType = 0
	// InterpTypeTiles: this is an accurate simulation of the PostScript image
	// operator without any interpolation enabled. Each pixel is rendered as a
	// tiny parallelogram of solid color, the edges of which are implemented
	// with antialiasing. It resembles nearest neighbor for enlargement, and
	// bilinear for reduction.
	InterpTypeTiles InterpType = 1
	// InterpTypeBilinear: best quality/speed balance; use this mode by default.
	// Bilinear interpolation. For enlargement, it is equivalent to
	// point-sampling the ideal bilinear-interpolated image. For reduction, it
	// is equivalent to laying down small tiles and integrating over the
	// coverage area.
	InterpTypeBilinear InterpType = 2
	// InterpTypeHyper: this is the slowest and highest quality reconstruction
	// function. It is derived from the hyperbolic filters in Wolberg's "Digital
	// Image Warping", and is formally defined as the hyperbolic-filter sampling
	// the ideal hyperbolic-filter interpolated image (the filter is designed to
	// be idempotent for 1:1 pixel mapping). **Deprecated**: this interpolation
	// filter is deprecated, as in reality it has a lower quality than the
	// @GDK_INTERP_BILINEAR filter (Since: 2.38)
	InterpTypeHyper InterpType = 3
)

func marshalInterpType(p uintptr) (interface{}, error) {
	return InterpType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PixbufAlphaMode: these values can be passed to
// gdk_pixbuf_xlib_render_to_drawable_alpha() to control how the alpha channel
// of an image should be handled. This function can create a bilevel clipping
// mask (black and white) and use it while painting the image. In the future,
// when the X Window System gets an alpha channel extension, it will be possible
// to do full alpha compositing onto arbitrary drawables. For now both cases
// fall back to a bilevel clipping mask.
type PixbufAlphaMode int

const (
	// PixbufAlphaModeBilevel: a bilevel clipping mask (black and white) will be
	// created and used to draw the image. Pixels below 0.5 opacity will be
	// considered fully transparent, and all others will be considered fully
	// opaque.
	PixbufAlphaModeBilevel PixbufAlphaMode = 0
	// PixbufAlphaModeFull: for now falls back to K_PIXBUF_ALPHA_BILEVEL. In the
	// future it will do full alpha compositing.
	PixbufAlphaModeFull PixbufAlphaMode = 1
)

func marshalPixbufAlphaMode(p uintptr) (interface{}, error) {
	return PixbufAlphaMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PixbufError: an error code in the K_PIXBUF_ERROR domain. Many gdk-pixbuf
// operations can cause errors in this domain, or in the FILE_ERROR domain.
type PixbufError int

const (
	// PixbufErrorCorruptImage: an image file was broken somehow.
	PixbufErrorCorruptImage PixbufError = 0
	// PixbufErrorInsufficientMemory: not enough memory.
	PixbufErrorInsufficientMemory PixbufError = 1
	// PixbufErrorBadOption: a bad option was passed to a pixbuf save module.
	PixbufErrorBadOption PixbufError = 2
	// PixbufErrorUnknownType: unknown image type.
	PixbufErrorUnknownType PixbufError = 3
	// PixbufErrorUnsupportedOperation: don't know how to perform the given
	// operation on the type of image at hand.
	PixbufErrorUnsupportedOperation PixbufError = 4
	// PixbufErrorFailed: generic failure code, something went wrong.
	PixbufErrorFailed PixbufError = 5
	// PixbufErrorIncompleteAnimation: only part of the animation was loaded.
	PixbufErrorIncompleteAnimation PixbufError = 6
)

func marshalPixbufError(p uintptr) (interface{}, error) {
	return PixbufError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PixbufRotation: the possible rotations which can be passed to
// gdk_pixbuf_rotate_simple(). To make them easier to use, their numerical
// values are the actual degrees.
type PixbufRotation int

const (
	// PixbufRotationNone: no rotation.
	PixbufRotationNone PixbufRotation = 0
	// PixbufRotationCounterclockwise: rotate by 90 degrees.
	PixbufRotationCounterclockwise PixbufRotation = 90
	// PixbufRotationUpsidedown: rotate by 180 degrees.
	PixbufRotationUpsidedown PixbufRotation = 180
	// PixbufRotationClockwise: rotate by 270 degrees.
	PixbufRotationClockwise PixbufRotation = 270
)

func marshalPixbufRotation(p uintptr) (interface{}, error) {
	return PixbufRotation(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PixbufSaveFunc specifies the type of the function passed to
// gdk_pixbuf_save_to_callback(). It is called once for each block of bytes that
// is "written" by gdk_pixbuf_save_to_callback(). If successful it should return
// true. If an error occurs it should set @error and return false, in which case
// gdk_pixbuf_save_to_callback() will fail with the same error.
type PixbufSaveFunc func(buf []uint8) (err *glib.Error, ok bool)

//export gotk4_PixbufSaveFunc
func gotk4_PixbufSaveFunc(arg0 *C.gchar, arg1 C.gsize, arg2 **C.GError, arg3 C.gpointer) C.gboolean {
	v := box.Get(box.Callback, uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var buf []uint8
	{
		buf = make([]uint8, arg1)
		for i := 0; i < uintptr(arg1); i++ {
			src := (C.guint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))
			buf[i] = uint8(src)
		}
	}

	err, ok := v.(PixbufSaveFunc)(buf)
}

func PixbufErrorQuark() glib.Quark {
	ret := C.gdk_pixbuf_error_quark()

	var ret0 glib.Quark
	{
		var tmp uint32
		tmp = uint32(ret)
		ret0 = glib.Quark(tmp)
	}

	return ret0
}

type PixbufFormat struct {
	native *C.GdkPixbufFormat
}

// WrapPixbufFormat wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPixbufFormat(ptr unsafe.Pointer) *PixbufFormat {
	p := (*C.GdkPixbufFormat)(ptr)
	v := PixbufFormat{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*PixbufFormat).free)

	return &v
}

func marshalPixbufFormat(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPixbufFormat(unsafe.Pointer(b))
}

func (p *PixbufFormat) free() {
	C.free(unsafe.Pointer(p.native))
}

// Native returns the pointer to *C.GdkPixbufFormat. The caller is expected to
// cast.
func (p *PixbufFormat) Native() unsafe.Pointer {
	return unsafe.Pointer(p.native)
}

// Pixbuf: this is the main structure in the gdk-pixbuf library. It is used to
// represent images. It contains information about the image's pixel data, its
// color space, bits per sample, width and height, and the rowstride (the number
// of bytes between the start of one row and the start of the next).
type Pixbuf interface {
	gextras.Objector

	// AddAlpha takes an existing pixbuf and adds an alpha channel to it. If the
	// existing pixbuf already had an alpha channel, the channel values are
	// copied from the original; otherwise, the alpha channel is initialized to
	// 255 (full opacity).
	//
	// If @substitute_color is true, then the color specified by (@r, @g, @b)
	// will be assigned zero opacity. That is, if you pass (255, 255, 255) for
	// the substitute color, all white pixels will become fully transparent.
	AddAlpha(substituteColor bool, r uint8, g uint8, b uint8) Pixbuf
	// ApplyEmbeddedOrientation takes an existing pixbuf and checks for the
	// presence of an associated "orientation" option, which may be provided by
	// the jpeg loader (which reads the exif orientation tag) or the tiff loader
	// (which reads the tiff orientation tag, and compensates it for the partial
	// transforms performed by libtiff). If an orientation option/tag is
	// present, the appropriate transform will be performed so that the pixbuf
	// is oriented correctly.
	ApplyEmbeddedOrientation() Pixbuf
	// Composite creates a transformation of the source image @src by scaling by
	// @scale_x and @scale_y then translating by @offset_x and @offset_y. This
	// gives an image in the coordinates of the destination pixbuf. The
	// rectangle (@dest_x, @dest_y, @dest_width, @dest_height) is then alpha
	// blended onto the corresponding rectangle of the original destination
	// image.
	//
	// When the destination rectangle contains parts not in the source image,
	// the data at the edges of the source image is replicated to infinity.
	//
	// ! (composite.png)
	Composite(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int)
	// CompositeColor creates a transformation of the source image @src by
	// scaling by @scale_x and @scale_y then translating by @offset_x and
	// @offset_y, then alpha blends the rectangle (@dest_x ,@dest_y,
	// @dest_width, @dest_height) of the resulting image with a checkboard of
	// the colors @color1 and @color2 and renders it onto the destination image.
	//
	// If the source image has no alpha channel, and @overall_alpha is 255, a
	// fast path is used which omits the alpha blending and just performs the
	// scaling.
	//
	// See gdk_pixbuf_composite_color_simple() for a simpler variant of this
	// function suitable for many tasks.
	CompositeColor(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int, checkX int, checkY int, checkSize int, color1 uint32, color2 uint32)
	// CompositeColorSimple creates a new Pixbuf by scaling @src to @dest_width
	// x @dest_height and alpha blending the result with a checkboard of colors
	// @color1 and @color2.
	CompositeColorSimple(destWidth int, destHeight int, interpType InterpType, overallAlpha int, checkSize int, color1 uint32, color2 uint32) Pixbuf
	// Copy creates a new Pixbuf with a copy of the information in the specified
	// @pixbuf. Note that this does not copy the options set on the original
	// Pixbuf, use gdk_pixbuf_copy_options() for this.
	Copy() Pixbuf
	// CopyArea copies a rectangular area from @src_pixbuf to @dest_pixbuf.
	// Conversion of pixbuf formats is done automatically.
	//
	// If the source rectangle overlaps the destination rectangle on the same
	// pixbuf, it will be overwritten during the copy operation. Therefore, you
	// can not use this function to scroll a pixbuf.
	CopyArea(srcX int, srcY int, width int, height int, destPixbuf Pixbuf, destX int, destY int)
	// CopyOptions: copy the key/value pair options attached to a Pixbuf to
	// another. This is useful to keep original metadata after having
	// manipulated a file. However be careful to remove metadata which you've
	// already applied, such as the "orientation" option after rotating the
	// image.
	CopyOptions(destPixbuf Pixbuf) bool
	// Fill clears a pixbuf to the given RGBA value, converting the RGBA value
	// into the pixbuf's pixel format. The alpha will be ignored if the pixbuf
	// doesn't have an alpha channel.
	Fill(pixel uint32)
	// Flip flips a pixbuf horizontally or vertically and returns the result in
	// a new pixbuf.
	Flip(horizontal bool) Pixbuf
	// BitsPerSample queries the number of bits per color sample in a pixbuf.
	BitsPerSample() int
	// ByteLength returns the length of the pixel data, in bytes.
	ByteLength() uint
	// Colorspace queries the color space of a pixbuf.
	Colorspace() Colorspace
	// HasAlpha queries whether a pixbuf has an alpha channel (opacity
	// information).
	HasAlpha() bool
	// Height queries the height of a pixbuf.
	Height() int
	// NChannels queries the number of channels of a pixbuf.
	NChannels() int
	// Option looks up @key in the list of options that may have been attached
	// to the @pixbuf when it was loaded, or that may have been attached by
	// another function using gdk_pixbuf_set_option().
	//
	// For instance, the ANI loader provides "Title" and "Artist" options. The
	// ICO, XBM, and XPM loaders provide "x_hot" and "y_hot" hot-spot options
	// for cursor definitions. The PNG loader provides the tEXt ancillary chunk
	// key/value pairs as options. Since 2.12, the TIFF and JPEG loaders return
	// an "orientation" option string that corresponds to the embedded TIFF/Exif
	// orientation tag (if present). Since 2.32, the TIFF loader sets the
	// "multipage" option string to "yes" when a multi-page TIFF is loaded.
	// Since 2.32 the JPEG and PNG loaders set "x-dpi" and "y-dpi" if the file
	// contains image density information in dots per inch. Since 2.36.6, the
	// JPEG loader sets the "comment" option with the comment EXIF tag.
	Option(key string) string
	// Options returns a Table with a list of all the options that may have been
	// attached to the @pixbuf when it was loaded, or that may have been
	// attached by another function using gdk_pixbuf_set_option().
	//
	// See gdk_pixbuf_get_option() for more details.
	Options() *glib.HashTable
	// Pixels queries a pointer to the pixel data of a pixbuf.
	Pixels() []uint8
	// PixelsWithLength queries a pointer to the pixel data of a pixbuf.
	PixelsWithLength() (length uint, guint8s []uint8)
	// Rowstride queries the rowstride of a pixbuf, which is the number of bytes
	// between the start of a row and the start of the next row.
	Rowstride() int
	// Width queries the width of a pixbuf.
	Width() int
	// NewSubpixbuf creates a new pixbuf which represents a sub-region of
	// @src_pixbuf. The new pixbuf shares its pixels with the original pixbuf,
	// so writing to one affects both. The new pixbuf holds a reference to
	// @src_pixbuf, so @src_pixbuf will not be finalized until the new pixbuf is
	// finalized.
	//
	// Note that if @src_pixbuf is read-only, this function will force it to be
	// mutable.
	NewSubpixbuf(srcX int, srcY int, width int, height int) Pixbuf
	// ReadPixelBytes provides a #GBytes buffer containing the raw pixel data;
	// the data must not be modified. This function allows skipping the implicit
	// copy that must be made if gdk_pixbuf_get_pixels() is called on a
	// read-only pixbuf.
	ReadPixelBytes() *glib.Bytes
	// ReadPixels provides a read-only pointer to the raw pixel data; must not
	// be modified. This function allows skipping the implicit copy that must be
	// made if gdk_pixbuf_get_pixels() is called on a read-only pixbuf.
	ReadPixels() uint8
	// Ref adds a reference to a pixbuf.
	Ref() Pixbuf
	// RemoveOption: remove the key/value pair option attached to a Pixbuf.
	RemoveOption(key string) bool
	// RotateSimple rotates a pixbuf by a multiple of 90 degrees, and returns
	// the result in a new pixbuf.
	//
	// If @angle is 0, a copy of @src is returned, avoiding any rotation.
	RotateSimple(angle PixbufRotation) Pixbuf
	// SaturateAndPixelate modifies saturation and optionally pixelates @src,
	// placing the result in @dest. @src and @dest may be the same pixbuf with
	// no ill effects. If @saturation is 1.0 then saturation is not changed. If
	// it's less than 1.0, saturation is reduced (the image turns toward
	// grayscale); if greater than 1.0, saturation is increased (the image gets
	// more vivid colors). If @pixelate is true, then pixels are faded in a
	// checkerboard pattern to create a pixelated image. @src and @dest must
	// have the same image format, size, and rowstride.
	SaturateAndPixelate(dest Pixbuf, saturation float32, pixelate bool)
	// SaveToBufferv saves pixbuf to a new buffer in format @type, which is
	// currently "jpeg", "tiff", "png", "ico" or "bmp". See
	// gdk_pixbuf_save_to_buffer() for more details.
	SaveToBufferv(_type string, optionKeys []string, optionValues []string) (buffer []uint8, bufferSize uint, ok bool)
	// SaveToCallbackv saves pixbuf to a callback in format @type, which is
	// currently "jpeg", "png", "tiff", "ico" or "bmp". If @error is set, false
	// will be returned. See gdk_pixbuf_save_to_callback () for more details.
	SaveToCallbackv(saveFunc PixbufSaveFunc, _type string, optionKeys []string, optionValues []string) bool
	// SaveToStreamv saves @pixbuf to an output stream.
	//
	// Supported file formats are currently "jpeg", "tiff", "png", "ico" or
	// "bmp". See gdk_pixbuf_save_to_stream() for more details.
	SaveToStreamv(stream gio.OutputStream, _type string, optionKeys []string, optionValues []string, cancellable gio.Cancellable) bool
	// SaveToStreamvAsync saves @pixbuf to an output stream asynchronously.
	//
	// For more details see gdk_pixbuf_save_to_streamv(), which is the
	// synchronous version of this function.
	//
	// When the operation is finished, @callback will be called in the main
	// thread. You can then call gdk_pixbuf_save_to_stream_finish() to get the
	// result of the operation.
	SaveToStreamvAsync(stream gio.OutputStream, _type string, optionKeys []string, optionValues []string, cancellable gio.Cancellable, callback gio.AsyncReadyCallback)
	// Savev saves pixbuf to a file in @type, which is currently "jpeg", "png",
	// "tiff", "ico" or "bmp". If @error is set, false will be returned. See
	// gdk_pixbuf_save () for more details.
	Savev(filename string, _type string, optionKeys []string, optionValues []string) bool
	// Scale creates a transformation of the source image @src by scaling by
	// @scale_x and @scale_y then translating by @offset_x and @offset_y, then
	// renders the rectangle (@dest_x, @dest_y, @dest_width, @dest_height) of
	// the resulting image onto the destination image replacing the previous
	// contents.
	//
	// Try to use gdk_pixbuf_scale_simple() first, this function is the
	// industrial-strength power tool you can fall back to if
	// gdk_pixbuf_scale_simple() isn't powerful enough.
	//
	// If the source rectangle overlaps the destination rectangle on the same
	// pixbuf, it will be overwritten during the scaling which results in
	// rendering artifacts.
	Scale(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType)
	// ScaleSimple: create a new Pixbuf containing a copy of @src scaled to
	// @dest_width x @dest_height. Leaves @src unaffected. @interp_type should
	// be K_INTERP_NEAREST if you want maximum speed (but when scaling down
	// K_INTERP_NEAREST is usually unusably ugly). The default @interp_type
	// should be K_INTERP_BILINEAR which offers reasonable quality and speed.
	//
	// You can scale a sub-portion of @src by creating a sub-pixbuf pointing
	// into @src; see gdk_pixbuf_new_subpixbuf().
	//
	// If @dest_width and @dest_height are equal to the @src width and height, a
	// copy of @src is returned, avoiding any scaling.
	//
	// For more complicated scaling/alpha blending see gdk_pixbuf_scale() and
	// gdk_pixbuf_composite().
	ScaleSimple(destWidth int, destHeight int, interpType InterpType) Pixbuf
	// SetOption attaches a key/value pair as an option to a Pixbuf. If @key
	// already exists in the list of options attached to @pixbuf, the new value
	// is ignored and false is returned.
	SetOption(key string, value string) bool
	// Unref removes a reference from a pixbuf.
	Unref()
}

type pixbuf struct {
	*externglib.Object
}

// WrapPixbuf wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbuf(obj *externglib.Object) Pixbuf {
	return pixbuf{*externglib.Object{obj}}
}

func marshalPixbuf(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbuf(obj), nil
}

func NewPixbuf(colorspace Colorspace, hasAlpha bool, bitsPerSample int, width int, height int) Pixbuf

func NewPixbufFromBytes(data *glib.Bytes, colorspace Colorspace, hasAlpha bool, bitsPerSample int, width int, height int, rowstride int) Pixbuf

func NewPixbufFromData(data []uint8, colorspace Colorspace, hasAlpha bool, bitsPerSample int, width int, height int, rowstride int) Pixbuf

func NewPixbufFromFile(filename string) Pixbuf

func NewPixbufFromFileAtScale(filename string, width int, height int, preserveAspectRatio bool) Pixbuf

func NewPixbufFromFileAtSize(filename string, width int, height int) Pixbuf

func NewPixbufFromInline(dataLength int, data []uint8, copyPixels bool) Pixbuf

func NewPixbufFromResource(resourcePath string) Pixbuf

func NewPixbufFromResourceAtScale(resourcePath string, width int, height int, preserveAspectRatio bool) Pixbuf

func NewPixbufFromStream(stream gio.InputStream, cancellable gio.Cancellable) Pixbuf

func NewPixbufFromStreamAtScale(stream gio.InputStream, width int, height int, preserveAspectRatio bool, cancellable gio.Cancellable) Pixbuf

func NewPixbufFromStreamFinish(asyncResult gio.AsyncResult) Pixbuf

func NewPixbufFromXpmData(data []string) Pixbuf

func (p pixbuf) AddAlpha(substituteColor bool, r uint8, g uint8, b uint8) Pixbuf

func (p pixbuf) ApplyEmbeddedOrientation() Pixbuf

func (p pixbuf) Composite(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int)

func (p pixbuf) CompositeColor(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int, checkX int, checkY int, checkSize int, color1 uint32, color2 uint32)

func (p pixbuf) CompositeColorSimple(destWidth int, destHeight int, interpType InterpType, overallAlpha int, checkSize int, color1 uint32, color2 uint32) Pixbuf

func (p pixbuf) Copy() Pixbuf

func (p pixbuf) CopyArea(srcX int, srcY int, width int, height int, destPixbuf Pixbuf, destX int, destY int)

func (p pixbuf) CopyOptions(destPixbuf Pixbuf) bool

func (p pixbuf) Fill(pixel uint32)

func (p pixbuf) Flip(horizontal bool) Pixbuf

func (p pixbuf) BitsPerSample() int

func (p pixbuf) ByteLength() uint

func (p pixbuf) Colorspace() Colorspace

func (p pixbuf) HasAlpha() bool

func (p pixbuf) Height() int

func (p pixbuf) NChannels() int

func (p pixbuf) Option(key string) string

func (p pixbuf) Options() *glib.HashTable

func (p pixbuf) Pixels() []uint8

func (p pixbuf) PixelsWithLength() (length uint, guint8s []uint8)

func (p pixbuf) Rowstride() int

func (p pixbuf) Width() int

func (p pixbuf) NewSubpixbuf(srcX int, srcY int, width int, height int) Pixbuf

func (p pixbuf) ReadPixelBytes() *glib.Bytes

func (p pixbuf) ReadPixels() uint8

func (p pixbuf) Ref() Pixbuf

func (p pixbuf) RemoveOption(key string) bool

func (p pixbuf) RotateSimple(angle PixbufRotation) Pixbuf

func (p pixbuf) SaturateAndPixelate(dest Pixbuf, saturation float32, pixelate bool)

func (p pixbuf) SaveToBufferv(_type string, optionKeys []string, optionValues []string) (buffer []uint8, bufferSize uint, ok bool)

func (p pixbuf) SaveToCallbackv(saveFunc PixbufSaveFunc, _type string, optionKeys []string, optionValues []string) bool

func (p pixbuf) SaveToStreamv(stream gio.OutputStream, _type string, optionKeys []string, optionValues []string, cancellable gio.Cancellable) bool

func (p pixbuf) SaveToStreamvAsync(stream gio.OutputStream, _type string, optionKeys []string, optionValues []string, cancellable gio.Cancellable, callback gio.AsyncReadyCallback)

func (p pixbuf) Savev(filename string, _type string, optionKeys []string, optionValues []string) bool

func (p pixbuf) Scale(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType)

func (p pixbuf) ScaleSimple(destWidth int, destHeight int, interpType InterpType) Pixbuf

func (p pixbuf) SetOption(key string, value string) bool

func (p pixbuf) Unref()

// PixbufAnimation: an opaque struct representing an animation.
type PixbufAnimation interface {
	gextras.Objector

	// Height queries the height of the bounding box of a pixbuf animation.
	Height() int
	// Iter: get an iterator for displaying an animation. The iterator provides
	// the frames that should be displayed at a given time. It should be freed
	// after use with g_object_unref().
	//
	// @start_time would normally come from g_get_current_time(), and marks the
	// beginning of animation playback. After creating an iterator, you should
	// immediately display the pixbuf returned by
	// gdk_pixbuf_animation_iter_get_pixbuf(). Then, you should install a
	// timeout (with g_timeout_add()) or by some other mechanism ensure that
	// you'll update the image after gdk_pixbuf_animation_iter_get_delay_time()
	// milliseconds. Each time the image is updated, you should reinstall the
	// timeout with the new, possibly-changed delay time.
	//
	// As a shortcut, if @start_time is nil, the result of g_get_current_time()
	// will be used automatically.
	//
	// To update the image (i.e. possibly change the result of
	// gdk_pixbuf_animation_iter_get_pixbuf() to a new frame of the animation),
	// call gdk_pixbuf_animation_iter_advance().
	//
	// If you're using PixbufLoader, in addition to updating the image after the
	// delay time, you should also update it whenever you receive the
	// area_updated signal and
	// gdk_pixbuf_animation_iter_on_currently_loading_frame() returns true. In
	// this case, the frame currently being fed into the loader has received new
	// data, so needs to be refreshed. The delay time for a frame may also be
	// modified after an area_updated signal, for example if the delay time for
	// a frame is encoded in the data after the frame itself. So your timeout
	// should be reinstalled after any area_updated signal.
	//
	// A delay time of -1 is possible, indicating "infinite."
	Iter(startTime *glib.TimeVal) PixbufAnimationIter
	// StaticImage: if an animation is really just a plain image (has only one
	// frame), this function returns that image. If the animation is an
	// animation, this function returns a reasonable thing to display as a
	// static unanimated image, which might be the first frame, or something
	// more sophisticated. If an animation hasn't loaded any frames yet, this
	// function will return nil.
	StaticImage() Pixbuf
	// Width queries the width of the bounding box of a pixbuf animation.
	Width() int
	// IsStaticImage: if you load a file with
	// gdk_pixbuf_animation_new_from_file() and it turns out to be a plain,
	// unanimated image, then this function will return true. Use
	// gdk_pixbuf_animation_get_static_image() to retrieve the image.
	IsStaticImage() bool
	// Ref adds a reference to an animation.
	Ref() PixbufAnimation
	// Unref removes a reference from an animation.
	Unref()
}

type pixbufAnimation struct {
	*externglib.Object
}

// WrapPixbufAnimation wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufAnimation(obj *externglib.Object) PixbufAnimation {
	return pixbufAnimation{*externglib.Object{obj}}
}

func marshalPixbufAnimation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufAnimation(obj), nil
}

func NewPixbufAnimationFromFile(filename string) PixbufAnimation

func NewPixbufAnimationFromResource(resourcePath string) PixbufAnimation

func NewPixbufAnimationFromStream(stream gio.InputStream, cancellable gio.Cancellable) PixbufAnimation

func NewPixbufAnimationFromStreamFinish(asyncResult gio.AsyncResult) PixbufAnimation

func (p pixbufAnimation) Height() int

func (p pixbufAnimation) Iter(startTime *glib.TimeVal) PixbufAnimationIter

func (p pixbufAnimation) StaticImage() Pixbuf

func (p pixbufAnimation) Width() int

func (p pixbufAnimation) IsStaticImage() bool

func (p pixbufAnimation) Ref() PixbufAnimation

func (p pixbufAnimation) Unref()

// PixbufAnimationIter: an opaque struct representing an iterator which points
// to a certain position in an animation.
type PixbufAnimationIter interface {
	gextras.Objector

	// Advance: possibly advances an animation to a new frame. Chooses the frame
	// based on the start time passed to gdk_pixbuf_animation_get_iter().
	//
	// @current_time would normally come from g_get_current_time(), and must be
	// greater than or equal to the time passed to
	// gdk_pixbuf_animation_get_iter(), and must increase or remain unchanged
	// each time gdk_pixbuf_animation_iter_get_pixbuf() is called. That is, you
	// can't go backward in time; animations only play forward.
	//
	// As a shortcut, pass nil for the current time and g_get_current_time()
	// will be invoked on your behalf. So you only need to explicitly pass
	// @current_time if you're doing something odd like playing the animation at
	// double speed.
	//
	// If this function returns false, there's no need to update the animation
	// display, assuming the display had been rendered prior to advancing; if
	// true, you need to call gdk_pixbuf_animation_iter_get_pixbuf() and update
	// the display with the new pixbuf.
	Advance(currentTime *glib.TimeVal) bool
	// DelayTime gets the number of milliseconds the current pixbuf should be
	// displayed, or -1 if the current pixbuf should be displayed forever.
	// g_timeout_add() conveniently takes a timeout in milliseconds, so you can
	// use a timeout to schedule the next update.
	//
	// Note that some formats, like GIF, might clamp the timeout values in the
	// image file to avoid updates that are just too quick. The minimum timeout
	// for GIF images is currently 20 milliseconds.
	DelayTime() int
	// Pixbuf gets the current pixbuf which should be displayed; the pixbuf
	// might not be the same size as the animation itself
	// (gdk_pixbuf_animation_get_width(), gdk_pixbuf_animation_get_height()).
	// This pixbuf should be displayed for
	// gdk_pixbuf_animation_iter_get_delay_time() milliseconds. The caller of
	// this function does not own a reference to the returned pixbuf; the
	// returned pixbuf will become invalid when the iterator advances to the
	// next frame, which may happen anytime you call
	// gdk_pixbuf_animation_iter_advance(). Copy the pixbuf to keep it (don't
	// just add a reference), as it may get recycled as you advance the
	// iterator.
	Pixbuf() Pixbuf
	// OnCurrentlyLoadingFrame: used to determine how to respond to the
	// area_updated signal on PixbufLoader when loading an animation.
	// area_updated is emitted for an area of the frame currently streaming in
	// to the loader. So if you're on the currently loading frame, you need to
	// redraw the screen for the updated area.
	OnCurrentlyLoadingFrame() bool
}

type pixbufAnimationIter struct {
	*externglib.Object
}

// WrapPixbufAnimationIter wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufAnimationIter(obj *externglib.Object) PixbufAnimationIter {
	return pixbufAnimationIter{*externglib.Object{obj}}
}

func marshalPixbufAnimationIter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufAnimationIter(obj), nil
}

func (p pixbufAnimationIter) Advance(currentTime *glib.TimeVal) bool

func (p pixbufAnimationIter) DelayTime() int

func (p pixbufAnimationIter) Pixbuf() Pixbuf

func (p pixbufAnimationIter) OnCurrentlyLoadingFrame() bool

// PixbufLoader: the GdkPixbufLoader struct contains only private fields.
type PixbufLoader interface {
	gextras.Objector

	// Close informs a pixbuf loader that no further writes with
	// gdk_pixbuf_loader_write() will occur, so that it can free its internal
	// loading structures. Also, tries to parse any data that hasn't yet been
	// parsed; if the remaining data is partial or corrupt, an error will be
	// returned. If false is returned, @error will be set to an error from the
	// K_PIXBUF_ERROR or FILE_ERROR domains. If you're just cancelling a load
	// rather than expecting it to be finished, passing nil for @error to ignore
	// it is reasonable.
	//
	// Remember that this does not unref the loader, so if you plan not to use
	// it anymore, please g_object_unref() it.
	Close() bool
	// Animation queries the PixbufAnimation that a pixbuf loader is currently
	// creating. In general it only makes sense to call this function after the
	// "area-prepared" signal has been emitted by the loader. If the loader
	// doesn't have enough bytes yet (hasn't emitted the "area-prepared" signal)
	// this function will return nil.
	Animation() PixbufAnimation
	// Format obtains the available information about the format of the
	// currently loading image file.
	Format() *PixbufFormat
	// Pixbuf queries the Pixbuf that a pixbuf loader is currently creating. In
	// general it only makes sense to call this function after the
	// "area-prepared" signal has been emitted by the loader; this means that
	// enough data has been read to know the size of the image that will be
	// allocated. If the loader has not received enough data via
	// gdk_pixbuf_loader_write(), then this function returns nil. The returned
	// pixbuf will be the same in all future calls to the loader, so simply
	// calling g_object_ref() should be sufficient to continue using it.
	// Additionally, if the loader is an animation, it will return the "static
	// image" of the animation (see gdk_pixbuf_animation_get_static_image()).
	Pixbuf() Pixbuf
	// SetSize causes the image to be scaled while it is loaded. The desired
	// image size can be determined relative to the original size of the image
	// by calling gdk_pixbuf_loader_set_size() from a signal handler for the
	// ::size-prepared signal.
	//
	// Attempts to set the desired image size are ignored after the emission of
	// the ::size-prepared signal.
	SetSize(width int, height int)
	// Write: this will cause a pixbuf loader to parse the next @count bytes of
	// an image. It will return true if the data was loaded successfully, and
	// false if an error occurred. In the latter case, the loader will be
	// closed, and will not accept further writes. If false is returned, @error
	// will be set to an error from the K_PIXBUF_ERROR or FILE_ERROR domains.
	Write(buf []uint8) bool
	// WriteBytes: this will cause a pixbuf loader to parse a buffer inside a
	// #GBytes for an image. It will return true if the data was loaded
	// successfully, and false if an error occurred. In the latter case, the
	// loader will be closed, and will not accept further writes. If false is
	// returned, @error will be set to an error from the K_PIXBUF_ERROR or
	// FILE_ERROR domains.
	//
	// See also: gdk_pixbuf_loader_write()
	WriteBytes(buffer *glib.Bytes) bool
}

type pixbufLoader struct {
	*externglib.Object
}

// WrapPixbufLoader wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufLoader(obj *externglib.Object) PixbufLoader {
	return pixbufLoader{*externglib.Object{obj}}
}

func marshalPixbufLoader(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufLoader(obj), nil
}

func NewPixbufLoader() PixbufLoader

func NewPixbufLoaderWithMIMEType(mimeType string) PixbufLoader

func NewPixbufLoaderWithType(imageType string) PixbufLoader

func (p pixbufLoader) Close() bool

func (p pixbufLoader) Animation() PixbufAnimation

func (p pixbufLoader) Format() *PixbufFormat

func (p pixbufLoader) Pixbuf() Pixbuf

func (p pixbufLoader) SetSize(width int, height int)

func (p pixbufLoader) Write(buf []uint8) bool

func (p pixbufLoader) WriteBytes(buffer *glib.Bytes) bool

// PixbufSimpleAnim: an opaque struct representing a simple animation.
type PixbufSimpleAnim interface {
	PixbufAnimation

	// AddFrame adds a new frame to @animation. The @pixbuf must have the
	// dimensions specified when the animation was constructed.
	AddFrame(pixbuf Pixbuf)
	// Loop gets whether @animation should loop indefinitely when it reaches the
	// end.
	Loop() bool
	// SetLoop sets whether @animation should loop indefinitely when it reaches
	// the end.
	SetLoop(loop bool)
}

type pixbufSimpleAnim struct {
	pixbufAnimation
}

// WrapPixbufSimpleAnim wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufSimpleAnim(obj *externglib.Object) PixbufSimpleAnim {
	return pixbufSimpleAnim{pixbufAnimation{*externglib.Object{obj}}}
}

func marshalPixbufSimpleAnim(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufSimpleAnim(obj), nil
}

func NewPixbufSimpleAnim(width int, height int, rate float32) PixbufSimpleAnim

func (p pixbufSimpleAnim) AddFrame(pixbuf Pixbuf)

func (p pixbufSimpleAnim) Loop() bool

func (p pixbufSimpleAnim) SetLoop(loop bool)

type PixbufSimpleAnimIter interface {
	PixbufAnimationIter
}

type pixbufSimpleAnimIter struct {
	pixbufAnimationIter
}

// WrapPixbufSimpleAnimIter wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufSimpleAnimIter(obj *externglib.Object) PixbufSimpleAnimIter {
	return pixbufSimpleAnimIter{pixbufAnimationIter{*externglib.Object{obj}}}
}

func marshalPixbufSimpleAnimIter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufSimpleAnimIter(obj), nil
}
