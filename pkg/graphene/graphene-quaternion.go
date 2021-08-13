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
		{T: externglib.Type(C.graphene_quaternion_get_type()), F: marshalQuaternion},
	})
}

// Quaternion: quaternion.
//
// The contents of the #graphene_quaternion_t structure are private and should
// never be accessed directly.
type Quaternion struct {
	nocopy gextras.NoCopy
	native *C.graphene_quaternion_t
}

func marshalQuaternion(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Quaternion{native: (*C.graphene_quaternion_t)(unsafe.Pointer(b))}, nil
}

// NewQuaternionAlloc constructs a struct Quaternion.
func NewQuaternionAlloc() *Quaternion {
	var _cret *C.graphene_quaternion_t // in

	_cret = C.graphene_quaternion_alloc()

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_quaternion, func(v *Quaternion) {
		C.graphene_quaternion_free((*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _quaternion
}

// Add adds two #graphene_quaternion_t a and b.
func (a *Quaternion) Add(b *Quaternion) Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_quaternion_t // out
	var _arg2 C.graphene_quaternion_t  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_quaternion_add(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Quaternion // out

	_res = *(*Quaternion)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Dot computes the dot product of two #graphene_quaternion_t.
func (a *Quaternion) Dot(b *Quaternion) float32 {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_quaternion_t // out
	var _cret C.float                  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_quaternion_dot(_arg0, _arg1)

	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether the given quaternions are equal.
func (a *Quaternion) Equal(b *Quaternion) bool {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_quaternion_t // out
	var _cret C._Bool                  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_quaternion_equal(_arg0, _arg1)

	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Init initializes a #graphene_quaternion_t using the given four values.
func (q *Quaternion) Init(x float32, y float32, z float32, w float32) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // out
	var _arg2 C.float                  // out
	var _arg3 C.float                  // out
	var _arg4 C.float                  // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)
	_arg4 = C.float(w)

	_cret = C.graphene_quaternion_init(_arg0, _arg1, _arg2, _arg3, _arg4)

	runtime.KeepAlive(q)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(z)
	runtime.KeepAlive(w)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitFromAngleVec3 initializes a #graphene_quaternion_t using an angle on a
// specific axis.
func (q *Quaternion) InitFromAngleVec3(angle float32, axis *Vec3) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // out
	var _arg2 *C.graphene_vec3_t       // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = C.float(angle)
	_arg2 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(axis)))

	_cret = C.graphene_quaternion_init_from_angle_vec3(_arg0, _arg1, _arg2)

	runtime.KeepAlive(q)
	runtime.KeepAlive(angle)
	runtime.KeepAlive(axis)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitFromAngles initializes a #graphene_quaternion_t using the values of the
// Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
//
// See also: graphene_quaternion_init_from_euler()
func (q *Quaternion) InitFromAngles(degX float32, degY float32, degZ float32) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // out
	var _arg2 C.float                  // out
	var _arg3 C.float                  // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = C.float(degX)
	_arg2 = C.float(degY)
	_arg3 = C.float(degZ)

	_cret = C.graphene_quaternion_init_from_angles(_arg0, _arg1, _arg2, _arg3)

	runtime.KeepAlive(q)
	runtime.KeepAlive(degX)
	runtime.KeepAlive(degY)
	runtime.KeepAlive(degZ)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitFromEuler initializes a #graphene_quaternion_t using the given
// #graphene_euler_t.
func (q *Quaternion) InitFromEuler(e *Euler) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_euler_t      // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_quaternion_init_from_euler(_arg0, _arg1)

	runtime.KeepAlive(q)
	runtime.KeepAlive(e)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitFromMatrix initializes a #graphene_quaternion_t using the rotation
// components of a transformation matrix.
func (q *Quaternion) InitFromMatrix(m *Matrix) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_matrix_t     // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_matrix_t)(gextras.StructNative(unsafe.Pointer(m)))

	_cret = C.graphene_quaternion_init_from_matrix(_arg0, _arg1)

	runtime.KeepAlive(q)
	runtime.KeepAlive(m)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitFromQuaternion initializes a #graphene_quaternion_t with the values from
// src.
func (q *Quaternion) InitFromQuaternion(src *Quaternion) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_quaternion_t // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_quaternion_init_from_quaternion(_arg0, _arg1)

	runtime.KeepAlive(q)
	runtime.KeepAlive(src)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitFromRadians initializes a #graphene_quaternion_t using the values of the
// Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
//
// See also: graphene_quaternion_init_from_euler()
func (q *Quaternion) InitFromRadians(radX float32, radY float32, radZ float32) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // out
	var _arg2 C.float                  // out
	var _arg3 C.float                  // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = C.float(radX)
	_arg2 = C.float(radY)
	_arg3 = C.float(radZ)

	_cret = C.graphene_quaternion_init_from_radians(_arg0, _arg1, _arg2, _arg3)

	runtime.KeepAlive(q)
	runtime.KeepAlive(radX)
	runtime.KeepAlive(radY)
	runtime.KeepAlive(radZ)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitFromVec4 initializes a #graphene_quaternion_t with the values from src.
func (q *Quaternion) InitFromVec4(src *Vec4) *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_vec4_t       // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_quaternion_init_from_vec4(_arg0, _arg1)

	runtime.KeepAlive(q)
	runtime.KeepAlive(src)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// InitIdentity initializes a #graphene_quaternion_t using the identity
// transformation.
func (q *Quaternion) InitIdentity() *Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _cret *C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	_cret = C.graphene_quaternion_init_identity(_arg0)

	runtime.KeepAlive(q)

	var _quaternion *Quaternion // out

	_quaternion = (*Quaternion)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _quaternion
}

// Invert inverts a #graphene_quaternion_t, and returns the conjugate quaternion
// of q.
func (q *Quaternion) Invert() Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.graphene_quaternion_t  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quaternion_invert(_arg0, &_arg1)
	runtime.KeepAlive(q)

	var _res Quaternion // out

	_res = *(*Quaternion)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Multiply multiplies two #graphene_quaternion_t a and b.
func (a *Quaternion) Multiply(b *Quaternion) Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_quaternion_t // out
	var _arg2 C.graphene_quaternion_t  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_quaternion_multiply(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res Quaternion // out

	_res = *(*Quaternion)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Normalize normalizes a #graphene_quaternion_t.
func (q *Quaternion) Normalize() Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.graphene_quaternion_t  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quaternion_normalize(_arg0, &_arg1)
	runtime.KeepAlive(q)

	var _res Quaternion // out

	_res = *(*Quaternion)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// Scale scales all the elements of a #graphene_quaternion_t q using the given
// scalar factor.
func (q *Quaternion) Scale(factor float32) Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // out
	var _arg2 C.graphene_quaternion_t  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	_arg1 = C.float(factor)

	C.graphene_quaternion_scale(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(q)
	runtime.KeepAlive(factor)

	var _res Quaternion // out

	_res = *(*Quaternion)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Slerp interpolates between the two given quaternions using a spherical linear
// interpolation, or SLERP (http://en.wikipedia.org/wiki/Slerp), using the given
// interpolation factor.
func (a *Quaternion) Slerp(b *Quaternion, factor float32) Quaternion {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 *C.graphene_quaternion_t // out
	var _arg2 C.float                  // out
	var _arg3 C.graphene_quaternion_t  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg2 = C.float(factor)

	C.graphene_quaternion_slerp(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(factor)

	var _res Quaternion // out

	_res = *(*Quaternion)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// ToAngleVec3 converts a quaternion into an angle, axis pair.
func (q *Quaternion) ToAngleVec3() (float32, Vec3) {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // in
	var _arg2 C.graphene_vec3_t        // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quaternion_to_angle_vec3(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(q)

	var _angle float32 // out
	var _axis Vec3     // out

	_angle = float32(_arg1)
	_axis = *(*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _angle, _axis
}

// ToAngles converts a #graphene_quaternion_t to its corresponding rotations on
// the Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
func (q *Quaternion) ToAngles() (degX float32, degY float32, degZ float32) {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // in
	var _arg2 C.float                  // in
	var _arg3 C.float                  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quaternion_to_angles(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(q)

	var _degX float32 // out
	var _degY float32 // out
	var _degZ float32 // out

	_degX = float32(_arg1)
	_degY = float32(_arg2)
	_degZ = float32(_arg3)

	return _degX, _degY, _degZ
}

// ToMatrix converts a quaternion into a transformation matrix expressing the
// rotation defined by the #graphene_quaternion_t.
func (q *Quaternion) ToMatrix() Matrix {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.graphene_matrix_t      // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quaternion_to_matrix(_arg0, &_arg1)
	runtime.KeepAlive(q)

	var _m Matrix // out

	_m = *(*Matrix)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _m
}

// ToRadians converts a #graphene_quaternion_t to its corresponding rotations on
// the Euler angles (http://en.wikipedia.org/wiki/Euler_angles) on each axis.
func (q *Quaternion) ToRadians() (radX float32, radY float32, radZ float32) {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.float                  // in
	var _arg2 C.float                  // in
	var _arg3 C.float                  // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quaternion_to_radians(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(q)

	var _radX float32 // out
	var _radY float32 // out
	var _radZ float32 // out

	_radX = float32(_arg1)
	_radY = float32(_arg2)
	_radZ = float32(_arg3)

	return _radX, _radY, _radZ
}

// ToVec4 copies the components of a #graphene_quaternion_t into a
// #graphene_vec4_t.
func (q *Quaternion) ToVec4() Vec4 {
	var _arg0 *C.graphene_quaternion_t // out
	var _arg1 C.graphene_vec4_t        // in

	_arg0 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))

	C.graphene_quaternion_to_vec4(_arg0, &_arg1)
	runtime.KeepAlive(q)

	var _res Vec4 // out

	_res = *(*Vec4)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}
