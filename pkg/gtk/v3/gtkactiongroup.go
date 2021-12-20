// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// gchar* _gotk4_gtk3_TranslateFunc(gchar*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_action_group_get_type()), F: marshalActionGrouper},
	})
}

// ActionGroupOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionGroupOverrider interface {
	// Action looks up an action in the action group by name.
	//
	// Deprecated: since version 3.10.
	//
	// The function takes the following parameters:
	//
	//    - actionName: name of the action.
	//
	// The function returns the following values:
	//
	//    - action: action, or NULL if no action by that name exists.
	//
	Action(actionName string) *Action
}

// ActionGroup actions are organised into groups. An action group is essentially
// a map from names to Action objects.
//
// All actions that would make sense to use in a particular context should be in
// a single group. Multiple action groups may be used for a particular user
// interface. In fact, it is expected that most nontrivial applications will
// make use of multiple groups. For example, in an application that can edit
// multiple documents, one group holding global actions (e.g. quit, about, new),
// and one group per document holding actions that act on that document (eg.
// save, cut/copy/paste, etc). Each window’s menus would be constructed from a
// combination of two action groups.
//
//
// Accelerators
//
// Accelerators are handled by the GTK+ accelerator map. All actions are
// assigned an accelerator path (which normally has the form
// <Actions>/group-name/action-name) and a shortcut is associated with this
// accelerator path. All menuitems and toolitems take on this accelerator path.
// The GTK+ accelerator map code makes sure that the correct shortcut is
// displayed next to the menu item.
//
//
// GtkActionGroup as GtkBuildable
//
// The ActionGroup implementation of the Buildable interface accepts Action
// objects as <child> elements in UI definitions.
//
// Note that it is probably more common to define actions and action groups in
// the code, since they are directly related to what the code can do.
//
// The GtkActionGroup implementation of the GtkBuildable interface supports a
// custom <accelerator> element, which has attributes named “key“ and
// “modifiers“ and allows to specify accelerators. This is similar to the
// <accelerator> element of Widget, the main difference is that it doesn’t allow
// you to specify a signal.
//
// A Dialog UI definition fragment. ##
//
//    <object class="GtkActionGroup" id="actiongroup">
//      <child>
//          <object class="GtkAction" id="About">
//              <property name="name">About</property>
//              <property name="stock_id">gtk-about</property>
//              <signal handler="about_activate" name="activate"/>
//          </object>
//          <accelerator key="F1" modifiers="GDK_CONTROL_MASK | GDK_SHIFT_MASK"/>
//      </child>
//    </object>.
type ActionGroup struct {
	_ [0]func() // equal guard
	*externglib.Object

	Buildable
}

var (
	_ externglib.Objector = (*ActionGroup)(nil)
)

