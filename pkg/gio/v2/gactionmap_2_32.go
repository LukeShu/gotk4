// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// GAction* _gotk4_gio2_ActionMap_virtual_lookup_action(void* fnptr, GActionMap* arg0, gchar* arg1) {
//   return ((GAction* (*)(GActionMap*, gchar*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gio2_ActionMap_virtual_add_action(void* fnptr, GActionMap* arg0, GAction* arg1) {
//   ((void (*)(GActionMap*, GAction*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gio2_ActionMap_virtual_remove_action(void* fnptr, GActionMap* arg0, gchar* arg1) {
//   ((void (*)(GActionMap*, gchar*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeActionMap = coreglib.Type(C.g_action_map_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeActionMap, F: marshalActionMap},
	})
}

// ActionMap interface is implemented by Group implementations that operate by
// containing a number of named #GAction instances, such as ActionGroup.
//
// One useful application of this interface is to map the names of actions from
// various action groups to unique, prefixed names (e.g. by prepending "app." or
// "win."). This is the motivation for the 'Map' part of the interface name.
//
// ActionMap wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ActionMap struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ActionMap)(nil)
)

// ActionMapper describes ActionMap's interface methods.
type ActionMapper interface {
	coreglib.Objector

	// AddAction adds an action to the action_map.
	AddAction(action Actioner)
	// AddActionEntries: convenience function for creating multiple Action
	// instances and adding them to a Map.
	AddActionEntries(entries []ActionEntry, userData unsafe.Pointer)
	// LookupAction looks up the action with the name action_name in action_map.
	LookupAction(actionName string) *Action
	// RemoveAction removes the named action from the action map.
	RemoveAction(actionName string)
}

var _ ActionMapper = (*ActionMap)(nil)

func wrapActionMap(obj *coreglib.Object) *ActionMap {
	return &ActionMap{
		Object: obj,
	}
}

func marshalActionMap(p uintptr) (interface{}, error) {
	return wrapActionMap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddAction adds an action to the action_map.
//
// If the action map already contains an action with the same name as action
// then the old action is dropped from the action map.
//
// The action map takes its own reference on action.
//
// The function takes the following parameters:
//
//    - action: #GAction.
//
func (actionMap *ActionMap) AddAction(action Actioner) {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.GAction    // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	C.g_action_map_add_action(_arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(action)
}

// AddActionEntries: convenience function for creating multiple Action instances
// and adding them to a Map.
//
// Each action is constructed as per one Entry.
//
//    static void
//    activate_quit (GSimpleAction *simple,
//                   GVariant      *parameter,
//                   gpointer       user_data)
//    {
//      exit (0);
//    }
//
//    static void
//    activate_print_string (GSimpleAction *simple,
//                           GVariant      *parameter,
//                           gpointer       user_data)
//    {
//      g_print ("s\n", g_variant_get_string (parameter, NULL));
//    }
//
//    static GActionGroup *
//    create_action_group (void)
//    {
//      const GActionEntry entries[] = {
//        { "quit",         activate_quit              },
//        { "print-string", activate_print_string, "s" }
//      };
//      GSimpleActionGroup *group;
//
//      group = g_simple_action_group_new ();
//      g_action_map_add_action_entries (G_ACTION_MAP (group), entries, G_N_ELEMENTS (entries), NULL);
//
//      return G_ACTION_GROUP (group);
//    }.
//
// The function takes the following parameters:
//
//    - entries: pointer to the first item in an array of Entry structs.
//    - userData (optional): user data for signal connections.
//
func (actionMap *ActionMap) AddActionEntries(entries []ActionEntry, userData unsafe.Pointer) {
	var _arg0 *C.GActionMap   // out
	var _arg1 *C.GActionEntry // out
	var _arg2 C.gint
	var _arg3 C.gpointer // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg2 = (C.gint)(len(entries))
	_arg1 = (*C.GActionEntry)(C.calloc(C.size_t(len(entries)), C.size_t(C.sizeof_GActionEntry)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GActionEntry)(_arg1), len(entries))
		for i := range entries {
			out[i] = *(*C.GActionEntry)(gextras.StructNative(unsafe.Pointer((&entries[i]))))
		}
	}
	_arg3 = (C.gpointer)(unsafe.Pointer(userData))

	C.g_action_map_add_action_entries(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(entries)
	runtime.KeepAlive(userData)
}

// LookupAction looks up the action with the name action_name in action_map.
//
// If no such action exists, returns NULL.
//
// The function takes the following parameters:
//
//    - actionName: name of an action.
//
// The function returns the following values:
//
//    - action (optional) or NULL.
//
func (actionMap *ActionMap) LookupAction(actionName string) *Action {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out
	var _cret *C.GAction    // in

	_arg0 = (*C.GActionMap)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_map_lookup_action(_arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)

	var _action *Action // out

	if _cret != nil {
		_action = wrapAction(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _action
}

// RemoveAction removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
//
// The function takes the following parameters:
//
//    - actionName: name of the action.
//
func (actionMap *ActionMap) RemoveAction(actionName string) {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_action_map_remove_action(_arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)
}

// addAction adds an action to the action_map.
//
// If the action map already contains an action with the same name as action
// then the old action is dropped from the action map.
//
// The action map takes its own reference on action.
//
// The function takes the following parameters:
//
//    - action: #GAction.
//
func (actionMap *ActionMap) addAction(action Actioner) {
	gclass := (*C.GActionMapInterface)(coreglib.PeekParentClass(actionMap))
	fnarg := gclass.add_action

	var _arg0 *C.GActionMap // out
	var _arg1 *C.GAction    // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	C._gotk4_gio2_ActionMap_virtual_add_action(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(action)
}

// lookupAction looks up the action with the name action_name in action_map.
//
// If no such action exists, returns NULL.
//
// The function takes the following parameters:
//
//    - actionName: name of an action.
//
// The function returns the following values:
//
//    - action (optional) or NULL.
//
func (actionMap *ActionMap) lookupAction(actionName string) *Action {
	gclass := (*C.GActionMapInterface)(coreglib.PeekParentClass(actionMap))
	fnarg := gclass.lookup_action

	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out
	var _cret *C.GAction    // in

	_arg0 = (*C.GActionMap)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_ActionMap_virtual_lookup_action(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)

	var _action *Action // out

	if _cret != nil {
		_action = wrapAction(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _action
}

// removeAction removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
//
// The function takes the following parameters:
//
//    - actionName: name of the action.
//
func (actionMap *ActionMap) removeAction(actionName string) {
	gclass := (*C.GActionMapInterface)(coreglib.PeekParentClass(actionMap))
	fnarg := gclass.remove_action

	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(coreglib.InternObject(actionMap).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C._gotk4_gio2_ActionMap_virtual_remove_action(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)
}

// ActionMapInterface: virtual function table for Map.
//
// An instance of this type is always passed by reference.
type ActionMapInterface struct {
	*actionMapInterface
}

// actionMapInterface is the struct that's finalized.
type actionMapInterface struct {
	native *C.GActionMapInterface
}
