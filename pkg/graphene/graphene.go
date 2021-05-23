// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <graphene-gobject.h>
//
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		// Enums
		// Skipped EulerOrder.
		// Skipped RayIntersectionKind.

		// Records
		{T: externglib.Type(C.graphene_box_get_type()), F: marshalBox},
		{T: externglib.Type(C.graphene_euler_get_type()), F: marshalEuler},
		{T: externglib.Type(C.graphene_frustum_get_type()), F: marshalFrustum},
		{T: externglib.Type(C.graphene_matrix_get_type()), F: marshalMatrix},
		{T: externglib.Type(C.graphene_plane_get_type()), F: marshalPlane},
		{T: externglib.Type(C.graphene_point_get_type()), F: marshalPoint},
		{T: externglib.Type(C.graphene_point3d_get_type()), F: marshalPoint3D},
		{T: externglib.Type(C.graphene_quad_get_type()), F: marshalQuad},
		{T: externglib.Type(C.graphene_quaternion_get_type()), F: marshalQuaternion},
		{T: externglib.Type(C.graphene_ray_get_type()), F: marshalRay},
		{T: externglib.Type(C.graphene_rect_get_type()), F: marshalRect},
		// Skipped Simd4F.
		// Skipped Simd4X4F.
		{T: externglib.Type(C.graphene_size_get_type()), F: marshalSize},
		{T: externglib.Type(C.graphene_sphere_get_type()), F: marshalSphere},
		{T: externglib.Type(C.graphene_triangle_get_type()), F: marshalTriangle},
		{T: externglib.Type(C.graphene_vec2_get_type()), F: marshalVec2},
		{T: externglib.Type(C.graphene_vec3_get_type()), F: marshalVec3},
		{T: externglib.Type(C.graphene_vec4_get_type()), F: marshalVec4},
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

// RayIntersectionKind: the type of intersection.
type RayIntersectionKind int

const (
	// RayIntersectionKindNone: no intersection
	RayIntersectionKindNone RayIntersectionKind = 0
	// RayIntersectionKindEnter: the ray is entering the intersected object
	RayIntersectionKindEnter RayIntersectionKind = 1
	// RayIntersectionKindLeave: the ray is leaving the intersected object
	RayIntersectionKindLeave RayIntersectionKind = 2
)

// BoxEmpty: a degenerate #graphene_box_t that can only be expanded.
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxEmpty() *Box {
	ret := C.graphene_box_empty()

	var ret0 *Box
	ret0 = wrapBox(ret)

	return ret0
}

// BoxInfinite: a degenerate #graphene_box_t that cannot be expanded.
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxInfinite() *Box {
	ret := C.graphene_box_infinite()

	var ret0 *Box
	ret0 = wrapBox(ret)

	return ret0
}

// BoxMinusOne: a #graphene_box_t with the minimum vertex set at (-1, -1, -1)
// and the maximum vertex set at (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxMinusOne() *Box {
	ret := C.graphene_box_minus_one()

	var ret0 *Box
	ret0 = wrapBox(ret)

	return ret0
}

// BoxOne: a #graphene_box_t with the minimum vertex set at (0, 0, 0) and the
// maximum vertex set at (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOne() *Box {
	ret := C.graphene_box_one()

	var ret0 *Box
	ret0 = wrapBox(ret)

	return ret0
}

// BoxOneMinusOne: a #graphene_box_t with the minimum vertex set at (-1, -1, -1)
// and the maximum vertex set at (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOneMinusOne() *Box {
	ret := C.graphene_box_one_minus_one()

	var ret0 *Box
	ret0 = wrapBox(ret)

	return ret0
}

// BoxZero: a #graphene_box_t with both the minimum and maximum vertices set at
// (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxZero() *Box {
	ret := C.graphene_box_zero()

	var ret0 *Box
	ret0 = wrapBox(ret)

	return ret0
}

// Point3DZero retrieves a constant point with all three coordinates set to 0.
func Point3DZero() *Point3D {
	ret := C.graphene_point3d_zero()

	var ret0 *Point3D
	ret0 = wrapPoint3D(ret)

	return ret0
}

// PointZero returns a point fixed at (0, 0).
func PointZero() *Point {
	ret := C.graphene_point_zero()

	var ret0 *Point
	ret0 = wrapPoint(ret)

	return ret0
}

// RectAlloc allocates a new #graphene_rect_t.
//
// The contents of the returned rectangle are undefined.
func RectAlloc() *Rect {
	ret := C.graphene_rect_alloc()

	var ret0 *Rect
	ret0 = wrapRect(ret)

	return ret0
}

