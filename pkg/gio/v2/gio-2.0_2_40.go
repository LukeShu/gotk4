// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_AppInfoMonitor_ConnectChanged(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeAppInfoMonitor = coreglib.Type(C.g_app_info_monitor_get_type())
	GTypeNotification   = coreglib.Type(C.g_notification_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAppInfoMonitor, F: marshalAppInfoMonitor},
		coreglib.TypeMarshaler{T: GTypeNotification, F: marshalNotification},
	})
}

// AppInfoMonitor is a very simple object used for monitoring the app info
// database for changes (ie: newly installed or removed applications).
//
// Call g_app_info_monitor_get() to get a InfoMonitor and connect to the
// "changed" signal.
//
// In the usual case, applications should try to make note of the change
// (doing things like invalidating caches) but not act on it. In particular,
// applications should avoid making calls to Info APIs in response to the change
// signal, deferring these until the time that the data is actually required.
// The exception to this case is when application information is actually
// being displayed on the screen (eg: during a search or when the list of all
// applications is shown). The reason for this is that changes to the list of
// installed applications often come in groups (like during system updates) and
// rescanning the list on every change is pointless and expensive.
type AppInfoMonitor struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AppInfoMonitor)(nil)
)

func wrapAppInfoMonitor(obj *coreglib.Object) *AppInfoMonitor {
	return &AppInfoMonitor{
		Object: obj,
	}
}

func marshalAppInfoMonitor(p uintptr) (interface{}, error) {
	return wrapAppInfoMonitor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged: signal emitted when the app info database for changes (ie:
// newly installed or removed applications).
func (v *AppInfoMonitor) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gio2_AppInfoMonitor_ConnectChanged), f)
}

// Notification is a mechanism for creating a notification to be shown to
// the user -- typically as a pop-up notification presented by the desktop
// environment shell.
//
// The key difference between #GNotification and other similar APIs is that, if
// supported by the desktop environment, notifications sent with #GNotification
// will persist after the application has exited, and even across system
// reboots.
//
// Since the user may click on a notification while the application is not
// running, applications using #GNotification should be able to be started as a
// D-Bus service, using #GApplication.
//
// User interaction with a notification (either the default action, or buttons)
// must be associated with actions on the application (ie: "app." actions).
// It is not possible to route user interaction through the notification itself,
// because the object will not exist if the application is autostarted as a
// result of a notification being clicked.
//
// A notification can be sent with g_application_send_notification().
type Notification struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Notification)(nil)
)

func wrapNotification(obj *coreglib.Object) *Notification {
	return &Notification{
		Object: obj,
	}
}

func marshalNotification(p uintptr) (interface{}, error) {
	return wrapNotification(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNotification creates a new #GNotification with title as its title.
//
// After populating notification with more details, it can be sent to the
// desktop shell with g_application_send_notification(). Changing any properties
// after this call will not have any effect until resending notification.
//
// The function takes the following parameters:
//
//   - title of the notification.
//
// The function returns the following values:
//
//   - notification: new #GNotification instance.
//
func NewNotification(title string) *Notification {
	var _arg1 *C.gchar         // out
	var _cret *C.GNotification // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_notification_new(_arg1)
	runtime.KeepAlive(title)

	var _notification *Notification // out

	_notification = wrapNotification(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _notification
}

// AddButton adds a button to notification that activates the action in
// detailed_action when clicked. That action must be an application-wide action
// (starting with "app."). If detailed_action contains a target, the action will
// be activated with that target as its parameter.
//
// See g_action_parse_detailed_name() for a description of the format for
// detailed_action.
//
// The function takes the following parameters:
//
//   - label of the button.
//   - detailedAction: detailed action name.
//
func (notification *Notification) AddButton(label, detailedAction string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_notification_add_button(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(label)
	runtime.KeepAlive(detailedAction)
}

// AddButtonWithTarget adds a button to notification that activates action
// when clicked. action must be an application-wide action (it must start with
// "app.").
//
// If target is non-NULL, action will be activated with target as its parameter.
//
// The function takes the following parameters:
//
//   - label of the button.
//   - action name.
//   - target (optional) to use as action's parameter, or NULL.
//
func (notification *Notification) AddButtonWithTarget(label, action string, target *glib.Variant) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // out
	var _arg3 *C.GVariant      // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(action)))
	defer C.free(unsafe.Pointer(_arg2))
	if target != nil {
		_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(target)))
	}

	C.g_notification_add_button_with_target_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(label)
	runtime.KeepAlive(action)
	runtime.KeepAlive(target)
}

// SetBody sets the body of notification to body.
//
// The function takes the following parameters:
//
//   - body (optional): new body for notification, or NULL.
//
func (notification *Notification) SetBody(body string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	if body != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(body)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_notification_set_body(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(body)
}

// SetDefaultAction sets the default action of notification to detailed_action.
// This action is activated when the notification is clicked on.
//
// The action in detailed_action must be an application-wide action
// (it must start with "app."). If detailed_action contains a target,
// the given action will be activated with that target as its parameter.
// See g_action_parse_detailed_name() for a description of the format for
// detailed_action.
//
// When no default action is set, the application that the notification was sent
// on is activated.
//
// The function takes the following parameters:
//
//   - detailedAction: detailed action name.
//
func (notification *Notification) SetDefaultAction(detailedAction string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_notification_set_default_action(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(detailedAction)
}

// SetDefaultActionAndTarget sets the default action of notification to action.
// This action is activated when the notification is clicked on. It must be an
// application-wide action (start with "app.").
//
// If target is non-NULL, action will be activated with target as its parameter.
//
// When no default action is set, the application that the notification was sent
// on is activated.
//
// The function takes the following parameters:
//
//   - action name.
//   - target (optional) to use as action's parameter, or NULL.
//
func (notification *Notification) SetDefaultActionAndTarget(action string, target *glib.Variant) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariant      // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(action)))
	defer C.free(unsafe.Pointer(_arg1))
	if target != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(target)))
	}

	C.g_notification_set_default_action_and_target_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(action)
	runtime.KeepAlive(target)
}

// SetIcon sets the icon of notification to icon.
//
// The function takes the following parameters:
//
//   - icon to be shown in notification, as a #GIcon.
//
func (notification *Notification) SetIcon(icon Iconner) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.GIcon         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	C.g_notification_set_icon(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(icon)
}

// SetPriority sets the priority of notification to priority. See Priority for
// possible values.
//
// The function takes the following parameters:
//
//   - priority: Priority.
//
func (notification *Notification) SetPriority(priority NotificationPriority) {
	var _arg0 *C.GNotification        // out
	var _arg1 C.GNotificationPriority // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	_arg1 = C.GNotificationPriority(priority)

	C.g_notification_set_priority(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(priority)
}

// SetTitle sets the title of notification to title.
//
// The function takes the following parameters:
//
//   - title: new title for notification.
//
func (notification *Notification) SetTitle(title string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_notification_set_title(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(title)
}

// SetUrgent: deprecated in favor of g_notification_set_priority().
//
// Deprecated: Since 2.42, this has been deprecated in favour of
// g_notification_set_priority().
//
// The function takes the following parameters:
//
//   - urgent: TRUE if notification is urgent.
//
func (notification *Notification) SetUrgent(urgent bool) {
	var _arg0 *C.GNotification // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(coreglib.InternObject(notification).Native()))
	if urgent {
		_arg1 = C.TRUE
	}

	C.g_notification_set_urgent(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(urgent)
}
