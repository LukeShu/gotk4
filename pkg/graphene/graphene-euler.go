// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"fmt"
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
		{T: externglib.Type(C.graphene_euler_get_type()), F: marshalEuler},
	})
}

// EulerOrder: specify the order of the rotations on each axis.
//
// The GRAPHENE_EULER_ORDER_DEFAULT value is special, and is used as an alias
// for one of the other orders.
type EulerOrder int

const (
	// EulerOrderDefault: rotate in the default order; the default order is one
	// of the following enumeration values
	EulerOrderDefault EulerOrder = -1
	// EulerOrderXYZ: rotate in the X, Y, and Z order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SXYZ
	EulerOrderXYZ EulerOrder = 0
	// EulerOrderYZX: rotate in the Y, Z, and X order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SYZX
	EulerOrderYZX EulerOrder = 1
	// EulerOrderZXY: rotate in the Z, X, and Y order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SZXY
	EulerOrderZXY EulerOrder = 2
	// EulerOrderXZY: rotate in the X, Z, and Y order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SXZY
	EulerOrderXZY EulerOrder = 3
	// EulerOrderYXZ: rotate in the Y, X, and Z order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SYXZ
	EulerOrderYXZ EulerOrder = 4
	// EulerOrderZYX: rotate in the Z, Y, and X order. Deprecated in Graphene
	// 1.10, it's an alias for GRAPHENE_EULER_ORDER_SZYX
	EulerOrderZYX EulerOrder = 5
	// EulerOrderSXYZ defines a static rotation along the X, Y, and Z axes
	// (Since: 1.10)
	EulerOrderSXYZ EulerOrder = 6
	// EulerOrderSXYX defines a static rotation along the X, Y, and X axes
	// (Since: 1.10)
	EulerOrderSXYX EulerOrder = 7
	// EulerOrderSXZY defines a static rotation along the X, Z, and Y axes
	// (Since: 1.10)
	EulerOrderSXZY EulerOrder = 8
	// EulerOrderSXZX defines a static rotation along the X, Z, and X axes
	// (Since: 1.10)
	EulerOrderSXZX EulerOrder = 9
	// EulerOrderSYZX defines a static rotation along the Y, Z, and X axes
	// (Since: 1.10)
	EulerOrderSYZX EulerOrder = 10
	// EulerOrderSYZY defines a static rotation along the Y, Z, and Y axes
	// (Since: 1.10)
	EulerOrderSYZY EulerOrder = 11
	// EulerOrderSYXZ defines a static rotation along the Y, X, and Z axes
	// (Since: 1.10)
	EulerOrderSYXZ EulerOrder = 12
	// EulerOrderSYXY defines a static rotation along the Y, X, and Y axes
	// (Since: 1.10)
	EulerOrderSYXY EulerOrder = 13
	// EulerOrderSZXY defines a static rotation along the Z, X, and Y axes
	// (Since: 1.10)
	EulerOrderSZXY EulerOrder = 14
	// EulerOrderSZXZ defines a static rotation along the Z, X, and Z axes
	// (Since: 1.10)
	EulerOrderSZXZ EulerOrder = 15
	// EulerOrderSZYX defines a static rotation along the Z, Y, and X axes
	// (Since: 1.10)
	EulerOrderSZYX EulerOrder = 16
	// EulerOrderSZYZ defines a static rotation along the Z, Y, and Z axes
	// (Since: 1.10)
	EulerOrderSZYZ EulerOrder = 17
	// EulerOrderRZYX defines a relative rotation along the Z, Y, and X axes
	// (Since: 1.10)
	EulerOrderRZYX EulerOrder = 18
	// EulerOrderRXYX defines a relative rotation along the X, Y, and X axes
	// (Since: 1.10)
	EulerOrderRXYX EulerOrder = 19
	// EulerOrderRYZX defines a relative rotation along the Y, Z, and X axes
	// (Since: 1.10)
	EulerOrderRYZX EulerOrder = 20
	// EulerOrderRXZX defines a relative rotation along the X, Z, and X axes
	// (Since: 1.10)
	EulerOrderRXZX EulerOrder = 21
	// EulerOrderRXZY defines a relative rotation along the X, Z, and Y axes
	// (Since: 1.10)
	EulerOrderRXZY EulerOrder = 22
	// EulerOrderRYZY defines a relative rotation along the Y, Z, and Y axes
	// (Since: 1.10)
	EulerOrderRYZY EulerOrder = 23
	// EulerOrderRZXY defines a relative rotation along the Z, X, and Y axes
	// (Since: 1.10)
	EulerOrderRZXY EulerOrder = 24
	// EulerOrderRYXY defines a relative rotation along the Y, X, and Y axes
	// (Since: 1.10)
	EulerOrderRYXY EulerOrder = 25
	// EulerOrderRYXZ defines a relative rotation along the Y, X, and Z axes
	// (Since: 1.10)
	EulerOrderRYXZ EulerOrder = 26
	// EulerOrderRZXZ defines a relative rotation along the Z, X, and Z axes
	// (Since: 1.10)
	EulerOrderRZXZ EulerOrder = 27
	// EulerOrderRXYZ defines a relative rotation along the X, Y, and Z axes
	// (Since: 1.10)
	EulerOrderRXYZ EulerOrder = 28
	// EulerOrderRZYZ defines a relative rotation along the Z, Y, and Z axes
	// (Since: 1.10)
	EulerOrderRZYZ EulerOrder = 29
)

