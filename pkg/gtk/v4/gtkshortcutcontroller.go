// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GTypeShortcutController returns the GType for the type ShortcutController.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeShortcutController() coreglib.Type {
	gtype := coreglib.Type(C.gtk_shortcut_controller_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalShortcutController)
	return gtype
}

// ShortcutControllerOverrider contains methods that are overridable.
type ShortcutControllerOverrider interface {
}

// ShortcutController: GtkShortcutController is an event controller that manages
// shortcuts.
//
// Most common shortcuts are using this controller implicitly, e.g. by adding a
// mnemonic underline to a GtkLabel, or by installing a key binding using
// gtk_widget_class_add_binding(), or by adding accelerators to global actions
// using gtk_application_set_accels_for_action().
//
// But it is possible to create your own shortcut controller, and add shortcuts
// to it.
//
// GtkShortcutController implements GListModel for querying the shortcuts that
// have been added to it.
//
//
// GtkShortcutController as a GtkBuildable
//
// GtkShortcutControllers can be creates in ui files to set up shortcuts in the
// same place as the widgets.
//
// An example of a UI definition fragment with GtkShortcutController:
//
//      <object class='GtkButton'>
//        <child>
//          <object class='GtkShortcutController'>
//            <property name='scope'>managed</property>
//            <child>
//              <object class='GtkShortcut'>
//                <property name='trigger'>&amp;lt;Control&amp;gt;k</property>
//                <property name='action'>activate</property>
//              </object>
//            </child>
//          </object>
//        </child>
//      </object>
//
//
// This example creates a gtk.ActivateAction for triggering the activate signal
// of the GtkButton. See gtk.ShortcutAction.ParseString for the syntax for other
// kinds of GtkShortcutAction. See gtk.ShortcutTrigger.ParseString to learn more
// about the syntax for triggers.
type ShortcutController struct {
	_ [0]func() // equal guard
	EventController

	*coreglib.Object
	gio.ListModel
	Buildable
}

var (
	_ EventControllerer = (*ShortcutController)(nil)
	_ coreglib.Objector = (*ShortcutController)(nil)
)

func classInitShortcutControllerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapShortcutController(obj *coreglib.Object) *ShortcutController {
	return &ShortcutController{
		EventController: EventController{
			Object: obj,
		},
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalShortcutController(p uintptr) (interface{}, error) {
	return wrapShortcutController(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewShortcutController creates a new shortcut controller.
//
// The function returns the following values:
//
//    - shortcutController: newly created shortcut controller.
//
func NewShortcutController() *ShortcutController {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_shortcut_controller_new()

	var _shortcutController *ShortcutController // out

	_shortcutController = wrapShortcutController(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcutController
}

// NewShortcutControllerForModel creates a new shortcut controller that takes
// its shortcuts from the given list model.
//
// A controller created by this function does not let you add or remove
// individual shortcuts using the shortcut controller api, but you can change
// the contents of the model.
//
// The function takes the following parameters:
//
//    - model: GListModel containing shortcuts.
//
// The function returns the following values:
//
//    - shortcutController: newly created shortcut controller.
//
func NewShortcutControllerForModel(model gio.ListModeller) *ShortcutController {
	var _arg1 *C.GListModel         // out
	var _cret *C.GtkEventController // in

	_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))

	_cret = C.gtk_shortcut_controller_new_for_model(_arg1)
	runtime.KeepAlive(model)

	var _shortcutController *ShortcutController // out

	_shortcutController = wrapShortcutController(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcutController
}

// AddShortcut adds shortcut to the list of shortcuts handled by self.
//
// If this controller uses an external shortcut list, this function does
// nothing.
//
// The function takes the following parameters:
//
//    - shortcut: GtkShortcut.
//
func (self *ShortcutController) AddShortcut(shortcut *Shortcut) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 *C.GtkShortcut           // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkShortcut)(unsafe.Pointer(coreglib.InternObject(shortcut).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(shortcut).Native()))

	C.gtk_shortcut_controller_add_shortcut(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(shortcut)
}

// MnemonicsModifiers gets the mnemonics modifiers for when this controller
// activates its shortcuts.
//
// The function returns the following values:
//
//    - modifierType controller's mnemonics modifiers.
//
func (self *ShortcutController) MnemonicsModifiers() gdk.ModifierType {
	var _arg0 *C.GtkShortcutController // out
	var _cret C.GdkModifierType        // in

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_shortcut_controller_get_mnemonics_modifiers(_arg0)
	runtime.KeepAlive(self)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// Scope gets the scope for when this controller activates its shortcuts. See
// gtk_shortcut_controller_set_scope() for details.
//
// The function returns the following values:
//
//    - shortcutScope controller's scope.
//
func (self *ShortcutController) Scope() ShortcutScope {
	var _arg0 *C.GtkShortcutController // out
	var _cret C.GtkShortcutScope       // in

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_shortcut_controller_get_scope(_arg0)
	runtime.KeepAlive(self)

	var _shortcutScope ShortcutScope // out

	_shortcutScope = ShortcutScope(_cret)

	return _shortcutScope
}

// RemoveShortcut removes shortcut from the list of shortcuts handled by self.
//
// If shortcut had not been added to controller or this controller uses an
// external shortcut list, this function does nothing.
//
// The function takes the following parameters:
//
//    - shortcut: GtkShortcut.
//
func (self *ShortcutController) RemoveShortcut(shortcut *Shortcut) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 *C.GtkShortcut           // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkShortcut)(unsafe.Pointer(coreglib.InternObject(shortcut).Native()))

	C.gtk_shortcut_controller_remove_shortcut(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(shortcut)
}

// SetMnemonicsModifiers sets the controller to have the given
// mnemonics_modifiers.
//
// The mnemonics modifiers determines which modifiers need to be pressed to
// allow activation of shortcuts with mnemonics triggers.
//
// GTK normally uses the Alt modifier for mnemonics, except in PopoverMenus,
// where mnemonics can be triggered without any modifiers. It should be very
// rarely necessary to change this, and doing so is likely to interfere with
// other shortcuts.
//
// This value is only relevant for local shortcut controllers. Global and
// managed shortcut controllers will have their shortcuts activated from other
// places which have their own modifiers for activating mnemonics.
//
// The function takes the following parameters:
//
//    - modifiers: new mnemonics_modifiers to use.
//
func (self *ShortcutController) SetMnemonicsModifiers(modifiers gdk.ModifierType) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 C.GdkModifierType        // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GdkModifierType(modifiers)

	C.gtk_shortcut_controller_set_mnemonics_modifiers(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(modifiers)
}

// SetScope sets the controller to have the given scope.
//
// The scope allows shortcuts to be activated outside of the normal event
// propagation. In particular, it allows installing global keyboard shortcuts
// that can be activated even when a widget does not have focus.
//
// With GTK_SHORTCUT_SCOPE_LOCAL, shortcuts will only be activated when the
// widget has focus.
//
// The function takes the following parameters:
//
//    - scope: new scope to use.
//
func (self *ShortcutController) SetScope(scope ShortcutScope) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 C.GtkShortcutScope       // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkShortcutScope(scope)

	C.gtk_shortcut_controller_set_scope(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(scope)
}
