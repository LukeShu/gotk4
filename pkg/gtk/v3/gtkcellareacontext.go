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
// extern void _gotk4_gtk3_CellAreaContextClass_allocate(GtkCellAreaContext*, gint, gint);
// extern void _gotk4_gtk3_CellAreaContextClass_reset(GtkCellAreaContext*);
import "C"

// glib.Type values for gtkcellareacontext.go.
var GTypeCellAreaContext = coreglib.Type(C.gtk_cell_area_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCellAreaContext, F: marshalCellAreaContext},
	})
}

// CellAreaContextOverrider contains methods that are overridable.
type CellAreaContextOverrider interface {
	// Allocate allocates a width and/or a height for all rows which are to be
	// rendered with context.
	//
	// Usually allocation is performed only horizontally or sometimes vertically
	// since a group of rows are usually rendered side by side vertically or
	// horizontally and share either the same width or the same height.
	// Sometimes they are allocated in both horizontal and vertical orientations
	// producing a homogeneous effect of the rows. This is generally the case
	// for TreeView when TreeView:fixed-height-mode is enabled.
	//
	// Since 3.0.
	//
	// The function takes the following parameters:
	//
	//    - width: allocated width for all TreeModel rows rendered with context,
	//      or -1.
	//    - height: allocated height for all TreeModel rows rendered with
	//      context, or -1.
	//
	Allocate(width, height int32)
	// Reset resets any previously cached request and allocation data.
	//
	// When underlying TreeModel data changes its important to reset the context
	// if the content size is allowed to shrink. If the content size is only
	// allowed to grow (this is usually an option for views rendering large data
	// stores as a measure of optimization), then only the row that changed or
	// was inserted needs to be (re)requested with
	// gtk_cell_area_get_preferred_width().
	//
	// When the new overall size of the context requires that the allocated size
	// changes (or whenever this allocation changes at all), the variable row
	// sizes need to be re-requested for every row.
	//
	// For instance, if the rows are displayed all with the same width from top
	// to bottom then a change in the allocated width necessitates a
	// recalculation of all the displayed row heights using
	// gtk_cell_area_get_preferred_height_for_width().
	//
	// Since 3.0.
	Reset()
}

// CellAreaContext object is created by a given CellArea implementation via its
// CellAreaClass.create_context() virtual method and is used to store cell sizes
// and alignments for a series of TreeModel rows that are requested and rendered
// in the same context.
//
// CellLayout widgets can create any number of contexts in which to request and
// render groups of data rows. However, it’s important that the same context
// which was used to request sizes for a given TreeModel row also be used for
// the same row when calling other CellArea APIs such as gtk_cell_area_render()
// and gtk_cell_area_event().
type CellAreaContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CellAreaContext)(nil)
)

func classInitCellAreaContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCellAreaContextClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCellAreaContextClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Allocate(width, height int32) }); ok {
		pclass.allocate = (*[0]byte)(C._gotk4_gtk3_CellAreaContextClass_allocate)
	}

	if _, ok := goval.(interface{ Reset() }); ok {
		pclass.reset = (*[0]byte)(C._gotk4_gtk3_CellAreaContextClass_reset)
	}
}

//export _gotk4_gtk3_CellAreaContextClass_allocate
func _gotk4_gtk3_CellAreaContextClass_allocate(arg0 *C.GtkCellAreaContext, arg1 C.gint, arg2 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Allocate(width, height int32) })

	var _width int32  // out
	var _height int32 // out

	_width = int32(arg1)
	_height = int32(arg2)

	iface.Allocate(_width, _height)
}

//export _gotk4_gtk3_CellAreaContextClass_reset
func _gotk4_gtk3_CellAreaContextClass_reset(arg0 *C.GtkCellAreaContext) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Reset() })

	iface.Reset()
}

func wrapCellAreaContext(obj *coreglib.Object) *CellAreaContext {
	return &CellAreaContext{
		Object: obj,
	}
}

