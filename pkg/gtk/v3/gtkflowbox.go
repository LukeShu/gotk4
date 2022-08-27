// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_FlowBox_ConnectUnselectAll(gpointer, guintptr);
// extern void _gotk4_gtk3_FlowBox_ConnectToggleCursorChild(gpointer, guintptr);
// extern void _gotk4_gtk3_FlowBox_ConnectSelectedChildrenChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_FlowBox_ConnectSelectAll(gpointer, guintptr);
// extern void _gotk4_gtk3_FlowBox_ConnectChildActivated(gpointer, GtkFlowBoxChild*, guintptr);
// extern void _gotk4_gtk3_FlowBox_ConnectActivateCursorChild(gpointer, guintptr);
// extern void _gotk4_gtk3_FlowBoxClass_unselect_all(GtkFlowBox*);
// extern void _gotk4_gtk3_FlowBoxClass_toggle_cursor_child(GtkFlowBox*);
// extern void _gotk4_gtk3_FlowBoxClass_selected_children_changed(GtkFlowBox*);
// extern void _gotk4_gtk3_FlowBoxClass_select_all(GtkFlowBox*);
// extern void _gotk4_gtk3_FlowBoxClass_child_activated(GtkFlowBox*, GtkFlowBoxChild*);
// extern void _gotk4_gtk3_FlowBoxClass_activate_cursor_child(GtkFlowBox*);
// extern void _gotk4_gtk3_FlowBoxChild_ConnectActivate(gpointer, guintptr);
// extern void _gotk4_gtk3_FlowBoxChildClass_activate(GtkFlowBoxChild*);
// extern gboolean _gotk4_gtk3_FlowBox_ConnectMoveCursor(gpointer, GtkMovementStep, gint, guintptr);
// extern gboolean _gotk4_gtk3_FlowBoxClass_move_cursor(GtkFlowBox*, GtkMovementStep, gint);
// gboolean _gotk4_gtk3_FlowBox_virtual_move_cursor(void* fnptr, GtkFlowBox* arg0, GtkMovementStep arg1, gint arg2) {
//   return ((gboolean (*)(GtkFlowBox*, GtkMovementStep, gint))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gtk3_FlowBoxChild_virtual_activate(void* fnptr, GtkFlowBoxChild* arg0) {
//   ((void (*)(GtkFlowBoxChild*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_FlowBox_virtual_activate_cursor_child(void* fnptr, GtkFlowBox* arg0) {
//   ((void (*)(GtkFlowBox*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_FlowBox_virtual_child_activated(void* fnptr, GtkFlowBox* arg0, GtkFlowBoxChild* arg1) {
//   ((void (*)(GtkFlowBox*, GtkFlowBoxChild*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_FlowBox_virtual_selected_children_changed(void* fnptr, GtkFlowBox* arg0) {
//   ((void (*)(GtkFlowBox*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_FlowBox_virtual_toggle_cursor_child(void* fnptr, GtkFlowBox* arg0) {
//   ((void (*)(GtkFlowBox*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeFlowBox      = coreglib.Type(C.gtk_flow_box_get_type())
	GTypeFlowBoxChild = coreglib.Type(C.gtk_flow_box_child_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFlowBox, F: marshalFlowBox},
		coreglib.TypeMarshaler{T: GTypeFlowBoxChild, F: marshalFlowBoxChild},
	})
}

// FlowBoxOverrides contains methods that are overridable.
type FlowBoxOverrides struct {
	ActivateCursorChild func()
	// The function takes the following parameters:
	//
	ChildActivated func(child *FlowBoxChild)
	// The function takes the following parameters:
	//
	//    - step
	//    - count
	//
	// The function returns the following values:
	//
	MoveCursor func(step MovementStep, count int) bool
	// SelectAll: select all children of box, if the selection mode allows it.
	SelectAll               func()
	SelectedChildrenChanged func()
	ToggleCursorChild       func()
	// UnselectAll: unselect all children of box, if the selection mode allows
	// it.
	UnselectAll func()
}

