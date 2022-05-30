// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gio2_ThreadedSocketServiceClass_run(GThreadedSocketService*, GSocketConnection*, GObject*);
import "C"

// glib.Type values for gthreadedsocketservice.go.
var GTypeThreadedSocketService = coreglib.Type(C.g_threaded_socket_service_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeThreadedSocketService, F: marshalThreadedSocketService},
	})
}

// ThreadedSocketServiceOverrider contains methods that are overridable.
type ThreadedSocketServiceOverrider interface {
	// The function takes the following parameters:
	//
	//    - connection
	//    - sourceObject
	//
	// The function returns the following values:
	//
	Run(connection *SocketConnection, sourceObject *coreglib.Object) bool
}

// ThreadedSocketService is a simple subclass of Service that handles incoming
// connections by creating a worker thread and dispatching the connection to it
// by emitting the SocketService::run signal in the new thread.
//
// The signal handler may perform blocking IO and need not return until the
// connection is closed.
//
// The service is implemented using a thread pool, so there is a limited amount
// of threads available to serve incoming requests. The service automatically
// stops the Service from accepting new connections when all threads are busy.
//
// As with Service, you may connect to SocketService::run, or subclass and
// override the default handler.
type ThreadedSocketService struct {
	_ [0]func() // equal guard
	SocketService
}

var (
	_ coreglib.Objector = (*ThreadedSocketService)(nil)
)

func classInitThreadedSocketServicer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GThreadedSocketServiceClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GThreadedSocketServiceClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Run(connection *SocketConnection, sourceObject *coreglib.Object) bool
	}); ok {
		pclass.run = (*[0]byte)(C._gotk4_gio2_ThreadedSocketServiceClass_run)
	}
}

//export _gotk4_gio2_ThreadedSocketServiceClass_run
func _gotk4_gio2_ThreadedSocketServiceClass_run(arg0 *C.GThreadedSocketService, arg1 *C.GSocketConnection, arg2 *C.GObject) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Run(connection *SocketConnection, sourceObject *coreglib.Object) bool
	})

	var _connection *SocketConnection  // out
	var _sourceObject *coreglib.Object // out

	_connection = wrapSocketConnection(coreglib.Take(unsafe.Pointer(arg1)))
	_sourceObject = coreglib.Take(unsafe.Pointer(arg2))

	ok := iface.Run(_connection, _sourceObject)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapThreadedSocketService(obj *coreglib.Object) *ThreadedSocketService {
	return &ThreadedSocketService{
		SocketService: SocketService{
			SocketListener: SocketListener{
				Object: obj,
			},
		},
	}
}

func marshalThreadedSocketService(p uintptr) (interface{}, error) {
	return wrapThreadedSocketService(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
