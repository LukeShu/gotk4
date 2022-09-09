// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern gssize _gotk4_gio2_SocketAddressClass_get_native_size(GSocketAddress*);
// extern gboolean _gotk4_gio2_SocketAddressClass_to_native(GSocketAddress*, gpointer, gsize, GError**);
// extern GSocketFamily _gotk4_gio2_SocketAddressClass_get_family(GSocketAddress*);
// GSocketFamily _gotk4_gio2_SocketAddress_virtual_get_family(void* fnptr, GSocketAddress* arg0) {
//   return ((GSocketFamily (*)(GSocketAddress*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_SocketAddress_virtual_to_native(void* fnptr, GSocketAddress* arg0, gpointer arg1, gsize arg2, GError** arg3) {
//   return ((gboolean (*)(GSocketAddress*, gpointer, gsize, GError**))(fnptr))(arg0, arg1, arg2, arg3);
// };
// gssize _gotk4_gio2_SocketAddress_virtual_get_native_size(void* fnptr, GSocketAddress* arg0) {
//   return ((gssize (*)(GSocketAddress*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeSocketAddress = coreglib.Type(C.g_socket_address_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSocketAddress, F: marshalSocketAddress},
	})
}

// SocketAddressOverrides contains methods that are overridable.
type SocketAddressOverrides struct {
	// Family gets the socket family type of address.
	//
	// The function returns the following values:
	//
	//    - socketFamily: socket family type of address.
	//
	Family func() SocketFamily
	// NativeSize gets the size of address's native struct sockaddr. You can use
	// this to allocate memory to pass to g_socket_address_to_native().
	//
	// The function returns the following values:
	//
	//    - gssize: size of the native struct sockaddr that address represents.
	//
	NativeSize func() int
	// ToNative converts a Address to a native struct sockaddr, which can be
	// passed to low-level functions like connect() or bind().
	//
	// If not enough space is available, a G_IO_ERROR_NO_SPACE error is
	// returned. If the address type is not known on the system then a
	// G_IO_ERROR_NOT_SUPPORTED error is returned.
	//
	// The function takes the following parameters:
	//
	//    - dest (optional): pointer to a memory location that will contain the
	//      native struct sockaddr.
	//    - destlen: size of dest. Must be at least as large as
	//      g_socket_address_get_native_size().
	//
	ToNative func(dest unsafe.Pointer, destlen uint) error
}

func defaultSocketAddressOverrides(v *SocketAddress) SocketAddressOverrides {
	return SocketAddressOverrides{
		Family:     v.family,
		NativeSize: v.nativeSize,
		ToNative:   v.toNative,
	}
}

// SocketAddress is the equivalent of struct sockaddr in the BSD sockets API.
// This is an abstract class; use SocketAddress for internet sockets, or
// SocketAddress for UNIX domain sockets.
type SocketAddress struct {
	_ [0]func() // equal guard
	*coreglib.Object

	SocketConnectable
}

var (
	_ coreglib.Objector = (*SocketAddress)(nil)
)

// SocketAddresser describes types inherited from class SocketAddress.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type SocketAddresser interface {
	coreglib.Objector
	baseSocketAddress() *SocketAddress
}

var _ SocketAddresser = (*SocketAddress)(nil)

func init() {
	coreglib.RegisterClassInfo[*SocketAddress, *SocketAddressClass, SocketAddressOverrides](
		GTypeSocketAddress,
		initSocketAddressClass,
		wrapSocketAddress,
		defaultSocketAddressOverrides,
	)
}