func defaultFlowBoxOverrides(v *FlowBox) FlowBoxOverrides {
	return FlowBoxOverrides{
		ActivateCursorChild:     v.activateCursorChild,
		ChildActivated:          v.childActivated,
		MoveCursor:              v.moveCursor,
		SelectAll:               v.selectAll,
		SelectedChildrenChanged: v.selectedChildrenChanged,
		ToggleCursorChild:       v.toggleCursorChild,
		UnselectAll:             v.unselectAll,
	}
}

// FlowBox positions child widgets in sequence according to its orientation.
//
// For instance, with the horizontal orientation, the widgets will be arranged
// from left to right, starting a new row under the previous row when necessary.
// Reducing the width in this case will require more rows, so a larger height
// will be requested.
//
// Likewise, with the vertical orientation, the widgets will be arranged from
// top to bottom, starting a new column to the right when necessary. Reducing
// the height will require more columns, so a larger width will be requested.
//
// The size request of a GtkFlowBox alone may not be what you expect; if you
// need to be able to shrink it along both axes and dynamically reflow its
// children, you may have to wrap it in a ScrolledWindow to enable that.
//
// The children of a GtkFlowBox can be dynamically sorted and filtered.
//
// Although a GtkFlowBox must have only FlowBoxChild children, you can add any
// kind of widget to it via gtk_container_add(), and a GtkFlowBoxChild widget
// will automatically be inserted between the box and the widget.
//
// Also see ListBox.
//
// GtkFlowBox was added in GTK+ 3.12.
//
// CSS nodes
//
//    flowbox
//    ├── flowboxchild
//    │   ╰── <child>
//    ├── flowboxchild
//    │   ╰── <child>
//    ┊
//    ╰── [rubberband]
//
// GtkFlowBox uses a single CSS node with name flowbox. GtkFlowBoxChild uses a
// single CSS node with name flowboxchild. For rubberband selection, a subnode
// with name rubberband is used.
type FlowBox struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Orientable
}

var (
	_ Containerer       = (*FlowBox)(nil)
	_ coreglib.Objector = (*FlowBox)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FlowBox, *FlowBoxClass, FlowBoxOverrides](
		GTypeFlowBox,
		initFlowBoxClass,
		wrapFlowBox,
		defaultFlowBoxOverrides,
	)
}