// RectZero returns a degenerate rectangle with origin fixed at (0, 0) and a
// size of 0, 0.
func RectZero() *Rect {
	ret := C.graphene_rect_zero()

	var ret0 *Rect
	ret0 = wrapRect(ret)

	return ret0
}

// SizeZero: a constant pointer to a zero #graphene_size_t, useful for equality
// checks and interpolations.
func SizeZero() *Size {
	ret := C.graphene_size_zero()

	var ret0 *Size
	ret0 = wrapSize(ret)

	return ret0
}

// Vec2One retrieves a constant vector with (1, 1) components.
func Vec2One() *Vec2 {
	ret := C.graphene_vec2_one()

	var ret0 *Vec2
	ret0 = wrapVec2(ret)

	return ret0
}

// Vec2XAxis retrieves a constant vector with (1, 0) components.
func Vec2XAxis() *Vec2 {
	ret := C.graphene_vec2_x_axis()

	var ret0 *Vec2
	ret0 = wrapVec2(ret)

	return ret0
}

// Vec2YAxis retrieves a constant vector with (0, 1) components.
func Vec2YAxis() *Vec2 {
	ret := C.graphene_vec2_y_axis()

	var ret0 *Vec2
	ret0 = wrapVec2(ret)

	return ret0
}

// Vec2Zero retrieves a constant vector with (0, 0) components.
func Vec2Zero() *Vec2 {
	ret := C.graphene_vec2_zero()

	var ret0 *Vec2
	ret0 = wrapVec2(ret)

	return ret0
}

// Vec3One provides a constant pointer to a vector with three components, all
// sets to 1.
func Vec3One() *Vec3 {
	ret := C.graphene_vec3_one()

	var ret0 *Vec3
	ret0 = wrapVec3(ret)

	return ret0
}

// Vec3XAxis provides a constant pointer to a vector with three components with
// values set to (1, 0, 0).
func Vec3XAxis() *Vec3 {
	ret := C.graphene_vec3_x_axis()

	var ret0 *Vec3
	ret0 = wrapVec3(ret)

	return ret0
}

// Vec3YAxis provides a constant pointer to a vector with three components with
// values set to (0, 1, 0).
func Vec3YAxis() *Vec3 {
	ret := C.graphene_vec3_y_axis()

	var ret0 *Vec3
	ret0 = wrapVec3(ret)

	return ret0
}

// Vec3ZAxis provides a constant pointer to a vector with three components with
// values set to (0, 0, 1).
func Vec3ZAxis() *Vec3 {
	ret := C.graphene_vec3_z_axis()

	var ret0 *Vec3
	ret0 = wrapVec3(ret)

	return ret0
}

// Vec3Zero provides a constant pointer to a vector with three components, all
// sets to 0.
func Vec3Zero() *Vec3 {
	ret := C.graphene_vec3_zero()

	var ret0 *Vec3
	ret0 = wrapVec3(ret)

	return ret0
}

// Vec4One retrieves a pointer to a #graphene_vec4_t with all its components set
// to 1.
func Vec4One() *Vec4 {
	ret := C.graphene_vec4_one()

	var ret0 *Vec4
	ret0 = wrapVec4(ret)

	return ret0
}

// Vec4WAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (0, 0, 0, 1).
func Vec4WAxis() *Vec4 {
	ret := C.graphene_vec4_w_axis()

	var ret0 *Vec4
	ret0 = wrapVec4(ret)

	return ret0
}

// Vec4XAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (1, 0, 0, 0).
func Vec4XAxis() *Vec4 {
	ret := C.graphene_vec4_x_axis()

	var ret0 *Vec4
	ret0 = wrapVec4(ret)

	return ret0
}

// Vec4YAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (0, 1, 0, 0).
func Vec4YAxis() *Vec4 {
	ret := C.graphene_vec4_y_axis()

	var ret0 *Vec4
	ret0 = wrapVec4(ret)

	return ret0
}

// Vec4ZAxis retrieves a pointer to a #graphene_vec4_t with its components set
// to (0, 0, 1, 0).
func Vec4ZAxis() *Vec4 {
	ret := C.graphene_vec4_z_axis()

	var ret0 *Vec4
	ret0 = wrapVec4(ret)

	return ret0
}

// Vec4Zero retrieves a pointer to a #graphene_vec4_t with all its components
// set to 0.
func Vec4Zero() *Vec4 {
	ret := C.graphene_vec4_zero()

	var ret0 *Vec4
	ret0 = wrapVec4(ret)

	return ret0
}

