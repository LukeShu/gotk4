// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_vec3_get_type()), F: marshalVec3},
	})
}

// Vec3: structure capable of holding a vector with three dimensions: x, y, and
// z.
//
// The contents of the #graphene_vec3_t structure are private and should never
// be accessed directly.
type Vec3 struct {
	native C.graphene_vec3_t
}

func marshalVec3(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Vec3)(unsafe.Pointer(b)), nil
}

// NewVec3Alloc constructs a struct Vec3.
func NewVec3Alloc() *Vec3 {
	var _cret *C.graphene_vec3_t // in

	_cret = C.graphene_vec3_alloc()

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_vec3, func(v *Vec3) {
		C.graphene_vec3_free((*C.graphene_vec3_t)(unsafe.Pointer(v)))
	})

	return _vec3
}

// Native returns the underlying C source pointer.
func (v *Vec3) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}

// Add adds each component of the two given vectors.
func (a *Vec3) Add(b *Vec3) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	C.graphene_vec3_add(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Cross computes the cross product of the two given vectors.
func (a *Vec3) Cross(b *Vec3) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	C.graphene_vec3_cross(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Divide divides each component of the first operand @a by the corresponding
// component of the second operand @b, and places the results into the vector
// @res.
func (a *Vec3) Divide(b *Vec3) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	C.graphene_vec3_divide(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Dot computes the dot product of the two given vectors.
func (a *Vec3) Dot(b *Vec3) float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	_cret = C.graphene_vec3_dot(_arg0, _arg1)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_vec3_t are equal.
func (v1 *Vec3) Equal(v2 *Vec3) bool {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v1))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(v2))

	_cret = C.graphene_vec3_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by @v
func (v *Vec3) free() {
	var _arg0 *C.graphene_vec3_t // out

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_free(_arg0)
}

// X retrieves the first component of the given vector @v.
func (v *Vec3) X() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	_cret = C.graphene_vec3_get_x(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// XY creates a #graphene_vec2_t that contains the first and second components
// of the given #graphene_vec3_t.
func (v *Vec3) XY() Vec2 {
	var _arg0 *C.graphene_vec3_t // out
	var _res Vec2

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_get_xy(_arg0, (*C.graphene_vec2_t)(unsafe.Pointer(&_res)))

	return _res
}

// XY0 creates a #graphene_vec3_t that contains the first two components of the
// given #graphene_vec3_t, and the third component set to 0.
func (v *Vec3) XY0() Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_get_xy0(_arg0, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// XYZ0 converts a #graphene_vec3_t in a #graphene_vec4_t using 0.0 as the value
// for the fourth component of the resulting vector.
func (v *Vec3) XYZ0() Vec4 {
	var _arg0 *C.graphene_vec3_t // out
	var _res Vec4

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_get_xyz0(_arg0, (*C.graphene_vec4_t)(unsafe.Pointer(&_res)))

	return _res
}

// XYZ1 converts a #graphene_vec3_t in a #graphene_vec4_t using 1.0 as the value
// for the fourth component of the resulting vector.
func (v *Vec3) XYZ1() Vec4 {
	var _arg0 *C.graphene_vec3_t // out
	var _res Vec4

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_get_xyz1(_arg0, (*C.graphene_vec4_t)(unsafe.Pointer(&_res)))

	return _res
}

// Xyzw converts a #graphene_vec3_t in a #graphene_vec4_t using @w as the value
// of the fourth component of the resulting vector.
func (v *Vec3) Xyzw(w float32) Vec4 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 C.float            // out
	var _res Vec4

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))
	_arg1 = C.float(w)

	C.graphene_vec3_get_xyzw(_arg0, _arg1, (*C.graphene_vec4_t)(unsafe.Pointer(&_res)))

	return _res
}

// Y retrieves the second component of the given vector @v.
func (v *Vec3) Y() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	_cret = C.graphene_vec3_get_y(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Z retrieves the third component of the given vector @v.
func (v *Vec3) Z() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	_cret = C.graphene_vec3_get_z(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes a #graphene_vec3_t using the given values.
//
// This function can be called multiple times.
func (v *Vec3) Init(x float32, y float32, z float32) *Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _arg3 C.float            // out
	var _cret *C.graphene_vec3_t // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)

	_cret = C.graphene_vec3_init(_arg0, _arg1, _arg2, _arg3)

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}

// InitFromFloat initializes a #graphene_vec3_t with the values from an array.
func (v *Vec3) InitFromFloat(src [3]float32) *Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.float
	var _cret *C.graphene_vec3_t // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))
	_arg1 = (*C.float)(unsafe.Pointer(&src))

	_cret = C.graphene_vec3_init_from_float(_arg0, _arg1)

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}