func wrapActionGroup(obj *externglib.Object) *ActionGroup {
	return &ActionGroup{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalActionGrouper(p uintptr) (interface{}, error) {
	return wrapActionGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectConnectProxy signal is emitted after connecting a proxy to an action
// in the group. Note that the proxy may have been connected to a different
// action before.
//
// This is intended for simple customizations for which a custom action class
// would be too clumsy, e.g. showing tooltips for menuitems in the statusbar.
//
// UIManager proxies the signal and provides global notification just before any
// action is connected to a proxy, which is probably more convenient to use.
func (actionGroup *ActionGroup) ConnectConnectProxy(f func(action Action, proxy Widgetter)) externglib.SignalHandle {
	return actionGroup.Connect("connect-proxy", f)
}

// ConnectDisconnectProxy signal is emitted after disconnecting a proxy from an
// action in the group.
//
// UIManager proxies the signal and provides global notification just before any
// action is connected to a proxy, which is probably more convenient to use.
func (actionGroup *ActionGroup) ConnectDisconnectProxy(f func(action Action, proxy Widgetter)) externglib.SignalHandle {
	return actionGroup.Connect("disconnect-proxy", f)
}

// ConnectPostActivate signal is emitted just after the action in the
// action_group is activated
//
// This is intended for UIManager to proxy the signal and provide global
// notification just after any action is activated.
func (actionGroup *ActionGroup) ConnectPostActivate(f func(action Action)) externglib.SignalHandle {
	return actionGroup.Connect("post-activate", f)
}

// ConnectPreActivate signal is emitted just before the action in the
// action_group is activated
//
// This is intended for UIManager to proxy the signal and provide global
// notification just before any action is activated.
func (actionGroup *ActionGroup) ConnectPreActivate(f func(action Action)) externglib.SignalHandle {
	return actionGroup.Connect("pre-activate", f)
}

// NewActionGroup creates a new ActionGroup object. The name of the action group
// is used when associating [keybindings][Action-Accel] with the actions.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - name of the action group.
//
// The function returns the following values:
//
//    - actionGroup: new ActionGroup.
//
func NewActionGroup(name string) *ActionGroup {
	var _arg1 *C.gchar          // out
	var _cret *C.GtkActionGroup // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_action_group_new(_arg1)
	runtime.KeepAlive(name)

	var _actionGroup *ActionGroup // out

	_actionGroup = wrapActionGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _actionGroup
}

// AddAction adds an action object to the action group. Note that this function
// does not set up the accel path of the action, which can lead to problems if a
// user tries to modify the accelerator of a menuitem associated with the
// action. Therefore you must either set the accel path yourself with
// gtk_action_set_accel_path(), or use gtk_action_group_add_action_with_accel
// (..., NULL).
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - action: action.
//
func (actionGroup *ActionGroup) AddAction(action *Action) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_add_action(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(action)
}

// AddActionWithAccel adds an action object to the action group and sets up the
// accelerator.
//
// If accelerator is NULL, attempts to use the accelerator associated with the
// stock_id of the action.
//
// Accel paths are set to <Actions>/group-name/action-name.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - action to add.
//    - accelerator (optional) for the action, in the format understood by
//      gtk_accelerator_parse(), or "" for no accelerator, or NULL to use the
//      stock accelerator.
//
func (actionGroup *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out
	var _arg2 *C.gchar          // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	if accelerator != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_action_group_add_action_with_accel(_arg0, _arg1, _arg2)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(action)
	runtime.KeepAlive(accelerator)
}

// AccelGroup gets the accelerator group.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - accelGroup: accelerator group associated with this action group or NULL
//      if there is none.
//
func (actionGroup *ActionGroup) AccelGroup() *AccelGroup {
	var _arg0 *C.GtkActionGroup // out
	var _cret *C.GtkAccelGroup  // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_accel_group(_arg0)
	runtime.KeepAlive(actionGroup)

	var _accelGroup *AccelGroup // out

	_accelGroup = wrapAccelGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _accelGroup
}

// Action looks up an action in the action group by name.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - actionName: name of the action.
//
// The function returns the following values:
//
//    - action: action, or NULL if no action by that name exists.
//
func (actionGroup *ActionGroup) Action(actionName string) *Action {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out
	var _cret *C.GtkAction      // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_action_group_get_action(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(actionName)

	var _action *Action // out

	_action = wrapAction(externglib.Take(unsafe.Pointer(_cret)))

	return _action
}

// Name gets the name of the action group.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - utf8: name of the action group.
//
func (actionGroup *ActionGroup) Name() string {
	var _arg0 *C.GtkActionGroup // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_name(_arg0)
	runtime.KeepAlive(actionGroup)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Sensitive returns TRUE if the group is sensitive. The constituent actions can
// only be logically sensitive (see gtk_action_is_sensitive()) if they are
// sensitive (see gtk_action_get_sensitive()) and their group is sensitive.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: TRUE if the group is sensitive.
//
func (actionGroup *ActionGroup) Sensitive() bool {
	var _arg0 *C.GtkActionGroup // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_sensitive(_arg0)
	runtime.KeepAlive(actionGroup)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible returns TRUE if the group is visible. The constituent actions can
// only be logically visible (see gtk_action_is_visible()) if they are visible
// (see gtk_action_get_visible()) and their group is visible.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - ok: TRUE if the group is visible.
//
func (actionGroup *ActionGroup) Visible() bool {
	var _arg0 *C.GtkActionGroup // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_get_visible(_arg0)
	runtime.KeepAlive(actionGroup)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListActions lists the actions in the action group.
//
// Deprecated: since version 3.10.
//
// The function returns the following values:
//
//    - list: allocated list of the action objects in the action group.
//
func (actionGroup *ActionGroup) ListActions() []Action {
	var _arg0 *C.GtkActionGroup // out
	var _cret *C.GList          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))

	_cret = C.gtk_action_group_list_actions(_arg0)
	runtime.KeepAlive(actionGroup)

	var _list []Action // out

	_list = make([]Action, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkAction)(v)
		var dst Action // out
		dst = *wrapAction(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// RemoveAction removes an action object from the action group.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - action: action.
//
func (actionGroup *ActionGroup) RemoveAction(action *Action) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_remove_action(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(action)
}

// SetAccelGroup sets the accelerator group to be used by every action in this
// group.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - accelGroup (optional) to set or NULL.
//
func (actionGroup *ActionGroup) SetAccelGroup(accelGroup *AccelGroup) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAccelGroup  // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	if accelGroup != nil {
		_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	}

	C.gtk_action_group_set_accel_group(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(accelGroup)
}

// SetSensitive changes the sensitivity of action_group
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - sensitive: new sensitivity.
//
func (actionGroup *ActionGroup) SetSensitive(sensitive bool) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_action_group_set_sensitive(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(sensitive)
}

// SetTranslateFunc sets a function to be used for translating the label and
// tooltip of ActionEntrys added by gtk_action_group_add_actions().
//
// If you’re using gettext(), it is enough to set the translation domain with
// gtk_action_group_set_translation_domain().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - fn: TranslateFunc.
//
func (actionGroup *ActionGroup) SetTranslateFunc(fn TranslateFunc) {
	var _arg0 *C.GtkActionGroup  // out
	var _arg1 C.GtkTranslateFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_TranslateFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_action_group_set_translate_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(fn)
}

// SetTranslationDomain sets the translation domain and uses g_dgettext() for
// translating the label and tooltip of ActionEntrys added by
// gtk_action_group_add_actions().
//
// If you’re not using gettext() for localization, see
// gtk_action_group_set_translate_func().
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - domain (optional): translation domain to use for g_dgettext() calls, or
//      NULL to use the domain set with textdomain().
//
func (actionGroup *ActionGroup) SetTranslationDomain(domain string) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	if domain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_action_group_set_translation_domain(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(domain)
}

// SetVisible changes the visible of action_group.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - visible: new visiblity.
//
func (actionGroup *ActionGroup) SetVisible(visible bool) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_action_group_set_visible(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(visible)
}

// TranslateString translates a string using the function set with
// gtk_action_group_set_translate_func(). This is mainly intended for language
// bindings.
//
// Deprecated: since version 3.10.
//
// The function takes the following parameters:
//
//    - str: string.
//
// The function returns the following values:
//
//    - utf8: translation of string.
//
func (actionGroup *ActionGroup) TranslateString(str string) string {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(actionGroup.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_action_group_translate_string(_arg0, _arg1)
	runtime.KeepAlive(actionGroup)
	runtime.KeepAlive(str)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ActionEntry structs are used with gtk_action_group_add_actions() to construct
// actions.
//
// Deprecated: since version 3.10.
//
// An instance of this type is always passed by reference.
type ActionEntry struct {
	*actionEntry
}

// actionEntry is the struct that's finalized.
type actionEntry struct {
	native *C.GtkActionEntry
}

// Name: name of the action.
func (a *ActionEntry) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.name)))
	return v
}

// StockID: stock id for the action, or the name of an icon from the icon theme.
func (a *ActionEntry) StockID() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.stock_id)))
	return v
}

