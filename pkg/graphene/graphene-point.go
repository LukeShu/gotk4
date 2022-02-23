// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

// glib.Type values for graphene-point.go.
var GTypePoint = externglib.Type(C.graphene_point_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePoint, F: marshalPoint},
	})
}

// Point: point with two coordinates.
//
// An instance of this type is always passed by reference.
type Point struct {
	*point
}

// point is the struct that's finalized.
type point struct {
	native *C.graphene_point_t
}

func marshalPoint(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Point{&point{(*C.graphene_point_t)(b)}}, nil
}

// NewPointAlloc constructs a struct Point.
func NewPointAlloc() *Point {
	var _cret *C.graphene_point_t // in

	_cret = C.graphene_point_alloc()

	var _point *Point // out

	_point = (*Point)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_point)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_point_free((*C.graphene_point_t)(intern.C))
		},
	)

	return _point
}

// X coordinate of the point.
func (p *Point) X() float32 {
	var v float32 // out
	v = float32(p.native.x)
	return v
}

// Y coordinate of the point.
func (p *Point) Y() float32 {
	var v float32 // out
	v = float32(p.native.y)
	return v
}

// Distance computes the distance between a and b.
//
// The function takes the following parameters:
//
//    - b: #graphene_point_t.
//
// The function returns the following values:
//
//    - dX (optional): distance component on the X axis.
//    - dY (optional): distance component on the Y axis.
//    - gfloat: distance between the two points.
//
func (a *Point) Distance(b *Point) (dX float32, dY float32, gfloat float32) {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 C.float             // in
	var _arg3 C.float             // in
	var _cret C.float             // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_point_distance(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _dX float32     // out
	var _dY float32     // out
	var _gfloat float32 // out

	_dX = float32(_arg2)
	_dY = float32(_arg3)
	_gfloat = float32(_cret)

	return _dX, _dY, _gfloat
}

// Equal checks if the two points a and b point to the same coordinates.
//
// This function accounts for floating point fluctuations; if you want to
// control the fuzziness of the match, you can use graphene_point_near()
// instead.
//
// The function takes the following parameters:
//
//    - b: #graphene_point_t.
//
// The function returns the following values:
//
//    - ok: true if the points have the same coordinates.
//
func (a *Point) Equal(b *Point) bool {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_point_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Init initializes p to the given x and y coordinates.
//
// It's safe to call this function multiple times.
//
// The function takes the following parameters:
//
//    - x: x coordinate.
//    - y: y coordinate.
//
// The function returns the following values:
//
//    - point: initialized point.
//
func (p *Point) Init(x float32, y float32) *Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 C.float             // out
	var _arg2 C.float             // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)

	_cret = C.graphene_point_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(p)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _point *Point // out

	_point = (*Point)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point
}

// InitFromPoint initializes p with the same coordinates of src.
//
// The function takes the following parameters:
//
//    - src to use.
//
// The function returns the following values:
//
//    - point: initialized point.
//
func (p *Point) InitFromPoint(src *Point) *Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_point_init_from_point(_arg0, _arg1)
	runtime.KeepAlive(p)
	runtime.KeepAlive(src)

	var _point *Point // out

	_point = (*Point)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point
}

// InitFromVec2 initializes p with the coordinates inside the given
// #graphene_vec2_t.
//
// The function takes the following parameters:
//
//    - src: #graphene_vec2_t.
//
// The function returns the following values:
//
//    - point: initialized point.
//
func (p *Point) InitFromVec2(src *Vec2) *Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_vec2_t  // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_vec2_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_point_init_from_vec2(_arg0, _arg1)
	runtime.KeepAlive(p)
	runtime.KeepAlive(src)

	var _point *Point // out

	_point = (*Point)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point
}

// Interpolate: linearly interpolates the coordinates of a and b using the given
// factor.
//
// The function takes the following parameters:
//
//    - b: #graphene_point_t.
//    - factor: linear interpolation factor.
//
// The function returns the following values:
//
//    - res: return location for the interpolated point.
//
func (a *Point) Interpolate(b *Point, factor float64) *Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 C.double            // out
	var _arg3 C.graphene_point_t  // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg2 = C.double(factor)

	C.graphene_point_interpolate(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(factor)

	var _res *Point // out

	_res = (*Point)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Near checks whether the two points a and b are within the threshold of
// epsilon.
//
// The function takes the following parameters:
//
//    - b: #graphene_point_t.
//    - epsilon: threshold between the two points.
//
// The function returns the following values:
//
//    - ok: true if the distance is within epsilon.
//
func (a *Point) Near(b *Point, epsilon float32) bool {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 C.float             // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_point_near(_arg0, _arg1, _arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(epsilon)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ToVec2 stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
//
// The function returns the following values:
//
//    - v: return location for the vertex.
//
func (p *Point) ToVec2() *Vec2 {
	var _arg0 *C.graphene_point_t // out
	var _arg1 C.graphene_vec2_t   // in

	_arg0 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_point_to_vec2(_arg0, &_arg1)
	runtime.KeepAlive(p)

	var _v *Vec2 // out

	_v = (*Vec2)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _v
}

// PointZero returns a point fixed at (0, 0).
//
// The function returns the following values:
//
//    - point: fixed point.
//
func PointZero() *Point {
	var _cret *C.graphene_point_t // in

	_cret = C.graphene_point_zero()

	var _point *Point // out

	_point = (*Point)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point
}