// Box: a 3D box, described as the volume between a minimum and a maximum
// vertices.
type Box struct {
	native *C.graphene_box_t
}

func wrapBox(p *C.graphene_box_t) *Box {
	v := Box{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Box).free)

	return &v
}

func marshalBox(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_box_t)(unsafe.Pointer(b))

	return wrapBox(c)
}

func (b *Box) free() {}

// Native returns the pointer to *C.graphene_box_t. The caller is expected to
// cast.
func (b *Box) Native() unsafe.Pointer {
	return unsafe.Pointer(b.native)
}

func NewBox() *Box

// Euler: describe a rotation using Euler angles.
//
// The contents of the #graphene_euler_t structure are private and should never
// be accessed directly.
type Euler struct {
	native *C.graphene_euler_t
}

func wrapEuler(p *C.graphene_euler_t) *Euler {
	v := Euler{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Euler).free)

	return &v
}

func marshalEuler(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_euler_t)(unsafe.Pointer(b))

	return wrapEuler(c)
}

func (e *Euler) free() {}

// Native returns the pointer to *C.graphene_euler_t. The caller is expected to
// cast.
func (e *Euler) Native() unsafe.Pointer {
	return unsafe.Pointer(e.native)
}

func NewEuler() *Euler

// Frustum: a 3D volume delimited by 2D clip planes.
//
// The contents of the `graphene_frustum_t` are private, and should not be
// modified directly.
type Frustum struct {
	native *C.graphene_frustum_t
}

func wrapFrustum(p *C.graphene_frustum_t) *Frustum {
	v := Frustum{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Frustum).free)

	return &v
}

func marshalFrustum(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_frustum_t)(unsafe.Pointer(b))

	return wrapFrustum(c)
}

func (f *Frustum) free() {}

// Native returns the pointer to *C.graphene_frustum_t. The caller is expected to
// cast.
func (f *Frustum) Native() unsafe.Pointer {
	return unsafe.Pointer(f.native)
}

func NewFrustum() *Frustum

// Matrix: a structure capable of holding a 4x4 matrix.
//
// The contents of the #graphene_matrix_t structure are private and should never
// be accessed directly.
type Matrix struct {
	native *C.graphene_matrix_t
}

func wrapMatrix(p *C.graphene_matrix_t) *Matrix {
	v := Matrix{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Matrix).free)

	return &v
}

func marshalMatrix(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_matrix_t)(unsafe.Pointer(b))

	return wrapMatrix(c)
}

func (m *Matrix) free() {}

// Native returns the pointer to *C.graphene_matrix_t. The caller is expected to
// cast.
func (m *Matrix) Native() unsafe.Pointer {
	return unsafe.Pointer(m.native)
}

func NewMatrix() *Matrix

// Plane: a 2D plane that extends infinitely in a 3D volume.
//
// The contents of the `graphene_plane_t` are private, and should not be
// modified directly.
type Plane struct {
	native *C.graphene_plane_t
}

func wrapPlane(p *C.graphene_plane_t) *Plane {
	v := Plane{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Plane).free)

	return &v
}

func marshalPlane(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_plane_t)(unsafe.Pointer(b))

	return wrapPlane(c)
}

func (p *Plane) free() {}

// Native returns the pointer to *C.graphene_plane_t. The caller is expected to
// cast.
func (p *Plane) Native() unsafe.Pointer {
	return unsafe.Pointer(p.native)
}

func NewPlane() *Plane

// Point: a point with two coordinates.
type Point struct {
	// X: the X coordinate of the point
	X float32
	// Y: the Y coordinate of the point
	Y float32

	native *C.graphene_point_t
}

func wrapPoint(p *C.graphene_point_t) *Point {
	var v Point

	v.X = float32(p.x)
	v.Y = float32(p.y)

	return &v
}

func marshalPoint(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_point_t)(unsafe.Pointer(b))

	return wrapPoint(c)
}

// Native returns the pointer to *C.graphene_point_t. The caller is expected to
// cast.
func (p *Point) Native() unsafe.Pointer {
	return unsafe.Pointer(p.native)
}

func NewPoint() *Point

// Point3D: a point with three components: X, Y, and Z.
type Point3D struct {
	// X: the X coordinate
	X float32
	// Y: the Y coordinate
	Y float32
	// Z: the Z coordinate
	Z float32

	native *C.graphene_point3d_t
}

func wrapPoint3D(p *C.graphene_point3d_t) *Point3D {
	var v Point3D

	v.X = float32(p.x)
	v.Y = float32(p.y)
	v.Z = float32(p.z)

	return &v
}

