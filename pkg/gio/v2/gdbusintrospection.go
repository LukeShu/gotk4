// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
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
		{T: externglib.Type(C.g_dbus_annotation_info_get_type()), F: marshalDBusAnnotationInfo},
		{T: externglib.Type(C.g_dbus_arg_info_get_type()), F: marshalDBusArgInfo},
		{T: externglib.Type(C.g_dbus_interface_info_get_type()), F: marshalDBusInterfaceInfo},
		{T: externglib.Type(C.g_dbus_method_info_get_type()), F: marshalDBusMethodInfo},
		{T: externglib.Type(C.g_dbus_node_info_get_type()), F: marshalDBusNodeInfo},
		{T: externglib.Type(C.g_dbus_property_info_get_type()), F: marshalDBusPropertyInfo},
		{T: externglib.Type(C.g_dbus_signal_info_get_type()), F: marshalDBusSignalInfo},
	})
}

// DBusAnnotationInfo: information about an annotation.
type DBusAnnotationInfo struct {
	native C.GDBusAnnotationInfo
}

// WrapDBusAnnotationInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusAnnotationInfo(ptr unsafe.Pointer) *DBusAnnotationInfo {
	return (*DBusAnnotationInfo)(ptr)
}

func marshalDBusAnnotationInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusAnnotationInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusAnnotationInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount: the reference count or -1 if statically allocated.
func (d *DBusAnnotationInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Key: the name of the annotation, e.g. "org.freedesktop.DBus.Deprecated".
func (d *DBusAnnotationInfo) Key() string {
	var v string // out
	v = C.GoString(d.native.key)
	return v
}

// Value: the value of the annotation.
func (d *DBusAnnotationInfo) Value() string {
	var v string // out
	v = C.GoString(d.native.value)
	return v
}

// Annotations: pointer to a nil-terminated array of pointers to
// BusAnnotationInfo structures or nil if there are no annotations.
func (d *DBusAnnotationInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusAnnotationInfo) {
				C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusAnnotationInfo) ref() *DBusAnnotationInfo {
	var _arg0 *C.GDBusAnnotationInfo // out
	var _cret *C.GDBusAnnotationInfo // in

	_arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_annotation_info_ref(_arg0)

	var _dBusAnnotationInfo *DBusAnnotationInfo // out

	_dBusAnnotationInfo = (*DBusAnnotationInfo)(unsafe.Pointer(_cret))
	C.g_dbus_annotation_info_ref(_cret)
	runtime.SetFinalizer(_dBusAnnotationInfo, func(v *DBusAnnotationInfo) {
		C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
	})

	return _dBusAnnotationInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusAnnotationInfo) unref() {
	var _arg0 *C.GDBusAnnotationInfo // out

	_arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(i))

	C.g_dbus_annotation_info_unref(_arg0)
}

// DBusArgInfo: information about an argument for a method or a signal.
type DBusArgInfo struct {
	native C.GDBusArgInfo
}

// WrapDBusArgInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusArgInfo(ptr unsafe.Pointer) *DBusArgInfo {
	return (*DBusArgInfo)(ptr)
}

func marshalDBusArgInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusArgInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusArgInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount: the reference count or -1 if statically allocated.
func (d *DBusArgInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name of the argument, e.g. @unix_user_id.
func (d *DBusArgInfo) Name() string {
	var v string // out
	v = C.GoString(d.native.name)
	return v
}

// Signature d-Bus signature of the argument (a single complete type).
func (d *DBusArgInfo) Signature() string {
	var v string // out
	v = C.GoString(d.native.signature)
	return v
}

// Annotations: pointer to a nil-terminated array of pointers to
// BusAnnotationInfo structures or nil if there are no annotations.
func (d *DBusArgInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusAnnotationInfo) {
				C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusArgInfo) ref() *DBusArgInfo {
	var _arg0 *C.GDBusArgInfo // out
	var _cret *C.GDBusArgInfo // in

	_arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_arg_info_ref(_arg0)

	var _dBusArgInfo *DBusArgInfo // out

	_dBusArgInfo = (*DBusArgInfo)(unsafe.Pointer(_cret))
	C.g_dbus_arg_info_ref(_cret)
	runtime.SetFinalizer(_dBusArgInfo, func(v *DBusArgInfo) {
		C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(unsafe.Pointer(v)))
	})

	return _dBusArgInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusArgInfo) unref() {
	var _arg0 *C.GDBusArgInfo // out

	_arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(i))

	C.g_dbus_arg_info_unref(_arg0)
}

