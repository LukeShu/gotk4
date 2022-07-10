// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
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

// GTypeCoverageLevel returns the GType for the type CoverageLevel.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCoverageLevel() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "CoverageLevel").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCoverageLevel)
	return gtype
}

// GTypeCoverage returns the GType for the type Coverage.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCoverage() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "Coverage").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCoverage)
	return gtype
}

// CoverageLevel: PangoCoverageLevel is used to indicate how well a font can
// represent a particular Unicode character for a particular script.
//
// Since 1.44, only PANGO_COVERAGE_NONE and PANGO_COVERAGE_EXACT will be
// returned.
type CoverageLevel C.gint

const (
	// CoverageNone: character is not representable with the font.
	CoverageNone CoverageLevel = iota
	// CoverageFallback: character is represented in a way that may be
	// comprehensible but is not the correct graphical form. For instance, a
	// Hangul character represented as a a sequence of Jamos, or a Latin
	// transliteration of a Cyrillic word.
	CoverageFallback
	// CoverageApproximate: character is represented as basically the correct
	// graphical form, but with a stylistic variant inappropriate for the
	// current script.
	CoverageApproximate
	// CoverageExact: character is represented as the correct graphical form.
	CoverageExact
)

func marshalCoverageLevel(p uintptr) (interface{}, error) {
	return CoverageLevel(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CoverageLevel.
func (c CoverageLevel) String() string {
	switch c {
	case CoverageNone:
		return "None"
	case CoverageFallback:
		return "Fallback"
	case CoverageApproximate:
		return "Approximate"
	case CoverageExact:
		return "Exact"
	default:
		return fmt.Sprintf("CoverageLevel(%d)", c)
	}
}

// Coverage structure is a map from Unicode characters to CoverageLevel values.
//
// It is often necessary in Pango to determine if a particular font can
// represent a particular character, and also how well it can represent that
// character. The Coverage is a data structure that is used to represent that
// information. It is an opaque structure with no public fields.
type Coverage struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Coverage)(nil)
)

func wrapCoverage(obj *coreglib.Object) *Coverage {
	return &Coverage{
		Object: obj,
	}
}

func marshalCoverage(p uintptr) (interface{}, error) {
	return wrapCoverage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCoverage: create a new PangoCoverage.
//
// The function returns the following values:
//
//    - coverage: newly allocated PangoCoverage, initialized to
//      PANGO_COVERAGE_NONE with a reference count of one, which should be freed
//      with pango_coverage_unref().
//
func NewCoverage() *Coverage {
	_info := girepository.MustFind("Pango", "Coverage")
	_gret := _info.InvokeClassMethod("new_Coverage", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _coverage *Coverage // out

	_coverage = wrapCoverage(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _coverage
}

// Copy an existing PangoCoverage.
//
// The function returns the following values:
//
//    - ret: newly allocated PangoCoverage, with a reference count of one, which
//      should be freed with pango_coverage_unref().
//
func (coverage *Coverage) Copy() *Coverage {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))

	_info := girepository.MustFind("Pango", "Coverage")
	_gret := _info.InvokeClassMethod("copy", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(coverage)

	var _ret *Coverage // out

	_ret = wrapCoverage(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _ret
}

// Max: set the coverage for each index in coverage to be the max (better) value
// of the current coverage for the index and the coverage for the corresponding
// index in other.
//
// Deprecated: This function does nothing.
//
// The function takes the following parameters:
//
//    - other PangoCoverage.
//
func (coverage *Coverage) Max(other *Coverage) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(other).Native()))

	_info := girepository.MustFind("Pango", "Coverage")
	_info.InvokeClassMethod("max", _args[:], nil)

	runtime.KeepAlive(coverage)
	runtime.KeepAlive(other)
}

// ToBytes: convert a PangoCoverage structure into a flat binary format.
//
// Deprecated: This returns NULL.
//
// The function returns the following values:
//
//    - bytes: location to store result (must be freed with g_free()).
//
func (coverage *Coverage) ToBytes() []byte {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(coverage).Native()))

	_info := girepository.MustFind("Pango", "Coverage")
	_info.InvokeClassMethod("to_bytes", _args[:], _outs[:])

	runtime.KeepAlive(coverage)

	var _bytes []byte // out

	defer C.free(unsafe.Pointer(*(**C.guchar)(unsafe.Pointer(&_outs[0]))))
	_bytes = make([]byte, *(*C.int)(unsafe.Pointer(&_outs[1])))
	copy(_bytes, unsafe.Slice((*byte)(unsafe.Pointer(*(**C.guchar)(unsafe.Pointer(&_outs[0])))), *(*C.int)(unsafe.Pointer(&_outs[1]))))

	return _bytes
}

// CoverageFromBytes: convert data generated from pango_coverage_to_bytes() back
// to a PangoCoverage.
//
// Deprecated: This returns NULL.
//
// The function takes the following parameters:
//
//    - bytes: binary data representing a PangoCoverage.
//
// The function returns the following values:
//
//    - coverage (optional): newly allocated PangoCoverage, or NULL if the data
//      was invalid.
//
func CoverageFromBytes(bytes []byte) *Coverage {
	var _args [2]girepository.Argument

	*(*C.int)(unsafe.Pointer(&_args[1])) = (C.int)(len(bytes))
	if len(bytes) > 0 {
		*(**C.guchar)(unsafe.Pointer(&_args[0])) = (*C.guchar)(unsafe.Pointer(&bytes[0]))
	}

	_info := girepository.MustFind("Pango", "from_bytes")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bytes)

	var _coverage *Coverage // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_coverage = wrapCoverage(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _coverage
}
