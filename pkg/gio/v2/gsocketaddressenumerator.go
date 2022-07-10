// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
// extern void* _gotk4_gio2_SocketAddressEnumeratorClass_next(void*, void*, GError**);
// extern void* _gotk4_gio2_SocketAddressEnumeratorClass_next_finish(void*, void*, GError**);
import "C"

// GTypeSocketAddressEnumerator returns the GType for the type SocketAddressEnumerator.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSocketAddressEnumerator() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "SocketAddressEnumerator").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSocketAddressEnumerator)
	return gtype
}

// SocketAddressEnumeratorOverrider contains methods that are overridable.
type SocketAddressEnumeratorOverrider interface {
	// Next retrieves the next Address from enumerator. Note that this may block
	// for some amount of time. (Eg, a Address may need to do a DNS lookup
	// before it can return an address.) Use
	// g_socket_address_enumerator_next_async() if you need to avoid blocking.
	//
	// If enumerator is expected to yield addresses, but for some reason is
	// unable to (eg, because of a DNS error), then the first call to
	// g_socket_address_enumerator_next() will return an appropriate error in
	// *error. However, if the first call to g_socket_address_enumerator_next()
	// succeeds, then any further internal errors (other than cancellable being
	// triggered) will be ignored.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//
	// The function returns the following values:
	//
	//    - socketAddress (owned by the caller), or NULL on error (in which case
	//      *error will be set) or if there are no more addresses.
	//
	Next(ctx context.Context) (SocketAddresser, error)
	// NextFinish retrieves the result of a completed call to
	// g_socket_address_enumerator_next_async(). See
	// g_socket_address_enumerator_next() for more information about error
	// handling.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - socketAddress (owned by the caller), or NULL on error (in which case
	//      *error will be set) or if there are no more addresses.
	//
	NextFinish(result AsyncResulter) (SocketAddresser, error)
}

// SocketAddressEnumerator is an enumerator type for Address instances. It is
// returned by enumeration functions such as g_socket_connectable_enumerate(),
// which returns a AddressEnumerator to list each Address which could be used to
// connect to that Connectable.
//
// Enumeration is typically a blocking operation, so the asynchronous methods
// g_socket_address_enumerator_next_async() and
// g_socket_address_enumerator_next_finish() should be used where possible.
//
// Each AddressEnumerator can only be enumerated once. Once
// g_socket_address_enumerator_next() has returned NULL, further enumeration
// with that AddressEnumerator is not possible, and it can be unreffed.
type SocketAddressEnumerator struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*SocketAddressEnumerator)(nil)
)

// SocketAddressEnumeratorrer describes types inherited from class SocketAddressEnumerator.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type SocketAddressEnumeratorrer interface {
	coreglib.Objector
	baseSocketAddressEnumerator() *SocketAddressEnumerator
}

var _ SocketAddressEnumeratorrer = (*SocketAddressEnumerator)(nil)

func classInitSocketAddressEnumeratorrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "SocketAddressEnumeratorClass")

	if _, ok := goval.(interface {
		Next(ctx context.Context) (SocketAddresser, error)
	}); ok {
		o := pclass.StructFieldOffset("next")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_SocketAddressEnumeratorClass_next)
	}

	if _, ok := goval.(interface {
		NextFinish(result AsyncResulter) (SocketAddresser, error)
	}); ok {
		o := pclass.StructFieldOffset("next_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_SocketAddressEnumeratorClass_next_finish)
	}
}

//export _gotk4_gio2_SocketAddressEnumeratorClass_next
func _gotk4_gio2_SocketAddressEnumeratorClass_next(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Next(ctx context.Context) (SocketAddresser, error)
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	socketAddress, _goerr := iface.Next(_cancellable)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(socketAddress).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(socketAddress).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_SocketAddressEnumeratorClass_next_finish
func _gotk4_gio2_SocketAddressEnumeratorClass_next_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		NextFinish(result AsyncResulter) (SocketAddresser, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	socketAddress, _goerr := iface.NextFinish(_result)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(socketAddress).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(socketAddress).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapSocketAddressEnumerator(obj *coreglib.Object) *SocketAddressEnumerator {
	return &SocketAddressEnumerator{
		Object: obj,
	}
}

func marshalSocketAddressEnumerator(p uintptr) (interface{}, error) {
	return wrapSocketAddressEnumerator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (enumerator *SocketAddressEnumerator) baseSocketAddressEnumerator() *SocketAddressEnumerator {
	return enumerator
}

// BaseSocketAddressEnumerator returns the underlying base object.
func BaseSocketAddressEnumerator(obj SocketAddressEnumeratorrer) *SocketAddressEnumerator {
	return obj.baseSocketAddressEnumerator()
}

// Next retrieves the next Address from enumerator. Note that this may block for
// some amount of time. (Eg, a Address may need to do a DNS lookup before it can
// return an address.) Use g_socket_address_enumerator_next_async() if you need
// to avoid blocking.
//
// If enumerator is expected to yield addresses, but for some reason is unable
// to (eg, because of a DNS error), then the first call to
// g_socket_address_enumerator_next() will return an appropriate error in
// *error. However, if the first call to g_socket_address_enumerator_next()
// succeeds, then any further internal errors (other than cancellable being
// triggered) will be ignored.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
// The function returns the following values:
//
//    - socketAddress (owned by the caller), or NULL on error (in which case
//      *error will be set) or if there are no more addresses.
//
func (enumerator *SocketAddressEnumerator) Next(ctx context.Context) (SocketAddresser, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "SocketAddressEnumerator")
	_gret := _info.InvokeClassMethod("next", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddresser)
			return ok
		})
		rv, ok := casted.(SocketAddresser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddresser")
		}
		_socketAddress = rv
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _socketAddress, _goerr
}

// NextAsync: asynchronously retrieves the next Address from enumerator and then
// calls callback, which must call g_socket_address_enumerator_next_finish() to
// get the result.
//
// It is an error to call this multiple times before the previous callback has
// finished.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - callback (optional) to call when the request is satisfied.
//
func (enumerator *SocketAddressEnumerator) NextAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[2])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[3] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "SocketAddressEnumerator")
	_info.InvokeClassMethod("next_async", _args[:], nil)

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// NextFinish retrieves the result of a completed call to
// g_socket_address_enumerator_next_async(). See
// g_socket_address_enumerator_next() for more information about error handling.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - socketAddress (owned by the caller), or NULL on error (in which case
//      *error will be set) or if there are no more addresses.
//
func (enumerator *SocketAddressEnumerator) NextFinish(result AsyncResulter) (SocketAddresser, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(enumerator).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "SocketAddressEnumerator")
	_gret := _info.InvokeClassMethod("next_finish", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(enumerator)
	runtime.KeepAlive(result)

	var _socketAddress SocketAddresser // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(SocketAddresser)
			return ok
		})
		rv, ok := casted.(SocketAddresser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddresser")
		}
		_socketAddress = rv
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _socketAddress, _goerr
}
