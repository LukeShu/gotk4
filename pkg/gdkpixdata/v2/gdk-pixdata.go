// Code generated by girgen. DO NOT EDIT.

package gdkpixdata

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// PIXBUF_MAGIC_NUMBER: magic number for Pixdata structures.
const PIXBUF_MAGIC_NUMBER = 1197763408

// PIXDATA_HEADER_LENGTH: length of a Pixdata structure without the pixel_data
// pointer.
//
// Deprecated: since version 2.32.
const PIXDATA_HEADER_LENGTH = 24

// PixdataDumpType: enumeration which is used by gdk_pixdata_to_csource() to
// determine the form of C source to be generated. The three values
// GDK_PIXDATA_DUMP_PIXDATA_STREAM, GDK_PIXDATA_DUMP_PIXDATA_STRUCT and
// GDK_PIXDATA_DUMP_MACROS are mutually exclusive, as are GDK_PIXBUF_DUMP_GTYPES
// and GDK_PIXBUF_DUMP_CTYPES. The remaining elements are optional flags that
// can be freely added.
//
// Deprecated: since version 2.32.
type PixdataDumpType C.guint

const (
	// PixdataDumpPixdataStream: generate pixbuf data stream (a single string
	// containing a serialized Pixdata structure in network byte order).
	PixdataDumpPixdataStream PixdataDumpType = 0b0
	// PixdataDumpPixdataStruct: generate Pixdata structure (needs the Pixdata
	// structure definition from gdk-pixdata.h).
	PixdataDumpPixdataStruct PixdataDumpType = 0b1
	// PixdataDumpMacros: generate <function>*_ROWSTRIDE</function>,
	// <function>*_WIDTH</function>, <function>*_HEIGHT</function>,
	// <function>*_BYTES_PER_PIXEL</function> and
	// <function>*_RLE_PIXEL_DATA</function> or
	// <function>*_PIXEL_DATA</function> macro definitions for the image.
	PixdataDumpMacros PixdataDumpType = 0b10
	// PixdataDumpGTypes: generate GLib data types instead of standard C data
	// types.
	PixdataDumpGTypes PixdataDumpType = 0b0
	// PixdataDumpCtypes: generate standard C data types instead of GLib data
	// types.
	PixdataDumpCtypes PixdataDumpType = 0b100000000
	// PixdataDumpStatic: generate static symbols.
	PixdataDumpStatic PixdataDumpType = 0b1000000000
	// PixdataDumpConst: generate const symbols.
	PixdataDumpConst PixdataDumpType = 0b10000000000
	// PixdataDumpRLEDecoder: provide a <function>*_RUN_LENGTH_DECODE(image_buf,
	// rle_data, size, bpp)</function> macro definition to decode run-length
	// encoded image data.
	PixdataDumpRLEDecoder PixdataDumpType = 0b10000000000000000
)

