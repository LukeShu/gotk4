// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_coverage_level_get_type()), F: marshalCoverageLevel},
		{T: externglib.Type(C.pango_coverage_get_type()), F: marshalCoverage},
	})
}

// CoverageLevel: `PangoCoverageLevel` is used to indicate how well a font can
// represent a particular Unicode character for a particular script.
//
// Since 1.44, only PANGO_COVERAGE_NONE and PANGO_COVERAGE_EXACT will be
// returned.
type CoverageLevel int

const (
	// None: the character is not representable with the font.
	CoverageLevelNone CoverageLevel = iota
	// Fallback: the character is represented in a way that may be
	// comprehensible but is not the correct graphical form. For instance, a
	// Hangul character represented as a a sequence of Jamos, or a Latin
	// transliteration of a Cyrillic word.
	CoverageLevelFallback
	// Approximate: the character is represented as basically the correct
	// graphical form, but with a stylistic variant inappropriate for the
	// current script.
	CoverageLevelApproximate
	// Exact: the character is represented as the correct graphical form.
	CoverageLevelExact
)

func marshalCoverageLevel(p uintptr) (interface{}, error) {
	return CoverageLevel(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Coverage structure is a map from Unicode characters to CoverageLevel values.
//
// It is often necessary in Pango to determine if a particular font can
// represent a particular character, and also how well it can represent that
// character. The Coverage is a data structure that is used to represent that
// information. It is an opaque structure with no public fields.
type Coverage interface {
	gextras.Objector

	// Copy an existing `PangoCoverage`.
	Copy() *CoverageClass
	// Get: determine whether a particular index is covered by @coverage.
	Get(index_ int) CoverageLevel
	// Max: set the coverage for each index in @coverage to be the max (better)
	// value of the current coverage for the index and the coverage for the
	// corresponding index in @other.
	//
	// Deprecated: since version 1.44.
	Max(other Coverage)
	// Ref: increase the reference count on the `PangoCoverage` by one.
	ref() *CoverageClass
	// ToBytes: convert a `PangoCoverage` structure into a flat binary format.
	//
	// Deprecated: since version 1.44.
	ToBytes() []byte
	// Unref: decrease the reference count on the `PangoCoverage` by one.
	//
	// If the result is zero, free the coverage and all associated memory.
	unref()
}

// CoverageClass implements the Coverage interface.
type CoverageClass struct {
	*externglib.Object
}

var _ Coverage = (*CoverageClass)(nil)

func wrapCoverage(obj *externglib.Object) Coverage {
	return &CoverageClass{
		Object: obj,
	}
}

func marshalCoverage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCoverage(obj), nil
}

// NewCoverage: create a new `PangoCoverage`
func NewCoverage() *CoverageClass {
	var _cret *C.PangoCoverage // in

	_cret = C.pango_coverage_new()

	var _coverage *CoverageClass // out

	_coverage = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*CoverageClass)

	return _coverage
}

// Copy an existing `PangoCoverage`.
func (c *CoverageClass) Copy() *CoverageClass {
	var _arg0 *C.PangoCoverage // out
	var _cret *C.PangoCoverage // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer((&c).Native()))

	_cret = C.pango_coverage_copy(_arg0)

	var _ret *CoverageClass // out

	_ret = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*CoverageClass)

	return _ret
}

// Get: determine whether a particular index is covered by @coverage.
func (c *CoverageClass) Get(index_ int) CoverageLevel {
	var _arg0 *C.PangoCoverage     // out
	var _arg1 C.int                // out
	var _cret C.PangoCoverageLevel // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer((&c).Native()))
	_arg1 = C.int(index_)

	_cret = C.pango_coverage_get(_arg0, _arg1)

	var _coverageLevel CoverageLevel // out

	_coverageLevel = (CoverageLevel)(_cret)

	return _coverageLevel
}

// Max: set the coverage for each index in @coverage to be the max (better)
// value of the current coverage for the index and the coverage for the
// corresponding index in @other.
//
// Deprecated: since version 1.44.
func (c *CoverageClass) Max(other Coverage) {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.PangoCoverage // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer((&c).Native()))
	_arg1 = (*C.PangoCoverage)(unsafe.Pointer((&other).Native()))

	C.pango_coverage_max(_arg0, _arg1)
}

// Ref: increase the reference count on the `PangoCoverage` by one.
func (c *CoverageClass) ref() *CoverageClass {
	var _arg0 *C.PangoCoverage // out
	var _cret *C.PangoCoverage // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer((&c).Native()))

	_cret = C.pango_coverage_ref(_arg0)

	var _ret *CoverageClass // out

	_ret = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*CoverageClass)

	return _ret
}

// ToBytes: convert a `PangoCoverage` structure into a flat binary format.
//
// Deprecated: since version 1.44.
func (c *CoverageClass) ToBytes() []byte {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.guchar
	var _arg2 C.int // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer((&c).Native()))

	C.pango_coverage_to_bytes(_arg0, &_arg1, &_arg2)

	var _bytes []byte

	_bytes = unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_bytes, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})

	return _bytes
}

// Unref: decrease the reference count on the `PangoCoverage` by one.
//
// If the result is zero, free the coverage and all associated memory.
func (c *CoverageClass) unref() {
	var _arg0 *C.PangoCoverage // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer((&c).Native()))

	C.pango_coverage_unref(_arg0)
}
