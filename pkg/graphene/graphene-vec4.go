// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_vec4_get_type()), F: marshalVec4},
	})
}

// Vec4: structure capable of holding a vector with four dimensions: x, y, z,
// and w.
//
// The contents of the #graphene_vec4_t structure are private and should never
// be accessed directly.
type Vec4 struct {
	nocopy gextras.NoCopy
	native *C.graphene_vec4_t
}

func marshalVec4(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Vec4{native: (*C.graphene_vec4_t)(unsafe.Pointer(b))}, nil
}

// NewVec4Alloc constructs a struct Vec4.
func NewVec4Alloc() *Vec4 {
	var _cret *C.graphene_vec4_t // in

	_cret = C.graphene_vec4_alloc()

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_vec4, func(v *Vec4) {
		C.graphene_vec4_free((*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _vec4
}

// Add adds each component of the two given vectors.
func (a *Vec4) Add(b *Vec4) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec4_add(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Divide divides each component of the first operand a by the corresponding
// component of the second operand b, and places the results into the vector
// res.
func (a *Vec4) Divide(b *Vec4) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec4_divide(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Dot computes the dot product of the two given vectors.
func (a *Vec4) Dot(b *Vec4) float32 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_vec4_dot(_arg0, _arg1)

	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_vec4_t are equal.
func (v1 *Vec4) Equal(v2 *Vec4) bool {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v1)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v2)))

	_cret = C.graphene_vec4_equal(_arg0, _arg1)

	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// W retrieves the value of the fourth component of the given #graphene_vec4_t.
func (v *Vec4) W() float32 {
	var _arg0 *C.graphene_vec4_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec4_get_w(_arg0)

	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// X retrieves the value of the first component of the given #graphene_vec4_t.
func (v *Vec4) X() float32 {
	var _arg0 *C.graphene_vec4_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec4_get_x(_arg0)

	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// XY creates a #graphene_vec2_t that contains the first two components of the
// given #graphene_vec4_t.
func (v *Vec4) XY() Vec2 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 C.graphene_vec2_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec4_get_xy(_arg0, &_arg1)
	runtime.KeepAlive(v)

	var _res Vec2 // out

	_res = *(*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// XYZ creates a #graphene_vec3_t that contains the first three components of
// the given #graphene_vec4_t.
func (v *Vec4) XYZ() Vec3 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 C.graphene_vec3_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec4_get_xyz(_arg0, &_arg1)
	runtime.KeepAlive(v)

	var _res Vec3 // out

	_res = *(*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Y retrieves the value of the second component of the given #graphene_vec4_t.
func (v *Vec4) Y() float32 {
	var _arg0 *C.graphene_vec4_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec4_get_y(_arg0)

	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Z retrieves the value of the third component of the given #graphene_vec4_t.
func (v *Vec4) Z() float32 {
	var _arg0 *C.graphene_vec4_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec4_get_z(_arg0)

	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes a #graphene_vec4_t using the given values.
//
// This function can be called multiple times.
func (v *Vec4) Init(x float32, y float32, z float32, w float32) *Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _arg3 C.float            // out
	var _arg4 C.float            // out
	var _cret *C.graphene_vec4_t // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)
	_arg4 = C.float(w)

	_cret = C.graphene_vec4_init(_arg0, _arg1, _arg2, _arg3, _arg4)

	runtime.KeepAlive(v)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(z)
	runtime.KeepAlive(w)

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// InitFromFloat initializes a #graphene_vec4_t with the values inside the given
// array.
func (v *Vec4) InitFromFloat(src [4]float32) *Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.float           // out
	var _cret *C.graphene_vec4_t // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = (*C.float)(unsafe.Pointer(&src))

	_cret = C.graphene_vec4_init_from_float(_arg0, _arg1)

	runtime.KeepAlive(v)
	runtime.KeepAlive(src)

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// InitFromVec2 initializes a #graphene_vec4_t using the components of a
// #graphene_vec2_t and the values of z and w.
func (v *Vec4) InitFromVec2(src *Vec2, z float32, w float32) *Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec2_t // out
	var _arg2 C.float            // out
	var _arg3 C.float            // out
	var _cret *C.graphene_vec4_t // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(src)))
	_arg2 = C.float(z)
	_arg3 = C.float(w)

	_cret = C.graphene_vec4_init_from_vec2(_arg0, _arg1, _arg2, _arg3)

	runtime.KeepAlive(v)
	runtime.KeepAlive(src)
	runtime.KeepAlive(z)
	runtime.KeepAlive(w)

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// InitFromVec3 initializes a #graphene_vec4_t using the components of a
// #graphene_vec3_t and the value of w.
func (v *Vec4) InitFromVec3(src *Vec3, w float32) *Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 C.float            // out
	var _cret *C.graphene_vec4_t // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(src)))
	_arg2 = C.float(w)

	_cret = C.graphene_vec4_init_from_vec3(_arg0, _arg1, _arg2)

	runtime.KeepAlive(v)
	runtime.KeepAlive(src)
	runtime.KeepAlive(w)

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// InitFromVec4 initializes a #graphene_vec4_t using the components of another
// #graphene_vec4_t.
func (v *Vec4) InitFromVec4(src *Vec4) *Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _cret *C.graphene_vec4_t // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_vec4_init_from_vec4(_arg0, _arg1)

	runtime.KeepAlive(v)
	runtime.KeepAlive(src)

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// Interpolate: linearly interpolates v1 and v2 using the given factor.
func (v1 *Vec4) Interpolate(v2 *Vec4, factor float64) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.double           // out
	var _arg3 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v1)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v2)))
	_arg2 = C.double(factor)

	C.graphene_vec4_interpolate(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(factor)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Length computes the length of the given #graphene_vec4_t.
func (v *Vec4) Length() float32 {
	var _arg0 *C.graphene_vec4_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_vec4_length(_arg0)

	runtime.KeepAlive(v)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Max compares each component of the two given vectors and creates a vector
// that contains the maximum values.
func (a *Vec4) Max(b *Vec4) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec4_max(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Min compares each component of the two given vectors and creates a vector
// that contains the minimum values.
func (a *Vec4) Min(b *Vec4) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec4_min(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Multiply multiplies each component of the two given vectors.
func (a *Vec4) Multiply(b *Vec4) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec4_multiply(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Near compares the two given #graphene_vec4_t vectors and checks whether their
// values are within the given epsilon.
func (v1 *Vec4) Near(v2 *Vec4, epsilon float32) bool {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.float            // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v1)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v2)))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_vec4_near(_arg0, _arg1, _arg2)

	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(epsilon)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Negate negates the given #graphene_vec4_t.
func (v *Vec4) Negate() Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec4_negate(_arg0, &_arg1)
	runtime.KeepAlive(v)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Normalize normalizes the given #graphene_vec4_t.
func (v *Vec4) Normalize() Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec4_normalize(_arg0, &_arg1)
	runtime.KeepAlive(v)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Scale multiplies all components of the given vector with the given scalar
// factor.
func (v *Vec4) Scale(factor float32) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 C.float            // out
	var _arg2 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))
	_arg1 = C.float(factor)

	C.graphene_vec4_scale(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(v)
	runtime.KeepAlive(factor)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Subtract subtracts from each component of the first operand a the
// corresponding component of the second operand b and places each result into
// the components of res.
func (a *Vec4) Subtract(b *Vec4) Vec4 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 *C.graphene_vec4_t // out
	var _arg2 C.graphene_vec4_t  // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_vec4_subtract(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// ToFloat stores the components of the given #graphene_vec4_t into an array of
// floating point values.
func (v *Vec4) ToFloat() [4]float32 {
	var _arg0 *C.graphene_vec4_t // out
	var _arg1 [4]C.float         // in

	_arg0 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(v)))

	C.graphene_vec4_to_float(_arg0, &_arg1[0])
	runtime.KeepAlive(v)

	var _dest [4]float32 // out

	_dest = *(*[4]float32)(unsafe.Pointer(&_arg1))

	return _dest
}

// Vec4One retrieves a pointer to a #graphene_vec4_t with all its components set
// to 1.
func Vec4One() *Vec4 {
	var _cret *C.graphene_vec4_t // in

	_cret = C.graphene_vec4_one()

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// Vec4WAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (0, 0, 0, 1).
func Vec4WAxis() *Vec4 {
	var _cret *C.graphene_vec4_t // in

	_cret = C.graphene_vec4_w_axis()

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// Vec4XAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (1, 0, 0, 0).
func Vec4XAxis() *Vec4 {
	var _cret *C.graphene_vec4_t // in

	_cret = C.graphene_vec4_x_axis()

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// Vec4YAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (0, 1, 0, 0).
func Vec4YAxis() *Vec4 {
	var _cret *C.graphene_vec4_t // in

	_cret = C.graphene_vec4_y_axis()

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// Vec4ZAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (0, 0, 1, 0).
func Vec4ZAxis() *Vec4 {
	var _cret *C.graphene_vec4_t // in

	_cret = C.graphene_vec4_z_axis()

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}

// Vec4Zero retrieves a pointer to a #graphene_vec4_t with all its components
// set to 0.
func Vec4Zero() *Vec4 {
	var _cret *C.graphene_vec4_t // in

	_cret = C.graphene_vec4_zero()

	var _vec4 *Vec4 // out

	_vec4 = (*Vec4)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _vec4
}
