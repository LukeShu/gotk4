// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

// GTypeFrustum returns the GType for the type Frustum.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFrustum() coreglib.Type {
	gtype := coreglib.Type(C.graphene_frustum_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalFrustum)
	return gtype
}

// Frustum: 3D volume delimited by 2D clip planes.
//
// The contents of the graphene_frustum_t are private, and should not be
// modified directly.
//
// An instance of this type is always passed by reference.
type Frustum struct {
	*frustum
}

// frustum is the struct that's finalized.
type frustum struct {
	native *C.graphene_frustum_t
}

func marshalFrustum(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Frustum{&frustum{(*C.graphene_frustum_t)(b)}}, nil
}

// NewFrustumAlloc constructs a struct Frustum.
func NewFrustumAlloc() *Frustum {
	var _cret *C.graphene_frustum_t // in

	_cret = C.graphene_frustum_alloc()

	var _frustum *Frustum // out

	_frustum = (*Frustum)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_frustum)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_frustum_free((*C.graphene_frustum_t)(intern.C))
		},
	)

	return _frustum
}

// ContainsPoint checks whether a point is inside the volume defined by the
// given #graphene_frustum_t.
//
// The function takes the following parameters:
//
//    - point: #graphene_point3d_t.
//
// The function returns the following values:
//
//    - ok: true if the point is inside the frustum.
//
func (f *Frustum) ContainsPoint(point *Point3D) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(f)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	_cret = C.graphene_frustum_contains_point(_arg0, _arg1)
	runtime.KeepAlive(f)
	runtime.KeepAlive(point)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Equal checks whether the two given #graphene_frustum_t are equal.
//
// The function takes the following parameters:
//
//    - b: #graphene_frustum_t.
//
// The function returns the following values:
//
//    - ok: true if the given frustums are equal.
//
func (a *Frustum) Equal(b *Frustum) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_frustum_t // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_frustum_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Planes retrieves the planes that define the given #graphene_frustum_t.
//
// The function returns the following values:
//
//    - planes: return location for an array of 6 #graphene_plane_t.
//
func (f *Frustum) Planes() [6]Plane {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 [6]C.graphene_plane_t // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(f)))

	C.graphene_frustum_get_planes(_arg0, &_arg1[0])
	runtime.KeepAlive(f)

	var _planes [6]Plane // out

	{
		src := &_arg1
		for i := 0; i < 6; i++ {
			_planes[i] = *(*Plane)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}

	return _planes
}

// Init initializes the given #graphene_frustum_t using the provided clipping
// planes.
//
// The function takes the following parameters:
//
//    - p0: clipping plane.
//    - p1: clipping plane.
//    - p2: clipping plane.
//    - p3: clipping plane.
//    - p4: clipping plane.
//    - p5: clipping plane.
//
// The function returns the following values:
//
//    - frustum: initialized frustum.
//
func (f *Frustum) Init(p0 *Plane, p1 *Plane, p2 *Plane, p3 *Plane, p4 *Plane, p5 *Plane) *Frustum {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_plane_t   // out
	var _arg2 *C.graphene_plane_t   // out
	var _arg3 *C.graphene_plane_t   // out
	var _arg4 *C.graphene_plane_t   // out
	var _arg5 *C.graphene_plane_t   // out
	var _arg6 *C.graphene_plane_t   // out
	var _cret *C.graphene_frustum_t // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(f)))
	_arg1 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p0)))
	_arg2 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p1)))
	_arg3 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p2)))
	_arg4 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p3)))
	_arg5 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p4)))
	_arg6 = (*C.graphene_plane_t)(gextras.StructNative(unsafe.Pointer(p5)))

	_cret = C.graphene_frustum_init(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(f)
	runtime.KeepAlive(p0)
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
	runtime.KeepAlive(p3)
	runtime.KeepAlive(p4)
	runtime.KeepAlive(p5)

	var _frustum *Frustum // out

	_frustum = (*Frustum)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _frustum
}

// InitFromFrustum initializes the given #graphene_frustum_t using the clipping
// planes of another #graphene_frustum_t.
//
// The function takes the following parameters:
//
//    - src: #graphene_frustum_t.
//
// The function returns the following values:
//
//    - frustum: initialized frustum.
//
func (f *Frustum) InitFromFrustum(src *Frustum) *Frustum {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_frustum_t // out
	var _cret *C.graphene_frustum_t // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(f)))
	_arg1 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_frustum_init_from_frustum(_arg0, _arg1)
	runtime.KeepAlive(f)
	runtime.KeepAlive(src)

	var _frustum *Frustum // out

	_frustum = (*Frustum)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _frustum
}

// InitFromMatrix initializes a #graphene_frustum_t using the given matrix.
//
// The function takes the following parameters:
//
//    - matrix: #graphene_matrix_t.
//
// The function returns the following values:
//
//    - frustum: initialized frustum.
//
func (f *Frustum) InitFromMatrix(matrix *Matrix) *Frustum {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_matrix_t  // out
	var _cret *C.graphene_frustum_t // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(f)))
	_arg1 = (*C.graphene_matrix_t)(gextras.StructNative(unsafe.Pointer(matrix)))

	_cret = C.graphene_frustum_init_from_matrix(_arg0, _arg1)
	runtime.KeepAlive(f)
	runtime.KeepAlive(matrix)

	var _frustum *Frustum // out

	_frustum = (*Frustum)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _frustum
}

// IntersectsBox checks whether the given box intersects a plane of a
// #graphene_frustum_t.
//
// The function takes the following parameters:
//
//    - box: #graphene_box_t.
//
// The function returns the following values:
//
//    - ok: true if the box intersects the frustum.
//
func (f *Frustum) IntersectsBox(box *Box) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_box_t     // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(f)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	_cret = C.graphene_frustum_intersects_box(_arg0, _arg1)
	runtime.KeepAlive(f)
	runtime.KeepAlive(box)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsSphere checks whether the given sphere intersects a plane of a
// #graphene_frustum_t.
//
// The function takes the following parameters:
//
//    - sphere: #graphene_sphere_t.
//
// The function returns the following values:
//
//    - ok: true if the sphere intersects the frustum.
//
func (f *Frustum) IntersectsSphere(sphere *Sphere) bool {
	var _arg0 *C.graphene_frustum_t // out
	var _arg1 *C.graphene_sphere_t  // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_frustum_t)(gextras.StructNative(unsafe.Pointer(f)))
	_arg1 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(sphere)))

	_cret = C.graphene_frustum_intersects_sphere(_arg0, _arg1)
	runtime.KeepAlive(f)
	runtime.KeepAlive(sphere)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
