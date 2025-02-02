// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
)

// #include <stdlib.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

// PixbufCalculateRowstride calculates the rowstride that an image created with
// those values would have.
//
// This function is useful for front-ends and backends that want to check image
// values without needing to create a GdkPixbuf.
//
// The function takes the following parameters:
//
//   - colorspace: color space for image.
//   - hasAlpha: whether the image should have transparency information.
//   - bitsPerSample: number of bits per color sample.
//   - width: width of image in pixels, must be > 0.
//   - height: height of image in pixels, must be > 0.
//
// The function returns the following values:
//
//   - gint: rowstride for the given values, or -1 in case of error.
//
func PixbufCalculateRowstride(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) int {
	var _arg1 C.GdkColorspace // out
	var _arg2 C.gboolean      // out
	var _arg3 C.int           // out
	var _arg4 C.int           // out
	var _arg5 C.int           // out
	var _cret C.gint          // in

	_arg1 = C.GdkColorspace(colorspace)
	if hasAlpha {
		_arg2 = C.TRUE
	}
	_arg3 = C.int(bitsPerSample)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	_cret = C.gdk_pixbuf_calculate_rowstride(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(colorspace)
	runtime.KeepAlive(hasAlpha)
	runtime.KeepAlive(bitsPerSample)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