func marshalCellAreaContext(p uintptr) (interface{}, error) {
	return wrapCellAreaContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Allocate allocates a width and/or a height for all rows which are to be
// rendered with context.
//
// Usually allocation is performed only horizontally or sometimes vertically
// since a group of rows are usually rendered side by side vertically or
// horizontally and share either the same width or the same height. Sometimes
// they are allocated in both horizontal and vertical orientations producing a
// homogeneous effect of the rows. This is generally the case for TreeView when
// TreeView:fixed-height-mode is enabled.
//
// Since 3.0.
//
// The function takes the following parameters:
//
//    - width: allocated width for all TreeModel rows rendered with context, or
//      -1.
//    - height: allocated height for all TreeModel rows rendered with context, or
//      -1.
//
func (context *CellAreaContext) Allocate(width, height int32) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _arg2 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)
	*(**CellAreaContext)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("allocate", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// Area fetches the CellArea this context was created by.
//
// This is generally unneeded by layouting widgets; however, it is important for
// the context implementation itself to fetch information about the area it is
// being used for.
//
// For instance at CellAreaContextClass.allocate() time it’s important to know
// details about any cell spacing that the CellArea is configured with in order
// to compute a proper allocation.
//
// The function returns the following values:
//
//    - cellArea this context was created by.
//
func (context *CellAreaContext) Area() CellAreaer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**CellAreaContext)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("get_area", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _cellArea CellAreaer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.CellAreaer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellAreaer)
			return ok
		})
		rv, ok := casted.(CellAreaer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellAreaer")
		}
		_cellArea = rv
	}

	return _cellArea
}

// PushPreferredHeight causes the minimum and/or natural height to grow if the
// new proposed sizes exceed the current minimum and natural height.
//
// This is used by CellAreaContext implementations during the request process
// over a series of TreeModel rows to progressively push the requested height
// over a series of gtk_cell_area_get_preferred_height() requests.
//
// The function takes the following parameters:
//
//    - minimumHeight: proposed new minimum height for context.
//    - naturalHeight: proposed new natural height for context.
//
func (context *CellAreaContext) PushPreferredHeight(minimumHeight, naturalHeight int32) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _arg2 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.gint(minimumHeight)
	_arg2 = C.gint(naturalHeight)
	*(**CellAreaContext)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("push_preferred_height", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(minimumHeight)
	runtime.KeepAlive(naturalHeight)
}

// PushPreferredWidth causes the minimum and/or natural width to grow if the new
// proposed sizes exceed the current minimum and natural width.
//
// This is used by CellAreaContext implementations during the request process
// over a series of TreeModel rows to progressively push the requested width
// over a series of gtk_cell_area_get_preferred_width() requests.
//
// The function takes the following parameters:
//
//    - minimumWidth: proposed new minimum width for context.
//    - naturalWidth: proposed new natural width for context.
//
func (context *CellAreaContext) PushPreferredWidth(minimumWidth, naturalWidth int32) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _arg2 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg1 = C.gint(minimumWidth)
	_arg2 = C.gint(naturalWidth)
	*(**CellAreaContext)(unsafe.Pointer(&args[1])) = _arg1
	*(*int32)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("push_preferred_width", args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(minimumWidth)
	runtime.KeepAlive(naturalWidth)
}

// Reset resets any previously cached request and allocation data.
//
// When underlying TreeModel data changes its important to reset the context if
// the content size is allowed to shrink. If the content size is only allowed to
// grow (this is usually an option for views rendering large data stores as a
// measure of optimization), then only the row that changed or was inserted
// needs to be (re)requested with gtk_cell_area_get_preferred_width().
//
// When the new overall size of the context requires that the allocated size
// changes (or whenever this allocation changes at all), the variable row sizes
// need to be re-requested for every row.
//
// For instance, if the rows are displayed all with the same width from top to
// bottom then a change in the allocated width necessitates a recalculation of
// all the displayed row heights using
// gtk_cell_area_get_preferred_height_for_width().
//
// Since 3.0.
func (context *CellAreaContext) Reset() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(**CellAreaContext)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("reset", args[:], nil)

	runtime.KeepAlive(context)
}