func marshalPoint3D(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_point3d_t)(unsafe.Pointer(b))

	return wrapPoint3D(c)
}

// Native returns the pointer to *C.graphene_point3d_t. The caller is expected to
// cast.
func (p *Point3D) Native() unsafe.Pointer {
	return unsafe.Pointer(p.native)
}

func NewPoint3D() *Point3D

// Quad: a 4 vertex quadrilateral, as represented by four #graphene_point_t.
//
// The contents of a #graphene_quad_t are private and should never be accessed
// directly.
type Quad struct {
	native *C.graphene_quad_t
}

func wrapQuad(p *C.graphene_quad_t) *Quad {
	v := Quad{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Quad).free)

	return &v
}

func marshalQuad(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_quad_t)(unsafe.Pointer(b))

	return wrapQuad(c)
}

func (q *Quad) free() {}

// Native returns the pointer to *C.graphene_quad_t. The caller is expected to
// cast.
func (q *Quad) Native() unsafe.Pointer {
	return unsafe.Pointer(q.native)
}

func NewQuad() *Quad

// Quaternion: a quaternion.
//
// The contents of the #graphene_quaternion_t structure are private and should
// never be accessed directly.
type Quaternion struct {
	native *C.graphene_quaternion_t
}

func wrapQuaternion(p *C.graphene_quaternion_t) *Quaternion {
	v := Quaternion{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Quaternion).free)

	return &v
}

func marshalQuaternion(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_quaternion_t)(unsafe.Pointer(b))

	return wrapQuaternion(c)
}

func (q *Quaternion) free() {}

// Native returns the pointer to *C.graphene_quaternion_t. The caller is expected to
// cast.
func (q *Quaternion) Native() unsafe.Pointer {
	return unsafe.Pointer(q.native)
}

func NewQuaternion() *Quaternion

// Ray: a ray emitted from an origin in a given direction.
//
// The contents of the `graphene_ray_t` structure are private, and should not be
// modified directly.
type Ray struct {
	native *C.graphene_ray_t
}

func wrapRay(p *C.graphene_ray_t) *Ray {
	v := Ray{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Ray).free)

	return &v
}

func marshalRay(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_ray_t)(unsafe.Pointer(b))

	return wrapRay(c)
}

func (r *Ray) free() {}

// Native returns the pointer to *C.graphene_ray_t. The caller is expected to
// cast.
func (r *Ray) Native() unsafe.Pointer {
	return unsafe.Pointer(r.native)
}

func NewRay() *Ray

// Rect: the location and size of a rectangle region.
//
// The width and height of a #graphene_rect_t can be negative; for instance, a
// #graphene_rect_t with an origin of [ 0, 0 ] and a size of [ 10, 10 ] is
// equivalent to a #graphene_rect_t with an origin of [ 10, 10 ] and a size of [
// -10, -10 ].
//
// Application code can normalize rectangles using graphene_rect_normalize();
// this function will ensure that the width and height of a rectangle are
// positive values. All functions taking a #graphene_rect_t as an argument will
// internally operate on a normalized copy; all functions returning a
// #graphene_rect_t will always return a normalized rectangle.
type Rect struct {
	// Origin: the coordinates of the origin of the rectangle
	Origin Point
	// Size: the size of the rectangle
	Size Size

	native *C.graphene_rect_t
}

func wrapRect(p *C.graphene_rect_t) *Rect {
	var v Rect

	v.Origin = wrapPoint(p.origin)
	v.Size = wrapSize(p.size)

	return &v
}

func marshalRect(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_rect_t)(unsafe.Pointer(b))

	return wrapRect(c)
}

// Native returns the pointer to *C.graphene_rect_t. The caller is expected to
// cast.
func (r *Rect) Native() unsafe.Pointer {
	return unsafe.Pointer(r.native)
}

type SIMD4F struct {
	native *C.graphene_simd4f_t
}

func wrapSIMD4F(p *C.graphene_simd4f_t) *SIMD4F {
	v := SIMD4F{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*SIMD4F).free)

	return &v
}

func marshalSIMD4F(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_simd4f_t)(unsafe.Pointer(b))

	return wrapSIMD4F(c)
}

func (s *SIMD4F) free() {}

// Native returns the pointer to *C.graphene_simd4f_t. The caller is expected to
// cast.
func (s *SIMD4F) Native() unsafe.Pointer {
	return unsafe.Pointer(s.native)
}

type SIMD4X4F struct {
	native *C.graphene_simd4x4f_t
}

func wrapSIMD4X4F(p *C.graphene_simd4x4f_t) *SIMD4X4F {
	v := SIMD4X4F{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*SIMD4X4F).free)

	return &v
}