// String returns the names in string for PixdataDumpType.
func (p PixdataDumpType) String() string {
	if p == 0 {
		return "PixdataDumpType(0)"
	}

	var builder strings.Builder
	builder.Grow(160)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PixdataDumpPixdataStream:
			builder.WriteString("PixdataStream|")
		case PixdataDumpPixdataStruct:
			builder.WriteString("PixdataStruct|")
		case PixdataDumpMacros:
			builder.WriteString("Macros|")
		case PixdataDumpCtypes:
			builder.WriteString("Ctypes|")
		case PixdataDumpStatic:
			builder.WriteString("Static|")
		case PixdataDumpConst:
			builder.WriteString("Const|")
		case PixdataDumpRLEDecoder:
			builder.WriteString("RLEDecoder|")
		default:
			builder.WriteString(fmt.Sprintf("PixdataDumpType(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if p contains other.
func (p PixdataDumpType) Has(other PixdataDumpType) bool {
	return (p & other) == other
}

// PixdataType: enumeration containing three sets of flags for a Pixdata struct:
// one for the used colorspace, one for the width of the samples and one for the
// encoding of the pixel data.
//
// Deprecated: since version 2.32.
type PixdataType C.guint

const (
	// PixdataColorTypeRGB: each pixel has red, green and blue samples.
	PixdataColorTypeRGB PixdataType = 0b1
	// PixdataColorTypeRGBA: each pixel has red, green and blue samples and an
	// alpha value.
	PixdataColorTypeRGBA PixdataType = 0b10
	// PixdataColorTypeMask: mask for the colortype flags of the enum.
	PixdataColorTypeMask PixdataType = 0b11111111
	// PixdataSampleWidth8: each sample has 8 bits.
	PixdataSampleWidth8 PixdataType = 0b10000000000000000
	// PixdataSampleWidthMask: mask for the sample width flags of the enum.
	PixdataSampleWidthMask PixdataType = 0b11110000000000000000
	// PixdataEncodingRaw: pixel data is in raw form.
	PixdataEncodingRaw PixdataType = 0b1000000000000000000000000
	// PixdataEncodingRLE: pixel data is run-length encoded. Runs may be up to
	// 127 bytes long; their length is stored in a single byte preceding the
	// pixel data for the run. If a run is constant, its length byte has the
	// high bit set and the pixel data consists of a single pixel which must be
	// repeated.
	PixdataEncodingRLE PixdataType = 0b10000000000000000000000000
	// PixdataEncodingMask: mask for the encoding flags of the enum.
	PixdataEncodingMask PixdataType = 0b1111000000000000000000000000
)

// String returns the names in string for PixdataType.
func (p PixdataType) String() string {
	if p == 0 {
		return "PixdataType(0)"
	}

	var builder strings.Builder
	builder.Grow(162)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PixdataColorTypeRGB:
			builder.WriteString("ColorTypeRGB|")
		case PixdataColorTypeRGBA:
			builder.WriteString("ColorTypeRGBA|")
		case PixdataColorTypeMask:
			builder.WriteString("ColorTypeMask|")
		case PixdataSampleWidth8:
			builder.WriteString("SampleWidth8|")
		case PixdataSampleWidthMask:
			builder.WriteString("SampleWidthMask|")
		case PixdataEncodingRaw:
			builder.WriteString("EncodingRaw|")
		case PixdataEncodingRLE:
			builder.WriteString("EncodingRLE|")
		case PixdataEncodingMask:
			builder.WriteString("EncodingMask|")
		default:
			builder.WriteString(fmt.Sprintf("PixdataType(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if p contains other.
func (p PixdataType) Has(other PixdataType) bool {
	return (p & other) == other
}

// PixbufFromPixdata converts a GdkPixdata to a GdkPixbuf.
//
// If copy_pixels is TRUE or if the pixel data is run-length-encoded, the pixel
// data is copied into newly-allocated memory; otherwise it is reused.
//
// Deprecated: Use GResource instead.
//
// The function takes the following parameters:
//
//    - pixdata to convert into a GdkPixbuf.
//    - copyPixels: whether to copy raw pixel data; run-length encoded pixel data
//      is always copied.
//
// The function returns the following values:
//
//    - pixbuf: new pixbuf.
//
func PixbufFromPixdata(pixdata *Pixdata, copyPixels bool) (*gdkpixbuf.Pixbuf, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(pixdata)))
	if copyPixels {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_gret := girepository.MustFind("GdkPixdata", "pixbuf_from_pixdata").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pixdata)
	runtime.KeepAlive(copyPixels)

	var _pixbuf *gdkpixbuf.Pixbuf // out
	var _goerr error              // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _pixbuf, _goerr
}

// Pixdata: pixel buffer suitable for serialization and streaming.
//
// Using GdkPixdata, images can be compiled into an application, making it
// unnecessary to refer to external image files at runtime.
//
// GdkPixbuf includes a utility named gdk-pixbuf-csource, which can be used to
// convert image files into GdkPixdata structures suitable for inclusion in C
// sources. To convert the GdkPixdata structures back into a GdkPixbuf, use
// gdk_pixbuf_from_pixdata().
//
// Deprecated: GdkPixdata should not be used any more. GResource should be used
// to save the original compressed images inside the program's binary.
//
// An instance of this type is always passed by reference.
type Pixdata struct {
	*pixdata
}

// pixdata is the struct that's finalized.
type pixdata struct {
	native unsafe.Pointer
}

// Deserialize deserializes (reconstruct) a Pixdata structure from a byte
// stream.
//
// The byte stream consists of a straightforward writeout of the GdkPixdata
// fields in network byte order, plus the pixel_data bytes the structure points
// to.
//
// The pixdata contents are reconstructed byte by byte and are checked for
// validity.
//
// This function may fail with GDK_PIXBUF_ERROR_CORRUPT_IMAGE or
// GDK_PIXBUF_ERROR_UNKNOWN_TYPE.
//
// Deprecated: Use GResource instead.
//
// The function takes the following parameters:
//
//    - stream of bytes containing a serialized Pixdata structure.
//
func (pixdata *Pixdata) Deserialize(stream []byte) error {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(pixdata)))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = (C.guint)(len(stream))
	if len(stream) > 0 {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(&stream[0]))
	}

	runtime.KeepAlive(pixdata)
	runtime.KeepAlive(stream)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Serialize serializes a Pixdata structure into a byte stream. The byte stream
// consists of a straightforward writeout of the Pixdata fields in network byte
// order, plus the pixel_data bytes the structure points to.
//
// Deprecated: Use #GResource instead.
//
// The function returns the following values:
//
//    - guint8s: a newly-allocated string containing the serialized Pixdata
//      structure.
//
func (pixdata *Pixdata) Serialize() []byte {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(pixdata)))

	_cret = *(**C.guint8)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(pixdata)

	var _guint8s []byte // out

	defer C.free(unsafe.Pointer(_cret))
	_guint8s = make([]byte, _outs[0])
	copy(_guint8s, unsafe.Slice((*byte)(unsafe.Pointer(_cret)), _outs[0]))

	return _guint8s
}
