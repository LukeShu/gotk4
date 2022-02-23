// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// glib.Type values for gtcpwrapperconnection.go.
var GTypeTCPWrapperConnection = externglib.Type(C.g_tcp_wrapper_connection_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTCPWrapperConnection, F: marshalTCPWrapperConnection},
	})
}

// TCPWrapperConnectionOverrider contains methods that are overridable.
type TCPWrapperConnectionOverrider interface {
}

// TCPWrapperConnection can be used to wrap a OStream that is based on a
// #GSocket, but which is not actually a Connection. This is used by Client so
// that it can always return a Connection, even when the connection it has
// actually created is not directly a Connection.
type TCPWrapperConnection struct {
	_ [0]func() // equal guard
	TCPConnection
}

var (
	_ IOStreamer = (*TCPWrapperConnection)(nil)
)

func classInitTCPWrapperConnectioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapTCPWrapperConnection(obj *externglib.Object) *TCPWrapperConnection {
	return &TCPWrapperConnection{
		TCPConnection: TCPConnection{
			SocketConnection: SocketConnection{
				IOStream: IOStream{
					Object: obj,
				},
			},
		},
	}
}

func marshalTCPWrapperConnection(p uintptr) (interface{}, error) {
	return wrapTCPWrapperConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTCPWrapperConnection wraps base_io_stream and socket together as a
// Connection.
//
// The function takes the following parameters:
//
//    - baseIoStream to wrap.
//    - socket associated with base_io_stream.
//
// The function returns the following values:
//
//    - tcpWrapperConnection: new Connection.
//
func NewTCPWrapperConnection(baseIoStream IOStreamer, socket *Socket) *TCPWrapperConnection {
	var _arg1 *C.GIOStream         // out
	var _arg2 *C.GSocket           // out
	var _cret *C.GSocketConnection // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(baseIoStream.Native()))
	_arg2 = (*C.GSocket)(unsafe.Pointer(socket.Native()))

	_cret = C.g_tcp_wrapper_connection_new(_arg1, _arg2)
	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(socket)

	var _tcpWrapperConnection *TCPWrapperConnection // out

	_tcpWrapperConnection = wrapTCPWrapperConnection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _tcpWrapperConnection
}

// BaseIOStream gets conn's base OStream.
//
// The function returns the following values:
//
//    - ioStream conn's base OStream.
//
func (conn *TCPWrapperConnection) BaseIOStream() IOStreamer {
	var _arg0 *C.GTcpWrapperConnection // out
	var _cret *C.GIOStream             // in

	_arg0 = (*C.GTcpWrapperConnection)(unsafe.Pointer(conn.Native()))

	_cret = C.g_tcp_wrapper_connection_get_base_io_stream(_arg0)
	runtime.KeepAlive(conn)

	var _ioStream IOStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}

	return _ioStream
}
