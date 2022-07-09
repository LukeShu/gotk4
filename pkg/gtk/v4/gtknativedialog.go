// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_NativeDialogClass_hide(void*);
// extern void _gotk4_gtk4_NativeDialogClass_response(void*, int);
// extern void _gotk4_gtk4_NativeDialogClass_show(void*);
// extern void _gotk4_gtk4_NativeDialog_ConnectResponse(gpointer, gint, guintptr);
import "C"

// GTypeNativeDialog returns the GType for the type NativeDialog.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeNativeDialog() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "NativeDialog").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalNativeDialog)
	return gtype
}

// NativeDialogOverrider contains methods that are overridable.
type NativeDialogOverrider interface {
	// Hide hides the dialog if it is visible, aborting any interaction.
	//
	// Once this is called the gtk.NativeDialog::response signal will *not* be
	// emitted until after the next call to gtk.NativeDialog.Show().
	//
	// If the dialog is not visible this does nothing.
	Hide()
	// The function takes the following parameters:
	//
	Response(responseId int32)
	// Show shows the dialog on the display.
	//
	// When the user accepts the state of the dialog the dialog will be
	// automatically hidden and the gtk.NativeDialog::response signal will be
	// emitted.
	//
	// Multiple calls while the dialog is visible will be ignored.
	Show()
}

// NativeDialog: native dialogs are platform dialogs that don't use GtkDialog.
//
// They are used in order to integrate better with a platform, by looking the
// same as other native applications and supporting platform specific features.
//
// The gtk.Dialog functions cannot be used on such objects, but we need a
// similar API in order to drive them. The GtkNativeDialog object is an API that
// allows you to do this. It allows you to set various common properties on the
// dialog, as well as show and hide it and get a gtk.NativeDialog::response
// signal when the user finished with the dialog.
//
// Note that unlike GtkDialog, GtkNativeDialog objects are not toplevel widgets,
// and GTK does not keep them alive. It is your responsibility to keep a
// reference until you are done with the object.
type NativeDialog struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*NativeDialog)(nil)
)

// NativeDialogger describes types inherited from class NativeDialog.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type NativeDialogger interface {
	coreglib.Objector
	baseNativeDialog() *NativeDialog
}

var _ NativeDialogger = (*NativeDialog)(nil)

func classInitNativeDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "NativeDialogClass")

	if _, ok := goval.(interface{ Hide() }); ok {
		o := pclass.StructFieldOffset("hide")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_NativeDialogClass_hide)
	}

	if _, ok := goval.(interface{ Response(responseId int32) }); ok {
		o := pclass.StructFieldOffset("response")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_NativeDialogClass_response)
	}

	if _, ok := goval.(interface{ Show() }); ok {
		o := pclass.StructFieldOffset("show")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_NativeDialogClass_show)
	}
}

//export _gotk4_gtk4_NativeDialogClass_hide
func _gotk4_gtk4_NativeDialogClass_hide(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Hide() })

	iface.Hide()
}

//export _gotk4_gtk4_NativeDialogClass_response
func _gotk4_gtk4_NativeDialogClass_response(arg0 *C.void, arg1 C.int) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Response(responseId int32) })

	var _responseId int32 // out

	_responseId = int32(arg1)

	iface.Response(_responseId)
}

//export _gotk4_gtk4_NativeDialogClass_show
func _gotk4_gtk4_NativeDialogClass_show(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Show() })

	iface.Show()
}

func wrapNativeDialog(obj *coreglib.Object) *NativeDialog {
	return &NativeDialog{
		Object: obj,
	}
}

func marshalNativeDialog(p uintptr) (interface{}, error) {
	return wrapNativeDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *NativeDialog) baseNativeDialog() *NativeDialog {
	return self
}

// BaseNativeDialog returns the underlying base object.
func BaseNativeDialog(obj NativeDialogger) *NativeDialog {
	return obj.baseNativeDialog()
}