func initFlowBoxClass(gclass unsafe.Pointer, overrides FlowBoxOverrides, classInitFunc func(*FlowBoxClass)) {
	pclass := (*C.GtkFlowBoxClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeFlowBox))))

	if overrides.ActivateCursorChild != nil {
		pclass.activate_cursor_child = (*[0]byte)(C._gotk4_gtk3_FlowBoxClass_activate_cursor_child)
	}

	if overrides.ChildActivated != nil {
		pclass.child_activated = (*[0]byte)(C._gotk4_gtk3_FlowBoxClass_child_activated)
	}

	if overrides.MoveCursor != nil {
		pclass.move_cursor = (*[0]byte)(C._gotk4_gtk3_FlowBoxClass_move_cursor)
	}

	if overrides.SelectAll != nil {
		pclass.select_all = (*[0]byte)(C._gotk4_gtk3_FlowBoxClass_select_all)
	}

	if overrides.SelectedChildrenChanged != nil {
		pclass.selected_children_changed = (*[0]byte)(C._gotk4_gtk3_FlowBoxClass_selected_children_changed)
	}

	if overrides.ToggleCursorChild != nil {
		pclass.toggle_cursor_child = (*[0]byte)(C._gotk4_gtk3_FlowBoxClass_toggle_cursor_child)
	}

	if overrides.UnselectAll != nil {
		pclass.unselect_all = (*[0]byte)(C._gotk4_gtk3_FlowBoxClass_unselect_all)
	}

	if classInitFunc != nil {
		class := (*FlowBoxClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFlowBox(obj *coreglib.Object) *FlowBox {
	return &FlowBox{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalFlowBox(p uintptr) (interface{}, error) {
	return wrapFlowBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivateCursorChild signal is a [keybinding signal][GtkBindingSignal]
// which gets emitted when the user activates the box.
func (box *FlowBox) ConnectActivateCursorChild(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(box, "activate-cursor-child", false, unsafe.Pointer(C._gotk4_gtk3_FlowBox_ConnectActivateCursorChild), f)
}

// ConnectChildActivated signal is emitted when a child has been activated by
// the user.
func (box *FlowBox) ConnectChildActivated(f func(child *FlowBoxChild)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(box, "child-activated", false, unsafe.Pointer(C._gotk4_gtk3_FlowBox_ConnectChildActivated), f)
}

// ConnectMoveCursor signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user initiates a cursor movement.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control the cursor programmatically.
//
// The default bindings for this signal come in two variants, the variant with
// the Shift modifier extends the selection, the variant without the Shift
// modifer does not. There are too many key combinations to list them all here.
//
// - Arrow keys move by individual children
//
// - Home/End keys move to the ends of the box
//
// - PageUp/PageDown keys move vertically by pages.
func (box *FlowBox) ConnectMoveCursor(f func(step MovementStep, count int) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(box, "move-cursor", false, unsafe.Pointer(C._gotk4_gtk3_FlowBox_ConnectMoveCursor), f)
}

// ConnectSelectAll signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to select all children of the box, if the selection mode permits it.
//
// The default bindings for this signal is Ctrl-a.
func (box *FlowBox) ConnectSelectAll(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(box, "select-all", false, unsafe.Pointer(C._gotk4_gtk3_FlowBox_ConnectSelectAll), f)
}

// ConnectSelectedChildrenChanged signal is emitted when the set of selected
// children changes.
//
// Use gtk_flow_box_selected_foreach() or gtk_flow_box_get_selected_children()
// to obtain the selected children.
func (box *FlowBox) ConnectSelectedChildrenChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(box, "selected-children-changed", false, unsafe.Pointer(C._gotk4_gtk3_FlowBox_ConnectSelectedChildrenChanged), f)
}

// ConnectToggleCursorChild signal is a [keybinding signal][GtkBindingSignal]
// which toggles the selection of the child that has the focus.
//
// The default binding for this signal is Ctrl-Space.
func (box *FlowBox) ConnectToggleCursorChild(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(box, "toggle-cursor-child", false, unsafe.Pointer(C._gotk4_gtk3_FlowBox_ConnectToggleCursorChild), f)
}

// ConnectUnselectAll signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted to unselect all children of the box, if the selection mode
// permits it.
//
// The default bindings for this signal is Ctrl-Shift-a.
func (box *FlowBox) ConnectUnselectAll(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(box, "unselect-all", false, unsafe.Pointer(C._gotk4_gtk3_FlowBox_ConnectUnselectAll), f)
}

func (box *FlowBox) activateCursorChild() {
	gclass := (*C.GtkFlowBoxClass)(coreglib.PeekParentClass(box))
	fnarg := gclass.activate_cursor_child

	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	C._gotk4_gtk3_FlowBox_virtual_activate_cursor_child(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(box)
}

// The function takes the following parameters:
//
func (box *FlowBox) childActivated(child *FlowBoxChild) {
	gclass := (*C.GtkFlowBoxClass)(coreglib.PeekParentClass(box))
	fnarg := gclass.child_activated

	var _arg0 *C.GtkFlowBox      // out
	var _arg1 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = (*C.GtkFlowBoxChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C._gotk4_gtk3_FlowBox_virtual_child_activated(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// The function takes the following parameters:
//
//    - step
//    - count
//
// The function returns the following values:
//
func (box *FlowBox) moveCursor(step MovementStep, count int) bool {
	gclass := (*C.GtkFlowBoxClass)(coreglib.PeekParentClass(box))
	fnarg := gclass.move_cursor

	var _arg0 *C.GtkFlowBox     // out
	var _arg1 C.GtkMovementStep // out
	var _arg2 C.gint            // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(coreglib.InternObject(box).Native()))
	_arg1 = C.GtkMovementStep(step)
	_arg2 = C.gint(count)

	_cret = C._gotk4_gtk3_FlowBox_virtual_move_cursor(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(step)
	runtime.KeepAlive(count)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (box *FlowBox) selectedChildrenChanged() {
	gclass := (*C.GtkFlowBoxClass)(coreglib.PeekParentClass(box))
	fnarg := gclass.selected_children_changed

	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	C._gotk4_gtk3_FlowBox_virtual_selected_children_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(box)
}

func (box *FlowBox) toggleCursorChild() {
	gclass := (*C.GtkFlowBoxClass)(coreglib.PeekParentClass(box))
	fnarg := gclass.toggle_cursor_child

	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(coreglib.InternObject(box).Native()))

	C._gotk4_gtk3_FlowBox_virtual_toggle_cursor_child(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(box)
}

// FlowBoxChildOverrides contains methods that are overridable.
type FlowBoxChildOverrides struct {
	Activate func()
}

func defaultFlowBoxChildOverrides(v *FlowBoxChild) FlowBoxChildOverrides {
	return FlowBoxChildOverrides{
		Activate: v.activate,
	}
}

type FlowBoxChild struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*FlowBoxChild)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*FlowBoxChild, *FlowBoxChildClass, FlowBoxChildOverrides](
		GTypeFlowBoxChild,
		initFlowBoxChildClass,
		wrapFlowBoxChild,
		defaultFlowBoxChildOverrides,
	)
}

func initFlowBoxChildClass(gclass unsafe.Pointer, overrides FlowBoxChildOverrides, classInitFunc func(*FlowBoxChildClass)) {
	pclass := (*C.GtkFlowBoxChildClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeFlowBoxChild))))

	if overrides.Activate != nil {
		pclass.activate = (*[0]byte)(C._gotk4_gtk3_FlowBoxChildClass_activate)
	}

	if classInitFunc != nil {
		class := (*FlowBoxChildClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFlowBoxChild(obj *coreglib.Object) *FlowBoxChild {
	return &FlowBoxChild{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalFlowBoxChild(p uintptr) (interface{}, error) {
	return wrapFlowBoxChild(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate signal is emitted when the user activates a child widget in a
// FlowBox, either by clicking or double-clicking, or by using the Space or
// Enter key.
//
// While this signal is used as a [keybinding signal][GtkBindingSignal], it can
// be used by applications for their own purposes.
func (child *FlowBoxChild) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(child, "activate", false, unsafe.Pointer(C._gotk4_gtk3_FlowBoxChild_ConnectActivate), f)
}

func (child *FlowBoxChild) activate() {
	gclass := (*C.GtkFlowBoxChildClass)(coreglib.PeekParentClass(child))
	fnarg := gclass.activate

	var _arg0 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C._gotk4_gtk3_FlowBoxChild_virtual_activate(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(child)
}

// FlowBoxChildClass: instance of this type is always passed by reference.
type FlowBoxChildClass struct {
	*flowBoxChildClass
}

// flowBoxChildClass is the struct that's finalized.
type flowBoxChildClass struct {
	native *C.GtkFlowBoxChildClass
}

func (f *FlowBoxChildClass) ParentClass() *BinClass {
	valptr := &f.native.parent_class
	var _v *BinClass // out
	_v = (*BinClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// FlowBoxClass: instance of this type is always passed by reference.
type FlowBoxClass struct {
	*flowBoxClass
}

// flowBoxClass is the struct that's finalized.
type flowBoxClass struct {
	native *C.GtkFlowBoxClass
}

func (f *FlowBoxClass) ParentClass() *ContainerClass {
	valptr := &f.native.parent_class
	var _v *ContainerClass // out
	_v = (*ContainerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
