// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_mount_get_type()), F: marshalMount},
	})
}

// MountOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MountOverrider interface {
	// CanEject checks if @mount can be ejected.
	CanEject() bool
	// CanUnmount checks if @mount can be unmounted.
	CanUnmount() bool
	Changed()
	// Eject ejects a mount. This is an asynchronous operation, and is finished
	// by calling g_mount_eject_finish() with the @mount and Result data
	// returned in the @callback.
	//
	// Deprecated: since version 2.22.
	Eject(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a mount. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// Deprecated: since version 2.22.
	EjectFinish(result AsyncResult) error
	// EjectWithOperation ejects a mount. This is an asynchronous operation, and
	// is finished by calling g_mount_eject_with_operation_finish() with the
	// @mount and Result data returned in the @callback.
	EjectWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a mount. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	EjectWithOperationFinish(result AsyncResult) error
	// DefaultLocation gets the default location of @mount. The default location
	// of the given @mount is a path that reflects the main entry point for the
	// user (e.g. the home directory, or the root of the volume).
	DefaultLocation() File
	// Drive gets the drive for the @mount.
	//
	// This is a convenience method for getting the #GVolume and then using that
	// object to get the #GDrive.
	Drive() Drive
	// Icon gets the icon for @mount.
	Icon() Icon
	// Name gets the name of @mount.
	Name() string
	// Root gets the root directory on @mount.
	Root() File
	// SortKey gets the sort key for @mount, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for @mount.
	SymbolicIcon() Icon
	// UUID gets the UUID for the @mount. The reference is typically based on
	// the file system UUID for the mount in question and should be considered
	// an opaque string. Returns nil if there is no UUID available.
	UUID() string
	// Volume gets the volume for the @mount.
	Volume() Volume
	// GuessContentType tries to guess the type of content stored on @mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is an asynchronous operation (see g_mount_guess_content_type_sync()
	// for the synchronous version), and is finished by calling
	// g_mount_guess_content_type_finish() with the @mount and Result data
	// returned in the @callback.
	GuessContentType(forceRescan bool, cancellable Cancellable, callback AsyncReadyCallback)
	// GuessContentTypeFinish finishes guessing content types of @mount. If any
	// errors occurred during the operation, @error will be set to contain the
	// errors and false will be returned. In particular, you may get an
	// G_IO_ERROR_NOT_SUPPORTED if the mount does not support content guessing.
	GuessContentTypeFinish(result AsyncResult) ([]string, error)
	// GuessContentTypeSync tries to guess the type of content stored on @mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is a synchronous operation and as such may block doing IO; see
	// g_mount_guess_content_type() for the asynchronous version.
	GuessContentTypeSync(forceRescan bool, cancellable Cancellable) ([]string, error)
	PreUnmount()
	// Remount remounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_remount_finish() with the @mount and Results
	// data returned in the @callback.
	//
	// Remounting is useful when some setting affecting the operation of the
	// volume has been changed, as these may need a remount to take affect.
	// While this is semantically equivalent with unmounting and then remounting
	// not all backends might need to actually be unmounted.
	Remount(flags MountMountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// RemountFinish finishes remounting a mount. If any errors occurred during
	// the operation, @error will be set to contain the errors and false will be
	// returned.
	RemountFinish(result AsyncResult) error
	// Unmount unmounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_unmount_finish() with the @mount and Result
	// data returned in the @callback.
	//
	// Deprecated: since version 2.22.
	Unmount(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// UnmountFinish finishes unmounting a mount. If any errors occurred during
	// the operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// Deprecated: since version 2.22.
	UnmountFinish(result AsyncResult) error
	// UnmountWithOperation unmounts a mount. This is an asynchronous operation,
	// and is finished by calling g_mount_unmount_with_operation_finish() with
	// the @mount and Result data returned in the @callback.
	UnmountWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// UnmountWithOperationFinish finishes unmounting a mount. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	UnmountWithOperationFinish(result AsyncResult) error
	Unmounted()
}

// Mount: the #GMount interface represents user-visible mounts. Note, when
// porting from GnomeVFS, #GMount is the moral equivalent of VFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in quotes
// because it's not the same as a unix mount, it might be a gvfs mount, but you
// can still access the files on it if you use GIO. Might or might not be
// related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To unmount
// a #GMount instance, first call g_mount_unmount_with_operation() with (at
// least) the #GMount instance and a ReadyCallback. The callback will be fired
// when the operation has resolved (either with success or failure), and a
// Result structure will be passed to the callback. That callback should then
// call g_mount_unmount_with_operation_finish() with the #GMount and the Result
// data to see if the operation was completed successfully. If an @error is
// present when g_mount_unmount_with_operation_finish() is called, then it will
// be filled with any error information.
type Mount interface {
	gextras.Objector

	// CanEject checks if @mount can be ejected.
	CanEject() bool
	// CanUnmount checks if @mount can be unmounted.
	CanUnmount() bool
	// Eject ejects a mount. This is an asynchronous operation, and is finished
	// by calling g_mount_eject_finish() with the @mount and Result data
	// returned in the @callback.
	//
	// Deprecated: since version 2.22.
	Eject(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a mount. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// Deprecated: since version 2.22.
	EjectFinish(result AsyncResult) error
	// EjectWithOperation ejects a mount. This is an asynchronous operation, and
	// is finished by calling g_mount_eject_with_operation_finish() with the
	// @mount and Result data returned in the @callback.
	EjectWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a mount. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	EjectWithOperationFinish(result AsyncResult) error
	// DefaultLocation gets the default location of @mount. The default location
	// of the given @mount is a path that reflects the main entry point for the
	// user (e.g. the home directory, or the root of the volume).
	DefaultLocation() File
	// Drive gets the drive for the @mount.
	//
	// This is a convenience method for getting the #GVolume and then using that
	// object to get the #GDrive.
	Drive() Drive
	// Icon gets the icon for @mount.
	Icon() Icon
	// Name gets the name of @mount.
	Name() string
	// Root gets the root directory on @mount.
	Root() File
	// SortKey gets the sort key for @mount, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for @mount.
	SymbolicIcon() Icon
	// UUID gets the UUID for the @mount. The reference is typically based on
	// the file system UUID for the mount in question and should be considered
	// an opaque string. Returns nil if there is no UUID available.
	UUID() string
	// Volume gets the volume for the @mount.
	Volume() Volume
	// GuessContentType tries to guess the type of content stored on @mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is an asynchronous operation (see g_mount_guess_content_type_sync()
	// for the synchronous version), and is finished by calling
	// g_mount_guess_content_type_finish() with the @mount and Result data
	// returned in the @callback.
	GuessContentType(forceRescan bool, cancellable Cancellable, callback AsyncReadyCallback)
	// GuessContentTypeFinish finishes guessing content types of @mount. If any
	// errors occurred during the operation, @error will be set to contain the
	// errors and false will be returned. In particular, you may get an
	// G_IO_ERROR_NOT_SUPPORTED if the mount does not support content guessing.
	GuessContentTypeFinish(result AsyncResult) ([]string, error)
	// GuessContentTypeSync tries to guess the type of content stored on @mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is a synchronous operation and as such may block doing IO; see
	// g_mount_guess_content_type() for the asynchronous version.
	GuessContentTypeSync(forceRescan bool, cancellable Cancellable) ([]string, error)
	// IsShadowed determines if @mount is shadowed. Applications or libraries
	// should avoid displaying @mount in the user interface if it is shadowed.
	//
	// A mount is said to be shadowed if there exists one or more user visible
	// objects (currently #GMount objects) with a root that is inside the root
	// of @mount.
	//
	// One application of shadow mounts is when exposing a single file system
	// that is used to address several logical volumes. In this situation, a
	// Monitor implementation would create two #GVolume objects (for example,
	// one for the camera functionality of the device and one for a SD card
	// reader on the device) with activation URIs
	// `gphoto2://[usb:001,002]/store1/` and `gphoto2://[usb:001,002]/store2/`.
	// When the underlying mount (with root `gphoto2://[usb:001,002]/`) is
	// mounted, said Monitor implementation would create two #GMount objects
	// (each with their root matching the corresponding volume activation root)
	// that would shadow the original mount.
	//
	// The proxy monitor in GVfs 2.26 and later, automatically creates and
	// manage shadow mounts (and shadows the underlying mount) if the activation
	// root on a #GVolume is set.
	IsShadowed() bool
	// Remount remounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_remount_finish() with the @mount and Results
	// data returned in the @callback.
	//
	// Remounting is useful when some setting affecting the operation of the
	// volume has been changed, as these may need a remount to take affect.
	// While this is semantically equivalent with unmounting and then remounting
	// not all backends might need to actually be unmounted.
	Remount(flags MountMountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// RemountFinish finishes remounting a mount. If any errors occurred during
	// the operation, @error will be set to contain the errors and false will be
	// returned.
	RemountFinish(result AsyncResult) error
	// Shadow increments the shadow count on @mount. Usually used by Monitor
	// implementations when creating a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Shadow()
	// Unmount unmounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_unmount_finish() with the @mount and Result
	// data returned in the @callback.
	//
	// Deprecated: since version 2.22.
	Unmount(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// UnmountFinish finishes unmounting a mount. If any errors occurred during
	// the operation, @error will be set to contain the errors and false will be
	// returned.
	//
	// Deprecated: since version 2.22.
	UnmountFinish(result AsyncResult) error
	// UnmountWithOperation unmounts a mount. This is an asynchronous operation,
	// and is finished by calling g_mount_unmount_with_operation_finish() with
	// the @mount and Result data returned in the @callback.
	UnmountWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// UnmountWithOperationFinish finishes unmounting a mount. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	UnmountWithOperationFinish(result AsyncResult) error
	// Unshadow decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Unshadow()
}

// MountInterface implements the Mount interface.
type MountInterface struct {
	*externglib.Object
}

var _ Mount = (*MountInterface)(nil)

func wrapMount(obj *externglib.Object) Mount {
	return &MountInterface{
		Object: obj,
	}
}

func marshalMount(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMount(obj), nil
}

// CanEject checks if @mount can be ejected.
func (m *MountInterface) CanEject() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_can_eject(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUnmount checks if @mount can be unmounted.
func (m *MountInterface) CanUnmount() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_can_unmount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Eject ejects a mount. This is an asynchronous operation, and is finished by
// calling g_mount_eject_finish() with the @mount and Result data returned in
// the @callback.
//
// Deprecated: since version 2.22.
func (m *MountInterface) Eject(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_mount_eject(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// EjectFinish finishes ejecting a mount. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
//
// Deprecated: since version 2.22.
func (m *MountInterface) EjectFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_eject_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// EjectWithOperation ejects a mount. This is an asynchronous operation, and is
// finished by calling g_mount_eject_with_operation_finish() with the @mount and
// Result data returned in the @callback.
func (m *MountInterface) EjectWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_mount_eject_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// EjectWithOperationFinish finishes ejecting a mount. If any errors occurred
// during the operation, @error will be set to contain the errors and false will
// be returned.
func (m *MountInterface) EjectWithOperationFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_eject_with_operation_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// DefaultLocation gets the default location of @mount. The default location of
// the given @mount is a path that reflects the main entry point for the user
// (e.g. the home directory, or the root of the volume).
func (m *MountInterface) DefaultLocation() File {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_default_location(_arg0)

	var _file File // out

	_file = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(File)

	return _file
}

// Drive gets the drive for the @mount.
//
// This is a convenience method for getting the #GVolume and then using that
// object to get the #GDrive.
func (m *MountInterface) Drive() Drive {
	var _arg0 *C.GMount // out
	var _cret *C.GDrive // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_drive(_arg0)

	var _drive Drive // out

	_drive = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Drive)

	return _drive
}

// Icon gets the icon for @mount.
func (m *MountInterface) Icon() Icon {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

// Name gets the name of @mount.
func (m *MountInterface) Name() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Root gets the root directory on @mount.
func (m *MountInterface) Root() File {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_root(_arg0)

	var _file File // out

	_file = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(File)

	return _file
}

// SortKey gets the sort key for @mount, if any.
func (m *MountInterface) SortKey() string {
	var _arg0 *C.GMount // out
	var _cret *C.gchar  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_sort_key(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SymbolicIcon gets the symbolic icon for @mount.
func (m *MountInterface) SymbolicIcon() Icon {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_symbolic_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

// UUID gets the UUID for the @mount. The reference is typically based on the
// file system UUID for the mount in question and should be considered an opaque
// string. Returns nil if there is no UUID available.
func (m *MountInterface) UUID() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_uuid(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Volume gets the volume for the @mount.
func (m *MountInterface) Volume() Volume {
	var _arg0 *C.GMount  // out
	var _cret *C.GVolume // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_volume(_arg0)

	var _volume Volume // out

	_volume = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Volume)

	return _volume
}

// GuessContentType tries to guess the type of content stored on @mount. Returns
// one or more textual identifiers of well-known content types (typically
// prefixed with "x-content/"), e.g. x-content/image-dcf for camera memory
// cards. See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is an asynchronous operation (see g_mount_guess_content_type_sync() for
// the synchronous version), and is finished by calling
// g_mount_guess_content_type_finish() with the @mount and Result data returned
// in the @callback.
func (m *MountInterface) GuessContentType(forceRescan bool, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.gboolean            // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	if forceRescan {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_mount_guess_content_type(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// GuessContentTypeFinish finishes guessing content types of @mount. If any
// errors occurred during the operation, @error will be set to contain the
// errors and false will be returned. In particular, you may get an
// G_IO_ERROR_NOT_SUPPORTED if the mount does not support content guessing.
func (m *MountInterface) GuessContentTypeFinish(result AsyncResult) ([]string, error) {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cret **C.gchar
	var _cerr *C.GError // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_mount_guess_content_type_finish(_arg0, _arg1, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

// GuessContentTypeSync tries to guess the type of content stored on @mount.
// Returns one or more textual identifiers of well-known content types
// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for camera
// memory cards. See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is a synchronous operation and as such may block doing IO; see
// g_mount_guess_content_type() for the asynchronous version.
func (m *MountInterface) GuessContentTypeSync(forceRescan bool, cancellable Cancellable) ([]string, error) {
	var _arg0 *C.GMount       // out
	var _arg1 C.gboolean      // out
	var _arg2 *C.GCancellable // out
	var _cret **C.gchar
	var _cerr *C.GError // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	if forceRescan {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_mount_guess_content_type_sync(_arg0, _arg1, _arg2, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

// IsShadowed determines if @mount is shadowed. Applications or libraries should
// avoid displaying @mount in the user interface if it is shadowed.
//
// A mount is said to be shadowed if there exists one or more user visible
// objects (currently #GMount objects) with a root that is inside the root of
// @mount.
//
// One application of shadow mounts is when exposing a single file system that
// is used to address several logical volumes. In this situation, a Monitor
// implementation would create two #GVolume objects (for example, one for the
// camera functionality of the device and one for a SD card reader on the
// device) with activation URIs `gphoto2://[usb:001,002]/store1/` and
// `gphoto2://[usb:001,002]/store2/`. When the underlying mount (with root
// `gphoto2://[usb:001,002]/`) is mounted, said Monitor implementation would
// create two #GMount objects (each with their root matching the corresponding
// volume activation root) that would shadow the original mount.
//
// The proxy monitor in GVfs 2.26 and later, automatically creates and manage
// shadow mounts (and shadows the underlying mount) if the activation root on a
// #GVolume is set.
func (m *MountInterface) IsShadowed() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_is_shadowed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remount remounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_remount_finish() with the @mount and Results data returned
// in the @callback.
//
// Remounting is useful when some setting affecting the operation of the volume
// has been changed, as these may need a remount to take affect. While this is
// semantically equivalent with unmounting and then remounting not all backends
// might need to actually be unmounted.
func (m *MountInterface) Remount(flags MountMountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountMountFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountMountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_mount_remount(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// RemountFinish finishes remounting a mount. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
func (m *MountInterface) RemountFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_remount_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Shadow increments the shadow count on @mount. Usually used by Monitor
// implementations when creating a shadow mount for @mount, see
// g_mount_is_shadowed() for more information. The caller will need to emit the
// #GMount::changed signal on @mount manually.
func (m *MountInterface) Shadow() {
	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_shadow(_arg0)
}

// Unmount unmounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_unmount_finish() with the @mount and Result data returned
// in the @callback.
//
// Deprecated: since version 2.22.
func (m *MountInterface) Unmount(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_mount_unmount(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// UnmountFinish finishes unmounting a mount. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
//
// Deprecated: since version 2.22.
func (m *MountInterface) UnmountFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_unmount_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// UnmountWithOperation unmounts a mount. This is an asynchronous operation, and
// is finished by calling g_mount_unmount_with_operation_finish() with the
// @mount and Result data returned in the @callback.
func (m *MountInterface) UnmountWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_mount_unmount_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// UnmountWithOperationFinish finishes unmounting a mount. If any errors
// occurred during the operation, @error will be set to contain the errors and
// false will be returned.
func (m *MountInterface) UnmountWithOperationFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_unmount_with_operation_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Unshadow decrements the shadow count on @mount. Usually used by Monitor
// implementations when destroying a shadow mount for @mount, see
// g_mount_is_shadowed() for more information. The caller will need to emit the
// #GMount::changed signal on @mount manually.
func (m *MountInterface) Unshadow() {
	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_unshadow(_arg0)
}