//export _gotk4_gtk4_NativeDialog_ConnectResponse
func _gotk4_gtk4_NativeDialog_ConnectResponse(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(responseId int32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(responseId int32))
	}

	var _responseId int32 // out

	_responseId = int32(arg1)

	f(_responseId)
}

// ConnectResponse is emitted when the user responds to the dialog.
//
// When this is called the dialog has been hidden.
//
// If you call gtk.NativeDialog.Hide() before the user responds to the dialog
// this signal will not be emitted.
func (self *NativeDialog) ConnectResponse(f func(responseId int32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "response", false, unsafe.Pointer(C._gotk4_gtk4_NativeDialog_ConnectResponse), f)
}

// Destroy destroys a dialog.
//
// When a dialog is destroyed, it will break any references it holds to other
// objects.
//
// If it is visible it will be hidden and any underlying window system resources
// will be destroyed.
//
// Note that this does not release any reference to the object (as opposed to
// destroying a GtkWindow) because there is no reference from the windowing
// system to the GtkNativeDialog.
func (self *NativeDialog) Destroy() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("destroy", _args[:], nil)

	runtime.KeepAlive(self)
}

// Modal returns whether the dialog is modal.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is set to be modal.
//
func (self *NativeDialog) Modal() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("get_modal", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the title of the GtkNativeDialog.
//
// The function returns the following values:
//
//    - utf8 (optional): title of the dialog, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (self *NativeDialog) Title() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("get_title", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TransientFor fetches the transient parent for this window.
//
// The function returns the following values:
//
//    - window (optional): transient parent for this window, or NULL if no
//      transient parent has been set.
//
func (self *NativeDialog) TransientFor() *Window {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("get_transient_for", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _window *Window // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_window = wrapWindow(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _window
}

// Visible determines whether the dialog is visible.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is visible.
//
func (self *NativeDialog) Visible() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("get_visible", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Hide hides the dialog if it is visible, aborting any interaction.
//
// Once this is called the gtk.NativeDialog::response signal will *not* be
// emitted until after the next call to gtk.NativeDialog.Show().
//
// If the dialog is not visible this does nothing.
func (self *NativeDialog) Hide() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("hide", _args[:], nil)

	runtime.KeepAlive(self)
}

// SetModal sets a dialog modal or non-modal.
//
// Modal dialogs prevent interaction with other windows in the same application.
// To keep modal dialogs on top of main application windows, use
// gtk.NativeDialog.SetTransientFor() to make the dialog transient for the
// parent; most window managers will then disallow lowering the dialog below the
// parent.
//
// The function takes the following parameters:
//
//    - modal: whether the window is modal.
//
func (self *NativeDialog) SetModal(modal bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if modal {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("set_modal", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(modal)
}

// SetTitle sets the title of the GtkNativeDialog.
//
// The function takes the following parameters:
//
//    - title of the dialog.
//
func (self *NativeDialog) SetTitle(title string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_args[1]))

	girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("set_title", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetTransientFor: dialog windows should be set transient for the main
// application window they were spawned from.
//
// This allows window managers to e.g. keep the dialog on top of the main
// window, or center the dialog over the main window.
//
// Passing NULL for parent unsets the current transient window.
//
// The function takes the following parameters:
//
//    - parent (optional) window, or NULL.
//
func (self *NativeDialog) SetTransientFor(parent *Window) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if parent != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	}

	girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("set_transient_for", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(parent)
}

// Show shows the dialog on the display.
//
// When the user accepts the state of the dialog the dialog will be
// automatically hidden and the gtk.NativeDialog::response signal will be
// emitted.
//
// Multiple calls while the dialog is visible will be ignored.
func (self *NativeDialog) Show() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	girepository.MustFind("Gtk", "NativeDialog").InvokeMethod("show", _args[:], nil)

	runtime.KeepAlive(self)
}
