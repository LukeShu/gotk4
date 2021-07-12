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
		{T: externglib.Type(C.graphene_sphere_get_type()), F: marshalSphere},
	})
}

// Sphere: sphere, represented by its center and radius.
type Sphere struct {
	native C.graphene_sphere_t
}

func marshalSphere(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Sphere)(unsafe.Pointer(b)), nil
}

// NewSphereAlloc constructs a struct Sphere.
func NewSphereAlloc() *Sphere {
	var _cret *C.graphene_sphere_t // in

	_cret = C.graphene_sphere_alloc()

	var _sphere *Sphere // out

	_sphere = (*Sphere)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_sphere, func(v *Sphere) {
		C.graphene_sphere_free((*C.graphene_sphere_t)(unsafe.Pointer(v)))
	})

	return _sphere
}

// Native returns the underlying C source pointer.
func (s *Sphere) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// ContainsPoint checks whether the given @point is contained in the volume of a
// #graphene_sphere_t.
func (s *Sphere) ContainsPoint(point *Point3D) bool {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point))

	_cret = C.graphene_sphere_contains_point(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Distance computes the distance of the given @point from the surface of a
// #graphene_sphere_t.
func (s *Sphere) Distance(point *Point3D) float32 {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C.float               // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point))

	_cret = C.graphene_sphere_distance(_arg0, _arg1)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether two #graphene_sphere_t are equal.
func (a *Sphere) Equal(b *Sphere) bool {
	var _arg0 *C.graphene_sphere_t // out
	var _arg1 *C.graphene_sphere_t // out
	var _cret C._Bool              // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(a))
	_arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(b))

	_cret = C.graphene_sphere_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by graphene_sphere_alloc().
func (s *Sphere) free() {
	var _arg0 *C.graphene_sphere_t // out

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))

	C.graphene_sphere_free(_arg0)
}

// BoundingBox computes the bounding box capable of containing the given
// #graphene_sphere_t.
func (s *Sphere) BoundingBox() Box {
	var _arg0 *C.graphene_sphere_t // out
	var _box Box

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))

	C.graphene_sphere_get_bounding_box(_arg0, (*C.graphene_box_t)(unsafe.Pointer(&_box)))

	return _box
}

// Center retrieves the coordinates of the center of a #graphene_sphere_t.
func (s *Sphere) Center() Point3D {
	var _arg0 *C.graphene_sphere_t // out
	var _center Point3D

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))

	C.graphene_sphere_get_center(_arg0, (*C.graphene_point3d_t)(unsafe.Pointer(&_center)))

	return _center
}

// Radius retrieves the radius of a #graphene_sphere_t.
func (s *Sphere) Radius() float32 {
	var _arg0 *C.graphene_sphere_t // out
	var _cret C.float              // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))

	_cret = C.graphene_sphere_get_radius(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes the given #graphene_sphere_t with the given @center and
// @radius.
func (s *Sphere) Init(center *Point3D, radius float32) *Sphere {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.float               // out
	var _cret *C.graphene_sphere_t  // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(center))
	_arg2 = C.float(radius)

	_cret = C.graphene_sphere_init(_arg0, _arg1, _arg2)

	var _sphere *Sphere // out

	_sphere = (*Sphere)(unsafe.Pointer(_cret))

	return _sphere
}

// InitFromPoints initializes the given #graphene_sphere_t using the given array
// of 3D coordinates so that the sphere includes them.
//
// The center of the sphere can either be specified, or will be center of the 3D
// volume that encompasses all @points.
func (s *Sphere) InitFromPoints(points []Point3D, center *Point3D) *Sphere {
	var _arg0 *C.graphene_sphere_t // out
	var _arg2 *C.graphene_point3d_t
	var _arg1 C.uint
	var _arg3 *C.graphene_point3d_t // out
	var _cret *C.graphene_sphere_t  // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))
	_arg1 = C.uint(len(points))
	_arg2 = (*C.graphene_point3d_t)(unsafe.Pointer(&points[0]))
	_arg3 = (*C.graphene_point3d_t)(unsafe.Pointer(center))

	_cret = C.graphene_sphere_init_from_points(_arg0, _arg1, _arg2, _arg3)

	var _sphere *Sphere // out

	_sphere = (*Sphere)(unsafe.Pointer(_cret))

	return _sphere
}

// InitFromVectors initializes the given #graphene_sphere_t using the given
// array of 3D coordinates so that the sphere includes them.
//
// The center of the sphere can either be specified, or will be center of the 3D
// volume that encompasses all @vectors.
func (s *Sphere) InitFromVectors(vectors []Vec3, center *Point3D) *Sphere {
	var _arg0 *C.graphene_sphere_t // out
	var _arg2 *C.graphene_vec3_t
	var _arg1 C.uint
	var _arg3 *C.graphene_point3d_t // out
	var _cret *C.graphene_sphere_t  // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))
	_arg1 = C.uint(len(vectors))
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(&vectors[0]))
	_arg3 = (*C.graphene_point3d_t)(unsafe.Pointer(center))

	_cret = C.graphene_sphere_init_from_vectors(_arg0, _arg1, _arg2, _arg3)

	var _sphere *Sphere // out

	_sphere = (*Sphere)(unsafe.Pointer(_cret))

	return _sphere
}

// IsEmpty checks whether the sphere has a zero radius.
func (s *Sphere) IsEmpty() bool {
	var _arg0 *C.graphene_sphere_t // out
	var _cret C._Bool              // in

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))

	_cret = C.graphene_sphere_is_empty(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Translate translates the center of the given #graphene_sphere_t using the
// @point coordinates as the delta of the translation.
func (s *Sphere) Translate(point *Point3D) Sphere {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _res Sphere

	_arg0 = (*C.graphene_sphere_t)(unsafe.Pointer(s))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point))

	C.graphene_sphere_translate(_arg0, _arg1, (*C.graphene_sphere_t)(unsafe.Pointer(&_res)))

	return _res
}