// String returns the name in string for EulerOrder.
func (e EulerOrder) String() string {
	switch e {
	case EulerOrderDefault:
		return "Default"
	case EulerOrderXYZ:
		return "XYZ"
	case EulerOrderYZX:
		return "YZX"
	case EulerOrderZXY:
		return "ZXY"
	case EulerOrderXZY:
		return "XZY"
	case EulerOrderYXZ:
		return "YXZ"
	case EulerOrderZYX:
		return "ZYX"
	case EulerOrderSXYZ:
		return "SXYZ"
	case EulerOrderSXYX:
		return "SXYX"
	case EulerOrderSXZY:
		return "SXZY"
	case EulerOrderSXZX:
		return "SXZX"
	case EulerOrderSYZX:
		return "SYZX"
	case EulerOrderSYZY:
		return "SYZY"
	case EulerOrderSYXZ:
		return "SYXZ"
	case EulerOrderSYXY:
		return "SYXY"
	case EulerOrderSZXY:
		return "SZXY"
	case EulerOrderSZXZ:
		return "SZXZ"
	case EulerOrderSZYX:
		return "SZYX"
	case EulerOrderSZYZ:
		return "SZYZ"
	case EulerOrderRZYX:
		return "RZYX"
	case EulerOrderRXYX:
		return "RXYX"
	case EulerOrderRYZX:
		return "RYZX"
	case EulerOrderRXZX:
		return "RXZX"
	case EulerOrderRXZY:
		return "RXZY"
	case EulerOrderRYZY:
		return "RYZY"
	case EulerOrderRZXY:
		return "RZXY"
	case EulerOrderRYXY:
		return "RYXY"
	case EulerOrderRYXZ:
		return "RYXZ"
	case EulerOrderRZXZ:
		return "RZXZ"
	case EulerOrderRXYZ:
		return "RXYZ"
	case EulerOrderRZYZ:
		return "RZYZ"
	default:
		return fmt.Sprintf("EulerOrder(%d)", e)
	}
}

// Euler: describe a rotation using Euler angles.
//
// The contents of the #graphene_euler_t structure are private and should never
// be accessed directly.
type Euler struct {
	nocopy gextras.NoCopy
	native *C.graphene_euler_t
}

func marshalEuler(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Euler{native: (*C.graphene_euler_t)(unsafe.Pointer(b))}, nil
}