// Label: label for the action. This field should typically be marked for
// translation, see gtk_action_group_set_translation_domain(). If label is NULL,
// the label of the stock item with id stock_id is used.
func (a *ActionEntry) Label() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.label)))
	return v
}

// Accelerator: accelerator for the action, in the format understood by
// gtk_accelerator_parse().
func (a *ActionEntry) Accelerator() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.accelerator)))
	return v
}

// Tooltip: tooltip for the action. This field should typically be marked for
// translation, see gtk_action_group_set_translation_domain().
func (a *ActionEntry) Tooltip() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.tooltip)))
	return v
}

// RadioActionEntry structs are used with gtk_action_group_add_radio_actions()
// to construct groups of radio actions.
//
// Deprecated: since version 3.10.
//
// An instance of this type is always passed by reference.
type RadioActionEntry struct {
	*radioActionEntry
}

// radioActionEntry is the struct that's finalized.
type radioActionEntry struct {
	native *C.GtkRadioActionEntry
}

// Name: name of the action.
func (r *RadioActionEntry) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.name)))
	return v
}

// StockID: stock id for the action, or the name of an icon from the icon theme.
func (r *RadioActionEntry) StockID() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.stock_id)))
	return v
}

// Label: label for the action. This field should typically be marked for
// translation, see gtk_action_group_set_translation_domain().
func (r *RadioActionEntry) Label() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.label)))
	return v
}

// Accelerator: accelerator for the action, in the format understood by
// gtk_accelerator_parse().
func (r *RadioActionEntry) Accelerator() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.accelerator)))
	return v
}

// Tooltip: tooltip for the action. This field should typically be marked for
// translation, see gtk_action_group_set_translation_domain().
func (r *RadioActionEntry) Tooltip() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.tooltip)))
	return v
}

// Value: value to set on the radio action. See
// gtk_radio_action_get_current_value().
func (r *RadioActionEntry) Value() int {
	var v int // out
	v = int(r.native.value)
	return v
}

// ToggleActionEntry structs are used with gtk_action_group_add_toggle_actions()
// to construct toggle actions.
//
// Deprecated: since version 3.10.
//
// An instance of this type is always passed by reference.
type ToggleActionEntry struct {
	*toggleActionEntry
}

// toggleActionEntry is the struct that's finalized.
type toggleActionEntry struct {
	native *C.GtkToggleActionEntry
}

// Name: name of the action.
func (t *ToggleActionEntry) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(t.native.name)))
	return v
}

// StockID: stock id for the action, or the name of an icon from the icon theme.
func (t *ToggleActionEntry) StockID() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(t.native.stock_id)))
	return v
}

// Label: label for the action. This field should typically be marked for
// translation, see gtk_action_group_set_translation_domain().
func (t *ToggleActionEntry) Label() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(t.native.label)))
	return v
}

// Accelerator: accelerator for the action, in the format understood by
// gtk_accelerator_parse().
func (t *ToggleActionEntry) Accelerator() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(t.native.accelerator)))
	return v
}

// Tooltip: tooltip for the action. This field should typically be marked for
// translation, see gtk_action_group_set_translation_domain().
func (t *ToggleActionEntry) Tooltip() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(t.native.tooltip)))
	return v
}

// IsActive: initial state of the toggle action.
func (t *ToggleActionEntry) IsActive() bool {
	var v bool // out
	if t.native.is_active != 0 {
		v = true
	}
	return v
}