// DBusInterfaceInfo: information about a D-Bus interface.
type DBusInterfaceInfo struct {
	native C.GDBusInterfaceInfo
}

// WrapDBusInterfaceInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusInterfaceInfo(ptr unsafe.Pointer) *DBusInterfaceInfo {
	return (*DBusInterfaceInfo)(ptr)
}

func marshalDBusInterfaceInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusInterfaceInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusInterfaceInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount: the reference count or -1 if statically allocated.
func (d *DBusInterfaceInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: the name of the D-Bus interface, e.g.
// "org.freedesktop.DBus.Properties".
func (d *DBusInterfaceInfo) Name() string {
	var v string // out
	v = C.GoString(d.native.name)
	return v
}

// Methods: pointer to a nil-terminated array of pointers to BusMethodInfo
// structures or nil if there are no methods.
func (d *DBusInterfaceInfo) Methods() []*DBusMethodInfo {
	var v []*DBusMethodInfo
	{
		var i int
		var z *C.GDBusMethodInfo
		for p := d.native.methods; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.methods, i)
		v = make([]*DBusMethodInfo, i)
		for i := range src {
			v[i] = (*DBusMethodInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_method_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusMethodInfo) {
				C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Signals: pointer to a nil-terminated array of pointers to BusSignalInfo
// structures or nil if there are no signals.
func (d *DBusInterfaceInfo) Signals() []*DBusSignalInfo {
	var v []*DBusSignalInfo
	{
		var i int
		var z *C.GDBusSignalInfo
		for p := d.native.signals; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.signals, i)
		v = make([]*DBusSignalInfo, i)
		for i := range src {
			v[i] = (*DBusSignalInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_signal_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusSignalInfo) {
				C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Properties: pointer to a nil-terminated array of pointers to BusPropertyInfo
// structures or nil if there are no properties.
func (d *DBusInterfaceInfo) Properties() []*DBusPropertyInfo {
	var v []*DBusPropertyInfo
	{
		var i int
		var z *C.GDBusPropertyInfo
		for p := d.native.properties; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.properties, i)
		v = make([]*DBusPropertyInfo, i)
		for i := range src {
			v[i] = (*DBusPropertyInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_property_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusPropertyInfo) {
				C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Annotations: pointer to a nil-terminated array of pointers to
// BusAnnotationInfo structures or nil if there are no annotations.
func (d *DBusInterfaceInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusAnnotationInfo) {
				C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// CacheBuild builds a lookup-cache to speed up
// g_dbus_interface_info_lookup_method(), g_dbus_interface_info_lookup_signal()
// and g_dbus_interface_info_lookup_property().
//
// If this has already been called with @info, the existing cache is used and
// its use count is increased.
//
// Note that @info cannot be modified until
// g_dbus_interface_info_cache_release() is called.
func (i *DBusInterfaceInfo) CacheBuild() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	C.g_dbus_interface_info_cache_build(_arg0)
}

// CacheRelease decrements the usage count for the cache for @info built by
// g_dbus_interface_info_cache_build() (if any) and frees the resources used by
// the cache if the usage count drops to zero.
func (i *DBusInterfaceInfo) CacheRelease() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	C.g_dbus_interface_info_cache_release(_arg0)
}

// LookupMethod looks up information about a method.
//
// The cost of this function is O(n) in number of methods unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusMethodInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_method(_arg0, _arg1)

	var _dBusMethodInfo *DBusMethodInfo // out

	_dBusMethodInfo = (*DBusMethodInfo)(unsafe.Pointer(_cret))
	C.g_dbus_method_info_ref(_cret)
	runtime.SetFinalizer(_dBusMethodInfo, func(v *DBusMethodInfo) {
		C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(unsafe.Pointer(v)))
	})

	return _dBusMethodInfo
}

// LookupProperty looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusPropertyInfo  // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_property(_arg0, _arg1)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	_dBusPropertyInfo = (*DBusPropertyInfo)(unsafe.Pointer(_cret))
	C.g_dbus_property_info_ref(_cret)
	runtime.SetFinalizer(_dBusPropertyInfo, func(v *DBusPropertyInfo) {
		C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(unsafe.Pointer(v)))
	})

	return _dBusPropertyInfo
}

// LookupSignal looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusSignalInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_signal(_arg0, _arg1)

	var _dBusSignalInfo *DBusSignalInfo // out

	_dBusSignalInfo = (*DBusSignalInfo)(unsafe.Pointer(_cret))
	C.g_dbus_signal_info_ref(_cret)
	runtime.SetFinalizer(_dBusSignalInfo, func(v *DBusSignalInfo) {
		C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(unsafe.Pointer(v)))
	})

	return _dBusSignalInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusInterfaceInfo) ref() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_interface_info_ref(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusInterfaceInfo) unref() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i))

	C.g_dbus_interface_info_unref(_arg0)
}