// NewEulerAlloc constructs a struct Euler.
func NewEulerAlloc() *Euler {
	var _cret *C.graphene_euler_t // in

	_cret = C.graphene_euler_alloc()

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_euler, func(v *Euler) {
		C.graphene_euler_free((*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _euler
}

// Equal checks if two #graphene_euler_t are equal.
func (a *Euler) Equal(b *Euler) bool {
	var _arg0 *C.graphene_euler_t // out
	var _arg1 *C.graphene_euler_t // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_euler_equal(_arg0, _arg1)

	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Alpha retrieves the first component of the Euler angle vector, depending on
// the order of rotation.
//
// See also: graphene_euler_get_x()
func (e *Euler) Alpha() float32 {
	var _arg0 *C.graphene_euler_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_euler_get_alpha(_arg0)

	runtime.KeepAlive(e)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Beta retrieves the second component of the Euler angle vector, depending on
// the order of rotation.
//
// See also: graphene_euler_get_y()
func (e *Euler) Beta() float32 {
	var _arg0 *C.graphene_euler_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_euler_get_beta(_arg0)

	runtime.KeepAlive(e)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Gamma retrieves the third component of the Euler angle vector, depending on
// the order of rotation.
//
// See also: graphene_euler_get_z()
func (e *Euler) Gamma() float32 {
	var _arg0 *C.graphene_euler_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_euler_get_gamma(_arg0)

	runtime.KeepAlive(e)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Order retrieves the order used to apply the rotations described in the
// #graphene_euler_t structure, when converting to and from other structures,
// like #graphene_quaternion_t and #graphene_matrix_t.
//
// This function does not return the GRAPHENE_EULER_ORDER_DEFAULT enumeration
// value; it will return the effective order of rotation instead.
func (e *Euler) Order() EulerOrder {
	var _arg0 *C.graphene_euler_t      // out
	var _cret C.graphene_euler_order_t // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_euler_get_order(_arg0)

	runtime.KeepAlive(e)

	var _eulerOrder EulerOrder // out

	_eulerOrder = EulerOrder(_cret)

	return _eulerOrder
}

// X retrieves the rotation angle on the X axis, in degrees.
func (e *Euler) X() float32 {
	var _arg0 *C.graphene_euler_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_euler_get_x(_arg0)

	runtime.KeepAlive(e)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Y retrieves the rotation angle on the Y axis, in degrees.
func (e *Euler) Y() float32 {
	var _arg0 *C.graphene_euler_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_euler_get_y(_arg0)

	runtime.KeepAlive(e)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Z retrieves the rotation angle on the Z axis, in degrees.
func (e *Euler) Z() float32 {
	var _arg0 *C.graphene_euler_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	_cret = C.graphene_euler_get_z(_arg0)

	runtime.KeepAlive(e)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes a #graphene_euler_t using the given angles.
//
// The order of the rotations is GRAPHENE_EULER_ORDER_DEFAULT.
func (e *Euler) Init(x float32, y float32, z float32) *Euler {
	var _arg0 *C.graphene_euler_t // out
	var _arg1 C.float             // out
	var _arg2 C.float             // out
	var _arg3 C.float             // out
	var _cret *C.graphene_euler_t // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)

	_cret = C.graphene_euler_init(_arg0, _arg1, _arg2, _arg3)

	runtime.KeepAlive(e)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(z)

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _euler
}

// InitFromEuler initializes a #graphene_euler_t using the angles and order of
// another #graphene_euler_t.
//
// If the #graphene_euler_t src is NULL, this function is equivalent to calling
// graphene_euler_init() with all angles set to 0.
func (e *Euler) InitFromEuler(src *Euler) *Euler {
	var _arg0 *C.graphene_euler_t // out
	var _arg1 *C.graphene_euler_t // out
	var _cret *C.graphene_euler_t // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	if src != nil {
		_arg1 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(src)))
	}

	_cret = C.graphene_euler_init_from_euler(_arg0, _arg1)

	runtime.KeepAlive(e)
	runtime.KeepAlive(src)

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _euler
}

// InitFromMatrix initializes a #graphene_euler_t using the given rotation
// matrix.
//
// If the #graphene_matrix_t m is NULL, the #graphene_euler_t will be
// initialized with all angles set to 0.
func (e *Euler) InitFromMatrix(m *Matrix, order EulerOrder) *Euler {
	var _arg0 *C.graphene_euler_t      // out
	var _arg1 *C.graphene_matrix_t     // out
	var _arg2 C.graphene_euler_order_t // out
	var _cret *C.graphene_euler_t      // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	if m != nil {
		_arg1 = (*C.graphene_matrix_t)(gextras.StructNative(unsafe.Pointer(m)))
	}
	_arg2 = C.graphene_euler_order_t(order)

	_cret = C.graphene_euler_init_from_matrix(_arg0, _arg1, _arg2)

	runtime.KeepAlive(e)
	runtime.KeepAlive(m)
	runtime.KeepAlive(order)

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _euler
}

// InitFromQuaternion initializes a #graphene_euler_t using the given normalized
// quaternion.
//
// If the #graphene_quaternion_t q is NULL, the #graphene_euler_t will be
// initialized with all angles set to 0.
func (e *Euler) InitFromQuaternion(q *Quaternion, order EulerOrder) *Euler {
	var _arg0 *C.graphene_euler_t      // out
	var _arg1 *C.graphene_quaternion_t // out
	var _arg2 C.graphene_euler_order_t // out
	var _cret *C.graphene_euler_t      // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	if q != nil {
		_arg1 = (*C.graphene_quaternion_t)(gextras.StructNative(unsafe.Pointer(q)))
	}
	_arg2 = C.graphene_euler_order_t(order)

	_cret = C.graphene_euler_init_from_quaternion(_arg0, _arg1, _arg2)

	runtime.KeepAlive(e)
	runtime.KeepAlive(q)
	runtime.KeepAlive(order)

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _euler
}

// InitFromRadians initializes a #graphene_euler_t using the given angles and
// order of rotation.
func (e *Euler) InitFromRadians(x float32, y float32, z float32, order EulerOrder) *Euler {
	var _arg0 *C.graphene_euler_t      // out
	var _arg1 C.float                  // out
	var _arg2 C.float                  // out
	var _arg3 C.float                  // out
	var _arg4 C.graphene_euler_order_t // out
	var _cret *C.graphene_euler_t      // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)
	_arg4 = C.graphene_euler_order_t(order)

	_cret = C.graphene_euler_init_from_radians(_arg0, _arg1, _arg2, _arg3, _arg4)

	runtime.KeepAlive(e)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(z)
	runtime.KeepAlive(order)

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _euler
}

// InitFromVec3 initializes a #graphene_euler_t using the angles contained in a
// #graphene_vec3_t.
//
// If the #graphene_vec3_t v is NULL, the #graphene_euler_t will be initialized
// with all angles set to 0.
func (e *Euler) InitFromVec3(v *Vec3, order EulerOrder) *Euler {
	var _arg0 *C.graphene_euler_t      // out
	var _arg1 *C.graphene_vec3_t       // out
	var _arg2 C.graphene_euler_order_t // out
	var _cret *C.graphene_euler_t      // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	if v != nil {
		_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(v)))
	}
	_arg2 = C.graphene_euler_order_t(order)

	_cret = C.graphene_euler_init_from_vec3(_arg0, _arg1, _arg2)

	runtime.KeepAlive(e)
	runtime.KeepAlive(v)
	runtime.KeepAlive(order)

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _euler
}

// InitWithOrder initializes a #graphene_euler_t with the given angles and
// order.
func (e *Euler) InitWithOrder(x float32, y float32, z float32, order EulerOrder) *Euler {
	var _arg0 *C.graphene_euler_t      // out
	var _arg1 C.float                  // out
	var _arg2 C.float                  // out
	var _arg3 C.float                  // out
	var _arg4 C.graphene_euler_order_t // out
	var _cret *C.graphene_euler_t      // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)
	_arg4 = C.graphene_euler_order_t(order)

	_cret = C.graphene_euler_init_with_order(_arg0, _arg1, _arg2, _arg3, _arg4)

	runtime.KeepAlive(e)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(z)
	runtime.KeepAlive(order)

	var _euler *Euler // out

	_euler = (*Euler)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _euler
}

// Reorder reorders a #graphene_euler_t using order.
//
// This function is equivalent to creating a #graphene_quaternion_t from the
// given #graphene_euler_t, and then converting the quaternion into another
// #graphene_euler_t.
func (e *Euler) Reorder(order EulerOrder) Euler {
	var _arg0 *C.graphene_euler_t      // out
	var _arg1 C.graphene_euler_order_t // out
	var _arg2 C.graphene_euler_t       // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))
	_arg1 = C.graphene_euler_order_t(order)

	C.graphene_euler_reorder(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(e)
	runtime.KeepAlive(order)

	var _res Euler // out

	_res = *(*Euler)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// ToMatrix converts a #graphene_euler_t into a transformation matrix expressing
// the extrinsic composition of rotations described by the Euler angles.
//
// The rotations are applied over the reference frame axes in the order
// associated with the #graphene_euler_t; for instance, if the order used to
// initialize e is GRAPHENE_EULER_ORDER_XYZ:
//
//    * the first rotation moves the body around the X axis with
//      an angle φ
//    * the second rotation moves the body around the Y axis with
//      an angle of ϑ
//    * the third rotation moves the body around the Z axis with
//      an angle of ψ
//
// The rotation sign convention is right-handed, to preserve compatibility
// between Euler-based, quaternion-based, and angle-axis-based rotations.
func (e *Euler) ToMatrix() Matrix {
	var _arg0 *C.graphene_euler_t // out
	var _arg1 C.graphene_matrix_t // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	C.graphene_euler_to_matrix(_arg0, &_arg1)
	runtime.KeepAlive(e)

	var _res Matrix // out

	_res = *(*Matrix)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// ToQuaternion converts a #graphene_euler_t into a #graphene_quaternion_t.
func (e *Euler) ToQuaternion() Quaternion {
	var _arg0 *C.graphene_euler_t     // out
	var _arg1 C.graphene_quaternion_t // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	C.graphene_euler_to_quaternion(_arg0, &_arg1)
	runtime.KeepAlive(e)

	var _res Quaternion // out

	_res = *(*Quaternion)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// ToVec3 retrieves the angles of a #graphene_euler_t and initializes a
// #graphene_vec3_t with them.
func (e *Euler) ToVec3() Vec3 {
	var _arg0 *C.graphene_euler_t // out
	var _arg1 C.graphene_vec3_t   // in

	_arg0 = (*C.graphene_euler_t)(gextras.StructNative(unsafe.Pointer(e)))

	C.graphene_euler_to_vec3(_arg0, &_arg1)
	runtime.KeepAlive(e)

	var _res Vec3 // out

	_res = *(*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}
