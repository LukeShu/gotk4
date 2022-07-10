// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeSizeGroup returns the GType for the type SizeGroup.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSizeGroup() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "SizeGroup").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSizeGroup)
	return gtype
}

// SizeGroupOverrider contains methods that are overridable.
type SizeGroupOverrider interface {
}

// SizeGroup provides a mechanism for grouping a number of widgets together so
// they all request the same amount of space. This is typically useful when you
// want a column of widgets to have the same size, but you can’t use a Grid
// widget.
//
// In detail, the size requested for each widget in a SizeGroup is the maximum
// of the sizes that would have been requested for each widget in the size group
// if they were not in the size group. The mode of the size group (see
// gtk_size_group_set_mode()) determines whether this applies to the horizontal
// size, the vertical size, or both sizes.
//
// Note that size groups only affect the amount of space requested, not the size
// that the widgets finally receive. If you want the widgets in a SizeGroup to
// actually be the same size, you need to pack them in such a way that they get
// the size they request and not more. For example, if you are packing your
// widgets into a table, you would not include the GTK_FILL flag.
//
// SizeGroup objects are referenced by each widget in the size group, so once
// you have added all widgets to a SizeGroup, you can drop the initial reference
// to the size group with g_object_unref(). If the widgets in the size group are
// subsequently destroyed, then they will be removed from the size group and
// drop their references on the size group; when all widgets have been removed,
// the size group will be freed.
//
// Widgets can be part of multiple size groups; GTK+ will compute the horizontal
// size of a widget from the horizontal requisition of all widgets that can be
// reached from the widget by a chain of size groups of type
// GTK_SIZE_GROUP_HORIZONTAL or GTK_SIZE_GROUP_BOTH, and the vertical size from
// the vertical requisition of all widgets that can be reached from the widget
// by a chain of size groups of type GTK_SIZE_GROUP_VERTICAL or
// GTK_SIZE_GROUP_BOTH.
//
// Note that only non-contextual sizes of every widget are ever consulted by
// size groups (since size groups have no knowledge of what size a widget will
// be allocated in one dimension, it cannot derive how much height a widget will
// receive for a given width). When grouping widgets that trade height for width
// in mode GTK_SIZE_GROUP_VERTICAL or GTK_SIZE_GROUP_BOTH: the height for the
// minimum width will be the requested height for all widgets in the group. The
// same is of course true when horizontally grouping width for height widgets.
//
// Widgets that trade height-for-width should set a reasonably large minimum
// width by way of Label:width-chars for instance. Widgets with static sizes as
// well as widgets that grow (such as ellipsizing text) need no such
// considerations.
//
//
// GtkSizeGroup as GtkBuildable
//
// Size groups can be specified in a UI definition by placing an <object>
// element with class="GtkSizeGroup" somewhere in the UI definition. The widgets
// that belong to the size group are specified by a <widgets> element that may
// contain multiple <widget> elements, one for each member of the size group.
// The ”name” attribute gives the id of the widget.
//
// An example of a UI definition fragment with GtkSizeGroup:
//
//    <object class="GtkSizeGroup">
//      <property name="mode">GTK_SIZE_GROUP_HORIZONTAL</property>
//      <widgets>
//        <widget name="radio1"/>
//        <widget name="radio2"/>
//      </widgets>
//    </object>.
type SizeGroup struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
}

var (
	_ coreglib.Objector = (*SizeGroup)(nil)
)

func classInitSizeGrouper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSizeGroup(obj *coreglib.Object) *SizeGroup {
	return &SizeGroup{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	return wrapSizeGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddWidget adds a widget to a SizeGroup. In the future, the requisition of the
// widget will be determined as the maximum of its requisition and the
// requisition of the other widgets in the size group. Whether this applies
// horizontally, vertically, or in both directions depends on the mode of the
// size group. See gtk_size_group_set_mode().
//
// When the widget is destroyed or no longer referenced elsewhere, it will be
// removed from the size group.
//
// The function takes the following parameters:
//
//    - widget to add.
//
func (sizeGroup *SizeGroup) AddWidget(widget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sizeGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_info := girepository.MustFind("Gtk", "SizeGroup")
	_info.InvokeClassMethod("add_widget", _args[:], nil)

	runtime.KeepAlive(sizeGroup)
	runtime.KeepAlive(widget)
}

// IgnoreHidden returns if invisible widgets are ignored when calculating the
// size.
//
// Deprecated: Measuring the size of hidden widgets has not worked reliably for
// a long time. In most cases, they will report a size of 0 nowadays, and thus,
// their size will not affect the other size group members. In effect, size
// groups will always operate as if this property was TRUE. Use a Stack instead
// to hide widgets while still having their size taken into account.
//
// The function returns the following values:
//
//    - ok: TRUE if invisible widgets are ignored.
//
func (sizeGroup *SizeGroup) IgnoreHidden() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sizeGroup).Native()))

	_info := girepository.MustFind("Gtk", "SizeGroup")
	_gret := _info.InvokeClassMethod("get_ignore_hidden", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sizeGroup)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Widgets returns the list of widgets associated with size_group.
//
// The function returns the following values:
//
//    - sList of widgets. The list is owned by GTK+ and should not be modified.
//
func (sizeGroup *SizeGroup) Widgets() []Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sizeGroup).Native()))

	_info := girepository.MustFind("Gtk", "SizeGroup")
	_gret := _info.InvokeClassMethod("get_widgets", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sizeGroup)

	var _sList []Widgetter // out

	_sList = make([]Widgetter, 0, gextras.SListSize(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	gextras.MoveSList(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))), false, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst Widgetter // out
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src)))
			if objptr == nil {
				panic("object of type gtk.Widgetter is nil")
			}

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			dst = rv
		}
		_sList = append(_sList, dst)
	})

	return _sList
}

// RemoveWidget removes a widget from a SizeGroup.
//
// The function takes the following parameters:
//
//    - widget to remove.
//
func (sizeGroup *SizeGroup) RemoveWidget(widget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sizeGroup).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_info := girepository.MustFind("Gtk", "SizeGroup")
	_info.InvokeClassMethod("remove_widget", _args[:], nil)

	runtime.KeepAlive(sizeGroup)
	runtime.KeepAlive(widget)
}

// SetIgnoreHidden sets whether unmapped widgets should be ignored when
// calculating the size.
//
// Deprecated: Measuring the size of hidden widgets has not worked reliably for
// a long time. In most cases, they will report a size of 0 nowadays, and thus,
// their size will not affect the other size group members. In effect, size
// groups will always operate as if this property was TRUE. Use a Stack instead
// to hide widgets while still having their size taken into account.
//
// The function takes the following parameters:
//
//    - ignoreHidden: whether unmapped widgets should be ignored when calculating
//      the size.
//
func (sizeGroup *SizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sizeGroup).Native()))
	if ignoreHidden {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "SizeGroup")
	_info.InvokeClassMethod("set_ignore_hidden", _args[:], nil)

	runtime.KeepAlive(sizeGroup)
	runtime.KeepAlive(ignoreHidden)
}