func initSocketAddressClass(gclass unsafe.Pointer, overrides SocketAddressOverrides, classInitFunc func(*SocketAddressClass)) {
	pclass := (*C.GSocketAddressClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeSocketAddress))))

	if overrides.Family != nil {
		pclass.get_family = (*[0]byte)(C._gotk4_gio2_SocketAddressClass_get_family)
	}

	if overrides.NativeSize != nil {
		pclass.get_native_size = (*[0]byte)(C._gotk4_gio2_SocketAddressClass_get_native_size)
	}

	if overrides.ToNative != nil {
		pclass.to_native = (*[0]byte)(C._gotk4_gio2_SocketAddressClass_to_native)
	}

	if classInitFunc != nil {
		class := (*SocketAddressClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSocketAddress(obj *coreglib.Object) *SocketAddress {
	return &SocketAddress{
		Object: obj,
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalSocketAddress(p uintptr) (interface{}, error) {
	return wrapSocketAddress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (address *SocketAddress) baseSocketAddress() *SocketAddress {
	return address
}

// BaseSocketAddress returns the underlying base object.
func BaseSocketAddress(obj SocketAddresser) *SocketAddress {
	return obj.baseSocketAddress()
}

// NewSocketAddressFromNative creates a Address subclass corresponding to the
// native struct sockaddr native.
//
// The function takes the following parameters:
//
//    - native: pointer to a struct sockaddr.
//    - len: size of the memory location pointed to by native.
//
// The function returns the following values:
//
//    - socketAddress: new Address if native could successfully be converted,
//      otherwise NULL.
//
func NewSocketAddressFromNative(native unsafe.Pointer, len uint) *SocketAddress {
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (C.gpointer)(unsafe.Pointer(native))
	_arg2 = C.gsize(len)

	_cret = C.g_socket_address_new_from_native(_arg1, _arg2)
	runtime.KeepAlive(native)
	runtime.KeepAlive(len)

	var _socketAddress *SocketAddress // out

	_socketAddress = wrapSocketAddress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketAddress
}

// Family gets the socket family type of address.
//
// The function returns the following values:
//
//    - socketFamily: socket family type of address.
//
func (address *SocketAddress) Family() SocketFamily {
	var _arg0 *C.GSocketAddress // out
	var _cret C.GSocketFamily   // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))

	_cret = C.g_socket_address_get_family(_arg0)
	runtime.KeepAlive(address)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// NativeSize gets the size of address's native struct sockaddr. You can use
// this to allocate memory to pass to g_socket_address_to_native().
//
// The function returns the following values:
//
//    - gssize: size of the native struct sockaddr that address represents.
//
func (address *SocketAddress) NativeSize() int {
	var _arg0 *C.GSocketAddress // out
	var _cret C.gssize          // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))

	_cret = C.g_socket_address_get_native_size(_arg0)
	runtime.KeepAlive(address)

	var _gssize int // out

	_gssize = int(_cret)

	return _gssize
}

// ToNative converts a Address to a native struct sockaddr, which can be passed
// to low-level functions like connect() or bind().
//
// If not enough space is available, a G_IO_ERROR_NO_SPACE error is returned. If
// the address type is not known on the system then a G_IO_ERROR_NOT_SUPPORTED
// error is returned.
//
// The function takes the following parameters:
//
//    - dest (optional): pointer to a memory location that will contain the
//      native struct sockaddr.
//    - destlen: size of dest. Must be at least as large as
//      g_socket_address_get_native_size().
//
func (address *SocketAddress) ToNative(dest unsafe.Pointer, destlen uint) error {
	var _arg0 *C.GSocketAddress // out
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(dest))
	_arg2 = C.gsize(destlen)

	C.g_socket_address_to_native(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(address)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(destlen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Family gets the socket family type of address.
//
// The function returns the following values:
//
//    - socketFamily: socket family type of address.
//
func (address *SocketAddress) family() SocketFamily {
	gclass := (*C.GSocketAddressClass)(coreglib.PeekParentClass(address))
	fnarg := gclass.get_family

	var _arg0 *C.GSocketAddress // out
	var _cret C.GSocketFamily   // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))

	_cret = C._gotk4_gio2_SocketAddress_virtual_get_family(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(address)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

// nativeSize gets the size of address's native struct sockaddr. You can use
// this to allocate memory to pass to g_socket_address_to_native().
//
// The function returns the following values:
//
//    - gssize: size of the native struct sockaddr that address represents.
//
func (address *SocketAddress) nativeSize() int {
	gclass := (*C.GSocketAddressClass)(coreglib.PeekParentClass(address))
	fnarg := gclass.get_native_size

	var _arg0 *C.GSocketAddress // out
	var _cret C.gssize          // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))

	_cret = C._gotk4_gio2_SocketAddress_virtual_get_native_size(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(address)

	var _gssize int // out

	_gssize = int(_cret)

	return _gssize
}

// toNative converts a Address to a native struct sockaddr, which can be passed
// to low-level functions like connect() or bind().
//
// If not enough space is available, a G_IO_ERROR_NO_SPACE error is returned. If
// the address type is not known on the system then a G_IO_ERROR_NOT_SUPPORTED
// error is returned.
//
// The function takes the following parameters:
//
//    - dest (optional): pointer to a memory location that will contain the
//      native struct sockaddr.
//    - destlen: size of dest. Must be at least as large as
//      g_socket_address_get_native_size().
//
func (address *SocketAddress) toNative(dest unsafe.Pointer, destlen uint) error {
	gclass := (*C.GSocketAddressClass)(coreglib.PeekParentClass(address))
	fnarg := gclass.to_native

	var _arg0 *C.GSocketAddress // out
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(dest))
	_arg2 = C.gsize(destlen)

	C._gotk4_gio2_SocketAddress_virtual_to_native(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(address)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(destlen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SocketAddressClass: instance of this type is always passed by reference.
type SocketAddressClass struct {
	*socketAddressClass
}

// socketAddressClass is the struct that's finalized.
type socketAddressClass struct {
	native *C.GSocketAddressClass
}