func marshalSIMD4X4F(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_simd4x4f_t)(unsafe.Pointer(b))

	return wrapSIMD4X4F(c)
}

func (s *SIMD4X4F) free() {}

// Native returns the pointer to *C.graphene_simd4x4f_t. The caller is expected to
// cast.
func (s *SIMD4X4F) Native() unsafe.Pointer {
	return unsafe.Pointer(s.native)
}

// Size: a size.
type Size struct {
	// Width: the width
	Width float32
	// Height: the height
	Height float32

	native *C.graphene_size_t
}

func wrapSize(p *C.graphene_size_t) *Size {
	var v Size

	v.Width = float32(p.width)
	v.Height = float32(p.height)

	return &v
}

func marshalSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_size_t)(unsafe.Pointer(b))

	return wrapSize(c)
}

// Native returns the pointer to *C.graphene_size_t. The caller is expected to
// cast.
func (s *Size) Native() unsafe.Pointer {
	return unsafe.Pointer(s.native)
}

func NewSize() *Size

// Sphere: a sphere, represented by its center and radius.
type Sphere struct {
	native *C.graphene_sphere_t
}

func wrapSphere(p *C.graphene_sphere_t) *Sphere {
	v := Sphere{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Sphere).free)

	return &v
}

func marshalSphere(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_sphere_t)(unsafe.Pointer(b))

	return wrapSphere(c)
}

func (s *Sphere) free() {}

// Native returns the pointer to *C.graphene_sphere_t. The caller is expected to
// cast.
func (s *Sphere) Native() unsafe.Pointer {
	return unsafe.Pointer(s.native)
}

func NewSphere() *Sphere

// Triangle: a triangle.
type Triangle struct {
	native *C.graphene_triangle_t
}

func wrapTriangle(p *C.graphene_triangle_t) *Triangle {
	v := Triangle{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Triangle).free)

	return &v
}

func marshalTriangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_triangle_t)(unsafe.Pointer(b))

	return wrapTriangle(c)
}

func (t *Triangle) free() {}

// Native returns the pointer to *C.graphene_triangle_t. The caller is expected to
// cast.
func (t *Triangle) Native() unsafe.Pointer {
	return unsafe.Pointer(t.native)
}

func NewTriangle() *Triangle

// Vec2: a structure capable of holding a vector with two dimensions, x and y.
//
// The contents of the #graphene_vec2_t structure are private and should never
// be accessed directly.
type Vec2 struct {
	native *C.graphene_vec2_t
}

func wrapVec2(p *C.graphene_vec2_t) *Vec2 {
	v := Vec2{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Vec2).free)

	return &v
}

func marshalVec2(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_vec2_t)(unsafe.Pointer(b))

	return wrapVec2(c)
}

func (v *Vec2) free() {}

// Native returns the pointer to *C.graphene_vec2_t. The caller is expected to
// cast.
func (v *Vec2) Native() unsafe.Pointer {
	return unsafe.Pointer(v.native)
}

func NewVec2() *Vec2

// Vec3: a structure capable of holding a vector with three dimensions: x, y,
// and z.
//
// The contents of the #graphene_vec3_t structure are private and should never
// be accessed directly.
type Vec3 struct {
	native *C.graphene_vec3_t
}

func wrapVec3(p *C.graphene_vec3_t) *Vec3 {
	v := Vec3{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Vec3).free)

	return &v
}

func marshalVec3(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_vec3_t)(unsafe.Pointer(b))

	return wrapVec3(c)
}

func (v *Vec3) free() {}

// Native returns the pointer to *C.graphene_vec3_t. The caller is expected to
// cast.
func (v *Vec3) Native() unsafe.Pointer {
	return unsafe.Pointer(v.native)
}

func NewVec3() *Vec3

// Vec4: a structure capable of holding a vector with four dimensions: x, y, z,
// and w.
//
// The contents of the #graphene_vec4_t structure are private and should never
// be accessed directly.
type Vec4 struct {
	native *C.graphene_vec4_t
}

func wrapVec4(p *C.graphene_vec4_t) *Vec4 {
	v := Vec4{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Vec4).free)

	return &v
}

func marshalVec4(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.graphene_vec4_t)(unsafe.Pointer(b))

	return wrapVec4(c)
}

func (v *Vec4) free() {}

// Native returns the pointer to *C.graphene_vec4_t. The caller is expected to
// cast.
func (v *Vec4) Native() unsafe.Pointer {
	return unsafe.Pointer(v.native)
}

func NewVec4() *Vec4
