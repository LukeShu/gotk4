// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtktoolpalette.go.
var (
	GTypeToolPaletteDragTargets = coreglib.Type(C.gtk_tool_palette_drag_targets_get_type())
	GTypeToolPalette            = coreglib.Type(C.gtk_tool_palette_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeToolPaletteDragTargets, F: marshalToolPaletteDragTargets},
		{T: GTypeToolPalette, F: marshalToolPalette},
	})
}

// ToolPaletteDragTargets flags used to specify the supported drag targets.
type ToolPaletteDragTargets C.guint

const (
	// ToolPaletteDragItems: support drag of items.
	ToolPaletteDragItems ToolPaletteDragTargets = 0b1
	// ToolPaletteDragGroups: support drag of groups.
	ToolPaletteDragGroups ToolPaletteDragTargets = 0b10
)

func marshalToolPaletteDragTargets(p uintptr) (interface{}, error) {
	return ToolPaletteDragTargets(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ToolPaletteDragTargets.
func (t ToolPaletteDragTargets) String() string {
	if t == 0 {
		return "ToolPaletteDragTargets(0)"
	}

	var builder strings.Builder
	builder.Grow(42)

	for t != 0 {
		next := t & (t - 1)
		bit := t - next

		switch bit {
		case ToolPaletteDragItems:
			builder.WriteString("Items|")
		case ToolPaletteDragGroups:
			builder.WriteString("Groups|")
		default:
			builder.WriteString(fmt.Sprintf("ToolPaletteDragTargets(0b%b)|", bit))
		}

		t = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if t contains other.
func (t ToolPaletteDragTargets) Has(other ToolPaletteDragTargets) bool {
	return (t & other) == other
}

// ToolPaletteOverrider contains methods that are overridable.
type ToolPaletteOverrider interface {
}

// ToolPalette allows you to add ToolItems to a palette-like container with
// different categories and drag and drop support.
//
// A ToolPalette is created with a call to gtk_tool_palette_new().
//
// ToolItems cannot be added directly to a ToolPalette - instead they are added
// to a ToolItemGroup which can than be added to a ToolPalette. To add a
// ToolItemGroup to a ToolPalette, use gtk_container_add().
//
//    static void
//    passive_canvas_drag_data_received (GtkWidget        *widget,
//                                       GdkDragContext   *context,
//                                       gint              x,
//                                       gint              y,
//                                       GtkSelectionData *selection,
//                                       guint             info,
//                                       guint             time,
//                                       gpointer          data)
//    {
//      GtkWidget *palette;
//      GtkWidget *item;
//
//      // Get the dragged item
//      palette = gtk_widget_get_ancestor (gtk_drag_get_source_widget (context),
//                                         GTK_TYPE_TOOL_PALETTE);
//      if (palette != NULL)
//        item = gtk_tool_palette_get_drag_item (GTK_TOOL_PALETTE (palette),
//                                               selection);
//
//      // Do something with item
//    }
//
//    GtkWidget *target, palette;
//
//    palette = gtk_tool_palette_new ();
//    target = gtk_drawing_area_new ();
//
//    g_signal_connect (G_OBJECT (target), "drag-data-received",
//                      G_CALLBACK (passive_canvas_drag_data_received), NULL);
//    gtk_tool_palette_add_drag_dest (GTK_TOOL_PALETTE (palette), target,
//                                    GTK_DEST_DEFAULT_ALL,
//                                    GTK_TOOL_PALETTE_DRAG_ITEMS,
//                                    GDK_ACTION_COPY);
//
//
// CSS nodes
//
// GtkToolPalette has a single CSS node named toolpalette.
type ToolPalette struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Orientable
	Scrollable
}

var (
	_ Containerer       = (*ToolPalette)(nil)
	_ coreglib.Objector = (*ToolPalette)(nil)
)

func classInitToolPaletter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapToolPalette(obj *coreglib.Object) *ToolPalette {
	return &ToolPalette{
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
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalToolPalette(p uintptr) (interface{}, error) {
	return wrapToolPalette(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewToolPalette creates a new tool palette.
//
// The function returns the following values:
//
//    - toolPalette: new ToolPalette.
//
func NewToolPalette() *ToolPalette {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("new_ToolPalette", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _toolPalette *ToolPalette // out

	_toolPalette = wrapToolPalette(coreglib.Take(unsafe.Pointer(_cret)))

	return _toolPalette
}

// DragItem: get the dragged item from the selection. This could be a ToolItem
// or a ToolItemGroup.
//
// The function takes the following parameters:
//
//    - selection: SelectionData.
//
// The function returns the following values:
//
//    - widget: dragged item in selection.
//
func (palette *ToolPalette) DragItem(selection *SelectionData) Widgetter {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(selection)))
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_drag_item", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)
	runtime.KeepAlive(selection)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
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
		_widget = rv
	}

	return _widget
}

// DropGroup gets the group at position (x, y).
//
// The function takes the following parameters:
//
//    - x position.
//    - y position.
//
// The function returns the following values:
//
//    - toolItemGroup (optional) at position or NULL if there is no such group.
//
func (palette *ToolPalette) DropGroup(x, y int32) *ToolItemGroup {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _arg2 C.gint  // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_drop_group", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _toolItemGroup *ToolItemGroup // out

	if _cret != nil {
		_toolItemGroup = wrapToolItemGroup(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _toolItemGroup
}

// DropItem gets the item at position (x, y). See
// gtk_tool_palette_get_drop_group().
//
// The function takes the following parameters:
//
//    - x position.
//    - y position.
//
// The function returns the following values:
//
//    - toolItem (optional) at position or NULL if there is no such item.
//
func (palette *ToolPalette) DropItem(x, y int32) *ToolItem {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _arg2 C.gint  // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_drop_item", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _toolItem *ToolItem // out

	if _cret != nil {
		_toolItem = wrapToolItem(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _toolItem
}

// Exclusive gets whether group is exclusive or not. See
// gtk_tool_palette_set_exclusive().
//
// The function takes the following parameters:
//
//    - group which is a child of palette.
//
// The function returns the following values:
//
//    - ok: TRUE if group is exclusive.
//
func (palette *ToolPalette) Exclusive(group *ToolItemGroup) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_exclusive", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Expand gets whether group should be given extra space. See
// gtk_tool_palette_set_expand().
//
// The function takes the following parameters:
//
//    - group which is a child of palette.
//
// The function returns the following values:
//
//    - ok: TRUE if group should be given extra space, FALSE otherwise.
//
func (palette *ToolPalette) Expand(group *ToolItemGroup) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_expand", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GroupPosition gets the position of group in palette as index. See
// gtk_tool_palette_set_group_position().
//
// The function takes the following parameters:
//
//    - group: ToolItemGroup.
//
// The function returns the following values:
//
//    - gint: index of group or -1 if group is not a child of palette.
//
func (palette *ToolPalette) GroupPosition(group *ToolItemGroup) int32 {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_group_position", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// HAdjustment gets the horizontal adjustment of the tool palette.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
//
// The function returns the following values:
//
//    - adjustment: horizontal adjustment of palette.
//
func (palette *ToolPalette) HAdjustment() *Adjustment {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	*(**ToolPalette)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_hadjustment", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// VAdjustment gets the vertical adjustment of the tool palette.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
//
// The function returns the following values:
//
//    - adjustment: vertical adjustment of palette.
//
func (palette *ToolPalette) VAdjustment() *Adjustment {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	*(**ToolPalette)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("get_vadjustment", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(palette)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// SetExclusive sets whether the group should be exclusive or not. If an
// exclusive group is expanded all other groups are collapsed.
//
// The function takes the following parameters:
//
//    - group which is a child of palette.
//    - exclusive: whether the group should be exclusive or not.
//
func (palette *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	if exclusive {
		_arg2 = C.TRUE
	}
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1
	*(**ToolItemGroup)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("set_exclusive", args[:], nil)

	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)
	runtime.KeepAlive(exclusive)
}

// SetExpand sets whether the group should be given extra space.
//
// The function takes the following parameters:
//
//    - group which is a child of palette.
//    - expand: whether the group should be given extra space.
//
func (palette *ToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1
	*(**ToolItemGroup)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("set_expand", args[:], nil)

	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)
	runtime.KeepAlive(expand)
}

// SetGroupPosition sets the position of the group as an index of the tool
// palette. If position is 0 the group will become the first child, if position
// is -1 it will become the last child.
//
// The function takes the following parameters:
//
//    - group which is a child of palette.
//    - position: new index for group.
//
func (palette *ToolPalette) SetGroupPosition(group *ToolItemGroup, position int32) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	_arg2 = C.gint(position)
	*(**ToolPalette)(unsafe.Pointer(&args[1])) = _arg1
	*(**ToolItemGroup)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("set_group_position", args[:], nil)

	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)
	runtime.KeepAlive(position)
}

// UnsetIconSize unsets the tool palette icon size set with
// gtk_tool_palette_set_icon_size(), so that user preferences will be used to
// determine the icon size.
func (palette *ToolPalette) UnsetIconSize() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	*(**ToolPalette)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("unset_icon_size", args[:], nil)

	runtime.KeepAlive(palette)
}

// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(), so
// that user preferences will be used to determine the toolbar style.
func (palette *ToolPalette) UnsetStyle() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(palette).Native()))
	*(**ToolPalette)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "ToolPalette").InvokeMethod("unset_style", args[:], nil)

	runtime.KeepAlive(palette)
}

// ToolPaletteGetDragTargetGroup: get the target entry for a dragged
// ToolItemGroup.
//
// The function returns the following values:
//
//    - targetEntry for a dragged group.
//
func ToolPaletteGetDragTargetGroup() *TargetEntry {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "get_drag_target_group").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _targetEntry
}

// ToolPaletteGetDragTargetItem gets the target entry for a dragged ToolItem.
//
// The function returns the following values:
//
//    - targetEntry for a dragged item.
//
func ToolPaletteGetDragTargetItem() *TargetEntry {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "get_drag_target_item").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _targetEntry
}
