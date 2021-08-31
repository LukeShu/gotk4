// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_monitor_get_type()), F: marshalFileMonitorrer},
	})
}

// FileMonitorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileMonitorOverrider interface {
	// Cancel cancels a file monitor.
	Cancel() bool
	Changed(file Filer, otherFile Filer, eventType FileMonitorEvent)
}

// FileMonitor monitors a file or directory for changes.
//
// To obtain a Monitor for a file or directory, use g_file_monitor(),
// g_file_monitor_file(), or g_file_monitor_directory().
//
// To get informed about changes to the file or directory you are monitoring,
// connect to the Monitor::changed signal. The signal will be emitted in the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that the monitor was created in (though if the global default main
// context is blocked, this may cause notifications to be blocked even if the
// thread-default context is still running).
type FileMonitor struct {
	*externglib.Object
}

// FileMonitorrer describes FileMonitor's abstract methods.
type FileMonitorrer interface {
	externglib.Objector

	// Cancel cancels a file monitor.
	Cancel() bool
	// EmitEvent emits the Monitor::changed signal if a change has taken place.
	EmitEvent(child Filer, otherFile Filer, eventType FileMonitorEvent)
	// IsCancelled returns whether the monitor is canceled.
	IsCancelled() bool
	// SetRateLimit sets the rate limit to which the monitor will report
	// consecutive change events to the same file.
	SetRateLimit(limitMsecs int32)
}

var _ FileMonitorrer = (*FileMonitor)(nil)

func wrapFileMonitor(obj *externglib.Object) *FileMonitor {
	return &FileMonitor{
		Object: obj,
	}
}

func marshalFileMonitorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileMonitor(obj), nil
}

// Cancel cancels a file monitor.
func (monitor *FileMonitor) Cancel() bool {
	var _arg0 *C.GFileMonitor // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_file_monitor_cancel(_arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EmitEvent emits the Monitor::changed signal if a change has taken place.
// Should be called from file monitor implementations only.
//
// Implementations are responsible to call this method from the [thread-default
// main context][g-main-context-push-thread-default] of the thread that the
// monitor was created in.
func (monitor *FileMonitor) EmitEvent(child Filer, otherFile Filer, eventType FileMonitorEvent) {
	var _arg0 *C.GFileMonitor     // out
	var _arg1 *C.GFile            // out
	var _arg2 *C.GFile            // out
	var _arg3 C.GFileMonitorEvent // out

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GFile)(unsafe.Pointer(otherFile.Native()))
	_arg3 = C.GFileMonitorEvent(eventType)

	C.g_file_monitor_emit_event(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(child)
	runtime.KeepAlive(otherFile)
	runtime.KeepAlive(eventType)
}

// IsCancelled returns whether the monitor is canceled.
func (monitor *FileMonitor) IsCancelled() bool {
	var _arg0 *C.GFileMonitor // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))

	_cret = C.g_file_monitor_is_cancelled(_arg0)
	runtime.KeepAlive(monitor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetRateLimit sets the rate limit to which the monitor will report consecutive
// change events to the same file.
func (monitor *FileMonitor) SetRateLimit(limitMsecs int32) {
	var _arg0 *C.GFileMonitor // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GFileMonitor)(unsafe.Pointer(monitor.Native()))
	_arg1 = C.gint(limitMsecs)

	C.g_file_monitor_set_rate_limit(_arg0, _arg1)
	runtime.KeepAlive(monitor)
	runtime.KeepAlive(limitMsecs)
}