// InitFromVec3 initializes a #graphene_vec3_t with the values of another
// #graphene_vec3_t.
func (v *Vec3) InitFromVec3(src *Vec3) *Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _cret *C.graphene_vec3_t // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(src))

	_cret = C.graphene_vec3_init_from_vec3(_arg0, _arg1)

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}

// Interpolate: linearly interpolates @v1 and @v2 using the given @factor.
func (v1 *Vec3) Interpolate(v2 *Vec3, factor float64) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 C.double           // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v1))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(v2))
	_arg2 = C.double(factor)

	C.graphene_vec3_interpolate(_arg0, _arg1, _arg2, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Length retrieves the length of the given vector @v.
func (v *Vec3) Length() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	_cret = C.graphene_vec3_length(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Max compares each component of the two given vectors and creates a vector
// that contains the maximum values.
func (a *Vec3) Max(b *Vec3) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	C.graphene_vec3_max(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Min compares each component of the two given vectors and creates a vector
// that contains the minimum values.
func (a *Vec3) Min(b *Vec3) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	C.graphene_vec3_min(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Multiply multiplies each component of the two given vectors.
func (a *Vec3) Multiply(b *Vec3) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	C.graphene_vec3_multiply(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Near compares the two given #graphene_vec3_t vectors and checks whether their
// values are within the given @epsilon.
func (v1 *Vec3) Near(v2 *Vec3, epsilon float32) bool {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 C.float            // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v1))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(v2))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_vec3_near(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Negate negates the given #graphene_vec3_t.
func (v *Vec3) Negate() Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_negate(_arg0, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Normalize normalizes the given #graphene_vec3_t.
func (v *Vec3) Normalize() Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_normalize(_arg0, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Scale multiplies all components of the given vector with the given scalar
// @factor.
func (v *Vec3) Scale(factor float32) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 C.float            // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))
	_arg1 = C.float(factor)

	C.graphene_vec3_scale(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// Subtract subtracts from each component of the first operand @a the
// corresponding component of the second operand @b and places each result into
// the components of @res.
func (a *Vec3) Subtract(b *Vec3) Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _res Vec3

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(b))

	C.graphene_vec3_subtract(_arg0, _arg1, (*C.graphene_vec3_t)(unsafe.Pointer(&_res)))

	return _res
}

// ToFloat copies the components of a #graphene_vec3_t into the given array.
func (v *Vec3) ToFloat() [3]float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 [3]C.float

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(v))

	C.graphene_vec3_to_float(_arg0, &_arg1[0])

	var _dest [3]float32

	_dest = *(*[3]float32)(unsafe.Pointer(&_arg1))

	return _dest
}

// Vec3One provides a constant pointer to a vector with three components, all
// sets to 1.
func Vec3One() *Vec3 {
	var _cret *C.graphene_vec3_t // in

	_cret = C.graphene_vec3_one()

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}

// Vec3XAxis provides a constant pointer to a vector with three components with
// values set to (1, 0, 0).
func Vec3XAxis() *Vec3 {
	var _cret *C.graphene_vec3_t // in

	_cret = C.graphene_vec3_x_axis()

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}

// Vec3YAxis provides a constant pointer to a vector with three components with
// values set to (0, 1, 0).
func Vec3YAxis() *Vec3 {
	var _cret *C.graphene_vec3_t // in

	_cret = C.graphene_vec3_y_axis()

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}

// Vec3ZAxis provides a constant pointer to a vector with three components with
// values set to (0, 0, 1).
func Vec3ZAxis() *Vec3 {
	var _cret *C.graphene_vec3_t // in

	_cret = C.graphene_vec3_z_axis()

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}

// Vec3Zero provides a constant pointer to a vector with three components, all
// sets to 0.
func Vec3Zero() *Vec3 {
	var _cret *C.graphene_vec3_t // in

	_cret = C.graphene_vec3_zero()

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(_cret))

	return _vec3
}