// DBusMethodInfo: information about a method on an D-Bus interface.
type DBusMethodInfo struct {
	native C.GDBusMethodInfo
}

// WrapDBusMethodInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusMethodInfo(ptr unsafe.Pointer) *DBusMethodInfo {
	return (*DBusMethodInfo)(ptr)
}

func marshalDBusMethodInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusMethodInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusMethodInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount: the reference count or -1 if statically allocated.
func (d *DBusMethodInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: the name of the D-Bus method, e.g. @RequestName.
func (d *DBusMethodInfo) Name() string {
	var v string // out
	v = C.GoString(d.native.name)
	return v
}

// InArgs: pointer to a nil-terminated array of pointers to BusArgInfo
// structures or nil if there are no in arguments.
func (d *DBusMethodInfo) InArgs() []*DBusArgInfo {
	var v []*DBusArgInfo
	{
		var i int
		var z *C.GDBusArgInfo
		for p := d.native.in_args; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.in_args, i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_arg_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusArgInfo) {
				C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// OutArgs: pointer to a nil-terminated array of pointers to BusArgInfo
// structures or nil if there are no out arguments.
func (d *DBusMethodInfo) OutArgs() []*DBusArgInfo {
	var v []*DBusArgInfo
	{
		var i int
		var z *C.GDBusArgInfo
		for p := d.native.out_args; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.out_args, i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_arg_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusArgInfo) {
				C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Annotations: pointer to a nil-terminated array of pointers to
// BusAnnotationInfo structures or nil if there are no annotations.
func (d *DBusMethodInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusAnnotationInfo) {
				C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusMethodInfo) ref() *DBusMethodInfo {
	var _arg0 *C.GDBusMethodInfo // out
	var _cret *C.GDBusMethodInfo // in

	_arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_method_info_ref(_arg0)

	var _dBusMethodInfo *DBusMethodInfo // out

	_dBusMethodInfo = (*DBusMethodInfo)(unsafe.Pointer(_cret))
	C.g_dbus_method_info_ref(_cret)
	runtime.SetFinalizer(_dBusMethodInfo, func(v *DBusMethodInfo) {
		C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(unsafe.Pointer(v)))
	})

	return _dBusMethodInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusMethodInfo) unref() {
	var _arg0 *C.GDBusMethodInfo // out

	_arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(i))

	C.g_dbus_method_info_unref(_arg0)
}

// DBusNodeInfo: information about nodes in a remote object hierarchy.
type DBusNodeInfo struct {
	native C.GDBusNodeInfo
}

// WrapDBusNodeInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusNodeInfo(ptr unsafe.Pointer) *DBusNodeInfo {
	return (*DBusNodeInfo)(ptr)
}

func marshalDBusNodeInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusNodeInfo)(unsafe.Pointer(b)), nil
}

// NewDBusNodeInfoForXML constructs a struct DBusNodeInfo.
func NewDBusNodeInfoForXML(xmlData string) (*DBusNodeInfo, error) {
	var _arg1 *C.gchar         // out
	var _cret *C.GDBusNodeInfo // in
	var _cerr *C.GError        // in

	_arg1 = (*C.gchar)(C.CString(xmlData))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_new_for_xml(_arg1, &_cerr)

	var _dBusNodeInfo *DBusNodeInfo // out
	var _goerr error                // out

	_dBusNodeInfo = (*DBusNodeInfo)(unsafe.Pointer(_cret))
	C.g_dbus_node_info_ref(_cret)
	runtime.SetFinalizer(_dBusNodeInfo, func(v *DBusNodeInfo) {
		C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusNodeInfo, _goerr
}

// Native returns the underlying C source pointer.
func (d *DBusNodeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount: the reference count or -1 if statically allocated.
func (d *DBusNodeInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Path: the path of the node or nil if omitted. Note that this may be a
// relative path. See the D-Bus specification for more details.
func (d *DBusNodeInfo) Path() string {
	var v string // out
	v = C.GoString(d.native.path)
	return v
}

// Interfaces: pointer to a nil-terminated array of pointers to BusInterfaceInfo
// structures or nil if there are no interfaces.
func (d *DBusNodeInfo) Interfaces() []*DBusInterfaceInfo {
	var v []*DBusInterfaceInfo
	{
		var i int
		var z *C.GDBusInterfaceInfo
		for p := d.native.interfaces; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.interfaces, i)
		v = make([]*DBusInterfaceInfo, i)
		for i := range src {
			v[i] = (*DBusInterfaceInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_interface_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusInterfaceInfo) {
				C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Nodes: pointer to a nil-terminated array of pointers to BusNodeInfo
// structures or nil if there are no nodes.
func (d *DBusNodeInfo) Nodes() []*DBusNodeInfo {
	var v []*DBusNodeInfo
	{
		var i int
		var z *C.GDBusNodeInfo
		for p := d.native.nodes; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.nodes, i)
		v = make([]*DBusNodeInfo, i)
		for i := range src {
			v[i] = (*DBusNodeInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_node_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusNodeInfo) {
				C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Annotations: pointer to a nil-terminated array of pointers to
// BusAnnotationInfo structures or nil if there are no annotations.
func (d *DBusNodeInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusAnnotationInfo) {
				C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// LookupInterface looks up information about an interface.
//
// The cost of this function is O(n) in number of interfaces.
func (i *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	var _arg0 *C.GDBusNodeInfo      // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_lookup_interface(_arg0, _arg1)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusNodeInfo) ref() *DBusNodeInfo {
	var _arg0 *C.GDBusNodeInfo // out
	var _cret *C.GDBusNodeInfo // in

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_node_info_ref(_arg0)

	var _dBusNodeInfo *DBusNodeInfo // out

	_dBusNodeInfo = (*DBusNodeInfo)(unsafe.Pointer(_cret))
	C.g_dbus_node_info_ref(_cret)
	runtime.SetFinalizer(_dBusNodeInfo, func(v *DBusNodeInfo) {
		C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(unsafe.Pointer(v)))
	})

	return _dBusNodeInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusNodeInfo) unref() {
	var _arg0 *C.GDBusNodeInfo // out

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i))

	C.g_dbus_node_info_unref(_arg0)
}

// DBusPropertyInfo: information about a D-Bus property on a D-Bus interface.
type DBusPropertyInfo struct {
	native C.GDBusPropertyInfo
}

// WrapDBusPropertyInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusPropertyInfo(ptr unsafe.Pointer) *DBusPropertyInfo {
	return (*DBusPropertyInfo)(ptr)
}

func marshalDBusPropertyInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusPropertyInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusPropertyInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount: the reference count or -1 if statically allocated.
func (d *DBusPropertyInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: the name of the D-Bus property, e.g. "SupportedFilesystems".
func (d *DBusPropertyInfo) Name() string {
	var v string // out
	v = C.GoString(d.native.name)
	return v
}

// Signature: the D-Bus signature of the property (a single complete type).
func (d *DBusPropertyInfo) Signature() string {
	var v string // out
	v = C.GoString(d.native.signature)
	return v
}

// Flags access control flags for the property.
func (d *DBusPropertyInfo) Flags() DBusPropertyInfoFlags {
	var v DBusPropertyInfoFlags // out
	v = (DBusPropertyInfoFlags)(d.native.flags)
	return v
}

// Annotations: pointer to a nil-terminated array of pointers to
// BusAnnotationInfo structures or nil if there are no annotations.
func (d *DBusPropertyInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusAnnotationInfo) {
				C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusPropertyInfo) ref() *DBusPropertyInfo {
	var _arg0 *C.GDBusPropertyInfo // out
	var _cret *C.GDBusPropertyInfo // in

	_arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_property_info_ref(_arg0)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	_dBusPropertyInfo = (*DBusPropertyInfo)(unsafe.Pointer(_cret))
	C.g_dbus_property_info_ref(_cret)
	runtime.SetFinalizer(_dBusPropertyInfo, func(v *DBusPropertyInfo) {
		C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(unsafe.Pointer(v)))
	})

	return _dBusPropertyInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusPropertyInfo) unref() {
	var _arg0 *C.GDBusPropertyInfo // out

	_arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(i))

	C.g_dbus_property_info_unref(_arg0)
}

// DBusSignalInfo: information about a signal on a D-Bus interface.
type DBusSignalInfo struct {
	native C.GDBusSignalInfo
}

// WrapDBusSignalInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusSignalInfo(ptr unsafe.Pointer) *DBusSignalInfo {
	return (*DBusSignalInfo)(ptr)
}

func marshalDBusSignalInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusSignalInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusSignalInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount: the reference count or -1 if statically allocated.
func (d *DBusSignalInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: the name of the D-Bus signal, e.g. "NameOwnerChanged".
func (d *DBusSignalInfo) Name() string {
	var v string // out
	v = C.GoString(d.native.name)
	return v
}

// Args: pointer to a nil-terminated array of pointers to BusArgInfo structures
// or nil if there are no arguments.
func (d *DBusSignalInfo) Args() []*DBusArgInfo {
	var v []*DBusArgInfo
	{
		var i int
		var z *C.GDBusArgInfo
		for p := d.native.args; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.args, i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_arg_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusArgInfo) {
				C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Annotations: pointer to a nil-terminated array of pointers to
// BusAnnotationInfo structures or nil if there are no annotations.
func (d *DBusSignalInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(unsafe.Pointer(src[i]))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(v[i], func(v *DBusAnnotationInfo) {
				C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
			})
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusSignalInfo) ref() *DBusSignalInfo {
	var _arg0 *C.GDBusSignalInfo // out
	var _cret *C.GDBusSignalInfo // in

	_arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i))

	_cret = C.g_dbus_signal_info_ref(_arg0)

	var _dBusSignalInfo *DBusSignalInfo // out

	_dBusSignalInfo = (*DBusSignalInfo)(unsafe.Pointer(_cret))
	C.g_dbus_signal_info_ref(_cret)
	runtime.SetFinalizer(_dBusSignalInfo, func(v *DBusSignalInfo) {
		C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(unsafe.Pointer(v)))
	})

	return _dBusSignalInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusSignalInfo) unref() {
	var _arg0 *C.GDBusSignalInfo // out

	_arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i))

	C.g_dbus_signal_info_unref(_arg0)
}
